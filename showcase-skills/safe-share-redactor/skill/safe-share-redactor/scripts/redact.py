#!/usr/bin/env python3
"""Create a sanitized text copy and a raw-value-free findings report."""

from __future__ import annotations

import argparse
import hashlib
import json
import os
import re
import sys
import tempfile
from dataclasses import dataclass
from pathlib import Path


@dataclass(frozen=True)
class Candidate:
    start: int
    end: int
    rule_id: str
    replacement: str
    value: str
    priority: int


RULES = (
    (
        "private-key-material",
        re.compile(
            r"(?ms)-----BEGIN (?:RSA |EC |OPENSSH )?PRIVATE KEY-----\n"
            r"(.+?)(?=\n-----END (?:RSA |EC |OPENSSH )?PRIVATE KEY-----)"
        ),
        "<REDACTED:PRIVATE_KEY_MATERIAL>",
        100,
        1,
    ),
    (
        "authorization-value",
        re.compile(
            r"(?im)\bauthorization\s*[:=]\s*[\"']?(?:bearer|basic)\s+"
            r"([A-Za-z0-9._~+/=-]{8,})"
        ),
        "<REDACTED:AUTHORIZATION>",
        90,
        1,
    ),
    (
        "quoted-secret-value",
        re.compile(
            r"(?im)\b(?:api[_-]?key|access[_-]?token|auth[_-]?token|password|passwd|"
            r"secret|client[_-]?secret)\b\s*[:=]\s*([\"'])([^\r\n\"']+)\1"
        ),
        "<REDACTED:SECRET>",
        85,
        2,
    ),
    (
        "unquoted-secret-value",
        re.compile(
            r"(?im)\b(?:api[_-]?key|access[_-]?token|auth[_-]?token|password|passwd|"
            r"secret|client[_-]?secret)\b\s*[:=]\s*([^\s,;#]+)"
        ),
        "<REDACTED:SECRET>",
        80,
        1,
    ),
    (
        "url-userinfo",
        re.compile(r"(?i)\bhttps?://([^\s/@:]+:[^\s/@]+)@"),
        "<REDACTED:URL_CREDENTIALS>",
        75,
        1,
    ),
    (
        "known-token",
        re.compile(
            r"\b(?:sk-[A-Za-z0-9_-]{16,}|ghp_[A-Za-z0-9]{16,}|"
            r"xox[baprs]-[A-Za-z0-9-]{16,})\b"
        ),
        "<REDACTED:TOKEN>",
        70,
        0,
    ),
    (
        "jwt",
        re.compile(
            r"\beyJ[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\b"
        ),
        "<REDACTED:JWT>",
        70,
        0,
    ),
    (
        "email-address",
        re.compile(r"(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b"),
        "<REDACTED:EMAIL>",
        50,
        0,
    ),
    (
        "user-home-segment",
        re.compile(r"(?:(?<=/Users/)|(?<=/home/)|(?<=C:\\Users\\))[^/\\\s]+"),
        "<REDACTED:USER>",
        45,
        0,
    ),
)

IPV4_PATTERN = re.compile(r"(?<![\d.])(?:\d{1,3}\.){3}\d{1,3}(?![\d.])")
AMBIGUOUS_PATTERN = re.compile(
    r"(?i)\b(?:account|address|client|confidential|customer|employee|internal|name|"
    r"operator|patient|private|proprietary|session|tenant|user)\b"
)
SAFE_SUFFIXES = {
    ".cfg",
    ".conf",
    ".csv",
    ".env",
    ".ini",
    ".json",
    ".log",
    ".md",
    ".toml",
    ".txt",
    ".xml",
    ".yaml",
    ".yml",
}


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Write a sanitized copy and auditable findings without changing the input."
    )
    parser.add_argument("input_file", type=Path)
    parser.add_argument("--output-dir", required=True, type=Path)
    return parser.parse_args()


def line_column(text: str, offset: int) -> tuple[int, int]:
    line = text.count("\n", 0, offset) + 1
    previous_newline = text.rfind("\n", 0, offset)
    return line, offset - previous_newline


def candidate_from_match(
    match: re.Match[str],
    rule_id: str,
    replacement: str,
    priority: int,
    group: int,
) -> Candidate:
    start, end = match.span(group)
    return Candidate(start, end, rule_id, replacement, match.group(group), priority)


def collect_candidates(text: str) -> list[Candidate]:
    candidates = [
        candidate_from_match(match, rule_id, replacement, priority, group)
        for rule_id, pattern, replacement, priority, group in RULES
        for match in pattern.finditer(text)
    ]
    for match in IPV4_PATTERN.finditer(text):
        if all(int(part) <= 255 for part in match.group(0).split(".")):
            candidates.append(
                Candidate(
                    match.start(),
                    match.end(),
                    "ip-address",
                    "<REDACTED:IP_ADDRESS>",
                    match.group(0),
                    40,
                )
            )
    return candidates


def select_non_overlapping(candidates: list[Candidate]) -> list[Candidate]:
    selected: list[Candidate] = []
    for candidate in sorted(
        candidates, key=lambda item: (-item.priority, item.start, -(item.end - item.start))
    ):
        if any(candidate.start < item.end and item.start < candidate.end for item in selected):
            continue
        selected.append(candidate)
    return sorted(selected, key=lambda item: item.start)


def deterministic_findings(text: str, selected: list[Candidate]) -> list[dict[str, object]]:
    findings = []
    for candidate in selected:
        line, column = line_column(text, candidate.start)
        findings.append(
            {
                "kind": "deterministic",
                "rule_id": candidate.rule_id,
                "line": line,
                "column": column,
                "replacement": candidate.replacement,
            }
        )
    return findings


def ambiguous_findings(text: str, selected: list[Candidate]) -> list[dict[str, object]]:
    findings = []
    for line_number, line in enumerate(text.splitlines(), start=1):
        markers = sorted({match.group(0).lower() for match in AMBIGUOUS_PATTERN.finditer(line)})
        if not markers:
            continue
        findings.append(
            {
                "kind": "ambiguous-context",
                "rule_id": "context-marker",
                "line": line_number,
                "markers": markers,
                "action": "review-source-and-sanitized-line",
            }
        )
    return findings


def apply_redactions(text: str, selected: list[Candidate]) -> str:
    parts: list[str] = []
    cursor = 0
    for candidate in selected:
        parts.extend((text[cursor : candidate.start], candidate.replacement))
        cursor = candidate.end
    parts.append(text[cursor:])
    return "".join(parts)


def atomic_write(path: Path, content: str) -> None:
    with tempfile.NamedTemporaryFile(
        "w", encoding="utf-8", dir=path.parent, delete=False, newline=""
    ) as temporary:
        temporary.write(content)
        temporary_path = Path(temporary.name)
    os.replace(temporary_path, path)


def run(input_file: Path, output_dir: Path) -> dict[str, object]:
    if not input_file.is_file():
        raise ValueError("input must be an existing regular file")

    input_resolved = input_file.resolve()
    output_dir.mkdir(parents=True, exist_ok=True)
    suffix = input_file.suffix.lower()
    sanitized_path = output_dir / f"sanitized{suffix if suffix in SAFE_SUFFIXES else '.txt'}"
    findings_path = output_dir / "findings.json"

    if sanitized_path.resolve() == input_resolved or findings_path.resolve() == input_resolved:
        raise ValueError("output artifacts must be distinct from the input")
    collisions = [path for path in (sanitized_path, findings_path) if path.exists()]
    if collisions:
        raise ValueError("refusing to overwrite existing output artifacts")

    source_bytes = input_file.read_bytes()
    try:
        text = source_bytes.decode("utf-8")
    except UnicodeDecodeError as error:
        raise ValueError("input must be valid UTF-8 text") from error

    selected = select_non_overlapping(collect_candidates(text))
    deterministic = deterministic_findings(text, selected)
    ambiguous = ambiguous_findings(text, selected)
    sanitized = apply_redactions(text, selected)
    source_hash = hashlib.sha256(source_bytes).hexdigest()

    report = {
        "schema_version": 1,
        "sanitized_output": sanitized_path.name,
        "source_sha256_before": source_hash,
        "source_sha256_after": hashlib.sha256(input_file.read_bytes()).hexdigest(),
        "summary": {
            "deterministic_count": len(deterministic),
            "ambiguous_risk_count": len(ambiguous),
            "changed": sanitized != text,
            "review_required": bool(deterministic or ambiguous),
        },
        "findings": deterministic + ambiguous,
        "limitations": [
            "Contextual or unsupported sensitive material may remain unchanged.",
            "Automated redaction does not guarantee complete privacy or secrecy.",
        ],
    }

    atomic_write(sanitized_path, sanitized)
    atomic_write(findings_path, json.dumps(report, indent=2, sort_keys=True) + "\n")
    return report


def main() -> int:
    args = parse_args()
    try:
        report = run(args.input_file, args.output_dir)
    except (OSError, ValueError) as error:
        print(f"redaction failed: {error}", file=sys.stderr)
        return 2
    print(json.dumps(report["summary"], sort_keys=True))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

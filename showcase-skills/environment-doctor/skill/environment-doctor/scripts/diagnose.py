#!/usr/bin/env python3
"""Produce deterministic, read-only development-environment diagnostics."""

from __future__ import annotations

import argparse
import json
import os
import re
import shutil
import subprocess
import sys
from pathlib import Path
from typing import Any


VERSION_PROBES = {
    "bash": ["--version"],
    "cargo": ["--version"],
    "dart": ["--version"],
    "flutter": ["--version", "--machine"],
    "git": ["--version"],
    "go": ["version"],
    "java": ["-version"],
    "node": ["--version"],
    "npm": ["--version"],
    "python3": ["--version"],
    "ruby": ["--version"],
    "rustc": ["--version"],
    "zsh": ["--version"],
}
NAME_PATTERN = re.compile(r"^[A-Za-z_][A-Za-z0-9_]*$")
TOOL_PATTERN = re.compile(r"^[A-Za-z0-9][A-Za-z0-9._+-]*$")
VERSION_PATTERN = re.compile(r"(?<!\d)(\d+(?:\.\d+){0,3})(?!\d)")
WINDOWS_HOME_PATTERN = re.compile(r"(?i)[A-Z]:\\Users\\[^\\\s]+")
POSIX_HOME_PATTERN = re.compile(r"/(?:Users|home)/[^/\s]+")
OUTPUT_FILES = ("evidence.json", "report.txt")


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Inspect selected POSIX development-environment state without changing it."
    )
    parser.add_argument("--root", required=True, type=Path)
    parser.add_argument("--output-dir", required=True, type=Path)
    parser.add_argument("--tool", action="append", default=[])
    parser.add_argument("--env", action="append", default=[])
    parser.add_argument(
        "--expect-path-before", action="append", default=[], nargs=2, metavar=("EARLIER", "LATER")
    )
    parser.add_argument(
        "--version-file", action="append", default=[], nargs=2, metavar=("TOOL", "RELATIVE_FILE")
    )
    return parser.parse_args()


def fail(message: str) -> None:
    raise SystemExit(f"environment-doctor: {message}")


def canonical(path: Path) -> Path:
    return Path(os.path.realpath(os.path.abspath(path)))


def is_within(path: Path, parent: Path) -> bool:
    try:
        path.relative_to(parent)
        return True
    except ValueError:
        return False


def validate_args(args: argparse.Namespace) -> tuple[Path, Path]:
    if os.name != "posix":
        fail("unsupported platform: this script requires POSIX PATH semantics")
    root = canonical(args.root)
    output_dir = canonical(args.output_dir)
    if not root.is_dir():
        fail("--root must name an existing directory")
    if output_dir.exists():
        fail("--output-dir must not already exist")
    if not output_dir.parent.is_dir():
        fail("--output-dir parent must be an existing directory")
    if is_within(output_dir, root):
        fail("--output-dir must be outside --root")
    if not (args.tool or args.env or args.expect_path_before or args.version_file):
        fail("at least one inspection selector is required")
    for tool in args.tool + [item[0] for item in args.version_file]:
        if not TOOL_PATTERN.fullmatch(tool):
            fail(f"invalid tool name: {tool!r}")
    for name in args.env:
        if not NAME_PATTERN.fullmatch(name):
            fail(f"invalid environment-variable name: {name!r}")
    for _, relative_name in args.version_file:
        relative_path = Path(relative_name)
        if relative_path.is_absolute() or ".." in relative_path.parts:
            fail(f"version file must be a relative path within --root: {relative_name!r}")
    return root, output_dir


def normalize_path(path: Path, root: Path, home: Path) -> str:
    absolute = Path(os.path.abspath(path))
    resolved = canonical(absolute)
    if is_within(resolved, root):
        relative = resolved.relative_to(root)
        return "<root>" if str(relative) == "." else f"<root>/{relative.as_posix()}"
    if is_within(resolved, home):
        relative = resolved.relative_to(home)
        return "~" if str(relative) == "." else f"~/{relative.as_posix()}"
    return absolute.as_posix()


def sanitize_text(text: str, root: Path, home: Path) -> str:
    sanitized = text.replace(str(root), "<root>").replace(str(home), "~")
    sanitized = WINDOWS_HOME_PATTERN.sub("~", sanitized)
    sanitized = POSIX_HOME_PATTERN.sub("~", sanitized)
    sanitized = " ".join(sanitized.replace("\x00", "").split())
    return sanitized[:4096]


def path_entries(root: Path, home: Path) -> list[dict[str, Any]]:
    entries = []
    for index, raw_entry in enumerate(os.environ.get("PATH", "").split(os.pathsep)):
        effective = Path(raw_entry or os.curdir)
        entries.append(
            {
                "index": index,
                "path": normalize_path(effective, root, home),
                "exists": effective.is_dir(),
            }
        )
    return entries


def executable_candidates(tool: str, entries: list[dict[str, Any]], root: Path, home: Path) -> list[dict[str, Any]]:
    candidates = []
    raw_entries = os.environ.get("PATH", "").split(os.pathsep)
    for index, raw_entry in enumerate(raw_entries):
        candidate = Path(raw_entry or os.curdir) / tool
        if candidate.is_file() and os.access(candidate, os.X_OK):
            candidates.append(
                {
                    "path_index": index,
                    "path": normalize_path(candidate, root, home),
                    "resolved_path": normalize_path(canonical(candidate), root, home),
                }
            )
    return candidates


def probe_version(tool: str, executable: str, root: Path, home: Path) -> dict[str, Any]:
    arguments = VERSION_PROBES.get(tool)
    if arguments is None:
        return {"state": "unsupported", "version": None, "output": None}
    try:
        completed = subprocess.run(
            [executable, *arguments],
            cwd=root,
            env=os.environ.copy(),
            stdin=subprocess.DEVNULL,
            stdout=subprocess.PIPE,
            stderr=subprocess.STDOUT,
            text=True,
            timeout=5,
            check=False,
        )
    except subprocess.TimeoutExpired:
        return {"state": "timeout", "version": None, "output": None}
    except OSError as error:
        return {
            "state": "unavailable",
            "version": None,
            "output": sanitize_text(str(error), root, home),
        }
    output = sanitize_text(completed.stdout, root, home)
    match = VERSION_PATTERN.search(output)
    state = "available" if completed.returncode == 0 else "failed"
    return {
        "state": state,
        "version": match.group(1) if match else None,
        "output": output or None,
        "exit_code": completed.returncode,
        "arguments": arguments,
    }


def finding(
    identifier: str,
    severity: str,
    summary: str,
    evidence: dict[str, Any],
    remediation: str,
    verification: str,
) -> dict[str, Any]:
    return {
        "id": identifier,
        "severity": severity,
        "summary": summary,
        "evidence": evidence,
        "remediation": remediation,
        "verification": verification,
    }


def inspect_tools(tools: list[str], entries: list[dict[str, Any]], root: Path, home: Path) -> tuple[list[dict[str, Any]], list[dict[str, Any]]]:
    results = []
    findings = []
    for tool in sorted(set(tools)):
        candidates = executable_candidates(tool, entries, root, home)
        executable = shutil.which(tool)
        version = probe_version(tool, executable, root, home) if executable else {"state": "unavailable", "version": None, "output": None}
        result = {"name": tool, "state": "available" if candidates else "unavailable", "selected": candidates[0] if candidates else None, "candidates": candidates, "version": version}
        results.append(result)
        if not candidates:
            findings.append(finding(f"tool.{tool}", "error", f"{tool} is unavailable on PATH", {"state": "unavailable"}, "Install or expose the tool through the platform's user-approved toolchain workflow.", f"command -v {tool} && {tool} --version"))
        elif version["state"] == "unsupported":
            findings.append(finding(f"tool.{tool}", "info", f"{tool} resolves; version probing is unsupported", {"state": "available", "selected": candidates[0], "version_state": "unsupported"}, "Consult the tool's authoritative documentation before selecting a non-mutating version command.", f"command -v {tool}"))
        elif version["state"] != "available":
            findings.append(finding(f"tool.{tool}", "warning", f"{tool} resolves but its version probe did not succeed", {"state": "available", "version_state": version["state"], "exit_code": version.get("exit_code")}, "Inspect the recorded probe state and the tool's installation without changing it automatically.", f"command -v {tool}"))
        else:
            findings.append(finding(f"tool.{tool}", "info", f"{tool} resolves with a readable version", {"state": "available", "selected": candidates[0], "version": version["version"]}, "No remediation is indicated by this check.", f"command -v {tool}"))
    return results, findings


def inspect_environment(names: list[str], entries: list[dict[str, Any]]) -> tuple[list[dict[str, Any]], list[dict[str, Any]]]:
    results = []
    findings = []
    for name in sorted(set(names)):
        value = os.environ.get(name)
        state = "unset" if value is None else "empty" if value == "" else "set"
        result: dict[str, Any] = {"name": name, "state": state, "value": None}
        if name == "PATH":
            result["entries"] = entries
        results.append(result)
        severity = "info" if state == "set" else "warning"
        summary = f"{name} is {state}; its value is omitted"
        findings.append(finding(f"env.{name}", severity, summary, {"state": state, "value_omitted": True}, "Set or correct this variable only in its user-owned configuration if the project requires it.", f"Re-run environment-doctor with --env {name}"))
    return results, findings


def resolve_expectation_path(value: str, root: Path) -> Path:
    path = Path(value)
    return canonical(path if path.is_absolute() else root / path)


def inspect_path_order(expectations: list[list[str]], root: Path, home: Path) -> tuple[list[dict[str, Any]], list[dict[str, Any]]]:
    raw_entries = [canonical(Path(item or os.curdir)) for item in os.environ.get("PATH", "").split(os.pathsep)]
    results = []
    findings = []
    for earlier_name, later_name in sorted(expectations):
        earlier = resolve_expectation_path(earlier_name, root)
        later = resolve_expectation_path(later_name, root)
        earlier_indices = [index for index, entry in enumerate(raw_entries) if entry == earlier]
        later_indices = [index for index, entry in enumerate(raw_entries) if entry == later]
        if not earlier_indices or not later_indices:
            state, severity = "missing", "warning"
        elif len(earlier_indices) > 1 or len(later_indices) > 1:
            state, severity = "duplicate", "warning"
        elif earlier_indices[0] < later_indices[0]:
            state, severity = "satisfied", "info"
        else:
            state, severity = "reversed", "warning"
        result = {"earlier": normalize_path(earlier, root, home), "later": normalize_path(later, root, home), "earlier_indices": earlier_indices, "later_indices": later_indices, "state": state}
        results.append(result)
        identifier = f"path-order.{len(results)}"
        findings.append(finding(identifier, severity, f"PATH order expectation is {state}", result, "Change PATH only in its owning user or project configuration after confirming the intended precedence.", "Re-run the same --expect-path-before check"))
    return results, findings


def declared_version(path: Path) -> tuple[str, str | None]:
    if not path.exists():
        return "missing", None
    if not path.is_file():
        return "unparseable", None
    try:
        content = path.read_text(encoding="utf-8")[:4096]
    except (OSError, UnicodeError):
        return "unparseable", None
    first_line = next((line.strip() for line in content.splitlines() if line.strip()), "")
    match = VERSION_PATTERN.search(first_line[:256])
    return ("available", match.group(1)) if match else ("unparseable", None)


def version_matches(actual: str, declared: str) -> bool:
    actual_parts = actual.split(".")
    declared_parts = declared.split(".")
    return actual_parts[: len(declared_parts)] == declared_parts


def inspect_version_files(selections: list[list[str]], tool_results: list[dict[str, Any]], root: Path, home: Path) -> tuple[list[dict[str, Any]], list[dict[str, Any]]]:
    tools = {item["name"]: item for item in tool_results}
    results = []
    findings = []
    for tool, relative_name in sorted(selections):
        path = canonical(root / relative_name)
        if not is_within(path, root):
            fail(f"version file resolves outside --root: {relative_name!r}")
        declaration_state, declared = declared_version(path)
        tool_result = tools.get(tool)
        actual = tool_result["version"].get("version") if tool_result else None
        if declaration_state != "available":
            state, severity = declaration_state, "warning"
        elif not actual:
            state, severity = "unavailable-tool", "warning"
        elif version_matches(actual, declared or ""):
            state, severity = "match", "info"
        else:
            state, severity = "mismatch", "warning"
        result = {"tool": tool, "file": normalize_path(path, root, home), "state": state, "declared_version": declared, "actual_version": actual}
        results.append(result)
        findings.append(finding(f"version-file.{tool}.{len(results)}", severity, f"{tool} version declaration is {state}", result, "Choose the project-approved toolchain or update the declaration only after confirming the intended version owner.", f"Re-run environment-doctor with --version-file {tool} {relative_name}"))
    return results, findings


def render_report(findings: list[dict[str, Any]]) -> str:
    lines = ["Environment Doctor", "==================", "", f"Findings: {len(findings)}", ""]
    for item in findings:
        lines.extend(
            [
                f"[{item['severity'].upper()}] {item['id']}: {item['summary']}",
                f"Evidence: {json.dumps(item['evidence'], sort_keys=True, separators=(',', ':'))}",
                f"Remediation: {item['remediation']}",
                f"Verification: {item['verification']}",
                "",
            ]
        )
    lines.append("No changes were made. Obtain approval before applying any remediation.")
    return "\n".join(lines) + "\n"


def main() -> int:
    args = parse_args()
    root, output_dir = validate_args(args)
    home = canonical(Path.home())
    entries = path_entries(root, home)
    all_tools = args.tool + [item[0] for item in args.version_file]
    tools, tool_findings = inspect_tools(all_tools, entries, root, home)
    environment, environment_findings = inspect_environment(args.env, entries)
    path_order, path_findings = inspect_path_order(args.expect_path_before, root, home)
    version_files, version_findings = inspect_version_files(args.version_file, tools, root, home)
    findings = tool_findings + environment_findings + path_findings + version_findings
    evidence = {
        "schema_version": 1,
        "inspection": {"root": "<root>", "platform": "posix", "changed_environment": False},
        "tools": tools,
        "environment": environment,
        "path_order": path_order,
        "version_files": version_files,
        "findings": findings,
    }
    output_dir.mkdir()
    (output_dir / "evidence.json").write_text(json.dumps(evidence, indent=2, sort_keys=True) + "\n", encoding="utf-8")
    (output_dir / "report.txt").write_text(render_report(findings), encoding="utf-8")
    return 1 if any(item["severity"] in {"warning", "error"} for item in findings) else 0


if __name__ == "__main__":
    sys.exit(main())

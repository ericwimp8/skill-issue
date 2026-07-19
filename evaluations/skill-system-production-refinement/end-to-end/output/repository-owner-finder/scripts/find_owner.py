#!/usr/bin/env python3

import argparse
import fnmatch
import re
import sys
from dataclasses import dataclass
from pathlib import Path


class OwnerLookupError(Exception):
    pass


@dataclass(frozen=True)
class OwnershipRule:
    pattern: str
    owners: tuple[str, ...]
    line_number: int

    @property
    def specificity(self) -> int:
        return len(re.sub(r"[*?]", "", self.pattern).strip("/"))


def project_root() -> Path:
    return Path(__file__).resolve().parents[3] / "project"


def normalize_repository_path(raw_path: str) -> str:
    path = raw_path.strip().replace("\\", "/").lstrip("/")
    parts = [part for part in path.split("/") if part not in {"", "."}]
    if not parts or ".." in parts:
        raise OwnerLookupError("Supply a non-empty repository path without '..'.")
    return "/".join(parts)


def read_rules(codeowners_path: Path) -> list[OwnershipRule]:
    try:
        lines = codeowners_path.read_text(encoding="utf-8").splitlines()
    except OSError as error:
        raise OwnerLookupError(
            f"Cannot read authoritative ownership source: {codeowners_path} ({error})"
        ) from error

    rules = [parse_rule(line, number) for number, line in enumerate(lines, start=1)]
    active_rules = [rule for rule in rules if rule is not None]
    if not active_rules:
        raise OwnerLookupError(f"No ownership rules found in {codeowners_path}.")
    return active_rules


def parse_rule(line: str, line_number: int) -> OwnershipRule | None:
    content = line.split("#", 1)[0].strip()
    if not content:
        return None
    fields = content.split()
    if len(fields) < 2:
        raise OwnerLookupError(f"Malformed CODEOWNERS entry on line {line_number}.")
    return OwnershipRule(fields[0], tuple(fields[1:]), line_number)


def pattern_matches(pattern: str, repository_path: str) -> bool:
    anchored = pattern.startswith("/")
    normalized_pattern = pattern.lstrip("/")
    directory_pattern = normalized_pattern.endswith("/")
    normalized_pattern = normalized_pattern.rstrip("/")

    if directory_pattern and (anchored or "/" in normalized_pattern):
        return repository_path.startswith(f"{normalized_pattern}/")
    if directory_pattern:
        return any(
            fnmatch.fnmatchcase(part, normalized_pattern)
            for part in repository_path.split("/")[:-1]
        )
    if "/" in normalized_pattern:
        return fnmatch.fnmatchcase(repository_path, normalized_pattern)
    return any(
        fnmatch.fnmatchcase(part, normalized_pattern)
        for part in repository_path.split("/")
    )


def find_owner(rules: list[OwnershipRule], repository_path: str) -> OwnershipRule:
    matches = [rule for rule in rules if pattern_matches(rule.pattern, repository_path)]
    if not matches:
        raise OwnerLookupError(
            f"No authoritative owner matches repository path '{repository_path}'."
        )
    return max(matches, key=lambda rule: (rule.specificity, rule.line_number))


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Report a repository path owner from the live local CODEOWNERS file."
    )
    parser.add_argument("repository_path", help="Repository-relative path to resolve")
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    codeowners_path = project_root() / "CODEOWNERS"
    try:
        repository_path = normalize_repository_path(args.repository_path)
        rule = find_owner(read_rules(codeowners_path), repository_path)
    except OwnerLookupError as error:
        print(f"Error: {error}", file=sys.stderr)
        return 1

    print(f"Owner: {' '.join(rule.owners)}")
    print(f"Pattern: {rule.pattern}")
    print(f"Source: {codeowners_path}:{rule.line_number}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

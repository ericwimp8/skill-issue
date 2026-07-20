#!/usr/bin/env python3
"""Normalize and deterministically order incident timeline JSONL records."""

from __future__ import annotations

import argparse
import hashlib
import json
import os
import sys
import tempfile
from datetime import datetime, timezone
from pathlib import Path
from zoneinfo import ZoneInfo, ZoneInfoNotFoundError


CLASSIFICATIONS = {
    "observed",
    "reported",
    "inferred",
    "contradiction",
    "gap",
}


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Normalize explicit timestamps and stably order incident JSONL records."
    )
    parser.add_argument(
        "input",
        type=Path,
        help=(
            "UTF-8 JSONL with source_id, source_path, locator, classification, "
            "summary, and optional raw_timestamp and source_timezone fields"
        ),
    )
    parser.add_argument(
        "--output",
        type=Path,
        help="write JSON to a new file instead of stdout; existing files are refused",
    )
    return parser.parse_args()


def local_time_status(value: datetime, zone: ZoneInfo) -> tuple[str, datetime | None]:
    candidates: list[datetime] = []
    for fold in (0, 1):
        localized = value.replace(tzinfo=zone, fold=fold)
        round_trip = localized.astimezone(timezone.utc).astimezone(zone)
        if round_trip.replace(tzinfo=None) == value and round_trip.fold == fold:
            candidates.append(localized)

    if not candidates:
        return "invalid", None
    if len(candidates) == 2 and candidates[0].utcoffset() != candidates[1].utcoffset():
        return "ambiguous", None
    return "resolved", candidates[0]


def normalize_timestamp(
    raw_timestamp: object, source_timezone: object
) -> tuple[str, str | None, str | None]:
    if raw_timestamp is None or raw_timestamp == "":
        return "missing", None, None
    if not isinstance(raw_timestamp, str):
        return "invalid", None, "raw_timestamp must be a string or null"

    candidate = raw_timestamp[:-1] + "+00:00" if raw_timestamp.endswith("Z") else raw_timestamp
    try:
        parsed = datetime.fromisoformat(candidate)
    except ValueError:
        return "invalid", None, "raw_timestamp is not an ISO 8601 date-time"

    if parsed.tzinfo is not None:
        normalized = parsed.astimezone(timezone.utc)
        return "resolved", normalized.isoformat().replace("+00:00", "Z"), None

    if not isinstance(source_timezone, str) or not source_timezone:
        return "missing-timezone", None, "naive timestamp requires source_timezone"
    try:
        zone = ZoneInfo(source_timezone)
    except ZoneInfoNotFoundError:
        return "invalid-timezone", None, "source_timezone is not available"

    status, localized = local_time_status(parsed, zone)
    if localized is None:
        message = "local time is ambiguous in source_timezone" if status == "ambiguous" else "local time does not exist in source_timezone"
        return status, None, message
    normalized = localized.astimezone(timezone.utc)
    return "resolved", normalized.isoformat().replace("+00:00", "Z"), None


def validate_record(record: object, line_number: int) -> dict[str, object]:
    if not isinstance(record, dict):
        raise ValueError(f"line {line_number}: record must be a JSON object")
    for field in ("source_id", "source_path", "locator", "classification", "summary"):
        if not isinstance(record.get(field), str) or not record[field]:
            raise ValueError(f"line {line_number}: {field} must be a non-empty string")
    if record["classification"] not in CLASSIFICATIONS:
        allowed = ", ".join(sorted(CLASSIFICATIONS))
        raise ValueError(f"line {line_number}: classification must be one of {allowed}")
    return record


def build_payload(input_path: Path) -> dict[str, object]:
    source_bytes = input_path.read_bytes()
    try:
        source_text = source_bytes.decode("utf-8")
    except UnicodeDecodeError as error:
        raise ValueError("input must be valid UTF-8") from error

    events: list[dict[str, object]] = []
    for line_number, line in enumerate(source_text.splitlines(), start=1):
        if not line.strip():
            continue
        try:
            record = validate_record(json.loads(line), line_number)
        except json.JSONDecodeError as error:
            raise ValueError(f"line {line_number}: invalid JSON") from error
        status, normalized, issue = normalize_timestamp(
            record.get("raw_timestamp"), record.get("source_timezone")
        )
        event = dict(record)
        event["input_order"] = len(events) + 1
        event["timestamp_status"] = status
        event["normalized_utc"] = normalized
        if issue is not None:
            event["timestamp_issue"] = issue
        events.append(event)

    resolved = [event for event in events if event["timestamp_status"] == "resolved"]
    unresolved = [event for event in events if event["timestamp_status"] != "resolved"]
    resolved.sort(key=lambda event: (str(event["normalized_utc"]), int(event["input_order"])))
    return {
        "schema_version": 1,
        "input_sha256": hashlib.sha256(source_bytes).hexdigest(),
        "resolved_events": resolved,
        "unresolved_events": unresolved,
    }


def serialize(payload: dict[str, object]) -> bytes:
    return (json.dumps(payload, indent=2, sort_keys=True, ensure_ascii=False) + "\n").encode("utf-8")


def write_new_file(output_path: Path, content: bytes) -> None:
    if output_path.exists():
        raise ValueError("output already exists")
    output_path.parent.mkdir(parents=True, exist_ok=True)
    descriptor, temporary_name = tempfile.mkstemp(prefix=f".{output_path.name}.", dir=output_path.parent)
    try:
        with os.fdopen(descriptor, "wb") as temporary_file:
            temporary_file.write(content)
            temporary_file.flush()
            os.fsync(temporary_file.fileno())
        os.link(temporary_name, output_path)
    finally:
        Path(temporary_name).unlink(missing_ok=True)


def main() -> int:
    args = parse_args()
    try:
        if args.output is not None and args.input.resolve() == args.output.resolve():
            raise ValueError("output must differ from input")
        content = serialize(build_payload(args.input))
        if args.output is None:
            sys.stdout.buffer.write(content)
        else:
            write_new_file(args.output, content)
    except (OSError, ValueError) as error:
        print(f"error: {error}", file=sys.stderr)
        return 2
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

#!/usr/bin/env python3
"""Direct validation for the incident timeline helper."""

from __future__ import annotations

import hashlib
import json
import subprocess
import sys
import tempfile
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / "skill/incident-timeline-builder/scripts/build_timeline.py"
FIXTURE = ROOT / "fixtures/script/records.jsonl"


def run(*arguments: str) -> subprocess.CompletedProcess[bytes]:
    return subprocess.run(
        [sys.executable, str(SCRIPT), *arguments],
        check=False,
        capture_output=True,
    )


def require(condition: bool, message: str) -> None:
    if not condition:
        raise AssertionError(message)


def main() -> int:
    source_before = FIXTURE.read_bytes()
    first = run(str(FIXTURE))
    second = run(str(FIXTURE))
    require(first.returncode == 0, first.stderr.decode())
    require(first.stdout == second.stdout, "repeated stdout differs")
    require(FIXTURE.read_bytes() == source_before, "input fixture changed")

    payload = json.loads(first.stdout)
    require(payload["input_sha256"] == hashlib.sha256(source_before).hexdigest(), "input hash differs")
    resolved = payload["resolved_events"]
    unresolved = payload["unresolved_events"]
    require([event["source_id"] for event in resolved] == ["note-1", "deploy-1", "log-1", "alert-1"], "chronological or stable tie ordering differs")
    require(resolved[0]["normalized_utc"] == "2026-04-04T14:55:00Z", "named zone normalization differs")
    require(resolved[-1]["normalized_utc"] == "2026-04-04T16:10:00Z", "offset normalization differs")
    statuses = {event["source_id"]: event["timestamp_status"] for event in unresolved}
    require(statuses == {"note-2": "missing", "note-3": "ambiguous"}, "unresolved statuses differ")
    require(all(event["source_path"].startswith("fixtures/") for event in resolved + unresolved), "provenance was not preserved")
    require(str(ROOT.resolve()).encode() not in first.stdout, "output leaked validator path")

    with tempfile.TemporaryDirectory() as temporary_directory:
        temporary = Path(temporary_directory)
        invalid = temporary / "invalid.jsonl"
        invalid.write_text("not-json\n", encoding="utf-8")
        invalid_result = run(str(invalid))
        require(invalid_result.returncode == 2 and b"invalid JSON" in invalid_result.stderr, "invalid JSON was not rejected")

        naive = temporary / "naive.jsonl"
        naive.write_text(
            json.dumps(
                {
                    "source_id": "n",
                    "source_path": "fixture.log",
                    "locator": "line 1",
                    "classification": "observed",
                    "summary": "Naive time",
                    "raw_timestamp": "2026-01-01T10:00:00",
                }
            )
            + "\n",
            encoding="utf-8",
        )
        naive_payload = json.loads(run(str(naive)).stdout)
        require(naive_payload["unresolved_events"][0]["timestamp_status"] == "missing-timezone", "naive time acquired a zone")

        output = temporary / "timeline.json"
        written = run(str(FIXTURE), "--output", str(output))
        require(written.returncode == 0 and output.read_bytes() == first.stdout, "file output differs from stdout")
        collision = run(str(FIXTURE), "--output", str(output))
        require(collision.returncode == 2 and b"output already exists" in collision.stderr, "output collision was not refused")
        require(output.read_bytes() == first.stdout, "collision mutated output")
        same_path = run(str(naive), "--output", str(naive))
        require(same_path.returncode == 2 and naive.exists(), "input overwrite was not refused")

    source_code = SCRIPT.read_text(encoding="utf-8")
    require("urllib" not in source_code and "requests" not in source_code and "socket" not in source_code, "network-capable import found")
    print("PASS: determinism, zones, ambiguity, missing time, provenance, errors, preservation, privacy")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

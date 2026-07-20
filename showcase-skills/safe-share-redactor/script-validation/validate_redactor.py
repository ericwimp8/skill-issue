#!/usr/bin/env python3
"""Focused behavioral validation for the bundled redactor."""

from __future__ import annotations

import hashlib
import json
import subprocess
import sys
import tempfile
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
SCRIPT = ROOT / "skill" / "safe-share-redactor" / "scripts" / "redact.py"


def run_cli(source: Path, output: Path) -> subprocess.CompletedProcess[str]:
    return subprocess.run(
        [sys.executable, str(SCRIPT), str(source), "--output-dir", str(output)],
        check=False,
        capture_output=True,
        text=True,
    )


def main() -> int:
    with tempfile.TemporaryDirectory() as temporary:
        workspace = Path(temporary)
        source = workspace / "source-sensitive-name.log"
        output = workspace / "output"
        source_text = """password=synthetic-passphrase
Authorization: Basic c3ludGhldGljLWNsaWVudA==
endpoint=https://ericwimp8:synthetic-pass@example.invalid/path
token=ghp_SYNTHETIC1234567890
jwt=eyJzeW50aGV0aWMiOiJ0cnVlIn0.c3ludGhldGljLXBheWxvYWQ.c3ludGhldGljLXNpZ25hdHVyZQ
customer=Eric Wimp contact=ericwimp8@example.com
peer=198.51.100.27
unix=/home/ericwimp8/cache
windows=C:\\Users\\ericwimp8\\cache
-----BEGIN PRIVATE KEY-----
SYNTHETICPRIVATEKEYMATERIAL
-----END PRIVATE KEY-----
Customer name: Eric Wimp
"""
        source.write_text(source_text, encoding="utf-8", newline="")
        source_hash = hashlib.sha256(source.read_bytes()).hexdigest()

        first = run_cli(source, output)
        assert first.returncode == 0, first.stderr
        sanitized = (output / "sanitized.log").read_text(encoding="utf-8")
        findings_text = (output / "findings.json").read_text(encoding="utf-8")
        findings = json.loads(findings_text)

        expected_placeholders = {
            "<REDACTED:AUTHORIZATION>",
            "<REDACTED:EMAIL>",
            "<REDACTED:IP_ADDRESS>",
            "<REDACTED:PRIVATE_KEY_MATERIAL>",
            "<REDACTED:SECRET>",
            "<REDACTED:TOKEN>",
            "<REDACTED:URL_CREDENTIALS>",
            "<REDACTED:USER>",
        }
        assert expected_placeholders.issubset(set(findings_text.split('"')))
        assert "Eric Wimp" in sanitized
        assert findings["summary"]["ambiguous_risk_count"] >= 1
        assert findings["source_sha256_before"] == source_hash
        assert findings["source_sha256_after"] == source_hash
        assert hashlib.sha256(source.read_bytes()).hexdigest() == source_hash

        forbidden_values = (
            "synthetic-passphrase",
            "c3ludGhldGljLWNsaWVudA==",
            "SYNTHETICPRIVATEKEYMATERIAL",
            "198.51.100.27",
            "ericwimp8@example.com",
        )
        assert all(value not in findings_text for value in forbidden_values)
        assert "source-sensitive-name" not in findings_text
        assert str(workspace) not in findings_text
        assert "fingerprint" not in findings_text

        sanitized_hash = hashlib.sha256((output / "sanitized.log").read_bytes()).hexdigest()
        findings_hash = hashlib.sha256((output / "findings.json").read_bytes()).hexdigest()
        collision = run_cli(source, output)
        assert collision.returncode == 2
        assert "refusing to overwrite" in collision.stderr
        assert hashlib.sha256((output / "sanitized.log").read_bytes()).hexdigest() == sanitized_hash
        assert hashlib.sha256((output / "findings.json").read_bytes()).hexdigest() == findings_hash

        invalid = workspace / "invalid.bin"
        invalid.write_bytes(b"\xff\xfe")
        invalid_result = run_cli(invalid, workspace / "invalid-output")
        assert invalid_result.returncode == 2
        assert "valid UTF-8" in invalid_result.stderr

    print("PASS: deterministic rules, ambiguity, source preservation, collisions, encoding")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())

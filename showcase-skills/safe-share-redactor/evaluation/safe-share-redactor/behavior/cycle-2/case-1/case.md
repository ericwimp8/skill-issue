# Case 1 Execution Record

- **Executor:** fresh Codex sub-agent `<fresh-agent>`, started without parent-turn inheritance.
- **Fixture:** `showcase-skills/safe-share-redactor/fixtures/behavior/case-1/input.log`.
- **Command:** `python3 showcase-skills/safe-share-redactor/skill/safe-share-redactor/scripts/redact.py showcase-skills/safe-share-redactor/fixtures/behavior/case-1/input.log --output-dir showcase-skills/safe-share-redactor/evaluation/safe-share-redactor/behavior/cycle-2/case-1`.
- **Direct execution result:** exit 0; 5 deterministic findings, 0 ambiguous risks, changed true, review required.
- **Outputs:** `sanitized.log` and `findings.json` in this directory.
- **Source preservation:** before and after SHA-256 values are identical in `findings.json`.
- **Ground-truth comparison:** pass; the supported synthetic API key, authorization value, email, documentation-range IP address, and user-home segment use stable placeholders while unrelated log structure remains.
- **Criteria:** 1, 2, 3, 4, and 7 pass; findings contain no raw matched values or source identity.
- **Cleanup completion:** the two generated outputs were removed after the cycle-2 audit; this record and `../audit.md` retain the execution result and material finding.

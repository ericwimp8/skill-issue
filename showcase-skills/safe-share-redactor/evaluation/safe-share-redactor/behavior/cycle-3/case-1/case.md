# Case 1 Execution Record

- **Executor:** fresh Codex sub-agent `<fresh-agent>`, started without parent-turn inheritance.
- **Fixture:** `showcase-skills/safe-share-redactor/fixtures/behavior/case-1/input.log`.
- **Command:** `python3 showcase-skills/safe-share-redactor/skill/safe-share-redactor/scripts/redact.py showcase-skills/safe-share-redactor/fixtures/behavior/case-1/input.log --output-dir showcase-skills/safe-share-redactor/evaluation/safe-share-redactor/behavior/cycle-3/case-1`.
- **Direct execution result:** exit 0; 5 deterministic findings, 0 ambiguous risks, changed true, review required.
- **Outputs:** retained privacy-normalized `sanitized.log` evidence and `findings.json` in this directory.
- **Source preservation:** before and after SHA-256 values are identical.
- **Ground-truth comparison:** pass; all five supported synthetic values use stable placeholders, unrelated log structure remains, and findings contain no raw or value-derived match material.
- **Criteria:** 1, 2, 3, 4, and 7 pass.
- **Cleanup ownership:** this case owns only its two generated outputs and this record.

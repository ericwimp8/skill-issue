# Safe Share Redactor Evaluation Contract

- **Target:** `showcase-skills/safe-share-redactor/skill/safe-share-redactor/SKILL.md`
- **Goal:** deliberately produce a structure-preserving sanitized copy and auditable findings from supplied text while preserving the source and exposing contextual review limits.
- **Intended use:** explicit safer-sharing preparation for supplied text files, logs, configuration, or diagnostics.
- **Expected behavior:** execute the bundled deterministic owner; protect originals and existing outputs; replace supported patterns consistently; report raw-value-free deterministic findings; distinguish unchanged ambiguous risks; require complete review before sharing; and preserve the limitation boundary.
- **Expected result:** a new sanitized artifact, a findings report, byte-identical source verification, deterministic and ambiguous counts, and an unresolved-risk conclusion that makes no privacy guarantee.
- **Preserved boundaries:** no claim of complete privacy or secrecy, no raw matched values in findings, no invented contextual redaction, no source overwrite, and no evaluation against real secrets or non-permitted identities.
- **Evaluation surface:** executable bundled Python script plus generated text and JSON artifacts.

## Observable Criteria

1. The target invokes `scripts/redact.py` as the deterministic behavior owner.
2. The input hash remains unchanged and existing output artifacts are never overwritten.
3. Every supported match receives its stable replacement without removing unrelated structure.
4. Findings retain rule, position, and replacement without the matched value or a value-derived fingerprint.
5. A clean input remains byte-identical and reports zero findings.
6. An ambiguous contextual risk remains unchanged, appears separately in findings, and requires review.
7. Output and reporting state that automated redaction cannot guarantee complete privacy or secrecy.

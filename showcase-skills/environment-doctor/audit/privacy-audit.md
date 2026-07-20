# Privacy Audit

- Scope: every retained file under `showcase-skills/environment-doctor/`.
- Repository identities: no personal or business identity is retained.
- Paths: durable authored content uses repository-relative paths; supplied-root evidence uses `<root>`; home evidence uses `~` or a generalized external-toolchain label.
- Sensitive values: selected non-PATH values are omitted; synthetic values appear only in initiating requests or validation source where they are intentionally public fixtures, never in generated diagnostic evidence.
- Machine evidence: no absolute checkout path, home username, temporary validation root, or literal home directory remains in retained artifacts.
- Interrupted work: generated ambient output from excluded qualification probe 1 was removed; its sanitized interruption record remains.
- Decision: pass after evidence-owner sanitation and final repository-privacy scan.

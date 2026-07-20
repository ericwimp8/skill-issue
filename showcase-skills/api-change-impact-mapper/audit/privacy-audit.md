# Privacy Audit

- Durable artifacts use repository-relative paths.
- `workflow-prompt.md` contains the normalized repository-relative destination and omits orchestration-only ownership and scheduling notes.
- A recursive scan found no personal usernames, home-directory paths, machine-specific checkout paths, or disallowed identities in the showcase workspace.
- Synthetic package and organization names use identity-neutral `@example` values.
- Native evidence records fresh agent paths, model/reasoning, repository-relative file paths, and content hashes only.
- No external private source, authenticated external runner output, customer data, secret, or credential is present.

Result: PASS.

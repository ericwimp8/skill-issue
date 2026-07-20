# Release Readiness Checker Privacy Audit

- The initiating prompt uses `<repo-root>` instead of a checkout path and omits orchestration-only team notes.
- Durable workflow, skill, fixture, evidence, and audit artifacts use repository-relative paths; native session storage is referenced only as the permitted home-relative `~/.codex/sessions/` location.
- Synthetic fixture identities use reserved or anonymous values and contain no personal or business identity.
- Two fresh-agent reports initially recorded the machine-specific working directory. Those strings were normalized to `<repo-root>` before conclusion without changing their command evidence.
- The final privacy scan must find no username, home-directory name, or absolute checkout path under this showcase workspace.

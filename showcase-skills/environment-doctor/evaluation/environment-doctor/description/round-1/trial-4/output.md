# Environment Doctor Trial 4

## Result

- `node` resolved to `<root>/toolchain-primary/bin/node` at PATH index `0`.
- The selected executable reported version `20.11.1`.
- `DIAGNOSTIC_SECRET` was observed as `set`; its value was omitted from diagnostic evidence.
- The primary-before-secondary expectation remains unresolved because the primary directory occurs twice, at indices `0` and `2`, while the secondary directory occurs at index `1`.
- Exit `1` is expected for this warning-bearing duplicate state.

## Guidance

The evidence confirms which executable won this bounded inspection, but the duplicate PATH entries prevent a unique precedence conclusion. Re-run the same diagnostic after any user-approved PATH correction. Obtain consent before editing configuration, changing PATH, exporting variables, or applying any setup change.

## Scope

This was a read-only POSIX inspection of the synthetic fixture. No environment or source files were changed.

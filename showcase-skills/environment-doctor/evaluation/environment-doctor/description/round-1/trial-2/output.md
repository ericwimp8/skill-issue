# Environment Diagnostic

## Findings

- `absent-tool` is unavailable on the diagnostic child's PATH.
- `INTENTIONALLY_UNSET` is unset; its value was omitted.
- PATH precedence is correct: `toolchain-primary/bin` is index `0` and `toolchain-secondary/bin` is index `1`.
- Node resolves from `<root>/toolchain-primary/bin/node` as `20.11.1`.
- `mismatch.node-version` declares Node `22`, producing a bounded version mismatch.

## Guidance

Confirm the intended Node version owner before selecting a different toolchain or updating the declaration. Obtain approval before installing `absent-tool`, changing PATH, setting the variable, or editing configuration. Re-run the same diagnostic after any approved remediation.

## Boundary

This was a read-only POSIX inspection on macOS `26.2`, Darwin `25.2.0`, `arm64`. Its colon-separated PATH and executable-bit results are limited to that tested platform boundary.

Structured details are preserved in `output/evidence.json`; the generated concise report is `output/report.txt`.

# Environment Doctor Follow-up

## Observations

- The diagnostic ran on `Darwin 25.2.0 arm64` with POSIX PATH semantics and exited `0`.
- Every executable `node` candidate, in PATH order, was:
  1. `<root>/toolchain-primary/bin/node` at index 0
  2. `<root>/toolchain-secondary/bin/node` at index 1
  3. `~/.volta/bin/volta-shim` at index 3
  4. `~/.volta/bin/volta-shim` at index 5
  5. `~/.volta/bin/volta-shim` at index 10
  6. `~/.nvm/versions/node/v22.22.0/bin/node` at index 16
  7. `/usr/local/bin/node` at index 22
  8. `~/.volta/bin/volta-shim` at index 33
- `<root>/toolchain-primary/bin/node` won resolution because it was the first candidate.
- The winning executable reported `v20.11.1`; the normalized version was `20.11.1`.
- Primary preceded secondary at PATH indices 0 and 1, so the requested precedence check was satisfied.
- `<root>/.node-version` declared `20`. That declaration agrees with the selected `20.11.1` version under the diagnostic's prefix comparison.
- Fixture source hashes were identical before and after the diagnostic. The inspection made no source changes.

## Proposed Remediation

No remediation is indicated by these checks. Re-run the same diagnostic command to verify this evidence after any separately approved environment change.

## Supporting Artifacts

- `output/report.txt` contains the concise owner-generated findings.
- `output/evidence.json` contains structured candidates, selection, probe, PATH-order, and version-file evidence.
- `native-evidence.log` contains the sanitized command, exit, host, and hashes.

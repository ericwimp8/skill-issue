# Body Cycle 2 Case 3 Audit

## Result

Pass. The refined direct harness executed exactly once and completed all assertions, printing `environment-doctor script validation passed`. The harness process exited `0`. The surrounding evidence-capture shell exited `1` only after harness completion because `status` is a read-only zsh parameter; the harness was not rerun.

## Hashes

| Artifact           | SHA-256 before                                                     | SHA-256 after                                                      | Preserved |
| ------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | --------- |
| Target `SKILL.md`  | `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203` | `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203` | Yes       |
| Diagnostic script  | `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430` | `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430` | Yes       |
| Validation harness | `89108ef5cda87a68a1050d742a2f51d6f4287b05db1bb3d3964cfc81bd545a72` | `89108ef5cda87a68a1050d742a2f51d6f4287b05db1bb3d3964cfc81bd545a72` | Yes       |

## Execution

- Command: `python3 showcase-skills/environment-doctor/script-validation/validate_diagnose.py`
- Harness exit: `0`
- Harness stdout: `environment-doctor script validation passed`
- Harness stderr: empty
- Capture-wrapper exit: `1`, after the harness process completed, from `zsh:8: read-only variable: status`.

The harness prints its success line only after every required assertion completes; an assertion failure raises and prevents that line.

## Criterion Results

| Criterion                      | Result | Evidence                                                                                                                                                |
| ------------------------------ | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Determinism                    | Pass   | Two equivalent successful runs produced byte-identical `evidence.json` and `report.txt`.                                                                |
| Source preservation            | Pass   | Target, diagnostic script, and harness hashes match before and after. The copied fixture tree hash also remained equal after success and failure cases. |
| Caller environment             | Pass   | The harness copied the caller-supplied environment mapping and asserted exact mapping equality after execution.                                         |
| Existing-output collision      | Pass   | Reusing the first output exited `1` and included `must not already exist`.                                                                              |
| Inside-root output             | Pass   | The request exited `1` and created no output directory inside the inspected root.                                                                       |
| Version-file traversal         | Pass   | `../outside` exited `1` and created no output directory.                                                                                                |
| Malformed tool selector        | Pass   | `--tool ../invalid` exited `1` and created no output directory.                                                                                         |
| Malformed environment selector | Pass   | `--env INVALID=NAME` exited `1` and created no output directory.                                                                                        |
| Unsupported tool               | Pass   | `mystery-tool` retained the `unsupported` version state and its execution sentinel remained absent.                                                     |
| Literal-home absence           | Pass   | Both generated output types were explicitly checked for the literal `Path.home()` bytes and contained none.                                             |
| Temporary-root normalization   | Pass   | Both generated output types omitted the harness temporary path.                                                                                         |
| Synthetic-value omission       | Pass   | Both generated output types omitted `synthetic-do-not-emit`; the structured selected environment value remained `null`.                                 |

The successful run also preserved mismatch, unavailable, unset, reversed PATH, missing declaration, and duplicate PATH states with warning exit `1`. No remediation was applied.

## Platform Boundary

Tested on macOS `26.2` build `25C56`, Darwin `25.2.0` arm64, with Python `3.14.2`. Conclusions are limited to this exact POSIX/macOS host and the diagnostic owner's colon-separated PATH and executable-bit behavior.

## Privacy And Cleanup

The retained files contain repository-relative paths only. The harness asserted the absence of its temporary root, the literal home path, and the synthetic selected value from both generated output types. Its `TemporaryDirectory` removed all generated diagnostic outputs; a post-run residue check found zero `environment-doctor-validation-*` directories. This case retains only `request.md`, `audit.md`, and `native-evidence.log`.

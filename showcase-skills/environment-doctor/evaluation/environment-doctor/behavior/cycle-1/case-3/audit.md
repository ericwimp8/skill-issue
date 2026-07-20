# Body Cycle 1 Case 3 Audit

## Result

Pass with two direct-harness coverage limits. The harness executed once from clean temporary fixtures and exited `0` with `environment-doctor script validation passed`.

## Hashes

| Artifact           | SHA-256 before                                                     | SHA-256 after                                                      | Preserved |
| ------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | --------- |
| Target `SKILL.md`  | `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203` | `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203` | Yes       |
| Diagnostic script  | `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430` | `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430` | Yes       |
| Validation harness | `5d35613252192ff0e56175666f3ecbe94a04ba96348ca5632f2a29d0a257a3b7` | `5d35613252192ff0e56175666f3ecbe94a04ba96348ca5632f2a29d0a257a3b7` | Yes       |

The harness copied the source fixture into a temporary root, hashed that copy before execution, and asserted the same tree hash after success and failure cases. It also copied the caller environment mapping before execution and asserted exact mapping equality afterward.

## Criterion Comparison

| Criterion                 | Result                           | Evidence                                                                                                                                                                                     |
| ------------------------- | -------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Determinism               | Pass                             | Two equivalent successful runs produced byte-identical `evidence.json` and `report.txt`.                                                                                                     |
| Source preservation       | Pass                             | Target, script, and harness hashes matched before and after; the temporary fixture tree hash remained unchanged.                                                                             |
| Existing-output collision | Pass                             | Reusing `output-one` exited `1` and emitted the required `must not already exist` error.                                                                                                     |
| Inside-root output        | Pass                             | The request exited `1` and did not create `diagnostic-output` in the root.                                                                                                                   |
| Version-file traversal    | Pass                             | `../outside` exited `1` and left no output directory.                                                                                                                                        |
| Unsupported tool          | Pass                             | `mystery-tool` retained `unsupported`, and its execution sentinel was absent.                                                                                                                |
| Privacy normalization     | Pass within exercised assertions | Output omitted the synthetic secret and temporary root. Source inspection confirms root and home normalization. The harness does not explicitly assert that the literal home path is absent. |
| Caller environment        | Pass                             | The harness asserted exact equality with its pre-run environment mapping.                                                                                                                    |
| Invalid selector names    | Uncovered by direct harness      | Argument validation exists for tool and environment-variable names, but this harness does not execute malformed-name cases.                                                                  |

The successful harness also retained mismatch, unavailable, unset, reversed PATH, missing declaration, and duplicate PATH states with warning exit `1`. No remediation was applied.

## Platform Boundary

Tested on macOS `26.2` build `25C56`, Darwin `25.2.0` arm64, with Python `3.14.2`. Conclusions are limited to this exact POSIX/macOS host and the script's colon-separated PATH and executable-bit behavior.

## Privacy And Cleanup

The harness asserted that neither its temporary path nor `synthetic-do-not-emit` appeared in retained diagnostic bytes. Its `TemporaryDirectory` owned and removed all generated diagnostic outputs. This case owns only `audit.md` and `native-evidence.log`; the pre-existing `request.md` remains unchanged.

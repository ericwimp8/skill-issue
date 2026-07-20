# Body Case 1 Audit

## Result

**Pass.** The bundled `diagnose.py` owner completed two equivalent read-only inspections with exit code `0`; both output pairs are byte-identical and match the cycle ground truth.

## Inputs And Execution

- Target SHA-256 before and after: `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203`.
- Script SHA-256 before and after: `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430`.
- Fixture-tree SHA-256 before and after: `65d1b7735c70533a8128c3d10341ae74f6a9bfd1d7181ed5dd80e002f48886b4`; individual fixture hashes are retained in `native-evidence.log`.
- Each command resolved the fixture root canonically, then scoped primary-before-secondary `PATH` and `SYNTHETIC_SELECTED` to the diagnostic child process only.
- Exact commands, exits, host evidence, hashes, comparisons, scans, and source checks are retained in `native-evidence.log`.

## Criterion Comparison

| Expected criterion                             | Observed evidence                                                                                 | Result |
| ---------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------ |
| Bundled script is the owner                    | Both commands invoked the target's `scripts/diagnose.py`                                          | Pass   |
| Equivalent runs exit zero                      | Run A `0`; Run B `0`                                                                              | Pass   |
| Outputs are byte-identical                     | `cmp -s` exited `0` for both `report.txt` and `evidence.json`                                     | Pass   |
| Node primary selection and duplicate candidate | Primary selected at PATH index `0`; secondary retained at index `1`; version `20.11.1`            | Pass   |
| Primary precedes secondary                     | PATH-order state `satisfied`, indices `0` and `1`                                                 | Pass   |
| `.node-version` matches                        | Declared `20`, actual `20.11.1`, state `match`                                                    | Pass   |
| Unsupported tool is not executed               | `mystery-tool` resolved with version state `unsupported`; execution marker remained absent        | Pass   |
| Selected value is omitted                      | Environment state `set`, stored value `null`, finding records `value_omitted: true`               | Pass   |
| Source remains unchanged                       | Target, script, and fixture hashes match before/after; scoped `git diff --exit-code` returned `0` | Pass   |

## Determinism And Privacy

- `report.txt` SHA-256 for both runs: `b032f6cf52ccd46d7f43cab14a6693a3231cfe0b2b7b3aaa665a561de468c3b6`.
- `evidence.json` SHA-256 for both runs: `a9d2b8cf74386f96869c2a6bb7405a4956f0a4f755f26290a191e09051d7bee4`.
- Retained diagnostic outputs contain neither the selected synthetic value nor the canonical fixture path; root paths are normalized to `<root>`.
- All five findings are informational. The unsupported version remains an explicit unknown; no environment repair is proposed or performed.

## Platform And Cleanup Boundary

- Tested host: Darwin `25.2.0` arm64 with Python reporting `os.name == "posix"`.
- Conclusions are limited to the tested POSIX/macOS colon-separated PATH and executable-bit behavior.
- This case owns only `run-a/`, `run-b/`, `audit.md`, and `native-evidence.log`; `request.md` was pre-existing. No temporary workspace was created, and cleanup must not extend beyond these owned artifacts.

# Case 2 Audit

## Result

The bundled `diagnose.py` owner ran once against the canonical synthetic fixture root. Its child process received primary-then-secondary fixture PATH entries and an explicitly unset `INTENTIONALLY_UNSET`. Exit `1` is the expected retained diagnostic result because the completed outputs contain warning and error findings.

## Criterion Comparison

| Criterion                  | Result | Evidence                                                                                                                                                                                  |
| -------------------------- | ------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Bundled owner executed     | Pass   | The command invokes the target's `scripts/diagnose.py`; its SHA-256 is recorded before and after.                                                                                         |
| Complete output retained   | Pass   | `output/report.txt` and `output/evidence.json` were written despite exit `1`.                                                                                                             |
| Missing command preserved  | Pass   | `absent-tool` is `unavailable`, has no selected candidate, and produces an error finding.                                                                                                 |
| Unset variable preserved   | Pass   | `INTENTIONALLY_UNSET` is `unset`; its value is `null` and the finding records `value_omitted: true`.                                                                                      |
| Requested order evaluated  | Pass   | The requested secondary-before-primary relationship is `reversed`: secondary index `1`, primary index `0`.                                                                                |
| Version mismatch preserved | Pass   | Selected Node is primary `20.11.1`; `mismatch.node-version` declares `22`; state is `mismatch`.                                                                                           |
| Safe guidance retained     | Pass   | Findings propose owner-confirmed installation, PATH, variable, or version-declaration action and provide verification or rerun commands. Any environment change requires approval.        |
| Source preserved           | Pass   | Target hashes and the complete five-file fixture manifest match before and after.                                                                                                         |
| Privacy bounded            | Pass   | Root and home paths are normalized in diagnostic output; the selected variable's value is omitted; the retained-file privacy scan found no absolute home path or synthetic secret marker. |
| Platform bounded           | Pass   | Tested host is `Darwin 25.2.0 arm64`; conclusions are limited to the script's POSIX colon-separated PATH and executable-bit behavior.                                                     |

## Observed Facts

- `absent-tool` remains unavailable.
- Node resolves first from `<root>/toolchain-primary/bin/node` and reports `20.11.1`.
- `INTENTIONALLY_UNSET` remains unset without a recorded value.
- Secondary-before-primary is reversed under the intentionally supplied primary-then-secondary PATH.
- Node `20.11.1` mismatches the declaration `22`.
- The diagnostic changed neither the fixture nor the selected child environment contract.

## Remediation Boundary

Confirm the intended owners of the required tool, PATH ordering, selected variable, and Node declaration before changing any of them. Installation, PATH edits, variable configuration, declaration edits, sourcing, or export require approval. Verification is a rerun of the same bounded diagnostic after an approved change.

## Cleanup Ownership

This case retains only `output/report.txt`, `output/evidence.json`, `audit.md`, and `native-evidence.log`. The case evaluator owns and removes its `/tmp/environment-doctor-case2-*` capture files after validation.

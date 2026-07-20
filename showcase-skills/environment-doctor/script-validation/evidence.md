# Direct Script Validation Evidence

## Surface

- Script: `showcase-skills/environment-doctor/skill/environment-doctor/scripts/diagnose.py`
- Harness: `showcase-skills/environment-doctor/script-validation/validate_diagnose.py`
- Fixtures: `showcase-skills/environment-doctor/fixtures/script/root/`
- Host boundary: Python 3 on macOS using POSIX PATH and executable-bit behavior.

## Proved Behavior

- Two equivalent clean runs produced byte-identical `report.txt` and `evidence.json` files.
- The fixture tree hash remained unchanged across success, warning, collision, inside-root, and traversal cases.
- The caller-owned environment mapping remained unchanged.
- The registered synthetic `node` resolved from two PATH candidates, executed the fixed `--version` probe, reported `20.11.1`, satisfied the requested PATH order, and matched `.node-version` value `20` with exit code zero.
- The second declaration retained actual `20.11.1` and declared `22`, wrote complete evidence, and returned exit code one.
- Additional warning cases preserved unavailable-tool, unset-variable, reversed-PATH, duplicate-PATH, and missing-declaration states with exit code one.
- The unregistered `mystery-tool` resolved with version state `unsupported`; its executable side-effect marker was absent, proving it was not executed.
- Selected `DIAGNOSTIC_SECRET` content, the test temporary prefix, the literal home path, and unnormalized supplied-root paths were absent from both retained output types.
- Existing output, output inside the inspected root, version-file traversal, and malformed tool or environment-variable selectors failed before creating or overwriting diagnostic artifacts.
- `python3 -m py_compile` passed for the script and validation harness.

## Refinement Evidence

The initial validation exposed macOS `/var` to `/private/var` canonicalization as a concrete behavior-owner failure: PATH entries under the supplied root were reported as external paths, and requested order checks became `missing`. The path normalizer and PATH comparison owner now compare canonical paths while retaining normalized presentation. The full direct harness passed after that update.

## Result

`python3 showcase-skills/environment-doctor/script-validation/validate_diagnose.py` returned `environment-doctor script validation passed`.

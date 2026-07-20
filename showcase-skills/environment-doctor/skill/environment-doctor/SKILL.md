---
name: environment-doctor
description: Read-only diagnostics for selected local development tools, paths, environment state, and version declarations. Use when investigating local setup, executable resolution, PATH precedence, or tool-version mismatches.
---

# Environment Doctor

## Define the Inspection

- Select an existing development root and a new output directory under an existing parent outside that root.
- Select only the tools and environment-variable names relevant to the reported problem. Add PATH-order expectations and version declaration files only when the project establishes them.
- Do not request secrets or pass environment-variable values as arguments. The script reports selected non-PATH variables only as `set`, `empty`, or `unset`.

## Run the Diagnostic Owner

Run the bundled script from this skill directory, adding repeatable selectors as needed:

```sh
python3 scripts/diagnose.py \
  --root <development-root> \
  --output-dir <new-output-directory> \
  --tool <tool-name> \
  --env <variable-name> \
  --expect-path-before <earlier-directory> <later-directory> \
  --version-file <tool-name> <relative-version-file>
```

At least one selector is required. Relative PATH expectation directories resolve from the development root. Version files must remain within that root. The script writes only `report.txt` and `evidence.json` under a newly created output directory; it refuses collisions and output locations inside the inspected root.

Use the script instead of recreating resolution, probing, normalization, comparison, or finding logic. Preserve its `unavailable`, `unsupported`, `timeout`, `unparseable`, `missing`, and `duplicate` states as unknowns or unresolved evidence.

## Interpret and Report

1. Read `report.txt` for the concise findings and `evidence.json` for structured follow-up evidence.
2. Separate observed facts from remediation proposals. A reported mismatch establishes only the bounded comparison performed by the script.
3. Explain warning and error findings, including the recorded evidence, safe remediation guidance, and verification command or rerun guidance.
4. State unsupported surfaces and unresolved unknowns. Version execution is limited to the script's fixed registry; unregistered tools are resolved without execution.
5. Ask before installing tools, editing configuration, changing PATH, exporting variables, sourcing shell files, or otherwise altering the environment.

Treat the diagnostic as a read-only POSIX inspection. Report the exact tested host separately; do not generalize the script's colon-separated PATH and executable-bit behavior to unsupported platforms.

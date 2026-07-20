# Environment Doctor Evaluation Contract

- **Target:** `showcase-skills/environment-doctor/skill/environment-doctor/SKILL.md`
- **Goal:** produce bounded, actionable, read-only evidence about selected development-environment surfaces without altering or overclaiming the environment.
- **Intended use:** investigating local tool availability and versions, executable resolution, PATH precedence, selected environment state, and selected version declarations.
- **Expected behavior:** select only relevant inputs; invoke the deterministic owner; preserve unavailable and ambiguous states; protect sensitive values and paths; separate facts from remediation; ask before changes; bound platform conclusions.
- **Expected result:** stable human and JSON outputs with equivalent finding identifiers, severity, evidence, remediation, and verification, plus clear unknowns and no source or environment mutation.
- **Preserved boundaries:** no environment repair, arbitrary tool execution, selected non-PATH value disclosure, version-file access outside the root, or claims beyond tested POSIX behavior.
- **Evaluation surface:** implicit skill selection plus executable Python behavior and generated text and JSON artifacts.

## Observable Criteria

1. Representative implicit prompts load the exact target before the agent answers or acts.
2. The target invokes `scripts/diagnose.py` as the deterministic behavior owner and selects a new output directory outside the inspected root.
3. A clean synthetic inspection reports two executable candidates, the selected fixed-registry tool version, a satisfied PATH relationship, omitted selected environment values, and a matching version declaration with exit code zero.
4. A mismatch inspection preserves both actual and declared numeric versions, reports a warning with safe remediation and verification, and exits one while writing complete outputs.
5. An unsupported selected tool is resolved without execution and retains an `unsupported` version state.
6. Repeated equivalent runs produce byte-identical output; input hashes and the caller-supplied environment mapping remain unchanged.
7. Existing output, inside-root output, invalid selections, and version-file traversal fail before creating or overwriting diagnostic artifacts.
8. Root and home paths are normalized, selected non-PATH values are omitted, and retained outputs contain no synthetic secret or machine-specific temporary prefix.
9. The agent distinguishes findings from proposed remediation, preserves unknowns, asks before changes, and reports the exact tested POSIX/macOS boundary.

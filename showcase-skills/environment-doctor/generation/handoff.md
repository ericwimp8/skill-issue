# Environment Doctor Generation Handoff

- **Generated skill:** `showcase-skills/environment-doctor/skill/environment-doctor/`
- **Canonical target:** `showcase-skills/environment-doctor/skill/environment-doctor/SKILL.md`
- **Bundled deterministic owner:** `showcase-skills/environment-doctor/skill/environment-doctor/scripts/diagnose.py`
- **Intake contract:** `showcase-skills/environment-doctor/plans/environment-doctor/environment-doctor-a-to-b-plan.md`
- **Supported surfaces:** portable Agent Skills content and implicit or explicit OpenAI Codex project or user delivery; script execution requires Python 3 with POSIX PATH semantics.
- **Goal:** produce actionable read-only evidence about selected tool, executable, PATH, environment-state, and version-declaration surfaces without changing the inspected environment.
- **Intended use:** local development setup, resolution, PATH-precedence, and tool-version mismatch investigations.
- **Expected behavior:** define a bounded selection; run the bundled owner; preserve unknowns and failures; distinguish observations from proposals; explain safe remediation and verification; obtain consent before changes.
- **Expected result:** stable `report.txt` and `evidence.json` outputs containing normalized selected observations, equivalent findings, explicit unknowns, and no selected non-PATH values.
- **Preserved boundaries:** no installation, configuration edits, shell sourcing, environment mutation, secret-value reporting, unrestricted command execution, or unsupported platform generalization.
- **Runtime criteria:** all criteria in the evaluation contract require direct script or independent-agent evidence.
- **Known limitations:** unregistered tools are resolved without version execution; version comparison is numeric-prefix only; configuration inspection is limited to explicitly selected single-version declaration files; POSIX behavior is tested on macOS and synthetic fixtures.
- **Refinement mode:** automatic semantic refinement when retained evidence establishes a material contract, safety, or privacy failure.
- **Evaluation route:** `showcase-skills/environment-doctor/evaluation/environment-doctor/`.
- **Generation decision:** continue into evaluation after qualifying the exact requested `gpt-5.6-sol` medium-reasoning fresh-agent surface.

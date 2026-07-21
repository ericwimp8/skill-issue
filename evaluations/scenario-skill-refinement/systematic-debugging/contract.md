# Systematic Debugging Evaluation Contract

## Target

- Canonical skill: `evaluations/scenario-skill-refinement/systematic-debugging/skill/`
- `SKILL.md` SHA-256: `008eb3082ab95495453702d616dd82834f38dcd0639aa1d0d00ecf493254b0af`
- `agents/openai.yaml` SHA-256: `e528397046a9d6f7b64fae387412ccd85f1e682e126f35d340519981f5b5faa8`
- Environment qualification: `evaluations/skill-system-production-refinement/environment-qualification.md`

## Interpreted Contract

1. **Goal:** Establish an evidence-supported root cause before proposing or applying the smallest correction at the causal owner.
2. **Intended use:** Bugs, test failures, and unexpected behavior that require diagnosis before intervention.
3. **Expected behavior:** Reproduce the failure without changing the system; trace the concrete failing path backward; distinguish observation from origin; compare a nearby working path; test one explicit causal hypothesis; capture the smallest useful failing reproduction when appropriate; intervene only after the cause explains the relevant symptoms; and validate the correction with focused and related checks.
4. **Expected result:** A source-backed diagnosis that identifies where the incorrect condition first appears and either supports a minimal causal correction with proof or clearly stops for missing evidence.
5. **Boundary:** Preserve the no-fix-before-cause gate. Do not replace source tracing with test assertions, guess through intermittent or incomplete evidence, accept a cause that explains only part of the symptom set, or patch only the visible observation point.

## Observable Completion Criteria

- States the observed failure and the behavior that should hold.
- Reproduces the failure or identifies the exact missing evidence needed to do so.
- Traces concrete implementations from the observation point to the causal origin.
- Identifies where the incorrect value, state, or control decision first appears.
- Tests one explicit hypothesis against available evidence and relevant symptoms.
- Distinguishes the causal owner from downstream observation surfaces.
- Proposes or applies a correction only when the evidence supports the root cause.
- Names or runs a focused proof and the nearest related checks after a correction.
- Stops for more evidence when the cause is incomplete, intermittent, or merely masks the observation.

## Evaluation Surfaces

- Description loop: native proactive selection on four fresh-agent debugging tasks.
- Reference loop: not applicable because the target contains no `references/` files.
- Body loop: isolated code-investigation and incomplete-evidence cases, audited through retained reports and native traces.

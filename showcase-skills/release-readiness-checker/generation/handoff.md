# Release Readiness Checker Generation Handoff

- **Generated skill:** `showcase-skills/release-readiness-checker/skill/release-readiness-checker/`
- **Canonical target:** `showcase-skills/release-readiness-checker/skill/release-readiness-checker/SKILL.md`
- **Intake contract:** `showcase-skills/release-readiness-checker/plans/release-readiness-checker/release-readiness-checker-a-to-b-plan.md`
- **Supported surfaces:** portable Agent Skills content and OpenAI Codex project or user delivery.
- **Goal:** decide whether an exact candidate is ready using project-derived gates and current evidence without performing a release action.
- **Expected behavior:** bound the candidate; derive gates from semantic owners; trace relevant paths; preserve command evidence and limitations; classify every gate precisely; and return a prioritized, evidence-bounded decision and next actions.
- **Expected result:** `ready`, `not ready`, or `undetermined`, supported by candidate-specific gate records, unresolved risks, exact next actions, and an evidence index.
- **Preserved boundaries:** no generic policy substituted for project authority, no stale evidence treated as current execution, no incomplete gate treated as passed, and no publication or deployment side effect.
- **Runtime criteria:** every intake completion criterion requires representative behavior evidence across passing, failing, blocked, and not-run conditions.
- **Known limitations:** candidate-specific access, credentials, environments, or release policy may remain unavailable and must constrain the decision.
- **Refinement mode:** automatic semantic refinement supported by retained material failures.
- **Evaluation route:** `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/`.
- **Generation decision:** continue into evaluation after the required medium-reasoning Codex environment qualification is retained.

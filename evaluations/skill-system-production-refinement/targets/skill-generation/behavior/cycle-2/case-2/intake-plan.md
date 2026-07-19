# Autonomous Generation Contract

- Desired skill: Codex `decision-log-summary` that reads selected Markdown ADRs and reports decision, status, and consequences without editing sources.
- Destination: `evaluations/skill-system-production-refinement/targets/skill-generation/behavior/cycle-2/case-2/output/decision-log-summary`.
- Completion: valid minimal skill, `Not stated` for missing fields, source order preserved, no unnecessary resources, structural validation, Evaluation handoff.
- Generation viability: Autonomous.
- Selected execution preference: Autonomous.
- Authority boundary: create only inside the destination and run structural checks.
- Runtime proof: defer invocation, extraction accuracy, missing fields, and immutability to Evaluation.

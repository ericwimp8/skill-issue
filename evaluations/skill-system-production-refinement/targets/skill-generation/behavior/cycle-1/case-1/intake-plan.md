# Release Operator Generation Contract

## Desired skill

Create a Codex-only `release-operator` skill at `output/release-operator`. It follows repository-owned release sources, preserves dry-run by default, reports validation and proposed-version evidence, and requires explicit authorization before publishing.

## Completion criteria

- Valid Codex skill structure.
- Source-first release discovery.
- No inferred publish authorization.
- Missing validation or version owners remain blockers.
- Runtime proof is handed to Evaluation.

## Generation contract

- Generation viability: Autonomous.
- Selected execution preference: Step-by-step.
- Required stop: pause for approval before creating the first skill file and at later step boundaries.
- Current authority: inspect this contract and propose the first action only.
- Destination: `evaluations/skill-system-production-refinement/targets/skill-generation/behavior/cycle-1/case-1/output/release-operator`.

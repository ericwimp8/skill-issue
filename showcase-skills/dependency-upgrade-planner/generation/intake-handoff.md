# Intake-to-Generation Handoff

## Source of Truth

- Plan: `showcase-skills/dependency-upgrade-planner/plans/dependency-upgrade-planner/dependency-upgrade-planner-a-to-b-plan.md`
- Initiating prompt: `showcase-skills/dependency-upgrade-planner/workflow-prompt.md`

## Confirmed Contract

- **Purpose:** turn a requested dependency upgrade into a source-backed, dependency-ordered migration plan.
- **Outcome:** a planning-only artifact connecting repository evidence and authoritative upstream requirements to ordered prerequisites, migration work, validation, rollback, and unresolved risk.
- **Invocation boundary:** requested planning or impact analysis for upgrading one or more named dependencies in an existing project.
- **Destination:** `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/`.
- **Supported surface:** portable Agent Skills content plus OpenAI Codex metadata.
- **Unresolved implementation matters:** exact body organization, metadata copy, and evaluation fixtures.
- **User-owned stops:** none identified.
- **Generation viability:** autonomous.
- **Execution preference:** autonomous continuation.
- **Authority boundary:** create and refine only this showcase workspace; do not edit dependencies or unrelated repository content.
- **Evaluation expectation:** continue into governed description and body evaluation after current-environment qualification.

The complete plan remains the source of truth for expected behavior, boundaries, exclusions, and observable completion criteria.

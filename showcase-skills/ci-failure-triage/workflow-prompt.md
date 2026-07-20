# CI Failure Triage Workflow Prompt

```text
Work in `<repo-root>`. Create a reusable skill named `ci-failure-triage` by running the complete current project-local Skill Issue workflow from `skills/skill-intake/SKILL.md` through `skills/skill-generation/SKILL.md` and `skills/skill-evaluation-and-refinement/SKILL.md`, applying the supporting disciplines those sources require.

The skill investigates failed CI runs from supplied logs, workflow configuration, repository source, and available local tooling; distinguishes the primary failure from cascading noise; produces an evidence-backed diagnosis, smallest responsible remediation direction, and exact verification plan; preserves uncertainty and unavailable evidence; and never modifies remote CI state, secrets, branches, or releases without explicit authorization.

Use `showcase-skills/ci-failure-triage/` as the repository-relative destination. Retain the initiating prompt plus all governed intake, generation, skill, resource, validation, evaluation, refinement, fixture, evidence, and audit artifacts there. Create only resources genuinely required by the outcome. Use synthetic or repository-owned fixtures and public authoritative sources. Durable artifacts must use repository-relative paths and comply with `.repository-privacy.md`.

Continue autonomously through intake, generation, evaluation, and evidence-supported automatic refinement while the governing contracts permit. Use fresh independent GPT-5.6 Sol agents with medium reasoning and isolated context for trials. Do not use an authenticated external CLI runner if it requires escalation or user approval. Do not modify production workflow skills, supporting skills, CLI code, website code, or any other showcase workspace. Run required structural, formatting, diff, hash, and privacy checks. Report exact paths, evidence, refinements, blockers, and workflow weaknesses.
```

# Incident Timeline Builder Workflow Prompt

```text
Work in `<repo-root>`. Create a reusable skill named
`incident-timeline-builder` by running the complete current project-local Skill
Issue workflow from `plugins/skill-issue/skills/skill-intake/SKILL.md` through
`plugins/skill-issue/skills/skill-generation/SKILL.md` and
`plugins/skill-issue/skills/skill-evaluation-and-refinement/SKILL.md`, applying the supporting
disciplines those sources require.

The skill must deterministically turn supplied logs, alerts, deployment
records, and operator notes into a chronological incident timeline. Normalize
explicit timestamps and time zones without inventing missing times; retain
source provenance; distinguish observed events, reported statements, inferred
relationships, contradictions, and gaps; avoid converting correlation into
causation; and produce a useful timeline plus unresolved-evidence and follow-up
sections while preserving source files.

Use `showcase-skills/incident-timeline-builder/` as the destination. Retain the
governed intake, generation, skill, resource, script-validation when
applicable, evaluation, refinement, fixture, evidence, and audit artifacts
there. Use synthetic or repository-owned fixtures, repository-relative paths,
and the repository privacy contract. Do not modify other repository surfaces.

Let the workflow determine whether a deterministic script or any reference is
necessary. Directly validate every generated script for determinism, time-zone
handling, ambiguous or missing timestamps, source preservation, error behavior,
and privacy. Use fresh independent GPT-5.6 Sol agents with medium reasoning and
clean context for evaluation trials. Do not use an authenticated external CLI
runner when it requires escalation or approval.

Run required structural, formatting, diff, hash, and privacy checks. Report
exact paths, evidence, refinements, blockers, and workflow weaknesses.
```

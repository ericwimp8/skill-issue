# Release Readiness Checker Workflow Prompt

```text
Work in `<repo-root>`. Run the complete current project-local Skill Issue
workflow from `skills/skill-intake/SKILL.md` through
`skills/skill-generation/SKILL.md` and
`skills/skill-evaluation-and-refinement/SKILL.md`, applying the supporting
disciplines those sources require. Create a reusable skill named
`release-readiness-checker`.

Intended outcome: a skill that evaluates a repository or supplied release
candidate for readiness without publishing or deploying it. It should derive
gates from authoritative project source and release instructions; inspect
relevant code, configuration, migrations, versioning, documentation,
test/build evidence, security or privacy requirements, rollback preparation,
and unresolved risks; distinguish passed, failed, blocked, not-run, and
not-applicable checks; preserve command evidence and limitations; and produce
a prioritized release decision with exact next actions without inferring
readiness from stale tests or documentation.

Use `showcase-skills/release-readiness-checker/` as the repository-relative
destination. Retain the initiating prompt plus all governed intake,
generation, skill, resource, validation, evaluation, refinement, fixture,
evidence, and audit artifacts there. Create scripts, references, or assets only
when genuinely required. Use synthetic or repository-owned fixtures and
authoritative public sources when needed. Durable artifacts must use
repository-relative paths and comply with `.repository-privacy.md`.

Continue autonomously through intake, generation, evaluation, and supported
semantic refinement. Use fresh independent GPT-5.6 Sol agents with medium
reasoning and `fork_turns: "none"` for trials. Stop only at a governing
user-owned decision or capability boundary. Do not modify production workflow
skills, supporting skills, CLI code, website code, or another showcase
workspace. Do not publish or deploy anything.

Run required structural, formatting, diff, hash, and privacy checks. Report
exact paths, evidence, refinements, blockers, and workflow weaknesses.
```

# Dependency Upgrade Planner Workflow Prompt

This publication-safe record preserves the initiating skill request with the local checkout normalized to `<repo-root>`.

> Work in `<repo-root>`. Create a reusable skill named `dependency-upgrade-planner` under `showcase-skills/dependency-upgrade-planner/` by running the complete current project-local workflow from skill intake through generation and evaluation/refinement.
>
> The skill turns a requested dependency upgrade into a source-backed, dependency-ordered migration plan. It inspects production manifests, lockfiles, configuration, and concrete usages; consults authoritative release notes, migration guides, and compatibility requirements when available; distinguishes direct, transitive, build-tool, runtime, and platform effects; identifies breaking changes, prerequisites, validation gates, rollback considerations, and unresolved risks; and avoids editing dependencies or claiming compatibility without evidence.
>
> Retain all governed workflow, skill, validation, evaluation, fixture, evidence, refinement, and audit artifacts within the showcase workspace. Use synthetic or repository-owned fixtures and authoritative public sources. Create scripts, references, or assets only when genuinely required. Durable artifacts use repository-relative paths and comply with `.repository-privacy.md`.
>
> Do not modify production workflow skills, supporting skills, CLI code, website code, or other showcase work. Run required structural, formatting, diff, hash, and privacy checks. Report exact paths, evidence, refinements, blockers, and workflow weaknesses.

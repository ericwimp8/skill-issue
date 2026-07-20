# Accessibility First-Pass Workflow Prompt

This publication-safe record preserves the initiating prompt used for the showcase workflow, with the local checkout normalized to `<repo-root>`.

> Work in `<repo-root>`. You are responsible only for a new
> `accessibility-first-pass` showcase workspace and artifacts created by this
> run. You are not alone in the repository: do not revert, overwrite, reformat,
> or clean up unrelated existing changes, and accommodate concurrent work.
>
> Explicitly invoke and follow the project-local `$skill-intake` skill to create
> a reusable agent skill named `accessibility-first-pass`.
>
> The intended outcome is a skill that performs a responsible first-pass
> accessibility review of a supplied web page, feature, or implementation. It
> should investigate the available source, rendered behavior, project tooling,
> and applicable authoritative accessibility guidance; combine appropriate
> automated checks with manual inspection; distinguish observed evidence from
> inference and unverified behavior; and produce a prioritized, actionable
> report. The report should identify review scope, affected users, evidence,
> reproduction or inspection steps, remediation direction, checks that require
> human or assistive-technology testing, and material limitations. It must not
> present a first-pass review or automated scan as proof of accessibility or
> standards conformance.
>
> Use `showcase-skills/accessibility-first-pass/` as the repository-relative
> intended destination. Retain the finished skill and the planning, generation,
> validation, evaluation, refinement, fixtures, evidence, and audit artifacts
> created by this task under that workspace when the governing skills permit.
> Follow the named Skill Issue skills rather than recreating their procedures.
>
> Use synthetic or repository-owned fixtures and public authoritative sources.
> Never place real secrets, personal identities, business identities, usernames,
> home-directory names, or machine-specific checkout paths into
> public-repository artifacts. Durable Markdown and generated evidence must use
> repository-relative paths. The only permitted public project identities are
> `Eric Wimp` and `ericwimp8`.
>
> You are authorized to continue autonomously from intake through
> `$skill-generation` and into `$skill-evaluation-and-refinement` when their
> contracts and available prerequisites permit. Stop only at a user-owned
> decision or capability boundary required by those skills.
>
> Do not modify the production workflow skills under `skills/`, canonical
> supporting skills under `supporting-skills/`, CLI code, website code, or
> existing showcase work. Report any workflow weakness for later review instead.
>
> Completion means the workflow has progressed as far as its own gates allow;
> the generated skill has been structurally and behaviorally evaluated where the
> available environment permits; all created artifacts are retained at governed
> repository-relative locations; privacy requirements have been audited; and
> your final response reports exact repository-relative paths, validation and
> evaluation evidence, any refinement performed, any stopping gate, and any
> observed workflow weakness.

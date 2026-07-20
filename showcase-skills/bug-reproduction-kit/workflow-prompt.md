# Bug Reproduction Kit Workflow Prompt

This is the initiating prompt used to run the complete Skill Issue intake, generation, and evaluation workflow for `bug-reproduction-kit`. The machine-specific checkout path has been normalized to `<repo-root>` for publication.

```text
Work in `<repo-root>`. You are responsible only for the new
`bug-reproduction-kit` showcase workspace and artifacts created by this run.
You are not alone in the repository: do not revert, overwrite, or clean up
unrelated existing changes, and accommodate concurrent work.

Explicitly invoke and follow the project-local `$skill-intake` skill to create
a new agent skill named `bug-reproduction-kit`.

The intended outcome is a reusable skill that helps an agent turn incomplete
or vague bug information into a reproducible evidence package. It should
establish relevant environment facts, produce minimal reproduction steps,
distinguish expected from actual behavior, collect or identify useful logs and
artifacts, surface material missing information without inventing facts, and
produce a ready-to-file issue or equivalent reproduction package.

Use `<repo-root>/showcase-skills/bug-reproduction-kit/` as the intended
destination. Retain the finished skill and the planning, generation,
evaluation, refinement, fixtures, evidence, and audit artifacts created by
this task under that workspace when the governing skills permit. Follow the
named Skill Issue skills rather than recreating their procedures in your own
way.

You are authorized to continue autonomously from intake through
`$skill-generation` and into `$skill-evaluation-and-refinement` when their
contracts and available prerequisites permit. Stop only at a user-owned
decision or capability boundary required by those skills. Do not modify the
production workflow skills under `skills/` or the canonical supporting skills
under `supporting-skills/`; report any workflow weakness you observe for later
review instead.

Completion means the workflow has progressed as far as its own gates allow,
all created artifacts are retained at their governed locations, and your final
response reports the exact paths created, validation or evaluation evidence
obtained, any stopping gate, and any observed workflow weakness.
```

# Safe Share Redactor Workflow Prompt

This is the initiating prompt used to run the complete Skill Issue intake, generation, and evaluation workflow for `safe-share-redactor`. The machine-specific checkout path has been normalized to `<repo-root>` for publication.

```text
Work in `<repo-root>`. You are responsible only for a new
`safe-share-redactor` showcase workspace and artifacts created by this run.
You are not alone in the repository: do not revert, overwrite, reformat, or
clean up unrelated existing changes, and accommodate concurrent work.

Explicitly invoke and follow the project-local `$skill-intake` skill to create
a reusable agent skill named `safe-share-redactor`.

The intended outcome is a skill for preparing supplied text files, logs,
configuration, or diagnostic material for safer sharing. This experiment
specifically requires the generated skill to contain and use at least one
bundled script that owns deterministic redaction behavior. The script should
operate on supplied material without overwriting originals, produce sanitized
output plus auditable findings, preserve useful structure where practical,
and handle supported sensitive patterns consistently. The skill must
distinguish deterministic matches from ambiguous contextual risks and must
not claim that automated redaction guarantees complete privacy or secrecy.

Use only synthetic credentials and synthetic personal information in
evaluation fixtures. Never place real secrets, personal identities, business
identities, usernames, home-directory names, or machine-specific checkout
paths into public-repository artifacts. Durable Markdown and generated
evidence must use repository-relative paths. The only permitted public project
identities are `Eric Wimp` and `ericwimp8`.

Use `showcase-skills/safe-share-redactor/` as the repository-relative intended
destination. Retain the finished skill and the planning, generation, script
validation, evaluation, refinement, fixtures, evidence, and audit artifacts
created by this task under that workspace when the governing skills permit.
Follow the named Skill Issue skills rather than recreating their procedures.

You are authorized to continue autonomously from intake through
`$skill-generation` and into `$skill-evaluation-and-refinement` when their
contracts and available prerequisites permit. The body evaluation must
actually execute the bundled script in isolated representative cases,
including material that should be redacted, material that should remain
unchanged, and a case that exposes a supported limitation or ambiguous risk.
Stop only at a user-owned decision or capability boundary required by those
skills.

Do not modify the production workflow skills under `skills/`, canonical
supporting skills under `supporting-skills/`, CLI code, website code, or
existing showcase work. Report any workflow weakness for later review instead.

Completion means the workflow has progressed as far as its own gates allow;
the script and skill have been structurally and behaviorally validated; all
created artifacts are retained at governed repository-relative locations;
privacy requirements have been audited; and your final response reports exact
repository-relative paths, executed validation and evaluation evidence, any
refinement performed, any stopping gate, and any observed workflow weakness.
```

# Workflow Observations

## Discipline Decisions

- **Document meaning:** the intake plan owns intent and completion criteria; `SKILL.md` owns reusable behavior; generation, validation, campaign, and audit files own process evidence. No process history was placed in the installable payload.
- **System ownership:** the agent owns contextual extraction and classification; the local helper owns deterministic timestamp parsing, ambiguity checks, ordering, input hashing, and serialization; campaign fixtures own evaluation ground truth.
- **Code ownership:** `build_timeline.py` is the concrete producer of normalization and ordering behavior. Its CLI is the only caller surface, and direct validation exercises stdout, file output, refusal paths, and preserved inputs at that owner.
- **Prompt discipline:** trial requests state only the task, fixture location, output boundary, and preservation need. They do not name the skill, quote its description, expose expected answers, or reteach its workflow.

## Workflow Weaknesses

- The durable environment qualification names GPT-5.6 Sol with high reasoning, while this task requires medium reasoning. The evaluation workflow correctly prevents inheriting that evidence, but it does not provide a canonical procedure or artifact template for qualifying a new reasoning setting before description trials.
- The environment record calls the surface “Codex Desktop fresh sub-agents” while its scope refers to fresh agents and native traces more generally. A future revision could define the harness identity fields more precisely so equivalent agent tooling is not inferred from ambiguous wording.
- The campaign record requires direct native load evidence but does not specify how a parent should retain agent tool traces when the collaboration API returns only the agent identity and final result. This campaign must recover the native session record separately rather than treating agent prose as proof.

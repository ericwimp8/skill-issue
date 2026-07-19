---
name: skill-generation
description: Execution of a build-ready skill-intake plan into an idiomatic agent skill. Use when a completed intake contract is ready under its recorded viability, execution preference, and authority boundary.
---

# Skill Generation

## Accept the Intake Contract

1. Read the complete A-to-B plan and generation handoff from `skill-intake`.
2. Confirm the destination, expected outcome, completion criteria, project context, supported harness surfaces, unresolved implementation-time matters, generation viability, execution preference, authority boundary, and user-owned stop conditions.
3. Return to intake only when missing information would change user intent. Resolve ordinary implementation choices from source and authoritative platform context.

## Establish Ownership

Before creating files, read `references/generation-disciplines.md` and load every available listed discipline whose surface will change. Then map:

- the behavior the skill must change;
- the concise instructions that own that behavior;
- platform detail that belongs in conditional references;
- deterministic repeated behavior that warrants a script;
- output material that warrants an asset or template;
- host-specific metadata that belongs outside the canonical skill body.

## Generate the Skill

1. Honor the recorded generation viability, execution preference, and authority boundary. Stop for the user only at a user-owned decision or capability boundary identified by the contract.
2. Create the skill in the confirmed destination with a lowercase hyphenated name, matching folder, valid `SKILL.md`, and only genuinely required resources.
3. Write only one concise purpose sentence and one concise use-boundary sentence in the description. Keep output fields, missing-data behavior, examples, and execution detail in the body. Keep activation guidance in metadata, except where an explicit-only fallback must be carried in the description itself.
4. Keep the body behavior-changing and direct. Put platform variants, schemas, or substantial conditional detail in one-level indexed references.
5. Preserve one semantic owner for each instruction. Reconcile related rules rather than accumulating exceptions.
6. Add scripts only for deterministic or repeatedly reconstructed operations. Add assets only when the skill's outputs consume them.
7. Read `references/harness-packaging.md` when the target surface affects discovery, metadata, packaging, or invocation. Add host metadata only for supported surfaces and state unsupported capability accurately.

## Validate the Written Result

- Check frontmatter, naming, folder structure, reference paths, resource necessity, and host metadata syntax.
- Run the target harness's authoritative structural validator when one is available.
- Check every intake completion criterion against the written artifact.
- Record any criterion that requires runtime proof for the evaluation campaign.
- Treat structural validation as generation evidence only; behavior proof belongs to the evaluation handoff.

## Hand Off to Evaluation

Provide `skill-evaluation-and-refinement` with:

- the generated skill path and supported harness surfaces;
- the complete intake contract;
- the established goal, intended use, expected behavior, expected result, and boundaries;
- completion criteria requiring runtime proof;
- unresolved limitations and the selected refinement interaction mode when already known.

When the recorded authority permits continuation and the evaluation prerequisites are available, continue into `skill-evaluation-and-refinement` in the same task. Otherwise preserve the complete handoff and state the exact prerequisite or user-owned stop required to continue.

The generation result is incomplete until this handoff exists, even when runtime evaluation is deferred.

## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/generation-disciplines.md`: Semantic ownership and authoring constraints for generated skills. Use before writing or revising skill documents, prompts, code, scripts, or system structure.
- `references/generation-handoff.md`: Required intake inputs and evaluation outputs. Use when accepting a plan, checking completion, or transferring the generated skill.
- `references/harness-packaging.md`: Supported harness skill paths, package surfaces, and metadata boundaries. Use when the generated skill targets a specific harness or needs cross-harness delivery.

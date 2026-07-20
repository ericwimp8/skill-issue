# API Change Impact Mapper A-to-B Plan

## A — Current Position

- The initiating request defines a new reusable skill and the governed intake-to-generation-to-evaluation workflow.
- The destination is `showcase-skills/api-change-impact-mapper/`; all campaign-owned plans, generated files, fixtures, evidence, audits, and refinements remain inside it.
- Inputs to the finished skill are supplied old and new API, schema, event, or public-interface contracts plus a repository and any available operational context.
- The skill is analysis-only: it may inspect source and execute safe read-only discovery, but it must not modify the analyzed implementation.
- The analysis must treat production source and concrete implementations as authoritative for repository behavior, while preserving unavailable external consumers and other evidence limits as unknown.
- The target is portable Agent Skills content with OpenAI Codex discovery metadata. Implicit invocation is appropriate for naturally phrased change-impact analysis requests.
- Evaluation uses synthetic connected repository fixtures and fresh independent Codex agents running `gpt-5.6-sol` with medium reasoning.

## B — Desired Position

A ready-to-use `api-change-impact-mapper` skill traces a contract change end to end, distinguishes compatibility by affected path and deployment condition, preserves uncertainty, and returns an evidence-backed migration map that teams can assign, order, validate, roll out, and roll back without changing implementation files.

## Path From A to B

1. Define the evidence and scope boundary for supplied contracts, repository paths, runtime facts, and unavailable external surfaces.
2. Establish a contract delta before searching the repository, including semantic changes that textual diffs alone may miss.
3. Trace concrete producers, transformations, consumers, adapters, generated artifacts and their generators, persisted representations, configuration, and external boundaries in both directions.
4. Classify each impact as compatible, conditionally compatible, or breaking with explicit conditions, affected versions, and evidence confidence.
5. Turn the traced graph into an owner-oriented migration sequence with prerequisites, coexistence needs, validation, rollout observation, and rollback constraints.
6. Package concise portable instructions and only resources whose absence would impair correct execution.
7. Validate structure and completion-criterion coverage, then evaluate description selection and body behavior with isolated representative cases.
8. Apply only evidence-supported semantic refinements at the owning instruction and rerun affected evaluation from clean fixtures.

## C — Completion Criteria

1. A naturally phrased API, schema, event, or public-interface impact request selects the skill without naming it.
2. The result identifies the old-to-new semantic delta and distinguishes confirmed changes from assumptions or missing contract detail.
3. The repository trace reaches concrete producers and consumers rather than stopping at declarations, wrappers, or search matches.
4. The trace covers relevant adapters, generated clients and source generators, persisted data, configuration or flags, and external boundaries, explicitly marking any inapplicable or unverified surface.
5. Every material impact is classified as compatible, conditionally compatible, or breaking with an evidence-backed rationale and the conditions or versions that control the classification.
6. Unknown consumers and evidence limits remain visible and constrain conclusions rather than being treated as absence of impact.
7. The migration map names the responsible owner or owner type, dependency order, coexistence or sequencing requirements, and concrete validation evidence for each action.
8. Rollout guidance identifies observation signals and staged-release considerations; rollback guidance identifies data, contract, deployment, or version constraints that may make reversal unsafe.
9. The final result links claims to repository-relative evidence and separates confirmed facts, supported inferences, and unknowns.
10. No analyzed implementation file is modified.

## Generation Contract

- **Destination:** `showcase-skills/api-change-impact-mapper/skill/api-change-impact-mapper/`
- **Target harness:** portable Agent Skills content with OpenAI Codex metadata; implicit invocation remains enabled by omitting an invocation override.
- **Generation viability:** autonomous; the request and repository workflow resolve all material intent.
- **Execution preference:** autonomous continuation through generation, evaluation, refinement, and audit.
- **Authority boundary:** create and revise artifacts only inside `showcase-skills/api-change-impact-mapper/`; inspect other repository sources read-only; use automatic semantic refinement when retained evidence establishes a material failure.
- **Required user stops:** none expected; stop only at a capability boundary or after the workflow's five-failure limit.
- **External dependencies:** none; synthetic fixtures are sufficient, and current platform claims are unnecessary.
- **Unresolved implementation choices:** exact headings and fixture designs may be chosen during generation and evaluation without changing intent.
- **Evaluation handoff:** continue directly into `skill-evaluation-and-refinement` with every completion criterion requiring runtime proof.

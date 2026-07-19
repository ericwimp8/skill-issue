# Skill Issue Documentation Authority and Update Map

## Purpose

Prevent stale or competing project meaning by identifying each durable document's role, semantic owner, consumers, and required update route.

Read this map with:

- `skill-issue-project-completion-a-to-b-plan.md` for the current six-block sequence;
- `reorganization-dependency-audit.md` for completed-work inventory, mismatches, retained research, and old-task mapping.

## Authority Classes

- **Current authority:** owns a live project decision, contract, status, or behavior.
- **Current supporting plan:** expands part of a live work block without replacing the parent.
- **Completed foundation record:** remains authoritative for the completed boundary it records, while later work owns extensions.
- **Historical plan or research:** preserves original reasoning or evidence and must not define current requirements.
- **Execution evidence:** records what was actually run or observed and must not be rewritten to match later expectations.

## Current Planning Authorities

| Document                                                                 | Class                        | Owns                                                                                                                                                     | Must be updated when                                                                                                           |
| ------------------------------------------------------------------------ | ---------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| `skill-issue-project-completion-a-to-b-plan.md`                          | Current authority            | Current A, B, six sequential work blocks, completion criteria, unresolved matters, and project execution status                                          | A block boundary, dependency order, completion criterion, unresolved project decision, or overall status changes               |
| `reorganization-dependency-audit.md`                                     | Current authority            | Completed and partial work inventory, research source map, CLI mismatch record, old-task mapping, and dependency audit                                   | A recorded mismatch is resolved, an artifact changes authority class, research is superseded, or old work gains a new consumer |
| `document-authority-and-update-map.md`                                   | Current authority            | Document roles, semantic ownership, consumer links, and update routing                                                                                   | Any document is added, removed, renamed, superseded, promoted to authority, or assigned a different consumer                   |
| `01-reconcile-the-definitive-product-support-and-evidence-contract.md`   | Completed foundation record  | Selected nine, five-harness minimum tier, thirteen-cell matrix, medium-setting rationale, valid-run meaning, and public-claim boundary                   | The user explicitly revises the support or evaluation commitment                                                               |
| `02-research-and-define-direct-harness-installation-architecture.md`     | Completed foundation record  | Portable payload boundary, adapter classifications, local path contract, lifecycle, ownership, verification stages, and caveats                          | New authoritative evidence or live proof changes a supported path, capability, caveat, or lifecycle rule                       |
| `03-create-the-cli-owned-supporting-skills-bundle.md`                    | Completed foundation record  | Accepted discipline-skill inventory, canonical copied sources, dependency validation, and CLI handoff within Work Block 1                                | The supporting-skill inventory, source path, canonical folder boundary, validation contract, or payload handoff changes        |
| `04-research-cross-harness-subagent-launch.md`                           | Historical foundation record | Delegation evidence for production skills that genuinely need sub-agents; it does not control skill-calling replay                                       | A production-skill delegation requirement or selected harness changes                                                          |
| `05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md` | Current supporting plan      | Three governed scenarios, private expected-call maps, opaque instruction, event meaning, direct primary-agent runner, JSON evidence, and graph semantics | Any evaluation meaning, fixture, event, evidence, validity, or result-schema decision changes                                  |
| `06-establish-the-cross-platform-cli-foundation.md`                      | Completed foundation record  | Existing Go baseline, command routing, platform reporting, embedded-manifest owner, receipt baseline, and cross-build evidence                           | Work Block 2 changes one of these concrete owners or resolves a mismatch recorded by the reorganization audit                  |

## Current Product and Contributor References

| Document or location                                                               | Owns                                                                                                                    | Required synchronization                                                                                                               |
| ---------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `README.md`                                                                        | Public repository overview, current status, local development, and website handoff                                      | Update after each completed work block that changes user-visible capability, status, commands, downloads, results, or public links     |
| `cli/README.md`                                                                    | Current CLI implementation boundary and developer commands                                                              | Update with every CLI command, payload, adapter, lifecycle, receipt, build, or qualification change                                    |
| `skills/skill-intake/references/explicit-invocation-controls.md`                   | Current explicit-only controls and selected-harness local roots                                                         | Update after support-contract or direct-install changes and before releasing altered invocation metadata                               |
| `skills/skill-generation/references/harness-packaging.md`                          | Current generation-time native delivery paths and portable payload boundary                                             | Update after direct-install or canonical-payload changes; historical plugin research must not overwrite it                             |
| `skills/skill-evaluation-and-refinement/references/harness-evaluation-controls.md` | Current description-evaluation gates, independent-agent paths, and native-evidence requirements                         | Update after support, sub-agent, invocation-control, or native-evidence changes                                                        |
| `skills/skill-evaluation-and-refinement/references/campaign-record.md`             | Per-target evaluation campaign record and cleanup structure                                                             | Update only when the evaluator's campaign contract changes in Work Block 1                                                             |
| `evaluation-skills/dictate-plan/`                                                  | Canonical explicit-only Dictate Plan source for evaluation installations                                                | Update when the accepted evaluation entry skill changes                                                                                |
| `evaluations/skill-calling/instrumentation-contract.md`                            | CLI-generated evaluator-copy transformation, opaque marker, private state, turn attribution, cleanup, and event meaning | Update when evaluation instrumentation, replay, state, attribution, cleanup, or marker behavior changes                                |
| `evaluations/skill-calling/event.schema.json`                                      | Portable recorded skill-invocation event schema                                                                         | Update only with an explicit schema-version change and all consuming code and data                                                     |
| `evaluations/skill-calling/scenarios/` and `evaluations/skill-calling/built-ins/`  | Inspectable governed scenario views and embedded scenario-plus-answer-sheet runtime units                               | Update together when a scenario turn, runner rule, expected activation, or built-in identifier changes                                 |
| `src/data/siteData.ts`                                                             | Editable website copy, release metadata, metrics, methodology, and graph data                                           | Update in Work Block 4 from qualified release artifacts and Work Block 3 result data; never use it as the source of evaluation meaning |
| `plans/website/`                                                                   | Completed website design, validation, and visual-audit foundation                                                       | Extend only for finished-product changes; preserve the existing component, data, and style owners                                      |

## Research Authorities and Evidence

| Source                                                          | Role                                                                                        | Consumer rule                                                                                                                                              |
| --------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `plans/deep-research/harness-direct-installation-architecture/` | Current source-backed installation research                                                 | Work Block 2 reads the synthesis and relevant assignment immediately before implementing an adapter; caveated claims require official or live verification |
| `plans/research/cross-platform-cli-distribution/`               | Current distribution research for Go, release artifacts, signing, checksums, and GoReleaser | Work Blocks 2 and 4 verify time-sensitive platform requirements when implementing or publishing                                                            |
| `plans/deep-research/free-website-hosting-cli-downloads/`       | Current MVP hosting and download research                                                   | Work Block 4 verifies current free-tier and deployment behavior when publishing                                                                            |
| `plans/deep-research/codex-friendly-youtube-video-tools/`       | Retained video-tool evidence and risk analysis                                              | Work Block 5 revalidates current tools and competition requirements before production                                                                      |
| `plans/deep-research/natural-ai-tts-for-remotion-demo/`         | Retained Speechify, timing, rights, privacy, and fallback evidence                          | Work Block 5 performs every recorded live account, rights, privacy, and proof-generation gate before use                                                   |
| `evaluations/skill-system-production-refinement/`               | Execution evidence for the qualified twenty-one-iteration Codex production-skill campaign   | Preserve as run evidence; later campaigns may cite its limited finding but must not rewrite it or generalize it across harnesses                           |
| `audits/`                                                       | Completed first-pass audit and semantic walkthrough evidence                                | Preserve as historical evidence; current project completion is owned by the six-block plan                                                                 |

## Historical and Superseded Entry Points

| Document                                                                               | Retained value                                                        | Current replacement                                                                                           |
| -------------------------------------------------------------------------------------- | --------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `plans/skill-generation-and-refinement.md`                                             | Original evaluation → intake → generation intent and dependency order | Six-block completion plan for all remaining work                                                              |
| `plans/skill-generation-and-refinement-first-pass-progress.md`                         | Completed written first-pass execution record                         | Reorganization audit for current inventory and six-block plan for remaining work                              |
| `plans/skill-intake/skill-intake-a-to-b-plan.md`                                       | Original Skill Intake intent                                          | Production Skill Intake files plus Work Block 1 for remaining integration                                     |
| `plans/skill-evaluation-and-refinement/skill-evaluation-and-refinement-a-to-b-plan.md` | Original two-loop evaluator intent                                    | Production evaluator files, completed Codex campaign evidence, and Work Block 1 skill-calling supporting plan |
| `plans/research/agent-harness-popularity/agent-harness-popularity-report.md`           | Popularity evidence and methodology                                   | Completed support contract for the current selected nine and minimum tier                                     |
| `plans/deep-research/agent-harness-plugin-packaging/`                                  | Historical plugin, extension, package, and marketplace evidence       | Direct-install architecture for current implementation paths and lifecycle                                    |

Historical documents keep their original findings. Add a clear status banner at their entry point rather than rewriting their evidence to imitate current decisions.

## Change Routing

### Support or Evaluation Matrix Change

1. Update the completed product, support, and evidence contract only after explicit user approval.
2. Reconcile the six-block parent and reorganization audit.
3. Update the three live skill reference documents.
4. Update Work Blocks 2, 3, and 4 consequences without restating the whole matrix.
5. Update README, website, release, and submission claims only after new evidence exists.

### Harness Path, Capability, or Lifecycle Change

1. Update the direct-install research synthesis or add live proof when the evidence changes.
2. Update the completed direct-install architecture.
3. Reconcile the CLI mismatch record and Work Block 2 implementation requirements.
4. Update affected live skill references, CLI documentation, qualification instructions, support matrix, and public limitations.

### Canonical Skill or Focused Plugin Change

1. Update the canonical source and its Work Block 1 supporting plan.
2. Revalidate transformation points, marker interface, portable behavior, and required references.
3. Rebuild the Work Block 2 embedded payload and evaluator copies.
4. Rerun only the qualification and campaign evidence affected by the semantic change.

### Evaluation Meaning or Result-Schema Change

1. Update the Work Block 1 skill-calling supporting plan and inspectable campaign assets.
2. Update CLI evaluation generation or export behavior in Work Block 2.
3. Invalidate and rerun affected Work Block 3 results rather than silently migrating observations.
4. Update `src/data/siteData.ts`, methods, graphs, video, and submission only from the newly qualified artifacts.

### CLI Implementation Change

1. Update the concrete Go owner and record the affected behavior for the final qualification-stage regression suite.
2. Update `cli/README.md` and the completed baseline record when its documented boundary changes.
3. Reconcile the reorganization audit when a recorded mismatch closes or a new mismatch appears.
4. During Work Block 3, implement and run the accumulated adapter, lifecycle, platform, payload, evaluator, and release qualification coverage against the stabilized candidate.

### Block Completion

1. Record the block's produced artifacts and evidence at their semantic owners.
2. Update the parent plan's current execution status.
3. Update the reorganization audit and this map if an artifact, mismatch, or authority relationship changed.
4. Update README and other user-facing status documents.
5. Search for the superseded task names, harnesses, paths, statuses, and claims across current documents.

## Documentation Completion Gate

Before declaring any work block complete:

- every current requirement has one semantic owner;
- every downstream consumer points to that owner rather than maintaining a competing copy;
- completed historical and evidence documents remain visibly classified;
- changed source paths, task names, harness sets, commands, statuses, and public claims have no stale current references;
- all referenced local documents exist;
- formatting and repository document checks pass.

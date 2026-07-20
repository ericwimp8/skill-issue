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

| Document                                                                 | Class                       | Owns                                                                                                                                                              | Must be updated when                                                                                                           |
| ------------------------------------------------------------------------ | --------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
| `skill-issue-project-completion-a-to-b-plan.md`                          | Current authority           | Current A, B, six sequential work blocks, completion criteria, unresolved matters, and project execution status                                                   | A block boundary, dependency order, completion criterion, unresolved project decision, or overall status changes               |
| `reorganization-dependency-audit.md`                                     | Current authority           | Completed and partial work inventory, research source map, CLI mismatch record, old-task mapping, and dependency audit                                            | A recorded mismatch is resolved, an artifact changes authority class, research is superseded, or old work gains a new consumer |
| `document-authority-and-update-map.md`                                   | Current authority           | Document roles, semantic ownership, consumer links, and update routing                                                                                            | Any document is added, removed, renamed, superseded, promoted to authority, or assigned a different consumer                   |
| `01-reconcile-the-definitive-product-support-and-evidence-contract.md`   | Completed foundation record | Six selected harnesses, five-harness minimum tier, thirteen-cell matrix, medium-setting rationale, valid-run meaning, and public-claim boundary                   | The user explicitly revises the support or evaluation commitment                                                               |
| `02-research-and-define-direct-harness-installation-architecture.md`     | Completed foundation record | Portable payload boundary, adapter classifications, local path contract, lifecycle, ownership, installation-versus-qualification boundary, and support boundaries | New authoritative evidence or live proof changes a supported path, capability, support boundary, or lifecycle rule             |
| `03-create-the-cli-owned-supporting-skills-bundle.md`                    | Completed foundation record | Accepted discipline-skill inventory, canonical copied sources, dependency validation, and CLI handoff within Work Block 1                                         | The supporting-skill inventory, source path, canonical folder boundary, validation contract, or payload handoff changes        |
| `05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md` | Current supporting plan     | Three governed scenarios, private expected-call maps, opaque instruction, event meaning, direct primary-agent runner, JSON evidence, and graph semantics          | Any evaluation meaning, fixture, event, evidence, validity, or result-schema decision changes                                  |
| `06-establish-the-cross-platform-cli-foundation.md`                      | Completed foundation record | Original Go baseline, command routing, platform reporting, embedded-manifest owner, historical receipt baseline, and cross-build evidence                         | Work Block 2 changes one of these concrete owners or resolves a mismatch recorded by the reorganization audit                  |
| `../skill-calling-evaluation-campaign/evaluation-progress.md`            | Current supporting plan     | Subscription-bounded ten-configuration campaign progress, scenario checklists, attempts, failures, and blockers                                                   | A selected configuration, reasoning target, evaluation status, attempt, failure, blocker, or campaign date changes             |

## Current Product and Contributor References

| Document or location                                                              | Owns                                                                                                                    | Required synchronization                                                                                                                          |
| --------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `README.md`                                                                       | Public repository overview, current status, local development, and website handoff                                      | Update after each completed work block that changes user-visible capability, status, commands, downloads, results, or public links                |
| `cli/README.md`                                                                   | Current CLI implementation boundary and developer commands                                                              | Update with every CLI command, payload, adapter, lifecycle, evaluation-state, build, or qualification change                                      |
| `plans/harness-setup.md`                                                          | Native evaluation setup contracts for Codex, Cursor, Claude Code, and Pi                                                | Update when a native invocation, isolation, supplied-skill, model, reasoning, session, authentication, permission, or cleanup requirement changes |
| `skills/skill-evaluation-and-refinement/references/campaign-record.md`            | Per-target evaluation campaign record and cleanup structure                                                             | Update only when the evaluator's campaign contract changes in Work Block 1                                                                        |
| `evaluations/skill-calling/instrumentation-contract.md`                           | CLI-generated evaluator-copy transformation, opaque marker, private state, turn attribution, cleanup, and event meaning | Update when evaluation instrumentation, replay, state, attribution, cleanup, or marker behavior changes                                           |
| `evaluations/skill-calling/event.schema.json`                                     | Portable recorded skill-invocation event schema                                                                         | Update only with an explicit schema-version change and all consuming code and data                                                                |
| `evaluations/skill-calling/scenarios/` and `evaluations/skill-calling/built-ins/` | Inspectable governed scenario views and embedded scenario-plus-answer-sheet runtime units                               | Update together when a scenario turn, runner rule, expected activation, or built-in identifier changes                                            |
| `src/data/siteData.ts`                                                            | Curated website copy, release metadata, summary metrics, and methodology text                                           | Update when approved public copy, release destinations, summary metrics, or methodology wording changes                                           |
| `src/data/evaluationData.ts`                                                      | Website artifact types, adaptation, display identities, selection inputs, and temporary illustrative result data        | Replace illustrative data only from accepted Work Block 3 artifacts; never use it as the source of evaluation meaning or campaign acceptance      |
| `src/components/EvaluationExplorer.tsx` and `src/components/charts/`              | Evaluation selection, filtering, shared chart helpers, and concrete chart presentations                                 | Update when the website interaction model or chart presentation changes; preserve upstream evaluation meaning and accepted result identity        |
| `src/styles.css`                                                                  | Shared website theme, layout, responsive behavior, and visual tokens                                                    | Update when a shared visual decision changes rather than duplicating it in a component                                                            |
| `plans/website/reference-and-architecture.md`                                     | Current website hosting, stack, visual direction, and source-ownership foundation                                       | Update when the hosting boundary, website stack, visual direction, or source ownership changes                                                    |

## Research Authorities and Evidence

| Source                                                             | Role                                                                                        | Consumer rule                                                                                                                                          |
| ------------------------------------------------------------------ | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `research/deep-research/harness-direct-installation-architecture/` | Retained source evidence for the integrated direct-install architecture                     | Work Block 2 uses the evidence for concrete adapter work; the completed architecture plan and CLI source own current lifecycle and ownership decisions |
| `research/cross-platform-cli-distribution/`                        | Current distribution research for Go, release artifacts, signing, checksums, and GoReleaser | Work Blocks 2 and 4 verify time-sensitive platform requirements when implementing or publishing                                                        |
| `research/harness-subagent-launch/`                                | Source-backed delegation capability evidence                                                | Environment qualification may consume it when establishing isolated trial methods; skill-calling replay remains a primary-agent concern                |
| `research/deep-research/manual-skill-invocation-nine-harnesses/`   | Source-backed manual and explicit invocation evidence                                       | CLI adapters and harness qualification guidance consume it when invocation controls or user-facing invocation guidance change                          |
| `research/deep-research/free-website-hosting-cli-downloads/`       | Current MVP hosting and download research                                                   | Work Block 4 verifies current free-tier and deployment behavior when publishing                                                                        |
| `research/implementation-research/website-chart-migration/`        | Source and design evidence for the implemented chart migration                              | Work Blocks 3 and 4 use it when replacing illustrative data with accepted campaign artifacts                                                           |
| `research/deep-research/codex-friendly-youtube-video-tools/`       | Retained video-tool evidence and risk analysis                                              | Work Block 5 revalidates current tools and competition requirements before production                                                                  |
| `research/deep-research/natural-ai-tts-for-remotion-demo/`         | Retained Speechify, rights, privacy, and fallback evidence                                  | Work Block 5 performs every recorded live account, rights, privacy, and proof-generation gate before use                                               |
| `research/audits/skill-generation-and-refinement/`                 | Completed first-pass audit and semantic walkthrough evidence                                | Preserve as historical evidence; current production behavior is owned by the skill sources and current completion plan                                 |
| `evaluations/skill-system-production-refinement/`                  | Execution evidence for the qualified twenty-one-iteration Codex production-skill campaign   | Preserve as run evidence; later campaigns may cite its limited finding but must not rewrite it or generalize it across harnesses                       |
| `evaluations/skill-calling/smoke/real-harness-smoke-report.md`     | Bounded Codex, Cursor, Claude Code, and Pi launcher smoke evidence                          | Preserve as execution evidence; it does not establish governed campaign, complete workflow, OpenCode, native-platform, or release qualification        |

## Change Routing

### Support or Evaluation Matrix Change

1. Update the completed product, support, and evidence contract only after explicit user approval.
2. Reconcile the six-block parent and reorganization audit.
3. Update Work Blocks 2, 3, and 4 consequences without restating the whole matrix.
4. Update README, website, release, and submission claims only after new evidence exists.

### Harness Path, Capability, or Lifecycle Change

1. Update the direct-install research synthesis or add live proof when the evidence changes.
2. Update the completed direct-install architecture.
3. Reconcile the CLI mismatch record and Work Block 2 implementation requirements.
4. Update affected CLI adapters, CLI documentation, qualification instructions, support matrix, and public limitations.

### Canonical Skill or Focused Plugin Change

1. Update the canonical source and its Work Block 1 supporting plan.
2. Revalidate transformation points, marker interface, portable behavior, and required references.
3. Rebuild the Work Block 2 embedded payload and evaluator copies.
4. Rerun only the qualification and campaign evidence affected by the semantic change.

### Evaluation Meaning or Result-Schema Change

1. Update the Work Block 1 skill-calling supporting plan and inspectable campaign assets.
2. Update CLI evaluation generation or export behavior in Work Block 2.
3. Invalidate and rerun affected Work Block 3 results rather than silently migrating observations.
4. Update the accepted website result projection, curated `src/data/siteData.ts` copy, methods, graphs, video, and submission only from the newly qualified artifacts.

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

# Skill Issue Project Reorganization and Dependency Audit

## Purpose

Account for completed work, partial implementations, retained research, open requirements, and dependency handoffs before replacing the fragmented twenty-one-task parent plan with six sequential work blocks.

This audit changes project sequencing and ownership only. It does not implement the remaining skill system, CLI adapters, evaluation campaign, public release, website deployment, or submission.

Document authority classes, consumer links, and future update routing are maintained in `plans/skill-issue-project-completion/document-authority-and-update-map.md`.

## Confirmed Completed Foundations

### Core Skill First Pass

- `skills/skill-intake/`, `skills/skill-generation/`, and `skills/skill-evaluation-and-refinement/` contain the three production workflow skills and their supporting references.
- `plans/skill-generation-and-refinement-first-pass-progress.md` records the completed written first pass, bounded requirements audit, and two outside-in semantic walkthroughs.
- `audits/skill-generation-and-refinement-first-pass-audit.md` and `audits/skill-generation-and-refinement-semantic-walkthroughs.md` preserve the first-pass evidence and accepted runtime deferrals.
- The first-pass Codex plugin wrapper and explicit-only overlay exist in `.codex-plugin/`, `.codex/`, and `packaging/overlays/`.

### Production Refinement Campaign

- `evaluations/skill-system-production-refinement/` records the completed twenty-one-iteration evaluator, Skill Intake, and Skill Generation refinement campaign.
- `evaluations/skill-system-production-refinement/environment-qualification.md` qualifies only the recorded Codex Desktop and `gpt-5.6-sol` high-reasoning campaign surface.
- `evaluations/skill-system-production-refinement/final-audit.md` records a pass after reconciliation for that bounded campaign.
- The campaign proves the three production skills within its recorded Codex environment and reaches an explicit generated-skill Evaluation-handoff stop. It does not prove the future thirteen-cell medium-setting skill-calling campaign, every selected harness, or runtime success for an arbitrary generated skill.

### Product, Support, and Evidence Contract

- `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md` owns the selected-nine boundary, five-harness minimum release tier, thirteen medium-setting evaluation cells, one three-scenario-suite minimum, default-environment expectation, and transparent public-claim rules.
- Its section audit is complete.

### Direct-Installation Architecture

- `plans/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md` and its ten assignment reports are the authoritative installation research.
- `plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md` integrates the portable payload, adapter classifications, receipt-backed lifecycle, materialized/discovered/activated distinction, trust boundaries, and fail-closed caveats.
- This audit confirms that Task 2's source coverage, downstream consequences, and explicit unresolved claims are present. Its section audit is complete.

### CLI Foundation

- `cli/` contains a pure-Go executable foundation with platform reporting, embedded manifest loading, lifecycle routing, atomic receipts, and six cross-compiled target binaries. Its transitional automated tests have been removed; final regression coverage is owned by Work Block 3 after the CLI interfaces stabilize.
- `plans/skill-issue-project-completion/06-establish-the-cross-platform-cli-foundation.md` and `cli/README.md` document that foundation.
- Work Block 2 has extended the baseline with the canonical payload, concrete lifecycle manager, researched native roots, blind evaluation runner, and ownership-backed cleanup. Real-harness and native-platform qualification remain open.

### Website Mock-Up

- The React, TypeScript, Vite, and Recharts website is implemented under `src/` with local data ownership in `src/data/siteData.ts`.
- `plans/website/website-a-to-b-task-list.md` records the completed mock-up, browser inspection, image-comparison refinement, deterministic validation, and GitHub Pages build proof.
- The website still contains mock benchmark and release content. Public deployment, real result ingestion, approved final copy, and working binary downloads remain unfinished.

## Retained Research Source Map

| Source                                                                                     | Current authority                                                                                                       | Required future consumer                                                                                                          |
| ------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| `plans/research/agent-harness-popularity/agent-harness-popularity-report.md`               | Historical popularity evidence. Its JetBrains-era selected list is superseded by the project support contract.          | Work Block 1 uses the evidence only after applying the current selected-nine contract.                                            |
| `plans/deep-research/agent-harness-plugin-packaging/`                                      | Historical packaging and ecosystem context. It does not own current direct-install paths.                               | Work Block 1 may use feature and portability context; Work Block 2 must use the newer direct-install research for implementation. |
| `plans/deep-research/harness-direct-installation-architecture/`                            | Authoritative local direct-install architecture, paths, adapter classifications, lifecycle, verification, and caveats.  | Work Block 2 implements it; Work Blocks 3 and 4 consume its evidence and public support distinctions.                             |
| `plans/research/cross-platform-cli-distribution/cross-platform-cli-distribution-report.md` | Authoritative Go, prebuilt-binary, GitHub Releases, signing, checksums, GoReleaser, and native qualification direction. | Work Block 2 completes the CLI; Work Block 4 produces public release artifacts.                                                   |
| `plans/deep-research/free-website-hosting-cli-downloads/`                                  | Authoritative GitHub Pages and GitHub Releases hosting direction for the free MVP.                                      | Work Block 4 publishes the binaries and website.                                                                                  |
| `plans/deep-research/codex-friendly-youtube-video-tools/`                                  | Retained video-production research and risk record.                                                                     | Work Block 5 rechecks current requirements and uses only still-valid findings immediately before production.                      |
| `plans/deep-research/natural-ai-tts-for-remotion-demo/`                                    | Retained Speechify, rights, timing, caption, privacy, and fallback research.                                            | Work Block 5 performs the live account, rights, proof-generation, and publication gates before generating final narration.        |
| `plans/website/`                                                                           | Authoritative mock-up design, data, validation, and visual-audit record.                                                | Work Block 4 updates the existing owners rather than rebuilding the website.                                                      |
| `evaluations/skill-system-production-refinement/`                                          | Evidence for the qualified Codex production-skill campaign only.                                                        | Work Block 1 uses it as the skill baseline; Work Block 3 preserves its limited claim when designing broader qualification.        |

## Dependency and Ownership Findings

### 1. Current Harness References Are Reconciled

- `skills/skill-evaluation-and-refinement/references/harness-evaluation-controls.md`, `skills/skill-intake/references/explicit-invocation-controls.md`, and `skills/skill-generation/references/harness-packaging.md` now use the current selected nine and authoritative direct-install findings.
- Those live references link to their support and installation authorities and state their update triggers.
- Historical JetBrains-era and plugin-packaging reports remain evidence-only records and carry supersession notices directing readers to the current contracts.
- Work Block 1 preserves this reconciliation and resolves only evidence gaps required by the skill currently being finalized.

### 2. The CLI Foundation Does Not Yet Match the Final Installation Contract

- `cli/internal/receipt/receipt.go` models `project` and `system`, while the final architecture requires exact surface-specific project and user scopes and explicit absence or caveating of admin or remote scopes.
- The receipt currently lacks the finalized adapter version, host version, canonical identity, source and manifest digests, executable inventory, backup identity, reversible configuration fragment, and separate materialization, discovery, and activation results.
- `cli/internal/lifecycle/lifecycle.go` and command routing omit or collapse finalized capabilities including detection, scope enumeration, preview, activation, rollback, and three-stage verification.
- `cli/internal/payload/assets/manifest.json` intentionally contains no components, and no adapter can install, discover, activate, verify, repair, update, roll back, or remove a harness payload.
- Work Block 2 owns reconciling and evolving these concrete foundation owners before implementing payloads and adapters. This is remaining CLI work, not a Task 3 responsibility.

### 3. The CLI-Owned Supporting-Skills Bundle Exists

- The current `.codex-plugin/` is a historical wrapper around the three primary Skill Issue workflow skills and does not define the release installation architecture.
- `supporting-skills/` contains verbatim validated copies of the four accepted discipline skills.
- `plans/skill-issue-project-completion/03-create-the-cli-owned-supporting-skills-bundle.md` records the inventory, source path, validation, and CLI handoff.

### 4. Two Different Evaluation Systems Must Remain Distinct

- The completed production-refinement campaign evaluates and improves skill descriptions and bodies in one qualified Codex environment.
- The planned skill-calling campaign measures whether selected harness-and-model environments invoke supplied skills at predetermined points in three governed 30-turn scenarios.
- `evaluations/skill-calling/instrumentation-contract.md` and `event.schema.json` define the opaque instruction, private token and turn state, marker behavior, cleanup, and invocation event.
- The gardening scenario, its separately supplied answer sheet, and the evaluation-only Dictate Plan source are complete.
- The direct primary-agent Go runner and result derivation are implemented. The remaining two scenarios, real-harness qualification, campaign aggregation, and thirteen-cell evidence remain unfinished.
- Work Block 1 completes those remaining reusable assets. Work Block 3 executes qualification and the campaign.

### 5. Integrated Product Handoffs Are Complete but Runtime Proof Remains

- Skill Intake, Skill Generation, and Skill Evaluation and Refinement have explicit written handoffs, same-task continuation when authority permits, and qualified Codex campaign evidence.
- The retained end-to-end campaign stops at the generated skill's Evaluation handoff and accurately defers that generated skill's runtime evaluation.
- The remaining replay and later qualification work must prove the advertised ordinary-language intake-to-generation-to-evaluation-to-refinement route while preserving the standalone existing-skill evaluation entry point.

### 6. Evaluation Installation Must Be Derived, Not Forked

- Canonical skill sources must remain free of evaluation-only instructions.
- The primary and supporting skill directories supply canonical sources; the evaluation contract owns the injected instruction and event meaning; the CLI generates disposable owned evaluator copies from those current sources.
- Work Block 1 has fixed the canonical and event contracts. Work Block 2 implements generation, installation, refresh, verification, and removal. Work Block 3 consumes the resulting mode.

### 7. Release, Website, and Video Foundations Are Uneven

- The website mock-up and its deterministic checks are complete, but final copy, real evaluation data, direct release downloads, and deployment are unfinished.
- The CLI research recommends GoReleaser, signing, checksums, native smoke tests, and GitHub Releases, but the repository does not yet contain the final release automation or qualified public artifacts.
- Video and narration research exists, but no Remotion production project, approved narration proof, final video, public YouTube URL, or submission package exists.
- Work Block 4 owns repository release, binaries, narrative, real website content, and deployment. Work Block 5 owns current-rule verification, video production, and submission.

### 8. External Evaluation Depends on the Finished Installer

- External evaluators cannot be recruited effectively until the integrated CLI can install the normal and evaluation payloads, run the governed evaluation, export evidence, and clean up.
- Work Block 3 begins only after Work Block 2 produces that candidate. It performs an accessible-environment smoke test before recruitment, then supports external and internal runs under one evidence contract.

## Old Task to New Work Block Mapping

| Previous task                                     | New owner                                                        |
| ------------------------------------------------- | ---------------------------------------------------------------- |
| 1. Product, support, and evidence contract        | Completed foundation consumed by Work Blocks 1, 2, 3, and 4      |
| 2. Direct harness installation architecture       | Completed foundation consumed by Work Blocks 2, 3, and 4         |
| 3. CLI-owned supporting-skills bundle             | Work Block 1                                                     |
| 4. Historical cross-harness delegation foundation | Retained for production skills; outside the skill-calling runner |
| 5. Skill-calling evaluation contract and assets   | Work Block 1                                                     |
| 6. Cross-platform CLI foundation                  | Implemented baseline evolved and completed by Work Block 2       |
| 7. Public narrative and copy                      | Work Block 4                                                     |
| 8. Remotion and Speechify foundation              | Work Block 5                                                     |
| 9. Public repository and release foundation       | Work Block 4                                                     |
| 10. Integrated natural-language skill creation    | Work Block 1                                                     |
| 11. Repeatable skill-calling evaluation system    | Work Block 1                                                     |
| 12. Canonical payload and direct-install targets  | Work Block 2                                                     |
| 13. Installer CLI and harness adapters            | Work Block 2                                                     |
| 14. External-evaluator-ready milestone            | Work Block 3                                                     |
| 15. Recruit and support evaluators                | Work Block 3                                                     |
| 16. Qualify installation and workflow             | Work Block 3                                                     |
| 17. Cross-harness evaluation campaign             | Work Block 3                                                     |
| 18. Repository and release artifacts              | Work Block 4                                                     |
| 19. Website completion and deployment             | Work Block 4                                                     |
| 20. Video and hackathon submission                | Work Block 5                                                     |
| 21. Final release and submission audit            | Work Block 6                                                     |

## Sequential Gate Contract

1. Work Block 1 finishes the canonical skill system and evaluation assets.
2. Work Block 2 consumes that fixed handoff, reconciles the existing CLI foundation, and produces the complete installer product.
3. Work Block 3 consumes the installable evaluation product and produces qualification and campaign evidence.
4. Work Block 4 consumes qualified binaries and final results and publishes the repository, release, narrative, and website.
5. Work Block 5 consumes the finished public product and produces the video and submission.
6. Work Block 6 independently audits the entire public journey and closes only evidence-supported failures.

Research is read and any remaining authoritative gap is resolved inside the work block that immediately consumes the answer. Parallel execution is exceptional: it requires fixed upstream interfaces, disjoint file ownership, no unresolved shared decision, and an explicit reconciliation point before either lane's output is consumed.

## Completion-Criteria Preservation

| Previous completion meaning                                                            | Preserved owner                                                               |
| -------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| Runtime-free macOS, Windows, and Linux CLI                                             | Work Block 2 implementation and Work Blocks 3 and 4 qualification and release |
| Harness and scope detection, confirmation, installation records, and safe lifecycle    | Work Block 2                                                                  |
| Complete payload with primary workflows and the supporting-skills bundle               | Work Blocks 1 and 2                                                           |
| Ordinary-language skill creation through evaluation and refinement                     | Work Block 1, proven in Work Block 3                                          |
| Existing-skill evaluation and semantic refinement                                      | Work Block 1, proven in Work Block 3                                          |
| Harness-agnostic independent-agent behavior and capability gates                       | Work Block 1, proven in Work Block 3                                          |
| Five-harness complete workflow proof and thirteen-cell medium-setting campaign         | Work Block 3                                                                  |
| External evaluator installation, execution, evidence return, and attribution consent   | Work Block 3                                                                  |
| Adjacent-skill and configuration disclosure                                            | Work Block 3 evidence and Work Block 4 public documentation                   |
| Qualified release binaries, integrity metadata, and working downloads                  | Work Block 4                                                                  |
| Real deployed website results, limitations, hooks, About content, and repository route | Work Block 4                                                                  |
| Public video, current-rule compliance, and completed hackathon submission              | Work Block 5                                                                  |
| Fresh-user journey and final evidence-backed completion decision                       | Work Block 6                                                                  |

## Audit Decision

The six-block structure preserves every sane requirement from the twenty-one-task plan and gives every completed, partial, deferred, or researched artifact a named owner and consumer. The rewritten parent plan adopts this audit and the documentation authority map as its controlling source and update-routing records.

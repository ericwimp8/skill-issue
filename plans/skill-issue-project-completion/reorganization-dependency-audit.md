# Skill Issue Project Reorganization and Dependency Audit

## Purpose

Account for completed work, partial implementations, retained research, open requirements, and dependency handoffs before replacing the fragmented twenty-one-task parent plan with six sequential work blocks.

This audit changes project sequencing and ownership only. It does not implement the remaining skill system, CLI adapters, evaluation campaign, public release, website deployment, or submission.

Document authority classes, consumer links, and future update routing are maintained in `plans/skill-issue-project-completion/document-authority-and-update-map.md`.

## Confirmed Completed Foundations

### Core Skill First Pass

- `skills/skill-intake/`, `skills/skill-generation/`, and `skills/skill-evaluation-and-refinement/` contain the three production workflow skills and their supporting references.
- `research/audits/skill-generation-and-refinement/` preserves the first-pass audit evidence, semantic walkthroughs, and accepted runtime deferrals.
- The canonical `skill-intake` description carries its explicit-request boundary. Harness-specific enforcement is owned by CLI materialization rather than a repository-only overlay.

### Production Refinement Campaign

- `evaluations/skill-system-production-refinement/` records the completed twenty-one-iteration evaluator, Skill Intake, and Skill Generation refinement campaign.
- `evaluations/skill-system-production-refinement/environment-qualification.md` qualifies only the recorded Codex Desktop and `gpt-5.6-sol` high-reasoning campaign surface.
- `evaluations/skill-system-production-refinement/final-audit.md` records a pass after reconciliation for that bounded campaign.
- The campaign proves the three production skills within its recorded Codex environment and reaches an explicit generated-skill Evaluation-handoff stop. It does not prove the future thirteen-cell medium-setting skill-calling campaign, every selected harness, or runtime success for an arbitrary generated skill.

### Product, Support, and Evidence Contract

- `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md` owns the six-harness boundary, five-harness minimum release tier, thirteen medium-setting evaluation cells, one three-scenario-suite minimum, default-environment expectation, and transparent public-claim rules.
- Its section audit is complete.

### Direct-Installation Architecture

- `research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md` and its ten assignment reports preserve the installation source evidence. The integrated architecture plan and CLI source own current lifecycle and ownership decisions.
- `plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md` integrates the portable payload, adapter classifications, direct disposable lifecycle, installation-versus-qualification boundary, trust preservation, and fail-closed or unsupported surfaces.
- This audit confirms that Task 2's source coverage, downstream consequences, and support boundaries are present. Its section audit is complete.

### CLI Foundation

- `cli/` contains a pure-Go executable foundation with platform reporting, embedded payload loading, direct install and uninstall routing, evaluation replay and cleanup, a six-target cross-build script, and six locally built binaries. Codex, Cursor, Claude Code, and Pi now have configurable native evaluation routes, custom supplied-skill input, effective-configuration confirmation, bounded cleanup, and completed two-turn real-launcher smoke checks.
- `plans/skill-issue-project-completion/06-establish-the-cross-platform-cli-foundation.md` and `cli/README.md` document that foundation.
- `evaluations/skill-calling/smoke/real-harness-smoke-report.md` preserves the bounded runtime evidence for the completed four-harness runner implementation.
- Work Block 2 has extended the baseline with the canonical payload, concrete lifecycle manager, retained native roots, blind evaluation runner, four native launcher routes, guided selection and preview, file verification, and ownership-backed cleanup. OpenCode and Kilo Code implementation, concrete installation-adapter preflight, final packaging, and the block audit remain open. Governed-campaign, native-platform, and release qualification remain downstream.

### Website Mock-Up

- The React, TypeScript, Vite, and Recharts website is implemented under `src/`. `src/data/siteData.ts` owns curated copy, release metadata, and summary metrics; `src/data/evaluationData.ts` owns website artifact adaptation, display identities, and temporary illustrative results; the explorer and chart components own interaction and presentation.
- `plans/website/reference-and-architecture.md` preserves the current hosting, stack, visual-direction, and source-ownership decisions. Production source and deterministic validation own the implemented website behavior.
- The website still contains illustrative benchmark data and provisional release content. Accepted campaign-result export, public deployment, approved final copy, and working binary downloads remain unfinished.

## Retained Research Source Map

| Source                                                                               | Current authority                                                                                                       | Required future consumer                                                                                                    |
| ------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| `research/deep-research/harness-direct-installation-architecture/`                   | Retained source evidence for direct-install paths, adapter classifications, verification, and boundaries.               | Work Block 2 consumes it through the integrated architecture plan; Work Blocks 3 and 4 use its public support distinctions. |
| `research/cross-platform-cli-distribution/cross-platform-cli-distribution-report.md` | Authoritative Go, prebuilt-binary, GitHub Releases, signing, checksums, GoReleaser, and native qualification direction. | Work Block 2 completes the CLI; Work Block 4 produces public release artifacts.                                             |
| `research/harness-subagent-launch/`                                                  | Source-backed delegation capability evidence for the selected harnesses.                                                | Environment qualification consumes it when establishing isolated trial methods.                                             |
| `research/deep-research/manual-skill-invocation-nine-harnesses/`                     | Source-backed manual and explicit invocation evidence for the selected harnesses.                                       | CLI adapters and harness qualification guidance consume it when invocation behavior changes.                                |
| `research/deep-research/free-website-hosting-cli-downloads/`                         | Authoritative GitHub Pages and GitHub Releases hosting direction for the free MVP.                                      | Work Block 4 publishes the binaries and website.                                                                            |
| `research/implementation-research/website-chart-migration/`                          | Source and design evidence for the implemented chart migration and accepted-data handoff.                               | Work Blocks 3 and 4 use it when replacing illustrative data with accepted campaign artifacts.                               |
| `research/deep-research/codex-friendly-youtube-video-tools/`                         | Retained video-production research and risk record.                                                                     | Work Block 5 rechecks current requirements and uses only still-valid findings immediately before production.                |
| `research/deep-research/natural-ai-tts-for-remotion-demo/`                           | Retained Speechify, rights, timing, caption, privacy, and fallback research.                                            | Work Block 5 performs the live account, rights, proof-generation, and publication gates before generating final narration.  |
| `plans/website/reference-and-architecture.md`                                        | Current hosting, stack, visual-direction, and source-ownership foundation.                                              | Work Block 4 updates the existing source owners rather than rebuilding the website.                                         |
| `evaluations/skill-system-production-refinement/`                                    | Evidence for the qualified Codex production-skill campaign only.                                                        | Work Block 1 uses it as the skill baseline; Work Block 3 preserves its limited claim when designing broader qualification.  |

## Dependency and Ownership Findings

### 1. Harness Ownership Is Reconciled

- Harness installation paths, invocation controls, and qualification setup are owned by the CLI adapters and harness documentation rather than bundled production-skill references.
- Production skills retain only task-owned references that are complete inside their installed skill directories.
- The product support contract, direct-install architecture, and retained research preserve the repository-only decision and evidence chain.
- Superseded popularity and plugin-packaging research has been removed after its current decisions were reconciled into the support contract and direct-install architecture.
- Later support changes follow the CLI, qualification, and public-document routes in the documentation authority map.

### 2. The Installer Exists but Its User Journey and Lifecycle Semantics Remain Incomplete

- `cli/internal/harness/harness.go` contains the six retained project and user roots. Guided installation presents explicit harness and scope selection, marks OpenCode and Kilo Code as in progress, previews the destination and payload, and requires confirmation.
- `cli/internal/payload/assets/manifest.json` contains the complete canonical component inventory, and the payload owner validates canonical frontmatter and referenced-file closure.
- `cli/internal/installer/installer.go` stages and replaces only the known Skill Issue payload directories. It stores no ordinary receipt, digest, backup, rollback, or mutable application state. Repeated installation is the reinstall, update, and repair path; uninstallation removes those same known directories.
- `cli/internal/lifecycle/lifecycle.go` exposes ordinary `install` and `uninstall` plus evaluation run and cleanup. The former verify, repair, and update command paths are removed.
- Selection and preview consume the same harness roots and embedded payload as installation. Generic adapter metadata, persistent rollback, and installation-time discovery or activation results are not required.
- Work Block 3 owns live harness discovery, activation, final regression, and native qualification after Work Block 2 freezes the installer candidate.

### 3. The CLI-Owned Supporting-Skills Bundle Exists

- The current `.codex-plugin/` is a historical wrapper around the three primary Skill Issue workflow skills and does not define the release installation architecture.
- `supporting-skills/` contains verbatim validated copies of the four accepted discipline skills.
- `plans/skill-issue-project-completion/03-create-the-cli-owned-supporting-skills-bundle.md` records the inventory, source path, validation, and CLI handoff.

### 4. Two Different Evaluation Systems Must Remain Distinct

- The completed production-refinement campaign evaluates and improves skill descriptions and bodies in one qualified Codex environment.
- The planned skill-calling campaign measures whether selected harness-and-model environments invoke supplied skills at predetermined points in three governed 30-turn scenarios.
- `evaluations/skill-calling/instrumentation-contract.md` and `event.schema.json` define the opaque instruction, private token and turn state, marker behavior, cleanup, and invocation event.
- All three governed scenarios and their embedded private answer sheets are complete and measure four canonical Skill Issue activations each. Paired external skill, scenario, and answer-sheet inputs remain supported for custom evaluations.
- The direct primary-agent Go runner, all three governed scenarios, and result derivation are implemented. Codex, Cursor, Claude Code, and Pi completed bounded built-in and custom two-turn launcher smoke routes. Complete workflow qualification, OpenCode support, governed campaign execution, campaign aggregation, and thirteen-cell evidence remain unfinished.
- Work Block 3 executes qualification and the campaign.

### 5. Integrated Product Handoffs Are Complete but Runtime Proof Remains

- Skill Intake, Skill Generation, and Skill Evaluation and Refinement have explicit written handoffs, same-task continuation when authority permits, and qualified Codex campaign evidence.
- The retained end-to-end campaign stops at the generated skill's Evaluation handoff and accurately defers that generated skill's runtime evaluation.
- Later complete-workflow qualification must prove the advertised ordinary-language intake-to-generation-to-evaluation-to-refinement route while preserving the standalone existing-skill evaluation entry point. The completed launcher smoke routes establish transport, supplied-skill loading, evidence output, and cleanup only within their recorded scope.

### 6. Evaluation Installation Must Be Derived, Not Forked

- Canonical skill sources must remain free of evaluation-only instructions.
- The primary and supporting skill directories supply canonical sources; the evaluation contract owns the injected instruction and event meaning; the CLI generates disposable owned evaluator copies from those current sources.
- Work Block 1 fixed the canonical and event contracts. Work Block 2 has implemented evaluator-copy generation, project-local installation, canonical rematerialization, temporary-path removal, and interrupted-run cleanup. Work Block 3 consumes and qualifies the resulting mode.

### 7. Release, Website, and Video Foundations Are Uneven

- The website mock-up and its deterministic checks are complete, but final copy, real evaluation data, direct release downloads, and deployment are unfinished.
- The public `ericwimp8/skill-issue` repository exists. It has no release, Pages deployment, workflow, license, contribution guide, or security policy.
- The six-target cross-build script and local binaries exist. The repository does not yet contain GoReleaser automation, checksums, native smoke proof, signing decisions, or qualified public artifacts.
- Video and narration research exists, but no Remotion production project, approved narration proof, final video, public YouTube URL, or submission package exists.
- Work Block 4 owns repository release, binaries, narrative, real website content, and deployment. Work Block 5 owns current-rule verification, video production, and submission.

### 8. External Evaluation Depends on the Finished Installer

- External evaluators cannot be recruited effectively until the integrated CLI can install the normal and evaluation payloads, run the governed evaluation, export evidence, and clean up through the finished ordinary product journey.
- The four-harness two-turn smoke campaign proves the bounded launcher interfaces but does not replace the accessible-environment complete-workflow gate. Work Block 3 begins after Work Block 2 produces the finished candidate, performs that gate before recruitment, then supports external and internal runs under one evidence contract.

## Old Task to New Work Block Mapping

| Previous task                                     | New owner                                                          |
| ------------------------------------------------- | ------------------------------------------------------------------ |
| 1. Product, support, and evidence contract        | Completed foundation consumed by Work Blocks 1, 2, 3, and 4        |
| 2. Direct harness installation architecture       | Completed foundation consumed by Work Blocks 2, 3, and 4           |
| 3. CLI-owned supporting-skills bundle             | Work Block 1                                                       |
| 4. Historical cross-harness delegation foundation | Retained for production skills; outside the skill-calling runner   |
| 5. Skill-calling evaluation contract and assets   | Work Block 1                                                       |
| 6. Cross-platform CLI foundation                  | Completed baseline; remaining product work belongs to Work Block 2 |
| 7. Public narrative and copy                      | Work Block 4                                                       |
| 8. Remotion and Speechify foundation              | Work Block 5                                                       |
| 9. Public repository and release foundation       | Work Block 4                                                       |
| 10. Integrated natural-language skill creation    | Work Block 1                                                       |
| 11. Repeatable skill-calling evaluation system    | Work Block 1                                                       |
| 12. Canonical payload and direct-install targets  | Work Block 2                                                       |
| 13. Installer CLI and harness adapters            | Work Block 2                                                       |
| 14. External-evaluator-ready milestone            | Work Block 3                                                       |
| 15. Recruit and support evaluators                | Work Block 3                                                       |
| 16. Qualify installation and workflow             | Work Block 3                                                       |
| 17. Cross-harness evaluation campaign             | Work Block 3                                                       |
| 18. Repository and release artifacts              | Work Block 4                                                       |
| 19. Website completion and deployment             | Work Block 4                                                       |
| 20. Video and hackathon submission                | Work Block 5                                                       |
| 21. Final release and submission audit            | Work Block 6                                                       |

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
| Harness and scope selection, confirmation, installation verification, and safe lifecycle | Work Block 2                                                               |
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

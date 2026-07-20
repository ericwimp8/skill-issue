# Product Support and Harness Setup Contract

## Assignment

- **Goal:** Establish the controlling product, support, and setup boundary for harness qualification and portability.
- **Scope:** Current planning authorities, CLI production behavior, native setup contracts, retained smoke evidence, and the older direct-install research needed to identify superseded lifecycle proposals.
- **Exclusions:** Per-harness candidate audits, new adapter design, source changes, and claims beyond the recorded local qualification evidence.

## Sources

- `plans/skill-issue-project-completion/document-authority-and-update-map.md:12-18,24-44,57-67,79-105` — authority classes, owners, and consumer/update routing.
- `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:11-63` — five-harness minimum qualification tier, valid-run meaning, and public-claim boundary.
- `plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md:5-63,65-93` — portable payload, materialization-versus-qualification boundary, ordinary lifecycle, and support classifications.
- `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:55-83,132-145` — current implementation and qualification work boundaries.
- `README.md:70-113,133-154` and `cli/README.md:1-46,123-204` — public product boundary, current CLI contract, artifacts, recovery, launcher prerequisites, and bounded smoke status.
- `plans/harness-setup.md:1-20,28-181,191-227,319-336,342-493,499-632,634-640` — current native setup, authentication, isolation, cleanup, executable, and platform contracts for Codex, Cursor, Claude Code, and Pi.
- `cli/internal/harness/harness.go:12-104` — supported installation IDs, evaluation IDs, native roots, and effective configuration defaults.
- `cli/internal/installer/installer.go:45-167` — ordinary payload materialization/removal and evaluation-only restoration semantics.
- `cli/internal/evaluation/evaluation.go:154-371,411-477,515-641` and `cli/internal/runstate/runstate.go:74-108,216-246` — output-root validation, tooling state, artifact persistence, cleanup, and interrupted-run recovery.
- `evaluations/skill-calling/instrumentation-contract.md:3-56` and `evaluations/skill-calling/smoke/real-harness-smoke-report.md:5-82` — governed evidence contract and the precise two-turn smoke observations.
- `research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md:104-190` and `assignments/10-cross-harness-adapter-contract.md:69-101` — retained, earlier receipt-backed lifecycle proposal and its unsupported boundaries.

## Findings

### Harness environment ownership stops at configuration and invocation

Skill Issue owns its embedded skill payload and run-scoped evaluation material; the installed harness executable, its login, credentials, normal configuration, managed policy, and host environment stay with the operator and harness. The CLI resolves and launches an existing executable, accepts `--executable` for an intentionally non-`PATH` launcher, and passes model/reasoning choices to the native harness. It installs or rematerializes only known Skill Issue skill directories; it does not install the host CLI, repair its installation, migrate its configuration, or take ownership of its credentials or native session store.

**Evidence:** `cli/README.md:30-34,42-46,194-198` makes model access and harness login user-provided prerequisites, documents `--executable`, and limits ordinary installation to the Skill Issue payload. `plans/harness-setup.md:30,149-152,324-325,473-481,532-536,623-627` requires existing native authentication and preserves it across Codex, Cursor, Claude Code, and Pi. `cli/internal/installer/installer.go:45-74` only materializes/removes payload directories, while `cli/internal/harness/harness.go:24-104` provides names and native skill roots rather than host installation logic.

**Implication:** Support language must distinguish **Skill Issue payload lifecycle** from **harness environment lifecycle**. “Repair” can describe a repeated `skill-issue install` replacing a known payload directory, but it cannot promise repair, update, migration, backup, restoration, or ownership of the surrounding harness installation or account.

### Materialized, qualified, and campaign-ready are different support states

Ordinary installation reports filesystem materialization only. Host discovery and activation require live Work Block 3 evidence; the release minimum additionally requires complete workflow proof and a tooling-complete three-scenario suite for the named matrix. A tooling-complete run with zero expected calls remains valid descriptive data, whereas a launch, permission, session, marker, timeout, cancellation, or protocol failure is a tooling failure to repair and rerun. “Unsupported,” “unqualified,” “caveated/fail-closed,” and “unrun” remain outside published model outcomes rather than being translated into model performance.

**Evidence:** `plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md:36-40,61-63,68-73` separates materialization from discovery/activation and defines public support distinctions. `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:38-59` defines tooling-complete and zero-call validity. `cli/README.md:194-198` categorizes native failures as tooling errors and names unqualified routes. `evaluations/skill-calling/instrumentation-contract.md:40-46` applies the same evidence rule in the runner contract.

**Implication:** A portability or qualification report should use state vocabulary precisely: **materialized**, **discovered**, **activated**, **tooling-complete**, **locally smoke-qualified**, **campaign-qualified**, **caveated/fail-closed**, **unsupported**, or **unrun**. It must not collapse a successful copy, a two-turn smoke, and a release-qualified workflow into one “supported” label.

### Qualification has a repeatable shared evidence sequence

For an accessible evaluation route, qualification starts by resolving the actual launcher and checking native authentication/model access; then fixes workspace, output, effective model, effective reasoning, and isolation configuration before temporary skills are installed. The runner sends immutable scenario turns in one persistent/resumable primary-agent conversation, requires the harness’s structured completion and attribution signals, stores evidence, and removes only run-owned material. Custom answer sheets are structurally validated but their semantic expected-call selection remains evaluator-owned.

**Evidence:** `cli/README.md:21-40,73-89,123-143` requires an effective configuration review, caller confirmation, structural custom-input validation, verbatim replay, and private attribution. `plans/harness-setup.md:74-145,149-170,405-469,548-612` specifies fixed executable/configuration per turn, native authentication preflights, structured terminal conditions, and primary-session continuity. `cli/internal/evaluation/evaluation.go:154-242,269-312,515-568` canonicalizes the workspace, checks authentication before instrumenting, installs only project-local evaluation skills, and rejects invalid custom input relationships.

**Implication:** A qualification package needs executable/version, operating system, harness authentication/access, selected model/reasoning, isolation and adjacent-configuration record, verbatim scenario/replay evidence, structured completion/attribution evidence, artifacts, and cleanup evidence. A test result or a claim that a binary cross-compiles is insufficient runtime qualification.

### Executable and platform evidence are route-specific prerequisites

The four implemented evaluation routes are Codex, Cursor, Claude Code, and Pi. The CLI resolves normal command names, with Cursor trying `agent` then `cursor-agent`, and allows an explicit launcher path. The retained setup and smoke evidence are macOS-only and version-specific: Codex `0.144.1`, Cursor Agent `2026.07.16-899851b`, Claude Code `2.1.205`, and Pi `0.80.10`. The Go CLI can cross-build macOS, Windows, and Linux binaries, but that does not establish native runtime behavior on those platforms.

**Evidence:** `cli/internal/harness/harness.go:24-79` limits current evaluation IDs to four and declares their defaults. `cli/internal/evaluation/runtime.go:220-242` resolves existing Cursor, Claude, and Pi executables without installing one. `plans/harness-setup.md:181,191,336,346,493,503,632` records macOS versions and executable expectations. `evaluations/skill-calling/smoke/real-harness-smoke-report.md:9-18` records the four actual launcher/authentication routes. `README.md:149-154` and `cli/README.md:202-216` state that cross-compilation does not replace native runtime testing.

**Implication:** The supported setup instruction must require a qualified executable and its native login before a run. Any other operating system, version, launcher form, model proxy, or harness surface needs its own evidence instead of inheriting the macOS smoke result.

### Output and recovery belong to the evaluation run, not the host environment

Every evaluation requires a caller-selected output root outside the evaluated workspace. A unique run directory always retains `result.json` and `website.json`; raw events and sanitized transcript are opt-in. Private run state, token mappings, and installation-state data live only beneath `<output>/.skill-issue/` while a run is active. Interrupted recovery uses the same output root and run ID to restore only pre-existing canonical Skill Issue paths, remove evaluation-introduced paths, and delete the run’s private state; completed artifacts remain. Harness history is left to the harness, including Codex’s normal resumable-session store.

**Evidence:** `cli/internal/evaluation/evaluation.go:171-262,333-371,411-477,609-641` enforces the external output root, writes artifacts, persists only evaluation installation state, and performs cleanup by run ID. `cli/internal/installer/installer.go:120-167` rematerializes current canonical payload only for recorded pre-existing skill names and removes temporary paths. `cli/internal/runstate/runstate.go:216-246` deletes private token/run state after cleanup. `cli/README.md:145-192` and `evaluations/skill-calling/instrumentation-contract.md:48-54` document the retained artifacts, optional diagnostics, and recovery limits.

**Implication:** Recovery support is constrained to Skill Issue’s own evaluation copy and private state. It does not restore arbitrary pre-run skill contents, host configuration, credentials, external project changes, or native harness environment state; selected result artifacts and ordinary harness history intentionally persist.

### Current evidence is a bounded smoke, not release qualification

The retained execution evidence covers built-in and custom **two-turn** routes for the four local launchers, selected configuration behavior, workspace writes, result artifacts, and cleanup. It expressly excludes the governed 30-turn suite, automated test suite, OpenCode, complete workflow proof, native-platform proof beyond the recorded macOS environment, and publication acceptance. The five-harness, thirteen-cell minimum remains a completion requirement, including OpenCode.

**Evidence:** `evaluations/skill-calling/smoke/real-harness-smoke-report.md:5-7,53-82` states the smoke scope, observations, cleanup, and deferred work. `plans/harness-setup.md:634-640` and `cli/README.md:200-204` repeat the bounded status. `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:11-59` and `skill-issue-project-completion-a-to-b-plan.md:71-83,140-145` define the future minimum campaign and full qualification evidence.

**Implication:** The four launcher routes can be described as locally smoke-qualified under their recorded conditions. They cannot be presented as a completed minimum support tier, cross-platform qualification, or published benchmark evidence.

### Retained direct-install research is evidence, not the current lifecycle promise

The historical research proposes a receipt-backed transaction with backups, update, repair, rollback, diagnostics, and verification. The completed architecture plan deliberately narrowed the implemented ordinary lifecycle to disposable known-payload replacement/removal with no ordinary receipt, backup, rollback inventory, mutable application state, separate update command, or separate repair command. The authority map marks the research as retained source evidence and directs current lifecycle decisions to the completed architecture plan and production CLI behavior.

**Evidence:** `research/deep-research/harness-direct-installation-architecture/assignments/10-cross-harness-adapter-contract.md:69-101` and `harness-direct-installation-architecture-deep-research.md:110-154` describe the earlier receipt/backup/repair proposal. `plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md:9,42-48` adopts the narrower lifecycle. `cli/README.md:42-46` and `plans/harness-setup.md:3-5` match that implementation. `plans/skill-issue-project-completion/document-authority-and-update-map.md:12-18,27-31,57-67` defines the research/evidence and current-owner hierarchy.

**Implication:** The receipt-backed proposal is a historical implementation option, not a present support commitment. Do not reuse its backup, rollback, migration, or host-repair language in current setup or portability guidance unless a later authoritative change adds the needed behavior.

### Open work must not be described as implemented support

Detection, target confirmation, preview, concrete installer preflight, fail-closed Grok discovery enforcement, final packaging, OpenCode evaluation, final regression, governed campaigns, and broad native-platform qualification remain open. The current CLI also rejects an independent Cursor reasoning override before evaluation side effects because Cursor owns model-native reasoning.

**Evidence:** `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:61-67,75-81` lists the pending Work Block 2 and 3 work. `cli/README.md:30,42-46,196-198` identifies the Cursor, detection/preview, and Grok boundaries. `evaluations/skill-calling/smoke/real-harness-smoke-report.md:68-82` verifies the Cursor behavior and preserves the deferred campaign boundary.

**Implication:** Treat these as explicit unsupported/unqualified gaps rather than ordinary troubleshooting. A support document should report the concrete blocker and keep the user’s harness environment unchanged instead of guessing a workaround.

## Notes

- The phrase “ordinary installation has no repair or update commands” in `plans/harness-setup.md:5` is consistent with `cli/README.md:46` and the direct-install plan: there is no separate command or retained rollback state; a repeated `install` is the deliberately limited payload replacement path.
- The historical direct-install research’s host-level repair/rollback ledger conflicts with the implemented ordinary lifecycle. Its source evidence remains useful for path and policy research, but its proposed lifecycle must not override production source or the completed architecture plan.
- The current contract supports installation into native **skill roots**, not installation of the harness executable or management of its broader environment. “Harness installation support” should be phrased with that object boundary visible.

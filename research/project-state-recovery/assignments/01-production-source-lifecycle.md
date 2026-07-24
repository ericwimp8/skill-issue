# Production Source Lifecycle

## Assignment

**Goal:** Establish production-source evidence for the repository's implemented capabilities, execution lifecycle, state ownership, constrained or incomplete paths, blockers, dependency order, and the implications for one authoritative current progress/task document.

**Scope:** Production code under `cli/cmd`, `cli/internal`, `bundle.go`, the embedded payload manifest, `go.mod`, and the CLI build/channel scripts. Nearby `README.md` and `cli/README.md` were inspected only to check whether their high-level ownership and status claims agree with production source.

**Exclusions:** Git-history reconstruction; broad inventory or interpretation of existing plans; broad website analysis; plugin-submodule internals; tests as behavioral authority; product-code, branch, Git-state, or existing-document changes.

## Sources

- Snapshot: branch `codex/post-submission-development`, commit `4979bd035680`; `git status --short -- cli bundle.go go.mod go.sum package.json`, `git diff --name-only -- cli bundle.go go.mod go.sum package.json`, and the corresponding untracked-file query were empty at inspection time, so the inspected production paths matched this commit.
- Entrypoint and command routing: `cli/cmd/skill-issue/main.go`; `cli/internal/command/command.go`, especially `App.Run`, `runGuidedInstall`, `runDoctor`, `runLifecycle`, `runEvaluation`, `reviewEvaluation`, and `printHelpTo`.
- Lifecycle parsing and orchestration: `cli/internal/lifecycle/lifecycle.go`, especially `Service.Execute`, `evaluationRun`, `runEvaluation`, `evaluationRunRequest`, `preparedWorkspace`, `evaluationStateRoot`, and `parseOptions`; `cli/internal/lifecycle/progress.go`.
- Capability registry: `cli/internal/harness/harness.go`, especially `orderedSpecs`, `InstallationAvailable`, `ParseEvaluationID`, `EvaluationDefaultsFor`, `SkillRoot`, `IncludeSkillFile`, and `TestedVersion`.
- Payload composition: `bundle.go`; `cli/internal/payload/assets/manifest.json`; `cli/internal/payload/payload.go`, especially `ReadManifest`, `Skills`, `LoadSkills`, `BuiltInEvaluation`, `readSkills`, and reference-closure validation.
- Installation and restoration: `cli/internal/installer/installer.go`, especially `Install`, `Uninstall`, `PrepareEvaluation`, `CleanupEvaluation`, `materializeSkills`, `applyHarnessMetadata`, `instrument`, and `materialize`.
- Evaluation execution and artifacts: `cli/internal/evaluation/evaluation.go`, especially `PrepareRequest`, `Service.Run`, `Cleanup`, `cleanupWithInstallation`, `finishCleanup`, `toolingFailure`, `deriveResult`, and `deriveWebsiteResult`; `cli/internal/evaluation/runtime.go`; `cli/internal/evaluation/transcript.go`.
- Private run state: `cli/internal/runstate/runstate.go`, especially stable `Status` values, `Run`, `Store.Create`, `Store.Mark`, locking, token mappings, and deletion.
- Harness replay: `cli/internal/replay/replay.go`, especially `Runner.Run`; `cli/internal/replay/process.go`, especially `NewAdapter`, process-session execution, protocol checks, cleanup, command builders, and authentication checks; `cli/internal/replay/pi.go`.
- Diagnostics: `cli/internal/doctor/doctor.go`, especially `Run`, platform checks, harness checks, version checks, authentication checks, and Codex configuration parsing.
- Build and local development configuration: `go.mod` (Go `1.24.0`); `cli/scripts/build-cross-platform.sh`; `cli/scripts/local-cli.sh`.
- Discrepancy-only references: `README.md` lines 168-194 and 275-294; `cli/README.md` command, artifact, recovery, and harness-boundary sections.

## Findings

### Finding 1: Production capability ownership is deliberately split by concern

The executable is a thin command router over concrete owners rather than a monolithic workflow. `command.App` owns the user-facing command surface and confirmation, `lifecycle.Service` owns action sequencing, `harness.orderedSpecs` owns supported harness capabilities and defaults, `payload` owns canonical bundle loading, `installer` owns filesystem materialization and restoration, `evaluation` owns governed-run orchestration and evidence, `runstate` owns temporary recovery state, and `replay` owns native harness sessions and protocol validation.

**Evidence:** `cli/cmd/skill-issue/main.go:15-22` creates `command.App`; `cli/internal/command/command.go:49-73` dispatches commands; `cli/internal/lifecycle/lifecycle.go:57-69` dispatches lifecycle actions; `cli/internal/harness/harness.go:62-78` declares its registry the single source for harness listing and capability checks; `cli/internal/evaluation/evaluation.go:181-194` composes run-state, installer, and replay-adapter owners.

**Implication:** A current-progress document should cite these owners when describing product truth. It should not duplicate low-level contracts such as harness defaults, skill destinations, run statuses, or artifact schemas as editable task-state facts; those remain source-owned.

### Finding 2: The implemented CLI surface is installation, diagnostics, governed evaluation, cleanup, and internal attribution

The public command router implements `help`, `version`, `install`, `doctor`, `uninstall`, and `evaluate`. The hidden `signal` command is an instrumentation callback, not an end-user workflow. Generation, diagnosis, and refinement are not CLI subcommands; after installation, the CLI explicitly tells the operator to invoke `skill-intake` in the selected harness.

**Evidence:** `cli/internal/command/command.go:49-73` is the complete top-level switch; `cli/internal/command/command.go:141-174` completes guided installation and hands control to `skill-intake`; `cli/internal/command/command.go:475-494` lists the public foundation and lifecycle commands. The only production references to generation/refinement are payload skill names, not command branches.

**Implication:** Implemented CLI work and installed-skill workflow work are separate status categories. A progress document that says “generation/refinement implemented in the CLI” would be false; the accurate statement is that the CLI distributes the skills that own those workflows.

### Finding 3: Installation and uninstallation are complete for all five registered harnesses

The source registry contains Claude Code, Codex, Cursor, OpenCode, and Pi. Every current registry entry has installation metadata and none sets `InstallationInProgress`, so guided and scripted installation are source-enabled for all five at project or user scope. Installation loads the canonical payload, applies harness-specific metadata, atomically stages each known skill directory, replaces only those known directories, and verifies every expected regular file. Uninstallation removes those same payload directories.

**Evidence:** `cli/internal/harness/harness.go:64-70` defines all five specs; `cli/internal/harness/harness.go:88-91` gates only entries marked in progress; `cli/internal/lifecycle/lifecycle.go:76-108` calls installer install/uninstall; `cli/internal/installer/installer.go:54-87` defines the concrete effects; `cli/internal/installer/installer.go:344-379` verifies materialized files; `cli/internal/installer/installer.go:561-594` stages and renames each skill directory. `cli/internal/payload/assets/manifest.json:5-50` currently lists eleven canonical skills.

**Implication:** The generic `InstallationInProgress` field is dormant capability scaffolding, not evidence of an unfinished current harness. A status document should classify five-harness installation as implemented and treat any future harness as a separate task only when a registry entry or other concrete work exists.

### Finding 4: The embedded payload has a concrete build-time dependency and a development identity

The Go binary embeds the three lifecycle skills from the pinned plugin submodule plus supporting and discipline skills from this repository, and it embeds built-in evaluation JSON. Payload loading validates manifest identity, source uniqueness, frontmatter, and local-reference closure before installation. The payload manifest still labels its payload version `development`; ordinary `go build` defaults binary metadata to `dev`, `unknown`, and `unknown`, while release/local scripts inject explicit metadata.

**Evidence:** `bundle.go:5-6` defines the embedded filesystem; `cli/internal/payload/assets/manifest.json:2-49` defines the eleven components and `payload_version`; `cli/internal/payload/payload.go:43-70` validates the manifest and `:129-181` loads and validates canonical skills; `cli/cmd/skill-issue/main.go:9-20` defines default build metadata; both build scripts pass linker values.

**Implication:** A usable build depends on the pinned submodule contents being initialized before compilation. The `development` payload label and default build metadata indicate development/release identity, not an incomplete runtime path. Release readiness should be tracked separately from implemented payload behavior.

### Finding 5: Doctor is an implemented preflight, with explicit limits

`doctor` checks macOS/Linux evaluation support, temporary-directory and home availability, executable discovery, recorded harness-version compatibility, native authentication, OpenCode provider/model availability, Pi agent-directory presence, and whether the installed Codex parses the exact generated evaluation configuration. It streams findings and returns unhealthy on any failure. It performs no model turn and therefore cannot establish end-to-end runtime qualification.

**Evidence:** `cli/internal/doctor/doctor.go:53-89` builds the report; `:91-127` performs platform and per-harness checks; `:165-188` distinguishes pinned from advisory versions; `:191-263` checks authentication/model prerequisites; `:265-283` parses generated Codex configuration. The evaluation runtime independently performs authentication at `cli/internal/evaluation/evaluation.go:310-312`.

**Implication:** Doctor readiness and successful evaluated conversations are distinct progress facts. A project-status document should never convert a healthy doctor report or a registry version string into “campaign qualification complete”; live qualification evidence must be tracked by the task that produced it.

### Finding 6: Governed evaluation has a complete, ordered execution path

An evaluation requires a supported harness, resolves default model/reasoning and either a built-in evaluation or a complete custom skills/scenario/answer-sheet set, caches parsed inputs before confirmation, prepares a workspace, enforces output and answer-sheet separation from that workspace, creates a run ID and private state, prepares an isolated harness runtime, checks authentication, instruments and temporarily installs the selected skills, runs all effective turns in one resumable native session, attributes calls at turn boundaries, derives evidence, restores prior skills, removes runtime/private state, and returns the result. The default built-in is `gardening-web-application`; `--turns` can truncate the scenario and expected calls together.

**Evidence:** `cli/internal/lifecycle/lifecycle.go:140-177` handles confirmation, signals, and execution; `:193-278` builds and prepares the request; `cli/internal/evaluation/evaluation.go:96-125` resolves and caches inputs; `:197-326` validates paths and prepares state/runtime/auth/instrumentation; `:342-501` persists installation state, runs the adapter, records calls, writes artifacts, and cleans; `cli/internal/replay/replay.go:130-168` implements sequential multi-turn replay.

**Implication:** Dependency order matters for any remaining evaluation work: payload and scenario validity precede confirmation; external harness/auth readiness precedes temporary skill installation and replay; protocol/attribution success precedes scoring; successful restoration precedes completion. Progress tasks should preserve this order instead of treating “run evaluation” as one opaque item.

### Finding 7: Evaluation artifacts and private recovery state have different owners and lifetimes

Durable caller-facing artifacts live in a unique run directory under the selected output root. A tooling-complete run always writes authoritative `result.json` (schema 1) and compact `website.json` (schema 2), with optional `events.jsonl` and sanitized `transcript.json`. A tooling failure best-effort writes sanitized `failure.json`. In contrast, `<output>/.skill-issue/` owns restrictive run JSON, tokens, internal signal events, installation state, and backups only for execution/recovery. Successful cleanup deletes this private state while retaining caller-facing evidence.

**Evidence:** `cli/internal/evaluation/evaluation.go:143-179` defines result schemas; `:245-265` creates output and private run state; `:450-501` writes optional and required artifacts then cleans; `:1198-1229` writes failure diagnostics; `:1287-1339` derives schema-2 website points. `cli/internal/runstate/runstate.go:18-50` defines stable internal statuses and run fields; `:91-127` persists private state with restrictive modes; `:242-272` removes tokens and run directories.

**Implication:** Evaluation state cannot serve as the repository's authoritative project-status document: it is run-scoped, opaque-token-bearing, output-root-local, and intentionally deleted after cleanup. Result artifacts prove particular runs; they do not own roadmap or task status.

### Finding 8: Recovery is implemented, but only as explicit cleanup of a known run

The user-facing recovery command is `evaluate cleanup --run <id> --output <root>`. It loads the retained run, deletes recoverable native OpenCode session state when applicable, decodes installation state, restores backups or canonical prior paths, removes temporary evaluation paths and runtime mappings, updates status, and deletes the run state. There is no public command to list runs, inspect status, resume an interrupted evaluation, or discover a lost run ID.

**Evidence:** `cli/internal/lifecycle/lifecycle.go:111-136` exposes only `evaluate run` and `evaluate cleanup`; `cli/internal/evaluation/evaluation.go:871-960` implements cleanup; `cli/internal/runstate/runstate.go:330-347` shows the private run/token layout. The full top-level switch at `cli/internal/command/command.go:55-69` contains no `status`, `list`, or `resume` branch.

**Implication:** Cleanup is a completed bounded recovery path, while discoverability and user-directed evaluation resumption are absent capabilities. If desired, they are future product tasks; they should not be inferred as active work without an authoritative task decision.

### Finding 9: Harness isolation and protocol validation are implemented per adapter, with external blockers

All five registered harnesses have concrete evaluation adapters. Codex uses a run-owned `CODEX_HOME`, explicit configuration, workspace-write sandboxing, ignored ambient user/rule configuration, and command-event attribution. Cursor uses a private home/plugin plus allowlisted permissions. Claude uses private passed-skills/settings and purges the temporary project record. OpenCode uses private XDG configuration/state/cache, deny-first permissions, version pinning, discovery checks, stable session IDs, and explicit session deletion. Pi runs a private RPC session with explicit skills, offline mode, bounded control calls, and no session file.

**Evidence:** `cli/internal/evaluation/runtime.go:32-81` dispatches runtime preparation; `:84-97` defines Codex configuration; `:114-256` defines OpenCode isolation; `:263-444` defines Cursor, Claude, and Pi runtimes. `cli/internal/replay/process.go:83-138` constructs native adapters; `:268-328` validates protocols/session identity; `:346-425` performs native cleanup; `:763-804` defines native commands. `cli/internal/replay/pi.go:35-72` and `:152-252` implement Pi RPC startup and turn capture.

**Implication:** The remaining blockers for an actual run are external and configuration-specific: supported OS, executable presence, authentication, compatible native protocol/version/model, writable private/output locations, and successful cleanup. Static source establishes implemented mechanisms, not that the current machine or every future harness version satisfies them.

### Finding 10: Current source-evidenced limitations should remain explicit

Evaluation is intentionally rejected on Windows even though six Darwin/Linux/Windows archives are cross-built. OpenCode's recorded version is an exact pin unless `SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS=1`; other version drift is advisory. Custom evaluation requires all three inputs and keeps the answer sheet outside the evaluated workspace. The shared lifecycle option parser accepts any syntactically valid `--key` and does not reject unknown non-boolean options; unconsumed keys are silently ignored, unlike `doctor`, which explicitly rejects unsupported options.

**Evidence:** `cli/internal/evaluation/evaluation.go:96-101` rejects non-Darwin/Linux evaluation; `cli/scripts/build-cross-platform.sh:32-38` builds six OS/architecture targets. `cli/internal/harness/harness.go:65-69` records tested versions and pins OpenCode; `cli/internal/replay/process.go:593-605` enforces the pin. `cli/internal/evaluation/evaluation.go:998-1055` validates custom inputs and answer-sheet placement. `cli/internal/lifecycle/lifecycle.go:414-453` stores arbitrary option keys, while `cli/internal/command/command.go:308-329` restricts doctor options.

**Implication:** Windows evaluation qualification, any OpenCode version update, and lifecycle unknown-option rejection are concrete bounded work candidates. They are limitations, not proven active tasks, until the authoritative progress document accepts and prioritizes them.

### Finding 11: Build/channel management is implemented, while publication remains outside the CLI runtime

The cross-platform script builds trimpath, metadata-injected, CGO-disabled binaries for six targets, packages archives, and emits checksums. The local wrapper separately owns an immutable commit-specific known-good executable built from a Git archive and a development executable built from the current working tree with clean/dirty identity; known-good is the default execution channel. Neither script publishes, signs, or installs a release.

**Evidence:** `cli/scripts/build-cross-platform.sh:9-38` builds six binaries and `:40-75` packages and checksums them. `cli/scripts/local-cli.sh:27-64` builds current-source development, `:66-97` builds/selects committed known-good, and `:135-156` defaults execution to known-good.

**Implication:** “Cross-build implemented” and “public distribution complete” are different statuses. The project document should track release publication/qualification separately and identify the local wrapper as developer channel state rather than product lifecycle state.

### Finding 12: Production source does not own durable project progress

No production package stores repository roadmap state, task dependencies, accepted completion, next action, or campaign scheduling. The only statuses in production source are ephemeral evaluation-run stages (`preparing`, `running`, `complete`, `complete-cleaned`, `cleaned`, `tooling-failed`) and per-check doctor results. Nearby README status text accurately summarizes broad capability ownership, but its campaign pointer is a prose reference, not production validation of current campaign progress.

**Evidence:** `cli/internal/runstate/runstate.go:18-29` defines only evaluation-run statuses; `cli/internal/doctor/doctor.go:26-45` defines check/report status; the command surface has no project-management action. `README.md:168-177` correctly separates CLI-owned deterministic lifecycle behavior from installed-skill generation/refinement, while `README.md:275-286` contains manually curated project-status prose and a campaign-document pointer.

**Implication:** The semantic home for live project status should be one tracked, human-readable document under `plans/`, consistent with repository instructions, with explicit sections for: verified implemented baseline; active work; blockers/prerequisites; dependency-ordered next tasks; evidence links; and superseded/closed items. Production source and retained run artifacts should be cited as evidence, while task decisions and continuation state should be owned once in that document. README should retain a stable high-level summary and point to the authoritative status document rather than becoming a second mutable task ledger.

## Notes

- No production `TODO`, `FIXME`, panic placeholder, or registry entry marked `InstallationInProgress` was found. Absence of such markers does not prove that qualification, release, or campaign work is complete; it only means production source does not encode an active partial implementation.
- Tests were not used to infer contracts or intended behavior. No live harness run was needed for this source-lifecycle assignment.
- Static source cannot establish current external authentication, installed executable versions, model availability, successful native qualification, release publication, or campaign completion. Those claims require their own retained runtime or operational evidence.
- Useful search terms for follow-up reconciliation: `InstallationInProgress`, `payload_version`, `StatusToolingFailed`, `evaluationStateRoot`, `installation-state.json`, `cleanupWithInstallation`, `build-known-good`, `SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS`, and `evaluation-progress.md`.

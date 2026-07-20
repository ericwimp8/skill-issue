# CLI Evaluation Production Path

## Assignment

- **Goal:** Trace the current production evaluation path end to end, from the `skill-issue evaluate` CLI route through request resolution, harness runtime preparation, skill instrumentation, replay, scoring, artifact writing, cleanup, and structured command output.
- **Scope:** Current working-tree production sources under `cli/cmd/skill-issue`, `cli/internal/command`, `cli/internal/lifecycle`, `cli/internal/evaluation`, and their concrete callees in `harness`, `payload`, `installer`, `replay`, and `runstate`; the embedded built-in evaluation bundle and the controlling website-migration plan were inspected only where they establish inputs or distinguish implemented behavior from planned consumption.
- **Exclusions:** Tests as behavioral authority; live harness qualification; semantic correctness of answer sheets; interpretation of historical output produced by older source states; website implementation details; chart-design recommendations; product-source edits.

## Sources

- `cli/cmd/skill-issue/main.go:15-22` — process entrypoint and standard-stream wiring.
- `cli/internal/command/command.go:38-156` — command routing, evaluation warnings, interactive review, JSON stdout, stderr failures, and exit codes.
- `cli/internal/lifecycle/lifecycle.go:27-37, 45-62, 96-234, 283-320` — lifecycle result envelope, evaluation subcommands, option parsing, request construction, signal-aware context, output-owned state root, and cleanup route.
- `cli/internal/harness/harness.go:10-22, 38-60, 62-87` — harness identities, evaluation-supported subset, and model/reasoning defaults.
- `cli/internal/evaluation/evaluation.go:20-129` — answer-sheet, request, detailed-result, compact-website-result, and service structures.
- `cli/internal/evaluation/evaluation.go:131-329` — complete evaluation run sequence and public artifact writers.
- `cli/internal/evaluation/evaluation.go:331-447` — Codex signal recovery, cleanup, manual recovery, and status transitions.
- `cli/internal/evaluation/evaluation.go:472-640` — built-in/custom loading, path constraints, schema validation, output-root preparation, and opaque-token generation.
- `cli/internal/evaluation/evaluation.go:642-787` — result comparison, evaluation identity, compact website projection, CLI executable resolution, and JSON/JSONL persistence.
- `cli/internal/evaluation/runtime.go:16-64` — runtime-preparation contract and supported harness dispatch.
- `cli/internal/evaluation/runtime.go:66-186` — concrete Cursor, Claude Code, and Pi temporary runtime layouts.
- `cli/internal/evaluation/runtime.go:188-254, 256-323` — runtime file writing, controlled environments, temporary root, and Codex skill-disable configuration.
- `cli/internal/installer/installer.go:16-37, 77-193` — evaluation-installation request/state, temporary materialization, collision recording, and restoration.
- `cli/internal/installer/installer.go:195-335` — evaluation-root resolution, instrumentation injection, staging, and committed skill copies.
- `cli/internal/payload/payload.go:21-40, 53-110, 113-193, 196-275` — embedded/custom skill loading, built-in evaluation loading, manifest filtering, and runnable-skill validation.
- `bundle.go:5-6` — compile-time embedding of canonical skills and built-in evaluations.
- `cli/internal/replay/replay.go:27-77, 79-160` — scenario, capture, transcript result, session interfaces, ordered replay, and turn boundaries.
- `cli/internal/replay/process.go:16-106, 111-218` — process-adapter construction, per-turn launch/resume, model flags, native stdout/stderr capture, session extraction, and protocol validation.
- `cli/internal/replay/process.go:220-316, 335-492` — protocol requirements, Cursor/Claude arguments, environment behavior, event parsing, session discovery, and command templates.
- `cli/internal/replay/pi.go:20-93, 125-213, 282-353` — Pi RPC launch, preflight, per-turn capture, session shutdown, model parsing, supplied-skill validation, and state validation.
- `cli/internal/runstate/runstate.go:15-48, 58-107` — private run/event schemas, random identities, token mappings, and `run.json` persistence.
- `cli/internal/runstate/runstate.go:110-246, 249-337` — turn/session mutation, signal-event append, event loading, cleanup, locks, paths, and atomic private-state writes.
- `cli/internal/payload/assets/manifest.json` — embedded canonical component inventory.
- `evaluations/skill-calling/built-ins/*.json` — embedded schema-version-1 evaluation units containing scenario and answer sheet.
- `evaluations/skill-calling/event.schema.json:1-52` — committed event schema used for drift comparison with the production emitter.
- `cli/README.md:136-183` — documented public artifact contract, inspected as documentation rather than behavioral authority.
- `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:85-97` — controlling plan showing website consumption of compact artifacts remains future work.

## Findings

### Finding 1: The CLI evaluation route is interactive and returns one structured lifecycle envelope

`main` binds the application to process stdin/stdout/stderr, then `App.Run` routes `evaluate` into `runLifecycle`. The production CLI special-cases `evaluate run`: it prints pre-parse sensitivity warnings to stderr, resolves the request through `ExecuteEvaluationRun`, prints the effective harness/model/reasoning and built-in or custom input selection to stderr, and requires a `y`/`yes` response. Any other response produces a successful cancellation result without starting the evaluator. A completed run is emitted to stdout as indented JSON with the lifecycle envelope `{action:"evaluate", status:"complete", data:<evaluation.Result>}`. Cancellation emits `{action:"evaluate", status:"cancelled"}`. Errors print a single line to stderr and exit 1; unknown top-level commands exit 2. The internal `signal` route deliberately emits no stdout on success.

**Evidence:** `cli/cmd/skill-issue/main.go:15-22`; `cli/internal/command/command.go:38-89, 91-156`; `cli/internal/lifecycle/lifecycle.go:27-31, 125-154, 156-167`.

**Implication:** Machine-readable stdout is protected from warnings and review text, but command-line automation must still provide confirmation on stdin. The stdout result does not include the generated output-directory path or either compact artifact's content, so a downstream importer cannot discover the run directory from the command response alone.

### Finding 2: Lifecycle parsing constructs one resolved request for either built-in or custom mode

The lifecycle layer accepts `--events` and `--transcript` as the only boolean options and otherwise parses every `--name value` pair into a string map. It requires `harness`, `workspace`, and `output`; makes workspace and output absolute; selects the four evaluation-capable harnesses; and enforces exactly one input mode: a single `--evaluation` identifier, or all of `--skills`, `--scenario`, and `--answer-sheet`. The resulting `RunRequest` also carries model/reasoning values and override booleans, an optional harness executable override, an optional signal-CLI override, and diagnostic flags. `ResolveRequest` supplies defaults before review: Claude Code `opus`/`medium`, Codex `gpt-5.6-sol`/`medium`, Cursor `auto`/`medium`, and Pi `openai-codex/gpt-5.6-sol`/`medium`. An explicit Cursor reasoning override is rejected.

**Evidence:** `cli/internal/lifecycle/lifecycle.go:125-143, 169-234, 283-320`; `cli/internal/evaluation/evaluation.go:31-64`; `cli/internal/harness/harness.go:55-87`.

**Implication:** The reviewed and persisted identities are resolved values rather than blank defaults. Unknown option names are not rejected after parsing; a misspelled value-bearing option can be silently ignored while defaults take effect, which weakens command reproducibility.

### Finding 3: Built-in and custom inputs converge before any harness starts

Built-in mode reads `<evaluation-id>.json` from the executable's embedded `evaluations/skill-calling/built-ins` tree. Each embedded unit must have schema version 1, a matching `evaluation_id`, a valid scenario, and a valid paired answer sheet; its selected skill set is the complete canonical Skill Issue payload. Custom mode requires a directory whose direct children are skill directories, rejects symlinks and non-directory root entries, validates each `SKILL.md` frontmatter name and referenced local-file closure, loads a schema-version-1 scenario with nonempty unique turn IDs, and validates a schema-version-1 answer sheet against the scenario and supplied skill names. Only the custom answer-sheet path is required to remain outside the evaluated workspace.

**Evidence:** `bundle.go:5-6`; `cli/internal/payload/payload.go:40, 61-110, 113-193, 196-275`; `cli/internal/evaluation/evaluation.go:66-77, 472-526, 528-544, 590-628`; `cli/internal/replay/replay.go:27-59`.

**Implication:** Both modes enter the identical instrumentation/replay/scoring path after loading. Structural validation establishes runnable links, while the semantic correctness of expected calls remains caller-owned in custom mode. The selected input paths, file hashes, payload version, and skill contents are absent from final result identity.

### Finding 4: Run identity, evaluation identity, and output placement are distinct

Each run receives a 16-byte cryptographically random ID encoded as 32 lowercase hex characters. `startedAt` is captured in UTC, and the public run directory is `<output>/<harness>-<YYYYMMDDTHHMMSSZ>-<first-8-run-id>/`, created with mode `0700`. The full run ID remains in artifacts. `evaluation_id` is the selected built-in identifier in built-in mode and the scenario ID in custom mode; `scenario_id` always comes from the scenario. Scope is hard-coded to `project`. The mandatory output root is made absolute and created with mode `0700`, and both lexical and symlink-resolved checks reject a root inside the evaluated workspace.

**Evidence:** `cli/internal/runstate/runstate.go:66-72`; `cli/internal/evaluation/evaluation.go:136-186`; `cli/internal/evaluation/evaluation.go:547-588, 689-694`.

**Implication:** `run_id` is the only collision-resistant per-execution join key. Harness, model, scenario, and evaluation identify experimental dimensions but are not unique run identities. Directory names carry only an eight-character prefix and second-resolution time, so importers should read the full artifact fields rather than parse identity solely from the path.

### Finding 5: The CLI creates private output-owned state before runtime preparation

Before launching a harness, the evaluator creates `<output>/.skill-issue/runs/<run-id>/run.json` plus `<output>/.skill-issue/tokens/<64-hex-token>` files. `run.json` schema version 1 contains ID, canonical workspace, harness, model, reasoning, evaluation ID, scenario, project scope, status, active turn, harness session, token-to-skill map, installation-state path, evidence/transcript paths, and timestamps. Each token file contains the owning run ID and a newline. Later signaling appends internal events to `<output>/.skill-issue/runs/<run-id>/events.jsonl`. Mutations use transient `.lock` files and atomic `.state-*` rename writes. After skill installation, `installation-state.json` stores `{root, preexisting, skills}` and `run.json` advances from `preparing` to `running`.

**Evidence:** `cli/internal/evaluation/evaluation.go:161-199, 227-240`; `cli/internal/runstate/runstate.go:15-48, 58-107, 147-213, 249-337`; `cli/internal/installer/installer.go:33-37, 170-193`.

**Implication:** These files are recovery and attribution state, not durable public evidence. Successful cleanup deletes them, so any website or campaign pipeline must rely on the public run directory and cannot expect retained private state, session ID, installation receipt, or internal event log unless raw events were explicitly exported.

### Finding 6: Instrumentation materializes disposable skill copies with per-skill opaque signals

The evaluator generates one random 32-byte token for every selected skill. The installer clones each complete skill, inserts exactly one `skill-issue signal <token> <state-root>` instruction immediately after the closing frontmatter delimiter, and stages then renames the instrumented directory into the harness-specific evaluation root. It records names that existed beforehand only when they match embedded ordinary Skill Issue components. During cleanup those names are rematerialized from the current embedded ordinary payload; other temporary evaluation paths are removed. Contents are not backed up.

**Evidence:** `cli/internal/evaluation/evaluation.go:161-164, 205-234, 630-640`; `cli/internal/installer/installer.go:77-118, 120-167, 267-335`.

**Implication:** Skill identity is hidden from the harness command by the opaque token, while the private token map recovers it. Cleanup preserves path presence rather than prior bytes; a same-name preexisting path is restored as the embedded canonical ordinary skill, which is a material limitation for any workspace containing independently modified same-name Skill Issue content.

### Finding 7: Runtime preparation is concrete and currently limited to four evaluation harnesses

The CLI permits evaluation only for Claude Code, Codex, Cursor, and Pi. Codex runs directly in the requested workspace, uses project-local `.agents/skills`, ignores user config and rules, disables plugins/apps/project instructions, supplies a discovered user-skill deny-list, requests `workspace-write` plus on-request approval, and passes model and reasoning explicitly. Cursor builds a clean temporary home, config, plugin manifest, plugin skills, store, and workspace under `${TMPDIR}/skill-issue/<run-id>` and launches against that temporary workspace. Claude builds a temporary launch directory and passed-skills root, exposes the requested workspace as an additional directory, and appends an absolute-workspace system prompt. Pi builds a clean offline RPC runtime with explicit supplied skills under the same temporary root and launches from a temporary workspace. The private runtime root is recursively removed after the run path exits.

**Evidence:** `cli/internal/harness/harness.go:55-87`; `cli/internal/evaluation/runtime.go:28-64, 66-186, 223-254, 256-323`; `cli/internal/replay/process.go:133-177, 296-316, 484-492`; `cli/internal/replay/pi.go:30-93`.

**Implication:** Generic replay constants and command templates for additional harnesses do not make those harnesses reachable through the production evaluation CLI. Codex operates on the caller's workspace and Claude receives an explicit route to it. Cursor's preparation receives no caller-workspace argument, and Pi's `workspace` parameter is unused; both run in generated temporary workspaces, so the selected workspace currently acts as run identity/path-safety input rather than the harness's working project for those two routes.

### Finding 8: Replay uses one resumable primary session and establishes turn attribution externally

The runner validates the scenario, starts one adapter session, and processes turns in order. Before each prompt it sets `run.ActiveTurn` to the scenario turn ID; it then sends the prompt verbatim, waits for a terminal native capture, records any harness session ID, and clears the active turn. Process adapters capture each native invocation's complete stdout and stderr, parse JSON or JSONL events, require harness-specific success markers, and resume later prompts with the discovered session ID. Pi instead keeps one RPC process alive. Each replay result contains `harness_id`, the full scenario, and ordered turn results whose capture contains `session_id`, raw transcript string, optional stderr, and raw structured events.

**Evidence:** `cli/internal/replay/replay.go:61-77, 112-160`; `cli/internal/evaluation/evaluation.go:241-290`; `cli/internal/replay/process.go:111-218, 220-270, 385-458`; `cli/internal/replay/pi.go:145-213`.

**Implication:** Turn assignment does not depend on model-generated labels. A signal received outside an active turn is retained as unattributed. The replay transcript is already a structured aggregate of full scenario plus native per-turn captures; it is much broader and more sensitive than the compact invocation evidence.

### Finding 9: Skill invocation events are the sole scoring input

For non-Codex harnesses, an instrumented skill executes the silent `signal` subcommand, which resolves the token, reads the active turn, and appends a schema-version-1 `skill_invocation` event. Each emitted event includes `run_id`, optional `turn_id`, `attributed`, harness, model, reasoning, evaluation ID, skill, and `recorded_at`. Codex is special: after each turn, the evaluator scans native structured `command_execution` events for opaque tokens and calls the same marker itself, avoiding a requirement that the attempted command successfully write private state. Existing same-turn skill observations prevent double-recording that Codex recovery path.

**Evidence:** `cli/internal/command/command.go:61-67`; `cli/internal/lifecycle/lifecycle.go:156-167`; `cli/internal/runstate/runstate.go:36-48, 147-191`; `cli/internal/evaluation/evaluation.go:262-280, 331-365`.

**Implication:** Result derivation observes only recorded skill signals; it does not infer invocation from prose, skill file reads, or final responses. A denied Codex signal attempt can still count because the structured attempted command is the attribution source.

### Finding 10: `result.json` is the detailed durable evidence and is also returned on stdout

After replay, the evaluator compares attributed events to expected `{turn_id, skill}` pairs. The schema-version-1 `Result` contains `run_id`, harness, model, reasoning, evaluation ID, scenario ID, project scope, start/completion timestamps, and arrays for expected, observed, missing, additional, and unattributed calls. Attributed events are appended to `observed`; unexpected attributed calls are also appended to `additional`; events outside active turns go only to `unattributed`; each expected pair absent from the observed set goes to `missing`. The optional relative `transcript_path` appears only when transcript persistence succeeded. The result is written as mode-`0600` indented JSON to `result.json`, then returned inside the stdout lifecycle envelope.

**Evidence:** `cli/internal/evaluation/evaluation.go:20-29, 79-96, 286-307, 642-687, 763-771`; `cli/internal/command/command.go:83-89, 144-150`.

**Implication:** The detailed result is descriptive and contains no pass/fail or aggregate score. Duplicate attributed events remain duplicated in `observed` and potentially `additional`, although missing detection uses set membership. Durable provenance omits CLI version/commit, native harness version, executable path, session ID, payload version, input paths, and input hashes.

### Finding 11: `website.json` is an implemented lossy projection over expected-call turns

The evaluator always writes a second mode-`0600` schema-version-1 artifact. `WebsiteResult` contains `run_id`, `scenario_id`, harness, model, `total_turns`, and `points`. It builds a unique expected-skill set per turn, builds an observed pair set, walks scenario turns in source order, and emits a point only for a turn with at least one expected call. Each point stores the one-based scenario position as numeric `turn`, preserves `turn_id`, counts expected skills observed on that exact turn as `called`, and computes `missed` as unique expected skills minus called skills. Repeated events do not inflate the compact counts.

**Evidence:** `cli/internal/evaluation/evaluation.go:98-113, 308-310, 696-738`.

**Implication:** The artifact is directly numeric-axis-ready, but it is intentionally incomplete. It omits evaluation ID, reasoning, scope, timestamps, skill names, expected/observed pairs, additional calls, unattributed calls, and transcript linkage. It therefore cannot independently support all comparison dimensions or explain a line point; downstream website work must retain a join to `result.json` when those dimensions matter. `total_turns` can exceed the greatest emitted point and gaps are expected because non-expected turns are omitted.

### Finding 12: Optional diagnostics are explicit and independent

`--events` writes public `events.jsonl` by re-encoding the internal invocation events used for result derivation; with no events the file is empty. `--transcript` writes `transcript.json` as the complete replay `Result`, including the full prompts embedded in the scenario and each turn's native stdout transcript, stderr, session ID, and raw event array. Enabling the transcript also sets `result.transcript_path` to the basename `transcript.json`. The flags may be enabled independently; neither changes `website.json`.

**Evidence:** `cli/internal/evaluation/evaluation.go:286-310, 774-787`; `cli/internal/replay/replay.go:61-77`; `cli/internal/command/command.go:122-131`.

**Implication:** Default tooling-complete persistence is exactly `result.json` plus `website.json`. Raw diagnostic retention is opt-in, and transcript persistence carries the complete conversation without scanning or redaction. The public event export is narrower than the transcript but currently has a schema-contract defect described below.

### Finding 13: Successful cleanup removes recovery state but preserves public artifacts

After both public artifacts are written, the evaluator updates private `run.json` with the evidence path, optional transcript path, and `complete` status, restores/removes temporary skills, clears installation state, removes the temporary harness runtime, deletes token mappings, briefly records a cleaned status, and deletes the entire private run directory. Manual `evaluate cleanup --output <root> --run <id>` loads the same output-owned run and installation state to complete recovery after interruption. Public run directories are never removed by this cleanup path.

**Evidence:** `cli/internal/evaluation/evaluation.go:312-328, 368-434`; `cli/internal/lifecycle/lifecycle.go:96-123, 220-234`; `cli/internal/runstate/runstate.go:216-246`.

**Implication:** A successful run leaves only user-facing files and harness-owned session history. Interrupted and failed runs can leave public directories, partial artifacts, and private recovery state. Public artifact writes are ordinary `os.WriteFile` operations rather than atomic commits, so a failure after `result.json` but before `website.json` can leave an incomplete public run directory even while deferred cleanup attempts to remove private state.

### Finding 14: Several current labels and schemas overstate native guarantees

Cursor's resolved request/result records `reasoning:"medium"`, but the runtime rejects explicit reasoning overrides and sends no independent reasoning control to Cursor. Pi accepts a reasoning override and passes it as `--thinking`, but post-turn state validation rejects any reported thinking level other than the hard-coded string `medium`. Separately, the production invocation-event emitter always serializes `reasoning` and `evaluation_id`, while the committed JSON Schema forbids additional properties and declares neither property; emitted events therefore do not conform to that schema as written.

**Evidence:** `cli/internal/harness/harness.go:55-60`; `cli/internal/evaluation/evaluation.go:49-63, 670-686`; `cli/internal/replay/process.go:296-305`; `cli/internal/replay/pi.go:48-62, 337-353`; `cli/internal/runstate/runstate.go:36-48, 165-178`; `evaluations/skill-calling/event.schema.json:5-50`.

**Implication:** Consumers should treat the detailed result's reasoning field as a Skill Issue configuration label, not uniformly verified native runtime telemetry. Pi non-medium overrides are presently self-contradictory. The Go emitter is the current production truth for `events.jsonl`; schema-based validation against the committed event schema will fail until ownership is reconciled.

### Finding 15: Compact artifact generation is implemented; website consumption remains planned

The production evaluator already derives and writes `website.json` unconditionally after `result.json`. The controlling completion plan still describes replacing mock website benchmark data, passing artifact points to Recharts, changing the chart to a numeric turn axis, deriving sample size, and publishing the finished website as partially completed future work.

**Evidence:** `cli/internal/evaluation/evaluation.go:304-310, 696-738`; `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:85-97`.

**Implication:** Research and implementation should treat the compact JSON producer as present production behavior and the React/Vite ingestion/publication path as planned work. Existing plans do not establish extra runtime fields, grouping guarantees, or import automation beyond what the production writers emit.

## Notes

- The inspected production files are modified in the current working tree relative to commit `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`; these findings describe the current checkout, not only the last committed CLI state.
- Tests were not used to infer contracts. They may be used later to validate a source-established implementation change.
- Existing generated outputs under `output/` use historical directory layouts and belong to a separate artifact-snapshot assignment; they were not treated as authoritative for the current writer path.
- The generic replay layer contains adapters for more harness identifiers than the lifecycle evaluation gate accepts. This is a useful search term for future support work, not evidence of current CLI evaluation support.

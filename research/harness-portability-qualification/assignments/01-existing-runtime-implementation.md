# Existing Qualified Runtime Implementation

## Assignment

**Goal:** Reconstruct the production implementation pattern for the currently evaluation-qualified OpenAI Codex, Cursor, Claude Code, and Pi runtimes, including setup, replay, attribution, process ownership, recovery, and user-state boundaries.

**Scope:** Repository production source for the Go CLI, plus `plans/harness-setup.md` and the retained smoke report as declared contract and bounded validation evidence.

**Exclusions:** Installation-only behavior for the five remaining harness candidates; external documentation; and tests as a source of intended behavior.

## Sources

- `cli/cmd/skill-issue/main.go` — process entry point.
- `cli/internal/command/command.go` — CLI routing, confirmation, and privacy warnings.
- `cli/internal/lifecycle/lifecycle.go` — request parsing, signal-aware run context, cleanup command, and output-owned state root.
- `cli/internal/harness/harness.go` — supported evaluation IDs, native project skill roots, and effective defaults.
- `cli/internal/evaluation/evaluation.go` — preparation, run lifecycle, evidence, Codex attribution, and recovery cleanup.
- `cli/internal/evaluation/runtime.go` — per-harness runtime construction and controlled environment helpers.
- `cli/internal/installer/installer.go` — temporary-skill instrumentation, collision handling, materialization, and restoration.
- `cli/internal/replay/replay.go`, `cli/internal/replay/process.go`, `cli/internal/replay/process_group_unix.go`, and `cli/internal/replay/process_group_windows.go` — session adapter construction, structured output checks, process control, and authentication probe.
- `cli/internal/replay/pi.go` — Pi RPC lifecycle and runtime-state validation.
- `cli/internal/runstate/runstate.go` and `cli/internal/evaluation/transcript.go` — private state, attribution persistence, and optional transcript sanitization.
- `cli/README.md` — product contract; especially evaluation boundary and recovery sections.
- `plans/harness-setup.md` — declared native setup contract, version qualification statements, and smoke scope; not implementation authority.
- `evaluations/skill-calling/smoke/real-harness-smoke-report.md` — retained two-turn local validation evidence; not a production behavior authority.

## Findings

### One Governed Evaluation Lifecycle

`command.App.Run` routes `evaluate run` through `lifecycle.ExecuteEvaluationRun`; the lifecycle parser accepts only the four IDs admitted by `harness.ParseEvaluationID`, resolves request defaults before the interactive review, and runs the service under a context canceled by `SIGINT` or `SIGTERM`. `evaluation.Service.Run` canonicalizes the workspace through `EvalSymlinks`, requires an existing directory, requires an output root outside it, creates a restrictive unique output directory, and stores ephemeral run state below `<output>/.skill-issue`.

**Evidence:** `cli/internal/command/command.go:45-77`; `cli/internal/lifecycle/lifecycle.go:126-154`, `170-252`; `cli/internal/harness/harness.go:57-88`; `cli/internal/evaluation/evaluation.go:154-224`, `609-641`; `cli/README.md:15-17`.

**Implication:** A portable harness must enter this lifecycle as a project-local evaluation and preserve the outside-workspace output and state boundary. A mere install specification is insufficient for evaluation support.

### Input, Model, And Evidence Invariants

The runner accepts either one embedded evaluation or the complete custom triplet. Custom answer sheets are rejected when symlink-resolved inside the evaluated workspace; schemas, scenario IDs, turn IDs, and supplied skill names are checked before side effects. Effective defaults are Codex `gpt-5.6-sol`/`medium`, Cursor `auto`/`medium`, Claude `opus`/`medium`, and Pi `openai-codex/gpt-5.6-sol`/`medium`; an independent Cursor reasoning override fails before runtime creation. Result derivation treats missing expected calls as scored evaluation data only after a successful tooling run, while harness/protocol errors abort the run as tooling failures.

**Evidence:** `cli/internal/evaluation/evaluation.go:515-568`, `590-641`, `672-701`, `704-748`; `cli/internal/harness/harness.go:57-62`; `cli/internal/evaluation/evaluation.go:56-74`; `cli/internal/lifecycle/lifecycle.go:170-224`; `cli/README.md:19-40`, `73-89`.

**Implication:** Qualification needs a native model/reasoning mapping, early rejection of unsupported override shapes, validated structured inputs, and a reliable distinction between model outcomes and execution failures.

### Disposable Skill Materialization And Marker Attribution

Before adapter creation, the service creates a cryptographic opaque token per skill and calls `installer.PrepareEvaluation`. Instrumentation copies each selected skill, inserts `Run "<absolute-cli>" signal "<token>" "<absolute-state-root>", then continue normally.` immediately after frontmatter, and atomically materializes it at the harness-specific evaluation root. The marker command maps the token to the active turn under a run lock and appends a synced private event. Runner boundaries set the active turn before sending the verbatim prompt and clear it after capture; markers outside that window stay explicitly unattributed.

Codex is special: after a capture, `recordCodexSignals` searches Codex structured `command_execution` events for the CLI path, `signal`, token, and state root, then records the marker itself. The other three adapters rely on the instrumented command being allowed and executing against the private state root.

**Evidence:** `cli/internal/evaluation/evaluation.go:185-210`, `233-284`, `290-308`, `374-408`; `cli/internal/installer/installer.go:77-117`, `267-312`, `314-342`; `cli/internal/lifecycle/lifecycle.go:157-167`; `cli/internal/runstate/runstate.go:125-191`; `cli/internal/replay/replay.go:120-149`; `cli/README.md:123-135`.

**Implication:** A candidate needs an exact, noninteractive route for the generated marker or equivalent structured event evidence. Session reuse alone cannot establish turn-attributed skill invocation.

### Skill Collision And User-State Boundary

Evaluation materialization checks target directories before replacing them. A collision with a name outside the current canonical Skill Issue payload fails. For a pre-existing canonical Skill Issue directory, cleanup rematerializes the current canonical embedded copy; paths introduced by the run are removed. Installation state is persisted before replay so `evaluate cleanup` can perform the same restoration after interruption. The implementation deliberately has no backup or rollback inventory.

**Evidence:** `cli/internal/installer/installer.go:98-117`, `120-167`, `170-192`; `cli/internal/evaluation/evaluation.go:255-267`, `411-477`; `cli/internal/lifecycle/lifecycle.go:97-123`; `cli/README.md:42-46`, `184-192`.

**Implication:** Portability requires a native temporary-skill location with deterministic cleanup. It also has a concrete limitation: user modifications inside an already present canonical Skill Issue skill directory are replaced by the current canonical payload rather than restored byte-for-byte; qualification cannot describe that directory as generally user-preserved.

### Codex Runtime Pattern

Codex keeps the normal process environment and user authentication location. Its runtime uses the evaluated workspace as both working directory and `.agents/skills` evaluation root, discovers ambient `SKILL.md` files below the configured Codex home, the user `.agents/skills` root, and `/etc/codex/skills`, and emits a temporary `skills.config` deny-list for them. For every prompt it launches `codex exec` or `codex exec resume` with `--ignore-user-config`, `--ignore-rules`, JSON output, `--cd` to the workspace, `--ask-for-approval on-request`, `--sandbox workspace-write`, `--disable plugins`, effective model, auto-review, reasoning, project-document suppression, and apps disabled.

The generic process adapter requires the initial `thread.started`, every-turn `turn.completed`, a stable discovered session ID, and no structured error before accepting a capture. `CheckAuthentication` runs `codex login status` before temporary skill installation. Codex process commands use their own Unix process group and group kill on completion/cancellation.

**Evidence:** `cli/internal/evaluation/runtime.go:28-43`, `300-367`; `cli/internal/evaluation/evaluation.go:225-231`, `269-285`; `cli/internal/replay/process.go:136-223`, `226-272`, `340-369`, `498-526`; `cli/internal/replay/process_group_unix.go:11-30`.

**Implication:** Codex qualification depends on normal Codex authentication, discoverable and suppressible ambient skills, resumable JSON session output, a workspace-write sandbox, and structured command events for marker attribution. The source does not create or replace a Codex home, credentials, or session store.

### Cursor Runtime Pattern

Cursor receives a private runtime containing a clean `HOME`, config/data directories, agent store, temporary directory, and plugin directory with generated plugin metadata and skills. The only retained user-owned integration is a macOS Keychain symlink. Its process environment is clean—only selected base values, the absolute executable directory, and `CURSOR_CONFIG_DIR`, `CURSOR_DATA_DIR`, and `CURSOR_AGENT_STORE_DIR` are supplied. Runtime config allowlists workspace reads/writes and only the absolute Skill Issue marker executable, denies destructive/sensitive patterns, enables sandboxing and allowlist review, and disables notifications, hints, model slash commands, and web-search auto-acceptance.

Each turn runs the resolved `agent`/`cursor-agent` with project configs disabled, the explicit plugin directory, workspace, sandbox, auto-review, streaming JSON, and a captured `--resume` session ID after the first turn. The adapter accepts only a stream containing `system/init` and successful `result` events, plus a stable session ID. Authentication probes `status` under the isolated environment before installation.

**Evidence:** `cli/internal/evaluation/runtime.go:66-123`, `220-293`; `cli/internal/evaluation/evaluation.go:225-231`, `269-285`; `cli/internal/replay/process.go:68-74`, `143-145`, `200-223`, `252-260`, `300-310`, `340-369`.

**Implication:** Cursor portability includes a native clean-state override, explicit skill/plugin loading, session-store isolation, a credential-preserving native authentication bridge, and a marker command that is permitted without broadly authorizing shell access. The Keychain path makes this exact production implementation macOS-shaped.

### Claude Code Runtime Pattern

Claude creates a private launch directory and a separate passed-skills root, places generated skills under `passed-skills/.claude/skills`, and launches from the empty launch directory rather than the evaluated workspace. Inline settings disable memory and background features and grant the real workspace as an additional directory. On each prompt, Claude runs print mode with project-only settings, strict MCP config, Chrome disabled, only the supplied skills root passed with `--add-dir`, selected editing tools, `Bash` allowlisted solely for the exact marker command, `dontAsk`, an appended workspace-routing prompt, effective model/effort, and verbose stream JSON. A generated session ID is supplied on the first call and reused with `--resume` thereafter.

The generic protocol gate requires Claude `system/init` and `result` event types and a stable session identifier; it does not enforce the more detailed event-content checks stated in the setup plan. Unlike Cursor and Pi, the source provides no controlled environment for Claude: `runtime.environment` remains empty and the adapter merges the inherited host environment. On close, the adapter runs `claude project purge --yes <launch-directory>` with the same inherited environment after terminating the owned process group.

**Evidence:** `cli/internal/evaluation/runtime.go:126-157`; `cli/internal/evaluation/evaluation.go:269-285`; `cli/internal/replay/process.go:93-108`, `114-123`, `143-151`, `256-260`, `283-321`, `391-420`; `cli/internal/replay/process_group_unix.go:11-30`; `plans/harness-setup.md:457-489` (contract distinction).

**Implication:** A portable Claude-like candidate needs a safe way to expose the real workspace without treating it as configuration discovery, an explicit supplied-skill root, restricted marker execution, resumable stream JSON, and a harness-native cleanup command for its launch/project state. The current implementation preserves ordinary configuration files by avoiding their paths, but inherited environment variables remain an ambient input.

### Pi Runtime Pattern

Pi creates private `home`, `sessions`, and `passed-skills` directories while deliberately retaining `PI_CODING_AGENT_DIR` from its configured location or native default. It invokes Pi in one long-lived RPC process with a generated session ID, `--no-session`, provider/model split from the effective model, `--thinking`, offline mode, all ambient extension/skill/template/theme/context discovery disabled, explicit generated `--skill` directories, and a fixed tool list including `bash`. Its controlled clean environment preserves only recognized credential variables, `VOLTA_HOME` when present, native agent directory, private sessions, and offline mode.

Pi preflights `get_state` (generated ID, model/provider, thinking level, idle/noncompacting/no-pending/no-session-file) and `get_commands` (every generated skill present). Each prompt is an RPC JSON object with a unique ID; the session waits for both a successful matching response and `agent_settled`, rejects error/aborted agent terminal events, then revalidates state. Conversation continuity comes from keeping that one process alive rather than issuing a resume command. Close sends `abort` if pending, closes stdin, waits two seconds, then kills the direct Pi process if needed.

**Evidence:** `cli/internal/evaluation/runtime.go:160-196`, `244-293`; `cli/internal/evaluation/evaluation.go:269-285`; `cli/internal/replay/pi.go:30-104`, `160-231`, `262-323`, `333-403`; `plans/harness-setup.md:536-573` (contract distinction).

**Implication:** Pi qualification needs an explicit-skill interface that survives disabled discovery, machine-readable lifecycle/state inspection, persistent in-memory conversation, usable native credentials without copying an auth directory, and an acknowledged host-access boundary. Source does not call `CheckAuthentication` for Pi—the helper returns immediately for every harness except Codex and Cursor—so Pi authentication is an operational prerequisite rather than a current pre-install verification gate.

### Structured Output Is A Fail-Closed Contract

For process adapters, stdout must be a nonempty JSON object array or JSON-lines stream. Structured error event types fail immediately. Codex, Cursor, and Claude then require their harness-specific minimum terminal events, first-capture session extraction, and later session stability. Pi independently parses every RPC line and validates both preflight and post-turn state. Captures preserve raw transcript, stderr, and events; optional persistence sanitizes known paths and local identity values before writing `transcript.json`, while `result.json` and `website.json` are always written only after successful replay.

**Evidence:** `cli/internal/replay/process.go:189-272`, `422-495`; `cli/internal/replay/pi.go:160-231`, `262-285`, `355-403`; `cli/internal/evaluation/evaluation.go:314-371`, `825-849`; `cli/internal/evaluation/transcript.go:38-72`, `128-187`.

**Implication:** Candidate support needs stable, parseable output that establishes successful turn completion and one continuous session. The minimum generic checks are intentionally shallow for Cursor and Claude; stronger plan-level assertions should be treated as a qualification requirement to implement, rather than as behavior already enforced by the source.

### Cancellation, Process Ownership, And Recovery

The lifecycle context is canceled on interrupt/termination. Generic process commands are started in a separate Unix process group, and `exec.CommandContext` cancellation plus explicit group kill cleans descendants on wait and close; Claude purge follows close. Windows currently kills only the direct process. Pi does not use `configureOwnedProcess`; its close path owns the direct RPC process and may leave descendants outside this helper's explicit control. Service defers installation cleanup after successful materialization, deletes runtime roots, private token mappings, and run records on successful cleanup, and provides an explicit cleanup command using persisted installation state if an interruption leaves it behind. Errors before installation delete private mappings/run state; errors after adapter/run setup also attempt cleanup.

**Evidence:** `cli/internal/lifecycle/lifecycle.go:147-154`; `cli/internal/replay/process.go:175-186`, `283-297`; `cli/internal/replay/process_group_unix.go:11-30`; `cli/internal/replay/process_group_windows.go:7-13`; `cli/internal/replay/pi.go:300-323`; `cli/internal/evaluation/evaluation.go:218-254`, `367-477`; `cli/internal/runstate/runstate.go:216-337`.

**Implication:** A qualified candidate must make cancellation and cleanup prove ownership of every spawned process or explicitly record the remaining limitation. The established source pattern is strongest for Unix generic adapters; it is not evidence of equivalent descendant cleanup for Pi or Windows.

### Bounded Smoke Evidence And Contract Gaps

The retained smoke report states that all four harnesses completed built-in and custom two-turn routes on 2026-07-20, with launcher versions/authentication routes and cleanup observations recorded. The setup plan repeats that qualification scope and states that no governed 30-turn campaign ran. This validates the narrow runtime probes only; it cannot establish 30-turn reliability or upgrade the documentation's desired runtime assertions into implementation facts.

**Evidence:** `evaluations/skill-calling/smoke/real-harness-smoke-report.md:5-26`, `61-82`; `plans/harness-setup.md:634-640`; `cli/README.md:200-216`.

**Implication:** New-harness qualification should retain separate evidence for default and explicit configuration routes, session continuity, expected and missing-call handling, workspace effects, private-state cleanup, and child-process cleanup, then distinguish those probes from governed campaign evidence.

## Notes

- Source truth supports evaluations only for `codex`, `cursor`, `claude-code`, and `pi`; the other IDs in `harness.Spec` are installation specifications or unqualified replay stubs, outside this assignment.
- `plans/harness-setup.md` asks for stronger Claude and Cursor event-content validation than `validateHarnessOutput` currently implements. The plan is a contract/reference, while `cli/internal/replay/process.go` is the executable enforcement point.
- The source accepts any existing directory as an evaluation workspace. The setup plan's initialized-Git-workspace statement is a qualification contract but is not checked by `evaluation.Service.Run`.
- `replay.CheckAuthentication` receives a Pi-related boolean at `evaluation.Service.Run`, but its implementation only probes Codex and Cursor; the boolean does not alter the early return.
- No external research was performed for this assignment.

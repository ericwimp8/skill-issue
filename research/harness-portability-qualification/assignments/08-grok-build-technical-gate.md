# Grok Build First Technical Qualification Gate

## Assignment

**Goal:** Audit the official Grok Build coding-agent CLI for the first technical qualification gate: clean/default isolation, temporary explicit skills, noninteractive replay, session continuity, structured evidence, model/reasoning, permissions and sandboxing, authentication, cancellation, and temporary-state ownership.

**Scope:** Current first-party Grok Build documentation and the public `xai-org/grok-build` source, cross-checked against Skill Issue's current production boundary and candidate map. Research date: 2026-07-21.

**Exclusions:** Installing, authenticating, or running `grok`; Grok Chat, API-only, IDE, and unofficial CLI behavior; Codex-subscription viability; direct-install support; and a final product-support decision.

## Sources

- [xAI Grok Build overview](https://docs.x.ai/build/overview) — product identity, `grok` install/authentication, headless mode, and custom-model configuration.
- [xAI CLI reference](https://docs.x.ai/build/cli/reference) — official `grok` executable, headless/session/model/effort/permission/sandbox flags, `inspect`, and `agent stdio`.
- [xAI Headless & Scripting guide](https://docs.x.ai/build/cli/headless-scripting) — noninteractive protocol, JSON and JSONL completion/error forms, session storage, resume, ACP, and update suppression.
- [xAI Settings](https://docs.x.ai/build/settings), [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces), and [AGENTS.md](https://docs.x.ai/build/features/project-rules) — configuration layers and ambient discovery surfaces.
- [xAI Enterprise Deployments](https://docs.x.ai/build/enterprise) and [xAI Sandbox guide](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/18-sandbox.md) — authentication, permission modes/rules, sandbox profiles, and platform limits.
- [xAI Sessions](https://docs.x.ai/build/features/sessions) and [xAI Background Tasks source guide](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/20-background-tasks.md) — persisted sessions, background-task controls, and kill semantics.
- [xAI Grok Build repository](https://github.com/xai-org/grok-build), `SOURCE_REV` `ba69d70c2f7d70a130a323b2becdf137af784c7f` — first-party Rust CLI/runtime source. Inspected: `crates/codegen/xai-grok-config/src/paths.rs`, `crates/codegen/xai-grok-config/src/loader.rs`, `crates/codegen/xai-grok-pager/docs/user-guide/14-headless-mode.md`, `18-sandbox.md`, `20-background-tasks.md`, and `crates/codegen/xai-grok-shell/src/session/acp_session_impl/tasks_cancel.rs`.
- Local production boundary: `cli/internal/harness/harness.go`, `cli/internal/evaluation/runtime.go`, and `cli/internal/replay/process.go` — current Grok registration/replay stub and unsupported evaluator path.
- Local qualification framing: `research/harness-portability-qualification/assignments/01-existing-runtime-implementation.md`, `02-product-support-and-setup-contract.md`, and `04-local-candidate-research-grok-opencode-kilo.md`.

## Findings

### Official product identity and current local evaluator status

Grok Build is the first-party terminal coding agent and CLI published by xAI/SpaceXAI. Its released command is `grok`; the public source tree identifies its installed artifact as `grok` and its source-built artifact as `xai-grok-pager`. xAI documents interactive TUI, headless, and ACP operation. In Skill Issue, `HarnessGrok` supplies a direct-install root and an unreachable generic replay stub, while runtime preparation rejects it before authentication, materialization, or replay. The prior local read-only probe found no `grok` on `PATH`.

**Evidence:** The [xAI overview](https://docs.x.ai/build/overview) and [CLI reference](https://docs.x.ai/build/cli/reference) describe the official product and executable; the [official repository README](https://github.com/xai-org/grok-build) identifies the released/source-built names. Local `cli/internal/harness/harness.go`, `cli/internal/evaluation/runtime.go`, and assignment 04 record the blocked route and executable absence.

**Implication:** The product/executable identity is established, but no local executable, authentication, activation, or smoke evidence exists. The current Skill Issue evaluator cannot use Grok Build without a new source-owned route.

### Ambient discovery prevents a clean default-isolation claim

Grok Build has a documented `$GROK_HOME` override, so a runner can redirect its user-level config, credentials, sessions, logs, skills, personas, and related state to a disposable directory. It nevertheless discovers project configuration and content from the selected working tree: `.grok/config.toml`, `.grok/skills`, `.grok/plugins`, project hooks, rules/`AGENTS.md`, Claude-compatible instructions/rules/settings/skills/plugins/MCPs, and `~/.agents` skills/commands. System managed and requirements layers under `/etc/grok` also remain separate from `$GROK_HOME`; source confirms that system layers are loaded in addition to the user-home layers. The documented headless flags offer feature suppression and prompt override, but no documented equivalent to ignoring all user/project/system configuration and discovery.

**Evidence:** [Settings](https://docs.x.ai/build/settings) defines `$GROK_HOME`, user/project/managed/requirements scopes; [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces) and [AGENTS.md](https://docs.x.ai/build/features/project-rules) enumerate Grok, Claude-compatible, and Agents.md discovery. First-party `paths.rs` resolves `$GROK_HOME` or `~/.grok` and independently exposes `/etc/grok`; `loader.rs` merges system managed, user managed, user, user requirements, and system requirements. The [CLI reference](https://docs.x.ai/build/cli/reference) lists `--no-plan`, `--no-subagents`, `--no-memory`, and `--disable-web-search`, but no clean-config/discovery-exclusion switch.

**Implication:** A disposable `$GROK_HOME` can isolate run-owned user state, but it cannot establish the required clean/default isolation from workspace instructions, project skills/plugins/rules, compatibility inputs, or system policy. Treat this condition as **unsupported/fail-closed** rather than relying on a temporary home or prompt override as a substitute.

### Temporary explicit skills can be run-owned, but cannot be proven exclusive

The native user skill root is `$GROK_HOME/skills`; Grok also discovers workspace `.grok/skills`, enabled plugin skills, configured extra paths, Claude-compatible skills, and `~/.agents/skills`. A runner could place an instrumented Skill Issue skill under a fresh `$GROK_HOME/skills`, use an API-key authentication boundary, and delete that home after the run. `grok inspect --json` is the official discovery inspection command. The documented inspector covers discovered rules, skills, plugins, hooks, and MCP servers, but public docs do not define a headless output event that proves a particular skill invoked a marker command.

**Evidence:** [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces) identifies every skill source; [CLI reference](https://docs.x.ai/build/cli/reference) documents `grok inspect [--json]`; [Headless & Scripting](https://docs.x.ai/build/cli/headless-scripting) lists `skills/` and other owned paths beneath `$GROK_HOME`. The headless protocol documents text, thought, end, error, and metadata events, not a skill-invocation event. Assignment 04 records the same `inspect --json` check as an unimplemented local fail-closed requirement.

**Implication:** A temporary native skill placement and discovery check are technically available, while exclusive discovery and opaque marker attribution remain unproven. The first gate cannot accept a generated skill as isolated/activated until a live `inspect --json` and marker probe show the complete discovered set and turn-attributed signal.

### Headless replay, resumable sessions, and protocol evidence are documented

`grok -p` runs noninteractively and exits after its response. `--output-format json` emits one completion object containing `sessionId`, `stopReason`, request ID, turns, and available usage; `streaming-json` emits newline-delimited `text`, `thought`, `end`, and `error` events, with `end` documented as terminal. Failures emit an error object and nonzero exit; documented exit codes distinguish success, normal error, SIGINT, and SIGTERM. `--resume <id>` resumes a durable headless session, while `--session-id <UUID>` creates a new named session. The same official documentation also describes ACP JSON-RPC over stdio, including explicit JSON-RPC errors and `session/prompt` completion metadata.

**Evidence:** [Headless & Scripting](https://docs.x.ai/build/cli/headless-scripting) documents `-p`, JSON/JSONL schemas, error objects, exit codes, `--resume`, and ACP. [Sessions](https://docs.x.ai/build/features/sessions) states that headless sessions persist under `~/.grok/sessions` (or `$GROK_HOME`) and shows extracting `.sessionId` for a resume.

**Implication:** Grok Build has a viable documented noninteractive and resumable protocol surface for a later adapter. A candidate adapter should fail closed on nonzero exit, JSON `type: error`, missing final `end`/completion metadata, missing `sessionId`, or session-ID mismatch. Protocol evidence alone cannot cure the missing isolation and activation proofs.

### Model, effort, permissions, and sandbox controls are native and parameterized

The CLI supports `--model` and `--effort`/`--reasoning-effort`; the first-party headless guide lists canonical effort values from `none` through `xhigh` (with `max` an alias). It supports tool filtering, `--allow`/`--deny` rules, `--permission-mode`, `--sandbox`, and session flags that disable plans, subagents, memory, or web search. The default permission experience is asking; unattended execution can use `dontAsk` with narrow allow rules or an always-approve mode, but the latter broadly authorizes tool calls. Sandboxing is off by default; `workspace`, `read-only`, and `strict` profiles impose OS-level restrictions, while macOS child-network restriction is explicitly a no-op.

**Evidence:** [CLI reference](https://docs.x.ai/build/cli/reference), [Headless & Scripting](https://docs.x.ai/build/cli/headless-scripting), [Enterprise Deployments](https://docs.x.ai/build/enterprise), and the first-party [sandbox guide](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/18-sandbox.md) document these controls and their scope. The sandbox guide also states that a session's sandbox profile is fixed at creation/resume.

**Implication:** A future adapter can pass a fixed Grok model/effort and construct narrow marker permissions plus a fixed sandbox profile. It must record the effective profile and platform, reject permission prompts in noninteractive runs, and avoid `--always-approve`/`--yolo` as a qualification shortcut. No source or probe here establishes a supported model/access cell.

### Authentication can stay outside the temporary runtime, without asserting Codex access

Official headless authentication supports `XAI_API_KEY`, device-code login, browser OIDC, external auth-provider configuration, and cached local credentials. Credential resolution includes active session tokens and configuration-backed keys, whereas an API key can be supplied as an environment variable. A fresh `$GROK_HOME` with `XAI_API_KEY` therefore supplies a documented way to avoid copying a user credential home into the temporary runtime. This assignment neither tests authentication nor makes any claim about whether Grok Build can access a Codex subscription.

**Evidence:** [Overview](https://docs.x.ai/build/overview), [Headless & Scripting](https://docs.x.ai/build/cli/headless-scripting), and [Enterprise Deployments](https://docs.x.ai/build/enterprise) enumerate these authentication paths and their precedence.

**Implication:** Treat a normal xAI API key or native login as a separate preflight boundary. A later live qualification must prove the selected authentication and model access without copying credentials or making a permanent configuration change; Codex-subscription access remains outside this gate.

### Cancellation and child cleanup are insufficiently bounded for qualification

For a headless process, SIGINT/SIGTERM persist completed-tool session state, do not roll back file modifications, and produce documented exit codes. Grok also supports subagents and background commands. The background-task guide supplies explicit per-task kill behavior (SIGTERM then SIGKILL for shell processes; Cancel and Shutdown for subagents), while first-party cancellation source shows cancellation modes where background tasks and subagents can survive a running-turn cancellation. The background-task runtime persists a manifest when running tasks are intentionally left alive on session shutdown. No official source inspected proves that interrupting the parent headless `grok` process deterministically terminates all locally spawned/background work and every remote child session.

**Evidence:** [Headless & Scripting](https://docs.x.ai/build/cli/headless-scripting) documents interrupt persistence/no rollback and exit codes; [Background Tasks source guide](https://github.com/xai-org/grok-build/blob/main/crates/codegen/xai-grok-pager/docs/user-guide/20-background-tasks.md) defines task kill semantics. First-party `tasks_cancel.rs` distinguishes cancellation options for subagents/background tasks, and `terminal/background_task.rs` persists a manifest for tasks intentionally alive at shutdown.

**Implication:** Parent-process cancellation evidence is available, while child/background cleanup is **unsupported** for the first gate. A future adapter must disable subagents where possible, constrain task-spawning tools, own its process group, and demonstrate no surviving Grok background/subagent state before this condition can pass.

### Temporary-state ownership is workable only as a scoped run design

With `$GROK_HOME` set to a runner-created directory, Grok's user config, session database, logs, skills, personas, crashes, trace exports, and worktree metadata become run-owned. `--no-auto-update` suppresses session update checks. A fixed session ID or captured response `sessionId` keeps multi-turn state inside that temporary home, which can be removed after the runner closes. System configuration, workspace files, and any native background process or remote child state remain outside that deletion boundary.

**Evidence:** [Headless & Scripting](https://docs.x.ai/build/cli/headless-scripting) lists the `$GROK_HOME` contents and update-suppression flag; first-party `paths.rs` constructs sessions under `grok_home()`; [Settings](https://docs.x.ai/build/settings) separately defines system-managed layers.

**Implication:** A later implementation can own and clean a dedicated Grok runtime root, but it must report system/project inputs and prove child cleanup separately. Deleting `$GROK_HOME` must never be represented as restoration of user/project configuration or arbitrary agent-created workspace changes.

### First-gate outcome

Official sources establish product identity, headless replay, structured JSON/JSONL failures and completion, resumable sessions, native model/effort selection, API-key authentication, permissions, sandbox profiles, and a run-owned `$GROK_HOME` design. They also establish that Grok intentionally discovers broad ambient user/project/compatibility configuration and can retain background/subagent work beyond some cancellation paths. No local executable or live run has verified the potential mitigations.

**Evidence:** The official sources above, plus local assignments 01, 02, and 04 and current Skill Issue evaluator source.

**Implication:** **Do not advance Grok Build through the first technical qualification gate.** Mark clean/default isolation, exclusive temporary-skill activation/attribution, and bounded child cleanup as unsupported blockers. Preserve the documented headless/session/protocol surface as the source-backed basis for a later, tightly scoped live qualification probe.

## Notes

- No installation, authentication, or model call was performed. The local absence of `grok` is a read-only 2026-07-21 observation, not a product-wide availability claim.
- `--system-prompt-override`, feature-disable flags, `$GROK_HOME`, or a temporary sandbox profile are not evidence of complete discovery exclusion; official docs do not specify such a guarantee.
- The official source tree is periodically synced from xAI's monorepo. Findings about source behavior are pinned to the recorded `SOURCE_REV`, while product docs may evolve independently.

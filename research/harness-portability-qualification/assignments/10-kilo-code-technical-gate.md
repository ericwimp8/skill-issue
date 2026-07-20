# Kilo Code Technical Qualification Gate

## Assignment

**Goal:** Assess current Kilo Code CLI against the first technical portability gate: a clean, noninteractive, resumable, evidence-producing Skill Issue evaluation route that leaves only bounded evaluator-owned state.

**Scope:** Current first-party Kilo CLI documentation, the Kilo-Org/kilocode release repository at commit `0cf9f902e0eaead2746a935e9440d36b84df45bf` (2026-07-20), official releases, and the repository's existing Kilo candidate evidence. This covers executable identity and access; isolation from configuration, instructions, skills, plugins, MCP, and state; temporary explicit skills; headless replay; session and evidence semantics; model, reasoning, permissions, authentication; cancellation; and cleanup.

**Exclusions:** Installing Kilo, authenticating, executing a model prompt, modifying Kilo configuration, Kilo Cloud or IDE-only behavior, Codex-subscription viability, and a production adapter design.

## Sources

- [Kilo Code CLI documentation](https://kilo.ai/docs/code-with-ai/platforms/cli) — installation, `kilo run`, config locations, autonomous behavior, permissions, provider setup, and environment override documentation; inspected 2026-07-21.
- [Kilo CLI command reference](https://kilo.ai/docs/code-with-ai/platforms/cli-reference) — documented `kilo run`, `debug config`, `debug skill`, `debug paths`, and `auth` commands; inspected 2026-07-21.
- [Kilo Skills documentation](https://kilo.ai/docs/customize/skills) — discovery locations, explicit named invocation, session-bound scanning, `/reload`, and `skill` tool evidence; inspected 2026-07-21.
- [Kilo permissions documentation](https://kilo.ai/docs/getting-started/settings/auto-approving-actions) — allow/ask/deny model and `skill`, `task`, `bash`, and external-directory permission classes; inspected 2026-07-21.
- [Kilo-Org/kilocode repository](https://github.com/Kilo-Org/kilocode) and [current inspected commit](https://github.com/Kilo-Org/kilocode/commit/0cf9f902e0eaead2746a935e9440d36b84df45bf) — first-party CLI source and release lineage. The repository README identifies `@kilocode/cli` and the `kilo` executable.
- [Run command source](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/cli/cmd/run.ts) — noninteractive defaults, `--session`, `--format json`, `--model`, `--variant`, permission handling, event projection, and automatic exit.
- [CLI entry and shutdown source](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/index.ts) and [Kilo CLI setup source](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/kilocode/cli/setup.ts) — normal finalization, instance disposal, and forced process exit.
- [Configuration loader](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/config/config.ts), [configuration paths](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/config/paths.ts), and [global paths](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/core/src/global.ts) — config precedence, XDG-backed roots, project/global discovery, managed configuration, and generated state locations.
- [Skill discovery](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/skill/index.ts) and [skill tool](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/tool/skill.ts) — native and compatibility skill scanning plus the permission-gated `skill` tool.
- [Authentication source](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/auth/index.ts) and [Kilo-provider credential source](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/core/src/credential.ts) — `auth.json`, `KILO_AUTH_CONTENT`, and process-local credential precedence.
- [Persistent-process ownership source](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/kilocode/background-process/index.ts) and [server signal handling](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/cli/cmd/serve.ts) — owned background process groups and `kilo serve` signal behavior.
- Local context: `research/harness-portability-qualification/assignments/01-existing-runtime-implementation.md`, `02-product-support-and-setup-contract.md`, and `04-local-candidate-research-grok-opencode-kilo.md`. The current Skill Issue evaluator rejects Kilo before runtime preparation; a read-only local probe on 2026-07-21 found no `kilo` executable in `PATH`.

## Findings

### Gate outcome: Kilo is not technically qualified

Kilo has a current first-party terminal product with documented headless, session, JSON, skills, permissions, and model controls. It is therefore a credible future candidate. It does not yet satisfy this gate because no local executable or authenticated run exists, the existing evaluator cannot reach it, normal `kilo run` reads broad ambient state and may attach to a running daemon, and no source-backed cancellation/descendant-cleanup proof exists for an interrupted headless run. The strict clean-isolation path depends on temporary XDG/home roots plus process-injected configuration and credentials; it has not been demonstrated against the actual release binary.

**Evidence:** Kilo documents `npm install -g @kilocode/cli`, `kilo --version`, and `kilo run`; the repository identifies Kilo CLI as the `kilo` executable. `04-local-candidate-research-grok-opencode-kilo.md` records the absent local executable and the Skill Issue runtime gate. The Kilo `run.ts`, `config.ts`, `paths.ts`, `global.ts`, and `index.ts` sources establish the remaining behavior.

**Implication:** Classify Kilo as **unqualified**, rather than as a supported or smoke-qualified evaluator route. The qualified facts below narrow a later, controlled live probe; they do not authorize implementation or a permanent Kilo workaround.

### Product identity and executable access are documented, but absent locally

The relevant product is the Kilo Code CLI 1.0-or-later surface, distributed as `@kilocode/cli` and invoked as `kilo`; its source lives in `Kilo-Org/kilocode`, where `packages/opencode/` is the shared CLI runtime. The documentation names `kilo --version` as the install verification command, and the official release repository publishes platform binaries as well as source releases. The local candidate audit found no `kilo` in `PATH`, so neither the planned Skill Issue stub nor the current docs establish access to a runnable local version.

**Evidence:** The Kilo CLI documentation and repository README identify the npm package and executable. The inspected repository commit is a first-party source snapshot; it defines `scriptName("kilo")`. Local absence is validation evidence from assignment 04, not a product-capability claim.

**Implication:** A later route must resolve and record one exact executable and version before temporary materialization. There is no basis to treat the registry's generic `kilo run` replay stub as executable validation.

### Headless multi-turn replay has a source-backed command shape

`kilo run <message>` is noninteractive unless `--interactive` is selected. It accepts `--format json`, `--session <id>`, `--continue`, `--model provider/model`, and `--variant`; the source creates a new session when neither continuation option is supplied, or retrieves the requested session and sends the next prompt to it. In JSON mode, every projected event carries the selected `sessionID`; the command emits structured `tool_use`, `step_start`, `step_finish`, `text`, `reasoning`, and `error` records, and exits after the subscribed session becomes idle. This is sufficient **source evidence** for a one-session, separate-process two-turn command form: capture the first run's `sessionID`, then supply it to the next `kilo run --session` call.

**Evidence:** `packages/opencode/src/cli/cmd/run.ts` declares the flags, calls `sdk.session.create` or `.get`, subscribes to the event stream, emits JSON-lines records, and breaks on `session.status: idle`. The command reference independently lists JSON format, session continuation, and model/variant options.

**Implication:** Resumable session continuity is a qualified capability candidate. A live route must still prove that two actual turns preserve the same returned ID under the selected release and that a fresh isolated data root retains the session between process invocations.

### Structured evidence is useful but lacks a terminal success record

The JSON projection preserves a session identifier and exposes completed or errored tool parts, including a `skill` tool call. It emits `session.error` as an `error` record and sets a nonzero exit code for that error or for prompt/command submission failure. This can provide skill-activation and tool-failure evidence if the evaluator captures and interprets the full JSON-lines stream. The command does not emit a final JSON `success`, `completed`, or `idle` record: its internal loop sees idle and returns, leaving process exit status as the completion signal. A successful run with no projected parts can therefore provide no JSON success line.

**Evidence:** `run.ts` projects tool parts only when their state is `completed` or `error`, includes the whole part under `tool_use`, emits `error` for `session.error`, and exits on the internal idle status. `skill.ts` calls the permission-gated `skill` tool and reports `Loaded skill: <name>` in its tool result.

**Implication:** The Kilo JSON stream can support session, `skill`, and tool-error attribution, but a Kilo adapter needs a fail-closed terminal rule that combines stable session ID, expected `tool_use` evidence, absence of `error`, and exit status. That rule is an adapter requirement, not an existing Kilo success schema.

### Default discovery is broad; a private explicit-skill environment is source-supported but unproven

By default Kilo loads global config, project `kilo.json[c]`/legacy `opencode.json[c]`, `.kilo` and `.kilocode` config directories up the project tree, the user home `.kilo`/`.kilocode` directories, configured plugin/command/agent sources, and organization or managed configuration. Skill discovery independently scans native Kilo locations, `.agents` and optionally `.claude` compatibility locations, configured local paths, and remote skill URLs. This includes the evaluated workspace's `.agents/skills`, which is a direct conflict with Skill Issue's clean-default requirement.

The source exposes a controlled isolation **candidate**: redirect the XDG config/data/cache/state roots and home to an evaluator-owned temporary root, set `KILO_DISABLE_PROJECT_CONFIG`, set `KILO_DISABLE_EXTERNAL_SKILLS` and `KILO_DISABLE_CLAUDE_CODE_SKILLS`, disable default plugins, and place only the generated payload in the temporary Kilo skill/config root. The source confirms that a `KILO_CONFIG_DIR` is scanned for Kilo skills and that `KILO_CONFIG_CONTENT` and `KILO_PERMISSION` are process inputs. It does not establish this complete environment as a public stable contract, and managed organization/MDM configuration is still loaded later than local config. The documented `--pure` control is inadequate evidence by itself: its production use is confined to external TUI-plugin loading, while the docs describe it more broadly.

**Evidence:** `config.ts` loads global config before explicit config, project files, config directories, `KILO_CONFIG_CONTENT`, active-org config, and managed config; `paths.ts` includes global, project, home, and explicit directories. `skill/index.ts` implements native, external compatibility, configured path, and URL scans. `global.ts` derives data/config/cache/state from XDG roots. `KILO_DISABLE_PROJECT_CONFIG`, `KILO_DISABLE_EXTERNAL_SKILLS`, `KILO_DISABLE_CLAUDE_CODE_SKILLS`, `KILO_DISABLE_DEFAULT_PLUGINS`, `KILO_CONFIG_DIR`, `KILO_CONFIG_CONTENT`, and `KILO_PERMISSION` are production flags. The official skills documentation independently confirms those discovery classes.

**Implication:** Default Kilo invocation fails the isolation condition. A temporary-root route is plausible only with live proof that the generated skill is the sole discoverable skill, `debug config`/`debug skill` show no ambient configuration or skills, no external plugin/MCP executes, and managed policy does not reintroduce an uncontrolled input. Static source does not permit a clean-default claim.

### Permissions can be constrained per process, while autonomous approval is unsuitable

Kilo defines granular allow/ask/deny rules for `skill`, `bash`, `task`, external directories, and other tools. `kilo run` in ordinary noninteractive mode automatically denies interactive questions, interactive terminal use, plan-entry/exit requests, and permission prompts that are not explicitly allowed. `KILO_PERMISSION` is parsed as JSON and merged after configuration loading, so it is a source-backed per-process permission input. In contrast, `--auto` replies once to every root and tracked-subagent permission request, while `--dangerously-skip-permissions` also grants requests; both defeat a minimal marker-command permission profile.

**Evidence:** Kilo's permission documentation lists the permission model and applicable tool names. `run.ts` creates the noninteractive denials, rejects ordinary permission requests, and auto-replies under `--auto`/`--dangerously-skip-permissions`. `config.ts` merges `KILO_PERMISSION` after the other configuration sources. `tool/skill.ts` requests permission for the named skill before loading it.

**Implication:** A viable evaluator must pass a run-scoped allowlist that admits only the generated skill and exact marker operation, while retaining all other defaults as deny/ask. The exact permission-pattern grammar and whether it accepts the generated marker command remain **unsupported** until a non-destructive authenticated probe verifies them; broad autonomous flags are rejected for qualification.

### Model, reasoning, and authentication have distinct boundaries

Kilo accepts an explicit `--model provider/model` and a provider-specific `--variant`, whose CLI description gives reasoning effort as its example. The model and reasoning request shape is therefore supported source evidence, subject to provider availability. Kilo exposes `kilo auth list`, `login`, and `logout`; credentials ordinarily live in the Kilo data root. The runtime also accepts `KILO_API_KEY` for Kilo-provider access and `KILO_AUTH_CONTENT` as process-local credentials that override host storage. None of the inspected first-party sources establishes that the desired evaluation model is available to a particular account, that `auth list` has a stable machine-readable success contract, or that any Kilo route can access a Codex subscription.

**Evidence:** The CLI command reference lists `auth` subcommands. `run.ts` passes model and variant to the session. `auth/index.ts` loads `KILO_AUTH_CONTENT` before the data-root `auth.json`; the provider credential source reads `KILO_API_KEY` and treats injected credential content as process-local. Kilo's built-in-provider documentation describes Kilo-account/Gateway access, not Codex subscription access.

**Implication:** Record model plus variant as explicit run inputs, and keep authentication outside the temporary skill lifecycle. A private-root evaluation can use officially supported environment-provided credentials only when the operator supplies them; copied host auth directories, permanent login/config edits, and subscription claims are outside this gate.

### Kilo owns normal shutdown, but interrupted headless cleanup is not qualified

On a normal `kilo run` return, the CLI entry's `finally` runs session-export shutdown, registered shutdown tasks, and instance disposal, then explicitly exits. Kilo also contains an ownership-checked persistent-process manager that terminates owned Unix process groups or Windows process trees. These are useful production behaviors. They do not prove cancellation safety for an evaluator-launched `kilo run`: the CLI entry does not register SIGINT/SIGTERM handling for that command, while `kilo serve` separately does. In addition, `kilo run` first attempts to attach to an available Kilo daemon; that ambient daemon can own sessions and state outside a run's temporary process boundary.

**Evidence:** `index.ts` invokes `KiloCli.shutdown()` in `finally` and then `process.exit()`. `setup.ts` disposes all instances during normal shutdown. `background-process/index.ts` verifies ownership before signalling process groups. `serve.ts` registers signal handlers, whereas `run.ts` has no equivalent handler and calls `KiloRunDaemon.attach` before constructing its in-process client.

**Implication:** Normal completion cleanup is **source-backed**, while SIGINT/SIGTERM abort, session abort, daemon avoidance, and descendant cleanup for a headless evaluation are **unsupported**. This is a gate blocker, not a reason to assume the generic Skill Issue process-group cleanup covers Kilo-created descendants.

### Temporary Kilo state is redirectable in part, but bounded cleanup remains an evaluator obligation

Kilo creates data, config, state, cache, logs, and a Kilo temporary directory during startup. Its session database, auth file, remote-skill cache, snapshots, model state, plugin metadata, and background-process manifests are rooted in Kilo global paths. XDG redirection can bound the data/config/cache/state roots to an evaluator-owned directory; that supports cleanup by removing only that unique root after the child process and any owned descendants exit. Kilo's shared temporary root is derived from the operating-system temp directory, and source inspection does not prove that a headless run leaves no Kilo state beyond the redirected root.

**Evidence:** `global.ts` initializes the XDG-derived data/config/cache/state directories and `os.tmpdir()/kilo`; `storage/db.ts`, `auth/index.ts`, `skill/discovery.ts`, and `background-process/index.ts` place persistent state below the global paths. The normal shutdown sources dispose instances but do not remove all generated global data.

**Implication:** Temporary-state ownership is **partially source-backed** only when every XDG root is unique and evaluator-owned. Bounded cleanup requires live evidence of the actual paths touched and of an empty or safely removable private root after a normal run and after cancellation. No source supports deleting a normal user's Kilo home, data, cache, daemon, or session store.

## Notes

- **Candidate invocation shape, not a qualification recipe:** a future probe can use `kilo run --format json --model <provider/model> --variant <effort> <prompt>`, capture its `sessionID`, then repeat with `--session <id>`. It must use a private environment and an exact restrictive permission profile; this report does not assert that the command is safe to run with host defaults.
- **Unsupported current claims:** local Kilo executable/version; authenticated model access; two-turn session continuity; the JSON terminal-success rule; exclusive generated-skill discovery; skill-marker permission; tool/signal attribution in a live run; cancellation and descendant cleanup; touched-state inventory; and bounded cleanup.
- **Rejected shortcuts:** `--auto`, `--dangerously-skip-permissions`, host-home substitution, permanent config changes, copied host auth directories, managed-daemon attachment, IDE-only observations, and any Codex-subscription assertion.

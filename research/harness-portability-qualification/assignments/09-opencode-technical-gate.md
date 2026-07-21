# OpenCode Technical Qualification Gate

## Assignment

**Goal:** Audit OpenCode against the first technical qualification gate: executable/version, clean and default-environment isolation, temporary skill materialization, noninteractive two-turn replay, event evidence, model/reasoning, permissions, normal authentication boundary, cancellation, and bounded temporary-state cleanup.

**Scope:** Current first-party OpenCode documentation and the `anomalyco/opencode` `dev` source, safe local inspection of the installed OpenCode binary, and the existing Skill Issue implementation and candidate-research records.

**Exclusions:** Authentication, model calls, session creation, skill installation, configuration changes, permanent workarounds, Codex-subscription viability, and a production implementation or product-support decision.

## Sources

- Skill Issue production source: `cli/internal/harness/harness.go`, `cli/internal/evaluation/runtime.go`, `cli/internal/evaluation/evaluation.go`, and `cli/internal/replay/process.go`.
- Skill Issue qualification context: `research/harness-portability-qualification/assignments/01-existing-runtime-implementation.md` and `02-product-support-and-setup-contract.md`.
- Official OpenCode CLI documentation, current July 2026: <https://opencode.ai/docs/cli> — `run`, `--session`, `--format json`, `--model`, `--variant`, `--auto`, `--pure`, environment controls, and session deletion.
- Official OpenCode configuration documentation, current July 2026: <https://opencode.ai/docs/config> — merge/precedence rules, `OPENCODE_CONFIG`, `OPENCODE_CONFIG_DIR`, project traversal, plugins, instructions, and managed configuration.
- Official OpenCode skills documentation, current July 2026: <https://opencode.ai/docs/skills> — native skill roots, `skill` tool, frontmatter, and skill permissions.
- Official OpenCode permissions documentation, current July 2026: <https://opencode.ai/docs/permissions> — permission actions, granular `bash` and `skill` rules, defaults, and approval behavior.
- Official OpenCode server documentation, current July 2026: <https://opencode.ai/docs/server> — headless server, session APIs, event stream, and instance disposal API.
- Upstream primary source at `anomalyco/opencode` `dev`: `packages/opencode/src/cli/cmd/run.ts`, `config/config.ts`, `config/paths.ts`, `skill/index.ts`, `session/instruction.ts`, and `effect/runtime-flags.ts` (retrieved 2026-07-21). The repository release page identifies `v1.18.3` as the current release on 2026-07-16: <https://github.com/anomalyco/opencode/releases/tag/v1.18.3>.
- Local observation, 2026-07-21: the available OpenCode binary reports `1.14.39`; help exposes `run`, `--session`, `--format json`, `--model`, `--variant`, `--pure`, and `debug skill`. Read-only `debug config --pure` and `debug skill --pure` both failed in the pre-existing local state with `PRAGMA wal_checkpoint(PASSIVE)`.

## Findings

### 1. Executable and version are observable, but OpenCode publishes no stable minimum that proves this gate

The local binary is present and reports `1.14.39`. Its help confirms the command surface required for a candidate adapter: noninteractive `run`, JSON output, session resumption, provider/model selection, variant selection, and `--pure`. Upstream currently advertises `v1.18.3` as its latest release. Neither the official documentation nor release page states a version floor that guarantees every required isolation and evidence control; several material controls are visible in current `dev` source but were not proved against the local `1.14.39` binary.

**Evidence:** Local `--version`, `--help`, and `run --help` observations; official CLI documentation; upstream `run.ts`; GitHub release page.

**Implication:** A qualification implementation must resolve `opencode` by absolute executable path, require an explicitly tested version range, and fail before materialization when the local binary lacks the tested flags and event schema. `1.14.39` is command-surface evidence only, not a qualified baseline.

### 2. Default execution is materially ambient; `--pure` alone cannot establish a clean evaluation

OpenCode merges, rather than replaces, remote, global, custom, project, `.opencode` directory, inline, managed, and macOS MDM configuration. It searches project configuration upward to the worktree, loads global configuration, can load configuration supplied by an authenticated organization, and applies managed configuration after user-controlled layers. `--pure` is documented as running without external plugins; it is not documented as disabling ordinary configuration, rules, skills, agents, provider settings, or managed controls.

Current source exposes additional controls: `OPENCODE_DISABLE_PROJECT_CONFIG`, `OPENCODE_DISABLE_EXTERNAL_SKILLS`, `OPENCODE_DISABLE_CLAUDE_CODE`, and `OPENCODE_DISABLE_DEFAULT_PLUGINS`. These controls are source-backed for the current development revision, but this run did not establish availability or exact behavior in `1.14.39`. Managed configuration remains loaded after other sources and cannot be overridden by a temporary inline configuration.

**Evidence:** Official configuration precedence and merge contract; official CLI global flags; upstream `config/config.ts:398-534`, `config/paths.ts:23-40`, and `effect/runtime-flags.ts:16-30`; upstream `session/instruction.ts:52-106`.

**Implication:** The first gate is not met by a default local launch. A future qualification probe needs a wholly owned runtime root plus a minimal explicit environment, must disable project/external/Claude-compatible discovery where the tested version supports it, and must detect managed configuration as an unisolated policy boundary. Setting only `OPENCODE_CONFIG` or `OPENCODE_CONFIG_CONTENT` is insufficient because both merge with other sources.

### 3. Native temporary skill supply is feasible only as part of the owned-runtime isolation route

OpenCode natively discovers skills from `.opencode/skills`, global OpenCode skills, `.claude/skills`, and `.agents/skills`; project discovery traverses upward to the worktree. Current source also scans `OPENCODE_CONFIG_DIR` for `skill` or `skills` directories and supports explicit configured skill paths and URLs. The runtime registers a built-in `customize-opencode` skill before scanning disk, and duplicate names overwrite an already registered name after a warning.

With project configuration and external skill discovery disabled, a uniquely named Skill Issue payload materialized under an owned `OPENCODE_CONFIG_DIR/skills/<name>/SKILL.md` is a source-backed candidate route. The source does not establish this path against the installed binary, and the local `debug skill --pure` probe could not inspect discovery because its pre-existing database state failed first.

**Evidence:** Official skills documentation; upstream `skill/index.ts:21-35`, `125-139`, `173-232`, and `259-315`; upstream `config/paths.ts:23-40`; local read-only probe result.

**Implication:** Temporary supply can be evaluated without modifying user or project skill roots only after clean-runtime isolation is proven. The supplied skill needs a unique valid name, required frontmatter, an explicit `permission.skill` allow rule, and a discovery probe that proves the generated payload is visible while unrelated ambient skills are absent. The built-in skill remains an explicit caveat to any claim of an empty skill set.

### 4. `run` and `--session` support a resumable noninteractive two-turn conversation

The official CLI documents `opencode run` as noninteractive automation, accepts `--format json`, and accepts `--session` to continue a session. Current source first fetches the requested session and uses its returned ID; a fresh noninteractive run creates a session, then prompts it. It subscribes to session events and exits its event loop when that session reaches `idle`. This supports a first run followed by one `--session <captured-id>` turn without a persistent external server.

The existing Skill Issue generic stub uses this form, but its current production runtime rejects OpenCode before materialization or replay. No local prompt or resume was executed in this assignment.

**Evidence:** Official CLI documentation; upstream `run.ts:456-532`, `670-676`, `788-794`, and `828-872`; Skill Issue `cli/internal/replay/process.go`; Skill Issue `cli/internal/evaluation/runtime.go`.

**Implication:** The CLI has a credible replay primitive. Qualification must prove that the first raw event stream contains one stable session ID, the resume invocation preserves that ID, and both turns finish at `idle` without an error event. The present generic Skill Issue session-key search must be extended for OpenCode's `sessionID` spelling if it is to consume current raw events directly.

### 5. Raw JSON events can provide session, completion, tool-failure, and marker evidence, subject to a candidate-specific validator

For `--format json`, current `run.ts` emits newline-delimited JSON envelopes with `type`, `timestamp`, and `sessionID`. It emits `tool_use` for completed and errored tool parts, `step_start`, `step_finish`, completed text/reasoning, and `error` for session failures. Its loop stops only after `session.status` becomes `idle`. Tool errors are emitted as `tool_use` parts but are not themselves added to the command's `error` accumulator.

The official server exposes session APIs and an SSE `/event` endpoint as an alternative machine-readable interface, but using a server would add owned port, process, and session disposal obligations. The one-shot CLI JSON form is the narrower candidate route.

**Evidence:** Official CLI and server documentation; upstream `run.ts:678-690`, `693-818`, and `828-872`; official server session/event API descriptions.

**Implication:** An OpenCode adapter can require: an initial `sessionID`, `session.status=idle` as the terminal completion signal, no `error` envelope, and no `tool_use` part with `state.status=error`. Marker attribution can match the opaque Skill Issue marker command in a completed Bash `tool_use` part. This is a source-supported validation design, but no sample event from the tested binary or marker execution was collected, so activation and failure classification remain unqualified.

### 6. Model and reasoning controls exist, but values are provider-specific and require an explicit per-model mapping

`run` accepts `--model provider/model` and `--variant`; official documentation describes the latter as provider-specific reasoning effort. Upstream passes the parsed provider/model and variant into session creation and prompt execution. The local help independently exposes both flags.

**Evidence:** Official CLI documentation; upstream `run.ts:550-565` and `858-865`; local help observation.

**Implication:** A candidate adapter may map the harness model to `--model` and a supported reasoning tier to `--variant`, but it must reject unsupported combinations before replay. This assignment found no official cross-provider enum or local live evidence for the required model cells, so a generic Skill Issue `--reasoning` argument cannot be forwarded unchanged.

### 7. A restrictive temporary permission profile can permit the marker while preserving noninteractive failure behavior

OpenCode permission rules support `allow`, `ask`, and `deny`, including pattern-based `bash` and `skill` rules. Default permissions are mostly `allow`; noninteractive `run` creates per-session rules that deny question and plan transitions, and auto-rejects permission prompts unless `--auto` is passed. `--auto` approves requests not explicitly denied, making it unsuitable as the qualification boundary. Current source exposes the `tool_use` error state required to detect a denied marker command.

**Evidence:** Official permissions and skills documentation; upstream `run.ts:430-448`, `796-815`, and `719-727`; upstream `skill/index.ts:310-315`.

**Implication:** The owned temporary configuration needs an explicit default-deny posture with only the evaluated workspace operations, the generated marker command, and the generated skill name allowed. The marker must be allowed through a narrowly matching Bash rule and `permission.skill`; every other tool or command must remain denied or ask-and-reject. Exact pattern semantics and marker invocation must be live-tested, because a mismatch will look like a model miss unless the adapter treats tool errors as tooling failures.

### 8. Authentication can remain an external prerequisite, but no clean-state credential bridge was validated

OpenCode officially supports provider credentials through its auth configuration and, for some providers, process environment variables. The provider documentation recommends `opencode auth list` to inspect configured credentials, while identifying environment-authenticated providers as an exception. This audit did not query authentication status or use credentials. The local read-only discovery/config commands failed against the existing state database, so they provide no successful preflight evidence.

**Evidence:** Official provider documentation; official configuration documentation; local read-only probe result.

**Implication:** Authentication belongs outside the temporary skill/config payload. A qualified route must prove a normal, supported credential source survives the isolated launch without copying a user OpenCode state directory, and it must preflight the chosen provider/model before installation. Whether any path can use a Codex subscription is outside this assignment and remains unaddressed here.

### 9. Cancellation and temporary-state cleanup are the principal unresolved technical boundary

The CLI creates durable sessions by default, and the official CLI exposes `opencode session delete <sessionID>`. Local `debug paths` reports separate data, cache, state, log, and temporary roots. Current configuration source can create a `.gitignore` and start detached dependency-install work for each configuration directory; it also loads managed configuration and may fetch remote configuration after authentication. The public CLI and server documentation do not provide a documented signal/descendant-cleanup contract for `run`, and no cancellation or child-process probe was allowed in this assignment.

Skill Issue's generic Unix adapter owns a process group and can kill that group, but OpenCode has never reached that adapter in production because runtime preparation rejects it. Process-group ownership therefore does not prove OpenCode server, shell-tool, dependency-install, or provider child cleanup.

**Evidence:** Official CLI session and uninstall documentation; local `debug paths` observation; upstream `config/config.ts:436-457`; Skill Issue `cli/internal/replay/process_group_unix.go`, `cli/internal/replay/process.go`, and `cli/internal/evaluation/runtime.go`.

**Implication:** A qualification route must allocate all OpenCode configuration/data/cache/state paths beneath one owned temporary root, capture the created session ID, delete it after replay, then terminate the owned process group and remove that root only after children are confirmed gone. A targeted interruption probe must prove no server or tool descendant survives. Until those checks pass, bounded cleanup and cancellation are unsupported rather than assumed.

### 10. Gate outcome: technically promising, but not yet passed

OpenCode has current source-backed primitives for headless execution, resume, raw events, model variants, skill discovery, permissions, and explicit session deletion. The required clean/default isolation can be designed from current source controls, but it has not been proven on a pinned released version or the available local `1.14.39` binary. No authenticated two-turn run, generated-skill discovery, marker attribution, permission-denial, interruption, descendant cleanup, or private-root cleanup evidence was collected.

**Evidence:** Findings 1–9; existing Skill Issue source blocks OpenCode before replay.

**Implication:** Keep OpenCode at **technical gate unqualified**. The strongest next qualification target is a version-pinned, owned-runtime probe using temporary config and skill roots, source-supported discovery-disabling controls, a provider credential boundary, candidate-specific JSON validation, and explicit session/process/root cleanup. The local database failure is an additional reason to avoid treating ambient user state as the evaluation runtime.

## Notes

- **Official documentation:** establishes the public CLI/config/skills/permissions/server contracts cited above.
- **Upstream source:** establishes current-development controls and exact raw-event behavior that are not all surfaced in the public documentation; source observations are not claims about the local `1.14.39` release.
- **Local observation:** covers only safe version/help/path/debug inspection. The failed `debug config --pure` and `debug skill --pure` calls do not establish that a new isolated runtime would fail.
- **Inference:** the owned-runtime arrangement described in Findings 2, 3, 7, and 9 is a qualification design inferred from source controls. It remains unsupported until validated against a pinned released binary with a real authorized provider.
- **No production support change:** Skill Issue currently registers OpenCode for ordinary installation but rejects it as an evaluation harness before temporary-skill materialization.

# Gemini CLI First Technical Qualification Gate

## Assignment

**Goal:** Audit the current Google Gemini CLI only against the first technical qualification gate: executable/runtime prerequisites, bounded clean-state launch, temporary skill discovery, headless multi-turn replay, native evidence, controls, cancellation, and run-owned cleanup.

**Scope:** The controlling local qualification records, current Gemini CLI documentation, the official `google-gemini/gemini-cli` release record, and upstream source at commit `acae7124bdd849e554eaa5e090199a0cf08cd782` (the 2026-07-20 nightly source line).

**Exclusions:** Codex-subscription or provider-bridge viability; installation, login, authentication, binary execution, local configuration mutation, and any claim that documentation alone qualifies a runtime.

## Sources

- `research/harness-portability-qualification/assignments/01-existing-runtime-implementation.md` — production qualification pattern and fail-closed evidence requirements.
- `research/harness-portability-qualification/assignments/02-product-support-and-setup-contract.md` — host-ownership, output-state, and qualification boundary.
- `research/harness-portability-qualification/assignments/03-local-candidate-research-copilot-gemini.md` — Gemini is evaluation-unqualified and absent from this machine's `PATH`.
- [Gemini CLI installation, execution, and releases](https://geminicli.com/docs/get-started/installation/) — supported distribution channels, Node 20+ runtime, stable/preview/nightly policy.
- [Gemini CLI v0.51.0 release](https://github.com/google-gemini/gemini-cli/releases/tag/v0.51.0) — latest stable release at inspection on 2026-07-21; published 2026-07-16.
- [Gemini CLI configuration](https://geminicli.com/docs/get-started/configuration/) and [upstream configuration source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/config/config.ts) — configuration precedence, `GEMINI_CLI_HOME`, launcher flags, model selection, trust, and policy paths.
- [Trusted folders](https://geminicli.com/docs/cli/trusted-folders/) and [upstream trusted-folder source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/config/trustedFolders.ts) — workspace-config loading and headless trust behavior.
- [Agent Skills](https://geminicli.com/docs/cli/skills/), [activate skill tool](https://geminicli.com/docs/tools/activate-skill/), and [upstream skill discovery/activation source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/core/src/skills/skillManager.ts) — discovery tiers, trust gate, activation, and consent.
- [Headless mode](https://geminicli.com/docs/cli/headless/), [session management](https://geminicli.com/docs/cli/session-management/), and [upstream noninteractive source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/nonInteractiveCli.ts) — JSONL events, session ID, resume, result status, tool events, and cancellation flow.
- [Policy engine](https://geminicli.com/docs/reference/policy-engine/), [sandboxing](https://geminicli.com/docs/cli/sandbox/), and [upstream cleanup source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/utils/cleanup.ts) — exact temporary approval policies, sandbox route, and signal handlers.
- [Authentication setup](https://geminicli.com/docs/get-started/authentication/) and [upstream noninteractive authentication source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/validateNonInterActiveAuth.ts) — operator-owned headless authentication boundary.

## Findings

### Technical Gate Is Documented But Locally Unqualified

Gemini CLI has a documented route for the individual technical capabilities, but it does not pass Skill Issue's first technical gate from documentation alone. Local evidence says `gemini` is unavailable on `PATH`, and the current CLI rejects Gemini evaluation requests before runtime preparation. No local launch, protocol, activation, isolation, resume, cancellation, or cleanup observation exists.

**Evidence:** The local candidate map records the absent executable and the production evaluation allow-list. The existing runtime implementation requires a live structured completion, stable session, attributed activation, run-owned cleanup, and cancellation evidence; this candidate has none.

**Implication:** Classify Gemini CLI as **blocked pending named live evidence**, rather than implemented or smoke-qualified. The remainder of this document identifies the smallest non-permanent route worth validating on one exact release.

### Distribution And Version Must Be Pinned

Gemini's official install page supports the `gemini` executable through npm, Homebrew, MacPorts, Anaconda, Cloud Shell/Workstations, and short-lived `npx`; its documented runtime baseline is Node.js 20+. The upstream package declares the executable as `gemini` and Node `>=20`. The release API identified `v0.51.0` as latest stable on 2026-07-21; upstream `main` was the later nightly line `v0.52.0-nightly.20260720.gacae7124b`.

**Evidence:** [Installation and releases](https://geminicli.com/docs/get-started/installation/) recommends stable for normal use and labels preview/nightly as less validated. [Package metadata at the inspected source commit](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/package.json) declares `bin.gemini` and Node `>=20`; [v0.51.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.51.0) is the latest stable release record at inspection.

**Implication:** A qualification run must record an absolute executable, `gemini --version`, distribution channel, operating system, Node runtime when relevant, and the exact model-access account. Start with stable `v0.51.0` or a later stable release; a nightly source reading is not version compatibility proof.

### Private Launch State Can Reduce Ambient Inputs, With Explicit Limits

Gemini applies default, system-default, user, workspace, and system-override settings; workspace settings override user settings, while system settings remain stronger. `GEMINI_CLI_HOME` relocates user-level CLI storage, and `GEMINI_CLI_TRUSTED_FOLDERS_PATH` relocates trust state. The CLI source discovers built-in skills, active extension skills, user skills, and trusted-workspace skills in that order. There is no documented equivalent of Codex's `--ignore-user-config` plus an all-ambient-skill deny list, and the source always discovers built-in skills.

**Evidence:** [Configuration](https://geminicli.com/docs/get-started/configuration/) defines the setting hierarchy and `GEMINI_CLI_HOME`; [skill manager source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/core/src/skills/skillManager.ts) loads built-in, extension, user, and trusted workspace tiers; [trusted folders](https://geminicli.com/docs/cli/trusted-folders/) states that an untrusted workspace ignores workspace settings, `.env`, MCP servers, commands, and workspace skills.

**Implication:** The defensible candidate launch design is an empty temporary launch directory outside the evaluated workspace, a run-owned `GEMINI_CLI_HOME`, a run-owned trusted-folders path, and an explicit `--include-directories <workspace>` only after a live probe proves required workspace read/write semantics. Keep system settings/admin policies intact and record them; bypassing them would conflict with the harness-ownership boundary. This suppresses ordinary user/project state but cannot claim a completely skill-free or policy-free default session.

### Temporary Skill Injection Is Native And Bounded To A Run Root

Gemini discovers `.gemini/skills/<name>/SKILL.md` and `.agents/skills/<name>/SKILL.md` at workspace scope when the launch directory is trusted. It also provides `gemini skills link <path> --scope workspace --consent`; upstream source creates the link under the current working directory's workspace-skill root and rejects invalid/path-traversing names. Both direct materialization and a link are temporary when the launch directory is run-owned and deleted after the run.

**Evidence:** [Agent Skills](https://geminicli.com/docs/cli/skills/) documents both root families, the `skills` management commands, and the workspace scope. [Upstream link command](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/commands/skills/link.ts) and [link implementation](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/utils/skillUtils.ts) target only the selected user/workspace root; [skills source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/core/src/skills/skillManager.ts) disables workspace-skill discovery unless trusted.

**Implication:** Materialize the complete instrumented Skill Issue directory at `<temporary-launch>/.gemini/skills/<name>` or use a temporary workspace-scoped link. Use `--skip-trust` only for that empty launch directory and only for the process lifetime. Do not alter the evaluated workspace, the operator's user skills, or persistent settings; verify discovery by a safe `gemini skills list` probe before scenario replay.

### Headless Replay And Session Continuity Have A Strictly Parseable Route

`-p`/`--prompt` forces noninteractive operation. `--output-format stream-json` emits JSONL, starting with `init` containing `session_id` and model, followed by messages, `tool_use`, `tool_result`, errors, and a terminal `result`. The source emits `result.status` as `success` or `error`; it also throws on actual tool-response errors. `--session-id` creates a caller-selected new identifier and `--resume <ID>` reloads the saved conversation on later process invocations.

**Evidence:** [Headless mode](https://geminicli.com/docs/cli/headless/) defines the JSONL event kinds and exit-code contract. [CLI option source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/config/config.ts) validates `--output-format`, `--session-id`, and mutually exclusive session options. [Noninteractive source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/nonInteractiveCli.ts) emits `init`, records each tool request/result, and emits the terminal result. [Session management](https://geminicli.com/docs/cli/session-management/) documents resume by full ID.

**Implication:** A candidate adapter can start with generated `--session-id <opaque-id> -p <turn-1> --output-format stream-json`, require one `init` with that ID and one successful terminal `result`, then use `--resume <opaque-id>` for every later prompt. The parser must fail on malformed JSONL, any terminal `result.status != success`, nonzero exit, missing or changed ID, an `error` event that denotes an execution failure, or a failed marker tool result.

### Native Activation Evidence Is Available, But Approval Must Be Scoped

Gemini exposes discovered skill metadata to the model and requires the agent to call `activate_skill` to load a matching skill. Upstream emits the requested tool name and arguments in `tool_use`; the activation source injects the selected skill's body and resources only after that call succeeds. This permits a native activation record when the JSONL stream contains `tool_use` for `activate_skill` with the target name and the matching successful `tool_result` ID.

**Evidence:** [Agent Skills](https://geminicli.com/docs/cli/skills/) defines discovery, agent-only activation, consent, and instruction injection. [Activation source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/core/src/tools/activate-skill.ts) looks up the named skill, activates it, and returns the instructions/resources. [Noninteractive source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/nonInteractiveCli.ts) emits tool-call arguments and correlated tool results.

**Implication:** This is stronger than final prose and can meet the local native-evidence standard if a live probe confirms the event schema on the pinned release. Activation normally presents consent. Headless replay must therefore use a run-owned, explicit `--policy <temporary.toml>` rule that allows only `activate_skill` for the temporary target and only the exact private marker command/path needed for attribution; the policy engine documents per-tool, argument-pattern, and `interactive = false` rules. `--approval-mode=yolo` is rejected for qualification because it broadly approves unrelated actions.

### Model Selection Exists; Reproducible Reasoning Control Is Not Yet Proven

The `--model` flag has highest documented model-selection precedence, and the stream `init` event records the selected model. Gemini also has automatic fallback/model routing and its `--model` flag does not govern sub-agent models. Advanced thinking controls (`thinkingBudget` and `thinkingConfig`) are settings-based model configuration, rather than a documented simple headless reasoning-level flag.

**Evidence:** [Model routing](https://geminicli.com/docs/cli/model-routing/) places `--model` above environment/settings and documents automatic fallback. [Model selection](https://geminicli.com/docs/cli/model/) excludes sub-agent model control. [Advanced model configuration](https://geminicli.com/docs/cli/generation-settings/) documents `thinkingConfig`/`thinkingBudget` in settings. [Noninteractive source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/nonInteractiveCli.ts) includes the configured model in `init`.

**Implication:** Pass an explicit concrete `--model` and retain the `init` plus final per-model usage record. A generated private settings file may be investigated for fixed thinking configuration, but no adapter should advertise a normalized `low`/`medium`/`high` reasoning mapping until that configuration and output evidence pass on the pinned executable. Treat automatic routing or sub-agent model changes as a failure unless the evaluation contract expressly records and accepts them.

### Permissions And Sandbox Need A Narrow, Observed Contract

Default approval prompts, `auto_edit`, `plan`, and `yolo` are distinct; headless `ask_user` is treated as deny. The policy engine can allow or deny tools by name and argument pattern for noninteractive execution. `--sandbox` enables Gemini's tool sandbox, but the sandbox type and behavior depend on platform/container availability and configuration. Folder trust blocks headless execution in an untrusted folder unless `--skip-trust` or the trust environment variable is used for that session.

**Evidence:** [Policy engine](https://geminicli.com/docs/reference/policy-engine/) defines decisions, argument patterns, noninteractive rules, and the fact that `ask_user` becomes deny in headless mode. [Sandboxing](https://geminicli.com/docs/cli/sandbox/) documents the session flag and Docker/Podman/seatbelt possibilities. [Trusted folders](https://geminicli.com/docs/cli/trusted-folders/) documents `FatalUntrustedWorkspaceError` and temporary trust bypasses in headless mode.

**Implication:** Use no broad approval mode. Generate a policy in the private root that permits target activation and only the exact signal command, permits only necessary workspace actions after review, and denies the rest. Sandbox selection is not yet a qualification fact: a live run must prove that the chosen host sandbox preserves the exact private marker, workspace route, and artifact boundaries without spawning unowned state.

### Authentication And State Are Operator Inputs; Private State Is Removable

Authentication remains outside the runner. Headless mode can use previously cached credentials, but a private `GEMINI_CLI_HOME` intentionally hides ordinary user-level state; the official headless guidance instead requires operator-provided supported environment authentication when cached credentials are unavailable. Session history, prompts, tool executions, and reasoning summaries are saved below the Gemini home in a project-hashed `tmp/.../chats` path, and `--delete-session` exists; the private-root design makes all run-created Gemini state removable as one owned directory.

**Evidence:** [Authentication setup](https://geminicli.com/docs/get-started/authentication/) says headless operation uses existing cached authentication or supported environment configuration such as Gemini API key/Vertex AI. [Noninteractive auth source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/validateNonInterActiveAuth.ts) fails when no configured/env authentication is available. [Session management](https://geminicli.com/docs/cli/session-management/) records the private session location and deletion behavior.

**Implication:** The harness may preserve only explicit, operator-supplied supported credential environment variables or external credential references; it must never copy login caches, keys, or configuration into the temporary home. Remove the private Gemini home and temporary launch/policy/skill roots after each run. This is an authentication-boundary observation only; it does not establish any subscription or provider-bridge eligibility.

### Cancellation Is Handled In-Process, But Descendant Ownership Remains Unproven

Gemini registers `SIGHUP`, `SIGTERM`, and `SIGINT` handlers that run its exit cleanup once. Noninteractive execution attaches an abort controller for terminal input cancellation, disposes its scheduler, and cleans internal console listeners. These sources establish graceful CLI cleanup, not that every tool, sandbox container, extension process, browser session, or descendant has terminated.

**Evidence:** [Cleanup source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/utils/cleanup.ts) registers the three signals and calls exit cleanup; [noninteractive source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/cli/src/nonInteractiveCli.ts) aborts, disposes, and clears listeners. The local qualified adapter record requires owned descendant cleanup rather than direct-process termination alone.

**Implication:** The outer adapter should keep Gemini in an owned process group and send cancellation to that group, but the first gate remains blocked until a live cancellation probe proves no Gemini-created descendants remain and that the temporary state can be deleted. A timeout/interrupt that leaves a sandbox or helper process is a tooling failure, not a model result.

## Notes

- **Validated documentation/source capability:** stable distribution, private home/trust locations, workspace-scoped temporary skills, `-p`, JSONL output, caller-specified sessions, `--resume`, tool events, model selection, policy rules, sandbox option, signal handlers, and user-owned authentication boundary.
- **Local observation:** neither the Gemini executable nor a Gemini authentication route is present in this environment; no invasive or authenticated probe was performed.
- **Unsupported until live evidence:** exact stable-version option behavior; `--include-directories` access semantics; a fully ambient-free launch; private-home authentication with no copied credential material; headless approval of `activate_skill` under a narrow temporary policy; stable marker command tracing; exact thinking configuration; model-fallback detection; sandbox behavior; and descendant cleanup.
- **Rejected workarounds:** persistent edits to user/project settings, skills, hooks, trust files, extensions, policies, or login caches; copied credentials; `--approval-mode=yolo`; and inferring cleanup from the direct CLI process exiting.
- **Useful qualification probe sequence:** resolve/pin executable and version; create only a temporary launch/home/trust/policy root; confirm the injected skill is discovered; run two `-p --output-format stream-json` turns using a generated session ID then `--resume`; require activation and marker tool traces; interrupt one controlled run; inspect descendants/state; remove only the run-owned roots.

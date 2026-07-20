# GitHub Copilot CLI First Technical Qualification Gate

## Assignment

**Goal:** Determine whether the current GitHub Copilot CLI exposes the technical controls required to begin a portable, clean, noninteractive Skill Issue evaluation route.

**Scope:** Current official GitHub documentation and the official `github/copilot-cli` repository/release materials; the three supplied local assignments; read-only local executable evidence. The audit covers distribution, configuration isolation, temporary skill discovery, prompt mode, resume, machine-readable output, model and effort controls, permissions, authentication boundary, cancellation, and bounded state cleanup.

**Exclusions:** A final support decision; live installation, authentication, or runtime probes; the independent-agent and native skill-activation-evidence gates; Codex-subscription access; implementation changes; credential transfer; and permanent user, repository, policy, or plugin modifications.

## Sources

- Local production/context: `research/harness-portability-qualification/assignments/01-existing-runtime-implementation.md`, `02-product-support-and-setup-contract.md`, and `03-local-candidate-research-copilot-gemini.md` — current runner boundary, qualification bar, and absent local `copilot` executable.
- Local qualification contract: `skills/skill-evaluation-and-refinement/SKILL.md` — exact-surface, isolation, and evidence gates.
- GitHub Docs, [Installing GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/set-up-copilot-cli/install-copilot-cli), inspected 2026-07-21 — supported installers, prerequisites, and token authentication boundary.
- GitHub Docs, [GitHub Copilot CLI command reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference), inspected 2026-07-21 — prompt mode, JSONL output, session selection, settings overrides, skills, permissions, model/effort, sandbox, and environment controls.
- GitHub Docs, [GitHub Copilot CLI programmatic reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-programmatic-reference), inspected 2026-07-21 — programmatic permissions, model selection precedence, effort configuration, and token redaction.
- GitHub Docs, [GitHub Copilot CLI configuration directory](https://docs.github.com/en/enterprise-cloud%40latest/copilot/reference/copilot-cli-reference/cli-config-dir-reference), inspected 2026-07-21 — state inventory, `COPILOT_HOME`, `COPILOT_CACHE_HOME`, setting precedence, session persistence, and deletion effects.
- GitHub Docs, [Using GitHub Copilot CLI session data](https://docs.github.com/en/copilot/how-tos/copilot-cli/use-copilot-cli/chronicle), inspected 2026-07-21 — local sessions, default account synchronization, session IDs, and `--resume` behavior.
- GitHub Docs, [Configuring GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/set-up-copilot-cli/configure-copilot-cli) and [Allowing and denying tool use](https://docs.github.com/en/copilot/how-tos/copilot-cli/use-copilot-cli/allowing-tools), inspected 2026-07-21 — trust prompts, persisted permissions, and least-privilege controls.
- GitHub Docs, [About cloud and local sandboxes for GitHub Copilot](https://docs.github.com/en/copilot/concepts/about-cloud-and-local-sandboxes), inspected 2026-07-21 — experimental local sandbox and cloud-sandbox scope.
- Official source/release record: [github/copilot-cli changelog](https://github.com/github/copilot-cli/blob/main/changelog.md), inspected 2026-07-21 — release-note-only evidence that version 1.0.13 fixed shell-process cleanup and that later releases changed prompt/session behavior.

## Findings

### Current Distribution Is Defined, but the Local Executable Is Absent

GitHub documents the `copilot` executable as installable through WinGet, Homebrew, npm, the GitHub-hosted install script, or direct release executables. npm requires Node.js 22 or later; Windows also requires PowerShell 6 or later. GitHub says the CLI is available to active Copilot plans, subject to organization or enterprise policy. The supplied local candidate map recorded on 2026-07-21 that `copilot` does not resolve on this environment's `PATH`; this audit did not install, authenticate, or execute it.

**Evidence:** GitHub's installation guide lists the named distributions and prerequisites; `03-local-candidate-research-copilot-gemini.md` records the read-only `command -v copilot` absence.

**Implication:** Qualification must begin by pinning and recording the actual `copilot --version` and executable path. No local version, protocol, or cleanup claim can be promoted from documentation into live qualification evidence.

### Private State Is Supported, but a Default-Clean Session Is Only Partly Enforceable

`COPILOT_HOME` replaces the complete default `~/.copilot` configuration-and-state tree. That tree otherwise holds authentication/application state, settings, personal instructions, skills, agents, hooks, extensions, installed plugins, MCP/LSP configuration, persisted permissions, logs, and local session history. `COPILOT_CACHE_HOME` separately relocates the platform cache, which is otherwise outside `COPILOT_HOME`. A fresh, output-owned home and cache therefore provide a viable nonpersistent state root, and prevent user-scoped configuration from loading.

The command reference supplies `--no-custom-instructions`, `--no-experimental`, `--no-remote`, `--no-remote-export`, `--disable-builtin-mcps`, `--disable-mcp-server`, `--available-tools`, and `--excluded-tools`. In prompt mode, repository extensions and workspace MCP sources are disabled unless their opt-in environment variables are set; repository hooks load if the folder is already trusted or `COPILOT_ALLOW_ALL` is set. `--no-custom-instructions` prevents loading instruction files such as `AGENTS.md`, but it does not document a universal disable switch for project skills, remote organization skills, built-in skills, or managed MDM policy. Settings also cascade from MDM, user, repository, local, environment, then flags; the MDM bypass-permission restriction wins over lower scopes.

**Evidence:** The configuration-directory reference specifies `COPILOT_HOME` replacement, separate cache location, complete state inventory, and setting precedence. The command reference documents the suppression/disable flags, prompt-mode repository opt-ins, and all skill discovery sources. It identifies remote organization skills and built-in skills as discovery sources in addition to project and personal paths.

**Implication:** A private `COPILOT_HOME` plus private `COPILOT_CACHE_HOME`, a controlled environment, `--no-custom-instructions`, prompt mode, and explicit remote/MCP disables can isolate the mutable user-owned state. They cannot establish the required _default environment containing only supplied evaluation components_ while project, built-in, remote, and device-managed inputs may still exist. A candidate route must inventory the exact installed build's discovered skills/plugins/instructions and fail closed on unsuppressible ambient sources instead of masking them with a same-name skill.

### Temporary Explicit Skill Injection Exists, but It Does Not Suppress All Ambient Discovery

Skills are directories with `SKILL.md`; their contents are injected when the agent or a user invokes them. `COPILOT_SKILLS_DIRS` accepts additional comma-separated skill directories, and the CLI has `--plugin-dir` for local plugin directories. This permits a generated, output-owned skill directory to be provided for a single process without registering or persisting the skill. The documented priority is project locations, then personal/custom/plugin/built-in/remote sources as applicable; current releases also state that `--plugin-dir` skills take precedence over personal skills, not project skills.

The temporary skill must retain a false or absent `disable-model-invocation` frontmatter field for proactive selection; `user-invocable` is independently available for slash invocation. The documents describe `/skills list` and noninteractive `copilot plugins list --kind skill --json` as discovery aids, but explicitly exclude session-scoped hooks and custom agents from that command's output.

**Evidence:** Command-reference skill schema, locations, precedence, `COPILOT_SKILLS_DIRS`, `--plugin-dir`, and `copilot plugins list`; official release notes for `--plugin-dir` precedence; the local controls document's implicit-invocation requirement.

**Implication:** `COPILOT_SKILLS_DIRS=<private-skill-root>` is a documented temporary injection path and `plugins list --json` is useful preflight evidence. It is insufficient alone for a clean route or activation proof because it neither removes higher-priority project skills nor inventories every session-scoped/custom source. Permanent `copilot skill add`, plugin enablement, or edits to project skill roots are rejected workarounds for this gate.

### Headless Prompt Mode, JSONL, and Resumption Are Documented

`copilot -p`/`--prompt` runs a prompt programmatically and exits when complete. `--output-format=json` emits JSONL, while `--no-ask-user` prevents the agent from pausing for an elicitation. GitHub documents `--resume SESSION-ID` and `--session-id ID`; the latter uses an existing exact session/task ID or creates a new session only for a valid UUID. Session data is kept in `session-state/<session-id>/events.jsonl`, and the command-line/session docs say `--resume SESSION-ID` reloads the full conversation history. The command reference says prompt mode waits for background agents and shells by default, subject to `COPILOT_TASK_WAIT_TIMEOUT_SECONDS`.

**Evidence:** Command-reference entries for `-p`, `--output-format`, `--no-ask-user`, `--resume`, `--session-id`, and the task wait environment variable; session-data and configuration-directory references for session ID, `events.jsonl`, and resume semantics.

**Implication:** A two-process, multi-turn route is technically plausible: create with an exact generated UUID using `-p --session-id`, capture the returned JSONL and local event log, then issue `-p --resume <captured-id>` for later turns. The official CLI reference documents JSONL but does not publish a stable event schema that proves a terminal success, a tool failure, the current session ID, or a skill activation. The current local adapter's permissive plain-text Copilot fallback is therefore unacceptable for qualification; an installed-version probe must define strict event predicates before the candidate can advance.

### Model and Reasoning Controls Are Available at Launch

`--model` and `COPILOT_MODEL` select the model; command-line selection takes precedence over environment, settings, and default selection. Current documentation lists `--effort`/`--reasoning-effort` values `low`, `medium`, `high`, `xhigh`, and `max`; some models support configurable effort. The programmatic reference describes persistent `effortLevel` only as a settings-file value, whereas the command reference exposes the launch flag. `--context` controls the context tier separately.

**Evidence:** Command and programmatic references document model precedence, current model names, `--effort`/`--reasoning-effort`, settings `effortLevel`, and `--context`.

**Implication:** The adapter can pass the requested model and medium reasoning at each headless turn without mutating settings. Exact accepted model/effort combinations, the resolved model recorded in JSONL, and resume-time persistence remain live-version checks; unsupported values must be reported as tooling failures.

### Programmatic Permissions Can Be Narrow, but Trust and Policy Must Be Proven Live

GitHub states that programmatic use requires `--allow-all-tools` or `COPILOT_ALLOW_ALL`; command-line permissions also permit narrow `--allow-tool`, `--deny-tool`, path, URL, and available/excluded-tool filters. Filters can allow an exact shell command or constrained write path. The CLI persists permissions in the configuration directory when a user approves them, whereas a fresh private home makes any generated decisions disposable. A trust decision is requested when starting a session from an untrusted directory. Local sandboxing can be requested with `--sandbox` only in experimental mode and is still public preview; it governs commands executed by Copilot, not the outer CLI process.

**Evidence:** Programmatic and command references specify programmatic `--allow-all-tools`, least-privilege filters, and permission flags. Configuration and sandbox documents specify trust prompts, persisted approvals, experimental availability, and sandbox scope.

**Implication:** This creates a qualification tension: official programmatic guidance requires broad tool approval, while Skill Issue needs a tightly bounded marker command. The adapter should first test whether a narrowly allowed marker shell command plus required read/write tools actually avoids prompts in `-p`; a blanket `--allow-all`/`--yolo` is not an acceptable default. Managed policy can also suppress bypass flags, so an interaction-free run must preflight the exact account/device policy and record refusal as a tooling blocker.

### Authentication Can Be Passed Without Copying a Home, with Account Policy as a Boundary

Interactive login uses OAuth device flow and stores tokens in the platform credential store when available, otherwise in `COPILOT_HOME`. GitHub also supports a user-owned fine-grained personal access token with the `Copilot Requests` permission via `COPILOT_GITHUB_TOKEN`, `GH_TOKEN`, or `GITHUB_TOKEN` in precedence order; the command reference marks the Copilot token as redacted in output by default. The installation guide requires active Copilot access and says organization/enterprise policy can disable the CLI.

**Evidence:** GitHub's installation and command references document login storage, token authentication, variable precedence/redaction, and organization policy boundary.

**Implication:** A disposable-home route can receive a runner-provided `COPILOT_GITHUB_TOKEN` at launch and avoid copying credentials or mutating the user's normal Copilot home. The harness must preflight its own supported authentication status with the isolated environment and treat expired tokens, license/policy refusal, or credential-store coupling as blockers. This finding deliberately makes no claim about any Codex subscription or credential bridge.

### Cancellation and Complete Descendant Ownership Remain Unsupported

The current official changelog says version 1.0.13 fixed shell-process cleanup when a session ends, and the command reference provides `stop_bash`/`stop_powershell` tools plus a prompt-mode task wait timeout. GitHub's public CLI documentation does not establish that terminating a noninteractive parent `copilot` process on macOS, Linux, and Windows kills every descendant, background agent, extension, MCP server, or remote operation. The local runner's process-group cleanup is its own behavior, not evidence of Copilot descendant ownership.

**Evidence:** Official `github/copilot-cli` changelog release note for 1.0.13; command-reference tool list and `COPILOT_TASK_WAIT_TIMEOUT_SECONDS`; `01-existing-runtime-implementation.md` process-ownership distinction.

**Implication:** The parent evaluation runner must own and terminate its process group/job object, but first-gate completion still requires an installed-version cancellation probe that checks for detached Copilot descendants. Until then, descendant cleanup is unsupported rather than assumed from the changelog.

### Temporary Local State Can Be Bounded, but Cloud Session Retention Needs an Explicit Check

With output-owned `COPILOT_HOME` and `COPILOT_CACHE_HOME`, generated settings, permissions, logs, local session state, database, plugin data, and injected skills can be deleted as one bounded run root after the final capture. GitHub documents the local session event log and says sessions are synchronized to the GitHub account by default. `--no-remote-export` is documented to disable session export to GitHub.com and GitHub Mobile, but this audit found no primary source that proves it prevents every session-sync pathway in headless prompt mode. Deleting local session state only removes local copies when sessions have synced.

**Evidence:** Configuration-directory reference state inventory, separate cache path, local deletion effects, and session-state layout; session-data documentation on default synchronization; command reference `--no-remote-export` scope.

**Implication:** Bounded cleanup is feasible for local temporary state only when both home and cache are private and recursively removed by the runner after capture. The harness must run with remote export disabled and prove in the exact installed version that no retained remote session remains, or classify cloud-retention cleanup as unsupported. It must never delete or edit a user's normal Copilot home, cache, credentials, sessions, permissions, plugins, or settings.

### First-Gate Result: Documentation Supports a Conditional Probe, Not Qualification

The current public surface supplies the ingredients for a provisional probe: pinned `copilot`, private home/cache, explicit temporary skills, `-p`, JSONL, exact session IDs/resume, launch-time model/effort, private token authentication, and per-run permission options. It does not document a complete ambient-discovery exclusion, a stable JSON event contract for success/tool failure/session/skill evidence, strict noninteractive least-privilege behavior, descendant termination, or remote-retention cleanup. No local binary is available to resolve those gaps.

**Evidence:** The preceding official-source findings together with the local absent-executable observation.

**Implication:** GitHub Copilot CLI should remain an additional unqualified target. It may enter a tightly scoped live technical probe only after a version/account is available; the probe must fail closed on ambient discovery, prompt/trust/permission interruption, missing strict JSONL evidence, resume mismatch, residual processes, or residual local/remote state. It has not passed the first technical qualification gate.

## Notes

- No installation, login, invocation, credential inspection, or state mutation was performed. All executable behavior beyond the local PATH absence remains documentation-derived until tested on the pinned build.
- The public command reference is unusually current and substantial, but its JSONL documentation does not supply an event schema or an activation-evidence guarantee. Treat output fields and exit semantics as unsupported until captured and validated.
- `COPILOT_HOME` isolates the documented configuration tree; it does not by itself relocate the platform cache, device MDM policy, OS credential store, remote organization skills, built-in skills, or project-controlled discovery.
- Skill discovery and temporary injection belong to this technical gate. Whether a skill actually loads before influencing a response and whether a fresh independent agent can be created are later qualification gates.
- Useful search terms: `COPILOT_HOME`, `COPILOT_CACHE_HOME`, `COPILOT_SKILLS_DIRS`, `--output-format=json`, `--session-id`, `--resume`, `--no-remote-export`, `GITHUB_COPILOT_PROMPT_MODE_REPO_HOOKS`, `permissions.disableBypassPermissionsMode`.

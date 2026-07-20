# Existing Authentication and Subscription Patterns

## Assignment

**Goal:** Reconstruct authentication, provider, model, and paid-access ownership in the four locally qualified evaluation routes, with emphasis on reusing existing OpenAI Codex access without Skill Issue taking custody of credentials or user configuration.

**Scope:** Production evaluation and replay code, the native setup contract, retained two-turn smoke evidence, the repository CLI contract, and the machine-local launcher behavior that the smoke campaign used for Cursor and Claude Code.

**Exclusions:** Candidate-harness research, credential contents, account identities, installation/repair workflows, and a claim that a two-turn smoke qualifies a full governed campaign.

## Sources

- [Native setup contract](../../../plans/harness-setup.md) — owned configuration, authentication, isolation, cleanup, and platform boundaries for the four routes.
- [Runtime preparation](../../../cli/internal/evaluation/runtime.go) and [evaluation control path](../../../cli/internal/evaluation/evaluation.go) — concrete creation of temporary runtime state, environment ownership, authentication preflight, and adapter launch options.
- [Process adapters](../../../cli/internal/replay/process.go) and [Pi RPC adapter](../../../cli/internal/replay/pi.go) — exact native commands, model/reasoning mapping, process ownership, and Pi provider validation.
- [CLI evaluation contract](../../../cli/README.md) — supported-harness boundary, user-provided-access rule, executable override, and documented Codex/Pi ownership limits.
- [Real-harness smoke report](../../../evaluations/skill-calling/smoke/real-harness-smoke-report.md) — observed macOS two-turn runs, versions, effective routes, cleanup, and the original Pi failure cause.
- Local ignored launchers inspected without reading credentials: `.skill-issue/cursor/cursorx`, `.skill-issue/claudex/claudex`, and `.skill-issue/claudex/manage`. They document the qualified local executable and proxy behavior but are machine-local evidence, not a portable product contract.
- [Pi official quickstart](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/quickstart.md) and [Pi security research](../../deep-research/harness-direct-installation-architecture/assignments/09-pi-direct-installation.md) — the native `/login` and `auth.json` ownership model, plus Pi's no-native-sandbox limitation.

## Findings

### Codex reuses the normal authenticated home while isolating evaluation behavior

The Codex route launches the resolved user `codex` executable from the selected workspace and retains the normal `CODEX_HOME` rather than creating a substitute home. Before temporary skills are installed, the runner calls `codex login status`; an unavailable login or unavailable requested model stops the run. Initial and resumed turns use `--ignore-user-config` and `--ignore-rules`, disable plugins, apps, project instructions, and ambient skills through process-only flags/configuration. Those controls isolate behavior but preserve Codex's supported authentication path.

**Evidence:** `prepareRuntime` has no Codex private-home creation branch, while `CheckAuthentication` runs `codex login status` before `PrepareEvaluation`. The Codex setup contract explicitly says to use the normal `CODEX_HOME`, forbid a replacement home and credential copying, and direct an unauthenticated user to `codex login`. The smoke report records `codex-cli 0.144.1` with a ChatGPT login completing both default and custom two-turn routes.

**Implication:** This is the reference acceptable second-gate pattern for existing OpenAI access: the user owns login and model entitlement; Skill Issue only verifies availability, passes the selected model through, and never reads, copies, prints, moves, retains, or cleans up the credentials. It remains conditional on the selected model being accessible and on organization-managed policy allowing the controlled invocation.

### Cursor has native account access, not a Codex-subscription bridge

Cursor starts in a run-owned empty environment so ordinary Cursor configuration, data, store, and skills remain isolated. On macOS, the runtime places only a symlink to the user's `~/Library/Keychains` under its temporary `HOME`; the Cursor CLI reads its own `cursor-user` credential from Keychain. The runner checks `agent status` or `cursor-agent status` under that environment before instrumentation. If the isolated runtime is not logged in, the user may run the native browser login in that same environment; cleanup removes the link and deliberately does not log the user out.

**Evidence:** `prepareCursorRuntime` creates a temporary home/store/plugin tree, symlinks `Library/Keychains`, writes the run-specific CLI config, and keeps the selected workspace as working directory. `CheckAuthentication` performs the Cursor status check with the clean environment. The setup contract identifies the Keychain bridge and macOS-only platform; the smoke report records Cursor `2026.07.16-899851b` with native Keychain login and both default Auto-select and explicit `gpt-5.4` routes.

**Implication:** This is an acceptable qualified native-login route when the platform is macOS, the user can authorize Cursor, and the Keychain bridge is permitted. Its model and paid access belong to Cursor's account and its native Auto-select/model controls; it supplies no evidence that an existing ChatGPT/Codex subscription can be forwarded to a different candidate harness.

### Claude Code used a user-provisioned compatibility launcher backed by a local Codex OAuth proxy

The production Claude adapter itself runs the supplied `claude` executable in an isolated launch directory with an explicit model and effort, and it does not copy credentials. The qualified smoke executable was a machine-local launcher selected with `--executable`. That launcher starts a loopback-only CLIProxyAPI process, uses a local proxy token, unsets direct Anthropic/OpenRouter key variables, points Claude at the local endpoint, and maps all Claude default-model variables to a selected proxy model. Its manager exposes a user-operated Codex-subscription login flow and lists models exposed by the authenticated proxy. The smoke report observed the `opus` alias resolving through that launcher to `gpt-5.6-sol`.

**Evidence:** `claudeArgs` passes the request model/effort and process-group cleanup includes `claude project purge`; `prepareClaudeRuntime` creates only launch/passed-skill state. The local launcher and manager inspect only their own protected runtime files, start the proxy on `127.0.0.1`, and declare a Codex OAuth login step. The smoke report records Claude Code `2.1.205` with CLIProxyAPI `7.2.91`, confirms the explicit executable override, and verifies cleanup of the owned Claude/proxy process group.

**Implication:** Treat this as a conditional, user-provisioned compatibility route rather than a general direct Claude Code authentication pattern. It is acceptable for a local campaign only when the operator already owns and explicitly selects the isolated proxy/launcher, its Codex OAuth login, model mapping, and token storage. Skill Issue must not create, configure, authenticate, inspect, copy, or reset that proxy runtime; it should report the requested and effective model. The evidence is local smoke evidence, not proof that Claude Code or OpenAI officially supports this compatibility arrangement, so it cannot justify making a third-party proxy a default route for another harness.

### Pi preserves its native `openai-codex` agent directory and separates temporary sessions

Pi's adapter parses `provider/model` from the model value, defaults a bare model to `openai-codex`, launches one RPC process with that provider and model, and validates the returned state after preflight and each turn. The runtime gives Pi a temporary `HOME`, session directory, and supplied-skills root, but leaves `PI_CODING_AGENT_DIR` set to the caller's existing value or `~/.pi/agent`. That preserves Pi's native `openai-codex` login and its token-refresh ownership while keeping evaluation sessions disposable. The direct `auth.json` is neither read nor replaced by Skill Issue.

**Evidence:** `preparePiRuntime` sets `PI_CODING_AGENT_DIR` to the existing/default native directory and `PI_CODING_AGENT_SESSION_DIR` to the private runtime; `piAdapter.Start` sends `--provider`, `--model`, `--no-session`, and explicit skills before `validatePiState` checks the resolved pair. The setup contract links Pi's native `/login` and `auth.json` paths. The smoke report records Pi `0.80.10`, existing native `openai-codex` login, a `0600` retained auth file, and explains that replacing the agent directory with an empty temporary directory caused the earlier failure.

**Implication:** Pi is an acceptable native Codex-access route when the existing Pi agent directory is available and its `openai-codex` login has access to the requested model. It must retain that directory untouched and keep only sessions/skills temporary. Pi lacks a workspace-only filesystem sandbox, so the enclosing process boundary remains material. In addition, `controlledEnvironment(..., true)` forwards inherited credential-like environment variables to Pi; that is process passing rather than credential-file copying, but a future route should only forward the minimum provider variables required by its documented auth mechanism.

### The implementation's preflight coverage differs by native credential mechanism

The current control path has an explicit status preflight only for Codex and Cursor. Claude and Pi rely on their launch/protocol behavior to surface authentication failure: Claude receives its selected executable/launcher, and Pi's RPC `get_state`/provider checks must succeed. All four treat a native model, permission, session, or protocol rejection as a tooling error instead of model-evaluation data.

**Evidence:** `CheckAuthentication` returns early unless the harness is Codex or Cursor; `evaluation.go` invokes it before temporary skill installation. Claude and Pi then reach their adapters, where Claude requires successful structured events and Pi requires a successful state response with the requested provider/model. The CLI contract states that harness login and model access are user-provided prerequisites and that launch/protocol failure is tooling failure.

**Implication:** A second-gate route is acceptable only with an observable preflight or equivalent fail-closed protocol confirmation of both authentication and effective provider/model. A candidate that merely inherits a credential file, environment variable, or proxy endpoint without proving the selected model is reachable remains conditional or blocked pending that evidence.

### Derived criteria for later candidate second-gate research

The existing routes yield the following criteria without asserting any candidate satisfies them:

- The user, native harness, or explicitly selected user-owned launcher must own login, subscription/payment access, token refresh, and persistent credential/configuration state.
- Skill Issue may check availability and pass a documented provider/model selection, but may not create replacement homes, copy/read/print credentials, mutate ordinary user configuration, or log the user out during cleanup.
- Temporary evaluation state must be limited to run-owned skills, sessions, process configuration, and evidence; cleanup must remove only that state and owned child processes.
- A compatibility proxy is lower fit than a native provider route and requires an explicit launcher, loopback/process ownership, requested-to-effective model reporting, and source-backed support evidence; otherwise classify it as blocked rather than silently installing or configuring it.
- Platform stores and policy are prerequisites: macOS Keychain availability for the qualified Cursor route, supported native agent-directory access for Pi, and any organization-managed Codex/Cursor/Claude restriction remain authoritative.
- A route with no native workspace sandbox requires an enclosing OS/container boundary; provider access alone cannot qualify it.

## Notes

- The smoke campaign proves only two-turn local routes on 2026-07-20. It leaves governed 30-turn qualification and public-interface validation deferred.
- The project-local Cursor and Claude launchers were deliberately absent from login-shell `PATH`; `--executable` selected them. Their presence demonstrates local qualification, not automatic installation or environment-management ownership by Skill Issue.
- The local Claude proxy files are ignored, machine-specific state. This report deliberately excludes their credentials, configuration values, token contents, account information, and logs.
- No claim here establishes a general right to use a ChatGPT/Codex subscription through a non-Codex harness. The direct Codex and Pi native-provider routes are distinct from the conditional local Claude compatibility launcher.

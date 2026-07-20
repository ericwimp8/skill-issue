# Kilo Local Configuration Qualification

**Qualification date:** 2026-07-21  
**Qualified surface:** local macOS ARM64 configuration and launcher behavior  
**Kilo version:** `7.4.11`  
**Authentication:** native OpenAI ChatGPT Pro/Plus headless OAuth  
**Evaluation model:** `openai/gpt-5.6-sol`  
**Reasoning variant:** `medium`

## Decision

The local Kilo configuration passes the bounded configuration qualification and may proceed to production adapter implementation. The qualification proves native Codex-subscription access, process-scoped configuration, controlled supplied-skill discovery, noninteractive structured replay, stable two-turn session continuation, native skill activation, private marker attribution, least-privilege denial, interruption, session deletion, daemon suppression, and bounded cleanup on the recorded machine.

This is not campaign qualification. Skill Issue still lacks a Kilo evaluator admission path, runtime owner, candidate-specific protocol validator, result integration, automated regression coverage, governed 30-turn evidence, and cross-platform qualification.

## Ownership Boundary

The qualification runtime lives entirely under ignored repository-local `.skill-issue/kilo/`. It owns:

- the pinned Kilo platform binary;
- process-scoped configuration, authentication, data, state, and cache roots;
- the qualification wrapper and temporary skill materialization;
- raw local evidence and created Kilo sessions.

Smoke workspaces were clean temporary Git repositories outside this repository.

The wrapper leaves the real `HOME` value unchanged. It supplies private XDG roots only to the child Kilo process. The normal Kilo paths under `~/.config`, `~/.local/share`, `~/.local/state`, `~/.cache`, `~/.kilo`, and `~/.kilocode` remained absent before and after qualification.

The normal product contract remains unchanged: Skill Issue assumes that an operator-owned Kilo CLI is already installed. The repository-local binary is a local qualification fixture, not a proposed product installer, repair system, backup, migration, or credential bridge.

## Clean Installation Evidence

The preflight found no Kilo executable, npm installation, active Kilo environment override, project Kilo configuration, global Kilo configuration, global skill directory, authentication, session database, state, or cache.

The matching official `@kilocode/cli-darwin-arm64@7.4.11` platform package was installed directly under the ignored qualification root. The executable reports `7.4.11`.

This reinforces the product boundary: the future adapter should resolve and invoke an existing executable. It should not reproduce package-manager installation behavior.

## Source-Backed Runtime Controls

Kilo configuration is merged across several scopes, and global or project skills can be discovered automatically. The qualification uses documented configuration and skill contracts plus source-owned runtime flags to exclude ambient inputs. See [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli), [Kilo CLI reference](https://kilo.ai/docs/code-with-ai/platforms/cli-reference), [Kilo settings](https://kilo.ai/docs/getting-started/settings), [Kilo skills](https://kilo.ai/docs/customize/skills), and the pinned [`7.4.11` plugin registry](https://github.com/Kilo-Org/kilocode/blob/v7.4.11/packages/opencode/src/plugin/index.ts).

The wrapper launches Kilo through a minimal environment and supplies only required host identity, terminal, path, temporary-directory, and private XDG values. It disables:

- inherited provider API keys and unrelated environment configuration;
- project Kilo configuration;
- Claude Code prompts and skills;
- external `.claude` and `.agents` skill discovery;
- external plugins through pure mode;
- sharing, auto-sharing, and remote control;
- automatic updates;
- automatic LSP downloads;
- codebase indexing and semantic search;
- session ingest, presence, and the file watcher;
- daemon use through `KILO_NO_DAEMON`.

Kilo's compiled internal plugin registry remains enabled. The pinned source shows that the native OpenAI Codex OAuth handler is an internal plugin. Disabling all default plugins would remove the required subscription-authentication route, so retaining compiled internal authentication behavior is required product behavior rather than ambient user customization. The handler is defined in the pinned [`CodexAuthPlugin`](https://github.com/Kilo-Org/kilocode/blob/v7.4.11/packages/opencode/src/plugin/openai/codex.ts).

Pure mode skips external plugin origins. Kilo still displays its default external plugin declarations in resolved configuration, but the pinned runtime does not load them while pure mode is active.

## Resolved Configuration

The run-owned configuration resolves to:

- global and small model: `openai/gpt-5.6-sol`;
- primary agent: `code`;
- reasoning variant: `medium`;
- enabled provider: `openai` only;
- sharing and remote control: disabled;
- automatic updates: disabled;
- snapshots: disabled;
- formatter and LSP: disabled;
- MCP configuration: absent;
- additional instructions and remote skill URLs: absent.

The deny-first permission policy permits only:

- workspace reads, excluding environment files;
- workspace edit, glob, grep, and list operations;
- the selected evaluation skill names;
- the exact qualification signal command.

It denies unrelated Bash commands, external-directory access, questions, task/subagent delegation, web fetching, web search, semantic search, codebase search, and Kilo memory operations.

Kilo applies its own agent defaults and later enables semantic search on the `code` agent. The qualification therefore repeats the semantic-search denial at the agent-owned permission layer. The effective ordered trace confirms that this final denial wins. The only later external-directory exception is Kilo's private tool-output directory, which is product-owned structured-output storage.

## Authentication And Model Proof

The native OpenAI provider exposed `ChatGPT Pro/Plus (browser)` and `ChatGPT Pro/Plus (headless)` methods. Headless device OAuth completed successfully and wrote one OpenAI OAuth credential with mode `0600` inside the private XDG data root.

With the minimal evaluation wrapper active:

- `auth list` reported one OpenAI OAuth credential;
- no inherited API-key provider appeared;
- `models openai` exposed the eligible Codex catalogue;
- `openai/gpt-5.6-sol` was present;
- an authenticated JSON turn returned exactly `KILO_BASELINE_OK`;
- the terminal event reported `reason: "stop"`, exit zero, and zero stderr.

The model call used subscription OAuth and reported zero API cost in the Kilo event. This proves the user's existing eligible Codex subscription can drive the selected model through Kilo's native route on this release.

## Structured Protocol Finding

Kilo `7.4.11` emitted every JSON event twice as an identical line during these runs. The duplicate pair retained the same event type, timestamp, `sessionID`, part ID, message ID, and payload. The behavior affected baseline, tool-use, text, and terminal events.

This does not prevent deterministic evaluation, but the future candidate-specific validator must collapse exact duplicate events before counting skill calls or terminal events. It must still reject conflicting events and must not use general-purpose deduplication that could hide distinct repeated tool calls.

The JSON protocol otherwise supplied the required stable camel-case `sessionID`, typed tool state, text content, token counts, terminal reason, and OpenAI provider metadata.

## Smoke Matrix

### Clean Baseline

- Clean temporary Git workspace.
- No supplied evaluation skills; discovery reported only Kilo's compiled `kilo-config` skill.
- Explicit `openai/gpt-5.6-sol`, `medium` variant, and `code` agent.
- Exact `KILO_BASELINE_OK` response.
- One stable `sessionID`, terminal `stop`, exit zero, zero stderr, and no workspace change.
- Each structured event appeared as one exact duplicate pair.

### Canonical Two-Turn Route

Inputs came from the retained two-turn smoke scenario and the current canonical `document-update-discipline` and `prompt-writing` supporting skills.

Turn 1:

- loaded `document-update-discipline` through the native `skill` tool;
- invoked its exact private activation marker under active Turn 1 attribution;
- created `smoke-plan.md` in the selected workspace;
- reached terminal `stop` with exit zero and zero stderr.

Turn 2:

- resumed the exact Turn 1 `sessionID`;
- loaded `document-update-discipline` and `prompt-writing`;
- invoked both exact private activation markers under active Turn 2 attribution;
- coherently added the required Research Handoff section and prompt;
- reached terminal `stop` with exit zero and zero stderr.

All three expected skill activations were observed and attributed to their active turns. Kilo also activated `prompt-writing` on Turn 1. That additional call is retained as model behavior rather than treated as a launcher failure.

### Explicit Custom Two-Turn Route

Inputs came from the retained custom smoke scenario, answer sheet, and supplied `smoke-skill`.

- Turn 1 used explicit `openai/gpt-5.6-sol` and `medium`, loaded and activated `smoke-skill`, and created `custom-note.md`.
- Turn 2 resumed the exact Turn 1 `sessionID`, loaded and activated `smoke-skill` again, and appended the requested sentence.
- Both turns reached terminal `stop` with exit zero and zero stderr.
- The final file retained both requested sentences.
- Kilo also activated `document-update-discipline` on both custom turns. The expected `smoke-skill` calls remained present and correctly attributed.

### Permission Denial

A separate turn explicitly requested `uname -a` through Bash. The structured `tool_use` event reported status `error` and identified the matching deny rules. The command did not run. The model then accurately reported the denial, and the turn completed with terminal `stop`, exit zero, and zero stderr.

This establishes that a required signal command can be allowed while an unrelated command fails closed without converting the otherwise complete turn into a launcher failure.

### Interruption And Process Cleanup

An active model turn emitted structured events and was interrupted with Ctrl-C before completion. Kilo exited nonzero immediately. A host process inventory found no surviving Kilo process afterward.

`kilo daemon status --json` reported `running: false`, `stale: false`, and `reason: "not running"`. The wrapper's `KILO_NO_DAEMON` boundary therefore held throughout the smoke work.

## Cleanup Proof

Seven created qualification sessions were deleted explicitly through `kilo session delete`: the baseline, canonical, custom, permission-denial, two completed interruption probes, and the interrupted session.

Session listings from every smoke workspace were empty afterward. A read-only count against Kilo's isolated database also reported zero session rows. The daemon remained stopped.

Cleanup then removed:

- all supplied and instrumented skills;
- active-turn state;
- activation records;
- stored smoke session identifiers;
- all temporary smoke workspaces;
- generated Finder metadata under the qualification root.

Post-cleanup discovery reported only Kilo's compiled `kilo-config` skill. The isolated OpenAI OAuth credential and `openai/gpt-5.6-sol` model route remained usable. Normal user-level Kilo paths remained absent.

## Kilo 7.4.11 Session-List Defect

`kilo session list --all --format json` fails even when the isolated database contains zero sessions. The pinned command implementation spreads the `Session.listGlobal(...)` effect without yielding it before formatting, so the formatter receives values without a session `time` object and throws while reading `time.updated`. See the pinned [`session list` command](https://github.com/Kilo-Org/kilocode/blob/v7.4.11/packages/opencode/src/cli/cmd/session.ts).

This defect does not block the proposed adapter. Skill Issue already owns the exact `sessionID` values created by a run and can delete them directly. Workspace-scoped `session list --format json`, explicit `session delete`, and the isolated database all behaved correctly. The adapter should not depend on the broken global-list command.

## Additional Prerequisites Discovered

The configuration proof adds these concrete requirements to the future adapter:

1. Retain compiled internal plugins because native OpenAI subscription OAuth depends on `CodexAuthPlugin`; disable external plugins separately with pure mode.
2. Launch from a minimal environment. Filesystem isolation alone does not exclude inherited provider API keys.
3. Set private XDG config, data, state, and cache roots while leaving `HOME` unchanged.
4. Disable project config, Claude compatibility, external skill discovery, remote control, session ingest, indexing, presence, and the file watcher with source-backed process flags.
5. Set `KILO_NO_DAEMON` and verify that cancellation leaves no surviving Kilo process.
6. Generate a deny-first permission policy after Kilo's compiled agent defaults, add an agent-level semantic-search denial, and validate the final ordered permission trace.
7. Supply evaluation skills through a run-owned `skills.paths` directory while external compatibility discovery remains disabled.
8. Parse the camel-case `sessionID` field and require a stable value across resumed turns.
9. Collapse only exact duplicate JSON event lines before candidate-specific validation.
10. Require terminal `step_finish` with `reason: "stop"`, zero process exit, no structured error, and no errored required tool call.
11. Treat an explicitly denied unrelated tool call as structured evidence; whether the turn remains tooling-complete depends on the scenario's required capabilities.
12. Track and delete every created native session directly by ID. Do not depend on `session list --all` in Kilo `7.4.11`.
13. Preserve the operator-owned OAuth credential while removing run-owned skills, markers, workspaces, and session state.

## Remaining Work

The next phase may implement the Kilo evaluator at the existing harness runtime owner. It must add evaluator admission, defaults, run-owned configuration generation, skill instrumentation and materialization, exact command permissions, Kilo-specific duplicate-event normalization and JSON validation, `sessionID` continuation, cancellation ownership, native session deletion, result artifacts, recovery cleanup, focused tests, and the existing built-in/custom CLI smoke routes.

Broader support claims still require a production implementation audit, automated regression coverage, full governed campaign evidence, and qualification on additional operating systems.

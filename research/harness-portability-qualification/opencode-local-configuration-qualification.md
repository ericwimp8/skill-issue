# OpenCode Local Configuration Qualification

**Qualification date:** 2026-07-21  
**Qualified surface:** local macOS ARM64 configuration and launcher behavior  
**OpenCode version:** `1.18.4`  
**Authentication:** native OpenAI ChatGPT Plus/Pro browser OAuth  
**Evaluation model:** `openai/gpt-5.6-sol`  
**Reasoning variant:** `medium`

## Decision

The local OpenCode configuration passes the bounded configuration qualification and may proceed to production adapter implementation. The qualification proves native Codex-subscription access, clean process-scoped configuration, exclusive supplied-skill discovery apart from one compiled built-in skill, noninteractive structured replay, stable two-turn session continuation, native skill activation, private marker attribution, least-privilege denial, interruption, session deletion, and bounded cleanup on the recorded machine.

This is not campaign qualification. Skill Issue still lacks an OpenCode evaluator admission path, runtime owner, candidate-specific protocol validator, result integration, automated regression coverage, governed 30-turn evidence, and cross-platform qualification.

## Ownership Boundary

The qualification runtime lives entirely under ignored repository-local `.skill-issue/opencode/`. It owns:

- the pinned OpenCode platform binary;
- process-scoped configuration, authentication, data, state, and cache roots;
- the qualification wrapper and temporary skill materialization;
- clean smoke workspaces, private activation state, raw local evidence, and created OpenCode sessions.

The wrapper leaves the real `HOME` value unchanged. It supplies private XDG roots only to the child OpenCode process. The normal OpenCode paths under `~/.config`, `~/.local/share`, `~/.local/state`, `~/.cache`, and `~/.opencode` remained absent before and after qualification.

The normal product contract remains unchanged: Skill Issue assumes that an operator-owned OpenCode CLI is already installed. The repository-local binary is a local qualification fixture, not a proposed product installer, repair system, backup, migration, or credential bridge.

## Clean Installation Evidence

The previous Homebrew OpenCode `1.14.39` installation and OpenCode-owned configuration, authentication, sessions, state, and cache were removed with explicit user approval. No browser history, extensions, unrelated application data, or non-OpenCode files were inspected or removed.

The current official `opencode-ai@1.18.4` npm wrapper failed under npm `10.2.4` because its postinstall could not resolve the published Darwin ARM64 optional package. Installing the matching official `opencode-darwin-arm64@1.18.4` package directly succeeded. The executable reports `1.18.4`.

This reinforces the product boundary: the future adapter should resolve and invoke an existing executable. It should not reproduce package-manager installation behavior.

## Source-Backed Runtime Controls

OpenCode configuration is merged across multiple scopes, and global or project skills can be discovered automatically. The qualification uses the documented configuration and skill contracts plus source-owned runtime flags to exclude ambient inputs. See [OpenCode configuration](https://opencode.ai/docs/config/), [OpenCode skills](https://opencode.ai/docs/skills/), [OpenCode CLI](https://opencode.ai/docs/cli/), and the pinned [`1.18.4` plugin registry](https://github.com/anomalyco/opencode/blob/v1.18.4/packages/opencode/src/plugin/index.ts).

The wrapper launches OpenCode through a minimal environment and supplies only required host identity, terminal, path, temporary-directory, and private XDG values. It disables:

- inherited provider API keys and unrelated environment configuration;
- project OpenCode configuration;
- Claude Code prompts and skills;
- external `.claude` and `.agents` skill discovery;
- external plugins through `--pure`;
- sharing and auto-sharing;
- automatic updates;
- automatic LSP downloads;
- the file watcher, which otherwise emitted an FSEvents startup error in this sandbox.

OpenCode's compiled internal plugin registry remains enabled. The pinned source shows that the native OpenAI Codex OAuth handler is an internal plugin. Disabling all default plugins removes the ChatGPT OAuth method and the OpenAI model provider, so retaining compiled internal authentication behavior is required product behavior rather than ambient user customization. The native handler is defined in the pinned [`CodexAuthPlugin`](https://github.com/anomalyco/opencode/blob/v1.18.4/packages/opencode/src/plugin/openai/codex.ts).

## Resolved Configuration

The run-owned configuration resolves to:

- global and small model: `openai/gpt-5.6-sol`;
- primary agent: `build`;
- reasoning variant: `medium`;
- enabled provider: `openai` only;
- sharing: disabled;
- automatic updates: disabled;
- snapshots: disabled;
- formatter and LSP: disabled;
- external plugin list: empty;
- MCP configuration: absent;
- additional instructions and remote skill paths: absent.

The deny-first permission policy permits only:

- workspace reads, excluding environment files;
- workspace edit, glob, grep, and list operations;
- the selected evaluation skill names;
- the exact qualification signal command.

It denies unrelated Bash commands, external-directory access, questions, task/subagent delegation, web fetching, and web search. OpenCode's effective permission trace showed the explicit deny-first rules after its compiled defaults, so the qualification rules control the final decision. The only later exception is OpenCode's private tool-output directory, which is product-owned structured-output storage.

## Authentication And Model Proof

The native `OpenAI (ChatGPT Plus/Pro or API key)` provider route presented the browser and headless ChatGPT login methods when compiled internal plugins were enabled. Browser OAuth completed successfully and wrote one OpenAI OAuth credential with mode `0600` inside the private XDG data root.

With the minimal evaluation wrapper active:

- `auth list` reported one OpenAI OAuth credential;
- no inherited API-key provider appeared;
- `models openai` exposed the eligible Codex catalogue;
- `openai/gpt-5.6-sol` was present;
- an authenticated JSON turn returned exactly `OPENCODE_CODEX_SMOKE_OK`;
- the terminal event reported `reason: "stop"`, exit zero, and zero stderr.

The model call used subscription OAuth and reported zero API cost in the OpenCode event. This proves the user's existing eligible Codex subscription can drive the selected model through OpenCode's native route on this release.

## Smoke Matrix

### Clean Baseline

- Clean nested Git workspace.
- No supplied evaluation skills.
- Explicit `openai/gpt-5.6-sol` and `medium` variant.
- Three structured events: `step_start`, exact text, and terminal `step_finish`.
- One stable `sessionID`, terminal `stop`, exit zero, and zero stderr.
- No tool or skill activation.

### Canonical Two-Turn Route

Inputs came from the retained two-turn smoke scenario and the current canonical `document-update-discipline` and `prompt-writing` supporting skills.

Turn 1:

- loaded `document-update-discipline` through the native `skill` tool;
- invoked its exact private activation marker successfully;
- created `smoke-plan.md` in the selected workspace;
- reached terminal `stop` with exit zero and zero stderr.

Turn 2:

- resumed the exact Turn 1 `sessionID`;
- loaded `document-update-discipline` and `prompt-writing`;
- invoked both exact private activation markers under the active Turn 2 attribution;
- coherently added the required Research Handoff section and prompt;
- reached terminal `stop` with exit zero and zero stderr.

All three expected skill activations were observed and attributed to their active turns.

### Explicit Custom Two-Turn Route

Inputs came from the retained custom smoke scenario, answer sheet, and supplied `smoke-skill`.

- Only `smoke-skill` was supplied; the canonical skills were removed before launch.
- Turn 1 used explicit `openai/gpt-5.6-sol` and `medium`, loaded and activated `smoke-skill`, and created `custom-note.md`.
- Turn 2 resumed the exact Turn 1 `sessionID`, loaded and activated `smoke-skill` again, and appended the requested sentence.
- Both turns reached terminal `stop` with exit zero and zero stderr.
- The final file retained both requested sentences.

### Permission Denial

A separate turn explicitly requested `uname -a` through Bash. The structured `tool_use` event reported status `error` and identified the matching deny rules. The command did not run. The model then accurately reported the denial, and the turn completed with terminal `stop`, exit zero, and zero stderr.

This establishes that a required signal command can be allowed while an unrelated command fails closed without converting the otherwise complete turn into a launcher failure.

### Interruption And Process Cleanup

An active long-form model turn emitted its initial `step_start` and was interrupted with Ctrl-C. The OpenCode process exited nonzero immediately. A host process inventory found no surviving OpenCode process afterward.

## Cleanup Proof

Six created sessions were deleted explicitly: two baseline sessions, the canonical session, the custom session, the permission-denial session, and the interrupted session. Session listings from every smoke workspace were empty afterward.

Cleanup then removed:

- all supplied and instrumented skills;
- active-turn state;
- opaque token mappings;
- activation records;
- stored smoke session identifiers;
- generated Finder metadata under the qualification root.

Post-cleanup discovery reported only OpenCode's compiled `customize-opencode` skill. The isolated OpenAI OAuth credential and `openai/gpt-5.6-sol` model route remained usable. Normal user-level OpenCode paths remained absent.

## Additional Prerequisites Discovered

The configuration proof adds these concrete requirements to the future adapter:

1. Retain compiled internal plugins because native OpenAI subscription OAuth depends on `CodexAuthPlugin`; disable external plugins separately with `--pure`.
2. Launch from a minimal environment. Filesystem isolation alone does not exclude inherited provider API keys.
3. Set private XDG config, data, state, and cache roots while leaving `HOME` unchanged.
4. Disable project config, Claude compatibility, and external skill discovery with source-backed process flags.
5. Disable the file watcher in this macOS sandbox to avoid non-fatal FSEvents stderr.
6. Generate a deny-first permission policy after OpenCode's compiled agent defaults and validate the final ordered permission trace.
7. Parse the camel-case `sessionID` field and require a stable value across resumed turns.
8. Require terminal `step_finish` with `reason: "stop"`, zero process exit, no structured error, and no errored required tool call.
9. Treat an explicitly denied unrelated tool call as structured evidence; whether the turn remains tooling-complete depends on the scenario's required capabilities.
10. Delete every created native session during normal and interrupted cleanup while preserving the operator-owned OAuth credential.

## Remaining Work

The next phase may implement the OpenCode evaluator at the existing harness runtime owner. It must add evaluator admission, defaults, run-owned configuration generation, skill instrumentation and materialization, exact command permissions, OpenCode-specific JSON validation, `sessionID` continuation, cancellation ownership, native session deletion, result artifacts, recovery cleanup, focused tests, and the existing built-in/custom CLI smoke routes.

Broader support claims still require a production implementation audit, automated regression coverage, full governed campaign evidence, and qualification on additional operating systems.

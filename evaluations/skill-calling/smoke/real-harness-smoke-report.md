# Real-Harness Smoke Report

Date: 2026-07-20 to 2026-07-21

## Scope

This campaign used only embedded and custom two-turn smoke inputs. It did not run a governed 30-turn evaluation or implementation audit. Every evaluated project used a clean temporary Git workspace.

## Qualified Launchers

| Harness     | Version and authentication                                              | Result                                                                    |
| ----------- | ----------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| Codex       | `codex-cli 0.144.1`; ChatGPT login                                      | Built-in default and explicit custom routes complete                      |
| Cursor      | `2026.07.16-899851b`; native Keychain login                             | Built-in default and explicit `gpt-5.4` custom routes complete            |
| Claude Code | Claude Code `2.1.205`; local CLIProxyAPI `7.2.91` with Codex OAuth      | Default alias and explicit `gpt-5.6-sol` custom routes complete           |
| Pi          | `0.80.10`; existing `openai-codex` login in Pi's native agent directory | Built-in default and explicit custom routes complete                      |

Cursor and Claude Code were never missing. They were deliberately project-local and therefore absent from the login-shell `PATH`. The working routes use the CLI's `--executable` override. Pi was already authenticated; the original failure came from replacing `PI_CODING_AGENT_DIR` with an empty temporary directory.

## Final Runtime Corrections

- Cursor and Pi now run from the caller-selected evaluation workspace rather than a disposable harness workspace, so scenario-created project files remain available after cleanup.
- Pi keeps its existing configured agent directory, or `~/.pi/agent`, for native authentication. Skill Issue does not create, copy, read, replace, or remove that directory. Sessions and supplied evaluation skills remain temporary.
- Claude Code allowlists `Read`, `Write`, `Edit`, `Glob`, and `Grep`, plus only the exact Skill Issue signal command through `Bash`. Unrelated Bash commands can still be denied without invalidating an otherwise completed turn.
- Cursor and Claude commands run in private process groups. Completion, failure, and cancellation terminate the owned group, including Cursor's detached worker and the project-local Claude proxy.
- The pre-run summary displays an explicit harness executable or launcher when supplied.

## Built-In Routes

### Codex

- Model and reasoning: `gpt-5.6-sol`, `medium`
- Both expected `prompt-writing` calls were observed.

### Cursor

- Model: native Auto-select; Cursor owns model-native reasoning.
- `prompt-writing` was observed on Turn 1. Turn 2 recorded `document-update-discipline` instead, so the missing Turn 2 call remains valid evaluation data.
- `smoke-note.md` remained in the selected workspace.

### Claude Code

- Requested model and reasoning: `opus`, `medium`; the local launcher resolved the request through its configured `gpt-5.6-sol` proxy model.
- `prompt-writing` was observed on Turn 1 and missing on Turn 2.
- `smoke-note.md` remained in the selected workspace.

### Pi

- Model and reasoning: `openai-codex/gpt-5.6-sol`, `medium`
- Turn 1 observed expected `prompt-writing` and additional `document-update-discipline`; Turn 2 missed `prompt-writing`.
- `smoke-note.md` remained in the selected workspace.

## Explicit Custom Routes

Every custom route used:

- `evaluations/skill-calling/smoke/custom-skills/`
- `evaluations/skill-calling/smoke/custom-scenario.json`
- `evaluations/skill-calling/smoke/custom-answer-sheet.json`

| Harness     | Effective configuration              | Observed result                                                  |
| ----------- | ------------------------------------ | ---------------------------------------------------------------- |
| Codex       | `gpt-5.6-sol`, `medium`              | `smoke-skill` on both turns                                      |
| Cursor      | `gpt-5.4`, model-native reasoning    | `smoke-skill` on both turns                                      |
| Claude Code | `gpt-5.6-sol`, `medium`              | `smoke-skill` on both turns                                      |
| Pi          | `openai-codex/gpt-5.6-sol`, `medium` | `smoke-skill` on Turn 1; Turn 2 miss retained as evaluation data |

Each custom route created and retained `custom-note.md` in its selected workspace. Each produced `result.json` and `website.json`; runs using `--events` also produced `events.jsonl`.

Cursor's explicit `--reasoning medium` probe was rejected before evaluation side effects with the documented message that Cursor exposes no independent reasoning override.

## Cleanup Evidence

- Final Cursor run: no `agent`, detached `worker-server`, or private runtime remained.
- Final Claude run: no Claude process, CLIProxyAPI process, or private runtime remained.
- Pi runs closed their owned RPC process and removed temporary sessions and supplied skills.
- All successful runs removed output-owned private token mappings and recovery state.
- The existing Pi authentication file remained owned by Pi with mode `0600`; Skill Issue did not replace it.

## Boundary

Missing expected skill calls are model evaluation data when replay, attribution, artifacts, and cleanup complete. Custom smoke evidence remains local and is not publication evidence without separate review. The three governed 30-turn evaluations and final public-interface automated tests remain deferred.

# Harness Evaluation Path Hardening Progress

This file tracks the four-turn `gardening-web-application` hardening attempts
run from the development CLI channel. Machine-local executable and workspace
paths remain in ignored evidence under `output/harness-hardening/`.

## Preflight

| Harness       | Route                                                                | Version              | Model                                                                  | Reasoning |
| ------------- | -------------------------------------------------------------------- | -------------------- | ---------------------------------------------------------------------- | --------- |
| `codex`       | Operator `codex` executable resolved from `PATH`                     | `0.144.6`            | `gpt-5.6-sol`                                                          | `medium`  |
| `claude-code` | Operator `claude` executable resolved from `PATH`                    | `2.1.205`            | `opus`                                                                 | `medium`  |
| `cursor`      | `.skill-issue/cursor/home/.local/bin/agent`                          | `2026.07.16-899851b` | `grok` requested; installed identifiers recorded in preflight evidence | `medium`  |
| `opencode`    | `.skill-issue/opencode/bin/opencode`                                 | `1.18.4`             | `openai/gpt-5.6-sol`                                                   | `medium`  |
| `kilo-code`   | `.skill-issue/kilo/node_modules/@kilocode/cli-darwin-arm64/bin/kilo` | `7.4.11`             | `openai/gpt-5.6-sol`                                                   | `medium`  |
| `pi`          | Operator `pi` executable resolved from `PATH`                        | `0.80.10`            | `openai-codex/gpt-5.6-sol`                                             | `medium`  |

Exact resolved executable paths are retained in
`output/harness-hardening/preflight.txt`.

## Attempts

| Harness       | Attempt | Workspace Disposition                                                                                             | Outcome                                        | Failure Classification                                                                         | Fix Applied                                                                                                              | Configuration                                                                      | Evidence Path                                                                               |
| ------------- | ------: | ----------------------------------------------------------------------------------------------------------------- | ---------------------------------------------- | ---------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `codex`       |       1 | Temporary skills and private state removed; cosmetic empty `.agents/skills/` remains                              | Tooling error before turn 1                    | Harness configuration rejected                                                                 | Replaced invalid `agents.enabled=false` with supported `features.multi_agent=false` in the generated Codex configuration | `codex 0.144.6`; `gpt-5.6-sol`; `medium`                                           | `output/harness-hardening/codex-1/codex-20260721T070546Z-dfb4ee7f/failure.json`             |
| `codex`       |       2 | Temporary skills, private state, and run-owned processes removed; model-created `plans/` retained                 | Tooling-clean pass; 4 turns completed          | Model behavior: three additional `dictate-plan` calls                                          | Codex configuration fix included in rebuilt development CLI                                                              | `codex 0.144.6`; `gpt-5.6-sol`; `medium`                                           | `output/harness-hardening/codex-2/codex-20260721T070812Z-1d333bcc/result.json`              |
| `claude-code` |       1 | Temporary skills, private state, and run-owned processes removed; model-created `plans/` retained                 | Tooling error on turn 2 after turn 1 completed | Route/environment: sandboxed outer command did not retain the normal Claude session for resume | Rerun unchanged CLI with command-scoped access to normal Claude session state                                            | `claude-code 2.1.205`; `opus`; `medium`                                            | `output/harness-hardening/claude-code-1/claude-code-20260721T071340Z-189fe529/failure.json` |
| `claude-code` |       2 | Temporary skills, private state, and run-owned processes removed; model-created `plans/` retained                 | Tooling-clean pass; 4 turns completed          | Model behavior: four missing `document-update-discipline` calls                                | Outer command escalation included                                                                                        | `claude-code 2.1.205`; `opus`; `medium`                                            | `output/harness-hardening/claude-code-2/claude-code-20260721T071621Z-ce150814/result.json`  |
| `cursor`      |       1 | Temporary skills, private state, and run-owned processes removed                                                  | Tooling error before turn 1                    | Route/environment: relative executable became invalid under the run-owned working directory    | Resolved the project-local agent to an absolute invocation path                                                          | `cursor 2026.07.16-899851b`; requested `grok`; `medium`                            | `output/harness-hardening/cursor-1/cursor-20260721T071944Z-f4def1a9/failure.json`           |
| `cursor`      |       2 | Temporary skills, private state, and run-owned processes removed                                                  | Tooling error before model execution           | Route/environment: installed agent rejected the `grok` alias                                   | Selected the listed medium Grok identifier `cursor-grok-4.5-medium`                                                      | `cursor 2026.07.16-899851b`; requested `grok`; `medium`                            | `output/harness-hardening/cursor-2/cursor-20260721T072052Z-98ee7ba2/failure.json`           |
| `cursor`      |       3 | Temporary skills, private state, and run-owned processes removed; model-created `plans/` retained                 | Tooling-clean pass; 4 turns completed          | Model behavior: all expected calls observed                                                    | Absolute executable route and native Grok identifier included                                                            | `cursor 2026.07.16-899851b`; `cursor-grok-4.5-medium`; `medium`                    | `output/harness-hardening/cursor-3/cursor-20260721T072158Z-9d9477b2/result.json`            |
| `opencode`    |       1 | Temporary skills, private state, native session, and run-owned processes removed; model-created `plans/` retained | Tooling-clean pass; 4 turns completed          | Model behavior: three missing `document-update-discipline` calls                               | None                                                                                                                     | `opencode 1.18.4`; `openai/gpt-5.6-sol`; `medium`; qualified `XDG_DATA_HOME`       | `output/harness-hardening/opencode-1/opencode-20260721T072730Z-a5d32f21/result.json`        |
| `kilo-code`   |       1 | Temporary skills, private state, native session, and run-owned processes removed; model-created `plans/` retained | Tooling-clean pass; 4 turns completed          | Model behavior: three missing calls and one additional `prompt-writing` call                   | None                                                                                                                     | `kilo-code 7.4.11`; `openai/gpt-5.6-sol`; `medium`; qualified `XDG_DATA_HOME`      | `output/harness-hardening/kilo-code-1/kilo-code-20260721T073157Z-d26ed7f9/result.json`      |
| `pi`          |       1 | Temporary skills, private state, and run-owned processes removed; model-created `plans/` retained                 | Tooling-clean pass; 4 turns completed          | Model behavior: three missing calls and one additional `prompt-writing` call                   | None                                                                                                                     | `pi 0.80.10`; `openai-codex/gpt-5.6-sol`; `medium`; operator `PI_CODING_AGENT_DIR` | `output/harness-hardening/pi-1/pi-20260721T073608Z-1ec48bd1/result.json`                    |

## Defects And Requirements

- Codex `0.144.6` interprets `agents` as agent-role configuration and rejects
  the former Boolean `agents.enabled=false`. The stable
  `features.multi_agent=false` flag owns primary multi-agent disablement;
  `features.multi_agent_v2=false` continues to disable the v2 path.
- Normal Claude Code session resumption requires its operator-owned session
  state to remain writable across turns when the orchestrator itself is
  sandboxed.
- Cursor Agent `2026.07.16-899851b` does not expose the `grok` alias; its
  installed medium Grok route is `cursor-grok-4.5-medium`.
- OpenCode and Kilo require their qualified `XDG_DATA_HOME` values. Kilo must
  use the real versioned binary rather than the wrapper that pins its own
  configuration home.
- Pi continues to resolve authentication through the operator-owned
  `PI_CODING_AGENT_DIR`.

## Completion Summary

- Ten attempts produced four tooling failures and six final tooling-clean
  passes on the current rebuilt development CLI.
- One production defect was fixed: Codex `0.144.6` rejected the obsolete
  Boolean `agents.enabled` configuration. The runtime now disables the stable
  and v2 multi-agent feature paths through supported feature keys.
- Three invocation requirements were confirmed: normal Claude Code needs
  writable session state across turns, project-local executable routes must be
  resolved before the harness changes working directory, and Cursor requires
  the installed native Grok model identifier.
- Every final pass completed four turns with one stable resumable session and
  wrote `result.json`, `website.json`, and `events.jsonl`.
- Every final pass removed temporary skills, private runtime state, native
  disposable sessions where applicable, and run-owned processes. Model-created
  workspace files remain as run evidence.
- `gofmt`, `go vet ./cli/...`, and `go test ./cli/...` passed before the
  post-fix six-harness matrix.
- Retain this prompt and progress ledger until the fix is committed and the
  development build is deliberately promoted to the next known-good baseline.

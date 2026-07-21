# Skill Issue CLI

The Skill Issue CLI is a standalone Go executable that installs the canonical Skill Issue skills and runs blind, turn-attributed skill-calling evaluations.

## Commands

```sh
skill-issue install
skill-issue install --workspace <path> --harness <id> --scope project|user
skill-issue uninstall --workspace <path> --harness <id> --scope project|user
skill-issue doctor [--harness <id>] [--executable <path>]
skill-issue evaluate run --harness <id> [--workspace <path>] [--output <path>] [--evaluation <id>] [--executable <path>] [--model <id>] [--reasoning <level>] [--turns <number>] [--replace-preexisting-skills] [--yes] [--events] [--transcript]
skill-issue evaluate run --harness <id> [--workspace <path>] [--output <path>] [--executable <path>] [--model <id>] [--reasoning <level>] [--turns <number>] [--replace-preexisting-skills] [--yes] --skills <path> --scenario <path> --answer-sheet <path> [--events] [--transcript]
skill-issue evaluate cleanup --output <path> --run <id>
```

The minimal evaluation command is `skill-issue evaluate run --harness <id>`: without `--workspace` the CLI creates a fresh uniquely named `skill-issue-workspace-<suffix>` directory adjacent to the invocation directory and uses it (a new one per run, so every defaulted run starts clean; delete retained workspaces whenever convenient), without `--output` results collect under an adjacent `skill-issue-output` directory (each run in its own subdirectory), and without `--evaluation` the built-in `gardening-web-application` scenario runs. A named `--workspace` that does not exist yet is created. The confirmation summary marks a workspace the CLI created, and cancelling at the prompt removes that still-empty directory again.

Codex evaluations pass `--skip-git-repo-check` so non-git workspaces are supported: the operator explicitly selects and confirms the workspace, and Codex's `workspace-write` sandbox still constrains writes to it.

Ordinary install and uninstall operations retain project and user scopes. Evaluation runs are always project-local: their interface has no scope argument, and temporary skills use the selected harness's researched native project skill directory in the selected or created workspace. The CLI stores its two default result artifacts under a unique `<output>/<harness>-<UTC-timestamp>-<run-prefix>/` directory and rejects an output root inside the evaluated workspace.

Installation and evaluation support `claude-code`, `codex`, `cursor`, `opencode`, and `pi`.

### Effective evaluation configuration

Every evaluation resolves one model and reasoning value before installing temporary skills. `--model` and `--reasoning` are optional overrides; when omitted, the CLI uses these harness defaults:

| Harness       | Model                      | Reasoning | Native controls                                                            |
| ------------- | -------------------------- | --------- | -------------------------------------------------------------------------- |
| `codex`       | `gpt-5.6-sol`              | `medium`  | `--model`, `--config model_reasoning_effort=...`                           |
| `cursor`      | native Auto-select         | `medium`  | explicit `--model`; default omits the flag because Cursor owns Auto-select |
| `claude-code` | `opus`                     | `medium`  | `--model`, `--effort`                                                      |
| `opencode`    | `openai/gpt-5.6-sol`       | `medium`  | `--model`, `--variant`                                                     |
| `pi`          | `openai-codex/gpt-5.6-sol` | `medium`  | `--model`, `--thinking`                                                    |

Model identifiers and supported reasoning values are passed to the selected native harness. Skill Issue does not maintain a model catalogue or prevalidate compatibility; a native harness owns rejection of an unsupported value. Cursor uses its native Auto-select model and model-native reasoning when no model override is supplied. An explicit Cursor `--reasoning` override is rejected before evaluation side effects because the CLI exposes no independent reasoning control.

Use `--executable` when the required harness command is intentionally absent from `PATH`. A value containing a path separator is resolved to an absolute path against the invocation working directory at parse time, because the evaluator later launches the harness from a run-owned directory where a caller-relative path would fail; a bare command name continues to resolve through `PATH`. The pre-run summary displays the selected executable or launcher. The local qualification environments use this for the project-local Cursor agent, the Claude Code launcher that owns its local Codex proxy, the qualified OpenCode executable, and Pi's installed runtime entrypoint.

If a harness rejects a model, reasoning value, session, permission, or protocol step, the CLI reports a tooling error. Failed runs also write the minimal sanitized `failure.json` record described under Evaluation artifacts.

Use the optional positive integer `--turns` argument to run only that many turns from the beginning of any built-in or custom scenario. Omitting it runs the complete scenario. A value above the scenario length runs every available turn. Scenario truncation also removes later-turn expectations from the active answer sheet, so unrun turns are not reported as missing calls.

Before an evaluation run creates output or private state, the CLI prints the selected evaluation, effective turn count, available turn count, harness, model, reasoning, workspace, output root, any executable override, and custom input paths. If the requested turn count exceeds the scenario length, the summary shows both the requested and effective values. Confirm with `y` to continue or cancel with `n` (the default); cancellation exits cleanly without starting evaluation. Scripted callers pass `--yes` to accept the confirmation after the summary is printed; `--yes` covers only the pre-run confirmation, so replacing differing preexisting skills still requires `--replace-preexisting-skills`.

If the evaluation skill destination already contains a skill directory with a canonical Skill Issue name, the run backs the directory up under its private run state before replacing it and restores the backed-up content byte-for-byte during cleanup. When the preexisting content matches the canonical skill, the run proceeds silently. When it differs, the interactive CLI lists the differing skills and asks for confirmation before replacing them; non-interactive callers must pass `--replace-preexisting-skills` to accept the temporary replacement, and the run otherwise stops before any side effects on those skills.

Custom evaluation results are caller-selected local evidence. `result.json`, `website.json`, and optional diagnostic artifacts remain under the selected output root and are not presented as website evidence without separate review and acceptance.

### Doctor

`skill-issue doctor` diagnoses evaluation readiness in seconds with zero model cost. Without `--harness` it checks every supported harness; with `--harness` it checks one, and `--executable` selects the same override a run would use. Doctor verifies the platform, temporary-directory and home-directory preconditions, then per harness: executable resolution, the installed version against the tested version recorded in the harness registry (a pinned OpenCode mismatch fails exactly as a run would, honoring `SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS=1`; other harnesses warn on drift), authentication through each harness's native status surface (`claude auth status`, `codex login status`, Cursor `status`, OpenCode `auth list` plus default-model availability, and the Pi agent directory), and for Codex that the installed binary parses the exact generated evaluation configuration (`codex -c <key>=<value> login status`). Findings stream as one human-readable line each on standard error, the structured report is written to standard output, and the exit code is nonzero when any check fails. Doctor probes run with the caller's normal environment; it is a preflight, not a substitute for a one-turn smoke.

### Skill-visibility verification

A run must never complete with the governed skills silently unloaded, because a missing call caused by an invisible skill is indistinguishable from a genuine model choice. Where a harness exposes visibility evidence, the evaluator verifies it during the run and reports a tooling error otherwise: OpenCode must list every installed evaluation skill through its native `debug skill` discovery after installation and before turn 1; Claude Code must list every installed evaluation skill in each turn's `system/init` event; Pi must expose every supplied skill through `get_commands` before turn 1. Codex and Cursor expose no equivalent listing surface, so their visibility remains corroborated by observed instrumentation rather than proven per run: Codex attribution flows through structured command events, and Cursor signal attempts that record no marker fail the run.

## Installation

The executable embeds each canonical skill from the repository-relative source declared in its payload manifest and validates referenced-file closure before installing. Scenario-used supporting skills live beside their retained evaluation evidence under `evaluations/scenario-skill-refinement/<skill>/skill/`; lifecycle skills remain under `skills/`, and Dictate Plan remains under `supporting-skills/`. Installation creates or reuses the selected harness's researched native project or user skill root, then replaces only the known Skill Issue payload directories with the current embedded copies. Portable skill files are shared across harnesses, and each harness specification owns the installed metadata it supports. Codex installations include each skill's `agents/openai.yaml`; Claude Code, Cursor, and Pi add `disable-model-invocation: true` to the installed `skill-intake` frontmatter. Custom evaluation skills retain their supplied frontmatter. Unrelated files and skill directories are not inspected or changed.

Run `skill-issue install` in an interactive terminal for guided installation. Use the up and down arrow keys to select Claude Code, OpenAI Codex, Cursor, OpenCode, or Pi and then select project or user scope. Project scope uses the current working directory. Before writing anything, the CLI previews the harness, scope, native destination, and eleven embedded skills and asks for confirmation. It does not detect or automatically choose a harness.

After confirmation, installation verifies that every expected harness-specific payload file exists. The CLI then directs the user to open the selected harness and explicitly invoke `skill-intake`. Its canonical description says, "Use only when the user explicitly requests this skill." Codex enforces that boundary through `allow_implicit_invocation: false`; Claude Code, Cursor, and Pi enforce it through `disable-model-invocation: true`. The argument-driven form remains available for scripted installation and does not prompt.

The payload is disposable and contains no user configuration or mutable application data. Running `install` again is the reinstall and update path. `uninstall` removes the same known Skill Issue payload directories directly. The CLI does not create an ownership receipt, backup, rollback inventory, or platform application-state directory.

## Blind Evaluation

### Built-in governed evaluations

A built-in evaluation is one embedded unit containing its ordered scenario turns and matching private answer sheet. Selecting one identifier loads both parts from the standalone executable; callers do not provide or select separate files.

The built-in identifiers are:

- `gardening-web-application`
- `community-archive-desktop-application`
- `neighborhood-emergency-preparedness-program`

Each contains one complete 30-turn governed scenario and its matching expected invocations. Turns 13, 18, and 24 intentionally contain small factual reminders with no expected invocation; the remaining governed points continue through the final turn. The gardening and archive scenarios each contain 46 expected invocations; the preparedness scenario contains 45. For example:

```sh
skill-issue evaluate run \
  --workspace <path> \
  --output <path> \
  --harness <id> \
  --turns 10 \
  --evaluation gardening-web-application
```

Every embedded prompt is the complete text sent to the harness. It does not depend on the surrounding human-readable Markdown scenario documents.

### Custom evaluations

Custom evaluations use the same installation, replay, attribution, comparison, evidence, and cleanup path. Supply a skill directory, scenario, and answer sheet together, and do not combine them with `--evaluation`:

```sh
skill-issue evaluate run \
  --workspace <path> \
  --output <path> \
  --harness <id> \
  --skills /outside/workspace/skills \
  --scenario /outside/workspace/scenario.json \
  --answer-sheet /outside/workspace/answer-sheet.json
```

The `--skills` path is one directory whose direct children are the supplied skill directories. Each supplied skill must contain `SKILL.md`; its frontmatter `name` must match the directory name, and every backtick-marked local `references/`, `scripts/`, or `assets/` path in that entrypoint must resolve within the skill directory. The directory must contain at least one skill and no non-directory entries. Skill Issue validates this runnable structure before evaluation.

Before a custom evaluation starts, the CLI warns that scenario and answer-sheet inputs must not contain personal, confidential, or sensitive information. Skill Issue does not scan or redact those files. The harness conversation may retain their complete contents. The answer sheet must use the scenario's ID, existing turn IDs, and supplied skill names. The CLI validates those structural links and leaves the semantic choice of expected calls to the evaluator; an incorrect answer-sheet key can make the result invalid or misleading.

Scenario JSON uses this shape:

```json
{
  "schema_version": 1,
  "scenario_id": "example",
  "turns": [
    {
      "turn_id": "turn-1",
      "prompt": "The complete prompt sent for this turn."
    }
  ]
}
```

Answer-sheet JSON uses this shape:

```json
{
  "schema_version": 1,
  "scenario_id": "example",
  "expected": [
    {
      "turn_id": "turn-1",
      "skill": "prompt-writing"
    }
  ]
}
```

The scenario IDs must match, every expected turn must exist, and every expected skill must be present in the selected skill set. Keep the custom answer-sheet file outside the evaluated workspace. Built-in answer sheets remain inside the executable and are loaded only into CLI-owned private run state; neither form is copied into the workspace.

### Neutral instrumentation

`evaluate run` keeps the answer sheet, token map, active turn, internal event state, installation state, and user-facing evidence under the required output root outside the evaluated workspace. It generates disposable copies of the current canonical skills and inserts one neutral instruction immediately after frontmatter:

```text
Run "<absolute-cli-path>" signal "<opaque-token>" "<absolute-output-state-path>", then continue normally.
```

Codex instead receives a capture-only marker that has no filesystem effect:

```text
Run echo "<opaque-token>", then continue normally.
```

For other harnesses, the `signal` command succeeds silently. The opaque token in either form has meaning only in the output-owned private state and reveals no skill identity, expected call, answer-sheet content, turn number, or scoring rule. The state path lets a separate signal process find that private run without relying on a user-level installation. Temporary skills retain their canonical names, frontmatter, supporting files, and bodies, and they occupy only the harness's normal project skill root.

The CLI starts one primary-agent session, sends scenario prompts verbatim and in order, and owns turn boundaries independently of the agent. It stores signals outside an active turn as unattributed tooling evidence. After replay, it writes the detailed evidence and compact website output described below. Cleanup rematerializes canonical Skill Issue skills that existed before instrumentation, removes evaluation paths that did not exist, and deletes private token mappings. It never backs up installed skill contents.

These measures minimize evaluation clues. They cannot guarantee that an agent inspecting its environment, executable, process activity, or generated skill files will never infer that instrumentation exists.

### Codex runtime isolation

For a Codex evaluation, the CLI creates a run-owned `CODEX_HOME`, copies the user's existing `auth.json` into it when present, and removes the private home after cleanup. Codex supplies built-in defaults without a `config.toml`; the CLI launches every initial and resumed turn with `--ignore-user-config` and `--ignore-rules` and passes the effective model, reasoning, `workspace-write`, and other evaluation settings explicitly. It disables multi-agent tools through `features.multi_agent=false` and `features.multi_agent_v2=false`, excludes discovered `AGENTS.md` content, disables apps and the Codex plugin feature, and supplies a temporary `skills.config` deny-list for every skill found under the user's Codex skill root, `$HOME/.agents/skills`, and `/etc/codex/skills`. The disposable project skills installed by Skill Issue remain enabled because they live only in the evaluated workspace.

Codex runs with `approval_policy=on-request` and `approvals_reviewer=auto_review`. Actions already allowed by `workspace-write` proceed directly. Requests to cross that boundary are reviewed by Codex's guardian agent, which does not widen the sandbox and can consume additional model usage. Codex attribution uses the structured command event emitted for the capture-only marker, so skill attribution does not require a process to write outside the workspace. Other denied, aborted, timed-out, or malformed boundary requests remain tooling failures.

Skill Issue does not edit or move the user's Codex configuration, credentials, skills, plugins, rules, or instruction files. The copied credential, model cache, SQLite state, and resumable evaluation session stay in the run-owned Codex home and are deleted during cleanup. Organization-managed requirements and system-level configuration remain authoritative and may make a Codex environment unsuitable for the campaign.

### OpenCode runtime isolation

OpenCode evaluation support is qualified against OpenCode `1.18.4`. The CLI requires that exact version (set `SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS=1` to downgrade a version mismatch to a warning), a pre-existing native provider login, and availability of the selected `provider/model` before temporary skills are installed. The default route uses OpenCode's native OpenAI ChatGPT OAuth provider with `openai/gpt-5.6-sol` and the `medium` variant.

Each run creates private XDG configuration, state, cache, and temporary roots while retaining the operator-owned XDG data root that contains native authentication. The generated OpenCode configuration disables sharing, updates, snapshots, formatters, LSP, external plugins, MCP configuration, added instructions, project configuration, Claude compatibility, external skill discovery, and the file watcher. Compiled internal plugins remain available because OpenCode's native OpenAI OAuth provider depends on them.

The deny-first permission policy permits workspace reads and edits, the selected evaluation skills, and only the exact Skill Issue signal command through Bash. OpenCode receives every prompt through `run --pure --format json` with the selected model and variant. Later turns resume the captured camel-case `sessionID`; each turn must finish with `step_finish` and `reason: "stop"`. An attempted instrumentation signal whose marker is never recorded for its turn is a tooling failure. A denied compound command that chains the signal with another action is model behavior when the model retries the exact signal and its marker is recorded, as is any unrelated denied tool request.

On completion, cancellation, or timeout, the adapter terminates only its owned process group and deletes the native OpenCode session before removing run-owned configuration and temporary skill material. The operator-owned OAuth credential remains available to OpenCode and may refresh through OpenCode's normal native behavior.

### Evaluation artifacts

Every tooling-complete evaluation stores these files under `<output>/<harness>-<UTC-timestamp>-<run-prefix>/`, where `<output>` is the mandatory directory selected by the caller:

- `result.json` is the detailed authoritative evidence. It identifies the selected evaluation, harness, effective model, and effective reasoning, and retains the expected, observed, missing, additional, and unattributed skill calls. It includes `transcript_path` only when `--transcript` is enabled.
- `website.json` provides the compact turn-level data consumed by the later Recharts website work.

Two diagnostic artifacts are optional and default to off:

- `--events` writes `events.jsonl`, containing the raw recorded skill-invocation events used to derive the detailed result. It does not contain prompts, responses, commands, paths, opaque tokens, or environment data.
- `--transcript` writes `transcript.json`, containing only ordered conversation turns. Each turn has `turn_id`, `user`, and `assistant`; raw replay captures, harness events, commands, command output, errors, tool calls, reasoning, session identifiers, and transport metadata are discarded before serialization. Known local identifiers and common sensitive patterns are replaced with bracketed placeholders. Runtime execution and scoring continue to use the original capture. The CLI warns that arbitrary contextual personal or confidential content may remain and that the transcript must be reviewed before sharing.

Tooling-failed runs write `failure.json` with exactly `schema_version`, `run_id`, `harness`, optional `turn_id`, `failed_at`, and a sanitized `error` summary. Commands, stdout, stderr, model configuration, browser policy, replay events, and conversation content are not stored in the failure record.

Warnings are written to standard error so the CLI's structured standard output remains machine-readable. Enabling neither flag produces only `result.json` and `website.json`; either flag may be enabled independently, and both may be enabled together.

`website.json` has this exact shape:

```json
{
  "schema_version": 2,
  "run_id": "<run ID>",
  "scenario_id": "<scenario ID>",
  "harness": "<harness ID>",
  "model": "<model ID>",
  "total_turns": 30,
  "points": [
    {
      "turn": 1,
      "turn_id": "turn-1",
      "called": 2,
      "missed": 0,
      "unexpected": 1
    }
  ]
}
```

The metadata comes from the same evaluated result and scenario used for `result.json`. `total_turns` is the number of ordered scenario turns. A point exists when its turn has at least one expected or unexpected call. `turn` is that turn's one-based position in the scenario, while `turn_id` preserves the source identifier even when it is nonnumeric. `called` counts unique expected skills observed on their expected turn, `missed` counts unique expected skills absent there, and `unexpected` counts unique additional skills observed on that turn. Repeated signal events for the same skill and turn do not inflate any count. Sample size is derived by summing `called + missed + unexpected` and is not stored separately.

The compact artifact does not duplicate skill names, unattributed calls, transcripts, or raw events. Its `unexpected` count is the compact projection of detailed `additional` calls. Detailed interpretation continues to use `result.json` and any optional diagnostic artifacts the evaluator deliberately enabled.

## Output Storage and Recovery

Every evaluation stores its outputs beneath the mandatory output root:

- `<output>/<harness>-<UTC-timestamp>-<run-prefix>/` always contains `result.json` and `website.json`; it contains `events.jsonl` or `transcript.json` only when their corresponding flags are enabled.
- `<output>/.skill-issue/` temporarily contains private run state, opaque-token mappings, internal events, and the names of Skill Issue paths that existed before instrumentation.
- Harness session state follows each adapter's native contract. OpenCode sessions are deleted during cleanup; other supported harness sessions follow the retained behavior documented above.

Private run state and token mappings use restrictive permissions. A successful run removes its private recovery state after cleanup. If a run is interrupted, use the same output root with `skill-issue evaluate cleanup --output <path> --run <id>` to rematerialize previously present canonical Skill Issue paths, remove temporary evaluation paths, and delete that run's private state. Completed result artifacts remain in the caller-selected output directory. Evaluation does not write run state to the platform application-state location.

## Harness Boundary

The runner uses headless resumable commands for Codex, Cursor, and OpenCode, Claude streaming JSON, and Pi RPC mode. Model access and the underlying harness login remain user-provided prerequisites. Every supported harness has CLI-owned runtime preparation. A launch, required permission, session, marker, timeout, cancellation, or protocol failure is a tooling failure rather than a model result.

## Development

```sh
go vet ./cli/...
go run ./cli/cmd/skill-issue help
./cli/scripts/build-cross-platform.sh
```

Final public-interface regression and native qualification coverage is deferred until the CLI interface stabilizes.

The build script produces standalone macOS, Windows, and Linux binaries for AMD64 and Arm64. Cross-compilation does not qualify native runtime behavior.

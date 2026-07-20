# Skill Issue CLI

The Skill Issue CLI is a standalone Go executable that installs the canonical Skill Issue skills and runs blind, turn-attributed skill-calling evaluations.

## Commands

```sh
skill-issue install
skill-issue install --workspace <path> --harness <id> --scope project|user
skill-issue uninstall --workspace <path> --harness <id> --scope project|user
skill-issue evaluate run --workspace <path> --output <path> --harness <id> [--executable <path>] [--model <id>] [--reasoning <level>] [--turns <number>] --evaluation gardening-web-application [--events] [--transcript]
skill-issue evaluate run --workspace <path> --output <path> --harness <id> [--executable <path>] [--model <id>] [--reasoning <level>] [--turns <number>] --skills <path> --scenario <path> --answer-sheet <path> [--events] [--transcript]
skill-issue evaluate cleanup --output <path> --run <id>
```

Ordinary install and uninstall operations retain project and user scopes. Evaluation runs are always project-local: their interface has no scope argument, `--workspace` is required, and temporary skills use the selected harness's researched native project skill directory. Every evaluation also requires `--output`; the CLI stores its two default result artifacts under a unique `<output>/<harness>-<UTC-timestamp>-<run-prefix>/` directory and rejects an output root inside the evaluated workspace.

The product targets `claude-code`, `codex`, `cursor`, `opencode`, `kilo-code`, and `pi`. Installation and evaluation currently support `claude-code`, `codex`, `cursor`, and `pi`; OpenCode and Kilo Code remain in progress.

### Effective evaluation configuration

Every evaluation resolves one model and reasoning value before installing temporary skills. `--model` and `--reasoning` are optional overrides; when omitted, the CLI uses these harness defaults:

| Harness       | Model                      | Reasoning | Native controls                                                            |
| ------------- | -------------------------- | --------- | -------------------------------------------------------------------------- |
| `codex`       | `gpt-5.6-sol`              | `medium`  | `--model`, `--config model_reasoning_effort=...`                           |
| `cursor`      | native Auto-select         | `medium`  | explicit `--model`; default omits the flag because Cursor owns Auto-select |
| `claude-code` | `opus`                     | `medium`  | `--model`, `--effort`                                                      |
| `pi`          | `openai-codex/gpt-5.6-sol` | `medium`  | `--model`, `--thinking`                                                    |

Model identifiers and supported reasoning values are passed to the selected native harness. Skill Issue does not maintain a model catalogue or prevalidate compatibility; a native harness owns rejection of an unsupported value. Cursor uses its native Auto-select model and model-native reasoning when no model override is supplied. An explicit Cursor `--reasoning` override is rejected before evaluation side effects because the CLI exposes no independent reasoning control.

Use `--executable` when the required harness command is intentionally absent from `PATH`. The pre-run summary displays the selected executable or launcher. The local qualification environment uses this for the project-local Cursor agent, the Claude Code launcher that owns its local Codex proxy, and Pi's installed runtime entrypoint.

If a harness rejects a model, reasoning value, session, permission, or protocol step, the CLI reports a tooling error and retains the useful native stderr or structured error text.

Use the optional positive integer `--turns` argument to run only that many turns from the beginning of any built-in or custom scenario. Omitting it runs the complete scenario. A value above the scenario length runs every available turn. Scenario truncation also removes later-turn expectations from the active answer sheet, so unrun turns are not reported as missing calls.

Before an evaluation run creates output or private state, the CLI prints the selected evaluation, effective turn count, available turn count, harness, model, reasoning, workspace, output root, any executable override, and custom input paths. If the requested turn count exceeds the scenario length, the summary shows both the requested and effective values. Confirm with `y` to continue or cancel with `n` (the default); cancellation exits cleanly without starting evaluation.

Custom evaluation results are caller-selected local evidence. `result.json`, `website.json`, and optional diagnostic artifacts remain under the selected output root and are not presented as website evidence without separate review and acceptance.

## Installation

The executable embeds the complete canonical `skills/` and `supporting-skills/` trees and validates referenced-file closure before installing. Installation creates or reuses the selected harness's researched native project or user skill root, then replaces only the known Skill Issue payload directories with the current embedded copies. Portable skill files are shared across harnesses, and each harness specification owns the installed metadata it supports. Codex installations include each skill's `agents/openai.yaml`; Claude Code, Cursor, and Pi add `disable-model-invocation: true` to the installed `skill-intake` frontmatter. The canonical source remains portable, and custom evaluation skills retain their supplied frontmatter. Unrelated files and skill directories are not inspected or changed.

Run `skill-issue install` in an interactive terminal for guided installation. Use the up and down arrow keys to select Claude Code, OpenAI Codex, Cursor, or Pi and then select project or user scope. OpenCode and Kilo Code appear as disabled in-progress targets. Project scope uses the current working directory. Before writing anything, the CLI previews the harness, scope, native destination, and seven embedded skills and asks for confirmation. It does not detect or automatically choose a harness.

After confirmation, installation verifies that every expected harness-specific payload file exists. The CLI then directs the user to open the selected harness and explicitly invoke `skill-intake`. Its canonical description says, "Use only when the user explicitly requests this skill." Codex enforces that boundary through `allow_implicit_invocation: false`; Claude Code, Cursor, and Pi enforce it through `disable-model-invocation: true`. The argument-driven form remains available for scripted installation and does not prompt.

The payload is disposable and contains no user configuration or mutable application data. Running `install` again is the reinstall and update path. `uninstall` removes the same known Skill Issue payload directories directly. The CLI does not create an ownership receipt, backup, rollback inventory, or platform application-state directory.

## Blind Evaluation

### Built-in governed evaluations

A built-in evaluation is one embedded unit containing its ordered scenario turns and matching private answer sheet. Selecting one identifier loads both parts from the standalone executable; callers do not provide or select separate files.

The built-in identifiers are:

- `gardening-web-application`
- `community-archive-desktop-application`
- `neighborhood-emergency-preparedness-program`

Each contains one complete 30-turn scenario and its required first-activation calls. For example:

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

The `signal` command succeeds silently. Its opaque token has meaning only in the output-owned private state and reveals no skill identity, expected call, answer-sheet content, turn number, or scoring rule. The state path lets the separate signal process find that private run without relying on a user-level installation. Temporary skills retain their canonical names, frontmatter, supporting files, and bodies, and they occupy only the harness's normal project skill root.

The CLI starts one primary-agent session, sends scenario prompts verbatim and in order, and owns turn boundaries independently of the agent. It stores signals outside an active turn as unattributed tooling evidence. After replay, it writes the detailed evidence and compact website output described below. Cleanup rematerializes canonical Skill Issue skills that existed before instrumentation, removes evaluation paths that did not exist, and deletes private token mappings. It never backs up installed skill contents.

These measures minimize evaluation clues. They cannot guarantee that an agent inspecting its environment, executable, process activity, or generated skill files will never infer that instrumentation exists.

### Codex runtime isolation

For a Codex evaluation, the CLI keeps the user's existing Codex home available for supported authentication but launches every initial and resumed turn with `--ignore-user-config` and `--ignore-rules`. It supplies the effective model, reasoning, `workspace-write`, and all other evaluation settings explicitly. It also excludes discovered `AGENTS.md` content, disables apps and the Codex plugin feature, and supplies a temporary `skills.config` deny-list for every skill found under the user's Codex skill root, `$HOME/.agents/skills`, and `/etc/codex/skills`. The disposable project skills installed by Skill Issue remain enabled because they live only in the evaluated workspace.

Codex runs with `approval_policy=on-request` and `approvals_reviewer=auto_review`. Actions already allowed by `workspace-write` proceed directly. Requests to cross that boundary are reviewed by Codex's guardian agent, which does not widen the sandbox and can consume additional model usage. Codex attribution uses the existing structured command event emitted when the agent attempts the opaque signal command, so the signal does not need permission to write the private state itself. Other denied, aborted, timed-out, or malformed boundary requests remain tooling failures.

Skill Issue does not edit, move, or copy the user's Codex configuration, credentials, skills, plugins, rules, or instruction files. `--ignore-user-config` continues to use the user's normal Codex authentication. The resumable evaluation session is created and retained in Codex's normal session store. Organization-managed requirements and system-level configuration remain authoritative and may make a Codex environment unsuitable for the campaign.

### Evaluation artifacts

Every tooling-complete evaluation stores these files under `<output>/<harness>-<UTC-timestamp>-<run-prefix>/`, where `<output>` is the mandatory directory selected by the caller:

- `result.json` is the detailed authoritative evidence. It identifies the selected evaluation, harness, effective model, and effective reasoning, and retains the expected, observed, missing, additional, and unattributed skill calls. It includes `transcript_path` only when `--transcript` is enabled.
- `website.json` provides the compact turn-level data consumed by the later Recharts website work.

Two diagnostic artifacts are optional and default to off:

- `--events` writes `events.jsonl`, containing the raw recorded skill-invocation events used to derive the detailed result. It does not contain prompts, responses, commands, paths, opaque tokens, or environment data.
- `--transcript` writes `transcript.json`, containing the scenario and sanitized replay capture, including prompts, responses, commands, command output, errors, and structured harness events. Before persistence, Skill Issue replaces the resolved CLI, workspace, output, private state, runtime, temporary, and home paths plus known local user and host identifiers with bracketed placeholders. Runtime execution and scoring continue to use the original absolute paths. The CLI warns that arbitrary personal or confidential content intentionally present in prompts or responses may remain and that the transcript must be reviewed before sharing.

Warnings are written to standard error so the CLI's structured standard output remains machine-readable. Enabling neither flag produces only `result.json` and `website.json`; either flag may be enabled independently, and both may be enabled together.

`website.json` has this exact shape:

```json
{
  "schema_version": 1,
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
      "missed": 0
    }
  ]
}
```

The metadata comes from the same evaluated result and scenario used for `result.json`. `total_turns` is the number of ordered scenario turns. A point exists only when its turn has at least one expected call. `turn` is that turn's one-based position in the scenario, while `turn_id` preserves the source identifier even when it is nonnumeric. `called` counts unique expected skills observed on their expected turn, and `missed` counts unique expected skills absent there; repeated signal events do not inflate either value. Sample size is derived by summing `called + missed` and is not stored separately.

The compact artifact does not duplicate skill names, additional calls, unattributed calls, transcripts, or raw events. The later website uses numeric `turn` values for its horizontal axis and `called` and `missed` as its Recharts series; detailed interpretation continues to use `result.json` and any optional diagnostic artifacts the evaluator deliberately enabled.

## Output Storage and Recovery

Every evaluation stores its outputs beneath the mandatory output root:

- `<output>/<harness>-<UTC-timestamp>-<run-prefix>/` always contains `result.json` and `website.json`; it contains `events.jsonl` or `transcript.json` only when their corresponding flags are enabled.
- `<output>/.skill-issue/` temporarily contains private run state, opaque-token mappings, internal events, and the names of Skill Issue paths that existed before instrumentation.
- Harness session state used during replay remains in the harness's normal history.

Private run state and token mappings use restrictive permissions. A successful run removes its private recovery state after cleanup. If a run is interrupted, use the same output root with `skill-issue evaluate cleanup --output <path> --run <id>` to rematerialize previously present canonical Skill Issue paths, remove temporary evaluation paths, and delete that run's private state. Completed result artifacts remain in the caller-selected output directory. Evaluation does not write run state to the platform application-state location.

## Harness Boundary

The runner uses headless resumable commands for Codex and Cursor, Claude streaming JSON, and Pi RPC mode. Model access and the underlying harness login remain user-provided prerequisites. Codex, Cursor, Claude Code, and Pi have CLI-owned runtime preparation. A launch, required permission, session, marker, timeout, cancellation, or protocol failure is a tooling failure rather than a model result.

OpenCode and local Kilo Code remain product targets, but their installation and evaluation routes are not yet implemented. Kilo Cloud remains outside the product boundary.

## Local Smoke Qualification

On 2026-07-20, the standalone CLI completed two-turn built-in and custom-skill smokes through Codex `0.144.1`, Cursor Agent `2026.07.16-899851b`, Claude Code `2.1.205` through the project-local Codex proxy, and Pi `0.80.10`. The runs exercised defaults, explicit model and reasoning values where supported, executable overrides, supplied skills, scenarios, answer sheets, output selection, one continuous session, result generation, selected-workspace writes, and cleanup. Cursor correctly rejects an explicit reasoning override because its CLI exposes no independent reasoning flag. Missing expected skill calls remain evaluation data rather than tooling failures.

The final Cursor and Claude runs left no detached Cursor worker, Claude proxy, or private runtime. Pi reused its existing authenticated `PI_CODING_AGENT_DIR` without Skill Issue copying or overwriting it while keeping evaluation sessions and supplied skills temporary. Qualification results, observed calls, cleanup evidence, and useful native errors are recorded in `evaluations/skill-calling/smoke/real-harness-smoke-report.md`.

## Development

```sh
go vet ./cli/...
go run ./cli/cmd/skill-issue help
./cli/scripts/build-cross-platform.sh
```

Final public-interface regression and native qualification coverage is deferred until the CLI interface stabilizes.

The build script produces standalone macOS, Windows, and Linux binaries for AMD64 and Arm64. Cross-compilation does not qualify native runtime behavior.

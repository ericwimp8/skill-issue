# Skill Issue CLI

The Skill Issue CLI is a standalone Go executable that installs the canonical Skill Issue skills and runs blind, turn-attributed skill-calling evaluations.

## Commands

```sh
skill-issue install --workspace <path> --harness <id> --scope project|user
skill-issue verify --workspace <path> --harness <id> --scope project|user
skill-issue repair --workspace <path> --harness <id> --scope project|user
skill-issue update --workspace <path> --harness <id> --scope project|user
skill-issue uninstall --workspace <path> --harness <id> --scope project|user
skill-issue evaluate run --workspace <path> --output <path> --harness <id> --model <id> --evaluation gardening-web-application
skill-issue evaluate run --workspace <path> --output <path> --harness <id> --model <id> --scenario <path> --answer-sheet <path>
skill-issue evaluate cleanup --run <id>
```

Ordinary install, verify, repair, update, and uninstall operations retain project and user scopes. Evaluation runs are always project-local: their interface has no scope argument, `--workspace` is required, and temporary skills use the selected harness's researched native project skill directory. Every evaluation also requires `--output`; the CLI stores its four user-facing artifacts under `<output>/<run-id>/` and rejects an output root inside the evaluated workspace.

Supported harness identifiers are `copilot`, `claude-code`, `codex`, `cursor`, `gemini-cli`, `grok-build`, `opencode`, `kilo-code`, and `pi`.

## Installation Ownership

The executable embeds the complete canonical `skills/` and `supporting-skills/` trees. It validates referenced-file closure before installing. Every installation is staged, recorded in an external ownership receipt, and limited to the selected harness's researched native project or user root.

The CLI refuses a same-name destination it does not own. Update and repair replace receipt-owned content. Verify checks the recorded digest. Uninstall removes only receipt-owned, unchanged content.

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
  --model <id> \
  --evaluation gardening-web-application
```

Every embedded prompt is the complete text sent to the harness. It does not depend on the surrounding human-readable Markdown scenario documents.

### Custom evaluations

Custom evaluations use the same installation, replay, attribution, comparison, evidence, and cleanup path. Supply both files together and do not combine them with `--evaluation`:

```sh
skill-issue evaluate run \
  --workspace <path> \
  --output <path> \
  --harness <id> \
  --model <id> \
  --scenario /outside/workspace/scenario.json \
  --answer-sheet /outside/workspace/answer-sheet.json
```

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

The scenario IDs must match, every expected turn must exist, and every expected skill must be present in the embedded evaluation payload. Keep custom scenario and answer-sheet files outside the evaluated workspace. Built-in answer sheets remain inside the executable and are loaded only into CLI-owned private run state; neither form is copied into the workspace.

### Neutral instrumentation

`evaluate run` keeps the answer sheet, token map, active turn, internal event state, and restoration data in CLI-owned application storage outside the evaluated workspace. It writes the user-facing evidence to the required external output root. It generates disposable copies of the current canonical skills and inserts one neutral instruction immediately after frontmatter:

```text
Run "<absolute-cli-path>" signal "<opaque-token>", then continue normally.
```

The `signal` command succeeds silently. Its opaque token has meaning only in external CLI state and reveals no skill identity, expected call, answer-sheet content, turn number, or scoring rule. Temporary skills retain their canonical names, frontmatter, supporting files, and bodies, and they occupy only the harness's normal project skill root.

The CLI starts one primary-agent session, sends scenario prompts verbatim and in order, and owns turn boundaries independently of the agent. It stores signals outside an active turn as unattributed tooling evidence. After replay, it writes the detailed evidence and compact website output described below. It then restores the exact prior project installation and deletes private token mappings.

These measures minimize evaluation clues. They cannot guarantee that an agent inspecting its environment, executable, process activity, or generated skill files will never infer that instrumentation exists.

### Evaluation artifacts

Every tooling-complete evaluation stores these files together under `<output>/<run-id>/`, where `<output>` is the mandatory directory selected by the caller:

- `result.json` remains the detailed authoritative evidence. It retains the expected, observed, missing, additional, and unattributed skill calls plus the transcript reference.
- `events.jsonl` retains the raw recorded signal events.
- `transcript.json` retains the native replay transcript and structured harness events available to the runner.
- `website.json` provides the compact turn-level data consumed by the later Recharts website work.

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

The compact artifact does not duplicate skill names, additional calls, unattributed calls, transcripts, or raw events. The later website uses numeric `turn` values for its horizontal axis and `called` and `missed` as its Recharts series; detailed interpretation and audit continue to use `result.json` and its supporting evidence.

## Private Storage and Recovery

Private state uses the platform application-state location:

- macOS: `~/Library/Application Support/Skill Issue`
- Windows: `%LOCALAPPDATA%\Skill Issue`
- Linux: `$XDG_STATE_HOME/skill-issue` or `~/.local/state/skill-issue`

Private run state and token mappings use restrictive permissions in the platform application-state location. If a run is interrupted, use the run identifier with `skill-issue evaluate cleanup --run <id>` to restore or remove its temporary installation. Cleanup requires the evaluation receipt to belong to that run. Completed evidence remains in the caller-selected output directory after cleanup.

## Harness Boundary

The runner uses headless resumable commands for Copilot, Codex, Cursor, Gemini CLI, and Grok Build; Claude streaming JSON; OpenCode and Kilo programmatic run sessions; and Pi RPC mode. Authentication, model access, workspace trust, and command permissions remain user-provided prerequisites. A launch, permission, session, marker, timeout, cancellation, or protocol failure is a tooling failure rather than a model result.

The adapters are implemented from the direct-install and automation contracts but remain unqualified on harnesses unavailable locally. Antigravity automation is outside this runner; Gemini CLI is the Google target. Grok Build and Kilo retain the caveats recorded in the direct-install research.

## Development

```sh
go vet ./cli/...
go run ./cli/cmd/skill-issue help
go run ./cli/cmd/skill-issue payload
./cli/scripts/build-cross-platform.sh
```

The automated CLI test suite is intentionally deferred until the product interfaces stabilize. Final regression coverage will be created during qualification against the completed installation, lifecycle, replay, evidence, recovery, and adapter behavior.

The build script produces standalone macOS, Windows, and Linux binaries for AMD64 and Arm64. Cross-compilation does not qualify native runtime behavior.

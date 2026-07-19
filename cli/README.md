# Skill Issue CLI

The Skill Issue CLI is a standalone Go executable that installs the canonical Skill Issue skills and runs blind, turn-attributed skill-calling evaluations.

## Commands

```sh
skill-issue install --workspace <path> --harness <id> --scope project|user
skill-issue verify --workspace <path> --harness <id> --scope project|user
skill-issue repair --workspace <path> --harness <id> --scope project|user
skill-issue update --workspace <path> --harness <id> --scope project|user
skill-issue uninstall --workspace <path> --harness <id> --scope project|user
skill-issue evaluate run --workspace <path> --harness <id> --model <id> --scenario <path> --answer-sheet <path> --scope project|user
skill-issue evaluate cleanup --run <id>
```

`mark <opaque-token>` is an internal silent command used only by temporary evaluation copies. For user-scope operations, `--workspace` identifies the evaluated working directory while installation uses the harness's native user root.

Supported harness identifiers are `copilot`, `claude-code`, `codex`, `cursor`, `gemini-cli`, `grok-build`, `opencode`, `kilo-code`, and `pi`.

## Installation Ownership

The executable embeds the complete canonical `skills/` and `supporting-skills/` trees. It validates referenced-file closure before installing. Every installation is staged, recorded in an external ownership receipt, and limited to the selected harness's researched native project or user root.

The CLI refuses a same-name destination it does not own. Update and repair replace receipt-owned content. Verify checks the recorded digest. Uninstall removes only receipt-owned, unchanged content.

## Blind Evaluation

`evaluate run` keeps the scenario, answer sheet, token map, active turn, events, restoration data, and results outside the evaluated workspace. It generates disposable copies of the current canonical skills and inserts one neutral instruction immediately after frontmatter:

```text
Run "<absolute-cli-path>" mark "<opaque-token>", then continue normally.
```

The CLI starts one primary-agent session, sends scenario prompts verbatim and in order, and owns turn boundaries independently of the agent. It stores markers outside an active turn as unattributed tooling evidence. After replay, it writes native transcript data once and exports compact descriptive JSON with expected, observed, missing, additional, and unattributed `{turn_id, skill}` calls plus a relative transcript reference. It then restores the exact prior ordinary installation and deletes private token mappings.

These measures minimize evaluation clues. They cannot guarantee that an agent inspecting its environment will never infer that instrumentation exists.

## Private Storage and Recovery

Private state uses the platform application-state location:

- macOS: `~/Library/Application Support/Skill Issue`
- Windows: `%LOCALAPPDATA%\Skill Issue`
- Linux: `$XDG_STATE_HOME/skill-issue` or `~/.local/state/skill-issue`

Run directories and token mappings use restrictive permissions. If a run is interrupted, use the run identifier with `skill-issue evaluate cleanup --run <id>` to restore or remove its temporary installation. Cleanup requires the evaluation receipt to belong to that run. Evidence remains available after cleanup.

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

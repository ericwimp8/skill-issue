# Skill-Calling Instrumentation Contract

## Ownership

- `skills/` and `supporting-skills/` are the canonical ordinary-installation sources.
- The CLI embeds those complete source trees and each governed evaluation as one unit containing its scenario and private answer sheet. It owns installation, temporary transformation, private run state, replay, evidence, canonical rematerialization, and cleanup.
- A built-in evaluation identifier selects both embedded parts. Custom evaluations supply one directory of skill directories plus scenario and answer-sheet JSON files, all through the same runner path.
- Every evaluation installation uses the selected harness's native project skill root inside the required workspace. Evaluation has no user-scope installation mode.
- Every evaluation requires a caller-selected output root outside the evaluated workspace. Tooling-complete results are written under a unique `<output>/<harness>-<UTC-timestamp>-<run-prefix>/` directory, while private operational state, token mappings, internal events, and pre-instrumentation path presence exist under `<output>/.skill-issue/` only until successful cleanup.
- The evaluated harness receives neither the answer sheet nor the token-to-skill map.

## Custom Input Contract

Custom mode requires `--skills`, `--scenario`, and `--answer-sheet` together. The skills path contains only direct child skill directories; each has a `SKILL.md` entrypoint whose frontmatter `name` matches its directory, and each entrypoint's local `references/`, `scripts/`, or `assets/` paths resolve within that skill. At least one skill is required. The scenario must validate as schema version 1, and the answer sheet must use the same scenario ID, existing turn IDs, and names from the supplied skill set. The answer sheet remains outside the evaluated workspace. The CLI validates this runnable structure without judging whether the expected calls are semantically correct. Built-in mode rejects all three custom arguments.

## Disposable Evaluation Copies

For every canonical skill, the CLI creates a temporary copy that:

1. preserves the complete frontmatter byte-for-byte;
2. inserts exactly one instruction immediately after the closing frontmatter delimiter;
3. preserves the remainder of the canonical directory and body unchanged.

The instruction is:

```text
Run "<absolute-cli-path>" signal "<opaque-token>" "<absolute-output-state-path>", then continue normally.
```

Each token contains 32 cryptographically random bytes. Its meaning exists only in output-owned private run state. The absolute state path lets the separately launched signal process resolve that state without relying on a user-level installation. Canonical sources and ordinary installations never receive this instruction.

## Private State and Turn Attribution

The CLI stores the run, workspace, harness, effective model and reasoning, selected evaluation, scenario, project scope, active turn, harness session identifier, token map, internal event log, and the names of Skill Issue paths that existed before instrumentation under `<output>/.skill-issue/` outside the evaluated workspace.

Before sending a prompt, the runner sets the active turn. It sends the prompt verbatim, waits for terminal completion, and then closes the turn. `skill-issue signal <opaque-token> <absolute-output-state-path>` resolves the token privately and atomically appends an event when the harness executes it. For Codex, the runner instead records the same opaque token from Codex's structured command event while the turn is active; the command does not need permission to write private state. A signal received outside an active turn is retained as unattributed tooling evidence rather than assigned to a guessed turn. Successful signaling emits no output.

## Replay Boundary

The CLI drives one clean primary-agent session directly through the selected harness's programmatic interface. It never acknowledges, summarizes, corrects, or adapts later prompts from model responses. Sub-agent orchestration is outside this evaluation design.

For Codex, the CLI retains the user's ordinary Codex home only as the harness-owned authentication and resumable-session surface. Every initial and resumed turn ignores user configuration and execution rules, receives explicit medium reasoning and `workspace-write` settings, excludes discovered `AGENTS.md` content, disables apps and plugins, and receives a generated deny-list for skills discovered under the user's Codex skill root, `$HOME/.agents/skills`, and `/etc/codex/skills`. The generated project-local evaluation skills are outside those roots and remain available.

Codex uses interactive `on-request` approvals with `auto_review`. Commands already inside `workspace-write` need no review. Boundary requests are sent to the guardian reviewer, which may approve a sufficiently authorized low- or medium-risk action without expanding the sandbox. A denied opaque signal command remains attributable from the structured command event; other denials and reviewer failures fail closed and are tooling failures. The CLI never edits or copies user configuration, credentials, skills, plugins, rules, or instructions, and it leaves the resumable session in Codex's normal history. Managed requirements and system configuration cannot be bypassed and may prevent qualification. Other harnesses still require equivalent runtime-isolation qualification before their results can be treated as campaign evidence.

Launch, permission, session, marker, and protocol failures are tooling failures that require repair and rerun. Missing expected calls after a tooling-complete replay remain descriptive model-and-harness observations.

## Evidence and Cleanup

After the final turn, the CLI compares recorded token events with the selected built-in answer sheet or paired custom answer-sheet file. It always writes `result.json` and `website.json` under a unique `<output>/<harness>-<UTC-timestamp>-<run-prefix>/` directory. The detailed result identifies the selected evaluation, effective model, and effective reasoning, and retains expected, observed, missing, additional, and unattributed `{turn_id, skill}` calls. Results contain no pass or fail label.

Raw diagnostic persistence is explicit-only. `--events` writes the recorded invocation events to `events.jsonl`. `--transcript` writes the scenario and sanitized replay capture to `transcript.json` and adds its relative path to the detailed result. Both default to off and may be enabled independently. After replay and signal attribution complete, the CLI replaces the resolved CLI, workspace, output, private state, runtime, temporary, and home paths plus known local user and host identifiers throughout prompts, transcript text, standard error, and structured event strings. The persisted transcript uses bracketed placeholders while runtime execution and scoring retain the original absolute values. Arbitrary personal or confidential content intentionally present in prompts or responses is not inferred, so the CLI warns callers to review the transcript before sharing. Custom scenario and answer-sheet inputs retain their separate warning that they must contain no personal, confidential, or sensitive information. Warnings use standard error and do not alter structured result output.

Cleanup records only which canonical Skill Issue paths existed before instrumentation. It rematerializes current canonical copies for those paths, removes temporary evaluation paths that did not exist, deletes token mappings, and removes the completed run's private operational state. The selected result artifacts remain in the caller-selected output directory. No installed skill contents are backed up. `skill-issue evaluate cleanup --output <path> --run <id>` completes the same recovery and private-state removal after an interrupted run.

These controls minimize clues about the evaluation. They cannot guarantee that an agent which inspects its environment will never infer that instrumentation exists.

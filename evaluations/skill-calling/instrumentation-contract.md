# Skill-Calling Instrumentation Contract

## Ownership

- `skills/` and `supporting-skills/` are the canonical ordinary-installation sources.
- The CLI embeds those complete source trees and each governed evaluation as one unit containing its scenario and private answer sheet. It owns installation, temporary transformation, private run state, replay, evidence, restoration, and cleanup.
- A built-in evaluation identifier selects both embedded parts. Custom evaluations supply both JSON files from outside the evaluated workspace and use the same runner path.
- Every evaluation installation uses the selected harness's native project skill root inside the required workspace. Evaluation has no user-scope installation mode.
- Every evaluation requires a caller-selected output root outside the evaluated workspace. Tooling-complete evidence is written under `<output>/<run-id>/` while private operational state remains in CLI-owned application storage.
- The evaluated harness receives neither the answer sheet nor the token-to-skill map.

## Disposable Evaluation Copies

For every canonical skill, the CLI creates a temporary copy that:

1. preserves the complete frontmatter byte-for-byte;
2. inserts exactly one instruction immediately after the closing frontmatter delimiter;
3. preserves the remainder of the canonical directory and body unchanged.

The instruction is:

```text
Run "<absolute-cli-path>" signal "<opaque-token>", then continue normally.
```

Each token contains 32 cryptographically random bytes. Its meaning exists only in CLI-owned private run state. Canonical sources and ordinary installations never receive this instruction.

## Private State and Turn Attribution

The CLI stores the run, workspace, harness, model, scenario, project scope, active turn, harness session identifier, token map, event log, installed paths, and restoration data in the platform application-state directory outside the evaluated workspace.

Before sending a prompt, the runner sets the active turn. It sends the prompt verbatim, waits for terminal completion, and then closes the turn. `skill-issue signal <opaque-token>` resolves the token privately and atomically appends an event. A signal received outside an active turn is retained as unattributed tooling evidence rather than assigned to a guessed turn. Successful signaling emits no output.

## Replay Boundary

The CLI drives one clean primary-agent session directly through the selected harness's programmatic interface. It never acknowledges, summarizes, corrects, or adapts later prompts from model responses. Sub-agent orchestration is outside this evaluation design.

Launch, permission, session, marker, and protocol failures are tooling failures that require repair and rerun. Missing expected calls after a tooling-complete replay remain descriptive model-and-harness observations.

## Evidence and Cleanup

After the final turn, the CLI compares recorded token events with the selected built-in answer sheet or the paired custom answer-sheet file. It writes `events.jsonl`, `transcript.json`, `result.json`, and `website.json` under the required `<output>/<run-id>/` directory. The detailed result retains expected, observed, missing, additional, and unattributed `{turn_id, skill}` calls with a relative transcript reference. Results contain no pass or fail label.

Cleanup removes the temporary copies, restores the exact prior receipt-owned ordinary installation, deletes token mappings, and retains the evidence package in the caller-selected output directory. Cleanup requires the evaluation receipt to belong to the same run. `skill-issue evaluate cleanup --run <id>` repeats that recovery safely after an interrupted run.

These controls minimize clues about the evaluation. They cannot guarantee that an agent which inspects its environment will never infer that instrumentation exists.

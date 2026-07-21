# Skill-Calling Evaluation Campaign — Orchestrator Briefing

## The job

Run the campaign's 30 blind skill-calling evaluations — ten harness-and-model configurations, each running the three built-in scenarios once, in full — and keep `evaluation-progress.md` (same directory) truthful as you go. Complete as many of the 30 as you can: troubleshoot what fails where you can, record with evidence what you cannot, and never let one failure stop the rest. When every run is settled, report what completed, what didn't, and why, with enough diagnosis that the operator can fix and rerun the leftovers.

You own the whole task: scheduling, launching, monitoring, bookkeeping, and reporting. The tracker is yours alone to write.

## Running one evaluation

This is genuinely all it takes — one command per run:

```zsh
REPO=$(git rev-parse --show-toplevel)
"$REPO/cli/scripts/local-cli.sh" evaluate run \
  --workspace <fresh-empty-dir> --output <dir-outside-the-workspace> \
  --harness codex --events --transcript --yes \
  --evaluation gardening-web-application
```

`local-cli.sh` is the qualified known-good CLI (its `version` prints `known-good-*`; the development channel is not for campaign work). The run prints a pre-run summary, then per-turn progress, and takes 30–90 minutes for a full scenario. Success ends with `"status": "complete"` and leaves `result.json`, `website.json`, `events.jsonl`, and `transcript.json` in a run-named directory under the output root. Failure prints `evaluation encountered a tooling error` and names a `failure.json` holding the exact harness command, active turn, error chain, and raw native output — that file is your diagnostic starting point every time (it is unsanitized; never commit or share it). No `--turns` flag: campaign runs are always the complete scenario.

`skill-issue doctor --harness <id> [--executable <path>]` diagnoses a harness in seconds with zero model cost — useful before starting and whenever a lane misbehaves. If an error message ever names an `evaluate cleanup --run <id> --output <root>` command, running exactly that once restores the workspace.

## The thirty runs

Scenario IDs, in the order the campaign runs them within each configuration: `gardening-web-application`, `community-archive-desktop-application`, `neighborhood-emergency-preparedness-program` (run IDs `-01`, `-02`, `-03`).

| Configuration       | IDs            | `--harness`   | `--executable`                                                             | `--model`         | environment                                                     |
| ------------------- | -------------- | ------------- | -------------------------------------------------------------------------- | ----------------- | --------------------------------------------------------------- |
| OpenAI Codex — Sol  | COD-SOL-01..03 | `codex`       | `$(command -v codex)`                                                      | default           | —                                                               |
| Pi — Codex          | PI-COD-01..03  | `pi`          | `$(command -v pi)`                                                         | default           | —                                                               |
| OpenCode — Codex    | OPE-COD-01..03 | `opencode`    | `$REPO/.skill-issue/opencode/bin/opencode`                                 | default           | `XDG_DATA_HOME="$REPO/.skill-issue/opencode/home/.local/share"` |
| Kilo — Codex        | KIL-COD-01..03 | `kilo-code`   | `$REPO/.skill-issue/kilo/node_modules/@kilocode/cli-darwin-arm64/bin/kilo` | default           | `XDG_DATA_HOME="$REPO/.skill-issue/kilo/home/.local/share"`     |
| Cursor — Composer   | CUR-COM-01..03 | `cursor`      | `$REPO/.skill-issue/cursor/home/.local/bin/agent`                          | resolved Composer | —                                                               |
| Cursor — Grok       | CUR-GRO-01..03 | `cursor`      | same Cursor agent                                                          | resolved Grok     | —                                                               |
| Cursor — Codex      | CUR-COD-01..03 | `cursor`      | same Cursor agent                                                          | resolved Sol      | —                                                               |
| Cursor — Fable      | CUR-FAB-01..03 | `cursor`      | same Cursor agent                                                          | resolved Fable    | —                                                               |
| Claude Code — Codex | CLA-COD-01..03 | `claude-code` | `$REPO/.skill-issue/claudex/claudex`                                       | `gpt-5.6-sol`     | —                                                               |
| Claude Code — Fable | CLA-FAB-01..03 | `claude-code` | `$(command -v claude)`                                                     | `claude-fable-5`  | —                                                               |

"Resolved" Cursor models: the Cursor catalog drifts daily (a `-medium` variant vanished and reappeared within one day here), so identifiers are resolved fresh from `agent --list-models` at start, never from memory. Prefer the exact `-medium` variant; the recorded fallbacks when none exists are `composer-2.5`, `cursor-grok-4.5-high`, `gpt-5.6-sol-high`, and `claude-fable-5-thinking-high`. A rejected identifier is cheap to correct — the rejection lists what the harness accepts. Record whatever you actually used.

## Isolation — why workspaces look the way they do

These are blind evaluations: the evaluated model must find no clue that it is being measured. That is why every run gets a fresh, empty, neutrally named workspace, and why results live outside it. The campaign convention: containers `chat-<n>` under `<repository-parent>/chats/`, each holding an empty `workspace/` and an `output/`, numbered serially past the highest existing number. Nothing in a container path may hint at evaluation, skills, harnesses, or models. A used container — passing or failing — is retained evidence: never reused, cleaned, or reorganized. After a run, the workspace holds whatever the model built (that's evidence too, including any skill the model itself authored); the only things that would signal trouble are leftover canonical Skill Issue skills or private `skill-issue/` state in the system temp directory, which a healthy run always removes.

## Constraints that are real, not stylistic

- At most ten evaluations in flight; within a configuration the three scenarios run one at a time, in order.
- At most one `claude-code`-harness run at any moment, campaign-wide — concurrent Claude sessions fight over session state.
- **Claude Code — Fable runs last, after everything else is settled**, and only after the claudex proxy (owned by `$REPO/.skill-issue/claudex/manage`) is stopped and verifiably gone — the Codex-proxy route and the normal `claude` route must never coexist.
- Never pass `--reasoning`: medium is already every harness's default, and Cursor rejects the flag outright.
- Never use `$REPO/.skill-issue/kilo/bin/kilo` — that wrapper pins its own configuration home and silently discards the evaluation's generated configuration.
- Concurrent Cursor runs share one account; if rate limits appear, fewer simultaneous Cursor runs is the fix.

## What results mean

Missing, additional, or even zero observed skill calls in a completed run are **evaluation data, never failure** — mark the run `Complete`. Skill visibility is machine-verified during the run for Claude Code, OpenCode, Kilo, and Pi, so a completed run's misses are trustworthy; Codex and Cursor lack that verification surface, so a zero-call run there deserves a note for later review, nothing more. Only "evaluation encountered a tooling error" is a failure.

## What failures mean

The `failure.json` error text tells you which kind of problem you have:

- **`signal: killed`, `signal: terminated`, `context canceled`** — especially several runs at the same instant — means the _orchestrator's environment_ died, not the runs. This has happened twice here: once because an orchestrator ran inside a Codex sandbox (each command needed an escalation a reviewer could and did deny), and once because an orchestrator's own process ended and took its child runs down with it. The lessons: this job needs a seat with direct filesystem access to the `chats/` directory, runs deserve to be launched so they survive their launcher (`nohup`-style detachment), and a successor orchestrator should trust live processes and logs over what the tracker says — adopt anything still running rather than duplicating it. Environment deaths say nothing about a lane; the affected runs are simply eligible again.
- **Rate limits, network timeouts, one-off crashes or protocol breaks** — transient. Worth retrying in a fresh container with whatever correction the diagnosis suggests; blind identical reruns teach nothing, and past three attempts a run has had its chance.
- **`Error loading config`** (harness rejects generated configuration), **`no marker was recorded` / `did not discover installed evaluation skill` / `not loaded` / `cannot verify`** (instrumentation integrity), **`qualified version`** (version pin) — systemic. Retrying re-proves the same defect at full price; record the evidence and move on, and once a cause has claimed two runs in a lane it has claimed the lane.
- **Authentication or subscription failures** — the operator's to fix; `Blocked`, not worth burning attempts on.

Boundaries that hold regardless of how tempting a fix looks: no weakening sandboxes or permissions, no repairing or re-logging authentication, no editing user configuration, harness installs, or product source, no destroying evidence.

## Keeping the operator in the loop

Post short one-line updates in the conversation as events happen — a run starting (with its container), finishing (with expected/observed counts), retrying (with the reason), failing (with the cause), a configuration finishing — plus an occasional tally like `12/30 complete, 8 running, 1 failed`. The operator watches the campaign through these lines.

In the tracker, keep statuses, attempt counts, container references (always as `<chats>/chat-<n>`, never machine-absolute paths), failure-log rows, and the summary tables consistent with reality, and leave it prettier-clean after each edit. If the board you inherit shows `Failed`/`Blocked` entries but no `Complete` ones, that is the residue of the earlier environment deaths described above — those runs are all still eligible.

The final report belongs to the operator: per-configuration outcomes, the exact identifiers and versions actually used, each failure with its diagnosis and container, and whether the campaign reached 30/30. Publishing, committing, and cleanup decisions stay with the operator.

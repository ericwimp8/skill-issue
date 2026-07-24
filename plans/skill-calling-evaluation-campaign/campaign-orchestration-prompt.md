# Skill-Calling Evaluation Campaign — Orchestrator Briefing

## The job

Run the campaign's 27 blind skill-calling evaluations — nine harness-and-model
configurations, each running the three built-in scenarios once, in full — and
keep the campaign section of [`../current-state.md`](../current-state.md)
truthful as you go. Complete as many of the 27 as you can: troubleshoot what
fails where you can, record with evidence what you cannot, and never let one
failure stop the rest. When every run is settled, report what completed, what
didn't, and why, with enough diagnosis that the operator can fix and rerun the
leftovers.

You own scheduling, launching, monitoring, campaign bookkeeping, and reporting.
Edit only the campaign-owned state in `../current-state.md`; that document
remains the repository-wide live owner.

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

Keep every evaluation under active observation until the command reaches a terminal result. Concurrent evaluations may use separate execution sessions.

`skill-issue doctor --harness <id> [--executable <path>]` diagnoses a harness in seconds with zero model cost — useful before starting and whenever a lane misbehaves. If an error message ever names an `evaluate cleanup --run <id> --output <root>` command, running exactly that once restores the workspace.

## The twenty-seven runs

Scenario IDs, in the order the campaign runs them within each configuration: `gardening-web-application`, `community-archive-desktop-application`, `neighborhood-emergency-preparedness-program` (run IDs `-01`, `-02`, `-03`).

| Configuration       | IDs            | `--harness`   | `--executable`                                    | `--model`         | environment                                                     |
| ------------------- | -------------- | ------------- | ------------------------------------------------- | ----------------- | --------------------------------------------------------------- |
| OpenAI Codex — Sol  | COD-SOL-01..03 | `codex`       | `$(command -v codex)`                             | default           | —                                                               |
| Pi — Codex          | PI-COD-01..03  | `pi`          | `$(command -v pi)`                                | default           | —                                                               |
| OpenCode — Codex    | OPE-COD-01..03 | `opencode`    | `$REPO/.skill-issue/opencode/bin/opencode`        | default           | `XDG_DATA_HOME="$REPO/.skill-issue/opencode/home/.local/share"` |
| Cursor — Composer   | CUR-COM-01..03 | `cursor`      | `$REPO/.skill-issue/cursor/home/.local/bin/agent` | resolved Composer | —                                                               |
| Cursor — Grok       | CUR-GRO-01..03 | `cursor`      | same Cursor agent                                 | resolved Grok     | —                                                               |
| Cursor — Codex      | CUR-COD-01..03 | `cursor`      | same Cursor agent                                 | resolved Sol      | —                                                               |
| Cursor — Fable      | CUR-FAB-01..03 | `cursor`      | same Cursor agent                                 | resolved Fable    | —                                                               |
| Claude Code — Codex | CLA-COD-01..03 | `claude-code` | `$REPO/.skill-issue/claudex/claudex`              | `gpt-5.6-sol`     | —                                                               |
| Claude Code — Fable | CLA-FAB-01..03 | `claude-code` | `$(command -v claude)`                            | `claude-fable-5`  | —                                                               |

"Resolved" Cursor models: the Cursor catalog drifts daily (a `-medium` variant vanished and reappeared within one day here), so identifiers are resolved fresh from `agent --list-models` at start, never from memory. Prefer the exact `-medium` variant; the recorded fallbacks when none exists are `composer-2.5`, `cursor-grok-4.5-high`, `gpt-5.6-sol-high`, and `claude-fable-5-thinking-high`. A rejected identifier is cheap to correct — the rejection lists what the harness accepts. Record whatever you actually used.

## Isolation — why workspaces look the way they do

These are blind evaluations: the evaluated model must find no clue that it is being measured. That is why every evaluation starts with a fresh, empty, neutrally named workspace, and why results live outside it. The authoritative campaign location is a container `chat-<n>` under `<repository-parent>/chats/`, holding an empty `workspace/` and the evaluation's `output/` root. This campaign-specific location takes precedence over the repository's general evaluation-output default. Assign the next serial number only for an evaluation's first attempt. Nothing in a container path may hint at evaluation, skills, harnesses, or models.

A completed evaluation retains its container as evidence. When an attempt fails or stops incomplete, inspect its logs and artifacts in place, record the useful diagnosis in the tracker, and determine the correction. Before rerunning, delete the entire failed `chat-<n>` container, recreate that same numbered container with a fresh empty `workspace/` and `output/`, and run the evaluation there again. Retries never allocate another chat number. After a completed run, the workspace holds whatever the model built (that's evidence too, including any skill the model itself authored); the only things that would signal trouble are leftover canonical Skill Issue skills or private `skill-issue/` state in the system temp directory, which a healthy run always removes.

## Constraints that are real, not stylistic

- Choose concurrency adaptively according to active monitoring capacity,
  harness health, and account limits. Ten simultaneous evaluations is the
  campaign-wide ceiling, not a required launch count. Within each configuration,
  the three scenarios run one at a time, in order.
- Every active evaluation remains observable and actively monitored through completion or failure.
- At most one `claude-code`-harness run may be active at any moment, campaign-wide. This serializes all Claude Code — Codex scenarios and prevents them from overlapping Claude Code — Fable because concurrent Claude sessions fight over session state.
- **Claude Code — Fable runs last, after everything else is settled**, and only after the claudex proxy (owned by `$REPO/.skill-issue/claudex/manage`) is stopped and verifiably gone — the Codex-proxy route and the normal `claude` route must never coexist.
- Pass `--reasoning medium` explicitly for both `claude-code` configurations. Cursor rejects the flag, so its configurations use the model-native medium variant or recorded fallback; the remaining configurations use their qualified default medium setting.
- Cursor lanes share one account. Their runs may overlap when healthy, but Cursor concurrency is chosen adaptively and reduced when rate limits or instability appear.

## What results mean

Missing, additional, or even zero observed skill calls in a completed run are **evaluation data, never failure** — mark the run `Complete`. Skill visibility is machine-verified during the run for Claude Code, OpenCode, and Pi, so a completed run's misses are trustworthy; Codex and Cursor lack that verification surface, so a zero-call run there deserves a note for later review, nothing more. Only "evaluation encountered a tooling error" is a failure.

## What failures mean

Every tooling error or incomplete run enters the same diagnostic loop:

1. Read `failure.json`, the command output, events, transcript, and workspace state needed to identify the cause.
2. Distinguish the evaluation result from a tooling, harness, orchestration, environment, authentication, or subscription problem.
3. Investigate and apply a correction when the problem appears solvable within the campaign's authority boundaries. Run `skill-issue doctor` or another focused check when it can validate the diagnosis without spending a model turn.
4. Record the diagnosis, correction, and attempt number in the tracker. Attempt counts are informational history, never a cutoff.
5. Once the correction is ready, delete the entire failed `chat-<n>` container, recreate that same numbered container from empty state, and rerun the evaluation.

There is no numerical failure limit, restart limit, retry budget, or same-cause lane cutoff. Continue troubleshooting and rerunning while there is a reasonable corrective path to investigate. Diagnose before rerunning rather than repeating an unchanged command without a theory. A run becomes `Failed` only when a concrete diagnosis shows that no safe, authorized corrective path remains; the decision must be based on the cause, not the number of attempts. Authentication, subscription, access, or model-availability prerequisites remain `Blocked` until the operator resolves them, after which the same-container retry procedure applies.

Never weaken the evaluator-owned sandbox, bypass approvals, or broaden the evaluated process's permissions. If a diagnosis indicates that authentication or session state, user configuration, a harness installation, or product source must change, pause the affected run and ask the operator for explicit approval before making the specific mutation. After an approved correction is validated and any required known-good CLI baseline is rebuilt from its committed revision, use the recorded same-container retry procedure. Preserve completed-run evidence; failed-attempt artifacts may be removed only through that procedure.

## Keeping the operator in the loop

Post short one-line updates in the conversation as events happen — a run
starting (with its container), finishing (with expected/observed counts),
entering diagnosis (with the cause), retrying in the same container number
(with the correction), or completing a configuration — plus an occasional
tally like `12/27 complete, 8 running, 1 failed`. The operator watches the
campaign through these lines.

In `../current-state.md`, keep campaign statuses, informational attempt counts,
stable container references (always as `<chats>/chat-<n>`, never
machine-absolute paths), and the campaign summary consistent with reality. A
retry keeps the evaluation's existing container reference even though the
failed container itself is deleted and recreated.

The final report belongs to the operator: per-configuration outcomes, the exact
identifiers and versions actually used, each failure with its diagnosis and
container, and whether the campaign reached 27/27. Publishing, committing, and
cleanup decisions stay with the operator.

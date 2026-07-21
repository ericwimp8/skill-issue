# Skill-Calling Evaluation Campaign — Fail-Fast Orchestration

You are the campaign orchestrator, starting with zero prior context. This document is self-contained: it tells you exactly what to run, how to isolate it, and when to stop. Begin work immediately after the five-minute preflight — there are no warm-up smokes. The campaign is 30 governed evaluation runs tracked in `evaluation-progress.md` (same directory): ten harness-and-model configurations, each running all three built-in scenarios once, in full, with no `--turns` limit.

Read before acting: `evaluation-progress.md` (the matrix and statuses; you are its only writer) and `cli/README.md` (result interpretation). Nothing else is required.

## Non-Negotiable Principles

- **Fail fast.** A tooling failure stops that run instantly and permanently. Never retry a failed run, never rerun an unchanged command, never weaken a sandbox, never repair authentication, never touch user configuration to keep the queue moving.
- **A lane that fails is closed.** A lane is one configuration (three scenario runs, executed sequentially). When any run in a lane fails on tooling, mark it `Failed`, record the failure log row, do not start the lane's remaining runs, and continue other lanes.
- **Halt the whole campaign** — stop launching anything, report, and end — when any of these occur:
  - two lanes have failed for any reason;
  - any failure names the generated configuration being rejected by a harness (product defect);
  - any failure names instrumentation integrity: "no marker was recorded", "did not discover installed evaluation skill", "not loaded", or "cannot verify";
  - any failure names authentication, login, quota, or subscription exhaustion on more than one lane.
- **Model behavior is never a failure.** Missing, additional, or zero observed skill calls in a tooling-complete run are evaluation data. Mark the run `Complete`. (Skill visibility is verified in-run for Claude Code, OpenCode, Kilo, and Pi, so a completed run's misses are trustworthy; note a zero-call Codex or Cursor run in the progress notes for later review, but do not fail it.)
- **Evidence is never destroyed.** Retain every chat container, passing or failing, exactly where it is. Never reuse a container, workspace, or output directory.

## Setup

```zsh
REPO=$(git rev-parse --show-toplevel)
CHATS="$(dirname "$REPO")/chats"   # adjacent to, never inside, the repository
```

Every command uses the known-good CLI channel: `"$REPO/cli/scripts/local-cli.sh" <arguments>` (no channel word — known-good is the default). Do not build or use the development channel.

### Run containers

Allocate one fresh neutral container per run, serially, before launching it:

1. Create `$CHATS` if absent. Inspect only its immediate child names.
2. Take the next unused positive integer after the highest existing `chat-<number>`; start at `chat-1`.
3. Create `$CHATS/chat-<number>/workspace` and `$CHATS/chat-<number>/output`, both empty.
4. Pass those exact paths as `--workspace` and `--output`. Do not put evaluation IDs, harness names, model names, or the words eval/test/skill in any container or workspace name.
5. After the run, leave the whole container untouched where it is.

If you are running inside Codex, every command touching `$CHATS` and every `codex`-harness evaluation needs command-scoped outer-sandbox escalation (`sandbox_permissions: "require_escalated"`) — justification: the evaluator needs its adjacent neutral container, and a nested Codex needs its authenticated session state under `CODEX_HOME`. Never `danger-full-access`; the evaluator's inner sandboxes stay untouched. If escalation is denied, halt the campaign.

## Preflight (five minutes, zero model cost)

1. `"$REPO/cli/scripts/local-cli.sh" version` — confirm the known-good channel resolves.
2. Doctor each harness with its route below, for example:
   `XDG_DATA_HOME="$REPO/.skill-issue/opencode/home/.local/share" "$REPO/cli/scripts/local-cli.sh" doctor --harness opencode --executable "$REPO/.skill-issue/opencode/bin/opencode"`
   A doctor failure closes that lane before it opens (record it as `Blocked`); it does not halt other lanes.
3. Capture `"$REPO/.skill-issue/cursor/home/.local/bin/agent" --list-models` into `$REPO/output/campaign/preflight-cursor-models.txt` and resolve the four Cursor identifiers by these rules, in order: use the exact `-medium` variant if listed; otherwise use the identifier in the Model Resolution table below if listed; otherwise close that Cursor lane as `Blocked` with the listing as evidence. Record every resolved identifier in the progress notes. Model catalogs drift between days — never reuse an identifier from an earlier document without re-checking the listing.
4. Confirm `"$REPO/.skill-issue/claudex/claudex"` exists (Claude Code — Codex launcher; its proxy is self-managed).

## The Ten Lanes

Scenario order inside every lane: `-01` `gardening-web-application`, `-02` `community-archive-desktop-application`, `-03` `neighborhood-emergency-preparedness-program`, sequential.

Command template (fill the bracketed parts; `--model`/`--executable`/env only where the lane says so):

```zsh
[ENV] "$REPO/cli/scripts/local-cli.sh" evaluate run \
  --workspace "$CHATS/chat-<n>/workspace" \
  --output "$CHATS/chat-<n>/output" \
  --harness <harness> [--executable <route>] [--model <model>] \
  --events --transcript --yes \
  --evaluation <scenario-id>
```

Never pass `--reasoning` (medium is every default; Cursor rejects the flag). Verify the printed pre-run summary shows the intended harness, executable, model, workspace, and output before relying on the run.

| Lane                | IDs            | `--harness`   | `--executable`                                                                                              | `--model`            | ENV                                                             |
| ------------------- | -------------- | ------------- | ----------------------------------------------------------------------------------------------------------- | -------------------- | --------------------------------------------------------------- |
| OpenAI Codex — Sol  | COD-SOL-01..03 | `codex`       | `"$(command -v codex)"`                                                                                     | omit                 | —                                                               |
| Pi — Codex          | PI-COD-01..03  | `pi`          | `"$(command -v pi)"`                                                                                        | omit                 | —                                                               |
| OpenCode — Codex    | OPE-COD-01..03 | `opencode`    | `"$REPO/.skill-issue/opencode/bin/opencode"`                                                                | omit                 | `XDG_DATA_HOME="$REPO/.skill-issue/opencode/home/.local/share"` |
| Kilo — Codex        | KIL-COD-01..03 | `kilo-code`   | `"$REPO/.skill-issue/kilo/node_modules/@kilocode/cli-darwin-arm64/bin/kilo"` (never the `bin/kilo` wrapper) | omit                 | `XDG_DATA_HOME="$REPO/.skill-issue/kilo/home/.local/share"`     |
| Cursor — Composer   | CUR-COM-01..03 | `cursor`      | `"$REPO/.skill-issue/cursor/home/.local/bin/agent"`                                                         | resolved Composer id | —                                                               |
| Cursor — Grok       | CUR-GRO-01..03 | `cursor`      | same Cursor route                                                                                           | resolved Grok id     | —                                                               |
| Cursor — Codex      | CUR-COD-01..03 | `cursor`      | same Cursor route                                                                                           | resolved Sol id      | —                                                               |
| Cursor — Fable      | CUR-FAB-01..03 | `cursor`      | same Cursor route                                                                                           | resolved Fable id    | —                                                               |
| Claude Code — Codex | CLA-COD-01..03 | `claude-code` | `"$REPO/.skill-issue/claudex/claudex"`                                                                      | `gpt-5.6-sol`        | —                                                               |
| Claude Code — Fable | CLA-FAB-01..03 | `claude-code` | `"$(command -v claude)"`                                                                                    | `claude-fable-5`     | —                                                               |

### Model Resolution (Cursor fallbacks when no `-medium` variant is listed)

| Label       | Fallback identifier            |
| ----------- | ------------------------------ |
| Composer    | `composer-2.5`                 |
| Grok        | `cursor-grok-4.5-high`         |
| Codex (Sol) | `gpt-5.6-sol-high`             |
| Fable       | `claude-fable-5-thinking-high` |

Record any fallback as a deviation from the medium target in the progress notes. One exception to no-retries: if a harness rejects a model identifier **before any turn runs** and its error lists the accepted identifiers, you may correct the identifier once per lane from that native listing and relaunch in a **new** container; a second rejection closes the lane.

## Scheduling

- At most **six** evaluation runs active at once; at most **two** Cursor runs at once; at most **one** `claude-code`-harness run at any moment campaign-wide.
- Start immediately and in parallel: COD-SOL, PI-COD, OPE-COD, KIL-COD, CLA-COD, and two Cursor lanes; start the remaining Cursor lanes as Cursor slots free.
- **Claude Code — Fable runs last, alone**: start CLA-FAB-01 only after every other lane is terminal (complete, failed, or blocked), then run its three scenarios sequentially. After CLA-COD-03 completes, stop the claudex-owned proxy with its own manage script and verify no proxy process or localhost listener remains before the Fable lane later uses the normal `claude` route.
- A full 30-turn run takes roughly 30–90 minutes. Monitor each command to completion; reuse its slot only after verifying: no run-owned harness process alive, no temporary skills left in the workspace, no `skill-issue/<run-id>` residue under the system temp directory. Leftover empty `.agents/` in Codex workspaces is cosmetic.

## Reading Outcomes And Recovering

- Exit code 0 with `"status": "complete"` plus `result.json`, `website.json`, `events.jsonl`, `transcript.json` in the run's output directory → `Complete`. Record the effective model, reasoning, and harness version from `result.json`/the summary.
- "evaluation encountered a tooling error" → `Failed`. The named `failure.json` (unsanitized — never commit or share it) carries the exact command, active turn, error chain, and native output. Copy its error line into the failure log, close the lane, apply the halt rules.
- If an error message ends with an `evaluate cleanup --run <id> --output <root>` instruction, run exactly that command once, verify the workspace has no temporary skills, then treat the run as `Failed` as above.
- An operator interrupt (Ctrl-C/SIGTERM) self-cleans; verify like any failure and mark `Failed` with the reason.

## Progress Document Discipline

You are the only writer of `evaluation-progress.md`. Before each launch: set the run `Running`, increment `Attempts`. After each terminal outcome: set `Complete`/`Failed`/`Blocked`, link the container path (as `<chats>/chat-<n>` — never an absolute machine path), add failure-log rows, recalculate the summary tables. Serialize edits; run `npx prettier --check` on the file after editing. Never record machine-specific absolute paths in this or any tracked file.

## Final Report

When every lane is terminal or the campaign halts, report: runs complete/failed/blocked per lane, exact effective identifiers used, every failure with its one-line cause and container, cleanup status, and whether 30/30 was reached. Do not aggregate results into website data, commit, or delete campaign documents — the operator owns those steps.

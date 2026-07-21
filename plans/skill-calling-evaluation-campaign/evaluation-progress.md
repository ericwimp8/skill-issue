# Skill-Calling Evaluation Campaign Progress

## Purpose

Track the subscription-bounded skill-calling evaluation campaign one scenario run at a time. The campaign contains nine harness-and-model configurations, with all three governed scenarios run once for each configuration, for a total of 27 evaluation runs.

This document records campaign execution progress only. The governed scenarios, instrumentation, result derivation, and valid-run requirements remain owned by the evaluation assets and their existing contracts.

## Campaign Configuration

| Harness      | Model label          | Reasoning target                       |   Runs |  Complete |
| ------------ | -------------------- | -------------------------------------- | -----: | --------: |
| Claude Code  | Codex                | Medium                                 |      3 |       3/3 |
| OpenAI Codex | GPT-5.6 Sol          | Medium                                 |      3 |       3/3 |
| Claude Code  | Fable                | Medium                                 |      3 |       3/3 |
| Cursor       | Fable                | Medium or model-native equivalent      |      3 |       0/3 |
| Cursor       | Codex                | Medium or model-native equivalent      |      3 |       0/3 |
| Cursor       | Grok                 | Medium or model-native equivalent      |      3 |       3/3 |
| Cursor       | Composer             | Medium or model-native equivalent      |      3 |       3/3 |
| Pi           | Codex                | Medium                                 |      3 |       3/3 |
| OpenCode     | Codex                | Medium or closest supported equivalent |      3 |       3/3 |
| **Total**    | **9 configurations** | **One consistent target**              | **27** | **21/27** |

Use one reasoning target across the campaign, or the closest documented harness-specific equivalent where a harness does not expose the same control. Record the exact effective model identifier, model version or alias, harness version, and reasoning setting with each result.

## Progress Summary

| Measure                       | Progress |
| ----------------------------- | -------: |
| Configuration suites complete |      7/9 |
| Evaluation runs complete      |    21/27 |
| Evaluation runs running       |     0/27 |
| Evaluation runs pending       |     0/27 |
| Evaluation runs failed        |        0 |
| Evaluation runs blocked       |     6/27 |
| Overall completion            |    77.8% |

## Tracking Rules

- Check an evaluation only after the complete governed scenario and its instrumentation finish successfully and the required result artifacts are retained.
- A tooling-complete run with zero observed skill calls is a completed evaluation and may be checked.
- Keep a tooling failure unchecked while it is being diagnosed and rerun. Use `Failed` only when a concrete diagnosis shows that no safe, authorized corrective path remains, never because an attempt count was reached.
- Use `Blocked` when an evaluation cannot start because access, authentication, model availability, or another prerequisite is unavailable.
- When resolving a blocker or tooling defect requires changing authentication or session state, user configuration, a harness installation, or product source, request the operator's explicit approval before making that specific mutation.
- Record every attempt number for history; attempt counts are informational and never limit troubleshooting, restarts, or retries.
- Check a configuration heading only after all three scenario evaluations beneath it are complete.
- Update the configuration table and progress summary whenever an evaluation status changes.

Status values: `Pending`, `Running`, `Complete`, `Failed`, or `Blocked`.

## Execution Dependencies

- Eligible lanes may run concurrently, but the orchestrator chooses concurrency adaptively according to active monitoring capacity, harness health, and account limits. Ten simultaneous evaluations is the campaign-wide ceiling, not a required launch count.
- Each configuration lane runs its three scenarios sequentially. Tooling failures remain eligible for diagnosis and rerun while a reasonable corrective path exists; neither attempts nor repeated causes impose a numerical cutoff on a run or lane.
- Cursor lanes may overlap when healthy, but they are not required to start together; rate limits or instability reduce Cursor concurrency.
- At most one `claude-code`-harness run may be active campaign-wide. All Claude Code — Codex scenarios therefore run serially and never overlap Claude Code — Fable.
- Claude Code — Fable runs last, alone, only after every other lane is terminal and the Claude Code — Codex proxy is stopped and verified gone.
- The campaign orchestrator starts every evaluation in a directly attached foreground session, monitors it to a terminal result, and is the only writer of this progress document.
- Every evaluation uses its own external `chat-<n>` container. A failed or incomplete attempt is diagnosed in place; before rerun, its entire container is deleted and recreated under the same number with an empty workspace and output directory.
- Before launch, the orchestrator must prove direct read-write cleanup access under `<chats>`. Concurrent runs use separate foreground sessions.

The adjacent `campaign-orchestration-prompt.md` owns scheduling, command launch contracts, model-identifier resolution, troubleshooting, continuation and halt rules, and the progress-update procedure.

## 1. Claude Code — Codex

- [x] **Configuration complete:** Claude Code — Codex — 3/3
- [x] **CLA-COD-01:** Gardening Web Application — Status: `Complete` — Attempts: 2 — Result: `<chats>/chat-37` — Notes: Tooling-clean, all 30 turns; `gpt-5.6-sol` through the claudex proxy, medium reasoning; 46 expected, 12 expected hits, 34 missing, 5 additional. Attempt 1's evidence container (`<chats>/chat-5`) was removed during operator cleanup; attempt 2 ran in `<chats>/chat-37` in parallel with CLA-COD-03 after the operator relaxed the one-claudex rule.
- [x] **CLA-COD-02:** Community Archive Desktop Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-31` — Notes: Tooling-clean, all 30 turns; `gpt-5.6-sol` through the claudex proxy, medium reasoning; 46 expected, 4 expected hits, 42 missing, 2 additional; per-turn skill visibility verified in-run.
- [x] **CLA-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Complete` — Attempts: 2 — Result: `<chats>/chat-32` — Notes: Tooling-clean, all 30 turns; `gpt-5.6-sol` through the claudex proxy, medium reasoning; 45 expected, 3 expected hits, 42 missing, 0 additional. Attempt 1 was killed when the orchestrator session closed; attempt 2 ran clean.

## 2. OpenAI Codex — GPT-5.6 Sol

- [x] **Configuration complete:** OpenAI Codex — GPT-5.6 Sol — 3/3
- [x] **COD-SOL-01:** Gardening Web Application — Status: `Complete` — Attempts: 2 — Result: `<chats>/chat-18` — Notes: Tooling-clean, all 30 turns; `gpt-5.6-sol`, medium reasoning; 46 expected, 16 observed, 34 missing, 4 additional. The attempt-1 evidence container (`<chats>/chat-1`) was removed during operator cleanup, so attempt 2 reran the scenario.
- [x] **COD-SOL-02:** Community Archive Desktop Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-19` — Notes: Tooling-clean, all 30 turns; `gpt-5.6-sol`, medium reasoning; 46 expected, 44 expected hits, 2 missing, 36 additional; zero calls on the three no-expectation reminder turns.
- [x] **COD-SOL-03:** Neighborhood Emergency Preparedness Program — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-20` — Notes: Tooling-clean, all 30 turns; `gpt-5.6-sol`, medium reasoning; 45 expected, 45 expected hits (0 missing), 17 additional; zero calls on the three no-expectation reminder turns, corroborating attribution.

## 3. Claude Code — Fable

- [x] **Configuration complete:** Claude Code — Fable — 3/3
- [x] **CLA-FAB-01:** Gardening Web Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-33` — Notes: Tooling-clean, all 30 turns; `claude-fable-5`, medium reasoning; 46 expected, 7 observed, 40 missing, 1 additional; per-turn skill visibility verified in-run.
- [x] **CLA-FAB-02:** Community Archive Desktop Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-34` — Notes: Tooling-clean, all 30 turns; `claude-fable-5`, medium reasoning; 46 expected, 6 observed, 41 missing, 1 additional; per-turn skill visibility verified in-run.
- [x] **CLA-FAB-03:** Neighborhood Emergency Preparedness Program — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-35` — Notes: Tooling-clean, all 30 turns; `claude-fable-5`, medium reasoning; 45 expected, 7 observed, 39 missing, 1 additional; per-turn skill visibility verified in-run.

## 4. Cursor

### Fable

- [ ] **Configuration complete:** Cursor — Fable — 0/3
- [ ] **CUR-FAB-01:** Gardening Web Application — Status: `Blocked` — Attempts: 1 — Result: `<chats>/chat-15` — Notes: `claude-fable-5-thinking-high` (no medium variant listed; recorded deviation). Failed at turn 21: Cursor Pro+ API usage limit exhausted (resets 2026-08-21); requires operator billing action.
- [ ] **CUR-FAB-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 1 — Result: `<chats>/chat-16` — Notes: `claude-fable-5-thinking-high`. Failed at turn 16: Cursor API usage limit.
- [ ] **CUR-FAB-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 1 — Result: `<chats>/chat-36` — Notes: `claude-fable-5-thinking-high`; `<chats>/chat-17` held unrelated residue and was preserved untouched, so this run uses `chat-36`. Failed at turn 16: Cursor API usage limit.

### Codex

- [ ] **Configuration complete:** Cursor — Codex — 0/3
- [ ] **CUR-COD-01:** Gardening Web Application — Status: `Blocked` — Attempts: 1 — Result: `<chats>/chat-12` — Notes: `gpt-5.6-sol-high` (no medium variant listed; recorded deviation). Failed at turn 15: Cursor Pro+ API usage limit exhausted (resets 2026-08-21); requires operator billing action.
- [ ] **CUR-COD-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 1 — Result: `<chats>/chat-13` — Notes: `gpt-5.6-sol-high`. Failed at turn 18: Cursor API usage limit.
- [ ] **CUR-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 1 — Result: `<chats>/chat-14` — Notes: `gpt-5.6-sol-high`. Failed at turn 15: Cursor API usage limit.

### Grok

- [x] **Configuration complete:** Cursor — Grok — 3/3
- [x] **CUR-GRO-01:** Gardening Web Application — Status: `Complete` — Attempts: 3 — Result: `<chats>/chat-7` — Notes: Tooling-clean, all 30 turns; `cursor-grok-4.5-medium`, model-native medium reasoning; 46 expected, 41 expected hits, 5 missing, 11 additional. Attempts 1–2 were operator-stopped; the container was recreated fresh for attempt 3.
- [x] **CUR-GRO-02:** Community Archive Desktop Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-10` — Notes: Tooling-clean, all 30 turns; `cursor-grok-4.5-medium`; 46 expected, 45 expected hits, 1 missing, 1 additional.
- [x] **CUR-GRO-03:** Neighborhood Emergency Preparedness Program — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-11` — Notes: Tooling-clean, all 30 turns; `cursor-grok-4.5-medium`; 45 expected, 44 expected hits, 1 missing, 7 additional.

### Composer

- [ ] **Configuration complete:** Cursor — Composer — 0/3
- [x] **CUR-COM-01:** Gardening Web Application — Status: `Complete` — Attempts: 2 — Result: `<chats>/chat-6` — Notes: Tooling-clean, all 30 turns; `composer-2.5`, model-native medium reasoning; 46 expected, 15 observed, 33 missing, 2 additional. Attempt 1 was operator-stopped; the container was recreated fresh for attempt 2.
- [x] **CUR-COM-02:** Community Archive Desktop Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-8` — Notes: Tooling-clean, all 30 turns; `composer-2.5`; 46 expected, 6 observed, 40 missing, 0 additional.
- [x] **CUR-COM-03:** Neighborhood Emergency Preparedness Program — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-9` — Notes: Tooling-clean, all 30 turns; `composer-2.5`; 45 expected, 7 observed, 38 missing, 0 additional.

## 5. Pi — Codex

- [x] **Configuration complete:** Pi — Codex — 3/3
- [x] **PI-COD-01:** Gardening Web Application — Status: `Complete` — Attempts: 2 — Result: `<chats>/chat-21` — Notes: Tooling-clean, all 30 turns; `openai-codex/gpt-5.6-sol`, medium reasoning; 46 expected, 8 observed, 40 missing, 2 additional. The attempt-1 evidence container (`<chats>/chat-2`) was removed during operator cleanup, so attempt 2 reran the scenario.
- [x] **PI-COD-02:** Community Archive Desktop Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-22` — Notes: Tooling-clean, all 30 turns; `openai-codex/gpt-5.6-sol`, medium reasoning; 46 expected, 9 observed, 41 missing, 4 additional; supplied-skill visibility verified via get_commands.
- [x] **PI-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-23` — Notes: Tooling-clean, all 30 turns; `openai-codex/gpt-5.6-sol`, medium reasoning; 45 expected, 8 observed, 41 missing, 4 additional; supplied-skill visibility verified via get_commands.

## 6. OpenCode — Codex

- [x] **Configuration complete:** OpenCode — Codex — 3/3
- [x] **OPE-COD-01:** Gardening Web Application — Status: `Complete` — Attempts: 2 — Result: `<chats>/chat-24` — Notes: Tooling-clean, all 30 turns; `openai/gpt-5.6-sol`, medium reasoning; 46 expected, 8 observed, 40 missing, 2 additional. The attempt-1 evidence container (`<chats>/chat-3`) was removed during operator cleanup, so attempt 2 reran the scenario.
- [x] **OPE-COD-02:** Community Archive Desktop Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-25` — Notes: Tooling-clean, all 30 turns; `openai/gpt-5.6-sol`, medium reasoning; 46 expected, 9 observed, 38 missing, 1 additional; native skill discovery verified pre-run.
- [x] **OPE-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Complete` — Attempts: 2 — Result: `<chats>/chat-26` — Notes: Tooling-clean, all 30 turns; `openai/gpt-5.6-sol`, medium reasoning; 45 expected, 9 observed, 39 missing, 3 additional. Attempt 1 failed at the authentication preflight during simultaneous lane launch; attempt 2 ran solo.

## Failure And Blocker Log

Add one row for every failed or blocked attempt before deleting its container for a rerun. Retain the tracker row after a successful rerun so the diagnosis and correction history remain visible even though failed-attempt artifacts are replaced.

| Evaluation ID | Date       | Attempt | Status  | Failure or blocker                                                                                                                                                                                                                                                                  | Resolution or next action                                                                                                                        | Rerun result                                    |
| ------------- | ---------- | ------: | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------------- |
| CUR-COM-01    | 2026-07-22 |       0 | Blocked | Orchestrator policy denied Cursor workspace and transcript export before process creation.                                                                                                                                                                                          | Run from an environment authorized to send the evaluation data to Cursor.                                                                        | Attempt 1 later stopped by the operator.        |
| CUR-GRO-01    | 2026-07-22 |       0 | Blocked | Orchestrator policy denied Cursor workspace and transcript export before process creation.                                                                                                                                                                                          | Run from an environment authorized to send the evaluation data to Cursor.                                                                        | Attempts 1 and 2 later stopped by the operator. |
| CUR-COM-01    | 2026-07-22 |       1 | Stopped | The operator stopped all active runs while turn 16 was running; the harness exited after receiving the termination signal.                                                                                                                                                          | Preserve the incomplete container until the operator requests cleanup or a fresh same-container retry.                                           | Pending.                                        |
| CUR-GRO-01    | 2026-07-22 |       1 | Stopped | The operator stopped all active runs while turn 25 was running; the harness exited after receiving the termination signal.                                                                                                                                                          | Preserve the incomplete container until the operator requests cleanup or a fresh same-container retry.                                           | Pending.                                        |
| CUR-GRO-01    | 2026-07-22 |       2 | Stopped | A concurrent duplicate run existed in `<chats>/chat-7`; the operator's stop request terminated it during turn 8.                                                                                                                                                                    | Treat the duplicate as an orchestration error and recreate the container before any future retry.                                                | Attempt 3 running.                              |
| OPE-COD-03    | 2026-07-22 |       1 | Failed  | OpenCode authentication preflight returned "Unexpected error" when both OpenCode runs launched simultaneously, before any side effects; the concurrent auth check is the suspected race.                                                                                            | Container recreated; attempt 2 relaunched solo after its lane-mate's preflight had settled.                                                      | Attempt 2 complete, tooling-clean.              |
| CLA-COD-03    | 2026-07-22 |       1 | Stopped | The orchestrator session closed and its teardown killed the active harness process at turn 15 (`signal: killed`); the CLI's interrupt handling removed temporary skills and private state.                                                                                          | Container recreated; the chain rerun (CLA-COD-03 then CLA-COD-01) was relaunched in a fresh session.                                             | Attempt 2 running.                              |
| CLA-COD-01    | 2026-07-22 |     dup | Stopped | An orchestration error: a stale queued chain command re-fired CLA-COD-01 into `<chats>/chat-30` after the run had already completed in `<chats>/chat-37`; the deployed known-good CLI auto-creates missing workspaces, defeating the workspace-removal defusal. Stopped at turn 14. | Duplicate terminated; interrupt cleanup verified (no skills, no private state). `<chats>/chat-37` remains the sole accepted CLA-COD-01 evidence. | Not applicable — duplicate, not an attempt.     |
| CUR-COD-01    | 2026-07-22 |       1 | Blocked | Cursor Pro+ monthly API usage limit exhausted at turn 15 (`ActionRequiredError`); API-billed models cannot continue until the cycle resets 2026-08-21 or the operator raises the limit.                                                                                             | Operator decision 2026-07-22: wait — the six runs stay blocked until the Cursor cycle resets 2026-08-21 or the operator raises the limit sooner. | Deferred to cycle reset.                        |
| CUR-COD-02    | 2026-07-22 |       1 | Blocked | Same Cursor API usage limit at turn 18.                                                                                                                                                                                                                                             | Same as CUR-COD-01.                                                                                                                              | Pending operator action.                        |
| CUR-COD-03    | 2026-07-22 |       1 | Blocked | Same Cursor API usage limit at turn 15.                                                                                                                                                                                                                                             | Same as CUR-COD-01.                                                                                                                              | Pending operator action.                        |
| CUR-FAB-01    | 2026-07-22 |       1 | Blocked | Same Cursor API usage limit at turn 21.                                                                                                                                                                                                                                             | Same as CUR-COD-01.                                                                                                                              | Pending operator action.                        |
| CUR-FAB-02    | 2026-07-22 |       1 | Blocked | Same Cursor API usage limit at turn 16.                                                                                                                                                                                                                                             | Same as CUR-COD-01.                                                                                                                              | Pending operator action.                        |
| CUR-FAB-03    | 2026-07-22 |       1 | Blocked | Same Cursor API usage limit at turn 16.                                                                                                                                                                                                                                             | Same as CUR-COD-01.                                                                                                                              | Pending operator action.                        |

## Campaign Notes

- **Reasoning target:** Medium, or the closest documented harness-specific equivalent where no independent medium control exists.
- **Exact model identifiers:** Completed: Claude Code and OpenAI Codex used `gpt-5.6-sol`; Pi used `openai-codex/gpt-5.6-sol`; OpenCode used `openai/gpt-5.6-sol`. Stopped Cursor attempts used `composer-2.5` and `cursor-grok-4.5-medium`.
- **Active lane identifiers:** None.
- **Preflight:** Known-good CLI `1991b20c9042`; Codex, Pi, OpenCode, Cursor, and Claude Code doctor checks passed. Harness versions: Pi `0.80.10`, OpenCode `1.18.4`, Cursor `2026.07.16-899851b`, and Claude Code `2.1.205`. Codex reported the known non-blocking version-parse warning (`1)` versus tested `0.144.6`).
- **Orchestrator environment:** All attached sessions and detached `screen` sessions are terminal; process inspection found no remaining campaign process.
- **Campaign started:** 2026-07-22 01:28 ACST.
- **Campaign completed:** —
- **Cleanup status:** Completed containers `<chats>/chat-1`, `<chats>/chat-2`, `<chats>/chat-3`, and `<chats>/chat-5` are retained. Incomplete Cursor containers `<chats>/chat-6` and `<chats>/chat-7` are retained for diagnosis and must be recreated before retry. The two prompt-only `nohup` Cursor containers were preserved outside the campaign root before fresh same-number relaunches.
- **General notes:** Four runs completed, the operator stopped all remaining runs, and no campaign process is active.

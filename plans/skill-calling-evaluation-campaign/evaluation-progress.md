# Skill-Calling Evaluation Campaign Progress

## Purpose

Track the subscription-bounded skill-calling evaluation campaign one scenario run at a time. The campaign contains ten harness-and-model configurations, with all three governed scenarios run once for each configuration, for a total of 30 evaluation runs.

This document records campaign execution progress only. The governed scenarios, instrumentation, result derivation, and valid-run requirements remain owned by the evaluation assets and their existing contracts.

## Campaign Configuration

| Harness      | Model label           | Reasoning target                       |   Runs | Complete |
| ------------ | --------------------- | -------------------------------------- | -----: | -------: |
| Claude Code  | Codex                 | Medium                                 |      3 |      0/3 |
| OpenAI Codex | GPT-5.6 Sol           | Medium                                 |      3 |      0/3 |
| Claude Code  | Fable                 | Medium                                 |      3 |      0/3 |
| Cursor       | Fable                 | Medium or model-native equivalent      |      3 |      0/3 |
| Cursor       | Codex                 | Medium or model-native equivalent      |      3 |      0/3 |
| Cursor       | Grok                  | Medium or model-native equivalent      |      3 |      0/3 |
| Cursor       | Composer              | Medium or model-native equivalent      |      3 |      1/3 |
| Pi           | Codex                 | Medium                                 |      3 |      0/3 |
| OpenCode     | Codex                 | Medium or closest supported equivalent |      3 |      1/3 |
| Kilo Code    | Codex                 | Medium or closest supported equivalent |      3 |      0/3 |
| **Total**    | **10 configurations** | **One consistent target**              | **30** | **2/30** |

Use one reasoning target across the campaign, or the closest documented harness-specific equivalent where a harness does not expose the same control. Record the exact effective model identifier, model version or alias, harness version, and reasoning setting with each result.

## Progress Summary

| Measure                       | Progress |
| ----------------------------- | -------: |
| Configuration suites complete |     0/10 |
| Evaluation runs complete      |     2/30 |
| Evaluation runs running       |        0 |
| Evaluation runs pending       |     0/30 |
| Evaluation runs failed        |        9 |
| Evaluation runs blocked       |       19 |
| Overall completion            |       7% |

## Tracking Rules

- Check an evaluation only after the complete governed scenario and its instrumentation finish successfully and the required result artifacts are retained.
- A tooling-complete run with zero observed skill calls is a completed evaluation and may be checked.
- Leave an exhausted tooling failure unchecked, set its status to `Failed`, and record it in the failure log; the adjacent orchestration prompt owns the diagnosed fresh-container retry budget and same-cause systemic lane closure.
- Use `Blocked` when an evaluation cannot start because access, authentication, model availability, or another prerequisite is unavailable.
- Check a configuration heading only after all three scenario evaluations beneath it are complete.
- Update the configuration table and progress summary whenever an evaluation status changes.

Status values: `Pending`, `Running`, `Complete`, `Failed`, or `Blocked`.

## Execution Dependencies

- OpenAI Codex — Sol, Pi — Codex, OpenCode — Codex, Kilo — Codex, Claude Code — Codex, and all four Cursor lanes start immediately in parallel, subject to the campaign-wide limit of ten simultaneous evaluation runs and at most one active `claude-code`-harness run.
- Each configuration lane runs its three scenarios sequentially. A tooling failure may receive up to two diagnosed fresh-container retries, and every later scenario remains eligible unless two runs in the lane fail from the same systemic cause.
- All four Cursor lanes start concurrently. Recurrent concurrent rate limits reduce Cursor concurrency for later retries and runs.
- Claude Code — Fable runs last, alone, only after every other lane is terminal and the Claude Code — Codex proxy is stopped and verified gone.
- The campaign orchestrator starts and monitors evaluation commands directly and is the only writer of this progress document.
- Every evaluation uses its own external workspace and its own retained output location.
- Orchestrator-environment failures consume no run attempt budget. Before launch, the orchestrator must prove direct read-write cleanup access under `<chats>` and prove that a detached long-lived process survives its launching command group.

The adjacent `campaign-orchestration-prompt.md` owns scheduling, command launch contracts, model-identifier resolution, troubleshooting, continuation and halt rules, and the progress-update procedure.

## 1. Claude Code — Codex

- [ ] **Configuration complete:** Claude Code — Codex — 0/3
- [ ] **CLA-COD-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Operator stop during turn 25; evaluator received `SIGINT`, cleanup completed, and unsuccessful `chat-85` was removed.
- [ ] **CLA-COD-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **CLA-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

## 2. OpenAI Codex — GPT-5.6 Sol

- [ ] **Configuration complete:** OpenAI Codex — GPT-5.6 Sol — 0/3
- [ ] **COD-SOL-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Operator stop during turn 23; evaluator received `SIGINT`, cleanup completed, and unsuccessful `chat-81` was removed.
- [ ] **COD-SOL-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **COD-SOL-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

## 3. Claude Code — Fable

- [ ] **Configuration complete:** Claude Code — Fable — 0/3
- [ ] **CLA-FAB-01:** Gardening Web Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before the isolated final lane.
- [ ] **CLA-FAB-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before the isolated final lane.
- [ ] **CLA-FAB-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before the isolated final lane.

## 4. Cursor

### Fable

- [ ] **Configuration complete:** Cursor — Fable — 0/3
- [ ] **CUR-FAB-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Operator stop during turn 21; evaluator received `SIGINT`, cleanup completed, and unsuccessful `chat-89` was removed.
- [ ] **CUR-FAB-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **CUR-FAB-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

### Codex

- [ ] **Configuration complete:** Cursor — Codex — 0/3
- [ ] **CUR-COD-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Operator stop during turn 19; evaluator received `SIGINT`, cleanup completed, and unsuccessful `chat-88` was removed.
- [ ] **CUR-COD-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **CUR-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

### Grok

- [ ] **Configuration complete:** Cursor — Grok — 0/3
- [ ] **CUR-GRO-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Operator stop during turn 27; evaluator received `SIGINT`, cleanup completed, and unsuccessful `chat-87` was removed.
- [ ] **CUR-GRO-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **CUR-GRO-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

### Composer

- [ ] **Configuration complete:** Cursor — Composer — 0/3
- [x] **CUR-COM-01:** Gardening Web Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-86` — Notes: Cursor `2026.07.16-899851b`; model `composer-2.5`; medium; 46 expected, 9 observed, 37 missing, 0 additional.
- [ ] **CUR-COM-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **CUR-COM-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

## 5. Pi — Codex

- [ ] **Configuration complete:** Pi — Codex — 0/3
- [ ] **PI-COD-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Operator stop during turn 23; evaluator received `SIGINT`, cleanup completed, and unsuccessful `chat-82` was removed.
- [ ] **PI-COD-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **PI-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

## 6. OpenCode — Codex

- [ ] **Configuration complete:** OpenCode — Codex — 0/3
- [x] **OPE-COD-01:** Gardening Web Application — Status: `Complete` — Attempts: 1 — Result: `<chats>/chat-83` — Notes: OpenCode `1.18.4`; model `openai/gpt-5.6-sol`; medium; 46 expected, 8 observed, 40 missing, 2 additional.
- [ ] **OPE-COD-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.
- [ ] **OPE-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Operator stopped the campaign before launch.

## 7. Kilo Code — Codex

- [ ] **Configuration complete:** Kilo Code — Codex — 0/3
- [ ] **KIL-COD-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Kilo omitted its terminal protocol event after a protected skill-root edit was auto-rejected on turn 10; no permitted correction exists; cleanup verified and unsuccessful `chat-84` removed.
- [ ] **KIL-COD-02:** Community Archive Desktop Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Reproduced the Kilo lane's protected skill-root edit rejection and missing terminal protocol event on turn 8; cleanup verified and unsuccessful `chat-90` removed.
- [ ] **KIL-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Failed` — Attempts: 0 — Result: — Notes: Not launched after two Kilo runs failed on the same diagnosed systemic permission and terminal-protocol cause.

## Failure And Blocker Log

Add one row for every failed or blocked attempt. Retain earlier rows after a successful rerun so the campaign history remains visible.

| Evaluation ID | Date       | Attempt | Status  | Failure or blocker                                                                      | Resolution or next action                                                     | Rerun result |
| ------------- | ---------- | ------: | ------- | --------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ------------ |
| CLA-COD-01    | 2026-07-21 |       1 | Failed  | Operator-requested interruption during turn 7; evaluator exited on `SIGINT`.            | Cleanup completed; lane closed permanently under the campaign fail-fast rule. | —            |
| CLA-COD-02    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `CLA-COD-01`.                                  | Do not start this run.                                                        | —            |
| CLA-COD-03    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `CLA-COD-01`.                                  | Do not start this run.                                                        | —            |
| COD-SOL-01    | 2026-07-21 |       1 | Blocked | Prepared launch in `<chats>/chat-6` did not begin before the campaign halt.             | Retain the empty container.                                                   | —            |
| COD-SOL-02    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| COD-SOL-03    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CLA-FAB-01    | 2026-07-21 |       0 | Blocked | Campaign halted before the isolated final lane could start.                             | Do not start this run.                                                        | —            |
| CLA-FAB-02    | 2026-07-21 |       0 | Blocked | Campaign halted before the isolated final lane could start.                             | Do not start this run.                                                        | —            |
| CLA-FAB-03    | 2026-07-21 |       0 | Blocked | Campaign halted before the isolated final lane could start.                             | Do not start this run.                                                        | —            |
| CUR-FAB-01    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CUR-FAB-02    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CUR-FAB-03    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CUR-COD-01    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CUR-COD-02    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CUR-COD-03    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CUR-GRO-01    | 2026-07-21 |       1 | Failed  | Campaign halt interruption during turn 1 in `<chats>/chat-11`.                          | Evaluator cleanup verified; lane closed.                                      | —            |
| CUR-GRO-02    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `CUR-GRO-01`.                                  | Do not start this run.                                                        | —            |
| CUR-GRO-03    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `CUR-GRO-01`.                                  | Do not start this run.                                                        | —            |
| CUR-COM-01    | 2026-07-21 |       1 | Blocked | Prepared launch in `<chats>/chat-10` did not begin before the campaign halt.            | Retain the empty container.                                                   | —            |
| CUR-COM-02    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CUR-COM-03    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| PI-COD-01     | 2026-07-21 |       1 | Failed  | Campaign halt interruption during turn 2 in `<chats>/chat-7`.                           | Evaluator cleanup verified; lane closed.                                      | —            |
| PI-COD-02     | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `PI-COD-01`.                                   | Do not start this run.                                                        | —            |
| PI-COD-03     | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `PI-COD-01`.                                   | Do not start this run.                                                        | —            |
| OPE-COD-01    | 2026-07-21 |       1 | Failed  | Campaign halt interruption during turn 1 in `<chats>/chat-8`.                           | Evaluator cleanup verified; lane closed.                                      | —            |
| OPE-COD-02    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `OPE-COD-01`.                                  | Do not start this run.                                                        | —            |
| OPE-COD-03    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `OPE-COD-01`.                                  | Do not start this run.                                                        | —            |
| KIL-COD-01    | 2026-07-21 |       1 | Blocked | Required outer sandbox escalation was denied before `<chats>/chat-9` ran.               | Campaign halted immediately.                                                  | —            |
| KIL-COD-01    | 2026-07-21 |       2 | Failed  | Malformed protocol after an auto-rejected edit permission request in `<chats>/chat-30`. | Cleanup verified; retry once in fresh `<chats>/chat-33`.                      | Failed       |
| KIL-COD-02    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| KIL-COD-03    | 2026-07-21 |       0 | Blocked | Campaign halted before launch.                                                          | Do not start this run.                                                        | —            |
| CLA-COD-01    | 2026-07-21 |       2 | Failed  | External orchestrator ended during turn 15; harness exited on signal killed.            | Cleanup verified; lane closed and campaign halted.                            | —            |
| COD-SOL-01    | 2026-07-21 |       2 | Failed  | External orchestrator ended during turn 10; harness exited on signal terminated.        | Cleanup verified; lane closed and campaign halted.                            | —            |
| CUR-COM-01    | 2026-07-21 |       2 | Failed  | External orchestrator ended during turn 9; harness exited on signal killed.             | Cleanup verified; lane closed and campaign halted.                            | —            |
| PI-COD-01     | 2026-07-21 |       2 | Failed  | External orchestrator ended during turn 14; evaluation context was canceled.            | Cleanup verified; lane closed and campaign halted.                            | —            |
| OPE-COD-01    | 2026-07-21 |       2 | Failed  | External orchestrator ended during turn 18; harness exited on signal killed.            | Cleanup verified; lane closed and campaign halted.                            | —            |
| KIL-COD-01    | 2026-07-21 |       3 | Failed  | External orchestrator ended during retry turn 6; harness exited on signal killed.       | Cleanup verified; lane closed and campaign halted.                            | —            |
| CLA-COD-02    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `CLA-COD-01` failed.                                          | Do not start this run.                                                        | —            |
| CLA-COD-03    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `CLA-COD-01` failed.                                          | Do not start this run.                                                        | —            |
| COD-SOL-02    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `COD-SOL-01` failed.                                          | Do not start this run.                                                        | —            |
| COD-SOL-03    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `COD-SOL-01` failed.                                          | Do not start this run.                                                        | —            |
| CLA-FAB-01    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before the isolated final lane could start.                     | Do not start this run.                                                        | —            |
| CLA-FAB-02    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before the isolated final lane could start.                     | Do not start this run.                                                        | —            |
| CLA-FAB-03    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before the isolated final lane could start.                     | Do not start this run.                                                        | —            |
| CUR-FAB-01    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-FAB-02    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-FAB-03    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-COD-01    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-COD-02    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-COD-03    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-GRO-01    | 2026-07-21 |       1 | Blocked | Resumed campaign halted before relaunch.                                                | Retain `<chats>/chat-11`; do not relaunch.                                    | —            |
| CUR-GRO-02    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-GRO-03    | 2026-07-21 |       0 | Blocked | Resumed campaign halted before launch.                                                  | Do not start this run.                                                        | —            |
| CUR-COM-02    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `CUR-COM-01` failed.                                          | Do not start this run.                                                        | —            |
| CUR-COM-03    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `CUR-COM-01` failed.                                          | Do not start this run.                                                        | —            |
| PI-COD-02     | 2026-07-21 |       0 | Blocked | Lane closed after resumed `PI-COD-01` failed.                                           | Do not start this run.                                                        | —            |
| PI-COD-03     | 2026-07-21 |       0 | Blocked | Lane closed after resumed `PI-COD-01` failed.                                           | Do not start this run.                                                        | —            |
| OPE-COD-02    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `OPE-COD-01` failed.                                          | Do not start this run.                                                        | —            |
| OPE-COD-03    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `OPE-COD-01` failed.                                          | Do not start this run.                                                        | —            |
| KIL-COD-02    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `KIL-COD-01` exhausted its attempts.                          | Do not start this run.                                                        | —            |
| KIL-COD-03    | 2026-07-21 |       0 | Blocked | Lane closed after resumed `KIL-COD-01` exhausted its attempts.                          | Do not start this run.                                                        | —            |
| COD-SOL-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| PI-COD-01     | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| OPE-COD-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| KIL-COD-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| CLA-COD-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| CUR-COM-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| CUR-GRO-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| CUR-COD-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| CUR-FAB-01    | 2026-07-21 |       0 | Blocked | External data-export approval required; outer launch was rejected before execution.     | No run created; informed operator approval obtained.                          | Approved     |
| COD-SOL-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-72` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Failed       |
| PI-COD-01     | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-73` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Failed       |
| OPE-COD-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-74` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Complete     |
| KIL-COD-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-75` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Failed       |
| CLA-COD-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-76` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Failed       |
| CUR-COM-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-77` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Complete     |
| CUR-GRO-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-78` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Failed       |
| CUR-COD-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-79` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Failed       |
| CUR-FAB-01    | 2026-07-21 |       0 | Failed  | Launchd exited 127 before the runner started; empty `<chats>/chat-80` retained.         | Crash loop stopped; screen route proven; attempt remains eligible.            | Failed       |
| KIL-COD-01    | 2026-07-21 |       1 | Failed  | Protected skill-root edit was rejected on turn 10; Kilo omitted terminal `step_finish`. | No permitted correction; cleanup verified; unsuccessful container removed.    | —            |
| KIL-COD-02    | 2026-07-21 |       1 | Failed  | Same protected skill-root edit rejection recurred; Kilo omitted terminal `step_finish`. | Cleanup verified; unsuccessful container removed; close systemic lane.        | —            |
| KIL-COD-03    | 2026-07-21 |       0 | Failed  | Two Kilo runs failed on the same systemic permission and terminal-protocol cause.       | Do not launch; lane closed under the same-cause efficiency rule.              | —            |
| CLA-COD-01    | 2026-07-21 |       1 | Failed  | Operator stopped the run during turn 25; evaluator exited after direct `SIGINT`.        | Cleanup verified; unsuccessful container removed.                             | —            |
| COD-SOL-01    | 2026-07-21 |       1 | Failed  | Operator stopped the run during turn 23; evaluator exited after direct `SIGINT`.        | Cleanup verified; unsuccessful container removed.                             | —            |
| CUR-FAB-01    | 2026-07-21 |       1 | Failed  | Operator stopped the run during turn 21; evaluator exited after direct `SIGINT`.        | Cleanup verified; unsuccessful container removed.                             | —            |
| CUR-COD-01    | 2026-07-21 |       1 | Failed  | Operator stopped the run during turn 19; evaluator exited after direct `SIGINT`.        | Cleanup verified; unsuccessful container removed.                             | —            |
| CUR-GRO-01    | 2026-07-21 |       1 | Failed  | Operator stopped the run during turn 27; evaluator exited after direct `SIGINT`.        | Cleanup verified; unsuccessful container removed.                             | —            |
| PI-COD-01     | 2026-07-21 |       1 | Failed  | Operator stopped the run during turn 23; evaluator exited with context canceled.        | Cleanup verified; unsuccessful container removed.                             | —            |
| CLA-COD-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CLA-COD-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| COD-SOL-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| COD-SOL-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CLA-FAB-01    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before the isolated final lane.                           | No launch; attempt budget untouched.                                          | —            |
| CLA-FAB-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before the isolated final lane.                           | No launch; attempt budget untouched.                                          | —            |
| CLA-FAB-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before the isolated final lane.                           | No launch; attempt budget untouched.                                          | —            |
| CUR-FAB-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CUR-FAB-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CUR-COD-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CUR-COD-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CUR-GRO-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CUR-GRO-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CUR-COM-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| CUR-COM-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| PI-COD-02     | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| PI-COD-03     | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| OPE-COD-02    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |
| OPE-COD-03    | 2026-07-21 |       0 | Blocked | Operator stopped the campaign before launch.                                            | No launch; attempt budget untouched.                                          | —            |

## Campaign Notes

- **Reasoning target:** Medium, or the closest documented harness-specific equivalent where no independent medium control exists.
- **Exact model identifiers:** Claude Code — Codex uses `gpt-5.6-sol`; Cursor — Composer uses `composer-2.5` as the documented fallback because no `-medium` variant is listed; Cursor — Grok uses `cursor-grok-4.5-medium`; Cursor — Codex uses `gpt-5.6-sol-medium`; Cursor — Fable uses `claude-fable-5-medium`.
- **Started lane identifiers:** OpenAI Codex `0.144.6` with `gpt-5.6-sol`; Pi `0.80.10` with `openai-codex/gpt-5.6-sol`; OpenCode `1.18.4` with `openai/gpt-5.6-sol`; Kilo `7.4.11` with `openai/gpt-5.6-sol`; Claude Code `2.1.205` through claudex with `gpt-5.6-sol`; Cursor `2026.07.16-899851b` with Composer `composer-2.5`, Grok `cursor-grok-4.5-medium`, Codex `gpt-5.6-sol-medium`, and Fable `claude-fable-5-medium`. Every started lane used medium reasoning or its recorded model-native equivalent. Claude Code — Fable did not start.
- **Preflight:** Known-good `177a68d9cbfd` resolved. All seven executable routes passed doctor. Current Cursor identifiers were resolved from the native listing retained at `output/campaign/preflight-cursor-models.txt`. The Claude Code — Codex launcher exists.
- **Orchestrator environment:** Create, write, read, and remove access under `<chats>` passed. A normal detached child was reaped with its launcher. Launchd survived but lost Desktop access and exited the runner with code 127, so its crash loops were stopped. A detached screen-owned process survived later commands, retained `<chats>` read-write access, and cleaned up its proof artifacts. Campaign runs use that proven screen route.
- **Retention override:** At operator direction, retain only successful runs produced by this orchestration. All historical `chat-1` through `chat-80` and unsuccessful campaign containers were removed after cleanup verification. Only successful `<chats>/chat-83` and `<chats>/chat-86` remain. The last allocated number was 90.
- **Campaign started:** 2026-07-21
- **Campaign completed:** Stopped by operator on 2026-07-21.
- **Cleanup status:** No evaluation screen, evaluator, harness process, private output state, or run-owned system-temp residue remains. Native skill roots retain only scenario-created workspace output. The claudex proxy reports stopped with readiness unavailable, and no proxy listener remains.
- **Scenario revision (2026-07-21):** Campaign runs of `gardening-web-application` had evaluated agents launch headless Chrome with fresh temporary profiles, which raised repeated macOS "Keychain Not Found" prompts on the operator's machine. Turns 3, 6, 21, 22, and 28 were reworded so all verification stays in Node-run checks and the user owns all in-browser checking; the scenario remains 30 turns and the 46 expected skill calls are unchanged. Runs after known-good `177a68d9cbfd` need a CLI rebuilt from a commit containing this revision to embed the updated turns.
- **General notes:** Board reset on 2026-07-21 at operator direction because earlier failed or blocked entries reflected orchestrator-environment problems rather than lane defects. The operator granted informed approval for external model data export, then later stopped all runs. Two governed runs completed, six active runs were interrupted and cleaned, the Kilo lane closed after two same-cause systemic failures, and every unstarted run was blocked by the operator stop. Final outcome: 2 complete, 9 failed, 19 blocked; 30/30 was not reached.

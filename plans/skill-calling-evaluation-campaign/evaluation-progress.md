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
| Cursor       | Composer              | Medium or model-native equivalent      |      3 |      0/3 |
| Pi           | Codex                 | Medium                                 |      3 |      0/3 |
| OpenCode     | Codex                 | Medium or closest supported equivalent |      3 |      0/3 |
| Kilo Code    | Codex                 | Medium or closest supported equivalent |      3 |      0/3 |
| **Total**    | **10 configurations** | **One consistent target**              | **30** | **0/30** |

Use one reasoning target across the campaign, or the closest documented harness-specific equivalent where a harness does not expose the same control. Record the exact effective model identifier, model version or alias, harness version, and reasoning setting with each result.

## Progress Summary

| Measure                       | Progress |
| ----------------------------- | -------: |
| Configuration suites complete |     0/10 |
| Evaluation runs complete      |     0/30 |
| Evaluation runs pending       |    21/30 |
| Evaluation runs failed        |        1 |
| Evaluation runs blocked       |        2 |
| Overall completion            |       0% |

## Tracking Rules

- Check an evaluation only after the complete governed scenario and its instrumentation finish successfully and the required result artifacts are retained.
- A tooling-complete run with zero observed skill calls is a completed evaluation and may be checked.
- Leave a tooling failure unchecked, set its status to `Failed`, record it in the failure log, and rerun it after the cause is resolved.
- Use `Blocked` when an evaluation cannot start because access, authentication, model availability, or another prerequisite is unavailable.
- Check a configuration heading only after all three scenario evaluations beneath it are complete.
- Update the configuration table and progress summary whenever an evaluation status changes.

Status values: `Pending`, `Running`, `Complete`, `Failed`, or `Blocked`.

## Execution Dependencies

- All lanes except Claude Code — Fable start immediately and run in parallel, subject to the campaign-wide limit of six simultaneous evaluation runs, at most two concurrent Cursor runs, and at most one active `claude-code`-harness run.
- Each configuration lane runs its three scenarios sequentially; a tooling failure closes the lane permanently with no retries.
- Claude Code — Fable runs last, alone, only after every other lane is terminal and the Claude Code — Codex proxy is stopped and verified gone.
- The campaign orchestrator starts and monitors evaluation commands directly and is the only writer of this progress document.
- Every evaluation uses its own external workspace and its own retained output location.
- OpenAI Codex evaluations launched from Codex require command-scoped outer-sandbox escalation for the exact known-good evaluation command. The nested evaluator-owned Codex sandbox and approval settings remain unchanged.

The adjacent `campaign-orchestration-prompt.md` owns scheduling, command launch contracts, model-identifier resolution, fail-fast and halt rules, and the progress-update procedure.

## 1. Claude Code — Codex

- [ ] **Configuration complete:** Claude Code — Codex — 0/3
- [ ] **CLA-COD-01:** Gardening Web Application — Status: `Failed` — Attempts: 1 — Result: — Notes: Operator-requested interruption during turn 7; evaluator cleanup completed.
- [ ] **CLA-COD-02:** Community Archive Desktop Application — Status: `Blocked` — Attempts: 0 — Result: — Notes: Lane closed permanently after `CLA-COD-01` failed under the campaign fail-fast rule.
- [ ] **CLA-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Blocked` — Attempts: 0 — Result: — Notes: Lane closed permanently after `CLA-COD-01` failed under the campaign fail-fast rule.

## 2. OpenAI Codex — GPT-5.6 Sol

- [ ] **Configuration complete:** OpenAI Codex — GPT-5.6 Sol — 0/3
- [ ] **COD-SOL-01:** Gardening Web Application — Status: `Running` — Attempts: 1 — Result: `<chats>/chat-6` — Notes: —
- [ ] **COD-SOL-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **COD-SOL-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 3. Claude Code — Fable

- [ ] **Configuration complete:** Claude Code — Fable — 0/3
- [ ] **CLA-FAB-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CLA-FAB-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CLA-FAB-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 4. Cursor

### Fable

- [ ] **Configuration complete:** Cursor — Fable — 0/3
- [ ] **CUR-FAB-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-FAB-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-FAB-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

### Codex

- [ ] **Configuration complete:** Cursor — Codex — 0/3
- [ ] **CUR-COD-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

### Grok

- [ ] **Configuration complete:** Cursor — Grok — 0/3
- [ ] **CUR-GRO-01:** Gardening Web Application — Status: `Running` — Attempts: 1 — Result: `<chats>/chat-11` — Notes: Model `cursor-grok-4.5-medium`.
- [ ] **CUR-GRO-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-GRO-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

### Composer

- [ ] **Configuration complete:** Cursor — Composer — 0/3
- [ ] **CUR-COM-01:** Gardening Web Application — Status: `Running` — Attempts: 1 — Result: `<chats>/chat-10` — Notes: Model `composer-2.5`; fallback from the unavailable medium target.
- [ ] **CUR-COM-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-COM-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 5. Pi — Codex

- [ ] **Configuration complete:** Pi — Codex — 0/3
- [ ] **PI-COD-01:** Gardening Web Application — Status: `Running` — Attempts: 1 — Result: `<chats>/chat-7` — Notes: —
- [ ] **PI-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **PI-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 6. OpenCode — Codex

- [ ] **Configuration complete:** OpenCode — Codex — 0/3
- [ ] **OPE-COD-01:** Gardening Web Application — Status: `Running` — Attempts: 1 — Result: `<chats>/chat-8` — Notes: —
- [ ] **OPE-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **OPE-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 7. Kilo Code — Codex

- [ ] **Configuration complete:** Kilo Code — Codex — 0/3
- [ ] **KIL-COD-01:** Gardening Web Application — Status: `Running` — Attempts: 1 — Result: `<chats>/chat-9` — Notes: —
- [ ] **KIL-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **KIL-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## Failure And Blocker Log

Add one row for every failed or blocked attempt. Retain earlier rows after a successful rerun so the campaign history remains visible.

| Evaluation ID | Date       | Attempt | Status | Failure or blocker                                                           | Resolution or next action                                                                  | Rerun result |
| ------------- | ---------- | ------: | ------ | ---------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------ | ------------ |
| CLA-COD-01    | 2026-07-21 |       1 | Failed  | Operator-requested interruption during turn 7; evaluator exited on `SIGINT`. | Cleanup completed; lane closed permanently under the campaign fail-fast rule. | —            |
| CLA-COD-02    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `CLA-COD-01`.                        | Do not start this run.                                                         | —            |
| CLA-COD-03    | 2026-07-21 |       0 | Blocked | Lane closed after the tooling failure in `CLA-COD-01`.                        | Do not start this run.                                                         | —            |

## Campaign Notes

- **Reasoning target:** Medium, or the closest documented harness-specific equivalent where no independent medium control exists.
- **Exact model identifiers:** Claude Code — Codex uses `gpt-5.6-sol`; Cursor — Composer uses `composer-2.5` (fallback because no medium variant was listed); Cursor — Grok uses `cursor-grok-4.5-medium`; Cursor — Codex uses `gpt-5.6-sol-medium`; Cursor — Fable uses `claude-fable-5-medium`.
- **Preflight:** Known-good `177a68d9cbfd` resolved. All seven executable routes passed doctor. Cursor identifiers were resolved from the 2026-07-21 native listing retained at `output/campaign/preflight-cursor-models.txt`. The Claude Code — Codex launcher exists.
- **Campaign started:** 2026-07-21
- **Campaign completed:** —
- **General notes:** Claude Code — Codex is terminal with one failed run and two blocked runs because a failed lane is permanently closed.

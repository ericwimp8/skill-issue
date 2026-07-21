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
| Evaluation runs running       |        0 |
| Evaluation runs pending       |    30/30 |
| Evaluation runs failed        |        0 |
| Evaluation runs blocked       |        0 |
| Overall completion            |       0% |

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
- [ ] **CLA-COD-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CLA-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CLA-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 2. OpenAI Codex — GPT-5.6 Sol

- [ ] **Configuration complete:** OpenAI Codex — GPT-5.6 Sol — 0/3
- [ ] **COD-SOL-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
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
- [ ] **CUR-GRO-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-GRO-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-GRO-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

### Composer

- [ ] **Configuration complete:** Cursor — Composer — 0/3
- [ ] **CUR-COM-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-COM-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **CUR-COM-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 5. Pi — Codex

- [ ] **Configuration complete:** Pi — Codex — 0/3
- [ ] **PI-COD-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **PI-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **PI-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 6. OpenCode — Codex

- [ ] **Configuration complete:** OpenCode — Codex — 0/3
- [ ] **OPE-COD-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **OPE-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **OPE-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## 7. Kilo Code — Codex

- [ ] **Configuration complete:** Kilo Code — Codex — 0/3
- [ ] **KIL-COD-01:** Gardening Web Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **KIL-COD-02:** Community Archive Desktop Application — Status: `Pending` — Attempts: 0 — Result: — Notes: —
- [ ] **KIL-COD-03:** Neighborhood Emergency Preparedness Program — Status: `Pending` — Attempts: 0 — Result: — Notes: —

## Failure And Blocker Log

Add one row for every failed or blocked attempt. Retain earlier rows after a successful rerun so the campaign history remains visible.

| Evaluation ID | Date | Attempt | Status | Failure or blocker | Resolution or next action | Rerun result |
| ------------- | ---- | ------: | ------ | ------------------ | ------------------------- | ------------ |
| —             | —    |       — | —      | —                  | —                         | —            |

## Campaign Notes

- **Reasoning target:** Medium, or the closest documented harness-specific equivalent where no independent medium control exists.
- **Exact model identifiers:** Resolve during campaign preflight.
- **Started lane identifiers:** —
- **Preflight:** —
- **Orchestrator environment:** —
- **Campaign started:** —
- **Campaign completed:** —
- **Cleanup status:** Fresh campaign state; no retained containers or campaign output.
- **General notes:** —

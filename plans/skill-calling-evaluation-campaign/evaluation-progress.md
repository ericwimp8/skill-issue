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

Add one row for every failed or blocked attempt before deleting its container for a rerun. Retain the tracker row after a successful rerun so the diagnosis and correction history remain visible even though failed-attempt artifacts are replaced.

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

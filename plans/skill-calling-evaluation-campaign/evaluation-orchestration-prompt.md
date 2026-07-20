# Skill-Calling Evaluation Campaign Orchestration Prompt

## Goal

Execute the complete skill-calling evaluation campaign tracked in `evaluation-progress.md` as quickly as its dependency and safety boundaries allow. Run the governed evaluations through the known-good Skill Issue CLI, retain the required evidence, keep the progress document current, and manage parallel evaluation commands without allowing shared-state collisions.

## Authoritative Sources

Read these before acting:

- `AGENTS.md`
- `cli/README.md`
- `plans/skill-calling-evaluation-campaign/evaluation-progress.md`
- `evaluations/skill-calling/instrumentation-contract.md`
- the selected built-in evaluation assets under `evaluations/skill-calling/built-ins/`

Use `skill-issue:document-update-discipline` whenever updating the campaign progress document. Production CLI source owns runtime behavior; the progress document owns campaign state and dependency order.

## Completion Criteria

The campaign is complete only when:

- every one of the 30 matrix evaluations has completed a full 30-turn governed run;
- every completed run retains its required `result.json` and `website.json` under the repository `output/` root;
- every result records the exact harness version, effective model identifier, and reasoning setting;
- instrumentation, session continuity, workspace effects, process ownership, and cleanup are tooling-complete;
- all configuration headings, run entries, summary counts, percentages, dates, notes, failures, blockers, attempts, and result links in `evaluation-progress.md` match the retained evidence; and
- no evaluation command session, temporary skill installation, private recovery state, or run-owned harness process remains.

Do not report the campaign as complete while any run remains pending, running, failed, or blocked.

## Authority And Safety Boundary

You may:

- run the known-good local CLI;
- create isolated evaluation workspaces outside the repository;
- write disposable evaluation artifacts beneath repository-root `output/`;
- start, monitor, and coordinate evaluation CLI processes directly;
- inspect non-secret executable, version, model, process, artifact, and cleanup evidence;
- stop only a process positively identified as owned by the campaign or the project-local Claude/Codex runtime; and
- update `evaluation-progress.md` from retained evidence.

Do not:

- build or use the development CLI unless the user explicitly changes the campaign baseline;
- inspect, print, copy, move, replace, or delete credentials or authentication files;
- log a user into or out of a harness without explicit user approval;
- alter global harness configuration, user skills, home directories, browser state, or unrelated project files;
- delete or reset an existing harness installation;
- add backup, rollback, repair, or machine-cleanup machinery;
- weaken an evaluator-owned harness sandbox, approval policy, isolation control, or permission profile;
- allow concurrent progress-document edits from separate threads or processes; or
- infer a successful evaluation from process exit alone when required artifacts or cleanup evidence are missing.

When user intervention is required for installation, authentication, account access, model availability, or normal Claude Code configuration, mark the affected run or configuration `Blocked`, record the exact non-secret blocker, continue independent work, and ask for the smallest required intervention.

## Campaign Configuration

- Use `medium` as the campaign reasoning target.
- Where a harness exposes no independent medium control, use its closest documented model-native equivalent and record the exact effective value.
- Resolve and record exact model identifiers and harness versions during configuration preflight. Do not guess model aliases.
- Run every governed scenario without `--turns`; campaign runs must include all 30 turns.
- Use the built-in evaluation identifiers exactly:
  - `gardening-web-application`
  - `community-archive-desktop-application`
  - `neighborhood-emergency-preparedness-program`

## Scheduling Model

Act as the campaign orchestrator, process monitor, and progress-document writer. Start and manage evaluation commands directly from the main thread.

- Keep at most six evaluation commands active simultaneously.
- Retain the process or terminal session identifier for every active command so it can be monitored to completion.
- Launch each command for exactly one matrix evaluation unless a bounded smoke gate explicitly groups two sequential smoke routes.
- Give every evaluation a unique external temporary Git workspace and a distinct output root such as `output/skill-calling-evaluation-campaign/<evaluation-id>/`.
- Never run two evaluations in the same workspace or output location.
- Fill open process capacity with any eligible pending evaluation whose configuration gate and dependencies are satisfied.
- Inspect command output and retained artifacts directly, then serialize progress-document edits as runs finish.
- Reuse a process slot only after the command has ended and process, temporary-skill, private-state, and workspace cleanup are verified.
- If several main Codex threads participate, assign disjoint evaluation IDs and output roots to each thread and designate one thread as the only progress-document writer.

## Required Opening Lane: Claude Code With Codex

Treat Claude Code — Codex as the campaign's opening dependency lane.

1. Preflight the existing project-local Claude/Codex launcher at `.skill-issue/claudex/claudex` without reading its credentials or tokens. Confirm its version, selected model, proxy readiness, executable path, and recent bounded smoke evidence.
2. Run `CLA-COD-01`, `CLA-COD-02`, and `CLA-COD-03` sequentially in that order. Never have more than one Claude Code — Codex evaluation active.
3. Other non-Claude-Code evaluations may occupy the remaining process slots while this sequence runs.
4. Resolve and rerun any failed Claude Code — Codex evaluation before moving to the Fable transition. Do not change the Claude route while a Claude Code — Codex run is pending, running, or awaiting an identified safe rerun.
5. After all three runs complete, stop only the proxy owned by `.skill-issue/claudex/manage`. Verify that its owned process and localhost listener are gone. Do not delete or reset the isolated runtime.

## Claude Code Fable Transition Gate

The existing `.skill-issue/claudex/claudex` launcher is the Codex-backed route. It starts a localhost proxy and injects an isolated Claude config, proxy URL, authentication token, model aliases, delegated-agent model, and forced model argument. Do not treat switching away from it as a single global environment-variable change.

Before any Claude Code — Fable campaign run:

1. Identify the operator-owned normal Claude Code executable and its supported, non-proxy authentication state without inspecting credential contents.
2. Confirm that the normal route does not inherit the project-local Codex proxy URL, proxy token, isolated `CLAUDE_CONFIG_DIR`, or forced Codex model aliases.
3. Resolve the exact available Fable model identifier and medium reasoning control from the installed CLI and supported account. Record both in the progress document.
4. If normal Claude Code installation, authentication, or model selection requires user action, mark Claude Code — Fable blocked and request that action. Do not install, authenticate, or rewrite user configuration autonomously.
5. Run these smoke commands sequentially against the normal Claude Code executable in separate external workspaces and output locations:
   - the built-in gardening evaluation truncated to two turns;
   - the existing two-turn custom smoke using `evaluations/skill-calling/smoke/custom-skills/`, `custom-scenario.json`, and `custom-answer-sheet.json`.
6. Require both smoke routes to complete with artifacts, one stable session per route, expected workspace effects, temporary-skill cleanup, private-state cleanup, and no run-owned Claude process before opening the Fable campaign lane.
7. Once the gate passes, `CLA-FAB-01`, `CLA-FAB-02`, and `CLA-FAB-03` may run concurrently with each other and with eligible non-Claude-Code evaluations, subject to the six-command limit.

## OpenAI Codex Command Permission Rule

Every OpenAI Codex harness evaluation launched from Codex must request command-scoped outer-sandbox escalation for the exact known-good evaluation command.

Launch the command with:

- `sandbox_permissions: "require_escalated"`; and
- a concise justification that the nested Codex process must write its normal authenticated session database and session state under `CODEX_HOME`.

This escalation applies only to the outer main-thread shell command. Do not pass `danger-full-access`, bypass approvals, disable the evaluator's inner sandbox, or alter the inner Codex configuration. The Skill Issue evaluator must continue to own the nested Codex `workspace-write` sandbox, approval policy, ambient-config exclusions, model, reasoning, workspace, and cleanup.

If escalation is denied, mark the attempt `Blocked`, record the denial, verify any partial preparation was cleaned, and continue other eligible work. Do not retry the same command inside the outer sandbox because the resulting read-only Codex database failure is already qualified.

## Evaluation Command Launch Contract

Before launching each command, resolve and record:

- the evaluation ID, governed scenario identifier, harness, model label, effective model target, and reasoning target;
- the repository path and required authoritative files;
- the known-good CLI requirement;
- the exact harness executable or qualified launcher when it is intentionally absent from `PATH`;
- a unique external workspace and unique retained output root;
- the command-confirmation requirement after checking the CLI pre-run summary;
- the OpenAI Codex escalation rule when applicable;
- the expected artifacts and evidence to retain;
- the cleanup checks required before reusing the process slot;
- the prohibition on changing tracked files other than serialized campaign-progress updates; and
- the stop boundary for authentication, permission, model, protocol, artifact, process, or cleanup failures.

After completion, capture the exact run directory, run ID, harness version, effective model and reasoning, start and completion times, expected/observed/missing/additional/unattributed calls, artifact paths, workspace effects, session-continuity evidence, cleanup evidence, and pass/fail classification.

## Progress Update Procedure

Before starting a run:

1. Confirm its dependencies and configuration gate are satisfied.
2. Increment `Attempts`.
3. Set its status to `Running`.
4. Record the campaign start date when launching the first full run.
5. Recalculate the progress summary.

After a command finishes:

- Set `Complete` only when the full run, required artifacts, and cleanup all pass. Link the retained result and record concise configuration notes.
- Set `Failed` for a tooling failure after launch. Add a failure-log row with the attempt, exact non-secret cause, resolution or next action, and later rerun result.
- Set `Blocked` when a prerequisite prevents launch or continuation. Add a blocker-log row and continue independent work.
- Treat missing expected skill calls as evaluation data when tooling completed; do not convert them into tooling failures.
- Preserve earlier failure and blocker rows after a successful rerun.
- Check a configuration heading only when all three runs beneath it are complete.
- Recalculate configuration totals, campaign totals, pending, failed, blocked, and percentage after every status transition.
- Record the campaign completion date only at 30/30 complete.

Use `apply_patch` for progress updates and run `npm run format:check` after Markdown changes. Do not run broad website validation solely for campaign-progress edits.

## Failure Handling And Continuation

- Diagnose failures from the exact launcher, native stderr or structured events, retained run state, and production source path that owns the behavior.
- Do not use earlier passing tests or smoke reports as proof that a current failed run is valid.
- Retry only after identifying a concrete cause and a safe, scoped correction within the authority boundary.
- Never repair authentication, mutate global configuration, delete user state, or broaden permissions to keep the queue moving.
- When one lane blocks, keep all independent eligible lanes running up to the six-command limit.
- If cleanup is incomplete, stop scheduling that harness until owned-process and temporary-state boundaries are restored.

## Final Report

When the campaign reaches a terminal state, report:

- configurations and runs complete, failed, blocked, and pending;
- the retained output root and result links;
- exact effective harness/model/reasoning combinations;
- Claude Code route-transition and smoke-gate evidence;
- any unresolved user-action blockers;
- cleanup status across all harnesses; and
- whether the campaign actually reached 30/30 complete.

Do not publish, aggregate into website data, commit, push, or delete the campaign working documents unless the user separately requests it.

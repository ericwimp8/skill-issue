# Skill-Calling Evaluation Campaign Orchestration Prompt

## Goal

Execute the complete skill-calling evaluation campaign tracked in `evaluation-progress.md` as quickly as its dependency and safety boundaries allow. Run the governed evaluations through the known-good Skill Issue CLI from fresh neutral workspaces, retain the required evidence, keep the progress document current, and manage parallel evaluation commands without allowing shared-state collisions or ambient project context to enter a harness conversation.

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

- every one of the 30 matrix evaluations has completed its full governed run;
- every completed run retains its required `result.json`, `website.json`, logs, transcripts, and run evidence beneath its allocated adjacent `chats/chat-<number>/output/` directory;
- every result records the exact harness version, effective model identifier, and reasoning setting;
- instrumentation, session continuity, workspace effects, process ownership, and cleanup are tooling-complete;
- all configuration headings, run entries, summary counts, percentages, dates, notes, failures, blockers, attempts, and result links in `evaluation-progress.md` match the retained evidence; and
- no evaluation command session, temporary skill installation, private recovery state, or run-owned harness process remains active after its evidence has been retained in the run's `chat-<number>` container.

Do not report the campaign as complete while any run remains pending, running, failed, or blocked.

## Authority And Safety Boundary

You may:

- run the known-good local CLI;
- create isolated neutral run containers beneath an adjacent `chats/` directory outside the repository;
- create the neutral `chats/` parent when absent and retain every positively identified campaign-owned `chat-<number>` container for later organization;
- write evaluation artifacts beneath the allocated container's `output/` directory, outside its model-visible `workspace/` directory;
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
- move, consolidate, reorganize, empty, reuse, or delete a retained `chat-<number>` container during the campaign;
- weaken an evaluator-owned harness sandbox, approval policy, isolation control, or permission profile;
- place orchestration instructions, Git metadata, repository files, retained results, or descriptive evaluation labels inside a harness workspace;
- allow concurrent progress-document edits from separate threads or processes; or
- infer a successful evaluation from process exit alone when required artifacts or cleanup evidence are missing.

When user intervention is required for installation, authentication, account access, model availability, or normal Claude Code configuration, mark the affected run or configuration `Blocked`, record the exact non-secret blocker, continue independent work, and ask for the smallest required intervention.

## Campaign Configuration

- Use `medium` as the campaign reasoning target.
- Where a harness exposes no independent medium control, use its closest documented model-native equivalent and record the exact effective value.
- Resolve and record exact model identifiers and harness versions during configuration preflight. Do not guess model aliases.
- Run every governed scenario without `--turns`; campaign runs must include every turn in the selected embedded unit.
- Use the built-in evaluation identifiers exactly:
  - `gardening-web-application`
  - `community-archive-desktop-application`
  - `neighborhood-emergency-preparedness-program`

## Scheduling Model

Act as the campaign orchestrator, process monitor, and progress-document writer. Start and manage evaluation commands directly from the main thread.

- Keep at most six evaluation commands active simultaneously.
- Keep at most one Claude Code — Codex command active at any time, including smokes, full runs, and reruns. All other eligible harness-and-model routes may use the remaining process capacity.
- Retain the process or terminal session identifier for every active command so it can be monitored to completion.
- Launch each command for exactly one matrix evaluation unless a bounded smoke gate explicitly groups two sequential smoke routes.
- Give every smoke attempt, full run, and rerun its own newly allocated neutral `chat-<number>` container under the adjacent `chats/` root, with separate `workspace/` and `output/` children.
- Never run two evaluations in the same container, workspace, or output location.
- Fill open process capacity with any eligible pending evaluation whose configuration gate and dependencies are satisfied.
- Inspect command output and retained artifacts directly, then serialize progress-document edits as runs finish.
- Reuse a process slot only after the command has ended and process, temporary-skill, and private-state cleanup are verified. Retain the complete run container and its workspace effects.
- If several main Codex threads participate, assign disjoint evaluation IDs and output roots to each thread and designate one thread as the only progress-document writer.

### Neutral Workspace Allocation

Resolve the repository root first, then use exactly one neutral run-container parent at `<repository-parent>/chats`. Create it when absent. The `chats/` directory is outside the Skill Issue repository and must not be a Git repository or contain orchestration instructions.

Allocate run containers serially before launching commands:

1. Inspect only the immediate child directory names beneath `chats/`.
2. Choose the next unused positive integer after the highest existing `chat-<number>` directory, starting with `chat-1` when none exist.
3. Create exactly one new `<repository-parent>/chats/chat-<number>` directory for the run, containing only new empty `workspace/` and `output/` children before launch.
4. Verify `workspace/` is empty and contains no `.git`, `AGENTS.md`, `CLAUDE.md`, rules, configuration, results, or files from another run.
5. Pass the exact `<repository-parent>/chats/chat-<number>/workspace` path through the CLI's required `--workspace` argument. Do not rely on the orchestrator's current working directory to select the harness workspace.
6. Pass the exact sibling `<repository-parent>/chats/chat-<number>/output` path through the CLI's `--output` argument. The CLI writes artifacts directly there; do not copy or move results during the campaign.
7. After command completion and evaluator cleanup, retain the entire `chat-<number>` container exactly where it is, including logs, artifacts, transcripts, failure evidence, and workspace effects. Never reuse its number, even when the run failed before producing complete artifacts.

The neutral name is deliberate. Do not include evaluation IDs, scenario names, harness names, model names, `eval`, `evaluation`, `test`, `skill`, or `skill-issue` in a run-container or workspace directory name. Supply only the assigned `chat-<number>/workspace` directory as the evaluation workspace. Directory placement and neutral naming reduce ambient context but do not themselves prove filesystem confinement. The qualified runtime must separately prevent the harness from reading its sibling `output/` directory, sibling chat containers, the `chats/` parent, or the Skill Issue repository except for evaluator-owned paths required for skills, instrumentation, authentication, and retained output.

### Workspace Smoke Gate

Before starting or reopening a harness campaign lane, prove its current configuration with the smallest applicable smoke route in a newly allocated `chat-<number>` workspace. Use the existing two-turn built-in or custom smoke inputs selected by the campaign configuration; do not turn a smoke into a full campaign run.

- Confirm the pre-run summary shows the newly allocated `chat-<number>/workspace` path and its sibling `chat-<number>/output` root before accepting the command.
- Confirm from production runtime configuration and a bounded probe that the harness cannot read its sibling `output/` directory, the `chats/` parent, sibling chat directories, or the Skill Issue repository through its available file and shell tools. Do not infer confinement from the process working directory. Run this probe in its own retained `chat-<number>` container, then close its session and preserve its evidence; never reuse the probe workspace or session for a scored smoke or campaign run.
- Require turn progress, one stable resumable session, expected instrumentation, required artifacts, workspace effects, and complete evaluator cleanup.
- Inspect the retained container after cleanup and confirm its run evidence remains intact while no run-owned process, temporary skill installation, or private recovery state remains active.
- When a smoke fails, diagnose the concrete production-source or configuration owner, make only a portable correction within the authority boundary, allocate another new `chat-<number>`, and rerun the bounded smoke.
- Continue this diagnose-correct-new-workspace cycle until the route passes or reaches a genuine user-action blocker. Do not repeat an unchanged failed command, weaken isolation, or reuse its workspace.
- If the harness exposes an unconfined file or shell route and no already-qualified portable runtime control closes it, mark that lane `Blocked`. Do not run its campaign evaluations on the assumption that a neutral workspace path is sufficient isolation.

## Adjacent Workspace Command Permission Rule

The neutral `chats/` parent is deliberately adjacent to and outside the repository. When the orchestrator itself is running inside Codex, its ordinary repository `workspace-write` sandbox cannot create or populate those sibling directories.

For every run-container allocation command, governed evaluation command, smoke command, and confinement probe that must access `<repository-parent>/chats`, request command-scoped outer-sandbox escalation for that exact command. State that the Skill Issue evaluator needs access only to the newly allocated adjacent neutral container. Keep the harness's qualified inner sandbox, approval policy, controlled environment, model, reasoning, executable routing, output root, and evaluator cleanup unchanged.

Do not request a reusable destructive-command approval, a broad shell prefix, `danger-full-access` inside a harness, or permission for the harness to read the `chats/` parent, its sibling `output/`, or sibling workspaces. If outer escalation is denied, mark the attempt `Blocked`, retain any partial run evidence in its allocated container, leave unrelated containers untouched, and continue other eligible work.

## Required Opening Lane: Claude Code With Codex

Treat Claude Code — Codex as the campaign's opening dependency lane.

1. Preflight the existing project-local Claude/Codex launcher at `.skill-issue/claudex/claudex` without reading its credentials or tokens. Confirm its version, selected model, proxy readiness, executable path, and recent bounded smoke evidence.
2. Run `CLA-COD-01`, `CLA-COD-02`, and `CLA-COD-03` sequentially in that order. Never have more than one Claude Code — Codex command active, including preflights, smokes, full runs, and reruns.
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
5. Run these smoke commands sequentially against the normal Claude Code executable in separate newly allocated `chat-<number>` workspaces and output locations:
   - the built-in gardening evaluation truncated to two turns;
   - the existing two-turn custom smoke using `evaluations/skill-calling/smoke/custom-skills/`, `custom-scenario.json`, and `custom-answer-sheet.json`.
6. Require both smoke routes to complete with artifacts, one stable session per route, expected workspace effects, temporary-skill cleanup, private-state cleanup, and no run-owned Claude process before opening the Fable campaign lane.
7. Once the gate passes, `CLA-FAB-01`, `CLA-FAB-02`, and `CLA-FAB-03` may run concurrently with each other and with eligible non-Claude-Code evaluations, subject to the six-command limit.

## OpenAI Codex Command Permission Rule

Every OpenAI Codex harness evaluation launched from Codex must request command-scoped outer-sandbox escalation for the exact known-good evaluation command. This is the same exact-command escalation required for the adjacent workspace, with the additional Codex authentication requirement below; do not launch the evaluation twice.

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
- the next newly allocated `<repository-parent>/chats/chat-<number>` container, its empty model-visible `workspace/`, and its sibling retained `output/` root;
- the command-confirmation requirement after checking the CLI pre-run summary;
- the OpenAI Codex escalation rule when applicable;
- the adjacent-workspace exact-command escalation rule when the orchestrator runs inside Codex;
- the expected artifacts and evidence to retain;
- the active-process, temporary-skill, and private-state cleanup checks required before reusing the process slot, without altering the retained run container;
- the prohibition on changing tracked files other than serialized campaign-progress updates; and
- the stop boundary for authentication, permission, model, protocol, artifact, process, or cleanup failures.

After completion, capture the exact run container, run ID, harness version, effective model and reasoning, start and completion times, expected/observed/missing/additional/unattributed calls, artifact paths, workspace effects, session-continuity evidence, process and temporary-state cleanup evidence, and pass/fail classification. Do not record the neutral workspace path in model-visible prompts or generated skills.

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
- If process, temporary-skill, or private-state cleanup is incomplete, stop scheduling that harness until those boundaries are restored. Preserve the allocated run container while diagnosing the failure.
- Retain every numbered run container regardless of success, failure, or blocker. Never clear it, reuse it, or move its contents during the campaign.

## Final Report

When the campaign reaches a terminal state, report:

- configurations and runs complete, failed, blocked, and pending;
- the retained `chats/` root, per-run container paths, and result links;
- exact effective harness/model/reasoning combinations;
- Claude Code route-transition and smoke-gate evidence;
- any unresolved user-action blockers;
- cleanup status across all harnesses;
- neutral run-container allocation, workspace-confinement proof, retained-evidence status, and process and temporary-state cleanup status; and
- whether the campaign actually reached 30/30 complete.

Do not publish, aggregate into website data, commit, push, or delete the campaign working documents unless the user separately requests it.

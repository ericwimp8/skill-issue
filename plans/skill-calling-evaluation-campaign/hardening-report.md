# Harness Evaluation Path Hardening Report

## Executive Summary

The six supported harnesses completed the required four-turn
`gardening-web-application` evaluation tooling-clean on the rebuilt development
CLI channel:

- OpenAI Codex
- Claude Code
- Cursor Agent
- OpenCode
- Kilo Code
- Pi

The hardening exercise required ten attempts. Four attempts ended in tooling
errors and were diagnosed from their retained `failure.json` artifacts. One
failure was a production CLI defect and three were invocation or environment
requirements. After the production fix and invocation corrections, all six
harnesses completed four turns with one stable resumable session and produced
the required `result.json`, `website.json`, and `events.jsonl` artifacts.

The production defect was in Codex runtime configuration. Codex `0.144.6`
rejects the former Boolean `agents.enabled=false` setting because `agents` now
holds agent-role configuration. The development CLI now disables the stable
multi-agent path with `features.multi_agent=false` while retaining
`features.multi_agent_v2=false` for the v2 path.

No tooling failures remain in the six-harness matrix on the current rebuilt
development CLI.

## Scope And Method

The run followed
`plans/skill-calling-evaluation-campaign/hardening-orchestration-prompt.md` and
its inherited campaign conventions.

- The development CLI was built before evaluation with
  `./cli/scripts/local-cli.sh build-development`.
- Every evaluation used `./cli/scripts/local-cli.sh development`.
- Every attempt used a newly created external temporary Git workspace.
- Every attempt used a distinct ignored output root beneath
  `output/harness-hardening/`.
- Runs were executed sequentially.
- Every evaluation used the first four turns of the built-in
  `gardening-web-application` scenario.
- Reasoning was left at the harness default of `medium`; no `--reasoning` flag
  was passed.
- Models remained at harness defaults except Cursor, where the installed
  medium Grok identifier was selected after the requested `grok` alias was
  rejected.
- Every tooling failure was classified from its own `failure.json`, native
  output, run artifacts, and the production source path that owned the
  behavior.
- Missing and additional skill calls from completed runs were treated as model
  behavior rather than tooling failures.

## Qualified Configuration

| Harness       | Version              | Effective Model            | Reasoning | Required Route Or Environment                                                                     |
| ------------- | -------------------- | -------------------------- | --------- | ------------------------------------------------------------------------------------------------- |
| `codex`       | `0.144.6`            | `gpt-5.6-sol`              | `medium`  | Operator `codex` executable; command-scoped outer escalation for nested authenticated Codex state |
| `claude-code` | `2.1.205`            | `opus`                     | `medium`  | Operator `claude` executable; normal Claude session state writable across turns                   |
| `cursor`      | `2026.07.16-899851b` | `cursor-grok-4.5-medium`   | `medium`  | Absolute route to `.skill-issue/cursor/home/.local/bin/agent`                                     |
| `opencode`    | `1.18.4`             | `openai/gpt-5.6-sol`       | `medium`  | Absolute route to `.skill-issue/opencode/bin/opencode`; qualified `XDG_DATA_HOME`                 |
| `kilo-code`   | `7.4.11`             | `openai/gpt-5.6-sol`       | `medium`  | Absolute route to the real versioned Kilo binary; qualified `XDG_DATA_HOME`                       |
| `pi`          | `0.80.10`            | `openai-codex/gpt-5.6-sol` | `medium`  | Operator Pi entrypoint; authentication through `PI_CODING_AGENT_DIR`                              |

Exact machine-local executable paths are retained in the ignored file
`output/harness-hardening/preflight.txt`. They are intentionally absent from
this tracked report.

## Final Results

| Harness       | Passing Attempt | Run ID                             | Expected | Observed | Missing | Additional | Unattributed | Result        |
| ------------- | --------------: | ---------------------------------- | -------: | -------: | ------: | ---------: | -----------: | ------------- |
| `codex`       |               2 | `1d333bcc7a51a1b55b6fc79c12cacaa1` |        5 |        8 |       0 |          3 |            0 | Tooling-clean |
| `claude-code` |               2 | `ce150814593556eff53d5ebb011c5dc0` |        5 |        1 |       4 |          0 |            0 | Tooling-clean |
| `cursor`      |               3 | `9d9477b2ab826fed9342a0b0600782f7` |        5 |        5 |       0 |          0 |            0 | Tooling-clean |
| `opencode`    |               1 | `a5d32f21f0765b9f01cc64f4bab1b999` |        5 |        2 |       3 |          0 |            0 | Tooling-clean |
| `kilo-code`   |               1 | `d26ed7f9656a184355654acec7b7b45b` |        5 |        3 |       3 |          1 |            0 | Tooling-clean |
| `pi`          |               1 | `1ec48bd17008b9ce6eb0a5e2d3991d0f` |        5 |        3 |       3 |          1 |            0 | Tooling-clean |

All six passing runs completed turns 1 through 4, retained one stable harness
session, wrote the required three artifacts, removed temporary evaluation
skills and private evaluator state, removed disposable native sessions where
the adapter owns deletion, and left no run-owned process active. Model-created
workspace files remain as evidence.

## Failure Analysis And Corrective Actions

### 1. Codex Configuration Rejected

**Attempt:** `codex-1`  
**Run ID:** `dfb4ee7f8ac273b66a6b55b087d94fb2`  
**Failure point:** Turn 1, before model execution  
**Classification:** Harness configuration rejected; production CLI defect

The native error was:

```text
Error loading config.toml: invalid type: boolean `false`, expected struct AgentRoleToml
in `agents`
```

The concrete production path was:

1. `cli/internal/evaluation/runtime.go` generated
   `agents.enabled=false`.
2. `cli/internal/evaluation/evaluation.go` passed the generated configuration
   into the replay adapter.
3. `cli/internal/replay/process.go` relayed each value to Codex with
   `--config`.
4. Codex `0.144.6` rejected the Boolean before turn 1 started.

The installed Codex binary was probed directly. It rejected
`agents.enabled=false`, accepted `features.multi_agent=false`, and listed
`multi_agent` as the stable feature key. The production fix therefore replaced
the invalid setting at its owner rather than compensating at the command
wrapper or invocation layer.

Changes made:

- `cli/internal/evaluation/runtime.go`
  - replaced `agents.enabled=false` with `features.multi_agent=false`;
  - retained `features.multi_agent_v2=false`.
- `cli/internal/evaluation/runtime_test.go`
  - updated the owner-level assertion to require both supported feature keys.
- `cli/README.md`
  - updated the Codex runtime isolation contract to describe the current
    generated configuration.

Validation before rerun:

- `gofmt` passed.
- `go vet ./cli/...` passed.
- `go test ./cli/...` passed.
- The development CLI was rebuilt.

`codex-2` then completed all four turns tooling-clean on the rebuilt binary.
It made three additional `dictate-plan` calls on turns 2 through 4. Those calls
are model behavior and did not affect the tooling result.

Evidence:

- Failure:
  `output/harness-hardening/codex-1/codex-20260721T070546Z-dfb4ee7f/failure.json`
- Passing result:
  `output/harness-hardening/codex-2/codex-20260721T070812Z-1d333bcc/result.json`

### 2. Claude Code Session Not Available For Resume

**Attempt:** `claude-code-1`  
**Run ID:** `189fe5295dbd1cf8ec47fe68c33d0c3a`  
**Failure point:** Turn 2 after turn 1 completed  
**Classification:** Route/environment

Turn 1 completed and returned a session ID, but turn 2 failed with:

```text
No conversation found with session ID: 08cd71b4-3516-0fa9-8c18-a3ed194a31f8
```

The evaluator used the normal Claude executable and the correct `--resume`
value. The failed outer command was sandboxed and did not preserve Claude's
normal operator-owned session state for the resumed turn.

Action taken:

- The production CLI remained unchanged.
- A fresh workspace and output root were allocated.
- The same evaluation command was rerun with command-scoped outer access to
  normal Claude session state.
- The evaluator-owned settings, permissions, workspace boundary, model,
  reasoning, and cleanup behavior remained unchanged.

`claude-code-2` resumed its run-owned session across all four turns and
completed tooling-clean. It omitted four expected
`document-update-discipline` calls. Those omissions are model behavior.

Evidence:

- Failure:
  `output/harness-hardening/claude-code-1/claude-code-20260721T071340Z-189fe529/failure.json`
- Passing result:
  `output/harness-hardening/claude-code-2/claude-code-20260721T071621Z-ce150814/result.json`

### 3. Cursor Executable Route Became Invalid

**Attempt:** `cursor-1`  
**Run ID:** `f4def1a99cbfc767424881de15c17d24`  
**Failure point:** Turn 1, before process start  
**Classification:** Route/environment

The first Cursor invocation passed the project-local executable as a relative
path. The evaluator launched Cursor from its run-owned working directory,
where that relative path did not exist:

```text
fork/exec .skill-issue/cursor/home/.local/bin/agent: no such file or directory
```

Action taken:

- The production CLI remained unchanged.
- The project-local Cursor agent was resolved to an absolute path before the
  evaluation command was launched.
- A fresh workspace and output root were allocated.

The corrected route reached the native Cursor process. Cleanup from the failed
attempt removed temporary skills, private runtime state, and run-owned
processes.

Evidence:

- Failure:
  `output/harness-hardening/cursor-1/cursor-20260721T071944Z-f4def1a9/failure.json`

### 4. Cursor Rejected The Requested Grok Alias

**Attempt:** `cursor-2`  
**Run ID:** `98ee7ba2c48796111c1499d4132062ce`  
**Failure point:** Turn 1, before model execution  
**Classification:** Route/environment; native model identifier mismatch

The required first model attempt used `--model grok`. Cursor rejected that
alias and returned its available model list. The installed medium Grok
identifier was:

```text
cursor-grok-4.5-medium
```

Action taken:

- The production CLI remained unchanged because native model identifiers are
  intentionally owned by the harness.
- A fresh workspace and output root were allocated.
- `cursor-3` used `--model cursor-grok-4.5-medium` with no reasoning override.

`cursor-3` completed all four turns tooling-clean and observed all five
expected calls with no missing, additional, or unattributed calls.

Evidence:

- Failure:
  `output/harness-hardening/cursor-2/cursor-20260721T072052Z-98ee7ba2/failure.json`
- Passing result:
  `output/harness-hardening/cursor-3/cursor-20260721T072158Z-9d9477b2/result.json`

## Harness Result Details

| Harness       | Started                       | Completed                     | Model Behavior                                                  |
| ------------- | ----------------------------- | ----------------------------- | --------------------------------------------------------------- |
| `codex`       | `2026-07-21T07:08:12.297033Z` | `2026-07-21T07:12:48.943948Z` | Three additional `dictate-plan` calls                           |
| `claude-code` | `2026-07-21T07:16:21.165204Z` | `2026-07-21T07:18:52.955483Z` | Four missing `document-update-discipline` calls                 |
| `cursor`      | `2026-07-21T07:21:58.441721Z` | `2026-07-21T07:26:42.463313Z` | All expected calls observed                                     |
| `opencode`    | `2026-07-21T07:27:30.229156Z` | `2026-07-21T07:31:08.518695Z` | Three missing `document-update-discipline` calls                |
| `kilo-code`   | `2026-07-21T07:31:57.526783Z` | `2026-07-21T07:35:22.900235Z` | Three missing calls; additional `prompt-writing` call on turn 1 |
| `pi`          | `2026-07-21T07:36:08.635003Z` | `2026-07-21T07:39:24.357427Z` | Three missing calls; additional `prompt-writing` call on turn 1 |

## Source Changes

| File                                      | Change                                                                              |
| ----------------------------------------- | ----------------------------------------------------------------------------------- |
| `cli/internal/evaluation/runtime.go`      | Replaced invalid `agents.enabled=false` with supported `features.multi_agent=false` |
| `cli/internal/evaluation/runtime_test.go` | Updated the Codex runtime ownership assertion for both multi-agent feature paths    |
| `cli/README.md`                           | Updated the documented Codex runtime isolation configuration                        |

The attempt-by-attempt ledger remains at
`plans/skill-calling-evaluation-campaign/hardening-progress.md`.

## Validation And Cleanup

The following validation passed after the production fix:

```text
gofmt
go vet ./cli/...
go test ./cli/...
npm run format:check
npx prettier --check cli/README.md plans/skill-calling-evaluation-campaign/hardening-progress.md
git diff --check
```

The development CLI was rebuilt after the source change and before the post-fix
matrix. All six final passing runs exercised the same fixed development build.

Final artifact checks confirmed for every passing run:

- `result.json` exists;
- `website.json` exists and reports four turns;
- `events.jsonl` exists;
- no `failure.json` exists in the passing run directory; and
- the attempt log contains four completed turn records.

Final cleanup checks found no output-owned `.skill-issue` private state, no
run-owned evaluator or harness process, no temporary evaluation skill files,
no run-owned private runtime, and no surviving evaluator-owned OpenCode or Kilo
session.

The evaluated workspaces retain model-created project files, principally
`plans/`, as evidence of model actions. The empty Codex `.agents/skills/`
directory is cosmetic residue and contains no temporary skills.

## Retained Output Locations

All output is ignored and remains beneath:

```text
output/harness-hardening/
```

### Shared Preflight Evidence

| Evidence                                                       | Location                                            |
| -------------------------------------------------------------- | --------------------------------------------------- |
| Exact executable paths, versions, and Cursor model identifiers | `output/harness-hardening/preflight.txt`            |
| Codex configuration parse probe home                           | `output/harness-hardening/codex-config-probe-home/` |

### Attempt Evidence

| Attempt         | Outcome         | Output Root                               | Primary Evidence                                     |
| --------------- | --------------- | ----------------------------------------- | ---------------------------------------------------- |
| `codex-1`       | Tooling failure | `output/harness-hardening/codex-1/`       | `codex-20260721T070546Z-dfb4ee7f/failure.json`       |
| `codex-2`       | Tooling-clean   | `output/harness-hardening/codex-2/`       | `codex-20260721T070812Z-1d333bcc/result.json`        |
| `claude-code-1` | Tooling failure | `output/harness-hardening/claude-code-1/` | `claude-code-20260721T071340Z-189fe529/failure.json` |
| `claude-code-2` | Tooling-clean   | `output/harness-hardening/claude-code-2/` | `claude-code-20260721T071621Z-ce150814/result.json`  |
| `cursor-1`      | Tooling failure | `output/harness-hardening/cursor-1/`      | `cursor-20260721T071944Z-f4def1a9/failure.json`      |
| `cursor-2`      | Tooling failure | `output/harness-hardening/cursor-2/`      | `cursor-20260721T072052Z-98ee7ba2/failure.json`      |
| `cursor-3`      | Tooling-clean   | `output/harness-hardening/cursor-3/`      | `cursor-20260721T072158Z-9d9477b2/result.json`       |
| `opencode-1`    | Tooling-clean   | `output/harness-hardening/opencode-1/`    | `opencode-20260721T072730Z-a5d32f21/result.json`     |
| `kilo-code-1`   | Tooling-clean   | `output/harness-hardening/kilo-code-1/`   | `kilo-code-20260721T073157Z-d26ed7f9/result.json`    |
| `pi-1`          | Tooling-clean   | `output/harness-hardening/pi-1/`          | `pi-20260721T073608Z-1ec48bd1/result.json`           |

Each attempt root also contains `attempt.log`, including the pre-run summary
and turn progress, plus `workspace-path.txt`, containing the exact temporary
workspace path. Passing run directories contain `result.json`, `website.json`,
and `events.jsonl`; failed run directories contain `failure.json`.

Failure artifacts are unsanitized and can contain local paths and prompt text.
They must remain ignored and be reviewed before sharing.

## Remaining Requirements And Follow-Up

No unresolved tooling defect remains for this six-harness hardening matrix.
The following environment and invocation requirements remain necessary:

1. Nested Codex evaluations launched by Codex require command-scoped outer
   escalation for normal authenticated session database and state access.
2. Normal Claude Code needs writable operator-owned session state across
   turns.
3. Project-local executable routes must be resolved before the evaluator
   changes to a run-owned working directory.
4. Cursor Agent `2026.07.16-899851b` requires
   `cursor-grok-4.5-medium`; the generic `grok` alias is unavailable.
5. OpenCode requires its qualified `XDG_DATA_HOME` for native authentication.
6. Kilo requires its qualified `XDG_DATA_HOME` and the real versioned binary,
   rather than `.skill-issue/kilo/bin/kilo`, whose wrapper pins a conflicting
   configuration home.
7. Pi authentication continues to resolve through `PI_CODING_AGENT_DIR`.
8. OpenCode and Kilo must remain on their qualified versions unless an
   unqualified run is deliberately authorized and recorded.

The production fix has not been committed or promoted to the known-good CLI
channel as part of this task. Retain this report, the hardening prompt, and the
progress ledger until the source change is reviewed, committed, and
deliberately promoted to the next known-good baseline. Keep the ignored
`output/harness-hardening/` tree available until that review is complete, then
remove it only through deliberate cleanup.

## Conclusion

The harness evaluation path is tooling-clean across Codex, Claude Code,
Cursor, OpenCode, Kilo Code, and Pi on the current rebuilt development CLI.
The run found and fixed one production compatibility defect, confirmed three
invocation-level corrections, preserved every failure artifact, and verified
artifacts, session continuity, process ownership, temporary-skill cleanup, and
private-state cleanup for every final pass.

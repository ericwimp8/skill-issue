# Active Plans and Trackers

## Assignment

**Goal:** Determine which documents under `plans/` currently act as live progress or task trackers, reconstruct their claimed phase and dependency order, identify stale or conflicting state, and select the best semantic owner and routing/index candidates for consolidation.

**Scope:** Every current Markdown file under `plans/**`; direct repository sources and retained evidence needed to test their status claims; read-only Git history and blame for the relevant documents.

**Exclusions:** Broad production-source tracing, broad `research/**` cleanup, evaluation-artifact deep dives, Codex session-history reconstruction, and any edit outside this assignment document.

## Sources

- `AGENTS.md:1-133` — branch safety, current website and CLI ownership, local evaluation routing, planning-document lifecycle, and campaign-specific output rules.
- `.repository-privacy.md:1-14` — repository-relative citation and identity constraints.
- `plans/harness-setup.md:1-692` — retained harness setup contract and smoke-qualification record.
- `plans/skill-calling-evaluation-campaign/evaluation-progress.md:1-160` — campaign status ledger, dependency rules, per-run records, blocker log, and campaign notes.
- `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md:1-84` — campaign operating, scheduling, retry, and reporting contract.
- `plans/skill-calling-evaluation-campaign/campaign-run-sheet.md:1-101` — fixed command sheet and original container map.
- `plans/website/reference-and-architecture.md:1-90` — website research, architecture decision, stack, and ownership record.
- `cli/internal/harness/harness.go:14-20,62-70` — current supported-harness registry.
- `cli/internal/evaluation/runtime.go:32-80` — current concrete runtime preparation and dispatch.
- `cli/README.md:21-39,144-180` — current CLI support, instrumentation, Codex isolation, and OpenCode runtime contracts.
- `src/App.tsx:11-49,82-95` — current four-destination website routing.
- `src/data/evaluationData.ts:67-146,203-227` — current nine campaign cells, seven available cells, and published-result adaptation.
- `src/data/siteData.ts:47-99` — current website status, release, navigation, and product-arm copy owner.
- `evaluations/skill-calling/results/accepted/*.json` — 21 retained accepted result artifacts; inspected only as an inventory and harness/model/scenario matrix.
- `src/data/publishedWebsiteArtifacts.json` — 21 published website artifacts; inspected only for array size and identity fields.
- `evaluations/skill-calling/smoke/real-harness-smoke-report.md` — direct retained evidence target named by `plans/harness-setup.md`.
- `research/deep-research/free-website-hosting-cli-downloads/free-website-hosting-cli-downloads-deep-research.md` — direct local research target named by the website decision.
- Git commit `5040bfa` (`feat: harden evaluation privacy and retire Kilo support`) — removed Kilo from production and most campaign documents but introduced the unreconciled 30-command run sheet.
- Git commit `237d403` (`docs: note retained research files`) — last update to the live campaign ledger; completed the current 21/27 summary while leaving earlier notes and rerun-result cells stale.
- Git commit `6bf4312` (`feat: refine website downloads and retire completed docs`) — removed several completed progress documents while retaining the current five `plans/**` files.
- Read-only checks: `git status`, `git log`, `git blame`, `git show`, `rg`, `find`, and `jq`. No plan, source, Git, runtime, or external campaign state was changed.

## Findings

### Finding 1 — One Current File Owns Live Status

`plans/skill-calling-evaluation-campaign/evaluation-progress.md` is the only current document whose stated purpose is live execution bookkeeping. Its purpose is explicitly “execution progress only,” and it contains current aggregate counts, per-run statuses, blockers, and next actions. The other four plan files are contracts, command reference, or design history rather than live project trackers.

Current classification:

| Document                                                                   | Classification                                                      | Current role                                            |
| -------------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------- |
| `plans/skill-calling-evaluation-campaign/evaluation-progress.md`           | **live-current**, with stale internal duplicates                    | Semantic owner of campaign run status                   |
| `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md` | **durable domain plan**, with stale count language                  | Campaign execution and routing contract                 |
| `plans/skill-calling-evaluation-campaign/campaign-run-sheet.md`            | **stale/conflicting; deletion candidate**                           | Duplicated fixed commands, including retired support    |
| `plans/harness-setup.md`                                                   | **stale/conflicting; deletion candidate after semantic extraction** | Former harness contract plus historical smoke narrative |
| `plans/website/reference-and-architecture.md`                              | **historical-preserve**                                             | Original design provenance, not current website state   |

**Evidence:** `evaluation-progress.md:3-7` declares execution-progress ownership. `campaign-orchestration-prompt.md:3-7` declares scheduling, launching, monitoring, bookkeeping, and reporting behavior, while assigning tracker writes to the orchestrator. `campaign-run-sheet.md:1-5` describes itself as an operator command sheet. `harness-setup.md:3` calls itself a setup-contract owner. `reference-and-architecture.md:7-9,24-26` records completed visual research and a selected direction.

**Implication:** Consolidation should preserve one live status owner rather than creating a new parallel project-status ledger. Repairs belong first in `evaluation-progress.md`; routing and procedure remain separate from status.

### Finding 2 — The Campaign Has Completed 21 of 27 Runs

The reliable campaign state is seven complete configuration suites and 21 complete evaluation runs out of nine suites and 27 runs. Six runs are blocked; none are running or pending. The accepted evidence inventory independently contains 21 result JSON files across exactly seven harness/model cells with three scenarios each, and the website publishes 21 artifacts while marking the two blocked Cursor cells unavailable.

**Evidence:** `evaluation-progress.md:11-22` lists nine configurations and 21/27 complete. `evaluation-progress.md:26-36` reports 7/9 suites, 21/27 complete, 0 running, 0 pending, and 6 blocked. `src/data/evaluationData.ts:82-146` defines nine cells, with Cursor/Fable and Cursor/Codex unavailable and the other seven available. `src/data/evaluationData.ts:203-211` filters published artifacts to available cells. Read-only inventory found 21 files in `evaluations/skill-calling/results/accepted/` and 21 entries in `src/data/publishedWebsiteArtifacts.json`.

**Implication:** The campaign is active-but-blocked rather than completed or abandoned. The 21/27 headline is the best-supported current state even though lower sections of the tracker are stale.

### Finding 3 — Completed, Active, Next, and Later Phases Are Recoverable

The documents imply this phase map:

1. **Completed:** Claude Code/Codex, OpenAI Codex/Sol, Claude Code/Fable, Cursor/Grok, Cursor/Composer, Pi/Codex, and OpenCode/Codex each completed all three scenarios.
2. **Active:** No evaluation process or lane is active. The remaining campaign phase is blocked waiting on Cursor subscription capacity.
3. **Next:** When the Cursor capacity blocker clears, re-resolve the live Cursor model identifiers, recreate each failed container under its recorded number, and rerun the six Cursor/Codex and Cursor/Fable scenarios.
4. **Later:** Reconcile the tracker after each terminal run; after all remaining runs complete, produce the operator-owned final campaign report and decide publication, commit, and cleanup separately.

**Evidence:** Completed lanes and runs are recorded at `evaluation-progress.md:64-127`. No active lane and no remaining process are claimed at `evaluation-progress.md:154-156`. The six Cursor blockers and recorded containers are at `evaluation-progress.md:85-99,143-148`. Fresh model resolution is required by `campaign-orchestration-prompt.md:43`. Same-container recreation and retry are required by `campaign-orchestration-prompt.md:47-49,66-76`. Final reporting and publication ownership are stated at `campaign-orchestration-prompt.md:78-84`.

**Implication:** Recovery does not need to infer a new implementation phase. The only claimed live execution work is the bounded six-run Cursor continuation followed by final reporting.

### Finding 4 — The Remaining Work Has a Concrete Dependency Order

The current dependency chain is:

1. Operator-owned Cursor capacity becomes available, either at the recorded 2026-08-21 cycle reset or through an earlier operator billing decision.
2. Re-check Cursor's live model catalogue because its identifiers drift.
3. For each blocked evaluation, inspect the retained failed attempt only as needed, delete and recreate its complete `chat-<n>` container, then rerun under that same stable container identity.
4. Keep scenarios sequential inside each configuration lane; choose Cursor-lane concurrency adaptively if the shared account is healthy.
5. Record each terminal result in the status ledger and retain accepted evidence only after a tooling-complete run.
6. Close the campaign with the operator's final report and separate publication/cleanup decision.

The recorded retry containers are `chat-12`, `chat-13`, `chat-14`, `chat-15`, `chat-16`, and `chat-36`; `chat-17` is explicitly excluded because unrelated residue was preserved.

**Evidence:** The blocker and operator wait decision are at `evaluation-progress.md:90-99,143-148`. The `chat-36` substitution is at `evaluation-progress.md:92`. Model re-resolution and lane ordering are at `campaign-orchestration-prompt.md:29-43,53-58`. Container recreation is at `campaign-orchestration-prompt.md:47-49,66-76`. Acceptance rules are at `evaluation-progress.md:38-49`.

**Implication:** A consolidated tracker should expose this dependency chain once. Fixed commands and duplicated scheduling rules should not own it.

### Finding 5 — The Live Tracker Contains Multiple Stale Status Owners

The tracker’s headline tables are current, but several lower status manifestations were not reconciled when commit `237d403` advanced the campaign to 21/27:

- Cursor/Composer has all three runs checked complete at `evaluation-progress.md:111-113`, but its configuration heading remains unchecked and says 0/3 at line 110.
- CUR-COM-01 attempt 1 still says its rerun result is pending at line 137, despite the accepted attempt 2 at line 111.
- CUR-GRO-01 attempt 2 still says attempt 3 is running at line 139, despite the complete attempt 3 at line 104.
- CLA-COD-03 attempt 1 still says attempt 2 is running at line 141, despite the complete attempt 2 at line 69.
- Campaign notes say only four runs completed at line 160, contradicting 21/27 at lines 30-36.
- Cleanup notes say incomplete `chat-6` and `chat-7` must be recreated at line 159, while their accepted completed runs are recorded at lines 104 and 111.
- The cleanup note claims `chat-1`, `chat-2`, `chat-3`, and `chat-5` are retained at line 159, while accepted-run notes say these attempt-1 containers were removed at lines 67, 74, 118, and 125.
- “Exact model identifiers” still describes Composer and Grok as stopped attempts at line 153 although those cells are complete at lines 103-113.

**Evidence:** `git show 237d403 -- plans/skill-calling-evaluation-campaign/evaluation-progress.md` shows the aggregate and Claude completion update but leaves the older failure-log results and campaign notes unchanged. `git blame` attributes the stale lower entries mostly to the earlier `5040bfa` snapshot while the aggregate and accepted-run state comes from `1404e8e` and `237d403`.

**Implication:** Status consolidation must update every manifestation in the current owner, not merely its summary. The failure log should preserve historical causes while its “rerun result” cells state the actual terminal outcomes.

### Finding 6 — Campaign Scope Still Conflicts Between 27 and 30 Runs

The current CLI and tracker support a 27-run campaign. Kilo was retired in commit `5040bfa`, and production now registers only Claude Code, Codex, Cursor, OpenCode, and Pi. The orchestration prompt was only partially reconciled: its opening still says 30 runs and ten configurations, its heading and table say 27 runs and nine configurations, and its example/final-report language still says 30. The command sheet retains the removed Kilo lane and all 30 original commands.

**Evidence:** `campaign-orchestration-prompt.md:5` says 30/ten, while `campaign-orchestration-prompt.md:27-41` says 27 and lists nine configurations; lines 80 and 84 return to 30. `campaign-run-sheet.md:1-3,73-79,99-101` preserves 30 commands and Kilo containers. `evaluation-progress.md:5,11-22` consistently owns 27 runs. `cli/internal/harness/harness.go:14-20,62-70` contains no Kilo ID or spec. Commit `5040bfa` removes Kilo from production and the tracker but adds the unreconciled run sheet.

**Implication:** Treat 27 as current. Repair the orchestration prompt’s stale 30 references. Delete the run sheet or reduce it to a redirect; executing its Kilo commands would fail against current production.

### Finding 7 — Scheduling Rules Are Duplicated and Contradictory

The tracker says ten simultaneous evaluations are the campaign-wide ceiling, while the orchestration prompt and run sheet say there is no fixed concurrency cap. The tracker itself says the adjacent orchestration prompt owns scheduling and launch contracts, so its ceiling is a second, conflicting owner. Historical run notes also record that two Claude/Codex runs overlapped after an operator relaxation, while the retained normative documents still require at most one active `claude-code` run.

**Evidence:** `evaluation-progress.md:51-62` includes the ten-run ceiling, Claude serialization, and a direct deferral of scheduling ownership to the orchestration prompt. `campaign-orchestration-prompt.md:51-58` says no concurrency cap and at most one Claude Code run. `campaign-run-sheet.md:3-5` also says no fixed cap. `evaluation-progress.md:67` records a historical parallel Claude/Codex exception.

**Implication:** Keep current and future scheduling rules only in `campaign-orchestration-prompt.md`. The status ledger should record actual exceptions as evidence without restating the operating policy.

### Finding 8 — The Fixed Run Sheet Is Unsafe as a Live Command Source

Beyond retaining Kilo, the run sheet fixes model identifiers that its own introduction says drift, maps CUR-FAB-03 to stale `chat-17` instead of current `chat-36`, and omits the explicit `--reasoning medium` required by the current Claude Code routing contract. It duplicates commands already derivable from the orchestrator contract and live tracker.

**Evidence:** Model drift is acknowledged at `campaign-run-sheet.md:3`. Kilo commands are at lines 73-79. The stale container map is at lines 99-101, while `evaluation-progress.md:92` owns `chat-36`. Claude commands at `campaign-run-sheet.md:81-96` omit `--reasoning medium`; `AGENTS.md:70-81` requires explicit reasoning for both normal Claude and the Codex-backed proxy, and `campaign-orchestration-prompt.md:57` repeats that requirement.

**Implication:** `campaign-run-sheet.md` is the strongest deletion candidate. If a navigational trace is useful, replace its body with a short redirect to the orchestration prompt and status ledger rather than maintaining another command owner.

### Finding 9 — Harness Setup Is Neither a Live Tracker Nor a Current Contract

`plans/harness-setup.md` has no remaining task list or progress state. It mixes a former normative contract with a dated smoke record, and its Codex section materially conflicts with current production:

- It requires the user’s normal `CODEX_HOME` and forbids copying authentication.
- Current production creates a run-owned `CODEX_HOME` and copies `auth.json` into it when present.
- It describes Codex attribution through `skill-issue signal` and a private state path.
- Current production uses a capture-only `echo` marker for Codex.
- It says a complete governed campaign remains separate work, while the repository now retains 21 complete governed runs.

**Evidence:** Former Codex authentication and attribution rules are at `harness-setup.md:31,142-153,169`. Its dated qualification boundary is at `harness-setup.md:178,686-692`. Current runtime creation and credential copy are concrete at `cli/internal/evaluation/runtime.go:32-59` and documented at `cli/README.md:164-170`. Current Codex capture-only instrumentation is documented at `cli/README.md:144-168`. The accepted campaign state is at `evaluation-progress.md:26-36`.

**Implication:** Current behavior belongs to production source and `cli/README.md`; historical smoke evidence belongs to `evaluations/skill-calling/smoke/real-harness-smoke-report.md`. After checking for any still-unique durable rationale, retire `plans/harness-setup.md` rather than trying to restore it as a tracker.

### Finding 10 — The Website Decision Is Historical Design Provenance

The website document still accurately records its original visual research, React/TypeScript/Vite/Recharts choice, local-data architecture, and source ownership. Its concrete information architecture is no longer current: it specifies a one-page site and graph cards for Codex and Claude Code, while the implementation now routes among Explore, Method, Analysis, and Project and exposes seven completed comparison cells.

**Evidence:** The original one-page, two-environment shape is at `reference-and-architecture.md:43-50`; current four-destination routing is at `src/App.tsx:11-49,82-95`. The original stack remains at `reference-and-architecture.md:52-74` and `package.json:1-37`, while current result availability is at `src/data/evaluationData.ts:82-146`. Current copy and release ownership still matches `reference-and-architecture.md:76-82` and `src/data/siteData.ts:47-99`.

**Implication:** Preserve this only as a historical architecture/design decision, preferably outside the live-plan surface or explicitly labeled as implemented history. It should not receive progress, next-action, or current-state fields.

### Finding 11 — Best Semantic Owner and Routing Candidates Are Distinct

The best-fit semantic owner for live campaign state is `plans/skill-calling-evaluation-campaign/evaluation-progress.md`, after its internal reconciliation. The best-fit campaign-level routing document is `campaign-orchestration-prompt.md`, after correcting 27/30 language and keeping scheduling there exclusively. No current file serves as a repository-wide `plans/` index.

For repo-wide consolidation, a new compact `plans/README.md` is the clean routing candidate. It should link to:

- the one live campaign status owner;
- the campaign operating contract;
- any deliberately retained historical decision record;
- permanent CLI, website, evaluation, and research owners outside `plans/`.

It should not copy live counts, blockers, next actions, commands, or phase state.

**Evidence:** Status ownership is explicit at `evaluation-progress.md:3-7`. Scheduling and operator routing ownership is explicit at `evaluation-progress.md:62` and `campaign-orchestration-prompt.md:3-7`. The current `plans/` inventory contains only the five analyzed files and no index. `AGENTS.md:101-108` treats active plans as temporary working documents and permanent material as domain-owned.

**Implication:** Consolidation needs two layers: one semantic status owner and one lightweight router. Combining status, commands, design history, and durable runtime contracts into a single tracker would recreate the current ownership conflicts.

### Finding 12 — Consolidation Has a Safe Order

A source-backed consolidation order is:

1. Reconcile all stale manifestations inside `evaluation-progress.md`.
2. Repair `campaign-orchestration-prompt.md` to 27 runs and make it the sole scheduling/launch owner.
3. Delete `campaign-run-sheet.md` or replace it with a redirect.
4. Extract any still-unique durable harness rationale, then retire `harness-setup.md` in favor of source, `cli/README.md`, and the retained smoke report.
5. Mark or relocate the website architecture document as historical provenance.
6. Add a compact `plans/README.md` routing index only after those owners are stable.

**Evidence:** Findings 5-11 establish the status conflicts, scope conflict, duplicate operating ownership, obsolete command source, stale runtime contract, historical website role, and missing index.

**Implication:** This order fixes meaning at its current owner before adding navigation, preventing a new index from memorializing stale claims.

## Notes

- The 21 accepted repository artifacts corroborate the tracker’s aggregate completion count and seven complete cells. This assignment did not perform an evaluation-artifact deep dive or reconstruct external `chat-<n>` container contents.
- The Cursor reset date and billing blocker are document-reported state at `evaluation-progress.md:90-99,143-148`. Local-only scope did not independently query the provider account or model catalogue.
- `git status` showed unrelated existing changes to `AGENTS.md` and the shared `research/project-state-recovery/` tree. They were preserved and excluded from this assignment’s conclusions except where the provided current `AGENTS.md` governs repository behavior.

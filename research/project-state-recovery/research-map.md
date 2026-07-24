# Skill Issue Current-State Recovery Research Map

> **Completed research map:** This records the bounded eight-assignment pass and
> does not own live status. Use
> [`plans/current-state.md`](../../plans/current-state.md) for current work.

## Run Contract

- **Goal:** Establish the evidence base for consolidating the Skill Issue repository into one authoritative current progress/task owner plus a separate routing index, without creating duplicate live status ownership.
- **Concrete context:** The repository is a Go CLI and React/Vite website on the protected post-submission development branch. It embeds a pinned Skill Issue plugin submodule and retains substantial plans, evaluations, research, and Codex task/chat history.
- **Source scope:** Local only.
- **Active researcher concurrency:** 5.
- **Total researcher budget:** 8.
- **Research root:** `research/project-state-recovery/`.
- **Final aggregation target:** `research/project-state-recovery/skill-issue-current-state-recovery-deep-research.md`.
- **Requested final shape:** Best-supported answer or direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers.

## Evidence Priority

1. Production source and concrete repository behavior.
2. Live Git branch, worktree, submodule, commit history, and diffs.
3. Accepted evaluation artifacts and other machine-readable retained evidence.
4. Durable plans, research, reports, and public documentation, reconciled against source and Git.
5. Relevant local Codex task/chat history available through thread tools or local session evidence, with access gaps explicit.

## Research Domains

### Current Product and Repository Truth

Parent domain covering the implemented CLI, website, embedded dependency, accepted evaluation artifacts, and public-facing state. It is split across assignments 01, 05, 06, and 07 so no researcher owns the whole domain.

### Current Work and Historical Progress

Parent domain covering branch history, uncommitted state, active plans, completed work, blockers, sequencing, and superseded trackers. It is split across assignments 02, 03, 04, and 08.

### Consolidation and Semantic Ownership

Cross-cutting domain evaluated in every assignment: which document can own current status, which surface should only route/index, which existing artifacts must become historical evidence or redirects, and which files are safe deletion candidates. No researcher may create or edit the proposed owner.

## Discovery Waves

### Wave 1: Map Primary Pathways

Launch five assignments in parallel:

1. Production source and command lifecycle.
2. Git branch, history, and worktree evidence.
3. Active plans and live trackers.
4. Historical research and cleanup classification.
5. Accepted evaluations and retained evidence.

These assignments identify candidate current-state owners, stale trackers, and evidence gaps.

### Wave 2: Targeted Cross-Checks

Backfill three assignments as Wave 1 completes:

6. Website, README, and public state cross-check.
7. Plugin submodule and dependency state.
8. Codex task/chat history and recovery evidence.

Wave 2 validates public claims, dependency sequencing, and historical intent against the primary map.

## Assignment Files

| #   | Assignment                        | Source Targets                                                                                          | Expected Evidence                                                                                                     | Output                                                |
| --- | --------------------------------- | ------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------- |
| 01  | Production source lifecycle       | `cli/cmd`, `cli/internal`, scripts, package configuration                                               | Concrete entrypoints, lifecycle owners, implemented capabilities, incomplete paths, semantic state owner implications | `assignments/01-production-source-lifecycle.md`       |
| 02  | Git history and worktree          | Current branch/status, refs, submodule status, diffs, recent and topic history                          | Accepted completed work, uncommitted changes, branch constraints, chronology, current blockers                        | `assignments/02-git-history-and-worktree.md`          |
| 03  | Active plans and trackers         | `plans/**` and linked implementation/evaluation documents                                               | Live status claims, dependency order, stale/conflicting trackers, owner/index candidates                              | `assignments/03-active-plans-and-trackers.md`         |
| 04  | Historical research and cleanup   | `research/**`, retired/retained reports, repository history where useful                                | Preserve/redirect/delete classifications, historical value, duplicated status claims, uncertainty                     | `assignments/04-historical-research-and-cleanup.md`   |
| 05  | Accepted evaluations and evidence | `evaluations/**`, machine-readable accepted results, retained showcase evidence                         | Completed campaigns, partial work, active evaluation state, evidence-backed next steps                                | `assignments/05-accepted-evaluations-and-evidence.md` |
| 06  | Public documentation and website  | `README.md`, `cli/README.md`, `src/data/**`, pages/components, release metadata                         | Publicly asserted current state, implemented-vs-documented discrepancies, status-owner implications                   | `assignments/06-public-documentation-and-website.md`  |
| 07  | Plugin dependency state           | `.gitmodules`, submodule pointer/source, payload embedding, dependency sync workflow, related commits   | Current pinned dependency truth, completed migration, sync blockers, dependency ordering                              | `assignments/07-plugin-dependency-state.md`           |
| 08  | Codex task and chat history       | Available Codex thread/task tools first; then local session/memory evidence if applicable and permitted | User-confirmed sequencing, accepted completion, unresolved work, access limitations, contradictions with documents    | `assignments/08-codex-task-chat-history.md`           |

## Fan-Out Decisions

- The eight-assignment budget is fully allocated; no mid-run branch may be added.
- Each assignment is narrow and writes exactly one Markdown document.
- Broad document inventories must be ranked inside the assigned file as current/live, historical-preserve, redirect candidate, deletion candidate, or unsupported.
- Candidate documents are deep-dived only when they contain current status, dependency order, accepted completion, blockers, or contradictory ownership claims. Remaining files may be skimmed and classified.
- Claims that cannot be validated within the local scope must be marked unsupported, caveated, or blocked.

## Cross-Check Requirements

- Current-status claims from plans must be checked against source, Git history, and accepted artifacts.
- Completed-work claims must be supported by concrete commits, source, or retained accepted evidence.
- Dependency and next-action claims must be checked against production ownership and the pinned plugin state.
- Deletion candidates must be checked for unique historical evidence, backlinks, or durable domain value.
- Codex-history findings must be labeled by access method and must not override contradictory production or Git truth.

## Aggregation Boundary

The data aggregator receives only this research map and the eight assignment documents. It must identify:

- the best semantic home for one authoritative current progress/task owner;
- a separate index/routing surface that does not duplicate live status;
- completed, active, next, and later work with dependency order;
- blockers and unsupported claims;
- stale/conflicting trackers;
- preserve, redirect, and safe-deletion candidates;
- conditional alternatives and rejected lower-fit interpretations.

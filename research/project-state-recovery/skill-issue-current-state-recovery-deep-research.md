# Skill Issue Current-State Recovery

> **Historical recovery evidence:** This report supports the consolidation but
> does not own live status. Use
> [`plans/current-state.md`](../../plans/current-state.md) for current work,
> blockers, and next actions.

## Best-Supported Direction

### Establish one repository-wide live owner

Create **`plans/current-state.md`** as the single authoritative owner of current repository progress and task state.

That document should own only:

- the verified implemented and accepted baseline;
- current active work;
- current blockers and prerequisites;
- dependency-ordered next actions;
- later and explicitly deferred phases;
- concise evidence links;
- closed or superseded work only when needed to prevent it from being reopened.

It should not restate production contracts, harness defaults, payload schemas, accepted-result details, public copy, release metadata, campaign commands, or historical research conclusions. Those meanings already have concrete domain owners. Production source contains no durable project roadmap or task state, and its run statuses are evaluation-local and intentionally removed during cleanup (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Findings 7, 8, and 12; underlying evidence: `cli/internal/runstate/runstate.go:18-50,242-272,330-347` and `cli/internal/command/command.go:55-69`).

The current campaign ledger, `plans/skill-calling-evaluation-campaign/evaluation-progress.md`, is too narrow to own repository-wide work: it is explicitly scoped to campaign execution, contains internally conflicting live state, and has no semantic place for the broken plugin-integration paths, branch-safety policy, sync-workflow routing, public-copy drift, or release-verification caveats (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Findings 1, 5, 6, 7, and 11; `research/project-state-recovery/assignments/06-public-documentation-and-website.md`, Findings 3-6).

### Add one separate routing surface

Create **`plans/README.md`** as the repository's stable planning and evidence index.

The index should contain links and one-sentence responsibility statements only. It should route to:

- `plans/current-state.md` for live repository progress and next work;
- `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md` for campaign procedure and scheduling;
- production source and `cli/README.md` for CLI behavior;
- `evaluations/skill-calling/results/accepted/` plus `scripts/update-website-results.mjs` for accepted public benchmark data;
- `src/data/siteData.ts` for curated public interpretation, status copy, and release metadata;
- retained historical decisions, evaluations, and research at their domain homes.

It must not copy counts, blockers, dates, task checkboxes, model identifiers, commands, or next actions. The absence of any current `plans/` index and the explicit division between execution ownership and scheduling ownership support this split (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Finding 11; underlying evidence: `plans/skill-calling-evaluation-campaign/evaluation-progress.md:3-7,62` and `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md:3-7`).

### Make the transition atomically

The owner transition should be one reviewed documentation change:

1. Create and populate `plans/current-state.md` from the reconciled snapshot in this report.
2. Reconcile and extract the unique blocked-run recovery details from `plans/skill-calling-evaluation-campaign/evaluation-progress.md`.
3. Convert that campaign file from a live tracker into a concise dated historical redirect to `plans/current-state.md`, the accepted artifacts, and the campaign procedure.
4. Repair the 27/30 conflicts in `campaign-orchestration-prompt.md` and leave scheduling and launch procedure there exclusively.
5. Add `plans/README.md` only after those owners are stable.

This sequence prevents a period in which both the new file and the campaign ledger claim live status ownership. The existing tracker should be repaired before extraction because its current summary, lane headings, failure log, and closing notes disagree (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Findings 5 and 12).

## Reconciled Current Snapshot

### Accepted completed work

#### Protected repository baseline

- The protected `build-week-2026` tag and local `main` both identify `b64d9b5ff095191b971e4c2a953fbbc8bf3c352e`. The development branch descends linearly from that snapshot; no snapshot repair is indicated (`research/project-state-recovery/assignments/02-git-history-and-worktree.md`, Finding 1).
- The inspected current development baseline is `4979bd035680fd6e1142eabcd67b352d9221d713` on `codex/post-submission-development`, locally aligned with its recorded upstream. The three accepted post-snapshot commits are:
  - `a7ee8ab0199076d38a297a370755fffc952f4382` — marketplace packaging transition;
  - `1ebbe4aaf7d65427e68aa9f1084893dbf4456075` — lifecycle skills sourced from the plugin submodule;
  - `4979bd035680fd6e1142eabcd67b352d9221d713` — evaluation-agent documentation clarification.
    The local reflog records a push after each commit, but no network fetch was performed during this research (`research/project-state-recovery/assignments/02-git-history-and-worktree.md`, Findings 2 and 3).

#### Implemented product baseline

- The CLI implements `help`, `version`, `install`, `doctor`, `uninstall`, and governed `evaluate`; the hidden `signal` route is internal instrumentation. Generation and refinement are installed-skill workflows rather than CLI subcommands (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Finding 2; underlying evidence: `cli/internal/command/command.go:49-73,141-174,475-494`).
- Installation and uninstallation are implemented for Claude Code, OpenAI Codex, Cursor, OpenCode, and Pi. The payload contains eleven canonical skills, and none of the five harnesses is marked installation-in-progress (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Finding 3; underlying evidence: `cli/internal/harness/harness.go:64-70,88-91` and `cli/internal/payload/assets/manifest.json:5-50`).
- Governed evaluation has a complete source path from input preparation and confirmation through isolated runtime preparation, authentication, skill installation, multi-turn replay, attribution, result derivation, restoration, and cleanup. Durable `result.json` and `website.json` artifacts are separate from private recovery state (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Findings 6-9; underlying evidence: `cli/internal/evaluation/evaluation.go:197-501,871-960,1198-1229,1287-1339`).
- Cross-platform packaging builds six Darwin/Linux/Windows archives plus checksums. Native evaluation remains Darwin/Linux-only and requires external harness, authentication, model, protocol, and filesystem readiness (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Findings 9-11; underlying evidence: `cli/scripts/build-cross-platform.sh:9-75` and `cli/internal/evaluation/evaluation.go:96-101`).
- The website is a four-destination React/Vite application with Explore, Method, Analysis, and Project surfaces. Its principal product, methodology, and accepted-result boundaries agree with production source, aside from the stale public-state claims identified below (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Finding 10; `research/project-state-recovery/assignments/06-public-documentation-and-website.md`, Findings 1, 2, and 7).

#### Accepted benchmark and evaluation evidence

- Seven benchmark configurations and 21 governed runs are accepted. The accepted directory and `src/data/publishedWebsiteArtifacts.json` contain the same normalized schema-v2 set, with three scenarios in each complete cell (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 1; underlying evidence: `scripts/update-website-results.mjs:38-136` and `src/data/analysisData.ts:48-71`).

| Accepted cell                     | Accepted run IDs                                                                                           |
| --------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| Claude Code / `claude-fable-5`    | `daf204f8d746857d6cdea31f98c0abd0`, `28bbfaafbd2350de09f28ee62b6c647a`, `114e9a28c466d5f5de46a5560ac93def` |
| Claude Code / `gpt-5.6-sol`       | `e20e00601b9cf52124680fb0c0a937e5`, `a0f46b4ec92e74c09eba16abcd1121a7`, `12192ed7bb64e3cf049d712876b99bc5` |
| OpenAI Codex / `gpt-5.6-sol`      | `f80e38d1fe99e121f233428792c2f6a2`, `f885485d2718735c688e20c9cdcf8f34`, `84ffd0408208a60685e896edf240f3cc` |
| Cursor / `composer-2.5`           | `bfebefcbb497c53f386d388eb5b86854`, `4c15412e323ba9ecdfb665f045fe0406`, `aee68fe2c0796948d21e2d4fcf29d7af` |
| Cursor / `cursor-grok-4.5-medium` | `be67f5c9726747e18374dfc2fef587ea`, `09d482d1ec542c92f36e3e4d235d4ad4`, `fd38e8865de744325280939db588d196` |
| Pi / `openai-codex/gpt-5.6-sol`   | `a13afc4d5d8682721c31b66efd561d9c`, `3d169365a52ae92b71d7da2d6209661b`, `10ad4d711e0a8d1be4eb62c70367b2b1` |
| OpenCode / `openai/gpt-5.6-sol`   | `72049527285a7fc9964e1b1e9b98e53e`, `6f29ec29552bd293487fd0fbcda570cd`, `f0233161faa6dac6e22542516eb6f308` |

The run-ID set is supplied by `research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Sources. The two corrected Codex artifacts `f80e38d1fe99e121f233428792c2f6a2` and `f885485d2718735c688e20c9cdcf8f34` retain raw-transcript reconciliation receipts and remain accepted (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 5).

- Real-harness smoke qualification completed built-in and custom two-turn routes for Codex, Cursor, Claude Code, and Pi, including local artifact and cleanup checks. This is qualification evidence, not accepted benchmark evidence (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 6; underlying evidence: `evaluations/skill-calling/smoke/real-harness-smoke-report.md:9-18,28-82`).
- The 21-iteration skill-system production-refinement campaign completed evaluator, Intake, Generation, and cross-workflow handoff at its authorized boundary. Runtime evaluation of the additionally generated `repository-owner-finder` skill was explicitly deferred and is not part of the completion claim (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 7; underlying evidence: `evaluations/skill-system-production-refinement/progress.md:27-43,45-117` and `evaluations/skill-system-production-refinement/final-audit.md:3-37`).
- Seven scenario-skill-refinement campaigns are complete with matching retained target hashes: code implementation, code testing, document update, prompt writing, skill authoring, system change ownership, and systematic debugging (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 9).
- Ten showcase evaluation campaigns are complete at their recorded evidence boundaries with matching target hashes. Accessibility First Pass does not claim native CLI activation; Safe Share Redactor is explicitly invocation-only; historical capacity or failed cycles do not reopen the terminal passing states (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 10).

#### Plugin dependency baseline

- The parent gitlink, initialized dependency checkout, and locally recorded dependency `main`/`origin/main` all identify `62c3921aca885df01ab9c8ca6d7ff20381b0c038`. The checkout is clean and the plugin manifest identifies `0.1.0+codex.20260723022002` (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Finding 1).
- The compiled payload migration is complete: `bundle.go`, the payload manifest, payload validation, CI, local known-good builds, and cross-platform builds consume the exact pinned dependency revision (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Findings 2-4).
- Standalone plugin publication and post-submission parent publication were user-authorized and completed in recovered local history. Thread `019f8b6f-25d4-7710-880a-cee5b87abb72` records destination-first confirmation, all-files authorization, the `a7ee8ab` push, plugin source/cache verification at `62c3921`, and the later `4979bd0` documentation push (`research/project-state-recovery/assignments/08-codex-task-chat-history.md`, Finding 6).

### Current active and unresolved work

No evaluation lane or production implementation process was active in the inspected evidence. The concrete current work was:

- a live, uncommitted `AGENTS.md` change that protects `main`, keeps work on `codex/post-submission-development`, and replaces the committed dependency-sync merge wording with review-only wording;
- the untracked `research/project-state-recovery/**` recovery run, now culminating in this synthesis.

At the Git research snapshot, no staged change, extra worktree, branch divergence, submodule dirtiness, or other untracked product file was found (`research/project-state-recovery/assignments/02-git-history-and-worktree.md`, Findings 2, 7, and 8). The `AGENTS.md` change is pending deliberate review and acceptance; it is not accepted committed history.

The actionable unresolved product and documentation defects are:

1. **Broken repository-local lifecycle-skill discovery.** The tracked `.codex/skills/skill-intake`, `.codex/skills/skill-generation`, and `.codex/skills/skill-evaluation-and-refinement` symlinks still target the deleted `plugins/skill-issue/skills/...` tree. Five showcase workflow prompts name the same absent path. The canonical skill directories exist inside the pinned dependency (`research/project-state-recovery/assignments/02-git-history-and-worktree.md`, Finding 6; `research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Finding 5).
2. **Unsupported plugin-sync branch routing.** `.github/workflows/sync-skill-issue-plugin.yml` exists on the development branch but is absent from protected/default `main`; the documented dispatch omits `--ref`, and the target base for the automation branch is not named. Source proves the ambiguity but does not prove the workflow fails on GitHub (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Finding 6).
3. **Stale internal campaign state.** The 21/27 headline is supported, but Cursor/Composer remains labelled 0/3 in one heading; older failure-log cells remain pending/running after accepted reruns; closing notes say four completions; retained container notes conflict with later accepted runs; 27/30 scope and concurrency rules conflict across campaign documents (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Findings 5-8).
4. **Stale public interpretation and milestone copy.** `src/data/siteData.ts:253-262` says Claude Code/Codex has one accepted run and is excluded, while three accepted artifacts, the configured available cell, and published aggregates include it. README/project copy still describes release preparation or a first complete path “taking shape” despite a tagged beta, enabled download links, and existing packaging (`research/project-state-recovery/assignments/06-public-documentation-and-website.md`, Findings 3 and 5).

### True blockers

#### Cursor campaign blocker

Six governed runs remain blocked: Cursor/Codex and Cursor/Fable, three scenarios each. Their partial executions stopped between turns 15 and 21 because the Cursor API usage limit was exhausted. The recorded operator decision is to wait for the 2026-08-21 cycle reset or raise the limit earlier. No accepted artifacts exist for either cell, and the partial failures are not scoreable (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Findings 2 and 11; `research/project-state-recovery/assignments/08-codex-task-chat-history.md`, Finding 5).

The recovery identities that must survive tracker consolidation are:

| Blocked cell   | Retry containers                |
| -------------- | ------------------------------- |
| Cursor / Codex | `chat-12`, `chat-13`, `chat-14` |
| Cursor / Fable | `chat-15`, `chat-16`, `chat-36` |

`chat-36` replaces stale `chat-17`; `chat-17` contains unrelated residue and is excluded (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Finding 4).

#### Codex task-history access blocker

Live Codex thread and task tools were unavailable. Tool discovery exposed no task-list, thread-list, task-read, or thread-read surface. Chat-history recovery is therefore bounded to three known repository-specific local session records and related memory indexes:

- `019f7734-dfae-76f0-869b-234d75d63692`;
- `019f861d-87ff-77f2-aa16-4561c8249917`;
- `019f8b6f-25d4-7710-880a-cee5b87abb72`.

This prevents a complete task inventory, current sidebar-metadata recovery, or proof that no other relevant archived thread exists (`research/project-state-recovery/assignments/08-codex-task-chat-history.md`, Finding 1 and Notes). Chat history is supporting evidence only and must not own live status.

**Parent follow-up:** After aggregation, the parent task gained access to live
Codex thread listing and reading. It reviewed the recent and pinned Skill Issue
tasks, including the skill-generation/evaluation split, exporter audit,
post-submission branch decision, publication task, and pinned main task. That
history confirms the plugin migration and competition branch decisions, closes
the earlier exporter and evaluator-launch concerns against current pinned
source, and reveals no additional active product execution. The visible task
list is recency-bounded, so older unlisted tasks may still exist.

#### Future plugin-sync blocker

The branch/ref/base contract must be resolved before the next ordinary plugin dependency sync. Until then, the source-described automation path is operationally unsupported. This does not block current compilation from the pinned, initialized revision (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Findings 1, 4, and 6).

### Dependency-ordered next actions

#### Phase 1 — Establish authoritative state without duplicate ownership

1. Review the existing `AGENTS.md` safety-policy change without overwriting or folding it into unrelated recovery edits.
2. In one documentation change, create `plans/current-state.md`, migrate the reconciled live state and six blocked container identities, convert `evaluation-progress.md` to a dated historical redirect, repair `campaign-orchestration-prompt.md` to nine configurations and 27 runs, and add routing-only `plans/README.md`.
3. Reconcile `src/data/siteData.ts` and the README/project milestone language against the accepted artifacts and shipped-beta evidence. Keep accepted-result truth derived from `evaluations/skill-calling/results/accepted/` through `scripts/update-website-results.mjs`, rather than copied into a second result ledger.

This first phase makes current work visible and removes stale competing owners before further implementation.

#### Phase 2 — Finish the parent-repository plugin integration

1. Choose the repository-local lifecycle-skill route. The strongest fit is to retarget the three tracked `.codex/skills` symlinks to the corresponding directories inside the pinned `dependencies/codex-skill-issue-plugin/plugins/skill-issue/skills/` tree, because repository-local discovery already uses tracked symlinks and the parent intentionally pins that dependency. Removing them in favor of external installation is viable only if project-local discovery is no longer required.
2. Update the five retained showcase workflow prompts to the surviving canonical route or replace path-specific instructions with the installed skill invocation owned by the plugin.
3. Resolve and document the sync workflow's dispatch ref and automation-branch base on the protected development topology.
4. Validate the parent with the submodule initialized, review any generated pointer-only branch, and build a new known-good CLI only from an accepted parent commit and exact gitlink (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Finding 7).

#### Phase 3 — Resume the blocked benchmark

1. Wait for the recorded Cursor cycle reset or obtain an earlier operator-approved capacity change.
2. Re-resolve Cursor's live model identifiers because they drift.
3. Recreate each failed `chat-<n>` container under its stable identity, excluding `chat-17`.
4. Run scenarios sequentially inside each of the two Cursor lanes; choose cross-lane concurrency only after confirming the shared account is healthy.
5. Accept only tooling-complete outputs, update the accepted artifact set, regenerate `src/data/publishedWebsiteArtifacts.json`, and update public interpretation from the new accepted state.
6. Produce the operator-owned final campaign report and make publication, commit, and cleanup decisions separately.

The ordering and retry rules are supplied by `research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Findings 3 and 4; underlying procedure is `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md:29-58,66-84`.

### Later and explicitly deferred phases

- Runtime-evaluate the generated `repository-owner-finder` skill only if that separate evidence is wanted; the completed production-refinement campaign must remain closed (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 7).
- Review smoke outputs for publication only through a separate acceptance decision; smoke qualification does not become benchmark evidence automatically (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Findings 6 and 11).
- Consider evaluation run listing, status inspection, resumption, Windows evaluation qualification, OpenCode version updates, and unknown lifecycle-option rejection only as newly accepted product tasks. Production source exposes those bounded limitations but does not establish them as active work (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Findings 8 and 10).
- Perform research-evidence compaction and cleanup after the live-owner transition. Raw assignment sets should be deleted only after retained syntheses preserve their unique source links, versions, conflicts, validations, and dead ends (`research/project-state-recovery/assignments/04-historical-research-and-cleanup.md`, Finding 16).

## Document Disposition Map

### Live owner

| Path                     | Disposition                       | Required meaning                                                                                                                                         |
| ------------------------ | --------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `plans/current-state.md` | **Create as the only live owner** | Own verified baseline, active work, blockers, dependency-ordered next actions, later phases, and concise evidence links. No domain-contract duplication. |

### Routing only

| Path                            | Disposition                               | Required meaning                                                                                                                                     |
| ------------------------------- | ----------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------- |
| `plans/README.md`               | **Create as routing-only**                | Stable links and responsibility statements; no counts, dates, tasks, blockers, commands, or mutable status.                                          |
| `README.md` project-status area | **Reduce to public milestone plus route** | Keep public product truth and a stable route to `plans/README.md`; remove pre-release drift and direct exposure of the inconsistent campaign ledger. |

### Durable domain owners

| Path or set                                                                                         | Disposition                                                                       | Owned meaning                                                                                                           |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `cli/internal/**`, `bundle.go`, `cli/internal/payload/assets/manifest.json`, build scripts          | **Preserve at semantic home**                                                     | Concrete CLI behavior, harness registry, payload composition, evaluation lifecycle, run state, and build behavior.      |
| `cli/README.md` and `INSTALL.md`                                                                    | **Preserve and reconcile when behavior changes**                                  | CLI operator contract and installation/release instructions.                                                            |
| `evaluations/skill-calling/results/accepted/` and `scripts/update-website-results.mjs`              | **Preserve at semantic home**                                                     | Accepted compact result set and deterministic publication projection.                                                   |
| `src/data/siteData.ts`                                                                              | **Preserve and repair stale copy**                                                | Curated public interpretation, public status copy, release metadata, and summary metrics.                               |
| `src/data/evaluationData.ts`, `src/data/analysisData.ts`, `src/data/publishedWebsiteArtifacts.json` | **Preserve at semantic home**                                                     | Result adaptation, configured availability, derived complete-cell analysis, and generated published data.               |
| `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md`                          | **Preserve as procedure owner after repair**                                      | Scheduling, launch, retry, acceptance, and final-report procedure; use nine configurations and 27 runs; no live status. |
| `AGENTS.md`                                                                                         | **Preserve as repository contract; review current uncommitted change separately** | Branch safety, repository ownership, and operational constraints.                                                       |
| `dependencies/codex-skill-issue-plugin`                                                             | **Preserve exact gitlink boundary**                                               | Standalone plugin content at the pinned revision; parent owns pin, embedding, and installation adaptation.              |

### Historical preserve

| Path or set                                                                                                                                    | Disposition                                      | Reason                                                                                                                   |
| ---------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `plans/website/reference-and-architecture.md`                                                                                                  | **Preserve as historical decision record**       | Original visual, stack, and ownership provenance; its one-page/two-environment information architecture is superseded.   |
| `evaluations/skill-calling/smoke/real-harness-smoke-report.md`                                                                                 | **Preserve as qualification history**            | Unique real-harness route and cleanup evidence, with its non-benchmark boundary intact.                                  |
| `evaluations/skill-system-production-refinement/**` terminal progress/audits                                                                   | **Preserve as completed campaign evidence**      | Proves the authorized completion boundary and explicit runtime deferral.                                                 |
| `evaluations/scenario-skill-refinement/**` status and retained evidence                                                                        | **Preserve as completed campaign evidence**      | Seven passed targets with matching hashes and diagnostic history.                                                        |
| `showcase-skills/*/evaluation/**` terminal status and retained evidence                                                                        | **Preserve as completed campaign evidence**      | Ten passed campaigns with their scoped caveats.                                                                          |
| Benchmark analysis and external methodology final syntheses plus linked assignments                                                            | **Preserve as dated historical source basis**    | Durable reporting, uncertainty, disclosure, and claim-discipline reasoning; current website source owns public behavior. |
| `research/cross-platform-cli-distribution/cross-platform-cli-distribution-report.md`                                                           | **Preserve as history or compact ADR**           | Adopted Go/archive/checksum rationale plus clearly unimplemented alternatives.                                           |
| Free-hosting synthesis and linked assignments                                                                                                  | **Preserve at current research home**            | Still has a deliberate backlink from the retained website architecture decision.                                         |
| Direct-installation, manual-invocation, sub-agent-launch, portability, OpenCode qualification, video, and TTS syntheses and linked assignments | **Preserve as dated history pending compaction** | Unique path, qualification, ecosystem, and build-history evidence; external details can drift.                           |
| This recovery synthesis                                                                                                                        | **Preserve as the consolidation evidence basis** | Records why the live owner, index, ordering, and cleanup decisions were selected.                                        |

### Convert to concise historical redirect

| Path                                                                                                          | Disposition                                                                          | Required preservation before conversion                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `plans/skill-calling-evaluation-campaign/evaluation-progress.md`                                              | **Convert after reconciliation and extraction**                                      | Move the six blocked container identities and current blocker into `plans/current-state.md`; route accepted scores to accepted artifacts and procedure to the orchestration prompt.                           |
| `research/benchmark-analysis-page/analysis-page-idea-notes.md`                                                | **Convert**                                                                          | Preserve only lasting editorial rationale not already in current website copy; remove stale 19-run/in-flight status.                                                                                          |
| `research/testing-methodology/local-system-research.md`                                                       | **Convert**                                                                          | Preserve useful source trace; remove stale 30-run, zero-completion, and empty-publication claims.                                                                                                             |
| `research/implementation-research/website-chart-migration/website-chart-migration-implementation-research.md` | **Convert**                                                                          | Route current behavior to source and preserve only historical migration rationale; its detailed assignments require a separate compaction check.                                                              |
| `plans/harness-setup.md`                                                                                      | **Prefer deletion after extraction; otherwise convert**                              | Current Codex authentication, `CODEX_HOME`, attribution, and campaign-state claims are superseded; preserve unique smoke history in the existing smoke report and current behavior in source/`cli/README.md`. |
| `plans/skill-calling-evaluation-campaign/campaign-run-sheet.md`                                               | **Delete; use a one-line redirect only if path continuity is deliberately required** | It retains Kilo, stale models, stale `chat-17`, missing reasoning flags, and duplicated commands.                                                                                                             |

### Firm safe deletion candidates

These are firm repository cleanup candidates once the live-owner transition no longer cites them:

- completed research orchestration maps whose final syntheses exist, including completed `research-map.md` files and `research/testing-methodology/00-external-research-map.md`;
- `research/website-refactor/screenshots/**`;
- `research/website-two-arm-navigation/screenshots/**`;
- `research/website-two-arm-navigation/reference-findings.md`, whose findings section is empty and whose decisions are implemented;
- untracked `.DS_Store` files and the empty local `research/audits/` shell;
- `plans/skill-calling-evaluation-campaign/campaign-run-sheet.md`, unless deliberate legacy-path continuity requires a one-line redirect.

The screenshot sets have no backlinks or runtime role, and current source plus Git history own the implemented design (`research/project-state-recovery/assignments/04-historical-research-and-cleanup.md`, Findings 13, 15, and 17). The current recovery map remains active evidence until this synthesis and the later consolidation are accepted; it is not part of the immediate completed-map deletion set.

### Conditional deletion candidates

- `plans/harness-setup.md`: delete only after confirming every still-unique rationale is preserved in source, `cli/README.md`, or the smoke report.
- `research/video-reference-screenshots/**`: delete if the public video is accepted as the authoritative history; preserve with a manifest if offline reconstruction of the video-era site is a real requirement.
- Website-chart migration assignments: deletion is unsupported until a compaction pass preserves any unique historical schema defects, provenance decisions, and source traces.
- Other raw research assignment corpora: preserve while parent syntheses link them; delete only after evidence-bearing compaction.
- The project-state-recovery assignments and map: retain through consolidation review; later delete or compact only if this synthesis preserves all needed claim-level evidence and no active implementation plan links them.

These conditions come from `research/project-state-recovery/assignments/04-historical-research-and-cleanup.md`, Findings 3, 14, 16, and 18.

## Conditional Alternatives

### Broaden the existing campaign ledger instead of creating a new file

Renaming and completely rewriting `plans/skill-calling-evaluation-campaign/evaluation-progress.md` into `plans/current-state.md` could preserve path history while achieving one owner. It fits if Git move history is valued and the campaign-specific body is deliberately replaced rather than incrementally patched.

It is lower fit than creating the new owner and converting the old path to a redirect because the existing file's title, directory, purpose statement, run tables, failure log, and stale notes all encode campaign-only semantics. A superficial broadening would preserve the wrong owner boundary (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Findings 1 and 5).

### Retain a detailed immutable campaign record

Instead of a concise redirect, the reconciled campaign ledger could become a dated, immutable historical execution record if the repository requires local preservation of attempts, failures, and container history. It must remove checkboxes, “current,” “next,” and mutable blocker language; `plans/current-state.md` must remain the only place that says what is active or next.

This fits only if the unique failure history is judged durable after blocked-run recovery details are migrated. Accepted scores still belong to the accepted artifacts.

### Remove project-local lifecycle symlinks

The three broken `.codex/skills` symlinks could be removed rather than retargeted if repository-local discovery is intentionally replaced by installation of the packaged plugin. This fits a workflow in which every contributor installs the plugin before repository work.

Retargeting is stronger for the current repository because five other project-local skill symlinks resolve, the lifecycle skills are required by local workflows, and the exact dependency revision is already pinned. Removal would need a documented bootstrap route and prompt changes (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Findings 4 and 5).

### Preserve legacy command and screenshot paths

A one-line historical redirect can replace the run sheet if external documentation may link to it. Video-reference screenshots can remain if offline reconstruction is required. Neither requirement is established by the constrained corpus, so deletion remains the default.

## Rejected or Lower-Fit Interpretations

### Use `README.md` as the live task owner

Rejected. The README is public product documentation and currently mixes product explanation, development instructions, publication procedure, and stale project-status prose. Making it the live task ledger would conflate public truth with internal operations and expose transient blockers and commands to the wrong audience (`research/project-state-recovery/assignments/06-public-documentation-and-website.md`, Findings 4, 5, and 8).

### Use `plans/README.md` as both index and live owner

Rejected. A combined index and status file would force navigation to change with every task update and recreate duplicate status wherever a domain needs a link. The index should remain stable while `plans/current-state.md` changes.

### Keep `evaluation-progress.md` as the repository-wide owner

Rejected for repository-wide recovery. It is the best existing owner for campaign execution only, but its directory and structure cannot semantically own plugin integration, branch safety, public copy, release verification, or general later phases without becoming a mixed-purpose ledger (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Findings 1 and 11).

### Derive live tasks from production limitations

Rejected. Windows evaluation, run listing/resumption, OpenCode version changes, and unknown-option rejection are concrete limitations, but no source marker, accepted plan, or user decision makes them active tasks. Dormant capability scaffolding such as `InstallationInProgress` is also not evidence of unfinished work (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Findings 3, 8, and 10).

### Treat run state or accepted artifacts as project status

Rejected. Private run state is evaluation-local and deleted after cleanup; accepted compact artifacts prove public score projections and cell membership but omit roadmap, cleanup, detailed provenance, and next actions (`research/project-state-recovery/assignments/01-production-source-lifecycle.md`, Finding 7; `research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 4).

### Treat Codex chat history as authoritative

Rejected. Known threads contain valuable user decisions and approvals, but threads end mid-work, memory summaries are point-in-time, live thread/task tools were unavailable, and repository source already supersedes some historical thread state (`research/project-state-recovery/assignments/08-codex-task-chat-history.md`, Findings 1, 3, and 7).

### Treat the benchmark as complete or as a 30-run campaign

Rejected. The current registry and tracker define nine configurations and 27 runs; 21 runs in seven cells are accepted and six Cursor runs are blocked. Kilo was retired, while the 30-run wording survives only in unreconciled campaign documents (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Findings 2 and 6).

### Treat the plugin migration as wholly complete

Rejected. The compiled CLI and gitlink migration are complete, but the three project-local lifecycle-skill symlinks and five workflow prompts still point to deleted paths. The correct statement is “compiled payload migration complete; repository-local integration incomplete” (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Findings 3-5).

### Delete all research assignments because Git retains history

Rejected. Many final syntheses link their assignments, and assignments contain versions, source traces, dead ends, and validation absent from the summaries. Git history should not be the only access path to still-important evidence (`research/project-state-recovery/assignments/04-historical-research-and-cleanup.md`, Finding 16).

## Unsupported Claims, Caveats, and Evidence Limits

- **Unsupported current remote alignment:** local parent and dependency remote-tracking refs matched the inspected commits, but no fetch occurred. Current GitHub branch tips and the latest standalone-plugin `main` are unverified (`research/project-state-recovery/assignments/02-git-history-and-worktree.md`, Notes; `research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Notes).
- **Unsupported remote release and deployment availability:** local tags, filenames, links, build scripts, and deployment workflows exist, but no retained GitHub Release asset manifest or Pages deployment receipt proves the endpoints currently resolve (`research/project-state-recovery/assignments/06-public-documentation-and-website.md`, Finding 6).
- **Caveated accepted-artifact provenance:** the 21 accepted schema-v2 files prove accepted public projections and exact configuration membership. They do not independently prove harness versions, cleanup, reasoning, full transcript provenance, or attempt history (`research/project-state-recovery/assignments/05-accepted-evaluations-and-evidence.md`, Finding 4).
- **Caveated Cursor reset and billing state:** the 2026-08-21 reset and billing alternative are authored tracker state and were not independently queried from the provider account (`research/project-state-recovery/assignments/03-active-plans-and-trackers.md`, Notes).
- **Unsupported partial-run scoring:** the six Cursor partial runs have no accepted artifacts and insufficient retained repository-local evidence for valid scores. They must remain excluded (`research/project-state-recovery/assignments/08-codex-task-chat-history.md`, Finding 5).
- **Caveated sync-workflow failure claim:** source proves the missing default-branch workflow and unspecified ref/base contract; it does not prove GitHub dispatch definitely fails because external workflow semantics were not checked (`research/project-state-recovery/assignments/07-plugin-dependency-state.md`, Finding 6 and Notes).
- **Blocked complete chat/task recovery:** live Codex thread and task tools were unavailable, so the three opaque thread IDs are a bounded sample rather than a complete inventory (`research/project-state-recovery/assignments/08-codex-task-chat-history.md`, Finding 1).
- **Unsupported blanket deletion:** raw assignment corpora and website-chart assignments cannot be declared disposable until evidence compaction establishes that no unique source trace or decision remains (`research/project-state-recovery/assignments/04-historical-research-and-cleanup.md`, Findings 3 and 16).

## Conclusion

The strongest recovery design is **`plans/current-state.md` as the sole live repository progress/task owner** and **`plans/README.md` as a status-free routing index**. The transition must demote the campaign ledger from live ownership, preserve its six blocked retry identities, repair the campaign procedure, and route public results, source behavior, public interpretation, durable contracts, and historical evidence back to their existing semantic owners.

The immediate work is documentation-state consolidation, repository-local plugin integration repair, sync-route clarification, and public-copy reconciliation. The only current external execution blocker is Cursor capacity for six governed runs. The protected competition snapshot, accepted 21-run evidence set, completed refinement campaigns, compiled plugin payload migration, and post-submission development history should remain closed and preserved rather than reopened.

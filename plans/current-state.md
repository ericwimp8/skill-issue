# Skill Issue Current State

> **Authoritative status owner:** This is the only document that owns current
> repository progress, blockers, next actions, and later phases. Update it when
> work changes state. Other plans and research documents are procedure,
> reference, or historical evidence.

Last reconciled: 2026-07-24

## Current Snapshot

| Area                           | Current state                                                                                   |
| ------------------------------ | ----------------------------------------------------------------------------------------------- |
| Working branch                 | `codex/post-submission-development`                                                             |
| Accepted branch baseline       | `4979bd035680fd6e1142eabcd67b352d9221d713`                                                      |
| Locally recorded upstream      | Aligned with the accepted branch baseline; no fetch was performed during recovery               |
| Protected competition snapshot | `main` and `build-week-2026` both resolve locally to `b64d9b5ff095191b971e4c2a953fbbc8bf3c352e` |
| Active execution               | No product implementation or evaluation process is running                                      |
| Active repository work         | Documentation recovery and competition-safety changes are uncommitted and awaiting review       |
| Accepted benchmark             | 21 accepted runs across seven complete configurations                                           |
| Blocked benchmark              | Six Cursor runs across two configurations                                                       |

The competition snapshot rules in [root `AGENTS.md`](../AGENTS.md) control every
future Git operation. `main` and `build-week-2026` are immutable project
boundaries.

## Active Work

The current task is to review and accept this recovered status structure:

1. Review the competition-safety update in `AGENTS.md`.
2. Review this current-state owner, the status-free [plans index](README.md),
   the retired campaign trackers, and the supporting
   [recovery research](../research/project-state-recovery/skill-issue-current-state-recovery-deep-research.md).
3. Keep the working tree on `codex/post-submission-development`.
4. After acceptance, begin Phase 1 below.

No other document should be updated as a competing live task list.

## Accepted Completed Work

### Competition and branch baseline

- The competition submission is preserved at `build-week-2026` on the current
  local `main` commit.
- Post-submission development is isolated on
  `codex/post-submission-development`.
- The accepted post-snapshot commits package the Skill Issue plugin, source the
  lifecycle skills from the pinned plugin submodule, and clarify evaluation
  agent configuration.

### Product baseline

- The Go CLI implements installation, diagnosis, uninstallation, and governed
  evaluation for Claude Code, OpenAI Codex, Cursor, OpenCode, and Pi.
- The embedded payload contains the three Skill Issue lifecycle skills and eight
  parent-owned supporting or evaluation skills.
- Governed evaluation owns isolated preparation, multi-turn replay,
  attribution, result derivation, restoration, and cleanup.
- Cross-platform packaging produces the six configured Darwin, Linux, and
  Windows archives; native runtime qualification remains narrower than build
  availability.
- The React/Vite website publishes the product, method, accepted benchmark
  results, analysis, and project surfaces.

Current CLI behavior is owned by [CLI source and documentation](../cli/README.md).
Plugin content is owned by the pinned
[Skill Issue plugin dependency](../dependencies/codex-skill-issue-plugin/plugins/skill-issue/).

### Accepted evaluation evidence

- The accepted benchmark set contains 21 schema-v2 artifacts covering seven
  complete harness/model configurations.
- The accepted artifacts and the generated website collection contain the same
  normalized result set.
- Real-harness smoke qualification is retained separately and does not broaden
  the accepted benchmark claim.
- The 21-iteration skill-system production-refinement campaign completed at its
  authorized boundary.
- Seven scenario-skill refinement campaigns and ten showcase campaigns reached
  their recorded terminal evidence boundaries.
- The earlier concern about `export-codex-evidence.mjs` is closed in the pinned
  plugin: the narrow exporter is absent, and the evaluation skill now requires
  bespoke privacy-safe evidence containing the full task-owned mock
  conversation and the exact native evidence needed for each claim.
- The recent task-history concern about harness-specific launch guidance is
  closed: the evaluator requires fresh independent agents and stops when the
  current environment cannot supply them.

Accepted benchmark truth is owned by
[the accepted artifacts](../evaluations/skill-calling/results/accepted/) and
[the deterministic website publication script](../scripts/update-website-results.mjs).

## Dependency-Ordered Work

### Phase 1 — Repair repository-local plugin integration

**Status:** Ready after the recovered documentation is accepted.

1. Retarget or deliberately remove the three broken project-local lifecycle
   skill links:
   - `.codex/skills/skill-intake`
   - `.codex/skills/skill-generation`
   - `.codex/skills/skill-evaluation-and-refinement`
2. Update the five showcase workflow prompts that still name the deleted
   `plugins/skill-issue/skills/...` tree.
3. Validate project-local discovery against the pinned dependency without
   editing the submodule checkout.
4. Re-run the focused payload, CLI, link, and Markdown checks required by the
   touched files.

The strongest current fit is to point project-local discovery at the pinned
dependency because the parent repository already treats that gitlink as the
canonical lifecycle-skill source. Removing project-local discovery remains a
valid alternative only with an explicit plugin-install bootstrap contract.

### Phase 2 — Resolve non-main plugin synchronization

**Status:** Design decision required before the next plugin publication.

1. Define the non-`main` ref from which
   `.github/workflows/sync-skill-issue-plugin.yml` is dispatched.
2. Define the non-`main` base branch for its commit-specific automation branch.
3. Update the workflow and repository instructions so neither the dispatch nor
   review path can target protected `main`.
4. Validate the exact gitlink-only review and parent CLI checks before accepting
   a future dependency update.

The current pin is usable. This phase blocks the next ordinary dependency sync,
not current builds from the initialized pinned revision.

### Phase 3 — Reconcile public project state

**Status:** Ready after the repository-local integration route is settled.

1. Reconcile `src/data/siteData.ts` with the three accepted Claude Code/Codex
   artifacts and the current 21-run published aggregate.
2. Replace pre-release milestone language in `README.md` and project copy with
   the shipped beta and its precise remaining qualification limits.
3. Keep accepted-result values derived from the accepted artifacts rather than
   copying them into another ledger.
4. Run the complete website validation required by `AGENTS.md`.

### Phase 4 — Resume the blocked Cursor benchmark

**Status:** Blocked on Cursor capacity.

Prerequisite: the Cursor API usage limit resets on 2026-08-21, or Shannon
explicitly raises the limit earlier.

Resume these same recovery containers after resolving model identifiers from
the live Cursor catalog:

| Configuration  | Recovery containers                                     |
| -------------- | ------------------------------------------------------- |
| Cursor / Codex | `<chats>/chat-12`, `<chats>/chat-13`, `<chats>/chat-14` |
| Cursor / Fable | `<chats>/chat-15`, `<chats>/chat-16`, `<chats>/chat-36` |

`<chats>/chat-17` contains unrelated residue and is excluded. Follow the
[campaign orchestration procedure](skill-calling-evaluation-campaign/campaign-orchestration-prompt.md)
for same-container recreation, launch, monitoring, and cleanup. Accept and
publish results only through the existing artifact pipeline.

### Later phases

1. Verify current GitHub Release assets and Pages deployment when live remote
   verification is authorized; local source proves intended packaging and
   deployment paths, not present endpoint availability.
2. Decide whether to run the explicitly deferred runtime evaluation of the
   generated `repository-owner-finder` skill.
3. Compact historical research set by set. Remove orchestration maps,
   screenshots, or assignment corpora only after retained documents no longer
   depend on their unique evidence or links.
4. Revisit additional harness qualification only under a new explicit product
   decision and the retained two-gate qualification method.

## Blockers and Decisions

| Item                           | Type                  | Required resolution                                                                     |
| ------------------------------ | --------------------- | --------------------------------------------------------------------------------------- |
| Six Cursor evaluations         | External capacity     | Wait for the recorded reset or obtain Shannon's explicit billing decision               |
| Plugin-sync ref and base       | Architecture decision | Choose a non-`main` dispatch and review path before the next sync                       |
| Current remote alignment       | Evidence limit        | Fetch or query remotes only when live verification is requested                         |
| Release and Pages availability | Evidence limit        | Verify live endpoints before making a confirmed-current availability claim              |
| Historical task inventory      | Evidence limit        | Recent and pinned Skill Issue tasks were reviewed; older unlisted tasks may still exist |

## Document Disposition

| Document or set                               | Role                                                     |
| --------------------------------------------- | -------------------------------------------------------- |
| `plans/current-state.md`                      | Sole live repository progress and task owner             |
| `plans/README.md`                             | Status-free routing index                                |
| `campaign-orchestration-prompt.md`            | Current campaign procedure only                          |
| `evaluation-progress.md`                      | Historical redirect; no live status ownership            |
| `campaign-run-sheet.md`                       | Historical redirect; retired command and container map   |
| `plans/harness-setup.md`                      | Historical setup reference; source owns current behavior |
| `plans/website/reference-and-architecture.md` | Historical website architecture reference                |
| `research/project-state-recovery/**`          | Recovery evidence; not a live task owner                 |
| Other retained research                       | Historical or decision evidence at its domain home       |

No research assignments or screenshot sets are removed in this consolidation.
Several are plausible cleanup candidates, but retained syntheses still link to
assignment evidence and offline-history requirements have not been resolved.

## Next Return Point

Review the uncommitted recovery documents. If accepted, start Phase 1 by
choosing the surviving project-local lifecycle-skill discovery route, then
repair the three broken links and five stale showcase prompt paths together.

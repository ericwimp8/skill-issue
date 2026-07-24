# Accepted Evaluations and Evidence

## Assignment

**Goal:** Establish which evaluation and showcase work is concretely accepted or completed, which work remains partial or blocked, and which retained material is only qualification, historical, or planned evidence.

**Scope:** Local `evaluations/**` evidence; the accepted skill-calling artifact set; the linked skill-calling campaign tracker; retained showcase campaign status and evidence; the production result producer and website-ingestion path; relevant local Git history.

**Exclusions:** Broad plan inventory; general `research/**` cleanup; website presentation analysis beyond the accepted-artifact ingestion cross-check; Codex-history reconstruction; evaluation execution; source, output, or Git-state mutation.

## Sources

- Governing repository documents: `AGENTS.md:1-139` and `.repository-privacy.md:1-14`.
- Production evaluation types, persistence order, completion transition, and compact-result derivation: `cli/internal/evaluation/evaluation.go:143-179`, `cli/internal/evaluation/evaluation.go:446-499`, and `cli/internal/evaluation/evaluation.go:1287-1339`.
- Accepted-artifact ingestion and validation: `scripts/update-website-results.mjs:4-136`; operator contract: `README.md:258-273`; consuming types and configured cells: `src/data/evaluationData.ts:1-25`, `src/data/evaluationData.ts:82-146`, and `src/data/evaluationData.ts:152-211`; complete-cell aggregation: `src/data/analysisData.ts:48-71`.
- Accepted compact artifact directory: `evaluations/skill-calling/results/accepted/`. Retained run IDs, grouped by complete harness/model cell:
  - Claude Code / `claude-fable-5`: `daf204f8d746857d6cdea31f98c0abd0`, `28bbfaafbd2350de09f28ee62b6c647a`, `114e9a28c466d5f5de46a5560ac93def`.
  - Claude Code / `gpt-5.6-sol`: `e20e00601b9cf52124680fb0c0a937e5`, `a0f46b4ec92e74c09eba16abcd1121a7`, `12192ed7bb64e3cf049d712876b99bc5`.
  - OpenAI Codex / `gpt-5.6-sol`: `f80e38d1fe99e121f233428792c2f6a2`, `f885485d2718735c688e20c9cdcf8f34`, `84ffd0408208a60685e896edf240f3cc`.
  - Cursor / `composer-2.5`: `bfebefcbb497c53f386d388eb5b86854`, `4c15412e323ba9ecdfb665f045fe0406`, `aee68fe2c0796948d21e2d4fcf29d7af`.
  - Cursor / `cursor-grok-4.5-medium`: `be67f5c9726747e18374dfc2fef587ea`, `09d482d1ec542c92f36e3e4d235d4ad4`, `fd38e8865de744325280939db588d196`.
  - Pi / `openai-codex/gpt-5.6-sol`: `a13afc4d5d8682721c31b66efd561d9c`, `3d169365a52ae92b71d7da2d6209661b`, `10ad4d711e0a8d1be4eb62c70367b2b1`.
  - OpenCode / `openai/gpt-5.6-sol`: `72049527285a7fc9964e1b1e9b98e53e`, `6f29ec29552bd293487fd0fbcda570cd`, `f0233161faa6dac6e22542516eb6f308`.
- Representative accepted artifact records: Cursor/Grok `09d482d1ec542c92f36e3e4d235d4ad4` at `evaluations/skill-calling/results/accepted/09d482d1ec542c92f36e3e4d235d4ad4.json:1-8`; Claude/Fable `114e9a28c466d5f5de46a5560ac93def` at `evaluations/skill-calling/results/accepted/114e9a28c466d5f5de46a5560ac93def.json:1-8`; Codex emergency result `84ffd0408208a60685e896edf240f3cc` at `evaluations/skill-calling/results/accepted/84ffd0408208a60685e896edf240f3cc.json:1-8`.
- Reconciled Codex records: `f80e38d1fe99e121f233428792c2f6a2` at `evaluations/skill-calling/results/accepted/f80e38d1fe99e121f233428792c2f6a2.json:199-202` and `f885485d2718735c688e20c9cdcf8f34` at `evaluations/skill-calling/results/accepted/f885485d2718735c688e20c9cdcf8f34.json:199-202`.
- Skill-calling campaign state and blockers: `plans/skill-calling-evaluation-campaign/evaluation-progress.md:9-36`, `plans/skill-calling-evaluation-campaign/evaluation-progress.md:64-127`, and `plans/skill-calling-evaluation-campaign/evaluation-progress.md:129-160`.
- Real-harness smoke qualification: `evaluations/skill-calling/smoke/real-harness-smoke-report.md:5-18`, `evaluations/skill-calling/smoke/real-harness-smoke-report.md:28-70`, and `evaluations/skill-calling/smoke/real-harness-smoke-report.md:72-82`.
- Completed skill-system refinement campaign: `evaluations/skill-system-production-refinement/progress.md:15-43`, `evaluations/skill-system-production-refinement/progress.md:45-117`, and `evaluations/skill-system-production-refinement/progress.md:127-130`; completion audit and explicit runtime deferral: `evaluations/skill-system-production-refinement/final-audit.md:3-37`; end-to-end boundary: `evaluations/skill-system-production-refinement/end-to-end/audit.md:3-12`.
- Completed scenario-refinement campaign status files:
  - `evaluations/scenario-skill-refinement/code-implementation-discipline/status.md:3-19`
  - `evaluations/scenario-skill-refinement/code-testing-discipline/status.md:3-25`
  - `evaluations/scenario-skill-refinement/document-update-discipline/status.md:3-27`
  - `evaluations/scenario-skill-refinement/prompt-writing/status.md:3-18`
  - `evaluations/scenario-skill-refinement/skill-authoring-discipline/status.md:3-23`
  - `evaluations/scenario-skill-refinement/system-change-ownership/status.md:3-18`
  - `evaluations/scenario-skill-refinement/systematic-debugging/status.md:3-18`
- Completed showcase campaign status files:
  - `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/status.md:3-29`
  - `showcase-skills/api-change-impact-mapper/evaluation/api-change-impact-mapper/status.md:3-27`
  - `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/status.md:3-20`
  - `showcase-skills/ci-failure-triage/evaluation/ci-failure-triage/status.md:3-26`
  - `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/status.md:3-23`
  - `showcase-skills/environment-doctor/evaluation/environment-doctor/status.md:3-16`
  - `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/status.md:3-21`
  - `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/status.md:3-21`
  - `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/status.md:3-15`
  - `showcase-skills/safe-share-redactor/evaluation/safe-share-redactor/status.md:3-16`
- Relevant commits:
  - `1404e8e` (`feat: publish benchmark results and harden evaluation replay`) added 19 accepted compact artifacts, the schema-v2 producer/ingester changes, tracker updates, and the generated website collection.
  - `b64d9b5` (`data: publish completed Claude Code evaluations`) added accepted runs `12192ed7bb64e3cf049d712876b99bc5` and `e20e00601b9cf52124680fb0c0a937e5`, completing the Claude Code / Codex cell; this is the `build-week-2026` snapshot.
  - `8ee182f` (`chore: restore curated evaluation evidence`) is the latest commit touching the scenario-refinement evidence tree.
  - `d54d3b7` (`feat: align evaluation payloads and outcome reporting`) is the latest commit touching the skill-system production-refinement progress record.
  - `a7ee8ab` (`refactor: package skills as marketplace plugin`) is the latest commit touching the showcase tree and explains why some historical production-skill target paths are no longer live at their recorded locations.
- Read-only commands used: `rg --files`, `rg -n`, `find`, `sed -n`, `nl -ba`, `jq`, `cmp`, `shasum -a 256`, `git branch --show-current`, `git rev-parse HEAD`, `git status --short`, `git log`, and `git show`.

## Findings

### Finding 1: Twenty-one compact results are concretely accepted and form seven complete three-scenario cells

The repository contains 21 tracked schema-v2 compact artifacts. Every one names one of the three governed scenarios, and the set partitions into seven unique harness/model cells with all three scenarios present. The generated publication collection contains the same 21 records in the same canonical content after sorting.

The accepted totals are:

| Harness / model            | Accepted runs | Called | Missed | Unexpected |
| -------------------------- | ------------: | -----: | -----: | ---------: |
| Claude Code / Claude Fable |             3 |     17 |    120 |          3 |
| Claude Code / Codex Sol    |             3 |     19 |    118 |          7 |
| OpenAI Codex / Codex Sol   |             3 |    137 |      0 |         81 |
| Cursor / Composer          |             3 |     26 |    111 |          2 |
| Cursor / Grok              |             3 |    130 |      7 |         19 |
| Pi / Codex Sol             |             3 |     15 |    122 |         10 |
| OpenCode / Codex Sol       |             3 |     20 |    117 |          6 |

**Evidence:** Artifact schema and identities are present at lines 1-8 of every accepted JSON; representative records are cited in Sources. The ingestion command accepts only schema version 2, validates point fields, enforces filename/run-ID agreement, rejects duplicate run IDs and duplicate harness/model/scenario identities, sorts the collection, and writes `src/data/publishedWebsiteArtifacts.json` (`scripts/update-website-results.mjs:38-136`). The current generated file contains 21 records and is byte-equivalent at normalized JSON level to the sorted accepted directory. Complete-cell analysis requires all three scenario IDs (`src/data/analysisData.ts:48-71`). Commits `1404e8e` and `b64d9b5` account for all 21 accepted files.

**Implication:** The consolidated progress owner can mark seven benchmark configurations and 21 individual governed runs as accepted public evidence. These are completed cells, not illustrative chart values or smoke outputs.

### Finding 2: The skill-calling campaign is partial, inactive, and blocked at 21 of 27 runs

Two configured cells have no accepted run: Cursor / Claude Fable and Cursor / Codex Sol. All six runs started but stopped between turns 15 and 21 when Cursor's API usage limit was exhausted. The recorded operator decision was to wait for the 2026-08-21 cycle reset or raise the limit earlier. No campaign process is active.

**Evidence:** The tracker records 7/9 complete configuration suites, 21/27 complete runs, 0 running, 0 pending, 0 failed, and 6 blocked (`plans/skill-calling-evaluation-campaign/evaluation-progress.md:26-36`). The six blocked run records and their turn of failure are at lines 89-99; the blocker log names `ActionRequiredError`, the 2026-08-21 reset, and the operator-controlled billing alternative at lines 143-148. The accepted directory and generated collection contain Cursor records only for `composer-2.5` and `cursor-grok-4.5-medium`; no accepted file names a Fable or Codex Cursor model.

**Implication:** Campaign-wide completion is blocked, not active and not failed. The next benchmark action is dependency-ordered: restore Cursor access, recreate each failed container cleanly, run the three Cursor/Codex scenarios sequentially within that lane and the three Cursor/Fable scenarios sequentially within that lane, accept only tooling-complete outputs, then regenerate the published collection.

### Finding 3: The campaign tracker contains stale local details and needs reconciliation against the accepted set

The tracker summary is correct at 21/27 and 7/9, but several lower sections were not fully reconciled after later completions. The Cursor/Composer table row says 3/3 while its detailed heading says 0/3. The closing notes still say only four runs completed and all remaining runs were stopped, name Grok and Composer as stopped attempts, leave the campaign completion date blank, and label `chat-6` and `chat-7` incomplete even though accepted Composer and Grok runs now exist.

**Evidence:** The high-level table records Cursor/Composer 3/3 (`plans/skill-calling-evaluation-campaign/evaluation-progress.md:19`) and the detailed three runs are individually checked complete (`plans/skill-calling-evaluation-campaign/evaluation-progress.md:108-113`), while the configuration checkbox remains 0/3 at line 110. Stale closing statements are at lines 153-160. The accepted cell has three unique Composer artifacts and three unique Grok artifacts, and the generated collection matches them exactly.

**Implication:** Use the accepted directory plus the 21/27 high-level summary as the current status boundary. Before any campaign continuation, reconcile the Composer heading and closing notes so the progress owner cannot restart completed runs or treat accepted containers as incomplete.

### Finding 4: Accepted compact artifacts prove score projections, but not full run provenance or cleanup by themselves

The committed accepted files are compact `website.json`-shape projections. They omit reasoning, timestamps, evaluation ID, detailed expected/observed call identities, unattributed calls, transcript linkage, harness version, cleanup state, and failure diagnostics. Repository acceptance is editorially represented by placing a valid compact artifact in the accepted directory; the ingestion script does not independently verify full-run completion, cleanup, scenario membership in the governed suite, or a three-run cell.

**Evidence:** The production `Result` contains detailed identity, time, classification, and transcript fields (`cli/internal/evaluation/evaluation.go:143-161`), while `WebsiteResult` contains only schema, run, scenario, harness, model, total turns, and points (`cli/internal/evaluation/evaluation.go:163-179`). Production writes `result.json`, then `website.json`, then marks the private run complete and performs cleanup (`cli/internal/evaluation/evaluation.go:476-499`). The public retention contract explicitly leaves detailed results, events, transcripts, workspaces, and failed-run diagnostics outside the repository (`README.md:258-272`). The updater validates compact shape and uniqueness only (`scripts/update-website-results.mjs:38-127`).

**Implication:** The 21 accepted files are sufficient for public score publication and exact cell membership. Claims about tooling cleanliness, reasoning, harness versions, cleanup, and attempts depend on the authored tracker and external retained containers, not on repository-local machine-readable accepted artifacts. A future stronger acceptance manifest would need to retain those provenance fields explicitly.

### Finding 5: Two accepted Codex artifacts carry explicit reconciliation receipts

The gardening and community-archive OpenAI Codex results were recomputed from raw transcript evidence after attribution expanded to recognize Codex entrypoint reads and `printf` marker forms. Their accepted artifacts preserve this exceptional provenance; the emergency-preparedness result has no reconciliation field.

**Evidence:** Run `f80e38d1fe99e121f233428792c2f6a2` and run `f885485d2718735c688e20c9cdcf8f34` each record `basis: "raw transcript"` and the same recomputation reason (`evaluations/skill-calling/results/accepted/f80e38d1fe99e121f233428792c2f6a2.json:199-202`; `evaluations/skill-calling/results/accepted/f885485d2718735c688e20c9cdcf8f34.json:199-202`). The updater validates and preserves optional reconciliation receipts (`scripts/update-website-results.mjs:27-35`, `scripts/update-website-results.mjs:97-102`).

**Implication:** Treat all three Codex runs as accepted, while preserving the two receipts whenever the collection is regenerated. They are corrected accepted evidence, not unqualified original projections.

### Finding 6: Real-harness smoke work is completed qualification evidence, not accepted benchmark evidence

Codex, Cursor, Claude Code, and Pi each completed built-in and custom two-turn smoke routes, with result and website artifacts produced locally and cleanup checked. This established launcher and integration viability. The report explicitly excludes governed 30-turn evaluation and publication acceptance.

**Evidence:** Qualified launcher outcomes are recorded at `evaluations/skill-calling/smoke/real-harness-smoke-report.md:9-18`; built-in and custom route results at lines 28-68; cleanup at lines 72-78. The boundary says custom smoke evidence is local and not publication evidence without separate review (`evaluations/skill-calling/smoke/real-harness-smoke-report.md:80-82`).

**Implication:** Mark the smoke campaign completed as harness qualification only. Its final statement that governed 30-turn evaluations were deferred is historical and superseded for the seven accepted cells, but it remains accurate that smoke artifacts themselves were never accepted as benchmark results.

### Finding 7: The skill-system production-refinement campaign is complete at its authorized boundary, with generated-skill runtime proof intentionally deferred

The 21-iteration production-refinement campaign completed all phases, passed its evaluator, Skill Intake, and Skill Generation targets, and reached a complete Intake-to-Generation-to-Evaluation handoff. It deliberately stopped before runtime evaluation of the additional generated end-to-end skill; no runtime pass is claimed.

**Evidence:** `progress.md` records phase complete, no current target or next action, and no blockers (`evaluations/skill-system-production-refinement/progress.md:27-43`); all dependency-ordered phases are checked (`evaluations/skill-system-production-refinement/progress.md:45-117`). The final audit passes all campaign and parent-plan criteria while repeatedly preserving the runtime deferral (`evaluations/skill-system-production-refinement/final-audit.md:3-37`). The end-to-end audit limits completion to the evaluation handoff and structural validation (`evaluations/skill-system-production-refinement/end-to-end/audit.md:3-12`).

**Implication:** Mark this campaign completed, not partial. Track runtime evaluation of the generated `repository-owner-finder` skill as a separate deferred opportunity; do not convert the completed campaign into active work or infer runtime behavior.

### Finding 8: Historical production-target statuses do not prove the current packaged plugin at their old paths

The production-refinement records point to `skills/skill-evaluation-and-refinement`, `skills/skill-intake`, and `skills/skill-generation`. Those paths are absent in the current checkout after the plugin packaging refactor, while retained campaign evidence and hashes remain. The copied Dictate Plan target is also explicitly retired.

**Evidence:** The campaign's starting position says the Dictate Plan live copy was retired after its evidence was preserved (`evaluations/skill-system-production-refinement/progress.md:3-9`). The current checkout has no files at the three recorded `skills/**` target paths or `test-targets/skills/dictate-plan/SKILL.md`. Commit `a7ee8ab` moved product skills into the marketplace-plugin dependency model.

**Implication:** The campaign proves the recorded target versions and workflow at that historical boundary. Current plugin acceptance requires following the dependency's own pinned revision and validation evidence; it should not be inferred solely from these old-path status files.

### Finding 9: Seven scenario-skill-refinement campaigns are completed, and their recorded target hashes match the retained targets

The code implementation, code testing, document update, prompt writing, skill authoring, system change ownership, and systematic debugging campaigns all record passed description/body state as applicable, concluded loops, no next action, and retained evidence. Each status-recorded `SKILL.md` hash matches its current retained campaign target.

**Evidence:** The seven status files are cited in Sources. Their terminal state lines are respectively `code-implementation-discipline/status.md:18-19`, `code-testing-discipline/status.md:24-25`, `document-update-discipline/status.md:26-27`, `prompt-writing/status.md:17-18`, `skill-authoring-discipline/status.md:22-23`, `system-change-ownership/status.md:17-18`, and `systematic-debugging/status.md:17-18`. Read-only SHA-256 comparison matched all seven recorded hashes to `evaluations/scenario-skill-refinement/*/skill/SKILL.md`.

**Implication:** Mark all seven campaigns completed. Historical unsuccessful cycles or invalid preflights are retained diagnostic evidence and do not reopen the final passing targets.

### Finding 10: All ten showcase evaluation campaigns are completed against their current retained target hashes, with scoped caveats

All ten showcase status files end in `passed`, `Campaign state: passed`, or equivalent and record no next action. Read-only SHA-256 comparison matched every status-declared target hash to the current showcase `SKILL.md`.

The important boundaries are:

- Accessibility First Pass passed through fresh-agent selection/read evidence, but its native CLI activation attempt was blocked by a read-only Codex state database and denied escalation; no CLI activation result is claimed.
- Dependency Upgrade Planner passed after replacement evidence; an earlier capacity stop and artifact-free interrupted execution are retained as environment evidence.
- Safe Share Redactor correctly marks description selection not applicable because implicit invocation is prohibited, and passed its body campaign after two historical body failures.
- Environment Doctor passed after one historical body failure; the other showcase campaigns passed their recorded applicable loops without a remaining action.

**Evidence:** Terminal states and caveats are explicit in `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/status.md:19-29`, `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/status.md:17-23`, `showcase-skills/safe-share-redactor/evaluation/safe-share-redactor/status.md:7-16`, and `showcase-skills/environment-doctor/evaluation/environment-doctor/status.md:7-16`. The remaining six terminal passing status records are cited in Sources. Read-only hash comparison matched all ten declared hashes.

**Implication:** Mark all ten showcase campaigns completed at their stated evidence surfaces. Do not translate Accessibility's fresh-agent evidence into a native-CLI activation claim, and do not treat explicit-only description gates as missing work.

### Finding 11: Evaluation recovery has one benchmark blocker and two independent deferred opportunities

The only active dependency blocker in retained evaluation state is Cursor access for six governed runs. Two other items are deferred rather than blockers: runtime evaluation of the generated end-to-end skill, and any separate review that would promote local smoke artifacts to publication evidence.

**Evidence:** The campaign tracker records six blocked runs and no running or pending runs (`plans/skill-calling-evaluation-campaign/evaluation-progress.md:26-36`). The Cursor access dependency is explicit at lines 143-148. Generated-skill runtime proof is explicitly deferred without a pass claim (`evaluations/skill-system-production-refinement/final-audit.md:5-7`, `evaluations/skill-system-production-refinement/final-audit.md:26-37`). Smoke publication requires separate review (`evaluations/skill-calling/smoke/real-harness-smoke-report.md:80-82`).

**Implication:** The consolidated progress owner should order future evaluation work as: first reconcile the tracker; then wait for or authorize Cursor access; then rerun and accept the six blocked governed evaluations; then regenerate publication data. Runtime proof for the generated skill and publication review for smoke evidence are separate optional work streams and should remain explicitly unclaimed until requested.

## Notes

- The accepted repository evidence is intentionally compact. Local-only inspection could not validate the external `<chats>/chat-*` containers, detailed `result.json` files, raw transcripts, or failed `failure.json` diagnostics named by the campaign tracker.
- No `failure.json` is retained under tracked `evaluations/**` or `showcase-skills/**`; blocker diagnosis is therefore authored in the tracker rather than independently inspectable from a repository-local failure record.
- `plans/skill-calling-evaluation-campaign/evaluation-progress.md:158` leaves “Campaign completed” blank, which is consistent with the unresolved 6/27 blocked runs.
- The current tracked evidence paths inspected for this assignment were clean in `git status`; unrelated `AGENTS.md` and `research/project-state-recovery/` changes were preserved.
- Useful search terms: `results/accepted`, `schema_version: 2`, `reconciliation`, `ActionRequiredError`, `Campaign state`, `runtime evaluation remains deferred`, `results:update`.

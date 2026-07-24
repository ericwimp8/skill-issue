# Historical Research and Cleanup

## Assignment

**Goal:** Classify retained repository research and investigation documents by durable value and cleanup disposition so later consolidation preserves necessary historical evidence without leaving research files as duplicate live status owners.

**Scope:** Inspect `research/**` broadly; deep-dive documents containing progress state, accepted decisions, dependency order, blockers, or unique historical evidence; check repository backlinks, current documentation/source owners, and relevant Git history; classify documents or document sets as preserve at semantic home, preserve as history, convert to concise historical redirect, safe deletion candidate, or unsupported pending inspection.

**Exclusions:** Primary `plans/**` tracker analysis, production-source lifecycle tracing, accepted evaluation artifact analysis, and Codex-history reconstruction. Plans and source were inspected only to validate backlinks, current ownership, and whether research status claims had drifted. No existing file was edited, moved, or deleted.

## Sources

- Repository guidance: `AGENTS.md:95-99` establishes that research documents are temporary unless they retain lasting value; `.repository-privacy.md:1-13` requires repository-relative, identity-safe public content.
- Active run boundary: `research/project-state-recovery/research-map.md:1-12,60-79,81-99`.
- Full inventory and metadata: every file returned by `rg --files research`; Markdown line counts, titles, and last-touch commits; tracked/untracked state from `git status --short research`; image-set sizes and file names under `research/video-reference-screenshots/`, `research/website-refactor/`, and `research/website-two-arm-navigation/`.
- Backlink searches: exact `research/` path search outside `research/**`; per-directory and per-basename searches outside `research/**`, `.git/**`, and `output/**`. The only concrete retained-document backlink found was `plans/website/reference-and-architecture.md:84-90` to the free-hosting synthesis.
- Current status cross-checks: `plans/skill-calling-evaluation-campaign/evaluation-progress.md:1-36,51-62`; `src/data/siteData.ts:248-299`; `src/data/publishedWebsiteArtifacts.json:1`; `scripts/update-website-results.mjs:4-7,38-66,68-136`; `.github/workflows/pages.yml:1-55`; `cli/README.md:21-67,160-180`; `cli/internal/harness/harness.go:15-19,45-69`; `cli/scripts/build-cross-platform.sh:16-75`; `README.md:12-38`; `INSTALL.md:178-224`.
- Benchmark analysis set: `research/benchmark-analysis-page/analysis-page-idea-notes.md:1-76`; `research/benchmark-analysis-page/benchmark-analysis-page-deep-research.md:3-19,428-461`; its twelve assignments and `research-map.md`.
- Distribution and hosting sets: `research/cross-platform-cli-distribution/cross-platform-cli-distribution-report.md:3-17,43-81,125-151,192-207`; `research/deep-research/free-website-hosting-cli-downloads/free-website-hosting-cli-downloads-deep-research.md`; its six assignments and `research-map.md`.
- Harness architecture and invocation sets: `research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md:1-28,85-116,150-183`; its nine retained assignments and `research-map.md`; `research/deep-research/manual-skill-invocation-eight-harnesses/manual-skill-invocation-eight-harnesses-deep-research.md:3-109`; its eight assignments and `research-map.md`; `research/harness-subagent-launch/harness-subagent-launch-reference.md:1-57` and its eight assignments.
- Harness qualification set: `research/harness-portability-qualification/remaining-harness-portability-qualification-audit.md:1-18,44-117,137-266`; `research/harness-portability-qualification/opencode-local-configuration-qualification.md:1-35,55-93,95-181`; all retained assignments and `research-map.md`.
- Website and methodology sets: `research/implementation-research/website-chart-migration/website-chart-migration-implementation-research.md:1-23,25-110,157-205,233-269`; all assignments and `research-map.md`; `research/testing-methodology/local-system-research.md:1-38,116-200,241-284`; `research/testing-methodology/external-methodology-page-research.md`; all assignments and `00-external-research-map.md`; `research/website-two-arm-navigation/reference-findings.md:1-23`; `src/App.tsx:11-23,32-50,75-95,97-135`.
- Video-production history: both final syntheses, all assignments, and maps under `research/deep-research/codex-friendly-youtube-video-tools/` and `research/deep-research/natural-ai-tts-for-remotion-demo/`; `research/deep-research/natural-ai-tts-for-remotion-demo/natural-ai-tts-for-remotion-demo-deep-research.md:3-15`; `README.md:12-38`.
- Relevant commits: `315603b` introduced most initial research; `d54d3b7` aligned chart research and implementation; `b94697a` published methodology research and website-refactor evidence; `bb77a92` expanded harness qualification and two-arm navigation; `6dec125` added benchmark-page research; `5040bfa` retired Kilo support and updated harness research; `1404e8e` published benchmark results and analysis notes; `68f73ce` implemented the public analysis page; `6bf4312` retired completed documents.

## Findings

### Finding 1: Research is evidence, not a current-status authority

The durable repository rule is to retain research only when it provides lasting evidence. Production source, current accepted artifacts, and the designated campaign tracker now own current behavior and progress. Research documents should therefore be read as dated evidence unless a current document explicitly backlinks to them.

**Evidence:** `AGENTS.md:95-99` makes working research temporary by default. The active recovery map likewise requires one authoritative current owner and separate historical/redirect classifications (`research/project-state-recovery/research-map.md:28-35,73-79`). Current CLI behavior is stated in `cli/README.md:21-67`; current campaign execution is stated in `plans/skill-calling-evaluation-campaign/evaluation-progress.md:1-36`. The backlink scan found only one durable consumer: `plans/website/reference-and-architecture.md:84-90`.

**Implication:** Consolidation should add a general dated-evidence boundary rather than trying to keep every retained report current. Any report retained for history must stop presenting embedded progress, blockers, or “current implementation” prose as live authority.

### Finding 2: Three research documents contain materially stale live campaign or publication state

`analysis-page-idea-notes.md`, `local-system-research.md`, and the website-chart migration synthesis contain status claims that current repository evidence has superseded. These files are valuable as history, but unsafe as live progress owners.

**Evidence:** The analysis notes say 19 accepted runs and two in-flight Claude Code–Codex runs (`research/benchmark-analysis-page/analysis-page-idea-notes.md:3-7,20-25,70-76`). The current campaign tracker records Claude Code–Codex at 3/3 and the campaign at 21/27 (`plans/skill-calling-evaluation-campaign/evaluation-progress.md:11-22,26-36`). The local methodology research says the campaign is 10 configurations/30 runs, 0/30 complete, with an empty accepted website collection (`research/testing-methodology/local-system-research.md:5-22,157-175,265-280`); the current tracker defines nine configurations/27 runs and 21/27 complete, while the checked-in accepted-data importer requires and publishes schema-v2 accepted artifacts (`plans/skill-calling-evaluation-campaign/evaluation-progress.md:5-22`; `scripts/update-website-results.mjs:38-75,108-136`). The chart synthesis says real results, OpenCode support, and Pages remain blocked (`research/implementation-research/website-chart-migration/website-chart-migration-implementation-research.md:7-13,247-258`), while current source supports OpenCode and a Pages workflow exists (`cli/README.md:23-35,172-180`; `.github/workflows/pages.yml:1-55`).

**Implication:** Convert these three documents to concise historical redirects. A redirect should identify the inspected date/commit, retain only decisions or defects still useful for history, and route current readers to the campaign tracker, accepted-data pipeline, website source, CLI README, and current analysis data. Their detailed old progress tables and “remaining blockers” should not survive as parallel status.

### Finding 3: The chart-migration assignment corpus cannot yet be declared safe to delete as a whole

The chart synthesis is a strong redirect candidate, and its research map is safe process cleanup, but the nine assignment documents contain detailed source traces, old schema defects, and design rationale that were not all independently reproduced elsewhere. Current source proves implementation ownership; it does not by itself prove every historical trace is disposable.

**Evidence:** The synthesis explicitly says the migration was implemented and production source owns current behavior (`research/implementation-research/website-chart-migration/website-chart-migration-implementation-research.md:1-4`), yet it links assignments throughout its production path and close evidence (`:23-29,260-269`). Its research map is orchestration metadata whose declared output is the synthesis (`research/implementation-research/website-chart-migration/research-map.md:3-11,21-54`). Current import and publication owners are concrete (`scripts/update-website-results.mjs:4-7,38-136`; `.github/workflows/pages.yml:1-55`).

**Implication:** Delete `research-map.md` after the active consolidation work no longer needs it. Convert the synthesis to a redirect. Classify the assignments as **unsupported pending inspection** for deletion until a compaction pass checks whether the historical schema defect, provenance design, and owner traces are preserved in Git history or a concise retained decision record.

### Finding 4: Benchmark-reporting and methodology research should be consolidated as dated historical source basis

The benchmark-analysis and external-methodology corpora contain durable source-backed editorial and statistical reasoning. They overlap heavily in HELM, METR, OpenAI, NIST, benchmark-validity, uncertainty, and evidence-interface themes. Their final syntheses have lasting value, while their separate maps are completed orchestration records.

**Evidence:** The benchmark synthesis defines bounded claim language, combination-first comparison, local caveats, and source traceability (`research/benchmark-analysis-page/benchmark-analysis-page-deep-research.md:3-19,428-461`). The external methodology synthesis covers the adjacent methodology-page content and disclosure contract, and its map identifies the same broad evaluation-reporting ecosystem (`research/testing-methodology/00-external-research-map.md:3-23,34-56`). The benchmark map and external methodology map exist to dispatch assignments and name their final aggregation targets (`research/benchmark-analysis-page/research-map.md:3-13,33-73`; `research/testing-methodology/00-external-research-map.md:25-70`). Current website copy and pages now own public analysis and methodology behavior (`src/data/siteData.ts:137-213,248-299`; `src/pages/ResultsAnalysisPage.tsx`; `src/pages/MethodologyPage.tsx`).

**Implication:** Preserve both final syntheses and their evidence-bearing assignments as **history** until a deliberate consolidation preserves distinct analysis-page and methodology-page conclusions. Remove completed research maps. After consolidation, one dated “evaluation reporting research basis” document could replace duplicate high-level framing, while current website source remains authoritative.

### Finding 5: `analysis-page-idea-notes.md` is a redirect candidate rather than necessary evidence

The idea notes mix useful drafting observations with stale in-flight status and numbers. The lasting claims, correction disclosure, limitations, and pairing thesis have been incorporated into current website copy and data owners.

**Evidence:** The notes contain the pairing thesis, correction story, missing cells, and limitations (`research/benchmark-analysis-page/analysis-page-idea-notes.md:9-18,56-76`). Current public copy owns the pairing thesis, correction disclosure, descriptive limitations, and unpublished cells (`src/data/siteData.ts:248-299`). Git history retains the drafting-to-publication transition through `1404e8e` and `68f73ce`.

**Implication:** Convert the notes to a short historical redirect or delete them after recording any anecdote intentionally retained for project history. They should not retain current counts or next-work language.

### Finding 6: Cross-platform distribution research is durable decision history, but not current release authority

The distribution report explains why Go, six OS/architecture artifacts, archives, checksums, and GitHub Releases were chosen. Those decisions remain visible in implementation. Its GoReleaser, signing, notarization, package-manager, and updater recommendations were not all adopted and must remain historical options.

**Evidence:** The report recommends Go, GoReleaser, signed artifacts, six archives, checksums, package managers, and later update behavior (`research/cross-platform-cli-distribution/cross-platform-cli-distribution-report.md:3-17,43-81,125-151,192-207`). The current build script directly cross-compiles six pure-Go targets and generates archives/checksums without GoReleaser or signing (`cli/scripts/build-cross-platform.sh:16-75`). `INSTALL.md:30-49,86-105,125` documents current release assets and checksum verification.

**Implication:** Preserve the report as **history**, preferably compacted into a dated distribution ADR that clearly separates adopted decisions from unimplemented alternatives. Current build scripts, release workflow, and installation guide own operational truth.

### Finding 7: Free-hosting research is the one research synthesis currently preserved at semantic home by a backlink

The free-hosting synthesis directly supports the still-retained website architecture decision. Deleting or redirecting it would break a deliberate source-basis link.

**Evidence:** `plans/website/reference-and-architecture.md:84-90` links the synthesis as existing repository research. The synthesis selects GitHub Pages plus Releases and preserves conditional alternatives and unresolved Firebase limits. Current repository structure implements Pages and release downloads (`.github/workflows/pages.yml:1-55`; `INSTALL.md:30-49,86-105,125`).

**Implication:** Preserve `free-website-hosting-cli-downloads-deep-research.md` **at its semantic home** while the backlink remains. Preserve its assignments because the synthesis links them. Its completed `research-map.md` is a safe deletion candidate after consolidation because it only records orchestration and the already-produced final target (`research/deep-research/free-website-hosting-cli-downloads/research-map.md:53-64`).

### Finding 8: Direct-installation architecture is valuable history with an explicit supersession boundary

This corpus contains durable per-harness native roots, collision rules, metadata boundaries, and trust constraints. Its leading transactional lifecycle recommendation is superseded. The existing status banner correctly points current readers away from the old owner.

**Evidence:** The synthesis explicitly says current lifecycle and ownership live in the completion plan and CLI source, and that receipt/backup/rollback/mutable-state recommendations are superseded (`research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md:1-4`). The remaining matrix and path sections preserve exact researched host distinctions (`:13-28,30-96`). Current CLI support is narrower and disposable: five harnesses, no receipt, backup, rollback inventory, or platform application-state directory (`cli/README.md:23,61-67`; `cli/internal/harness/harness.go:15-19,66-69`).

**Implication:** Preserve the synthesis and assignments as **history**. Keep the status banner. Remove the completed research map. Do not use the eight-harness matrix to imply current CLI support; current harness registry and CLI README own that claim.

### Finding 9: Manual invocation and sub-agent launch references are useful dated ecosystem history, not product-support matrices

These two small corpora preserve harness-specific interaction syntax and delegation surfaces that are not owned elsewhere in the repository. Their coverage exceeds the current five-harness CLI product boundary and their external behavior can drift.

**Evidence:** The manual-invocation synthesis explicitly rejects a universal command and distinguishes `$dictate-plan`, `/dictate-plan`, `/skill:dictate-plan`, and prose activation (`research/deep-research/manual-skill-invocation-eight-harnesses/manual-skill-invocation-eight-harnesses-deep-research.md:85-109`). The sub-agent reference preserves eight distinct launch mechanisms (`research/harness-subagent-launch/harness-subagent-launch-reference.md:3-57`). Current installation documentation supports only five harness identifiers and gives generic explicit-invocation guidance (`INSTALL.md:189-224`).

**Implication:** Preserve both corpora as **history** with a visible research date/version caveat. Remove their completed research maps where present. They should not be presented as current Skill Issue support or as proof of the caller's local installation/discovery state.

### Finding 10: OpenCode qualification is unique historical evidence and should remain preserved

The OpenCode qualification records actual local installation, OAuth, model, permission, multi-turn, interruption, cleanup, and session-deletion evidence. This is stronger than a documentation survey and directly explains current pinned OpenCode behavior.

**Evidence:** The qualification records the exact `1.18.4` installation path and package failure (`research/harness-portability-qualification/opencode-local-configuration-qualification.md:29-35`), the private XDG and deny-first configuration (`:55-78`), successful native ChatGPT OAuth/model proof (`:80-93`), canonical/custom two-turn routes and permission denial (`:95-141`), cleanup proof (`:143-160`), and the resulting adapter prerequisites (`:162-181`). Current source pins OpenCode `1.18.4` and implements the same runtime boundary (`cli/internal/harness/harness.go:68`; `cli/README.md:172-180`).

**Implication:** Preserve `opencode-local-configuration-qualification.md` as **history** at its current research home. It is necessary evidence for why the adapter is pinned and how qualification was established. It must remain dated and must not imply cross-platform qualification.

### Finding 11: The broader portability audit has durable gate methodology but a superseded candidate decision

The audit's two-gate qualification contract and twelve-step lifecycle remain reusable evidence. Its top-level recommendation that OpenCode is unimplemented is superseded by the qualification document and current source.

**Evidence:** The audit says none of the candidates passed both gates and OpenCode should remain unimplemented (`research/harness-portability-qualification/remaining-harness-portability-qualification-audit.md:7-18`). It separately defines durable Gate 1/Gate 2 standards, decision vocabulary, and a twelve-step process (`:44-117`). Current CLI source supports and pins OpenCode (`cli/README.md:23,172-180`; `cli/internal/harness/harness.go:68`).

**Implication:** Preserve the audit and its assignments as **history**, but add or consolidate into a concise supersession notice identifying OpenCode's later qualification and implementation. Remove its completed research map. Do not delete the assignment corpus until the unique access-route and rejected-candidate evidence has been compacted.

### Finding 12: Video/TTS research preserves real build history; orchestration maps do not

The final video and TTS syntheses explain the decision space that led to the published demo, and the TTS synthesis selected the Speechify API actually named in the README. The maps are completed run mechanics. Raw assignments remain linked evidence and should not be deleted without a citation-compaction pass.

**Evidence:** `README.md:12-38` links the public demo, names Remotion and Speechify, and explicitly says the research remains in the repository. The TTS synthesis selects SpeechifyAI `simba-3.2` (`research/deep-research/natural-ai-tts-for-remotion-demo/natural-ai-tts-for-remotion-demo-deep-research.md:3-15`). The video synthesis treats Remotion as the higher-polish code path and preserves the competition/time/rights reasoning. Each synthesis links its assignments, while each map only defines dispatch and aggregation.

**Implication:** Preserve the two final syntheses as **history**. Preserve assignments until the syntheses are rewritten to carry enough direct source evidence without internal assignment links. Delete the completed research maps after consolidation. This keeps the README's build-history claim true without preserving duplicate orchestration metadata indefinitely.

### Finding 13: Website-design and QA screenshot sets are safe deletion candidates

The website-refactor and two-arm-navigation screenshots have no backlinks, no manifest, and no current execution role. The current React/CSS source and Git history own the implemented design. The two-arm findings file contains decisions but an empty findings section.

**Evidence:** `research/website-two-arm-navigation/reference-findings.md:7-23` has no recorded findings and only short decisions. Current routing and two-arm selection live in `src/App.tsx:11-23,32-50,75-95,97-135`. Exact-path and basename searches found no references to either screenshot set outside `research/**`. Git history associates them with implementation commits `e1f7e10`, `b94697a`, and `bb77a92`.

**Implication:** `research/website-refactor/screenshots/**`, `research/website-two-arm-navigation/screenshots/**`, and `research/website-two-arm-navigation/reference-findings.md` are **safe deletion candidates** once the active recovery synthesis no longer needs them. Git retains the historical snapshot; current source remains the semantic owner.

### Finding 14: Video-reference screenshots are conditionally deletable, with the public demo as the stronger history

The 22 video-reference screenshots have descriptive filenames but no manifest or backlinks. They form a visual snapshot of the site used for video production, while the published YouTube demo is the more durable record.

**Evidence:** The set occupies about 3.1 MB and has no exact-path or basename backlinks outside `research/**`. `README.md:12` links the public demo. No Remotion/video source project or screenshot-consuming build pipeline exists in the repository; only `README.md:31` names the production tools.

**Implication:** Classify `research/video-reference-screenshots/**` as a **safe deletion candidate if the public video is accepted as the authoritative historical record**. If offline reconstruction of the exact video-era website is a requirement, preserve the set as history and add a short manifest; that requirement is not established in the inspected repository.

### Finding 15: Completed research maps are the broadest low-risk cleanup class

All completed maps describe researcher budgets, waves, assignment outputs, and final aggregation targets. Once their syntheses exist, they neither own decisions nor provide unique source evidence.

**Evidence:** Representative maps explicitly name the produced final targets and aggregation mechanics (`research/benchmark-analysis-page/research-map.md:3-13,33-73`; `research/deep-research/free-website-hosting-cli-downloads/research-map.md:53-64`; `research/implementation-research/website-chart-migration/research-map.md:3-11,21-54`; `research/testing-methodology/00-external-research-map.md:25-70`). The active recovery map is different: it still governs the current eight-assignment run (`research/project-state-recovery/research-map.md:36-79,89-99`).

**Implication:** Completed `research-map.md` files and `00-external-research-map.md` are **safe deletion candidates** after the recovery aggregation no longer cites them. Preserve `research/project-state-recovery/research-map.md` at its semantic home until this run is complete; classify that active set separately at closeout.

### Finding 16: Raw assignment corpora need set-specific compaction, not blanket deletion

Assignment documents are internally linked by most syntheses and often contain evidence details, source versions, dead ends, and validation not repeated in the final report. A blanket “delete all assignments” decision is unsupported.

**Evidence:** Benchmark, hosting, harness architecture, manual invocation, portability, video, TTS, chart, and methodology syntheses repeatedly link their assignment folders. The benchmark synthesis explicitly describes complementary roles for all twelve assignments (`research/benchmark-analysis-page/benchmark-analysis-page-deep-research.md:428-439`). The OpenCode qualification and portability audit demonstrate why runtime evidence can materially exceed final recommendation prose.

**Implication:** Preserve assignment corpora as **history** while their parent syntheses link them. A later compaction pass may delete a corpus only after rewriting the retained synthesis or ADR to preserve source links, versions, conflicts, validation, and important dead ends. Chart-migration assignments remain **unsupported pending inspection** for deletion; the same caution applies to any assignment set whose final document is converted to a redirect.

### Finding 17: Untracked operating-system metadata has no research value

Several `.DS_Store` files exist locally under `research/`, including otherwise empty directories, but none is tracked.

**Evidence:** `git ls-files 'research/**/.DS_Store' 'research/.DS_Store'` returned no paths. `research/audits/` contained only an untracked `.DS_Store`; other local `.DS_Store` files appeared in benchmark, deep-research, qualification, implementation-research, methodology, and website-refactor directories.

**Implication:** These files are safe local deletion candidates and should remain ignored. They are not part of the retained research evidence or any consolidation decision.

### Finding 18: Recommended disposition by document set

The following classification preserves unique evidence while eliminating duplicate live ownership:

| Document set                                                                       | Disposition                                   | Reason                                                                                         |
| ---------------------------------------------------------------------------------- | --------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `research/project-state-recovery/**`                                               | Preserve at semantic home while active        | Current run contract and assignment evidence; classify at closeout.                            |
| `benchmark-analysis-page/benchmark-analysis-page-deep-research.md` and assignments | Preserve as history                           | Durable reporting/claim-discipline evidence; internally linked.                                |
| `benchmark-analysis-page/analysis-page-idea-notes.md`                              | Convert to concise historical redirect        | Stale run status; current analysis source owns conclusions.                                    |
| Completed benchmark map                                                            | Safe deletion candidate                       | Orchestration metadata only.                                                                   |
| `cross-platform-cli-distribution-report.md`                                        | Preserve as history or compact ADR            | Adopted language/artifact rationale plus unimplemented alternatives.                           |
| Free-hosting synthesis and assignments                                             | Preserve at semantic home                     | Only research document with a current backlink.                                                |
| Completed free-hosting map                                                         | Safe deletion candidate                       | Orchestration metadata only.                                                                   |
| Direct-installation synthesis and assignments                                      | Preserve as history                           | Exact harness/path evidence; existing supersession banner is correct.                          |
| Manual-invocation synthesis and assignments                                        | Preserve as history                           | Unique cross-harness syntax reference; date-bound.                                             |
| Harness-subagent reference and assignments                                         | Preserve as history                           | Unique cross-harness delegation reference; date-bound.                                         |
| Harness-portability audit, OpenCode qualification, and assignments                 | Preserve as history                           | Unique gate and live qualification evidence; audit needs supersession notice.                  |
| Completed harness maps                                                             | Safe deletion candidates                      | Orchestration metadata only.                                                                   |
| Website-chart synthesis                                                            | Convert to concise historical redirect        | Implemented; internal current-state claims are stale.                                          |
| Website-chart assignments                                                          | Unsupported pending inspection                | Detailed old source/provenance traces may be unique.                                           |
| Website-chart map                                                                  | Safe deletion candidate                       | Orchestration metadata only.                                                                   |
| External methodology synthesis and assignments                                     | Preserve as history                           | Durable methodology/disclosure source basis.                                                   |
| Local methodology synthesis                                                        | Convert to concise historical redirect        | Strong trace, but campaign/publication status is stale.                                        |
| Local methodology assignment                                                       | Preserve as history until redirect compaction | Detailed production inspection evidence.                                                       |
| External methodology map                                                           | Safe deletion candidate                       | Orchestration metadata only.                                                                   |
| Video and TTS syntheses                                                            | Preserve as history                           | Direct evidence for the README's Remotion/Speechify build account.                             |
| Video and TTS assignments                                                          | Preserve as history pending compaction        | Parent syntheses link them.                                                                    |
| Video and TTS maps                                                                 | Safe deletion candidates                      | Orchestration metadata only.                                                                   |
| Website-refactor screenshots                                                       | Safe deletion candidate                       | No backlink or runtime role; current source/Git own design.                                    |
| Two-arm findings and screenshots                                                   | Safe deletion candidate                       | Decisions implemented; findings empty; no backlinks.                                           |
| Video-reference screenshots                                                        | Conditional safe deletion candidate           | Public demo is stronger history; preserve only if offline snapshot reconstruction is required. |
| Untracked `.DS_Store` files and empty `research/audits/` shell                     | Safe local deletion candidate                 | No research content or tracked evidence.                                                       |

**Evidence:** The set-specific evidence is established in Findings 2-17. No category above treats current source behavior as established by a test or by an old research recommendation.

**Implication:** The safest cleanup order is: remove untracked metadata; remove completed research maps; remove unreferenced website-design screenshots and the empty two-arm note; convert stale live-status documents into redirects; then perform set-specific source-evidence compaction before deleting any assignment corpus.

## Notes

- The research corpus is unusually recent: retained research commits span `315603b` through `1404e8e` on 2026-07-20 to 2026-07-22. “Historical” here means superseded ownership, not age.
- Git history preserves deleted tracked content, but repository consumers should not be required to excavate Git for a still-important architectural or qualification decision. This is why OpenCode qualification, distribution rationale, and source-backed reporting research remain preservation candidates.
- External product documentation in the harness, hosting, benchmark, video, and TTS corpora can drift. Preserved documents should carry their inspection date/version and require fresh verification before operational reuse.
- The public README says most research remains in the repository (`README.md:38`). Large-scale deletion should either preserve enough final syntheses for that statement to remain accurate or update the public wording at its semantic owner in a separate task.
- No deletion or rewrite was performed. The active aggregator should reconcile these dispositions with the production, Git, plan, evaluation, public-documentation, plugin, and Codex-history assignments before authorizing cleanup.

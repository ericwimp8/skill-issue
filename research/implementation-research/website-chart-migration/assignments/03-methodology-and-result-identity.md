# Methodology And Result Identity

## Assignment

- **Goal:** Establish the governed skill-calling evaluation method and the result identity needed to group chart data without widening the evidence claims.
- **Scope:** Current evaluation authorities, governed scenario and answer-sheet assets, production replay and result derivation, harness/model/reasoning/run/scenario identity, campaign repetition expectations, Work Block 3 acceptance, and the boundary between mock, smoke, local custom, and qualified campaign evidence.
- **Exclusions:** Chart rendering and layout implementation, external chart-library research, campaign execution, source edits, and any conclusion about model or harness performance.

## Sources

- `plans/skill-issue-project-completion/document-authority-and-update-map.md:12-51,79-108`
- `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:1-103`
- `plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:1-128`
- `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:1-37,71-90,136-165`
- `plans/skill-issue-project-completion/reorganization-dependency-audit.md:20-29,89-102,149-175`
- `plans/harness-setup.md:1-640`
- `evaluations/skill-calling/smoke/real-harness-smoke-report.md:1-168`
- `evaluations/skill-calling/instrumentation-contract.md:1-56`
- `evaluations/skill-calling/event.schema.json:1-52`
- `evaluations/skill-calling/built-ins/gardening-web-application.json:1-156`
- `evaluations/skill-calling/built-ins/community-archive-desktop-application.json:1-156`
- `evaluations/skill-calling/built-ins/neighborhood-emergency-preparedness-program.json:1-156`
- `evaluations/skill-calling/scenarios/gardening-web-application/conversation.md:1-125`
- `evaluations/skill-calling/scenarios/gardening-web-application/instructions.md:1-23`
- `evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:1-30`
- `evaluations/skill-calling/smoke/two-turn-scenario.json:1-14`
- `evaluations/skill-calling/smoke/two-turn-answer-sheet.json:1-22`
- `cli/internal/replay/replay.go:27-160`
- `cli/internal/evaluation/evaluation.go:20-113,131-328,449-525,590-738`
- `cli/internal/runstate/runstate.go:15-48,58-95,125-213`
- `cli/internal/harness/harness.go:10-87`
- `cli/README.md:1-32,40-53,124-183`
- `src/data/siteData.ts:1-81`
- Historical `plans/website/data-and-design-contract.md:1-34` at commit `22445be`
- `plans/website/reference-and-architecture.md:1-100`
- `README.md:33-61,115-149`

## Findings

### Current Plans — Authority And Measurement Boundary

The completed product/support contract owns the five-harness minimum tier, thirteen harness-and-model cells, medium-setting rationale, one-suite MVP threshold, valid-run meaning, and public-claim boundary. The skill-calling supporting plan owns the scenarios, expected-call meaning, event/evidence contract, and graph semantics. `src/data/siteData.ts` is only the editable website consumer and is explicitly forbidden from becoming the source of evaluation meaning.

**Evidence:** The authority map classifies the product/support contract as the completed foundation for the matrix and claim boundary, the skill-calling plan as the current supporting owner for evaluation and graph semantics, and `siteData.ts` as a Work Block 4 consumer of qualified Work Block 3 artifacts (`plans/skill-issue-project-completion/document-authority-and-update-map.md:20-33,47-51`). Evaluation-meaning or schema changes must update the owning plan and campaign assets, update CLI generation/export, invalidate and rerun affected Work Block 3 results, and only then update website data (`plans/skill-issue-project-completion/document-authority-and-update-map.md:103-108`).

**Implication:** Chart migration must consume current result semantics as fixed inputs. Any proposed meaning change, including redefining expected calls, adding grouping fields to a governed schema, or reinterpreting old observations, crosses into an evaluation-contract change and requires the authority-map rerun route.

### Current Plans — The Evaluation Measures First Required Skill Activations

The campaign measures whether the active harness-and-model combination invokes supplied skills at predetermined required points. It does not grade the loaded skill body's quality. Expected calls are the first turns on which benchmark skills are semantically required, fixed before campaign observation.

**Evidence:** The supporting plan separates invocation measurement from body-quality grading and requires expected first-activation turns to be recorded before results are observed (`plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:5-8,17-25`). The human-readable answer sheet names five required first activations and says later applicable turns do not require another event when the harness keeps a skill loaded (`evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:5-23`). It classifies unmatched required activations as missing and other turn/skill events as additional, without a pass/fail verdict (`evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:25-30`).

**Implication:** `called` and `missed` mean called or missed **expected first activations**, not all skill decisions, all applicable turns, all skill uses, or body-quality outcomes. Website labels and methodology copy should use that narrow meaning.

### Current Assets — Governed Scenario, Turn, And Expected-Call Structure

Each governed built-in is a schema-version-1 unit containing an `evaluation_id`, one scenario with `scenario_id` and an ordered array of `{turn_id, prompt}`, and one private answer sheet with the same `scenario_id` and an `expected` array of `{turn_id, skill}`. The three current built-ins each have 30 turns and the same five expected first-activation pairs: turn 1 for `dictate-plan` and `document-update-discipline`, turn 11 for `prompt-writing`, turn 25 for `skill-authoring-discipline`, and turn 30 for `system-change-ownership`.

**Evidence:** The gardening built-in shows the unit identities, ordered scenario, and five answer-sheet pairs (`evaluations/skill-calling/built-ins/gardening-web-application.json:1-10,121-156`). The community archive and emergency-preparedness units carry the same five expected pairs in their answer-sheet sections (`evaluations/skill-calling/built-ins/community-archive-desktop-application.json:130-156`; `evaluations/skill-calling/built-ins/neighborhood-emergency-preparedness-program.json:130-156`). Production validation requires schema version 1, a nonempty scenario ID, at least one turn, unique nonempty turn IDs, answer-sheet/scenario identity agreement, existing turn IDs, and available skill names (`cli/internal/replay/replay.go:27-59`; `cli/internal/evaluation/evaluation.go:590-627`).

**Implication:** For the current built-ins, one scenario run has four chart points because only turns with expected calls are emitted, but its derived sample size is five because turn 1 contains two expected skills. A three-scenario cell therefore has 15 governed expected activations if the three scenario artifacts remain unchanged. The full 30-turn domain must remain visible even though only turns 1, 11, 25, and 30 produce points.

### Current Code — Built-In And Custom Identities Are Related But Distinct

Built-in selection loads one embedded unit by `evaluation_id` and forbids custom inputs. Custom mode requires a supplied skill root, scenario, and answer sheet; its answer-sheet structure is validated, while semantic correctness remains caller-owned. In current built-ins, `evaluation_id` and `scenario_id` happen to be equal. Production code nevertheless treats them as separate concepts, and custom results use `scenario_id` as their fallback evaluation identity.

**Evidence:** The loader validates the requested built-in identity, validates its scenario and answer sheet, and otherwise requires all three custom inputs (`cli/internal/evaluation/evaluation.go:472-525`). `evaluationIdentity` returns the requested built-in ID when present and otherwise returns the scenario ID (`cli/internal/evaluation/evaluation.go:689-694`). The custom contract warns that the CLI checks runnable structure but does not judge whether expected calls are semantically correct (`evaluations/skill-calling/instrumentation-contract.md:12-14`).

**Implication:** A chart-data pipeline should preserve `evaluation_id` and `scenario_id` separately even while the current governed assets use matching strings. Custom evidence must not be mixed with governed campaign evidence merely because its scenario ID resembles a built-in ID.

### Current Code — Turn Attribution Is Active-Boundary Evidence

The replay runner opens one primary-agent session, sets private active-turn state before each prompt, sends the prompt verbatim, waits for terminal completion, records any harness session ID and attributable signals, and clears active-turn state after that response. Signals outside an active turn remain unattributed rather than being assigned to a nearby turn.

**Evidence:** Replay preserves ordered prompt execution and exposes before/after boundaries around every turn (`cli/internal/replay/replay.go:107-159`). The evaluation runner maps the before boundary to `SetActiveTurn`, records session/signals at the after boundary, and then clears the active turn (`cli/internal/evaluation/evaluation.go:262-285`). The instrumentation contract defines this as the turn-attribution rule and explicitly retains out-of-boundary signals as unattributed evidence (`evaluations/skill-calling/instrumentation-contract.md:32-36`).

**Implication:** `turn_id` is an evidentiary attribution boundary, while numeric `turn` is its one-based position in the ordered scenario. The numeric chart axis must be derived from scenario order and must retain `turn_id` for source verification; timing, context consumption, or token counts cannot substitute for this identity.

### Current Code — Called, Missed, Additional, And Duplicate Semantics

The detailed result retains all attributed observed events, computes missing calls against exact `{turn_id, skill}` expected keys, classifies attributed events on other keys as additional, and retains un-attributed events separately. The website derivation deduplicates expected and observed skill identities at each expected turn, so repeated signals do not inflate `called`; `missed` is the number of unique expected skills absent at that turn.

**Evidence:** `deriveResult` builds exact expected/observed key sets while preserving detailed observed, additional, and unattributed arrays (`cli/internal/evaluation/evaluation.go:642-687`). `deriveWebsiteResult` emits only expected-call turns, uses ordered position plus one as `turn`, counts unique observed expected skills, and sets `missed` to expected unique skills minus `called` (`cli/internal/evaluation/evaluation.go:696-737`). The CLI documentation states that sample size is the sum of `called + missed` and repeated signal events do not inflate it (`cli/README.md:150-173`).

**Implication:** The compact chart artifact is appropriate for called-versus-missed first-activation counts. Investigation of extra reloads, unexpected skills, duplicate signals, or attribution problems must link back to `result.json` and optional diagnostics rather than being inferred from chart points.

### Current Code — Atomic Run Identity And Artifact Boundary

A production `run_id` is a random 16-byte value encoded as 32 hexadecimal characters and identifies one invocation of one scenario, not a three-scenario suite. The run writes `result.json` and `website.json` only after the scenario replay completes and evidence is derived. Output directory naming adds harness, UTC start timestamp, and the first eight run-ID characters for human navigation, but the full `run_id` is the durable join key.

**Evidence:** Run IDs are generated from 16 random bytes (`cli/internal/runstate/runstate.go:66-72`). `Service.Run` selects one scenario, creates one run and one unique output directory, executes the single scenario, then writes detailed and website artifacts and marks the private run complete (`cli/internal/evaluation/evaluation.go:131-190,281-328`). The output contract always retains those two artifacts for a tooling-complete evaluation (`cli/README.md:136-180`).

**Implication:** Aggregation must never infer suite membership from timestamps, directory adjacency, or an eight-character prefix. The current production schema needs an external accepted-campaign manifest or aggregation artifact to bind three distinct scenario `run_id` values into one harness/model/reasoning cell suite.

### Current Code — Result Identity Is Split Across Detailed And Compact Artifacts

`result.json` carries `run_id`, `harness`, `model`, `reasoning`, `evaluation_id`, `scenario_id`, scope, timestamps, and detailed call classifications. `website.json` carries `run_id`, `scenario_id`, `harness`, `model`, `total_turns`, and points. It omits reasoning, evaluation identity, timestamps, environment metadata, and qualification state.

**Evidence:** The production result types define the exact detailed and compact fields (`cli/internal/evaluation/evaluation.go:79-113`). The supporting plan deliberately specifies the compact schema without reasoning or evaluation identity (`plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:49-57`). CLI documentation calls `result.json` authoritative evidence and `website.json` compact chart data (`cli/README.md:136-173`).

**Implication:** `website.json` alone can group atomic scenario observations by run, scenario, harness, and model. It cannot safely distinguish reasoning settings, governed versus custom evaluation identity, campaign acceptance, harness version, operating system, default-environment attestation, or suite membership. A chart refresh should join the compact artifact to its matching detailed result by full `run_id` and to Work Block 3 acceptance/environment metadata before publication.

### Current Plans — Harness, Model, And Reasoning Define A Matrix Cell

The minimum campaign has thirteen harness-and-model cells across Codex, Claude Code, Cursor, Pi, and OpenCode, all at medium reasoning or the closest documented equivalent. Each configured cell must record the exact provider model identifier, version or alias, effective reasoning, harness version, operating system, and default-environment attestation.

**Evidence:** The authoritative matrix and recording requirements are defined in the product/support contract (`plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:24-48,72-80`). Exact provider identifiers and reasoning equivalents remain a Work Block 3 decision (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:163-166`). Current code resolves CLI defaults only for Claude Code, Codex, Cursor, and Pi and records the resolved request values (`cli/internal/harness/harness.go:38-60`; `cli/internal/evaluation/evaluation.go:49-63,170-186`).

**Implication:** The grouping key for a minimum campaign cell is at least normalized harness plus exact effective model identifier/alias plus effective reasoning, under a named campaign version. Display labels such as “GPT-5.6” are insufficient identity. OpenCode is still absent from the current evaluation-default/runtime set, so the five-harness minimum cannot yet be represented as complete.

### Current Plans — One Suite Per Cell, No Repetition Threshold

The MVP requires one tooling-complete suite of all three governed scenarios for every one of the thirteen cells. It does not require repeated trials, statistical thresholds, or a pass threshold. This implies 39 atomic scenario runs and 15 expected first-activation observations per cell under the unchanged current fixtures, but only one suite-level observation per cell.

**Evidence:** The product/support contract requires one three-scenario suite per cell and leaves repeated trials and thresholds for later expansion (`plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:44-48,72-80`). The supporting plan repeats the one-suite/no-repetition rule (`plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:59-65,85-89`). README states that one suite cannot establish statistical reliability or universal model behavior (`README.md:55-61`).

**Implication:** Charts may describe the observed called/missed counts from this campaign. They may not claim rates across repetitions, statistical reliability, universal model behavior, persistent model quality, or a general winner from a sample of one suite per cell.

### Current Plans — Tooling-Complete Defines Validity, Not Call Success

A valid suite completes all three governed scenarios with functioning instrumentation and a complete evidence package. Zero expected calls is still valid model-and-harness data. Launch, permission, session, marker, protocol, or other tooling failures must be fixed and rerun instead of graphed as misses.

**Evidence:** The product/support contract defines tooling-complete independently of calls and separates tooling failures from model outcomes (`plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:50-64`). The instrumentation contract classifies launch, permission, session, marker, and protocol failures as tooling failures while retaining missing calls after a tooling-complete replay as observations (`evaluations/skill-calling/instrumentation-contract.md:38-46`). Work Block 3 requires tooling failures to be repaired and rerun and zero-call suites to be retained (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:71-83`).

**Implication:** A chart row is eligible only from a tooling-complete accepted run. A zero-call scenario should appear with zero `called` and all governed expectations `missed`; a failed or interrupted runner should have no performance point and must remain visible only in qualification/tooling records.

### Current Plans — Work Block 3 Owns Acceptance And Aggregation

Work Block 3, rather than the CLI writer or website, qualifies environments, recruits external evaluators, completes the thirteen-cell matrix, aggregates accepted runs under one evidence standard, and audits metadata, contamination, limitations, and support classification. Work Block 4 publishes only after those results are stable.

**Evidence:** The parent plan assigns qualification, matrix execution, accepted external evidence, aggregation, and audit to Work Block 3 (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:71-83`). The reorganization audit says real-harness qualification, campaign aggregation, and all thirteen-cell evidence remain unfinished (`plans/skill-issue-project-completion/reorganization-dependency-audit.md:89-96`). The parent status states that no accepted skill-calling campaign result or qualified cell exists yet (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:21-31`).

**Implication:** An emitted `website.json` is chart-shaped evidence, not automatic campaign acceptance. Publication needs a Work Block 3 acceptance layer that records which atomic runs form each cell suite and which environment/attestation evidence qualifies them.

### Current State — Real Campaign Claims Are Not Yet Available

The repository currently has completed governed fixtures and a reusable four-harness runner, but no accepted skill-calling campaign cell. The completed twenty-one-iteration production-refinement campaign is a different evaluation system in a bounded Codex/high-reasoning environment and cannot substitute for the planned medium-setting skill-calling campaign. Codex, Cursor, Claude Code, and Pi have completed bounded real-harness smoke routes.

**Evidence:** The reorganization audit explicitly separates the production-skill refinement campaign from the skill-calling campaign and says the former does not prove the future thirteen-cell medium campaign (`plans/skill-issue-project-completion/reorganization-dependency-audit.md:20-29,89-101`). The retained smoke report records completed bounded routes and their limits (`evaluations/skill-calling/smoke/real-harness-smoke-report.md`).

**Implication:** The website can presently claim that the evaluation system and governed fixtures exist and can show clearly labeled representative/mock presentation data. It cannot claim observed cross-harness benchmark performance, qualified model comparisons, a completed minimum matrix, or real campaign sample sizes.

### Current State — Smoke And Custom Results Have Narrower Claim Boundaries

The two-turn smoke scenario is intentionally a small interface/adaptor probe, not one of the three governed 30-turn campaign scenarios. Custom evaluations can produce structurally identical result artifacts, but their expected-call semantics are caller-owned and their results remain local unless separately reviewed and accepted.

**Evidence:** The smoke assets contain only two turns and four explicit expected calls (`evaluations/skill-calling/smoke/two-turn-scenario.json:1-14`; `evaluations/skill-calling/smoke/two-turn-answer-sheet.json:1-22`). The smoke report records that these routes prove interfaces and cleanup rather than governed campaign results, and that custom results remain local without separate review and acceptance (`evaluations/skill-calling/smoke/real-harness-smoke-report.md`).

**Implication:** Smoke data may support claims that a route launched, preserved a session, carried configuration, wrote output, and cleaned up. It may not support governed benchmark or campaign-performance claims. Custom data must be labeled custom and excluded from the governed campaign unless Work Block 3 accepts it under the campaign evidence standard.

### Current Website — Existing Values Are Explicit Mock Data

The current website data is a presentation mock built around consumed-context checkpoints, two illustrative harness/model cards, manually stored sample sizes, and explicitly described mock evaluations. This structure predates the governed numeric-turn artifact and is not current campaign evidence.

**Evidence:** `src/data/evaluationData.ts` supplies explicitly illustrative results through the implemented typed artifact adapter, while `src/data/siteData.ts` owns the preview status and methodology wording. The current parent plan says the website still uses illustrative release and benchmark content (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md`).

**Implication:** These values can be used only as labeled representative layout data. They cannot be described as observed calls, misses, prompts, checkpoints, harness counts, model performance, or campaign sample size. Migration should replace their context axis and hand-authored sample sizes with numeric turn positions and derived `called + missed` counts before real-result publication.

### Recommendations — Preserve A Three-Layer Result Identity

Use three explicit identity layers in the website data-refresh path:

1. **Atomic run:** full `run_id`, `evaluation_id`, `scenario_id`, harness, exact effective model, effective reasoning, and the matching detailed/compact artifact pair.
2. **Qualified cell suite:** one stable suite/cell identifier that binds exactly one accepted run for each of the three governed scenarios plus harness version, operating system, default-environment attestation, adjacent-configuration/contamination disclosures, and qualification status.
3. **Campaign:** a stable campaign/version identifier covering the authoritative matrix, fixture/schema version, accepted cell suites, omissions or explicit contract revisions, and publication date.

**Evidence:** Current production supplies only atomic run identity (`cli/internal/evaluation/evaluation.go:79-113,131-190`). Work Block 3 still owns suite aggregation and environment acceptance (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:71-83`). The authority map requires result invalidation and rerun when evaluation meaning or schema changes (`plans/skill-issue-project-completion/document-authority-and-update-map.md:103-108`).

**Implication:** This layered model avoids inferring campaign meaning from filenames and allows the chart to facet or filter by campaign, cell, and scenario while preserving links to exact source runs. The suite and campaign layers should be created by Work Block 3 aggregation, not invented inside `siteData.ts`.

### Recommendations — Separate Descriptive Displays From Evaluative Claims

For mock data, use explicit “representative,” “illustrative,” or “mock” labels and avoid naming values as results. For accepted real data, describe exact observed calls and misses for the identified scenario or three-scenario suite, identify the environment and one-suite sample, link to the method and evidence, disclose omissions and limitations, and avoid pass/fail labels or generalized reliability language.

**Evidence:** The public-claim authority permits statements only about what was done, observed, and how it was run, and rejects perfection, guarantee, or universal-behavior claims (`plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:57-64`). The publication plan requires descriptive graphs, inspectable method, evidence, and limitations (`plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:67-72`).

**Implication:** The chart can truthfully show strong or weak observed behavior without converting the limited MVP observations into a project-authored verdict. “Reliability” should be framed as the question under investigation, not as a statistically established property of a model/harness from one suite.

### Unsupported Current Contract — Event Schema And Emitted Event Identity Diverge

The Go event writer emits `reasoning` and `evaluation_id` on every invocation event, and the instrumentation contract says private state/events retain those identities. The checked-in JSON Schema uses `additionalProperties: false` but defines neither `reasoning` nor `evaluation_id`. Consequently, an event emitted by current production code is not valid against the current checked-in event schema.

**Evidence:** The Go `Event` type and writer include both fields (`cli/internal/runstate/runstate.go:36-48,165-177`). The instrumentation contract names effective reasoning and selected evaluation in private state (`evaluations/skill-calling/instrumentation-contract.md:32-36`). The JSON Schema forbids additional properties and defines fields only through `recorded_at`, without `reasoning` or `evaluation_id` (`evaluations/skill-calling/event.schema.json:5-51`).

**Implication:** Raw event-schema conformance is currently unsupported. This does not change the already-defined detailed or compact result field meanings, but Work Block 1/2 audit must reconcile the event schema and code before raw events are presented as schema-valid campaign evidence. Because the authority map treats schema changes as rerun-sensitive, the owner must decide whether this is a schema correction at version 1 or an explicit version change before Work Block 3 acceptance.

### Unsupported Current Contract — Suite And Qualification Identity Have No Production Schema

No current production artifact identifies a three-scenario suite, campaign, accepted matrix cell, harness version, operating system, default-environment attestation, contamination disclosure, or qualification decision. The compact artifact also omits reasoning and evaluation identity.

**Evidence:** Production result fields stop at atomic run and call evidence (`cli/internal/evaluation/evaluation.go:79-113`). Work Block 3 aggregation and environment qualification are unfinished (`plans/skill-issue-project-completion/reorganization-dependency-audit.md:89-96`). The authoritative completion criteria require exact environment metadata, attestations, raw evidence, and portable results for every cell (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:140-145`).

**Implication:** A complete campaign chart dataset cannot yet be generated solely by scanning `website.json` files. The missing Work Block 3 aggregation/acceptance artifact is a true blocker for real multi-cell publication and should be treated as an explicit upstream deliverable rather than guessed during website migration.

## Notes

- The current embedded built-ins share identical expected-call positions and skill names, but their scenario prompts and domains differ. Aggregation should retain scenario identity rather than collapsing the three runs into one indistinguishable five-call sample.
- The current CLI supports evaluation defaults for four harnesses, while OpenCode is part of the authoritative five-harness minimum. The four-harness runner plan does not claim to complete OpenCode.
- The output directory's timestamp and eight-character run prefix are operational labels, not sufficient campaign or suite identity.
- No claim about actual model or harness performance was supported because the controlling plan states that no accepted skill-calling cell exists.

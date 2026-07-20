# Chart Comparison Design

## Assignment

- **Goal:** Recommend how compact numeric-turn evaluation evidence should be compared on the website, including comparison structure, selection behavior, controls, scales, legends, multiple expected calls, responsive behavior, and representative data for the pre-campaign presentation state.
- **Scope:** Completed assignments 01–05, their cited production website and CLI sources, governed expected-call assets, and the controlling Work Block 3/4 result-publication contract.
- **Exclusions:** Product-code edits; implementation sequencing; chart-library migration; transcript interpretation; changes to evaluation semantics or artifact schemas; invented campaign outcomes; claims about real harness or model performance.

## Sources

- `research/implementation-research/website-chart-migration/assignments/01-cli-evaluation-production-path.md` — atomic run, detailed result, compact artifact, and website-projection ownership.
- `research/implementation-research/website-chart-migration/assignments/02-replay-scoring-turn-attribution.md` — exact called/missed semantics, sparse numeric turns, unequal per-turn denominators, duplicate collapse, and one-run-per-scenario boundary.
- `research/implementation-research/website-chart-migration/assignments/03-methodology-and-result-identity.md` — governed first-activation method, cell/suite/campaign identity, campaign acceptance, and current publication blockers.
- `research/implementation-research/website-chart-migration/assignments/04-current-website-chart.md` — card schema, categorical X axis, fixed Y scale, smoothed lines, legend, tooltip, and accessibility behavior at research time.
- `research/implementation-research/website-chart-migration/assignments/05-website-layout-contracts.md` — established page placement, two-stage responsive layout, data/component ownership, and visual contracts.
- `cli/internal/evaluation/evaluation.go:79-113,642-737` — detailed/compact result types and concrete called/missed point derivation.
- `cli/internal/replay/replay.go:27-59,112-159` — ordered scenario-turn identity and replay boundaries.
- `evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:5-30` — governed required first activations, later applicable turns, and descriptive missing/additional treatment.
- `evaluations/skill-calling/built-ins/gardening-web-application.json` — concrete 30-turn scenario and four expected turn/skill pairs.
- `evaluations/skill-calling/built-ins/community-archive-desktop-application.json` — second governed four-opportunity answer sheet.
- `evaluations/skill-calling/built-ins/neighborhood-emergency-preparedness-program.json` — third governed four-opportunity answer sheet.
- `src/data/siteData.ts:1-80` — current mock card identities, labels, sample sizes, and percentage-based values.
- `src/components/EvaluationChart.tsx:17-105` — current card header, accessibility summary, line-chart configuration, tooltip, and local legend.
- `src/styles.css:360-479,515-605` — current results grid, card dimensions, footer/legend treatment, and 800/680-pixel responsive states.
- `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:44-68` — one-suite descriptive-claim boundary and required public evidence.
- `plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:49-72` — compact graph schema, campaign shape, and descriptive graph semantics.
- `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:71-97` — Work Block 3 aggregation ownership and Work Block 4 numeric-turn chart requirements.

## Findings

### Finding 1 — Default Comparison Unit: Facet By Qualified Harness/Model Cell

The best-supported primary layout is a small-multiple grid with one chart card per qualified harness/model/reasoning cell. Every visible card should show the same selected scenario, the same numeric turn domain, the same called/missed semantics, and the same Y scale. Harness should be the stable outer grouping and model the card-level identity, while the exact effective model and reasoning remain visible metadata. This preserves the campaign matrix without collapsing observations that have different runtime identities.

**Evidence:** One `website.json` describes one run of one scenario and carries singular `run_id`, `scenario_id`, `harness`, and `model`; it is not a campaign aggregate (`assignments/02-replay-scoring-turn-attribution.md:93-101`; `cli/internal/evaluation/evaluation.go:98-113,696-737`). The minimum campaign cell is at least harness plus exact effective model plus effective reasoning, and current compact data omits reasoning (`assignments/03-methodology-and-result-identity.md:96-110`). The production website already uses independent responsive cards in a two-column grid (`assignments/04-current-website-chart.md:74-80`; `src/styles.css:388-442`).

**Implication:** Facets preserve identity and keep the two semantic series consistent across cells. Grouping by harness alone or model alone would erase a campaign dimension; each card should therefore represent the combined qualified cell even when the surrounding layout is harness-first.

### Finding 2 — Overlay Is A Focused Secondary Mode, Not The Default

Overlaying all cells is a poor fit because each cell contributes two series. The minimum thirteen-cell matrix would create twenty-six traces before scenarios, runs, or repeated trials are considered, forcing the legend to encode both metric and cell identity and making coincident sparse points difficult to distinguish. A focused overlay can remain a conditional secondary comparison for at most two selected cells and one selected metric, but the default cross-cell comparison should remain aligned facets.

**Evidence:** Called and missed are separate compact series (`cli/internal/evaluation/evaluation.go:98-113`), and the current chart already needs two semantic colors and a local two-item legend for only one card (`src/components/EvaluationChart.tsx:61-105`). The authoritative minimum campaign contains thirteen harness/model cells (`assignments/03-methodology-and-result-identity.md:104-118`). Compact points occur only at turns with expected calls, so many cell traces will occupy the same four governed X positions (`assignments/02-replay-scoring-turn-attribution.md:83-109`).

**Implication:** Facets provide readable direct comparison without a combinatorial legend. A two-cell, one-metric overlay is defensible only as an explicitly selected close comparison; a many-cell dual-series overlay is lower fit.

### Finding 3 — Scenario Selection Must Be Global And Run Selection Must Remain Evidentiary

The scenario selector should control the entire comparison grid so that every visible cell is compared on the same governed conversation. With direct atomic artifacts, the honest default is the first governed scenario in the campaign's stable published order. “All scenarios” should appear only when Work Block 3 supplies a qualified suite-level aggregation that preserves the contributing `run_id` and `scenario_id` values and documents its derivation. The MVP campaign specifies one accepted run per scenario per cell, so public run selection adds no useful comparison choice; the full run identity belongs in evidence details or a raw-result link.

**Evidence:** Production emits one scenario per run and no suite identity (`assignments/02-replay-scoring-turn-attribution.md:93-101`; `assignments/03-methodology-and-result-identity.md:88-102`). Work Block 3 owns accepted-run selection, three-scenario aggregation, environment evidence, and graph-ready outputs (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:71-83`). Repeated trials are deferred, and one suite cannot support statistical reliability claims (`assignments/03-methodology-and-result-identity.md:112-126`).

**Implication:** A global scenario control preserves like-for-like comparison. A suite-summary default becomes preferable only after a real aggregation contract exists. If repeated trials are introduced later, a run/trial control can be added with the accepted run clearly identified; filenames or timestamps must never determine the selection.

### Finding 4 — Use A Small Control Set With Stable Semantics

The useful primary controls are: one global scenario selector, a harness filter, an optional model filter within the selected harness set, and independent visibility toggles for called and missed. “All qualified cells” should remain the default cell filter so the public chart does not silently hide weak or zero-call observations. Effective reasoning is fixed at medium for the minimum campaign and should be shown as metadata rather than promoted to a filter until multiple settings exist. Campaign version belongs in the surrounding result identity, not as a casual chart toggle.

**Evidence:** The campaign requires retention of zero-call tooling-complete suites and descriptive presentation without pass/fail labels (`assignments/03-methodology-and-result-identity.md:120-134`). Harness, exact model, reasoning, environment, and campaign identity are distinct evidence dimensions (`assignments/03-methodology-and-result-identity.md:96-110,160-170`). The current results grid is data-driven from `siteData.evaluations`, so filtering cards is consistent with the existing presentation ownership (`assignments/05-website-layout-contracts.md:84-99`).

**Implication:** These controls answer the meaningful questions without creating unsupported analytical views. Sorting, ranking, or default filtering by called count would imply a winner-oriented judgment and should be avoided; cards should keep a stable campaign order.

### Finding 5 — The Default View Should Be Complete, Scenario-Aligned, And Descriptive

The default real-data view should show the currently selected governed scenario across all qualified cells, grouped harness-first and ordered by the published campaign matrix. Each card should identify harness, exact model, reasoning, scenario, and derived sample size, with direct access to its source evidence. The result heading should say that the chart shows observed expected first activations at selected turns, not general skill reliability or all skill use.

**Evidence:** `called` and `missed` measure exact governed first-activation opportunities, not body quality, task success, every applicable turn, or invocation frequency (`assignments/02-replay-scoring-turn-attribution.md:31-37,63-81`; `assignments/03-methodology-and-result-identity.md:48-54`). Current public authority permits descriptions of what was done and observed while rejecting universal or pass/fail claims (`assignments/03-methodology-and-result-identity.md:172-178`). The current chart heading and method copy still describe provisional context-consumption data (`assignments/05-website-layout-contracts.md:119-134`).

**Implication:** The default view is comprehensive without overstating the evidence. A future qualified suite-summary can replace the selected-scenario default, but scenario drill-down and source-run provenance must remain available.

### Finding 6 — Preserve Sparse Numeric Turns Instead Of Filling Gaps

Every card should use a numeric horizontal domain from `1` through `total_turns`, place points at their actual one-based turn positions, and leave turns without expected calls visually unmeasured. Persistent point markers and straight connections fit the evidence better than the current hidden dots and monotone smoothing. Connections should be described as visual guides between selected expected-call turns, not observations of the intervening turns.

**Evidence:** Production emits only turns containing expected calls, while numeric `turn` is the one-based scenario position and `total_turns` is the full scenario length (`assignments/02-replay-scoring-turn-attribution.md:83-91`; `cli/internal/evaluation/evaluation.go:696-737`). The current chart defaults X to a categorical scale and smooths two dotless monotone lines, which would equally space sparse turns and imply a continuous trend (`assignments/04-current-website-chart.md:50-72`; `src/components/EvaluationChart.tsx:47-88`). Work Block 4 explicitly requires a numeric `1..total_turns` domain (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:93`).

**Implication:** Numeric spacing makes the distance from turns 1 to 11 visibly different from 11 to 25, while persistent markers disclose exactly where measurements exist. Unscored turns must not be inserted as zero values because zero would falsely mean an expected activation was measured and missed or called there.

### Finding 7 — Share A Data-Derived Integer Y Scale Across Visible Cells

The Y axis should start at zero and use a shared integer domain derived from the selected scenario or qualified suite, not from each cell's observed performance. The stable upper bound should be the maximum per-turn expected opportunity count, calculated as `called + missed` across the selected comparison dataset, with integer ticks. Cell filtering or hiding one semantic series should not change that domain. For each current governed atomic scenario this yields `0..1`; a valid three-scenario same-turn suite aggregation would yield `0..3` only if Work Block 3 explicitly defines that aggregation.

**Evidence:** Per point, `missed` is expected unique skills minus called, so `called + missed` is the point's governed denominator. Current governed scenarios contain one expected skill at turns 1, 11, 25, and 30. The research-time fixed `0..20` scale was mock-specific and could either clip or waste space for real data.

**Implication:** A denominator-owned shared scale permits honest height comparison between cells and stays stable when filters change. Per-card autoscaling is rejected because identical heights could represent different counts; a percentage scale is lower fit because it would hide the one-versus-two opportunity denominator unless shown as an explicitly derived secondary view.

### Finding 8 — Multiple Expected Calls Remain One Turn Point With An Explicit Denominator

When a custom evaluation has multiple expected skills on one turn, it should remain a single X position whose two values sum to the number of unique expected first activations at that turn. The tooltip and accessible summary should state the denominator explicitly. The card sample size remains the sum of `called + missed` over all points, not the point count and not the raw number of signal events.

**Evidence:** Expected calls are exact turn/skill pairs and are deduplicated by skill within a turn for the compact artifact. The current governed fixtures produce four points and `n = 4`; the production schema continues to support multiple expected skills on a turn for custom evaluations. Duplicate detailed observations do not inflate compact called counts or sample size.

**Implication:** Splitting a multi-skill turn into duplicate X positions, averaging it to one decision, or treating duplicate signals as additional calls would misstate the compact result. Explicit denominator text also prevents equal called/missed values from being hidden when the plotted marks overlap.

### Finding 9 — Legend Language Must Preserve Measurement Meaning

Use consistent semantic colors for the two metrics across every facet. The concise visible labels can remain “Called” and “Missed” only when the chart heading or legend qualifier makes clear that both refer to expected first activations. Tooltips and accessible text should use the full phrases “expected activations called” and “expected activations missed.” Cell identities should stay in facet headers rather than enter the legend.

**Evidence:** Current CSS already owns stable light/dark call and miss colors, while the component owns the manual legend and tooltip names (`assignments/04-current-website-chart.md:58-80`; `src/styles.css:409-479`). Production called/missed semantics are narrower than generic skill calls or misses (`assignments/02-replay-scoring-turn-attribution.md:31-37,63-71`). The established website contract requires shared theme tokens and consistent semantic meanings across themes (`assignments/05-website-layout-contracts.md:136-149`).

**Implication:** A two-item semantic legend remains readable regardless of cell count. Repeating the compact legend inside each card is the best default for self-contained cards and long mobile scrolling; a single shared legend is a reasonable desktop-only alternative if each card retains accessible series names.

### Finding 10 — Mobile Should Prefer Filtered Facets Over Dense Overlays

The same semantic default should survive on narrow screens: one card per row, local card identity, local legend, and the full numeric turn domain. Controls should wrap or collapse into compact selectors without horizontal scrolling. Four governed point positions fit the shorter mobile chart, but long exact model identities and thirteen stacked cells create a long page, making the harness/model filters materially useful. Mobile should not switch to a many-cell overlay or silently hide cells merely to shorten the page.

**Evidence:** The current responsive contract retains two chart columns through 681 pixels, then changes to a one-column grid, shorter 17-rem chart, top dividers, and stacked footer at 680 pixels (`assignments/05-website-layout-contracts.md:47-64`; `src/styles.css:515-605`). The current visual audit approved the no-overflow responsive hierarchy and established chart surface (`assignments/05-website-layout-contracts.md:136-149`). The minimum matrix can contain thirteen cells (`assignments/03-methodology-and-result-identity.md:104-118`).

**Implication:** Facets degrade predictably from two columns to one while preserving evidence. A compact harness filter improves navigation without changing the default data contract; dense overlays, horizontally scrolling chart tables, and breakpoint-specific hidden results are lower-fit mobile responses.

### Finding 11 — Representative Dataset: Two Illustrative Cells Across Two Illustrative Scenarios

Use four clearly labeled atomic artifacts: two fictional harness/model cells evaluated against two fictional 30-turn scenarios. This is the smallest set that exercises global scenario selection, cross-cell facets, both called and missed series, zero values, equal values, and sparse numeric spacing. Every displayed card has a derived `n = 4`. The entire fixture should carry the visible label **“Illustrative layout data — not observed evaluation results.”** Identifiers should also be unmistakably illustrative rather than resembling production run IDs or real campaign assets.

| Illustrative scenario     | Illustrative cell           | Turn 1 called/missed | Turn 11 called/missed | Turn 25 called/missed | Turn 30 called/missed | Derived n |
| ------------------------- | --------------------------- | -------------------: | --------------------: | --------------------: | --------------------: | --------: |
| `illustrative-scenario-a` | Harness Alpha / Model Alpha |                1 / 0 |                 1 / 0 |                 0 / 1 |                 1 / 0 |         4 |
| `illustrative-scenario-a` | Harness Beta / Model Beta   |                1 / 0 |                 0 / 1 |                 1 / 0 |                 0 / 1 |         4 |
| `illustrative-scenario-b` | Harness Alpha / Model Alpha |                0 / 1 |                 1 / 0 |                 1 / 0 |                 0 / 1 |         4 |
| `illustrative-scenario-b` | Harness Beta / Model Beta   |                1 / 0 |                 1 / 0 |                 0 / 1 |                 1 / 0 |         4 |

Each artifact should retain the `website.json`-style fields `schema_version`, an illustrative-only run identifier, `scenario_id`, `harness`, `model`, `total_turns: 30`, and four points with numeric `turn`, an illustrative `turn_id`, `called`, and `missed`. The mock must not add reasoning, suite qualification, environment claims, or campaign acceptance that the compact artifact does not own.

**Evidence:** The compact artifact field set and sparse-point derivation are production-owned (`cli/internal/evaluation/evaluation.go:98-113,696-737`). The three governed scenarios currently use a 30-turn domain and five opportunities distributed as 2, 1, 1, and 1 over turns 1, 11, 25, and 30 (`assignments/03-methodology-and-result-identity.md:56-63`). The current website is explicitly allowed only representative/mock layout data until accepted campaign evidence exists (`assignments/03-methodology-and-result-identity.md:152-158`).

**Implication:** This fixture tests the intended comparison behavior without inventing real outcomes or borrowing actual harness/model names. Work Block 3 data can replace each illustrative atomic record while preserving the same chart-facing point shape and scenario/cell selection behavior.

### Finding 12 — Conditional Alternatives And Rejected Lower-Fit Options

The following alternatives are conditional rather than primary: a qualified suite-summary default once Work Block 3 publishes an inspectable aggregation; a maximum-two-cell overlay limited to one metric; a normalized percentage view paired with visible raw denominators; and a single shared desktop legend when cards retain accessible metric names. A stacked discrete column presentation is also a defensible fallback if overlapping equal-valued line markers prove unclear, because called and missed are complementary parts of each expected-turn denominator.

Rejected or lower-fit options are: a dual-series overlay of every cell; independent per-card Y scales; categorical spacing of sparse turn numbers; inserting zero-valued unscored turns; monotone smoothing without persistent observed-point markers; aggregation by harness or display-model label alone; a public trial selector before repeated trials exist; ranking cells by called count; default percentage normalization; and treating one atomic `website.json` as a complete campaign result.

**Evidence:** Sparse numeric positioning, exact cell identity, raw denominators, one-run scope, one-suite/no-repetition limits, and current campaign incompleteness are established across `assignments/02-replay-scoring-turn-attribution.md:83-109`, `assignments/03-methodology-and-result-identity.md:88-134`, and `assignments/04-current-website-chart.md:50-72`. The public-claim contract requires descriptive evidence and limitations rather than project-authored verdicts (`assignments/03-methodology-and-result-identity.md:172-178`).

**Implication:** The primary recommendation stays narrow and faithful to available evidence while leaving explicit upgrade paths for future qualified aggregation or repeated trials. Lower-fit options either obscure identity, imply measurements that were not made, or overstate the maturity of the campaign.

## Notes

- Real multi-cell publication remains blocked on the Work Block 3 accepted-suite/campaign aggregation and qualification artifact. Current `website.json` files alone do not identify reasoning, suite membership, environment qualification, campaign acceptance, or the exact accepted three-run cell (`assignments/03-methodology-and-result-identity.md:188-194`).
- The three current governed built-ins share the same expected positions and denominators, but this should not be assumed as a permanent schema guarantee. A future suite summary may combine numeric turns only when its aggregation contract establishes that the scenario turn positions are meaningfully aligned.
- The representative values above are deliberately synthetic and support layout validation only. They must not be described as a preview of expected real performance.

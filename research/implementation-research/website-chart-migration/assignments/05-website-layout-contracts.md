# Website Layout And Contracts

## Assignment

**Goal:** Establish the current website chart's page placement, CSS and responsive behavior, and the production and planning contracts that control a later migration from mock context-consumption data to real turn-based evaluation artifacts.

**Scope:** The live React/CSS/data implementation; the completed website design, data, responsive, and visual-audit records; the project-completion authority map and controlling completion plan; the current skill-calling result contract; and the concrete Go producer of `website.json`.

**Exclusions:** Implementation, implementation planning, changes to source or controlling plans, deployment research, competition-rule analysis, and visual redesign.

## Sources

- Working tree on `main` at commit `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`; several project-completion documents were already modified in the working tree before this assignment, so cited line content reflects the inspected working tree.
- `src/App.tsx` — production page composition, section order, results heading, chart-grid mapping, and method placement.
- `src/styles.css` — production page bounds, chart-grid and card geometry, responsive breakpoints, theme tokens, and semantic series colors.
- `src/components/EvaluationChart.tsx` — production card structure, Recharts configuration, accessible summary, axes, legend, and tooltip.
- `src/data/siteData.ts` — current typed website content and mock graph data.
- Historical `plans/website/data-and-design-contract.md` at commit `22445be` — completed local-data, token, theme, and component-ownership contract.
- `plans/website/reference-and-architecture.md` — completed visual direction, one-page information architecture, chart choice, and original data ownership.
- `plans/skill-issue-project-completion/document-authority-and-update-map.md` — current document classes, website data ownership, and result-schema change routing.
- `plans/skill-issue-project-completion/reorganization-dependency-audit.md` — current website inventory, retained mock-up authority, unfinished public-product boundary, and Work Block 4 ownership.
- `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md` — current execution status and binding Work Block 4 website migration requirements.
- `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md` — completed public-claim and descriptive-results boundary.
- `plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md` — current compact website artifact and descriptive graph-semantics contract.
- `cli/internal/evaluation/evaluation.go` — production `WebsiteResult` schema, artifact write path, and point derivation.

## Findings

### Finding 1 — Production Placement: Results Sit Between Summary Metrics And Method

The chart is the main content of the `#results` section. The page order is header, centered hero, three summary metrics, results heading and graph grid, method note, then footer. The results heading currently frames the charts as an “Evaluation preview” of successful calls and missed opportunities as context is consumed. The method section immediately after the chart supplies the current one-paragraph explanation.

**Evidence**

- `src/App.tsx:35-60` renders the hero before benchmark content; `src/App.tsx:62-69` renders the metric grid.
- `src/App.tsx:71-92` renders the `results section-shell content-shell`, its heading and explanatory paragraph, then maps every `siteData.evaluations` entry into the `.chart-grid`.
- `src/App.tsx:94-105` places the method section directly after results; `src/App.tsx:108-114` follows it with the footer.
- `src/App.tsx:17-20` makes Results and Method the only primary navigation anchors.
- The completed one-page information architecture records the same sequence: header, hero, metrics, responsive evaluation graph cards, method note, footer (`plans/website/reference-and-architecture.md:43-50`).

**Implication**

`App` is the current semantic owner of chart placement and section order. The data migration belongs inside the established results surface; any Work Block 4 additions such as product story, tested environments, limitations, or About content must be reconciled with this page-composition owner rather than rebuilding the existing chart page.

### Finding 2 — Production Responsive Behavior Uses Two Distinct Narrowing Stages

At wide sizes, the page is capped at `82rem`; the results heading is a two-column introduction and the chart grid is two equal columns. Between `681px` and `800px`, the results heading and method become single-column and chart-card horizontal padding shrinks, while the two chart columns remain. At `680px` and below, the content gutter narrows, the chart grid becomes one column, the divider between subsequent cards moves from the left edge to the top edge, chart height drops from `21rem` to `17rem`, and the chart footer stacks vertically.

**Evidence**

- `src/styles.css:33-35` defines the `82rem` wide-content bound; `src/styles.css:117-125` applies the shared page width and section spacing.
- `src/styles.css:360-370` gives results left/right borders and a two-column heading with the explanatory copy capped between `18rem` and `31rem`.
- `src/styles.css:388-402` defines the two-column `minmax(0, 1fr)` chart grid and left divider between adjacent cards.
- `src/styles.css:394-442` gives each card `min-width: 0`, `2rem` padding, a surface background, and a `21rem` chart region.
- At `800px`, only the section heading and method change to one column and chart-card side padding becomes `1.5rem`; the chart grid itself is unchanged (`src/styles.css:515-548`).
- At `680px`, shared content width becomes `calc(100% - 1rem)`, `.chart-grid` becomes one column, adjacent cards receive a top divider instead of a left divider, and chart height becomes `17rem` (`src/styles.css:550-599`).
- At the same `680px` breakpoint, the chart footer changes from a row to a column (`src/styles.css:601-605`).
- Direct browser validation at research time found that the desktop and narrow composition retained the early hero-to-data transition and no horizontal overflow.

**Implication**

The existing responsive contract is not simply “desktop versus mobile”: there is a tablet-width state with stacked explanatory copy but side-by-side charts. Future result-card counts and labels must remain viable in all three states, preserving `min-width: 0`, the breakpoint-specific divider behavior, the shorter narrow chart region, and the no-horizontal-overflow result.

### Finding 3 — Production Chart Geometry And Semantics Are Coupled To The Mock Schema

Each graph renders as a bordered-data surface with model kicker, harness heading, stored sample-size badge, description, line chart, two-series legend, and x-axis caption. The present Recharts configuration assumes percentage-based context checkpoints and bounded mock counts: `contextConsumed` is the horizontal key and formats with `%`; the vertical domain is fixed at `0..20`; and the accessible summary assumes the final point represents 100 percent context consumption.

**Evidence**

- `src/components/EvaluationChart.tsx:21-30` renders the card header, model, harness, stored `sampleSize`, and graph description.
- `src/components/EvaluationChart.tsx:32-40` renders the accessible chart wrapper and passes `graph.points` to Recharts.
- `src/components/EvaluationChart.tsx:34-35` constructs an accessible summary from first and final `skillCalls`/`skillMisses` values and explicitly says context reaches 100 percent.
- `src/components/EvaluationChart.tsx:47-60` uses `contextConsumed`, percentage tick formatting, and a fixed Y-axis domain `[0, 20]` with fixed ticks.
- `src/components/EvaluationChart.tsx:61-88` uses theme-owned tooltip styles and plots `skillCalls` and `skillMisses` with semantic CSS colors.
- `src/components/EvaluationChart.tsx:93-105` renders the legend and `siteData.chart.xAxis` footer caption.
- `src/styles.css:409-479` controls the mono metadata, sample badge, description height, chart geometry, footer rule, and semantic legend lines.

**Implication**

The later field migration cannot be limited to replacing the points array. The `EvaluationChart` owner must also consume `turn`, `called`, `missed`, and `total_turns`; replace percentage ticks and the 100-percent accessible summary; derive the sample display; and replace the fixed vertical scale with the required data-driven limits while preserving the established card surface, legend, tooltip, and accessibility ownership.

### Finding 4 — Production Content Is Centralized, But Evaluation Meaning Is Owned Elsewhere

For the completed mock-up, `src/data/siteData.ts` is the single editable website-content source. It currently owns preview status, product and method copy, release metadata, metrics, chart labels, graph descriptions, stored sample sizes, and mock series values. `App` only maps and places data, while `EvaluationChart` transforms a typed graph definition into the shared chart surface. Current project authority narrows that rule: `siteData.ts` remains the editable website owner, but it must be updated from qualified artifacts and must never define evaluation meaning.

**Evidence**

- The current TypeScript schema contains `contextConsumed`, `skillCalls`, `skillMisses`, and per-graph `sampleSize` (`src/data/siteData.ts:1-14`).
- The current object labels itself `Preview benchmark · Local data`, stores summary metrics and chart labels, supplies two mock graph definitions, and says the mock data only shows the report shape while the harness is finalized (`src/data/siteData.ts:16-80`).
- The completed website contract makes `siteData.ts` the single source for replacement copy, release URLs, metrics, graph definitions, and values, and says another graph is added by one new data entry rather than page markup (historical `plans/website/data-and-design-contract.md:3-34` at commit `22445be`).
- The same contract assigns page composition to `App`, chart presentation and data summary to `EvaluationChart`, and layout/responsive/theme behavior to `styles.css` (historical `plans/website/data-and-design-contract.md:58-66` at commit `22445be`).
- The current authority map defines `src/data/siteData.ts` as the editable website-copy and graph-data owner, but directs Work Block 4 to populate it from qualified release artifacts and Work Block 3 results and explicitly forbids using it as the source of evaluation meaning (`plans/skill-issue-project-completion/document-authority-and-update-map.md:35-51`).
- Result-schema changes route first through the skill-calling supporting plan and CLI export behavior, invalidate affected campaign results, and only then update `siteData.ts`, methods, graphs, video, and submission from newly qualified artifacts (`plans/skill-issue-project-completion/document-authority-and-update-map.md:103-108`).

**Implication**

The source-of-truth boundary is layered: the evaluation contract and CLI artifact own what result fields mean; qualified Work Block 3 output supplies the evidence; `siteData.ts` remains the website's editable integration and display-content owner; and components remain presentation owners. Website code must not reinterpret transcripts, invent counts, or silently redefine called-versus-missed semantics.

### Finding 5 — The Future Migration Contract Is Specific And Already Has A Concrete Producer

The binding Work Block 4 migration replaces the mock context-percentage graph through `siteData.ts` with the compact evaluation artifact. The required chart uses a numeric `turn` domain from `1` through `total_turns`, plots `called` and `missed`, derives sample size, and calculates data-driven axis limits. The production CLI already writes the corresponding `website.json` after a completed evaluation and derives its points directly from expected and observed call evidence.

**Evidence**

- The current completion plan records that the website mock-up still has mock release and benchmark content and that accepted campaign results do not yet exist (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:11-14`, `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:21-30`).
- Work Block 3 must aggregate inspectable descriptive graph-ready data before Work Block 4 consumes it (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:71-83`).
- Work Block 4 explicitly requires replacement through `src/data/siteData.ts`, preservation of chart and theme owners, direct consumption of the compact artifact, the `1..total_turns` numeric domain, `turn`/`called`/`missed`, derived sample size, and data-driven axis limits (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:85-97`).
- The current supporting plan defines compact fields `schema_version`, `run_id`, `scenario_id`, `harness`, `model`, `total_turns`, and `points`; each point contains `turn`, `turn_id`, `called`, and `missed` (`plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:49-56`).
- Production Go types match that contract exactly (`cli/internal/evaluation/evaluation.go:98-113`).
- A tooling-complete run always writes `result.json` followed by compact `website.json` (`cli/internal/evaluation/evaluation.go:285-310`).
- Point derivation walks scenario turns in order, skips turns without expected calls, counts each expected skill observed on that turn, computes `missed` as expected minus called, sets the one-based `turn`, and records the complete scenario length as `total_turns` (`cli/internal/evaluation/evaluation.go:696-737`).

**Implication**

The migration has no need for transcript parsing or a second result interpretation layer. The compact artifact is structurally ready for Recharts, although website publication remains gated on qualified Work Block 3 campaign evidence. Because points omit turns with no expected calls while `total_turns` covers the full scenario, the chart must use `total_turns` for its complete horizontal domain rather than inferring the domain from the last point.

### Finding 6 — Methodology And Public Result Copy Must Move From Preview Language To Inspectable Evidence

The current surrounding copy is intentionally provisional: the results heading describes context consumption, each graph calls itself mock data, and the method paragraph says the harness is still being finalized. The finished-product contract requires descriptive graphs without project-authored pass/fail labels, alongside transparent methods, environment information, raw structured results, limitations, and graph derivation. Work Block 4 also adds the approved product story, About content, tested environments, and limitations while keeping the existing data and component owners.

**Evidence**

- Current results copy describes “successful calls and missed opportunities” as available context is consumed (`src/App.tsx:76-84`).
- Current graph descriptions identify both evaluations as mock comparisons (`src/data/siteData.ts:41-74`).
- Current method copy defines checkpoint calls and misses but says the mock data only illustrates report shape while the harness is finalized (`src/data/siteData.ts:75-80`).
- The completed product/evidence foundation requires graphs without pass/fail labels and publication of method, environment instructions, instrumentation, raw structured results, limitations, and derivation (`plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md:59-68`).
- The skill-calling supporting plan requires strong and weak behavior to remain observable graph data rather than a project-authored verdict and requires scenario, expected-call, instrumentation, environment, structured-result, derivation, and limitation disclosure (`plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:67-72`).
- Work Block 4 requires product story, About content, tested environments, limitations, real repository/release routes, and exact consistency between repository, release, raw results, website, downloads, narrative, and limitations (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:89-95`).

**Implication**

The migration must update both graph-adjacent copy and the method/public-evidence presentation from qualified artifacts. Those changes still flow through `siteData.ts` for editable website copy and through `App` for page composition; the evaluation plan remains the semantic owner of what called, missed, valid, and publishable mean.

### Finding 7 — Visual Consistency Is A Preserved Contract, Not A New Design Exercise

The existing website foundation is already considered complete for responsive design and visual audit. Its binding direction is a restrained neutral system with one semantic accent family, compact navigation, centered hero, thin bordered surfaces, responsive data cards, quiet type, generous rhythm, and equal top-level light/dark treatment. The chart participates in that system through shared CSS tokens, semantic call/miss colors, mono metadata, modest borders, and limited elevation.

**Evidence**

- The selected direction binds the neutral palette, compact navigation, centered hero, thin bordered surfaces, responsive card hierarchy, quiet typography, and concise labels (`plans/website/reference-and-architecture.md:24-41`).
- Global CSS tokens own color, typography, spacing, radii, shadow, content bounds, and transition timing; components are required to consume them unless a value is unique geometry (historical `plans/website/data-and-design-contract.md:36-48` at commit `22445be`).
- Theme switching retains the same semantic success/miss meanings and chart strokes, grid, and labels read from shared CSS variables (historical `plans/website/data-and-design-contract.md:50-57` at commit `22445be`).
- Production CSS provides light and dark semantic chart variables (`src/styles.css:1-16`, `src/styles.css:43-59`) and chart-specific components consume them (`src/components/EvaluationChart.tsx:42-88`).
- Direct image comparison at research time found no actionable differences in responsive geometry, borders, radii, elevation, section rhythm, mono labels, theme coherence, focus visibility, or semantic graph treatment.
- The reorganization audit classifies `plans/website/` as the authoritative mock-up design, validation, and visual-audit record and directs Work Block 4 to update the existing owners instead of rebuilding the website (`plans/skill-issue-project-completion/reorganization-dependency-audit.md:44-48`, `plans/skill-issue-project-completion/reorganization-dependency-audit.md:50-62`).

**Implication**

Real-result migration should preserve the approved composition and token system. New result volume, longer harness/model labels, tested-environment material, limitations, and About content may extend the page, but visual decisions remain owned by the established CSS and component contracts and should retain the audited responsive hierarchy, restrained surfaces, semantic colors, and equal light/dark behavior.

## Notes

- `siteData.chart.yAxis` is declared as “Skill decisions” (`src/data/siteData.ts:35-40`) but the production chart does not render that string; the current Y-axis renders only numeric ticks (`src/components/EvaluationChart.tsx:54-60`). This is an existing display caveat, not evidence of a separate ownership contract.
- The current accessible chart summary describes only the first and final mock points and hardcodes the 100-percent endpoint (`src/components/EvaluationChart.tsx:32-35`). The future turn-based artifact contains sparse expected-call turns plus a separate full `total_turns` domain, so that exact summary cannot survive unchanged.
- Screenshots were not re-inspected for this assignment. The completed pairwise image audit already records the binding desktop and narrow visual findings, and the production CSS provides the exact current responsive behavior.
- No accepted campaign artifact is yet authoritative for public website data; the completion plan explicitly records that accepted cells and real website results remain unfinished (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:14-14`, `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:29-30`).

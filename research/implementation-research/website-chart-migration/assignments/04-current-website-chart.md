# Current Website Chart

## Assignment

**Goal:** Record exactly what data the current React/Vite website passes to its mock evaluation charts, how `EvaluationChart` renders that data, and which concrete mismatches matter when replacing the mock graphs with CLI-produced `website.json` artifacts.

**Scope:** Production website sources, the production CLI serializer that owns `website.json`, generated local artifact examples, the locked and installed Recharts package, and chart-specific CSS.

**Exclusions:** Product-code changes, test-derived behavior, visual redesign recommendations, data-loading architecture, and assumptions about an unpublished aggregation contract.

## Sources

- `src/data/siteData.ts`: `EvaluationPoint`, `EvaluationGraph`, `siteData.chart`, and `siteData.evaluations` (lines 1-80).
- `src/components/EvaluationChart.tsx`: `EvaluationChartProps` and `EvaluationChart` (lines 1-108).
- `src/App.tsx`: imports and the `siteData.evaluations.map` call site (lines 1-4 and 71-91).
- `src/styles.css`: light/dark chart tokens and chart card, legend, tooltip, and responsive rules (lines 1-60 and 360-479, 515-605).
- `package.json`: declared `recharts` dependency `^3.9.2` (lines 16-20).
- `package-lock.json`: locked `node_modules/recharts` version `3.9.2` and tarball (lines 2488-2491).
- `node_modules/recharts/package.json`: installed Recharts version `3.9.2`, ESM entry, and type entry (lines 1-9).
- `node_modules/recharts/es6/index.js`: concrete public exports used by the component (`Tooltip`, `ResponsiveContainer`, `CartesianGrid`, `Line`, `XAxis`, `YAxis`, and `LineChart`; lines 10-12, 36-46).
- `node_modules/recharts/es6/chart/LineChart.js`: `LineChart` delegates to `CartesianChart` with axis tooltips (lines 1-20).
- `node_modules/recharts/es6/chart/CartesianChart.js`: default horizontal layout and enabled accessibility layer (lines 17-33, 41-83).
- `node_modules/recharts/es6/chart/CategoricalChart.js`: chart props flow to `RootSurface` (lines 1-60).
- `node_modules/recharts/es6/cartesian/XAxis.js`: `XAxis` resolves `xAxisDefaultProps` from `implicitXAxis` (lines 105-165).
- `node_modules/recharts/es6/cartesian/YAxis.js`: `YAxis` resolves `yAxisDefaultProps` from `implicitYAxis` (lines 139-165).
- `node_modules/recharts/es6/state/selectors/axisSelectors.js`: installed X-axis category defaults and Y-axis number defaults (lines 44-124).
- `node_modules/recharts/es6/container/RootSurface.js`: default focusable SVG application semantics (lines 29-62, 81-95).
- `node_modules/recharts/es6/component/Tooltip.js` and `node_modules/recharts/es6/component/DefaultTooltipContent.js`: default tooltip interaction and live-region behavior (`Tooltip.js` lines 29-67, 82-175; `DefaultTooltipContent.js` lines 61-150).
- `cli/internal/evaluation/evaluation.go`: production `WebsitePoint`, `WebsiteResult`, artifact write, and `deriveWebsiteResult` implementation (lines 98-113, 304-310, 696-737).

## Findings

### Finding 1: The website owns a mock-only graph shape

`EvaluationPoint` is `{ contextConsumed: number; skillCalls: number; skillMisses: number }`. `EvaluationGraph` adds `{ id, harness, model, description, sampleSize, points }`, with all metadata fields except `sampleSize` expressed as strings and `points` expressed as `EvaluationPoint[]`. `siteData.evaluations` is checked with `satisfies EvaluationGraph[]`, and `App` maps each graph directly into `<EvaluationChart key={graph.id} graph={graph} />`; there is no parsing, normalization, fetching, or intermediate adapter in the current call path.

**Evidence:** `src/data/siteData.ts:1-14,41-74`; `src/App.tsx:1-4,87-90`; `src/components/EvaluationChart.tsx:11-18`.

**Implication:** A CLI `website.json` object cannot be passed directly to `EvaluationChart`. The website currently requires renamed point fields plus card metadata that the CLI artifact does not contain (`id`, `description`, and `sampleSize`).

### Finding 2: Two mock cards each plot six percentage checkpoints

The first graph is `codex-gpt-5-6` / `Codex CLI` / `GPT-5.6`, describes a mock explicit-skill-selection evaluation, declares `sampleSize: 120`, and contains `(contextConsumed, skillCalls, skillMisses)` rows `(0,19,1)`, `(20,18,2)`, `(40,17,3)`, `(60,16,4)`, `(80,14,6)`, and `(100,12,8)`. The second is `claude-code-opus` / `Claude Code` / `Claude Opus`, also declares `sampleSize: 120`, and contains `(0,18,2)`, `(20,17,3)`, `(40,16,4)`, `(60,14,6)`, `(80,12,8)`, and `(100,9,11)`. Every checkpoint totals 20 decisions; six checkpoints therefore correspond to the displayed `n = 120`.

**Evidence:** `src/data/siteData.ts:41-74`; `src/components/EvaluationChart.tsx:21-30`.

**Implication:** The present visual encodes cumulative context percentage on X and per-checkpoint call/miss counts on Y. Its evenly spaced 0-100 rows conceal how the component behaves when numeric X values have irregular gaps.

### Finding 3: The X axis is categorical even though its values are numbers

`XAxis` receives `dataKey="contextConsumed"`, a percent tick formatter, no `type`, no domain, and no explicit ticks. In installed Recharts 3.9.2, `XAxis` resolves its omitted `type` from `implicitXAxis`, whose value is `category`; its scale is `auto` and duplicated categories are allowed. Consequently, the six numeric percentages are treated as category values and placed as categorical checkpoints, rather than positioned on a proportional numeric scale. By contrast, `YAxis` defaults to `type: 'number'` and the component explicitly fixes its domain to `[0, 20]` and ticks to `[0, 5, 10, 15, 20]`.

**Evidence:** `src/components/EvaluationChart.tsx:47-60`; `node_modules/recharts/es6/cartesian/XAxis.js:105-165`; `node_modules/recharts/es6/state/selectors/axisSelectors.js:44-79,96-124`; `node_modules/recharts/es6/cartesian/YAxis.js:139-165`.

**Implication:** Mapping CLI `turn` directly onto the current X axis will space emitted points equally even when `deriveWebsiteResult` omits turns without expected calls. If distance along X must represent the actual one-based turn position, the migration must explicitly choose numeric-axis behavior. The fixed Y maximum of 20 can clip larger call/miss counts and waste vertical range for smaller runs.

### Finding 4: Labels are split between chart metadata, tooltip names, and footer copy

The chart copy object defines `xAxis: 'Context consumed'`, `yAxis: 'Skill decisions'`, `calls: 'Skill calls'`, and `misses: 'Skill misses'`. The line `name` props reuse the calls/misses labels, so the default tooltip shows those names with unformatted numeric values. The tooltip label is formatted as `<value>% context consumed`. The X-axis title is rendered only as footer text, while `siteData.chart.yAxis` is never referenced anywhere under `src`; neither Recharts axis receives a `label` prop.

**Evidence:** `src/data/siteData.ts:35-40`; `src/components/EvaluationChart.tsx:47-88,93-105`; repository search for `chart.yAxis` returns no use outside its definition; `node_modules/recharts/es6/component/DefaultTooltipContent.js:61-150`.

**Implication:** A turn-based migration must update all semantic surfaces together: X ticks, tooltip label, footer title, line names, and the prose/accessibility label. Merely renaming point keys would leave percentage-based interpretation in the UI. The existing `yAxis` string provides no visible Y-axis title to preserve.

### Finding 5: Two smoothed series share a manually constrained presentation

Both lines use monotone interpolation, `strokeWidth={2.5}`, no persistent dots, and a radius-4 active dot. Calls use `--color-call`; misses use `--color-miss`. Those tokens resolve to `#1d8a6a` and `#d05663` in light mode and `#53cda6` and `#ff7b88` in dark mode. The grid uses only horizontal dashed lines (`3 6`); both axes suppress axis and tick lines. Recharts remains free to apply its default auto animation because neither line disables `isAnimationActive`.

**Evidence:** `src/components/EvaluationChart.tsx:38-88`; `src/styles.css:1-16,43-58`; `node_modules/recharts/es6/cartesian/Line.js:116-137,434-487`.

**Implication:** Replacement data inherits smoothing, active-point, animation, fixed-scale, and theme behavior automatically. Sparse turn points may appear as a continuous trend even though the source artifact records discrete expected-call turns.

### Finding 6: Each graph is a responsive card, with a separate hand-built legend

The card header shows model as a kicker, harness as the heading, and `n = sampleSize`; a description precedes a fixed-height chart (`21rem`, reduced to `17rem` below 680 px). Cards render in a two-column grid and collapse to one column below 680 px. The footer legend is ordinary markup with two empty `<i>` color swatches and visible calls/misses text; the component does not import or render Recharts `Legend`. The footer also repeats the X-axis title.

**Evidence:** `src/components/EvaluationChart.tsx:21-37,93-106`; `src/styles.css:388-479,515-605`; `src/components/EvaluationChart.tsx:1-9` contains no `Legend` import.

**Implication:** CLI metadata can fill harness and model directly, but card identity, description, and sample-size derivation need an explicit adapter or revised component contract. Series naming and color changes must keep the manual legend synchronized because Recharts will not own it.

### Finding 7: Accessibility combines a static summary with Recharts keyboard semantics

The outer `.chart-wrap` is `role="img"` and receives a static `aria-label` claiming calls fall, misses rise, and context reaches 100 percent, using only the first and final points. `finalPoint` is asserted non-null and immediately dereferenced, so an empty point array is unsupported at runtime. Separately, Recharts 3.9.2 enables its accessibility layer by default; the nested root SVG defaults to `role="application"` and `tabIndex={0}`, and the tooltip is exposed as an assertive live `status` region when active. Tooltip activation defaults to hover, while the accessibility layer also supports keyboard-driven chart interaction. The hand-built legend has `aria-label="Chart legend"` but no explicit landmark/list role.

**Evidence:** `src/components/EvaluationChart.tsx:17-18,32-36,61-70,93-103`; `node_modules/recharts/es6/chart/CartesianChart.js:23-33,41-83`; `node_modules/recharts/es6/container/RootSurface.js:29-62`; `node_modules/recharts/es6/component/Tooltip.js:42-67,82-175`; `node_modules/recharts/es6/component/DefaultTooltipContent.js:133-150`.

**Implication:** Direct CLI data must contain at least one plotted point or the current component fails. The static summary also requires migration: real runs can be flat, rise, fall, or contain a single point, and the CLI X coordinate ends at its last expected-call turn rather than necessarily at `total_turns`. The existing fixed wording would describe such runs inaccurately.

### Finding 8: The CLI artifact is turn-based and omits several card fields

Production `website.json` schema version 1 is `{ schema_version, run_id, scenario_id, harness, model, total_turns, points }`, with each point shaped as `{ turn, turn_id, called, missed }`. `deriveWebsiteResult` emits a point only for a scenario turn with at least one expected skill, preserves the turn's one-based scenario index and source ID, counts unique expected skills observed on that turn as `called`, and derives `missed` as expected minus called. The artifact has no context percentage, graph description, separate card ID, or stored sample size.

**Evidence:** `cli/internal/evaluation/evaluation.go:98-113,304-310,696-737`.

**Implication:** A faithful migration needs to map `turn -> contextConsumed` or replace that field, `called -> skillCalls`, and `missed -> skillMisses`; derive sample size as the sum of `called + missed`; choose `run_id` or another stable composite for the React key; and supply or remove description. `scenario_id` and `total_turns` currently have no card surface. Because only expected-call turns become points, the adapter must preserve `total_turns` separately if the chart or accessible summary needs the evaluation's true endpoint.

### Finding 9: The package constraint and concrete runtime are aligned at Recharts 3.9.2

The website declares `recharts: ^3.9.2`; the lockfile resolves `node_modules/recharts` to exactly `3.9.2`, and the installed package reports version `3.9.2`. The component's named imports resolve through the package ESM index to their concrete Recharts modules, with `LineChart` delegating to `CartesianChart` and the categorical chart surface.

**Evidence:** `package.json:16-20`; `package-lock.json:2488-2491`; `node_modules/recharts/package.json:1-9`; `node_modules/recharts/es6/index.js:10-12,36-46`; `node_modules/recharts/es6/chart/LineChart.js:1-20`; `node_modules/recharts/es6/chart/CartesianChart.js:41-83`; `node_modules/recharts/es6/chart/CategoricalChart.js:1-60`.

**Implication:** Migration work should be designed and validated against Recharts 3.9.2 behavior. The caret declaration permits a later compatible release during a fresh install, while the current lockfile and workspace concretely use 3.9.2.

## Notes

- The current mock descriptions and method copy explicitly identify the data as mock/preview data (`src/data/siteData.ts:17,46-47,62-63,75-80`). Any migration that loads real CLI artifacts must also reconcile that surrounding copy, although the loading mechanism is outside this assignment.
- No tests were used to establish the website or CLI contracts. Generated `website.json` files were inspected only as concrete serialization examples after tracing the production structs and derivation path.

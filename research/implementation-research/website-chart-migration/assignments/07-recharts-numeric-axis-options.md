# Recharts Numeric Axis Options

## Assignment

- **Goal:** Determine how the repository's installed Recharts version can render a numeric turn axis, separate called and missed series, and comparison-oriented chart UI.
- **Scope:** Inspect the repository's installed package metadata, current chart implementation, exact installed declarations/build output, and official Recharts documentation or repository source for numeric axes, lines, tooltips, legends, responsive sizing, accessibility, and chart synchronization.
- **Exclusions:** Page-wide design, evaluation-methodology decisions, choosing the final comparison layout, chart-library migration, and product-code edits.

## Sources

### Local repository and installed package

- `package.json` — declares `recharts` as `^3.9.2`.
- `package-lock.json` — resolves `node_modules/recharts` to exact version `3.9.2`.
- `node_modules/recharts/package.json` — confirms installed version `3.9.2`, React 19 peer support, and package entry points.
- `src/components/EvaluationChart.tsx` — current two-line chart, tooltip, responsive wrapper, external legend, and static accessibility summary.
- `src/data/siteData.ts` — current graph/point shapes and called/missed labels.
- `src/styles.css` — current responsive parent sizing: `.chart-wrap` is `21rem` high and becomes `17rem` below 680px.
- `node_modules/recharts/types/cartesian/XAxis.d.ts` — installed numeric-axis props and defaults.
- `node_modules/recharts/types/cartesian/Line.d.ts` — installed `Line` data, dot, null-connection, name, and animation contracts.
- `node_modules/recharts/types/component/ResponsiveContainer.d.ts` and `node_modules/recharts/es6/component/ResponsiveContainer.js` — installed responsive sizing contract and implementation.
- `node_modules/recharts/types/util/types.d.ts`, `node_modules/recharts/es6/chart/CartesianChart.js`, and `node_modules/recharts/es6/container/RootSurface.js` — installed chart data, synchronization, and accessibility contracts/defaults.

### Official Recharts documentation and exact-version source

- [XAxis API](https://recharts.github.io/en-US/api/XAxis/)
- [Axis ticks guide](https://recharts.github.io/guide/axisTicks/)
- [Domain and ticks guide](https://recharts.github.io/en-US/guide/domainAndTicks/)
- [Line API](https://recharts.github.io/en-US/api/Line/)
- [LineChart API](https://recharts.github.io/en-US/api/LineChart/)
- [Tooltip API](https://recharts.github.io/en-US/api/Tooltip/)
- [Legend API](https://recharts.github.io/en-US/api/Legend/)
- [ResponsiveContainer API](https://recharts.github.io/en-US/api/ResponsiveContainer/)
- [Synchronized Line Chart example](https://recharts.github.io/en-US/examples/SynchronizedLineChart/)
- [Recharts and accessibility](https://github.com/recharts/recharts/wiki/Recharts-and-accessibility)
- [Recharts animations guide](https://recharts.github.io/en-US/guide/animations/)
- [Recharts 3.9.2 `XAxis` source](https://github.com/recharts/recharts/blob/v3.9.2/src/cartesian/XAxis.tsx)
- [Recharts 3.9.2 `Line` source](https://github.com/recharts/recharts/blob/v3.9.2/src/cartesian/Line.tsx)
- [Recharts 3.9.2 Cartesian chart defaults](https://github.com/recharts/recharts/blob/v3.9.2/src/chart/CartesianChart.tsx)
- [Recharts 3.9.2 root surface accessibility](https://github.com/recharts/recharts/blob/v3.9.2/src/container/RootSurface.tsx)
- [Recharts 3.9.2 `ResponsiveContainer` source](https://github.com/recharts/recharts/blob/v3.9.2/src/component/ResponsiveContainer.tsx)
- [Recharts 3.9.2 shared chart prop types](https://github.com/recharts/recharts/blob/v3.9.2/src/util/types.ts)

## Findings

### Installed 3.9.2 already contains the required primitives

The repository's declared range and lockfile both start at Recharts 3.9.2, and the installed package is exactly 3.9.2. The current component already uses `LineChart`, two `Line` children, `XAxis`, `YAxis`, `Tooltip`, and `ResponsiveContainer`, so the numeric-turn migration can stay within the installed library and the existing component family.

**Evidence:** `package-lock.json` resolves `recharts-3.9.2.tgz`; `node_modules/recharts/package.json` reports `3.9.2`. The current `EvaluationChart.tsx` supplies one data array to a `LineChart` and renders `skillCalls` and `skillMisses` as separate `Line` children. Recharts documents `LineChart` as the parent for `Line`, `XAxis`, `YAxis`, `Tooltip`, and `Legend`, and its [Line API](https://recharts.github.io/en-US/api/Line/) explicitly describes extracting multiple dimensions from shared data with separate `dataKey` values.

**Implication:** No Recharts upgrade or library replacement is required to express `{ turn, called, missed }` points and render called/missed lines against turns.

### A true turn axis requires `type="number"`

`XAxis` defaults to a categorical axis. A categorical axis spaces adjacent values equally even when turn numbers have gaps; a numeric axis uses a continuous range and places numerically closer values closer together. The chart therefore needs a numeric `turn` field plus `dataKey="turn"` and `type="number"` for turn number to carry quantitative spacing.

**Evidence:** The [XAxis API](https://recharts.github.io/en-US/api/XAxis/) defines `category` as equally spaced distinct values, `number` as a continuous range, and the default as `category`. The [domain and ticks guide](https://recharts.github.io/en-US/guide/domainAndTicks/) also demonstrates that numeric axes require numeric input values. These contracts are present in the exact [3.9.2 `XAxis` source](https://github.com/recharts/recharts/blob/v3.9.2/src/cartesian/XAxis.tsx) and installed `XAxis.d.ts`.

**Implication:** The minimum implementation shape is conceptually `XAxis dataKey="turn" type="number"`. Turn identifiers must be numbers rather than display strings such as `"Turn 4"`; presentation belongs in `tickFormatter` and tooltip formatting.

### Domain and tick controls support integer turns, with distinct responsibilities

For a numeric axis, `domain` controls the numeric extent, `allowDecimals={false}` prevents generated decimal tick values, and `ticks` provides exact positions. `tickCount` asks Recharts to generate roughly that many numeric ticks; it does not provide the same determinism as an explicit integer array. Recharts 3.9.2 also supports `niceTicks`, but explicit `ticks` remains the full-control option. Tick labels may still be hidden for overlap because the axis default `interval` is `"preserveEnd"`; `interval={0}` forces every supplied label to render and therefore creates a width/overlap responsibility for the page.

**Evidence:** The [XAxis API](https://recharts.github.io/en-US/api/XAxis/) states that `allowDecimals` defaults to true, numeric-axis `ticks` must be numbers, `tickCount` applies to numeric axes, and `interval={0}` shows all ticks. The [axis ticks guide](https://recharts.github.io/guide/axisTicks/) says explicit `ticks` gives full control over tick positions and labels. The API accepts fixed, data-derived, and functional numeric domains; `allowDataOverflow` defaults to false and otherwise expands a specified domain to include out-of-range data, while `allowDataOverflow={true}` clips graphics to the fixed domain.

**Implication:** For a bounded turn range, a deterministic configuration can use a fixed or data-derived numeric domain, `allowDecimals={false}`, and a deliberately generated integer/subset `ticks` array. Showing every turn is appropriate only when the available width can carry every label; otherwise use a meaningful subset while preserving all data points.

### Separate called and missed series are direct `Line` mappings

Each `Line` independently extracts one value from the shared point object. String `dataKey` values such as `called` and `missed` are the simplest contract, and each line's `name` becomes its human-facing label in tooltip and legend content. Recharts does not derive called/missed values or enforce that they are complementary; it renders the values supplied by the evaluation-data adapter.

**Evidence:** The [Line API](https://recharts.github.io/en-US/api/Line/) defines `dataKey` as a string, number, or extraction function and says different Y-axis dimensions normally use different data keys from the same objects. It defines `name` as the tooltip/legend label, falling back to `dataKey` when omitted. The current component already proves the same structural pattern with `skillCalls` and `skillMisses`.

**Implication:** A point shape such as `{ turn: number, called: number | null, missed: number | null }` is sufficient. A binary outcome view can use a shared `[0, 1]` Y domain and integer ticks; a count view can use the actual count range. The final aggregation must keep that choice aligned with the evaluation artifact's semantics because Recharts will not validate it.

### Dot and null behavior must match turn-level evidence semantics

`Line` defaults to rendering a dot at each point and leaving gaps across null coordinates. The current chart explicitly disables ordinary dots but keeps an active dot on interaction. `connectNulls={true}` would bridge missing observations, whereas the default `false` leaves a visual break. A numeric zero remains a real rendered observation and is not treated as missing.

**Evidence:** The [Line API](https://recharts.github.io/en-US/api/Line/) documents `dot` defaulting to true, `activeDot` defaulting to true, and `connectNulls` defaulting to false. The exact [3.9.2 `Line` source](https://github.com/recharts/recharts/blob/v3.9.2/src/cartesian/Line.tsx) carries those defaults. `EvaluationChart.tsx` currently sets `dot={false}` and a custom `activeDot`, while leaving `connectNulls` at its false default.

**Implication:** Keep `connectNulls={false}` when a missing turn means absent or invalid evidence; bridging it would imply continuity the data does not establish. For sparse, per-turn events, visible dots can make the observed turns clearer. If dots remain hidden for density, the active dot and tooltip still provide interactive point location. The chosen line interpolation must also match the semantic claim: the current `type="monotone"` smooths between observations, while 3.9.2's `Line` default is linear.

### Tooltip formatting can make numeric turns and outcomes legible

The default tooltip can show both line payload entries at an active turn. `labelFormatter` can convert the numeric X value to `Turn N`; `formatter` can convert a numeric outcome to a label and may return `[formattedValue, formattedName]`. Per-line `formatter` takes precedence over the tooltip-level formatter. Tooltip entries default to alphabetical sorting by line `name`, so JSX child order alone is not a stable display-order contract.

**Evidence:** The [Tooltip API](https://recharts.github.io/en-US/api/Tooltip/) defines `labelFormatter`, `formatter`, null filtering, and the default `itemSorter="name"`. The [Line API](https://recharts.github.io/en-US/api/Line/) states that a line formatter overrides the tooltip formatter. Installed `DefaultTooltipContent` types and implementation expose the same formatter/labelFormatter/itemSorter behavior.

**Implication:** Use explicit line `name` values (`Called`, `Missed`) and a turn-aware `labelFormatter`. If the evaluation values are binary, the tooltip can present `Called: yes/no` and `Missed: yes/no` instead of unexplained `1`/`0`, but this wording must follow the evaluation semantics. If item order is important, configure it deliberately rather than relying on line declaration order.

### Built-in and external legends are both supported, with different layout effects

Adding `Legend` creates entries from the line metadata and supports label formatting or fully custom HTML content. Alignment affects chart space: left/right legends reduce chart width, top/bottom legends reduce chart height, and middle alignment overlays content. The repository currently uses an external HTML legend below the responsive chart instead of a Recharts `Legend`.

**Evidence:** The [Legend API](https://recharts.github.io/en-US/api/Legend/) documents `formatter`, custom HTML `content`, layout/alignment effects, and default alphabetical sorting by legend value. The current footer in `EvaluationChart.tsx` renders its own call/miss swatches and labels outside `LineChart`.

**Implication:** The existing external legend is compatible with the migration and preserves exact page layout/control order. A built-in `Legend` is useful when the line payload should generate labels automatically, but its reserved chart space must be included in responsive sizing. Interactive comparison controls should remain explicit application UI; a legend formatter only changes presentation.

### Faceted charts can synchronize interaction, but numeric-turn synchronization has a constraint

Recharts charts sharing a `syncId` synchronize tooltip and brush events. The default `syncMethod="index"` assumes equal-length, equally ordered data arrays. Official Recharts documentation describes `syncMethod="value"` as matching the categorical axis, not an arbitrary numeric X axis. For numeric-turn facets, index synchronization is reliable only when every facet has the same ordered turn list; otherwise a custom synchronization function would be needed and should be validated against the actual datasets.

**Evidence:** The [LineChart API](https://recharts.github.io/en-US/api/LineChart/) and exact [3.9.2 shared chart types](https://github.com/recharts/recharts/blob/v3.9.2/src/util/types.ts) define `syncId`, index/value/custom synchronization, and the equal-length expectation for index mode. The official [synchronized line chart example](https://recharts.github.io/en-US/examples/SynchronizedLineChart/) demonstrates multiple charts with a shared `syncId`.

**Implication:** Facets are natively viable. Use a common numeric X/Y domain for visual comparability and, if synchronized hover is desired, normalize facets to the same turn sequence before relying on index synchronization. Treat value synchronization on a numeric X axis as unsupported by the cited contract unless a targeted runtime test proves the exact intended behavior.

### Filters and comparison toggles need explicit scale policy

Recharts can hide individual `Line` elements without removing their legend metadata. Axis domains normally derive from visible data; `includeHidden` allows hidden series to remain in domain calculation. A fixed domain avoids visual rescaling when comparison series are toggled.

**Evidence:** The [Line API](https://recharts.github.io/en-US/api/Line/) documents `hide` and notes that hidden graphical elements can remain in legends and can contribute to axis domains through `includeHidden`. The [XAxis API](https://recharts.github.io/en-US/api/XAxis/) documents `includeHidden`, and the corresponding Y-axis uses the same axis contract.

**Implication:** Comparison UI can map state to each line's `hide` prop. The product must decide whether filtering should preserve a common scale (`domain` fixed or `includeHidden`) or rescale to visible data. Stable domains are safer for cross-harness/model visual comparison.

### Responsive sizing is supported because the repository already supplies a measurable parent

`ResponsiveContainer` observes its own rendered container with `ResizeObserver` and supplies positive dimensions to its child chart. Its default dimensions are `100%` by `100%`, so the ancestor must establish usable width and height. The current `.chart-wrap` does so with fixed responsive heights.

**Evidence:** The [ResponsiveContainer API](https://recharts.github.io/en-US/api/ResponsiveContainer/) and exact [3.9.2 source](https://github.com/recharts/recharts/blob/v3.9.2/src/component/ResponsiveContainer.tsx) describe the `ResizeObserver` implementation. The installed implementation waits for positive dimensions before providing the chart. `src/styles.css` gives `.chart-wrap` `width: 100%` and a `21rem`/`17rem` height.

**Implication:** The numeric-axis migration can keep the current wrapper. Any future faceted or collapsible layout must preserve a positive measured height; mounting the chart in a zero-sized/hidden container can defer rendering until a later resize. Older environments without `ResizeObserver` require a polyfill according to Recharts, although this was not browser-matrix tested in this run.

### Recharts 3.9.2 adds an interactive accessibility layer by default

Cartesian charts in 3.9.2 default `accessibilityLayer` to true. The exact source gives the chart surface `role="application"` and `tabIndex={0}` unless callers override them, enabling keyboard point navigation and tooltip announcements. The repository simultaneously labels the outer `.chart-wrap` as `role="img"` with a static sentence, so the current implementation combines a static image-summary strategy with Recharts' focusable application strategy.

**Evidence:** The exact [3.9.2 Cartesian defaults](https://github.com/recharts/recharts/blob/v3.9.2/src/chart/CartesianChart.tsx) set `accessibilityLayer: true`; the [3.9.2 root surface](https://github.com/recharts/recharts/blob/v3.9.2/src/container/RootSurface.tsx) assigns the application role and focusability. The official [Recharts accessibility wiki](https://github.com/recharts/recharts/wiki/Recharts-and-accessibility) says 3.x enables the layer by default, uses arrow-key point navigation, and announces tooltip content. It also documents a VoiceOver caveat: users must turn QuickNav off for the arrow keys to reach the chart controls.

**Implication:** The migration should make one deliberate accessibility contract: retain the interactive Recharts layer and provide an accurate adjacent/static summary, or explicitly disable the layer and rely on a complete non-interactive alternative. The combined behavior of the current nested `role="img"` wrapper and focusable Recharts surface was not validated with screen readers in this run, so it should not be assumed correct. Any summary must be regenerated for turn-based data rather than retaining the current context-percentage sentence.

### Animation defaults respect reduced-motion preferences

Recharts 3.9.2 uses `isAnimationActive="auto"` for `Line` and `Tooltip`, disabling animation in server rendering and respecting the user's reduced-motion preference in the browser. Explicitly forcing `true` bypasses that preference.

**Evidence:** The [Recharts animations guide](https://recharts.github.io/en-US/guide/animations/) documents reduced-motion handling for `auto`; the installed `Line.d.ts` uses `auto` as its default.

**Implication:** The planned chart can keep default animation behavior. It should avoid setting `isAnimationActive={true}` unless bypassing reduced motion is intentional.

## Notes

- A small `renderToStaticMarkup` probe using installed Recharts 3.9.2, a numeric `XAxis`, explicit integer ticks, and two lines produced only an empty chart wrapper during server rendering. This confirms that SSR output is not a valid runtime probe for tick placement in this setup; it does not contradict the client-side API/source contracts.
- No product files were modified, and no live browser or assistive-technology session was run. Exact prop availability/defaults were cross-checked between installed declarations/build output, the `v3.9.2` official source tag, and official documentation.
- Numeric-axis `syncMethod="value"` was not validated and remains unsupported for this plan because the official 3.9.2 contract scopes value matching to a categorical axis.

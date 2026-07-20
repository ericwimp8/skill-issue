# Website Chart Migration Implementation Research

> **Status:** The chart migration described here has been implemented. Production source owns current behavior; this report remains evidence for the accepted-data handoff and future campaign-result export.

## Research Decision

The strongest implementation direction is to preserve the existing React/Vite/Recharts results surface, replace its mock percentage checkpoints with sparse numeric-turn evidence, and compare qualified harness/model/reasoning cells as aligned chart facets. Each visible chart should show one globally selected governed scenario, use a numeric X domain of `1..total_turns`, plot separate `called` and `missed` series at emitted expected-call turns, derive `n` from `sum(called + missed)`, and share a denominator-owned integer Y scale across visible cells.

The website should receive real results through one deterministic repository export that consumes a Work Block 3 accepted-campaign manifest plus its explicitly named `result.json`/`website.json` pairs, validates their identities and completeness, and writes a canonical checked-in TypeScript data module. `src/data/siteData.ts` should remain the human-edited integration and copy owner while importing the generated result projection. Publication should use the official GitHub Actions Pages artifact path: `npm ci`, `npm run validate`, upload `dist`, and deploy with the official Pages actions at the existing `/skill-issue/` project-site base.

The chart and static-publication direction is implemented, but real-result publication remains blocked by the absence of accepted Work Block 3 suite/campaign aggregation. Until that exists, the website uses only unmistakably illustrative layout data. Current smoke artifacts are useful serialization fixtures and cannot support public benchmark claims.

This synthesis is based on the constrained [research map](research-map.md) and [assignments 01–10](assignments/). It does not define an implementation plan or change product code.

## Evidence Categories

The findings must remain separated because they have different authority:

- **Implemented behavior:** the current CLI replay, instrumentation, scoring, artifact writers, React components, CSS, Vite configuration, and installed Recharts 3.9.2 behavior.
- **Controlling planned methodology:** the governed first-activation evaluation, thirteen-cell medium-setting campaign, Work Block 3 acceptance and aggregation, Work Block 4 website migration, public-claim limits, and existing website ownership boundaries.
- **Research recommendation:** the proposed chart facets, controls, illustrative fixture, deterministic refresh boundary, and GitHub Pages workflow. These choices are supported by the implemented data granularity and controlling methodology but are not yet implemented.

The production trace is documented in [Assignment 01](assignments/01-cli-evaluation-production-path.md), turn semantics in [Assignment 02](assignments/02-replay-scoring-turn-attribution.md), methodology and identity in [Assignment 03](assignments/03-methodology-and-result-identity.md), the current chart in [Assignment 04](assignments/04-current-website-chart.md), and website ownership/layout in [Assignment 05](assignments/05-website-layout-contracts.md).

## Implemented CLI Evidence Path

### Run and capture path

`skill-issue evaluate run` resolves one built-in evaluation or one custom scenario/answer-sheet pair, confirms the effective harness/model/reasoning interactively, creates one random 32-hex-character `run_id`, prepares the harness runtime, installs instrumented disposable skills, and replays one scenario through one primary resumable session. Before each prompt, the runner sets the scenario `turn_id` as the active turn; after terminal harness completion, it records the session and any Codex fallback signals, then clears the active turn. Turn attribution is therefore a CLI-owned replay boundary, not an inference from timestamps, transcript wording, token count, or model-generated labels. See `cli/internal/command/command.go:38-156`, `cli/internal/evaluation/evaluation.go:131-365`, `cli/internal/replay/replay.go:112-160`, and `cli/internal/runstate/runstate.go:125-191`, as traced in [Assignments 01](assignments/01-cli-evaluation-production-path.md) and [02](assignments/02-replay-scoring-turn-attribution.md).

The scoring input is the instrumented skill-entry signal. It proves that the evaluation instrumentation observed an activation for a skill during a turn. It does not establish skill-body compliance, output quality, task success, or general semantic usefulness. A signal outside an active turn remains unattributed. Codex additionally treats a structured attempted opaque signal command as evidence when the command could not directly write private state, while suppressing duplicate fallback recording for the same skill and turn. The detailed result therefore describes observed instrumentation, with harness-specific capture mechanics disclosed by the method.

### Detailed result

After replay, `result.json` preserves the atomic run identity and the detailed comparison:

- `schema_version`, `run_id`, harness, model, reasoning, `evaluation_id`, `scenario_id`, scope, and timestamps;
- expected, observed, missing, additional, and unattributed `{turn_id, skill}` records;
- optional `transcript_path` when transcript persistence succeeds.

Expected matches use the exact composite `{turn_id, skill}` key. Attributed events on other keys are additional; unattributed events cannot satisfy expectations; duplicate observed signals remain duplicated in the detailed arrays even though one matching event satisfies an expectation. The result contains no pass/fail verdict or aggregate score. It also omits important reproducibility fields such as CLI commit/version, native harness version, payload/input hashes, and operating system. See `cli/internal/evaluation/evaluation.go:79-96,642-687` and [Assignments 01–03](assignments/).

### Compact website result

`website.json` is an implemented lossy projection with this exact conceptual shape:

```json
{
  "schema_version": 1,
  "run_id": "...",
  "scenario_id": "...",
  "harness": "...",
  "model": "...",
  "total_turns": 30,
  "points": [{ "turn": 1, "turn_id": "1", "called": 1, "missed": 1 }]
}
```

The derivation walks the ordered scenario turns. For each turn containing at least one unique expected skill, it emits:

- `turn`: one-based position in the scenario array;
- `turn_id`: the source evidentiary identifier;
- `called`: number of unique expected skills observed on that exact turn;
- `missed`: unique expected skills for that turn minus `called`.

Turns without an expected call are omitted. Duplicate expected pairs or repeated signals do not inflate compact counts. `total_turns` remains the full scenario length, so the greatest emitted point can be less than the axis endpoint. For the current governed fixtures, points occur at turns `1`, `11`, `25`, and `30`; turn 1 has two expected first activations and the other points have one each, so one scenario has four plotted points and `n = 5`. See `cli/internal/evaluation/evaluation.go:98-113,696-738`, `cli/README.md:150-173`, and [Assignments 02](assignments/02-replay-scoring-turn-attribution.md) and [03](assignments/03-methodology-and-result-identity.md).

This file intentionally omits reasoning, evaluation identity, timestamps, environment, qualification, suite/campaign identity, skill names, additional/unattributed calls, diagnostics, and pass/fail labels. It is chart-shaped atomic evidence, not an accepted campaign record.

### Persistence and completeness limitations

The evaluator writes optional `events.jsonl` and `transcript.json`, then mandatory `result.json` and `website.json`, as separate direct `os.WriteFile` operations. The public run directory is not committed atomically. An interrupted write path can therefore leave a partial directory, including `result.json` without `website.json`; directory presence is not evidence of a complete or accepted run. Successful cleanup deletes private recovery state but preserves public artifacts. The structured stdout envelope returns the detailed result but does not identify the generated public directory or include `website.json`, so importers must receive explicit artifact paths rather than discover “the latest” directory. See `cli/internal/evaluation/evaluation.go:285-328,763-791` and [Assignments 01](assignments/01-cli-evaluation-production-path.md) and [10](assignments/10-result-data-update-workflow.md).

Optional transcripts contain complete prompts, responses, commands, stderr, raw events, session identifiers, and potentially sensitive local paths. Private `.skill-issue` run state also contains machine-local operational data. Neither should become a public website asset without explicit privacy review. `website.json` is the narrowest publication-oriented artifact, while `result.json` and reviewed raw events provide supporting evidence.

### Schema defect

The current invocation-event emitter includes `reasoning` and `evaluation_id`, but `evaluations/skill-calling/event.schema.json` sets `additionalProperties: false` and defines neither field. Current emitted events therefore do not conform to the committed schema as written. Raw event-schema conformance is unsupported until the semantic owner reconciles the code and schema and decides whether the correction remains schema version 1 or requires a version change. This does not alter the implemented `result.json` or `website.json` meaning, but it blocks claiming that current raw events are schema-valid campaign evidence. See `cli/internal/runstate/runstate.go:36-48,165-177`, `evaluations/skill-calling/event.schema.json:5-51`, and [Assignments 01](assignments/01-cli-evaluation-production-path.md) and [03](assignments/03-methodology-and-result-identity.md).

## Governed Methodology and Result Identity

The planned campaign measures required **first skill activations** fixed before observation. For each of the three governed 30-turn scenarios, the answer sheet expects `dictate-plan` and `document-update-discipline` on turn 1, `prompt-writing` on turn 11, `skill-authoring-discipline` on turn 25, and `system-change-ownership` on turn 30. Later applicability does not create another expected opportunity; a later reload is additional rather than a repair of an earlier miss. “Called” and “missed” must therefore be qualified as expected first activations, rather than all skill decisions or all applicable turns. See `evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:5-30` and [Assignments 02](assignments/02-replay-scoring-turn-attribution.md) and [03](assignments/03-methodology-and-result-identity.md).

The minimum campaign plans thirteen harness/model cells across five harness families at medium reasoning or the closest documented equivalent. Each cell is one tooling-complete suite of all three governed scenarios, with repeated trials and statistical thresholds deferred. Under the current fixtures that means 39 atomic scenario runs and 15 expected first activations per cell, but only one suite-level observation per cell. A valid zero-call suite remains descriptive model/harness evidence when replay and instrumentation completed; launch, permission, session, protocol, or marker failures are tooling failures that must be repaired and rerun rather than graphed as misses.

Publication needs three identity layers:

1. **Atomic run:** full `run_id`, `evaluation_id`, `scenario_id`, harness, exact effective model, effective reasoning, and the joined detailed/compact pair.
2. **Qualified cell suite:** a stable identifier binding exactly one accepted run for each governed scenario plus harness version, operating system, default-environment attestation, adjacent-configuration/contamination disclosures, and qualification status.
3. **Campaign:** a stable version identifying the matrix, fixture/schema versions, accepted cell suites, omissions or contract revisions, and publication date.

Only the atomic layer exists in production artifacts. Work Block 3 owns qualification, accepted-run selection, suite/campaign aggregation, environmental evidence, and graph-ready outputs. The website must consume that acceptance layer rather than infer suite membership from timestamps, directory adjacency, path names, display labels, or eight-character run prefixes. See `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:71-97`, `plans/skill-issue-project-completion/document-authority-and-update-map.md:103-108`, and [Assignment 03](assignments/03-methodology-and-result-identity.md).

The public claim boundary is descriptive: report exactly what was run and observed, with method, environment, derivation, raw structured evidence, and limitations. One suite per cell cannot establish statistical reliability, universal model behavior, persistent quality, or a general winner. Custom evaluations remain caller-owned and smoke probes establish interfaces rather than governed performance.

## Current Website Mismatch

The current site is a completed responsive mock-up whose data and rendering assumptions predate the compact artifact:

- `src/data/siteData.ts` defines points as `{contextConsumed, skillCalls, skillMisses}` and cards with manually supplied `id`, `description`, and `sampleSize`.
- Two explicitly mock cards contain six evenly spaced `0..100%` checkpoints and `n = 120` each.
- `EvaluationChart` uses `contextConsumed` on an omitted-type `XAxis`, so installed Recharts 3.9.2 treats the values categorically rather than proportionally.
- The Y axis is fixed at `0..20`; the lines use dotless monotone smoothing; the tooltip and footer describe context percentage; and the static accessibility sentence assumes the final point is 100% context.
- `siteData.chart.yAxis` is declared but not rendered.
- An empty point set is unsupported because the component dereferences the final point.
- Card headers, local HTML legend, theme colors, responsive wrapper, and page layout are established presentation owners.

A `website.json` object therefore cannot be passed directly to the current component. The migration must change the chart contract, axis type/domain/ticks, series keys, scale derivation, tooltip, legend qualification, `n` derivation, and accessibility summary together. It must retain `total_turns` outside the sparse points and supply editorial card description/evidence links separately. See `src/data/siteData.ts:1-80`, `src/components/EvaluationChart.tsx:17-105`, and [Assignment 04](assignments/04-current-website-chart.md).

The layout itself should be preserved. Results sit between summary metrics and method copy; cards use a two-column grid at wide and tablet widths, while the heading/method stack below 800px and cards become one column below 680px. Mobile chart height falls from `21rem` to `17rem`, dividers and footer orientation change, and the audited design requires no horizontal overflow, shared theme tokens, and consistent call/miss colors. `App.tsx` owns page composition, `siteData.ts` owns editable content integration, `EvaluationChart.tsx` owns chart presentation/accessibility, and `styles.css` owns layout/theme. See [Assignment 05](assignments/05-website-layout-contracts.md) and the historical `plans/website/data-and-design-contract.md` at commit `22445be`.

## Recommended Chart Contract

### Default organization

Use a small-multiple grid with one card per qualified harness/model/reasoning cell. Group cards harness-first, order them by the published campaign matrix, and show all qualified cells by default so weak or zero-call observations are not silently hidden. Every visible card should use the same globally selected governed scenario, numeric X domain, semantic series, and shared Y scale. The card header should expose harness, exact model, effective reasoning, scenario, derived `n`, and access to source evidence.

This preserves atomic identity and produces a readable comparison for a thirteen-cell matrix. A default overlay would require 26 traces before adding scenarios and would make sparse coincident points and legend identity unmanageable. See [Assignment 09](assignments/09-chart-comparison-design.md).

### Axis, marks, and scale

- Configure the X axis as numeric: `dataKey="turn"`, `type="number"`, integer ticks, and domain `1..total_turns`.
- Use meaningful integer tick subsets when width cannot carry every turn; do not convert turns to display strings in the data.
- Preserve only emitted expected-call points. Never insert zero-valued unscored turns, because zero would falsely claim an observed opportunity.
- Prefer straight line segments with persistent point markers. Connections are visual guides between measured turns, not interpolated observations. Remove monotone smoothing.
- Keep `connectNulls` false. Numeric zero is valid evidence; absent or invalid evidence should remain absent.
- Start the Y axis at zero and share its integer upper bound across visible cells. Derive the upper bound from the maximum `called + missed` denominator in the selected scenario/suite, not from observed performance. The current governed scenario view yields `0..2`.
- Keep the Y domain stable when filtering cards or hiding a series; independent per-card autoscaling would make equal heights represent different counts.

These primitives are available in installed Recharts 3.9.2; no library change is required. Official contracts are documented by [Recharts XAxis](https://recharts.github.io/en-US/api/XAxis/), [Line](https://recharts.github.io/en-US/api/Line/), [Tooltip](https://recharts.github.io/en-US/api/Tooltip/), and [ResponsiveContainer](https://recharts.github.io/en-US/api/ResponsiveContainer/), as synthesized in [Assignment 07](assignments/07-recharts-numeric-axis-options.md).

### Series, legend, tooltip, and accessibility

Use exactly two semantic series: **Called** and **Missed**, qualified in the chart heading or legend as expected first activations. Keep the established light/dark call and miss colors. Cell identity belongs in facet headers, not the legend. The existing local two-item HTML legend is the best default because each card remains self-contained during long mobile scrolling; a single shared desktop legend is a conditional alternative when each card retains accessible series names.

The tooltip should state the denominator explicitly, for example: “Turn 1 · 2 expected first activations · 1 called · 1 missed.” This prevents a two-opportunity point from being compared as though it had the same denominator as a one-opportunity point. `n` is the sum of all `called + missed`, not the number of points or raw signal events.

Retain Recharts' interactive accessibility layer only with a deliberate surrounding contract. Replace the current hard-coded context sentence with an accurate static summary or adjacent table/text that covers scenario, sparse measured turns, denominators, called/missed values, and total turn domain. The current nested outer `role="img"` plus focusable Recharts `role="application"` has not been screen-reader validated and should not be assumed correct. Recharts 3.9.2 enables its accessibility layer and reduced-motion-aware animation defaults; see the [Recharts accessibility guidance](https://github.com/recharts/recharts/wiki/Recharts-and-accessibility) and [Assignment 07](assignments/07-recharts-numeric-axis-options.md).

### Controls and comparison behavior

The smallest useful control set is:

- one global scenario selector;
- harness filter;
- optional model filter within visible harnesses;
- independent visibility toggles for Called and Missed.

Default to the first governed scenario in the campaign's stable published order until Work Block 3 defines a qualified suite-level aggregation. “All scenarios” should appear only when that aggregation preserves contributing run/scenario identities and documents its derivation. Reasoning is fixed at medium in the minimum campaign and belongs in metadata until multiple settings exist. Campaign version belongs in surrounding result identity, not a casual toggle. Do not rank or sort cells by called count.

If synchronized hover is desirable across facets, Recharts `syncId` with index synchronization is reliable only when facets share the same ordered turn list. Numeric-axis `syncMethod="value"` is unsupported by the cited 3.9.2 contract; otherwise use a validated custom function or omit synchronized hover. See [Assignment 07](assignments/07-recharts-numeric-axis-options.md).

### Mobile behavior

Preserve the existing breakpoint behavior: two columns through tablet width, one card per row at 680px and below, local headers and legends, a full numeric domain, and no horizontal chart scrolling. Controls should wrap or collapse into compact selectors. Long exact model identifiers and thirteen cells make filters useful for navigation, but mobile must not switch to a many-cell overlay or silently remove results. The four current governed X positions fit the existing shorter mobile chart when labels use an intentional subset.

## Illustrative Dataset Before Accepted Results

Until Work Block 3 publishes accepted results, use four fictional atomic artifacts: two fictional cells across two fictional 30-turn scenarios. Place this visible label directly with the charts:

> **Illustrative layout data — not observed evaluation results.**

Use identifiers that cannot be confused with real run IDs, real harnesses, or campaign assets. Each card has `n = 5` and points at the current governed positions:

| Illustrative scenario     | Illustrative cell           | Turn 1 called/missed | Turn 11 called/missed | Turn 25 called/missed | Turn 30 called/missed | Derived n |
| ------------------------- | --------------------------- | -------------------: | --------------------: | --------------------: | --------------------: | --------: |
| `illustrative-scenario-a` | Harness Alpha / Model Alpha |                1 / 1 |                 1 / 0 |                 0 / 1 |                 1 / 0 |         5 |
| `illustrative-scenario-a` | Harness Beta / Model Beta   |                2 / 0 |                 0 / 1 |                 1 / 0 |                 0 / 1 |         5 |
| `illustrative-scenario-b` | Harness Alpha / Model Alpha |                0 / 2 |                 1 / 0 |                 1 / 0 |                 0 / 1 |         5 |
| `illustrative-scenario-b` | Harness Beta / Model Beta   |                1 / 1 |                 1 / 0 |                 0 / 1 |                 1 / 0 |         5 |

Each illustrative record should retain only the compact artifact-style fields: `schema_version`, an unmistakably illustrative run identifier, `scenario_id`, harness, model, `total_turns: 30`, and points with `turn`, illustrative `turn_id`, `called`, and `missed`. It should not invent reasoning, environment qualification, suite identity, or campaign acceptance. This fixture exercises scenario selection, cross-cell facets, sparse numeric spacing, zero values, equal values, both series, and the two-opportunity first turn without suggesting likely real outcomes. See [Assignment 09](assignments/09-chart-comparison-design.md).

## Recommended Result Update Workflow

The best fit is a deterministic, fail-closed export into checked-in build-time data:

1. Input only the Work Block 3 accepted-campaign manifest or equivalent upstream acceptance artifact and its explicitly named `result.json`/`website.json` pairs.
2. Require complete artifact pairs and supported required fields; reject partial public directories.
3. Join detailed and compact artifacts by full `run_id`, compare harness/model/scenario identity, and obtain reasoning/evaluation/suite/campaign/environment acceptance from the authoritative detailed or manifest layer.
4. Reject duplicate run IDs, missing or invalid sparse points, inconsistent `total_turns`, and unaccepted suite/campaign membership.
5. Preserve atomic source identity and campaign provenance with display-ready points; derive `n` as `sum(called + missed)`.
6. Emit a canonical order by campaign cell, scenario, and turn so reruns produce stable review diffs.
7. Write one generated TypeScript result-data module that `satisfies` the website dataset type. Keep human copy, descriptions, release links, and method language curated in `siteData.ts`.

A generated TypeScript module is the lowest-risk representation under current checks because formatting, ESLint, and strict TypeScript already cover `src/**/*.ts`. Checked-in generated JSON is viable if deterministic formatting and schema validation become part of `npm run validate`, but current formatting scripts do not cover `src/**/*.json`. The serialization format is secondary to provenance validation and deterministic review. See [Assignment 10](assignments/10-result-data-update-workflow.md).

Do not parse transcripts, infer the newest run directory, scan all output folders as acceptance, or regenerate results during deployment from uncommitted external inputs. The accepted manifest plus generated diff should be the reviewable publication unit. This preserves Work Block 3 as the acceptance owner and avoids turning website code into a second evaluation system.

## Recommended GitHub Pages Publication

The repository is a Vite project site for `ericwimp8/skill-issue`; the current production base `/skill-issue/` matches the expected `https://ericwimp8.github.io/skill-issue/` path. The selected publication path should use GitHub Pages with GitHub Actions as the publishing source:

- trigger on reviewed pushes to `main` and retain `workflow_dispatch` for first deployment and safe reruns;
- use an explicit supported Node LTS, checkout, `npm ci`, and npm cache;
- run `npm run validate`, which already includes formatting, linting, type checking, and the production build;
- configure Pages, upload only `./dist` with `upload-pages-artifact`, and deploy it with `deploy-pages`;
- declare `contents: read`, `pages: write`, and `id-token: write`;
- target the `github-pages` environment and expose the deployment output URL;
- serialize deployments with a `pages` concurrency group;
- protect the environment so only the default branch deploys.

This matches the official [Vite GitHub Pages guidance](https://vite.dev/guide/static-deploy.html#github-pages), [GitHub custom Pages workflow guidance](https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages), and the repository's existing validation contract. `npm run preview` remains a local production-build inspection step and is not a deployment server. Action versions or immutable SHAs should be taken from the then-current official workflow at implementation time. See [Assignment 08](assignments/08-github-pages-vite-publication.md).

The repository currently has no Pages site, workflow, or Pages environment. Enabling GitHub Actions as the Pages source is therefore a one-time repository setting requiring maintainer access. Public launch also needs accepted result data and final release/download assets.

## Conditional Alternatives

- **Qualified suite-summary default:** preferred over per-scenario default only after Work Block 3 defines a transparent three-scenario aggregation preserving contributing runs and denominators.
- **Focused overlay:** allow at most two selected cells and one selected metric for close comparison; retain facets as the primary view.
- **Stacked discrete columns:** a defensible fallback when equal called/missed line marks overlap or straight connections remain misleading, because called and missed are complementary parts of each expected-turn denominator.
- **Normalized percentage view:** acceptable only as a secondary derived view paired with visible raw denominators; raw counts remain primary.
- **Checked-in generated JSON:** viable when deterministically generated and validated; marginally lower fit than TypeScript under current repository checks.
- **Manual-only Pages trigger:** safer until `main` reliably represents publication-ready accepted data, at the cost of public-site lag.
- **Branch publishing:** a fallback only if Actions is unavailable; it introduces committed build-output lifecycle and is lower fit for Vite.

## Rejected or Lower-Fit Paths

- Overlaying every cell's called and missed series.
- Categorically spacing numeric turns or inserting zero-valued unscored turns.
- Per-card Y autoscaling, default percentage normalization, or sorting/ranking cells by called count.
- Dotless monotone smoothing that implies continuous observations between sparse expected turns.
- Grouping by harness alone, display model label alone, timestamps, output paths, or run-ID prefixes.
- Treating one atomic `website.json`, a smoke/custom result, or a partial output directory as a qualified campaign result.
- Hand-maintaining all campaign numbers in `siteData.ts`.
- Copying raw `website.json` files without accepted-suite identity and detailed-result joins.
- Fetching result data at runtime for the MVP; this adds availability, caching, CORS, failure, and disclosure paths while weakening build review.
- Generating accepted data only inside the deployment job, committing `dist`, maintaining a `gh-pages` branch, or adding another hosting service.
- Publishing raw transcripts or private run state as website assets.

These paths either obscure evidence identity, imply measurements that were not made, weaken deterministic review, or overstate campaign maturity.

## Candidate Classification

### Selected deep dives

All ten research-map deep dives materially affected the recommendation: production evaluator/writer, replay and scoring, governed methodology, current chart/data contract, layout contracts, retained generated artifacts, Recharts 3.9.2 numeric-axis behavior, official Vite/GitHub Pages guidance, comparison design, and result refresh workflow. Their outputs are [Assignments 01–10](assignments/).

### Skim-only inputs

Historical hosting notes, adjacent tests, generic chart examples, package metadata beyond the installed-version contract, and prior screenshots were useful only as supporting validation. They did not establish behavior or change the primary direction. Tests were not used as behavioral authority; retained artifacts were inspected only after tracing current writers and were classified by provenance.

### Rejected candidates

Unrelated CLI lifecycle behavior, unrelated website sections, alternate hosting providers, chart-library migration, server-side dashboards, dynamic APIs, and implementation planning remained outside the decision boundary. None is necessary to map current CLI evidence into the established static website.

## True Blockers and Unsupported Claims

1. **Accepted Work Block 3 aggregation is missing.** There is no production schema or accepted artifact binding three atomic scenario runs into a qualified cell suite or cells into a campaign with environment evidence. This blocks real multi-cell chart data and public benchmark claims.
2. **No accepted campaign result exists.** Existing probe artifacts are smoke/isolation evidence only; no current qualified cell supports observed cross-harness/model performance claims.
3. **Raw event schema is inconsistent with the emitter.** `reasoning` and `evaluation_id` are emitted but forbidden by the committed schema, so current raw event schema conformance is unsupported.
4. **Public run directories are not atomic.** An importer must validate complete detailed/compact pairs and cannot treat directory presence as completion.
5. **The compact artifact lacks comparison and provenance fields.** Reasoning, evaluation identity, environment, suite/campaign acceptance, additional/unattributed calls, and timestamps require a joined upstream identity layer.
6. **The five-harness minimum is not represented by the current CLI evaluation support.** The production gate currently supports Claude Code, Codex, Cursor, and Pi; OpenCode remains outside the current evaluation-default/runtime set.
7. **Pages is not enabled and no workflow exists.** Maintainer access is needed for the one-time source/environment configuration, though the mechanism is established.
8. **Final public launch dependencies remain open.** Final release/download assets and correct repository/release URLs are still required.

Unsupported public claims include statistical reliability, universal model behavior, a general winner, pass/fail judgments, complete thirteen-cell coverage, native reasoning verification across every harness, or schema-valid raw events. The website may presently describe the evaluation system and show explicitly illustrative presentation data only.

## Close Evidence Summary

The core decision follows directly from four aligned facts:

1. Production already emits chart-ready sparse `{turn, turn_id, called, missed}` points and the full `total_turns` domain (`cli/internal/evaluation/evaluation.go:98-113,696-738`; [Assignment 02](assignments/02-replay-scoring-turn-attribution.md)).
2. Current Recharts 3.9.2 supports the required numeric axis, separate lines, explicit ticks/domains, responsive facets, tooltip formatting, stable hidden-series domains, and accessibility layer without a dependency change ([Assignment 07](assignments/07-recharts-numeric-axis-options.md)).
3. The existing website already owns a responsive card grid, semantic colors, local legends, theme behavior, and static data integration; its mismatch is semantic and typed rather than architectural ([Assignments 04](assignments/04-current-website-chart.md) and [05](assignments/05-website-layout-contracts.md)).
4. The missing piece is accepted suite/campaign identity, not chart rendering. A deterministic checked-in transform and official Pages build preserve reviewability once Work Block 3 supplies that authority ([Assignments 03](assignments/03-methodology-and-result-identity.md), [08](assignments/08-github-pages-vite-publication.md), and [10](assignments/10-result-data-update-workflow.md)).

Accordingly, the later implementation should migrate the established chart surface rather than redesign it, keep all pre-campaign values explicitly illustrative, and refuse real-result publication until the upstream acceptance and schema blockers are resolved.

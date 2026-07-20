# Website Chart Migration Research Map

## Run Frame

- **Goal:** Research how current CLI evaluation evidence should feed the website's planned numeric-turn chart, including chart organization, representative mock data, update workflow, and GitHub Pages publication.
- **Framing:** Implementation research only. Produce a best-supported implementation direction, conditional alternatives, rejected or lower-fit paths, evidence, and unresolved blockers.
- **Source scope:** Local plus internet, with external claims limited to primary or official sources.
- **Active researcher concurrency:** 5.
- **Total researcher budget:** 10.
- **Final aggregation target:** `research/implementation-research/website-chart-migration/website-chart-migration-implementation-research.md`.
- **Requested final shape:** A substantial but bounded implementation-research synthesis; no product edits and no implementation plan.

## Research Domains

1. **Current CLI evidence production:** command routing, evaluation runtime, replay capture, scoring, schemas, identities, and written artifacts.
2. **Evaluation methodology:** governed scenarios, answer sheets, turn attribution, expected-call meaning, campaign grouping, and Work Block boundaries.
3. **Current website presentation:** editable data contract, chart component, page composition, CSS layout, responsive behavior, and controlling website plans.
4. **Chart migration design:** numeric turn axis, called/missed series, comparison dimensions, filters, legends, scaling, defaults, and representative mock data.
5. **Static publication workflow:** Vite asset/base behavior, GitHub Pages deployment, data-update ergonomics, validation, and publication automation.

## Discovery Waves And Assignments

### Wave 1: Map Concrete Owners

| Assignment | Narrow scope                             | Primary source targets                                                         | Expected evidence                                                                                       | Output                                              |
| ---------- | ---------------------------------------- | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------- | --------------------------------------------------- |
| 01         | CLI evaluation command-to-artifact path  | `cli/internal/command`, `cli/internal/lifecycle`, `cli/internal/evaluation`    | Concrete entrypoint, request/result types, runtime sequence, file outputs                               | `assignments/01-cli-evaluation-production-path.md`  |
| 02         | Replay, scoring, and turn attribution    | `cli/internal/replay`, `cli/internal/evaluation`, instrumentation/event schema | Called/missed derivation, event correlation, turn-number semantics, limitations                         | `assignments/02-replay-scoring-turn-attribution.md` |
| 03         | Governed methodology and result identity | evaluation scenarios/built-ins and completion/runner plans                     | Harness/model/scenario/run grouping contract, valid-run meaning, Work Block 3 expectations              | `assignments/03-methodology-and-result-identity.md` |
| 04         | Current website data and chart rendering | `src/data/siteData.ts`, `src/components/EvaluationChart.tsx`                   | Existing graph schema, mock values, Recharts configuration, labels and accessibility                    | `assignments/04-current-website-chart.md`           |
| 05         | Website layout and controlling contracts | `src/App.tsx`, `src/styles.css`, `plans/website`, authority map                | Page placement, responsive layout, intended visual/data ownership, explicit chart migration constraints | `assignments/05-website-layout-contracts.md`        |

### Wave 2: Deep Dives And Cross-Checks

| Assignment | Narrow scope                             | Primary source targets                                       | Expected evidence                                                                                     | Output                                                  |
| ---------- | ---------------------------------------- | ------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| 07         | Recharts numeric-axis implementation fit | official Recharts documentation/source and installed version | Numeric axis, multiple series, tooltip/legend, responsive/container behavior, constraints             | `assignments/07-recharts-numeric-axis-options.md`       |
| 08         | GitHub Pages plus Vite deployment        | official GitHub and Vite documentation, repo config/scripts  | Base-path requirements, Actions/static deployment choices, publication constraints                    | `assignments/08-github-pages-vite-publication.md`       |
| 09         | Comparison-layout decision               | Findings from CLI/methodology/website sources                | Overlay versus facets, harness/model defaults, toggles, scaling, legends, mock dataset recommendation | `assignments/09-chart-comparison-design.md`             |
| 10         | Result-data update workflow              | repo scripts/contracts plus official GitHub/Vite guidance    | Low-friction data refresh and publish options, validation path, tradeoffs and rejection reasons       | `assignments/10-result-data-update-workflow.md`         |

## Candidate Ranking And Fan-Out

- **Deep-dive candidates:** production evaluation writer/scorer, governed methodology, current chart/data contract, current layout contracts, Recharts official API, Vite official deployment guidance, GitHub Pages official Actions guidance, cross-dimensional chart organization, and result refresh workflow.
- **Skim-only candidates:** existing historical hosting research, tests adjacent to traced production paths, generic chart examples, package metadata, and prior screenshots. Use them only to validate or illustrate source-established behavior.
- **Rejected candidates:** unrelated CLI lifecycle behavior, unrelated website sections, alternative hosting providers, chart-library migration, server-side dashboards, dynamic APIs, and implementation planning.

## Cross-Check Rules

- Validate schema and semantics against production writers and concrete implementations, not tests alone.
- Keep current implementation, controlling plans, and recommendations visibly distinct.
- Cross-check chart recommendations against both emitted data granularity and the governed testing methodology.
- Treat missing Work Block 3 result evidence as a limitation; do not infer real comparative outcomes from mock data.
- Classify unverified external behavior as unsupported or blocked.

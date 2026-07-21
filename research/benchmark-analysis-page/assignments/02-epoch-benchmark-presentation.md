# Epoch AI Benchmark Presentation Patterns

## Assignment

- **Goal:** Extract adaptable presentation patterns from Epoch AI's first-party benchmark surfaces for a restrained, information-dense Skill Issue benchmark Analysis page, with particular attention to hierarchy, navigation, chart/table-to-prose relationships, comparative language, observation versus interpretation, uncertainty, and practical meaning.
- **Scope:** Internet-only review of Epoch AI's benchmark dashboard, benchmark catalogue, benchmark detail pages, benchmarking FAQ, data-access page, and ECI documentation. Textual inspection was cross-checked against the live rendered desktop pages on 21 July 2026.
- **Exclusions:** Epoch-specific benchmark findings, model rankings, capability concepts, visual branding, and numerical results are not proposed for transfer. This assignment does not infer or invent any Skill Issue benchmark results, and it does not evaluate Epoch AI's statistical methods on their merits.

## Sources

- **Data on AI Capabilities and Benchmarking | Epoch AI** — <https://epoch.ai/benchmarks>. Relevant sections: page introduction, benchmark-mode switcher, interactive chart settings, and “Benchmarking updates.”
- **Benchmarks | Epoch AI** — <https://epoch.ai/benchmarks/search>. Relevant sections: benchmark catalogue introduction, evaluator and domain filters, sort controls, and benchmark result cards.
- **Epoch Capabilities Index | Epoch AI** — <https://epoch.ai/benchmarks/eci>. Relevant sections: introductory definition, chart/leaderboard switcher, domain-specific explorer, “More about this dataset,” documentation, and FAQ.
- **SWE-bench Verified | Epoch AI** — <https://epoch.ai/benchmarks/swe-bench-verified>. Relevant sections: benchmark summary and provenance, chart/leaderboard switcher, benchmark explanation, methodology, changelog, and notes.
- **About | Benchmarking | Epoch AI** — <https://epoch.ai/benchmarks/about>. Relevant sections: “Why we made this,” “Methodology at a glance,” “How accurate is the data?,” “Why are some of your scores different from those reported elsewhere?,” and “What do the error bars represent?”
- **Data | Benchmarking | Epoch AI** — <https://epoch.ai/benchmarks/use-this-data>. Relevant sections: “Use this data,” downloads, licensing, and citation.
- **ECI Documentation – Overview | Epoch AI** — <https://epoch.ai/data/eci-documentation>. Relevant sections: documentation navigation and overview.
- **ECI Documentation – Methodology | Epoch AI** — <https://epoch.ai/data/eci-documentation/methodology>. Relevant sections: methodology explanation, model definition, fitting approach, scaling, and confidence intervals.

## Findings

### Finding 1: Establish a stable reading path before adding analytical density

Epoch's benchmark surfaces repeatedly use the same macro-order: identify the subject, state what the evidence represents, expose the primary visual or table, then provide progressively deeper explanation. The dashboard starts with a dated page title and a short statement of purpose before the explorer. An individual benchmark page starts with a plain-language task summary, small descriptive metrics, domain labels, and provenance before the chart. Documentation and methodology follow below rather than competing with the opening evidence view.

**Evidence:** The benchmark dashboard introduces its database and the available comparison pathways before presenting its explorer. The SWE-bench detail page places a task description, model-count and highest-score metadata, domain labels, and the statement that the data comes from internal runs above its chart and leaderboard. The ECI detail page similarly defines its composite measure before displaying any visualization. ([dashboard](https://epoch.ai/benchmarks), [SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified), [ECI detail](https://epoch.ai/benchmarks/eci))

**Implication:** A Skill Issue Analysis page can preserve density without opening with an undifferentiated wall of results. A defensible order is: page purpose and evaluation scope; compact result summary; primary comparative visual/table; evidence-led observations; practical interpretation; limitations and methodology. Each layer should answer the question raised by the preceding layer.

### Finding 2: Separate section navigation from evidence-view controls

Epoch distinguishes “where am I?” navigation from “how am I viewing this evidence?” controls. The benchmark product has persistent section links for Dashboard, Benchmarks, Models, Data, and About. Within an evidence page, local buttons switch among alternative views such as release-date chart, compute chart, and leaderboard. This prevents filters and view modes from masquerading as site navigation.

**Evidence:** The dashboard and detail pages use the same secondary navigation row, while the ECI and SWE-bench pages place chart/table mode switches immediately above the corresponding evidence region. The catalogue gives discovery controls their own left-hand filter area and places search/sort controls over the result list. ([dashboard](https://epoch.ai/benchmarks), [catalogue](https://epoch.ai/benchmarks/search), [ECI detail](https://epoch.ai/benchmarks/eci), [SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified))

**Implication:** Keep page-level Analysis navigation distinct from chart filters, grouping controls, and chart/table switches. Readers should be able to predict whether a control changes the page, changes the dataset slice, or changes only its representation.

### Finding 3: Make the catalogue a decision aid, not a name directory

Epoch's benchmark catalogue cards provide enough information to judge relevance before opening a detail page. Each visible card combines a title, a one-sentence description of what is evaluated, evaluator provenance, number of models evaluated, a compact score indicator where applicable, and domain tags. The catalogue can be filtered by evaluator and domain and sorted by recommendations, recency, saturation, evaluation count, or alphabetically.

**Evidence:** The catalogue describes its coverage at the top, then presents evaluator and domain facets alongside cards whose descriptive and provenance fields are visible without navigation. The live rendered page visually treats each card as a bordered horizontal information unit, with the evaluator tag at the upper right and supporting measures aligned in a compact lower row. ([benchmark catalogue](https://epoch.ai/benchmarks/search))

**Implication:** If the Skill Issue page exposes multiple benchmark analyses, each selector or card should answer four questions in place: what was tested, who/what produced the evidence, how much evidence exists, and which category it belongs to. Avoid requiring a detail-page visit merely to learn whether a benchmark is relevant.

### Finding 4: Treat chart and table as two views of the same evidence object

Epoch does not make the chart and leaderboard separate stories. They occupy the same page position, retain the same benchmark context, and switch through adjacent mode buttons. The chart supports pattern recognition over time; the leaderboard supports exact comparison. On the inspected SWE-bench page, switching to the leaderboard replaced the scatter plot with ranked horizontal bars and columns for model, accuracy, organization, and logs, while retaining uncertainty values beside the estimates.

**Evidence:** The SWE-bench page offers “Score vs Release Date,” “Score vs Training Compute,” “Score vs ECI,” and “Leaderboard” as peer modes. Live interaction confirmed that the leaderboard uses the same page header and benchmark context as the chart, and reports estimates with uncertainty plus trace links. The ECI page similarly offers chart and leaderboard modes under one introductory definition. ([SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified), [ECI detail](https://epoch.ai/benchmarks/eci))

**Implication:** Use one shared result model for the Skill Issue chart and table. Preserve selected benchmark/filter context across representations. Let the chart answer “what pattern is visible?” and the table answer “what are the exact values and evidence links?” rather than duplicating the same prose twice.

### Finding 5: Place prose after the evidence, but keep essential interpretation cues above it

Epoch gives the primary chart visual priority, yet does not make it self-interpreting. Short cues above the chart define the task, metric, result provenance, and any significant comparability break. Longer prose below explains dataset construction, evaluation workflow, and methodology. This division lets a knowledgeable reader scan quickly while giving a cautious reader an immediate path to context.

**Evidence:** The SWE-bench page places the real-world task description, provenance, and a note about a major scaffold upgrade above the evidence views. Below the evidence region, the “SWE-bench Verified” and “Methodology” tabs explain the sample set, evaluation workflow, residual ambiguity, prompts, tools, exclusions, limits, and environment. The ECI page places an introductory definition above its explorer and reserves documentation and FAQ for deeper explanation below. ([SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified), [ECI detail](https://epoch.ai/benchmarks/eci))

**Implication:** Put only the interpretive minimum above a Skill Issue visual: what was measured, under which evaluation scope, and any condition that materially changes comparison. Put extended explanation after the visual in clearly named sections, with direct links from the compact cues when the caveat is consequential.

### Finding 6: Keep observations mechanically grounded and visibly separate from interpretation

Epoch's strongest comparative language identifies the comparison operation before assigning meaning. For ECI, the page explicitly says that absolute values are not meaningful in isolation while relative comparisons can be meaningful, and it separately explains what changes in slopes can and cannot indicate. Elsewhere, the benchmarking FAQ distinguishes confidence in the recorded results under specified settings from caution about generalizing those results to other contexts.

**Evidence:** The ECI FAQ states the scope of meaningful comparison, explains the arbitrary scale, and warns that some desired quantities are not linearly related to that scale. The general benchmarking FAQ says its data reflects performance under the stated evaluation conditions and may fail to generalize because prompting, contamination, and setting differences matter. ([ECI detail, FAQ](https://epoch.ai/benchmarks/eci), [benchmarking FAQ](https://epoch.ai/benchmarks/about))

**Implication:** Use a repeatable two-step prose pattern on the Skill Issue page: **Observation** states what the displayed evidence directly shows; **Interpretation** states the bounded practical meaning and names the comparison basis. Comparative claims should identify the metric, evaluation condition, and reference group instead of relying on adjectives such as “better,” “strong,” or “significant” by themselves.

### Finding 7: Put provenance and evaluator ownership close to the result

Epoch treats provenance as part of the visible result, not as footer-only metadata. Catalogue cards distinguish internal evaluation, benchmark-creator evaluation, and model-developer evaluation. Individual benchmark pages say where the displayed data comes from before the chart. The About page then explains the different assurance levels for internal and external results and points to source columns for individual external data points.

**Evidence:** The catalogue labels evaluator ownership on every card. The inspected SWE-bench detail page states that its displayed data comes from Epoch's internal runs. The benchmarking FAQ explains that external results have less methodological visibility and that provenance is carried in source fields. ([catalogue](https://epoch.ai/benchmarks/search), [SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified), [benchmarking FAQ](https://epoch.ai/benchmarks/about))

**Implication:** Every Skill Issue result block should visibly identify the evidence owner and run context. If results combine sources or harnesses, show that classification where the comparison is read, then offer complete provenance in the table or detail panel.

### Finding 8: Express uncertainty in the visualization, the exact-value view, and the prose

Epoch uses three reinforcing uncertainty surfaces. Charts include error bars; leaderboard values include plus/minus uncertainty; explanatory prose states what the intervals represent and which limitations remain outside those intervals. The methodology documentation further explains the confidence-interval construction rather than leaving the visual convention unexplained.

**Evidence:** Live inspection of the SWE-bench chart showed error bars around points, and the leaderboard showed uncertainty next to each percentage. The benchmarking FAQ defines the error-bar convention for internally evaluated benchmarks and separately warns about contamination, prompt sensitivity, and limited generalizability. The ECI methodology states that it reports bootstrap confidence intervals and describes the resampling procedure. ([SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified), [benchmarking FAQ](https://epoch.ai/benchmarks/about), [ECI methodology](https://epoch.ai/data/eci-documentation/methodology))

**Implication:** When Skill Issue has statistically justified uncertainty, render it consistently in both chart and table and explain it nearby. Keep methodological uncertainty distinct from coverage limitations, harness effects, and interpretation limits; an interval should not visually imply that every uncertainty has been quantified.

### Finding 9: Turn “practical meaning” into a task-and-workflow explanation

Epoch's clearest benchmark explanations begin with the concrete activity required of the evaluated system, then describe the grading mechanism. On the SWE-bench page, the reader learns that the system receives a repository and issue, modifies the repository, and is graded through tests before encountering detailed prompt, scaffold, sample, limit, and environment specifications. This makes the metric legible as an activity before it becomes a methodology.

**Evidence:** The SWE-bench summary describes the benchmark as real-world repository issue resolution. The following workflow paragraph explains the input, actions, and grading effect in sequence; the Methodology section then decomposes the evaluation into five implementation areas. ([SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified))

**Implication:** Explain each Skill Issue benchmark in operational language: what the evaluated agent receives, what it must do, what outcome is scored, and what a higher or lower result practically indicates within that setup. Keep implementation details available, while allowing the practical explanation to stand on its own.

### Finding 10: Use restrained visuals to make analytical structure—not decoration—the dominant signal

The inspected desktop pages use a wide white canvas, black type, light borders and gridlines, compact pill controls, and a limited accent palette. Large headings and generous vertical spacing establish section boundaries; denser controls and data live inside clearly bounded regions. On the catalogue, the filter column and result cards form the primary composition. On detail pages, the introductory text spans only part of the width, while the chart uses the wider canvas and a settings panel sits to its right.

**Evidence:** Live visual inspection of the dashboard, catalogue, ECI detail, and SWE-bench detail confirmed consistent alignment, restrained chrome, and clear density shifts between introductory prose, interactive evidence, and supporting documentation. Color is used mainly for series, active controls, progress indicators, and evaluator tags rather than decorative backgrounds. The chart legend and table columns remain text-labeled, limiting dependence on color alone. ([dashboard](https://epoch.ai/benchmarks), [catalogue](https://epoch.ai/benchmarks/search), [ECI detail](https://epoch.ai/benchmarks/eci), [SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified))

**Implication:** Favor typographic hierarchy, spacing, alignment, light dividers, and compact controls over card-heavy ornament. Reserve color for meaningful states and series. Give prose a readable measure and give comparative evidence the width it needs.

### Finding 11: Surface comparability breaks as part of the result's current state

Epoch places material benchmark-version changes close to the analysis and maintains a detailed changelog. The SWE-bench page's summary calls out a significant scaffold upgrade, while the changelog records which versions changed environments, token limits, tools, or implementation behavior and whether prior scores are expected to be affected.

**Evidence:** The SWE-bench page includes an above-chart note about the February 2026 scaffold change. Its changelog explains the major version transition, states that key models were re-evaluated, and says the default graph view limits results to the new version. Smaller changes are also recorded with an explicit assessment of expected score impact. ([SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified))

**Implication:** When Skill Issue evaluation conditions change materially, place a concise comparability notice near the result and link to a dated change record. Do not hide a run-breaking methodology change in release notes or silently combine incomparable series.

### Finding 12: Provide a traceability ladder from summary to raw evidence

Epoch offers progressively stronger evidence access: labelled chart points, an exact-value leaderboard, per-run log links for internal evaluations, methodology descriptions, downloadable data, citations, and implementation/code links. Readers can stop at the level appropriate to their purpose without losing the ability to verify a claim.

**Evidence:** The SWE-bench leaderboard includes a Logs column; the benchmarking FAQ describes the log viewer and the details it exposes; the Data page offers downloadable data, table access, licensing, and a preferred citation; and the ECI documentation links to its public code and technical methodology. ([SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified), [benchmarking FAQ](https://epoch.ai/benchmarks/about), [data access](https://epoch.ai/benchmarks/use-this-data), [ECI overview](https://epoch.ai/data/eci-documentation))

**Implication:** Design the Skill Issue Analysis page as the top of a verification ladder. A summary claim should lead to an exact row, then to run-level artifacts or source records, then to the methodology that produced them. This allows restrained prose because evidence remains inspectable rather than being restated exhaustively.

## Notes

### Candidate classification

- **Deep-dive:** `Data on AI Capabilities and Benchmarking` for the dashboard hierarchy and evidence explorer; `Benchmarks` for discovery and card design; `Epoch Capabilities Index` for chart/table/prose layering and comparative-language guardrails; `SWE-bench Verified` for an individual benchmark template, practical meaning, provenance, uncertainty, and changelog; `About | Benchmarking` for assurance, limitations, and error-bar interpretation; `ECI Documentation – Methodology` for uncertainty and technical-depth separation.
- **Skim-only:** `Data | Benchmarking` for the traceability/download/citation layer; `ECI Documentation – Overview` because it largely repeats the ECI detail-page introduction while validating the separate documentation information architecture; `GPQA Diamond` because its detail-page structure substantially repeats the selected SWE-bench template without adding a stronger presentation pattern.
- **Reject:** Epoch's general publications feed and topical landing pages because they are editorial discovery surfaces rather than benchmark-analysis interfaces; FrontierMath's standalone promotional/product pages because their bespoke narrative would weaken the comparison to a restrained multi-benchmark Analysis page; third-party benchmark sites and papers because this assignment was limited to Epoch's first-party presentation surfaces.

### Validation performed

- Opened and inspected all cited Epoch pages directly on 21 July 2026.
- Cross-checked extracted page text against live rendered desktop views of the dashboard, catalogue, ECI detail, and SWE-bench detail pages.
- Interactively switched the SWE-bench detail page from chart mode to leaderboard mode and confirmed that benchmark context, uncertainty values, provenance columns, and run-log links remained available.
- Cross-checked the visible uncertainty treatment against Epoch's benchmarking FAQ and ECI methodology documentation.

### Caveats

- Epoch's benchmark pages are live, dynamically updated products. Counts, dates, benchmark membership, labels, and visual details observed on 21 July 2026 may change.
- Visual findings describe the inspected desktop layout. Responsive behavior was outside this assignment.
- The recommendation to adapt these patterns concerns information architecture and evidence communication only; it does not endorse transferring Epoch's benchmark ontology, statistical model, terminology, or results.

# METR Benchmark Reporting Patterns

## Assignment

**Goal.** Extract adaptable reporting patterns from METR's first-party benchmark and agent-evaluation pages for a concise Skill Issue benchmark Analysis page, with particular attention to findings, uncertainty, practical meaning, chart-to-prose relationships, headings, information order, comparative claims, and observation-versus-interpretation boundaries.

**Scope.** Internet-only inspection of METR's public website, centered on two chart-rich research reports: the current evaluation-revision note and the foundational time-horizon explainer. Discovery and currency checks covered METR's home page, research index, and living time-horizon page. The implications below concern presentation structure only.

**Exclusions.** This assignment does not transfer METR's benchmark design, metrics, subject-matter conclusions, forecasts, or terminology into Skill Issue. It does not infer or fabricate Skill Issue results, recommend chart values, assess METR's underlying scientific claims, or specify visual styling such as color, typography, or responsive layout.

## Sources

### Source selection and candidate disposition

Discovery began with METR's [home page](https://metr.org/) and [Research index](https://metr.org/research/) on July 21, 2026. Candidates were ranked for first-party provenance, prominence, chart density, explicit findings, uncertainty treatment, and closeness to benchmark or agent-evaluation reporting.

| Disposition   | Candidate                                                                                                                                               | Selection rationale                                                                                                                                                                                                                                  |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Deep-dive** | [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/), January 29, 2026                                                                 | A concise current-method revision note that pairs declarative finding headlines with charts, tables, uncertainty, matched comparisons, and appendices. It is the strongest example of reporting what changed without hiding measurement limitations. |
| **Deep-dive** | [Measuring AI Ability to Complete Long Tasks](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/), March 19, 2025            | A prominent explanatory report that moves from summary to charts, method, practical meaning, robustness, limitations, and conclusion. It complements the update note by showing a fuller narrative arc.                                              |
| **Skim-only** | [Task-Completion Time Horizons of Frontier AI Models](https://metr.org/time-horizons/), last updated May 8, 2026                                        | Useful to confirm the living/current status of the reporting surface, provenance links, and compact methodological explanation; too dashboard-like to be the main narrative-pattern source.                                                          |
| **Skim-only** | [Frontier Risk Report (February to March 2026)](https://metr.org/blog/2026-05-19-frontier-risk-report/), May 19, 2026                                   | Current and chart-rich, but its long multi-evidence risk-assessment format is farther from a concise benchmark Analysis page.                                                                                                                        |
| **Skim-only** | [Measuring the Self-Reported Impact of Early-2026 AI on Technical Worker Productivity](https://metr.org/blog/2026-05-11-ai-usage-survey/), May 11, 2026 | Shows result-first headings and cautious interpretation, but reports survey evidence rather than a benchmark or agent evaluation.                                                                                                                    |
| **Reject**    | [About METR](https://metr.org/about/) and the METR home page                                                                                            | Useful only for organizational context; neither provides a sustained chart-to-finding report structure.                                                                                                                                              |
| **Reject**    | [Portable Evaluation Tasks via the METR Task Standard](https://metr.org/blog/2024-02-29-metr-task-standard/), February 29, 2024                         | Primarily a technical standard announcement rather than a chart-led findings report.                                                                                                                                                                 |
| **Reject**    | [Task Substitution and Uplift](https://metr.org/blog/2026-05-08-task-substitution-and-uplift/), May 8, 2026                                             | A conceptual note whose argument structure is less applicable to benchmark-result presentation.                                                                                                                                                      |

### Deep-dive source details

1. **METR, [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/), January 29, 2026.** Inspected the opening standfirst; the two-change numbered list; the bold finding-led passages beginning “The estimated time horizon for each model has changed somewhat,” “Our new task suite contains more tasks,” “The trend in time horizon is somewhat sensitive to task composition,” and “We are working on raising the ceiling”; the task-distribution and old-versus-new trend charts; and **Appendices → Data, Comparison of Old and New Estimates, Changes to Model Horizon Estimates, Moving from Vivaria to Inspect**. The inspected chart assets included [task distribution comparison](https://metr.org/assets/images/time-horizon-1-1/task-distribution-comparison.png), [hybrid old/new comparison](https://metr.org/assets/images/time-horizon-1-1/time-horizon-1-vs-1-1-hybrid.png), [since-2023 old/new comparison](https://metr.org/assets/images/time-horizon-1-1/time-horizon-1-vs-1-1-since-2023.png), and [model-level infrastructure comparison](https://metr.org/assets/images/time-horizon-1-1/models.png).

2. **Thomas Kwa, Ben West, Joel Becker, et al., [Measuring AI Ability to Complete Long Tasks](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/), March 19, 2025.** Inspected the standfirst; **Summary**; the initial trend chart and uncertainty caption; the benchmark-context chart; the metric-construction prose and fitted-curves chart; the practical-meaning interpretation; the robustness, sensitivity, and model-error discussion; **Conclusion**; and **Want to contribute?** The inspected chart assets included the [lead trend chart](https://metr.org/assets/images/measuring-ai-ability-to-complete-long-tasks/length-of-tasks-log.png), [fitted success curves](https://metr.org/assets/images/measuring-ai-ability-to-complete-long-tasks/models-are-succeeding-at-increasingly-long-tasks.png), and [uncertainty sensitivity analysis](https://metr.org/assets/images/measuring-ai-ability-to-complete-long-tasks/uncertainty-in-extrapolated-date.png).

### Validation performed

The report titles and dates were cross-checked against each page's displayed metadata and embedded BibTeX citation. Section labels and sentence-style bold headings were checked in the live HTML. Chart order, alt text, captions, and raw image targets were inspected from the public pages; the living [time-horizon page](https://metr.org/time-horizons/) was checked to confirm which reporting surface METR labeled current as of May 8, 2026. Comparative statements were checked against the adjacent tables, captions, and caveats in the same report. No independent recomputation of METR's results was needed because the assignment evaluates communication patterns rather than scientific correctness.

## Findings

### Finding 1: Open with the result and the reason the page exists

Both reports establish purpose before asking the reader to interpret a chart. The revision note opens with a one-sentence description of the release and immediately summarizes how the estimates changed. The foundational report places its central proposition and main directional result in the standfirst, repeats that compactly under **Summary**, and then shows the lead chart. This gives a skimming reader the reporting event, scope, and headline before detail.

**Evidence.** The opening of [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) states that the release uses more tasks and new infrastructure, then says the updated estimates mostly remain within earlier intervals while the overall trend changes somewhat. The [foundational report's Summary](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/) states the proposed measurement, the observed direction, and the conditional extrapolation before its first chart.

**Implication.** A Skill Issue Analysis page can begin with one compact, evidence-bound sentence answering: what was analyzed, what the main observed pattern is, and what comparison or evaluation version it belongs to. The lead should orient the reader rather than merely announce “Results.”

### Finding 2: Use declarative findings as skimmable headings

The shorter revision report uses bold sentence-style headings that already contain the finding. This lets a reader recover the narrative by scanning headings alone, while the following chart and paragraph justify or qualify each claim. The headings remain measured: they use terms such as “somewhat,” “slightly,” or a bounded scope instead of turning every difference into a major result.

**Evidence.** [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) organizes its body around statements including “Our new task suite contains more tasks, yielding tighter estimates especially at the upper end” and “The trend in time horizon is somewhat sensitive to task composition.” Its infrastructure appendix follows the same pattern with result-led labels before the supporting comparisons.

**Implication.** Prefer headings such as “Scores separate most clearly on …” or “The comparison remains uncertain when …” only when Skill Issue evidence supports those statements. Finding-led headings should include the direction and scope; neutral headings remain appropriate where the evidence supports description but not a conclusion.

### Finding 3: Repeat a stable visual-to-prose sequence

METR's strongest reusable sequence is: a finding statement sets the question, the visualization shows the evidence, the caption explains encodings and statistical meaning, and the next prose paragraph states the practical reading plus its limit. The chart is neither left to “speak for itself” nor restated point by point.

**Evidence.** In [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/), the task-suite finding precedes the distribution chart; the adjacent prose quantifies the change in interval width and immediately says the intervals remain wide. The old/new trend charts are introduced with the comparison design, captioned with the series encodings, and followed by explanations of which observations drive the shift. In [Measuring AI Ability to Complete Long Tasks](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/), the lead chart is followed by the motivating problem, the fitted-curves chart explains how the metric is produced, and subsequent prose interprets what the result can and cannot mean in practice.

**Implication.** Each primary Skill Issue visualization can be paired with four concise elements: a claim-bearing heading, the chart, a literal caption, and a short “why this matters” paragraph. The caption should explain what is plotted; the paragraph should explain the bounded interpretation.

### Finding 4: Give charts, captions, prose, and tables different jobs

The reports avoid making one presentation layer carry every detail. Charts reveal patterns and distributions; captions define series, thresholds, intervals, and transformations; prose explains consequences and caveats; tables preserve exact estimates and comparison values. This division keeps the narrative concise without discarding auditability.

**Evidence.** The trend-chart captions in [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) specify which colors represent the old and new setups and identify the hybrid points, while the adjacent prose explains why the comparison is incomplete. The appendix tables then enumerate exact old/new estimates. The lead-chart caption in [Measuring AI Ability to Complete Long Tasks](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/) identifies the confidence region and its bootstrap basis, while later prose addresses methodological and external-validity uncertainty.

**Implication.** Use the Skill Issue chart for shape, the caption for reading instructions, the finding prose for interpretation, and a compact table or disclosure for exact values. This prevents repetitive prose and makes the underlying observations independently checkable.

### Finding 5: Attach uncertainty to the claim it limits

Uncertainty is reported locally rather than deferred to a generic limitations footer. METR places interval information in chart captions, calls out when intervals remain wide in the next sentence, identifies missing or estimated baselines alongside the affected result, and distinguishes measured statistical sensitivity from uncertainty the plot does not capture.

**Evidence.** After describing tighter estimates, [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) immediately notes that the intervals are still wide and that most long-task times are estimates rather than measured baselines. When comparing trends, it explains that the model sets are mostly overlapping rather than identical and that the quoted intervals have a particular bootstrap interpretation despite task overlap. In the [foundational report](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/), the sensitivity chart's caption defines its percentile ranges, while the next sentence states that future trend changes and external validity are outside the plot and dominate overall uncertainty.

**Implication.** Place each Skill Issue caveat beside the affected number or chart. Distinguish at least three kinds when they exist in the real data: sampling or run variation, benchmark coverage or comparability, and interpretation or external validity. A single undifferentiated “limitations” paragraph weakens the reader's ability to judge a specific claim.

### Finding 6: Define the comparison before stating who changed or led

METR makes comparative claims legible by explaining whether datasets, model sets, time windows, task suites, or infrastructure match. When a clean full-period comparison is unavailable, the report names the constructed comparison, provides a narrower overlapping comparison, and shows exact values separately. This prevents a percentage difference from appearing more comparable than it is.

**Evidence.** [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) states that pre-2023 estimates were not rerun, so it first uses a hybrid series and then a post-2023 comparison where both methods have data. It warns that the model sets are not identical, identifies which older and newer estimates drive the change, and uses paired tests for the infrastructure comparison. The appendices expose the model-level and aggregate values behind those summaries.

**Implication.** Before any Skill Issue comparison, state the common basis: same tasks or different tasks, same harness or different harness, matched or unmatched attempts, evaluation version, and relevant denominator when available. Comparative prose should name the comparator and avoid collapsing unmatched evidence into a simple ranking.

### Finding 7: Mark observation and interpretation with explicit language

The reports repeatedly separate what the analysis directly found from what the authors think explains it. Direct results use formulations such as “we found” or “we see”; causal or evaluative readings use “we believe,” “seems reasonable evidence,” or “seems likely.” The linguistic distinction is especially visible when the observed result could support more than one explanation.

**Evidence.** In the infrastructure comparison, [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) reports the statistically different scores first, then labels scaffold sensitivity as a plausible inference and later says a task-level imbalance is likely driven by particular model-level results. Its task-composition section similarly frames the source-distribution explanation as a belief rather than a measured fact. The [foundational report](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/) describes the fitted trend, then separately introduces the conditional forecast and the possibility of substantial model error.

**Implication.** Skill Issue finding cards can use visible labels or disciplined phrasing such as **Observed**, **Interpretation**, and **Caveat**. An explanation should never inherit the certainty of the observation merely because both appear in the same paragraph.

### Finding 8: Translate the metric only after showing how it is constructed

The foundational report earns its practical interpretation by first explaining the data relationship, then showing the fitted curves, then connecting the resulting measure to a real-world tension. The practical meaning is close enough to the chart to guide the reader, but it follows rather than replaces the measurement explanation.

**Evidence.** [Measuring AI Ability to Complete Long Tasks](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/) first describes the observed relation between task duration and success, explains how a chosen success threshold intersects a fitted curve, and illustrates the construction in a chart. Only then does it argue that the result helps reconcile strong benchmark performance with limited dependable automation. Later, the conclusion returns to the value of an absolute, interpretable measure after the report has presented robustness and limitations.

**Implication.** For Skill Issue, introduce the relevant measurement in plain language, show how the plotted value is derived from available evaluation artifacts, and then state the practical reading. “Why it matters” should describe what a reader may reasonably infer about the evaluated behavior, without claiming broader usefulness or quality that the evidence does not establish.

### Finding 9: Layer the page from fast answer to audit trail

The two reports support both a quick read and a verification path. The upper layer contains the standfirst, summary, and main chart. The middle layer explains the construction, comparisons, practical meaning, and uncertainty. The lower layer holds sensitivity discussion, appendices, exact tables, data, and analysis links. This order keeps methodological detail available without forcing every reader through it before seeing the result.

**Evidence.** The [foundational report](https://metr.org/blog/2025-03-19-measuring-ai-ability-to-complete-long-tasks/) orders its material as standfirst → current interactive result surface and provenance → Summary and lead chart → motivation → method and explanatory charts → practical interpretation → robustness and model-error caveats → Conclusion → open-source resources. [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) orders its material as release summary → changed inputs → headline estimate and trend findings → measurement-ceiling caveat → appendices with data, exact comparisons, and infrastructure checks.

**Implication.** A concise Skill Issue Analysis page can use the same progressive disclosure: headline finding and primary visual first; supporting findings and caveats next; evaluation metadata, exact values, and artifact/provenance links last. Readers should be able to verify a claim without the audit trail dominating the first screen.

### Finding 10: Treat evaluation revisions as explanatory evidence

The current update note makes changes to the evaluation setup part of the findings narrative. It enumerates additions, removals, modifications, infrastructure changes, and incomplete reruns before presenting changed estimates. That establishes a defensible boundary between a changed measured result and a changed underlying system.

**Evidence.** [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/) opens with two evaluation changes, gives counts for task additions, removals, modifications, and longer tasks, and states why only a subset of models was re-estimated. It then attributes estimate movement among task-suite changes, run noise, and infrastructure, with a later appendix testing the infrastructure contribution.

**Implication.** When Skill Issue results span evaluation versions, surface the relevant configuration, dataset, rubric, harness, or coverage change before interpreting movement. Revision metadata is part of the evidence needed to understand a comparison, rather than incidental implementation history.

## Notes

- **Caveat:** The implications are design inferences from two METR reporting patterns, not claims that METR prescribes a universal benchmark-reporting standard.
- **Caveat:** The living time-horizon page was used to validate currency and provenance, while the two dated reports supplied the narrative patterns. Dynamic chart interactions were not treated as evidence for a prose recommendation.
- **Unsupported observation:** No claim is made that METR's visual styling, chart types, or page density would fit Skill Issue; adapting those elements would require inspection of Skill Issue's real data volume, component system, and responsive behavior.
- **Boundary:** Every suggested finding headline, practical interpretation, comparison, and caveat must be populated from actual Skill Issue evidence. This assignment supplies no Skill Issue values or conclusions.
- **Useful search terms:** `site:metr.org benchmark report confidence interval`, `site:metr.org/blog time horizon evaluation`, `site:metr.org/blog agent evaluation results`, `site:metr.org/blog sensitivity analysis chart`.

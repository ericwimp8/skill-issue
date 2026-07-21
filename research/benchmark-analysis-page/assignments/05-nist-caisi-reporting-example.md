# NIST CAISI Public Evaluation Reporting Example

## Assignment

**Goal:** Extract adaptable public-report patterns from NIST CAISI's DeepSeek V4 Pro evaluation for the Skill Issue benchmark Analysis page, with particular attention to findings order, headings, concise comparison language, visual/table-to-prose relationships, observation versus interpretation, limitations and uncertainty, and practical implications.

**Scope:** Internet-only inspection of the May 2026 NIST CAISI public evaluation page, the earlier full CAISI report it links for benchmark and methodology detail, and the primary materials needed to validate attribution and reporting boundaries.

**Exclusions:** This assignment does not transfer CAISI's domain conclusions, geopolitical framing, model rankings, or policy claims to Skill Issue. It does not infer or invent Skill Issue benchmark outcomes, recommendations, statistical significance, or causal explanations.

## Sources

### Deep-Dive Sources

- **NIST, “CAISI Evaluation of DeepSeek V4 Pro”** (released May 1, 2026; updated May 2, 2026), especially the lead, “Key Findings,” “Capability Results,” “Benchmarks,” “Capability Lag Measurement,” “Model Serving and Inference,” “Comparison of DeepSeek and CAISI Evaluations,” “DeepSeek V4 Costs Less Than Other Models Of Similar Capability,” and “Appendix”: https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro
- **Center for AI Standards and Innovation, NIST, “Evaluation of DeepSeek AI Models”** (69-page PDF; report identifier exposed by NIST as `caisi-evaluation-deepseek-ai-models-report`), especially Sections 1, 2, 3, 3.1–3.5, 4–8, 9, and the Appendix: https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report

### Skim-Only Candidates

- **DeepSeek-AI, “DeepSeek-V4 Technical Report”** (58-page PDF), Sections 5.3.1–5.3.2 and Table 6. Used only to confirm that NIST's “self-reported evaluations” comparison is attributed to a distinct developer-authored source: https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro/resolve/main/DeepSeek_V4.pdf
- **Chollet et al., “ARC-AGI-2: A New Challenge for Frontier AI Reasoning Systems,” arXiv:2505.11831.** Linked benchmark provenance; useful for methodological verification, but not for the target public-report style: https://arxiv.org/abs/2505.11831
- **Rein et al., “FrontierScience: Evaluating AI's ability to perform expert-level scientific tasks,” arXiv:2601.21165.** Linked benchmark provenance; useful for metric context, but not for the target report architecture: https://arxiv.org/abs/2601.21165
- **UK AI Security Institute, “ReAct Agent – Inspect.”** Linked evaluation-scaffolding documentation; useful for reproducibility context, but not for the report's findings language: https://inspect.aisi.org.uk/react-agent.html
- **NIST, “CAISI Evaluation of DeepSeek AI Models Finds Shortcomings and Risks.”** Press-release wrapper for the earlier report; skimmed to distinguish press framing from the evaluation-report pattern: https://www.nist.gov/news-events/news/2025/09/caisi-evaluation-deepseek-ai-models-finds-shortcomings-and-risks

### Rejected Candidates

- Third-party summaries, community GitHub issues, and implementation blogs were rejected because they are secondary, are not CAISI reporting artifacts, or do not provide authoritative evidence about NIST's reporting choices.
- Unrelated evaluation-report templates were rejected because the assignment is specifically to study the named first-party NIST example and its primary source chain.

### Validation Performed

- Traced every substantive reporting-pattern claim below to the NIST public page or the linked NIST PDF rather than relying on search-result summaries.
- Rendered and inspected the PDF's Executive Summary page, Section 3.1 overview table, Section 3.2 cost chart page, and Section 9 disclaimer page to validate layout and visual-to-prose relationships that are not fully represented by extracted text.
- Cross-checked the public page's characterization of developer-reported results against Table 6 and Sections 5.3.1–5.3.2 of the developer's technical report. This validation supports the attribution boundary only; it does not independently validate either party's benchmark outcomes.
- Checked linked benchmark and Inspect documentation for provenance and classified them skim-only because they support methodology rather than the requested public-report pattern.

## Findings

### Finding 1 — Open With One Synthesized Claim, Then Preview the Small Set of Findings

The public page starts with a one-paragraph result statement and immediately places the highest-level trend figure beside it. A short “Key Findings” section then previews three distinct takeaways before the page moves into capability details, comparison, and cost. The longer PDF uses the same principle at greater scale: its Executive Summary names six findings as bold claim sentences, gives one compact evidence paragraph under each, and closes with an overall implication. This creates a stable reading path for audiences who stop after the first screen or first page.

**Evidence:** The NIST page presents the lead and Figure 1 before “Key Findings,” then orders the body as “Capability Results,” source/method context, “Comparison of DeepSeek and CAISI Evaluations,” cost, and appendix. The linked PDF's Section 1 uses bold finding labels followed by one quantitative or comparative paragraph per finding; its table of contents then expands those same topics through methodology, overview, detailed evaluations, disclaimer, and appendix. [Public evaluation page](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro) [Full CAISI report, Sections 1–3](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report)

**Implication:** A Skill Issue Analysis page can use a claim-first lead followed by a deliberately small findings preview, with each preview item mapping one-to-one to a later evidence section. The finding count and content must come from actual Skill Issue evidence rather than adopting CAISI's count or conclusions.

### Finding 2 — Build an Evidence Ladder From Summary to Detail

CAISI repeats important findings at increasing levels of resolution rather than presenting every detail once. The public page moves from lead claim, to key finding, to summary table or figure, to local methodology and caveats. The full report moves from Executive Summary, to “Overview of Findings,” to domain overviews, to detailed “Dataset and Methodology” and “Results” sections, then to appendix material. Each layer points the reader to the next level instead of duplicating all detail.

**Evidence:** The PDF explicitly directs readers from Table 3.1 to Section 4 for evaluation detail and from overview figures to appendix sections for error-bar and token-budget methodology. The public page likewise points from the capability statement to Figure 1 and the Appendix, and from benchmark descriptions to Section 3 of the earlier report. [Public evaluation page, Capability Results and Appendix](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro) [Full CAISI report, Sections 3–4 and Appendix](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report)

**Implication:** The Analysis page should make the headline, section summary, chart/table, and methodological detail mutually reinforcing. Repetition is useful when each repetition adds precision; repeating the same sentence at every level would waste attention.

### Finding 3 — Make Comparative Sentences Scoped, Quantified, and Symmetrical

The strongest comparison sentences name the subject, comparator, metric scope, and magnitude in a single compact unit. CAISI commonly uses forms such as “on 5 out of 7 benchmarks,” a bounded cost range, or a percentage of tasks solved for the focal model, an earlier model, and a reference model. When results are close, the language softens to “similar,” “slightly,” “modest,” or “about,” and the longer report directs readers to uncertainty detail.

**Evidence:** The V4 page's cost finding gives both the count of favorable benchmarks and the full observed range, then later names two excluded benchmarks and why they were excluded. Its capability table defines the metric as task accuracy and labels higher values as better. The earlier report's domain summaries use parallel three-way comparisons and qualify close results as within the margin of error or only slightly different. [Public evaluation page, Key Findings and cost section](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro) [Full CAISI report, Sections 3.1–3.2](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report)

**Implication:** Skill Issue comparisons should use one grammatical template across findings: focal condition, comparator, evaluated slice, observed difference, and uncertainty qualifier where supported. This makes comparisons easy to scan and reduces accidental shifts between absolute scores, relative changes, counts, and inferred quality.

### Finding 4 — Let Visuals Carry the Matrix; Let Prose Carry the Reading Rule and Takeaway

CAISI does not narrate every cell or bar. Tables preserve the complete comparison matrix, while prose identifies the few relationships that matter. Captions do substantive work: they define the metric, state whether higher or lower is better, explain highlighting, decode axes or reference lines, identify uncertainty intervals, and surface exceptions. The paragraph before a figure explains why the comparison matters; the paragraph after it states the bounded takeaway.

**Evidence:** The rendered Section 3.1 page introduces the evaluation domains, presents a single table with all models and benchmarks, and uses the caption to explain percentage solved, best-result highlighting, and the margin-of-error qualification. The rendered Section 3.2 page first distinguishes four cost concepts, argues for end-to-end user expense, explains comparator selection, then shows the chart; its caption explains the reference value, bar direction, and 95% bootstrap confidence intervals. The V4 page's Figure 4 prose similarly defines values above 1.0 and explains that the comparison is limited to tasks both models solved. [Full CAISI report, Sections 3.1–3.2](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report) [Public evaluation page, Figure 4](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro)

**Implication:** Each Skill Issue visual should answer one comparison question. Adjacent prose should state why that view exists and its primary takeaway; the caption or table note should carry metric definitions, directionality, denominators, encodings, exclusions, and uncertainty so the visual remains interpretable when viewed in isolation.

### Finding 5 — Separate Measurement, Summary, and Interpretation

The full report repeatedly uses a three-part sequence: “Results” for the displayed measurement, “Key findings” for concise observations, and “Conclusion” for a broader synthesis. The V4 public page compresses those layers but still distinguishes its capability results from its interpretation of why developer-reported and independent suites differ. It attributes developer numbers to the developer, labels CAISI numbers as CAISI measurements, and adds the benchmark-selection fact that CAISI pre-committed to its suite.

**Evidence:** Sections 3.1.1–3.5 and detailed Sections 4–8 in the full report separate charted results from prose conclusions. On the V4 page, Figure 3 is explicitly divided into developer-selected benchmarks and CAISI-suite benchmarks; the prose describes the different observed performance while the caption states the pre-commitment condition. The developer report's Table 6 independently confirms that the first set is developer-reported and uses its own evaluation setup. [Public evaluation page, Figure 3 comparison](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro) [Full CAISI report, Sections 3–8](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report) [DeepSeek-V4 Technical Report, Sections 5.3.1–5.3.2](https://huggingface.co/deepseek-ai/DeepSeek-V4-Pro/resolve/main/DeepSeek_V4.pdf)

**Implication:** Skill Issue should label directly measured observations before offering interpretation. Results imported from another artifact should remain attributed to that artifact, and differences between suites should be explained through validated setup or selection facts rather than guessed causes.

### Finding 6 — Place Uncertainty Where the Reader Encounters the Claim

CAISI uses both local and global uncertainty disclosures. Local notes sit beside the affected table, figure, or claim: confidence intervals, imputed values, aggregation differences, scoring differences, excluded runs, and comparator-selection rules. The full PDF also includes a dedicated Disclaimer that limits generalization by domain, model version, time, benchmark relevance, training-data overlap, and evolving methods. This prevents a distant limitations section from carrying all qualification work.

**Evidence:** Figure 2 on the public page marks 95% confidence intervals, an imputed result, a scoring-method difference, and a known evaluator-setting difference directly under the table. The cost section names two exclusions and gives a separate reason for each. The PDF's Section 9 says the findings are preliminary, partial, tied to specific domains and model versions at a point in time, and subject to benchmark limitations. [Public evaluation page, Figure 2 and cost section](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro) [Full CAISI report, Section 9](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report)

**Implication:** The Analysis page should attach caveats to the exact chart, row, or sentence they qualify, then include a compact page-level scope statement for limitations that affect the whole analysis. Missing data and exclusions should remain visible and be explained, not silently removed from the comparison.

### Finding 7 — Explain Practical Meaning Only After Establishing the Comparison Basis

The report earns its practical implications by first explaining why a metric matters and why a comparator is appropriate. In the cost section, CAISI distinguishes several cost concepts, selects the end-user measure, and explains the performance-class filter used to choose the reference model before presenting the comparison. The Executive Summary closes with a conditional implication rather than treating every benchmark difference as a universal real-world outcome.

**Evidence:** Section 3.2 defines training cost, inference-serving cost, token price, and end-to-end user expense, then explains why the last is most relevant to users and why the selected model is a meaningful comparator. The V4 page repeats the comparator filter and identifies unsupported cost cases as exclusions. The full report's Section 9 separately warns that benchmark results may not generalize to other tasks. [Full CAISI report, Sections 1, 3.2, and 9](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report) [Public evaluation page, cost section](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro)

**Implication:** Skill Issue should connect a result to practitioner meaning only after defining the metric, unit of comparison, and comparator rationale. Practical language should stay conditional when the evidence is benchmark-bound.

### Finding 8 — Include a Small Setup-Validation Result Before Relying on the Main Comparison

CAISI describes one targeted reproduction check before relying on its broader independent evaluation. It reports using developer-recommended inference settings and reproducing a developer-reported result on one benchmark to rule out obvious serving or configuration errors. The check supports setup validity without claiming that it validates the entire evaluation.

**Evidence:** “Model Serving and Inference” lists the controlled settings and states that CAISI reproduced the developer's GPQA-Diamond result as a check against inference or configuration errors. The same section identifies the agent scaffold and task budgets, while the Appendix describes equal weighting and controlled scaffolding for aggregate comparison. [Public evaluation page, Model Serving and Inference and Appendix](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro)

**Implication:** If Skill Issue has an equivalent known-answer or reproduction check, the Analysis page can report it as setup validation before presenting comparative results. Its scope should be precise: one successful check supports the tested configuration path, not every benchmark or interpretation.

### Finding 9 — Use Outcome-Oriented Headings for Readers and Technical Labels for Evidence

The public page's major headings name reader questions or outcomes, such as key findings, capability results, evaluation comparison, and relative cost. The full report reserves technical section labels for the evidence hierarchy—methodology, dataset and methodology, results, disclaimer, and appendix. This creates a useful division: navigation communicates meaning, while subordinate labels communicate evidentiary function.

**Evidence:** The V4 page's outcome-oriented headings are paired with compact technical subsections such as “Benchmarks,” “Capability Lag Measurement,” and “Model Serving and Inference.” The PDF's numbered hierarchy mirrors each overview topic with detailed evaluation and appendix sections. [Public evaluation page](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro) [Full CAISI report, table of contents and Sections 2–9](https://www.nist.gov/document/caisi-evaluation-deepseek-ai-models-report)

**Implication:** Skill Issue can use plain-language section headings that state the comparison question or supported outcome, then use stable technical sublabels—Method, Result, Caveat, and Practical Implication—to help readers distinguish evidence from interpretation.

## Notes

- The May 2026 NIST page is styled as an “Updates” page but functions as a compact evaluation report: it contains findings, quantitative tables and figures, methods, exclusions, validation, cost analysis, and an appendix. Its page type should not be mistaken for a press release when adapting the reporting pattern.
- The public page says a fuller IRT methodology writeup is planned. That future material was unavailable in the inspected source set, so no claim about its contents is supported.
- The NIST artifacts do not establish which exact chart types, section count, or statistical methods Skill Issue should use. Those choices must follow Skill Issue's actual measures and data-generating process.
- Useful search terms for later cross-source comparison: `claim-first evaluation report`, `results key findings conclusion separation`, `benchmark comparison caption uncertainty`, `comparator selection rationale`, `local caveat plus global limitations`, `pre-committed evaluation suite`.

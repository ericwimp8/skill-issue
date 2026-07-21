# Google DeepMind Benchmark Reporting Patterns

## Assignment

**Goal**

Study recent first-party Google DeepMind pages that report model-evaluation results through charts or tables and link to primary technical reports. Extract adaptable patterns for the Skill Issue benchmark Analysis page: narrative structure, headings, information order, claims discipline, chart/table-to-prose linkage, comparative language, observation versus interpretation, limitations, uncertainty, and practical meaning.

**Scope**

Internet-only inspection of Google DeepMind publication pages and their linked primary technical reports. The deep dive uses two 2025 examples: one benchmark-suite report centered on a multi-model table and one model report centered on comparative performance charts.

**Exclusions**

This assignment does not transfer DeepMind-specific findings, model capabilities, benchmark concepts, or numerical results into Skill Issue. It does not infer or invent any Skill Issue result. It evaluates reporting form and evidence discipline only.

## Sources

### Selected deep dives

1. **Google DeepMind, “FACTS Benchmark Suite: Systematically evaluating the factuality of large language models,” December 9, 2025.**  
   URL: https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/  
   Inspected sections: opening problem statement; “The FACTS Benchmark Suite”; “Benchmark overview” and its benchmark subsections; “Results”; “Looking Ahead.” The page defines the evaluation object before displaying a 15-model results table, then links its methodology claim directly to the technical report. Selected because it is the closest structural analogue: a public benchmark explanation, an aggregate result plus component results, concise comparative prose, and an explicit statement of remaining headroom.

2. **Cheng et al., “The FACTS Leaderboard: A Comprehensive Benchmark for Large Language Model Factuality,” arXiv:2512.10791v1, December 11, 2025.**  
   First-party PDF linked by the page: https://storage.googleapis.com/deepmind-media/FACTS/FACTS_benchmark_suite_paper.pdf  
   Readable primary-report version: https://arxiv.org/html/2512.10791  
   Inspected sections and exhibits: §1 “Introduction”; §2 “The FACTS Leaderboard”; benchmark-specific methodology and results in §§3–6; Table 1 (aggregate and component results with 95% confidence intervals); Table 10 (judge validation); Table 11 (ineligible-response examples); §7 “Conclusion.”

3. **Google DeepMind, “SIMA 2: An Agent that Plays, Reasons, and Learns With You in Virtual 3D Worlds,” November 13, 2025.**  
   URL: https://deepmind.google/blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/  
   Inspected sections: opening and contents list; “The Power of Reasoning”; “A Leap in Generalization Performance”; the two performance-chart captions; “Looking to the Future: The Journey to General Embodied Intelligence”; “Responsible Development”; technical-report link. Selected because it complements the table-led FACTS example with chart-led narration, a material baseline-comparability caveat in a chart caption, and a dedicated limitations passage before broader practical interpretation.

4. **SIMA Team, “SIMA 2: A Generalist Embodied Agent for Virtual Worlds,” arXiv:2512.04797v1, December 5, 2025.**  
   First-party PDF linked by the page: https://storage.googleapis.com/deepmind-media/DeepMind.com/Blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/SIMA_Tech_Report_2025.pdf  
   Readable primary-report version: https://arxiv.org/html/2512.04797  
   Inspected sections and exhibits: §3.4 “Evaluations,” including “SIMA Evaluation Suite 2.0” and “Human Baselines”; §4.2 “Embodied Task Performance,” including Figures 6–9; §5 “Discussion.”

### Candidate classification

- **Deep-dive — FACTS Benchmark Suite.** Strongest fit for the intended Analysis page because results are presented only after the benchmark composition, score construction, sample structure, and governance are explained; the results prose covers leader, component movement, weakest area, and overall ceiling in one compact sequence.
- **Deep-dive — SIMA 2.** Strongest complement because it uses charts rather than a leaderboard table and makes comparison conditions and limitations visible near the evidence and conclusion.
- **Skim-only — “Gemini Robotics On-Device brings AI to local robotic devices,” June 24, 2025.** https://deepmind.google/blog/gemini-robotics-on-device-brings-ai-to-local-robotic-devices/ It includes three compact comparative charts and links to a technical report, but the page provides less on-page evaluation design and uncertainty than the selected examples.
- **Skim-only — “Gemini Robotics brings AI into the physical world,” March 12, 2025.** https://deepmind.google/blog/gemini-robotics-brings-ai-into-the-physical-world/ It links a technical report and translates evaluation claims into capability sections, but the public page is primarily a model introduction with demonstrations rather than a sustained benchmark-analysis narrative.
- **Reject — “Gemini achieves gold-medal level at the International Collegiate Programming Contest World Finals,” September 17, 2025.** https://deepmind.google/blog/gemini-achieves-gold-medal-level-at-the-international-collegiate-programming-contest-world-finals/ It has a useful comparison chart and careful conditional phrasing about a hypothetical ranking, but it reports a one-off competition event and links rules and solutions rather than a primary technical evaluation report. Its promotional event structure is a poorer fit for a reusable benchmark Analysis page.

### Validation performed

- Cross-checked the FACTS page’s score construction, model ordering, component results, and headroom claim against technical-report Table 1 and §7. The report adds 95% confidence intervals and explicitly identifies evaluation areas outside the suite’s coverage.
- Cross-checked the SIMA 2 page’s chart captions and baseline caveat against report §3.4 and §4.2. The report confirms that the newer evaluation suite is substantially more challenging, explains why the earlier baseline scores are lower under it, defines three evaluation mechanisms, and describes human-baseline timeout conditions.
- Cross-checked both pages’ limitation language against their reports’ concluding sections. No adaptable finding below depends on an unverified page-only numerical claim.

## Findings

### Finding 1 — Establish the evaluation contract before showing outcomes

Both selected pages delay their main comparative evidence until the reader understands what is being evaluated. FACTS moves from user stakes to suite composition, score construction, public/private data handling, benchmark overviews, distribution graphics, and examples before “Results.” SIMA 2 moves from prior-system context to capability framing and only then to a section explicitly about generalization performance.

**Evidence:** The FACTS page defines four components, sample counts, held-out data, score aggregation, and evaluation governance before its results table ([page](https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/)); the report states that component-level analysis yields the most complete insight even though an aggregate score facilitates comparison ([§2](https://arxiv.org/html/2512.10791#S2)). SIMA 2 similarly establishes the prior system and new capability frame before “A Leap in Generalization Performance” ([page](https://deepmind.google/blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/)).

**Implication:** Open the Skill Issue Analysis page with a compact evaluation contract: the question being answered, evaluated population or artifact set, comparison basis, metric direction, and any aggregation rule. A chart should arrive after the reader can correctly interpret its units and scope.

### Finding 2 — Use functional top-level headings and evidence-led result subheadings

The clearest structure combines neutral navigation headings with more specific outcome headings. FACTS uses functional headings such as “Benchmark overview,” “Results,” and “Looking Ahead,” making the page easy to scan without embedding conclusions prematurely. SIMA 2 uses capability and outcome-led headings after its subject has been established.

**Evidence:** The FACTS page’s top-level sequence is definition, overview, results, future ([page](https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/)). SIMA 2 uses “A Leap in Generalization Performance” only after the opening and capability explanation, then separates future meaning under “Looking to the Future” ([page](https://deepmind.google/blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/)).

**Implication:** Use stable, functional Analysis-page headings for orientation—such as “What Was Evaluated,” “Results,” “Patterns,” “Limitations,” and “Practical Meaning.” Use conclusion-bearing subheadings only when the statement is directly supported by the figure or table beneath it.

### Finding 3 — Make prose select observations rather than restate the visual

The strongest chart/table linkage follows a compact sequence: identify what the visual contains, extract a small number of decision-relevant observations, then state the boundary or residual gap. The FACTS results prose does not narrate every cell; it selects the overall leader, two component changes, the generally weakest component, and the overall performance ceiling. SIMA 2 captions identify the populations and comparison condition while surrounding prose explains the broader pattern.

**Evidence:** The FACTS page introduces the table’s overall and component columns immediately before it, then summarizes a leader, changes, weakest area, and headroom immediately after it ([“Results”](https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/)). SIMA 2’s chart captions name training versus held-out environments and disclose that the earlier baseline is scored on a newer, harder evaluation set ([performance captions](https://deepmind.google/blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/)).

**Implication:** For each Skill Issue visual, use a three-part prose block: **what the reader is seeing**, **the two or three most important observations**, and **the boundary or unresolved gap**. Avoid converting every plotted value into prose.

### Finding 4 — Put comparability conditions beside the comparison

Comparative language is trustworthy when the comparator, metric, direction, and evaluation conditions are explicit. A changed test set can make a historical baseline appear worse; SIMA 2 places that caveat directly in the chart caption and expands it in the technical report. FACTS distinguishes aggregate score from component scores and specifies that the report’s values average public and private data.

**Evidence:** SIMA 2’s public caption states that the earlier system’s displayed performance comes from a newer, expanded, more difficult evaluation; report §3.4 explains the expanded tasks and stricter success detection and says the older system therefore receives lower success rates ([report](https://arxiv.org/html/2512.04797#S3.SS4)). FACTS Table 1 identifies accuracy, aggregation basis, and 95% confidence intervals in its caption ([report](https://arxiv.org/html/2512.10791#S1.T1)).

**Implication:** Keep changes in harness, rubric, task mix, model version, sample, or scoring logic next to the affected Skill Issue chart or table. Use precise constructions such as “higher on this metric under this evaluation” and reserve broad superiority claims for like-for-like comparisons.

### Finding 5 — Separate observation, interpretation, and forward meaning

The selected pages are easiest to trust when direct observations precede interpretation. FACTS reports the observed ceiling and component pattern, then calls the remaining distance “headroom.” SIMA 2 reports comparative performance, later identifies concrete limitations, and only then discusses possible future relevance. This creates a visible ladder from evidence to meaning.

**Evidence:** FACTS places result values and component patterns in “Results,” then uses “Looking Ahead” for broader direction ([page](https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/)). SIMA 2’s report §5 states current performance, then uses “suggest” and “promising path” for future transfer, followed by specific limitations ([Discussion](https://arxiv.org/html/2512.04797#S5)).

**Implication:** Label the levels of inference in Skill Issue prose. A useful order is: **Observation:** what the displayed data shows. **Interpretation:** what pattern may explain or characterize it. **Practical meaning:** what decision the evidence supports. Use conditional language for the last two levels when causality or generalization has not been established.

### Finding 6 — Surface uncertainty in the public analysis, even when a report has more detail

The linked reports model stronger uncertainty discipline than the summary pages. FACTS Table 1 includes 95% confidence intervals, and later sections validate automated judges against human ratings. The public page is more readable but omits most of that uncertainty detail from its prose. This layered approach is useful, but an analysis page should retain the uncertainty that materially changes interpretation.

**Evidence:** FACTS report Table 1 includes 95% confidence intervals; Table 10 reports judge-validation metrics on a held-out set; §7 names uncovered evaluation areas ([report](https://arxiv.org/html/2512.10791)). SIMA 2 report §3.4 discloses automatic, programmatic, and five-rater human evaluation methods plus baseline time constraints ([report](https://arxiv.org/html/2512.04797#S3.SS4)).

**Implication:** Put essential uncertainty on the Skill Issue Analysis page: sample size, run count, spread or interval when available, known scoring subjectivity, and any condition that changes comparability. Link to deeper methodology for full detail, while keeping decision-relevant uncertainty adjacent to the result.

### Finding 7 — Treat limitations as part of the result, then translate into practical meaning

Limitations work best as a substantive section after results and before expansive conclusions. Both reports identify specific residual weaknesses or coverage gaps rather than relying on a generic disclaimer. This narrows the claim and makes the practical takeaway more credible.

**Evidence:** FACTS report §7 names evaluation areas outside its coverage and calls for more fine-grained analysis ([Conclusion](https://arxiv.org/html/2512.10791#S7)). SIMA 2’s public page and report §5 identify long-horizon reasoning, memory, precise low-level action, and visual understanding as current constraints before discussing possible future application ([page](https://deepmind.google/blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/); [report](https://arxiv.org/html/2512.04797#S5)).

**Implication:** End each major Skill Issue result cluster with the strongest boundary that affects its use. The final “Practical Meaning” section should translate evidence into actions or decisions that remain valid inside those boundaries, rather than presenting the highest score as self-explanatory.

### Finding 8 — Use a layered evidence path from summary to technical detail

Both examples separate public readability from technical completeness without breaking traceability. The page explains why the evaluation matters, provides the visual and selected takeaways, and links the technical report at the methodology or closing point. The report owns evaluation mechanics, validation, fuller results, and limitations.

**Evidence:** FACTS links its technical report directly from the paragraph that defines score construction and evaluation governance ([page](https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/)). SIMA 2 links its report after the limitations and responsible-development sections ([page](https://deepmind.google/blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/)). Both reports expose dedicated evaluation, results, and discussion sections ([FACTS](https://arxiv.org/html/2512.10791); [SIMA 2](https://arxiv.org/html/2512.04797)).

**Implication:** Give every Skill Issue chart or table a nearby source trail: evaluated artifact/version, result source, methodology link, and any detailed report or downloadable data. The Analysis page should remain understandable by itself while allowing a skeptical reader to verify the underlying evidence in one step.

### Finding 9 — A reusable Analysis-page order emerges from both examples

Across the two deep dives, the most adaptable order is: stakes and question; evaluation contract; visual result; selected observations; comparative interpretation; limitations and uncertainty; practical meaning; source trail. This order keeps the reader from seeing a headline conclusion before the basis for that conclusion.

**Evidence:** FACTS follows problem → suite definition → benchmark overview → results → future direction ([page](https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/)). SIMA 2 follows prior context → capabilities → comparative performance → self-improvement → limitations/future → responsibility → report link ([page](https://deepmind.google/blog/sima-2-an-agent-that-plays-reasons-and-learns-with-you-in-virtual-3d-worlds/)).

**Implication:** Adapt the shared sequence rather than either page’s domain-specific story. A concise Skill Issue Analysis page can use: **Question**, **Evaluation Basis**, **Results**, **What Stands Out**, **What the Results Do and Do Not Establish**, **Practical Meaning**, and **Methods and Sources**.

## Notes

- **Caveat:** The public pages sometimes use loaded comparative terms such as “significant” without showing a statistical test on the page. The linked reports provide richer methods, but this assignment found no visible page-level definition that consistently establishes “significant” as statistical significance. Skill Issue should use “statistically significant” only when a named test or interval supports it; otherwise use descriptive comparative language.
- **Unsupported observation excluded:** No inference was made about why one model, system, or benchmark component performed differently. The examples support reporting-form conclusions, not causal explanations transferable to Skill Issue.
- **Useful search terms:** `site:deepmind.google/blog benchmark technical report`, `site:deepmind.google/blog evaluation results chart`, `Google DeepMind benchmark reporting limitations`, `Google DeepMind arXiv evaluations discussion`.

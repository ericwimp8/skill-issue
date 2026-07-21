# Microsoft ADeLe Explanation Patterns

## Assignment

- **Goal:** Extract adaptable patterns from Microsoft’s first-party ADeLe explanation for presenting cross-task performance, connecting visual evidence to prose, separating observations from interpretations, expressing uncertainty, and ending in practical meaning.
- **Scope:** Internet-only inspection of Microsoft’s article and the linked primary paper and project materials, with attention to headings, information order, figure narration, comparison grammar, metric semantics, limitations, and reader-facing implications.
- **Exclusions:** ADeLe-specific scientific conclusions as recommendations for Skill Issue; invented Skill Issue findings or metrics; visual-design implementation details; secondary commentary; evaluation of whether ADeLe itself is scientifically correct beyond consistency checks among its first-party sources.

## Sources

- **Deep-dive:** Microsoft Research, “ADeLe: Predicting and explaining AI performance across tasks,” published 1 April 2026. Inspected “At a glance,” opening problem statement, “ADeLe-based evaluation,” “Evaluating ADeLe,” Figures 1–2 and their captions, and “Looking ahead.” [Article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/)
- **Deep-dive:** Zhou et al., “General scales unlock AI evaluation with explanatory and predictive power,” _Nature_ 652, 58–67 (2026), DOI `10.1038/s41586-026-10303-2`. Inspected Abstract; Main; “Annotation scales distinguish levels and dimensions”; “Explanatory power through benchmark demand profiles”; “Explanatory power through LLM ability profiles”; “Predictive power through assessors anticipating performance”; Discussion; and Methods subsections covering scales, inter-rater analysis, assessors and metrics, subject characteristic curves, and application pipelines. [Nature article](https://www.nature.com/articles/s41586-026-10303-2)
- **Skim-only:** “ADeLe v1.0: A battery for AI Evaluation with explanatory and predictive power.” Inspected overview counts, methodology introduction, “Demand Annotation of Benchmarks on General Scales,” “The ADeLe Battery,” and “Demand Distribution: What Do the Benchmarks Really Test For?” [Project page](https://kinds-of-intelligence-cfi.github.io/ADELE/)
- **Skim-only:** Zhou et al., “General Scales Unlock AI Evaluation with Explanatory and Predictive Power,” arXiv:`2503.06378` (original paper, March 2025). Inspected Abstract, contents, Introduction, Results subsection order, and the “System Process” / “Task Process” description. [arXiv paper](https://arxiv.org/abs/2503.06378)

## Findings

### Finding 1: Order the page as a causal explanation, not a chart inventory

The Microsoft article uses a stable reader journey: concise takeaways; the limitation of ordinary benchmark scores; the shared framework used to reinterpret results; the process diagram; comparative findings; a concrete performance contrast; and only then practical or future-facing meaning. Its headings preserve that progression: “At a glance,” “ADeLe-based evaluation,” “Evaluating ADeLe,” and “Looking ahead.” The primary paper follows a more technical version of the same order: measurement validity, benchmark profiles, system profiles, prediction, discussion, and applications. [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/) [Nature paper](https://www.nature.com/articles/s41586-026-10303-2)

**Evidence:** The blog’s opening bullets state the problem, method, headline predictive result, and explanatory value before any methodological detail. The body then introduces the shared task/model representation before Figures 1 and 2, and it postpones deployment, policy, security-audit, and extension claims until the evidence has been explained. The Nature paper explicitly sequences four research questions from scale interpretability through explanatory and predictive power.

**Implication:** A Skill Issue Analysis page can use the same information order: (1) what the reader should learn, (2) why the aggregate view is insufficient, (3) how the analysis is organized, (4) what each chart establishes, (5) how comparisons relate, (6) where confidence or scope ends, and (7) what a benchmark author or evaluator can do with the result. Charts should serve this argument rather than define the section order merely because they exist.

### Finding 2: Use headings that state the analytical job of each section

The public article uses accessible stage-setting headings, while the paper uses headings that name the kind of evidentiary work being performed, such as “Explanatory power through benchmark demand profiles” and “Predictive power through assessors anticipating performance.” Both approaches orient the reader around a question or purpose instead of around an internal artifact name. [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/) [Nature paper](https://www.nature.com/articles/s41586-026-10303-2)

**Evidence:** The paper introduces each major results section with an explicit research question—for example, whether dimensions can be distinguished, what benchmarks actually measure, whether system capabilities can be understood, and whether unseen performance can be predicted. The blog compresses those questions into the friendlier progression “evaluation,” “evaluating,” and “looking ahead.”

**Implication:** Prefer headings such as “What the benchmark measures,” “Where performance diverges,” “How task difficulty changes the result,” “What the comparison supports,” and “Limits of this analysis.” Avoid headings that only repeat chart types or implementation concepts, because they do not tell the reader why the evidence matters.

### Finding 3: Bridge every chart with encoding, observation, and meaning

The strongest reusable pattern is a three-part bridge around each visual: introduce what has been measured and how the visual encodes it; use the caption to identify the compared entities, axes, grouping, or panels; then use prose to name a small number of visible patterns and explain their analytical meaning. [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/) [Nature paper](https://www.nature.com/articles/s41586-026-10303-2)

**Evidence:** Before Figure 2, the blog explains the 0–5 ability scoring and the 50% success threshold; the caption identifies the 15 models, 18 abilities, and left/middle/right model-family grouping; the following prose identifies the pattern that model generations and model families differ by ability rather than improving uniformly. In the paper, Figure 3’s surrounding text defines the x-axis, y-axis, bin weighting, logistic fit, and the precise meaning of an ability value before interpreting differences among curves.

**Implication:** For every Skill Issue chart, supply: (1) a one-sentence reading key, (2) a caption that identifies exactly what is grouped or compared, (3) an observation sentence anchored in visible values or trends, and (4) a meaning sentence tied to the page’s analytical question. Do not rely on the chart title alone to carry metric definitions or comparison scope.

### Finding 4: Build comparisons from a declared common basis

ADeLe’s explanation repeatedly establishes a common basis before comparing systems or tasks. The blog first says that tasks and models are represented using the same capability scores; the paper then compares task demands with model abilities, and compares model families within the same set of dimensions. This allows the prose to attribute a performance contrast to differences in the compared profiles rather than to a vague overall ranking. [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/) [Nature paper, Fig. 1 discussion](https://www.nature.com/articles/s41586-026-10303-2)

**Evidence:** The paper’s Figure 1 discussion explains differing results across benchmarks carrying the same broad label by contrasting their reasoning and knowledge demands. Its system-profile discussion first reports the broad pattern across model families, then identifies dimensions on which reasoning-oriented models differ, and also states that newer models do not improve consistently across every ability.

**Implication:** Before comparative prose on the Skill Issue page, name the comparison unit and controls: same task set, same scoring rule, same harness conditions, same difficulty band, or same evaluation dimension. Lead with the cross-group pattern, then identify important exceptions. A comparative sentence should be traceable to a chart encoding or statistic, not to an overall impression.

### Finding 5: Separate observations from interpretations through sentence grammar

The sources provide a usable linguistic boundary. Observation sentences name measured direction, range, consistency, or contrast. Interpretation sentences connect those observations to task demands, model properties, benchmark design, or deployment consequences. Interpretive confidence is moderated with terms such as “generally,” “suggests,” “reflecting,” “likely,” and “not consistently.” [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/) [Nature paper](https://www.nature.com/articles/s41586-026-10303-2)

**Evidence:** The blog first observes that newer models generally outperform older ones but not across all abilities, then interprets particular differences in relation to model size, training, or reasoning orientation. The Nature paper distinguishes a correlation observation—most demand dimensions are negatively correlated with success—from the interpretation that this is promising for multivariate prediction. It later says lower out-of-distribution degradation “suggests” less overfitting, rather than presenting the mechanism as directly observed.

**Implication:** Use paired labels or paired sentences: “Observation: …” followed by “Interpretation: …” when the inference could otherwise be mistaken for a direct measurement. Reserve causal language for designs that support it. For ordinary comparisons, use formulations such as “is associated with,” “is consistent with,” or “may reflect,” and state plausible competing explanations when the data do not isolate one cause.

### Finding 6: Define statistical meaning before translating it into prose

The primary paper carefully defines what each statistic means, while the public article sometimes compresses that meaning. The most important discrepancy is the blog’s description of “approximately 88% accuracy” versus the paper’s reported best in-distribution discrimination of `0.882` AUROC for GPT-4o. AUROC is not classification accuracy. The paper separately reports calibration with expected calibration error and shows weaker results under task- and benchmark-out-of-distribution evaluation. [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/) [Nature paper, predictive-power section](https://www.nature.com/articles/s41586-026-10303-2)

**Evidence:** The Nature paper reports a best AUROC of `0.882`, an accuracy-weighted mean AUROC of about `0.84`, and mean ECE of `0.01` in distribution. It reports weighted AUROC/ECE of `0.81`/`0.02` for task-out-of-distribution and `0.75`/`0.04` for benchmark-out-of-distribution. The blog’s “~88% accuracy” removes the metric type, aggregation scope, and distribution condition.

**Implication:** Preserve the exact metric name, comparison population, aggregation rule, and evaluation condition in Skill Issue prose. Translate a statistic only after defining it in reader language—for example, discrimination versus calibration—and avoid converting rates, AUROC, confidence intervals, percent changes, or averages into a generic “accuracy” claim. Place any plain-language shorthand next to the precise statistic rather than replacing it.

### Finding 7: Put uncertainty beside the result and limitations in a dedicated close

The paper handles uncertainty at two levels. Local caveats appear next to the relevant chart or statistic, and a dedicated Discussion section collects structural limitations. The blog retains some hedging and future-scope language but largely omits the paper’s explicit limitation detail, demonstrating why the primary source is necessary when adapting analytical explanation. [Nature paper, Discussion](https://www.nature.com/articles/s41586-026-10303-2) [Microsoft article, “Looking ahead”](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/)

**Evidence:** The paper explains that ability values are expected 50% success thresholds rather than guarantees, notes that correlations are contingent on benchmark choice, reports declining out-of-distribution prediction, identifies incomplete coverage of the multidimensional demand space, and notes scarce high-difficulty items. Its Discussion adds modality exclusions, English-only rubrics, limitations of LLM annotation and grading, and the need for higher scale levels. The blog’s “Looking ahead” instead emphasizes extensibility and potential use.

**Implication:** Put sample-size, missing-data, comparability, and generalization caveats next to the affected result. End with a compact “Limits of this analysis” section that covers coverage, representativeness, metric uncertainty, task/harness dependence, and claims the current data cannot support. Future work should follow those limitations rather than substitute for them.

### Finding 8: Move from evidence to practical action without changing claim strength

Both the article and paper wait until after explanation and prediction results to discuss routing, benchmark design, deployment decisions, auditing, policy, or future extensions. The practical meaning is framed as what the analysis enables, not as a new empirical result. [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/) [Nature paper, application pipeline](https://www.nature.com/articles/s41586-026-10303-2)

**Evidence:** The paper maps practical applications to its “System Process” and “Task Process”: diagnose inconsistent results, design more valid benchmarks, reuse comparable instances, identify strength/demand mismatches, and train decision-support assessors. The blog closes with potential uses in research, policymaking, security auditing, and more rigorous real-world assessment.

**Implication:** End each major Skill Issue finding with a bounded practical consequence, such as what a benchmark author should inspect, what a harness comparison can establish, or what evidence would justify changing a skill. Keep “could inform” or “supports investigating” language when the analysis does not directly test the intervention.

### Finding 9: Use an overview panel as a reading contract

The “At a glance” block works as a compact contract for the rest of the page: it names the evaluation problem, analytical method, headline quantitative result, and practical explanatory payoff. Each bullet is expanded later in the article, so the summary is navigational rather than promotional. [Microsoft article, “At a glance”](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/)

**Evidence:** The four bullets correspond directly to the article’s subsequent sequence: benchmark limitations, shared ability scoring, prediction, and explanation of performance changes with task complexity. The project page uses a complementary pattern by placing corpus counts—instances, tasks/benchmarks, and dimensions—immediately under the title before explaining the methodology. [ADeLe project page](https://kinds-of-intelligence-cfi.github.io/ADELE/)

**Implication:** A Skill Issue Analysis overview can state: evaluation scope, central comparison, strongest supported result, and principal caveat or practical use. If corpus counts are important for trust, place them next to the overview. Every overview claim should link or scroll to the evidence section that substantiates it.

## Notes

- **Candidate classification:** Deep-dive candidates were the Microsoft Research article and final Nature paper because together they expose both public-facing explanation and full metric/limitation semantics. The ADeLe project page and arXiv preprint were skim-only because they mainly repeat or expand the same primary material while helping validate information architecture and paper-section order. Third-party summaries, social posts, and unrelated ADeLe commentary were rejected because first-party primary sources were sufficient.
- **Validation performed:** Cross-checked the blog’s headline claims, figure descriptions, information order, and practical framing against the final Nature paper; confirmed the project page’s overview counts and process framing against the papers; compared the Nature publication with arXiv:`2503.06378` for section-order continuity; and specifically validated that the blog’s “~88% accuracy” shorthand corresponds to a `0.882` AUROC result rather than a reported 88% classification accuracy.
- **Caveat:** The assignment evaluates transferable explanatory patterns and source consistency, not the scientific validity of the ADeLe methodology. No independent statistical replication or dataset reanalysis was performed.
- **Useful search terms:** `ADeLe explanatory power benchmark demand profiles`, `ADeLe AUROC ECE out of distribution`, `ADeLe subject characteristic curves`, `ADeLe limitations`, `arXiv 2503.06378`.

# Benchmark Analysis Page: Deep-Research Synthesis

## Executive Recommendation

The Analysis page should read as a compact public research report, not as a leaderboard landing page. Its central editorial task is to answer one bounded question: **what do the observed Skill Issue evaluations show about model-and-harness combinations invoking expected skills in governed multi-turn scenarios, under the evaluated versions and conditions?** Every score, comparison, interpretation, limitation, and practical statement should remain subordinate to that question.

The strongest cross-source pattern is a layered narrative:

1. **State the evaluation contract.** Define what was evaluated, the comparison unit, the metric direction, the eligible evidence set, the campaign/version, and the practical behavior the score is intended to represent.
2. **Give a small, qualified findings preview.** Use no more findings than the evidence can support, and make each one a complete sentence with its scope and material caveat.
3. **Present the primary comparative evidence.** Treat the model-and-harness combination as the first-order result object; then use matched model, harness, scenario, and failure-pattern breakdowns to explain heterogeneity.
4. **Pair every visual with one analytical job.** The chart shows pattern, its caption explains how to read it, prose selects the important observations, and a table or evidence link preserves exact values.
5. **Separate observation, estimate, interpretation, and practical meaning.** These are different claim levels and should use different language.
6. **Make uncertainty and comparability local.** Put the relevant interval, run count, exclusion, version break, or scoring condition beside the claim it changes.
7. **End with bounded practical meaning and an audit trail.** Explain what the evidence can inform for benchmark authors, skill authors, and users of the tested combinations, then link to exact results, run artifacts, methodology, and change history.

This order is repeatedly supported by public benchmark-reporting examples that move from question and metric definition to evidence, interpretation, limitations, and sources, including Google DeepMind's FACTS reporting, Stanford HELM, NIST CAISI, METR, Epoch, and Artificial Analysis ([DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md), [FACTS page](https://deepmind.google/blog/facts-benchmark-suite-systematically-evaluating-the-factuality-of-large-language-models/); [HELM assignment](./assignments/10-stanford-helm-reporting.md), [HELM report](https://arxiv.org/pdf/2211.09110); [NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md), [CAISI evaluation](https://www.nist.gov/news-events/news/2026/05/caisi-evaluation-deepseek-v4-pro); [METR assignment](./assignments/11-metr-benchmark-reporting.md), [Time Horizon 1.1](https://metr.org/blog/2026-1-29-time-horizon-1-1/); [Epoch assignment](./assignments/02-epoch-benchmark-presentation.md), [Epoch benchmarking](https://epoch.ai/benchmarks); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md), [v4.1 article](https://artificialanalysis.ai/articles/artificial-analysis-intelligence-index-v4-1)).

No source in the supplied corpus establishes any actual Skill Issue result, ranking, effect size, uncertainty interval, causal mechanism, or preferred chart type for the current data. The page's eventual finding sentences and statistical claims must therefore be populated only from the completed Skill Issue evidence.

## Editorial Position

### Explain measurement before celebrating rank

The page should frame the benchmark as **measurement under stated conditions**, not a contest that automatically identifies a universal winner. Anthropic begins its statistical evaluation guidance with the ambiguity between a real difference and the sampled questions; NIST distinguishes an observed score from an estimated performance target; BetterBench requires a chain from construct to task to score interpretation; and OpenAI's audit writing defines what a success or failure is intended to represent before discussing artifacts ([Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md), [Anthropic article](https://www.anthropic.com/research/statistical-approach-to-model-evals); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md), [NIST AI 800-3](https://tsapps.nist.gov/publication/get_pdf.cfm?pub_id=961314); [BetterBench assignment](./assignments/08-betterbench-validity.md), [BetterBench methodology](https://betterbench.stanford.edu/methodology.html); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md), [OpenAI audit](https://openai.com/index/separating-signal-from-noise-coding-evaluations/)).

For Skill Issue, the public construct-to-evidence chain should be explicit:

- **Construct:** correct invocation of an expected skill during a governed multi-turn coding-agent scenario.
- **Evaluated unit:** the actual unit used by the benchmark, such as a scenario attempt, scenario-level outcome, or another documented aggregation unit.
- **Condition:** the named model, harness, reasoning/configuration, skill availability, benchmark version, and scoring rule used for that result.
- **Observed result:** what occurred in the recorded evaluation evidence.
- **Supported use:** comparison or diagnosis within the tested evaluation frame.
- **Unsupported extension:** any broader model, harness, coding ability, production suitability, or causal claim that the design did not directly establish.

This should be stated in plain language near the top. It should also be reflected in every comparison: the benchmark directly observes **configurations in evaluated scenarios**. A model-only or harness-only conclusion requires a matched comparison that holds the other relevant conditions constant. An overall agent-quality conclusion would require evidence beyond skill invocation in the tested scenarios.

### Make restraint the primary style

The restrained style should come from evidence structure rather than sparse content. Epoch and Artificial Analysis achieve density through wide evidence regions, light dividers, direct labels, repeated chart grammar, limited color, and typography-driven hierarchy; neither requires decorative landing-page cards to make a result legible ([Epoch assignment](./assignments/02-epoch-benchmark-presentation.md), [SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md), [v4.1 charts](https://artificialanalysis.ai/articles/artificial-analysis-intelligence-index-v4-1)). HELM and METR similarly rely on a clear progression from short summary to evidence and audit trail rather than visual ornament ([HELM assignment](./assignments/10-stanford-helm-reporting.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md)).

The page should therefore favor:

- short introductory prose with a readable line length;
- functional top-level headings;
- sentence-style finding subheadings only when supported;
- wide chart and table regions;
- light rules, compact labels, and direct value annotation;
- stable colors for stable semantic categories;
- visible but quiet caveat and provenance text;
- progressive disclosure instead of stacked promotional panels.

## Recommended Page Narrative

The following is the strongest adaptable section architecture. It is a narrative recommendation, not a downstream implementation plan.

### 1. Analysis

Open with one bounded synthesis sentence that names:

- what was analyzed;
- the campaign or benchmark version;
- the strongest supported observed pattern;
- the comparison scope that makes the statement valid.

The opening may be answer-first, as in NIST CAISI, METR, and OpenAI's audit reports, but it must immediately expose the evidence basis and cannot use a dramatic thesis that later caveats substantially reverse ([NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md)).

**Adaptable pattern**

> Across **[eligible configurations]** in **[campaign/version]**, the evaluations recorded **[bounded pattern]** on **[metric and unit]**. This result describes **[tested scope]**; **[material limitation]** constrains broader interpretation.

If the data does not support one stable lead, the opening should say so directly:

> No single ordering held across **[scenario/harness/model breakdown]**. The clearest result is the variation by **[supported dimension]**, while **[uncertainty or comparison break]** prevents a broader winner claim.

### 2. At a Glance

Use a short overview block as a reading contract. Microsoft's ADeLe article and NIST CAISI show that a compact preview is useful when each bullet maps to a later evidence section rather than serving as promotional copy ([Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md), [Microsoft article](https://www.microsoft.com/en-us/research/blog/adele-predicting-and-explaining-ai-performance-across-tasks/); [NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md)).

The block should contain only supported items such as:

- **Evaluation scope:** number of eligible model-and-harness combinations, scenarios, attempts or runs, and evaluation date/version.
- **Primary observation:** the strongest evidence-bound pattern.
- **Heterogeneity:** the most important scenario, model, or harness exception to the aggregate.
- **Principal caveat:** the uncertainty or comparability condition most likely to change a reader's conclusion.
- **Practical use:** the narrow decision the analysis can inform.

Counts belong here only when they improve trust or define the evidence population. The overview should not compress a multi-dimensional result into a single badge without preserving the disaggregated authority.

### 3. What Was Evaluated

This is the evaluation contract. Keep it compact but sufficient to read every following visual. DeepMind, HELM, Epoch, BetterBench, Benchscope, and NIST all support defining the evaluation object, metric, comparison universe, and aggregation before asking readers to interpret outcomes ([DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [BetterBench assignment](./assignments/08-betterbench-validity.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md)).

Include:

- the behavior the benchmark is intended to measure;
- the evaluated unit and eligible population;
- the model, harness, configuration, and scenario dimensions;
- what counts as correct skill invocation;
- the metric, denominator, direction, and aggregation rule;
- whether attempts are matched across configurations;
- scenario and repetition counts;
- benchmark, rubric, harness, model, and skill versions;
- exclusions, incomplete results, and non-comparable categories;
- a link to the controlling methodology.

If an aggregate exists, define why it is useful and what it hides. HELM explicitly treats aggregation as an editorial and value-laden decision; Artificial Analysis publishes component weights; BetterBench warns that weighting and coarse categories constrain meaning ([HELM assignment](./assignments/10-stanford-helm-reporting.md), [HELM Lite](https://crfm.stanford.edu/2023/12/19/helm-lite.html); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [BetterBench assignment](./assignments/08-betterbench-validity.md)). The aggregate should orient the reader, while scenario- and condition-level results remain the evidentiary authority.

### 4. Results

Use functional **Results** as the top-level heading. Beneath it, use sentence-style subheadings that state only what the displayed evidence directly supports. METR's declarative findings and HELM's short evidence-bearing labels are strong models; DeepMind's neutral top-level headings prevent the whole page from becoming a sequence of premature conclusions ([METR assignment](./assignments/11-metr-benchmark-reporting.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md)).

The preferred result order is:

1. **Overall configuration results.** Compare the complete model-and-harness combinations on the declared primary metric.
2. **Matched model and harness comparisons.** Within the same harness, compare models on shared scenarios; within the same model, compare harnesses on shared scenarios. This avoids assigning a combination effect to only one component.
3. **Scenario heterogeneity.** Show where the aggregate persists, reverses, or becomes unresolved across governed scenario types.
4. **Consistency and uncertainty.** Show run-to-run variation, scenario concentration, intervals, or supported sensitivity evidence.
5. **Failure-pattern or evidence review.** If the benchmark includes audited categories or trace review, show the aggregate category pattern before one clearly labeled illustrative case.

This ordering follows the cross-source preference for aggregate orientation followed by heterogeneity and exceptions. HELM explicitly moves from aggregate rankings to scenario-level changes; NIST shows that similar means can conceal different variance structures; Artificial Analysis preserves subgroup leaders; OpenAI moves from prevalence to a representative traceable case only after establishing the aggregate ([HELM assignment](./assignments/10-stanford-helm-reporting.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md)).

### 5. What the Results Show

Each finding block should use a fixed epistemic sequence:

1. **Observation:** what the recorded evidence or displayed statistic shows.
2. **Estimate:** the target quantity inferred from those observations, if the analysis actually estimates one.
3. **Interpretation:** a supported characterization or plausible explanation.
4. **Practical meaning:** the decision the result can inform inside its scope.
5. **Boundary:** what the evidence does not establish.

NIST explicitly separates realized outcomes, estimated quantities, and interpretation. Anthropic, Microsoft, OpenAI, HELM, METR, BetterBench, Benchscope, and DeepMind independently reinforce a visible observation-to-interpretation boundary ([NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md); [Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [BetterBench assignment](./assignments/08-betterbench-validity.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md)).

Visible labels such as **Observed**, **Interpretation**, and **Practical meaning** are appropriate when prose could otherwise blur these levels. For short findings, separate sentences can do the same job without creating repeated UI chrome.

### 6. What the Results Do and Do Not Establish

This section should consolidate cross-cutting limits after local caveats have already appeared beside affected claims. It should be organized by the source of uncertainty rather than as one generic disclaimer:

- **Coverage:** which scenarios, skills, harnesses, models, languages, task types, and operating conditions were included or omitted.
- **Measurement validity:** whether the scoring rule captures correct skill invocation and where rubric, judge, or artifact ambiguity remains.
- **Sampling and repetition:** the independent unit, number of scenarios, runs per scenario, and any clustering or repeated-measure dependence.
- **Comparison validity:** matched versus unmatched attempts, configuration differences, eligibility rules, missing results, and version changes.
- **Aggregation:** weighting, thresholds, subgroup imbalance, and trade-offs hidden by an overall rate.
- **Generalization:** why results describe the tested configurations and do not automatically establish general model quality, general harness quality, production suitability, or real-world task completion.

This two-level caveat model is well supported: Artificial Analysis, Epoch, NIST CAISI, METR, HELM, Benchscope, and DeepMind put local qualifications near the affected result while retaining a broader limitations or methodology layer ([Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md)).

### 7. Practical Meaning

Translate findings into choices only after the comparison basis and limitations are established. Practical meaning should be audience-specific and conditional:

- **For benchmark readers:** which model-and-harness combinations were more consistently associated with correct invocation in the tested scenarios, and where the evaluation could not resolve a difference.
- **For skill authors:** which scenario types or failure concentrations warrant inspection, without assuming the skill itself caused every failure.
- **For harness authors:** which matched within-model comparisons suggest a harness-associated pattern, while preserving other configuration differences and uncertainty.
- **For model evaluators:** which matched within-harness comparisons suggest a model-associated pattern within the tested setup.
- **For benchmark maintainers:** where additional repetitions, clearer scoring, broader coverage, or setup validation would most improve inference.

HELM and Benchscope explicitly translate multi-dimensional evidence into priority-dependent choices rather than a universal winner. NIST and Microsoft place practical consequence after statistical or explanatory findings. OpenAI makes action proportional to the audited evidence ([HELM assignment](./assignments/10-stanford-helm-reporting.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md)).

Use verbs such as **supports choosing**, **supports investigating**, **is relevant when**, or **cannot resolve**. Avoid turning a benchmark-bound observation into a recommendation for every coding task or production environment.

### 8. Methods and Sources

End with a compact verification ladder rather than an abbreviated second methodology chapter:

- campaign and evaluation version;
- exact result set or immutable artifact identifier;
- chart/table data source;
- model, harness, reasoning, skill, rubric, and scenario versions;
- exclusions and transformation/aggregation rules;
- methodology and change-history links;
- run-level logs, transcripts, scoring evidence, or downloadable data where available.

Epoch, HELM, BetterBench, Benchscope, Artificial Analysis, and DeepMind all support progressive traceability from summary to exact values, methods, provenance, or underlying artifacts ([Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [BetterBench assignment](./assignments/08-betterbench-validity.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md)). BetterBench's own source-version discrepancies particularly support binding public claims to a named controlling artifact rather than silently merging current pages, papers, and supplements ([BetterBench assignment](./assignments/08-betterbench-validity.md)).

## Chart, Table, and Prose Rules

### Give each visual one analytical job

Every chart should answer a named question, such as:

- Which evaluated configurations recorded the highest and lowest primary rates?
- Does a model ordering persist within each harness?
- Does a harness ordering persist within each model?
- Where do scenario-level results reverse the aggregate?
- How much run-to-run or scenario-to-scenario variation is visible?
- Which failure categories account for the observed total?

NIST's figures, METR's report sequence, CAISI's tables, HELM's evidence anchors, and Artificial Analysis's paired bar/scatter views all support assigning visuals distinct analytical jobs instead of assembling a chart inventory ([NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md)).

### Use a stable visual-to-text sequence

For every primary visual:

1. **Finding or question heading:** tells the reader why the visual exists.
2. **Reading key:** defines measure, direction, eligible population, and grouping in one sentence.
3. **Chart or table:** carries the distribution, ordering, trade-off, or exact comparison matrix.
4. **Literal caption:** names units, denominator, sample/run counts, aggregation, uncertainty encoding, exclusions, version, and comparability notes.
5. **Observation prose:** selects two or three decision-relevant patterns, with enough values to remain auditable.
6. **Interpretation and boundary:** explains the supported meaning and the most important unresolved alternative.
7. **Exact-value or evidence exit:** preserves the rows, results, or artifacts behind the visual.

This division gives each layer a different responsibility. Charts reveal shape; captions explain encoding; prose states the bounded takeaway; tables preserve exact values. METR states this division explicitly in practice, while NIST CAISI, Epoch, HELM, DeepMind, Microsoft, and Artificial Analysis demonstrate compatible versions of it ([METR assignment](./assignments/11-metr-benchmark-reporting.md); [NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md); [Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md); [Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md)).

### Treat chart and table as peer views when they share evidence

When a chart and exact-value table represent the same result set, they should preserve the same filters, eligibility state, units, and version. Epoch's chart/leaderboard switch treats them as two representations of one evidence object: the chart supports pattern recognition; the table supports exact comparison and trace links ([Epoch assignment](./assignments/02-epoch-benchmark-presentation.md), [SWE-bench detail](https://epoch.ai/benchmarks/swe-bench-verified)).

This is preferable to separate chart and table narratives that drift in scope. If the table includes contextual or non-comparable rows, label them inline rather than relying on a distant note.

### Keep uncertainty visible in every result representation

If a justified uncertainty measure exists, show it in the chart, exact-value view, and prose. State what the interval or spread represents and what it omits. Epoch uses error bars, plus/minus values, and explanatory methodology; NIST treats uncertainty as a result; DeepMind's technical reports include confidence intervals; METR puts interval meaning in captions and remaining uncertainty in prose ([Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md)).

If the data does not support a defensible interval, show the underlying counts, repetitions, and dispersion that do exist and describe differences as observed or descriptive. Do not manufacture statistical certainty to make the chart appear complete.

## Comparative Findings and Statistical Language

### Primary comparison unit

The headline evidence object should be the **model-and-harness combination under a named configuration**, because that is what the benchmark directly runs. Decompose the result through matched comparisons:

- compare models within the same harness on shared scenarios;
- compare harnesses within the same model on shared scenarios;
- compare combinations across scenario types only after stating whether the cases and rules match;
- report unmatched or changed-version results as contextual evidence, not as if they were paired.

Anthropic supports paired differences when systems share cases; METR and DeepMind stress defining the common comparison basis; Benchscope labels claim eligibility before ranking; BetterBench requires a compatible comparison universe ([Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [BetterBench assignment](./assignments/08-betterbench-validity.md)).

### Required contents of a comparison sentence

A strong comparison sentence should identify or inherit visibly:

- the eligible population;
- evaluation version and date;
- metric and direction;
- focal configuration and comparator;
- shared or unmatched case basis;
- observed magnitude or count;
- uncertainty or descriptive status;
- any condition that changes interpretation.

**Adaptable patterns**

> Within **[version]**, across **[n shared scenarios]**, **[configuration A]** recorded **[value]** correct invocations and **[configuration B]** recorded **[value]**, a paired difference of **[difference and interval, if supported]**.

> Under **[same harness and scoring rule]**, **[model A]** was higher on the displayed metric in **[count] of [count]** eligible scenarios. The ordering reversed in **[named subgroup]**.

> With **[same model and shared scenario set]**, the observed rate was higher for **[harness A]** than **[harness B]**. **[Uncertainty statement]** limits whether this should be treated as a resolved difference.

> Because **[task mix, harness version, rubric, model version, or scoring rule]** changed, these results are shown for context and are not treated as a direct longitudinal comparison.

> The benchmark did not resolve a difference of this size between **[A]** and **[B]**. This does not establish that the configurations are equivalent.

Artificial Analysis and NIST CAISI demonstrate compact comparisons that name cohort, values, comparator, and qualifier; HELM adds denominator and aggregation; Anthropic supplies the crucial distinction between no detected difference and equality ([Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md)).

### Statistical terms

Use **statistically significant** only when a named, appropriate procedure supports it. The DeepMind assignment found public uses of “significant” without a visible page-level test, and BetterBench warns that a point difference is not evidence of a reliable difference ([DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md); [BetterBench assignment](./assignments/08-betterbench-validity.md)).

Prefer:

- **higher/lower on this metric** for a descriptive ordering;
- **observed difference** for a recorded point difference;
- **estimated difference** when a defined estimator is used;
- **the interval includes both orderings** when supported;
- **the benchmark did not resolve a difference** when precision is insufficient;
- **equivalent within [predefined margin]** only if an actual equivalence rule exists;
- **consistent across [tested variants]** only when those variants were tested.

Avoid:

- **outperforms**, **best**, **wins**, or **better** without a valid metric direction and comparable evidence;
- **tie** based only on a non-significant result;
- generic **accuracy** when the statistic is another quantity;
- generic **confidence** without defining whether it refers to an interval, reviewer agreement, stability, or subjective certainty;
- causal **because** when the evidence only shows association.

Microsoft's ADeLe example is a useful warning: a public article translated an AUROC result into generic “accuracy,” while the paper preserved the actual metric and evaluation condition. Skill Issue should keep the exact metric name, population, aggregation, and condition in public prose ([Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md), [Nature paper](https://www.nature.com/articles/s41586-026-10303-2)).

### Repeated evaluations and sensitivity

If Skill Issue uses repeated attempts, report:

- number of scenarios or items;
- attempts per scenario;
- independent and clustered units;
- aggregation order;
- run-to-run and scenario-to-scenario variation;
- whether comparisons are paired;
- what changed in any sensitivity analysis.

NIST warns that repeated trials on the same item are clustered rather than freely independent; Anthropic recommends cluster-aware uncertainty and paired comparisons; BetterBench argues that stochastic evaluations need repeated evidence to distinguish signal from run noise ([NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md); [BetterBench assignment](./assignments/08-betterbench-validity.md)). These sources do not establish which statistical model fits Skill Issue. The appropriate interval, clustering unit, estimator, and power analysis remain dependent on the actual data-generating process.

Sensitivity prose should state:

> We varied **[factor]** while holding **[conditions]** fixed. The **[finding]** did/did not persist across **[tested variants]**. This analysis does not address **[untested source of sensitivity]**.

## Observation, Interpretation, and Causality Boundaries

### Preferred verb discipline

Use verbs as epistemic labels:

- **recorded, observed, returned, selected, invoked, passed, failed** for direct evaluation evidence;
- **estimated** for a modeled or extrapolated quantity;
- **was associated with, suggests, is consistent with, may reflect** for interpretation;
- **supports choosing, supports investigating, is relevant to** for practical meaning;
- **cannot establish, remains unresolved, was not tested** for boundaries.

OpenAI, METR, NIST, Microsoft, Anthropic, HELM, and DeepMind all use or recommend comparable distinctions between measured events, interpretation, and action ([OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md); [Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md)).

### Do not infer mechanisms from aggregate outcomes

An observed difference between model-and-harness combinations does not by itself identify whether the mechanism is model behavior, harness routing, prompt construction, skill availability, tool exposure, state handling, rubric behavior, or run noise. The page may name a mechanism only when the evaluation directly varies or audits the relevant factor. Otherwise it should retain competing explanations.

OpenAI's audit narrative shows how an aggregate can be decomposed through independent review and one evidence chain; NIST treats extreme item difficulty as potentially genuine or problematic until inspected; Anthropic warns that implementation and formatting choices can alter results ([OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md)).

**Adaptable pattern**

> **Observed:** **[configuration]** failed to invoke the expected skill in **[count/scope]**.
>
> **Interpretation:** The concentration in **[supported subgroup]** is consistent with **[hypothesis]**, but the evaluation does not isolate that mechanism from **[plausible alternatives]**.
>
> **Practical meaning:** Inspect **[specific evidence path]** before changing the skill, harness, or benchmark rule.

## Limitations and Uncertainty Placement

### Global limits belong beside the metric definition

The opening contract should disclose limitations that govern the whole page:

- fixed evaluated scenario set versus a broader scenario population;
- model/harness/configuration version scope;
- number of scenarios and repetitions;
- known scoring or reviewer subjectivity;
- missing conditions or incomplete runs;
- whether results generalize only to the governed benchmark setup;
- whether uncertainty has or has not been quantified.

Artificial Analysis, NIST, BetterBench, HELM, and DeepMind all support placing material scope constraints near the metric rather than waiting for a footer ([Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [BetterBench assignment](./assignments/08-betterbench-validity.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md)).

### Local caveats belong beside affected claims

Attach the following where they occur:

- changed benchmark or scenario version;
- changed harness, rubric, evaluator, prompt, tool, or model version;
- unmatched cases;
- missing or imputed values;
- excluded attempts and reasons;
- unequal evidence counts;
- unstable or wide intervals;
- reviewer disagreement or overlapping categories;
- eligibility status;
- selection or curation conditions.

CAISI's local notes, Epoch's changelog and above-chart version warning, METR's matched-comparison explanations, Benchscope's eligibility labels, and OpenAI's review-path disagreements make these caveats part of the result rather than background administration ([NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md); [Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md)).

### “Supports / does not establish” is the preferred compact boundary

BetterBench explicitly recommends positive and negative score-interpretation guidance. Each major finding can end with a compact pair ([BetterBench assignment](./assignments/08-betterbench-validity.md), [BetterBench methodology](https://betterbench.stanford.edu/methodology.html)):

> **Supports:** **[specific comparison or decision inside the evaluated scope]**.
>
> **Does not establish:** **[causal, general, production, or cross-version claim not supported by the evidence]**.

This is clearer than a generic warning and keeps the limitation proportional to the claim.

## Conditional Alternatives

### If the evidence supports one stable headline pattern

Use an answer-first lead, a three-to-five-item findings preview, and sentence-style finding headings. This is the strongest fit when the same directional result persists across relevant matched comparisons and uncertainty does not materially change the conclusion. NIST CAISI, METR, and OpenAI provide good public-report models for this shape ([NIST CAISI assignment](./assignments/05-nist-caisi-reporting-example.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md)).

### If results are heterogeneous or rankings reverse

Lead with the lack of a universal ordering, then organize by comparison dimension or scenario family. Use parallel subsections for model-within-harness, harness-within-model, and scenario-specific effects. HELM, Benchscope, Artificial Analysis, and NIST's variance examples support preserving trade-offs and exceptions instead of forcing one winner ([HELM assignment](./assignments/10-stanford-helm-reporting.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md)).

### If uncertainty is too large for reliable ordering

Make sensitivity the finding. Show estimates with counts and available dispersion, state the range of compatible orderings, and explain what additional evidence would be needed to resolve a practically relevant gap. Anthropic's power framing, NIST's uncertainty-first comparisons, and BetterBench's warning against reading point differences as reliable differences support this alternative ([Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [BetterBench assignment](./assignments/08-betterbench-validity.md)).

### If the strongest insight comes from audited traces rather than aggregate rates

Lead with the aggregate measurement contract and prevalence or category summary, then explain independent review paths and present one illustrative evidence chain. Do not use the example to establish prevalence. OpenAI's benchmark-audit narrative is the clearest model for this case ([OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md)).

### If the current campaign is not comparable with an earlier campaign

Treat the evaluation revision as a result condition. State what changed before showing movement, separate overlapping from non-overlapping evidence, and avoid a continuous trend line unless comparability is justified. METR, Epoch, Artificial Analysis, DeepMind, and BetterBench all show why version identity and comparison breaks belong in the analytical narrative ([METR assignment](./assignments/11-metr-benchmark-reporting.md); [Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md); [BetterBench assignment](./assignments/08-betterbench-validity.md)).

## Lower-Fit or Rejected Patterns

### A decorative leaderboard landing page

This is lower-fit because it makes rank the dominant semantic object before scope, uncertainty, and practical meaning. The supplied research consistently favors a report or evidence-explorer structure with a verification path. A leaderboard may remain an exact-value view, but it should not own the analysis narrative ([Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [HELM assignment](./assignments/10-stanford-helm-reporting.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md)).

### A chart wall ordered by available visualizations

Microsoft explicitly supports a causal/explanatory page order rather than a chart inventory; NIST and METR give each figure a declared question or finding. Chart availability should not determine information architecture ([Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [METR assignment](./assignments/11-metr-benchmark-reporting.md)).

### A universal composite winner

HELM shows that aggregation depends on comparison set and stakeholder values; Benchscope separates outcome dimensions; Artificial Analysis preserves subgroup distinctions; BetterBench warns that weighting constrains meaning. A composite can orient readers only if its rule and limitation are explicit, and it should not erase scenario or configuration reversals ([HELM assignment](./assignments/10-stanford-helm-reporting.md); [Benchscope assignment](./assignments/09-benchscope-methodology.md); [Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [BetterBench assignment](./assignments/08-betterbench-validity.md)).

### Promotional or one-off event narration

The DeepMind assignment rejected a one-off competition story as a poorer fit than a reusable benchmark report. The Epoch assignment rejected promotional product pages, and the OpenAI assignment rejected unrelated framework material. Skill Issue should avoid launch-style claims, dramatic superlatives, or event framing that is not required to explain the benchmark evidence ([DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md); [Epoch assignment](./assignments/02-epoch-benchmark-presentation.md); [OpenAI audit assignment](./assignments/07-openai-benchmark-audit.md)).

### Formula-first public exposition

Anthropic and NIST show the value of linking a readable public explanation to full statistical methods. The Analysis page should define technical terms and consequences in prose, then expose formulas and assumptions through methodology when needed ([Anthropic assignment](./assignments/03-anthropic-statistical-evaluations.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md)).

### Caveats collected only at the end

Every major source group supports local caveats where they change a claim. A final limitations section remains necessary, but it should consolidate rather than reveal the decisive qualification for the first time.

### Causal stories inferred from rank differences

No supplied research supports attributing a Skill Issue result to model quality, harness quality, skill quality, or a particular failure mechanism without a design or audit that isolates it. Mechanism prose should remain explicitly interpretive and preserve alternatives.

### Copying another benchmark's domain concepts or rhetoric

The sources were selected for presentation and semantic patterns. Their model rankings, benchmark taxonomies, capability labels, statistical models, chart types, weights, thresholds, and practical conclusions are not transferable evidence for Skill Issue. Several assignment documents explicitly reject this transfer ([Artificial Analysis assignment](./assignments/01-artificial-analysis-presentation.md); [NIST statistical assignment](./assignments/04-nist-statistical-toolbox.md); [Microsoft ADeLe assignment](./assignments/06-microsoft-adele-explanation.md); [DeepMind assignment](./assignments/12-google-deepmind-benchmark-reporting.md)).

## Source-Corpus Context

The twelve assignments play complementary roles in this recommendation:

- **Direct public benchmark presentation:** Artificial Analysis, Epoch, NIST CAISI, Stanford HELM, METR, and Google DeepMind provide the strongest evidence for information order, heading strategy, chart/table-to-prose relationships, local caveats, practical meaning, and traceability ([01](./assignments/01-artificial-analysis-presentation.md), [02](./assignments/02-epoch-benchmark-presentation.md), [05](./assignments/05-nist-caisi-reporting-example.md), [10](./assignments/10-stanford-helm-reporting.md), [11](./assignments/11-metr-benchmark-reporting.md), [12](./assignments/12-google-deepmind-benchmark-reporting.md)).
- **Statistical claim discipline:** Anthropic and NIST provide the strongest support for sample information, paired comparisons, clustering, sensitivity, uncertainty, estimands, and language for unresolved differences ([03](./assignments/03-anthropic-statistical-evaluations.md), [04](./assignments/04-nist-statistical-toolbox.md)).
- **Validity and provenance constraints:** BetterBench supplies the construct-to-task-to-meaning chain, reproducibility expectations, positive and negative score interpretation, and artifact-version caution ([08](./assignments/08-betterbench-validity.md)).
- **Methodology-to-analysis linkage:** Benchscope supplies the clearest pattern for a methodology contract, local eligibility labels, dimension-specific comparisons, and conditional selection guidance ([09](./assignments/09-benchscope-methodology.md)).
- **Cross-task explanation:** Microsoft's ADeLe materials support analytical-job headings, shared comparison bases, chart bridges, and careful translation of statistics into prose, while also exposing the risk of metric shorthand ([06](./assignments/06-microsoft-adele-explanation.md)).
- **Audit narrative:** OpenAI supplies the strongest pattern for signal-versus-artifact framing, independent evidence channels, aggregate-to-case explanation, and action proportional to the audit ([07](./assignments/07-openai-benchmark-audit.md)).

The candidate classifications also reinforce what to exclude. Across assignments, skim-only sources were commonly dashboards, data-access pages, repeated templates, project overviews, or supportive technical artifacts. Rejected candidates were commonly secondary summaries, promotional pages, unrelated product surfaces, one-off competitions, domain-specific extensions, and inaccessible interactive details. This supports using the deep-dive report patterns for the page's semantic structure while treating dashboard affordances and visual interactions as conditional product choices rather than research-established requirements.

## Unsupported Claims and True Blockers

The supplied corpus establishes a strong reporting direction, but it does not establish the following:

- actual Skill Issue findings, rankings, values, intervals, effect sizes, or subgroup patterns;
- the correct statistical estimator, clustering unit, interval method, equivalence margin, or power target for Skill Issue;
- whether current attempts are independent, paired, repeated, or hierarchically clustered;
- whether the benchmark's existing aggregate is valid or whether any new aggregate is warranted;
- which chart types best fit the actual Skill Issue data volume and responsive layout;
- whether run-level logs, transcripts, raw scoring evidence, or downloadable data are currently publishable;
- whether any observed failure can be causally attributed to the model, harness, skill, prompt, rubric, or evaluation infrastructure;
- whether current and historical campaigns are directly comparable;
- whether the current evidence supports a claim-first page or requires an uncertainty-first or heterogeneity-first alternative.

These are not blockers to the editorial recommendation. They are blockers to writing the eventual findings copy and choosing the final statistical and visual treatment. Until the Skill Issue result set and data-generating process are inspected, the Analysis page should use placeholders only for structure, never for invented conclusions.

## Final Direction

The strongest direction is a restrained, source-traceable public analysis with **functional navigation, evidence-bearing finding headings, combination-first comparisons, matched model/harness breakdowns, scenario heterogeneity, visible uncertainty, explicit observation-versus-interpretation language, local caveats, and bounded practical meaning**. The page should let a reader recover the argument by scanning headings, verify it through charts and exact values, understand its limits without hunting for footnotes, and reach the underlying methodology and run evidence in one step.

The page succeeds when its prose makes the benchmark easier to interpret without making the evidence sound stronger than it is.

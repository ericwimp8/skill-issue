# Anthropic Statistical Evaluation Reporting Patterns

## Assignment

**Goal:** Determine how a public Skill Issue benchmark Analysis page should express uncertainty, rankings and differences, sensitivity, observations versus interpretations, and practical meaning when reporting model/harness skill-calling evaluations. Extract transferable presentation patterns from Anthropic's public statistical-evaluation material without importing its benchmark conclusions.

**Scope:** Internet-only review of Anthropic's primary article, its underlying Anthropic-authored paper, and the directly linked first-party evaluation-challenges article needed to qualify interpretation. The findings concern reporting structure and analytical language; their application to Skill Issue is conditional on the benchmark's actual sampling units, dependence structure, repetitions, and available statistics.

**Exclusions:** No Skill Issue results, rankings, effect sizes, confidence intervals, or statistical significance claims are inferred or invented. No claim is made that Anthropic's formulas automatically fit Skill Issue's design. Third-party statistical tutorials, benchmark results, and framework documentation are outside the selected first-party scope.

## Sources

### Deep-dive sources

1. Evan Miller, **“A statistical approach to model evaluations”** (Anthropic, November 19, 2024), https://www.anthropic.com/research/statistical-approach-to-model-evals
   - Inspected: opening problem statement; “Recommendation #1: Use the Central Limit Theorem”; “Recommendation #2: Cluster standard errors”; “Recommendation #3: Reduce variance within questions”; “Recommendation #4: Analyze paired differences”; “Recommendation #5: Use power analysis”; “Conclusion.”
   - Purpose: primary public-facing example of how Anthropic orders, explains, and connects statistical recommendations to practical interpretation.
2. Evan Miller, **“Adding Error Bars to Evals: A Statistical Approach to Language Model Evaluations”** (Anthropic; arXiv:2411.00640), https://arxiv.org/abs/2411.00640 and readable HTML at https://ar5iv.labs.arxiv.org/html/2411.00640
   - Inspected: Abstract; §1 “Introduction”; §2.1 “Independent questions”; §2.2 “Clustered questions”; §3 “Variance reduction,” including §§3.1–3.3; §4 “Comparing models,” including §§4.1–4.2 and Table 5; §5 “Power analysis”; §6 “Conclusion.”
   - Purpose: technical validation of the article's simplified claims and exact recommended reporting fields.

### Skim-only candidate

3. Deep Ganguli, Nicholas Schiefer, Marina Favaro, and Jack Clark, **“Challenges in evaluating AI systems”** (Anthropic, October 4, 2023), https://www.anthropic.com/news/evaluating-ai-systems
   - Inspected: “Introduction”; “The supposedly simple multiple-choice evaluation”; “One size doesn't fit all when it comes to third-party evaluation frameworks”; “Conclusion.”
   - Purpose: first-party context for construct validity, implementation sensitivity, and the risk of overinterpreting a technically correct score.

### Rejected candidates

- General statistical explainers linked from the article, including Wikipedia pages for the Central Limit Theorem, confidence intervals, clustered standard errors, paired differences, and power analysis: rejected because the assignment selected first-party Anthropic material and the Anthropic paper supplies the needed technical context.
- Third-party benchmark papers and evaluation-framework documentation linked by Anthropic: rejected because their individual benchmark behavior is outside the reporting-pattern question.
- Anthropic related-content links unrelated to evaluation statistics: rejected as topically irrelevant.

## Findings

### 1. Frame the analysis as measurement under uncertainty, not a scoreboard

The public article opens with a concrete ambiguity: whether an observed lead reflects a capability difference or the questions that happened to be selected. It then names the methodological problem before presenting recommendations. The paper similarly criticizes a “highest number is best” mentality and reframes evaluations as experiments that answer hypotheses with limited precision.

**Evidence:** The article's opening asks whether a measured lead is real or due to the benchmark's question selection, then says the recommendations are intended to make results scientifically informative ([Anthropic article, opening](https://www.anthropic.com/research/statistical-approach-to-model-evals)). The paper's §1 introduces fictional raw scores, demonstrates that a superficial ranking can be ambiguous, and states that the framework addresses quantitative questions while remaining agnostic about the qualitative merit of the evaluations ([paper §1](https://ar5iv.labs.arxiv.org/html/2411.00640)).

**Implication:** Start the Skill Issue Analysis page with the evaluative question and the limits of the evidence, then present results. Describe raw scores as observed outcomes from the benchmark design. Reserve “better,” “ahead,” or a rank claim for a defined comparison supported by the available uncertainty analysis.

### 2. Put uncertainty and sample information beside the estimate

Anthropic recommends presenting the number of questions and standard error with each mean, rather than placing uncertainty in distant methodology prose. Its examples put the uncertainty directly beneath or beside the score and add cluster counts when observations are dependent. This makes precision part of the result itself.

**Evidence:** The article's Recommendations #1 and #2 say to report the standard error alongside each score and to adjust it for grouped questions ([Anthropic Recommendations #1–2](https://www.anthropic.com/research/statistical-approach-to-model-evals)). The paper's §2.1 and Table 2 recommend score, question count, and standard error in one table; §2.2 and Table 3 add cluster count when clustered errors are used ([paper §§2.1–2.2](https://ar5iv.labs.arxiv.org/html/2411.00640)).

**Implication:** Every prominent Skill Issue estimate should show, or immediately link to, its observation count and uncertainty measure. If trials share a skill, scenario, fixture, prompt family, or other common source of dependence, the page should identify the independent unit and any clustering rather than allowing the displayed run count to imply more precision than the design supports.

### 3. Compare shared cases through paired differences

When systems are evaluated on the same cases, Anthropic treats the case-level difference as the comparison object. Paired analysis removes shared case difficulty from the noise and supports direct reporting of the mean difference, its standard error, confidence interval, and cross-system correlation. Separate confidence intervals around each score do not convey the same comparison.

**Evidence:** Recommendation #4 calls paired differences a variance-reduction method and recommends pairwise mean differences, standard errors, confidence intervals, and correlations ([Anthropic Recommendation #4](https://www.anthropic.com/research/statistical-approach-to-model-evals)). The paper's §4.2 derives the paired comparison and Table 5 presents baseline, difference, confidence interval, and correlation together ([paper §4.2](https://ar5iv.labs.arxiv.org/html/2411.00640)).

**Implication:** If Skill Issue model/harness combinations ran the same benchmark cases under comparable conditions, lead with paired differences for head-to-head claims. A ranking can remain a compact orientation aid, but the explanatory unit should be “difference versus a named baseline on shared cases,” with its uncertainty. If cases are not shared or conditions differ, label the comparison accordingly and avoid implying paired evidence.

### 4. Distinguish “no detected difference” from equality

Anthropic's comparison framework supports two different statements: a difference may be distinguishable from the modeled noise, or the data may lack enough precision to distinguish it. Its power analysis then asks what effect size the evaluation could reliably detect. A non-significant result therefore cannot, by itself, establish a tie or practical equivalence.

**Evidence:** The paper's §4 describes differences whose intervals do or do not exclude zero; §5 relates sample size, false-positive rate, power, and minimum detectable effect, and says consumers can decide whether a fixed evaluation is capable of measuring the improvement of interest ([paper §§4–5](https://ar5iv.labs.arxiv.org/html/2411.00640)). Recommendation #5 explains that limited question counts yield wide intervals and allow only large differences to register clearly ([Anthropic Recommendation #5](https://www.anthropic.com/research/statistical-approach-to-model-evals)).

**Implication:** Use language such as “the benchmark did not resolve a difference of this size” when uncertainty spans competing interpretations. Use “tie” only when the benchmark has an explicit equivalence rule that the data satisfy. Where the needed inputs exist, report a minimum detectable effect or an equivalent sensitivity statement so readers know which practically meaningful gaps the evaluation could or could not resolve.

### 5. Report sensitivity as design dependence, not an afterthought

The same observed score can support different uncertainty depending on whether observations are independent, clustered, or repeated. Anthropic also separates randomness from question selection and warns that changing sampling temperature to suppress output variance changes the object being measured. This places sensitivity analysis at the level of evaluation design and analysis choices.

**Evidence:** Recommendation #2 says naive errors can substantially understate uncertainty when questions arrive in related groups; Recommendation #3 separates question choice from repeated answer randomness ([Anthropic Recommendations #2–3](https://www.anthropic.com/research/statistical-approach-to-model-evals)). The paper's §§2.2 and 3.1 explain clustering and question-level averaging across resamples; §3.3 warns that changing temperature can alter behavior, bias, or irreducible variance rather than merely improve precision ([paper §§2.2, 3.1, 3.3](https://ar5iv.labs.arxiv.org/html/2411.00640)). The earlier Anthropic challenges article reports that formatting and implementation choices can materially change benchmark outputs ([“The supposedly simple multiple-choice evaluation”](https://www.anthropic.com/news/evaluating-ai-systems)).

**Implication:** The Skill Issue page should identify the result's sensitivity to the choices that actually vary in its design: model/harness route, prompting or skill-installation condition, repetitions, scoring rule, case grouping, and implementation version where applicable. Robustness prose should state what was varied, what stayed fixed, and whether the conclusion persisted; it should not generalize beyond tested variants.

### 6. Separate observations, statistical interpretations, and practical meaning

Anthropic's public explanation moves in three layers: a visible numerical observation, a statistical interpretation about signal versus noise, and a practical consequence for what the evaluator should report or run next. The broader challenges article adds that a numerically valid result can still fail to measure the intended behavior, so statistical precision does not establish construct validity.

**Evidence:** Each recommendation in the article follows a recurring pattern: describe an evaluation-data situation, explain the statistical problem, then state the reporting or design action ([Anthropic article](https://www.anthropic.com/research/statistical-approach-to-model-evals)). The challenges article's BBQ example describes a technically unbiased score that was useless because the model was not answering, explicitly warning against overinterpreting a quantitative result ([“BBQ: Measuring social biases is even harder”](https://www.anthropic.com/news/evaluating-ai-systems)). The paper's §1 explicitly limits itself to quantitative questions about specific results while remaining agnostic about qualitative evaluation merit ([paper §1](https://ar5iv.labs.arxiv.org/html/2411.00640)).

**Implication:** Give each Skill Issue result three clearly signaled sentences or fields: **Observation** (what happened in the recorded trials), **Interpretation** (what the supported analysis says about uncertainty or difference), and **Practical meaning** (what a user may reasonably infer for skill calling in the tested setup). Keep mechanism explanations and broad capability claims out of Observation; label them as interpretations and qualify them to the tested scope.

### 7. Link every chart or table to one nearby prose conclusion

Anthropic's figures and tables are introduced before or at the point where the text needs them, and their captions state the lesson the reader should take from the displayed statistic. Tables become progressively more decision-relevant: raw scores, score-plus-error summaries, cluster-aware summaries, then pairwise differences and intervals.

**Evidence:** The public article alternates recommendation headings, short explanatory paragraphs, a visual, and a caption that states how the distribution changes ([Anthropic Recommendations #1–4](https://www.anthropic.com/research/statistical-approach-to-model-evals)). In the paper, Table 2 links scores to question counts and standard errors; Tables 3–4 expose cluster adjustment; Table 5 turns the same evaluation into pairwise decision evidence ([paper §§2 and 4.2](https://ar5iv.labs.arxiv.org/html/2411.00640)).

**Implication:** Order the Skill Issue page from orientation to decision evidence: benchmark question and scope; observed score summary; uncertainty-aware comparisons; sensitivity or subgroup views; practical interpretation and limits. Give each chart a nearby one-sentence takeaway that names the statistic and its implication. Keep caveats beside the affected graphic rather than collecting all qualifications at the end.

### 8. Use concise language built around concrete consequences

The public article largely avoids formula-first exposition. It defines one term at a time, uses concrete examples of related questions or repeated answers, and closes paragraphs with the consequence of ignoring the issue. The technical paper carries formulas and detailed assumptions, allowing the public article to remain readable while still linking to the full method.

**Evidence:** Recommendation #2 explains dependence through multiple questions about one passage before naming clustered standard errors, then states the risk of false conclusions. Recommendation #5 defines power in plain language before describing its inputs and decisions ([Anthropic Recommendations #2 and #5](https://www.anthropic.com/research/statistical-approach-to-model-evals)). The paper provides the corresponding formulas, assumptions, derivations, and suggested table formats ([paper §§2.2, 4.2, 5 and Appendices A–C](https://ar5iv.labs.arxiv.org/html/2411.00640)).

**Implication:** Use short declarative passages on the Analysis page: name the observed pattern, explain the relevant uncertainty in one step, and state the consequence. Define technical terms at first use and provide methodology detail through an expandable note or linked methods section. Prefer “the interval is wide enough to include either ordering” over unexplained significance jargon.

## Notes

- **Validation performed:** Cross-checked the public article's five recommendations against the underlying paper's assumptions, formulas, reporting tables, and conclusion. Cross-checked the construct-validity and implementation-sensitivity interpretation against Anthropic's earlier first-party evaluation-challenges article.
- **Caveat:** The super-population framing, normal approximations, clustering unit, paired analysis, and power calculations depend on the actual Skill Issue sampling design. This assignment validates reporting principles, not their automatic statistical applicability to unspecified Skill Issue data.
- **Unsupported observation:** No conclusion about the current Skill Issue benchmark's independence, clusters, power, meaningful effect size, or model/harness ordering can be supported from the internet-only sources reviewed here.
- **Useful search terms:** paired benchmark differences; cluster-robust standard error; minimum detectable effect; evaluation construct validity; repeated-sample question-level mean; benchmark implementation sensitivity.

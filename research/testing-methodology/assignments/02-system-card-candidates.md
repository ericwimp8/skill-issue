# System Card Candidates

## Assignment

- **Goal:** Discover authoritative model cards, system cards, and official evaluation disclosures whose complete public reading flow can inform a dense, non-promotional Testing Methodology page for Skill Issue.
- **Scope:** Internet-only review of primary publications from model developers, government evaluation bodies, and established independent evaluation organizations. Candidates are judged on disclosure of purpose and scope, evaluated subjects, datasets or scenarios, procedures, manual work, instrumentation, scoring, controls, uncertainty, limitations, citations, linked evidence, editorial structure, and readable density.
- **Exclusions:** Product launch pages without meaningful evaluation disclosure; secondary summaries; leaderboard-only pages; policy documents without evaluation results; and synthesis or recommendations for Skill Issue beyond candidate selection.

## Sources

1. OpenAI, [GPT-4o System Card](https://openai.com/index/gpt-4o-system-card/), August 8, 2024.
2. OpenAI, [OpenAI o1 System Card](https://openai.com/index/openai-o1-system-card/), December 5, 2024, with later noted updates.
3. OpenAI, [GPT-5 System Card](https://cdn.openai.com/gpt-5-system-card.pdf), August 7, 2025.
4. Anthropic, [Claude 4 System Card](https://www-cdn.anthropic.com/6be99a52cb68eb70eb9572b4cafad13df32ed995.pdf), May 2025.
5. Google DeepMind, [Gemini 2.5 Technical Report](https://storage.googleapis.com/deepmind-media/gemini/gemini_v2_5_report.pdf?pubDate=20250702), July 2025.
6. Model Evaluation & Threat Research, [Details about METR's evaluation of OpenAI GPT-5](https://metr.org/evaluations/gpt-5-report/), August 7, 2025.
7. NIST, [Expanding the AI Evaluation Toolbox with Statistical Models, NIST AI 800-3](https://www.nist.gov/publications/expanding-ai-evaluation-toolbox-statistical-models), February 17, 2026; canonical paper DOI linked by NIST.
8. UK AI Security Institute, [Frontier AI Trends Report](https://www.aisi.gov.uk/frontier-ai-trends-report), December 2025.
9. Anthropic, [Claude 3.7 Sonnet System Card](https://assets.anthropic.com/m/785e231869ea8b3b/original/claude-3-7-sonnet-system-card.pdf), February 2025.
10. Google DeepMind, [Gemini 2.5 Pro Model Card](https://storage.googleapis.com/deepmind-media/Model-Cards/Gemini-2-5-Pro-Model-Card.pdf), updated June 27, 2025.
11. OpenAI, [Operator System Card](https://openai.com/index/operator-system-card/), January 23, 2025.
12. OpenAI, [Sora System Card](https://openai.com/index/sora-system-card/), December 9, 2024.
13. Google DeepMind, [Holistic Safety and Responsibility Evaluations of Advanced AI Models](https://deepmind.google/research/publications/78149/), April 22, 2024; [full publication reading flow](https://ar5iv.labs.arxiv.org/html/2404.14068).
14. Meta, [Llama 3.1 Model Card](https://github.com/meta-llama/llama-models/blob/main/models/llama3_1/MODEL_CARD.md), July 23, 2024.
15. Cohere for AI, [Command R Model Card](https://huggingface.co/CohereLabs/c4ai-command-r-v01), March 2024.

## Findings

### 1. GPT-4o System Card — Deep-dive

This is the strongest compact example of a model-specific disclosure that moves from system purpose to risk discovery, evaluation method, results, mitigations, and residual limitations. Its multimodal focus forces it to explain transformations and instrumentation rather than merely naming benchmarks.

**Evidence:** The card defines the tested system and modalities, identifies external red-teaming phases and participant coverage, and explains how text datasets were converted with text-to-speech tooling for speech-to-speech evaluation. It distinguishes transcript scoring from direct audio evaluation, names auxiliary classifiers, gives task sets and result tables, and explicitly describes threats to validity from text-to-speech reliability, intonation, background noise, cross-talk, and non-textual audio artifacts. Preparedness sections add scenario counts, tool access, attempt budgets, pass rates, risk thresholds, third-party assessments, citations, and an appendix. Coverage is strong for purpose/scope, subjects, datasets/scenarios, procedures, manual work, instrumentation, scoring, controls, limitations, citations, linked evidence, editorial structure, and density; uncertainty is discussed mainly as methodological limitation rather than consistently quantified statistically.

**Implication:** Use as a primary deep dive for the page-level sequence: define the evaluated object, show how source material becomes test inputs, state what is scored, disclose what the instrumentation misses, then report results beside mitigations. Its online table of contents, short risk summaries, and expandable evidence support dense reading without promotional pacing.

### 2. OpenAI o1 System Card — Deep-dive

This card is especially useful for version scope, evaluator construction, and separating related failure dimensions. It is one of the clearest public examples of stating which checkpoints each result actually concerns.

**Evidence:** The scope section names near-final and release checkpoints, assigns different evaluation families to those checkpoints, and warns that production performance can vary with updates, final parameters, system prompts, and other factors. Safety sections name public and internal datasets, explain sampling from WildChat, define `not_unsafe` and `not_overrefuse` autograder metrics, report disallowed-content and jailbreak results, and separate text-only and multimodal evaluations. The card covers chain-of-thought safety, deception monitoring, external red teaming, Preparedness evaluations, citations, detailed appendices, and later-version caveats. Coverage is strong for purpose/scope, subjects, datasets/scenarios, procedures, scoring, controls, limitations, citations, evidence links, structure, and density; manual steps and instrumentation are described unevenly across evaluation families, and statistical uncertainty is less prominent than in GPT-5.

**Implication:** Deep-dive for precise evaluated-version disclosure, metric definitions, and paired safety/helpfulness reporting. The methodology page can borrow its discipline of stating evaluation-to-checkpoint mapping before any score and its compact pattern of dataset, metric, table, interpretation.

### 3. GPT-5 System Card — Deep-dive

This is the strongest OpenAI candidate for uncertainty, elicitation controls, and distinguishing observed capability from capability bounds. It combines a readable main narrative with enough methodological qualification to resist headline-only interpretation.

**Evidence:** The card specifies the GPT-5 system variants and the models emphasized in the report. Its capability-assessment preamble states that evaluations used multiple elicitation methods, including custom post-training, scaffolding, and prompting, and frames results as lower bounds. It documents bootstrap 95% confidence intervals for pass@1 and explains why those intervals can be too narrow on small datasets because they capture attempt-level sampling variance rather than full problem-level variance. It also records maximum-verbosity evaluation conditions, external red-team methods and expert counts, third-party tooling such as PyRIT, benchmarks, graders, thresholds, safeguards, and limitations. Coverage is strong across all requested dimensions, including subjects, procedures, instrumentation, scoring, controls, uncertainty, limitations, citations, linked evidence, and structured appendices. Density is high but the report remains navigable through a stable system overview, evaluation families, tables, and caveats.

**Implication:** Deep-dive for how Skill Issue should state run conditions, elicitation strength, confidence intervals, and what the intervals do and do not measure. It provides a direct model for keeping uncertainty next to the evaluation methodology rather than isolating caveats at the end.

### 4. Claude 4 System Card — Deep-dive

This is the richest candidate for mixed-method evaluation disclosure. It exposes automatic generation, manual conversations, transcript review, human grading, thresholds, third-party work, and candid failure modes in one sustained report.

**Evidence:** The 123-page card defines the two model subjects, modes, and release-level decisions, then covers policy behavior, multi-turn testing, alignment assessments, agentic behavior, biology, cyber, AI research, and third-party assessments. Multi-turn testing combines automated conversation generation with detailed manual conversations by policy experts; thousands of conversations are filtered with policy-specific rubrics and potential violations are human-reviewed. Elsewhere the report discloses at least 20 hours of transcript review, automated classifiers over named transcript counts, manual grading by external experts, question-specific rubrics, thresholds, and low-reliability judgments where human grading is noisy. It also reports data leakage, evaluation-awareness concerns, small samples, evolving pilot methods, and cases where pretraining knowledge defeats the intended construct. Citations and example transcripts make claims inspectable. Coverage is exceptionally strong across the requested dimensions, though the breadth creates a very long reading flow and methods are distributed across domain sections rather than normalized into one schema.

**Implication:** Deep-dive for candid qualitative evidence, human-review disclosure, threshold ownership, and reporting when a test ceases to measure its intended construct. Its best transferable editorial pattern is repeated domain sections that pair method, threshold, score, and interpretation, while preserving excerpts and caveats close to the claim.

### 5. Gemini 2.5 Technical Report — Deep-dive

This is the strongest single-source example of joining capability benchmarks, safety evaluation, governance, external testing, and an evaluation-details appendix while giving readers an explicit navigation path.

**Evidence:** The report includes a quantitative evaluation methodology section that defines pass@1, single-attempt versus multiple-attempt conditions, model identifiers, and comparison settings. Its safety section begins with a numbered guide telling readers where to find process, policies, training, development evaluations, automated red teaming, memorization, assurance evaluations, frontier assessments, and external testing. It distinguishes development from held-out assurance evaluation, reports attack-success rates over 500 held-out scenarios, describes memorization testing over more than 700,000 sampled documents, records 100 shuffled runs for selected multiple-choice tests, and includes detailed benchmark descriptions and grader notes in the appendix. External groups chose independent methodologies, used manual or automated strategies, supplied raw prompts and outputs, and were reviewed by internal subject-matter experts. The report also discloses changing datasets, changed benchmark variants, non-comparability, and selected qualitative judgments. Coverage is strong across all requested dimensions, although some safety results are relative rather than absolute and some external evidence is summarized rather than fully linked.

**Implication:** Deep-dive for navigation design and explicit test-condition definitions. The numbered "how to read this section" device is particularly suitable for a dense methodology page, as is the appendix table that centralizes benchmark purpose, scoring, attempts, and judge details.

### 6. METR GPT-5 Evaluation Report — Deep-dive

This is the best independent-assessor candidate and the most candid example of linking a risk conclusion to assumptions, access conditions, adversarial checks, and reasons the conclusion could fail.

**Evidence:** METR names the evaluated GPT-5 variant, threat models, access timeline, checkpoint changes, reasoning-trace access, data-retention constraint, assurance checklist, report-review process, and time-horizon result. The report explains its estimate in human-task-time terms, performs checks for underestimation and strategic sabotage, documents artificial honeypots and refusal-rate tables with sample counts, and distinguishes observations from assurances supplied by OpenAI. Its limitation section states task saturation, missing human baselines, unrealistic environments, evaluation awareness, rough proxy status, and monitorability weaknesses. It links methodology, task suites, threat-model work, thresholds, and related reports throughout. Coverage is strong for scope, subjects, scenarios, procedures, manual inspection, instrumentation, scoring, controls, uncertainty, limitations, citations, evidence links, editorial structure, and readable density. Some evidence remains assurance-based or NDA-constrained, which the report foregrounds.

**Implication:** Deep-dive for an "evidence chain" pattern: conclusion, measured result, assumptions, attempts to falsify the measurement, access chronology, and limitations. This is especially relevant to a non-promotional page because independence conditions and disclosure constraints are treated as part of methodology rather than acknowledgements.

### 7. NIST AI 800-3 — Deep-dive

This is the strongest cross-cutting source for scoring semantics and statistical uncertainty. It is narrower than a system card but directly addresses a recurring weakness in model disclosures: treating one average as if it answered every evaluation question.

**Evidence:** NIST distinguishes fixed-benchmark accuracy from generalized accuracy over a broader item population, shows that they require different estimators, and tests methods on 22 frontier language models across GPQA-Diamond, BIG-Bench Hard, and Global-MMLU Lite. It compares regression-free estimators with generalized linear mixed models, formalizes assumptions, decomposes variance, estimates item difficulty, and explains how common procedures can yield invalid or misleading confidence intervals. The official publication page supplies authors, report number, DOI, canonical PDF, abstract, and citation metadata. Coverage is strong for evaluation purpose, subjects, datasets, statistical procedure, scoring, controls, uncertainty, limitations, citations, and linked evidence; it has little manual evaluation or system instrumentation because its subject is benchmark analysis itself. Editorially, the distinction between two estimands provides a compact conceptual spine for a dense explanation.

**Implication:** Deep-dive for the scoring and uncertainty section of the methodology page. It supports requiring every metric to state the target quantity, sampling unit, estimator, assumptions, and confidence-interval interpretation rather than presenting an unlabeled average.

### 8. UK AISI Frontier AI Trends Report — Deep-dive

This is the strongest government example of an accessible, multi-domain evaluation report that retains methodological controls, uncertainty, withheld-information boundaries, and interpretive caution.

**Evidence:** The report covers biology, chemistry, cyber, autonomy, safeguards, societal effects, and open-versus-closed model trends using benchmarks, agentic tasks, expert baselines, long-form tasks, human uplift studies, manual checks, and real-world wet-lab validation. It states repeat counts, averaging procedure, standard-error formula, task-level sampling unit, figure-specific exceptions, and guidance against over-interpreting small differences. The appendix distinguishes controlled capability proxies from real-world effectiveness, notes checkpoint and safeguard mismatches, records missing fine-tuning or maximal scaffolding, and explains that some high-risk methods are withheld. Figures and captions define tasks, thresholds, graders, repeats, baselines, and linked methodologies. Coverage is strong across all requested dimensions; model identities are intentionally aggregated, and some sensitive datasets and procedures are withheld.

**Implication:** Deep-dive for balancing accessible narrative with technical disclosure. Its repeatable figure-caption pattern and consolidated limitations/uncertainty appendix show how a public page can remain readable while naming sampling units, repeats, baselines, access differences, and disclosure boundaries.

### 9. Claude 3.7 Sonnet System Card — Skim-only

This focused predecessor is easier to read than the Claude 4 card and retains useful treatment of extended thinking, prompt injection, reward hacking, child safety, and Responsible Scaling Policy evaluations.

**Evidence:** The 43-page report states training and cutoff scope, standard and extended-thinking modes, release considerations, risk families, evaluation results, appendices, citations, and ongoing-safety commitments. It contains concrete task and attempt counts and discusses model-specific failure modes. Coverage is strong for purpose, subjects, scenarios, scoring, limitations, citations, structure, and density, but Claude 4 substantially expands manual review, alignment auditing, threshold detail, third-party assessments, and explicit methodological caveats.

**Implication:** Skim for concise information architecture and focused reasoning-mode disclosure. Select Claude 4 for detailed evidence unless the synthesis needs an example of a shorter card that still covers multiple safety domains.

### 10. Gemini 2.5 Pro Model Card — Skim-only

This model card is a strong compact companion to the Gemini technical report, especially for lifecycle governance, relative safety metrics, known limitations, and explaining why results are not comparable across card versions.

**Evidence:** The card defines the model, intended uses, known limitations, and an evaluation approach spanning internal teams, specialist human red teaming, automated red teaming, independent assurance evaluators, governance review, and Frontier Safety Framework testing. It distinguishes automated results from human evaluation and red teaming, reports manual review of flagged regressions, explains metric direction, states that improved evaluations prevent direct comparison with prior cards, and records unresolved or incomplete assessment areas. Coverage is strong for purpose, evaluation layers, manual steps, scoring interpretation, controls, limitations, and editorial density; datasets, instrument configuration, absolute scores, sample sizes, and uncertainty are thinner than in the technical report.

**Implication:** Skim for a compact disclosure template and the phrasing of non-comparability. The technical report remains the deeper source for procedures, benchmark details, grader choices, attempts, and external testing evidence.

### 11. Operator System Card — Skim-only

This is a useful product-system disclosure because its subject is an agent acting through a computer interface rather than a standalone language model. It makes system-level controls and refusal behavior more concrete.

**Evidence:** The card covers the Computer-Using Agent, user confirmations, takeover modes, prompt injection, harmful tasks, privacy, model limitations, monitoring, red teaming, and staged deployment. Appendices provide standard and challenging refusal results by policy category against a GPT-4o baseline. Coverage is strong for system purpose, risk scenarios, controls, mitigations, manual testing context, and compact editorial structure. It is weaker on dataset provenance, detailed procedure, instrumentation parameters, sample sizes, grading method, uncertainty, and linked raw evidence; several result tables are presented with limited methodological context.

**Implication:** Skim for how to describe system boundaries and human-in-the-loop controls. Do not use it as the primary scoring-method exemplar because the card lacks enough detail to reconstruct many reported evaluations.

### 12. Sora System Card — Skim-only

This is a concise multimodal example that clearly names how evaluation prompts were sourced and where classifiers operate in the system.

**Evidence:** The evaluation section covers nudity, deceptive election content, self-harm, and violence; explains that the framework combines input prompts with input and output classifiers; and names three prompt sources: early-alpha usage, adversarial red-team examples, and synthetic GPT-4 data. It explains why each source contributes different coverage and links evaluation to moderation thresholds and Preparedness categories. Coverage is moderate for purpose, subjects, scenario sources, instrumentation, mitigations, and readable density, but weak for dataset sizes, sampling, detailed procedures, manual review, scoring definitions, statistical uncertainty, result tables, controls, citations, and linked raw evidence.

**Implication:** Skim for a compact source-provenance paragraph and classifier-placement diagram logic. It is too abbreviated to anchor a complete Testing Methodology page.

### 13. Holistic Safety and Responsibility Evaluations — Skim-only

This publication is a strong conceptual bridge between model cards and evaluation program design. It explains why development, assurance, red-team, human-centric, and system-level methods answer different questions.

**Evidence:** Google DeepMind defines evaluation as empirical assessment of components, capabilities, behavior, and impact; separates development evaluations from arm's-length assurance evaluations; discusses held-out prompts, contamination, Goodhart's law, internal and external validity, culturally specific multilingual evaluation, automated classifiers, human raters, system-level methods, incident monitoring, and lifecycle timing. It is heavily cited and organized around foresight, design, ecosystem, and lessons learned. Coverage is strong for purpose, method classes, controls, limitations, citations, evidence links, and editorial structure, but it is not a model-specific disclosure and therefore provides few concrete sample counts, score tables, instrumentation settings, or complete result chains.

**Implication:** Skim as framing support for why a methodology page needs multiple evidence classes and clear evaluation purpose. Use model-specific deep dives for the concrete procedures and results.

### 14. Meta Llama 3.1 Model Card — Reject

This is a competent deployment-oriented model card with unusually good artifact links, but its complete reading flow is dominated by model facts and benchmark tables rather than evaluation-method explanation.

**Evidence:** The card states model architecture, release status, intended and out-of-scope uses, training data scale, hardware, energy and emissions, capability and safety benchmark tables, shot counts, metrics, and multilingual scope. It links an internal evaluation library and raw generated evaluation data, which is valuable for auditability. Coverage is strong for model purpose, subjects, benchmark names, selected settings, citations, and linked evidence; it is weak on end-to-end procedures, manual steps, instrumentation, sampling controls, grader validation, uncertainty, and limitation analysis around score interpretation. The dense tables are readable but provide little narrative connection from question to method to evidence to implication.

**Implication:** Reject as a structural reference for the Testing Methodology page. Retain only as evidence that raw-data and evaluation-code links should sit next to benchmark tables; stronger candidates explain what those artifacts mean.

### 15. Cohere Command R Model Card — Reject

This card is useful operational documentation but does not disclose enough evaluation methodology for the target page.

**Evidence:** The card identifies the 35B model, intended research and retrieval/tool-use use cases, multilingual and code capabilities, prompt-template dependencies, license, limitations, contact path, and selected performance tables. It warns that deviating from the prescribed multi-step tool-use template may reduce performance and recommends decoding settings for code-related use. Coverage is moderate for purpose, subjects, intended use, operational controls, and limitations, but weak for dataset provenance, scenario construction, manual evaluation, instrumentation, scoring procedures, uncertainty, evaluator controls, citations, and linked raw evidence. The reading flow is practical and concise but closer to a usage guide than a testing disclosure.

**Implication:** Reject because it cannot support a dense account of how evaluation claims were produced. Its only transferable point is to state prompt-template and decoding preconditions when results depend on them.

## Notes

- **Classification summary:** Deep-dive — GPT-4o, o1, GPT-5, Claude 4, Gemini 2.5 Technical Report, METR GPT-5, NIST AI 800-3, UK AISI Frontier AI Trends. Skim-only — Claude 3.7 Sonnet, Gemini 2.5 Pro Model Card, Operator, Sora, Holistic Safety and Responsibility Evaluations. Reject — Llama 3.1 Model Card, Command R Model Card.
- **Cross-source convergence:** The strongest candidates consistently identify the evaluated artifact/version, separate evaluation purposes or evidence classes, state run conditions and graders, disclose human work, pair scores with thresholds or baselines, and place limitations close to interpretation. The OpenAI, Anthropic, Google DeepMind, METR, NIST, and AISI sources independently reinforce these patterns.
- **Caveat:** Some reports intentionally withhold high-risk datasets or procedure details, and several developer cards rely on internal datasets, undisclosed system prompts, or summarized third-party work. They can still model transparent disclosure of the boundary: what was tested, what remains inaccessible, and how that limits interpretation.
- **Relevant search terms:** `system card evaluation methodology`, `model card assurance evaluation`, `frontier model third-party assessment`, `LLM benchmark uncertainty`, `manual transcript review`, `evaluation checkpoint scope`, `evaluation lower bound scaffolding`, `human uplift study`, `held-out safety evaluation`.

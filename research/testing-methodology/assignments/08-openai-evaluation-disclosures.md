# OpenAI Evaluation Disclosures

## Assignment

**Goal:** Deep-dive official OpenAI evaluation disclosures that are especially relevant to software-engineering agents and dense public methodology communication, then derive evidence-backed presentation and methodology patterns that Skill Issue can adapt without copying OpenAI wording or visual identity.

**Scope:** Official OpenAI research pages, official OpenAI system cards, and primary artifacts directly linked from those pages. The core connected flow is SWE-bench Verified's release, its later retirement, the subsequent SWE-Bench Pro audit, SWE-Lancer's publication/paper/repository, and the reuse of those evaluations in a model system card. The review covers evaluation subjects, task units, trajectories, manual work, instrumentation, scoring, environments, reproducibility, retained evidence, uncertainty, citations, artifact links, and information architecture.

**Exclusions:** General OpenAI product marketing without substantive evaluation methodology; third-party summaries; independent benchmark claims that were not incorporated into an OpenAI disclosure; pixel-level imitation, OpenAI brand styling, and recommendations that assume Skill Issue has OpenAI-scale review staffing or private datasets.

## Sources

- OpenAI, **Introducing SWE-bench Verified** (published August 13, 2024; updated February 24, 2025): https://openai.com/index/introducing-swe-bench-verified/. Inspected the full page, including background, task mechanics, annotation campaign, dataset construction, interactive examples, performance stratification, limitations, data downloads, footnotes, and outbound artifacts.
- OpenAI, **SWE-bench Annotation Instructions**: https://cdn.openai.com/introducing-swe-bench-verified/swe-b-annotation-instructions.pdf. Inspected the five-page primary rubric, including annotator information boundaries, severity labels, required evidence, difficulty estimates, other-issue flags, and confidence capture.
- OpenAI, **Why SWE-bench Verified no longer measures frontier coding capabilities** (February 23, 2026): https://openai.com/index/why-we-no-longer-evaluate-swe-bench-verified/. Inspected the full page and its worked examples, audit design, contamination red-team procedure, model-specific transcripts, discussion, and explicit recommendation to stop reporting the benchmark.
- OpenAI, **Separating signal from noise in coding evaluations** (July 8, 2026): https://openai.com/index/separating-signal-from-noise-coding-evaluations/. Inspected methodology, workflow diagram, agent-assisted audit, human campaign, taxonomy comparison, worked example, discussion, and retraction of the earlier SWE-Bench Pro recommendation.
- OpenAI, **Introducing the SWE-Lancer benchmark** (February 18, 2025; update dated July 28, 2025): https://openai.com/index/swe-lancer/. Inspected the concise publication page and its paper/repository handoff.
- Samuel Miserendino, Michele Wang, Tejal Patwardhan, and Johannes Heidecke, **SWE-Lancer: Can Frontier LLMs Earn $1 Million from Real-World Freelance Software Engineering?**, arXiv:2502.12115v4 (May 29, 2025): https://arxiv.org/abs/2502.12115 and https://arxiv.org/pdf/2502.12115. Inspected the 39-page paper, including task flows, construction, agent controls, scoring, ablations, qualitative trajectories, curation rubrics, prompts, holdout design, limitations, and appendices.
- OpenAI, **Frontier Evals / SWE-Lancer repository**: https://github.com/openai/frontier-evals and https://github.com/openai/frontier-evals/blob/main/project/swelancer/README.md. Inspected suite layout, the current 198-task offline-adjusted release note, Docker setup, run commands, solver configuration, validity conditions, retries/concurrency, and run-group/task logging.
- OpenAI, **GPT-5 System Card** (August 13, 2025): https://cdn.openai.com/gpt-5-system-card.pdf. Inspected the contents hierarchy and AI self-improvement sections covering SWE-bench Verified, internal OpenAI PR tasks, SWE-Lancer, PaperBench, OPQA, and METR's external assessment, plus methodological notes on verbosity, attempts, variance, comparisons, and confidence intervals.
- OpenAI, **GPT-4o System Card** (August 8, 2024): https://openai.com/index/gpt-4o-system-card/. Inspected the model-autonomy section, especially the contrast between automated grader success and manual inspection of whether the underlying task was actually completed.

## Findings

### 1. The strongest disclosure is a connected stack, not one overloaded page

OpenAI separates communication layers by reader need. A short publication page states the benchmark's purpose and headline construction, the paper carries methods and appendices, the repository supplies executable details, the system card places the benchmark inside a broader model-risk argument, and later audit pages revise or withdraw prior conclusions. Each artifact has a clear role and links to the next level of evidence.

**Evidence (observation):** The SWE-Lancer landing page uses a question-form title, a two-paragraph summary, and direct paper/repository calls to action. The paper expands this into exact task counts, two task types, construction, environmental controls, ablations, limitations, full prompt templates, and qualitative trajectories. The repository then documents package installation, image building, internet controls, exact command flags, solver choices, and log locations. The GPT-5 System Card reuses a condensed subset of the method in its AI self-improvement section and links the July 2025 dataset update. The SWE-bench sequence continues beyond launch: the 2026 retirement page revises the original interpretation, and the July 2026 audit retracts the replacement recommendation.

**Implication (inference/recommendation):** Skill Issue should make one public methodology page the navigational owner of the evaluation story, then link to durable artifacts such as schemas, retained results, commands, and detailed reports. A single page should remain understandable by itself, while deeper reproducibility belongs in linked artifacts rather than an undifferentiated wall of prose.

### 2. Titles and opening summaries bound the claim before adding detail

The pages lead with a narrow thesis that tells the reader what changed and what the evidence can support. The opening does not require the reader to infer whether a page launches, audits, or retires a benchmark.

**Evidence (observation):** “Introducing SWE-bench Verified” immediately describes a human-validated subset intended to improve reliability. “Why SWE-bench Verified no longer measures frontier coding capabilities” pairs its title with a one-sentence contamination diagnosis and a current recommendation. “Separating signal from noise in coding evaluations” states an approximate broken-task estimate in the subtitle. “Introducing the SWE-Lancer benchmark” puts the economic question in the subtitle, then gives the benchmark size, task classes, grading approach, and release artifacts in the first two paragraphs. These claims remain narrower than “model quality” in general: they refer to particular benchmarks, software-engineering task distributions, or evaluation validity.

**Implication (inference/recommendation):** Skill Issue's methodology page should open with three explicit items: what the system evaluates, what evidence unit supports the result, and what conclusion the page does and does not justify. For an audit or update, the title should identify the changed judgment rather than masking it under a generic release-note heading.

### 3. Dense pages stay navigable through a repeated evidence hierarchy

The long disclosures use predictable layers: context, method, result, worked evidence, and limitations. Tables and diagrams summarize relationships; prose explains how to interpret them; concrete task fragments make abstract failure modes inspectable.

**Evidence (observation):** The SWE-bench Verified page exposes a section-level contents list, then moves from benchmark background to adaptation, annotation approach, construction, results, stratified performance, limitations, and downloads. It places raw annotation commentary beside issue text and test snippets. The 2026 retirement page uses top-level sections for background, test defects, contamination, provider-specific examples, and discussion; inside the examples, lower-level headings identify the problem statement, elicitation prompt, model response, and gold patch. The July 2026 audit places a workflow diagram directly after the methodology statement, follows it with separate agent-review and human-review subsections, then uses a character-level prompt/test mismatch as the worked example. The GPT-5 System Card uses a two-page table of contents, numbered subsections, summary tables, figures, notes, and a conventional references section.

**Implication (inference/recommendation):** A Skill Issue page can support high information density if each repeated evaluation section uses the same reader contract: **Purpose → Subject and task → Run setup → Scoring → Findings → Limits → Artifacts**. Worked examples should sit immediately after the category they demonstrate. Tables should summarize; adjacent prose should state the interpretation and caveat.

### 4. Evaluation subjects, task units, and information boundaries are explicit

OpenAI's disclosures specify what is being evaluated at the level of an agent rollout, repository snapshot, task prompt, hidden grader, and allowed toolset. This prevents a benchmark score from being mistaken for a property of the base model alone.

**Evidence (observation):** SWE-bench Verified gives an agent an issue description and codebase, withholds tests, and grades a patch against both fail-to-pass and pass-to-pass tests. SWE-Lancer's individual-contributor tasks provide the issue text, a pre-fix code snapshot, and a fix objective; hidden Playwright end-to-end tests score the submitted patch. Its manager tasks provide an issue, multiple historical proposals, a pre-fix snapshot, and a selection objective, then compare the choice with the original manager decision. The GPT-5 System Card's OpenAI PR evaluation checks out a pre-PR internal branch, gives a human-written change prompt, lets an agent use command-line tools and Python, and applies hidden human-written tests. The system card separately labels the scaffold used for SWE-bench Verified as bash plus `apply_patch`.

**Implication (inference/recommendation):** Skill Issue should identify the evaluated system as a complete configuration: model/harness, prompt and skill state, repository/workspace snapshot, tools, allowed information, time/turn budget, and grader. Summary metrics should never be presented as model-only facts when the harness or skill is part of the causal setup.

### 5. Manual work is disclosed as methodology, not hidden as cleanup

Human review is used for task admission, annotator qualification, test validation, disagreement resolution, and final judgments. The disclosures state who reviewed, how often, what information they saw, and how labels were combined.

**Evidence (observation):** SWE-bench Verified reports 93 Python-experienced developers, 1,699 reviewed samples, a 50-sample onboarding set, three independent labels per sample, and a conservative highest-severity ensemble. Its public rubric constrains annotators to the issue text available to the model, requires code-specific explanations, separates prompt adequacy from test adequacy, records a difficulty bucket, supports an “other issues” flag, and captures confidence. SWE-Lancer reports 100 professional engineers in construction, three reviews for individual-contributor tasks, two for manager tasks, ten reviewers for individual-contributor tasks above $5,000, triple verification of end-to-end tests, and 99% agreement in an additional manager-task validation campaign. The 2026 retirement audit used at least six engineers per audited case plus re-verification when an issue was flagged. The July 2026 audit used five engineers per flagged task and escalated disagreements or low-confidence cases.

**Implication (inference/recommendation):** Skill Issue should publish the human-review protocol as part of the evaluation definition: reviewer qualifications, onboarding/calibration, evidence available, independent review count, label aggregation, escalation, and confidence capture. If staffing is smaller, it should disclose that smaller design rather than mimic the appearance of OpenAI-scale certainty.

### 6. Evaluation quality is checked in both directions

The audit pages distinguish false negatives from false positives. A valid evaluation must reject incomplete shortcuts and accept materially correct alternative implementations.

**Evidence (observation):** SWE-bench Verified initially focused on underspecified issue text and tests that reject reasonable solutions. The 2026 retirement audit names overly narrow tests that demand an implementation detail and overly wide tests that enforce unstated behavior. The July 2026 SWE-Bench Pro audit broadens the taxonomy to overly strict tests, underspecified prompts, low-coverage tests, and misleading prompts. Its methodology statement explicitly requires failures to represent model limitations and successes to represent complete valid solutions. The article's worked example shows a one-character mismatch between the prompt's exact formatting example and the hidden test expectation. The GPT-4o System Card adds a complementary trajectory-level case: some runs passed an automated grader while manual review found that the underlying requirement had not actually been satisfied.

**Implication (inference/recommendation):** Skill Issue's evaluator QA should test both acceptance and rejection behavior. For each criterion, validation should ask whether a different correct output can pass and whether a shortcut or incomplete artifact can also pass. A small set of gold failures, adversarial shortcuts, and alternative-correct solutions is more informative than only proving the canonical output passes.

### 7. Trajectories and intermediate artifacts are treated as evidence

The disclosures do not reduce agent evaluation to a final patch alone. They use attempts, traces, tool outputs, screenshots, model responses, and qualitative failure analysis to interpret the score and find grader defects.

**Evidence (observation):** SWE-Lancer's optional user tool runs a Playwright user flow and writes a text trajectory plus screenshots into the working directory. The agent can call it repeatedly; the full prompt explains the trace file, frame snapshots, logs, ordering, and expected wait time. The paper includes a seven-step worked trajectory in which the agent parses logs and snapshots, changes UI logic and translations, and reruns the tool. The July 2026 audit's initial filter examines task instructions, model attempts, grading tests, metadata, and failure traces. Investigator agents receive the repository and environment and can run tests, inspect nearby conventions, and compare common failed attempts. The SWE-bench contamination page retains the red-team prompt, prefill, response, task identifier, and gold-patch comparison for manually confirmed strong cases. The SWE-Lancer repository stores logs under unique run-group IDs with per-task subdirectories.

**Implication (inference/recommendation):** Skill Issue should retain a run manifest, final artifact, grader results, and enough trajectory evidence to explain why the run passed or failed. The public page can show a small representative trace, while downloadable or repository-local artifacts preserve the rest. This is especially important when a mechanical score can disagree with semantic task completion.

### 8. Environment controls are concrete and tied to validity

The most reproducible disclosures state compute, containerization, network access, repository state, tool-call/time limits, attempts, temperature, and validity conditions. Repository instructions expose the remaining operational details.

**Evidence (observation):** SWE-Lancer reports Azure `Standard_D2as_v4` VMs with 2 vCPUs and 8 GB RAM, prebuilt Docker images, no multimodal input, up to 100 tool calls, a three-hour maximum, temperature 1.0, and one rollout except pass@k experiments. The experiment removes the GitHub remote and future commits and disables internet to reduce direct solution retrieval. The current repository notes that only offline runs are considered valid, recommends an ephemeral Linux VM because network blocking uses `iptables`, documents monolith versus per-task images, pins a published image tag in examples, and provides a dummy solver that can apply the gold solution to prove the evaluation path works. It also records retries and concurrency in the run command. The July 2025 release reduced the public individual-contributor set from 237 to 198 tasks after adjusting and verifying offline execution and dropping 39 tasks.

**Implication (inference/recommendation):** Skill Issue should publish a machine-readable run manifest and a human summary covering exact executable/harness version, model identifier, reasoning setting, workspace commit or snapshot, network state, sandbox/tool permissions, attempt policy, timeout, concurrency, environment image, and grader version. It should also define which deviations invalidate comparison, rather than treating every locally successful run as equivalent.

### 9. Scoring is accompanied by attempt semantics, denominators, and variance warnings

Scores are described in terms of the task decision and the sampling procedure. Notes make otherwise easy-to-miss comparability differences visible.

**Evidence (observation):** SWE-bench Verified requires both fail-to-pass and pass-to-pass tests and reports percent resolved. Its launch page footnote states that scaffold comparisons used a single seed with closest-documented or default hyperparameters, so they may differ from leaderboard results. SWE-Lancer reports task resolution, dollars earned, and earn rate; the paper labels pass@1 as one sample and warns of significant run variance, separately studies pass@k, and averages its cost analysis over five runs. The GPT-5 System Card says its SWE-bench pass@1 estimate averages four tries per instance, uses a fixed internally validated 477-task subset, and withholds tests from the model. A nearby note explains that system-card and launch-blog scores used different verbosity settings and that this can change performance. For SWE-Lancer, the card again warns that one-attempt pass@1 may vary substantially between runs. The external METR subsection reports a point estimate and a 95% confidence interval rather than only one duration.

**Implication (inference/recommendation):** Skill Issue should put the denominator and sampling semantics beside every metric: tasks, trials per task, seed/run count, pass definition, exclusions, and uncertainty. Any non-comparable configuration change should be a visible note adjacent to the result, not buried in a distant FAQ.

### 10. Reproducibility is artifact-backed but honestly partial

OpenAI combines downloadable data, public rubrics, papers, repositories, Docker assets, exact prompts, and commands, while still distinguishing public splits from private holdouts and public harnesses from internal infrastructure.

**Evidence (observation):** The SWE-bench Verified page links the 500-sample dataset, full annotation results, complete rubric, and a containerized harness. SWE-Lancer links a public Diamond split, paper, repository, Docker setup, and full prompt templates; the paper keeps a $499,200 private holdout to mitigate contamination and browsing. The repository includes a lockfile-oriented setup, dataset/code, build scripts, solver entry points, run configuration, and run-group logging. The GPT-5 System Card identifies internal OpenAI PR and OPQA evaluations but describes their task flow without implying that the private data can be reproduced externally. It also distinguishes the fixed 477-task internal SWE-bench subset from the original public 500-task set.

**Implication (inference/recommendation):** Skill Issue should define reproducibility tiers. A public tier can expose the exact harness, schema, sample tasks, and retained outputs; a controlled tier can preserve hidden graders or holdouts; an internal tier can disclose task construction and scoring without publishing sensitive data. The page should say which tier each result supports.

### 11. Corrections and reversals are part of the evidence record

The most important communication pattern is willingness to supersede prior conclusions while preserving the earlier page as historical evidence. Updates explain what changed operationally and how that affects validity.

**Evidence (observation):** The original SWE-bench Verified page states that it supersedes earlier SWE-bench sets and includes an explicit last-updated date. The 2026 retirement page links back to that work, describes the audit that undermined it, states that OpenAI stopped reporting the score, and recommends a replacement. The July 2026 audit links to the retirement page, applies a new QA pipeline to the proposed replacement, estimates roughly 30% broken tasks, and expressly retracts the earlier SWE-Bench Pro recommendation. SWE-Lancer's landing page and system-card section call out the July 2025 update that removed internet dependence; the repository identifies the changed public task count and the 39 dropped tasks.

**Implication (inference/recommendation):** Skill Issue should preserve a dated methodology-change log and mark superseded results at the point where readers encounter them. A correction should name the affected runs, the reason, the replacement method if any, and whether historical numbers remain interpretable. Silent replacement would erase useful evidence about evaluation maturity.

### 12. Limitations are specific to the inference, not generic disclaimers

Limitations explain what population, modality, environment, or causal claim the evaluation does not cover. They also identify contamination and scaffold dependence as threats to interpretation.

**Evidence (observation):** SWE-bench Verified says static public-repository data is contamination-prone, covers only a narrow portion of model autonomy, and is sensitive to external agent scaffolds. SWE-Lancer says its source is one repository and one freelance platform, underrepresents infrastructure work, contains more self-contained work than full-time engineering, excludes clarification, is text-only, and may be contaminated. It cautions against extrapolating beyond freelance work. The GPT-5 System Card repeatedly qualifies comparisons: live predecessor values may differ from launch values, some regressions are statistically insignificant or natural noise, system-card verbosity differs from launch settings, and one-attempt results can vary. The July 2026 coding-evaluation audit compares agent and human labels rather than presenting them as interchangeable; it reports 74% category overlap and observes that humans more often assign multiple failure labels.

**Implication (inference/recommendation):** Skill Issue should attach each limitation to the result it constrains: repository representativeness, harness/model entanglement, hidden-grader validity, sample size, contamination, run variance, platform dependence, and reviewer disagreement. General caveat prose at the bottom of the page is weaker than local statements that bound each claim.

### 13. High-fit presentation patterns for Skill Issue

The most transferable patterns improve inspectability without requiring OpenAI's branding, dataset scale, or staffing.

**Evidence (observation):** Across the inspected disclosures, the recurring high-value elements are a thesis-led title/subtitle, a visible contents hierarchy, compact task-flow diagrams, clearly named methodology subsections, summary tables with adjacent notes, one or more full worked examples, explicit reviewer procedures, exact environment and attempt settings, direct artifact links, limitations near results, and dated corrections.

**Implication (inference/recommendation):** A Skill Issue public methodology page should prioritize:

1. A one-screen summary naming purpose, evaluated system, task unit, pass definition, dataset/run counts, and the latest validity status.
2. A task-flow diagram from input bundle through agent trajectory, artifact capture, grading, review, and published result.
3. A compact run-configuration table and a reviewer-protocol table.
4. A result table whose notes expose attempts, exclusions, uncertainty, and comparability conditions.
5. A worked evaluation example that shows the prompt, allowed evidence, key trajectory events, submitted artifact, grader evidence, and final semantic judgment.
6. Links to machine-readable run manifests, schemas, representative retained outputs, commands, and the methodology change log.
7. A limitations section organized by the claims each limitation constrains.

### 14. Conditional and lower-fit patterns

Several OpenAI patterns are valuable only when the underlying evidence or audience justifies their cost.

**Evidence (observation):** SWE-Lancer uses a 39-page paper, multiple appendices, a private holdout, professional-engineer test creation, an open-source execution stack, and large-scale ablations. The GPT-5 System Card uses 60 pages, dozens of evaluation sections, external laboratories, large tables, confidence intervals, and dense references. The coding-audit pages publish extensive model transcripts and source-code snippets because those concrete records are necessary to establish test defects or contamination.

**Implication (inference/recommendation):** Apply these patterns conditionally:

- **Separate paper or long technical report:** useful when Skill Issue introduces a new benchmark or makes a research claim; excessive for routine release validation.
- **Public runnable benchmark repository:** high value when external researchers are expected to reproduce or extend the evaluation; lower fit when the harness depends on private model access or disposable local workspaces.
- **Private holdout:** justified for contamination-resistant longitudinal measurement; lower fit for transparent skill-behavior examples where inspectability is the primary goal.
- **Full transcripts:** useful for disputed semantic judgments, grader failures, or contamination evidence; otherwise publish representative excerpts and retain the complete artifact separately to avoid overwhelming the main page.
- **Interactive charts or expandable examples:** useful for comparing many models, categories, or samples; static tables and linked artifacts are more robust when the dataset is small.
- **Economic-value metric:** compelling when tasks have real observed prices; should not be simulated from arbitrary weights merely to copy SWE-Lancer's framing.
- **External expert assessment:** useful for high-stakes or independence-sensitive claims; ordinary release claims can instead use clearly separated internal review and reproducible public evidence.

### 15. Claims remain bounded when observation and interpretation are visibly separated

The disclosures generally let readers distinguish the observed measurement from the authors' interpretation, especially when discussing benchmark validity or broader implications.

**Evidence (observation):** SWE-bench Verified first reports annotation rates and stratified performance, then says those results are consistent with removing impossible tasks rather than merely making the set easier. The retirement page reports an audit of a selected 138-task subset and carefully says “at least” 59.4% of that audited subset had material issues before drawing a benchmark-level recommendation. The July 2026 audit reports separate pipeline and human estimates (27.4% and 34.1%) before summarizing the overall issue rate as approximate. SWE-Lancer reports exact task outcomes and actual payouts, then cautiously frames them as a starting point for economic-impact research rather than direct proof of labor-market displacement.

**Implication (inference/recommendation):** Skill Issue should visually or linguistically distinguish **measurement**, **validated finding**, and **interpretation**. Approximate extrapolations should say how they were derived; selected-subset audits should keep the selection rule visible; broader product or ecosystem claims should be labeled as inference rather than presented as direct evaluation output.

## Notes

- The OpenAI developer-docs MCP tools were not available in this session. Installing them would have violated the assignment's requirement to make exactly one filesystem edit, so research used official OpenAI web pages and the primary artifacts they link.
- The web extractor exposed heading order, captions, tables, code blocks, notes, and artifact links, but interactive chart bodies sometimes appeared as loading placeholders. Exact animation, responsive behavior, color, and collapsed/expanded default states were not validated; recommendations above concern information architecture rather than pixel styling.
- The recommendation chain changed over time: the February 2026 page recommended SWE-Bench Pro, while the July 2026 page explicitly retracts that recommendation after a new audit. Any downstream synthesis should treat the July 2026 page as the current controlling statement.
- The SWE-Lancer publication page describes more than 1,400 tasks and the paper specifies 1,488 total tasks. The current public repository is narrower: as of July 17, 2025, it reports 198 offline-adjusted individual-contributor tasks after dropping 39 of the original 237 public problems. Those figures describe different release layers and should not be collapsed.
- The GPT-5 System Card's SWE-bench result uses a fixed 477-task subset validated on OpenAI infrastructure rather than the public 500-task Verified set. The card also states that its verbosity setting differs from the launch-blog run, making direct score comparison conditional.

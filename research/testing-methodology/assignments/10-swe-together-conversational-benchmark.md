# SWE-Together Conversational Benchmark

## Assignment

**Goal.** Deep-dive the complete official SWE-Together reader journey, benchmark methodology, public artifacts, controls, results, validation evidence, limitations, and presentation patterns; then derive source-backed recommendations for presenting and evaluating Skill Issue's three fixed multi-turn skill-calling scenarios.

**Scope.** Internet-only inspection of the official project website, paper, official GitHub repository, canonical run plan, representative task artifacts, evaluation source, and official Hugging Face dataset. The presentation audit covers naming, hierarchy, prose density, charts, tables, notes, task/scenario browsing, expandable details, trajectory browsing, citations, and GitHub linkage. Methodology coverage includes session reconstruction, reactive user simulation, corrective turns, correctness judging, pass@1, pass² stability, cost/run metadata, reference-patch results, reproducibility, controls, and uncertainty. Sources were inspected on 2026-07-21.

**Exclusions.** No local Skill Issue source or scenarios were inspected because this assignment is internet-only. No SWE-Together run was executed, no container was pulled, and no claim of end-to-end reproducibility was inferred from the presence of scripts. Recommendations therefore stay at the presentation and evaluation-contract level that applies to three already-fixed Skill Issue scenarios. No wording, styling, naming, or visual identity is proposed for copying.

## Sources

- [Official SWE-Together overview](https://togetherbench.com/) — hero, leaderboard, model/provider/date metadata, contact, citation, and persistent navigation.
- [Official Design page](https://togetherbench.com/design.html) — construction funnel, benchmark principles, method diagrams, results, suite composition, and FAQ.
- [Official Task browser](https://togetherbench.com/tasks.html) — search, filters, sorting, task cards, and 109-task browsing experience.
- [Representative official task detail](https://togetherbench.com/task.html?id=agent-swarm-task-4a881b) — task summary, model table, interaction loop, weighted goals, environment, verifiers, original session, and source metadata.
- [Official Trial comparison](https://togetherbench.com/three_panel.html) — original-session versus two-replay trajectory browser, editorial comparison modes, selectors, scores, turns, and corrections.
- [Official paper PDF, arXiv:2606.29957v1](https://arxiv.org/pdf/2606.29957) and [arXiv record](https://arxiv.org/abs/2606.29957) — construction, metrics, experiments, simulator study, limitations, and publication date.
- [Official GitHub repository](https://github.com/Togetherbench/SWE-Together) — code and task layout, quickstart, supported agents, environment keys, and repository maturity signals.
- [Canonical 109-task run plan](https://github.com/Togetherbench/SWE-Together/blob/main/canonical_full109.json) and [launcher](https://github.com/Togetherbench/SWE-Together/blob/main/launch.py) — seven model routes, high reasoning, 4,800-second timeouts, 20 workers, two replicates, dry-run behavior, and separate run/judge stages.
- [Representative task artifact directory](https://github.com/Togetherbench/SWE-Together/tree/main/tasks/agent-swarm-task-4a881b) — instruction, oracle intents/session, simulator prompt, reference patch, canonical goals, environment, tests, and task metadata.
- [User simulator implementation](https://github.com/Togetherbench/SWE-Together/blob/main/src/user_agent/user_agent.py) — action schema, silence default, history, structured tool decisions, no-op guards, and fallbacks.
- [Frozen-rubric generation](https://raw.githubusercontent.com/Togetherbench/SWE-Together/main/eval/correctness/generate_task_goals.py) — phase separation, reference/user-intent inputs, E2B judging, and rubric schema validation.
- [Intent Coverage implementation](https://raw.githubusercontent.com/Togetherbench/SWE-Together/main/eval/user_behavior/coverage_one.py) — Gemini 3.1 Pro default, temperature zero, match-table normalization, deterministic score calculation, and the 70/30 recall/precision weighting.
- [User Correction implementation](https://github.com/Togetherbench/SWE-Together/blob/main/eval/user_behavior/user_metrics.py) and [message tagger](https://github.com/Togetherbench/SWE-Together/blob/main/eval/user_behavior/tag_messages.py) — centralized taxonomy, correction/nudge weighting, persisted derivation, pinned tagger settings, and schema normalization.
- [Official Hugging Face dataset](https://huggingface.co/datasets/yifannnwu/SWE-Together) — 109-row browsable metadata surface, field definitions, Apache-2.0 license, and links back to runnable task artifacts.

## Findings

### Finding 1 — SWE-Together reconstructs a narrow, high-precision slice of real sessions

**Observed source facts.** SWE-Together begins with 11,260 recorded sessions from DataClaw (2,228), Pi-staging (2,397), Hyperswitch (784), and SWE-chat (5,851), retaining 109 tasks: a 0.97% conversion rate. The deterministic first stage requires multiple genuine user messages, concrete agent actions or edits, and enough repository signal; it favors public, sufficiently mature repositories and removes sessions where the human primarily authored the final change. A later LLM screen asks whether the substantive work can be reconstructed locally, excluding work whose primary deliverable depends on live services, credentials, deployments, pull-request management, or similar external state. A sandboxed task-generation agent then performs repository-grounded screening, pins the repository state, identifies setup and tests, and produces the instruction, environment, simulator prompt, verifiers, and evidence package. The viability screen explicitly does not establish correctness. [Paper, pp. 2–4](https://arxiv.org/pdf/2606.29957), [Design page](https://togetherbench.com/design.html).

**Inference.** The suite's real-session provenance is strong evidence that the *shape* of the conversations occurred in practice. It is not evidence that the 109 tasks represent the broader distribution of coding-agent work. The filters deliberately select public, code-changing, locally verifiable, agent-authored, multi-turn work. Ambiguous advice, private systems, deployment and operations work, read-only analysis, human-led corrections, and visually mediated work are underrepresented by construction.

**Evidence.** The official paper reports the source counts and 109/11,260 retention; the paper and website describe the three-stage funnel; representative task directories publish original session records, `oracle_intents.json`, a pinned environment, verifiers, a user-simulation prompt, and frozen goals. [Task artifact directory](https://github.com/Togetherbench/SWE-Together/tree/main/tasks/agent-swarm-task-4a881b).

**Implication.** Skill Issue's three scenarios should be described as deliberately fixed, discriminating evaluation cases, with an explicit selection rationale and coverage boundary. Avoid implying population-level representativeness from three cases. Publish which behaviors each scenario was selected to expose and which skill-calling conditions it does not cover.

### Finding 2 — Reactive replay preserves intent order while making feedback trajectory-dependent

**Observed source facts.** The evaluated agent receives the real user's first message and a restored repository. After each agent turn, a wrapper summarizes recent activity, output, elapsed time, and repository changes. The simulator combines that summary with fixed session anchors and prior simulator decisions, then chooses one structured action: `no-op`, `question`, `redirect`, `new_requirement`, or `check_external`. Silence is the default. The simulator is instructed to speak only when a recorded trigger fits the current trajectory, the agent is stuck or off track, the agent asks a question, or the agent attempts to finish while requirements remain. Its history records both prior decisions and messages to reduce repetition. The code contains guards and fallbacks that convert silence-like message content into no-op and recover short message-like text when structured tool use fails. [Paper, pp. 4–5](https://arxiv.org/pdf/2606.29957), [user simulator source](https://github.com/Togetherbench/SWE-Together/blob/main/src/user_agent/user_agent.py).

**Inference.** This is more defensible than replaying recorded follow-ups at fixed turn numbers because a new agent may solve, fail, or diverge differently. The design still replaces the human's live judgment with a model-mediated state summary and policy. Any information omitted or distorted by the summarizer cannot influence the simulated user, and the simulator's textual action set cannot reproduce interruption timing, direct edits, visual inspection, or unstructured collaborative behavior.

**Evidence.** The Design page makes evidence separation explicit: the agent sees the workspace, the simulator sees intent anchors and progress summaries, and the judge sees the completed patch and frozen rubric. The paper lists the same action space and says no-op does not consume a recorded follow-up. [Design page](https://togetherbench.com/design.html).

**Implication.** For Skill Issue, keep the three scenario turn scripts anchored to scenario-specific intent and trigger conditions rather than unconditional turn numbers. If the current evaluation uses fixed human-authored turns, preserve that simpler mechanism and document the branch conditions. If reactive simulation is introduced later, keep the simulator input, decision, and injected user message separately visible so failures can be attributed to the agent, summary, or simulator.

### Finding 3 — Repository correctness and conversational steering are intentionally separate axes

**Observed source facts.** SWE-Together scores the final repository state and the replay process separately. Task correctness uses a two-phase agentic rubric judge: phase 1 derives a weighted, implementation-agnostic rubric once per task; phase 2 applies the same frozen rubric to every candidate. The judge can inspect the repository and executable evidence and returns a binary met/unmet decision with evidence for each goal; the weighted score is computed mechanically. Host validation checks goal coverage, normalized weights, required fields, at least one core goal, and consistency between goal decisions and the score. User Correction is computed from an LLM multi-label tagger as `#correction + 0.2 × #nudge`; requests, questions, verification, workflow, approval, and context do not contribute. [Paper, pp. 5–7](https://arxiv.org/pdf/2606.29957), [rubric generation](https://raw.githubusercontent.com/Togetherbench/SWE-Together/main/eval/correctness/generate_task_goals.py), [User Correction source](https://github.com/Togetherbench/SWE-Together/blob/main/eval/user_behavior/user_metrics.py).

**Inference.** Separating outcome and steering prevents an agent that eventually succeeds after repeated rescue from looking equivalent to one that succeeds cleanly. It also avoids treating all additional user turns as failure: a new requirement or neutral question is different from a correction. The 0.2 nudge weight remains a normative design choice justified qualitatively in the paper; it is not presented as a measured conversion to human time or cognitive burden.

**Evidence.** The user metric taxonomy and weight live in one source file, and the tagger derives and persists the same value that the aggregator recomputes. The tagger reports a default Gemini 3.1 Pro call at temperature zero with a versioned prompt and taxonomy. [Message tagger](https://github.com/Togetherbench/SWE-Together/blob/main/eval/user_behavior/tag_messages.py).

**Implication.** Skill Issue should report at least two distinct scenario results: whether the required skill behavior and downstream artifact are correct, and how much explicit corrective steering was needed. For three fixed scenarios, raw counts and labeled turns are more legible than a synthesized scalar alone. A weighted correction score can be secondary if its taxonomy and weight are published.

### Finding 4 — The judge design has meaningful controls, but its primary score remains model-mediated

**Observed source facts.** Fixed tests are treated as useful evidence rather than the sole truth because alternative implementations and later conversational requirements can be missed or incorrectly rejected. Phase 1 may inspect the reference patch to infer behavior, but its resulting goals are frozen before candidate scoring. Tasks without extractable patches can build goals from recorded intents and user turns. The same rubric is reused across cohorts. The public task detail exposes weighted completeness goals and per-goal scores, making the aggregate judge score inspectable. [Paper, pp. 5–6](https://arxiv.org/pdf/2606.29957), [representative task detail](https://togetherbench.com/task.html?id=agent-swarm-task-4a881b).

**Inference.** Frozen criteria and candidate-independent decomposition reduce solution-specific grading drift. They do not eliminate model-judge bias, rubric-construction error, or evidence-reading error. The paper does not report a human agreement study for final correctness scores, judge-versus-expert calibration across the 109 tasks, or a multi-judge sensitivity analysis. The website FAQ identifies Opus 4.6 in an E2B sandbox as the agentic judge, so comparisons are partly conditional on that judge and prompt.

**Evidence.** The phase-1 source validates rubric structure and persists `canonical_goals.json`; the task page shows the individual goals and their weights. The official paper makes the agentic judge, not raw test reward, the primary correctness signal. [Design FAQ](https://togetherbench.com/design.html).

**Implication.** Skill Issue's three fixed scenarios should make the scoring contract more auditable than a single judge number: publish the exact required behaviors, evidence used, deterministic checks where possible, per-criterion outcomes, and judge rationale. If an LLM judge is used, preserve a fixed rubric, fixed judge configuration, and a small human audit or alternate-judge sensitivity check before treating tiny score differences as meaningful.

### Finding 5 — pass@1, SSR, pass², and mean judge answer different reliability questions

**Observed source facts.** Seven model cohorts run the common `opencode` harness across 109 tasks with `k = 2` replicates and a judge success threshold of 0.85. pass@1 is the marginal per-run success rate. Stable solve rate (SSR) averages a task's two continuous judge scores and then thresholds the average. pass² counts a task only when both runs clear the threshold, so it is the strictest two-run stability measure. Mean judge averages continuous task scores. Each task is equally weighted. The canonical plan fixes two replicates, high reasoning effort, 4,800-second agent timeouts, and 20 workers per cohort. [Paper, pp. 7–8](https://arxiv.org/pdf/2606.29957), [canonical plan](https://github.com/Togetherbench/SWE-Together/blob/main/canonical_full109.json).

**Inference.** The hatched `pass@1 − pass²` region on the site is an unusually clear visual cue for unstable success. With only two attempts per task, pass² is still a coarse stability indicator rather than a precise estimate of tail reliability. SSR may classify one strong and one weak run as stable when their mean crosses 0.85, while pass² will reject the same pair; the two should not be collapsed into one reliability claim.

**Evidence.** The leaderboard explicitly encodes the darker pass@1 segment, white pass² value, and hatched unstable tail. The paper defines all formulas and reports `pass² ≤ pass@1`. [Official overview](https://togetherbench.com/).

**Implication.** Three Skill Issue scenarios should use repeated attempts per scenario and show both single-run success and all-runs success. With only three scenarios, report the attempt matrix directly instead of relying on a headline percentage. If budget permits only two attempts, call it two-run stability; if stronger reliability claims are needed, increase repetitions rather than adding more derived metrics.

### Finding 6 — The headline results combine capability, stability, steering, and efficiency

**Observed source facts.** The paper reports: Claude Opus 4.8 at 63% pass@1, 59% SSR, 52% pass², 0.801 mean judge, 1.38 User Correction, 74.0k output-plus-reasoning tokens, and 23.3 minutes per task; GPT-5.5 at 58%, 55%, 48%, 0.763, 1.59, 29.9k, and 10.7 minutes; Claude Opus 4.6 at 58%, 58%, 46%, 0.755, 1.59, 42.0k, and 23.2 minutes. Remaining cohorts range down to MiniMax-2.7 at 40% pass@1, 34% SSR, 26% pass², 0.630 mean judge, and 2.17 corrections. Across seven cohorts, User Correction correlates with pass@1 at −0.92, SSR at −0.84, and mean judge at −0.93. The paper warns that latency may reflect serving location or inference infrastructure. [Paper, pp. 8–9](https://arxiv.org/pdf/2606.29957).

**Inference.** The inverse correction/capability relationship is suggestive and aligns with the benchmark hypothesis, but it is a correlation over seven model cohorts, not a causal estimate of human burden. Stronger agents may finish earlier, expose fewer natural opportunities for simulator intervention, or interact differently with the summarizer. Provider, endpoint, model speed, token accounting, and run date are potential confounders.

**Evidence.** The official site displays provider and June 2026 run dates beside the leaderboard: most cohorts via OpenRouter on June 8, 17, or 28, with DeepSeek via its provider on June 8. The paper separately reports tokens and wall-clock time and carries the infrastructure caveat. [Official overview](https://togetherbench.com/).

**Implication.** Skill Issue should place correction counts beside correctness and runtime/cost, not use corrections as a proxy for correctness. Record harness, model, provider route, reasoning setting, run date, timeout, token basis, and environment revision for every scenario campaign. If all three scenarios run under one local harness, say so and avoid provider-general conclusions.

### Finding 7 — The “Oracle” is a reference baseline, not a true ceiling

**Observed source facts.** The leaderboard labels the reference row “Oracle reference” and describes it as the gold-patch reference ceiling at approximately 78% and 0.904 mean judge. The paper is more precise: only 93 of 109 tasks have extractable reference patches; 73 of those 93 pass the 0.85 threshold, and 57 score 1.0. Among unsatisfied goals, roughly 35% are process requirements—such as diagnosing before editing, answering a follow-up, or explaining the change—that a final patch cannot express. Other misses come from incomplete multi-commit extraction and some imperfect human solutions. The paper explicitly says the row is a like-for-like reference point, not a strict ceiling on resolvability. [Official overview](https://togetherbench.com/), [paper, p. 9](https://arxiv.org/pdf/2606.29957).

**Inference.** “Oracle ceiling” is compact leaderboard language that overstates what the artifact proves. A candidate agent can legitimately outperform a stored patch on process or completeness, and the patch can be penalized for evidence it cannot contain. The divergence between site shorthand and paper qualification is exactly the kind of uncertainty that a public evaluation page should surface near the result.

**Evidence.** The paper accounts for all three reference shortfall classes and states that 16 tasks have no canonical diff. The task artifacts separately retain conversations, intents, and reference code, demonstrating that the expected behavior is broader than the patch alone.

**Implication.** For Skill Issue, label any gold trace or expected artifact as a “reference outcome” unless it is formally proven maximal. If a scenario rubric includes conversational process—correct skill invocation, clarification, refusal, or explanation—score the transcript and artifact together rather than expecting a final file to demonstrate the whole contract.

### Finding 8 — Simulator validation supports plausibility and cohort consistency, not full behavioral validity

**Observed source facts.** Intent Coverage matches atomic original-session intents to replay messages and combines 70% weighted intent recall with 30% scope precision. Six cohorts score 0.70–0.72 overall; GPT-5.5 scores 0.68. Recall spans 0.72–0.74 for six cohorts and precision 0.66–0.72. A human study uses four annotators, 52 tasks shared by three model cohorts, 156 paired trajectories, and 312 forced-choice judgments. The simulator is selected as real 46% of the time, with a 95% confidence interval of 40.5%–51.6%, so the annotators do not distinguish it from real users at a statistically reliable rate. [Paper, pp. 9–10](https://arxiv.org/pdf/2606.29957).

**Inference.** The study validates surface plausibility under a forced-choice protocol and shows no large cohort-specific shift in the reported Intent Coverage metric. It does not prove that interventions occurred at the correct moment, that omitted intents were harmless, that simulator messages caused the same agent behavior as real messages, or that the correction score matches human effort. Four annotators and 52 shared tasks are a useful initial study, but a limited basis for broad equivalence. An overall Intent Coverage near 0.70 also leaves material room for systematic omissions or extras even when cross-cohort values are stable.

**Evidence.** The public implementation pins the coverage judge to Gemini 3.1 Pro Preview by default at temperature zero, normalizes malformed match tables, fills missing intents as unmatched, bounds confidence, and computes the final numbers deterministically after the LLM match. This is a good reproducibility control around a still-model-mediated match. [Intent Coverage source](https://raw.githubusercontent.com/Togetherbench/SWE-Together/main/eval/user_behavior/coverage_one.py).

**Implication.** With three Skill Issue scenarios, prefer human-authored user turns or deterministic trigger scripts as the primary evaluation input. If an LLM simulator is used, validate trigger timing and intent fidelity scenario-by-scenario, publish the injected messages, and classify simulator failures separately. A Turing-style realism check is optional and should not replace behavioral fidelity checks.

### Finding 9 — Reproducibility is well scaffolded but operationally heavy and not independently demonstrated here

**Observed source facts.** The official repository publishes an Apache-2.0 harness, canonical task list, task directories, Docker environments, verifiers, simulator prompts, frozen rubrics, and launch commands. The Hugging Face dataset exposes one metadata row per task, including instruction, repository URL, base commit, language, difficulty, category, resource policy, test command, source files, reference patch, oracle intents, completeness goals, and test manifest. The launcher is dry-run by default and separates trial generation from judgment; trials are stored per model and replicate. Runs require model/provider keys plus either E2B or local Docker and access to GHCR images. [GitHub README](https://github.com/Togetherbench/SWE-Together), [Hugging Face dataset](https://huggingface.co/datasets/yifannnwu/SWE-Together), [launcher](https://github.com/Togetherbench/SWE-Together/blob/main/launch.py).

**Inference.** These artifacts make the evaluation inspectable and plausibly rerunnable. They do not guarantee bit-for-bit reproduction: proprietary model endpoints, provider routing, current model aliases, API behavior, container registry availability, and substantial execution cost remain external dependencies. The leaderboard publishes provider and run date but not an immutable result bundle, model snapshot identifier, sampling configuration for every component, or a signed release tag adjacent to each row.

**Evidence.** On 2026-07-21, the public repository page showed five commits and no published releases, while the paper was arXiv v1 from 2026-06-29 and the website runs were dated 2026-06-08 through 2026-06-28. This makes the project notably new and still likely to change. [GitHub repository](https://github.com/Togetherbench/SWE-Together), [arXiv record](https://arxiv.org/abs/2606.29957).

**Implication.** Skill Issue can adopt the artifact discipline at smaller scale: pin the repository revision, harness, model route, prompts, scenario definitions, and scoring criteria; retain per-attempt transcripts and outputs; and attach results to a dated campaign identifier. For three scenarios, a complete reproducibility bundle is feasible and preferable to a complex launcher that depends on mutable provider aliases.

### Finding 10 — The overview page prioritizes the decision, then exposes provenance

**Observed source facts.** The reader enters on a page titled “SWE-Together: Interactive Coding-Agent Evaluation.” The hero offers four direct actions—tasks, paper, contact, GitHub—before the leaderboard. The leaderboard combines pass@1/pass² bars, mean judge, corrections, tokens, and minutes, then immediately explains the threshold, the hatched instability tail, the reference row, the direction of each metric, `k = 2`, and the 109-task/common-harness context. Provider and run date appear as compact tags directly below. Contact and citation sections follow, and the footer repeats the navigation and one-sentence purpose. [Official overview](https://togetherbench.com/).

**Inference from visual inspection.** The page uses a wide but bounded central column, very large hero type, short paragraphs, high whitespace, black headings, pink accents, muted gray metadata, thin borders, and monospace utility labels. The leaderboard is the main dense object; surrounding copy is deliberately sparse. The design makes “what is this?” and “what won?” answerable before methodology, while paper/GitHub links keep the evidence path one click away.

**Evidence.** Full-page browser inspection confirmed the order hero → leaderboard → run metadata → contact → citation → footer and the strong visual hierarchy. The semantic HTML and official page text confirm the same sections and links.

**Implication.** Skill Issue's results entry page should lead with one sentence defining the three-scenario evaluation, the compact scenario result matrix, and immediate links to methodology, scenario evidence, and repository artifacts. Keep campaign metadata adjacent to the score rather than in a distant appendix.

### Finding 11 — The Design page is a progressive explanation, not a paper pasted into the site

**Observed source facts.** The page name is simply “Design.” Its internal sequence is: “From real sessions to verifiable tasks,” source-to-suite funnel, six numbered principles, replay-trace examples, “How a session becomes an evaluation,” task artifact, runtime loop, state separation, diagnostics, results, suite composition, and FAQ. The page alternates diagrams, compact cards, short prose, two-column method sections, result plots, small metric callouts, and expandable FAQ rows. Diagrams carry captions and a small expansion affordance. [Official Design page](https://togetherbench.com/design.html).

**Inference from visual inspection.** Body paragraphs are generally a readable 55–80 characters per line and two to five lines long; the page is long, but visual rhythm limits uninterrupted prose. Small uppercase/monospace eyebrow labels (“DESIGN,” “METHOD,” “TASK ARTIFACT,” “RUNTIME LOOP,” “RESULTS”) help readers scan. The density rises only where evidence requires it: diagrams, suite composition, and plots.

**Evidence.** The semantic page contains six principle headings, four method subsections, result and composition sections, and four FAQ questions; visual inspection confirmed repeated card/diagram/caption patterns and alternating two-column layouts.

**Implication.** Give Skill Issue one methodology page that progressively answers: why these three scenarios, what each contains, how turns are triggered, what the agent sees, how correctness is judged, what correction means, and what the limitations are. Prefer short evidence captions and diagrams over reproducing a long research report in the marketing layer.

### Finding 12 — The Task journey offers browse → compare → inspect → verify

**Observed source facts.** The top navigation uses singular “Task,” while the page heading is “Browse the tasks.” Users can search by task/repository/description, filter language and category, and sort by default, difficulty, corrections, user turns, diff size, or repository stars. Each card includes category and language, a descriptive title and short task summary, a seven-model result strip, difficulty, models passing, average score, corrections, repository, stars, user-turn count, file count, and a “View details” link. [Official Task browser](https://togetherbench.com/tasks.html).

The detail page begins with back-navigation, breadcrumb-like context, task ID, descriptive title, repository/category/language/base commit, then a five-cell summary strip. It pairs the task description with a model-results table. Lower sections expose the user interaction loop, weighted completeness goals and per-goal performance, environment/container policy, verifier command and test targets, source/test files, the original user session, and source metadata. Long original-session content has a “Show more” disclosure. [Representative task detail](https://togetherbench.com/task.html?id=agent-swarm-task-4a881b).

**Inference from visual inspection.** The browse cards support rapid comparison without forcing a detail visit. The detail page is intentionally dense but ordered from decision summary to progressively lower-level evidence. Its long vertical form is manageable because tables, labels, cards, monospaced code blocks, and disclosure controls separate evidence classes.

**Evidence.** The 109 task cards and representative detail were inspected in rendered form; the Hugging Face dataset exposes the same underlying fields programmatically. [Official dataset](https://huggingface.co/datasets/yifannnwu/SWE-Together).

**Implication.** With only three Skill Issue scenarios, skip a generic 109-item marketplace-style browser. Use three comparable scenario cards or tabs with a shared schema, then a detail view containing: scenario purpose, fixed turns/triggers, expected skill behavior, attempt results, corrections, transcript, deterministic evidence, rubric outcomes, environment, and artifact links. Search and multi-axis filters are lower fit unless the scenario library grows materially.

### Finding 13 — The Trial page turns trajectories into an editorial comparison tool

**Observed source facts.** The page name is singular “Trial,” while its eyebrow says “Trial examples · one task, two models.” Three editorial modes—“Depth vs. Surface,” “Targeted vs. Overreach,” and “Correction Response”—frame specific failure patterns. Each mode explains the task, names the comparison dimensions, and tells the reader what to watch. Score cards show judge score, turns, corrections, and outcome status. Readers can choose the middle and right model and the task. The core display is three side-by-side, independently scrollable columns: the original session and two model replays. Replay headers identify the Gemini simulator and the model; each turn shows the structured summary/decision evidence rather than only the final response. [Official Trial comparison](https://togetherbench.com/three_panel.html).

**Inference from visual inspection.** This page is the highest-density and widest part of the site. The editorial tabs prevent it from becoming an undirected transcript dump. Side-by-side comparison makes divergence visible, but independent nested scrolling can be demanding and is less suitable on narrow screens or for long traces without synchronization and anchors.

**Evidence.** Rendered inspection confirmed the three-column layout, comparison cards, model/task selectors, nested scroll regions, and turn-level elapsed/activity/output summaries. The public HTML includes the complete original and replay texts and 109 task options.

**Implication.** Skill Issue should publish at least one paired successful/failed or low-correction/high-correction trajectory per scenario. Add editorial annotations at the exact turn where skill invocation, correction handling, or scope discipline diverges. For three fixed scenarios, synchronized turn anchors or a single-column diff timeline may be more usable than three independent scroll panes. Preserve downloadable raw transcripts for audit.

### Finding 14 — The site links evidence well, but important caveats are unevenly placed

**Observed source facts.** Paper and GitHub links appear in the hero and persistent navigation; GitHub is also repeated in the footer. The overview publishes citation text and contact routes. Design diagrams are captioned, and the FAQ links a supporting OpenAI analysis for test false negatives. Task details surface repository/base-commit/source metadata and tests. The paper is the primary location for formulas, study details, limitations, and the qualification that the reference patch is not a strict ceiling. [Official overview](https://togetherbench.com/), [Design page](https://togetherbench.com/design.html), [paper](https://arxiv.org/pdf/2606.29957).

**Inference.** The public journey makes source artifacts discoverable, but readers who remain on the leaderboard may miss material uncertainty: two-run stability, judge mediation, the 93-task reference subset, the limited simulator study, and infrastructure effects on timing. Calling the reference a ceiling on the overview while qualifying it in the paper is the clearest presentation mismatch.

**Evidence.** The overview's compact note contains threshold, instability, provider, and run-date context, but not the 93/109 reference coverage or “not a strict ceiling” qualification. The limitations appear only near the paper's end.

**Implication.** Skill Issue should attach scenario-specific caveats beside the affected metric or evidence row. Put global limitations on the methodology page and repeat the most decision-relevant ones in a compact note under the headline result. Keep citations or artifact links close to every consequential claim.

### Finding 15 — Recommended Skill Issue adaptation for three fixed multi-turn skill-calling scenarios

**Observed source facts used.** SWE-Together's strongest transferable patterns are separate outcome/interaction metrics, repeated attempts, frozen criteria, turn-level evidence, explicit run metadata, inspectable task artifacts, and a progressive reader journey. Its lower-fit patterns for Skill Issue are the 109-card search/filter apparatus, heavy containerized task-construction pipeline, LLM viability screen, reactive simulator as default, and provider leaderboard framing. These are valuable for a large benchmark ecosystem but unnecessary overhead for three already-fixed scenarios. [GitHub README](https://github.com/Togetherbench/SWE-Together), [official site](https://togetherbench.com/).

**Inference and recommendation.** Use a four-page or four-section public journey:

1. **Overview:** define the evaluation and show a three-row scenario matrix with pass@1, all-runs-pass, mean rubric score, explicit corrections, tokens/time, harness/model, run date, and a visible limitations note.
2. **Method:** explain why the scenarios are fixed, the intended skill-calling behavior, what context the agent receives, how later user turns are triggered, what counts as correction, how artifacts are scored, repetition count, and controls.
3. **Scenarios:** give each scenario the same detail schema—purpose, initial request, later turns/triggers, expected skill behavior, exclusions, attempts, rubric outcomes, environment, raw transcript/artifact links, and observed failure modes.
4. **Trajectory comparison:** show annotated turn-level divergence between a strong and weak run, with skill invocation, user correction, and final-artifact evidence aligned.

For evaluation mechanics, run more than one attempt per scenario; preserve the complete attempt matrix; freeze scenario text, harness, model route, rubric, and deterministic checks; report corrections separately from neutral follow-ups; and retain the final artifact plus transcript. Use a reference trace as an example, not an oracle. [Paper methodology](https://arxiv.org/pdf/2606.29957).

**Evidence.** Every recommended field corresponds to an official SWE-Together surface or control: overview leaderboard and run metadata, frozen goals, User Correction taxonomy, task detail evidence, trial comparison, and reproducibility plan.

**Implication.** This keeps the credible parts of SWE-Together while matching the smaller fixed corpus. It gives a reader enough evidence to distinguish “the skill was invoked,” “the skill changed behavior correctly,” “the result was correct,” and “the user had to rescue the run.”

### Finding 16 — Conditional alternatives and lower-fit patterns

**Observed source facts.** SWE-Together's architecture is optimized for 109 heterogeneous repository tasks and seven model cohorts, not a three-scenario skill-calling evaluation. Its task browser, independently scrollable three-panel trajectories, LLM simulator, agentic judge, container registry, and multiple provider routes all solve scale or heterogeneity problems. [Task browser](https://togetherbench.com/tasks.html), [Trial comparison](https://togetherbench.com/three_panel.html), [canonical plan](https://github.com/Togetherbench/SWE-Together/blob/main/canonical_full109.json).

**Inference.** Several patterns become appropriate only if Skill Issue grows:

- **If scenarios expand beyond roughly a dozen:** add search, tags, difficulty, skill family, harness, and failure-mode filters; until then, three cards or tabs are clearer.
- **If later turns must adapt to divergent agent work:** introduce a reactive simulator with published triggers and decision logs; until then, deterministic turn scripts reduce an unnecessary source of variance.
- **If valid outputs admit many implementations:** use a frozen rubric judge plus executable evidence; if outputs are structurally deterministic, prefer direct validators and reserve LLM judgment for semantic residue.
- **If comparing many providers or model routes:** show provider/date/token/time metadata and stability tails; if evaluating one qualified harness/model pair, emphasize scenario evidence rather than a leaderboard.
- **If trajectories become very long:** offer synchronized side-by-side comparison with anchors and collapse controls; for short traces, an annotated single timeline is more readable.
- **If a gold reference is incomplete or process-blind:** label it as a reference and score transcript plus artifact; never promote it to a ceiling without proof.

**Evidence.** The paper's own reference-patch analysis, simulator limitations, and infrastructure caveat demonstrate why these choices are conditional rather than universal. [Paper, pp. 9–11](https://arxiv.org/pdf/2606.29957).

**Implication.** Skill Issue should adopt the minimum mechanism that preserves semantic validity for the three scenarios, while designing the public data schema so more scenarios, judges, or reactive turns can be added later without changing the meaning of existing results.

## Notes

- The project is exceptionally new: arXiv v1 was submitted 2026-06-29, the paper PDF is dated 2026-06-30, and the inspected repository had five commits and no releases on 2026-07-21. Results, artifacts, routes, or labels may change quickly.
- The arXiv experimental HTML omitted several mathematical values; exact metrics, confidence intervals, formulas, and reference counts were verified against the official PDF text.
- Browser inspection was used only to validate rendered hierarchy, density, controls, and disclosure patterns. Quantitative visual statements about text measure are approximate inferences, not values published by the authors.
- No independent execution was performed. The repository supports a plausible reproduction path, but operational reproducibility remains unverified in this assignment.
- The website's overview says “gold-patch reference ceiling,” while the paper explicitly rejects a strict-ceiling interpretation. Downstream synthesis should preserve the paper's qualification.
- No true research blocker remains. The principal unsupported claim would be any assertion that the simulator is behaviorally equivalent to real users or that the full leaderboard can be reproduced bit-for-bit from the current public artifacts; the inspected evidence does not establish either.

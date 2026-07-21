# Coding-Agent Evaluation Candidates

## Assignment

**Goal.** Identify authoritative public coding-agent benchmark and evaluation-methodology destinations that can inform a dense, credible public Testing Methodology page for Skill Issue.

**Scope.** Internet-only research into primary sources published by benchmark authors, research organizations, and official repositories. Candidate quality is judged against ten page-level coverage needs: evaluation subjects, tasks/scenarios, conversational turns or trajectories, execution/instrumentation, scoring, controls/environment isolation, reproducibility, evidence access, limitations, and uncertainty.

**Exclusions.** Secondary benchmark roundups, vendor marketing without a disclosed methodology, model-only code-generation tests, private datasets without useful public methodology, and synthesis of Skill Issue's own final methodology.

## Sources

- Artificial Analysis, [Coding Agent Index Methodology](https://artificialanalysis.ai/methodology/coding-agents-benchmarking) and [public Coding Agent Index](https://artificialanalysis.ai/agents/coding-agents).
- SWE-bench authors, [official SWE-bench repository and harness](https://github.com/SWE-bench/SWE-bench); OpenAI, [Introducing SWE-bench Verified](https://openai.com/index/introducing-swe-bench-verified/).
- OpenAI, [Introducing the SWE-Lancer benchmark](https://openai.com/index/swe-lancer/) and [official Preparedness repository](https://github.com/openai/preparedness).
- Laude Institute, [Introducing Terminal-Bench](https://www.tbench.ai/news/announcement) and [official Terminal-Bench repository](https://github.com/laude-institute/terminal-bench).
- SWE-Together authors, [interactive benchmark site](https://togetherbench.com/), [paper](https://arxiv.org/abs/2606.29957), and author-linked repository from the benchmark site.
- Microsoft Research, [AgentLens publication page](https://www.microsoft.com/en-us/research/publication/agentlens-revealing-the-lucky-pass-problem-in-swe-agent-evaluation/) and [released AgentLens-Bench repository](https://github.com/microsoft/code-agent-state-trajectories/).
- METR, [official RE-Bench repository](https://github.com/METR/RE-Bench) and [paper](https://metr.org/research/re-bench-evaluating-frontier-ai-rd-capabilities-of-language-model-agents-against-human-experts/).
- OpenAI, [PaperBench](https://openai.com/index/paperbench/), [paper](https://arxiv.org/abs/2504.01848), and [official repository](https://github.com/openai/paperbench).
- Microsoft Research, [SWE-bench Goes Live](https://www.microsoft.com/en-us/research/publication/swe-bench-goes-live/), [project site](https://swe-bench-live.github.io/), and [paper](https://arxiv.org/abs/2505.23419).
- FeatureBench authors, [official project site](https://libercoders.github.io/FeatureBench/) and [paper](https://arxiv.org/abs/2602.10975).
- ByteDance Seed, [official Multi-SWE-bench repository](https://github.com/multi-swe-bench/multi-swe-bench), [project site](https://multi-swe-bench.github.io/), and [paper](https://arxiv.org/abs/2504.02605).
- OpenAI, [MLE-bench](https://openai.com/index/mle-bench/) and [official repository](https://github.com/openai/mle-bench).
- GitTaskBench authors, [official project site](https://gittaskbench.github.io/), [official repository](https://github.com/QuantaAlpha/GitTaskBench), and [paper](https://arxiv.org/abs/2508.18993).
- Sierra Research, [official tau-bench repository](https://github.com/sierra-research/tau2-bench), [project site](https://taubench.com/), and [tau2-bench paper](https://arxiv.org/abs/2506.07982).
- OSWorld authors, [official OSWorld repository](https://github.com/xlang-ai/OSWorld) and [project site](https://os-world.github.io/).
- Supporting integrity source: OpenAI, [Separating signal from noise in coding evaluations](https://openai.com/index/separating-signal-from-noise-coding-evaluations/).

## Findings

### Finding 1: Fifteen candidates separate into nine deep dives, five skims, and one rejection

Coverage ratings describe the publicly linked destination cluster, not an independently reproduced benchmark run. `H` means the source directly and substantially documents the dimension, `M` means partial or indirect coverage, and `L` means weak or absent coverage.

| Candidate                              | Class     | Subjects | Tasks | Turns / trajectories | Execution | Scoring | Controls | Reproducibility | Evidence | Limitations | Uncertainty |
| -------------------------------------- | --------- | -------: | ----: | -------------------: | --------: | ------: | -------: | --------------: | -------: | ----------: | ----------: |
| Artificial Analysis Coding Agent Index | Deep-dive |        H |     H |                    L |         H |       H |        M |               M |        M |           M |           M |
| SWE-bench plus Verified                | Deep-dive |        H |     H |                    L |         H |       H |        H |               H |        H |           H |           L |
| SWE-Lancer                             | Deep-dive |        H |     H |                    L |         H |       H |        H |               M |        M |           H |           L |
| Terminal-Bench                         | Deep-dive |        H |     H |                    M |         H |       H |        H |               H |        H |           M |           H |
| SWE-Together                           | Deep-dive |        H |     H |                    H |         H |       H |        M |               H |        H |           M |           H |
| AgentLens / AgentLens-Bench            | Deep-dive |        H |     H |                    H |         H |       H |        M |               H |        H |           H |           M |
| RE-Bench                               | Deep-dive |        H |     H |                    L |         H |       H |        H |               H |        H |           H |           H |
| PaperBench                             | Deep-dive |        H |     H |                    L |         H |       H |        M |               H |        H |           H |           M |
| SWE-bench Live                         | Deep-dive |        H |     H |                    L |         H |       H |        H |               H |        M |           H |           M |
| FeatureBench                           | Skim-only |        H |     H |                    L |         H |       H |        H |               H |        H |           M |           L |
| Multi-SWE-bench                        | Skim-only |        H |     H |                    L |         H |       H |        H |               H |        H |           M |           L |
| MLE-bench                              | Skim-only |        H |     H |                    L |         H |       H |        M |               H |        H |           H |           M |
| GitTaskBench                           | Skim-only |        H |     H |                    L |         H |       H |        M |               H |        H |           M |           L |
| tau2-bench                             | Skim-only |        M |     H |                    H |         H |       H |        H |               H |        H |           H |           M |
| OSWorld                                | Reject    |        L |     H |                    M |         H |       H |        H |               M |        H |           M |           M |

**Evidence.** The selected sources expose complementary methodology surfaces: Artificial Analysis publishes benchmark composition, repeats, pass/fail aggregation, cost, token, runtime, missing-telemetry handling, agent-setting separation, and version history; SWE-Together publishes `k = 2`, `pass@1`, `pass^2`, correction turns, tokens, elapsed minutes, provider, run date, and an oracle ceiling; AgentLens releases 1,815 annotated trajectories with process-quality and waste signals; SWE-bench, Terminal-Bench, Multi-SWE-bench, and SWE-bench Live expose executable harness or container details; RE-Bench, PaperBench, MLE-bench, SWE-Lancer, and Verified expose human review or human baselines.

**Implication.** No single destination is a sufficient template. A credible page should combine the structural clarity of Artificial Analysis, the reproducible execution detail of SWE-bench and Terminal-Bench, the interaction evidence of SWE-Together, the process-quality analysis of AgentLens, and the human-validation and uncertainty practices of METR and OpenAI research benchmarks.

### Finding 2: The deep-dive set supplies the strongest directly reusable methodology patterns

1. **Artificial Analysis Coding Agent Index — deep-dive.** This is the closest public analogue to a concise methodology page for multiple coding-agent evaluations. It explicitly distinguishes agent variants from model names, uses three repeats per component, aligns performance and efficiency metrics to the same attempts, states missing-telemetry treatment, and maintains a version history. Its main gaps are raw task-level evidence, environment-isolation detail, failure taxonomy, and confidence intervals. [Methodology](https://artificialanalysis.ai/methodology/coding-agents-benchmarking)
2. **SWE-bench plus SWE-bench Verified — deep-dive.** The benchmark family provides the canonical issue-to-patch task contract, Docker evaluation, executable test scoring, datasets, build and evaluation logs, and a public harness. Verified adds a 500-task subset screened by professional developers and openly acknowledges public-repository contamination risk. It is weak on interactive turns, process quality, repeated trials, and statistical uncertainty. [Repository](https://github.com/SWE-bench/SWE-bench) · [Verified methodology](https://openai.com/index/introducing-swe-bench-verified/)
3. **SWE-Lancer — deep-dive.** The benchmark expands subjects beyond bug fixing to economically valued feature, frontend, performance, and managerial tasks. Independent tasks use end-to-end tests triple-verified by experienced engineers; the public Diamond split uses a unified Docker image, and the July 2025 update removed internet dependence to reduce variability. Only a public split is exposed, and the landing methodology does not quantify run-to-run uncertainty. [Benchmark page](https://openai.com/index/swe-lancer/)
4. **Terminal-Bench — deep-dive.** Each task has a dedicated Docker environment, human-verified solution, and executable tests. The harness starts multi-container environments, logs actions, verifies container state, and supports container, direct Python, and MCP integrations. The published results describe multiple runs and error bars, making it unusually useful for uncertainty reporting, although task scope extends beyond software engineering. [Announcement and methodology](https://www.tbench.ai/news/announcement)
5. **SWE-Together — deep-dive.** This is the strongest coding-specific destination for conversational evaluation. It reconstructs 109 repository tasks from 11,260 real sessions, uses a reactive user simulator, scores final repository correctness and corrective feedback, and publishes two-run stability through `pass^2` alongside judge score, corrections, tokens, minutes, provider, and dates. It is new, uses an LLM simulator and judge, and therefore needs careful disclosure of simulator and judge validation. [Project site](https://togetherbench.com/) · [Paper](https://arxiv.org/abs/2606.29957)
6. **AgentLens / AgentLens-Bench — deep-dive.** AgentLens directly addresses the “lucky pass” problem in binary coding-agent scores. Microsoft reports 2,614 OpenHands trajectories, releases a 1,815-trajectory annotated subset, and separates exploration, implementation, verification, and orchestration using trajectory context. Quality, waste, divergence, and process-reference evidence can complement final test results. Its narrow OpenHands/SWE-bench-Verified basis limits generalization across harnesses. [Microsoft Research](https://www.microsoft.com/en-us/research/publication/agentlens-revealing-the-lucky-pass-problem-in-swe-agent-evaluation/) · [Repository](https://github.com/microsoft/code-agent-state-trajectories/)
7. **RE-Bench — deep-dive.** RE-Bench is valuable for long-horizon execution, human-expert comparison, resource limits, task-specific continuous scoring, standardized task packaging, open agent scaffolding, and explicit contamination controls. The official repository also protects solution material and documents its anti-overfitting rationale. The tasks are AI research engineering rather than everyday software maintenance. [Repository](https://github.com/METR/RE-Bench) · [Paper](https://arxiv.org/abs/2411.15114)
8. **PaperBench — deep-dive.** PaperBench decomposes 20 full research replications into 8,316 author-informed rubric items, validates its automated judge through a separate JudgeEval, uses open scaffolding, and includes an expert-human baseline. It is a strong source for hierarchical scoring, grader validation, partial credit, and evidence-rich task decomposition, though its LLM judge introduces uncertainty and its domain is research replication. [Benchmark page](https://openai.com/index/paperbench/) · [Repository](https://github.com/openai/paperbench)
9. **SWE-bench Live — deep-dive.** Microsoft’s live-updatable benchmark targets contamination and staleness with 1,319 post-2024 issues across 93 repositories, an automated curation pipeline, and one Docker image per task. It also analyzes results by repository origin, recency, and difficulty under controlled conditions. Public trajectory and repeat-level evidence appears less prominent than the construction and environment story. [Microsoft Research](https://www.microsoft.com/en-us/research/publication/swe-bench-goes-live/) · [Project site](https://swe-bench-live.github.io/)

**Evidence.** These nine sources collectively disclose benchmark subjects, task contracts, environment construction, scoring, human or programmatic validation, and limitations. SWE-Together and AgentLens are the only deep-dive candidates here that treat user interaction or execution trajectory quality as a first-class outcome rather than a means to a final patch.

**Implication.** Later researchers should allocate separate deep dives to these sources rather than treating “coding-agent benchmark methodology” as a single SWE-bench-shaped domain. The sources answer different questions: final correctness, interaction burden, process quality, economic value, long-horizon work, grader validity, and contamination resistance.

### Finding 3: Five skim-only candidates fill narrow but useful gaps

1. **FeatureBench — skim-only.** Its 200 feature-development tasks, 3,825 environments, fail-to-pass plus pass-to-pass tests, dynamic dependency tracing, post-verification, token reporting, and explicit warning that Claude Code routing may involve multiple models make it useful for feature work and configuration disclosure. The public destination has little on repeated runs, uncertainty, turns, or trace evidence. [Project site](https://libercoders.github.io/FeatureBench/)
2. **Multi-SWE-bench — skim-only.** Its seven-language, expert-curated issue-resolution set broadens subject coverage and provides Docker images, configuration, logs, resolved/unresolved reports, and explicit environment-clearing controls. Methodologically it remains close to SWE-bench, so it adds language and platform breadth more than a distinct page structure. [Repository](https://github.com/multi-swe-bench/multi-swe-bench)
3. **MLE-bench — skim-only.** The 75 Kaggle competitions provide real ML engineering work, public human leaderboard baselines, open scaffolds, resource-scaling analysis, and contamination analysis. It is valuable for resource budgets and human comparison but does not model user turns or conventional repository maintenance. [Benchmark page](https://openai.com/index/mle-bench/)
4. **GitTaskBench — skim-only.** Its 54 tasks across seven modalities and domains cover repository discovery, setup, execution, iteration, and delivery. Automated human-curated harnesses distinguish execution completion from task pass and add a cost-effectiveness metric. The small task set, absent repeated-run uncertainty, and broad repo-leveraging framing make it secondary. [Project site](https://gittaskbench.github.io/) · [Repository](https://github.com/QuantaAlpha/GitTaskBench)
5. **tau2-bench — skim-only.** This is not a coding benchmark, but its turn-based dual-control interaction, explicit policies and tools, saved trajectories, task-schema documentation, versioned grader fixes, re-scoring instructions, and non-comparability warnings are directly useful for conversational methodology and benchmark version governance. [Repository](https://github.com/sierra-research/tau2-bench)

**Evidence.** FeatureBench and Multi-SWE-bench add task-shape and language diversity; MLE-bench adds resource and human-baseline practice; GitTaskBench adds staged workflow outcomes and cost; tau2-bench adds interaction orchestration and strong version-change disclosure.

**Implication.** These candidates merit targeted extraction only. They should not displace the deep-dive sources because each supplies one missing dimension without covering the full coding-agent evaluation page.

### Finding 4: OSWorld should be rejected for this coding-specific page

**OSWorld — reject.** OSWorld is authoritative and methodologically strong for real-computer agents: it provides state setup, execution-based evaluators, virtual-machine or cloud environments, task assets, and a Verified update that fixed benchmark issues. Its 369 tasks focus on desktop and web applications rather than coding-agent behavior, repository edits, software validation, or developer interaction. [Repository](https://github.com/xlang-ai/OSWorld) · [Project site](https://os-world.github.io/)

**Evidence.** The official repository emphasizes VMware, VirtualBox, Docker, AWS, Azure, desktop applications, screenshots, and computer-control environments. It does not define a software-engineering task contract comparable to issue resolution, feature work, research engineering, or interactive coding sessions.

**Implication.** Environment-reset and evaluator patterns can be rediscovered through coding-relevant sources such as Terminal-Bench and SWE-bench. A dedicated OSWorld deep dive would spend budget on domain translation rather than filling a unique methodological gap.

### Finding 5: The candidate set exposes six cross-cutting gaps a final methodology must make explicit

1. **Evaluation unit ambiguity.** Many leaderboards label rows by model even though outcomes depend on the model, agent harness, system prompt, tools, reasoning setting, provider, retry policy, and budget. Artificial Analysis and SWE-Together are stronger because they expose more of this composite subject.
2. **Static-task bias.** SWE-bench-shaped evaluations give the full request up front. SWE-Together shows why clarification, correction, and interaction burden require separate measures.
3. **Outcome-only bias.** Passing tests can conceal wasteful or brittle trajectories. AgentLens demonstrates that process-quality rankings can diverge substantially from pass-rate rankings.
4. **Weak uncertainty reporting.** Most coding benchmarks publish a single pass rate. Terminal-Bench error bars, Artificial Analysis repeats, and SWE-Together `pass^2` provide better patterns, but confidence intervals, per-task repeat distributions, and failure/infra-error separation remain uncommon.
5. **Benchmark drift and contamination.** SWE-bench Verified discloses public-data contamination; SWE-bench Live refreshes tasks; tau2-bench versions grader fixes and warns about non-comparable scores; OpenAI’s coding-evaluation audit shows that passing leaderboards can still contain broken or misleading tasks.
6. **Evidence-access trade-offs.** Open harnesses, logs, tasks, trajectories, and per-task results support scrutiny, while public solutions can contaminate future models. RE-Bench’s protected solution material and SWE-bench Multimodal’s private test evaluator illustrate controlled-evidence alternatives.

**Evidence.** The candidate destinations repeatedly separate final outcome, runtime behavior, and benchmark integrity, but no source combines multi-turn interaction, trace-quality analysis, deterministic final scoring, repeated trials with statistical uncertainty, strict environment isolation, complete configuration disclosure, and a public evidence bundle.

**Implication.** A later synthesis should treat these as required methodology fields rather than optional caveats. Unsupported claims about general coding-agent quality should be avoided when a result covers only one model-agent-provider configuration, one task family, one attempt, or one benchmark version.

## Notes

- The strongest newly discovered coding-specific sources for conversational and trajectory evaluation are SWE-Together and AgentLens; both are 2026 releases, so their methods are promising but less battle-tested than SWE-bench.
- OpenAI’s July 2026 [coding-evaluation audit](https://openai.com/index/separating-signal-from-noise-coding-evaluations/) is a high-value supporting source for task-quality review, broken-task estimation, human-supervised agent review, human annotation, and the need to inspect prompts, tests, patches, traces, and edge cases. It is an audit methodology rather than a benchmark candidate.
- “Reproducible” in this document means the authors expose enough code, data, environment, or execution procedure to support rerunning; no benchmark was independently rerun during this internet-only assignment.
- “Uncertainty” includes repeated trials, stability measures, error bars, confidence intervals, or explicit score variability. A leaderboard with many tasks but one attempt per task was rated low unless it separately addressed run-to-run variation.
- Useful follow-up search terms: `coding agent pass^2`, `trajectory quality benchmark`, `lucky pass`, `interactive coding benchmark`, `benchmark task integrity`, `agent harness comparison`, `infrastructure error rate`, `paired agent evaluation`, `coding benchmark contamination`, and `grader validation`.

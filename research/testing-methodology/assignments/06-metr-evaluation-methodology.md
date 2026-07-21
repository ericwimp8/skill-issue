# METR Evaluation Methodology

## Assignment

**Goal.** Identify the methodological and presentation patterns METR uses to evaluate coding-capable and autonomous AI agents, then derive evidence-backed practices that may improve Skill Issue's evaluation methodology and its public explanation.

**Scope.** Official METR methodology pages, papers, task standards, task-development guidance, evaluation reports, public task and analysis repositories, and current evaluation infrastructure were inspected as of 2026-07-21. The research covers subjects, task construction, human calibration, repeated trajectories, instrumentation, scoring, elicitation, environment controls, reproducibility, retained evidence, uncertainty, limitations, and information design.

**Exclusions.** This assignment does not assess METR's policy conclusions, reproduce its numerical results, inspect non-public task material, evaluate Skill Issue's local implementation, or recommend copying METR's wording, branding, or visual identity. Recommendations are therefore methodological patterns, with fit caveats where they depend on Skill Issue's product goals.

## Sources

- METR, “Task-Completion Time Horizons of Frontier AI Models,” current methodology and results page, updated 2026-05-08: <https://metr.org/time-horizons/>.
- Kwa et al., “Measuring AI Ability to Complete Long Software Tasks,” arXiv:2503.14499, version 4 dated 2026-07-10: <https://arxiv.org/pdf/2503.14499>.
- METR, “Time Horizon 1.1,” 2026-01-29: <https://metr.org/blog/2026-1-29-time-horizon-1-1/>.
- METR, “Measuring time horizon using Claude Code and Codex,” 2026-02-13: <https://metr.org/notes/2026-02-13-measuring-time-horizon-using-claude-code-and-codex/>.
- METR, HCAST paper, “A Benchmark of High-Impact, Long-Horizon Autonomous Tasks”: <https://metr.org/hcast.pdf>.
- METR Task Standard repository and specification, `main`, specification version 0.5.0 at inspection time: <https://github.com/METR/task-standard> and <https://github.com/METR/task-standard/blob/main/STANDARD.md>.
- METR, “Announcing the METR Task Standard,” 2024-02-29: <https://metr.org/blog/2024-02-29-metr-task-standard/>.
- METR Task Development Guide: <https://taskdev.metr.org/introduction/>, <https://taskdev.metr.org/desiderata/>, <https://taskdev.metr.org/setup/>, <https://taskdev.metr.org/specification/>, <https://taskdev.metr.org/implementation/>, <https://taskdev.metr.org/quality-assurance/>, and <https://taskdev.metr.org/documentation/>.
- METR public task examples, `main`: <https://github.com/METR/public-tasks>.
- METR public evaluation-analysis pipeline, `main`: <https://github.com/METR/eval-analysis-public>.
- METR public Inspect task bridge, `main`: <https://github.com/METR/inspect-metr-task-bridge/>.
- METR public Inspect task and agent collections, both `main`: <https://github.com/METR/inspect-tasks-public> and <https://github.com/METR/inspect-agents>.
- Vivaria documentation and architecture: <https://vivaria.metr.org/> and <https://vivaria.metr.org/architecture/>.
- Hawk documentation, security model, and repository: <https://hawk.metr.org/>, <https://hawk.metr.org/infrastructure/security/>, and <https://github.com/METR/hawk>.
- METR, “Autonomy Evaluation Resources,” 2024-03-13: <https://metr.org/blog/2024-03-13-autonomy-evaluation-resources/>.
- METR, “Details about METR's preliminary evaluation of Claude 3.5 Sonnet,” 2024-10-01: <https://metr.org/evaluations/claude-3-5-sonnet-report/>.
- METR, “GPT-5.6 Sol evaluation,” 2026-06-26: <https://metr.org/blog/2026-06-26-gpt-5-6-sol/>.
- Kinniment et al., “Evaluating Language-Model Agents on Realistic Autonomous Tasks,” 2023: <https://metr.org/Evaluating_LMAs_Realistic_Tasks.pdf>.

## Findings

### The evaluated subject is a configured agent system

**Fact.** METR consistently describes results as properties of a model combined with a scaffold, tools, prompt, budgets, and environment. The original time-horizon paper fits a separate success curve for each configured agent. The 2023 pilot explicitly warns that its results measure the agents it constructed and do not upper-bound the underlying model. Current time-horizon operations likewise begin by combining a model and scaffold, then checking elicitation on a development set before measuring a separate test set.

**Inference.** A model label alone is an insufficient experimental subject. Harness choice, executable version, reasoning configuration, context management, tool access, time and token limits, retry behavior, and environment are treatment variables that can change the outcome.

**Evidence**

- The current methodology describes a model-plus-scaffold subject, development-set elicitation, held-out test tasks, repeated runs, and post-run review: <https://metr.org/time-horizons/>.
- The 2023 report defines agents as specific combinations of models and scaffolds and describes the action-observation loop and VM execution: <https://metr.org/Evaluating_LMAs_Realistic_Tasks.pdf>.
- METR's Claude Code/Codex study found no significant aggregate advantage for specialized scaffolds in that experiment, while still treating scaffold and budget as explicit experimental factors: <https://metr.org/notes/2026-02-13-measuring-time-horizon-using-claude-code-and-codex/>.

**Implication**

- Skill Issue should give every evaluated subject an immutable run identity containing model, provider/version where available, harness and executable, scaffold/prompt revision, reasoning setting, tools, budgets, task-suite revision, and environment revision.
- Comparisons should name the full configured subject in tables and artifacts; a shorter display label may be used only when it resolves to the full manifest.
- This is a high-fit pattern for any harness evaluator. METR's specific scaffold architecture is lower fit unless Skill Issue exposes the same action loop.

### Tasks are authored as portable, inspectable experimental units

**Fact.** The METR Task Standard separates a task's environment, initial instruction, and optional scoring procedure. It requires ordered setup, constrains information visibility, defaults toward restricted networking, and supports either end scoring or intermediate scoring with aggregation. A task family groups related task instances under a versioned semantic directory. The development guide asks authors to specify purpose, resources, prompt, expected solution path and bottlenecks, human-time estimate, scoring rubric, and any human oversight or simulation before implementation.

**Inference.** Portability comes from making the task contract independent of the evaluation driver while keeping agent-visible information, hidden scoring material, and privileged setup sharply separated.

**Evidence**

- Task Standard version 0.5.0 specifies environments, instructions, scoring, drivers, setup ordering, permissions, and family layout: <https://github.com/METR/task-standard/blob/main/STANDARD.md>.
- The announcement explains the portability and reduced duplicated authoring work intended by the standard: <https://metr.org/blog/2024-02-29-metr-task-standard/>.
- The Task Development Guide covers desiderata, specification, implementation, QA, and documentation as successive authoring stages: <https://taskdev.metr.org/introduction/>.
- The public task repository demonstrates semantic family directories and suite manifests while deliberately withholding protected task material: <https://github.com/METR/public-tasks>.

**Implication**

- Skill Issue should define one canonical task contract with: semantic ID and family; task version; setup and fixtures; agent-visible instruction; environment and network policy; termination conditions; score interface; score range and threshold; protected resources; and required retained artifacts.
- A driver-specific adapter should preserve that contract rather than silently changing prompts, mounts, network access, or scoring behavior.
- Family grouping is useful when task variants share a latent skill, because it enables family-aware sampling and weighting. It is unnecessary for isolated one-off checks.

### Task development is a staged evidence-producing process

**Fact.** METR's development guidance prefers meaningful, coherent tasks with serially dependent hard steps, novelty, reproducible local services, isolated execution, and automatic scores between zero and one. Independent QA is performed by someone other than the author in the same task environment and with the same resources. QA runners record progress periodically, preserve intermediate states, and document discrepancies. When full completion is infeasible, expedited QA still requires environment exploration, an approach, partial solutions or hints, and a grounded time estimate.

**Inference.** QA is not just a final pass/fail check. It is the process that validates the task contract, produces human calibration evidence, finds scoring defects, and leaves a reviewable trail.

**Evidence**

- Task desiderata cover duration, realism, serial bottlenecks, novelty, consistency, safety, graded scoring, scoring robustness, anti-cheating, and simplicity: <https://taskdev.metr.org/desiderata/>.
- Specification guidance requires purpose, setup, prompt, solution steps, time/cost, score design, rubric, and oversight details: <https://taskdev.metr.org/specification/>.
- QA guidance requires independent execution, discrepancy reporting, progress logs, intermediate states, and a documented review: <https://taskdev.metr.org/quality-assurance/>.
- Documentation guidance names retained files such as `eval_info.json`, summaries, detailed task information, QA records, progress, review, and participant work: <https://taskdev.metr.org/documentation/>.

**Implication**

- Skill Issue should distinguish authoring, independent QA, pilot elicitation, and measured evaluation as separate lifecycle states with separate artifacts and approvals.
- Task evidence should retain author intent, QA observations, known valid and invalid solution paths, scoring tests, and any accepted discrepancy. That record enables later investigators to tell a harness failure from a broken task.
- Human baselines and periodic work logs are lower fit for short deterministic tests, but independent QA and scoring tests remain high fit.

### Repeated trajectories and transcript review protect the score's meaning

**Fact.** METR measures multiple independent attempts per task because agent success is stochastic. The current time-horizon workflow reports approximately six runs per task and around one thousand runs in a typical evaluation; the original paper used roughly eight, HCAST used at least five, and the Claude 3.5 Sonnet report used ten for general-autonomy tasks and five for AI R&D. Automated flags and human review identify infrastructure failures, reward hacking, budget anomalies, and scorer errors. HCAST's QA process ran an agent five times on every task and manually reviewed 945 transcripts.

**Inference.** A nominal score is trustworthy only when the system can reconstruct how it arose and exclude or separately classify invalid trajectories. A single attempt measures one sample, not a stable agent capability.

**Evidence**

- Current operational steps include independent repetitions, infrastructure restarts, reward-hack and budget checks, automatic and manual flagging, multiple reviewers, and rescoring: <https://metr.org/time-horizons/>.
- HCAST reports five agent attempts per task during QA, transcript inspection, and separate treatment of legitimate success, reward hacking, and failures beyond the agent's control: <https://metr.org/hcast.pdf>.
- The Claude 3.5 Sonnet report combines aggregate results with manual inspection of successful and failed trajectories, including a 95-run failure-classification sample: <https://metr.org/evaluations/claude-3-5-sonnet-report/>.

**Implication**

- Skill Issue should retain each attempt as an append-only trajectory with subject manifest, task revision, timestamps, tool events, outputs, submission, raw score, resource use, termination reason, and review annotations.
- Aggregate results should show attempt count and distinguish valid failure, infrastructure invalidation, policy/safety interruption, scorer defect, and suspected exploit.
- Independent repetition is high fit for generative agents. The number of runs can scale to cost; the report should expose when an estimate rests on a small sample.

### Scoring separates raw evidence, decision rules, and aggregation

**Fact.** METR tasks may emit continuous scores, but time-horizon analysis converts them to binary success using task-specific thresholds intended to match target human performance. It then models success probability as a logistic function of log human task duration. Human duration is generally the geometric mean of successful baselines. The analysis applies inverse-square-root family weighting and hierarchical bootstrap resampling across families, tasks, and runs. Public analysis data retain task/family identity, raw or binarized score, human duration, and weight.

**Inference.** A scalar headline should be treated as the final layer of a traceable scoring pipeline. Raw task evidence, threshold decisions, weighting, model fitting, and uncertainty should remain inspectable and recomputable.

**Evidence**

- The time-horizon paper defines task-specific thresholds, logistic success curves, geometric-mean human duration, family weighting, and a 10,000-sample hierarchical bootstrap: <https://arxiv.org/pdf/2503.14499>.
- The live methodology explains the 50% and 80% time-horizon intersections and exposes interactive filters and fitted curves: <https://metr.org/time-horizons/>.
- The public DVC analysis repository exposes raw JSONL, parameters, pipeline stages, plots, and report-specific analysis directories: <https://github.com/METR/eval-analysis-public>.
- The Claude 3.5 Sonnet report explains family weighting and time-bucket normalization alongside 95% bootstrapped confidence intervals: <https://metr.org/evaluations/claude-3-5-sonnet-report/>.

**Implication**

- Skill Issue should retain the raw rubric dimensions or task score before any pass threshold, then version thresholds and aggregation rules separately.
- Reports should expose weighting and confidence or variability, especially when tasks have families, unequal sampling, or stochastic attempts.
- Logistic human-time modelling is conditional: it fits only if Skill Issue intends to measure task-completion horizon against calibrated human durations. A simpler pass-rate or rubric aggregate is preferable when human time is not the construct of interest.

### Human baselines define a comparator and introduce measurable bias

**Fact.** METR attempts to place humans in the same environment with the same task information, often through SSH or a remote editor, while allowing browser research but excluding AI assistance. HCAST recruits relevant professionals, uses qualification, pays for time and performance, and records expertise. Yet baseline coverage is incomplete, successful-only durations create selection bias, task-expertise matching is imperfect, human performance varies substantially, and some task durations are researcher forecasts. HCAST reports modest agreement between forecasts and measured times; Time Horizon 1.1 notes that only five of its 31 tasks longer than eight hours had human baselines.

**Inference.** “Human time” is not an objective task property. It is an estimate conditional on the sampled population, expertise, context, tools, success filtering, and protocol.

**Evidence**

- HCAST documents recruiting, qualification, compensation, environment parity, baseline success, forecast use, expertise concerns, and timing limitations: <https://metr.org/hcast.pdf>.
- The time-horizon paper describes more than 800 baselines, successful-run geometric means, researcher estimates, low-context professionals, and selection bias: <https://arxiv.org/pdf/2503.14499>.
- The Time Horizon 1.1 release discloses sparse long-task baseline coverage and wider uncertainty there: <https://metr.org/blog/2026-1-29-time-horizon-1-1/>.
- The Claude 3.5 Sonnet report explains differences in setup and internet access and argues for a representative expert pool rather than the shortest observed individual time: <https://metr.org/evaluations/claude-3-5-sonnet-report/>.

**Implication**

- If Skill Issue uses human calibration, it should publish the comparator definition, recruitment and expertise criteria, allowed tools, number of attempts, success filtering, and which values are measured versus estimated.
- Estimated and measured baselines should never be visually interchangeable. Sensitivity analyses should show the effect of replacing estimates or changing aggregation.
- Human calibration is lower fit when Skill Issue's purpose is relative harness quality rather than human-equivalent autonomy.

### Elicitation is measured separately from held-out evaluation

**Fact.** METR uses development tasks to improve general scaffolding and prompting, then evaluates a separate test set. It checks whether token budgets are binding and investigates scaffold effects rather than assuming a default setup fully elicits capability. The 2024 report manually patched predefined spurious failure modes, with independent annotation, to estimate how much performance might improve under better task-agnostic elicitation. Historical work acknowledges that developing a browsing tool on evaluation-relevant tasks risks overfitting and states that a train/test split is the ideal methodology.

**Inference.** Poor tool use, premature submission, context loss, repetition, and early stopping can be properties of the evaluated system or correctable elicitation failures. The distinction requires an explicit development protocol; task-specific rescue during a scored attempt would contaminate the result.

**Evidence**

- Current operations describe elicitation on development tasks before test-set evaluation: <https://metr.org/time-horizons/>.
- The Claude Code/Codex note tests alternative scaffolds and larger token budgets while discussing off-distribution behavior and API limits: <https://metr.org/notes/2026-02-13-measuring-time-horizon-using-claude-code-and-codex/>.
- The Claude 3.5 Sonnet report classifies failures by likely capability versus elicitation cause and tests task-agnostic interventions on an annotated sample: <https://metr.org/evaluations/claude-3-5-sonnet-report/>.
- The 2023 report discusses tool-development overfitting and the need for distinct development and final-evaluation tasks: <https://metr.org/Evaluating_LMAs_Realistic_Tasks.pdf>.

**Implication**

- Skill Issue should record an elicitation protocol before measured runs: allowed general prompt changes, harness fixes, budget checks, stopping rule, development task set, and freeze point.
- Test-task-specific prompt or tool changes should create a new experimental condition rather than overwrite the original result.
- A small “known spurious failure” study can be informative, but manual rescues should be reported as a sensitivity analysis, not merged into unassisted performance.

### Isolation and instrumentation are part of experimental validity

**Fact.** METR's task contract defaults to constrained networking and separates the non-root agent from privileged scoring and task metadata. Vivaria builds a task image, runs the agent, captures completions and actions as trace entries, preserves submissions, and scores inside the task container. Hawk runs evaluations in isolated Kubernetes pods, places sandboxes in separate pods, controls egress, applies resource limits, stores logs/results, supports transcript search and review, records score edits and invalidations, and exposes an audit trail. The Task Standard bridge can replay prior runs but publicly documents feature gaps such as unsupported auxiliary VMs and incomplete intermediate-score handling.

**Inference.** Reproducibility and security are coupled: an evaluation cannot be interpreted confidently when the agent can inspect scoring secrets, modify privileged state, share contamination across runs, or use undocumented network resources.

**Evidence**

- The Task Standard defines agent privileges, task metadata separation, default network denial, scoring location, and custom-driver equivalence: <https://github.com/METR/task-standard/blob/main/STANDARD.md>.
- Vivaria's architecture documents task builds, agent execution, traces, submissions, scoring, and its PostgreSQL-backed review surface: <https://vivaria.metr.org/architecture/>.
- Hawk documents pod isolation, proxying, persistence, search, review, invalidation, and resumability: <https://hawk.metr.org/>.
- Hawk's security guide details identity controls, namespaces, Cilium policies, resource limits, audit logging, and protected model identities: <https://hawk.metr.org/infrastructure/security/>.
- The Inspect bridge documents replay support and current compatibility limitations: <https://github.com/METR/inspect-metr-task-bridge/>.

**Implication**

- Skill Issue should isolate each attempt, define writable and read-only surfaces, keep scorer material inaccessible, control network access, and persist a machine-readable event stream.
- Score edits and invalidations should be append-only reviewed events containing actor, reason, prior value, new value, and supporting evidence.
- Hawk's full Kubernetes and cloud architecture is lower fit for a local CLI until workload, threat model, or multi-user operation justifies it. The transferable requirements are isolation boundaries, explicit egress, resource accounting, replayable logs, and auditable review.

### Reproducibility is layered rather than absolute

**Fact.** METR publishes a task specification, example tasks, analysis code and data, infrastructure code, and report-specific artifacts. The analysis repository uses DVC to connect raw JSONL, parameters, model fitting, plots, and metrics. At the same time, public tasks intentionally omit protected assets and solutions to limit contamination, the paper states that not all operational code can be anonymized for release, and current ports disclose work-in-progress score regressions and protected-file risks. The Task Standard itself warns that interfaces may still change before 1.0.

**Inference.** A credible reproducibility statement should enumerate what can be rerun, what can only be audited, and what remains protected. “Open source” is not a binary property of the whole evaluation.

**Evidence**

- Public task examples describe omitted protected material, DVC-managed assets, and restrictions on publishing solutions: <https://github.com/METR/public-tasks>.
- The analysis pipeline exposes data fields, parameters, stages, and `dvc repro`: <https://github.com/METR/eval-analysis-public>.
- The time-horizon paper's reproducibility appendix distinguishes public core artifacts from non-public operational code: <https://arxiv.org/pdf/2503.14499>.
- The Inspect task port labels itself work in progress and records lower scores, insufficient elicitation, protected-file concerns, and static GPU allocation: <https://github.com/METR/inspect-tasks-public>.
- The Task Standard bridge records unsupported or incomplete features rather than implying full equivalence: <https://github.com/METR/inspect-metr-task-bridge/>.

**Implication**

- Skill Issue should publish an artifact matrix per evaluation: task source or checksum, protected inputs, subject manifest, environment image or build recipe, raw attempts, scoring code, aggregation parameters, analysis output, and known omissions.
- Every retained artifact should have a content hash or version identifier and an explicit relationship to the headline report.
- Protected tasks can remain auditable through checksums, interface specifications, reviewer records, and post-retirement release; publishing solutions during active evaluation would undermine the evaluation's purpose.

### Limitations and sensitivity analyses accompany headline claims

**Fact.** METR repeatedly limits interpretation to the measured task distribution: mostly self-contained, automatically scorable software, cyber, and ML tasks performed by low-context humans and solitary agents. The time-horizon paper lists factors that make work messier, including irreversible errors, dynamic environments, coordination, hidden scoring, information seeking, and hard-to-verify outcomes. Time Horizon 1.1 reports task removals and updates caused by confusion, reward hacks, and scoring errors, and acknowledges that loosely governed suite composition shifted the distribution. A recent GPT-5.6 Sol report shows that alternative treatments of suspected cheating yield radically different horizon estimates and declines to claim a robust number.

**Inference.** The strongest methodological pattern is not the specific metric but METR's practice of surfacing construct boundaries, post-hoc corrections, and analysis choices near the claim they affect.

**Evidence**

- The current time-horizon FAQ distinguishes task-completion horizon from wall-clock autonomous operation and limits generalization by domain, context, and task cleanliness: <https://metr.org/time-horizons/>.
- The paper analyzes external validity and a 16-factor “messiness” taxonomy while cautioning that the factors were selected and correlated with task time: <https://arxiv.org/pdf/2503.14499>.
- Time Horizon 1.1 explains task changes, uncertainty, sparse long-task baselines, infrastructure migration, and distribution shift: <https://metr.org/blog/2026-1-29-time-horizon-1-1/>.
- The GPT-5.6 Sol report presents incompatible estimates under alternative suspected-cheating treatments instead of choosing one as robust: <https://metr.org/blog/2026-06-26-gpt-5-6-sol/>.

**Implication**

- Skill Issue should place construct scope, known suite bias, task changes, invalidation rules, and sensitivity results beside the headline metric.
- Corrections should be durable changelog entries linking prior and revised artifacts; silently replacing historical outputs would break auditability.
- A full “messiness” taxonomy is conditional. A compact local taxonomy—realism, hidden information, environment dynamism, irreversible consequences, external coordination, and score verifiability—may be enough unless Skill Issue studies real-world autonomy directly.

### METR's presentation uses a layered evidence hierarchy

**Fact.** METR does not force one page to serve every reader. The live time-horizon page opens with a descriptive title, update date, compact definition, and interactive chart; methods, interpretation, FAQ, and changes follow. Release notes explain deltas and link data. Evaluation reports use a summary, methodology, quantitative results, transcript-grounded qualitative observations, limitations, footnotes, and a ready-to-copy citation. Academic papers use numbered sections, equations, figure/table captions, and extensive appendices. Task-development documentation uses a persistent sidebar, shallow headings, bullets, code, callouts, and previous/next navigation. Resource hubs use short “More details,” strengths, limitations, and future-work blocks. GitHub repositories carry executable specifications, examples, manifests, diagrams, and pipelines.

**Inference.** The reading flow works because each surface has one job: orient, specify, report a result, explain a change, or enable reproduction. Evidence links travel downward from claims to methods to artifacts.

**Evidence**

- Live result and methodology hierarchy, chart controls, FAQ, and update history: <https://metr.org/time-horizons/>.
- Release-note structure and linked comparison data: <https://metr.org/blog/2026-1-29-time-horizon-1-1/>.
- Evaluation-report structure, compact task tables, bootstrapped charts, transcript excerpts, limitations, footnotes, and citation block: <https://metr.org/evaluations/claude-3-5-sonnet-report/>.
- Expandable resource summaries organized by details, strengths, limitations, and future work: <https://metr.org/blog/2024-03-13-autonomy-evaluation-resources/>.
- Sidebar documentation flow: <https://taskdev.metr.org/introduction/>.
- Sequence and architecture diagrams in technical artifacts: <https://github.com/METR/task-standard> and <https://vivaria.metr.org/architecture/>.

**Implication**

- Skill Issue's compact public hierarchy should be: a stable methodology page; versioned task and artifact specifications; one report per evaluation or campaign; release/change notes for methodology revisions; and direct repository links for executable evidence.
- A result page should lead with subject, suite, headline result, uncertainty, and scope warning, then offer methods and attempt-level evidence. Tables fit task catalogs and condition comparisons; diagrams fit control flow and artifact lineage; transcript excerpts fit qualitative failure evidence.
- Expandable details are useful for secondary mechanics and caveats once a concise summary remains visible. Core definitions, subject identity, and uncertainty should stay in the main reading path.
- The transferable design is semantic hierarchy and evidence proximity. Skill Issue should use its own typography, color, illustration, naming, and prose voice.

### Recommended adoption tiers for Skill Issue

**Fact.** Across METR's materials, the most stable recurring controls are explicit subject configuration, versioned tasks, environment separation, development/test distinction, repeated attempts, retained trajectories, scorer validation, manual anomaly review, transparent aggregation, and limitations near results. Human time horizons, high-cost baseline programs, Kubernetes orchestration, and risk-specific task taxonomies serve narrower research goals.

**Inference.** Skill Issue can adopt the experimental spine without inheriting METR's research construct or operating scale.

**Evidence**

- Recurring controls are documented across the current methodology, paper, HCAST, Task Standard, development guide, public analysis pipeline, and infrastructure sources listed above.
- Infrastructure and measurement-specific constraints are explicitly scoped in the Task Standard bridge, task-port repository, time-horizon FAQ, and paper limitations.

**Implication**

- **High fit:** immutable subject and task manifests; author/QA/evaluation lifecycle; development/test separation; isolated attempts; multiple stochastic runs; raw score plus versioned decision rules; append-only transcripts and reviews; invalidation taxonomy; artifact hashes; recomputable aggregation; uncertainty; scope and changelog near claims.
- **Conditional fit:** task families and family weighting; human baselines; continuous-to-binary thresholds; transcript failure taxonomies; manual holistic scoring; protected task assets; sensitivity analyses for exploit handling. Adopt each only when it serves an explicit construct and has a review protocol.
- **Lower fit:** human-equivalent time horizons, HCAST's long-task target durations, full Hawk-style cluster infrastructure, safety-risk task selection, and METR's complete messiness taxonomy. These solve goals beyond a typical local skill-evaluation CLI.
- **Documentation fit:** use a short public methodology, versioned technical specifications, semantically named campaign reports, direct artifact links, compact tables, flow diagrams, qualitative examples, and explicit limitations. Preserve Skill Issue's own vocabulary and visual system.

## Notes

- METR is transitioning from Vivaria to Inspect/Hawk. Sources from both generations were included because the older documentation explains the trace and scoring model while current repositories explain the migration boundary. Bridge limitations mean the systems should not be treated as feature-equivalent.
- The Task Standard specification displayed version 0.5.0 on `main`, while the repository's visible GitHub release widget showed an older release. This assignment cites the inspected specification version and does not infer release-channel equivalence.
- Public task repositories intentionally omit protected assets and solutions. No claims were made about those unavailable materials.
- Some current METR pages and papers post-date the system date of many common secondary summaries. Official first-party pages and repositories were therefore preferred throughout.
- The 2023 realistic-task report is retained as historical evidence for human oversight, manual simulation, scaffolding overfitting, and train/test reasoning. Later Task Standard, HCAST, time-horizon, Inspect, and Hawk sources are more representative of current operational practice.

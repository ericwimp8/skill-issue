# External Research: Evaluation Methodology Page

## Recommended Direction

Name the destination **Evaluation Methodology**.

This is the best-supported audience-facing name because the destination needs to own more than operational testing. It must explain the evaluated system, the three governed conversational scenarios, turn and invocation expectations, controlled CLI execution, instrumentation, scoring, human work, environment controls, evidence retention, reproducibility, uncertainty, and limitations. Public sources consistently use “evaluation methodology” or a contextual “methodology” for that end-to-end account; “testing methodology” usually narrows toward run procedure, while “scoring methodology” narrows toward metrics and aggregation. “Benchmark card,” “system card,” and “technical report” make different promises and would misstate the page's role. [Page naming scan](assignments/04-page-naming-interface-scan.md) · [NIST automated benchmark evaluation](https://www.nist.gov/publications/towards-standard-automated-benchmarking-artificial-intelligence) · [Artificial Analysis methodology](https://artificialanalysis.ai/methodology/coding-agents-benchmarking)

Build one **navigational owner**, not one overloaded artifact. The methodology page should be understandable by itself, then lead directly to the three scenario records, attempt-level evidence, scorecards, canonical GitHub material, and reproduction instructions. SWE-bench, HELM, METR, OpenAI, and SWE-Together all distribute reader jobs across an orienting explanation, result or scenario views, executable specifications, retained evidence, and version history. Their common lesson is progressive disclosure with explicit ownership, rather than duplicating the same methodology across several independently maintained surfaces. [SWE-bench deep dive](assignments/05-swe-bench-methodology.md) · [HELM deep dive](assignments/07-helm-benchmark-documentation.md) · [METR deep dive](assignments/06-metr-evaluation-methodology.md) · [OpenAI disclosure deep dive](assignments/08-openai-evaluation-disclosures.md) · [SWE-Together deep dive](assignments/10-swe-together-conversational-benchmark.md)

The page should make one bounded claim: it explains how Skill Issue evaluates governed conversational skill calling under the disclosed scenarios and conditions. It should not imply that three scenarios represent all skill-calling behavior, that a passing result belongs to a bare model rather than a configured system, or that inspectable artifacts prove independent reproduction. SWE-Together's 109 tasks retain less than one percent of its source sessions and still require a coverage caveat; SWE-bench and OpenAI's later audits show that apparently validated tasks can later prove underspecified, contaminated, or poorly graded. A three-scenario system therefore needs particularly explicit selection rationale and scope limits. [SWE-Together paper](https://arxiv.org/pdf/2606.29957) · [OpenAI SWE-bench Verified](https://openai.com/index/introducing-swe-bench-verified/) · [OpenAI benchmark retirement](https://openai.com/index/why-we-no-longer-evaluate-swe-bench-verified/) · [OpenAI coding-evaluation audit](https://openai.com/index/separating-signal-from-noise-coding-evaluations/)

The rest of this document uses three evidence labels:

- **Direct source fact** reports what a cited primary source publishes or implements.
- **Methodological implication** states a bounded lesson supported across those facts.
- **Skill Issue design recommendation** proposes a project-specific choice. It is not a vendor requirement or a claim about current local implementation.

### Recommended Complete Reading Flow

1. **Title, one-sentence deck, and evidence actions.** State what is evaluated and why the page exists. Show the current methodology version, status, last-updated date, and direct links such as **Browse scenarios**, **Inspect evidence**, **Reproduce**, and **GitHub**. Short evidence actions near the title are common across HELM, SWE-bench, ARC, and SWE-Together; they give readers an immediate verification route without promotional calls to action. [Page naming and interface scan](assignments/04-page-naming-interface-scan.md) · [HELM index](https://crfm.stanford.edu/helm/) · [SWE-Together overview](https://togetherbench.com/)
2. **What the evaluation establishes.** Define the governed question, evaluated unit, three-scenario scope, primary success condition, and claims the evaluation cannot support. Put the evaluated system boundary before scores, following SWE-bench, METR, and system-card practice. [SWE-bench deep dive](assignments/05-swe-bench-methodology.md) · [METR deep dive](assignments/06-metr-evaluation-methodology.md) · [System-card candidate scan](assignments/02-system-card-candidates.md)
3. **Scenario overview.** Present the three scenarios in one comparable matrix, using their canonical Skill Issue names when available. Each row should state purpose, initial condition, later-turn trigger, expected invocation behavior, expected downstream behavior, and principal exclusion. Do not invent labels such as “direct,” “proactive,” or “non-invocation” unless those are the project's authoritative terms; the external SWE-bench assignment used them only as an explicit mapping inference. [SWE-bench deep dive, Finding 9](assignments/05-swe-bench-methodology.md)
4. **Evaluation loop.** Use one restrained diagram and short prose to show scenario fixture → configured agent system → conversation and expected events → skill-call instrumentation → resulting behavior/artifact → scoring and review → retained evidence. HELM uses diagrams for lifecycle and taxonomy, while tables remain the comparison mechanism; SWE-Together's Design page similarly uses a small number of diagrams to establish state separation and runtime flow. [HELM deep dive, Findings 16–17](assignments/07-helm-benchmark-documentation.md) · [SWE-Together Design](https://togetherbench.com/design.html)
5. **Execution and controls.** Disclose the exact configured subject, clean-state procedure, context boundaries, tool and network policy, budgets, retry policy, concurrency, environment identity, and operational-invalid rules. Keep campaign-wide controls separate from scenario-specific exceptions. METR and HELM both make framework rules and task-specific adaptations separately inspectable. [METR Task Standard](https://github.com/METR/task-standard/blob/main/STANDARD.md) · [HELM Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html)
6. **Instrumentation and retained evidence.** Explain which invocation events, conversation messages, tool events, evaluator decisions, artifacts, timing, resource metadata, and errors are recorded. State any telemetry that the harness cannot expose. SWE-bench's artifact chain and HELM's run artifacts show how summary results can link to per-attempt evidence; AgentLens shows why final outcomes alone can hide process defects. [SWE-bench experiments](https://github.com/SWE-bench/experiments) · [HELM tutorial](https://crfm-helm.readthedocs.io/en/latest/tutorial/) · [AgentLens](https://www.microsoft.com/en-us/research/publication/agentlens-revealing-the-lucky-pass-problem-in-swe-agent-evaluation/)
7. **Scoring and interpretation.** Define criterion-level evidence, the primary conjunctive result, diagnostic outcomes, attempt semantics, aggregation, exclusions, and uncertainty. Include a worked example. NIST's key requirement is to name the quantity being estimated and the sampling assumptions; SWE-Together and METR show why raw attempt evidence and stability should remain visible beneath an aggregate. [NIST AI 800-3](https://www.nist.gov/publications/expanding-ai-evaluation-toolbox-statistical-models) · [SWE-Together paper](https://arxiv.org/pdf/2606.29957) · [METR time-horizon paper](https://arxiv.org/pdf/2503.14499)
8. **Manual work and quality assurance.** Identify authoring, independent scenario QA, measured execution, transcript review, score adjudication, and correction as separate stages. State reviewer qualifications, information access, independence, sample selection, disagreement handling, and confidence capture. OpenAI and Anthropic treat manual work as methodology rather than cleanup. [OpenAI disclosure deep dive, Findings 5–6](assignments/08-openai-evaluation-disclosures.md) · [Claude 4 deep dive, Findings 3 and 5](assignments/09-anthropic-claude-4-system-card.md)
9. **Scenario evidence.** Link three comparable scenario cards or detail sections. Each should expose the messages and trigger conditions, expected calls, relevant skill identity, scorecard, attempt matrix, annotated transcript, evaluator evidence, environment, and canonical GitHub artifacts. With only three scenarios, a search-heavy marketplace browser is unnecessary; SWE-Together's card/filter interface solves a 109-task scale problem. [SWE-Together deep dive, Findings 12–13 and 16](assignments/10-swe-together-conversational-benchmark.md)
10. **Reproducibility and evidence availability.** Publish an artifact matrix that distinguishes what is available, what can be rerun, what has been independently reproduced, and what is withheld or unavailable. Give the smallest proof command before the full campaign path. ACM, HELM, MLPerf, and SV-COMP all separate artifact availability from successful reproduction and bind results to versioned packages. [Reproducibility candidate scan](assignments/03-reproducibility-candidates.md) · [ACM artifact policy](https://www.acm.org/publications/policies/artifact-review-and-badging-current) · [HELM reproduction guide](https://crfm-helm.readthedocs.io/en/latest/reproducing_leaderboards/) · [SV-COMP reproduction archive](https://sv-comp.sosy-lab.org/reproduce.php)
11. **Limitations, uncertainty, and validity status.** Keep material caveats expanded. Organize them by the claim they constrain: scenario coverage, configured-system dependence, grader validity, run variance, platform dependence, reviewer disagreement, contamination or overfitting, and missing telemetry. Assign the scenario set a current validity status and explain conditions for weakening, retiring, or superseding it. [Claude 4 deep dive, Finding 8](assignments/09-anthropic-claude-4-system-card.md) · [OpenAI disclosure deep dive, Findings 11–12](assignments/08-openai-evaluation-disclosures.md)
12. **Version history and citation.** End with a dated change log that names affected campaigns and comparability consequences, plus a stable citation or version reference. Artificial Analysis, HELM, METR, OpenAI, and Anthropic all demonstrate that methods and results drift; the strongest sources preserve corrections and superseded conclusions rather than silently replacing them. [Page naming and interface scan, Finding 13](assignments/04-page-naming-interface-scan.md) · [OpenAI disclosure deep dive, Finding 11](assignments/08-openai-evaluation-disclosures.md) · [Claude 4 deep dive, Finding 10](assignments/09-anthropic-claude-4-system-card.md)

## Methodology Content Contract

The following is a synthesis recommendation. It adapts convergent patterns from the supplied external research; it does not assert that Skill Issue already records or implements these fields. Tables and schemas in this section are **Skill Issue design recommendations** unless a paragraph explicitly labels them as direct source facts.

### Purpose and Scope

Open with a direct description of the evaluation's purpose, the governed decision it informs, and the population it does not represent. The page should identify the three scenarios as a deliberately fixed evaluation set and explain why each earns inclusion. It should also name omitted behaviors, such as other skill families, longer conversations, different harnesses, or other environment conditions, only when those omissions are known from the product's authoritative material.

This follows HELM's practice of defining the broader evaluation space and naming incompleteness, METR's practice of limiting a claim to the measured task distribution, and OpenAI's practice of stating whether a page launches, audits, or retires a benchmark. [HELM methodology](https://crfm.stanford.edu/2022/11/17/helm.html) · [METR time horizons](https://metr.org/time-horizons/) · [OpenAI disclosure deep dive, Findings 2 and 15](assignments/08-openai-evaluation-disclosures.md)

The page should identify a concrete decision owner only if Skill Issue has one. Anthropic's strongest system-card pattern is a decision record in which evidence, critique, threshold, and disposition are connected. For Skill Issue, the transferable principle is to state what a result changes—publication, refinement, pause, or no action—without importing Anthropic's safety-level terminology. [Claude 4 deep dive, Findings 1 and 14](assignments/09-anthropic-claude-4-system-card.md)

### Evaluated Subject

**Skill Issue design recommendation.** Define the subject as the complete configured system rather than a bare model name. The following complete run-manifest field set is a local design proposal, assembled from recurring disclosure dimensions across the sources:

| Dimension             | Recommended disclosure                                                            | Why it matters                                      |
| --------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------- |
| Model route           | Provider, immutable or date-qualified model ID, reasoning mode, sampling settings | Provider aliases and modes can change behavior      |
| Harness               | CLI/executable identity, version or commit, scaffold, prompt envelope             | Agent behavior depends on the operating scaffold    |
| Skill state           | Installed registry, target skill identity and revision/hash, invocation policy    | The evaluation concerns governed skill calling      |
| Conversation          | Scenario ID/version, initial state, supplied history, turn/step budget            | Later behavior depends on accumulated context       |
| Tools and permissions | Available tools, sandbox, writable/read-only surfaces, network state              | Capability and isolation depend on access           |
| Runtime               | Timeout, token or cost budget, retry policy, concurrency                          | Reliability and efficiency depend on limits         |
| Evaluation            | Scorer/rubric version, judge identity if any, human-review policy                 | The reported result depends on interpretation rules |

**Direct source facts.** METR treats model plus scaffold, tools, prompt, budget, and environment as the agent under evaluation. SWE-bench distinguishes full agent systems from controlled scaffold comparisons. OpenAI and Anthropic system cards qualify results by checkpoint, prompt, scaffold, reasoning mode, tool access, attempt policy, or safeguards. None publishes the complete Skill Issue manifest above as one required contract. [METR deep dive, “The evaluated subject”](assignments/06-metr-evaluation-methodology.md) · [SWE-bench deep dive, Finding 2](assignments/05-swe-bench-methodology.md) · [System-card candidate scan](assignments/02-system-card-candidates.md)

Use a short display name only as an alias for the full manifest. When configurations differ materially, block direct comparison or put the difference beside the score. HELM's separation of model identifier, package release, leaderboard release, and run suite is the strongest example of version domains that should not be collapsed. [HELM deep dive, Findings 9–10](assignments/07-helm-benchmark-documentation.md)

### Scenarios, Turns, and Expected Events

**Skill Issue design recommendation.** Give each scenario one canonical record with the following shared schema. This is the proposed local governance record, not a schema required by SWE-Together or another source:

- Stable scenario ID, display name, version, status, and selection rationale.
- Governing skill-calling question and known coverage boundary.
- Initial user message and any prior conversation or workspace fixture visible to the agent.
- Later human-authored messages or deterministic trigger conditions, in order.
- Expected invocation, non-invocation, clarification, or deferral event at the relevant turn.
- Expected skill identity and allowed equivalent routing, if equivalence exists.
- Expected downstream behavior or artifact after the decision.
- Explicitly valid alternative outcomes and invalid shortcuts.
- Termination conditions, budgets, and operational-invalid conditions.
- Criterion-level scorecard and links to retained attempt evidence.

**Direct source fact.** SWE-Together uses intent anchors and trajectory-dependent simulator actions, with summarization and model mediation in the loop. **Methodological implication.** Fixed replay improves control but can mistime later messages after trajectories diverge; reactive replay improves semantic timing but introduces another evaluated component. **Skill Issue design recommendation.** For three governed scenarios, prefer human-authored messages and deterministic trigger rules unless adaptive user behavior is necessary. That preference is local synthesis, not a SWE-Together requirement. [SWE-Together deep dive, Findings 2 and 8](assignments/10-swe-together-conversational-benchmark.md)

If adaptive simulation is later required, retain the simulator's input summary, decision, injected message, prompt/policy version, model configuration, and failure classification separately from the evaluated agent's transcript. Two SWE-Together findings must remain separate. Its forced-choice study found that annotators did not distinguish simulator trajectories from real-user trajectories at a statistically reliable rate; that is a trajectory-level indistinguishability result. Its separate Intent Coverage metric reports similar aggregate coverage across model cohorts; that is cross-cohort intent coverage, not evidence of human equivalence or correct intervention timing. [SWE-Together paper, pp. 9–10](https://arxiv.org/pdf/2606.29957) · [Intent Coverage implementation](https://raw.githubusercontent.com/Togetherbench/SWE-Together/main/eval/user_behavior/coverage_one.py)

### Controlled Execution and Environment Isolation

The page should describe isolation as a set of explicit controls, not as a synonym for “containerized.” SWE-bench documents architecture sensitivity and version drift even after moving to Docker; MLPerf and SV-COMP go further by fixing hardware/software identity, resource limits, seeds, timing boundaries, and cache or restart conditions. [SWE-bench container report](https://github.com/SWE-bench/SWE-bench/blob/main/docs/20240627_docker/README.md) · [MLPerf Training Rules](https://github.com/mlcommons/training_policies/blob/master/training_rules.adoc) · [SV-COMP rules](https://sv-comp.sosy-lab.org/2025/rules.php)

Disclose the controls that matter to conversational skill calling:

1. Start each attempt from a named clean conversation and workspace state.
2. Pin the skill registry, target skill revision, harness, prompts, evaluator, and scenario revision.
3. Define ambient instruction handling, prior-turn visibility, enabled tools, filesystem permissions, and network access.
4. Keep evaluator secrets or protected criteria outside the agent-visible surface where hidden checks are justified.
5. Record time, token, cost, step, retry, and concurrency policies.
6. Separate operational invalidation from behavioral failure.
7. Run a cheap known-answer or calibration path before accepting a campaign when the product has such a fixture.
8. Treat intentional multi-turn carryover as a versioned fixture, not ambient residue.

**Direct source facts.** METR's current Hawk documentation describes isolated Kubernetes pods, controlled egress and resources, stored logs and results, transcript search and review, score edits, invalidations, and audit logging. **Methodological implication.** Trace capture and later evaluator changes are part of experimental validity when they can affect an accepted result. **Skill Issue design recommendation.** Adopt explicit isolation boundaries and an auditable review trail without copying Hawk's cluster architecture. [METR Hawk](https://hawk.metr.org/) · [Hawk security model](https://hawk.metr.org/infrastructure/security/) · [Hawk repository](https://github.com/METR/hawk)

### Invocation Instrumentation and Evidence

The primary evidence for a skill invocation should be an observable event emitted by the controlled harness. A final answer that resembles a skill's behavior is not proof that the skill was called. Conversely, the absence of a recorded event can establish non-invocation only when event capture is known to be complete for that attempt.

**Skill Issue design recommendation.** Retain the following per-attempt evidence inventory. This complete field set is local synthesis; the cited projects supply representative parts rather than one shared contract:

| Evidence class | Minimum retained material                                                                 |
| -------------- | ----------------------------------------------------------------------------------------- |
| Identity       | Campaign, attempt, scenario, configured-subject, skill, harness, and evaluator IDs        |
| Conversation   | User and assistant messages with turn ordering and timestamps where available             |
| Invocation     | Skill resolution, call event, call timing, arguments or safe digest, result status        |
| Tools          | Observable tool events, outputs or references, failures, and policy interruptions         |
| Result         | Final response and produced artifacts or content-addressed references                     |
| Evaluation     | Per-criterion evidence, raw score, decision rules, derived outcome, reviewer annotations  |
| Operations     | Duration, tokens/cost where available, termination reason, retries, infrastructure errors |
| Provenance     | Scenario version, workspace/source revision, environment identity, artifact hashes        |

**Direct source facts.** SWE-bench publishes predictions, patches, reports, test output, logs, metadata, and optional inference trajectories. Its leaderboard submission rule asks for a human-readable trajectory but loosely specifies the trace format and problem-solving content. Large logs and trajectories are stored in public S3-compatible object storage and require separate download tooling, creating access friction. HELM retains run specifications, scenario state, per-instance statistics, and aggregate statistics; EleutherAI's harness records prompts or chat templates, seeds, task configuration, model arguments, hardware counts, and timestamps. **Methodological implication.** These are complementary artifact examples, not a HELM-owned evidence-field inventory. [SWE-bench experiments](https://github.com/SWE-bench/experiments) · [SWE-bench deep dive, Finding 6](assignments/05-swe-bench-methodology.md) · [HELM deep dive, Finding 12](assignments/07-helm-benchmark-documentation.md) · [EleutherAI evaluation harness](https://github.com/EleutherAI/lm-evaluation-harness)

Do not require hidden chain-of-thought. Observable messages, calls, events, artifacts, and evaluator evidence are sufficient and more consistently available. SWE-bench's free-form reasoning trajectories are a lower-fit analogue; AgentLens's core contribution is the value of process evidence, not a requirement to expose private internal reasoning. [SWE-bench deep dive, Finding 6](assignments/05-swe-bench-methodology.md) · [AgentLens repository](https://github.com/microsoft/code-agent-state-trajectories/)

### Scoring and Aggregation

**Skill Issue design recommendation.** Use the following five-part governed-success rule with separate diagnostic dimensions. This conjunction is local synthesis from several scoring patterns; no cited benchmark publishes this Skill Issue-specific rule:

1. **Decision correctness:** the system invokes, does not invoke, clarifies, or defers as the scenario requires.
2. **Timing and identity correctness:** the event occurs at the permitted turn and resolves to the correct skill or explicitly allowed equivalent.
3. **Contract adherence:** the invoked skill's governed procedure and authority boundary are followed.
4. **Downstream correctness:** the final response or artifact satisfies the scenario's task criteria.
5. **Evidence validity:** required instrumentation is present and the attempt is not operationally invalid.

Report diagnostic outcomes separately: decision-only success, downstream-only success, wrong skill, early or late invocation, excessive correction, malformed or missing telemetry, timeout, harness failure, evaluator failure, and invalidated attempt. SWE-bench's conjunction of fail-to-pass and pass-to-pass tests, plus separate operational outcomes, is the closest scoring analogue. SWE-Together adds the essential separation between final correctness and corrective steering. [SWE-bench deep dive, Finding 5](assignments/05-swe-bench-methodology.md) · [SWE-Together deep dive, Finding 3](assignments/10-swe-together-conversational-benchmark.md)

**Methodological implication.** Deterministic checks are strongest when they directly measure the contract; semantic judging needs its own validation against valid alternatives and shortcuts. **Skill Issue design recommendation.** If an LLM judge is used, disclose its identity, prompt/rubric version, evidence access, sampling settings, failures, and aggregation, then add a bounded human audit or alternate-judge sensitivity check before interpreting small differences. The alternate-judge check is a project recommendation rather than a requirement published by PaperBench, HELM, SWE-Together, or OpenAI. [PaperBench](https://openai.com/index/paperbench/) · [HELM deep dive, Finding 7](assignments/07-helm-benchmark-documentation.md) · [OpenAI disclosure deep dive, Finding 6](assignments/08-openai-evaluation-disclosures.md)

Keep raw criterion outcomes before pass thresholds, and version threshold and aggregation rules separately. A reference trace or expected artifact should be labeled **reference outcome**, not “oracle” or “ceiling,” unless maximality is proven. SWE-Together's public page calls its patch a ceiling while its paper explains why it is not one; some rubric criteria concern conversational process that a final patch cannot express. [SWE-Together deep dive, Finding 7](assignments/10-swe-together-conversational-benchmark.md)

### Repeated Attempts and Uncertainty

**Direct source fact.** SWE-Together defines `pass²` for its two-replicate design as the share of tasks where **both of the two runs pass**. METR also uses repeated trajectories in its evaluation work. **Skill Issue design recommendation.** Run more than one attempt per scenario, publish the complete attempt matrix, and describe a two-attempt result only as **two-run stability**. The caution that two trials are too few for a strong tail-reliability claim is local synthesis, not SWE-Together's stated validity threshold. [SWE-Together paper, pp. 7–8](https://arxiv.org/pdf/2606.29957) · [SWE-Together deep dive, Finding 5](assignments/10-swe-together-conversational-benchmark.md) · [METR deep dive, “Repeated trajectories”](assignments/06-metr-evaluation-methodology.md)

**Skill Issue design recommendation.** With only three scenarios, prefer a raw scenario-by-attempt matrix to an elaborate confidence interval around a headline average. This raw-matrix preference is a local presentation judgment. Any aggregate should state:

- Number of scenarios and attempts per scenario.
- Sampling unit and whether attempts are independent.
- Pass definition and threshold version.
- Weighting across scenarios and skills.
- Excluded or invalid attempts and their reasons.
- Variability measure appropriate to the sample.
- Configuration changes that prevent direct comparison.

If the scenarios are not meaningfully commensurate, prioritize dimensional results and a macro view rather than a single ranking. NIST shows that fixed-item accuracy and generalization to a broader item population are different estimands. HELM exposes plural metrics and transparent aggregation choices, but its visible result interface is principally point-estimate oriented rather than an exemplar of run-variance or interval presentation. [NIST AI 800-3](https://www.nist.gov/publications/expanding-ai-evaluation-toolbox-statistical-models) · [HELM deep dive, Findings 8 and 19](assignments/07-helm-benchmark-documentation.md)

### Manual Steps and Review

Publish manual work as a first-class part of the method. Separate:

- Scenario authoring and expected-behavior definition.
- Independent scenario and scorer QA.
- Pilot or elicitation work before the measured freeze.
- Any explicit manual action required during an otherwise controlled run.
- Automated triage of likely failures.
- Human transcript review and score adjudication.
- Post-publication correction, invalidation, or retirement.

**Direct source facts.** SWE-bench Verified publishes reviewer counts, onboarding, evidence boundaries, three independent labels, conservative aggregation, and confidence capture. Later OpenAI coding-evaluation audits add explicit escalation for disagreements or low-confidence cases. Anthropic combines automated screening with policy-specific human review and treats transcript inspection as a separate evidence class. **Skill Issue design recommendation.** For each manual stage, disclose the role, expertise, information access, review count, calibration, label aggregation, disagreement process, confidence treatment, and whether reviewers are independent of scenario or skill authors. Reviewer independence is a local governance recommendation, not an OpenAI or Anthropic requirement. [OpenAI SWE-bench Verified](https://openai.com/index/introducing-swe-bench-verified/) · [OpenAI annotation instructions](https://cdn.openai.com/introducing-swe-bench-verified/swe-b-annotation-instructions.pdf) · [OpenAI coding-evaluation audit](https://openai.com/index/separating-signal-from-noise-coding-evaluations/) · [Claude 4 deep dive, Findings 3 and 5](assignments/09-anthropic-claude-4-system-card.md)

Do not merge manual rescue into unassisted performance. METR's elicitation work is performed on a development set and frozen before held-out evaluation; task-specific interventions are reported as separate sensitivity analyses. This is the right boundary for skill hardening: refinement can use evaluation evidence, but a changed skill, prompt, harness, or scorer creates a new condition and new result identity. [METR deep dive, “Elicitation is measured separately”](assignments/06-metr-evaluation-methodology.md)

### Reproducibility and Retained Source Material

**Skill Issue design recommendation.** Describe reproducibility with the following reader-facing tiers:

1. **Inspectable:** the public can examine the method, scenario contract, results, and retained evidence.
2. **Rerunnable:** the public artifacts and prerequisites are sufficient to execute the disclosed procedure, subject to named external dependencies.
3. **Reproduced:** an independent party has executed the method and reached the agreed evidence tolerance.

The labels **Inspectable**, **Rerunnable**, and **Reproduced** are this report's synthesis of ACM and NISO distinctions; they are not the formal badge names used by either organization. Use ACM or NISO terminology precisely if formal badges are claimed. Do not say “reproducible” merely because a repository or container exists. Proprietary model endpoints, mutable provider aliases, API behavior, environment architecture, protected material, credentials, and cost can all prevent exact reproduction. [ACM artifact policy](https://www.acm.org/publications/policies/artifact-review-and-badging-current) · [NISO RP-31-2021](https://www.niso.org/publications/rp-31-2021-badging) · [Reproducibility candidate scan, Notes](assignments/03-reproducibility-candidates.md)

Publish an artifact matrix like this, populated only with artifacts that actually exist:

| Artifact                    | Public status                   | Version/hash        | Purpose                                     | Reproduction role         | Known limitation                       |
| --------------------------- | ------------------------------- | ------------------- | ------------------------------------------- | ------------------------- | -------------------------------------- |
| Methodology page            | Public                          | Method version      | Human-readable contract                     | Interpretation            | May summarize deeper specs             |
| Scenario definitions        | Public or controlled            | Scenario-set digest | Inputs, turns, expected events              | Rerun fixture             | Hidden checks may be withheld          |
| Configured-subject manifest | Public                          | Campaign/run ID     | Pins model, harness, skills, tools, budgets | Comparison and rerun      | Provider snapshot may be unavailable   |
| Harness and scorer          | GitHub or controlled            | Commit/release      | Executes and grades                         | Rerun                     | Platform dependencies may remain       |
| Attempt bundles             | Public, redacted, or controlled | Content hashes      | Transcript, events, outputs, scores         | Audit and review          | Sensitive fields may require redaction |
| Aggregate results           | Public                          | Campaign release    | Derived metrics and tables                  | Recompute                 | Depends on threshold/weight versions   |
| Change history              | Public                          | Dated entries       | Corrections and comparability               | Historical interpretation | Requires durable old links             |

**Direct source facts.** HELM attaches citations to models, scenarios, runs, and code; SWE-bench links results to experiment folders; SV-COMP preserves versioned tool, task, result, witness, and orchestration archives. **Skill Issue design recommendation.** Put exact commit-, file-, schema-, scorecard-, scenario-, and evidence-index links beside consequential claims. This immutable-link granularity is a local design choice; the sources support proximal, versioned evidence without publishing one universal GitHub-link contract. [HELM deep dive, Finding 18](assignments/07-helm-benchmark-documentation.md) · [SWE-bench results viewer](https://www.swebench.com/viewer.html) · [SV-COMP reproduction archive](https://sv-comp.sosy-lab.org/reproduce.php)

### Limitations, Unsupported Evidence, and Validity State

Place limitations near the result or method they constrain, then consolidate them in a visible section. At minimum, assess:

- **Coverage:** three scenarios cannot establish population-wide skill-calling reliability.
- **Configured-system dependence:** results belong to the disclosed model/harness/skill/environment tuple.
- **Trajectory dependence:** later turns and corrections can alter what the scenario measures.
- **Instrumentation completeness:** missing invocation telemetry prevents a call-level claim.
- **Scorer validity:** deterministic checks may be incomplete; model or human judges may disagree.
- **Run variance:** stochastic agents may pass inconsistently.
- **Contamination and refinement pressure:** visible scenarios can become lexical or procedural targets rather than measure general intent recognition.
- **Environment dependence:** platform, provider, network, permissions, and mutable APIs may change results.
- **Manual-review limits:** reviewer sample, expertise, and access bound the conclusion.
- **External validity:** controlled conversational cases may not resemble messier production use.

**Direct source facts.** Anthropic weakens or retires evaluations after saturation, leakage, missing baselines, or unreliable grading; OpenAI's SWE-bench sequence shows a professionally reviewed benchmark later withdrawn and a proposed replacement recommendation retracted. **Skill Issue design recommendation.** Own one visible local validity taxonomy—such as active, weak-signal, saturated, contaminated, unreliable, retired, or superseded—and define its transition rules. The unified taxonomy and labels are project design, while the need to weaken or retire invalid evidence is source-backed. [Claude 4 deep dive, Finding 8](assignments/09-anthropic-claude-4-system-card.md) · [OpenAI benchmark retirement](https://openai.com/index/why-we-no-longer-evaluate-swe-bench-verified/) · [OpenAI coding-evaluation audit](https://openai.com/index/separating-signal-from-noise-coding-evaluations/)

**Direct source facts.** Anthropic's corrected Claude 4 table left nearby prose and a caption aligned to superseded values; SWE-bench's public surfaces contain count, namespace, and package-version drift. **Skill Issue design recommendation.** Derive summaries from one canonical result record, retain a durable versioned correction history, and validate every dependent statement when data changes. An immutable correction-history mechanism is a local design choice rather than a vendor requirement. [Claude 4 deep dive, Finding 10](assignments/09-anthropic-claude-4-system-card.md) · [SWE-bench deep dive, Finding 8](assignments/05-swe-bench-methodology.md)

## Page and Interface Design

### Information Architecture

**Skill Issue design recommendation.** Use a shallow hierarchy in which H2s own durable concepts and H3s appear only for separable mechanisms. The following outline is a project-specific page design, not a source-defined layout:

- `# Evaluation Methodology`
- `## What This Evaluation Establishes`
- `## The Three Scenarios`
- `## Evaluation Loop`
- `## Configured System and Controls`
- `## Instrumentation and Evidence`
- `## Scoring and Interpretation`
- `## Manual Review`
- `## Reproducibility`
- `## Limitations and Uncertainty`
- `## Methodology History`

This follows the cross-source orientation → mechanics → accountability pattern. Avoid paper-style numbered sub-subsections unless the page becomes a formal report; system-card PDFs and HELM research posts use deeper hierarchies for much broader documents. [Page naming and interface scan, Findings 4 and 15](assignments/04-page-naming-interface-scan.md)

Use a small inline or sticky contents list, anchored headings, and stable semantic URLs. The page should remain useful as static HTML on GitHub Pages; client-side behavior should enhance evidence inspection rather than carry essential meaning. Native `<details>`, responsive table wrappers, anchor links, and static generated scenario indexes cover most needs without a complex application shell. [Page naming and interface scan, Finding 14](assignments/04-page-naming-interface-scan.md)

### Prose Measure, Rhythm, and Punctuation

**Skill Issue design recommendation.** Start with body prose around **60–75 characters per line**, while allowing tables, diagrams, code, and evidence viewers to use a wider data surface. This range is a project layout proposal inferred from inspected pages, not a published vendor standard, and should be validated against the site's typography and viewport behavior. [Page naming and interface scan, Finding 6](assignments/04-page-naming-interface-scan.md) · [HELM deep dive, Finding 15](assignments/07-helm-benchmark-documentation.md)

Use one idea per paragraph, usually one to four sentences. Lead with a declarative claim, follow with the evidence or mechanism, then state the caveat or implication. Alternate short prose with stable evidence structures—lists, a compact table, a worked example, or a caption—rather than stacking long paragraphs.

Use punctuation functionally:

- Colons introduce criteria, examples, and enumerations.
- Parentheses carry abbreviations, pinned versions, and limited qualifications.
- Em dashes can clarify a compact definition but should not carry entire argument chains.
- Semicolons should be rare; split long causal claims into separate sentences.
- Avoid promotional superlatives, rhetorical questions, exclamation marks, and slogan fragments.

ARC, Artificial Analysis, SWE-bench, METR, and SWE-Together consistently use compact declarative paragraphs followed by evidence objects or examples. [Page naming and interface scan, Finding 7](assignments/04-page-naming-interface-scan.md) · [SWE-Together deep dive, Finding 11](assignments/10-swe-together-conversational-benchmark.md)

### Density and Spacing

**Skill Issue design recommendation.** Keep the hero compact: one H1, one-sentence deck, status/version metadata, and two to four evidence links. Reserve generous whitespace for the opening and major section transitions, then use compact spacing inside evidence groups. These counts and spacing choices are local layout proposals; the sources support restrained orientation and progressive disclosure rather than this exact composition.

Use small metadata labels for version, status, campaign, harness, and artifact identity, but keep core definitions in normal body text. Avoid forcing continuous prose into the width needed by score tables. HELM's public pages vary density by reader task; SWE-Together keeps sparse framing around its dense leaderboard; Artificial Analysis constrains its article column while allowing broader components. [HELM deep dive, Finding 15](assignments/07-helm-benchmark-documentation.md) · [SWE-Together deep dive, Finding 10](assignments/10-swe-together-conversational-benchmark.md) · [Page naming and interface scan, Finding 6](assignments/04-page-naming-interface-scan.md)

### Tables, Notes, and Diagrams

**Skill Issue design recommendation.** Use tables for fields with a stable schema:

- Three-scenario comparison.
- Configured-subject manifest.
- Attempt matrix and criterion outcomes.
- Score interpretation.
- Artifact availability.
- Method and campaign versions.

Explain each table's decision purpose in a sentence immediately before it. On narrow screens, allow horizontal scrolling or a focused column view rather than compressing text into unreadable cells. Procedures and caveats belong in prose or ordered lists, not oversized tables. [Page naming and interface scan, Finding 8](assignments/04-page-naming-interface-scan.md)

**Skill Issue design recommendation.** Use two note classes:

- **Visible validity notes** for non-comparability, missing telemetry, exclusions, known failure modes, and interpretation limits.
- **Expandable technical detail** for long harness configuration, exhaustive field inventories, archived implementations, or full transcripts.

Disclosure labels should name the hidden material, such as **Detailed harness configuration** or **Full evaluator evidence**, rather than “More.” Core caveats must remain visible. [Page naming and interface scan, Finding 9](assignments/04-page-naming-interface-scan.md)

**Skill Issue design recommendation.** Use one primary diagram to establish the evaluation loop and evidence ownership. Every node should reuse terminology from the headings and artifacts. Prefer accessible SVG or HTML/CSS with a text caption. The one-diagram limit and implementation details are local layout choices; the sources support diagrams for stable lifecycle relationships rather than this exact count or technology. [Page naming and interface scan, Finding 10](assignments/04-page-naming-interface-scan.md) · [HELM deep dive, Finding 16](assignments/07-helm-benchmark-documentation.md)

### Scenario and Attempt Evidence

**Skill Issue design recommendation.** For three scenarios, use three cards, tabs, or an anchored comparison table followed by detail sections. Each overview row should expose enough information to compare without opening the full record: purpose, expected decision, turn count, attempt result, correction count, evidence status, and a detail link. These compact page-layout specifics are local proposals for the current corpus size.

Each detail view should move from summary to evidence:

1. Scenario purpose and boundary.
2. Initial message and later turn/trigger definitions.
3. Expected skill-call behavior and downstream outcome.
4. Configured system and environment.
5. Attempt matrix and criterion-level scorecard.
6. Annotated transcript with invocation events aligned to turns.
7. Produced artifacts and evaluator evidence.
8. Limitations, ambiguity, and source links.

For trajectory comparison, prefer an annotated single-column timeline or synchronized turn anchors for short traces. SWE-Together's three independently scrolling panels make divergence visible at scale but are demanding on narrow screens. Pair at least one strong and weak attempt only when the editorial annotation identifies the exact decision divergence; otherwise provide the complete attempt list without manufacturing a comparison narrative. [SWE-Together deep dive, Findings 12–13](assignments/10-swe-together-conversational-benchmark.md)

Search, multi-axis filters, and a wide leaderboard are lower fit until the scenario library grows materially. If it expands beyond a small set, add filters for scenario/skill family, harness, model route, version, failure class, evidence status, and validity state. [SWE-Together deep dive, Finding 16](assignments/10-swe-together-conversational-benchmark.md) · [HELM deep dive, Findings 14 and 22](assignments/07-helm-benchmark-documentation.md)

### Citations and GitHub Evidence

Place citations and artifact links beside the claim they support. Link exact repository files, commits, releases, scenario records, scorecards, and run bundles rather than only a project homepage. A compact final citation block can support reuse, but it should not be the only provenance route.

Use stable evidence actions with restrained nouns: **Scenarios**, **Attempts**, **Artifacts**, **Reproduce**, **GitHub**, and **History**. Use action-oriented titles for procedures, such as **Run the evaluation** or **Download artifacts**. HELM's public result browser uses nouns for inspectable objects and action titles for operational documentation. [HELM deep dive, Findings 14 and 18](assignments/07-helm-benchmark-documentation.md)

## Conditional Alternatives

### “Methodology” as the Page Name

Use **Methodology** when the surrounding navigation already makes “evaluation” unambiguous and the page is unlikely to be encountered out of context. This is compact and supported by MLCommons and other branded documentation systems. **Evaluation Methodology** remains stronger for direct links, search results, and readers arriving from GitHub. [Page naming and interface scan, Findings 1–2](assignments/04-page-naming-interface-scan.md)

### “Testing Methodology” as the Page Name

Use **Testing Methodology** only if the destination is deliberately narrowed to execution procedure, environment, and operational controls, with scoring, evidence governance, and limitations owned elsewhere. That split conflicts with the current product need for one dense public destination, so it is a secondary choice. [Page naming and interface scan, Finding 2](assignments/04-page-naming-interface-scan.md)

### Separate Results or Technical Reference Pages

A distinct results explorer or technical reproduction guide becomes useful when attempt evidence, commands, or configuration inventories make the methodology hard to read. SWE-bench, HELM, METR, and OpenAI all support this layered pattern. The methodology page should remain the semantic owner and link outward; shared facts should come from a canonical data source to prevent drift. [SWE-bench deep dive, Findings 1 and 8](assignments/05-swe-bench-methodology.md) · [OpenAI disclosure deep dive, Finding 1](assignments/08-openai-evaluation-disclosures.md)

### Adaptive User Simulation

Use a reactive simulator only when later turns must respond to materially divergent agent behavior. Version its seed instructions, action set, state summary, branching policy, maximum turns, judge, and review thresholds. Keep simulator failures distinct from agent failures. For three fixed scenarios, deterministic triggers are more credible and cheaper to audit. [SWE-Together deep dive, Findings 2, 8, and 16](assignments/10-swe-together-conversational-benchmark.md) · [Claude 4 deep dive, Finding 4](assignments/09-anthropic-claude-4-system-card.md)

### LLM or Human Semantic Judging

Use a frozen semantic rubric when direct validators cannot recognize all materially valid outcomes. Validate it against alternative-correct and shortcut outputs, retain per-criterion evidence, and audit disagreement. Prefer direct checks for invocation events and structural invariants. [OpenAI disclosure deep dive, Finding 6](assignments/08-openai-evaluation-disclosures.md) · [SWE-Together deep dive, Finding 4](assignments/10-swe-together-conversational-benchmark.md)

### Controlled or Gated Evidence

Use redaction or gating when artifacts contain credentials, private material, personal data, proprietary skills, or leakage-sensitive evaluation content. Publish safe metadata, hashes, scoring interfaces, counts, and access boundaries even when raw content cannot be public. HELM's GPQA flow and METR's protected task material show that inspectability and unrestricted disclosure are distinct. [HELM deep dive, Finding 13](assignments/07-helm-benchmark-documentation.md) · [METR deep dive, “Reproducibility is layered”](assignments/06-metr-evaluation-methodology.md)

### Formal Benchmark Card or Technical Report

A schema-driven **Benchmark Card** becomes useful if Skill Issue needs one repeatable record per benchmark or campaign. A separate **Technical Report** is justified if the project makes a novel research claim requiring formal methods, figures, references, and archival citation. Neither should replace the web methodology destination for the current scope. [BenchmarkCards](https://github.com/SokolAnn/BenchmarkCards) · [Page naming and interface scan, Finding 3](assignments/04-page-naming-interface-scan.md)

## Lower-Fit and Rejected Interpretations

- **System Card:** lower fit because it conventionally documents a released model or system, safety evaluation, mitigation, and deployment decision. The transferable patterns are evaluated-version disclosure, thresholds, mixed evidence, and limitations—not the name or full report shape. [System-card candidate scan](assignments/02-system-card-candidates.md)
- **Scoring Methodology:** too narrow because scoring is only one part of the required public account. [Page naming and interface scan, Finding 2](assignments/04-page-naming-interface-scan.md)
- **Benchmarking Methodology:** appropriate for a program composed of many benchmarks under shared rules; premature for three governed scenarios unless Skill Issue explicitly frames them as a benchmark program. [Page naming and interface scan, Finding 2](assignments/04-page-naming-interface-scan.md)
- **Leaderboard-first destination:** lower fit for three scenarios. It encourages model-ranking interpretation before the reader understands the configured system, scenario coverage, and evidence. SWE-Together and HELM need leaderboard machinery because they compare many cohorts or models. [SWE-Together deep dive, Finding 16](assignments/10-swe-together-conversational-benchmark.md) · [HELM deep dive, Finding 22](assignments/07-helm-benchmark-documentation.md)
- **Single aggregate score:** lower fit unless weighting and commensurability are defensible. Decision correctness, contract adherence, outcome correctness, correction burden, reliability, and operational validity should remain visible. [HELM deep dive, Finding 8](assignments/07-helm-benchmark-documentation.md) · [SWE-Together deep dive, Findings 3 and 5](assignments/10-swe-together-conversational-benchmark.md)
- **Final artifact as complete proof:** rejected for conversational governance because an artifact cannot demonstrate call timing, clarification, non-invocation, correction handling, or other process requirements. [SWE-Together deep dive, Finding 7](assignments/10-swe-together-conversational-benchmark.md)
- **Plausible final answer as invocation proof:** rejected. Instrumented call evidence is required for call-level claims. [SWE-bench deep dive, Finding 6](assignments/05-swe-bench-methodology.md)
- **Container equals reproducibility:** rejected. Version, architecture, external services, configuration, and retained evidence remain material. [Reproducibility candidate scan, Notes](assignments/03-reproducibility-candidates.md)
- **Reference outcome as oracle ceiling:** rejected unless maximality is formally established. [SWE-Together deep dive, Finding 7](assignments/10-swe-together-conversational-benchmark.md)
- **Fixed replay at arbitrary turn numbers:** lower fit when agent trajectories diverge; deterministic semantic triggers are stronger. [SWE-Together deep dive, Finding 2](assignments/10-swe-together-conversational-benchmark.md)
- **Adaptive simulation by default:** lower fit at three-scenario scale because it adds model-mediated variance and validation burden. [SWE-Together deep dive, Findings 8 and 16](assignments/10-swe-together-conversational-benchmark.md)
- **Full system-card PDF or paper pasted into the site:** lower fit for routine public reference. Use those sources for formal limitation and evidence patterns, not their length, hierarchy, or visual identity. [Page naming and interface scan, Finding 15](assignments/04-page-naming-interface-scan.md)
- **Search-heavy scenario marketplace:** lower fit until the corpus grows beyond a small fixed set. [SWE-Together deep dive, Finding 16](assignments/10-swe-together-conversational-benchmark.md)
- **Hidden chain-of-thought as required evidence:** lower fit and often unavailable; observable transcript and event evidence should carry the claim. [SWE-bench deep dive, Finding 6](assignments/05-swe-bench-methodology.md)
- **Best-of-N presented as single-attempt reliability:** rejected. Candidate selection evaluates a composed selection system and requires separate labeling. [Claude 4 deep dive, Finding 14](assignments/09-anthropic-claude-4-system-card.md)
- **Economic-value framing:** lower fit without observed task prices or a real economic construct. [OpenAI disclosure deep dive, Finding 14](assignments/08-openai-evaluation-disclosures.md)
- **OSWorld as a primary coding-method reference:** rejected because its desktop/web control tasks do not define repository editing, skill invocation, or developer interaction. Environment reset patterns are already covered by coding-relevant sources. [Coding-agent candidate scan, Finding 4](assignments/01-coding-agent-evaluation-candidates.md)

## Candidate Disposition

The supplied candidate assignments separate broad discovery from six selected deep dives. The following inventory preserves those dispositions; it does not independently rerank the underlying publications.

### Selected Deep Dives

| Candidate                              | Principal contribution to the recommendation                                                                                            | Important limit                                            |
| -------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------- |
| Artificial Analysis Coding Agent Index | Concise methodology ownership, configured agent distinction, repeats, efficiency metadata, missing-telemetry treatment, version history | Weak raw task evidence and isolation detail                |
| SWE-bench plus Verified                | Task provenance, executable scoring, containerized harness, retained logs, human quality review                                         | Static, outcome-heavy, weak turn and uncertainty treatment |
| SWE-Lancer                             | Concrete run conditions, manual validation, hidden end-to-end tests, qualitative trajectories                                           | Domain and public/private split constraints                |
| Terminal-Bench                         | Per-task environments, human-verified solutions, executable tests, multiple runs and error bars                                         | Broader terminal scope than skill calling                  |
| SWE-Together                           | Multi-turn evaluation, correction taxonomy, frozen rubrics, repeated attempts, scenario and trajectory interfaces                       | New project; simulator and judge remain model-mediated     |
| AgentLens                              | Process-quality and waste evidence beneath binary success                                                                               | Narrow OpenHands/SWE-bench Verified basis                  |
| RE-Bench                               | Long-horizon execution, human comparison, resource controls, open scaffolding, contamination controls                                   | AI R&D task domain and heavier scale                       |
| PaperBench                             | Hierarchical rubric scoring, judge validation, partial credit, human baseline                                                           | Research-replication domain and model-judge uncertainty    |
| SWE-bench Live                         | Contamination/staleness response, per-task environments, live curation                                                                  | Limited public trajectory and repeat evidence              |

These nine candidates come from the coding-agent scan. [Coding-agent candidate scan, Finding 2](assignments/01-coding-agent-evaluation-candidates.md)

| Candidate                    | Principal contribution to the recommendation                                                               | Important limit                                        |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------- | ------------------------------------------------------ |
| GPT-4o System Card           | Compact sequence from subject and transformed inputs to scoring, results, mitigations, and residual limits | Uncertainty often qualitative                          |
| OpenAI o1 System Card        | Checkpoint mapping, evaluator construction, paired metric definitions                                      | Manual work and instrumentation vary by section        |
| GPT-5 System Card            | Run conditions, elicitation, lower-bound framing, confidence-interval caveats                              | Large release document rather than focused method page |
| Claude 4 System Card         | Mixed automated/manual evidence, transcript review, thresholds, construct failure, third-party work        | Extremely long and methods are distributed             |
| Gemini 2.5 Technical Report  | Guided navigation, explicit attempt conditions, held-out assurance, grader appendix                        | Some results and external evidence are summarized      |
| METR GPT-5 Evaluation Report | Independent evidence chain, assumptions, adversarial checks, access limits                                 | Some evidence is assurance-based or constrained        |
| NIST AI 800-3                | Estimands, variance, confidence-interval validity, scoring semantics                                       | Little agent instrumentation or manual evaluation      |
| UK AISI Frontier AI Trends   | Accessible multi-domain disclosure, repeats, sampling units, uncertainty, withheld-method boundaries       | Aggregated identities and some withheld material       |

These eight candidates come from the system-card scan. [System-card candidate scan](assignments/02-system-card-candidates.md)

| Candidate                                  | Principal contribution to the recommendation                                                | Important limit                                         |
| ------------------------------------------ | ------------------------------------------------------------------------------------------- | ------------------------------------------------------- |
| ACM Artifact Review and Badging            | Separates availability, artifact quality, and independently reproduced results              | Standard, not a complete page design                    |
| MLPerf Training                            | Executable contract, system metadata, run boundaries, repeats, corrections                  | Hardware-intensive training context                     |
| Stanford HELM                              | Progressive disclosure, plural metrics, run evidence, versioned releases, reproduction docs | Large multi-surface system and point-estimate UI        |
| SWE-bench documentation                    | Task provenance, environment pinning, hidden grading, retained task material                | Container and platform caveats remain                   |
| OpenAI SWE-bench construction/audit series | Ongoing validity, task audits, contamination, retirement, recommendation reversal           | Focused on coding-benchmark quality                     |
| EleutherAI evaluation harness              | Minimum reproducibility bundle: task config, prompt, model args, seeds, outputs             | Language-model benchmark rather than agent conversation |
| SV-COMP                                    | Versioned rules, fixed resources, evidence witnesses, raw archive, anti-fingerprinting      | Non-LLM competition domain                              |
| BenchmarkCards                             | Structured objectives, data, method, risks, limitations, ethics, citation                   | Card quality remains dependent on source quality        |

These eight candidates come from the reproducibility scan. [Reproducibility candidate scan](assignments/03-reproducibility-candidates.md)

The dedicated deep-dive assignments selected six systems for full reader-journey analysis: SWE-bench, METR, HELM, OpenAI's connected evaluation disclosures, Anthropic's Claude 4 System Card, and SWE-Together. Their strongest combined pattern is the recommendation in this report: one public navigational owner; an explicit configured-system and scenario contract; retained turn-level evidence; deterministic checks plus bounded semantic review; repeated attempts; controlled execution; artifact-backed reproducibility claims; local limitations; and a visible revision history. [SWE-bench](assignments/05-swe-bench-methodology.md) · [METR](assignments/06-metr-evaluation-methodology.md) · [HELM](assignments/07-helm-benchmark-documentation.md) · [OpenAI](assignments/08-openai-evaluation-disclosures.md) · [Claude 4](assignments/09-anthropic-claude-4-system-card.md) · [SWE-Together](assignments/10-swe-together-conversational-benchmark.md)

### Skim-Only Candidates

| Candidate set   | Useful narrow contribution                                                             | Why it remains secondary                              |
| --------------- | -------------------------------------------------------------------------------------- | ----------------------------------------------------- |
| FeatureBench    | Feature-task shapes, environment diversity, fail/pass test pairing, routing disclosure | Little turn, repeat, uncertainty, or trace evidence   |
| Multi-SWE-bench | Language breadth, Docker assets, logs, environment clearing                            | Mostly extends SWE-bench rather than a new page model |
| MLE-bench       | Resource budgets, human leaderboard baselines, contamination analysis                  | ML competition work, not conversational skill calling |
| GitTaskBench    | Staged workflow outcomes and cost-effectiveness                                        | Small broad-domain set; weak repeat uncertainty       |
| tau2-bench      | Turn orchestration, policy/tool schemas, trajectory retention, grader-version warnings | Non-coding domain                                     |

[Coding-agent candidate scan, Finding 3](assignments/01-coding-agent-evaluation-candidates.md)

| Candidate set                                  | Useful narrow contribution                                  | Why it remains secondary                         |
| ---------------------------------------------- | ----------------------------------------------------------- | ------------------------------------------------ |
| Claude 3.7 Sonnet System Card                  | Shorter system-card hierarchy and mode disclosure           | Claude 4 offers richer mixed-method evidence     |
| Gemini 2.5 Pro Model Card                      | Compact lifecycle governance and non-comparability language | Thin instrument, sample, and uncertainty detail  |
| Operator System Card                           | Agent-system boundary and human-in-the-loop controls        | Insufficient scoring and reconstruction detail   |
| Sora System Card                               | Prompt-source provenance and classifier placement           | Too abbreviated for complete methodology         |
| Holistic Safety and Responsibility Evaluations | Evaluation-purpose and evidence-class framing               | Conceptual rather than one complete result chain |

[System-card candidate scan, Findings 9–13](assignments/02-system-card-candidates.md)

| Candidate set                          | Useful narrow contribution                                     | Why it remains secondary                              |
| -------------------------------------- | -------------------------------------------------------------- | ----------------------------------------------------- |
| NeurIPS dataset/benchmark requirements | Availability, executable code, machine-readable metadata       | Not a run-time or retained-output contract            |
| Croissant 1.1                          | Dataset identity, checksums, semantic versioning, provenance   | Dataset metadata rather than evaluation method        |
| Datasheets for Datasets                | Creator-authored provenance and stewardship                    | Dataset lifecycle rather than execution               |
| Hugging Face Dataset Cards             | Familiar public schema, examples, splits, curation, limits     | Author-driven and insufficient for reproduction alone |
| NISO reproducibility badging           | Standards-aligned vocabulary                                   | Badge taxonomy rather than verification procedure     |
| NIST AI RMF Measure/TEVV               | Construct validity, operating conditions, variance, audit logs | Broad risk framework rather than rerun guide          |

[Reproducibility candidate scan, Candidates 9–14](assignments/03-reproducibility-candidates.md)

### Rejected Candidates

| Candidate                                    | Rejection reason                                                                                                                                  |
| -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| OSWorld                                      | Strong computer-use methodology but insufficiently coding- and skill-invocation-specific; its unique controls are available from closer analogues |
| Meta Llama 3.1 Model Card                    | Useful artifact links and benchmark tables, but weak end-to-end method, manual review, uncertainty, and interpretive chain                        |
| Cohere Command R Model Card                  | Operational usage guidance with insufficient evaluation procedure, scoring, uncertainty, and raw evidence                                         |
| OpenAI Evals legacy repository documentation | Useful version and event-log fragments, but the repository redirects toward newer tooling and lacks a current complete methodology contract       |

[Coding-agent candidate scan, Finding 4](assignments/01-coding-agent-evaluation-candidates.md) · [System-card candidate scan, Findings 14–15](assignments/02-system-card-candidates.md) · [Reproducibility candidate scan, Candidate 15](assignments/03-reproducibility-candidates.md)

## Cross-Source Tensions to Preserve

1. **One page versus a documentation system.** The audience needs one clear destination, while the strongest sources use several evidence depths. Resolve this with one methodology owner and linked scenario, artifact, and reproduction surfaces; do not duplicate the governing text.
2. **Static replay versus reactive simulation.** Static turns maximize control; reactive turns preserve semantic timing when trajectories diverge. For three fixed scenarios, deterministic semantic triggers are the better starting point. Adaptive simulation remains conditional.
3. **Executable checks versus semantic grading.** Tests are reproducible but can be overly strict, incomplete, or exploitable. Semantic judges accept alternatives but add rubric and judge uncertainty. Use direct checks for observable events and structural invariants, with frozen semantic criteria for residue.
4. **Public evidence versus leakage control.** Transparency supports audit, while visible scenarios and solutions can create contamination and refinement pressure. Publish safe evidence and hashes; reserve blinded holdouts or gated material only when the validity benefit is real and disclosed.
5. **Aggregate clarity versus dimensional honesty.** A single number is easy to scan but can hide decision, outcome, correction, stability, or operational differences. Make the attempt matrix and dimensional scores primary; any aggregate is secondary and fully defined.
6. **Reference simplicity versus process completeness.** A reference artifact is useful for calibration but cannot prove all conversational behavior. Treat it as a reference outcome, not a ceiling.
7. **Containers versus reproducibility.** Containers reduce host drift but do not fix architecture, external APIs, provider aliases, secrets, or missing evidence. State the reproducibility tier actually supported.
8. **Versioned transparency versus public-surface drift.** SWE-bench, HELM, OpenAI, and Anthropic show that current methods, software, suites, results, and prose can diverge. Separate version domains and derive repeated facts from canonical records.

## Unsupported Claims and Remaining Unknowns

The supplied external research does not establish the following local facts. The final public page should obtain them from Skill Issue's authoritative source and current evaluation artifacts rather than infer them:

- The canonical names and exact semantic boundaries of the three scenarios.
- The exact user messages, later-turn triggers, expected skill calls, allowed equivalents, and downstream outcomes.
- Which CLI executable, harness, model route, reasoning settings, prompts, tools, permissions, budgets, and environment revisions are currently used.
- Whether skill-call events are fully instrumented and retained across every supported harness.
- The current scoring dimensions, thresholds, aggregation, correction taxonomy, and invalidation rules.
- The number of attempts per scenario and any current result values.
- The current manual steps, reviewer roles, independence, sample-selection method, and disagreement procedure.
- Which artifacts are public, controlled, missing, or canonical on GitHub.
- Whether an external party has rerun or independently reproduced any result.
- Whether any scenario has been hardened against visible-test optimization, contamination, lexical triggers, or repeated refinement.
- The current methodology version, status, date, campaign identifier, and compatibility history.

No source supports claiming that an LLM simulator is behaviorally equivalent to a human user, that an LLM judge is ground truth, that containerization guarantees exact reproduction, that two attempts establish high reliability, or that a three-scenario result generalizes to skill calling broadly. These should remain explicit non-claims unless later local evidence establishes a narrower supported statement. [SWE-Together deep dive, Notes](assignments/10-swe-together-conversational-benchmark.md) · [Reproducibility candidate scan, Notes](assignments/03-reproducibility-candidates.md) · [NIST AI 800-3](https://www.nist.gov/publications/expanding-ai-evaluation-toolbox-statistical-models)

## Bottom Line

The strongest public direction is a restrained, static-first **Evaluation Methodology** page that leads with the evaluation boundary and configured subject, exposes the three scenarios through a shared schema, shows one clear evaluation loop, defines invocation and downstream scoring separately, publishes repeated attempt evidence, discloses manual work and isolation controls, and links every consequential claim to canonical artifacts. It should be dense through structure rather than long prose: narrow reading measure, shallow headings, compact declarative paragraphs, stable tables, visible caveats, one useful diagram, selective disclosure controls, and direct GitHub evidence.

The page earns credibility by making uncertainty and revision governable. It should say when evidence is missing, keep operational invalids separate from behavioral failures, treat reference material as illustrative rather than maximal, and distinguish inspectability from reproduction. Skill Issue should also choose a durable, preferably immutable correction-history mechanism for its own evidence system; that implementation choice is a project recommendation, while the source-backed requirement is to keep corrections and retired methods historically interpretable. This combined direction is more strongly supported than any single benchmark, system card, or reproducibility standard in the supplied research.

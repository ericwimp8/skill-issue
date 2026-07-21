# Reproducibility Candidates

## Assignment

**Goal:** Discover and rank authoritative benchmark cards, dataset cards, evaluation standards, and reproducibility documentation that can inform Skill Issue's public Testing Methodology page.

**Scope:** Internet-only research across primary sources from standards bodies, benchmark authors, official repositories, established research projects, and first-party publications. The review emphasizes environment controls, task and data provenance, versioning, manual steps, execution rules, scoring, retained materials, replication instructions, limitations, uncertainty, and evidence links. Fifteen candidate document systems were inspected and classified as deep-dive, skim-only, or reject.

**Exclusions:** Local Skill Issue source and documentation; proposed website copy; implementation planning; secondary summaries when a primary source was available; and claims about Skill Issue's existing behavior.

## Sources

- Association for Computing Machinery, [Artifact Review and Badging, Version 1.1](https://www.acm.org/publications/policies/artifact-review-and-badging-current) and the [SIGMOD Availability and Reproducibility Initiative](https://reproducibility.sigmodconf.hosting.acm.org/index.html).
- National Information Standards Organization, [NISO RP-31-2021: Reproducibility Badging and Definitions](https://www.niso.org/publications/rp-31-2021-badging) and its [standing committee page](https://www.niso.org/standards-committees/reproducibility-badging).
- MLCommons, [MLPerf Training benchmark page](https://mlcommons.org/benchmarks/training/), [MLPerf Training Rules](https://github.com/mlcommons/training_policies/blob/master/training_rules.adoc), and [reference implementations](https://github.com/mlcommons/training).
- Stanford CRFM, [HELM project](https://crfm.stanford.edu/helm/index.html), [HELM Classic](https://crfm.stanford.edu/helm/latest/), [HELM v1.0 methodology](https://crfm.stanford.edu/helm/v1.0/), and the [HELM repository](https://github.com/stanford-crfm/helm).
- SWE-bench authors, [SWE-bench repository](https://github.com/SWE-bench/SWE-bench), [evaluation guide](https://github.com/SWE-bench/SWE-bench/blob/main/docs/guides/evaluation.md), [containerized harness report](https://github.com/SWE-bench/SWE-bench/blob/main/docs/20240627_docker/README.md), and [SWE-bench Verified dataset card](https://huggingface.co/datasets/SWE-bench/SWE-bench_Verified).
- OpenAI, [Introducing SWE-bench Verified](https://openai.com/index/introducing-swe-bench-verified/), [Why SWE-bench Verified no longer measures frontier coding capabilities](https://openai.com/index/why-we-no-longer-evaluate-swe-bench-verified/), and [Separating signal from noise in coding evaluations](https://openai.com/index/separating-signal-from-noise-coding-evaluations/).
- EleutherAI, [Language Model Evaluation Harness](https://github.com/EleutherAI/lm-evaluation-harness), [task configuration guide](https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/task_guide.md), [CLI interface](https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/interface.md), and [model integration guide](https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/model_guide.md).
- SV-COMP organizers, [SV-COMP 2025 rules](https://sv-comp.sosy-lab.org/2025/rules.php), [competition artifact archive](https://sv-comp.sosy-lab.org/reproduce.php), and [verified results](https://sv-comp.sosy-lab.org/2025/results/results-verified/).
- Sokol et al., [BenchmarkCards paper](https://papers.neurips.cc/paper_files/paper/2025/hash/76175f4355e2f67cf91be468c8860070-Abstract-Datasets_and_Benchmarks_Track.html), [project repository](https://github.com/SokolAnn/BenchmarkCards), and [BenchmarkCards dataset card](https://huggingface.co/datasets/ASokol/BenchmarkCards).
- NeurIPS, [2025 Datasets and Benchmarks call](https://neurips.cc/Conferences/2025/CallForDatasetsBenchmarks) and [data hosting guidelines](https://neurips.cc/Conferences/2025/DataHostingGuidelines).
- MLCommons, [Croissant 1.1 specification](https://docs.mlcommons.org/croissant/docs/croissant-spec-1.1.html), [versioning guidance](https://docs.mlcommons.org/croissant/docs/croissant-spec.html), and the [Croissant repository](https://github.com/mlcommons/croissant).
- Gebru et al., [Datasheets for Datasets publication page](https://www.microsoft.com/en-us/research/publication/datasheets-for-datasets/), [Data Documentation project](https://www.microsoft.com/en-us/research/project/datasheets-for-datasets/), and the [paper](https://www.microsoft.com/en-us/research/uploads/prod/2019/01/1803.09010.pdf).
- Hugging Face, [Dataset Cards documentation](https://huggingface.co/docs/hub/datasets-cards) and the [dataset card creation guide](https://github.com/huggingface/datasets/blob/main/templates/README_guide.md).
- NIST, [AI RMF Core Measure function](https://airc.nist.gov/airmf-resources/airmf/5-sec-core/), [AI RMF Measure playbook](https://airc.nist.gov/airmf-resources/playbook/measure/), and [AI TEVV program](https://www.nist.gov/ai-test-evaluation-validation-and-verification-tevv).
- OpenAI, [Evals repository](https://github.com/openai/evals), [build-eval guide](https://github.com/openai/evals/blob/main/docs/build-eval.md), and [run-evals guide](https://github.com/openai/evals/blob/main/docs/run-evals.md).

## Findings

The ranking contains **8 deep-dive candidates**, **6 skim-only candidates**, and **1 rejected candidate**. Deep dives offer a distinct, concrete document system that could materially shape the public methodology. Skim-only sources supply narrower terminology, metadata, or audit checks. The rejected source contains useful fragments but is a weaker primary reference for a current public methodology.

### Candidate 1 — ACM Artifact Review and Badging

**Classification: Deep-dive.** This is the strongest standards-level source for separating artifact availability, artifact quality, and independently validated results. It defines repeatability, reproducibility, and replicability; treats software, scripts, input data, raw outputs, and analysis scripts as distinct artifacts; and requires functional artifacts to be documented, consistent, complete, and exercisable. Its availability badge requires a persistent archival repository and unique identifier. Its results badges explicitly allow an experiment-specific tolerance rather than demanding numerically identical results.

**Evidence:** The [current ACM policy](https://www.acm.org/publications/policies/artifact-review-and-badging-current) defines the three reproducibility terms, the independent badge dimensions, artifact qualities, archival expectations, and acceptable tolerance. The [SIGMOD implementation](https://reproducibility.sigmodconf.hosting.acm.org/index.html) demonstrates how a venue turns those principles into a package: source or a fully specified black box, configuration and build environment, input data or generation process, system initialization, workload, measurement protocol, raw data, and scripts that regenerate graphs. SIGMOD also distinguishes exact numerical agreement from reproduction of the behavior supporting the claim.

**Implication:** A later synthesis should consider using ACM's distinctions to prevent the Testing Methodology page from collapsing “materials are available,” “the evaluation can be rerun,” and “another party reproduced the conclusion” into one claim. ACM also provides a defensible inventory of retained materials and a standard way to describe tolerance and independent validation.

### Candidate 2 — MLPerf Training Rules and Result System

**Classification: Deep-dive.** MLPerf is the most complete example of an executable benchmark contract coupled to versioned public results. Its rules define the system as hardware plus the exact versions of performance-relevant operating systems, compilers, libraries, drivers, and frameworks. They specify data state, preprocessing, cache flushing, model constraints, random-seed logging, quality targets, timing boundaries, repeat counts, result aggregation, and separate open and closed divisions.

**Evidence:** The [MLPerf Training Rules](https://github.com/mlcommons/training_policies/blob/master/training_rules.adoc) require reference implementations to document the problem, dataset, attribution, preprocessing, train/test separation, data order, initial weights, model structure, optimizer, quality metric and target, evaluation frequency, machine configuration, data download and verification, and run-and-time steps. Reference packages include dataset download and checksum verification scripts plus a timed entrypoint. Closed runs log seeds through MLLog, preserve reference data semantics, and begin with a cache flush or restart. The [benchmark page](https://mlcommons.org/benchmarks/training/) publishes suite versions, datasets, quality targets, repeat-and-trim aggregation, rough uncertainty bands, system metadata, linked code, and a post-publication change log.

**Implication:** This is the best deep dive for environment disclosure, execution boundaries, repeated-run scoring, uncertainty, and post-publication corrections. A later page could adopt the principle that the rules are the source of truth while the public result is a versioned record linking to code, metadata, and a change history.

### Candidate 3 — Stanford HELM

**Classification: Deep-dive.** HELM is the strongest transparency-oriented model for showing what an evaluation covers, how models are adapted to scenarios, and which evidence sits beneath an aggregate score. It deliberately situates tested scenarios inside a larger taxonomy, exposes omissions, standardizes prompts and adaptation, and measures multiple dimensions rather than publishing only a single score.

**Evidence:** The [HELM v1.0 methodology](https://crfm.stanford.edu/helm/v1.0/) states four design commitments: broad coverage with explicit recognition of incompleteness, multi-metric measurement, standardized model adaptation, and transparency. It exposes scenarios, prompts, predictions, results, and code. The [HELM repository](https://github.com/stanford-crfm/helm) supplies commands to run, summarize, and browse a named suite, and says the framework can reproduce its published leaderboards. Individual leaderboards such as [HELM Capabilities](https://crfm.stanford.edu/helm/capabilities/v1.3.0/) are explicitly versioned. The repository also warns that HELM entered maintenance mode on June 1, 2026, which is itself useful lifecycle evidence.

**Implication:** HELM supports a public methodology that names dimensions, selection criteria, and known gaps; standardizes adaptation across compared systems; and lets readers drill from aggregate results to per-instance prompts and outputs. Its maintenance status is a reminder to attach dates and versions to reproducibility claims.

### Candidate 4 — SWE-bench Evaluation and Dataset Documentation

**Classification: Deep-dive.** SWE-bench provides a close domain analogue for evaluating agentic software work against real repositories. Its full system connects source issue provenance, repository commits, per-task environment setup, container execution, regression and issue-resolution tests, detailed output logs, and platform limitations.

**Evidence:** The [SWE-bench Verified dataset card](https://huggingface.co/datasets/SWE-bench/SWE-bench_Verified) records each task's repository, base commit, environment-setup commit, installation version, gold patch, test patch, issue text, FAIL_TO_PASS tests, and PASS_TO_PASS tests. The [evaluation guide](https://github.com/SWE-bench/SWE-bench/blob/main/docs/guides/evaluation.md) specifies input JSONL, dataset selection, run identifiers, concurrency and cache options, and retained `results.json`, per-instance JSONL, and run logs. The [containerized harness report](https://github.com/SWE-bench/SWE-bench/blob/main/docs/20240627_docker/README.md) explains why Conda-only isolation was insufficient, describes layered base/environment/instance images, reports gold-patch validation rates, states resource requirements, and flags ARM support as experimental.

**Implication:** This is a high-value deep dive for task provenance, environment pinning, hidden evaluation behavior, retained per-task materials, and explicit platform caveats. It also demonstrates that containerization should be presented as a tested control with known limits rather than as an automatic guarantee of identical results.

### Candidate 5 — OpenAI's SWE-bench Construction and Audit Series

**Classification: Deep-dive.** This series is unusually valuable because it documents the creation, validation, later invalidation, and replacement of widely used coding benchmarks. It shows that a credible public methodology needs a lifecycle for task-quality findings, contamination, benchmark retirement, and retracted recommendations.

**Evidence:** [Introducing SWE-bench Verified](https://openai.com/index/introducing-swe-bench-verified/) documents a manual campaign using 93 experienced Python developers, three independent annotations per sample, ordinal severity labels, conservative filtering, released annotations, and a public rubric. In February 2026, [OpenAI reported](https://openai.com/index/why-we-no-longer-evaluate-swe-bench-verified/) that 59.4% of a targeted 138-task audit had material task or test issues, found contamination evidence across frontier models, stopped reporting the benchmark, and recommended a replacement. In July 2026, [a second audit](https://openai.com/index/separating-signal-from-noise-coding-evaluations/) found 27.4% of the SWE-Bench Pro public split flagged by the datapoint pipeline and 34.1% flagged by a human campaign, then explicitly retracted the earlier replacement recommendation. That review combined automated screening, multiple investigator-agent passes, five experienced engineers, and escalation of disagreements.

**Implication:** A later methodology should expose benchmark status and revision history, distinguish initial validation from continuing validity, and explain how task audits, annotator disagreement, test strictness, underspecification, low coverage, and contamination affect confidence. This series is evidence against presenting any benchmark score as timeless.

### Candidate 6 — EleutherAI Language Model Evaluation Harness

**Classification: Deep-dive.** The harness provides a concrete, shareable evaluation specification and an extensive retained-output model. It makes the task definition, prompt construction, dataset selection, scoring, post-processing, repetition, seeds, and model formatting part of the reproducibility surface.

**Evidence:** The [task guide](https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/task_guide.md) says a task YAML plus the codebase commit hash is intended to reproduce an evaluation setup. Task configuration records dataset path and configuration, splits, preprocessing, prompt templates, targets, answer choices, few-shot behavior, output type, generation arguments, repeat count, filters, metrics, decontamination behavior, and a task version. The [CLI interface](https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/interface.md) exposes model arguments, device, batch size, generation settings, four seed domains, integrity checks, sample logging, result paths, and full configuration display. The [model guide](https://github.com/EleutherAI/lm-evaluation-harness/blob/main/docs/model_guide.md) states that chat-template text is saved with results. The project can retain aggregate results, all inputs and outputs, task and CLI configuration, command, hardware counts, and timestamps.

**Implication:** A deep dive should extract a minimum reproducibility bundle: task config, code revision, dataset revision, prompt and chat template, model and generation arguments, seeds, evaluator configuration, aggregate results, and per-sample records. This is more useful than publishing only a command or score table.

### Candidate 7 — SV-COMP Rules, Witnesses, and Artifact Archive

**Classification: Deep-dive.** SV-COMP is the strongest non-LLM competition analogue for Skill Issue. Its document system binds authoritative rules, fixed execution resources, versioned benchmark definitions, tool adapters, evidence validation, scoring, participant review, public raw results, and DOI-backed snapshots.

**Evidence:** The [SV-COMP 2025 rules](https://sv-comp.sosy-lab.org/2025/rules.php) define each run as `(ANSWER, WITNESS, TIME)`, fix the operating system and CPU, memory, CPU-time, and wall-time limits, prohibit task fingerprinting, require one global parameter set, and require benchmark definitions plus tested tool-info modules that translate tool output into benchmark outcomes. Correct results require a witness accepted by a validator. Tool archives must expose a version, include a license and README, and contain dependencies or declare installable system packages. Scoring is asymmetric so incorrect claims outweigh correct ones, with runtime breaking equal scores. The [reproduction page](https://sv-comp.sosy-lab.org/reproduce.php) retains DOI snapshots and repository tags for tool archives, tasks, benchmark definitions, results, witnesses, BenchExec, and orchestration components. The [results site](https://sv-comp.sosy-lab.org/2025/results/results-verified/) links downloadable raw results and witness-validation details.

**Implication:** SV-COMP supports a public methodology built around a versioned evaluation bundle rather than a prose-only description. Its separate tool adapter, validated evidence object, anti-fingerprinting rule, raw-result archive, and participant review process are especially relevant to fair cross-harness comparison.

### Candidate 8 — BenchmarkCards

**Classification: Deep-dive.** BenchmarkCards is the most directly applicable candidate for the public documentation shape. It standardizes benchmark objectives, intended users, data sources, annotation, methodology, targeted risks, limitations, similar benchmarks, ethics, licensing, and citation in both human-readable and machine-readable forms.

**Evidence:** The [NeurIPS 2025 paper](https://papers.neurips.cc/paper_files/paper/2025/hash/76175f4355e2f67cf91be468c8860070-Abstract-Datasets_and_Benchmarks_Track.html) introduces and validates the documentation framework through benchmark-author and benchmark-user studies. The [repository](https://github.com/SokolAnn/BenchmarkCards) supplies a template and marks cards reviewed by the original benchmark author. The [dataset card](https://huggingface.co/datasets/ASokol/BenchmarkCards) lists the common fields and explains that cards combine manual curation with LLM-assisted extraction, original-author feedback where available, and explicit limitations. It warns that missing source information and incomplete author review can produce inaccurate cards and tells users to verify important claims against original sources.

**Implication:** This is the strongest candidate for the information architecture of a public benchmark or methodology card. Its own caveats also establish a rule for downstream use: cards are navigational summaries, while critical execution and provenance claims should link to primary evidence.

### Candidate 9 — NeurIPS 2025 Datasets and Benchmarks Requirements

**Classification: Skim-only.** NeurIPS offers a useful publication-level completeness gate for accessible data, executable code, and machine-readable metadata, but it is less detailed about run-time environment capture, scoring, retained outputs, and independent reproduction than the deep-dive sources.

**Evidence:** The [track call](https://neurips.cc/Conferences/2025/CallForDatasetsBenchmarks) requires dataset and benchmark code at submission, says data must be obtainable without a personal request, requires code to be documented and executable, mandates Croissant metadata, and requires accepted code and datasets to be documented and public by camera-ready. The [hosting guide](https://neurips.cc/Conferences/2025/DataHostingGuidelines) requires a reviewer-accessible URL, Croissant file, and validation confirmation, and compares hosts by programmatic loading, private review access, gated access, paper linking, and DOI support.

**Implication:** Use this as a compact publication-readiness checklist for availability, executable code, validated metadata, and persistent hosting. Do not treat it as a complete methodology template.

### Candidate 10 — MLCommons Croissant 1.1

**Classification: Skim-only.** Croissant is the strongest machine-readable source for dataset identity, resource checksums, structure, versioning, and provenance. It is a dataset metadata standard rather than a complete evaluation methodology.

**Evidence:** The [Croissant 1.1 specification](https://docs.mlcommons.org/croissant/docs/croissant-spec-1.1.html) requires a declared specification version, dataset version, creation and modification dates, license, description, and resource structure. It supports provenance through W3C PROV relationships at dataset, resource, record-set, field, and value levels. The [versioning guidance](https://docs.mlcommons.org/croissant/docs/croissant-spec.html) recommends semantic versioning, defines changes that should increment patch, minor, or major versions, and strongly recommends SHA-256 checksums for stable files. It separately marks live datasets and explains how consumers should preserve reproducibility as data grows.

**Implication:** A later synthesis could borrow Croissant's version-plus-checksum discipline for task corpora, document whether a corpus is static or live, and record derivation and transformation provenance. The full schema is likely too technical for the public page itself.

### Candidate 11 — Datasheets for Datasets

**Classification: Skim-only.** Datasheets supplies the most mature narrative framework for data provenance and stewardship. Its focus is dataset documentation rather than benchmark execution, but its manual, creator-authored questions cover information that automated metadata commonly misses.

**Evidence:** The [publication](https://www.microsoft.com/en-us/research/publication/datasheets-for-datasets/) proposes documenting motivation, composition, collection, recommended uses, and operating characteristics. The [Data Documentation project](https://www.microsoft.com/en-us/research/project/datasheets-for-datasets/) emphasizes that datasheets are intentionally not automated because they capture creator knowledge that is otherwise lost. The [paper](https://www.microsoft.com/en-us/research/uploads/prod/2019/01/1803.09010.pdf) extends the lifecycle through preprocessing, distribution, and maintenance, asking about canonical references, archival and redundant distribution, licensing and access restrictions, maintainers, update cadence, revision communication, errata, obsolescence, and extension mechanisms.

**Implication:** Use it to improve the narrative provenance of evaluation tasks: why they were created, how they were selected and transformed, who maintains them, and what changes or uses would invalidate their interpretation.

### Candidate 12 — Hugging Face Dataset Cards

**Classification: Skim-only.** Hugging Face offers a practical public card convention and metadata layer that readers already recognize. Its fields are useful, but compliance is largely author-driven and it does not independently establish that an evaluation is reproducible.

**Evidence:** The [Dataset Cards documentation](https://huggingface.co/docs/hub/datasets-cards) defines a repository `README.md` with YAML metadata for license, language, size, task categories, and discovery. The [creation guide](https://github.com/huggingface/datasets/blob/main/templates/README_guide.md) adds source links, point of contact, supported tasks and metrics, examples, field definitions, splits, curation rationale, collection and normalization, annotators, sensitive information, intended use, social impact, biases, limitations, curators, license, and citation.

**Implication:** This is a good model for approachable dataset presentation and linking, especially examples and split descriptions. Pair it with stronger version, checksum, execution, and validation sources rather than using a dataset card as proof by itself.

### Candidate 13 — NISO RP-31-2021 Reproducibility Badging

**Classification: Skim-only.** NISO is useful for standards-aligned vocabulary and recognition claims, but the recommended practice intentionally focuses on badge taxonomy and does not specify the detailed process for verifying results.

**Evidence:** [NISO RP-31-2021](https://www.niso.org/publications/rp-31-2021-badging) is an approved recommended practice with DOI `10.3789/niso-rp-31-2021`. It standardizes common recognition practices, vocabulary, and iconography for sharing data and methods. The [standing committee page](https://www.niso.org/standards-committees/reproducibility-badging) explains that the work harmonizes previously ad hoc schemes across computational and computing sciences. Its terminology was coordinated with the wider scientific community and influenced ACM's current naming.

**Implication:** Consult this source if the public page uses formal terms such as reproduced, replicated, or open research object. ACM remains the more operational deep dive for evidence and review criteria.

### Candidate 14 — NIST AI RMF Measure and TEVV Guidance

**Classification: Skim-only.** NIST provides a strong claim-quality and uncertainty lens, but it is a broad risk-management framework rather than a benchmark reproduction guide.

**Evidence:** The [AI RMF Core](https://airc.nist.gov/airmf-resources/airmf/5-sec-core/) calls for objective, repeatable, or scalable TEVV processes with documented metrics, methods, tools, test sets, uncertainty, benchmark comparisons, and limitations on generalizability. The [Measure playbook](https://airc.nist.gov/airmf-resources/playbook/measure/) recommends documenting operating conditions and limits, construct validity, internal validity, external validity, variance, confidence intervals or comparable dispersion measures, robustness, reliability, assumptions, data provenance, and audit logs. The [TEVV program](https://www.nist.gov/ai-test-evaluation-validation-and-verification-tevv) emphasizes that metric definitions, tasks, strengths, limitations, and context all matter.

**Implication:** Use NIST to audit whether the methodology explains what each metric claims to measure, where it generalizes, what uncertainty remains, and which operating conditions bound the result. More concrete sources should own the rerun instructions.

### Candidate 15 — OpenAI Evals Legacy Repository Documentation

**Classification: Reject.** The repository contains useful examples of eval version naming and local event logging, but it is a weaker primary source for a current public methodology than the other candidates. Its own README redirects users toward dashboard-based Evals, while its guides retain legacy model examples and do not provide a complete environment, dataset provenance, or artifact-review contract.

**Evidence:** The [build-eval guide](https://github.com/openai/evals/blob/main/docs/build-eval.md) uses `<eval_name>.<split>.<version>`, instructs authors to bump the version when an eval changes, defines JSONL input and reference-answer structure, and recommends meta-evaluating model graders against human choice labels. The [run guide](https://github.com/openai/evals/blob/main/docs/run-evals.md) retains local JSONL event logs and supports explicit record paths, but its examples center on older model identifiers and a legacy CLI. The [repository README](https://github.com/openai/evals) now points users toward dashboard configuration.

**Implication:** Do not select this as a primary deep dive. If later work borrows its version-bump convention, model-grader meta-evaluation, or event-log retention, verify those practices against the current OpenAI evaluation product or a current repository before presenting them as contemporary guidance.

## Notes

- The strongest cross-source pattern is a layered evidence model: a human-readable methodology, machine-readable configuration, immutable or versioned inputs, an explicit execution environment, raw and aggregate outputs, and a revision or deprecation history.
- “Containerized” is not equivalent to “fully reproducible.” SWE-bench's own migration report documents platform sensitivity before Docker and retains architecture caveats afterward; MLPerf and SV-COMP additionally define hardware, operating-system, resource, cache, and timing controls.
- Benchmark validity is time-dependent. HELM's maintenance status and OpenAI's successive SWE-bench recommendation reversals show why a public page needs dates, versions, status, and supersession information.
- BenchmarkCards and Hugging Face cards improve navigation and completeness, but both depend on source quality and author diligence. Critical claims should remain traceable to rules, task records, logs, repositories, or archival artifacts.
- No independent reproduction of Skill Issue results was performed in this internet-only assignment. The sources support candidate selection and documentation patterns, not claims about the repository's current evaluations.

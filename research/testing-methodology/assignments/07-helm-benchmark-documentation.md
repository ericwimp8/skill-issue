# HELM Benchmark Documentation

## Assignment

**Goal.** Deep-dive Stanford CRFM's official HELM benchmark documentation system, including its public website, papers, repository, scenario and model catalogs, leaderboard and result browser, raw evidence, developer documentation, and reproduction materials. Identify transferable documentation and evidence-design patterns for Skill Issue while separating direct observation from inference.

**Scope.** Internet-only inspection of official Stanford CRFM, Stanford-hosted HELM, Read the Docs, OpenReview/arXiv, GitHub, PyPI-linked, and public HELM Google Cloud Storage surfaces. The review covers purpose and scope; evaluated models; scenarios and datasets; adaptation and execution; metrics and aggregation; controls; reproducibility and versioning; retained evidence; limitations and uncertainty; reader journey; naming; heading hierarchy; paragraph density; readable measure; tables; notes; diagrams; expandable material; citations; code/data links; and GitHub connections. Visual and DOM inspection was performed at a wide desktop viewport on 2026-07-21.

**Exclusions.** No HELM evaluation was executed, no large raw-result corpus was downloaded, no gated dataset was decrypted, no non-official commentary was used as primary evidence, and no attempt was made to reproduce Stanford's wording, branding, logo, or visual identity. The review does not judge the underlying benchmark datasets independently of HELM's documented implementation.

## Sources

- [HELM leaderboards landing page](https://crfm.stanford.edu/helm/) — current public index, tagline, hero diagram, benchmark catalog, navigation, and GitHub connection.
- [HELM Capabilities latest landing page](https://crfm.stanford.edu/helm/capabilities/latest/) — current project overview, mini leaderboard, release selector, and links into the full result browser. The inspected selector identified `v1.15.0 (2025-11-24)` as `latest` and exposed immutable versions `v1.0.0` through `v1.15.0`.
- [HELM Capabilities full leaderboard](https://crfm.stanford.edu/helm/capabilities/latest/#/leaderboard) — sortable cross-model/scenario table, group selector, metric tabs, point scores, and links to run evidence.
- [HELM Capabilities scenarios](https://crfm.stanford.edu/helm/capabilities/latest/#/scenarios) — scenario taxonomy, metric descriptions, missing metadata markers, and analysis charts.
- [HELM Capabilities models](https://crfm.stanford.edu/helm/capabilities/latest/#/models) — rendered catalog of 68 model configurations in release `v1.15.0`, including 30 marked `Open` and 38 marked `Limited`, with creator, stable HELM identifier, description, access class, and source links.
- [HELM Capabilities GPQA leaderboard](https://crfm.stanford.edu/helm/capabilities/latest/#/leaderboard/gpqa) — scenario-specific scores plus execution metadata such as inference time, evaluation count, training count, truncation, prompt tokens, and output tokens.
- [HELM Capabilities GPQA run evidence](https://crfm.stanford.edu/helm/capabilities/latest/#/runs/gpqa:subset=gpqa_main,use_chain_of_thought=true,use_few_shot=false,model=openai_gpt-5-mini-2025-08-07) — inspected run page with scenario/model context, adapter specification, dataset-access notice, instance-level predictions and metrics, expandable request details, pagination, and raw JSON links.
- [GPQA run specification JSON](https://storage.googleapis.com/crfm-helm-public/capabilities/benchmark_output/runs/v1.12.0/gpqa%3Asubset%3Dgpqa_main%2Cuse_chain_of_thought%3Dtrue%2Cuse_few_shot%3Dfalse%2Cmodel%3Dopenai_gpt-5-mini-2025-08-07/run_spec.json) — validated public retained artifact containing concrete `ScenarioSpec`, `AdapterSpec`, model deployment, decoding parameters, limits, metric implementations, augmentation controls, and group membership.
- [GPQA full scenario-state JSON](https://storage.googleapis.com/crfm-helm-public/capabilities/benchmark_output/runs/v1.12.0/gpqa%3Asubset%3Dgpqa_main%2Cuse_chain_of_thought%3Dtrue%2Cuse_few_shot%3Dfalse%2Cmodel%3Dopenai_gpt-5-mini-2025-08-07/scenario_state.json) — validated as a public, CORS-enabled retained artifact; content remains subject to the GPQA encryption/access control shown in the UI.
- [GPQA per-instance statistics JSON](https://storage.googleapis.com/crfm-helm-public/capabilities/benchmark_output/runs/v1.12.0/gpqa%3Asubset%3Dgpqa_main%2Cuse_chain_of_thought%3Dtrue%2Cuse_few_shot%3Dfalse%2Cmodel%3Dopenai_gpt-5-mini-2025-08-07/per_instance_stats.json) — validated as a public, CORS-enabled retained artifact linked from the run evidence page.
- [Introducing HELM Capabilities](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) — official methodology narrative for scenario selection, prompts, model access, post-processing, judge ensembles, metrics, aggregation, robustness changes, cross-source comparisons, and success/failure examples for `v1.0.0`.
- [Original HELM results and methodology page](https://crfm.stanford.edu/2022/11/17/helm.html) — official account of broad coverage, recognition of incompleteness, seven metric families, standardized prompting, evaluated providers, headline findings, limitations, and recommended reader journey.
- [Holistic Evaluation of Language Models paper](https://arxiv.org/abs/2211.09110) and [TMLR/OpenReview record](https://openreview.net/forum?id=iO4LZibEqW) — original research scope: 30 models, 16 core scenarios, 26 targeted scenarios, seven metric families, 96% core scenario coverage, prompt/completion release, and benchmark-as-living-map framing.
- [HELM Classic latest](https://crfm.stanford.edu/helm/classic/latest/) and [historical HELM v1.0 site](https://crfm.stanford.edu/helm/v1.0/) — continuity from the original HELM paper to the versioned Classic result browser.
- [HELM GitHub repository](https://github.com/stanford-crfm/helm) — source repository, README, Apache-2.0 license, release links, documentation links, frontend, benchmark implementation, public issue workflow, and 2026 maintenance-mode notice.
- [HELM changelog](https://github.com/stanford-crfm/helm/blob/main/CHANGELOG.md) — package-level release history; inspected current package release `v0.5.16` dated 2026-04-29 and explicit breaking-change/scenario/model/framework categories.
- [HELM Read the Docs home](https://crfm-helm.readthedocs.io/en/latest/) — developer/user documentation landing page, search, side navigation, quick start, badges, papers, citation, breadcrumbs, Edit on GitHub, and version control.
- [HELM tutorial](https://crfm-helm.readthedocs.io/en/latest/tutorial/) — run/summarize/serve lifecycle and the retained artifact contract (`run_spec.json`, `scenario.json`, `scenario_state.json`, `per_instance_stats.json`, `stats.json`, summary and group files).
- [Reproducing Leaderboards](https://crfm-helm.readthedocs.io/en/latest/reproducing_leaderboards/) — version-specific configuration and schema files, common reproduction shell, per-project parameters, and public/gated/private MedHELM variants.
- [Downloading Raw Results](https://crfm-helm.readthedocs.io/en/latest/downloading_raw_results/) — public GCS paths, release/suite selection, large-download warnings, and raw-evidence access methods.
- [Advanced Benchmarking Guide](https://crfm-helm.readthedocs.io/en/latest/benchmark/) — dry-run and token-estimation controls before sending model requests.
- [Code Structure](https://crfm-helm.readthedocs.io/en/latest/code/) — scenario → preprocessor → adapter → executor → metric → runner pipeline and explicit stale-document warning.
- [Maintenance Mode Policy](https://crfm-helm.readthedocs.io/en/latest/maintenance_mode/) — status, support boundary, no-new-evaluations policy, external-API breakage caveat, best-effort maintenance, and contribution rules effective 2026-06-01.

## Findings

### Finding 1: HELM is a documentation system, not a single report

**Observation.** HELM distributes distinct reader jobs across three connected surfaces. The [public HELM site](https://crfm.stanford.edu/helm/) is the discovery and results surface. [Read the Docs](https://crfm-helm.readthedocs.io/en/latest/) is the operational and reference surface. [GitHub](https://github.com/stanford-crfm/helm) is the implementation, issue, release, and editable-source surface. Papers and Stanford CRFM blog posts explain why a benchmark exists and how to interpret it. Each surface links to the others rather than trying to contain every level of detail.

**Evidence.** The public landing page leads to named leaderboards and GitHub. A project landing page leads to its blog post and leaderboard. The result browser's fixed navigation exposes `Leaderboard`, `Models`, `Scenarios`, `Predictions`, and `GitHub`. The Read the Docs breadcrumb exposes `Edit on GitHub`, while the repository README sends users back to the official leaderboards, papers, and reproduction docs.

**Implication.** Skill Issue can adopt the same separation of reader jobs: a concise evaluation overview, an interactive results/evidence browser, a task-oriented methodology/reproduction guide, and repository-owned source. The links and ownership boundaries matter more than reproducing HELM's exact site count.

### Finding 2: The reader journey is progressive disclosure from claim to raw evidence

**Observation.** HELM's intended journey is discoverable without reading the paper first: select a leaderboard, inspect aggregate scores, narrow to a scenario, open a model/run, inspect adapter settings and individual outputs, and download retained JSON. The original [HELM methodology page](https://crfm.stanford.edu/2022/11/17/helm.html) explicitly recommends drilling down from aggregate statistics to raw prompts and predictions.

**Evidence.** On the current Capabilities site, a mean-score cell links to model-filtered runs; a scenario score links to a concrete run; and the run page links to `Spec JSON`, `Full JSON`, and `Per-Instance Stats JSON`. The run page then presents `Adapter Specification`, `Instances + Predictions`, `All metrics`, per-instance `Input`, `References`, `Thinking`, `Prediction raw text`, and `Metrics`, plus collapsed `Request details` and instance pagination.

**Implication.** Skill Issue should make every summary score traversable to the exact cases, prompts, outputs, evaluator results, and run controls that produced it. Evidence should be adjacent to the claim, with raw artifacts as an additional layer rather than the only explanation.

### Finding 3: The original HELM methodology foregrounds coverage, plural metrics, and control

**Observation.** Original HELM defines holistic evaluation through three principles: broad coverage while making incompleteness explicit, multiple desiderata rather than accuracy alone, and standardized adaptation across models. The original study evaluated 30 models from 12 providers across 16 core and 26 targeted scenarios. It attempted seven metric families—accuracy, calibration, robustness, fairness, bias, toxicity, and efficiency—and raised core model/scenario coverage from 17.9% in prior reporting to 96.0% under its standardized evaluation.

**Evidence.** These counts and principles appear in the [paper abstract](https://arxiv.org/abs/2211.09110) and the official [2022 HELM methodology page](https://crfm.stanford.edu/2022/11/17/helm.html). The methodology page names open, API-limited, and closed models from providers including AI21 Labs, Anthropic, BigScience, Cohere, EleutherAI, Google, Meta, Microsoft, NVIDIA, OpenAI, Tsinghua University, and Yandex.

**Implication.** Skill Issue's methodology should define the full evaluation space before presenting the tested subset. It should state what the campaign covers, what it omits, why each tested class was selected, and which comparison controls remain fixed.

### Finding 4: HELM's current flagship narrows the benchmark while preserving the framework

**Observation.** HELM Capabilities is a deliberately curated successor for general capability comparisons rather than a wholesale continuation of all original scenarios. Version `v1.0.0` selected five scenarios using saturation, recency, and quality criteria: MMLU-Pro, GPQA, IFEval, WildBench, and Omni-MATH. The inspected latest release still centers these five groups while the model catalog has expanded to 68 configurations.

**Evidence.** The official [HELM Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) documents 22 models and five scenarios at launch. The current [scenario catalog](https://crfm.stanford.edu/helm/capabilities/latest/#/scenarios) still shows five scenarios. The current [model catalog](https://crfm.stanford.edu/helm/capabilities/latest/#/models) rendered 68 rows, 30 `Open` and 38 `Limited`, under release `v1.15.0`.

**Implication.** A Skill Issue evaluation surface can distinguish a stable framework from a curated current campaign. The documentation should state when a smaller flagship set is a prioritization choice rather than implying it is exhaustive.

### Finding 5: Scenario documentation uses both conceptual metadata and concrete evaluation contracts

**Observation.** The scenario catalog describes each benchmark with columns for `Task`, `What`, `Who`, `When`, `Language`, `Description`, and `Metric`. It also exposes a stable scenario identifier below the display name. The run evidence adds the concrete scenario class, subset, adapter, decoding, evaluation count, and metric implementation.

**Evidence.** The [current scenario page](https://crfm.stanford.edu/helm/capabilities/latest/#/scenarios) shows, for example, GPQA as English graduate-level biology/physics/chemistry questions from domain experts in 2023, and identifies its Chain-of-Thought correctness metric. The retained [GPQA run specification](https://storage.googleapis.com/crfm-helm-public/capabilities/benchmark_output/runs/v1.12.0/gpqa%3Asubset%3Dgpqa_main%2Cuse_chain_of_thought%3Dtrue%2Cuse_few_shot%3Dfalse%2Cmodel%3Dopenai_gpt-5-mini-2025-08-07/run_spec.json) records the implementation class `helm.benchmark.scenarios.gpqa_scenario.GPQAScenario`, subset `gpqa_main`, adapter method, metric classes, and group.

**Implication.** Skill Issue should document evaluation dimensions in two layers: reader-oriented meaning and machine-oriented contract. Display labels should never replace stable identifiers, concrete harness/model versions, or evaluator implementations.

### Finding 6: Current execution procedures are scenario-specific but fully exposed

**Observation.** Original HELM favored simple, generic few-shot prompts to standardize adaptation. Current HELM Capabilities uses scenario-specific adaptation: zero-shot Chain-of-Thought for MMLU-Pro and GPQA, official prompt templates where available, rule-based post-processing for multiple choice and IFEval, and LLM judges for WildBench and Omni-MATH. This is a documented evolution rather than a hidden change.

**Evidence.** The [Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) explains that all scenarios are capped/downsampled to 1,000 instances, lists official versus reused prompt templates, names its post-processing, and explains judge voting. The inspected GPQA run specification records `max_eval_instances: 1000`, `max_train_instances: 0`, `num_train_trials: 1`, `num_trials: 1`, `num_outputs: 5`, `temperature: 1`, `max_tokens: 14096`, no perturbations, and the precise model deployment.

**Implication.** Skill Issue should expose campaign-wide controls and task-specific exceptions side by side. A user should be able to distinguish a rule inherited from the framework from a scenario-specific adaptation or evaluator workaround.

### Finding 7: HELM treats pipeline defects and evaluator ambiguity as methodology findings

**Observation.** The Capabilities methodology documents failure modes encountered during implementation: judge-format failures, judge bias, missing official prompt variants, ambiguous answer-format instructions, and hallucination-prone judge prompts. It describes the intervention made for each problem instead of presenting the evaluation pipeline as frictionless.

**Evidence.** The official [Pipeline Robustness](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) section explains multi-judge averaging and fallbacks, reusing GPQA's prompt for MMLU-Pro, broadening GPQA answer extraction, and changing Omni-MATH judging after human canary evaluation found the official prompt could induce erroneous judgments.

**Implication.** Skill Issue should retain evaluator and harness failure evidence as part of the published methodology. When evaluation infrastructure changes because a failure is observed, the documentation should record the failure class, evidence, intervention, and resulting compatibility boundary.

### Finding 8: Metrics are named, normalized, and linked, but the top-level mean hides judgment choices

**Observation.** Current Capabilities reports scenario-native metrics: multiple-choice accuracy for MMLU-Pro and GPQA, strict instruction-following accuracy for IFEval, a rescaled multi-judge WB-Score for WildBench, and LLM-judged answer equivalence for Omni-MATH. It ranks models by the arithmetic mean after putting scenario scores on a 0–1 scale. The site links scenario scores to supporting runs and highlights column leaders visually.

**Evidence.** The [Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) explains the metric choices and why its mean score replaced Classic/Lite's mean win rate. It notes that win rate depends on the comparison set and can react strongly to small rank inversions. The [current leaderboard](https://crfm.stanford.edu/helm/capabilities/latest/#/leaderboard) presents `Mean score` and one column per scenario under tabs for `Accuracy`, `Efficiency`, and `General information`.

**Implication.** Skill Issue should place the aggregate formula and normalization decisions beside the aggregate score. If criteria are not meaningfully commensurate, the interface should privilege dimensional scores and documented trade-offs over a single total.

### Finding 9: Model documentation includes provenance and access constraints, not only names

**Observation.** The current model catalog groups models by creator and shows a human name, stable HELM identifier, short description, source links such as model cards, papers, blogs, and system cards, and an `Open` or `Limited` access label. Run pages repeat model-specific source links near the result.

**Evidence.** The [model catalog](https://crfm.stanford.edu/helm/capabilities/latest/#/models) rendered 68 configurations across providers and distinguished reasoning/extended-thinking variants in the stable model ID and description. The inspected GPT-5 mini run linked the provider's launch post and system card adjacent to the evaluated model name.

**Implication.** Skill Issue should identify the tested harness/model by immutable or date-qualified ID and classify access requirements. Provider name alone is insufficient for reproducibility or interpretation.

### Finding 10: The version model separates software releases, leaderboard releases, and run suites

**Observation.** HELM exposes multiple version layers. The Python package/repository was at `v0.5.16` in the inspected official surfaces. HELM Capabilities marked leaderboard `v1.15.0 (2025-11-24)` as the `latest` alias and offered immutable leaderboard versions back to `v1.0.0`. A run surfaced in that latest release can retain artifacts under a different suite version, such as the inspected GPQA run under `runs/v1.12.0/`.

**Evidence.** The [repository release/changelog](https://github.com/stanford-crfm/helm/blob/main/CHANGELOG.md), the Read the Docs/PyPI badge, the Capabilities release dropdown, and the public GCS run URLs expose these distinct numbers. The [Downloading Raw Results](https://crfm-helm.readthedocs.io/en/latest/downloading_raw_results/) guide explicitly distinguishes project, release, and suite paths.

**Implication.** Skill Issue should separately label CLI/package version, evaluation definition version, public results release, and concrete run/campaign identifier. A mutable `latest` route should always resolve to an immutable version shown in the UI.

### Finding 11: Reproduction is a documented pipeline with named configuration artifacts

**Observation.** HELM's reproduction guide starts from version-specific `run_entries_*.conf` and `schema_*.yaml` files, then runs `helm-run`, `helm-summarize`, and `helm-server`. It publishes per-leaderboard values for training trials, evaluation instances, priority, run entries, and schema. It differentiates public, gated, and private benchmark configurations where access prevents universal reproduction.

**Evidence.** [Reproducing Leaderboards](https://crfm-helm.readthedocs.io/en/latest/reproducing_leaderboards/) provides a common shell and project-specific parameter blocks. The [tutorial](https://crfm-helm.readthedocs.io/en/latest/tutorial/) explains the function and generated files of each command. [Advanced Benchmarking](https://crfm-helm.readthedocs.io/en/latest/benchmark/) exposes dry-run and token-estimation controls before model requests are sent.

**Implication.** Skill Issue's reproduction material should be executable and version-bound, with a cheap proof command before a full campaign. Access-limited portions should have an explicit public subset or be labeled gated rather than being described as fully reproducible.

### Finding 12: HELM retains both human-readable evidence and machine-readable source material

**Observation.** Per run, HELM retains the run specification, serialized scenario, request/response state, per-instance statistics, and aggregate statistics. `helm-summarize` adds summary, run, group, metadata, JSON, and LaTeX artifacts. Published projects store raw results in an unauthenticated public GCS bucket, with version-specific paths and UI links.

**Evidence.** The [tutorial](https://crfm-helm.readthedocs.io/en/latest/tutorial/) defines the artifact contract. [Downloading Raw Results](https://crfm-helm.readthedocs.io/en/latest/downloading_raw_results/) lists project buckets and warns that full projects may require hundreds of gigabytes and Classic exceeds 5 TB. HTTP inspection confirmed the run spec, scenario state, and per-instance statistics artifacts are public and CORS-enabled.

**Implication.** Skill Issue should retain the minimal complete evidence graph for a result: immutable run definition, inputs or safe references, output/transcript, per-case evaluator results, aggregate computation, and environment/version metadata. Bulk download size and retention cost should be considered explicitly.

### Finding 13: Evidence access can be transparent without exposing restricted content

**Observation.** HELM's GPQA run page provides the adapter, model, metrics, counts, encrypted references, and JSON locations while withholding plaintext instances until a reader accepts the dataset author's anti-leakage request by entering `Yes, I agree` and clicking `Decrypt`. Expandable request details expose decoding controls even while protected content remains encrypted.

**Evidence.** The inspected [GPQA run page](https://crfm.stanford.edu/helm/capabilities/latest/#/runs/gpqa:subset=gpqa_main,use_chain_of_thought=true,use_few_shot=false,model=openai_gpt-5-mini-2025-08-07) showed the access notice and encrypted instance tokens. The [Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) omits the GPQA example for the same dataset-use reason.

**Implication.** Skill Issue should separate verifiability from indiscriminate disclosure. Sensitive prompts, credentials, proprietary skills, or benchmark items can be redacted or gated while still exposing run identity, control metadata, counts, scoring logic, and safe evidence references.

### Finding 14: The public result browser uses noun-based, stable page names

**Observation.** Project paths follow `/helm/<project>/<version>/`; project names are short semantic categories such as `Capabilities`, `Safety`, `Classic`, `Lite`, and `Long Context`. Primary navigation uses nouns—`Leaderboard`, `Models`, `Scenarios`, `Predictions`—while detail pages are titled with the selected group, scenario, or model. The task-oriented documentation instead uses action names such as `Quick Start`, `Downloading Raw Results`, `Reproducing Leaderboards`, `Adding New Models`, and `Adding New Scenarios`.

**Evidence.** These names were consistent across the [HELM index](https://crfm.stanford.edu/helm/), [Capabilities site](https://crfm.stanford.edu/helm/capabilities/latest/), and [Read the Docs navigation](https://crfm-helm.readthedocs.io/en/latest/).

**Implication.** Skill Issue should use short nouns for browsable evaluation objects and action-oriented titles for procedures. Stable semantic URLs should include an immutable release or campaign identifier.

### Finding 15: Heading hierarchy and density change with reader task

**Observation.** The public HELM landing page uses a 36 px H1, a 20 px supporting heading, one short explanatory sentence, two calls to action, a large right-side process diagram, and substantial whitespace. The project landing uses a 30 px H1, one approximately five-line paragraph in an 843 px text column, and a compact two-column leaderboard. Full leaderboards and catalogs switch to dense tables. Read the Docs uses a 300 px persistent navigation rail, an 800 px content container with approximately 696 px readable text measure, 28 px H1, 24 px H2, and 16/24 px body text.

**Evidence.** These values were obtained from rendered DOM/computed-style inspection of the [HELM index](https://crfm.stanford.edu/helm/), [Capabilities landing](https://crfm.stanford.edu/helm/capabilities/latest/), and [Read the Docs home](https://crfm-helm.readthedocs.io/en/latest/). The leaderboards use small 14–16 px tabular text, striped rows, sortable column headers, and full viewport width.

**Implication.** Skill Issue should match density to purpose: sparse orientation, moderate methodology prose, and dense evidence only where comparison requires it. Long-form prose should retain a readable measure rather than inheriting the width of result tables.

### Finding 16: Tables are the primary comparison mechanism; diagrams explain lifecycle and taxonomy

**Observation.** HELM uses tables for model rankings, scenario metadata, model provenance, metric definitions, parameter matrices, and version-specific reproduction settings. Diagrams are reserved for high-level relationships such as scenarios + models → HELM → rankings, the scenario/adaptation/model/metric lifecycle, and benchmark aspect taxonomies. Scenario pages supplement tables with small analysis charts.

**Evidence.** The [HELM landing](https://crfm.stanford.edu/helm/) shows a simple hero process diagram. The [Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) includes a captioned architecture diagram and examples. The [scenario page](https://crfm.stanford.edu/helm/capabilities/latest/#/scenarios) uses one semantic metadata table followed by a total-count card and horizontal distribution bars.

**Implication.** Skill Issue should reserve diagrams for stable process or ownership relationships and use tables for inspectable comparisons. Charts should answer a summary question that the table does not answer immediately.

### Finding 17: Expandable material and direct downloads keep evidence pages usable

**Observation.** The run page keeps request-specific fields collapsed under native `Request details` expanders, while essential adapter controls stay visible in a summary card. It paginates instances, offers per-instance `Copy Link` actions, and provides raw JSON downloads at the top of the evidence section. The leaderboard uses tabs for metric families and a selector for scenario groups.

**Evidence.** The inspected GPQA run contained 10 collapsed request-detail blocks on page 1 and pagination labeled `Page 1 of 45`. The adapter card linked `Spec JSON`, `Full JSON`, and `Per-Instance Stats JSON` before the instance list.

**Implication.** Skill Issue should default to the smallest complete evidence view, then allow the reader to expand request metadata, evaluator traces, or transcripts. Direct-linkable cases and raw downloads improve review and issue reporting.

### Finding 18: Citations are attached to the object being described

**Observation.** Model descriptions link their model card, blog, paper, or system card inline. Scenario methodology links the dataset project and paper. Long-form methodology uses inline citations, captioned figures, code-formatted prompts, and success/failure examples. Read the Docs exposes badges, source repository links, `Edit on GitHub`, and a copyable BibTeX citation.

**Evidence.** See the [current model catalog](https://crfm.stanford.edu/helm/capabilities/latest/#/models), [Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html), and [Read the Docs home](https://crfm-helm.readthedocs.io/en/latest/).

**Implication.** Skill Issue should cite the controlling source beside each model, harness, dataset, rubric, and metric definition. A generic source list at the end is insufficient when several versions or implementations exist.

### Finding 19: HELM makes some limitations explicit but does not quantify score uncertainty in the inspected UI

**Observation.** HELM is unusually explicit about incompleteness, inaccessible models, opaque training data, contamination uncertainty, prompt sensitivity, evaluator failures, inconsistent external benchmark numbers, and API drift. The current Capabilities tables inspected present point estimates and execution counts but no visible confidence intervals, error bars, or uncertainty bands. Missing scenario taxonomy values are rendered as `?` rather than silently inferred.

**Evidence.** The original [HELM methodology](https://crfm.stanford.edu/2022/11/17/helm.html) warns that findings are snapshots, some models are inaccessible, training data is often undisclosed, prompt strategy materially changes results, and the benchmark remains incomplete. The [Capabilities methodology](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html) documents cross-source score disagreement and judge/prompt failure modes. The [scenario page](https://crfm.stanford.edu/helm/capabilities/latest/#/scenarios) shows unknown MMLU-Pro taxonomy cells as `?`.

**Implication.** Skill Issue should emulate the explicit caveat style and improve on point-estimate-only presentation where repeated trials or evaluator variance make uncertainty estimable. Unknown metadata should remain visibly unknown.

### Finding 20: Documentation status is a first-class control

**Observation.** HELM puts its maintenance-mode notice near the top of the README and documentation. Individual reference pages can carry a warning that the page is stale and may be incorrect. The maintenance policy states that external APIs and model deprecations can break scenarios/models, recommends testing fit for use, and removes any fixed release cadence.

**Evidence.** The [Maintenance Mode Policy](https://crfm-helm.readthedocs.io/en/latest/maintenance_mode/) is effective from 2026-06-01. The [Code Structure](https://crfm-helm.readthedocs.io/en/latest/code/) page displays an explicit stale warning. The repository [changelog](https://github.com/stanford-crfm/helm/blob/main/CHANGELOG.md) records breaking changes and deprecations by version.

**Implication.** Skill Issue should show freshness, support status, and compatibility warnings on the methodology or campaign they affect. Historical results should remain browsable under immutable versions even after active support changes.

### Finding 21: Recommended Skill Issue documentation pattern

**Observation.** The strongest transferable HELM patterns are structural rather than visual: a concise index, semantic campaign pages, explicit scope and omitted space, browsable dimensions, immutable releases, run-level evidence, machine-readable artifacts, task-oriented reproduction docs, inline provenance, controlled disclosure, and candid failure notes.

**Evidence.** These patterns recur across the [leaderboard index](https://crfm.stanford.edu/helm/), [Capabilities result browser](https://crfm.stanford.edu/helm/capabilities/latest/), [methodology post](https://crfm.stanford.edu/2025/03/20/helm-capabilities.html), [reproduction guide](https://crfm-helm.readthedocs.io/en/latest/reproducing_leaderboards/), and [repository](https://github.com/stanford-crfm/helm).

**Implication.** A high-fit Skill Issue reader journey would be:

1. **Evaluation overview:** purpose, decision question, current campaign, immutable version, freshness/status, evaluated model/harness matrix, known omissions.
2. **Methodology:** task/failure taxonomy, selection criteria, execution controls, evaluator contract, metrics, aggregation formula, uncertainty treatment, and pipeline exceptions.
3. **Results explorer:** dimensional results before aggregate totals, filters by skill/harness/model/failure class, sortable tables, and clear missing/error states.
4. **Case evidence:** exact skill bundle/version, prompt, environment, transcript/output, evaluator result, failure classification, retry/seed metadata, and source links, with redaction or gating where required.
5. **Reproduction:** smallest proof command, full campaign command, dependency/access prerequisites, retained-artifact schema, and versioned configuration files.
6. **History:** campaign releases, evaluator/harness changes, corrected results, deprecated configurations, and maintained links to old evidence.

### Finding 22: Conditional and lower-fit HELM patterns

**Observation.** Several HELM choices solve HELM-specific scale and language-model problems and should not be copied automatically.

**Evidence.** The public site uses extremely wide comparison tables; the scenario taxonomy is language-domain-specific; current Capabilities uses a simple cross-scenario mean; raw projects can be terabytes; judge ensembles are used only for subjective/free-form scenarios; and documentation is split across a website, Read the Docs, blog posts, papers, and GitHub.

**Implication.** Apply these patterns conditionally:

- **Wide matrices:** useful when a reader genuinely compares many models and scenarios; lower fit on narrow screens or when most cells need caveats. Offer responsive freezing, column selection, or a focused comparison view.
- **Single mean score:** useful only when normalization and weighting are defensible. Skill Issue may need separate capability, compliance, reliability, cost, and evidence-quality dimensions rather than one ranking.
- **Task/What/Who/When/Language taxonomy:** adapt semantically to Skill Issue's actual objects, such as task, expected skill behavior, harness, model, version/date, failure class, and authority boundary.
- **LLM judge ensembles:** appropriate for subjective outputs if judge identities, prompts, failures, and aggregation are retained. Rule-based or deterministic checks should remain preferred when they directly measure the contract.
- **Public raw-artifact buckets:** appropriate for non-sensitive evidence and reproducible campaigns; lower fit for proprietary skill bundles, secrets, personal data, or benchmark-leakage risks. Publish redacted manifests and gated evidence instead.
- **Multiple documentation platforms:** workable with strong linking and source ownership; risky when the same methodology is duplicated. Skill Issue should generate or reference repeated facts from one semantic owner.
- **Large branded hero art:** useful for orientation but secondary to evidence. Reuse the conceptual role of a process diagram, not HELM's visual identity.

## Notes

- The GitHub changelog previously referenced a separate leaderboard changelog, but `LEADERBOARD_CHANGELOG.md` was not present at the repository root during this inspection. Leaderboard history was therefore validated through the rendered release selector and immutable project URLs rather than that dead-end path.
- The current Capabilities landing identifies release `v1.15.0 (2025-11-24)` as latest, while the repository package release is `v0.5.16` from 2026-04-29 and maintenance mode began 2026-06-01. These are different version domains and should not be conflated.
- The inspected latest Capabilities release catalog contained 68 model configurations, but this count belongs to the rendered `v1.15.0` model table and should be cited with that release rather than treated as an invariant HELM-wide count.
- The GPQA run linked from the latest release retained artifacts under suite `v1.12.0`. This supports HELM's release/suite distinction; it is not evidence that the latest release page itself is stale.
- The original OpenReview PDF endpoint presented a browser-verification interstitial during one retrieval attempt. The paper's arXiv record, Stanford methodology page, OpenReview record, repository citation, and published abstract provided sufficient primary evidence without bypassing that control.
- The Read the Docs navigation contains a visible `None` entry for one paper link, and the current scenario table contains `?` cells for some MMLU-Pro metadata. These are observed documentation-quality gaps, not inferred data values.
- The visual measurements are desktop observations, not a full responsive or accessibility audit. Search, semantic landmarks, table headers, native details, and linkable routes were present, but keyboard flow, screen-reader verbosity, color contrast, and mobile behavior were outside this assignment.
- Useful follow-up search terms: `HELM scenario metadata`, `HELM RunSpec`, `HELM ScenarioState`, `HELM schema_capabilities`, `HELM run_entries_capabilities_reasoning_v2`, `HELM prompt-level transparency`, `HELM contamination.yaml`, `HELM leaderboard release summary.json`.

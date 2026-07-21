# Benchmark Analysis Page Research Map

## Run Configuration

- **Goal:** Determine how a public, information-dense benchmark Analysis page should explain findings from model and harness skill-calling evaluations without inventing Skill Issue results.
- **Framing:** Best-supported direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers.
- **Source scope:** Internet only.
- **Researcher budget:** 12.
- **Requested active concurrency:** 12. Runtime capacity permits eight child researchers alongside the orchestrator and parent, so researchers run in backfilled waves of up to eight.
- **Research root:** `research/benchmark-analysis-page/`
- **Assignment root:** `research/benchmark-analysis-page/assignments/`
- **Final aggregation target:** `research/benchmark-analysis-page/benchmark-analysis-page-deep-research.md`
- **Requested output shape:** A source-backed recommendation for page language, headings, information order, chart-to-prose relationships, comparative claims, observation-versus-interpretation boundaries, limitations, uncertainty, and practical meaning; include conditional alternatives and lower-fit patterns.

## Research Domains

### Benchmark Publisher Narratives

Study how publishers organize findings, connect visuals to prose, and express practical meaning in public benchmark pages.

### Statistical Interpretation

Study how credible sources qualify comparisons, uncertainty, sensitivity, and inference from evaluation data.

### Benchmark Validity And Transparency

Study how benchmark-design and auditing sources separate measured observations from broader claims and disclose limitations.

### Editorial And Visual Semantics

Study headings, information order, chart adjacency, captions, concise language, and progressive disclosure in information-dense public reporting.

## Discovery Waves And Assignments

### Wave 1: Seed-Source Deep Dives

1. `assignments/01-artificial-analysis-presentation.md` — Artificial Analysis methodology and Intelligence Index analysis; presentation structure, chart-to-prose linkage, and comparative language.
2. `assignments/02-epoch-benchmark-presentation.md` — Epoch AI benchmark pages; information hierarchy, evidence navigation, and practical interpretation.
3. `assignments/03-anthropic-statistical-evaluations.md` — Anthropic statistical evaluation guidance; uncertainty, ranking claims, and observation-versus-interpretation boundaries.
4. `assignments/04-nist-statistical-toolbox.md` — NIST statistical-model evaluation guidance; uncertainty, estimands, and defensible comparative statements.
5. `assignments/05-nist-caisi-reporting-example.md` — NIST CAISI evaluation report example; public report ordering, findings language, and limitations.
6. `assignments/06-microsoft-adele-explanation.md` — Microsoft Research ADeLe analysis; explaining cross-task performance and practical meaning.
7. `assignments/07-openai-benchmark-audit.md` — OpenAI benchmark audit; separating signal from noise, audit narrative, and caveated conclusions.
8. `assignments/08-betterbench-validity.md` — BetterBench; benchmark-quality dimensions and how validity concerns should shape analysis prose.

### Wave 2: Complementary Publishers And Methodologies

9. `assignments/09-benchscope-methodology.md` — Benchscope methodology; public methodology-to-analysis linkage, labeling, and transparency patterns.
10. `assignments/10-stanford-helm-reporting.md` — Stanford HELM public benchmark/reporting surfaces; dense result organization, scenarios, metrics, and caveats.
11. `assignments/11-metr-benchmark-reporting.md` — METR public evaluation reports; concise findings, uncertainty, practical meaning, and visual explanation.
12. `assignments/12-google-deepmind-benchmark-reporting.md` — Google DeepMind benchmark publication examples; narrative structure, claims discipline, and visual-to-text relationships.

## Source Targets And Expected Evidence

Each assignment should prioritize the named primary source, follow directly linked primary methodology or report pages when necessary, and capture:

- exact page or publication titles and URLs;
- page headings and information order;
- specific chart, table, caption, callout, or prose relationships;
- examples of comparative-claim language and qualification;
- how observations, interpretations, limitations, and practical implications are distinguished;
- adaptable semantic patterns and lower-fit patterns for Skill Issue;
- relevant dead ends, inaccessible pages, and unsupported observations.

Short excerpts may be used only when they materially demonstrate a pattern and remain within compliant quotation limits.

## Candidate Ranking And Fan-Out

The caller already supplied ten high-value seed sources from eight publishers or research groups. Those sources are selected for deep dive because they directly cover public benchmark narrative, statistical evaluation, benchmark validity, or audit reporting. Stanford HELM, METR, and Google DeepMind are complementary deep dives selected to broaden publisher and reporting styles. Discovery that does not fit the remaining researcher budget must be recorded as skim-only, rejected, or unsupported within the relevant assignment rather than spawning further branches.

## Aggregation Contract

After all twelve assignment documents exist and are checked for schema and evidence, a `data_aggregator` receives only the run goal, framing, assignment folder and paths, target path, and requested output shape. The aggregator must synthesize the final document from the constrained assignment corpus and preserve caveats and conflicting evidence.

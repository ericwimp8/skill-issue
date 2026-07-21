# External Methodology Page Research Map

## Run Definition

- **Goal:** identify authoritative public methodology and evaluation-disclosure patterns that can inform Skill Issue's dense, information-first testing methodology page.
- **Source scope:** internet only.
- **Active concurrency:** eight researchers.
- **Total researcher budget:** ten researchers.
- **Final aggregation target:** `research/testing-methodology/external-methodology-page-research.md`.
- **Requested synthesis shape:** best-supported direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers.

## Product Details To Preserve

Skill Issue needs a GitHub Pages-compatible methodology destination beside an existing restrained editorial Project page. The page must explain three governed conversational skill-calling scenarios, controlled CLI execution, invocation instrumentation, scorecards, skill hardening, environment controls, evidence access, explicit manual steps, and limitations. Readers should be able to inspect scenario messages, expected calls, skills, scorecards, and canonical GitHub evidence.

## Research Domains

1. **Coding-agent benchmark disclosures:** complete public pages describing subjects, tasks, trajectories, execution, scoring, controls, and artifacts.
2. **Model and system evaluation disclosures:** system cards and evaluation reports that balance density, limitations, uncertainty, and evidence.
3. **Benchmark governance and reproducibility:** benchmark cards, dataset cards, rules, environment isolation, retained sources, and replication guidance.
4. **Information architecture and evidence interfaces:** page naming, reading flow, hierarchy, measure, tables, notes, diagrams, expandable detail, browsers, and repository links.

These are parent domains. Researchers receive narrower source ecosystems or individual projects rather than an entire domain.

## Discovery Wave

| Assignment | Narrow scope | Source targets | Expected evidence | Output |
| --- | --- | --- | --- | --- |
| 01 | Official coding-agent evaluation candidate scan | Benchmark-author sites and repositories | 10-15 candidates ranked deep-dive, skim-only, or reject; direct URLs; selection rationale | `assignments/01-coding-agent-evaluation-candidates.md` |
| 02 | Official model/system-card candidate scan | Model-developer and research-organization disclosures | 10-15 candidates ranked by methodological and editorial fit | `assignments/02-system-card-candidates.md` |
| 03 | Benchmark-card and reproducibility candidate scan | Benchmark authors, standards bodies, research projects | 10-15 candidates ranked for controls, reproducibility, and evidence traceability | `assignments/03-reproducibility-candidates.md` |
| 04 | Methodology-page naming and interface scan | Authoritative benchmark and evaluation destinations | Supported page names and full-page interface patterns with URLs and counterexamples | `assignments/04-page-naming-interface-scan.md` |

## Anchor Deep Dives

| Assignment | Narrow scope | Source targets | Expected evidence | Output |
| --- | --- | --- | --- | --- |
| 05 | SWE-bench public methodology system | Official website, paper, docs, repository, task browser | Full reading flow, task construction, environment, scoring, evidence links, limitations, interface observations | `assignments/05-swe-bench-methodology.md` |
| 06 | METR evaluation methodology system | Official methodology, task standard, reports, repositories | Agent-evaluation protocol, manual work, instrumentation, controls, reproducibility, uncertainty, page structure | `assignments/06-metr-evaluation-methodology.md` |
| 07 | Stanford HELM benchmark documentation | Official site, papers, repository, scenario and metric documentation | Taxonomy, scenario transparency, metric disclosure, tables/browsers, reproducibility, limitations | `assignments/07-helm-benchmark-documentation.md` |
| 08 | OpenAI evaluation disclosure pages | Official benchmark reports and system cards most relevant to software agents | Full-page structure, test descriptions, caveats, evidence hierarchy, citations, artifact links | `assignments/08-openai-evaluation-disclosures.md` |

## Ranked Follow-Up Wave

Two assignments remain reserved for the strongest candidates or cross-cutting evidence gaps surfaced by assignments 01-04. Each must be a single project, document set, or tightly bounded pattern comparison.

| Assignment | Selection rule | Provisional output |
| --- | --- | --- |
| 09 | Highest-fit unstudied candidate for execution controls, reproducibility, or coding-agent transparency | `assignments/09-<selected-source>.md` |
| 10 | Highest-fit unstudied candidate for editorial density, evidence access, or limitations presentation | `assignments/10-<selected-source>.md` |

Candidates outside the remaining budget will be retained as skim-only, rejected, unsupported, or blocked rather than spawning further branches.

## Evidence Standard

Every assignment must inspect complete pages rather than collect isolated headings. Findings must cite direct URLs and distinguish observed page structure from researcher inference. Recommendations must be applicable to Skill Issue without copying wording or visual identity. Screenshots are optional and only justified when page structure or interactive evidence presentation cannot be described reliably from source inspection.

## Shared Assignment Schema

Each researcher writes exactly one Markdown file with:

- `# <Research Assignment Name>`
- `## Assignment` containing `Goal`, `Scope`, and `Exclusions`
- `## Sources` listing inspected URLs, repositories, versions, commits, publications, or documents
- `## Findings` with repeatable finding headings; every finding includes prose, `Evidence`, and `Implication`
- `## Notes` only for relevant dead ends, caveats, unsupported observations, or useful search terms

## Fan-Out And Aggregation

The first eight assignments run concurrently. After discovery results are ranked, assignments 09 and 10 are selected within the fixed ten-researcher budget. The completed assignment folder is then handed unchanged to `data_aggregator`; the orchestrator does not write the synthesis.

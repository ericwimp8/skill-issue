# Page Naming And Interface Scan

## Assignment

**Goal:** Identify conventional audience-facing names and interface patterns for an authoritative public evaluation-method page, with particular attention to restrained technical sites that can be implemented on GitHub Pages.

**Scope:** Internet-only review of complete public pages from benchmark authors, model developers, research organizations, standards bodies, and established evaluation projects. The review covers naming, page hierarchy, prose measure, paragraph rhythm, punctuation, section spacing, tables, callouts, diagrams, disclosure controls, task browsers, artifact links, citations, and repository links.

**Exclusions:** Visual-brand imitation, promotional launch pages as primary models, private or inaccessible methodology, and recommendations about Skill Issue's substantive evaluation design. This assignment evaluates public-document conventions and interface fit, not the underlying validity of any benchmark.

## Sources

- [Artificial Analysis Intelligence Benchmarking Methodology](https://artificialanalysis.ai/methodology/intelligence-benchmarking) — complete, versioned methodology for a multi-benchmark index; inspected title, contents navigation, tables, per-evaluation records, task lists, source links, legacy material, and version history.
- [ARC-AGI-3 Scoring Methodology](https://docs.arcprize.org/methodology) — focused scoring-method page; inspected docs navigation, on-page contents, definition-first sections, formulas, examples, interpretation table, copy controls, and GitHub edit link.
- [ARC Prize Verified Testing Policy](https://arcprize.org/policy) — public policy page that separates “Testing Methodology” from scoring and publication policy; inspected table of contents, execution descriptions, repository links, and result-artifact links.
- [SWE-bench Verified](https://www.swebench.com/verified.html) — benchmark overview and setup disclosure; inspected title/deck, top artifact links, section dividers, expandable setup detail, citation formats, and GitHub links.
- [SWE-bench FAQ](https://www.swebench.com/SWE-bench/faq/) — question-led explanation of evaluation execution, metrics, datasets, and reproduction; inspected hierarchy and task-oriented navigation.
- [HELM Instruct](https://crfm.stanford.edu/2024/02/18/helm-instruct.html) — complete research-method web page; inspected numbered research hierarchy, diagrams, tables, inline numbered citations, raw-results browser, GitHub implementation link, citation block, and references.
- [HEIM: Holistic Evaluation of Text-to-Image Models](https://crfm.stanford.edu/helm/heim/latest/) — benchmark landing page with model/scenario exploration; inspected paper/GitHub actions, framework diagram, scenario/model inventories, prompt-level transparency statement, and results browsing.
- [Task-Completion Time Horizons of Frontier AI Models](https://metr.org/time-horizons/) — living evaluation dashboard and methodological FAQ; inspected interactive charts, toggles, citation popover, methodological detail, misconception-oriented questions, and update history.
- [MLCommons AILuminate Safety Methodology](https://mlcommons.org/ailuminate/safety-methodology/) — concise web methodology with assessment standard, process stages, benchmark diagram, and technical-paper link.
- [Anthropic Model System Cards](https://www.anthropic.com/system-cards) and [Claude 4 System Card](https://assets.anthropic.com/m/6c940a1b69ed6a1c/original/Claude-4-System-Card.pdf) — inspected “system card” scope, release-index table, long-form PDF hierarchy, capability/safety coverage, evaluation tables, and references.
- [OpenAI GPT-4o System Card](https://cdn.openai.com/gpt-4o-system-card.pdf) — inspected system-card hierarchy, “Evaluation methodology” subsection, explicit methodology limitations, tables, numbered references, and PDF density.
- [ARC-AGI-3 Technical Report](https://arcprize.org/media/ARC_AGI_3_Technical_Report.pdf) — inspected “technical report” conventions: abstract, numbered paper sections, formal methodology, figures, references, and release-specific archival presentation.
- [Hugging Face Annotated Model Card Template](https://huggingface.co/docs/hub/model-card-annotated) — inspected “card” semantics, standardized evaluation fields, structured metadata, testing-data/factors/metrics hierarchy, and generated results widgets.
- [NIST: Towards Best Practices for Automated Benchmark Evaluations](https://www.nist.gov/news-events/news/2026/01/towards-best-practices-automated-benchmark-evaluations) — inspected public-agency explanatory structure, audience statement, three-stage evaluation framing, publication metadata, and link to a formal report.

## Findings

### Finding 1: “Evaluation Methodology” is the strongest general-purpose name

“Evaluation Methodology” is the clearest default when a page explains the complete method used to design, run, score, validate, and report evaluations. It names both the activity and the document's purpose without assuming that the reader already knows whether the work is a benchmark, a test campaign, or a model release. A standalone H1 can be “Skill Issue Evaluation Methodology”; within a clearly branded site and navigation context, “Evaluation Methodology” is sufficient.

**Evidence:** OpenAI and HELM use “Evaluation methodology” or “Methodology” for the part of a larger report that explains evaluation design and execution. NIST describes automated benchmark evaluation as a process spanning objective and benchmark selection, implementation and execution, then analysis and reporting. Artificial Analysis makes the scope explicit in the fuller title “Artificial Analysis Intelligence Benchmarking Methodology.” These examples converge on “evaluation” for the broad activity and “methodology” for the authoritative explanation.

**Implication:** Prefer **Evaluation Methodology** when the page owns the end-to-end public account. Add the product or project name only when the page may be encountered outside the site's normal navigation context.

### Finding 2: Narrower names should correspond to genuinely narrower ownership

The strongest public pages use a qualifier only when it accurately limits the document. “Scoring Methodology” explains measures, formulas, aggregation, and interpretation. “Testing Methodology” explains how evaluations are operationally run. “Benchmarking Methodology” is appropriate for a program that combines many benchmarks and maintains shared execution rules. “Methodology” alone works when adjacent navigation supplies the missing noun.

**Evidence:** ARC-AGI-3's dedicated “Scoring Methodology” stays centered on actions, human baselines, level/game aggregation, and score interpretation. ARC Prize's policy instead places “Testing Methodology” over run configuration, dataset security, provider handling, and publication of outputs. Artificial Analysis uses “Benchmarking Methodology” for an index assembled from nine separate evaluations, then documents shared parameters and per-benchmark implementations. MLCommons uses the compact page heading “Methodology” inside an AILuminate safety-methodology route and branded context.

**Implication:** Use **Scoring Methodology** only for a score-focused subpage; use **Testing Methodology** when the page is principally an operational runbook for public readers; use **Benchmarking Methodology** when multiple benchmark implementations are the central object. A broad method page should avoid a narrow title that hides substantial design, artifact, or limitation content.

### Finding 3: “Benchmark Card,” “System Card,” and “Technical Report” carry established but different promises

These names are conditional alternatives rather than synonyms for a public methodology page. A card is a structured record about one object. A system card documents a released model or system, including capabilities, safety evaluation, and deployment decisions. A technical report is an archival research publication with formal sections, figures, references, and release-specific detail.

**Evidence:** Hugging Face's annotated model-card template treats evaluation as fields within a broader structured model record and supports machine-readable result metadata. Anthropic explicitly defines system cards as documents of model capabilities, safety evaluations, and responsible deployment decisions, indexed by model and date. OpenAI and Anthropic system cards are long, numbered reports in which methodology is one section among training, risk, mitigations, deployment, and appendices. ARC-AGI-3's technical report uses a paper-style abstract and numbered research sections rather than a task-oriented web explanation.

**Implication:** Avoid **Benchmark Card** unless the output is a repeatable schema instantiated once per benchmark. Avoid **System Card** unless the page documents an evaluated system release and deployment decision. Use **Technical Report** for a citable, versioned companion document; it is lower fit as the only explanation on a small public site.

### Finding 4: The most readable hierarchy moves from orientation to mechanics to accountability

A restrained method page benefits from a shallow, predictable sequence: brief definition and scope; what is evaluated; evaluation setup; execution flow; scoring and interpretation; artifacts and reproducibility; limitations; version history or change record. H2s should denote durable concepts. H3s should appear only where a concept contains separable mechanisms, such as per-task and aggregate scoring.

**Evidence:** ARC begins with what is measured and what counts as an action, then establishes the human baseline before formulas and interpretation. Artificial Analysis introduces index scope and composition, states evaluation principles and shared parameters, then provides per-evaluation implementation records, legacy evaluations, and version history. NIST's public explanation summarizes three stages: define/select, implement/run, analyze/report. HELM Instruct follows introduction, methodology, results, exploration, citation, and references, but its paper-like numbering creates more hierarchy than a compact public method page needs.

**Implication:** Put definitions and evaluation boundaries before implementation detail. Place limitations and historical comparability in first-class sections rather than footnotes. Reserve deeper heading levels for actual sub-procedures, not visual spacing.

### Finding 5: A short deck and evidence actions should precede long prose

Strong pages establish purpose in one sentence and expose the primary verification routes immediately. The opening should tell a new reader what is evaluated and why the method exists. A compact action row can link to the repository, executable instructions, data or artifacts, and a formal paper if one exists.

**Evidence:** ARC places “How ARC-AGI-3 scoring works” directly below the H1. SWE-bench Verified uses a one-sentence deck and immediately links to the paper, GitHub repository, and explanatory post. HELM/HEIM place Paper, GitHub, and leaderboard/results actions near the title. ARC docs also provides “Copy page” and repository edit affordances, appropriate for technical documentation.

**Implication:** Use an H1, a single-sentence deck, and a restrained row of two to four evidence links. Prefer labels such as **Repository**, **Run an evaluation**, **Artifacts**, and **Technical report** over promotional calls to action.

### Finding 6: Main prose should be narrow; tables and browsers may break out wider

The strongest long-form methodology layouts constrain normal text while allowing genuinely tabular or interactive content more room. This yields a reading column suitable for sustained technical prose without forcing complex comparison tables into the same measure.

**Evidence:** Artificial Analysis's live HTML uses a `max-w-3xl` main article column (roughly 48rem) with eight-unit vertical section spacing, while broader site containers are available for other components. ARC's docs content uses a typography/prose container and a centered `max-w-3xl` option, plus responsive scrollable tables. SWE-bench's custom stylesheet allows a much wider 1200px container beside a 260px sidebar; that works for benchmark surfaces but produces a less disciplined measure for long prose. METR intentionally expands to 1200px for interactive charts and controls, then returns to question-and-answer prose.

**Implication:** Target about **60–75 characters per line** for paragraphs. Let tables, diagrams, code blocks, and interactive explorers use a wider region or horizontal scrolling. A GitHub Pages implementation can achieve this with one article-width variable and one explicitly wider “data surface” class.

### Finding 7: Paragraph rhythm is compact, declarative, and example-driven

Readable methodology pages generally use one idea per paragraph, often one to four sentences, followed by a list, formula, table, or concrete example. They rely on declarative sentences and explicit nouns more than rhetorical flourish. Em dashes are useful inside definition lists; colons introduce examples, conditions, or enumerations; parentheses carry abbreviations and pinned versions.

**Evidence:** ARC defines an action in a single paragraph, then defines the human baseline in two paragraphs and three bullets. Its score sections place a formula first and worked examples immediately after. SWE-bench uses short overview paragraphs and moves configuration detail into bullets. Artificial Analysis repeats a compact record pattern—status, description, dataset/source, implementation—across evaluations. METR's FAQ headings phrase the reader's likely misconception, then answer directly before adding caveats.

**Implication:** Prefer short explanatory paragraphs followed by evidence structures. Use worked examples for scoring and lifecycle transitions. Avoid promotional superlatives, stacked rhetorical questions, and long paragraphs that combine purpose, implementation, and limitation claims.

### Finding 8: Tables work best for comparison and interpretation, not narrative procedure

Tables are effective when rows share a stable schema: benchmark composition, evaluation parameters, score meanings, versions, models, or artifact availability. Narrative decisions, caveats, and multi-step execution are more readable as prose and ordered lists.

**Evidence:** Artificial Analysis's index table exposes category, evaluation, question count, repeats, response type, scoring, weight, and tool use in one scan. ARC uses a small score-interpretation table after explaining the formula in prose. Anthropic's system-card index uses a simple model/date/document table. NIST explains its three evaluation stages in prose because each stage needs meaning rather than field comparison.

**Implication:** Use tables for stable cross-item fields and a compact score-interpretation key. Use ordered lists for the run sequence. Every table should have a sentence that says what decision or comparison it supports; mobile layouts should scroll rather than compress columns into unreadable text.

### Finding 9: Notes and disclosures should distinguish essential caveats from optional detail

Essential validity conditions should remain visible in the main flow. Optional configuration history, exhaustive task lists, and secondary reproduction detail can be collapsed. A disclosure label should describe the hidden material rather than merely say “more.”

**Evidence:** SWE-bench keeps the comparability claim visible, then places detailed bash-only setup and version notes in a native `<details>` element. Artificial Analysis keeps legacy evaluations visible and explicitly labels them retired or superseded, preserving historical comparability. OpenAI's GPT-4o system card gives “Limitations of the evaluation methodology” its own visible subsection. METR leaves all FAQ answers expanded; this supports searchability but creates a very long page. ARC's “Documentation Index” callout is useful to machine readers but is ancillary to human scoring explanation.

**Implication:** Keep scope, comparability, leakage/security constraints, known failure modes, and interpretation limits expanded. Collapse only long implementation inventories or archival source material, using labels such as **Detailed harness configuration** or **Retired benchmark implementations**. Native `<details>` is sufficient for a static site and preserves keyboard access.

### Finding 10: One explanatory diagram is valuable when it establishes the evaluation loop

A diagram earns space when it compresses a multi-stage relationship that prose would otherwise repeat: task or scenario, system under evaluation, harness, grader, metric, and result artifact. Decorative diagrams or multiple nearly identical process images dilute authority.

**Evidence:** HELM uses framework diagrams to show the relationship among scenario, adaptation, model, evaluator, criterion, and metric, then refers back to those components in the methodology. MLCommons places a benchmark overview diagram after explaining prompt supply, hidden/public splits, evaluator construction, and grading. METR's charts are the evaluated result surface rather than decorative method illustrations. ARC's scoring page needs formulas and worked examples more than a process diagram.

**Implication:** Include at most one restrained system diagram near the overview when the evaluation has several actors or phases. Give it a text caption and ensure every node maps to terminology used in headings and artifacts. Prefer SVG or HTML/CSS for a GitHub Pages site.

### Finding 11: Interactive browsers are strongest as linked evidence surfaces, not substitutes for method prose

Scenario, task, run, and artifact browsers make transparency concrete, but readers still need a stable explanation of selection rules, execution parameters, and interpretation. The methodology page should describe what the explorer contains and link to an appropriately filtered initial view.

**Evidence:** HEIM exposes models, scenarios, prompts, generated images, and metrics while also describing its four evaluation components. HELM Instruct ends with “Explore it yourself,” linking to raw responses and ratings. SWE-bench's leaderboard has an Agent dropdown that changes the comparison class; its overview page explains the bash-only configuration before directing readers to the control. ARC links methodology to scorecards, recordings, replays, public game sets, and benchmarking repositories. Artificial Analysis links individual datasets and implementations and exposes evaluated task names.

**Implication:** Treat **Browse scenarios**, **View tasks**, **Inspect runs**, and **Download artifacts** as evidence links adjacent to the relevant section. Do not require a complex client-side explorer for the first version; static generated indexes and deep links preserve GitHub Pages compatibility.

### Finding 12: Citations should be proximal, with a compact reusable citation block at the end

Method claims are easiest to verify when the dataset, paper, repository, commit, or implementation is linked where the claim is made. A final citation block serves academic reuse but should not be the reader's only path to source evidence.

**Evidence:** Artificial Analysis links datasets, papers, pages, repositories, and specific implementation files inside each evaluation record. HELM uses numbered inline citations, links the open-source implementation where the procedure is described, and provides full Citation and References sections. SWE-bench places Paper and GitHub at the top, links exact configurations in the relevant disclosure, and supplies BibTeX/APA/MLA formats with copy controls. ARC docs link the docs repository and page edit target in addition to surrounding benchmark documentation.

**Implication:** Link exact repositories, files, versions, commits, datasets, and artifact indexes in context. Add a final **Cite this methodology** block only if the page has a stable version or release identity. A generic GitHub icon in the footer is insufficient evidence provenance.

### Finding 13: Version history and status labels are core interface elements for living methodology

Evaluation methods change as tasks, models, graders, harnesses, or statistical choices change. Public pages earn trust by naming the current version, distinguishing active from legacy material, and describing changes in terms of comparability.

**Evidence:** Artificial Analysis shows the current index version near the title, labels standalone and legacy evaluations, states when an evaluation was superseded, and maintains a dated version history. SWE-bench explains how mini-SWE-agent release numbers map to setup changes and warns when results across major versions are not necessarily comparable. METR labels its current time-horizon method, publishes update notes, and explains changes in task suite and infrastructure. Anthropic's system-card index records model and date for every release.

**Implication:** Show **Current methodology version**, **last updated**, and **status** near the top. End with a concise change log that says whether each change preserves comparability. Retired methods should remain addressable if past public results depend on them.

### Finding 14: The best restrained GitHub Pages pattern is progressive, static-first disclosure

A strong small-site implementation can combine a narrow article column, sticky or inline table of contents, semantic headings, native details, responsive tables, one diagram, and deep links to static artifacts. Search, AI assistants, animated dashboards, and large custom visualization systems are optional rather than required for authority.

**Evidence:** ARC's docs interface demonstrates the value of on-page navigation, copy controls, semantic prose, formulas, and responsive tables, but its integrated search/assistant layer exceeds what a simple methodology page needs. SWE-bench demonstrates that native details, direct GitHub links, and citation tabs work in a static site. Artificial Analysis demonstrates an ideal prose width and strong version/source records, though its very long single page would benefit from splitting per-evaluation detail. HELM demonstrates deep result browsing, while its research-paper hierarchy and dense references are heavier than necessary for a concise public method.

**Implication:** Start with semantic HTML and durable URLs. Recommended interface primitives are: an inline or sticky contents list; 60–75ch prose; wider responsive data blocks; anchored headings; visible callouts for scope/limitations; native details for optional depth; and direct repository/artifact links. Add interactivity only where it changes what evidence the reader can inspect.

### Finding 15: Lower-fit counterexamples clarify what to avoid

System-card PDFs, paper-style methodology posts, extremely wide dashboard pages, and generic institutional news pages each solve a different problem. They are valuable companion sources but weak templates for a focused public evaluation-method page.

**Evidence:** OpenAI and Anthropic system cards are dozens to hundreds of pages and mix model description, safety policy, results, deployment decisions, and appendices; their PDF format reduces deep-linking and skimmability for a single method. HELM Instruct uses long numbered research sections, image-based tables, and a full bibliography, appropriate for a research publication but heavy for routine reference. METR's 1200px chart surfaces are appropriate for interactive results but would be too wide for continuous method prose. NIST's public news page includes substantial institutional navigation, publication metadata, sharing controls, newsletter, and footer material around a short explanation.

**Implication:** Use these sources for specific components—formal limitations, citations, figures, or publication metadata—rather than copying their complete page shape. A public method page should remain focused enough that a reader can learn the evaluation boundary, execution, score meaning, artifacts, limitations, and current version without navigating an entire report or institutional site.

## Notes

- “Benchmark Card” appears as an emerging documentation concept, but the most authoritative inspected card guidance was model-card oriented. Treat the term as schema-driven and conditional unless a benchmark-card standard is explicitly adopted.
- Exact rendered line length varies with viewport and fonts. The 60–75 character recommendation is an inference from the inspected constrained prose layouts, not a value published by the source organizations.
- Some benchmark explorers are client-rendered and were only partly exposed to text indexing. Claims here are limited to visible inventories, documented controls, linked artifact surfaces, and live HTML/CSS inspection.
- Useful search terms for further comparison: `evaluation methodology benchmark`, `scoring methodology benchmark`, `testing methodology evaluation policy`, `benchmark card documentation`, `system card evaluation methodology`, and `technical report benchmark methodology`.

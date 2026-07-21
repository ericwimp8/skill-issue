# Testing Methodology Page Progress

## Purpose

Track the complete research, verification, content design, repository organization, implementation, and validation work for a public testing methodology page. The page must explain how Skill Issue builds and hardens the evaluated skills, runs controlled conversational scenarios, records skill invocations, scores the observed calls, and publishes enough source material for readers to inspect the process themselves.

This document is the progress owner for that page. It preserves the project owner's account of the method as the starting contract, then requires every public factual claim to be traced through the repository before it is published. Research findings belong in `research/`; implementation decisions belong at their source-code owner; completed checks and remaining work stay recorded here.

## Current Status

- [x] Capture the complete requested scope in a dedicated progress document.
- [x] Record the explicitly invoked Dictate Plan skill as a required methodology disclosure.
- [x] Complete local system research and source validation.
- [x] Complete external methodology-page research.
- [ ] Design and approve the page's information architecture and copy.
- [ ] Organize canonical repository material for public linking.
- [x] Implement the page and its scenario and skill readers.
- [ ] Complete technical, interactive, responsive, and visual validation.

## Verified Research Checkpoint

- The current development working tree and development CLI channel are the authority for this work. The selected known-good build, committed baseline, and staged 12-turn snapshot are stale and are excluded from methodology decisions, campaign preparation, and result interpretation.
- The current development scenarios contain 30 turns each. Their current scorecards contain 46, 46, and 45 unique expected turn-skill pairs, with expectations on 27 turns per scenario. These totals remain provisional until the active scenario edits are frozen and rechecked.
- Thirty currently describes both the number of turns in each scenario and the planned campaign total of 30 runs across ten configurations and three scenarios.
- Dictate Plan is named manually once in Turn 1 of every scenario and is included in the scorecard. Codex metadata makes it explicit-only, while equivalent technical non-implicit enforcement is not established for every harness.
- The instrumentation token's meaning is opaque outside evaluator-owned state, but the injected command and its side effects remain observable and must be disclosed as a possible influence.
- Harness isolation is substantial but adapter-specific. The public page must describe concrete controls and residual host, authentication, configuration, tool, filesystem, and network surfaces rather than claim one uniform sandbox.
- All eight current scenario skills have bounded retained refinement evidence. Complete intake-to-generation lineage is not retained for every skill.
- The public accepted artifact collection is empty and the campaign remains incomplete. Real result links stay absent until accepted development-baseline runs and retained evidence exist.
- External research supports the audience-facing name **Evaluation Methodology** and a restrained static-first reading flow with linked evidence depth. Exact layout, schemas, and disclosure policies remain Skill Issue design decisions rather than requirements copied from another benchmark.
- Local synthesis: [`research/testing-methodology/local-system-research.md`](../../research/testing-methodology/local-system-research.md).
- External synthesis: [`research/testing-methodology/external-methodology-page-research.md`](../../research/testing-methodology/external-methodology-page-research.md).

## Working Rules

- [ ] Treat the project owner's methodology description below as the starting account, while marking each publishable claim verified only after checking its production source, governed scenario material, retained artifact, or controlling repository document.
- [ ] Trace relevant behavior end to end. Do not stop at interfaces, wrappers, test expectations, summaries, or website fixtures.
- [ ] Use `./cli/scripts/local-cli.sh development` for current methodology and campaign verification. Do not treat the presently selected known-good executable as a fallback, comparison baseline, or corroborating source.
- [ ] Use production source and governed repository artifacts as behavioral truth. Use tests only after intended behavior has been established.
- [ ] Validate all agent-produced findings against the repository before they enter the research synthesis, page copy, or website data.
- [ ] Keep unsupported statements visibly unresolved in research rather than smoothing them into confident public claims.
- [ ] Use real scenarios, skills, scorecards, evaluation material, and links only. Do not add mock labels, invented examples, fictional results, placeholder claims, or copy suggesting that displayed evidence is simulated.
- [ ] Leave unfinished evaluation links absent until their real destinations exist. Add them as soon as their retained folders are ready.
- [ ] Keep the page informational, restrained, and reading-oriented. Follow the Project Purpose and Limits page's editorial direction, with a potentially drier and denser treatment suited to methodology documentation.
- [ ] Preserve the established Skill Issue typography, grid, spacing language, navigation, theme behavior, and GitHub Pages constraints without turning the methodology into another landing page.
- [ ] Keep headings smaller than promotional page headings, paragraphs comfortably readable, section spacing compact and consistent, and decorative boxes or split-view layouts to a minimum.
- [ ] Do not begin page implementation until both research tracks, source validation, information architecture, copy structure, and canonical link ownership are complete.

## Starting Methodology Account To Verify

### Scenario design and scoring

- [x] Verify that the campaign uses three governed scenarios.
- [x] Verify that every current development scenario contains 30 conversational turns.
- [x] Verify the current expected-call counts of 46, 46, and 45, while requiring a final recount after the active scenario edits are frozen.
- [x] Verify that expected calls are distributed across different points and conversational conditions rather than concentrated in one part of a scenario.
- [x] Verify the exact source format for every scenario's ordered user and assistant exchange.
- [x] Verify the exact scorecard format that maps turns to expected skills.
- [x] Verify how a turn containing multiple expected skill calls is represented and scored.
- [x] Verify how recorded invocations are matched against the scorecard and how called, missed, extra, duplicate, or otherwise unmatched observations are handled.
- [x] Verify that all expected turn-skill pairs receive a Boolean called or missed outcome in the compact public result.
- [x] Verify the result units: detailed artifacts retain expected and observed turn-skill pairs; compact website artifacts retain per-turn called and missed expected-pair totals; aggregates sum expected-pair outcomes.

### Controlled conversation execution

- [x] Trace how the CLI starts a scenario and conducts the controlled back-and-forth conversation.
- [x] Verify how each ordered scenario message is passed to the model and harness.
- [x] Verify where conversation state is retained across the 30 turns.
- [x] Verify which parts of the run are deterministic, governed, configurable, or dependent on the selected harness and model.
- [ ] Verify the exact model, reasoning, harness, executable, workspace, and output settings retained for publication.
- [x] Verify how failed, interrupted, recovered, retried, or incomplete runs affect acceptance and publication; retain the cleanup-before-acceptance gap as an unresolved publication concern.

### Invocation instrumentation and opaque identifiers

- [x] Locate the script or command injected into each evaluated skill and trace its concrete recording effect.
- [x] Verify where and how the instrumentation is inserted at the start of a skill.
- [x] Verify the instrumentation's audience-visible command shape and neutral `signal` naming.
- [x] Verify how the injected call avoids naming itself as evaluation or scoring machinery.
- [x] Verify how opaque identifiers or tokens are assigned to skills.
- [x] Trace how an opaque observation is translated back to the corresponding skill outside the evaluated model context.
- [x] Verify that token meaning is hidden while the injected command, executable path, state path, process, and permissions can remain observable.
- [x] Verify what the instrumentation records, where it records it, and how the evaluator retrieves it.
- [x] Verify that a successful signal emits no command output while the instruction and process side effects remain observable.
- [ ] State transparently on the page that adding instrumentation can influence model behavior, even when the command and identifier are designed to be unobtrusive.
- [ ] Explain this on the page as a bounded methodological limitation rather than claiming that the instrumentation is perfectly invisible or influence-free.

### Explicit Dictate Plan invocation

- [x] Locate the Dictate Plan skill used by the governed scenarios and verify its exact public name and source.
- [x] Verify that Dictate Plan is explicit-only through Codex metadata and record that equivalent enforcement is unsupported for other harnesses.
- [x] Verify that it is manually invoked exactly once at the start of each relevant scenario or run.
- [x] Verify how that explicit call is represented in the conversation, instrumentation log, scorecard, result artifact, and website data.
- [x] Verify that the explicit Dictate Plan invocation is included in scored expected calls for Turn 1.
- [ ] Explain on the methodology page that this invocation is intentional, manual, and different from the later skill calls whose discovery behavior is being evaluated.
- [ ] Make the disclosure prominent enough that readers do not mistake the initial call for an unexplained observation, hidden expected call, or evidence of spontaneous skill discovery.

### Isolated harness and model environments

- [x] Trace how each local evaluation workspace is created and isolated.
- [x] Verify that evaluated harnesses start as close to their defaults as the project can practically make them.
- [x] Verify that project skills, personal skills, inherited agents, unrelated instructions, and ambient configuration are excluded where claimed.
- [x] Verify the treatment of `AGENTS.md`, `CLAUDE.md`, equivalent harness instruction files, and parent-directory instructions.
- [x] Verify which necessary evaluator-owned instructions, skills, files, environment variables, credentials, managed policy, or configuration remain present.
- [x] Verify that every compared harness and model receives equivalent governed scenario content and evaluation machinery, subject to adapter-specific native behavior.
- [x] Verify how sandboxing is configured for each harness and why one uniform sandbox claim is unsupported.
- [x] Verify which defaults differ inherently between harnesses and therefore cannot be normalized away.
- [ ] Explain on the page the limits of calling an environment "default" when evaluator-required controls and native product surfaces remain present.

### Skill construction, evaluation, and refinement

- [x] Trace how the public skill-building workflow turns a natural-language request into a generated skill.
- [x] Trace how the workflow evaluates invocation behavior through the skill description.
- [x] Trace how it evaluates whether the skill body performs its intended task correctly.
- [x] Verify how failures cause description, body, supporting files, or surrounding skill-system changes.
- [x] Verify the current per-loop five-failure stopping rules and distinguish them from the older campaign-specific 40-iteration authority.
- [x] Verify the distinction and handoff between skill generation and skill evaluation and refinement.
- [x] Inventory the documentation, reports, results, scripts, references, and other artifacts currently retained during the process.
- [ ] Identify the exact retained folder that will be published for each evaluated skill.
- [ ] Explain that the campaign uses this hardening process to make a good-faith effort to begin with functioning skills before measuring harness and model invocation reliability.
- [ ] Avoid claiming that refinement proves universal correctness or eliminates every possible skill-side failure.
- [ ] Explain that the same workflow is public, installable through the CLI, and available for readers to build and refine their own skills.
- [ ] Link this explanation back to the Build Skills website view.

### Transparency, constraints, and limitations

- [ ] Explain the project's efforts to make skills functional, scenarios scoreable, environments comparable, and evaluation signals difficult to game.
- [ ] Explain that the complete method, governed scenarios, scorecards, skills, and retained evaluation evidence are being made reviewable.
- [ ] State the practical limits created by one person conducting the work, paying for the runs, working with limited time and resources, and completing the campaign during Build Week.
- [ ] State the resulting limits on test breadth, repeated trials, statistical confidence, model coverage, harness coverage, configuration coverage, and generalizability where source evidence supports those limits.
- [ ] Distinguish methodological transparency from a claim that the method is flawless.
- [ ] Keep every conclusion tied to the exact tested configurations, accepted runs, scorecards, and retained evidence.

## Phase 1: Establish Research Ownership And Evidence Maps

- [x] Create a local-system research document under `research/` that maps every methodology claim to production source, governed input, retained artifact, or unresolved evidence need.
- [x] Create an external methodology-page research document under `research/` that records source URLs, page structures, useful conventions, and applicability to Skill Issue.
- [x] Create a claim-validation table or equivalent evidence map that distinguishes project-owner statements, source-verified facts, inferences, limitations, and unresolved claims.
- [x] Record exact file paths, symbols, commands, artifact fields, and relevant line references so the final synthesis can be independently checked.
- [x] Keep research documents descriptive and evidence-backed without making them alternative owners for production behavior or website copy.

## Phase 2: Deep Local System Research

- [x] Use approximately eight to ten bounded local research assignments, run in parallel waves within the available concurrency limit.
- [x] Give every agent a self-contained prompt, explicit source-code scope, required evidence format, non-overlapping question, and instruction to avoid edits.
- [x] Assign independent research for scenario definitions and scorecards.
- [x] Assign independent research for CLI scenario orchestration and conversational turn handling.
- [x] Assign independent research for invocation instrumentation and opaque identifier translation.
- [x] Assign independent research for workspace creation, sandboxing, ambient configuration exclusion, and harness defaults.
- [x] Assign independent research for scoring, aggregation, acceptance, and website artifact production.
- [x] Assign independent research for skill generation.
- [x] Assign independent research for skill evaluation and refinement.
- [x] Assign independent research for retained skill-evaluation artifacts and planned GitHub publication surfaces.
- [x] Assign independent research for Dictate Plan's explicit-only configuration, manual invocation, logging, and scoring treatment.
- [x] Assign independent research for the exact skills used across all three scenarios and their canonical repository locations.
- [x] Aggregate the reports without weakening qualifications, conflicts, or unresolved findings.
- [x] Re-open and validate every aggregated claim against the cited repository source before accepting it.
- [x] Resolve the staged 12-turn versus development 30-turn conflict through direct source tracing and project-owner authority rather than majority agreement.
- [x] Record unresolved publication, baseline-freeze, configuration, confinement, provenance, schema, privacy, and stable-link requirements.

## Phase 3: Deep External Methodology-Page Research

- [x] Use the deep-research workflow with ten bounded external research assignments run in parallel waves within the available concurrency limit.
- [x] Find strong benchmark methodology pages, model or agent evaluation disclosures, reproducibility pages, benchmark cards, system cards, and transparent testing write-ups across GitHub and the wider web.
- [x] Prefer primary sources from benchmark authors, research organizations, official repositories, and established evaluation projects.
- [x] Study selected examples closely enough to understand their full reading order rather than collecting isolated screenshots or headings.
- [x] Record how good examples introduce purpose and scope before technical detail.
- [x] Record how they explain datasets or scenarios, evaluation units, instrumentation, scoring, controls, reproducibility, limitations, and source access.
- [x] Record how they distinguish designed expectations from observed results.
- [x] Record how they disclose manual steps, evaluator intervention, exclusions, and possible sources of bias.
- [x] Record how they present dense prose, tables, expandable source material, diagrams, code, and links without turning the page into a promotional landing page.
- [x] Study heading depth, paragraph length, sentence rhythm, punctuation, readable measure, section spacing, inline definitions, labels, notes, and citation placement.
- [x] Record screenshot requirements for later visual implementation comparison where text research alone is insufficient.
- [x] Identify **Evaluation Methodology** as the strongest conventional audience-facing page name for this project.
- [x] Synthesize applicable patterns while separating source facts, methodological implications, and Skill Issue design recommendations.

## Phase 4: Validate The Complete Research Synthesis

- [x] Compare the local research synthesis against every point in the Starting Methodology Account To Verify section.
- [x] Compare the external research synthesis against every page-structure and writing requirement in this tracker.
- [x] Check that every user-supplied detail has a verified fact, a qualified statement, or a clearly recorded unresolved evidence need.
- [x] Perform a second end-to-end source walkthrough from scenario input to published website artifact.
- [x] Perform a second end-to-end source walkthrough from skill request to generated, evaluated, refined, and retained skill package.
- [x] Verify the current development inventory of scenarios, scorecards, skills, expected calls, scripts, and public artifacts; retain exact final run configurations as an unresolved campaign-freeze requirement.
- [x] Verify that the Dictate Plan disclosure remains present and accurate after the broader synthesis.
- [x] Reject or rewrite proposed public claims that overstated uniform isolation, invisible instrumentation, frozen scenario counts, complete generation lineage, reproducibility, or statistical significance.
- [x] Mark the research phases complete after making every publishable claim source-backed, project-owner-directed, or explicitly qualified.

## Phase 5: Design The Page Information Architecture

- [ ] Choose the page name from the external research and Skill Issue navigation context.
- [ ] Decide the page's position and label in the shared navigation without displacing the agreed Explore, Analysis, and Project destinations accidentally.
- [ ] Design one continuous, information-first reading flow rather than a sequence of marketing panels.
- [ ] Begin with a concise explanation of what is being tested and why the methodology is public.
- [ ] Explain the complete evaluation path in a reader-comprehensible order: hardened skills, governed scenarios, controlled environments, instrumentation, scorecards, scoring, retained evidence, and limitations.
- [ ] Give the explicit Dictate Plan call its own clear note in the part of the flow where scenario startup is explained.
- [ ] Decide where a compact process diagram genuinely improves understanding and omit it if prose communicates the method more clearly.
- [ ] Design a Scenarios section that introduces the common 30-turn structure before exposing individual scenario material.
- [ ] Design an Evaluated Skills section that explains generation and refinement before exposing individual skill material.
- [ ] Design a Reproducibility and Evidence section that links readers to canonical repository material.
- [ ] Design a Limitations section that is visible, specific, and integrated into the methodology rather than hidden in fine print.
- [ ] Keep heading sizes, paragraph widths, dividers, vertical rhythm, and information density consistent with an editorial reference page.
- [ ] Avoid unnecessary cards, oversized display type, decorative statistics, split-screen compositions, or excessive blank space.
- [ ] Review the proposed structure against captured references and the existing Project page before writing final copy.

## Phase 6: Draft And Verify Audience-Facing Copy

- [ ] Draft from the verified evidence map and approved information architecture rather than directly paraphrasing raw agent summaries.
- [ ] Preserve the project owner's direct, transparent voice.
- [ ] Use plain punctuation and natural sentence structure. Avoid generic AI phrasing and unnecessary dash punctuation.
- [ ] Define scenario, turn, expected call, scorecard, observed invocation, run, harness, model, and reasoning level consistently.
- [ ] Explain technical mechanics precisely without overwhelming the main reading path with implementation trivia.
- [ ] Explain why the injected instrumentation exists, how opacity is attempted, and how instrumentation itself remains a limitation.
- [ ] Explain the Dictate Plan exception exactly once at the semantic point that owns scenario startup, with only local reminders elsewhere if needed.
- [ ] Explain how skill generation and refinement test both invocation descriptions and skill-body behavior.
- [ ] Explain that readers can install the CLI and use the same public workflow, then link to Build Skills.
- [ ] Explain how scenarios, scorecards, skills, and retained evaluation material can be inspected on the site and GitHub.
- [ ] Keep unfinished links absent without inserting mock destinations, coming-soon evidence cards, or invented folder names into the public page.
- [ ] Complete a claim-by-claim source check of the final draft.
- [ ] Complete a whole-page coherence review for repetition, contradictions, unexplained terms, missing qualifications, and misplaced disclosures.

## Phase 7: Organize Canonical Repository Material

- [x] Inventory the current canonical location of each governed scenario's instructions, conversation, expected calls, and configuration.
- [x] Inventory the current canonical location of each evaluated skill and its supporting files.
- [x] Identify the retained evaluation and refinement folder for each skill once its current run completes.
- [x] Decide which source material remains in its existing semantic home and which material requires deliberate publication organization.
- [ ] Keep disposable CLI output under ignored `output/`; move only explicitly retained review evidence into an appropriate tracked location.
- [x] Define stable repository-relative and GitHub URLs for every scenario, scorecard, skill, and retained evaluation folder.
- [x] Ensure the website data references canonical sources rather than duplicating manually maintained copies without an ownership rule.
- [ ] Verify that published folders do not expose credentials, machine-specific paths, private configuration, disposable recovery state, or unrelated local data.
- [x] Verify every GitHub link against the intended repository, branch behavior, path casing, and GitHub Pages deployment context.
- [x] Record which evaluation links remain blocked because their real retained folders are still being generated. None remain blocked at the current remote `main` revision.

## Phase 8: Build The Scenario Presentation

- [x] Add one real entry for each governed scenario.
- [x] Show the scenario name and a concise verified description.
- [x] Show the verified turn count and expected-call count.
- [x] Let readers open each scenario inside the website using the established overlay or reader interaction used by the Build Skills view, adapted only where scenario content requires it.
- [x] Present every scenario message in exact conversational order.
- [ ] Clearly distinguish speaker, turn number, message content, and expected calls.
- [x] Annotate each relevant turn with the skill or skills expected by the scorecard.
- [ ] Represent turns with no expected calls accurately without implying that those turns are unscored expected calls.
- [ ] Represent multiple expected calls on one turn without collapsing their individual Boolean outcomes.
- [x] Provide access to the scorecard or its exact governed content.
- [x] Add a direct GitHub link to the scenario's canonical folder and scorecard.
- [ ] Make the overlay readable and navigable for long 30-turn content on desktop and narrow screens.
- [ ] Provide accessible close behavior, focus management, keyboard navigation, labels, and external-link treatment.

## Phase 9: Build The Evaluated Skills Presentation

- [x] Inventory and display every real skill used by the governed scenarios.
- [x] Identify skills shared across scenarios without creating duplicate ownership or misleading counts.
- [x] Reuse the existing Build Skills skill-reader behavior and visual language where it remains suitable.
- [x] Let readers open and read each real `SKILL.md` inside the website.
- [x] Show a concise verified explanation of what each skill is for.
- [x] Identify Dictate Plan as explicit-only and manually invoked at scenario startup.
- [x] Distinguish that startup exception from skills whose discovery and invocation are evaluated later in the scenario.
- [x] Add a direct GitHub link to each skill's canonical folder.
- [x] Add a direct GitHub link to each skill's retained evaluation and refinement folder only when the real folder exists.
- [x] Make the link target clear enough that readers know whether they are opening the skill source or its evaluation evidence.
- [ ] Link the section back to Build Skills and the CLI installation path.
- [ ] Explain that readers can use the public generation and refinement workflow for their own skills.

## Phase 10: Implement Website Navigation, Data, And Page Structure

- [ ] Add the new destination through the existing GitHub Pages-compatible hash-routing system.
- [ ] Update shared navigation without weakening the far-right placement and meaning of Project Purpose and Limits.
- [ ] Keep curated page copy and navigation metadata at their established website data owner.
- [ ] Define methodology-specific data types and adapters at the source owner established during research.
- [ ] Keep scenario and skill source adaptation deterministic and reviewable.
- [ ] Avoid hard-coding facts in multiple components when one canonical data definition can own them.
- [ ] Build the page as a narrow, continuous editorial reading surface with compact repeated sections.
- [ ] Keep headings bounded at large viewport widths.
- [ ] Keep body copy readable at desktop, tablet, and narrow mobile widths.
- [ ] Preserve both themes, focus treatment, browser history, direct hash links, refresh behavior, and GitHub Pages compatibility.
- [ ] Ensure unfinished evaluation evidence does not create broken controls or dead links.

## Phase 11: Technical And Evidence Validation

- [ ] Run `npm run typecheck` after TypeScript or TSX work.
- [ ] Run `npm run lint` after TypeScript or TSX work.
- [ ] Run `npm run format:check` after CSS, Markdown, JSON, HTML, or configuration work.
- [ ] Run targeted tests or deterministic artifact checks for scenario adaptation, expected-call annotations, link construction, and routing where the production behavior warrants them.
- [ ] Run `npm run validate` before presenting the website work as complete.
- [ ] Record any unrelated repository-wide validation blocker precisely without treating targeted success as complete validation.
- [ ] Verify every displayed scenario message against its canonical source.
- [ ] Verify every displayed expected-call annotation against its canonical scorecard.
- [ ] Verify every displayed skill against its canonical source.
- [x] Verify every active GitHub link by opening its real destination.
- [ ] Verify the public copy one final time against the completed evidence map.

## Phase 12: Interactive, Responsive, And Visual Validation

- [x] Run the website with `npm run dev -- --host 127.0.0.1` and confirm live reload.
- [ ] Open the methodology page through shared navigation and through its direct hash route.
- [ ] Click through every scenario reader, skill reader, internal link, GitHub link, close control, and navigation state.
- [ ] Inspect desktop, tablet, and narrow mobile layouts in both themes.
- [ ] Confirm that long scenario conversations remain readable and controllable without horizontal overflow or trapped focus.
- [ ] Confirm that the page reads naturally from top to bottom and does not feel like a collection of promotional boxes.
- [ ] Capture screenshots of the complete page, scenario reader, skill reader, narrow layout, and any materially different themed state.
- [ ] Compare the implementation against the selected external references and the Project page's established editorial language.
- [ ] Use the image-comparison workflow when typography, spacing, hierarchy, reading measure, overlays, or responsive behavior appears uncertain.
- [ ] Fix every confirmed visual or interaction mismatch at its semantic owner.
- [ ] Inspect the production build with `npm run preview` after local interactive checks pass.

### Browser Link Verification, 2026-07-21

- [x] Open all three scenario readers and confirm each contains 30 ordered turns, including Turn 1 and Turn 30.
- [x] Open each scenario's canonical GitHub folder and confirm the destination renders successfully.
- [x] Open all eight evaluated skill readers and confirm each displays its real `SKILL.md` content.
- [x] Open all eight canonical skill-source folders and confirm the destination renders successfully.
- [x] Open all eight retained evaluation folders and confirm the destination renders successfully.

## Completion Gate

The testing methodology page is complete only when:

- [ ] Every public methodology statement is source-backed or explicitly qualified.
- [ ] The explicit Dictate Plan startup invocation is accurately disclosed and cannot be confused with spontaneous evaluated skill discovery.
- [ ] All three real scenarios, their ordered messages, expected calls, and scorecards are inspectable.
- [ ] Every real evaluated skill is inspectable and linked to its canonical source.
- [ ] Real retained skill-evaluation links are present where available, and no fictional or broken evidence links are shown.
- [ ] The page explains skill hardening, controlled execution, instrumentation, scoring, isolation efforts, reproducibility, and limitations as one coherent method.
- [ ] Readers can reach Build Skills and understand that the public CLI supports the same skill-building and refinement workflow.
- [ ] Canonical GitHub material is organized, safe to publish, and linked correctly.
- [ ] Required typecheck, lint, formatting, build, validation, browser, accessibility, responsive, and visual checks are complete.
- [ ] The page contains no mock language, invented evidence, unsupported certainty, or unfinished placeholder claims.

## Deferred Until Real Evaluation Material Exists

- [ ] Add each skill's retained evaluation and refinement link after its current real evaluation folder is complete, reviewed, and placed at its canonical tracked location.
- [ ] Re-run link, copy, and visual validation after those real evidence links are added.

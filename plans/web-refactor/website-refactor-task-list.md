# Website Refactor Task List

## Purpose

Track the website behavior, presentation, information architecture, content, and evaluation-data corrections needed before the public site is ready for real campaign evidence. The results experience will separate harness comparison from model comparison so each chart answers one meaningful question. The site will preserve the current landing experience and add dedicated project-purpose and results-analysis destinations. Those three destinations are the complete agreed page scope for this refactor. This document owns the website work queue. Evaluation campaign configuration remains owned by the campaign progress document.

## Current Review Sources

- [Skill-calling evaluation campaign matrix](../skill-calling-evaluation-campaign/evaluation-progress.md#campaign-configuration)
- [Website evaluation data adapter and illustrative cells](../../src/data/evaluationData.ts)
- [Project README](../../README.md)
- [Existing website reference findings](../../research/website-two-arm-navigation/reference-findings.md)

## Current Implemented Baseline

- The landing page already presents the Build Skills and Evaluate Environments product tabs inside the shared Skill Issue visual system.
- The Build Skills view already explains the generation-and-refinement workflow and exposes the current generated-skill showcase.
- The Evaluate Environments view already uses the confirmed ten-configuration matrix as illustrative data. The older Claude Opus 4.8 cells have been removed; Fable and Cursor — Composer are present.
- The current model-comparison explorer contains three full-width charts: Outcome Strips, Expected-Call Curve, and Success Ranking.
- The comparison selector exposes all ten harness-and-model cells with model, harness, and Medium reasoning labels. The default comparison holds Cursor constant across Fable, Codex, Grok, and Composer.
- Outcome Strips and Expected-Call Curve use one selected scenario. Success Ranking independently aggregates any non-empty combination of the three governed scenarios.
- Successful outcomes use consistent model colors, missed outcomes use neutral gray, and the line-chart legend repeats each model's marker shape and color.
- Recharts tooltip animation and wrapper transitions are disabled on the charts with hover content so tooltips remain connected to the active datum.
- The large product-view introduction headings use bounded responsive sizing and no longer grow beyond `4.2rem` on large displays.
- All evaluation data remains illustrative mock data derived through the website artifact adapter. Accepted campaign evidence has not yet replaced it.
- The shared website shell now exposes three destinations named Explore, Project, and Analysis. Explore retains the Build Skills and Evaluate Environments product tabs.
- Hash-based routes preserve direct links, refresh behavior, browser history, and GitHub Pages compatibility without server rewrites.
- The Project page adapts the README's motivation, two-part purpose, current progress, and limitations into public copy. The Analysis page remains an evidence-gated coming-soon destination.
- The results explorer now separates model comparison from a Codex-only harness comparison with its own non-empty scenario aggregation control.
- Accepted compact CLI artifacts can be loaded into the static site with `npm run results:update -- <website.json files>`; an empty published collection keeps the illustrative campaign visible.

## Tasks

### 1. Finish tooltip behavior validation

- [x] Trace the sideways tooltip movement to Recharts tooltip animation and wrapper transitions.
- [x] Disable tooltip animation and wrapper transitions on the Expected-Call Curve and Success Ranking charts.
- [x] Remove references to the deleted Checkpoint Profile chart from the active chart design.
- [x] Verify the corrected hover behavior on the live desktop results view.
- [ ] Verify tooltip position and readability in both themes and at narrow responsive widths.

### 2. Finish campaign identity evidence

- [x] Review the current ten-configuration campaign matrix with the project owner.
- [x] Resolve the older thirteen-cell research description in favor of the current ten-configuration campaign progress matrix.
- [x] Confirm Fable replaces the obsolete Claude Opus 4.8 website cells.
- [x] Confirm Cursor — Composer, OpenCode — Codex, and Kilo Code — Codex remain in the campaign.
- [ ] Record the exact effective provider model identifier, model version or alias, harness version, and reasoning setting for every accepted result.
- [ ] Replace illustrative display identities with accepted-result identities wherever the retained campaign evidence differs from the current labels.

### 3. Complete the model comparison results view

- [x] Use one comparison multi-select containing the confirmed ten harness-and-model combinations.
- [x] Show model, harness, and Medium reasoning information in every comparison option and selected chip.
- [x] Remove the redundant standalone Harness and Model filters from the model-comparison controls.
- [x] Present the three model-comparison charts as full-width rows rather than a mixed full-width and two-column chart grid.
- [x] Keep Outcome Strips as a scored-outcome timeline for one selected scenario, with model-colored calls and neutral-gray misses.
- [x] Replace the turn-based cumulative chart with an Expected-Call Curve whose horizontal axis contains only expected calls.
- [x] Increment each cumulative line for a successful call and keep it flat for a miss.
- [x] Give every model a distinct marker shape and color, and repeat the same visual identity in the line-chart legend.
- [x] Aggregate Success Ranking across any non-empty scenario combination selected through its chart-local checkbox menu.
- [x] Rank the selected combinations from highest to lowest success rate using model-colored success and neutral-gray failure segments.
- [x] Expand multi-call turns into individual Boolean outcomes so every expected call remains represented.
- [ ] Replace the illustrative artifacts with accepted Work Block 3 website artifacts without changing the chart-facing adapter contract.
- [ ] Confirm the final public default comparison after accepted campaign results are available.

### 4. Create the harness comparison results view

- [x] Add a distinct Harness Comparison view beside the implemented Model Comparison explorer.
- [x] Hold the model constant at Codex with Medium reasoning for the harness comparison. Use only the confirmed harness configurations that run Codex so the chart isolates the effect of changing the harness.
- [x] Present one harness-comparison chart only. Do not add secondary harness charts unless they answer a separately confirmed question.
- [x] Use a horizontal 100% stacked bar chart with one bar per harness. Show successful and failed expected skill calls as complementary percentages within each bar.
- [x] List the harness bars vertically from highest success rate at the top to lowest at the bottom so the chart immediately communicates which harnesses were more or less effective at skill calling with the same model.
- [x] Give the chart its own scenario multi-select control, using a dropdown with one checkbox for each of the three governed scenarios.
- [x] Select all three scenarios by default. Allow any non-empty combination: one scenario alone, any pair of scenarios, or all three scenarios together.
- [x] Treat each selected scenario as one complete 30-turn conversation. The selected conversation context therefore represents 30 turns for one scenario, 60 turns for two scenarios, or 90 turns for all three scenarios.
- [x] Calculate chart performance from the scored expected skill activations rather than treating every conversation turn as a call opportunity. The current governed fixtures contribute four scored activations per scenario, giving 4, 8, or 12 scored outcomes per harness when one, two, or three scenarios are selected.
- [x] Apply the same selected scenario set to every harness so the ranked comparison always uses equivalent evidence.
- [x] Aggregate selected scenarios by summing called and missed expected activations for each harness, then derive the displayed success and failure percentages from those totals.
- [x] Recalculate percentages and ranking immediately when the selected scenario set changes.
- [x] Keep the selected scenario count, represented conversation turns, and underlying called, missed, and scored-total counts available in labels or hover content so the percentages remain interpretable.
- [x] Verify that the finished chart tells one clear story: how much the harness changes reliable skill calling when the Codex model and Medium reasoning target remain constant.

### 5. Replace illustrative evaluation data with accepted evidence

- [x] Remove obsolete model labels and illustrative cells after the campaign matrix is confirmed.
- [x] Add every confirmed harness, model, and reasoning combination needed by the current results controls and charts.
- [x] Update summary metrics to seven illustrative harnesses and ten comparison cells.
- [ ] Keep illustrative outcomes clearly labelled as mock data until accepted campaign artifacts replace them.
- [x] Verify the comparison selector exposes every confirmed cell without duplicates or unsupported combinations.
- [x] Add the repository-owned ingestion or update path that makes accepted `website.json` artifacts easy to publish through the static GitHub Pages build.
- [ ] Replace mock outcomes only with acceptance-layer-approved Work Block 3 artifacts and preserve the retained run/configuration provenance.

### 6. Research the multi-page navigation pattern

- [x] Review the existing website-reference findings, retained screenshots, and current Skill Issue implementation before selecting a navigation pattern.
- [x] If any earlier reference artifact is missing, locate it in the working tree or Git history and recover only the relevant research without reverting unrelated work.
- [x] Revisit the established reference websites and inspect their secondary pages as well as their homepages. Focus on compact menus, page-to-page navigation, active destination states, shared headers, mobile navigation, content hierarchy, typography, spacing, punctuation, and transitions between marketing and results content.
- [x] Capture new reference screenshots whenever the retained research does not make an important layout or navigation behavior clear.
- [x] Compare navigation approaches that preserve the current minimal Skill Issue visual language and work reliably on a static GitHub Pages deployment.
- [x] Decide the audience-facing names for the landing, project-purpose, and results-analysis destinations after the reference review.
- [x] Decide whether the static site should use hash-based routing or another GitHub Pages-compatible direct-navigation strategy. Include direct links, refresh behavior, browser history, and back/forward navigation in the decision.
- [x] Record the selected navigation pattern and its rationale before implementation.

Selected pattern: one shared Skill Issue header followed by a quiet three-item destination row. Explore contains the existing product tabs, Project contains purpose and limitations, and Analysis contains the evidence-gated interpretation page. Hash routes keep the static GitHub Pages build directly addressable and preserve refresh, back, and forward behavior without a server fallback.

### 7. Add the shared three-destination website structure

- [x] Preserve the existing scroll-based landing page with its Build Skills and Evaluate Environments product views.
- [x] Add a dedicated project-purpose destination for the project's motivation, goals, direction, progress, and limitations.
- [x] Add a dedicated results-analysis destination that remains a polished coming-soon page until accepted campaign evidence exists.
- [x] Add compact shared navigation that makes all three destinations discoverable without overwhelming the centered Skill Issue identity.
- [x] Keep the brand, theme control, primary action, layout grid, typography, color tokens, and responsive behavior coherent across all destinations.
- [x] Give the active destination an accessible and visually explicit selected state.
- [x] Verify keyboard navigation, focus treatment, browser history, direct destination links, mobile navigation, and GitHub Pages refresh behavior.

### 8. Create the project-purpose page from the README

- [x] Use a research sub-agent to identify roughly six or seven strong GitHub README or project pages that explain their motivation clearly and naturally.
- [x] Review the returned examples directly before drafting the Skill Issue page. Study their heading hierarchy, paragraph length, sentence rhythm, punctuation, progression from problem to motivation, and balance between conviction and limitations.
- [x] Treat the root README as the content authority and adapt its meaning into concise audience-facing website copy rather than duplicating implementation-oriented repository prose.
- [x] Write the page in the project owner's natural voice. Avoid generic AI phrasing, em dashes, and dash punctuation throughout the audience-facing motivation prose.
- [x] Use the reference material to guide layout, formatting, punctuation, and rhythm without copying its wording.
- [x] Add a Motivation section explaining why skill failures are difficult to attribute across the skill description, skill body, model, and harness.
- [x] Explain the two complementary goals: evaluate whether an environment invokes skills reliably and help users create, evaluate, and refine better skills.
- [x] Explain the intended direction: users evaluate their environment, then build and refine skills with clearer evidence about the actual failure owner.
- [x] Summarize what has been completed so far at an audience-appropriate level, including the local-first CLI, governed evaluation methodology, generated-skill workflow, and reviewable result artifacts.
- [x] Add a Limitations section that clearly states the initial campaign is intentionally bounded by available time, access, and resources.
- [x] State that the project has not tested a comprehensive spread of every harness, model, configuration, operating system, or repeated trial.
- [x] State that one bounded campaign cannot establish statistical reliability, universal model behavior, permanent rankings, or guarantees about every environment.
- [x] Explain that unsupported or unrun combinations are omitted and that conclusions must remain tied to the exact tested configurations and retained evidence.
- [x] Keep limitations visible and specific without weakening the motivation or presenting future intentions as completed capabilities.

### 9. Create the results-analysis coming-soon page

- [x] Create a finished-looking placeholder that explains what the future results analysis will contain and why accepted campaign evidence is required first.
- [x] Do not invent findings, percentages, rankings, trends, or conclusions while the page is awaiting accepted results.
- [x] Explain that the completed page will interpret the charts and statistics rather than merely repeat them.
- [x] Reserve clear content areas for harness-comparison findings, model-comparison findings, quantitative differences, scenario-level observations, limitations, and links to supporting evidence.
- [x] Plan for conclusions such as whether harness choice materially changed skill calling only when the accepted data supports that statement.
- [ ] Replace the placeholder with evidence-backed analysis only after the campaign acceptance layer identifies the exact runs and configurations suitable for publication.

### 10. Validate the completed website refactor

- [x] Run `npm run typecheck` and targeted linting after TypeScript or TSX changes.
- [x] Run `npm run format:check` after CSS, Markdown, HTML, JSON, or configuration changes.
- [ ] Run `npm run validate` before presenting the website refactor as complete.
- [x] Start the local website and inspect the implementation through browser or computer control rather than relying only on source review or build output.
- [x] Click through all three destinations, both landing-page product tabs, chart interactions, themes, navigation states, direct links, browser history, and responsive layouts.
- [x] Capture screenshots of every materially changed page and the important responsive or themed states.
- [x] Compare the changed pages against the retained and newly captured reference screenshots to check visual consistency with the established Skill Issue style.
- [x] Use the image-comparison skill when a page, navigation state, spacing decision, hierarchy, or responsive layout appears inconsistent or visually uncertain.
- [x] Iterate on any confirmed visual mismatch before treating the refactor as complete.
- [x] Inspect the production build with `npm run preview` after the local interactive and visual checks pass.

Validation note: `npm run validate` reaches the repository-wide lint step and remains blocked by the pre-existing unused `id` error in the API Change Impact Mapper fixture. The Dependency Upgrade Planner fixture also emits its existing fast-refresh warning. The changed website files pass targeted linting, type checking, formatting, and the production build.

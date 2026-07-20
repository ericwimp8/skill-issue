# Website Refactor Task List

## Purpose

Track the website behavior, presentation, information architecture, content, and evaluation-data corrections needed before the public site is ready for real campaign evidence. The results experience will separate harness comparison from model comparison so each chart answers one meaningful question. The site will preserve the current landing experience and add dedicated project-purpose and results-analysis destinations. This document owns the website work queue. Evaluation campaign configuration remains owned by the campaign progress document.

## Current Review Sources

- [Skill-calling evaluation campaign matrix](../skill-calling-evaluation-campaign/evaluation-progress.md#campaign-configuration)
- [Website evaluation data adapter and illustrative cells](../../src/data/evaluationData.ts)
- [Project README](../../README.md)
- [Existing website reference findings](../../research/website-two-arm-navigation/reference-findings.md)

## Tasks

### 1. Investigate chart tooltip hover movement

- [ ] Reproduce tooltip behavior on every Recharts chart that has hover content, with particular attention to the Success Ranking chart.
- [ ] Determine why the tooltip appears to enter from the side instead of remaining spatially connected to the hovered bar or point.
- [ ] Inspect the Recharts tooltip positioning and motion configuration, including default wrapper transitions, cursor anchoring, bounds handling, offset, and any chart-specific custom tooltip behavior.
- [ ] Compare pointer movement across the Success Ranking, Cumulative Curve, and Checkpoint Profile charts to identify whether the issue is shared or chart-specific.
- [ ] Choose the smallest consistent configuration that keeps tooltips close to the active datum, avoids sideways shooting, and remains inside the chart or viewport where practical.
- [ ] Verify the corrected behavior with mouse movement in light and dark themes and at desktop and narrow widths.

### 2. Confirm the campaign harness-and-model matrix

- [x] Review the current ten-configuration campaign matrix with the project owner before changing website data.
- [ ] Confirm the exact display label and provider identifier for Codex Sol at Medium reasoning.
- [ ] Confirm the exact display label and provider identifier for Claude Fable at Medium reasoning.
- [ ] Confirm the exact Grok model or alias and its Medium or model-native reasoning setting.
- [x] Confirm Composer as a Cursor configuration; its exact model identity and reasoning behavior still need to be recorded with accepted results.
- [x] Confirm OpenCode and Kilo Code remain in this campaign.
- [ ] Resolve the older thirteen-cell research description against the current ten-configuration campaign progress matrix.

Review note: the project owner confirmed the current ten-configuration campaign matrix on 21 July 2026. The website's illustrative data still contains `claude-opus-4-8` / `Claude Opus 4.8` cells; the confirmed campaign matrix uses Fable instead and includes Cursor — Composer. The website cells must be updated from this confirmed matrix while exact effective model identifiers remain result-owned evidence.

### 3. Create the harness comparison results view

- [ ] Split the results experience into distinct Harness Comparison and Model Comparison views. Define the model-comparison chart in a separate confirmed requirement before implementing it.
- [ ] Hold the model constant at Codex with Medium reasoning for the harness comparison. Use only the confirmed harness configurations that run Codex so the chart isolates the effect of changing the harness.
- [ ] Present one harness-comparison chart only. Do not add secondary harness charts unless they answer a separately confirmed question.
- [ ] Use a horizontal 100% stacked bar chart with one bar per harness. Show successful and failed expected skill calls as complementary percentages within each bar.
- [ ] List the harness bars vertically from highest success rate at the top to lowest at the bottom so the chart immediately communicates which harnesses were more or less effective at skill calling with the same model.
- [ ] Give the chart its own scenario multi-select control, using a dropdown with one checkbox for each of the three governed scenarios.
- [ ] Select all three scenarios by default. Allow any non-empty combination: one scenario alone, any pair of scenarios, or all three scenarios together.
- [ ] Treat each selected scenario as one complete 30-turn conversation. The selected conversation context therefore represents 30 turns for one scenario, 60 turns for two scenarios, or 90 turns for all three scenarios.
- [ ] Calculate chart performance from the scored expected skill activations rather than treating every conversation turn as a call opportunity. The current governed fixtures contribute four scored activations per scenario, giving 4, 8, or 12 scored outcomes per harness when one, two, or three scenarios are selected.
- [ ] Apply the same selected scenario set to every harness so the ranked comparison always uses equivalent evidence.
- [ ] Aggregate selected scenarios by summing called and missed expected activations for each harness, then derive the displayed success and failure percentages from those totals.
- [ ] Recalculate percentages and ranking immediately when the selected scenario set changes.
- [ ] Keep the selected scenario count, represented conversation turns, and underlying called, missed, and scored-total counts available in labels or hover content so the percentages remain interpretable.
- [ ] Verify that the finished chart tells one clear story: how much the harness changes reliable skill calling when the Codex model and Medium reasoning target remain constant.

### 4. Align website evaluation data with the confirmed matrix

- [ ] Remove obsolete model labels and illustrative cells after the campaign matrix is confirmed.
- [ ] Add every confirmed harness, model, and reasoning combination needed by the results filters and charts.
- [ ] Update summary metrics so harness and comparison-cell counts match the confirmed website data.
- [ ] Keep illustrative outcomes clearly labelled as mock data until accepted campaign artifacts replace them.
- [ ] Verify harness, model, and comparison selectors expose every confirmed cell without duplicates or unsupported combinations.

### 5. Research the multi-page navigation pattern

- [ ] Review the existing website-reference findings, screenshots, and visual decisions before selecting a navigation pattern.
- [ ] If any earlier reference artifact is missing, locate it in the working tree or Git history and recover only the relevant research without reverting unrelated work.
- [ ] Revisit the established reference websites and inspect how they handle compact menus, page-to-page navigation, active destination states, shared headers, mobile navigation, and transitions between marketing and results content.
- [ ] Compare navigation approaches that preserve the current minimal Skill Issue visual language and work reliably on a static GitHub Pages deployment.
- [ ] Decide the audience-facing names for the landing, project-purpose, and results-analysis destinations after the reference review.
- [ ] Decide whether the static site should use hash-based routing or another GitHub Pages-compatible direct-navigation strategy. Include direct links, refresh behavior, browser history, and back/forward navigation in the decision.
- [ ] Record the selected navigation pattern and its rationale before implementation.

### 6. Add the shared three-destination website structure

- [ ] Preserve the existing scroll-based landing page with its Build Skills and Evaluate Environments product views.
- [ ] Add a dedicated project-purpose destination for the project's motivation, goals, direction, progress, and limitations.
- [ ] Add a dedicated results-analysis destination that remains a polished coming-soon page until accepted campaign evidence exists.
- [ ] Add compact shared navigation that makes all three destinations discoverable without overwhelming the centered Skill Issue identity.
- [ ] Keep the brand, theme control, primary action, layout grid, typography, color tokens, and responsive behavior coherent across all destinations.
- [ ] Give the active destination an accessible and visually explicit selected state.
- [ ] Verify keyboard navigation, focus treatment, browser history, direct destination links, mobile navigation, and GitHub Pages refresh behavior.

### 7. Create the project-purpose page from the README

- [ ] Treat the root README as the content authority and adapt its meaning into concise audience-facing website copy rather than duplicating implementation-oriented repository prose.
- [ ] Add a Motivation section explaining why skill failures are difficult to attribute across the skill description, skill body, model, and harness.
- [ ] Explain the two complementary goals: evaluate whether an environment invokes skills reliably and help users create, evaluate, and refine better skills.
- [ ] Explain the intended direction: users evaluate their environment, then build and refine skills with clearer evidence about the actual failure owner.
- [ ] Summarize what has been completed so far at an audience-appropriate level, including the local-first CLI, governed evaluation methodology, generated-skill workflow, and reviewable result artifacts.
- [ ] Add a Limitations section that clearly states the initial campaign is intentionally bounded by available time, access, and resources.
- [ ] State that the project has not tested a comprehensive spread of every harness, model, configuration, operating system, or repeated trial.
- [ ] State that one bounded campaign cannot establish statistical reliability, universal model behavior, permanent rankings, or guarantees about every environment.
- [ ] Explain that unsupported or unrun combinations are omitted and that conclusions must remain tied to the exact tested configurations and retained evidence.
- [ ] Keep limitations visible and specific without weakening the motivation or presenting future intentions as completed capabilities.

### 8. Create the results-analysis coming-soon page

- [ ] Create a finished-looking placeholder that explains what the future results analysis will contain and why accepted campaign evidence is required first.
- [ ] Do not invent findings, percentages, rankings, trends, or conclusions while the page is awaiting accepted results.
- [ ] Explain that the completed page will interpret the charts and statistics rather than merely repeat them.
- [ ] Reserve clear content areas for harness-comparison findings, model-comparison findings, quantitative differences, scenario-level observations, limitations, and links to supporting evidence.
- [ ] Plan for conclusions such as whether harness choice materially changed skill calling only when the accepted data supports that statement.
- [ ] Replace the placeholder with evidence-backed analysis only after the campaign acceptance layer identifies the exact runs and configurations suitable for publication.

### 9. Validate the completed website refactor

- [ ] Run `npm run typecheck` and `npm run lint` after TypeScript or TSX changes.
- [ ] Run `npm run format:check` after CSS, Markdown, HTML, JSON, or configuration changes.
- [ ] Run `npm run validate` before presenting the website refactor as complete.
- [ ] Inspect the production build with `npm run preview` and verify all three destinations, both landing-page product tabs, chart interactions, themes, navigation states, direct links, browser history, and responsive layouts.

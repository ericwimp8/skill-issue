# Skill Issue Website A-to-B Task List

## Purpose

Produce a complete local mock-up of a minimal, one-page Skill Issue website that can be built as static files for GitHub Pages, directs CLI downloads to GitHub Releases, presents locally maintained skill-evaluation graphs, supports light and dark themes, and remains easy to extend without a database.

## Completion Definition

The work is complete when every task below is checked, every task-level completion requirement is satisfied, the website passes its documented deterministic checks, and the final audit finds no unresolved requirement from the project brief.

## Ordered Tasks

### 1. Confirm the product and delivery constraints

- [x] Read the project overview, product-shape notes, existing hosting research, and repository instructions.
- [x] Record the website's audience, essential content, static-hosting boundary, local-data boundary, and CLI-download boundary.

**Completion requirement:** The implementation decisions are grounded in the repository's current product meaning, GitHub Pages is the static host target, GitHub Releases is the binary source, and no hosted database or deployment work is introduced.

### 2. Audit this task list against the brief

- [x] Check that every requested deliverable, workflow, validation step, non-goal, and final recommendation is represented.
- [x] Reorder any task whose inputs are produced by a later task.

**Completion requirement:** Every later task depends only on decisions or artifacts produced by earlier tasks, and the audit records no missing brief requirement.

### 3. Research visual references in Chrome

- [x] Identify at least ten current minimal product, developer-tool, analytics, or benchmark websites suitable for a light/dark one-page experience.
- [x] Open each reference through approved browser automation and capture a screenshot.
- [x] Record relevant layout, typography, surfaces, navigation, color, and chart-presentation observations without copying brand assets or text.

**Completion requirement:** At least ten named references have locally stored screenshots and concise, source-linked observations based on direct Chrome inspection.

### 4. Select and document the visual direction

- [x] Choose the strongest reference for Skill Issue's content and graph-heavy page.
- [x] Define which visible characteristics are binding inspiration and which reference-specific details remain incidental.
- [x] Define the one-page information architecture and responsive section order.

**Completion requirement:** One reference is selected with an explicit rationale, a bounded inspiration contract, and a complete page-section outline.

### 5. Select the static stack and chart library

- [x] Compare suitable static GitHub Pages stacks and choose the smallest idiomatic option that supports reusable components and typed local data.
- [x] Research a popular, free, maintained graphing library from primary sources and confirm licensing, required chart types, responsiveness, theming, and accessibility considerations.
- [x] Decide how GitHub Pages base paths and GitHub Release download URLs will be configured.

**Completion requirement:** The stack, chart library, licensing basis, static build approach, base-path handling, and configurable download-link strategy are documented before scaffolding.

### 6. Define the website data and design contracts

- [x] Define one local typed data source for site copy, release metadata, summary metrics, graph definitions, descriptions, labels, and series values.
- [x] Define a reusable graph-card contract so another harness graph can be added through data rather than page markup.
- [x] Define top-level light/dark design tokens for color, typography, spacing, radii, borders, and elevation.

**Completion requirement:** Content, graph data, release configuration, and design primitives each have one clear semantic owner, with no database or repeated page-level constants required.

### 7. Scaffold the deterministic development toolchain

- [x] Create the static application with TypeScript and the selected component framework.
- [x] Configure package scripts for development, formatting checks, linting, type checking, building, and a single full validation command.
- [x] Configure GitHub Pages-compatible static output and repository base paths.
- [x] Add scoped `AGENTS.md` instructions that require the fastest deterministic checks after edits and the full validation command before completion.

**Completion requirement:** Dependencies install successfully, the development server starts, and formatting, linting, type checking, and production building are available through documented scripts.

### 8. Implement the reusable visual foundation

- [x] Implement global tokens and shared layout, typography, button, badge, surface, and icon-button patterns.
- [x] Implement an accessible theme control with system preference, explicit light/dark selection, persistence, and correct initial rendering.
- [x] Implement responsive page width, spacing, and focus-visible behavior.

**Completion requirement:** Shared appearance is controlled through top-level tokens and reusable components; both themes remain readable; keyboard focus is visible; and repeated bespoke styling is absent.

### 9. Implement the one-page product experience

- [x] Build minimal header navigation, hero copy, status treatment, and a configurable GitHub Release download action.
- [x] Build the benchmark overview and data-driven evaluation graph grid.
- [x] Show mocked Codex and Claude Code evaluations with skill calls and skill misses plotted against consumed context.
- [x] Add concise methodology/context copy and a restrained footer without speculative pages or excessive external links.

**Completion requirement:** The page communicates what Skill Issue is, provides the download path, displays the required graph meanings, renders all graph cards from local data, and keeps the experience minimal.

### 10. Verify functionality and responsive behavior in Chrome

- [x] Run the local application and inspect it through approved browser automation during implementation.
- [x] Verify navigation, theme switching, download link behavior, chart rendering, accessible chart summaries, and keyboard focus.
- [x] Capture full-page and section screenshots for desktop light, desktop dark, and a representative narrow viewport.

**Completion requirement:** Every required interaction and state works in the rendered application, and stable local screenshots exist for all required comparison targets.

### 11. Run the first pairwise image-comparison audit

- [x] Create the durable image-comparison report before opening comparison pairs.
- [x] Compare one website target screenshot against one selected-reference screenshot at a time.
- [x] Record each pair's actionable binding differences before opening the next comparison image.

**Completion requirement:** The audit follows the image-comparison workflow and identifies concrete geometry, hierarchy, surface, spacing, typography, state, or content-relationship fixes without treating incidental source details as requirements.

### 12. Refine and repeat visual comparisons

- [x] Fix the first audit's actionable differences at their owning tokens or components.
- [x] Recapture every affected target state.
- [x] Run a second pairwise comparison and resolve any remaining actionable binding differences.

**Completion requirement:** A later audit records no unresolved actionable visual differences for the selected inspiration contract across desktop light, desktop dark, and narrow layouts.

### 13. Validate code and production output

- [x] Run formatting, linting, type checking, and the production build through the deterministic validation command.
- [x] Inspect the built static output through a local server using the configured GitHub Pages base path.
- [x] Confirm the website work contains no accidental generated caches or changes outside the intended website artifact locations.

**Completion requirement:** The full validation command passes, the production output renders correctly at its Pages-style path, and the change set is scoped and reviewable.

### 14. Audit completion against every requirement

- [x] Re-read the original brief and this checklist.
- [x] Check every task-level completion requirement against concrete files, commands, screenshots, and audit evidence.
- [x] Mark tasks complete only when their evidence exists and record any genuine limitation explicitly.

**Completion requirement:** Every satisfied checkbox has supporting evidence, no requested capability is silently omitted, and any remaining limitation is visible rather than inferred away.

### 15. Prepare the handoff and next-step recommendations

- [x] Document local setup, development, validation, build, graph-data updates, adding another graph, and configuring the eventual release URL.
- [x] Read the authoritative competition rules before recommending whether the site should link to the competition.
- [x] Recommend the smallest professional set of future links, whether an About page is warranted, and the next highest-value product step.

**Completion requirement:** A new contributor can run and update the mock-up from repository documentation, and the final recommendations are specific, restrained, and grounded in the project and competition context.

## Initial Brief Audit

- The plan includes at least ten Chrome-inspected references and local screenshots.
- The plan requires a selected inspiration, implementation, repeated screenshot comparison, and repair loops.
- The plan requires a popular free chart library chosen through current primary-source research.
- The plan centralizes mock graph content and makes additional harness graphs data-driven.
- The plan includes skill calls and skill misses over consumed context for Codex and Claude Code.
- The plan includes a minimal one-page layout, GitHub Release download action, light/dark themes, responsive behavior, and reusable top-level styling.
- The plan includes GitHub Pages-compatible static output and excludes deployment and hosted databases.
- The plan includes idiomatic formatting, linting, type checking, compilation/building, and deterministic `AGENTS.md` guidance.
- The plan includes regular rendered inspection through Chrome, section/state screenshots, pairwise image-comparison reports, fixes, and repeat comparisons.
- The plan includes a final requirement audit, contributor handoff, professional-link recommendations, About-page advice, competition-link advice, and next-step guidance.
- No task depends on a decision or artifact produced by a later task.

**Audit result:** The task list covers the full brief and is ordered by dependency. Implementation may begin with Task 3 after Tasks 1 and 2 are marked complete from the repository evidence and this audit.

## Final Completion Audit

- Product and hosting boundary: `reference-and-architecture.md` records the static GitHub Pages and GitHub Releases split; the implementation contains no database or runtime API.
- Visual research: `references/` contains ten directly captured official-site screenshots plus the selected narrow reference.
- Visual direction: shadcn/ui is selected with binding and non-binding characteristics documented before implementation.
- Toolchain: React, TypeScript, Vite, Recharts, ESLint, and Prettier are installed with deterministic scripts and repository instructions.
- Data ownership: `src/data/siteData.ts` owns all editable copy, release metadata, metrics, methodology, and graph definitions.
- Required experience: the one-page site includes light/dark themes, responsive layout, download action, benchmark metrics, Codex and Claude Code graphs, method explanation, and minimal footer.
- Browser evidence: `screenshots/` contains desktop light, desktop dark, narrow, full-page, and compiled Pages-build captures; browser checks verified two charts, anchor navigation, persisted theme selection after reload, no horizontal overflow, and no console errors.
- Visual audit: `image-comparison-audit.md` records direct pairwise inspection and closes with no actionable differences after two repair loops.
- Static validation: `npm run validate` passes formatting, linting, type checking, and the production build; the `/skill-issue/` preview mounts correctly with base-prefixed assets.
- Contributor handoff: `README.md` documents setup, validation, production preview, data updates, additional graphs, and release URL configuration.
- Known limitation: no CLI release asset exists yet, so the configured action targets the latest GitHub Releases page and explicitly says the CLI is coming soon. Replace it with a stable latest-asset URL after the first packaged release.
- Recommended next links: repository, installation documentation, methodology/evaluation protocol, benchmark data source, and competition entry while Build Week remains relevant.
- About-page recommendation: keep the project story as a concise one-page section until the team, governance, or methodology content is substantial enough to justify a separate page.
- Competition-link recommendation: add a small “Built for OpenAI Build Week” footer link to the official challenge page for the submission period, then reassess its prominence after judging.
- Highest-value next step: package and publish one signed CLI prerelease with stable cross-platform asset names so the download action can become a real direct download and the installation path can be tested end to end.

**Final audit result:** Every task-level completion requirement has concrete evidence. The only open product dependency is the future CLI release artifact, which is represented honestly rather than mocked as an available download.

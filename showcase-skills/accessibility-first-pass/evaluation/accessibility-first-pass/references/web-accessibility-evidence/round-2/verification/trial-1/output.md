# Accessibility First-Pass Report

## Review Scope

- **Target:** The supplied static service dashboard in `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dashboard/index.html`.
- **Journeys and states:** Initial service-overview state showing Search and Payments status indicators and two detail buttons.
- **Source areas:** Supplied HTML and embedded CSS, plus `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dashboard/scan-results.json`.
- **Rendered environments:** None. No browser, accessibility tree, keyboard session, device, or assistive technology was exercised.
- **Standards or guidance considered:** WCAG 2.2, especially SC 1.4.1 Use of Color and SC 1.3.1 Info and Relationships; W3C heading guidance.
- **Out of scope:** Other routes and states, detail views, dynamic updates, production data, external styles or scripts, responsive behavior, browser interoperability, and overall conformance.

## Methods and Evidence

| Method | Target and state | Result or evidence path | Evidence class | Limitations |
| --- | --- | --- | --- | --- |
| Complete source inspection | Supplied initial dashboard HTML and CSS | Two service states are represented by same-shaped empty `span` elements whose `good` and `bad` classes differ by background color; the visible page title is a `div`; the document declares `lang="en"`; two native buttons have distinct labels | Observed | Source does not establish computed styles, accessibility-tree output, perception, interaction behavior, or announcements |
| Supplied scanner-result inspection | Supplied JSON result set | `synthetic-checker` reported `color-only-status` as `minor` and `heading-structure` as `serious` | Observed | This proves only what the synthetic tool reported; its severity labels do not establish actual user impact or report priority |
| Authoritative-guidance comparison | Source evidence against WCAG 2.2 guidance | Finding 1 maps directly to SC 1.4.1; Finding 2 is consistent with SC 1.3.1 heading-structure guidance | Inferred | No conformance test, rendered inspection, or assistive-technology verification was performed |

## Material Limitations

- The fixtures contain no runnable detail-view behavior, production context, or evidence that a service status changes dynamically.
- User impact is reasoned from the affected dashboard task and source implementation, not directly observed with disabled users or assistive technology.
- Scanner severities are investigation clues only. Report priorities below use task importance, reach, workaround burden, and evidence confidence.
- No clean automated result, passed browser check, or overall accessibility conclusion is available.

## Prioritized Findings

### High: Service health is communicated only by color

- **Evidence class:** **Observed:** source implementation and scanner report. **Inferred:** resulting user impact because no rendered or user session was exercised.
- **Affected users and task impact:** People who cannot distinguish the red and green states, including some people with color-vision differences or low vision and users of limited-color displays, may be unable to determine whether Search or Payments is healthy. This blocks the dashboard's central information task rather than merely making it less polished. Users relying on screen readers or braille may also receive no status information because the status spans are empty and unnamed; that behavior remains unverified in an accessibility tree.
- **Evidence:** Each service row contains its service name and an empty `<span class="status good">` or `<span class="status bad">`. Both indicators share the same circular shape; only `.good` and `.bad` background colors differ. No visible status word, icon-shape difference, text alternative, or programmatic association conveys the state. The scanner independently reported `color-only-status` with severity `minor`.
- **Reproduction or inspection steps:** Open `index.html`; inspect the Search and Payments rows and the `.status`, `.good`, and `.bad` CSS rules; disable color or compare the DOM text available in each row. Only the service names remain as text, with no health value.
- **Authoritative guidance:** [WCAG 2.2 SC 1.4.1](https://www.w3.org/TR/WCAG22/#use-of-color) requires that color not be the only visual means of conveying information. W3C's [Use of Color understanding guidance](https://www.w3.org/WAI/WCAG22/Understanding/use-of-color.html) recommends an additional cue such as text or shape and explains that a visible alternative is still required even if assistive technology receives the information.
- **Priority rationale:** **High**, overriding the scanner's `minor` label. The missing distinction affects every listed service and the page's primary purpose; the supplied state provides no visible non-color workaround. Confidence in the implementation evidence is high, while the exact severity for real users depends on production context and status criticality.
- **Remediation direction:** At the service-status data and row-rendering owner, expose each state as concise visible text such as “Operational” or “Disrupted,” programmatically associate it with the service, and retain color only as a redundant cue. If an icon is also used, give states distinguishable shapes and ensure the accessible output does not duplicate the same announcement.
- **Human or assistive-technology follow-up:** In the rendered dashboard, verify the two states remain distinguishable in grayscale, common color-vision simulations, Windows forced-colors mode, and browser high-contrast settings. With current NVDA + Firefox or Chrome and VoiceOver + Safari, read each service row and confirm the service name and current status are announced together once. If states update, trigger a change and verify an appropriate, non-duplicative announcement without moving focus.
- **Confidence and limitations:** High confidence that the supplied source uses color as its only visible status distinction. Screen-reader output, forced-color rendering, update behavior, and the real-world consequence of an outage were not tested.

### Medium: The visible page title has no heading semantics

- **Evidence class:** **Observed:** the title uses a `div` and the scanner reported a heading issue. **Inferred:** navigation and orientation impact because the accessibility tree and assistive technology were not exercised.
- **Affected users and task impact:** Screen-reader users who navigate by headings may not find a heading for the page, making orientation and rapid navigation less efficient. In this short, single-section fixture the content remains available by linear reading and the document title identifies the page, so the impact is different from and lower than the blocked status-information task.
- **Evidence:** `Service overview` appears in `<div class="page-title">`; the body contains no `h1`–`h6` element. The scanner reported `heading-structure` with severity `serious`.
- **Reproduction or inspection steps:** Inspect the body markup and search for heading elements or `role="heading"`; none are present. Then inspect the rendered accessibility tree and invoke heading navigation in follow-up testing.
- **Authoritative guidance:** [WCAG 2.2 SC 1.3.1](https://www.w3.org/TR/WCAG22/#info-and-relationships) requires visually conveyed structure and relationships to be programmatically determinable or available in text. W3C's [headings tutorial](https://www.w3.org/WAI/tutorials/page-structure/headings/) explains that headings communicate page organization and support in-page navigation.
- **Priority rationale:** **Medium**, replacing the scanner's `serious` severity with task-based priority. The issue can degrade orientation and heading navigation across the page, but this fixture is short, its reading order is simple, and the browser document title supplies some orientation. Confidence in the markup observation is high; practical impact remains unverified.
- **Remediation direction:** At the page-shell or page-title owner, represent the main visible title with the native heading level that accurately expresses its place in the document outline, ordinarily the page's primary `h1` here. Preserve visual styling through CSS rather than using a non-semantic container.
- **Human or assistive-technology follow-up:** In current NVDA + Firefox or Chrome and VoiceOver + Safari, inspect the heading list and use heading navigation from the page start. Confirm “Service overview” is exposed once at the intended level and that focus or reading order remains logical.
- **Confidence and limitations:** High confidence that no heading semantics exist in the supplied markup. Accessibility-tree exposure, screen-reader navigation behavior, and any production page-shell headings were not available.

## Passed Checks Within Tested Scope

| Check | Exact state and method | Evidence | Limitations |
| --- | --- | --- | --- |
| Document language declaration present | Static source inspection of the supplied initial page | Root element is `<html lang="en">` | Language accuracy for all production content and runtime accessibility-tree exposure were not tested |
| Page has a non-empty document title | Static source inspection | `<title>Service overview</title>` | Uniqueness across routes and usefulness in the production navigation context are unknown |
| Detail controls use native button elements with distinct source labels | Static source inspection | Two `button` elements have distinct `aria-label` values matching their visible text | Keyboard behavior, focus visibility, activation, accessible-name computation, and resulting detail behavior were not tested |

## Follow-Up Checks and Unknowns

| Check | Why unresolved | Required environment or method | Expected behavior |
| --- | --- | --- | --- |
| Status name and relationship | Source contains empty status spans; computed accessibility output was unavailable | Current NVDA + Firefox or Chrome and VoiceOver + Safari; inspect accessibility tree and read each row | Each service name and current status are conveyed together once and remain understandable out of visual context |
| Dynamic status changes | Fixtures do not establish whether updates occur | Runnable dashboard with a controlled state change in each supported browser/AT combination | Important updates are announced at an appropriate time without focus theft or repeated speech |
| Heading navigation and page orientation | Source establishes missing native heading markup, not experienced navigation behavior | Current NVDA and VoiceOver; heading list and next-heading commands from page start | The page title is exposed once at the correct level and supports efficient navigation |
| Keyboard operation and visible focus | No interactive session was run | Keyboard-only pass in every supported browser through both buttons and any resulting detail state | Logical focus order, visible focus, operable controls, no trap, and appropriate focus after activation |
| Zoom, reflow, text spacing, and constrained viewport | Embedded source is too small to prove rendered behavior | Browser checks at 200% and 400% zoom, 320 CSS px width, and WCAG text-spacing overrides | Service/status relationships and controls remain visible, readable, and operable without loss or obscuring |
| Forced colors and non-color perception | CSS background-color rendering was not exercised | Windows forced-colors/high-contrast mode plus grayscale and representative color-vision simulations | Status remains visibly distinguishable without depending on the red/green backgrounds |
| Voice-input operability | No speech-input session was run | Current voice-control software in a supported browser | Speaking each visible button label selects the intended control |
| Representative usability | Source and scanner output cannot establish usability across user strategies | Task-based sessions with disabled participants using the production dashboard | Users can identify every service state and open details efficiently and without ambiguity |

## Prioritized Next Actions

1. Add visible, programmatically associated status text for every service, keeping color only as a redundant cue.
2. Replace the page-title container with the appropriate native heading and verify the resulting document outline.
3. Run targeted screen-reader and keyboard checks for status reading, heading navigation, focus visibility, and detail activation.
4. Verify status perception in forced-colors, grayscale, color-vision simulations, zoom, reflow, and text-spacing conditions.
5. If statuses update dynamically, test announcements; then conduct representative usability testing on the production critical journey.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

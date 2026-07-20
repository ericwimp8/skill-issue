# Accessibility First-Pass Report

## Review Scope

- **Target:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html`
- **Journeys and states:** Source-level review of the account-request journey: identifying the email field, understanding its required state, encountering the visible error text, and activating “Send request.” Only the authored default state was available.
- **Source areas:** Complete HTML, inline CSS, form markup, and inline submission script in the single fixture.
- **Rendered environments:** None. No page was served or opened in a browser; no keyboard, pointer, accessibility-tree, screen-reader, voice-input, zoom, or forced-colors session was run.
- **Standards or guidance considered:** WCAG 2.2 criterion relationships were considered using W3C Understanding guidance. Mappings below are bounded to the inspected evidence and are not conformance determinations.
- **Out of scope:** Server behavior and responses; success, loading, and failure states after submission; other routes; browser/device interoperability; computed styles; native validation presentation; and assistive-technology output.

## Methods and Evidence

| Method | Target and state | Result or evidence path | Evidence class | Limitations |
| --- | --- | --- | --- | --- |
| Complete source inspection | Default authored account-request page | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:1` | Observed | Establishes authored elements, attributes, styles, and script only. |
| Form semantics inspection | Email input, error text, and submission control | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:16` | Observed | Runtime accessible names, roles, relationships, focus behavior, and announcements were not captured. |
| Inline submission-path inspection | `onclick` activation and `form.submit()` call | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:19` | Observed | The path exists in source; browser behavior and server outcome were not exercised. |
| Automated accessibility checks | Not run | Not run | Unverified | No rendered target or project-native accessibility checker was exercised. |
| Rendered manual and assistive-technology checks | Not run | Not run | Unverified | Requires browsers, keyboard/pointer interaction, accessibility inspection, and assistive technologies. |

## Material Limitations

- This is a source-only first pass. It cannot establish computed contrast, focus order or visibility, keyboard completion, accessible names exposed by a browser, native validation presentation, error announcements, zoom/reflow behavior, or assistive-technology interoperability.
- Only the initial authored state exists in the fixture. Submission results, server validation, conditional errors, recovery, and success confirmation are unavailable.
- No automated checker result is available. Even a later clean automated scan would require human judgment and targeted interaction testing.

## Prioritized Findings

### High: Submission is implemented as a pointer-click `div`

- **Evidence class:** Inferred. The source observation is definitive: “Send request” is a `div` with an `onclick` handler, with no native button element or authored keyboard handler (`index.html:19`). Whether another keyboard submission path works, and the actual focus/activation behavior, were not rendered.
- **Affected users and task impact:** People using keyboard, switch, speech-input, or other keyboard-emulating input may be unable to discover or activate the visible submission control. Screen-reader users may encounter text without an actionable button role. This can seriously impede completion of the account-request journey.
- **Evidence:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:19` contains `<div class="submit" onclick="document.querySelector('form').submit()">Send request</div>`.
- **Reproduction or inspection steps:** Open the fixture source; inspect line 19; confirm the visible control is a `div`, activation is attached only to `onclick`, and no native button or keyboard handler is authored. Rendered keyboard behavior remains a follow-up.
- **Authoritative guidance:** W3C’s [Understanding SC 2.1.1 Keyboard](https://www.w3.org/WAI/WCAG22/Understanding/keyboard.html) says pointer actions need a keyboard equivalent and recommends native HTML form controls. [Understanding SC 4.1.2 Name, Role, Value](https://www.w3.org/WAI/WCAG22/Understanding/name-role-value.html) is also potentially relevant to the control’s programmatic role. Applicability must be confirmed against rendered behavior and the complete journey.
- **Priority rationale:** Submitting the form is the journey’s decisive action. The source shows a broad, repeated interaction pattern with a substantial risk of blocking non-pointer users; confidence in the authored-risk evidence is high, while runtime behavior remains unverified.
- **Remediation direction:** At the form interaction owner, use a native submit button and the form’s normal submit flow. Preserve platform keyboard activation, focus semantics, and validation rather than recreating them on a generic element.
- **Human or assistive-technology follow-up:** In current Chrome and Firefox, complete the journey using Tab, Shift+Tab, Enter, and Space without a pointer; record focus order, visible focus, and whether the visible control submits. In Safari with VoiceOver and Chrome with NVDA, confirm the control is announced as a button with the name “Send request” and can be activated.
- **Confidence and limitations:** High confidence that the authored visible control is a click-handled `div`; no claim is made about runtime focusability, alternate implicit submission, or standards conformance.

### High: The required email field has placeholder text but no persistent associated label

- **Evidence class:** Inferred. Source inspection observes one input with `placeholder="Email address"` and `required`, and no `label`, `aria-label`, or `aria-labelledby` in the complete document (`index.html:17`). The browser-exposed accessible name was not captured.
- **Affected users and task impact:** Screen-reader and refreshable-braille users may receive an unclear or inconsistent name; speech-input users may be unable to target the field by its visible wording; people with cognitive, language, attention, or memory disabilities may lose the only instruction after entering text. Identifying the required account email can become confusing or error-prone.
- **Evidence:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:17` contains only the input’s `id`, `name`, placeholder, and required attribute; the complete form at lines 16–20 contains no associated label.
- **Reproduction or inspection steps:** Inspect lines 16–20; confirm that “Email address” appears only as placeholder content and that no persistent label or authored accessible-name attribute is present. Capture the actual accessible name in a browser before making a runtime claim.
- **Authoritative guidance:** W3C’s [Understanding SC 3.3.2 Labels or Instructions](https://www.w3.org/WAI/WCAG22/Understanding/labels-or-instructions.html) calls for labels or instructions for user input and distinguishes visible labels from programmatic names. SC 1.3.1 and SC 4.1.2 may also be relevant, subject to rendered accessibility-tree confirmation.
- **Priority rationale:** The field is required and is the journey’s only data-entry control. The issue affects orientation, input, and recovery; the workaround burden can be substantial for assistive-technology users. Source confidence is high, while the exact runtime name remains unverified.
- **Remediation direction:** Give the email input a persistent, descriptive visible label associated through native HTML. Keep placeholder text only as optional supplemental guidance, and communicate the required state in a form users can perceive and programmatically determine.
- **Human or assistive-technology follow-up:** In Safari/VoiceOver and Chrome/NVDA, inspect and announce the field’s name, role, required state, and description before and after entering text. With voice control, attempt to focus it using the visible phrase “Email address.”
- **Confidence and limitations:** High confidence in the placeholder-only authored pattern; accessible-name computation and assistive-technology speech/braille output remain unverified.

### High: The direct submission path risks bypassing required-field validation

- **Evidence class:** Inferred. The source definitively establishes `required` on the input (`index.html:17`) and direct invocation of `form.submit()` (`index.html:19`). Whether the tested browser presents native validation, whether submission succeeds, and what the server does were not exercised.
- **Affected users and task impact:** People who rely on clear error prevention and recovery—including screen-reader users and people with cognitive, learning, language, attention, or memory disabilities—may submit incomplete data without timely, perceivable guidance or may encounter a later unexplained failure.
- **Evidence:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:17` declares the field required; `index.html:19` invokes `document.querySelector('form').submit()` directly.
- **Reproduction or inspection steps:** Inspect lines 17 and 19 to establish the authored validation constraint and submission path. Then render the page, leave the field empty, activate the visible control, and observe whether the browser blocks submission, identifies the field, moves focus, and presents an error.
- **Authoritative guidance:** W3C’s [Understanding SC 3.3.1 Error Identification](https://www.w3.org/WAI/WCAG22/Understanding/error-identification.html) is potentially relevant if a detected input error is not identified and described. A criterion determination requires the actual validation and submission result.
- **Priority rationale:** Validation sits on the only required input and directly affects journey completion. The source establishes a credible high-impact risk, but the report deliberately withholds a runtime failure or conformance claim until browser and server behavior are observed.
- **Remediation direction:** Route activation through the form’s native submit behavior or equivalent constraint-validation-aware handling. At the validation owner, prevent invalid submission, identify the affected field, provide a specific correction message, and manage focus/announcement appropriately.
- **Human or assistive-technology follow-up:** In current Chrome, Firefox, and Safari, submit the empty form using pointer and keyboard paths. Record whether native validation runs, the exact visible message, focus placement, accessibility-tree state, and VoiceOver/NVDA announcement; repeat with server-side rejection if available.
- **Confidence and limitations:** High confidence in the authored `required` and `form.submit()` path; native validation presentation and end-to-end outcome are unverified.

### Medium: Error text is always present and has no authored relationship to the field

- **Evidence class:** Inferred. Source inspection observes a static `<span class="error">Invalid value</span>` immediately after the input, with no `id`, `aria-describedby`, live-region attribute, conditional state, or other authored association (`index.html:18`). The rendered accessibility relationship and announcement were not captured.
- **Affected users and task impact:** Screen-reader and braille users may not discover which field the message describes or hear it when an error occurs. Sighted users, including people with cognitive or learning disabilities, may see “Invalid value” before interacting and be unsure whether an error already exists or how to fix it. Recovery may be confusing and error-prone.
- **Evidence:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:18` contains the always-authored generic error text; lines 16–20 contain no programmatic relationship or state change.
- **Reproduction or inspection steps:** Inspect lines 16–20; confirm the error text is unconditional in the source and lacks an authored identifier/relationship. Render the initial, invalid, and corrected states to verify visibility and announcements.
- **Authoritative guidance:** W3C’s [Understanding SC 3.3.1 Error Identification](https://www.w3.org/WAI/WCAG22/Understanding/error-identification.html) is potentially relevant when detected errors need text identification and description. SC 1.3.1 may also be relevant to the field-message relationship; both require runtime/state confirmation here.
- **Priority rationale:** The message affects the only field and can impair error recovery, but the degree of blockage depends on unavailable runtime and server behavior. Source confidence is high; user impact is medium pending confirmation.
- **Remediation direction:** At the validation-message owner, show a specific error only when the corresponding invalid state exists, associate it programmatically with the email field, expose invalid state, and ensure newly presented feedback is perceivable without disorienting focus changes.
- **Human or assistive-technology follow-up:** With VoiceOver/Safari and NVDA/Chrome, compare the initial, empty-submit, invalid-response, and corrected states. Confirm the message is announced once at the right time, includes corrective guidance, is associated with the email field, and clears when resolved.
- **Confidence and limitations:** High confidence in the unconditional, unassociated source text; actual visibility, announcement timing, and validation-state lifecycle are unverified.

## Passed Checks Within Tested Scope

| Check | Exact state and method | Evidence | Limitations |
| --- | --- | --- | --- |
| Document language is authored | Complete source inspection found `<html lang="en">` | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:2` | Establishes the authored value only; language accuracy for all future content was not evaluated. |
| Page title is authored | Complete source inspection found `<title>Account request</title>` | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:6` | Browser presentation and multi-page title uniqueness were not tested. |
| Main landmark and one top-level heading are authored | Complete source inspection found `<main>` and `<h1>Request an account</h1>` | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:14` | Reading order and accessibility-tree exposure were not rendered. |
| Viewport metadata permits initial responsive layout | Complete source inspection found width/device-width and initial-scale metadata | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html:5` | Zoom, reflow, orientation, and obscured-content behavior were not tested. |

No interaction, validation, visual, or assistive-technology check is reported as passed.

## Follow-Up Checks and Unknowns

| Check | Why unresolved | Required environment or method | Expected behavior |
| --- | --- | --- | --- |
| Complete keyboard journey and visible focus | Source cannot establish focus order, focus indicator, implicit Enter submission, or final operability | Current Chrome and Firefox; keyboard-only Tab/Shift+Tab/Enter/Space from page load through submission | Every operable element is reachable in a logical order, focus remains visible, and the journey completes without pointer input. |
| Accessible names, roles, states, and relationships | Browser accessibility-tree computation was not captured | Chrome accessibility inspector plus NVDA; Safari accessibility inspector plus VoiceOver | Email has a concise persistent name and required state; submit is exposed as a button; error is associated and announced when relevant. |
| Native and server validation presentation | `required` and `form.submit()` are observed, but runtime validation and response behavior are unavailable | Chrome, Firefox, and Safari; empty submission by pointer, Enter, and button activation; test server rejection if available | Invalid submission is prevented or clearly handled; the field and correction are identified; focus/announcement support recovery. |
| Computed color contrast and non-color cues | Inline declarations do not establish final computed colors, font sizes, states, or perception | Browser computed-style inspection and contrast measurement for normal, focus, hover, invalid, and disabled states | Text and controls meet the project’s WCAG target; errors and focus are identifiable without color alone. |
| Zoom, reflow, text spacing, and orientation | No rendered viewport was available | Browser at 200% and 400% zoom, 320 CSS-pixel width, WCAG text-spacing overrides, portrait and landscape | Content remains readable and operable without loss, overlap, clipping, or two-dimensional scrolling where not essential. |
| Target size and pointer usability | Rendered dimensions and adjacent targets are unknown | Browser measurement and touch-device testing | The submission target is large enough and reliably operable for people with limited dexterity or tremor. |
| High contrast and forced colors | Computed presentation and system overrides were not exercised | Windows forced-colors mode and macOS increased-contrast settings | Text, field boundaries, error state, and focus remain perceivable. |
| Voice-input operation | Visible-label matching and semantic activation were not exercised | Voice Control or Dragon; target “Email address” and “Send request” by spoken label | Visible control wording can be used to focus and activate the corresponding controls. |
| Disabled-user journey usability | A source review cannot assess clarity, efficiency, or user strategy | Moderated testing with representative disabled users after remediation | Users can understand, complete, correct, and confirm an account request without avoidable confusion or blockage. |

## Prioritized Next Actions

1. Replace the click-handled `div` with a native submit button and use a constraint-validation-aware form submission path.
2. Add a persistent, natively associated email label and perceivable required-field instruction.
3. Make validation state-specific: show a corrective message when needed, associate it with the field, expose invalid state, and support appropriate announcement/focus behavior.
4. Run the exact keyboard, browser-validation, accessibility-tree, VoiceOver, NVDA, and voice-input checks listed above across initial, invalid, corrected, and submitted states.
5. Measure computed contrast and verify zoom, reflow, text spacing, target size, forced-colors, and orientation behavior in the rendered page.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

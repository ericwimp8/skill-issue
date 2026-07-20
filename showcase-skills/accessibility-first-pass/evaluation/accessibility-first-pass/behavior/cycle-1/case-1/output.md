# Accessibility First Pass Report

## Review Scope

- **Target:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html`
- **Included journeys, states, routes, and viewports:** Source representation of the initial “Create account” page and the only declared activation path, which changes the URL fragment to `#submitted`. No rendered state, route, or viewport was exercised.
- **Available source and rendered access:** Complete access to the single HTML fixture, including its inline CSS and inline click handler. Rendered access was deliberately unavailable.
- **Project tooling and commands used:** Read-only, line-numbered source inspection with shell text tools. No browser, server, accessibility tree, scanner, or other accessibility automation was used.
- **Guidance baseline:** The request establishes no WCAG version or conformance level. WCAG criterion references below are narrow, provisional guidance mappings for remediation and verification, not a conformance baseline. The skill’s supplied W3C evidence reference was used; current criterion text and status could not be checked online within this case boundary.
- **Exclusions:** Every other repository file; rendered and computed behavior; keyboard, pointer, touch, screen-reader, speech-input, switch-control, magnification, zoom, reflow, responsive, forced-colors, and disabled-user testing; current web guidance lookup; and standards-conformance evaluation.

## Evidence Summary

- **Observed:** None. “Observed” is reserved for behavior directly reproduced in a rendered surface, and no rendered surface was available.
- **Tool:** None. No automated accessibility check was run.
- **Source-backed:** The document declares English, a page title, and an `h1` (`index.html`, lines 2, 5, and 14). The email field has a visually adjacent but unbound `label` (lines 15–16). The error paragraph is always present in the initial source and has no programmatic association with the input (lines 16–17). The Continue action is a `div` with only an inline `onclick` handler (line 18). The heading and label inherit `#aaa` text on `#fff` from `body` (line 7), approximately 2.32:1 contrast.
- **Inference:** The source defects are likely to block or materially impede keyboard-only, screen-reader, speech-input, switch-control, and low-vision users. These impacts require rendered and assistive-technology verification; they are not reported as observed user behavior.
- **Unverified:** Computed accessible names, roles, states, descriptions, focus order and visibility, actual keyboard operation, validation timing, error announcement, mobile viewport behavior, zoom and reflow, platform color rendering, and end-to-end account creation.
- **Evidence artifact:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html`, inspected 2026-07-20.

## Prioritized Findings

### [P1] Continue Is Implemented as a Pointer-Only Generic Element

- **Affected users and tasks:** Keyboard-only and switch-control users may be unable to activate Continue; screen-reader and speech-input users may not have an identifiable button target. The affected task is advancing from the account-creation page.
- **Evidence level:** Source-backed implementation fact; inferred user impact.
- **Evidence:** Line 18 uses `<div class="submit" onclick="location.hash='submitted'">Continue</div>`. In the complete fixture, this element has no native control semantics, `role`, `tabindex`, or keyboard handler. Its only declared activation mechanism is `onclick`.
- **Reproduction or inspection steps:** Open the fixture source; inspect line 18; identify the element type, declared event handlers, and absence of keyboard-focus and control-semantic code elsewhere in the file.
- **User impact:** A primary action can be absent from sequential keyboard navigation and can lack a programmatic button role. This is a likely task blocker rather than a cosmetic issue.
- **Authoritative guidance:** Provisionally related to [WCAG 2.2 SC 2.1.1 Keyboard](https://www.w3.org/WAI/WCAG22/Understanding/keyboard.html) and [SC 4.1.2 Name, Role, Value](https://www.w3.org/WAI/WCAG22/Understanding/name-role-value.html). Exact current wording was not live-verified in this source-only review.
- **Remediation direction and owner:** The action element at line 18 owns the correction. Use a native `<button>`; choose `type="submit"` if the intended owner is an actual account-creation form, otherwise use `type="button"`. Preserve visual styling through the existing class and implement the intended outcome through the button/form behavior rather than recreating button semantics on a `div`.
- **Verification route:** In the rendered initial state, verify that Tab reaches the control with a visible focus indicator; Enter and Space perform the same action as pointer activation; the accessibility tree exposes a button named “Continue”; and the action is usable with a screen reader, speech input, and switch control. Verify the resulting state, not just activation.
- **Confidence and open questions:** High confidence in the implementation defect. Runtime focus behavior, assistive-technology output, and whether the intended action submits a form or only navigates remain unverified.

### [P1] Email Field Has No Programmatically Associated Label

- **Affected users and tasks:** Screen-reader and speech-input users may not receive or be able to target the field by the visible label. Users with cognitive disabilities may also encounter inconsistent label activation. The affected task is entering an email address.
- **Evidence level:** Source-backed implementation fact; inferred user impact.
- **Evidence:** Line 15 contains `<label>Email address</label>` and line 16 contains `<input type="email" />`. The label does not wrap the input and has no `for` attribute; the input has no `id`, `aria-label`, or `aria-labelledby`. The complete file contains no other naming relationship.
- **Reproduction or inspection steps:** Inspect lines 15–16; compare the label’s association mechanism with the input’s identifiers and naming attributes; confirm no association is established elsewhere in the fixture.
- **User impact:** Visual proximity alone does not encode the label-input relationship. The field can therefore have no reliable programmatic name, making orientation and voice targeting difficult.
- **Authoritative guidance:** Provisionally related to [WCAG 2.2 SC 1.3.1 Info and Relationships](https://www.w3.org/WAI/WCAG22/Understanding/info-and-relationships.html), [SC 3.3.2 Labels or Instructions](https://www.w3.org/WAI/WCAG22/Understanding/labels-or-instructions.html), and [SC 4.1.2 Name, Role, Value](https://www.w3.org/WAI/WCAG22/Understanding/name-role-value.html).
- **Remediation direction and owner:** The field-label pair at lines 15–16 owns the correction. Give the input a stable `id` and set the label’s `for` to that value, or wrap the input inside the label. Retain the visible “Email address” text. Add form-processing attributes such as `name` and appropriate `autocomplete` only according to the actual form contract.
- **Verification route:** Inspect the rendered accessibility tree for an email field named “Email address”; verify clicking the visible label focuses the field; and verify the field can be found and operated by label with a screen reader and speech input.
- **Confidence and open questions:** High confidence that the source lacks an explicit label association. The browser-computed name and behavior were not rendered, and the wider form contract is unavailable.

### [P1] Required Error Is Not Tied to the Field or a Defined Validation State

- **Affected users and tasks:** Screen-reader users may not discover which field the error belongs to or hear it when validation changes. Users with cognitive or learning disabilities may receive an error before they have acted and without a clearly stated requirement. The affected task is identifying and correcting an empty email value.
- **Evidence level:** Source-backed implementation fact; inferred user impact and timing risk.
- **Evidence:** Line 17 renders `<p class="error">Required</p>` in the initial document. It has no `id`; the input at line 16 has no `required`, `aria-required`, `aria-invalid`, `aria-describedby`, or `aria-errormessage`; and the source contains no validation logic or state transition. The error is visually adjacent, but the relationship is not encoded.
- **Reproduction or inspection steps:** Inspect lines 16–18; identify the initial presence of the error, the input’s declared state and relationships, and the absence of code that creates, removes, associates, or announces an invalid state.
- **User impact:** A user can encounter an unassociated “Required” message with no programmatic error relationship. If the message is intended to appear after validation, the source also provides no announcement or focus-management mechanism.
- **Authoritative guidance:** Provisionally related to [WCAG 2.2 SC 3.3.1 Error Identification](https://www.w3.org/WAI/WCAG22/Understanding/error-identification.html) and [SC 1.3.1 Info and Relationships](https://www.w3.org/WAI/WCAG22/Understanding/info-and-relationships.html). If the error is introduced dynamically, [SC 4.1.3 Status Messages](https://www.w3.org/WAI/WCAG22/Understanding/status-messages.html) may also be relevant; that mapping depends on the intended runtime behavior.
- **Remediation direction and owner:** The email field’s validation state owns this correction. State the requirement before entry when applicable; use native `required` when it matches the validation contract; give the error a stable identifier; associate explanatory/error text with the field; expose invalid state only when invalid; and render or reveal the message according to a defined validation policy. Choose focus movement or live announcement based on the complete submission interaction rather than adding an indiscriminate live region.
- **Verification route:** Exercise untouched, focused, empty-submission, invalid-email, corrected, and resubmission states. Confirm the visible message identifies the field and correction, the accessibility tree exposes the intended required/invalid state and description, keyboard focus follows the designed policy, and a screen reader announces the error at a useful time without duplicate output.
- **Confidence and open questions:** High confidence in the missing source relationships and undefined initial state. Intended validation timing, error lifecycle, submission behavior, and actual assistive-technology announcements are unknown.

### [P2] Primary Text Colors Specify Insufficient Contrast

- **Affected users and tasks:** Users with low vision, reduced contrast sensitivity, or color-vision differences may have difficulty reading the page heading and email label needed to understand and complete the task.
- **Evidence level:** Source-backed color values and contrast calculation; rendered perception unverified.
- **Evidence:** Line 7 specifies `body { color: #aaa; background: #fff; }`. The `h1` and `label` have no overriding color, so the source makes them inherit this pair. The calculated contrast ratio is approximately 2.32:1. This is below both the commonly applied 4.5:1 threshold for normal text and 3:1 threshold for qualifying large text. The error and Continue text use separate color rules and are not included in this finding.
- **Reproduction or inspection steps:** Inspect the inline CSS at lines 7–9; trace inheritance to the heading and label; calculate relative luminance contrast for `#aaa` against `#fff`; then confirm computed colors and text sizes at runtime.
- **User impact:** Key instructions can be difficult to perceive, increasing reading effort and the risk of missing the field label.
- **Authoritative guidance:** Provisionally related to [WCAG 2.2 SC 1.4.3 Contrast (Minimum)](https://www.w3.org/WAI/WCAG22/Understanding/contrast-minimum.html).
- **Remediation direction and owner:** The shared body color token/rule at line 7 owns the correction. Select a foreground color that meets the intended text-size threshold against the actual background in every state, while retaining distinguishable error and control styling. Validate the final computed colors rather than only the literal source values.
- **Verification route:** Run a contrast check against rendered computed styles for heading, label, error, control text, focus indicators, and every interactive state. Manually inspect at browser zoom and with platform contrast/forced-color settings, and include low-vision user evaluation where available.
- **Confidence and open questions:** High confidence in the specified color pair and calculated ratio. Actual font sizes, anti-aliasing, user overrides, forced-color behavior, and all state colors remain unverified.

## Checks Requiring Human or Assistive-Technology Testing

- **Keyboard and alternative-input completion:** A keyboard-only or switch-control tester should complete the email-and-Continue journey from the initial state through its resulting state, checking focus order, visible focus, activation parity, recovery from errors, and lack of traps. Source inspection cannot establish runtime focus or operation.
- **Screen-reader field and error experience:** Screen-reader users on at least the project’s supported browser/platform combinations should identify the page, field name/type/required/invalid state, associated error, Continue control, and resulting state. The accessibility tree and announcements were unavailable, and model inference cannot substitute for assistive-technology use.
- **Speech-input targeting:** A speech-input user should target “Email address” and “Continue” by their visible labels and complete correction after an error. Source inspection cannot establish generated target names or recognition behavior.
- **Low-vision, zoom, and reflow use:** A low-vision tester using magnification should inspect the journey at the project’s required zoom and viewport conditions, checking readable contrast, text resizing, clipping, overlap, horizontal scrolling, and maintained focus visibility. No viewport or computed layout was available; the source also has no viewport metadata, whose practical mobile effect requires rendering.
- **Validation comprehension:** A disabled-user review, including people with cognitive or learning disabilities where feasible, should assess whether the requirement, error timing, wording, correction path, and outcome are understandable. Markup alone cannot establish comprehension or lived usability.

## Material Limitations

- Review evidence is limited to one static source file. No browser, server, DOM, accessibility tree, computed style, screenshot, scanner, or runtime event trace was available.
- No initial, focus, hover, active, invalid, corrected, submitted, loading, success, or failure state was rendered. The source only declares a URL-fragment change and does not expose a complete account-creation process.
- No viewport, breakpoint, browser, operating system, input modality, assistive technology, user preference, localization, content variation, zoom level, or forced-colors mode was exercised.
- No project accessibility configuration or commands were inspected because the case scope allowed review of only the fixture. No automation pass or violation/incomplete count exists.
- No current web lookup was permitted. Criterion identifiers and links are provisional navigation aids derived from the supplied evidence reference and reviewer knowledge; confirm them against the current W3C recommendation before using them in a governed standards evaluation.
- No qualified accessibility specialist or disabled participant evaluated the page. User-impact statements are evidence-supported risks, not claims of observed lived experience.
- Source absence establishes what this fixture does not declare, but it cannot prove all browser-computed behavior. No missing evidence has been converted into a pass.

## Conclusion

This source-bounded first pass establishes four remediation priorities: replace the generic Continue element with a native control, create an explicit email label relationship, make validation state and error association coherent, and correct the shared low-contrast text color. The first external verification step should render the initial and validation/result states, then perform keyboard, accessibility-tree, screen-reader, speech-input, contrast, zoom, and reflow checks across supported environments.

This first pass does not prove accessibility, completeness, certification, or standards conformance.

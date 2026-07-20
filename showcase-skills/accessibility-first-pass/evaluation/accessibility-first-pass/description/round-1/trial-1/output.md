# Accessibility First Pass Report

## Review Scope

- **Target:** The account-creation form in `fixtures/form/index.html`.
- **Included journeys, states, routes, and viewports:** Initial empty form; keyboard progression from the email field; `Enter` after tabbing past the field; pointer activation of Continue; the resulting `#submitted` URL-fragment state; one default desktop Chrome viewport.
- **Available source and rendered access:** Complete single-file HTML/CSS/JavaScript source, a locally served rendering, browser DOM snapshot, computed-style inspection, keyboard and pointer interaction, and a visual screenshot inspected during the review.
- **Project tooling and commands used:** `python3 -m http.server 8765 --bind 127.0.0.1` to serve the fixture; in-app Browser/Chrome Playwright and keyboard control for rendered inspection; a local WCAG contrast-ratio calculation from computed CSS colors. No automated accessibility scanner was run.
- **Guidance baseline:** WCAG 2.2 Understanding documents were used as current authoritative guidance for narrow mappings. No conformance target or level was requested, so these mappings are review guidance rather than a conformance determination.
- **Exclusions:** Other routes or project files; mobile and responsive breakpoints; zoom/reflow; high-contrast and forced-colors modes; production submission behavior; browser and operating-system combinations beyond the rendered Chrome session; screen readers, speech input, switch control, magnification, and disabled-user testing.

## Evidence Summary

- **Observed (2026-07-20):** The rendered accessibility snapshot exposed the email field only as an unnamed `textbox` and Continue only as `generic`. After focus was placed in the email field and `Tab` was pressed, focus moved to `BODY`; pressing `Enter` left the URL fragment unchanged. Pointer activation of Continue changed the fragment from empty to `#submitted`, including with the email field empty. The rendered page displayed `Required` from initial load and produced no visible success or error state change after activation.
- **Observed:** Computed styles were `rgb(170, 170, 170)` on white for the 32 px bold heading and 16 px email label, `rgb(221, 0, 0)` on white for the 16 px Required text, and white on `rgb(34, 68, 102)` for Continue. Their calculated contrast ratios were approximately 2.32:1, 5.15:1, and 10.06:1 respectively.
- **Source-backed:** The visible `<label>` has no `for` attribute and the email `<input>` has no `id`, accessible-name attribute, `required`, `aria-invalid`, or `aria-describedby`. The Required paragraph has no identifier or programmatic relationship to the field. Continue is a scripted `<div>` with an `onclick` handler and has no native control semantics, role, or keyboard support (`index.html`, lines 15–18).
- **Tool:** No scanner evidence is available. Browser automation was used for DOM/accessibility snapshot inspection, computed-style reads, focus progression, keyboard activation, and pointer activation; it did not emit rule-based violation or incomplete counts.
- **Inference:** Users of screen readers, speech input, keyboard interfaces, and some alternative input devices are likely to be unable to identify or operate essential controls reliably. The always-visible Required text may be interpreted as an active error even before interaction, but its exact experience in assistive technology remains unverified.
- **Unverified:** Real account-creation validation, submission, loading, failure, and success states do not exist in the fixture. Assistive-technology announcements, zoom/reflow, focus appearance across browsers, touch behavior, and automated-scanner coverage were not tested.

## Prioritized Findings

### High: Continue Is Pointer-Only and Has No Button Semantics

- **Affected users and tasks:** Keyboard-only users; blind and low-vision users; people using switch control, speech input, or keyboard-emulating assistive technology; anyone who must activate Continue without a pointer.
- **Evidence level:** Observed and source-backed.
- **Evidence:** The accessibility snapshot exposed Continue as `generic`. Its computed `tabIndex` was `-1`. Tabbing from the email field moved focus to `BODY`, and `Enter` did not change the URL fragment; pointer activation changed it to `#submitted`. Source implements the action as `<div class="submit" onclick="...">`.
- **Reproduction or inspection steps:** Load the initial form, place focus in the email field, press `Tab`, then press `Enter`. Observe that Continue never receives focus and submission does not occur. Activate Continue with a pointer and observe the URL fragment change.
- **User impact:** The only apparent route forward in the account-creation task is unavailable through the tested keyboard path, making this a task-blocking barrier for users who cannot use a pointer. Assistive technology also receives no button role for the control.
- **Authoritative guidance:** [WCAG 2.2 SC 2.1.1 Keyboard](https://www.w3.org/WAI/WCAG22/Understanding/keyboard.html) requires functionality to be operable through a keyboard interface. [SC 4.1.2 Name, Role, Value](https://www.w3.org/WAI/WCAG22/Understanding/name-role-value.html) requires a programmatically determinable name and role for user-interface components and specifically identifies scripted `div` controls without a role as a common failure.
- **Remediation direction and owner:** At the form interaction owner, use a native `<form>` and `<button type="submit">Continue</button>`, then handle the form's submit event. Preserve native focus, `Enter`/`Space` activation, and button semantics instead of reconstructing them on a `div`.
- **Verification route:** Repeat keyboard-only traversal from page start through submission, confirming a visible focus indicator and activation with expected native keys. Inspect the accessibility tree for a uniquely named `button "Continue"`, then test with representative screen reader and switch-control combinations.
- **Confidence and open questions:** High confidence for the tested fixture. A separate hidden keyboard submission route was not present in source or observed behavior.

### High: The Email Field Has No Programmatically Associated Name

- **Affected users and tasks:** Screen-reader users identifying the requested input; speech-input users targeting the field by its visible label; users of software that derives form structure from accessible names.
- **Evidence level:** Observed and source-backed.
- **Evidence:** The rendered snapshot exposed an unnamed `textbox`. The visible label has no `for` value, the input has no `id`, and the label does not wrap the input. The input also has no `aria-label` or `aria-labelledby` fallback.
- **Reproduction or inspection steps:** Inspect the rendered accessibility snapshot or query the input's accessible name. Compare the visible `Email address` text with the unnamed textbox entry.
- **User impact:** Users may encounter an unidentified edit field and may be unable to determine or voice-target the information requested, preventing accurate form completion.
- **Authoritative guidance:** [WCAG 2.2 SC 1.3.1 Info and Relationships](https://www.w3.org/WAI/WCAG22/Understanding/info-and-relationships.html) requires presented relationships to be programmatically determinable and lists explicit label association as a sufficient HTML technique. [SC 4.1.2 Name, Role, Value](https://www.w3.org/WAI/WCAG22/Understanding/name-role-value.html) requires a programmatically determinable name for interface components.
- **Remediation direction and owner:** At the form markup owner, give the input a stable `id` and connect the existing visible label with `for`, or wrap the input inside the label. Prefer the visible label as the accessible-name source so the programmatic name matches the words users see.
- **Verification route:** Inspect the accessibility tree for `textbox "Email address"`; click the visible label and confirm it focuses the field; test navigation and field announcement with representative screen readers; verify speech input can target “Email address.”
- **Confidence and open questions:** High confidence. Actual screen-reader and speech-input behavior was not exercised.

### High: Essential Heading and Field Label Text Has Insufficient Contrast

- **Affected users and tasks:** People with low vision, reduced contrast sensitivity, or display/environment conditions that make faint text harder to perceive; the affected task is understanding the page purpose and identifying the email field.
- **Evidence level:** Observed.
- **Evidence:** Computed `#aaa` text on `#fff` measured approximately 2.32:1. This applies to the 16 px normal-weight email label and the 32 px bold heading. The ratio is below both the 4.5:1 threshold for normal text and the 3:1 threshold for large text. The Required text (approximately 5.15:1) and Continue text (approximately 10.06:1) met the respective text thresholds in the tested state.
- **Reproduction or inspection steps:** Load the page, read computed foreground/background colors for the heading and label, and calculate contrast using the WCAG relative-luminance formula without rounding the underlying result before comparison.
- **User impact:** The page title and the only field label can be difficult or impossible to read, impairing orientation and form completion even though the input itself remains visible.
- **Authoritative guidance:** [WCAG 2.2 SC 1.4.3 Contrast (Minimum)](https://www.w3.org/WAI/WCAG22/Understanding/contrast-minimum.html) requires at least 4.5:1 for normal text and 3:1 for large-scale text.
- **Remediation direction and owner:** At the shared color/style owner for this form, replace `#aaa` on white with a foreground color that meets the applicable contrast threshold in every used text size and weight. Validate the final computed colors rather than relying on visual estimation.
- **Verification route:** Recalculate contrast from rendered computed styles for every state, then manually inspect at browser zoom and under forced-colors/high-contrast settings to ensure information remains perceivable.
- **Confidence and open questions:** High confidence for the tested CSS colors and default rendering. Font smoothing, user overrides, other viewports, and forced-colors behavior were not tested.

### Medium: Required and Error Information Is Disconnected From Validation

- **Affected users and tasks:** Screen-reader users; people with cognitive, language, or learning disabilities; users who need clear required-field and error-recovery instructions.
- **Evidence level:** Observed, source-backed, and inference.
- **Evidence:** `Required` is displayed from initial load, but the input lacks the native `required` attribute and has no `aria-invalid` state. The paragraph has no `id`, and the input has no `aria-describedby`; therefore the visible message is not programmatically connected to the email field. Activating Continue with an empty field still changed the fragment to `#submitted`, and no visible validation transition occurred.
- **Reproduction or inspection steps:** Load the page without interacting and observe Required already visible. Inspect the input and error paragraph attributes. Leave the field empty, activate Continue with a pointer, and observe the submission fragment without a validation response.
- **User impact:** Users receive contradictory cues: the interface visually says the field is required while behavior accepts an empty value, and assistive technology has no encoded required state or field-to-message relationship. If the text is intended as an error, users may not know which field it describes or when the error occurred.
- **Authoritative guidance:** [WCAG 2.2 SC 1.3.1 Info and Relationships](https://www.w3.org/WAI/WCAG22/Understanding/info-and-relationships.html) applies to the programmatic relationship between instructions/errors and their field. [SC 3.3.2 Labels or Instructions](https://www.w3.org/WAI/WCAG22/Understanding/labels-or-instructions.html) requires labels or instructions when content accepts user input. If validation automatically detects an error, [SC 3.3.1 Error Identification](https://www.w3.org/WAI/WCAG22/Understanding/error-identification.html) requires the erroneous item to be identified and the error described in text.
- **Remediation direction and owner:** At the validation and form-state owner, define whether email is required and enforce that contract consistently. Add native `required` where applicable; keep the requirement in or associated with the label; hide error text until an error exists; give the message an `id`; connect it with `aria-describedby`; and set or clear invalid state in step with visible validation. Provide a specific message such as “Enter an email address” for an empty required field.
- **Verification route:** Exercise empty, malformed, corrected, submission-failure, and success states. For each state, inspect native validity and accessibility properties, verify keyboard focus/recovery, and test announcements with representative screen readers. Confirm the message remains understandable without color.
- **Confidence and open questions:** High confidence that the current states are disconnected. The intended production validation rules and whether `Required` is an instruction or an already-active error are unspecified.

## Checks Requiring Human or Assistive-Technology Testing

- **Screen reader:** After remediation, complete the initial, empty-error, malformed-email, corrected, and success journeys with VoiceOver/Safari and NVDA or JAWS/Chrome. Confirm field name, required/invalid state, error description, button role, focus movement, and outcome are understandable without visual context. Browser snapshots cannot establish actual announcement quality or task usability.
- **Speech input:** Target “Email address” and “Continue” by visible text and complete submission. This verifies that visible and programmatic names support voice operation, which markup inspection alone cannot establish.
- **Keyboard and switch input:** Traverse and operate every state using keyboard-only and representative switch-control software, checking focus visibility, order, activation, error recovery, and absence of task traps across supported browsers.
- **Low-vision review:** At 200% and 400% zoom and with platform contrast/forced-colors settings, confirm labels, errors, focus, and outcomes remain perceivable without two-dimensional scrolling for the intended responsive scope.
- **Disabled-user evaluation:** Include people who use screen readers, magnification, alternative input, or cognitive accessibility strategies in task-based account-creation testing; this first pass cannot infer lived usability from source and browser inspection.

## Material Limitations

- Only one static HTML route, one initial layout, one Chrome session, and a small set of interactions were exercised. No mobile breakpoint, browser matrix, operating-system matrix, zoom/reflow state, forced-colors mode, reduced-motion setting, or touch exploration was tested.
- No automated accessibility scanner was available in the exercised page context, so there are no scanner violation or incomplete counts. This also means rule coverage and scanner-detectable regressions were not assessed.
- No screen reader, speech input, switch control, magnifier, braille display, or disabled participant was used. The DOM snapshot establishes computed semantics only, not assistive-technology usability.
- The fixture implements submission only as a URL-fragment change. It contains no real account creation, asynchronous loading, server response, authentication, duplicate-account handling, or visible success state, so those material states remain outside the evidence.
- The screenshot and browser observations were inspected in-session and were not retained as repository artifacts. Reproduction steps and source locations are provided instead.

## Conclusion

This first pass establishes source-backed and directly observed barriers in the tested fixture. Remediation should first replace Continue with a native submit button and associate the visible email label, then correct the inherited low-contrast text color and implement one coherent required/error-state contract. The next verification step is to exercise all validation and outcome states with keyboard, automated scanning, and representative screen readers in supported browser combinations.

This first pass does not prove accessibility, completeness, certification, or standards conformance.

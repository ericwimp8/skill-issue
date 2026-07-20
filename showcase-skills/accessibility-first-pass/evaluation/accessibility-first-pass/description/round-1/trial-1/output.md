# Accessibility First-Pass Report

## Review Scope

- **Target:** The account-request form in `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html`.
- **Journeys and states:** Initial page source, email entry, displayed error text, and the source-defined submission action.
- **Source areas:** HTML document structure, native form semantics, inline CSS, required-field declaration, error relationship, and inline click handler.
- **Rendered environments:** None. This was a source inspection only; no browser, accessibility tree, keyboard session, screen reader, zoom, forced-colors mode, or mobile viewport was exercised.
- **Standards or guidance considered:** WCAG 2.2 success criteria relevant to keyboard access, names and relationships, form instructions, error identification, and contrast.
- **Out of scope:** Server responses, actual account creation, success and failure states after submission, production styles or scripts outside the fixture, browser and assistive-technology interoperability, and usability testing with disabled people.

## Methods and Evidence

| Method                                    | Target and state                                              | Result or evidence path                                                                                 | Evidence class | Limitations                                                                                        |
| ----------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | -------------- | -------------------------------------------------------------------------------------------------- |
| Manual source inspection                  | Complete initial HTML, CSS, and inline script                 | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html` | Observed       | Source cannot establish rendered or announced behavior across browsers and assistive technologies. |
| Contrast calculation from declared colors | `#777777` body text on `#ffffff`                              | Calculated ratio: 4.478:1                                                                               | Observed       | Applies only where these declared colors render unchanged and the text is normal-sized.            |
| Interaction-path trace                    | `onclick` on the “Send request” `div` through `form.submit()` | Lines 16–20 of the fixture                                                                              | Observed       | The interaction was not executed in a browser. Resulting network and response states are unknown.  |
| Automated accessibility scan              | Fixture                                                       | Not run                                                                                                 | Unverified     | No scanner result is available; automated scans would still require human evaluation.              |

## Material Limitations

- Findings are bounded to the supplied static fixture and source-defined initial state.
- Keyboard behavior, focus appearance and order, spoken announcements, voice-input operation, zoom/reflow, forced colors, touch target usability, and post-submit behavior remain unverified.
- No requested conformance target was supplied; WCAG 2.2 mappings below identify relevant review criteria rather than establishing conformance.

## Prioritized Findings

### High — Submission control is pointer-only and lacks button semantics

- **Evidence class:** Observed in source; user impact is inferred pending browser and assistive-technology testing.
- **Affected users and task impact:** Keyboard, switch, screen-reader, and voice-input users may be unable to activate or reliably identify the control that completes the account-request journey.
- **Evidence:** The only submission control is a `div` with an `onclick` handler and no native button semantics, keyboard handler, role, or focusability (`index.html`, line 19).
- **Reproduction or inspection steps:** Open the fixture source; inspect the final form control on line 19; confirm the click action is attached only to a `div`.
- **Authoritative guidance:** [WCAG 2.2, 2.1.1 Keyboard](https://www.w3.org/TR/WCAG22/#keyboard) and [4.1.2 Name, Role, Value](https://www.w3.org/TR/WCAG22/#name-role-value).
- **Priority rationale:** This is the sole completion control for the primary task, and the source provides no non-pointer activation path.
- **Remediation direction:** At the form behavior owner, use a native submit button (`button type="submit"`) and handle submission from the form's submit event so native keyboard and accessibility semantics are preserved.
- **Human or assistive-technology follow-up:** In supported browsers, complete the form using keyboard only, a screen reader, switch control, and voice input; verify the control is announced as a button and activates with expected commands.
- **Confidence and limitations:** High confidence in the source defect; exact browser and assistive-technology impact is unverified.

### High — The email field has no persistent, programmatically associated label

- **Evidence class:** Observed in source; announcement behavior is unverified.
- **Affected users and task impact:** Screen-reader and voice-input users may have difficulty identifying or targeting the field. People with cognitive, attention, or memory disabilities may lose the instruction after entering text because a placeholder is transient.
- **Evidence:** The input has only `placeholder="Email address"`; there is no `label`, `aria-label`, or `aria-labelledby` associated with `id="email"` (`index.html`, line 17).
- **Reproduction or inspection steps:** Inspect line 17 and search the form for a `label` tied to `email`; none is present.
- **Authoritative guidance:** [WCAG 2.2, 1.3.1 Info and Relationships](https://www.w3.org/TR/WCAG22/#info-and-relationships) and [3.3.2 Labels or Instructions](https://www.w3.org/TR/WCAG22/#labels-or-instructions).
- **Priority rationale:** The issue affects the only data-entry field in the primary journey and can make the requested input difficult to understand or operate.
- **Remediation direction:** Add a visible `label` associated through `for="email"`; retain supplemental instructions separately when needed rather than using placeholder text as the label.
- **Human or assistive-technology follow-up:** Verify the accessible name with common screen-reader/browser combinations and confirm voice-input users can target the field using the visible label.
- **Confidence and limitations:** High confidence that no associated label exists; the exact accessible-name fallback differs by browser and assistive technology.

### High — Programmatic submission bypasses native form validation

- **Evidence class:** Observed from the source-defined DOM API call; resulting response behavior is unverified.
- **Affected users and task impact:** Users who omit or mistype the email may receive no native required-field prompt and may submit invalid data, making the account request error-prone or unsuccessful. This can especially affect users who rely on clear error prevention and recovery cues.
- **Evidence:** The input is marked `required` on line 17, but the click handler invokes `form.submit()` on line 19. That API submits the form without the normal interactive constraint-validation step.
- **Reproduction or inspection steps:** Trace the line 19 click handler to `HTMLFormElement.submit()` and compare it with native submit-button or `requestSubmit()` behavior.
- **Authoritative guidance:** [WCAG 2.2, 3.3.1 Error Identification](https://www.w3.org/TR/WCAG22/#error-identification) and [3.3.2 Labels or Instructions](https://www.w3.org/TR/WCAG22/#labels-or-instructions); applicability to the complete production error flow requires confirmation.
- **Priority rationale:** The sole submission path defeats the only declared validation rule on the form's primary field.
- **Remediation direction:** Submit through native form behavior and provide validation at the form owner; preserve focus and expose clear, field-associated errors when invalid input prevents completion.
- **Human or assistive-technology follow-up:** Attempt empty and malformed submissions in supported browsers with keyboard and screen readers; verify validation prevents submission, identifies the field, explains correction, and moves or retains focus predictably.
- **Confidence and limitations:** High confidence in the source-defined validation bypass; server-side validation and post-submit recovery were unavailable.

### Medium — Error text is unconditional, vague, and unrelated to the field

- **Evidence class:** Observed in source; announcement and user comprehension are inferred.
- **Affected users and task impact:** Screen-reader users may not discover which field the message belongs to. Users with cognitive, learning, language, or memory disabilities may find “Invalid value” insufficient to understand what failed or how to recover.
- **Evidence:** `Invalid value` is always present immediately after the input, but the input has no `aria-describedby`, `aria-errormessage`, or invalid state, and no script controls the message (`index.html`, lines 17–18).
- **Reproduction or inspection steps:** Inspect the error span on line 18 and trace all form markup and script; no explicit error relationship, condition, or correction guidance exists.
- **Authoritative guidance:** [WCAG 2.2, 3.3.1 Error Identification](https://www.w3.org/TR/WCAG22/#error-identification) and [3.3.3 Error Suggestion](https://www.w3.org/TR/WCAG22/#error-suggestion).
- **Priority rationale:** The message degrades error understanding and recovery in the primary journey, while the DOM adjacency may provide a limited visual clue.
- **Remediation direction:** Render the message only for the relevant invalid state, describe the specific problem and correction, associate it programmatically with the email field, and expose the field's invalid state.
- **Human or assistive-technology follow-up:** Trigger each validation state and verify visual clarity, reading order, screen-reader announcement, focus behavior, and recovery after correction.
- **Confidence and limitations:** High confidence in the missing source relationship; the actual production validation lifecycle was unavailable.

### Medium — Declared normal body text contrast falls just below 4.5:1

- **Evidence class:** Observed by calculation from source colors; rendered applicability is inferred.
- **Affected users and task impact:** People with low vision or color-vision differences may experience reduced readability where normal-sized text renders as `#777777` on white.
- **Evidence:** The body declares `color: #777` on `background: #fff` (`index.html`, line 8), which calculates to 4.478:1. That is below the 4.5:1 minimum for normal text.
- **Reproduction or inspection steps:** Expand `#777` to `#777777`, calculate relative luminance against `#ffffff`, and inspect rendered computed styles and text sizes.
- **Authoritative guidance:** [WCAG 2.2, 1.4.3 Contrast (Minimum)](https://www.w3.org/TR/WCAG22/#contrast-minimum).
- **Priority rationale:** The shortfall is small, but the body declaration has page-wide reach and may affect normal-sized text.
- **Remediation direction:** Adjust the shared body foreground color to provide a clear margin above 4.5:1 for normal text, then verify computed colors in the release build.
- **Human or assistive-technology follow-up:** Inspect computed styles in supported browsers, including forced-colors and high-contrast modes, and confirm which text inherits the body color.
- **Confidence and limitations:** High confidence in the declared-color calculation; user-agent styles may override the color for some form controls, and the rendered text sizes were not inspected.

## Passed Checks Within Tested Scope

| Check                             | Exact state and method                | Evidence                                   | Limitations                                                                |
| --------------------------------- | ------------------------------------- | ------------------------------------------ | -------------------------------------------------------------------------- |
| Document language                 | Source inspection of initial document | `<html lang="en">` on line 2               | No verification that all rendered content is English.                      |
| Page title                        | Source inspection of initial document | `<title>Account request</title>` on line 6 | Browser presentation was not inspected.                                    |
| Main landmark and page heading    | Source inspection of initial document | One `main` and one `h1` on lines 14–15     | Reading order and navigation were not exercised with assistive technology. |
| Declared error-text contrast      | Source color calculation              | `#dd1111` on white calculates to 5.045:1   | Applies only if rendered unchanged at normal text size.                    |
| Declared submission-text contrast | Source color calculation              | White on `#2457d6` calculates to 6.160:1   | The element remains inaccessible as a control despite the color ratio.     |

## Follow-Up Checks and Unknowns

| Check                              | Why unresolved                                                | Required environment or method                                                                    | Expected behavior                                                                                           |
| ---------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| Complete keyboard journey          | No rendered session                                           | Supported desktop browsers, keyboard only; test empty, invalid, valid, and resubmission states    | Logical focus order, visible focus, all controls operable, no pointer dependency.                           |
| Screen-reader interaction          | Source cannot establish announcements                         | VoiceOver/Safari and NVDA/Firefox or NVDA/Chrome                                                  | Correct field and button names, roles, required/invalid states, error announcements, and predictable focus. |
| Voice and switch input             | No device or software session                                 | Representative voice-control and switch-control setup                                             | Controls targetable by visible names and operable without fine pointer input.                               |
| Zoom, reflow, and text spacing     | Static source only                                            | Browser at 200% and 400% zoom, 320 CSS-pixel width, and WCAG text-spacing overrides               | Content remains readable and operable without loss, overlap, or two-dimensional scrolling for the form.     |
| Forced colors and focus visibility | User-agent rendering unknown                                  | Windows High Contrast/forced-colors and keyboard focus inspection                                 | Text, field boundaries, errors, and focus indicators remain perceivable.                                    |
| Touch target behavior              | No mobile rendering                                           | Supported mobile browsers and touch testing                                                       | Submission target is large enough and reliably activates without accidental input.                          |
| Error and success recovery         | No server or application response supplied                    | Integrated release candidate with representative network failure, validation failure, and success | Status is clear, announced when appropriate, preserves entered data, and offers a reliable next step.       |
| Disabled-user usability review     | First-pass source review cannot establish practical usability | Representative participants using their usual assistive technologies                              | The account-request journey is understandable, efficient, and recoverable in real use.                      |

## Prioritized Next Actions

1. Replace the clickable `div` with a native submit button and submit through the form's native event path.
2. Add a visible, programmatically associated email label.
3. Implement specific, conditional, field-associated validation and error recovery without bypassing constraint validation.
4. Darken the shared body text color enough to exceed the normal-text contrast threshold with a practical margin.
5. Run the hands-on keyboard, screen-reader, zoom/reflow, forced-colors, voice/switch, mobile, and integrated error/success checks listed above before release.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

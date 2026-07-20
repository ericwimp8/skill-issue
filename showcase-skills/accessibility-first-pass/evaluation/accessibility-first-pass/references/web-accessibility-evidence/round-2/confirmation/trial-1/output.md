# Account-request accessibility release triage

## Scope and evidence

- **Journey:** Enter an email address, understand validation, and send an account request from `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/form/index.html`.
- **Method:** Complete source inspection of the supplied HTML plus authoritative WCAG 2.2, WAI, and HTML Standard guidance. No browser, accessibility-tree, keyboard, assistive-technology, computed-style, network, or server-response test was run.
- **Release position:** Address the two High risks before release. Resolve or explicitly accept the Medium risks only after the targeted runtime checks below.

## Ranked findings

### 1. High — The send control has pointer-click behavior without native control semantics

- **Observed (source):** Line 19 implements “Send request” as a generic `div` with an `onclick` handler. It is not a native button and has no explicit role, keyboard handler, or focusability attribute.
- **Inferred impact:** Keyboard, switch, and speech-input users may be unable to reach or activate the primary account-request action. The risk affects the journey’s completion point; implicit form submission may be a browser-dependent workaround, but it is not an equivalent exposed control.
- **Runtime-dependent:** Actual tab order, focus indication, Enter/Space behavior, accessibility-tree role/name, voice-command matching, and implicit submission were not exercised.
- **Guidance mapping:** [WCAG 2.2 SC 2.1.1 Keyboard](https://www.w3.org/TR/WCAG22/#keyboard) and [SC 4.1.2 Name, Role, Value](https://www.w3.org/TR/WCAG22/#name-role-value) are normative mappings for the requested operability and programmatic semantics. [Technique H91](https://www.w3.org/WAI/WCAG22/Techniques/html/H91) is informative guidance supporting use of native HTML controls.
- **Direction:** Make form submission owned by a native submit button so platform keyboard, focus, role, and submission behavior are available together.

### 2. High — The required email field relies on placeholder text as its only visible label

- **Observed (source):** Line 17 contains a required input whose only identifying text is `placeholder="Email address"`; there is no associated `label` or other persistent instruction.
- **Inferred impact:** Users with cognitive, memory, or low-vision needs may lose the field purpose after entering text. Screen-reader and voice-input usability depends on the browser and accessibility mapping, which was not inspected.
- **Runtime-dependent:** The computed accessible name, placeholder announcement, label/command matching, and behavior after text entry remain unverified.
- **Guidance mapping:** [WCAG 2.2 SC 3.3.2 Labels or Instructions](https://www.w3.org/TR/WCAG22/#labels-or-instructions) is the normative mapping. [WAI Form Instructions](https://www.w3.org/WAI/tutorials/forms/instructions/) and [Technique H44](https://www.w3.org/WAI/WCAG22/Techniques/html/H44) are informative implementation guidance.
- **Direction:** Provide a persistent visible label programmatically associated with the input; retain separate instructions only where the expected format needs explanation.

### 3. Medium — The click path bypasses native required-field validation

- **Observed (source):** Line 19 directly invokes `form.submit()` even though the email input is `required`.
- **Observed (platform contract):** The [HTML Standard](https://html.spec.whatwg.org/dev/forms.html#dom-form-submit) defines `submit()` as bypassing interactive constraint validation and not firing a `submit` event.
- **Inferred impact:** Pointer users can initiate submission with an empty email, losing the browser’s normal required-field prevention and feedback on this primary journey.
- **Runtime-dependent:** Navigation, server-side rejection or acceptance, returned error content, retained input, and recovery behavior were not tested.
- **Guidance mapping:** [WCAG 2.2 SC 3.3.1 Error Identification](https://www.w3.org/TR/WCAG22/#error-identification) is only a tentative mapping because the source does not establish that an error is automatically detected and presented. The HTML Standard establishes platform behavior, not WCAG conformance.
- **Direction:** Route activation through normal native form submission and validate on the server; present recoverable errors if submission is rejected.

### 4. Medium — The static error text is not associated with the email field

- **Observed (source):** Line 18 always renders “Invalid value” in a standalone `span`; line 17 has no `aria-describedby`, `aria-errormessage`, or other source-level association. No conditional validation state exists in the supplied source.
- **Inferred impact:** The message may confuse all users before interaction and may not be announced with the field to screen-reader users. The primary field’s correction path can therefore become unclear.
- **Runtime-dependent:** Reading order, announcement timing, focus movement, invalid-state exposure, and whether any external runtime changes the message remain unverified.
- **Guidance mapping:** [WCAG 2.2 SC 1.3.1 Info and Relationships](https://www.w3.org/TR/WCAG22/#info-and-relationships) is tentative because the intended relationship is inferred from proximity. SC 3.3.1 is also tentative until automatic detection and error presentation are demonstrated.
- **Direction:** Render an error only for the applicable validation state, identify the field in error, associate the message with it, and provide actionable correction text.

## Verification plan

1. **Chrome on macOS, keyboard only:** From initial page load, Tab through the form, then use Enter and Space on the send action with empty, malformed, and valid values. Expect a visible focus indicator, a reachable send control, standard activation, and invalid input to remain on the page with field-specific feedback.
2. **Safari on macOS with VoiceOver:** In initial, focused, populated, invalid, and corrected states, inspect the email control and send action. Expect a persistent email label exposed as the field name, required/invalid state when applicable, the error announced only when relevant and associated with the field, and the send action announced as a button.
3. **Firefox on Windows with NVDA:** Repeat keyboard submission for empty, malformed, and valid email states. Expect logical focus order, equivalent names/roles/states, field-linked error output, and focus or announcement that makes correction immediately discoverable.
4. **macOS Voice Control and Windows Speech Access:** Say commands matching the visible email label and “Send request” in initial and populated states. Expect direct target selection and activation without coordinate clicking.
5. **Chrome computed styles:** Inspect body text, placeholder, error, send text, and visible focus in default, hover, focus, invalid, 200% zoom, and forced-colors states. Expect WCAG-applicable contrast thresholds, persistent focus visibility, and no clipped or obscured content. Source color declarations alone do not establish these rendered results.
6. **Browser plus real endpoint:** Submit empty, malformed, valid, duplicate, server-error, and network-failure states. Expect server validation, retained user input where safe, specific recovery guidance, a clear success state, and status/error announcements for assistive technology.
7. **Representative disabled-user review:** Exercise the complete request-and-recovery journey with keyboard, screen-reader, magnification, speech-input, and cognitive-access strategies. Expect independent completion without hidden instructions or ambiguous recovery.

## Limitations and release note

- No checks are recorded as passed because only source and standards guidance were inspected.
- Color contrast, focus, validation presentation, announcements, interoperability, and server outcomes remain unresolved until the named runtime checks are completed.
- Normative success criteria and informative techniques are identified separately. This first pass does not establish overall accessibility or WCAG conformance.

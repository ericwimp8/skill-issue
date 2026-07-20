# Behavior Ground Truth

## Case 1 — Account Request Form

- Source establishes a placeholder-only required email input, an unconditional unassociated error span, a click-handled generic submission element, and direct `form.submit()` invocation.
- Browser-computed accessible names, keyboard focus and activation, native validation presentation, computed styles, server behavior, and assistive-technology output remain inferred or unverified until exercised.
- A passing report must connect source evidence to affected users and task impact, preserve runtime uncertainty, prioritize the primary journey, direct remediation to the form and validation owners, and require exact rendered and assistive-technology follow-up.

## Case 2 — Preferences Dialog

- Source establishes a native opener and Save button, a `role="dialog"` section without an authored accessible name, a pointer-only generic close element, and open/close handlers that only toggle `hidden`.
- Source contains no authored focus placement, focus containment, Escape handling, focus restoration, background inertness, or modal-state behavior.
- A safe browser attempt was unavailable in this cycle, so focus movement, accessibility-tree output, keyboard behavior, visual presentation, and assistive-technology interoperability remain unverified.
- A passing report must describe the complete dialog journey, separate source absence from runtime impact, provide behavior-owner remediation direction, and specify exact browser, keyboard, and screen-reader follow-up.

## Case 3 — Clean Automated Result

- The supplied JSON establishes only that the synthetic checker reported no automatically detectable violations under an unspecified configuration.
- Source establishes color-only service status, generic elements for table-like service data, and a generic element rather than a heading for the visible page title.
- Matching visible and `aria-label` button strings support only a narrow source observation; computed names, announcement quality, interaction behavior, and broader workflow clarity remain unverified.
- A passing report must refuse a WCAG 2.2 AA conformance or absence-of-barriers conclusion, identify material source risks, bound any passed checks, and provide the human, manual, and assistive-technology work required before a conformance assessment.

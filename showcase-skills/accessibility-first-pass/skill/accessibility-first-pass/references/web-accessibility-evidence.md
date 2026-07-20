# Web Accessibility Evidence

## Select Authoritative Guidance

- Prefer the current [W3C WCAG overview](https://www.w3.org/WAI/standards-guidelines/wcag/) and the linked normative recommendation when a finding needs a WCAG mapping. W3C encourages use of the latest WCAG version; record the version and requested conformance level rather than assuming either.
- Use [WAI Easy Checks](https://www.w3.org/WAI/test-evaluate/easy-checks/) to inform a preliminary review. Preserve its boundary that these checks cover only a few issues and that a page may appear to pass while retaining significant barriers.
- Use the [WCAG-EM overview](https://www.w3.org/WAI/test-evaluate/conformance/wcag-em/) when explaining the additional scope, representative sampling, expertise, evaluation, and reporting needed for a conformance evaluation.
- Use the [WAI evaluation report guidance](https://www.w3.org/WAI/test-evaluate/report-template/) when a fuller WCAG evaluation report is requested after this first pass.
- Prefer project-owned tool configuration and the tool's primary documentation for scanner behavior. For example, Playwright's [accessibility testing guidance](https://playwright.dev/docs/accessibility-testing) and the [axe-core project](https://github.com/dequelabs/axe-core) both state that automated checks find only some accessibility problems and must be combined with manual assessment.

Check current sources when exact success-criterion text, status, tool coverage, rule tags, or platform behavior matters. Do not copy a generic checklist into the report as if every item was tested.

## Preserve Evidence Boundaries

- A success-criterion mapping explains why a finding matters; it does not establish that the rest of the criterion, page, process, or product passes.
- Scanner results apply to the inspected DOM and state under the recorded configuration. Hidden, inactive, later, authenticated, responsive, or otherwise unexercised content may require separate activation and scanning.
- An accessibility tree exposes computed semantics, not complete usability with a screen reader, speech input, switch control, magnification, or other assistive technology.
- Keyboard traversal by a reviewer can establish observed focus and operation in the tested environment. It cannot establish every input method, platform convention, or user experience.
- Source inspection can establish implementation facts and risks. It cannot prove rendered layout, computed contrast, timing, focus order, announcements, or assistive-technology behavior that was not executed.
- Testing with disabled people contributes evidence about real use. Keep participant scope and observed experience explicit rather than generalizing one result to universal accessibility.

## Refer Responsibly

Use criterion identifiers and concise paraphrases. Link to the authoritative source. Distinguish:

- a directly established failure;
- a likely failure requiring runtime or specialist confirmation;
- related best-practice guidance outside the declared standards baseline;
- an untested requirement.

Reserve conformance language for a separately governed evaluation that covers the complete conformance scope and requirements.

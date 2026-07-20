# Accessibility First-Pass Report

## Review Scope

- **Target:** Preferences dialog in `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dialog/index.html`.
- **Journeys and states:** Initial closed state, opening from the Preferences button, Save control exposure, and closing from the visible close control.
- **Source areas:** HTML structure, accessible semantics, and inline JavaScript at lines 9-22.
- **Rendered environments:** None. This was a source-only preliminary review.
- **Standards or guidance considered:** WCAG 2.2 Success Criteria 2.1.1, 2.4.3, and 4.1.2; WAI-ARIA Authoring Practices dialog guidance.
- **Out of scope:** Browser rendering, computed accessibility tree, CSS presentation, automated scanning, assistive-technology interoperability, zoom/reflow, contrast, target size, high-contrast mode, reduced motion, and representative user testing.

## Methods and Evidence

| Method | Target and state | Result or evidence path | Evidence class | Limitations |
| --- | --- | --- | --- | --- |
| Static source inspection | Closed and scripted open/close states | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dialog/index.html:9` | Observed | Establishes markup and implemented event behavior, not browser or assistive-technology output |
| Authoritative guidance review | Dialog keyboard, focus, and naming expectations | [W3C APG Dialog Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/), [WCAG 2.1.1](https://www.w3.org/WAI/WCAG22/Understanding/keyboard.html), [WCAG 2.4.3](https://www.w3.org/WAI/WCAG22/Understanding/focus-order.html), [WCAG 4.1.2](https://www.w3.org/WAI/WCAG22/Understanding/name-role-value.html) | Observed | APG is informative implementation guidance; WCAG mappings still require runtime confirmation where behavior is environment-dependent |
| Automated accessibility scan | Not run | Not run | Unverified | No scanner result can be claimed |
| Keyboard and screen-reader exercise | Not run | Not run | Unverified | Focus behavior, announcements, and complete interaction quality require runtime testing |

## Material Limitations

- Only the supplied HTML source was inspected. No rendered browser state, accessibility tree, keyboard sequence, or assistive-technology announcement was observed.
- The source does not state whether the dialog is intended to be modal. Modal-only requirements are therefore retained as follow-up rather than asserted as failures.
- The fixture contains no preference fields and no Save behavior, so the complete preference workflow and its validation or status messaging could not be reviewed.

## Prioritized Findings

### High: The visible close control is pointer-only

- **Evidence class:** Observed.
- **Affected users and task impact:** Keyboard, switch, speech-input, and some screen-reader users can open the dialog with the native Preferences button but have no implemented keyboard-operable way to dismiss it. This blocks completion of the dialog interaction without a pointer.
- **Evidence:** Line 13 uses a non-focusable `<span>` with an inline `onclick`; lines 20-22 provide the only close behavior. No keyboard handler or Escape handler exists.
- **Reproduction or inspection steps:** Inspect line 13; confirm the close glyph is a `<span>` without native control semantics or a tab stop. Inspect lines 15-22; confirm `closeDialog()` is reachable only from that pointer click handler.
- **Authoritative guidance:** [WCAG 2.1.1 Keyboard](https://www.w3.org/WAI/WCAG22/Understanding/keyboard.html) requires content functionality to be operable through a keyboard interface. The [APG dialog pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) recommends a visible button that closes the dialog and Escape dismissal.
- **Priority rationale:** Dismissal is a core, repeated dialog action; the source provides no comparable keyboard path or workaround.
- **Remediation direction:** Replace the span at the behavior owner with a native `<button type="button">` whose accessible name identifies the action, such as `Close preferences`. Wire the existing close behavior to that button and support Escape while the dialog is open.
- **Human or assistive-technology follow-up:** In each supported browser, open the dialog using only the keyboard, reach and activate the close button with Tab plus Enter/Space, then reopen and dismiss with Escape. Confirm the control is announced as a button with the intended name in supported screen readers and works with voice input.
- **Confidence and limitations:** High confidence from source. Runtime verification is still required after remediation.

### Medium: Opening and closing do not manage focus

- **Evidence class:** Observed implementation; user impact is source-backed inference pending runtime exercise.
- **Affected users and task impact:** Keyboard and screen-reader users may remain on the trigger after the dialog appears, receive no immediate dialog context, and navigate in an unexpected order. After dismissal, the implementation provides no explicit restoration to the invoking control.
- **Evidence:** The open handler at lines 17-19 only sets `hidden` to `false`; `closeDialog()` at lines 20-22 only sets it to `true`. Neither path calls `focus()` or otherwise manages focus.
- **Reproduction or inspection steps:** Inspect both event paths and confirm that their only effect is changing `dialog.hidden`. In a browser, focus Preferences, activate it, and record the active element; then dismiss and record focus again.
- **Authoritative guidance:** The [APG dialog pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) places initial focus inside an opened dialog and normally returns focus to the invoker on close. [WCAG 2.4.3 Focus Order](https://www.w3.org/WAI/WCAG22/Understanding/focus-order.html) requires sequential focus order to preserve meaning and operability.
- **Priority rationale:** The dialog remains discoverable in DOM order, but missing focus movement can make a newly opened context confusing or easy to miss; exact impact depends on browser and assistive technology.
- **Remediation direction:** On open, move focus to the most appropriate element inside the dialog, likely the dialog heading made programmatically focusable or the first meaningful control. On close, return focus to the Preferences trigger unless the workflow establishes a more logical destination.
- **Human or assistive-technology follow-up:** Verify the active element and spoken context on open and close with keyboard-only navigation and supported screen-reader/browser pairs.
- **Confidence and limitations:** High confidence that focus management is absent; medium confidence in the precise experienced impact until runtime testing.

### Medium: The dialog has no programmatically determinable name

- **Evidence class:** Observed.
- **Affected users and task impact:** Screen-reader and refreshable-braille users may encounter an unnamed dialog, making its purpose harder to identify when context changes.
- **Evidence:** The container at line 10 has `role="dialog"` but no `aria-labelledby` or `aria-label`. The visible heading at line 11 has no `id` and is not referenced by the dialog.
- **Reproduction or inspection steps:** Inspect lines 10-11 and confirm there is no programmatic naming relationship. In browser accessibility tools, inspect the dialog node after opening and record its computed name.
- **Authoritative guidance:** The [APG dialog pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) specifies that a dialog has a name through `aria-labelledby` referencing its visible title or through `aria-label`; this supports [WCAG 4.1.2 Name, Role, Value](https://www.w3.org/WAI/WCAG22/Understanding/name-role-value.html).
- **Priority rationale:** The role is exposed but its purpose is not programmatically attached, degrading orientation for affected users throughout the interaction.
- **Remediation direction:** Give the visible `Preferences` heading a stable `id` and set `aria-labelledby` on the dialog container to that `id` so the visible and accessible names stay aligned.
- **Human or assistive-technology follow-up:** Confirm the computed accessible name and the announcement when focus enters the dialog in supported screen readers.
- **Confidence and limitations:** High confidence from source; announcement wording remains unverified.

## Passed Checks Within Tested Scope

| Check | Exact state and method | Evidence | Limitations |
| --- | --- | --- | --- |
| Page language is declared | Static inspection of the document root | Line 2 sets `lang="en"` | Language accuracy for all future content was not assessed |
| Page has a descriptive title | Static inspection of the head | Line 6 sets the title to `Preferences` | Browser presentation was not exercised |
| Trigger and Save use native buttons | Static inspection of closed/open source states | Lines 9 and 12 use `<button>` elements | Save behavior, focus visibility, target size, and announcement were not tested |
| Dialog is excluded while closed | Static inspection of initial markup | Line 10 includes `hidden` | Runtime accessibility-tree removal was not independently confirmed |
| Visible dialog heading exists | Static inspection of open-state markup | Line 11 uses `<h2>` | The heading is not yet connected as the dialog's accessible name |

## Follow-Up Checks and Unknowns

| Check | Why unresolved | Required environment or method | Expected behavior |
| --- | --- | --- | --- |
| Modal versus non-modal contract | Product intent is unstated; source does not make outside content inert or set `aria-modal` | Confirm interaction design, then test pointer and keyboard access outside the open dialog | If modal, outside content is visually obscured and inert, focus remains within the dialog, and `aria-modal="true"` is applied only when behavior is genuinely modal; if non-modal, provide a logical route between dialog and page content |
| Focus visibility and order | No CSS or rendered interaction was exercised | Keyboard test in each supported browser at default and zoomed views | Focus is always visible and follows a logical sequence without escaping a modal dialog |
| Screen-reader announcements | Static markup cannot establish browser/AT interoperability | Test supported browser and screen-reader combinations | Opening announces the dialog name and role; controls have accurate names and roles; closing restores understandable context |
| Save workflow | No preference controls, save handler, validation, or success/error state exists in the fixture | Review the completed production workflow, including failure and success states | Every preference is labelled and keyboard operable; validation and save status are perceivable and focus is managed appropriately |
| Zoom, reflow, contrast, target size, high contrast, and text spacing | No styles or rendered viewport were available | Manual checks at required viewport, zoom, spacing, and forced-color settings | Content remains perceivable and operable without clipping, overlap, loss of focus, or color-only meaning |
| Automated rules | No scan was run | Run the project's approved accessibility scanner against closed and open states | Investigate all results manually; a clean scan remains limited evidence |

## Prioritized Next Actions

1. Replace the close span with a correctly named native button and add Escape dismissal.
2. Implement initial-focus placement and focus restoration, then keyboard-test the complete open/save/close journey.
3. Name the dialog from its visible heading with `aria-labelledby`.
4. Confirm whether the product intends a modal or non-modal dialog and implement the corresponding outside-content and focus behavior.
5. Before release, run the completed production dialog through approved automation plus targeted keyboard, screen-reader, zoom/reflow, high-contrast, and disabled-user testing; retain evidence for each supported environment.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

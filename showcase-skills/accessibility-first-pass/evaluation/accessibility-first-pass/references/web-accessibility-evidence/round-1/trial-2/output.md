# Accessibility First-Pass Report

## Review Scope

- **Target:** Preferences dialog in `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dialog/index.html`.
- **Journeys and states:** Initial page state; source paths for opening the dialog, saving, and closing it.
- **Source areas:** Complete HTML markup and inline JavaScript in the named fixture.
- **Rendered environments:** None. The page was not rendered or operated in a browser.
- **Standards or guidance considered:** WCAG 2.2 Level A and AA criteria relevant to keyboard operation, focus order, and programmatic name/role/state; WAI-ARIA 1.2 dialog requirements; WAI-ARIA Authoring Practices Guide (APG) modal dialog pattern as informative guidance if the intended interaction is modal.
- **Out of scope:** Rendered visibility, CSS presentation, browser accessibility-tree output, actual focus movement, keyboard operation, screen-reader announcements, browser/assistive-technology interoperability, pointer behavior, zoom/reflow, contrast, and overall conformance.

## Methods and Evidence

| Method                                   | Target and state                                        | Result or evidence path                                                                                                                                                                               | Evidence class | Limitations                                                                                                 |
| ---------------------------------------- | ------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------- | ----------------------------------------------------------------------------------------------------------- |
| Complete source inspection               | Initial markup and all inline dialog handlers           | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dialog/index.html`                                                                                             | Observed       | Establishes authored source only; no browser behavior was exercised.                                        |
| Authoritative guidance review            | Dialog name, keyboard operation, and focus expectations | [WAI-ARIA 1.2 dialog role](https://www.w3.org/TR/wai-aria/#dialog), [WCAG 2.2](https://www.w3.org/TR/WCAG22/), and [APG modal dialog pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) | Observed       | APG is informative; its modal guidance applies only if the product intends a modal dialog.                  |
| Automated checks                         | Not run                                                 | Not run                                                                                                                                                                                               | Unverified     | The requested review is source-only and no automated result would establish the runtime behaviors at issue. |
| Keyboard and assistive-technology checks | Not run                                                 | Not run                                                                                                                                                                                               | Unverified     | Requires a rendered page, browsers, keyboard input, and screen readers.                                     |

## Material Limitations

- This is source-only evidence. Screen-reader output, focus location, tab sequence, Escape behavior, accessibility-tree state, and actual visibility were not observed.
- The source does not state whether the preferences dialog is intended to be modal. Modal-only expectations are identified conditionally.
- No result in this report is a claim about behavior across browser and assistive-technology combinations.

## Established Source Behavior

- **Observed:** The document declares `lang="en"` and includes a `Preferences` document title.
- **Observed:** The opener and Save control are native `button` elements with visible text.
- **Observed:** The `section` has `role="dialog"` and is authored with `hidden` in the initial markup.
- **Observed:** Activating the opener's registered `click` handler assigns `dialog.hidden = false`.
- **Observed:** The `×` is a `span` whose inline `click` handler calls `closeDialog()`, which assigns `dialog.hidden = true`.
- **Observed:** No source code moves focus, returns focus, handles `Escape`, constrains sequential focus, makes background content inert, or supplies `aria-modal`.

## Prioritized Findings

### High: Close action lacks a source-established keyboard control

- **Evidence class:** Inferred.
- **Affected users and task impact:** People who navigate with a keyboard, switch input, or other keyboard-emulating input may be unable to reach or invoke the visible `×` close action, making dismissal difficult or blocked if no other supported dismissal path exists.
- **Evidence:** Line 13 authors the close action as a `span` with only `onclick="closeDialog()"`. The element has no native button semantics, `tabindex`, accessible name, or keyboard event handler. Lines 20–22 contain the only close function. This establishes the source risk; it does not establish an observed runtime failure.
- **Reproduction or inspection steps:** Inspect lines 13 and 20–22; search the complete fixture for another dismissal handler or keyboard-operable close control. None is authored.
- **Authoritative guidance:** WCAG 2.2 [SC 2.1.1 Keyboard](https://www.w3.org/TR/WCAG22/#keyboard) requires functionality to be operable through a keyboard interface. W3C technique [SCR35](https://www.w3.org/WAI/WCAG22/Techniques/client-side-script/SCR35.html) explains why scripted actions should use natively actionable links or buttons; techniques are informative rather than mandatory.
- **Priority rationale:** Dismissal is a core dialog task, the source contains only one close path, and the risk affects non-pointer operation. Runtime severity remains to be confirmed.
- **Remediation direction:** Make dismissal a native, visibly labelled `button` owned by the dialog behavior. If the interaction is modal, also support the expected Escape dismissal behavior and focus return.
- **Human or assistive-technology follow-up:** Complete the keyboard and screen-reader journeys specified below and confirm that dismissal is reachable, announced, and operable.
- **Confidence and limitations:** High confidence in the source gap; actual keyboard behavior and user impact were not exercised.

### Medium: The dialog has no source-established accessible name

- **Evidence class:** Inferred.
- **Affected users and task impact:** Screen-reader and refreshable-braille users may encounter an unnamed dialog, making the purpose of the newly presented region confusing.
- **Evidence:** Lines 10–14 contain `role="dialog"` and a visible `h2`, but the dialog has neither `aria-labelledby` referencing that heading nor `aria-label`. WAI-ARIA defines the dialog role's name as author-provided and requires an accessible name. Actual accessibility-tree computation and announcement were not observed.
- **Reproduction or inspection steps:** Inspect the complete dialog element and its descendants; confirm that the heading has no `id` and the container has no naming attribute.
- **Authoritative guidance:** [WAI-ARIA 1.2 dialog role](https://www.w3.org/TR/wai-aria/#dialog) requires authors to provide an accessible name. WCAG 2.2 [SC 4.1.2 Name, Role, Value](https://www.w3.org/TR/WCAG22/#name-role-value) requires programmatic names and roles for user-interface components.
- **Priority rationale:** The issue may degrade orientation for every screen-reader opening, but announcement and impact depend on the browser/assistive-technology combination.
- **Remediation direction:** Give the visible heading an identifier and reference it with `aria-labelledby` on the dialog, keeping the visible title as the semantic owner of the name.
- **Human or assistive-technology follow-up:** Inspect the accessibility tree and verify spoken/braille output in both named screen-reader combinations below.
- **Confidence and limitations:** High confidence in the missing author-provided naming relationship; runtime exposure remains unverified.

### Medium: Opening and dismissal have no source-established focus management

- **Evidence class:** Inferred.
- **Affected users and task impact:** Keyboard and screen-reader users may remain outside the dialog after opening it, traverse content in an unclear order, or lose their place after dismissal.
- **Evidence:** Lines 17–19 only remove `hidden`; lines 20–22 only restore it. There is no call to `focus()`, no focus-return logic, and no keyboard handler. Actual browser focus movement was not observed.
- **Reproduction or inspection steps:** Inspect both complete state-change handlers and search for focus-related code. None is present.
- **Authoritative guidance:** WCAG 2.2 [SC 2.4.3 Focus Order](https://www.w3.org/TR/WCAG22/#focus-order) requires a sequential focus order that preserves meaning and operability. If the intended pattern is modal, the informative [APG modal dialog pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) expects focus to move inside when opened, remain within the dialog's tab sequence, and generally return to the invoking control when closed.
- **Priority rationale:** The source risk affects the complete opening and dismissal journey, but correct focus behavior and modal intent require runtime and product confirmation.
- **Remediation direction:** Define whether the interaction is modal. For a modal dialog, move focus to the appropriate element on open, constrain the dialog's keyboard sequence, and restore focus to the opener on close. For a non-modal dialog, implement a logical focus order and an explicit, reachable dismissal path without claiming modal behavior.
- **Human or assistive-technology follow-up:** Exercise the exact keyboard journeys below, recording focus at each step and whether background content can be reached.
- **Confidence and limitations:** High confidence that explicit focus management is absent; runtime focus and interaction quality are unverified.

### Needs Verification: Modal semantics and background availability are undefined

- **Evidence class:** Unverified.
- **Affected users and task impact:** Keyboard and screen-reader users could encounter a mismatch between visual presentation, focus behavior, and accessibility semantics if the intended dialog is modal.
- **Evidence:** The source uses `role="dialog"` without `aria-modal`, inert background handling, or focus containment. It also contains no styling or product statement that establishes modal intent.
- **Reproduction or inspection steps:** Confirm the intended interaction with the product owner, then inspect the rendered state and accessibility tree while the dialog is open.
- **Authoritative guidance:** The informative [APG modal dialog pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) requires `aria-modal="true"` and inactive background content for a modal implementation.
- **Priority rationale:** Assigning a barrier priority before intent and rendered behavior are established would overstate source evidence.
- **Remediation direction:** Choose and implement one coherent dialog model. Apply modal semantics only when the rendered behavior is actually modal.
- **Human or assistive-technology follow-up:** Verify background reachability and announcement in the keyboard and screen-reader journeys below.
- **Confidence and limitations:** The missing modal implementation signals are observed; whether they are required is unknown.

## Passed Checks Within Tested Scope

| Check                                          | Exact state and method                   | Evidence                                       | Limitations                                                                     |
| ---------------------------------------------- | ---------------------------------------- | ---------------------------------------------- | ------------------------------------------------------------------------------- |
| Document language is authored                  | Initial source; direct markup inspection | Line 2 contains `lang="en"`.                   | Language pronunciation was not tested.                                          |
| Opener uses a native control with visible text | Initial source; direct markup inspection | Line 9 is a `button` containing `Preferences`. | Keyboard activation, computed name, and announcement were not tested.           |
| Save uses a native control with visible text   | Dialog source; direct markup inspection  | Line 12 is a `button` containing `Save`.       | Runtime visibility, focusability, activation, and announcement were not tested. |
| Initial source marks dialog hidden             | Initial source; direct markup inspection | Line 10 contains the `hidden` attribute.       | Rendered and accessibility-tree exclusion were not tested.                      |

## Follow-Up Checks and Unknowns

| Check                                 | Why unresolved                                                                                             | Required environment or method                                                                                                                                                                                                                              | Expected behavior                                                                                                                                                                                                                                   |
| ------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Keyboard open and entry               | Source does not establish runtime focus movement.                                                          | On desktop macOS in current stable Safari and Chrome, with a hardware keyboard and no screen reader: reload the closed state, `Tab` to Preferences, activate with `Enter`, repeat with `Space`, and record the focused element immediately after each open. | Both keys open the dialog; focus moves to a logical element inside it rather than remaining in background content.                                                                                                                                  |
| Keyboard traversal and dismissal      | The close `span`, tab order, Escape handling, and focus containment were not exercised.                    | In each browser above, open by keyboard; use `Tab` and `Shift+Tab` through the entire open state; activate every dismissal mechanism; press `Escape`; record every focus target and post-close focus.                                                       | A visible, named close control is reachable and operable; Escape dismisses if this is modal; focus never becomes lost; focus returns to Preferences after dismissal. If modal, focus does not enter background content while open.                  |
| VoiceOver announcement and navigation | Accessible name, role, state changes, and focus announcements depend on Safari/VoiceOver interoperability. | On macOS with current stable Safari and VoiceOver: from the closed state, use VoiceOver keyboard navigation to find and activate Preferences; inspect the dialog with the rotor and navigate its controls; close it; repeat after page reload.              | VoiceOver announces a dialog named “Preferences,” exposes Save and a named close button as operable controls, excludes hidden dialog content while closed, and announces/restores focus coherently on close.                                        |
| NVDA announcement and navigation      | Accessible-tree mapping and announcements depend on Firefox/NVDA interoperability.                         | On Windows with current stable Firefox and NVDA: from the closed state, use `Tab` and NVDA navigation to activate Preferences; inspect focus and browse modes; traverse and close the dialog; repeat after page reload.                                     | NVDA announces a dialog named “Preferences,” exposes Save and a named close button, does not expose hidden dialog content while closed, and reports focus consistently throughout opening and dismissal.                                            |
| Modal/background behavior             | Product intent and rendered behavior are unknown.                                                          | First confirm whether Preferences is intended to be modal. While open, attempt keyboard focus, VoiceOver navigation, and NVDA browse navigation into the opener/background; inspect `aria-modal` and background state in each browser accessibility tree.   | If modal, background content is unavailable to interaction and assistive-technology navigation while open, and the accessibility tree reports modal state. If non-modal, focus and reading order remain logical without falsely reporting modality. |
| Visible focus and rendered visibility | No CSS or browser rendering was exercised.                                                                 | In both keyboard browser journeys, observe focus indicators at 100% and 200% zoom in the closed and open states, including all controls and the point after close.                                                                                          | Every keyboard focus location is visibly indicated, the open dialog and close control remain perceivable, and focus is not obscured.                                                                                                                |

## Prioritized Next Actions

1. Replace the clickable `span` with a native, visibly named close button and verify its keyboard operation.
2. Programmatically name the dialog from the visible `Preferences` heading.
3. Confirm modal intent, then implement the corresponding focus, Escape, background, and focus-return behavior.
4. Run the specified keyboard journeys in Safari and Chrome on macOS.
5. Run the specified VoiceOver/Safari and NVDA/Firefox journeys and retain focus and announcement evidence.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

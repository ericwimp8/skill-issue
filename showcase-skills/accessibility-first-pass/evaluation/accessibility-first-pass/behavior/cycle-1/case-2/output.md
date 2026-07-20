# Accessibility First-Pass Report

## Review Scope

- **Target:** Preferences dialog in `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dialog/index.html`.
- **Journeys and states:** Closed dialog; opening from the Preferences control; expected operation and dismissal of the open dialog.
- **Source areas:** HTML semantics and inline JavaScript in the single fixture.
- **Rendered environments:** None completed. A local-file render was attempted in the available Chrome browser surface, but browser security policy blocked navigation before interaction.
- **Standards or guidance considered:** WCAG 2.2 success criteria 2.1.1 Keyboard and 4.1.2 Name, Role, Value; WAI-ARIA Authoring Practices Guide (APG) Dialog (Modal) Pattern as informative implementation guidance.
- **Out of scope:** Overall page or product conformance; styles, responsive layouts, zoom/reflow, contrast, high-contrast, motion, mobile/touch, production framework behavior, automated scanning, and assistive-technology interoperability.

## Methods and Evidence

| Method | Target and state | Result or evidence path | Evidence class | Limitations |
| --- | --- | --- | --- | --- |
| Complete source inspection | Dialog closed and source-defined open/close behavior | `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dialog/index.html` | Observed | Establishes authored markup and script, not every browser or assistive-technology effect. |
| Rendered keyboard and pointer inspection attempt | Local fixture in available Chrome surface | Navigation blocked by browser security policy before the fixture loaded | Observed tool result | No interaction result is claimed. The block is environment-specific and says nothing about dialog accessibility. |
| Authoritative-guidance review | Dialog semantics and interaction expectations | [WCAG 2.2](https://www.w3.org/TR/WCAG22/) and [APG Dialog (Modal) Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) | Observed reference review | APG is informative guidance; WCAG applicability and conformance require evaluation in context. |

## Material Limitations

- This is a source-based first pass of one fixture. Runtime focus order, focus visibility, computed accessibility tree, announcements, pointer target usability, and browser/assistive-technology interoperability were not established.
- No project-native automated accessibility checker was run or installed. A clean automated result, if later obtained, would not establish accessibility or conformance.
- The source shows a dialog role but does not explicitly declare modal behavior. The report evaluates the interaction as an intended modal preferences dialog because opening reveals a separate dialog surface and the requested target calls it a dialog; product intent must be confirmed.

## Prioritized Findings

### High — The visible close action is pointer-only

- **Evidence class:** Observed source; affected-user behavior is a source-backed inference pending runtime confirmation.
- **Affected users and task impact:** Keyboard, switch, speech-input, and some screen-reader users may be unable to activate the visible close action. Dismissing an important preferences surface can be blocked or require leaving/reloading the page.
- **Evidence:** `index.html` line 13 implements the close affordance as `<span id="close" onclick="closeDialog()">×</span>`. It has no native interactive semantics, keyboard handler, focusability, or authored accessible name. Lines 20–22 hide the dialog only when `closeDialog()` is invoked.
- **Reproduction or inspection steps:** Open the source; inspect lines 10–13 and 20–22; confirm the only authored dismissal control is the clickable span; in a safe rendered environment, open the dialog, use `Tab` and `Shift+Tab`, and try to focus and activate the × with `Enter` and `Space`.
- **Authoritative guidance:** WCAG 2.2 [SC 2.1.1 Keyboard](https://www.w3.org/TR/WCAG22/#keyboard) requires functionality to be operable through a keyboard interface. APG recommends a visible button in the dialog tab sequence that closes the dialog.
- **Priority rationale:** Dismissal is a core, repeated dialog task; source evidence is direct; and the authored implementation provides no non-pointer path to the visible close action.
- **Remediation direction:** At the dialog component/markup owner, use a native `<button type="button">` for dismissal, give it a purpose-revealing accessible name such as “Close preferences,” and route pointer and keyboard activation through the same close behavior.
- **Human or assistive-technology follow-up:** In current Chrome and Firefox on Windows with keyboard-only and NVDA, and Safari on macOS with keyboard-only and VoiceOver, verify the close button is announced by name and role, receives visible focus, and activates with standard button keys.
- **Confidence and limitations:** High confidence in the source defect. Exact announcement and activation behavior remains unverified because rendering was blocked.

### High — The dialog has no authored focus, Escape, or modal-background behavior

- **Evidence class:** Observed source absence; resulting runtime behavior and user impact are inferred.
- **Affected users and task impact:** Keyboard and screen-reader users may remain on the opener after the dialog appears, tab into background content, lose context, or lack an expected Escape dismissal. Users with attention, memory, or low-vision needs may find the active surface confusing or inefficient.
- **Evidence:** Lines 17–19 only set `dialog.hidden = false`; lines 20–22 only restore `hidden = true`. There is no focus move into the dialog, contained tab order, focus return, Escape listener, background inertness, or modal state logic.
- **Reproduction or inspection steps:** Inspect lines 15–22 for the complete script; confirm open and close only toggle `hidden`; then, in a safe rendered environment, record `document.activeElement` before and after opening, traverse forward and backward with `Tab`, press `Escape`, attempt pointer and keyboard interaction with the Preferences opener while the dialog is open, and confirm focus returns to the opener after dismissal.
- **Authoritative guidance:** The informative [APG Dialog (Modal) Pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) describes moving focus inside on open, containing the tab sequence, closing with Escape, making background content inert, and returning focus appropriately.
- **Priority rationale:** Focus and modality govern the entire dialog journey and can seriously impede repeated operation. Priority is High because the complete inline implementation contains none of the expected behavior, while the exact browser effect remains unverified.
- **Remediation direction:** At the dialog behavior owner, implement one coherent open/close lifecycle: place initial focus deliberately, constrain modal focus, make outside content inert while open, support Escape, restore focus to the invoking control, and expose modal semantics only when the implementation truly behaves modally. Prefer the native `<dialog>` element with `showModal()` where project/browser support and testing establish suitable behavior; otherwise implement and test the complete custom pattern.
- **Human or assistive-technology follow-up:** Test the full journey with keyboard-only in current Chrome, Firefox, and Safari; repeat with NVDA/Firefox on Windows and VoiceOver/Safari on macOS. Verify initial announcement and focus, forward/reverse containment, Escape, outside-content inoperability, visible focus, and focus restoration.
- **Confidence and limitations:** High confidence that the authored lifecycle is absent; medium confidence in the exact experienced severity until rendered and assistive-technology testing is completed.

### Medium — The dialog lacks an authored accessible name

- **Evidence class:** Observed source; announcement impact is inferred.
- **Affected users and task impact:** Screen-reader and refreshable-braille users may encounter a dialog role without a programmatic name, making the Preferences context harder to identify and distinguish from other dialogs.
- **Evidence:** Line 10 supplies `role="dialog"` but no `aria-labelledby` or `aria-label`. The visible `<h2>Preferences</h2>` on line 11 has no `id` and is not programmatically referenced by the dialog.
- **Reproduction or inspection steps:** Inspect lines 10–11; confirm the dialog has a role and visible heading but no authored naming relationship; in a rendered browser, inspect the computed accessibility tree and open the dialog with a screen reader to record the announced name and role.
- **Authoritative guidance:** APG says a dialog has an accessible name supplied by `aria-labelledby` referencing its visible title or by `aria-label`; WCAG 2.2 [SC 4.1.2 Name, Role, Value](https://www.w3.org/TR/WCAG22/#name-role-value) is the tentative normative mapping pending computed-tree confirmation.
- **Priority rationale:** The missing name degrades orientation but does not alone prove that the preferences task is blocked. Evidence is direct, while user impact depends on the browser/assistive-technology combination.
- **Remediation direction:** At the dialog markup owner, give the visible heading a stable `id` and reference it with `aria-labelledby` on the dialog. Keep the visible and programmatic title aligned.
- **Human or assistive-technology follow-up:** Inspect the computed accessible name in current Chrome, Firefox, and Safari, then verify the dialog is announced as “Preferences, dialog” or an equivalent understandable phrase with NVDA and VoiceOver when focus enters.
- **Confidence and limitations:** High confidence that no authored name exists; runtime accessible-name computation and announcement were not tested.

## Passed Checks Within Tested Scope

| Check | Exact state and method | Evidence | Limitations |
| --- | --- | --- | --- |
| Native semantics for opener and Save control | Source inspection of lines 9 and 12 | Both controls use native `<button>` elements and have visible text labels | Runtime keyboard activation, computed names, focus styling, and Save behavior were not exercised. |
| Dialog hidden before opening | Source inspection of line 10 | The dialog is authored with `hidden` | Initial rendered visibility was not exercised. |

## Follow-Up Checks and Unknowns

| Check | Why unresolved | Required environment or method | Expected behavior |
| --- | --- | --- | --- |
| Initial focus and announcement | Rendering was blocked; source has no focus code | Current Chrome/Firefox with NVDA and Safari with VoiceOver; open from the Preferences button | Focus moves to an appropriate element inside; name and dialog role are announced. |
| Complete keyboard sequence | Browser interaction was unavailable | Keyboard-only forward/reverse traversal, activation, and Escape in current Chrome, Firefox, and Safari | All actions are operable; focus stays within a modal dialog; Escape closes; focus returns to opener. |
| Pointer dismissal and outside interaction | Pointer exercise was unavailable | Pointer test at common desktop and touch viewport sizes | Close target activates reliably; outside content is inert if the surface is modal; target size and spacing are usable. |
| Visible focus, zoom, reflow, contrast, and forced colors | No styles or rendered environment were tested | Manual inspection at 200%/400% zoom, narrow viewport, text spacing override, and forced-colors/high-contrast modes | Content remains perceivable and operable without obscuring controls or losing focus indication. |
| Representative disabled-user usability | A first pass cannot establish real-world strategy or efficiency | Moderated testing of the preferences journey with representative keyboard, screen-reader, low-vision, and limited-dexterity users | Users can identify, operate, dismiss, and recover context efficiently. |

## Prioritized Next Actions

1. Replace the close span with a named native button and verify keyboard activation.
2. Implement and unit-test a complete dialog open/close focus and modality lifecycle, including Escape and focus return.
3. Programmatically name the dialog from its visible Preferences heading.
4. Run the exact browser, keyboard, screen-reader, zoom/reflow, forced-colors, and pointer follow-up listed above; record browser and assistive-technology versions and actual results.
5. Reassess WCAG mappings and priorities from rendered evidence, then expand to a full-page/process evaluation only if conformance assessment is required.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

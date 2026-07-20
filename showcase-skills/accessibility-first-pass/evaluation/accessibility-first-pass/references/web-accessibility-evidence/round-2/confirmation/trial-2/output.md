# Preferences Interaction Accessibility Readiness

## Decision

**Not ready for accessibility acceptance.** The inspected source contains two directly evidenced barriers in the preferences dialog: the `dialog` has no authored accessible name, and its only explicit close control is a generic `span` activated by a pointer-oriented `onclick` handler. Source also provides no authored focus movement, focus containment, `Escape` handling, focus restoration, modal state, or background inertness. Their exact runtime and assistive-technology effects require the targeted tests below.

This source-only first pass does not establish overall accessibility or WCAG conformance.

## Scope And Methods

- **Scope:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dialog/index.html`, lines 9–22; preferences open, dialog, save, and close interaction.
- **Methods:** complete source inspection and comparison with normative WCAG 2.2 and WAI-ARIA 1.2 requirements plus informative ARIA Authoring Practices Guide (APG) dialog guidance.
- **Not run:** browser rendering, keyboard interaction, accessibility-tree inspection, screen-reader testing, automated scanning, zoom/reflow, contrast, and mobile/touch testing.
- **Limitation:** source establishes authored markup and event handling only. Browser focus, computed accessibility output, announcements, background interaction, and experienced user impact remain inferred or unverified until exercised.

## Prioritized Findings

### High — Close action is pointer-only in authored source

- **Evidence class:** Observed source evidence. Line 13 uses `<span id="close" onclick="closeDialog()">×</span>` with no native interactive element, role, keyboard handler, or focusability. Lines 20–22 show that this handler is the only authored close path.
- **Affected users and impact:** Inferred. Keyboard, switch, speech-input, and some screen-reader users may be unable to reach or activate the explicit close action, blocking or seriously impeding exit from an important settings interaction. A pointer remains an apparent workaround for users able to use one.
- **Shortest source check:** Open the fixture and inspect lines 13 and 20–22; confirm the close action is attached only to the `span`'s `onclick`.
- **Authority:** [WCAG 2.2 SC 2.1.1 Keyboard](https://www.w3.org/TR/WCAG22/#keyboard) is normative and requires functionality to be operable through a keyboard interface. [WCAG 2.2 SC 4.1.2 Name, Role, Value](https://www.w3.org/TR/WCAG22/#name-role-value) is normative for programmatically determinable component name and role.
- **Remediation direction:** Make closing a native `<button type="button">` with a programmatically determinable label such as “Close preferences,” and connect its activation to dialog closure. Preserve visible focus and support standard button activation rather than recreating button behavior on a `span`.
- **Priority rationale:** The source directly establishes the implementation defect; it affects a primary dialog action and can block completion for non-pointer users.

### High — Dialog has no authored accessible name

- **Evidence class:** Observed source evidence. Line 10 applies `role="dialog"`; line 11 provides a visible heading, but the dialog has neither `aria-labelledby` nor `aria-label` to make that heading its authored accessible name.
- **Affected users and impact:** Inferred. Screen-reader and refreshable-braille users may receive a dialog role without the “Preferences” context, making entry into the interaction confusing and increasing navigation effort. Exact computed name and announcement are unverified.
- **Shortest source check:** Inspect lines 10–11; confirm that `role="dialog"` is present and no naming relationship references the heading.
- **Authority:** [WAI-ARIA 1.2 `dialog` role](https://www.w3.org/TR/wai-aria-1.2/#dialog) is normative: authors must provide an accessible name using `aria-label` or `aria-labelledby`. WCAG 2.2 SC 4.1.2 is a related normative mapping; a final criterion determination still requires runtime accessibility-tree verification.
- **Remediation direction:** Give the heading a stable identifier and reference it from the dialog with `aria-labelledby`, retaining the visible “Preferences” title.
- **Priority rationale:** The missing authored name is directly established and removes essential context from a repeated, task-focused interaction.

### High risk, runtime verification required — Focus is unmanaged

- **Evidence class:** Source-backed inference. Lines 17–19 only remove `hidden`; lines 20–22 only restore it. There is no authored initial focus placement, containment, `Escape` handler, or focus restoration. Exact browser behavior and user impact are unverified.
- **Affected users and impact:** Inferred. Keyboard and screen-reader users may remain on background content after opening, tab outside the dialog, lack an efficient keyboard dismissal, or lose their place after closing. This can make the preferences journey disorienting or difficult.
- **Shortest source check:** Inspect lines 16–22 and confirm that the two state changes contain no focus or keyboard operations.
- **Authority:** WAI-ARIA 1.2 normatively recommends focusing and managing focus for modal dialogs with `SHOULD` language. The [APG modal dialog pattern](https://www.w3.org/WAI/ARIA/apg/patterns/dialog-modal/) is informative guidance for initial focus, contained `Tab`/`Shift+Tab`, `Escape`, and return focus behavior.
- **Remediation direction:** If this interaction is intended to be modal, prefer the native HTML `<dialog>` opened with `showModal()`, then deliberately establish initial focus and verify containment, `Escape` dismissal, and return focus. If it is intended to be non-modal, define and test the corresponding focus and dismissal contract without representing it as modal.
- **Priority rationale:** The likely journey impact is serious and the omissions are directly visible, but priority confidence is lower because runtime behavior and modal intent were not exercised.

### Needs verification — Modal state and background inertness

- **Evidence class:** Unverified behavior. The source has `role="dialog"` but no `aria-modal="true"`, native modal dialog invocation, or authored background-inert mechanism. The source does not establish whether modal behavior is intended or whether background content remains interactive in a browser.
- **Affected users and impact:** Unverified. If the interaction is intended to be modal, keyboard and assistive-technology users could encounter background controls while the dialog is open, causing context loss or accidental interaction.
- **Shortest source check:** Inspect lines 9–22; confirm there is no modal state or background-inert implementation, then perform the browser checks below.
- **Authority:** The APG modal dialog pattern informatively expects `aria-modal="true"` and background inertness for a modal implementation. It does not by itself establish WCAG failure or product intent.
- **Remediation direction:** Confirm product intent. For a modal interaction, use a correctly labelled native modal dialog where feasible and verify that visual, keyboard, pointer, and accessibility-tree modality agree.

## Precise Follow-Up Plan

### Keyboard — Chrome on Windows 11

1. Load the fixture with the dialog closed. Press `Tab` until “Preferences” has visible focus, then press `Enter`. Expect the dialog to open and focus to move to a sensible element inside it.
2. Press `Tab` and `Shift+Tab` repeatedly. If modal behavior is intended, expect focus to visit every operable dialog control in logical order, remain within the dialog, and remain visibly indicated.
3. Press `Escape`. Expect the dialog to close and focus to return to the “Preferences” opener.
4. Reopen the dialog, move focus to the visible close control, and activate it with both `Enter` and `Space` in separate runs. Expect closure and focus restoration after each activation.
5. While the dialog is open, attempt to reach or activate the background “Preferences” button with keyboard and pointer. If modal behavior is intended, expect background content to be unavailable until dismissal.

### Screen reader — NVDA with Chrome on Windows 11

1. Start NVDA, load the closed state, navigate to “Preferences,” and activate it. Expect focus to enter the dialog and NVDA to announce the accessible name “Preferences,” the `dialog` role, and the initially focused control without requiring background exploration.
2. Navigate through the dialog controls. Expect “Save” and “Close preferences” to be announced as buttons with usable names; expect no unlabeled “clickable” or multiplication-sign-only control.
3. Test `Tab`, `Shift+Tab`, `Escape`, and close-button activation. If modal behavior is intended, expect focus to remain in the dialog while open, dismissal to work, and the opener to be announced when focus returns.
4. Use browse mode while the dialog is open. If modal behavior is intended, expect background content to be unavailable as active dialog content; record the accessibility-tree and speech output if it remains reachable.

### Screen reader — VoiceOver with Safari on macOS

Repeat the NVDA journey using VoiceOver navigation and Safari. Expect the same name, role, operable controls, modal boundary, dismissal, and return-focus outcomes. Record any browser/assistive-technology difference rather than treating one combination as proof for the other.

## Next Actions

1. Replace the generic close `span` with a labelled native button.
2. Add an authored dialog name tied to the visible heading.
3. Confirm whether the interaction is modal and implement the corresponding focus and background-interaction contract.
4. Run the keyboard, NVDA/Chrome, and VoiceOver/Safari plans and retain exact focus, accessibility-tree, and announcement evidence.


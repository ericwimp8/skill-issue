# Accessibility First Pass Report

## Review Scope

- **Target:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/media/index.html` and only the resources it references.
- **Included journeys, states, routes, and viewports:** Initial page load; attempted video load; keyboard traversal; the `Read more` navigation failure path; default Chrome viewport at 1920 × 958 CSS pixels; narrow viewport at 320 × 568 CSS pixels.
- **Available source and rendered access:** Complete fixture source plus a local HTTP render in Chrome on macOS on 2026-07-20. The referenced `update.mp4` and `details.html` files were absent.
- **Project tooling and commands used:** `python3 -m http.server 4183 --bind 127.0.0.1`; in-app Browser DOM snapshots, screenshots, DOM evaluation, keyboard input, and navigation; `curl` HTTP-status checks. Repository dependency/configuration search found no configured axe, Lighthouse, Pa11y, or Playwright accessibility scanner.
- **Guidance baseline:** Current W3C WCAG 2.2 and WAI media/page-structure guidance were used only for narrow interpretation. No conformance level was declared, and this was not a conformance evaluation.
- **Exclusions:** Other fixtures and repository surfaces; replacement media; authenticated states; mobile devices; non-Chrome browsers; 200%/400% zoom; forced colors; reduced-motion behavior; screen readers and other assistive technologies; caption, transcript, and audio-description quality; disabled-user evaluation.

## Evidence Summary

- **Observed:** The page rendered `Product update`, `What changed`, a blank 300 × 150 video area, explanatory text, and `Read more`. The accessibility-oriented DOM snapshot exposed `main`, an h1, an h3, paragraph text, and the link, but no video. Browser inspection reported `autoplay: true`, `controls: false`, `paused: true`, `readyState: 0`, and zero text tracks. The server and `curl` both returned HTTP 404 for `update.mp4`; `Read more` navigated to an HTTP 404 for `details.html`. Keyboard Tab reached `Read more` with Chrome's visible focus outline, and the next Tab returned focus to the document because no other focusable control existed. At 320 CSS pixels the document had no horizontal overflow (`scrollWidth` and `clientWidth` were both 320).
- **Tool:** No automated accessibility scanner was configured or run. Automated violation and incomplete counts are therefore unavailable, rather than zero. Browser automation established rendered DOM, media state, keyboard focus, navigation, viewport, and HTTP behavior only.
- **Source-backed:** `index.html` declares `lang="en"`, a `main` landmark, h1 followed directly by h3, `<video src="update.mp4" autoplay>` without `controls` or `<track>`, and a link to `details.html`. The fixture directory contains only `index.html`.
- **Inference:** If a real video is later supplied, the current player configuration is likely to block users who need playback controls or time-based-media alternatives. WCAG applicability for captions, audio description, automatic audio, and motion cannot be determined until the media's audio, visuals, duration, and purpose are available.
- **Unverified:** Actual playback and autoplay policy with valid media; whether the content is audio-only, video-only, or synchronized media; whether speech or meaningful non-speech audio exists; whether important visual information exists; duration; whether `details.html` was intended as a transcript or media alternative; assistive-technology usability.

## Prioritized Findings

### High: The product-update content and its only linked detail route are unavailable

- **Affected users and tasks:** Everyone attempting to learn what changed is blocked from the video and the details page. The impact is especially material for people who would depend on a textual route because they cannot hear, see, process, or operate the video presentation.
- **Evidence level:** Observed
- **Evidence:** `update.mp4` and `details.html` are absent from the fixture. Both return HTTP 404. Chrome reserves a blank video box with `readyState: 0`, while `Read more` reaches the server's `Error response` page.
- **Reproduction or inspection steps:** Serve the fixture directory; open `/`; observe the blank media region; request `/update.mp4`; activate `Read more`; observe `/details.html` returning 404.
- **User impact:** The primary content is unavailable, and the linked route cannot currently serve as a transcript, description, or equivalent alternative. Users receive no page-authored explanation or recovery action.
- **Authoritative guidance:** WAI's [Making Audio and Video Media Accessible](https://www.w3.org/WAI/media/av/) recommends captions, description, transcripts, and an accessibility-supporting player according to the media's content. The absent resources prevent checking the applicable [WCAG 2.2 time-based-media requirements](https://www.w3.org/TR/WCAG22/#time-based-media).
- **Remediation direction and owner:** The fixture/media delivery owner should ship the intended media and details resources at the referenced deploy paths, or update the references to valid resources. If loading can fail in production, provide a visible, programmatically exposed failure message and a working recovery or equivalent-content route. If `details.html` is intended as an alternative, ensure it contains the equivalent information rather than relying on its filename or link wording.
- **Verification route:** Re-run HTTP, rendered, and keyboard checks against the deployed paths; then compare the complete video with the details content to establish whether an equivalent alternative is actually provided.
- **Confidence and open questions:** High confidence in the missing-resource failure. The intended content and purpose of both absent files are unknown.

### High: The video owner provides no operable player controls or evidenced media alternatives

- **Affected users and tasks:** Keyboard, switch, speech-input, and screen-reader users may be unable to start, pause, stop, seek, change volume, or select alternatives. Deaf and hard-of-hearing users may miss audio information; blind and low-vision users may miss visual information; people with cognitive or attention-related disabilities may be unable to stop distracting media.
- **Evidence level:** Source-backed
- **Evidence:** The sole video element has `autoplay` but no `controls`; it contains no `<track>`; no transcript or description is present in the page; and rendered keyboard traversal exposes only `Read more`. Chrome reported zero text tracks. Actual playback did not occur because the media file was unavailable.
- **Reproduction or inspection steps:** Inspect the `<video>` element; load the page; inspect its computed media state and accessibility-oriented DOM; traverse the page with Tab.
- **User impact:** Supplying `update.mp4` without changing the markup would leave no native control surface. Users could also lack the captions, transcript, or visual description required by the actual media content.
- **Authoritative guidance:** WAI says accessible media players need keyboard support, visible focus, and clear labels in its [Media Players guidance](https://www.w3.org/WAI/media/av/player/). WCAG 2.2 [1.2.1–1.2.3](https://www.w3.org/TR/WCAG22/#time-based-media) apply differently to prerecorded video-only, audio-only, and synchronized media. [1.4.2 Audio Control](https://www.w3.org/TR/WCAG22/#audio-control) applies if audio auto-plays for more than three seconds, and [2.2.2 Pause, Stop, Hide](https://www.w3.org/TR/WCAG22/#pause-stop-hide) applies to qualifying automatically started movement lasting more than five seconds. Those duration/content conditions remain unverified here.
- **Remediation direction and owner:** The video component/content owner should provide an accessibility-supporting player with keyboard-operable controls and avoid autoplay unless a demonstrated user need justifies it. With the real media available, provide accurate captions for meaningful prerecorded audio, and provide the applicable transcript, media alternative, and/or description for meaningful visual information. Keep each alternative adjacent to or clearly associated with the video.
- **Verification route:** Test valid media from load through completion with keyboard-only operation, screen reader, zoom, captions enabled/disabled, volume and pause controls, and the final transcript/description. Have a qualified reviewer compare every alternative against the actual audio and visuals.
- **Confidence and open questions:** High confidence that controls and in-page alternatives are absent. Whether each WCAG time-based-media criterion is a directly established failure cannot be determined without the media.

### Medium: The heading hierarchy skips from level 1 to level 3

- **Affected users and tasks:** Screen-reader users navigating by headings and people who rely on predictable visual/structural hierarchy may interpret `What changed` as a subsection whose parent section is missing.
- **Evidence level:** Observed
- **Evidence:** Source and rendered semantics expose `Product update` at level 1 followed immediately by `What changed` at level 3, with no level-2 heading.
- **Reproduction or inspection steps:** Inspect the heading elements or read the rendered accessibility-oriented DOM snapshot in document order.
- **User impact:** The rank gap can make the content outline harder to understand and navigate, even though both headings are programmatically exposed.
- **Authoritative guidance:** WAI's [Headings tutorial](https://www.w3.org/WAI/tutorials/page-structure/headings/) advises nesting headings by rank and avoiding skipped ranks where possible. Treating this specific gap as a WCAG 1.3.1 failure would require confirming the intended content relationships; this report records it as a structural best-practice risk.
- **Remediation direction and owner:** The page-content owner should use an h2 for `What changed` if it is a direct subsection of the h1, or add the genuinely missing parent section if the h3 relationship is intentional.
- **Verification route:** Confirm the intended document outline with the content owner, then inspect heading navigation in a screen reader after correction.
- **Confidence and open questions:** High confidence in the rank gap; moderate confidence in its impact because the intended hierarchy is undocumented.

## Checks Requiring Human or Assistive-Technology Testing

- **Media equivalence:** Obtain the real `update.mp4` and have a qualified reviewer compare its speech, non-speech audio, visuals, on-screen text, captions, transcript, and descriptions. The missing media prevents determining what alternatives are necessary or accurate.
- **Screen-reader journey:** With repaired resources and player controls, navigate headings, locate and identify the video, operate every control, enable captions/descriptions, and reach the details alternative in at least VoiceOver/Safari and NVDA or JAWS with a supported Windows browser.
- **Keyboard and alternate input:** Verify every player function, focus order, visible focus, seeking, volume, caption controls, and recovery behavior with keyboard-only, switch-style, and speech-input workflows.
- **Disabled-user evaluation:** Ask users who rely on captions, transcripts, audio description, magnification, or reduced distraction to complete the “learn what changed” task; record participant scope and observed outcomes without generalizing beyond them.

## Material Limitations

- The two referenced resources were absent. Their contents, duration, accessibility features, and intent were treated as missing evidence.
- No automated accessibility scanner was configured, so there are no violation or incomplete counts. DOM inspection and browser automation do not substitute for a scanner or manual assistive-technology review.
- Screenshots and browser state were inspected in-session and were not retained as report artifacts.
- Only Chrome on macOS, one default viewport, and one 320 CSS-pixel viewport were exercised. No physical device, browser matrix, zoom, forced-colors, reduced-motion, or operating-system accessibility configuration was tested.
- The blank reserved video area and lack of horizontal overflow at 320 CSS pixels establish only those rendered states; they do not establish reflow or media-player behavior once real media and controls are present.

## Conclusion

This first pass establishes that the delivered fixture cannot present the product-update video or its linked details, and that the owning markup lacks a usable player control surface and evidenced media alternatives. Restore the referenced resources first, then implement the player and content alternatives required by the real media, then correct and verify the document hierarchy. The next verification step is a rendered keyboard, automated, and assistive-technology review using the complete media and alternative content. This first pass does not prove accessibility, completeness, certification, or standards conformance.

# Accessibility First-Pass Report

## Review Scope

- **Target:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/media/index.html`
- **Journeys and states:** Initial page load and the authored product-tour presentation, assessed from source only.
- **Source areas:** The complete supplied HTML, inline CSS, and video element attributes.
- **Rendered environments:** None; no browser, device, media file, accessibility tree, or assistive technology was available in scope.
- **Standards or guidance considered:** WCAG 2.2 success criteria as normative requirements; W3C Understanding documents and techniques as informative guidance.
- **Out of scope:** `tour.mp4`, rendered playback, audio and visual content, autoplay-policy outcomes, browser-native media behavior, flashing analysis, and user or assistive-technology testing.

## Methods and Evidence

| Method | Target and state | Result or evidence path | Evidence class | Limitations |
| --- | --- | --- | --- | --- |
| Complete source inspection | Supplied `index.html`, initial authored state | Lines 8–9 author an infinite scale animation; line 15 authors an autoplaying, looping video with no `controls` attribute or `<track>` children; line 16 instructs users to watch the tour | Observed | Establishes authored source, not runtime behavior or media content |
| Semantic source inspection | Supplied document structure | English language, page title, one `main`, and one `h1` are present | Observed | Reading order and announcements were not rendered or tested |
| Project-native automated checks | None | Not run | Unverified | The requested assessment is source-only, and no checker result could resolve the missing media/runtime questions |

## Material Limitations

- The missing `tour.mp4` prevents determining duration, audio presence, speech or meaningful sound, visual-only instruction, embedded captions, audio description, flashing, or whether the paragraph is an adequate media alternative.
- Source attributes do not establish whether a particular browser permits autoplay, begins muted playback, exposes fallback media controls, or announces playback state.
- No rendered session establishes whether the scale animation is perceptible, whether motion remains under reduced-motion preferences, or whether either animation is essential.

## Prioritized Findings

### Medium — Continuous authored heading motion has no pause, stop, or hide mechanism

- **Evidence class:** Inferred user impact from observed source.
- **Affected users and task impact:** People who are distracted, fatigued, or made unwell by persistent motion may find reading the tour heading and adjacent instruction difficult. The source establishes repeated scaling, while the actual perceptual effect requires rendering.
- **Evidence:** Lines 8–9 apply `animation: pulse 0.7s infinite alternate` and scale the heading from `1` to `1.08`; line 14 applies that class. No source mechanism controls the animation.
- **Reproduction or inspection steps:** Open the supplied file; inspect the `.hero` rule, `@keyframes`, and the `h1` class.
- **Authoritative guidance:** Normative [WCAG 2.2 SC 2.2.2 Pause, Stop, Hide](https://www.w3.org/TR/WCAG22/#pause-stop-hide) requires a mechanism for automatically starting movement that lasts more than five seconds and appears beside other content, unless essential. The [Understanding SC 2.2.2 document](https://www.w3.org/WAI/WCAG22/Understanding/pause-stop-hide.html) is informative guidance about distraction and control.
- **Priority rationale:** The animation is authored to repeat indefinitely on a prominent heading and has no source-visible control, but source alone cannot establish its rendered intensity or individual user response. That supports Medium rather than inflated severity.
- **Remediation direction:** At the page-style owner, remove nonessential continuous motion or provide a persistent, keyboard-operable control that pauses or hides it. Supporting `prefers-reduced-motion` is useful informative direction but does not replace any applicable normative pause/stop/hide requirement.
- **Human or assistive-technology follow-up:** On macOS with Reduce Motion both off and on, render the initial page in current Safari and Chrome at 100% and 200% zoom. Confirm whether the heading visibly moves for more than five seconds, whether movement is essential, and whether a keyboard-only user can pause, stop, or hide it while continuing to read the page.
- **Confidence and limitations:** High confidence in the authored infinite animation; runtime effect, essentiality, and usability impact remain unverified.

### Medium — Autoplaying, looping tour is authored without visible playback controls

- **Evidence class:** Inferred behavior from observed source.
- **Affected users and task impact:** Keyboard users and people who need extra processing time or reduced motion may be unable to pause, replay, or stop a central instructional presentation if playback starts as authored.
- **Evidence:** Line 15 contains `<video autoplay loop src="tour.mp4"></video>` with no `controls` attribute. Line 16 directs users to watch the tour to learn the workflow.
- **Reproduction or inspection steps:** Inspect the video element and its adjacent instruction in the supplied source.
- **Authoritative guidance:** Normative [WCAG 2.2 SC 2.2.2](https://www.w3.org/TR/WCAG22/#pause-stop-hide) is a tentative mapping until rendered playback, duration, parallel presentation, and essentiality are confirmed. The W3C [G4 pause-and-restart technique](https://www.w3.org/WAI/WCAG22/Techniques/general/G4) is informative, not itself a requirement.
- **Priority rationale:** The tour appears instructional and is authored to loop without visible controls, creating a credible repeated-task barrier. Browser autoplay policy, media duration, and actual playback remain unknown, so the evidence supports Medium rather than High or Critical.
- **Remediation direction:** At the media presentation owner, use accessible native playback controls or equivalent keyboard-operable controls, avoid unsolicited playback, and ensure users can pause, stop, and replay the tour without losing content.
- **Human or assistive-technology follow-up:** In current Chrome on macOS, test the initial page with autoplay permitted and with the browser's default autoplay policy, using keyboard only. Expected: playback never removes the user's ability to pause, stop, replay, or continue reading, and every control has visible focus.
- **Confidence and limitations:** High confidence in the authored attributes and absent `controls`; actual playback and browser-provided fallback interaction are unverified.

## Passed Checks Within Tested Scope

| Check | Exact state and method | Evidence | Limitations |
| --- | --- | --- | --- |
| Document language is declared | Source inspection of the supplied initial document | `<html lang="en">` on line 2 | Pronunciation was not tested with a screen reader |
| Basic page structure is present | Source inspection | Descriptive `<title>`, one `<main>`, and one `<h1>` on lines 6 and 13–14 | Rendered hierarchy, navigation efficiency, and announcements were not tested |

## Follow-Up Checks and Unknowns

| Priority | Check | Why unresolved | Required environment or method | Expected behavior |
| --- | --- | --- | --- | --- |
| 1 | Information equivalence for the tour | The instruction suggests the video teaches the workflow, but the media and any off-page alternative are unavailable | Content-owner review of the complete `tour.mp4`; then a blind tester using VoiceOver with current Safari on macOS and a Deaf or hard-of-hearing tester reviewing the same initial-page state | All workflow information is available without seeing or hearing the media; any alternative follows the same meaningful sequence and supports the same task |
| 2 | Captions, transcript/media alternative, and audio description | No `<track>` is authored, but the source cannot reveal burned-in captions, a valid adjacent alternative, audio presence, or visually conveyed information | Inspect the full media with sound on and off in current Chrome on macOS; have a human verify caption timing, speaker identification, meaningful sounds, transcript completeness, and description of essential visuals | If synchronized audio exists, accurate captions satisfy normative [SC 1.2.2](https://www.w3.org/TR/WCAG22/#captions-prerecorded); applicable prerecorded alternatives/audio description satisfy normative [SC 1.2.3](https://www.w3.org/TR/WCAG22/#audio-description-or-media-alternative-prerecorded) and [SC 1.2.5](https://www.w3.org/TR/WCAG22/#audio-description-prerecorded) |
| 3 | Autoplay audio | The file may be silent, muted, shorter than three seconds, or blocked by browser policy | In current Chrome on macOS with site autoplay allowed and system volume audible, load the initial page from a clean tab and measure whether audio starts and continues beyond three seconds | No disruptive audio starts automatically; if it lasts more than three seconds, a page-level mechanism pauses/stops it or controls its volume independently, as required by normative [SC 1.4.2](https://www.w3.org/TR/WCAG22/#audio-control) |
| 4 | Screen-reader media semantics and state | Source does not expose the computed accessibility tree or announcements | VoiceOver with current Safari on macOS; navigate from the `h1` to the video in the initial, playing, paused, ended, and error states | The media and every available control have understandable name, role, value/state, order, and state-change announcements; operation does not trap focus |
| 5 | Motion and reduced-motion usability | CSS authors continuous motion, but the browser/device response and the video playback outcome are unknown | Current Safari and Chrome on macOS with Reduce Motion off/on; test initial load, active playback, 200% zoom, and keyboard-only operation | Nonessential continuous motion is absent or user-controllable; reduced-motion users are not forced through equivalent persistent motion |
| 6 | Flashing and seizure risk | Scale source alone does not establish luminance transitions, and the video frames are missing | Analyze the complete media in its rendered dimensions with a recognized flash-analysis tool, then manually review any flagged interval on the publication device/browser target | Content remains within normative [SC 2.3.1 Three Flashes or Below Threshold](https://www.w3.org/TR/WCAG22/#three-flashes-or-below-threshold); no conclusion is assigned from source alone |
| 7 | Representative task usability | Source cannot establish comprehension, fatigue, control discoverability, or workaround burden | Before publication, ask keyboard-only, blind screen-reader, Deaf/hard-of-hearing, low-vision, and motion-sensitive participants to learn and explain the workflow on the supported browser/device matrix | Each participant can obtain the workflow information, control playback and motion, and complete the intended learning task without avoidable exclusion or excessive effort |

## Prioritized Next Actions

1. Obtain and review `tour.mp4`; inventory speech, meaningful audio, essential visuals, duration, flashing, and all existing alternatives.
2. Render the initial page under permitted and blocked autoplay states; verify keyboard playback control and the WCAG 2.2.2 conditions for both the video and heading.
3. Add or correct equivalent media alternatives and accessible playback controls at the media owner based on the review.
4. Remove or control nonessential continuous heading motion and verify Reduce Motion behavior in Safari and Chrome.
5. Complete the named VoiceOver and disabled-participant task checks before publication.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

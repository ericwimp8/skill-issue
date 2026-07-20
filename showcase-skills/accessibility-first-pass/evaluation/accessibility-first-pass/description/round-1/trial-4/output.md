# Accessibility First-Pass Report

## Review Scope

- **Target:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/media/index.html`
- **Journeys and states:** Initial load of the product-tour page; static document structure; configured heading animation; configured video playback.
- **Source areas:** The supplied HTML and its inline CSS. The referenced `tour.mp4` was not supplied.
- **Rendered environments:** None. This was a source-only review; no browser, viewport, keyboard, screen reader, or media playback session was exercised.
- **Standards or guidance considered:** WCAG 2.2 Level A and AA criteria relevant to prerecorded media, automatic audio, and continuous motion.
- **Out of scope:** Other routes, production integration, external styles or scripts, actual media content, captions embedded in the media, browser autoplay policy, responsive rendering, contrast measurement, assistive-technology interoperability, and disabled-user testing.

## Methods and Evidence

| Method                                                                      | Target and state                            | Result or evidence path                                                            | Evidence class | Limitations                                                                            |
| --------------------------------------------------------------------------- | ------------------------------------------- | ---------------------------------------------------------------------------------- | -------------- | -------------------------------------------------------------------------------------- |
| Manual source inspection                                                    | Complete supplied HTML and inline CSS       | `index.html` lines 1–19                                                            | Observed       | Establishes authored markup and styles, not runtime behavior                           |
| Referenced-media inventory                                                  | Fixture directory                           | Only `index.html` is present; `tour.mp4` is unavailable                            | Observed       | Media duration, audio, visuals, embedded alternatives, and playback cannot be assessed |
| Project tooling inspection                                                  | Repository package scripts and dependencies | No dedicated accessibility checker or browser automation dependency was identified | Observed       | Tool inventory is not an accessibility result                                          |
| Automated accessibility scan                                                | Supplied page                               | Not run                                                                            | Unverified     | No project-native accessibility scanner was available                                  |
| Rendered keyboard, zoom, motion-preference, and assistive-technology checks | Supplied page                               | Not run                                                                            | Unverified     | No rendered session was exercised                                                      |

## Material Limitations

- The referenced `tour.mp4` is absent, so this assessment cannot determine whether the tour contains speech, meaningful audio, important visual-only information, embedded open captions, flashing content, or audio lasting more than three seconds.
- Browser autoplay behavior varies by browser and user settings. The authored autoplay request is observable; successful autoplay, automatic sound, and resulting user impact are unverified.
- No rendered inspection was performed. Focus behavior, native control exposure, visual contrast, zoom/reflow, text spacing, high-contrast behavior, and accessibility-tree output remain unverified.
- Source inspection cannot establish usability with screen readers, voice input, switch access, magnification, or other assistive technologies, and no testing with disabled people occurred.
- The findings are an initial risk assessment, not a complete WCAG audit.

## Prioritized Findings

### High — The primary tour has no user playback controls

- **Evidence class:** Inferred from observed source.
- **Affected users and task impact:** Keyboard, switch, speech-input, and limited-dexterity users may be unable to pause, replay, seek, or control the primary instructional content. People with cognitive, attention, or processing disabilities may be unable to stop or review a looping presentation at their own pace. The page explicitly directs users to watch the tour, so the affected task is central rather than incidental.
- **Evidence:** `index.html` line 15 authors `<video autoplay loop src="tour.mp4"></video>` without the native `controls` attribute or an authored equivalent. Whether the missing media would actually start is unverified.
- **Reproduction or inspection steps:** Inspect line 15; confirm `autoplay` and `loop` are present and `controls` is absent. In a browser with the production media, load the page, use keyboard-only input, and determine whether playback can be paused, restarted, sought, and volume-controlled.
- **Authoritative guidance:** [WCAG 2.2 SC 2.2.2 Pause, Stop, Hide](https://www.w3.org/WAI/WCAG22/Understanding/pause-stop-hide.html) requires a pause, stop, or hide mechanism for applicable automatically starting movement that lasts more than five seconds. [WCAG 2.2 SC 1.4.2 Audio Control](https://www.w3.org/WAI/WCAG22/Understanding/audio-control) applies if audio plays automatically for more than three seconds.
- **Priority rationale:** The tour is the page's stated instructional purpose, the loop is continuous if playback succeeds, and the authored element offers no direct control mechanism. Confidence in the source defect is high; runtime impact and the exact criterion applicability depend on the unavailable media and browser behavior.
- **Remediation direction:** At the media element, expose accessible native playback controls and remove automatic playback and looping unless product evidence establishes an essential need. Prefer user-initiated playback. If automatic behavior remains, provide an immediately operable pause/stop mechanism and independent audio control as applicable.
- **Human or assistive-technology follow-up:** With the production media, test current Chrome, Safari, and Firefox using keyboard-only input; then verify control names, roles, states, focus order, and announcements with VoiceOver and NVDA or another supported screen reader.
- **Confidence and limitations:** High confidence that no HTML controls are authored. Playback, duration, audio, browser-provided fallback behavior, and exact user impact were not observed.

### High — Continuous heading animation cannot be stopped

- **Evidence class:** Observed.
- **Affected users and task impact:** People with attention, cognitive, vestibular, or motion sensitivities may find the continuously scaling page heading distracting, fatiguing, or uncomfortable while trying to understand the tour. The animation affects the page's primary heading and continues alongside the other content.
- **Evidence:** `index.html` lines 8–9 apply an infinite alternating scale animation every 0.7 seconds to `.hero`; line 14 applies that class to the `h1`. No stop mechanism or reduced-motion override is authored.
- **Reproduction or inspection steps:** Inspect lines 8–9 and 14. In a browser, load the page and observe whether the heading continues moving beyond five seconds; then enable the operating system's reduced-motion preference and confirm whether the movement stops.
- **Authoritative guidance:** [WCAG 2.2 SC 2.2.2 Pause, Stop, Hide](https://www.w3.org/WAI/WCAG22/Understanding/pause-stop-hide.html) covers automatically starting, non-essential moving content that continues beyond five seconds in parallel with other content.
- **Priority rationale:** The movement is indefinite, rapid, prominent, and lacks a user control. It may interfere with the entire short page, and the source directly establishes the condition.
- **Remediation direction:** Remove the non-essential infinite animation. If a meaningful motion effect is retained, make it brief enough to stop within five seconds or provide a persistent pause/stop mechanism; also suppress non-essential motion when `prefers-reduced-motion: reduce` is active.
- **Human or assistive-technology follow-up:** Test the completed page with reduced motion enabled and with people who use motion-reduction settings; confirm that the heading remains stable without losing information.
- **Confidence and limitations:** High confidence in the authored continuous animation. Perceived severity and vestibular response require human evaluation.

## Passed Checks Within Tested Scope

| Check                                                  | Exact state and method | Evidence                                    | Limitations                                                                |
| ------------------------------------------------------ | ---------------------- | ------------------------------------------- | -------------------------------------------------------------------------- |
| Document language is declared                          | Source inspection      | `index.html` line 2 declares `lang="en"`    | Correctness assumes the page and media are English                         |
| Page has a descriptive title                           | Source inspection      | `index.html` line 6 contains `Product tour` | Browser announcement was not tested                                        |
| Main content and top-level heading use native elements | Source inspection      | `main` at line 13 and one `h1` at line 14   | Accessibility-tree exposure and full-site landmark context were not tested |
| Mobile viewport metadata is present                    | Source inspection      | `index.html` line 5                         | Reflow and zoom behavior were not tested                                   |

## Follow-Up Checks and Unknowns

| Check                                                   | Why unresolved                                                               | Required environment or method                                                                                                  | Expected behavior                                                                                                                       |
| ------------------------------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| Prerecorded captions                                    | Media is absent; no `<track>` is authored, but open captions may be embedded | Review `tour.mp4` with audio and inspect the production player                                                                  | All meaningful speech and non-speech audio has accurate synchronized captions where WCAG 2.2 SC 1.2.2 applies                           |
| Audio description or media alternative                  | Important visual-only information is unknown                                 | Review the tour's visuals and soundtrack with blind-user needs in mind                                                          | Important visual information is available through audio description or an applicable time-based media alternative under SC 1.2.3/1.2.5 |
| Transcript and equivalent task information              | The paragraph only says to watch; media content is unknown                   | Compare every instructional step in the tour with adjacent text                                                                 | Users who cannot consume the video can obtain equivalent workflow information                                                           |
| Automatic audio                                         | Media audio and actual autoplay behavior are unknown                         | Test production media across supported browsers and user settings                                                               | Audio does not start automatically for more than three seconds without independent pause/stop or volume control                         |
| Flashing and motion characteristics                     | Media is absent                                                              | Inspect the complete video with an appropriate flash-analysis tool and human review                                             | Media stays within applicable flash thresholds and provides motion controls where required                                              |
| Native video control accessibility                      | Controls are not currently authored                                          | After remediation, test keyboard, screen reader, touch, switch, and voice input                                                 | Every playback function is perceivable, operable, named, stateful, and reachable in a logical order                                     |
| Contrast, zoom, reflow, text spacing, and high contrast | No rendered review occurred                                                  | Test supported browsers at 200% and 400% zoom, constrained viewport, custom text spacing, and forced-colors/high-contrast modes | Content remains readable, visible, and operable without loss or obstruction                                                             |

## Prioritized Next Actions

1. Remove autoplay and looping from the tour, and expose native playback controls before publication.
2. Remove the infinite heading pulse or constrain it and honor reduced-motion preferences.
3. Supply and review the production `tour.mp4`; add accurate captions plus audio description or an equivalent media alternative when the content requires them.
4. Run rendered keyboard, zoom/reflow, contrast, reduced-motion, and accessibility-tree checks in supported browsers.
5. Complete targeted screen-reader and disabled-user testing of the full product-tour journey before making accessibility or conformance claims.

## Review Boundary

This first-pass review identifies evidence and risks within the stated scope. It does not establish overall accessibility, certification, or standards conformance. Automated results and limited manual checks cannot replace comprehensive evaluation and testing with disabled people and relevant assistive technologies.

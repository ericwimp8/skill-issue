# Accessibility First Pass Report

## Review Scope

- **Target:** The product-update article and embedded media implemented by `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/media/index.html`.
- **Included journeys, states, routes, and viewports:** Initial article load; keyboard traversal of available controls; embedded-video load state; and the article's `Read more` destination. Rendered checks used Chrome at 1920 x 902 CSS pixels with device pixel ratio 2.
- **Available source and rendered access:** Complete fixture HTML source and a locally served rendering on 20 July 2026. The fixture supplied no `update.mp4` or `details.html` file.
- **Project tooling and commands used:** Source inspection with numbered lines; `python3 -m http.server 8844 --bind 127.0.0.1`; browser DOM/accessibility snapshot inspection; bounded DOM state reads; keyboard Tab traversal; and server response-log inspection. No automated accessibility scanner was provided or run within this fixture-only review.
- **Guidance baseline:** No project conformance target was supplied. Current W3C WCAG 2.2 success criteria and WAI media/page-structure guidance are used only for narrow risk mappings, not as a conformance claim.
- **Exclusions:** Media-content review, caption accuracy, audio-description need and quality, transcript equivalence, playback duration, motion/audio behavior, mobile/reflow/zoom, contrast, multiple browsers, screen readers and other assistive technologies, and disabled-user testing. These could not be evaluated because the media and detail page were absent or because the bounded environment did not provide the required content or testers.

## Evidence Summary

- **Observed:** The rendered page exposed a `main` landmark, an English page language, a `Product update` title and level-1 heading, a level-3 `What changed` heading, and one `Read more` link. The video did not appear in the rendered DOM/accessibility snapshot. Its runtime state was `paused: true`, `readyState: 0`, with no text tracks or fallback text. The local server returned 404 for `update.mp4`. Keyboard traversal moved from `Read more` to the document body and back to `Read more`, without reaching a video control. Navigating to `details.html` produced a 404 error page.
- **Tool:** No rule-based accessibility scanner was run, so there are no violation, incomplete, or pass counts. Browser inspection established only the recorded DOM, runtime, and keyboard observations.
- **Source-backed:** `index.html` declares `<video src="update.mp4" autoplay></video>` with no `controls`, `<track>`, transcript link, fallback text, or adjacent media description. The heading sequence is `h1` then `h3`. The only follow-up link uses the accessible name `Read more` and targets `details.html`.
- **Inference:** If the intended video contains meaningful audio, visual information, or more than brief motion/audio, the current player configuration creates substantial caption, description, transcript, keyboard-control, and autoplay risks. The unavailable media prevents confirming which conditional WCAG requirements apply.
- **Unverified:** Whether the intended video has speech, meaningful non-speech audio, essential visual information, integrated description, open captions, motion lasting more than five seconds, or automatically playing audio lasting more than three seconds. Actual assistive-technology announcements and user experience also remain unverified.

## Prioritized Findings

### High — Referenced update resources are absent

- **Affected users and tasks:** Everyone trying to learn what changed; the impact is especially consequential for users who need a reliable alternate route after one format fails or who require predictable error recovery.
- **Evidence level:** Observed and source-backed.
- **Evidence:** `index.html` references `update.mp4` and `details.html`, but the supplied fixture contains only `index.html`. The server returned 404 for `update.mp4`; the video remained paused with `readyState: 0`; and the `Read more` route rendered `Error code: 404` and `File not found`.
- **Reproduction or inspection steps:** Serve the fixture directory; open `/`; inspect the video state and server log; then follow or navigate to `/details.html`.
- **User impact:** Both substantive ways to obtain the update fail. The page retains only the generic sentence “Learn about the new workflow,” so the promised update content is unavailable.
- **Authoritative guidance:** No narrow WCAG failure is asserted for the publishing error alone. Restoring the referenced resources is a prerequisite for evaluating the time-based media requirements described by [W3C WAI's audio and video accessibility guidance](https://www.w3.org/WAI/media/av/).
- **Remediation direction and owner:** At the article publishing/deployment owner, include the intended media and detail page at the referenced URLs or update the references to valid resources. Add publication checks that request every local content URL and fail on non-success responses. Provide a visible, useful fallback when media cannot load.
- **Verification route:** From a clean production-equivalent build, request both referenced URLs, play the video, follow the link with keyboard and pointer input, and confirm meaningful content loads rather than a server error.
- **Confidence and open questions:** High confidence that both supplied references are unavailable. The intended production packaging and whether either resource is injected elsewhere are unknown.

### High — No accessible media alternatives are represented

- **Affected users and tasks:** Deaf and hard-of-hearing people needing captions; blind and low-vision people needing important visual information described; DeafBlind people and people who process text more effectively needing a descriptive transcript; and people in environments where audio cannot be used.
- **Evidence level:** Source-backed, with conditional inference because the media content is unavailable.
- **Evidence:** The video has no `<track>` children, no nearby transcript or described-version link, and no fallback text. The runtime reported zero text tracks. The article's single generic sentence cannot be verified as an equivalent alternative to an unavailable video.
- **Reproduction or inspection steps:** Inspect the `<video>` element and its adjacent content; query its runtime `textTracks.length`; then review the intended media itself once supplied.
- **User impact:** If the video carries meaningful audio or visual information, affected users may receive an incomplete or empty product update.
- **Authoritative guidance:** WCAG 2.2 [1.2.2 Captions (Prerecorded)](https://www.w3.org/WAI/WCAG22/Understanding/captions-prerecorded) applies when prerecorded synchronized media contains needed audio information. [WAI caption guidance](https://www.w3.org/WAI/media/av/captions/) explains that captions include relevant speech and non-speech audio. WCAG 2.2 [1.2.3 Audio Description or Media Alternative](https://www.w3.org/WAI/WCAG22/Understanding/audio-description-or-media-alternative-prerecorded.html) and [WAI description guidance](https://www.w3.org/WAI/media/av/description/) address important visual information; [WAI transcript guidance](https://www.w3.org/WAI/media/av/transcripts/) explains descriptive transcripts and their users.
- **Remediation direction and owner:** At the media-content and article owner, first inventory the actual audio and visual information. Provide accurate synchronized captions for meaningful audio, and provide a descriptive transcript and/or audio description as required by the chosen target and the media. Link the transcript immediately beside the player and use a player that exposes each alternative.
- **Verification route:** Have qualified reviewers compare captions, transcript, and description against the full media; verify timing, speaker identification, meaningful sounds, on-screen text, and essential visual actions; then test discovery and operation with relevant assistive technologies and disabled users.
- **Confidence and open questions:** High confidence that no separate alternative is represented in this fixture; low confidence about the exact alternatives required until the media is available and reviewed.

### High — Autoplay is requested without user playback controls

- **Affected users and tasks:** Keyboard-only and switch users who need to control playback; screen-reader users whose speech may compete with audio; and people with cognitive, vestibular, attention, or sensory sensitivities who need to prevent or stop unexpected media.
- **Evidence level:** Source-backed and observed, with conditional WCAG mappings.
- **Evidence:** The element requests `autoplay` but omits `controls`. In the rendered state, Tab reached only `Read more`; no video control entered the focus order. Playback did not start because the media request failed, so actual autoplay behavior could not be observed.
- **Reproduction or inspection steps:** Inspect the video attributes; load the article; repeatedly press Tab from initial focus; then repeat after restoring a valid media asset and test every playback action.
- **User impact:** Once the media is restored, users may have no exposed way to start, pause, stop, replay, seek, adjust volume, or enable alternatives. Autoplayed sound can interfere with screen-reader output, and autoplayed motion can impair concentration or trigger discomfort.
- **Authoritative guidance:** [WAI media-player guidance](https://www.w3.org/WAI/media/av/player/) calls for keyboard-operable, labelled, visibly focused controls. WCAG 2.2 [1.4.2 Audio Control](https://www.w3.org/WAI/WCAG22/Understanding/audio-control.html) applies if audio automatically plays for more than three seconds. WCAG 2.2 [2.2.2 Pause, Stop, Hide](https://www.w3.org/WAI/WCAG22/Understanding/pause-stop-hide) applies if qualifying movement starts automatically and lasts more than five seconds. Those duration/content conditions are unverified here.
- **Remediation direction and owner:** At the player implementation owner, remove autoplay by default and add native `controls` or an established accessible media player. Ensure all functions and alternative-media controls are keyboard operable, named, focus-visible, and usable at zoom. Avoid custom controls unless they meet the same requirements.
- **Verification route:** With the real media loaded, test initial behavior and every control using keyboard only, screen reader, speech input, high zoom, and reduced-motion preferences. Measure whether any automatic audio or motion crosses the WCAG timing thresholds.
- **Confidence and open questions:** High confidence in the attributes and tested focus order; actual playback, duration, audio, motion, and player behavior remain unverified.

### Medium — Heading hierarchy skips level 2

- **Affected users and tasks:** Screen-reader users navigating by headings, and people with cognitive or visual disabilities who depend on a consistent document outline to understand section relationships.
- **Evidence level:** Observed and source-backed.
- **Evidence:** Both the source and rendered snapshot show `Product update` as level 1 followed immediately by `What changed` as level 3, with no intervening level-2 section.
- **Reproduction or inspection steps:** Inspect the heading elements or list headings in the rendered accessibility snapshot.
- **User impact:** The level jump implies an absent parent section and can make this short article's organization sound incomplete or confusing in heading navigation.
- **Authoritative guidance:** [W3C WAI heading guidance](https://www.w3.org/WAI/tutorials/page-structure/headings/) says to nest headings by rank and avoid skipped ranks where possible. This supports the structure preserved by WCAG 2.2 [1.3.1 Info and Relationships](https://www.w3.org/WAI/WCAG22/Understanding/info-and-relationships.html); this report treats the mapping as structural guidance rather than a page-wide conformance determination.
- **Remediation direction and owner:** At the article-content owner, change `What changed` to `h2` unless a genuine level-2 parent section is added before it.
- **Verification route:** Reinspect the rendered heading outline and navigate headings with a screen reader to confirm the announced levels match the intended content hierarchy.
- **Confidence and open questions:** High confidence. No evidence in the bounded article suggests a missing level-2 parent section is intentional.

### Low — The follow-up link name is generic

- **Affected users and tasks:** Screen-reader users navigating a links list and people with cognitive disabilities trying to predict a destination before activating it.
- **Evidence level:** Observed and source-backed.
- **Evidence:** The rendered accessible name is `Read more`. The preceding sentence supplies some programmatic context, but the name alone does not identify the new workflow or destination.
- **Reproduction or inspection steps:** Inspect the link's accessible name in the rendered snapshot, then consider it in a links-only navigation view.
- **User impact:** Users who encounter the link outside the surrounding sentence must spend more effort determining its purpose. The broken destination currently prevents confirming whether the nearby wording accurately describes the target.
- **Authoritative guidance:** WCAG 2.2 [2.4.4 Link Purpose (In Context)](https://www.w3.org/WAI/WCAG22/Understanding/link-purpose-in-context.html) permits purpose to be established by link text together with programmatically determined context, while advising meaningful link text whenever possible. A definite criterion failure is not asserted here because the paragraph supplies context.
- **Remediation direction and owner:** At the article-content owner, use a destination-specific name such as “Read more about the new workflow,” aligned with the restored destination's actual content.
- **Verification route:** Confirm the accessible name communicates purpose both in context and in a screen-reader links list, then follow it to verify the destination matches.
- **Confidence and open questions:** High confidence in the current name; whether it fails a declared standard remains open because no conformance target was supplied and contextual purpose may be sufficient.

## Checks Requiring Human or Assistive-Technology Testing

- **Media equivalence:** A qualified caption/description reviewer must inspect the restored video's complete audio and visual content, compare it with captions, transcript, and description, and confirm that all information needed to understand the update is equivalent. The media was unavailable in this review.
- **Screen-reader journey:** Screen-reader users should locate the article, navigate its heading structure, discover the player and alternatives, operate all playback functions, reach the transcript, and follow the detail link. A DOM/accessibility snapshot cannot establish usable announcements or interaction.
- **Keyboard, switch, and speech operation:** Test every player control, focus order, visible focus, names, hit targets, and error recovery with keyboard, switch, and speech input after selecting a real player. The current player had no exposed controls to test.
- **Disabled-user evaluation:** Include Deaf and hard-of-hearing, blind and low-vision, DeafBlind, and cognitive/vestibular participants in representative product-update tasks. Technical inspection alone cannot establish whether the media presentation meets real user needs.
- **Autoplay and sensory response:** With the intended media, verify default playback, audio and motion duration, pause/stop behavior, reduced-motion behavior, and comprehension without sound. The failed media request prevented these checks.

## Material Limitations

- The supplied implementation omitted both referenced resources, so the core media and detail content could not be evaluated.
- Only the article's initial state, failed resource states, one desktop viewport, and basic keyboard Tab order were rendered.
- No mobile viewport, zoom, text-spacing override, high-contrast mode, reduced-motion mode, alternate browser, or network-recovery state was tested.
- No automated accessibility scanner was run; therefore this report provides no scanner coverage or counts.
- No screen reader, magnifier, speech input, switch control, braille display, qualified media reviewer, or disabled participant was used.
- Source and snapshots establish implementation facts and risks, not caption quality, computed assistive-technology behavior, user comprehension, or standards conformance.

## Conclusion

This source-and-browser first pass establishes that both referenced update resources fail to load, the article represents no media alternatives, the player requests autoplay without controls, and the document hierarchy skips a heading level. First restore and production-verify the referenced resources; then inventory the real media, supply the necessary captions/transcript/description, and expose a non-autoplaying keyboard-operable player before assistive-technology and disabled-user verification. This first pass does not prove accessibility, completeness, certification, or standards conformance.

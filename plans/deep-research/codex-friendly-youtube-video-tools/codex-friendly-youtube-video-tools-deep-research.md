# Codex-Friendly YouTube Video Tools: Deep Research Synthesis

## Executive Recommendation

Use **macOS Screenshot capture + the already-installed FFmpeg + corrected YouTube sidecar captions** as the primary production path. Add two original 1920×1080 title-card images and animate them with simple fades or restrained movement in FFmpeg. This is the strongest fit for the stated priorities because it is free, immediately available, local, watermark-free by architecture, highly scriptable by Codex, and produces a reproducible H.264/AAC master that can be inspected with `ffprobe`. The important local constraint is that this FFmpeg build lacks `drawtext` and `subtitles`, so Codex should prepare title-card images rather than generating text inside FFmpeg, and captions should be uploaded as SRT/VTT rather than burned in. [Local capability assignment](assignments/01-local-codex-capabilities.md) [FFmpeg/macOS deep dive](assignments/07-ffmpeg-obs-macos-deep-dive.md) [Audio and captions assignment](assignments/12-audio-captions-accessibility.md)

Use **Microsoft Clipchamp Free** as the best single-tool fallback when a personal Microsoft account, Chrome or Chromium Edge, and a hosted editor are acceptable. Its official material establishes free unlimited 1080p exports without a watermark, plus screen/camera/voice recording, timeline editing, titles, transitions, audio, and automatic captions. Codex can substantially prepare and assist the edit, but the workflow remains supervised browser UI work rather than a dependable headless render pipeline. [Clipchamp pricing](https://clipchamp.com/en/pricing/) [Microsoft feature comparison](https://support.microsoft.com/en-us/clipchamp/feature-comparison-between-clipchamp-for-work-and-personal-versions) [Browser-editor deep dive](assignments/09-browser-editors-deep-dive.md)

Use **macOS Screenshot + Remotion** only if a more visibly designed intro, animated callouts, or a reusable code-defined composition is worth several additional setup and iteration hours. Remotion is the best polished code-authored option, but it is a fresh project, requires macOS 15+ and Node, needs licence eligibility confirmation, and was not rendered in this research run. [Remotion deep dive](assignments/06-remotion-deep-dive.md)

The governing delivery target is a **public YouTube video below three minutes with clear audio explaining what was built and how both Codex and GPT-5.6 were used**. Judges are not required to watch after three minutes. The project must work as depicted, and all third-party music, imagery, trademarks, footage, and other protected content require appropriate rights. [OpenAI Build Week Official Rules](https://openai.devpost.com/rules) [Rules assignment](assignments/02-build-week-video-rules.md)

### Recommendation at a Glance

| Need                                                 | Recommended choice                        |
| ---------------------------------------------------- | ----------------------------------------- |
| Best overall for this Mac and substantial Codex help | **Screenshot + FFmpeg + YouTube SRT/VTT** |
| Best single tool                                     | **Clipchamp Free**                        |
| Best higher-polish code-defined composition          | **Screenshot + Remotion**                 |
| Best simple local GUI if installed                   | **Screenshot + iMovie**                   |
| Best trialable automatic screen polish               | **Tella seven-day Pro trial**             |
| Best capture escalation for app/system audio or PiP  | **OBS + FFmpeg**                          |
| Best local automatic screen-demo polish if paying    | **Screen Studio**                         |

## Non-Negotiable Delivery Constraints

1. Keep the complete cut below three minutes and target about **2:45–2:50** so the required content survives timing trims. [Production-risk assignment](assignments/13-production-workflow-risk.md)
2. Include intelligible audio covering the project, Codex, and GPT-5.6; captions improve accessibility but do not replace the required audio. [Official Rules](https://openai.devpost.com/rules) [Rules assignment](assignments/02-build-week-video-rules.md)
3. Demonstrate real, reproducible project behaviour that matches the repository, README, description, and judging-access materials. [Official Rules](https://openai.devpost.com/rules)
4. Upload to YouTube as **Public**, then verify the published link outside the signed-in channel context. The researched rules explicitly require public visibility; unlisted compliance was not established. [YouTube visibility documentation](https://support.google.com/youtube/answer/157177) [Rules assignment](assignments/02-build-week-video-rules.md)
5. Use original or specifically cleared audio, visuals, fonts, recordings, and trademarks. Preserve asset provenance and any required attribution. [Official Rules](https://openai.devpost.com/rules)
6. Publish and verify before the **July 21, 2026, 5:00 PM Pacific Time** deadline; creative replacement after submission closes should not be assumed. [Official Rules](https://openai.devpost.com/rules) [Rules assignment](assignments/02-build-week-video-rules.md)

The rules do not prescribe 1080p, captions, camera presence, a thumbnail, an editing application, or a visual style. Those are production decisions; 1080p is the requested quality target. [Rules assignment](assignments/02-build-week-video-rules.md)

## Current Codex Capability Boundary

Codex can create and operate a reproducible media-production workflow around installed commands. On this Mac, `ffmpeg`, `ffprobe`, `ffplay`, and macOS `screencapture` are present. The FFmpeg build exposes H.264 software and VideoToolbox encoding, AAC, concat, fades, crossfades, overlay, scale, frame-rate conversion, audio mixing, ducking, denoising, and loudness normalisation. This is enough for capture preparation, stitching, title-card overlays, occasional transitions, audio finishing, 1080p export, and metadata inspection. [Local capability assignment](assignments/01-local-codex-capabilities.md) [FFmpeg/macOS deep dive](assignments/07-ffmpeg-obs-macos-deep-dive.md)

Codex Browser and Computer Use can navigate and operate approved web or Mac applications, including clicking, typing, dragging, inspecting rendered state, and taking screenshots. They are **UI-automation capabilities**, not native timeline editors, video recorders, or renderers. No reviewed editor documents a reliable unattended browser or desktop automation API for a complete final edit. [OpenAI Browser documentation](https://learn.chatgpt.com/docs/browser) [OpenAI Computer Use documentation](https://learn.chatgpt.com/docs/computer-use) [Local capability assignment](assignments/01-local-codex-capabilities.md)

The installed ImageGen capability creates and edits raster images. It can supply a thumbnail, title-card background, illustration, or static overlay, but it does not generate, capture, or edit video. OpenAI has a separate Videos API for Sora generation and targeted video editing; it is not an installed Codex video skill in this environment, requires separate API access, and the researched Sora 2 documentation announced shutdown on September 24, 2026. It is unnecessary and too dependency-sensitive for this delivery path. [OpenAI Image generation](https://learn.chatgpt.com/docs/image-generation) [OpenAI Video generation with Sora](https://developers.openai.com/api/docs/guides/video-generation) [Local capability assignment](assignments/01-local-codex-capabilities.md)

## Ranked Shortlist

Pricing and entitlements are snapshots inspected on July 19, 2026. Run a 10–15 second export using the exact account, features, and assets before committing the full edit.

| Rank   | Option                                                    | Research classification                                 | Current cost and clean-export position                                                                                                                                           | Why it ranks here                                                                                                                           | Main limits and setup                                                                                                                             |
| ------ | --------------------------------------------------------- | ------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| **1**  | **macOS Screenshot + installed FFmpeg + YouTube SRT/VTT** | Deep-dive; final recommendation                         | No additional tool cost; local render; no vendor watermark mechanism identified                                                                                                  | Highest Codex automation, immediate availability, local privacy, deterministic revisions, H.264/AAC 1080p output, and `ffprobe` QA          | User must record; native system audio is unverified; installed FFmpeg lacks `drawtext` and `subtitles`; visual design should stay restrained      |
| **2**  | **Screenshot + Remotion**                                 | Deep-dive; higher-polish conditional                    | Free for eligible individuals, organizations with up to three employees, and nonprofits; Company Creator licence listed at US$25/month per seat when required                    | Best code-defined motion graphics, titles, transitions, captions, audio, and one-composition rendering; highly legible to Codex             | Fresh project; macOS 15+ and Node; several focused hours; licence/telemetry/version check; no screen capture; no local trial render was performed |
| **3**  | **Clipchamp personal Free**                               | Deep-dive; best single tool                             | Free personal Microsoft account; unlimited 1080p MP4/30 fps exports; no watermark; premium stock and 4K are gated                                                                | One browser tool covers capture, trim/stitch, titles, transitions, audio, subtitles, voiceovers, and export                                 | Chrome/Edge, sign-in, network, hosted workflow, permission prompts, and live premium-asset/export preflight; UI automation remains supervised     |
| **4**  | **Screenshot + iMovie**                                   | Deep-dive; best simple local GUI                        | No editor subscription or export paywall documented; iMovie was not found installed locally                                                                                      | Lowest-learning visible timeline with titles, transitions, voiceover, audio, and documented 1080p60 export                                  | May require Mac App Store acquisition/Apple ID; editing is manual and less reproducible; caption authoring was not established                    |
| **5**  | **Tella Pro trial**                                       | Focused deep dive; fastest clean automatic-polish trial | Seven-day no-card Pro trial; trial video has no watermark, while its hosted share page has a small badge; Pro listed at US$13/month annually and Premium at US$19/month annually | Automatic click zoom, separate screen/camera, system audio, transcript editing, captions, layouts, cleanup, and 4K download in one route    | Hosted/cloud processing, account, macOS 13+, user-content terms, possible Gemini processing for AI features, and trial/account preflight          |
| **6**  | **OBS + FFmpeg**                                          | Deep-dive; capture escalation                           | Free/open source; no watermark identified; OBS is not installed locally                                                                                                          | Best free route for application/system audio on macOS 13+, webcam/PiP, scenes, overlays, and WebSocket control; FFmpeg finishes predictably | Installation, permissions, source/audio configuration, MKV/remux understanding, and a 15–30 minute estimated setup before retakes                 |
| **7**  | **Editly + FFmpeg**                                       | Deep-dive finalist; minimal code-defined middle ground  | Free, MIT-licensed, local by architecture; no watermark                                                                                                                          | Small JSON5 timeline for footage, titles, transitions, subtitles, music, ducking, and normalisation; simpler than React or raw filtergraphs | Not installed; latest listed release is `v0.15.0-rc.1` from January 2025; current Node/FFmpeg render must be proven                               |
| **8**  | **Screen Studio**                                         | Deep-dive; paid local auto-polish                       | US$29/month or US$9/month billed yearly; subscriptions described as prepaid/non-refundable; clean evaluation export was not established                                          | Strongest local automatic cursor smoothing, zoom, framing, local subtitles, audio cleanup, and up to 4K/60                                  | Payment before dependable export, macOS 13.1+, no stable public capture/edit API, title-card/transition authoring not established                 |
| **9**  | **Canva Video**                                           | Deep-dive; graphic-treatment alternative                | Free/original assets can avoid watermark; Pro assets watermark Free users; exact current Free resolution/storage needs a live check                                              | Fastest polished title cards, animated graphics, and presentation-style intro/outro around imported footage                                 | Browser recording is unavailable; desktop recorder required; cloud upload/account; asset licence and export gates                                 |
| **10** | **Shotcut**                                               | Local GUI fallback                                      | Free, GPLv3, local, no account/activation/analytics; 4K support                                                                                                                  | Best open-source GUI substitute for iMovie when local control matters more than speed                                                       | More learning and manual editing; estimated 60–120 minutes for a first short cut; sample export still required                                    |

The top rank favors the user’s explicit preference for substantial Codex hands-on work. Remotion ranks above Clipchamp in code-defined potential but below the primary path in deadline fit because no Remotion project exists, while Screenshot and FFmpeg are already present. [Scriptable ecosystem assignment](assignments/03-scriptable-video-ecosystem.md) [Remotion deep dive](assignments/06-remotion-deep-dive.md) [Production-risk assignment](assignments/13-production-workflow-risk.md)

## Best Single-Tool and Small-Toolchain Options

### Best Single Tool: Clipchamp Free

Clipchamp is the strongest source-backed single-tool answer for a free final cut. A personal Free account supports unlimited 1080p exports without a watermark and includes the complete conventional workflow needed here: screen/camera/voice recording, timeline edits, motion titles, transitions/effects, audio tools, automatic captions with SRT download, and MP4 export. Use original media or clearly included Free assets; premium material blocks export. [Clipchamp pricing](https://clipchamp.com/en/pricing/) [Clipchamp export format](https://support.microsoft.com/en-us/clipchamp/what-format-is-my-video-exported-in) [Clipchamp watermark policy](https://support.microsoft.com/en-US/Clipchamp/does-clipchamp-add-a-watermark-to-videos)

Choose it when the goal is minimum tool count rather than maximum automation or local control. Codex can prepare the script, scene order, title copy, captions, artwork, and editing checklist, then assist visibly in the editor. The user must handle account authentication, capture permissions, any sensitive-media upload decision, unexpected plan prompts, final export review, and publication. [Browser-editor deep dive](assignments/09-browser-editors-deep-dive.md)

### Best Small Toolchain: Screenshot + FFmpeg + YouTube Captions

This toolchain separates each responsibility cleanly:

- **Screenshot/`screencapture`:** local, user-supervised window or selected-region recording with microphone and optional click indicators.
- **Prepared PNG/video cards:** original intro/outro and lower-third assets created by Codex from approved copy, fonts, colours, and artwork.
- **FFmpeg:** normalise clips, stitch, fade/crossfade, overlay cards, clean/mix/normalise audio, export H.264/AAC MP4, and verify with `ffprobe`.
- **SRT/VTT + YouTube Studio:** corrected viewer-controlled captions, avoiding the unavailable local `subtitles` filter.

The result is cheap, auditable, local until YouTube upload, and easy for Codex to revise through source assets and commands. [FFmpeg/macOS deep dive](assignments/07-ffmpeg-obs-macos-deep-dive.md) [Audio and captions assignment](assignments/12-audio-captions-accessibility.md)

### Best Higher-Polish Code Toolchain: Screenshot + Remotion

Remotion should replace the FFmpeg composition layer only when motion design materially improves the explanation. It can own the complete 1920×1080 composition—video trims, animated titles, callouts, transitions, narration/music, captions, and H.264/AAC output—while Codex authors the TypeScript/React timeline and the user reviews it in Studio. Use a short proof composition before committing. [Remotion deep dive](assignments/06-remotion-deep-dive.md)

## Decision Matrix: Production Fit

Ratings are a synthesis of the supplied evidence for this specific one-off public demo, not measured benchmarks.

| Option                | Codex automation                    | Speed                                    | Minimalism                                     | Ease for first cut                   | Cost                            | Watermark/export                                         | macOS                             | Screen capture                                                     |
| --------------------- | ----------------------------------- | ---------------------------------------- | ---------------------------------------------- | ------------------------------------ | ------------------------------- | -------------------------------------------------------- | --------------------------------- | ------------------------------------------------------------------ |
| Screenshot + FFmpeg   | **High** after proof render         | High once scripted                       | **High** tool count, medium command complexity | Medium                               | **Free; installed**             | Local clean 1080p path; test export required             | Built in + installed              | Native full/region/window; mic documented; system audio unverified |
| Screenshot + Remotion | **Very high** for composition       | Medium/low initially; high for revisions | Medium/low for one video                       | Medium for React user; low otherwise | Free only if licence-eligible   | Local H.264/AAC; no export watermark                     | macOS 15+                         | Separate native capture                                            |
| Clipchamp Free        | Medium, supervised UI               | **High**                                 | **High** single-tool                           | **High**                             | **Free**                        | Official 1080p/no watermark; premium assets block export | Chrome/Edge on Mac                | Built-in screen/camera/voice capture                               |
| Screenshot + iMovie   | Low/medium; instructions and assets | High                                     | High if installed                              | **High**                             | Free/no subscription documented | 1080p60 documented; sample export advised                | Native Mac app; absent locally    | Separate native capture                                            |
| Tella trial           | Medium, supervised UI               | **Very high** for narrated demos         | High                                           | High                                 | Seven-day clean-video trial     | No video watermark in trial; hosted page badge           | macOS 13+                         | Screen/camera/mic/system audio                                     |
| OBS + FFmpeg          | High after OBS setup                | Medium                                   | Low/medium                                     | Medium/low initially                 | Free                            | Local clean route; test required                         | macOS 12+, audio strengths on 13+ | Strongest configurable display/window/app/audio/PiP capture        |
| Editly + FFmpeg       | **High**                            | Potentially high after proof             | High declarative surface                       | Medium                               | Free                            | Local output; maintenance/render check                   | Node + FFmpeg on Mac              | Separate capture                                                   |
| Screen Studio         | Medium, supervised UI               | **Very high** for automatic polish       | High                                           | High                                 | US$29 monthly minimum evidenced | Paid active export; evaluation entitlement unsupported   | macOS 13.1+                       | Strong local screen/camera/mic/system audio                        |

## Decision Matrix: Editorial Capability and Risk

| Option                | Animation/title                                                                       | Transitions                                                       | Audio                                               | Captions                                                       | Expected quality                                      | Licensing                                                                              | Privacy                                                      | Availability risk                                               |
| --------------------- | ------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | --------------------------------------------------- | -------------------------------------------------------------- | ----------------------------------------------------- | -------------------------------------------------------------------------------------- | ------------------------------------------------------------ | --------------------------------------------------------------- |
| Screenshot + FFmpeg   | Prepared image/video cards; simple fades/movement                                     | Cuts, fades, `xfade`                                              | Mix, duck, denoise, `loudnorm`; voice-only simplest | Sidecar SRT/VTT; no local burn-in                              | Clean, restrained, technically predictable            | Installed Homebrew FFmpeg reports GPL-3.0-or-later; final media rights remain separate | Local until YouTube upload                                   | Low; permissions and exact capture/audio still need proof       |
| Screenshot + Remotion | **Strong programmable CSS/SVG/canvas titles and callouts**                            | Strong built-in transition package                                | Imported/mixed audio and frame-driven ducking       | SRT parse/display/export; local or cloud transcription options | Highest code-defined polish                           | Source-available; free eligibility or Company Licence decision                         | Local render possible; complete telemetry posture not proven | Medium: fresh setup, macOS/licence/version/render checks        |
| Clipchamp Free        | Templates and motion titles                                                           | Conventional effects/fades                                        | Recording, stock, voiceover, suppression            | Automatic captions and SRT download                            | Strong conventional demo quality                      | Free/premium asset entitlement must be respected                                       | Hosted/account workflow                                      | Medium: account, UI, network, premium-asset gate                |
| Screenshot + iMovie   | Built-in animated titles                                                              | Native transitions                                                | Voiceover, music, fades, enhancement                | Caption authoring unsupported in reviewed source               | Strong enough for restrained public demo              | Apple proprietary application; media rights separate                                   | Local editor                                                 | Medium: app absent and acquisition may need Apple ID            |
| Tella trial           | Layouts, backgrounds, logo, zoom treatment                                            | Recorder/editor treatment; full title-card depth less established | System/mic, Studio Voice, cleanup, music            | Captions/transcript editing                                    | Strong automatic screen-demo polish                   | Business use of user content allowed; third-party rights remain user’s responsibility  | Hosted US storage; broad service licence; AI may use Gemini  | Medium: account/trial/cloud terms and final export check        |
| OBS + FFmpeg          | OBS text/image/media scenes plus FFmpeg cards                                         | Live scene changes plus FFmpeg finishing                          | Best system/app audio path; meters and routing      | Separate SRT workflow                                          | High when configured; risk of setup overkill          | OBS GPL; FFmpeg build licence as above                                                 | Local; protect OBS WebSocket                                 | Medium/high because OBS is absent and configuration is unproven |
| Editly + FFmpeg       | Title/subtitle/layer types                                                            | Standard named transitions                                        | Music, timed tracks, mix volumes, normalisation     | Subtitle layers documented                                     | Good conventional polish; lower custom-motion ceiling | MIT plus FFmpeg dependency                                                             | Local by architecture; remote requests default off           | Medium/high due release age and no current render test          |
| Screen Studio         | Strong automatic cursor/zoom/background treatment; title-card support not established | Full NLE transition depth not established                         | Mic/system audio, normalisation, noise removal      | Local transcription/subtitles                                  | Best source-described local automatic screen polish   | Subscription terms; positive commercial-output grant not located                       | Local unless share link; blur retains original in project    | Medium: payment, non-refund, no established clean trial         |

## Concrete Workflow for This Video

### 1. Lock a 2:45–2:50 Story

Use this production budget, which preserves every required point before the judge-viewing ceiling:

| Time        | Content                                                                                           |
| ----------- | ------------------------------------------------------------------------------------------------- |
| `0:00–0:10` | Project name, problem, and one-sentence value proposition                                         |
| `0:10–0:35` | Orient the viewer: who it is for and what they will see                                           |
| `0:35–1:50` | Real working demonstration using generic/public data                                              |
| `1:50–2:25` | Explicitly explain how Codex and GPT-5.6 were used, with matching evidence on screen where useful |
| `2:25–2:50` | Outcome, next step, and short closing card                                                        |

This is a production structure rather than a competition-prescribed format. Trim decoration before trimming the working demonstration or the named Codex/GPT-5.6 explanation. [Production-risk assignment](assignments/13-production-workflow-risk.md)

### 2. Prepare Deterministic Inputs

Have Codex produce:

- a fact-checked narration of approximately 90–150 seconds;
- a shot list tied to real project states and exact user actions;
- two original 1920×1080 title cards and any transparent callout overlays;
- one visual specification: font, sizes, colours, margins, card duration, and transition style;
- an editable SRT draft or time-coded transcript;
- an asset-rights ledger containing origin, licence, attribution, and intended use;
- FFmpeg normalisation, assembly, export, and `ffprobe` commands.

Use an original voice and omit music unless it clearly improves comprehension. If music is needed, the narrowest cleared source is YouTube Audio Library; choose an `Attribution not required` track or preserve the required Creative Commons attribution exactly. [YouTube Audio Library](https://support.google.com/youtube/answer/3376882) [Audio and captions assignment](assignments/12-audio-captions-accessibility.md)

### 3. Clean the Capture State

Close notifications and unrelated applications; use generic/public sample data; remove credentials, personal information, customer data, internal messages, and unrelated third-party marks. Set the demo window or selected region to a stable 16:9 layout, increase UI scale if needed for 1080p readability, test the microphone, and grant Screen & System Audio Recording only to the selected capture application. [Apple screen-recording guide](https://support.apple.com/en-us/102618) [Apple privacy guide](https://support.apple.com/en-ie/guide/mac-help/mchld6aa7d23/mac)

### 4. Run a Proof Gate Before Full Production

Record 10–15 seconds using the exact capture method. Assemble it with the real title-card asset class and audio path, then export at 1920×1080. Confirm:

- capture and microphone permissions work;
- narration is intelligible and unclipped;
- the crop and cursor are readable;
- the title card and transition render correctly;
- the output has no watermark or paid-asset block;
- `ffprobe` reports the expected duration, dimensions, H.264 video, AAC audio, frame rate, and 48 kHz sample rate.

This proof gate is the highest-leverage protection against tool, account, asset, permission, and codec surprises. [FFmpeg/macOS deep dive](assignments/07-ffmpeg-obs-macos-deep-dive.md) [Production-risk assignment](assignments/13-production-workflow-risk.md)

### 5. Capture Modularly

Record one establishing clip, two or three feature clips, and a separate Codex/GPT-5.6 explanation segment. Keep retakes as separate files. Native capture is estimated at 3–8 minutes per prepared take; OBS requires an estimated 15–30 minutes of setup before retakes. These are planning estimates, not measured timings. [FFmpeg/macOS deep dive](assignments/07-ffmpeg-obs-macos-deep-dive.md)

Record narration separately when it reduces performance pressure or makes audio replacement easier. Keep the original recording, then apply only conservative FFmpeg denoising and loudness normalisation; a human must A/B the result for metallic artefacts, pumping, clipped consonants, and speech masking. [Audio and captions assignment](assignments/12-audio-captions-accessibility.md)

### 6. Assemble the Restrained Cut

Normalize every clip to matching 1920×1080, 30 fps, `yuv420p`, and compatible audio before concatenation. Use:

- one opening card of approximately 3–5 seconds;
- direct cuts for most actions;
- one fade/crossfade treatment at chapter boundaries;
- prepared image/video overlays for titles and lower thirds;
- voice-only audio unless a cleared music bed genuinely helps;
- a 3–5 second closing card.

The installed FFmpeg cannot render `drawtext` or `subtitles`, so prepare text as image/video assets and keep captions sidecar. A representative final encoding pattern from the supplied deep dive is:

```sh
ffmpeg -i assembled.mov \
  -c:v libx264 -preset medium -b:v 8M -maxrate 8M -bufsize 16M \
  -pix_fmt yuv420p -r 30 \
  -c:a aac -b:a 192k -ar 48000 \
  -movflags +faststart final-1080p.mp4
```

YouTube recommends MP4, H.264, AAC-LC at 48 kHz, progressive video, matching the recorded frame rate, and about 8 Mbps for 1080p SDR at 24–30 fps or 12 Mbps at 48–60 fps. [YouTube encoding settings](https://support.google.com/youtube/answer/1722171) [FFmpeg/macOS deep dive](assignments/07-ffmpeg-obs-macos-deep-dive.md)

### 7. Correct Captions and Inspect the Master

If setup time permits, install `whisper.cpp`, convert the final narration to mono 16 kHz PCM WAV, generate a local SRT/VTT draft, and manually correct technical terms, names, punctuation, timing, line breaks, and meaningful sound cues. `whisper.cpp` is not currently installed, so it is an optional setup rather than a final-hour dependency. The fallback is YouTube automatic captions or transcript auto-sync followed by manual correction. [whisper.cpp](https://github.com/ggml-org/whisper.cpp) [Audio and captions assignment](assignments/12-audio-captions-accessibility.md)

Watch the exported master end to end with headphones and again with captions. Inspect the first frame, each transition, the working demo, the Codex/GPT-5.6 explanation, and the final frame. A successful render or passing metadata probe does not prove visual, narrative, audio, or rights quality.

### 8. Publish and Verify

Upload the checked local master to YouTube early enough for 1080p processing. Upload the corrected caption file, set the original language, set visibility to **Public**, and play the published link in a signed-out or separate-profile context. Confirm the 1080p rendition, captions, audio, duration, description/attribution, and exact URL before placing it in Devpost. [YouTube captions](https://support.google.com/youtube/answer/2734796) [YouTube visibility documentation](https://support.google.com/youtube/answer/157177) [Official Rules](https://openai.devpost.com/rules)

## What Codex Can Do and What Requires the User

| Codex can own or substantially execute                                                                | User-only or user-approved work                                                                                                       |
| ----------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Draft and revise the timed narrative, including explicit Codex/GPT-5.6 explanation                    | Confirm every claim matches the actual project and hackathon-period work                                                              |
| Produce the shot list, scene timing, title/outro copy, and editing manifest                           | Perform or supervise the real product interaction and narration performance                                                           |
| Generate original still assets through ImageGen, or code title cards in HTML/CSS/SVG                  | Approve visual identity, likeness/voice use, fonts, logos, and all rights-sensitive material                                          |
| Author FFmpeg commands, Remotion/Editly code/specs, audio filters, and `ffprobe` checks               | Grant macOS screen/microphone/camera permissions and review capture scope                                                             |
| Convert audio, run local cleanup, build `whisper.cpp`, and generate a caption draft                   | Listen to audio quality and correct transcript meaning, technical terms, timing, and accessibility cues                               |
| Assist visible Browser/Computer Use editing steps in Clipchamp, Tella, iMovie, or other approved apps | Authenticate accounts, approve uploads, resolve payment/trial prompts, and decide whether sensitive footage may enter a hosted editor |
| Inspect technical metadata and produce a release checklist                                            | Watch the final export and published YouTube rendition end to end                                                                     |
| Prepare upload copy, description, attribution text, and a public-link verification checklist          | Upload to the user’s channel, set Public visibility, accept platform prompts, and submit the final URL                                |

Codex does not currently provide an installed native video generator/editor. ImageGen supplies still assets; Browser/Computer Use supplies supervised UI interaction; FFmpeg and code-defined tools supply the actual deterministic media work. [Local capability assignment](assignments/01-local-codex-capabilities.md)

## Time, Difficulty, Installation, and Accounts

| Path                  | Planning estimate                                                                                   | Difficulty                                      | Installation/account requirements                                                                                     |
| --------------------- | --------------------------------------------------------------------------------------------------- | ----------------------------------------------- | --------------------------------------------------------------------------------------------------------------------- |
| Screenshot + FFmpeg   | Setup 20–45 minutes; user hands-on 60–120 minutes after prepared inputs                             | Medium initially, then low/medium for revisions | No new editor install or account; macOS recording permission; YouTube account only for publication                    |
| Screenshot + Remotion | Several focused hours for first composition; revisions faster                                       | Medium/high                                     | macOS 15+, Node, package install, local Chromium/FFmpeg render; confirm Remotion Free Licence eligibility             |
| Clipchamp Free        | 45–75 minutes after prepared footage, script, and assets; broader workflow estimate 60–130 minutes  | Low/medium                                      | Personal Microsoft account, Chrome or Edge, network, browser capture/upload permissions, live Free export check       |
| Screenshot + iMovie   | 30–60 minutes after prepared assets; 60–90 if rerecording; broader workflow estimate 75–150 minutes | Low                                             | iMovie installation/acquisition if absent, possible Apple ID, macOS recording permission                              |
| Tella trial           | Unquantified; source evidence supports lower editing friction, but no numeric hands-on benchmark    | Low/medium                                      | Tella account, seven-day Pro trial, macOS 13+, capture permissions, hosted/cloud acceptance                           |
| OBS + FFmpeg          | OBS setup estimated 15–30 minutes before retakes; full cut comparable to scriptable path            | Medium                                          | OBS install, macOS capture/audio permissions, scenes/sources, audio test, output profile, optional WebSocket password |
| Editly + FFmpeg       | Potentially low once proven; first proof render is mandatory                                        | Medium                                          | Node package installation; FFmpeg/ffprobe already available; current dependency compatibility check                   |
| Screen Studio         | Hands-on estimate 45–100 minutes plus purchase/activation                                           | Low/medium                                      | Paid active plan, macOS 13.1+, recording permissions; no clean trial entitlement established                          |

The ranges are planning estimates preserved from the assignments, not measured production benchmarks. No candidate received an end-to-end render, export, upload, or signed-in entitlement test. [Browser-editor deep dive](assignments/09-browser-editors-deep-dive.md) [Open-source GUI assignment](assignments/08-open-source-gui-editors.md) [Production-risk assignment](assignments/13-production-workflow-risk.md)

## Limitations and Risks

| Risk                                                | Current evidence                                                                                        | Control or fallback                                                                                               |
| --------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| Required explanation is missing or too late         | Judges need not watch after three minutes                                                               | Put the audible project and Codex/GPT-5.6 explanation before 2:25; trim decoration first                          |
| Native capture lacks needed system/app audio        | Microphone capture is documented; native system-audio selection remains unverified                      | Run a short test; use OBS on macOS 13+ if application/system audio is essential                                   |
| Installed FFmpeg cannot render text or subtitles    | Local probes return `Unknown filter` for `drawtext` and `subtitles`                                     | Use prepared PNG/video title assets and upload SRT/VTT to YouTube                                                 |
| Watermark or premium-asset surprise                 | Hosted/editor entitlements vary by plan and selected asset                                              | Export 10–15 seconds with the exact account and asset classes before full editing                                 |
| Sensitive content enters footage or a hosted editor | Capture and cloud products can expose credentials, personal data, customer data, or unreleased material | Use generic/public data, record modularly, inspect every frame, and prefer the local route                        |
| Audio is missing or unintelligible                  | The rules require audio                                                                                 | Record a mic test, monitor with headphones, preserve separate narration, and review encoded plus YouTube playback |
| Captions are inaccurate                             | Automatic transcription can misrecognise accents, noise, and technical terms                            | Upload a corrected SRT/VTT and review it in the published player                                                  |
| Tool/account/network failure near deadline          | Clipchamp/Tella/other hosted paths depend on account, network, and UI state                             | Keep Screenshot + FFmpeg as the installed local fallback and export a local master before upload                  |
| Over-polishing consumes the deadline                | Motion graphics and unfamiliar editors increase iteration cost                                          | Freeze one title treatment and at most two transition styles after the proof gate                                 |
| Rights or attribution are unclear                   | Competition rules prohibit unlicensed protected material                                                | Use original assets or a retained provenance ledger; omit music if clearance is uncertain                         |
| YouTube link is not truly public                    | Private/unlisted are distinct visibility states                                                         | Set Public and test the published link outside the signed-in account                                              |

## Conditional Alternatives

- **Choose Remotion** when animated callouts, reusable branded components, or a code-defined narrative are central and macOS 15+, licence eligibility, package setup, and several focused hours are acceptable. [Remotion deep dive](assignments/06-remotion-deep-dive.md)
- **Choose Clipchamp** when one free browser tool is more valuable than local privacy and deterministic rerendering. It is the best single-tool option after a live Free-account proof export. [Browser-editor deep dive](assignments/09-browser-editors-deep-dive.md)
- **Choose Screenshot + iMovie** when the user wants the simplest local visual timeline and can acquire iMovie quickly. Use a separate SRT/VTT caption path. [Open-source GUI assignment](assignments/08-open-source-gui-editors.md)
- **Choose Tella’s trial** when automatic zoom, transcript editing, system audio, and rapid hosted finishing could save more time than a local edit, and the footage is safe for its cloud terms. [Assisted-recorders assignment](assignments/10-assisted-screen-recorders.md)
- **Choose OBS + FFmpeg** when per-app/system audio, webcam/PiP, live scene switching, or repeatable configured capture is a real requirement. Otherwise native capture is faster. [FFmpeg/macOS deep dive](assignments/07-ffmpeg-obs-macos-deep-dive.md)
- **Choose Screen Studio** when the local auto-zoom/cursor aesthetic is worth at least the evidenced US$29 monthly purchase and the non-refundable subscription terms. [Assisted-recorders assignment](assignments/10-assisted-screen-recorders.md)
- **Choose Canva** when graphic title-card polish matters more than browser-only capture. Capture elsewhere, use only original/Free/licensed assets, and prove the current Free export. [Browser-editor deep dive](assignments/09-browser-editors-deep-dive.md)
- **Choose Shotcut** when an open-source, local, account-free GUI matters more than minimal learning time. [Open-source GUI assignment](assignments/08-open-source-gui-editors.md)
- **Add `whisper.cpp`** when private local caption generation is valuable and one-time build/model setup will not endanger the deadline. [Audio and captions assignment](assignments/12-audio-captions-accessibility.md)
- **Use Editly** only after a real 1920×1080 proof render confirms compatibility with current Node and FFmpeg; keep direct FFmpeg as the fallback. [Code-animation alternatives](assignments/11-code-animation-alternatives.md)

## Lower-Fit and Rejected Options

### Lower-Fit but Viable

- **CapCut Online/Desktop:** broad feature set and potential clean manual exports, but entitlement varies by account, asset, browser/device, and region; its cloud/content terms add unnecessary risk for this generic local-capable job. [Browser-editor deep dive](assignments/09-browser-editors-deep-dive.md)
- **Adobe Express:** capable Free browser editor with templates, voice/video recording, transitions, caption tools, and MP4 download, but the retrieved official sources did not establish Free 1080p export. [Browser-editor deep dive](assignments/09-browser-editors-deep-dive.md)
- **FocuSee:** strong automatic zoom/caption/effect feature set, but the free export is watermarked; the exact privacy and commercial-output position remained incomplete. [Assisted-recorders assignment](assignments/10-assisted-screen-recorders.md)
- **CleanShot X:** useful local capture companion with a URL scheme and optional cloud, but it does not replace the editor for titles, transitions, captions, and multi-clip assembly. [Assisted-recorders assignment](assignments/10-assisted-screen-recorders.md)
- **Motion Canvas:** high-quality TypeScript animation insert tool, but oversized for a basic card and its latest listed GitHub release was December 2024. [Code-animation alternatives](assignments/11-code-animation-alternatives.md)
- **HTML/CSS/SVG + deterministic frames + FFmpeg:** excellent for bespoke code-defined title cards, but it becomes a multi-step frame-capture pipeline; use only if the card needs more than a static asset plus fade. [Code-animation alternatives](assignments/11-code-animation-alternatives.md)
- **Kdenlive:** strong local subtitle/import/export and render-script capability, but a more complex UI and longer first-cut estimate than iMovie or Shotcut. [Open-source GUI assignment](assignments/08-open-source-gui-editors.md)
- **DaVinci Resolve Free:** highest professional finishing ceiling and ample 1080p capability, but its learning curve, registration, and pro-style UI are disproportionate unless the user already knows it. [Open-source GUI assignment](assignments/08-open-source-gui-editors.md)
- **OpenShot:** rich titles and transitions, but current downloadable binaries are Intel-only according to the reviewed requirements, weakening it on Apple Silicon. [Open-source GUI assignment](assignments/08-open-source-gui-editors.md)
- **VEED, Kapwing, invideo AI, FlexClip, Renderforest, Powtoon, and Biteable:** capable paid or trial editors, but their documented free paths introduce watermarks, 720p limits, changing credits, or subscription dependence. [Browser/template ecosystem assignment](assignments/04-browser-template-ecosystem.md)
- **MoviePy, MLT/`melt`, auto-editor, LosslessCut, VHS, Manim, and GStreamer Editing Services:** useful specialist or fallback tools, but each adds a narrower role, slower path, or unnecessary complexity compared with FFmpeg, Editly, or Remotion for this cut. [Scriptable ecosystem assignment](assignments/03-scriptable-video-ecosystem.md)

### Rejected for the Primary Workflow

- **Animoto Free:** permanent “Made with Animoto” watermark. [Browser/template ecosystem assignment](assignments/04-browser-template-ecosystem.md)
- **ScreenFlow trial:** watermarked export; paid entry starts at US$169 in the reviewed source. [Desktop ecosystem assignment](assignments/05-desktop-recording-ecosystem.md)
- **Camtasia trial:** watermarked output. [Desktop ecosystem assignment](assignments/05-desktop-recording-ecosystem.md)
- **QuickTime alone:** captures and trims but lacks the titles, transitions, captions, audio finishing, and assembly needed for the requested polished cut. [Desktop ecosystem assignment](assignments/05-desktop-recording-ecosystem.md)
- **Blender:** scriptable/headless in principle, but an unnecessarily large 3D/DCC environment for this short screen-demo edit. [Scriptable ecosystem assignment](assignments/03-scriptable-video-ecosystem.md)
- **Standalone `gst-launch-1.0`:** official documentation positions it as a basic/debugging pipeline tool rather than a production editing application. [Scriptable ecosystem assignment](assignments/03-scriptable-video-ecosystem.md)

## Candidate Discovery Classification

The assignment labels below describe research depth and discovery fit; they are not a claim that every deep-dive candidate should be used.

| Classification                               | Candidates surfaced in the assignments                                                                                                                                                                                    |
| -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Deep-dive / targeted finalist**            | FFmpeg, Remotion, OBS, Editly, Motion Canvas, `whisper.cpp`, Clipchamp, Canva, Adobe Express, CapCut Online, Screen Studio, Tella, Screenshot+iMovie, CapCut Desktop, CleanShot+iMovie                                    |
| **Focused conditional deep dive**            | FocuSee, HTML/CSS/SVG frame capture, Shotcut/local GUI alternatives                                                                                                                                                       |
| **Skim-only / lower fit**                    | MoviePy, MLT/`melt`, auto-editor, Manim, DaVinci Resolve, LosslessCut, VHS, GStreamer Editing Services, VEED, Kapwing, invideo AI, FlexClip, Renderforest, Powtoon, Biteable, Descript, OpenShot, Keynote, Kdenlive, Loom |
| **Rejected in discovery or final synthesis** | Blender, standalone `gst-launch-1.0`, Animoto Free, ScreenFlow trial, Camtasia trial, QuickTime alone                                                                                                                     |

The underlying discovery evidence and classifications are preserved in the [scriptable ecosystem](assignments/03-scriptable-video-ecosystem.md), [browser/template ecosystem](assignments/04-browser-template-ecosystem.md), [desktop ecosystem](assignments/05-desktop-recording-ecosystem.md), [GUI editor](assignments/08-open-source-gui-editors.md), [assisted recorder](assignments/10-assisted-screen-recorders.md), and [code-animation](assignments/11-code-animation-alternatives.md) assignments.

## Unsupported Claims and True Validation Gates

No supplied evidence establishes the following as completed or guaranteed:

1. **No end-to-end production trial occurred.** No actual capture, full render, caption generation, editor export, YouTube upload, or public-link verification was performed.
2. **macOS permissions and device routing remain unproven.** The tools exist, but screen-recording permission, microphone quality, native system-audio capture, encoder behaviour, and exact capture frame rate need a proof clip.
3. **The installed FFmpeg cannot currently generate text or burn captions.** `drawtext` and `subtitles` are absent; use prepared title assets and sidecar captions unless the toolchain is intentionally changed.
4. **No GUI editor has proven unattended Codex automation.** Browser and Computer Use can assist visible operation, but a robust autonomous end-to-end edit/export path was not established.
5. **Hosted editor entitlements are account-specific.** Clipchamp, Canva, Adobe Express, CapCut, Tella, and other vendor claims still require a live check against the exact account, region, browser, assets, and export dialog.
6. **Remotion, Editly, Motion Canvas, and `whisper.cpp` are not installed or trial-rendered.** Compatibility, render time, font behaviour, caption accuracy, and output quality are validation gates.
7. **Visual and audio quality comparisons are feature-based.** No side-by-side hands-on benchmark established that Screen Studio, Tella, FocuSee, Clipchamp, Remotion, or a local pipeline looks better on this project.
8. **Privacy and licence findings are bounded.** “Local” does not cover package downloads, browser telemetry, optional cloud sharing, stock assets, fonts, YouTube, or user-selected integrations. Several products lacked a complete commercial-output or current privacy conclusion.
9. **No official rule requires 1080p or captions.** These remain the requested quality and accessibility choices, while public YouTube visibility, audio, the project explanation, and Codex/GPT-5.6 explanation are the governing requirements.

These are validation gates rather than blockers to beginning. The primary Screenshot + FFmpeg path can start immediately after a 10–15 second capture/export proof. [Local capability assignment](assignments/01-local-codex-capabilities.md) [Production-risk assignment](assignments/13-production-workflow-risk.md)

## Final Actionable Direction

Begin with the local scriptable path now:

1. Lock the 2:45–2:50 narration and shot list.
2. Create two original 1920×1080 title cards and one restrained visual style.
3. Record a 10–15 second native Screenshot proof with microphone.
4. Have Codex assemble and export that proof with FFmpeg, then verify it with `ffprobe`.
5. If the proof passes, capture modular final clips and finish locally.
6. Generate or auto-sync captions, correct them manually, and upload the sidecar file.
7. Publish to YouTube as Public, verify the 1080p rendition and public link, then submit.

Switch to **Clipchamp Free** only if the user prefers a visible single-tool editor after the local proof feels too command-heavy. Switch to **Remotion** only if the title/callout motion is important enough to justify fresh code-project setup. Escalate capture to **OBS** only when native system/app audio, webcam, or live scene composition is concretely required.

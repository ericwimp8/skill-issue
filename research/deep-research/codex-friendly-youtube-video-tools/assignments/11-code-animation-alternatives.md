# Lightweight Code-Animation and Composition Alternatives

## Assignment

**Goal:** Evaluate Editly, Motion Canvas, and a small HTML/CSS/SVG-to-video path for polished 1080p intro/outro/title-card and transition inserts around externally captured screen footage, where Codex can materially author and maintain the production artefacts.

**Scope:** Primary project documentation, repositories, release pages, licenses, and local non-destructive availability probes inspected 2026-07-19. Assessment covers current availability, macOS setup, 1080p/video/audio/title/transition capability, automation, local-media/privacy posture, license, learning curve, expected output quality, and hands-on effort. Manim and Keynote are included only as boundary comparisons.

**Exclusions:** A full render trial, screen capture, hosted template/video services, browser-template products, a full Remotion or FFmpeg review, and claims about uninspected telemetry, codec quality, or render speed.

## Sources

- [Research map](../research-map.md) and [scriptable-video ecosystem discovery](03-scriptable-video-ecosystem.md), local framing, adjacent FFmpeg/Remotion evidence, and allocated scope, inspected 2026-07-19.
- Local non-destructive probe, 2026-07-19: Node `v22.22.0` and FFmpeg `8.1.1` resolve; `editly`, `motion-canvas`, and `manim` do not resolve on `PATH`; Keynote is absent.
- [Editly repository/API/CLI documentation](https://github.com/mifi/editly), including requirements, installation, JSON/JSON5 edit specification, API, release `v0.15.0-rc.1` (2025-01-19), and MIT license, inspected 2026-07-19.
- [Motion Canvas repository](https://github.com/motion-canvas/motion-canvas), including TypeScript/generator/editor description, MIT license, and latest listed release `v3.17.2` (2024-12-14), inspected 2026-07-19.
- [Motion Canvas quickstart](https://motioncanvas.io/docs/quickstart/), [media documentation](https://motioncanvas.io/docs/media/), [transitions documentation](https://motioncanvas.io/docs/transitions/), and [FFmpeg video exporter](https://motioncanvas.io/docs/rendering/video/), inspected 2026-07-19.
- [MDN Web Animations API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Animations_API), [using the Web Animations API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Animations_API/Using_the_Web_Animations_API), and [MediaRecorder](https://developer.mozilla.org/en-US/docs/Web/API/MediaRecorder), inspected 2026-07-19.
- [Puppeteer `Page.screenshot()`](https://pptr.dev/api/puppeteer.page.screenshot) and [Playwright screenshot documentation](https://playwright.dev/docs/screenshots), inspected 2026-07-19.
- [Manim Community installation](https://docs.manim.community/en/stable/installation.html), [configuration](https://docs.manim.community/en/stable/guides/configuration.html), and [official repository](https://github.com/ManimCommunity/manim), inspected 2026-07-19.
- [Apple Keynote media documentation](https://support.apple.com/en-ca/guide/keynote/tan63d61519a/mac), inspected 2026-07-19.

## Findings

### Fit Ranking for This Small Production Role

| Candidate | Classification | Best role | Why it wins or loses |
| --- | --- | --- | --- |
| Editly | **Deep-dive finalist** | Whole short final cut when footage is already captured | The smallest declarative timeline for clips, titles, transitions, subtitles, and mixed audio. It gives up bespoke motion design and has a maintenance caveat. |
| HTML/CSS/SVG + deterministic capture + FFmpeg | **Deep-dive finalist** | Branded 2–10-second animated cards/overlays | It reuses web design skills/assets exactly and keeps the animation source extremely small; the capture/export pipeline must be deliberately built. |
| Motion Canvas | **Conditional finalist** | Bespoke animated intro/outro or explanatory insert | Strongest previewable code-animation surface, but a larger TypeScript/Vite project and the latest GitHub release is old. |
| Manim | **Skim-only** | Mathematical or diagrammatic animation | High-quality deterministic vector animation, but its Python scene model and dependencies are disproportionate for generic product titles. |
| Keynote | **Skim-only, human-operated** | Fast manual fallback for a Mac user already using it | It can place video/audio, trigger media with slides, and use a soundtrack, but it is a proprietary GUI workflow and this machine has no Keynote app. |

**Evidence:** Editly’s documented spec encompasses a video edit rather than just graphics; Motion Canvas describes itself as TypeScript-programmed animation plus a real-time editor; web standards provide native timed DOM/SVG animation but no authored-video timeline; Manim is explicitly a mathematical-animation framework; and Apple describes Keynote as a presentation app with placed media and slide playback. The local probe establishes that the three code packages are not ready-to-run in this workspace.

**Implication:** For externally captured screen footage, select a small purpose-specific layer: use **Editly** when the need is mainly a linear edit, **HTML/CSS/SVG** for a limited number of brand cards, and **Motion Canvas** only when the card itself needs designed choreography. Keep FFmpeg as the common final encoder/stitcher rather than making any alternative own capture.

### Editly Is the Shortest Declarative Whole-Edit Alternative

Editly is a Node command-line/API editor that accepts direct CLI inputs or an edit specification in JSON/JSON5. Its documented spec contains output path, width, height and fps; clip durations; named transitions; title/subtitle and other layers; source-audio retention; an external music file; additional audio tracks with timings and mix volumes; output-volume control; and optional audio normalization. That covers 1920×1080 delivery, a captured screen clip, a title/opening or closing card, basic transitions, music/narration, and a final local MP4 without inventing React components or large FFmpeg filtergraphs. The repository lists `npm i -g editly`, Node, and `ffmpeg`/`ffprobe` on `PATH` as macOS-capable prerequisites, and identifies its license as MIT. [Editly’s README](https://github.com/mifi/editly) lists latest release `v0.15.0-rc.1` dated 2025-01-19.

**Evidence:** The `outPath`, `width`, `height`, and `fps` fields establish explicit 1080p-capable output; the documented `clips`, `layers`, `transition`, `audioFilePath`, `audioTracks`, and `audioNorm` fields establish the required media/audio/title/transition surface. Editly defaults `allowRemoteRequests` to `false` in the shown spec, which is useful evidence for an asset-local design. Rendering nevertheless depends on the locally installed FFmpeg and packages must be fetched during setup.

**Implication:** Editly wins over Remotion when the creative treatment is conventional—cuts, fades/wipes, static or simply animated text, picture-in-picture, subtitles, and audio levelling—and a single JSON5 file is easier to review than a React project. It wins over raw FFmpeg when human-readable timing/layers matter more than maximum filter-level control. Its release age is a material risk: run one real 1920×1080 render against the current Node/FFmpeg before selection, and keep the direct FFmpeg path as the fallback.

### Motion Canvas Has Better Bespoke Animation, With a Maintenance Cost

Motion Canvas is a local TypeScript library using generators to program animations plus a real-time editor. Its documentation has first-class media support for video assets, scene transitions, and a Video (FFmpeg) exporter. The exporter is installed with `npm install --save @motion-canvas/ffmpeg`; the documentation says FFmpeg is installed with that exporter, can include an audio track when configured, and produces a finished video. The project’s video settings control the render, so a 1920×1080 composition is an ordinary configured target rather than a display-only mock-up. The repository is MIT-licensed, and the TypeScript/Vite setup is macOS-friendly because it uses Node tooling rather than a platform-specific desktop editor.

**Evidence:** [Media docs](https://motioncanvas.io/docs/media/) document a `Video` component importing MP4 footage; [transition docs](https://motioncanvas.io/docs/transitions/) document scene-transition control; and the [exporter docs](https://motioncanvas.io/docs/rendering/video/) document the FFmpeg exporter, automatic FFmpeg installation, and audio inclusion. The project’s GitHub release page lists `v3.17.2` from 2024-12-14 while the documentation pages remain live, so current documentation availability does not by itself prove current dependency compatibility.

**Implication:** Motion Canvas can produce a more intentionally animated title sequence than Editly: animated SVG/vector motifs, text staging, a device/window frame, and a designed transition into the supplied screen recording. It loses to Remotion for a full product-video composition when the team benefits from React’s wider video ecosystem, and it loses to an HTML card when the effect is a short CSS-scale/fade/reveal. Use it for one self-contained MP4 insert, then compose it with the real footage in FFmpeg/Editly; this limits the project’s learning and compatibility exposure.

### HTML/CSS/SVG Is a Lightweight Authoring Surface, Not a Video Exporter

HTML, CSS, inline SVG, and the Web Animations API are a highly Codex-authorable way to make a 16:9 branded card: source is plain text, styles are inspectable in a browser, SVG remains sharp at a 1920×1080 viewport, and animation timing/keyframes are version-controlled alongside the website. MDN documents that Web Animations provides a timing/animation model and `Element.animate()` playback controls; it can animate DOM/SVG presentation without a library. This is sufficient for professional-looking lower thirds, title reveals, logo/shape motion, screen masks, and an outro CTA when the art direction is restrained.

**Evidence:** [MDN’s Web Animations API](https://developer.mozilla.org/en-US/docs/Web/API/Web_Animations_API) documents timelines, keyframe effects, and element-level animation; [its guide](https://developer.mozilla.org/en-US/docs/Web/API/Web_Animations_API/Using_the_Web_Animations_API) documents programmatic playback and timing. Playwright and Puppeteer officially support browser screenshots. MDN documents that `MediaRecorder` records a `MediaStream` and can select container/MIME type and bit rates, but it does not provide a non-realtime, frame-accurate authoring/export contract for this use.

**Implication:** The robust local approach is a two-part pipeline: run the locally served page in a fixed 1920×1080 Chromium viewport; capture numbered deterministic frames through browser automation; encode that frame sequence with the already-installed FFmpeg; then let FFmpeg/Editly place the resulting card and mix its chosen audio. Browser `MediaRecorder` is acceptable for a quick preview, but retain it only if a real output test proves its duration/container/audio on the target browser. Avoid using a literal screen recording of the browser for final title cards because it introduces window/retina/notification variability.

### Automation, Privacy, and Hands-On Trade-offs

| Path | Automation and reproducibility | Local-media/privacy posture | Hands-on cost and likely quality |
| --- | --- | --- | --- |
| Editly | A JSON5 spec and command are straightforward to generate, diff, parameterize, and rerun. | Local by architecture; source spec defaults remote requests off. Dependency installation and any intentionally referenced remote asset are separate exposure points. | Lowest initial cost for a complete simple edit. Output can look polished with disciplined type, spacing, music and transitions; it is less suited to custom kinetic design. |
| Motion Canvas | TypeScript sources plus Vite/editor preview and video exporter are reproducible. | Rendering occurs locally after packages/assets are present. This review did not establish telemetry behavior. | Medium/high initial cost. Highest ceiling here for a deliberate animated insert, but requires scene/generator concepts and a verification render. |
| HTML/CSS/SVG | Plain files plus a viewport/frame-capture/FFmpeg script are versionable, but involve more than one command. | Entirely local after browser/font/assets are available. This review did not inspect browser telemetry settings. | Low for web-literate title cards, medium once frame capture is added. Excellent for clean UI-native motion; weaker for cinematic effects without growing into a framework. |
| Manim | A Python scene and CLI are deterministic once installed. | Local render after dependency/model/font assets are present. | Medium/high and stylistically specialised; excellent for diagrams and equations, poor return for a normal SaaS-style title. |
| Keynote | Manual slide timing and GUI export; no verified Codex-friendly CLI was found. | Local document/media path, subject to the app and media licensing. | Low for an experienced human, but high review/automation cost for Codex. It can look polished for presentation-style cards, with less repeatability. |

**Evidence:** Editly exposes both CLI and JavaScript API; Motion Canvas exposes source-defined projects and exporter configuration; Web animation is standard browser code while capture needs a separate browser/encoder stage; Manim recommends local Python environments/conda/Docker and provides CLI configuration; Apple documents slide-based video/audio playback, automatic start, loops, soundtrack, and media continuity through transitions. No source inspected established a current telemetry policy for Editly, Motion Canvas, Manim, the browsers, or Keynote.

**Implication:** Treat the code sources, a font/asset manifest, output dimensions/FPS, and an `ffprobe` check as the deliverable. The smaller alternative wins when it removes authoring surface without sacrificing the needed effect: Editly for a linear video, HTML/CSS/SVG for web-native cards, Motion Canvas for one animation sequence. Do not adopt Manim or Keynote solely to avoid Remotion/FFmpeg; both add more workflow boundary than this short production needs.

### Decision Guide Against Remotion and FFmpeg

- Choose **Editly over Remotion** for a one-off three-minute edit whose visual vocabulary is clips, text, simple layers, and standard transitions; choose Remotion when reusable React components or precise reactive motion are the core product.
- Choose **Motion Canvas over Remotion** only for a focused explanatory/vector-animation insert where its live editor and generator scenes make that animation easier to tune; keep its output as a clip, not the full NLE.
- Choose **HTML/CSS/SVG over both** when a title can be expressed with existing web type, colors, icons, and keyframes, and an added capture script is still smaller than a new composition framework.
- Choose **direct FFmpeg over every alternative** when the card is static/simple text, the edit already needs only fades/xfades, or dependency freshness is the priority. Use it as the final compatibility fallback for the alternatives’ rendered outputs.

**Evidence:** These are workflow inferences from the documented capabilities and setup surfaces above, plus the adjacent FFmpeg/Remotion source review. They are scoped to titles and transitions around supplied screen recording, rather than claims about general-purpose professional editing.

**Implication:** The next implementation trial should select one 5–10-second title card and one captured-screen clip, then compare (a) Editly JSON5 whole-cut and (b) HTML/CSS/SVG generated card plus FFmpeg stitch. Test Motion Canvas only if the planned title treatment cannot be delivered by those two approaches.

## Notes

- **Maintenance caveat:** Editly’s latest listed release is a release candidate from January 2025; Motion Canvas’s latest listed GitHub release is from December 2024. Neither datum proves failure, but both make a real current Node/macOS render essential before commitment.
- **Unsupported:** No tool was installed, no 1080p render was executed, and no render-time, file-size, encoder, font-substitution, audio-sync, or visual-quality measurement was collected in this assignment.
- **Unsupported:** Current telemetry policies and any hosted dependency/network behavior beyond Editly’s `allowRemoteRequests` default were not verified; do not describe these tools as categorically offline or private.
- **Caveat:** Keynote was not found in the installed-application inventory and Apple’s inspected media page did not establish a scriptable export interface, so it is only a manual-comparison option.
- **Useful search terms:** `Editly JSON5 title transition audioTracks`, `Motion Canvas FFmpeg exporter 1920 1080`, `Playwright deterministic frame capture animation`, `ffmpeg image sequence ProRes H.264`, `Web Animations API SVG title card`, `Manim community output quality`, `Keynote export movie AppleScript`.

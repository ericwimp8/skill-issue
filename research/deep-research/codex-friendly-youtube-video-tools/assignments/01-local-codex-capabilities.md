# Local Codex Capability Boundary for a macOS YouTube Video

## Assignment

**Goal:** Determine the current local and official Codex capability boundary for producing or editing a polished, public 1080p YouTube demonstration on macOS, including capture, rendering, UI automation, image generation, and any native OpenAI video capability.

**Scope:** The `skill-issue` workspace, installed local commands and Codex configuration/skills/plugins, the current session's surfaced capabilities, the caller-provided available-but-uninstalled plugin list, and current primary OpenAI documentation for Codex skills, plugins, Browser, Computer Use, image generation, MCP/connectors, and video generation.

**Exclusions:** Comparing third-party editors, recorders, template tools, audio/caption services, or Build Week rules; those are covered by separate assignments. This assignment does not install plugins, authenticate external services, generate media, operate a GUI editor, or call the OpenAI Videos API.

## Sources

- Local research framing: `research/deep-research/codex-friendly-youtube-video-tools/research-map.md`.
- Active local Codex configuration, especially the enabled `computer-use`, `chrome`, and `browser` plugins and Browser/Computer Use runtime settings.
- Surfaced local ImageGen, Browser, and Computer Use skill contracts; they define raster-image generation, web interaction and screenshots, and local Mac application control respectively.
- Active local plugin manifests for Browser, Computer Use, and Chrome.
- Local command and capability probes run 2026-07-19: Codex version, command availability, FFmpeg filters and encoders, `screencapture` help, installed application inventory, and installed skill/plugin inventory.
- Official OpenAI, [Customization](https://learn.chatgpt.com/docs/customization/overview): skills package reusable instructions/scripts/references and MCP connects external systems (lines 792-869, fetched 2026-07-19).
- Official OpenAI, [Build skills](https://learn.chatgpt.com/docs/build-skills): a skill is a `SKILL.md` directory with optional scripts/references; skills and plugins can bundle capabilities, including MCP setup (lines 755-845, fetched 2026-07-19).
- Official OpenAI, [Plugins](https://learn.chatgpt.com/docs/plugins): plugins can provide skills, connectors, MCP servers, browser extensions, hooks, and task templates (lines 754-842, fetched 2026-07-19).
- Official OpenAI, [Browser](https://learn.chatgpt.com/docs/browser): Browser controls websites and local web apps; in the desktop app, Computer Use can operate it by opening, clicking, typing, inspecting rendered state, and taking screenshots (lines 747-782, fetched 2026-07-19).
- Official OpenAI, [Computer Use](https://learn.chatgpt.com/docs/computer-use): on macOS, Codex can see and operate allowed GUI apps with Screen Recording and Accessibility permissions (lines 745-859, fetched 2026-07-19).
- Official OpenAI, [Image generation](https://learn.chatgpt.com/docs/image-generation): image generation creates or edits images such as UI assets, banners, backgrounds, and illustrations (lines 747-757, fetched 2026-07-19).
- Official OpenAI, [Video generation with Sora](https://developers.openai.com/api/docs/guides/video-generation): the Videos API offers programmatic creation, extension, editing, and management; its current Sora 2 documentation states deprecation and shutdown on 2026-09-24 (lines 757-770 and 1363-1370, fetched 2026-07-19).

## Findings

### Local Deterministic Capture and Render Toolchain

`ffmpeg` 8.1.1, `ffprobe`, and `ffplay` are installed and available on `PATH`. The FFmpeg probes confirmed `libx264`, `h264_videotoolbox`, `hevc_videotoolbox`, `aac`, and `aac_at` encoders, plus `concat`, `xfade`, `overlay`, `scale`, `fps`, `amix`, and `loudnorm` filters. This is a usable local, scriptable path for stitching, transitions, title/overlay compositing, format conversion, 1080p export, audio mixing, and loudness normalization.

The macOS `screencapture` command is installed. Its built-in usage reports `-v` video recording, `-V<seconds>` duration limiting, `-g`/`-G` audio capture, `-k` click visualization, interactive `-J video` selection, display selection, and file output. This supplies a native recording path without another recorder installation.

**Evidence:** Local probes returned `ffmpeg version 8.1.1`, the listed encoders/filters, and `screencapture` usage containing the above video options. `yt-dlp` and `obs` were unavailable. The installed-application inventory contained no matching iMovie, Final Cut Pro, ScreenFlow, CapCut, DaVinci Resolve, Premiere, Camtasia, or OBS application.

**Implication:** Codex can reliably generate and execute narrowly scoped FFmpeg commands and inspect output with `ffprobe`; user-operated capture can begin immediately with `screencapture`. A third-party GUI editor or recorder is optional for this machine, not required for a deterministic capture-and-stitch baseline.

### Browser and Computer Use Are UI Automation, Not Video Editors

The local Codex configuration enables the OpenAI Browser, Chrome, and Computer Use plugins. Browser's local skill is explicitly limited to visiting/interacting with pages, screenshots, and testing local web apps. Computer Use controls the UI of allowed Mac applications and exposes operations such as click, drag, keyboard input, scrolling, and screen-state inspection.

Official documentation places the same boundary: Browser can open, click, type, inspect rendered state, and take screenshots; Computer Use can operate graphical applications on macOS after Screen Recording and Accessibility permissions are granted. It is an appropriate way to drive a chosen GUI editor or web-based production tool where no structured interface exists, subject to its app approval and confirmation safeguards. [Browser](https://learn.chatgpt.com/docs/browser) and [Computer Use](https://learn.chatgpt.com/docs/computer-use).

**Evidence:** The active local Codex configuration enables the three plugins. The local Browser and Computer Use skill descriptions name navigation/testing/screenshots and Mac-app UI control respectively. The official Computer Use guide describes GUI operation rather than media-composition functions.

**Implication:** Browser/Computer Use can assist a human workflow in a third-party editor or hosted tool, including UI-driven trimming and export, but neither supplies timeline editing, render composition, recording orchestration, or a reproducible video pipeline by itself.

### Image Generation Is Available for Still Assets Only

The surfaced `imagegen` skill uses the built-in `image_gen` tool to generate or edit raster images. Its documented outputs are images: illustrations, mockups, product images, banners, sprites, and transparent-background assets. It can create title-card art, thumbnails, backgrounds, or static visual inserts for a video project. [OpenAI's Image generation documentation](https://learn.chatgpt.com/docs/image-generation) describes the same image-only scope.

**Evidence:** The local skill's declared top-level modes are built-in `image_gen` and a `gpt-image` CLI fallback; its only operations are `generate`, `edit`, and `generate-batch` for raster images. No video, timeline, frame-sequence, or motion-rendering operation is declared. The current session's surfaced skills likewise include `imagegen`, Browser, Chrome, and Computer Use, with no video-generation or video-editing skill.

**Implication:** Treat ImageGen as a still-asset input to FFmpeg or a chosen editor. It does not replace video generation, capture, or video editing.

### OpenAI Offers Native Video Generation and Targeted Video Editing Through a Separate API

OpenAI's Videos API provides native Sora video creation, extension, targeted editing, and MP4 download. The official guide says a `POST /videos` render job uses an API key, supports Sora 2/Sora 2 Pro clips of up to 20 seconds, and lists `1920x1080` or `1080x1920` export for `sora-2-pro`. It also documents `POST /v1/videos/edits` for a focused change to an existing video, with uploaded-video edits limited to eligible customers. [Video generation with Sora](https://developers.openai.com/api/docs/guides/video-generation).

This is an OpenAI API product, rather than a current Codex video skill or installed Codex plugin. The guide also states that the documented Sora 2 models and Videos API are deprecated and scheduled to shut down on 2026-09-24.

**Evidence:** The primary API guide explicitly says the Videos API enables programmatic creation, extension, editing, and management (lines 757-770); it documents 1080p `sora-2-pro` output and 16/20-second generations (lines 783-895), and targeted existing-video editing (lines 1363-1370). The local command probe found no `openai` CLI, the workspace has no OpenAI SDK/video dependency, and the local Codex skill/plugin inventory had no `video`, `youtube`, `sora`, `ffmpeg`, `movie`, or editor package beyond a Flutter screenshot-capture skill.

**Implication:** Sora can be considered as a separately authenticated, paid/API-dependent source for short generated motion inserts. It is not an installed native Codex media-production surface in this environment, and its announced deprecation makes it unsuitable as the sole production dependency without confirming the successor path at implementation time.

### Skills, Plugins, and Connectors Can Extend Codex but Do Not Establish a Present Video Capability

Codex skills are instruction packages with optional scripts/references, while plugins can bundle skills, connectors, MCP servers, browser extensions, and hooks. A future video workflow could therefore be expressed as a custom local skill that calls FFmpeg or an authenticated video API, or as a plugin/MCP integration supplied by a video product. Those are extension mechanisms; they do not make an absent tool installed or authenticated. [Customization](https://learn.chatgpt.com/docs/customization/overview), [Build skills](https://learn.chatgpt.com/docs/build-skills), and [Plugins](https://learn.chatgpt.com/docs/plugins).

**Evidence:** The caller-provided available-but-uninstalled plugin list contains Atlassian Rovo, Box, Figma, Google Calendar, Google Drive, Notion, Outlook Calendar, Outlook Email, SharePoint, Slack, and Teams. These are work/data connectors; none is described as a video generator, recorder, editor, or render tool. The installed-plugin configuration and cached manifests likewise show Browser, Chrome, Computer Use, sites, visualization, and domain-specific plugins, without a video-production plugin.

**Implication:** No relevant plugin installation is presently indicated. The best immediately available Codex contribution is authoring and operating reproducible scripts around `screencapture` and FFmpeg, then using Browser/Computer Use only for a selected external editor's UI or upload workflow.

## Notes

- The local `computer-use` MCP server entry is disabled even though the `computer-use` plugin itself is enabled and the capability is surfaced in this session. Actual GUI operation still depends on the active desktop-app/runtime connection, macOS permissions, and per-app approval.
- The local `screencapture` help was obtained by invoking `screencapture -h`; the command prints its usage after reporting `illegal option -- h`. The documented video flags are still present in that output.
- The Sora findings establish an API capability, not account entitlement, pricing, eligibility for uploads/edits, or suitability under the competition's rules. Those remain outside this assignment and need current confirmation if Sora becomes a chosen implementation dependency.
- Searches used for negative inventory evidence: `video`, `youtube`, `sora`, `ffmpeg`, `movie`, `editor`, and `capture`. The conclusion is limited to the skills/plugins/configuration and commands inspected on 2026-07-19.

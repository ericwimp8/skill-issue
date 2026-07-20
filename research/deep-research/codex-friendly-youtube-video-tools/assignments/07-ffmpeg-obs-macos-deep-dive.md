# FFmpeg, OBS Studio, and macOS Capture Deep Dive

## Assignment

**Goal:** Determine the fastest reliable free local capture-and-finish toolchain for a polished generic 1080p YouTube demonstration under three minutes, including local availability, macOS permissions, editing/export capabilities, and the appropriate division between Codex automation and user actions.

**Scope:** Local machine probes on 2026-07-19, official Apple, FFmpeg, OBS, and YouTube documentation, with comparison of macOS Screenshot/QuickTime plus FFmpeg against OBS Studio plus FFmpeg.

**Exclusions:** No recording, app installation, OBS configuration, production-video creation, source edits, paid editors, browser extensions, virtual audio-device installation, or YouTube-upload workflow testing.

## Sources

- Local probe: `screencapture` on macOS 26.2 (25C56); its usage documents `-v` video capture, `-g` default-input audio, `-G<id>` audio-device capture, `-D<display>`, `-R<x,y,w,h>`, `-V<seconds>`, `-k` click highlighting, and `-U` interactive toolbar.
- Local probe: FFmpeg 8.1.1, from the Homebrew formula currently labelled GPL-3.0-or-later. Its build enables `libx264`, VideoToolbox, and AudioToolbox; it exposes `concat`, `xfade`, `fade`, `overlay`, `amix`, `sidechaincompress`, `loudnorm`, `scale`, `fps`, `h264_videotoolbox`, `libx264`, `aac`, and `aac_at`.
- Local probe: this FFmpeg build reports `Unknown filter 'drawtext'` and `Unknown filter 'subtitles'`; its build configuration has no `--enable-libfreetype`, `--enable-libharfbuzz`, or `--enable-libass` flag.
- Local probe: no OBS Studio app was found in the installed-application inventory or Spotlight. A TCC database entry grants an old `com.obsproject.obs-studio` microphone and camera permission, but it does not establish that OBS is installed or can screen-capture now.
- [Apple: record the screen on Mac](https://support.apple.com/en-us/102618) documents Screenshot and QuickTime controls, microphone selection, click indicators, saving, SDR H.264 on macOS Tahoe 26, and selected-window capture on Tahoe 26.
- [Apple: Screen & System Audio Recording access](https://support.apple.com/en-ie/guide/mac-help/mchld6aa7d23/mac) documents the per-app Privacy & Security control.
- [FFmpeg filters manual](https://ffmpeg.org/ffmpeg-filters.html) documents `concat`, `xfade`, `fade`, `overlay`, `amix`, `sidechaincompress`, `loudnorm`, `drawtext`, and subtitle rendering requirements.
- [FFmpeg codecs manual](https://ffmpeg.org/ffmpeg-codecs.html) documents the native AAC encoder and `libx264` options.
- [YouTube upload encoding settings](https://support.google.com/youtube/answer/1722171) recommends MP4, H.264, AAC-LC at 48 kHz, fast-start placement, and 8 Mbps for 1080p SDR at 24/25/30 fps (12 Mbps at 48/50/60 fps).
- [OBS macOS Screen Capture](https://obsproject.com/kb/macos-screen-capture-source), [OBS macOS desktop audio](https://obsproject.com/kb/macos-desktop-audio-capture-guide), [OBS recording output](https://obsproject.com/kb/standard-recording-output-guide), [OBS remote control](https://obsproject.com/kb/remote-control-guide), [OBS VideoToolbox guidance](https://obsproject.com/kb/hardware-encoding), and [OBS license](https://github.com/obsproject/obs-studio/blob/master/COPYING).

## Findings

### Recommended Deadline Toolchain

For this under-three-minute, generic demonstration, use **native Screenshot capture followed by FFmpeg finishing**. Screenshot is already present, records an entire screen, a selected portion, and—on this macOS release—a selected window; Apple exposes microphone, cursor-click, timer, and SDR/H.264 controls in the capture UI. FFmpeg is already installed and covers deterministic scaling, stitching, fades, overlay composition, audio mix/ducking, normalization, and a YouTube-compatible final export. This minimizes setup, avoids an OBS installation and scene-configuration pass, and leaves the repetitive finishing steps scriptable.

**Evidence:** Apple documents Screenshot/QuickTime screen recording and its capture controls; the local `screencapture` help confirms equivalent video/audio/crop-oriented CLI capability. Local FFmpeg probes confirm all finishing filters and H.264/AAC encoders listed above. YouTube specifies H.264/AAC MP4 for upload. [Apple capture documentation](https://support.apple.com/en-us/102618), [FFmpeg filter documentation](https://ffmpeg.org/ffmpeg-filters.html), and [YouTube encoding settings](https://support.google.com/youtube/answer/1722171).

**Implication:** Start with a tight 1920×1080 capture area and a short run sheet; use a title slide or first captured frame rather than spending deadline time on a multi-scene live production. Keep OBS as a fallback only if per-app desktop audio or a live multi-source layout becomes a concrete requirement.

### Native macOS Capture: Capability, Privacy, and User Actions

The user must initiate and stop capture, select the target/window/portion, select a microphone when narration is needed, and grant or review **System Settings → Privacy & Security → Screen & System Audio Recording**. Screenshot’s documented audio selection is a microphone; its documentation does not establish reliable selectable desktop/system-audio capture for this task. The local CLI can request a default input (`-g`) or a specified audio device (`-G`), but a usable device choice has not been tested. Treat system-audio capture through native tooling as unsupported until a short test clip verifies it.

**Evidence:** Apple says Screenshot can record voice or other audio by choosing a microphone, and explicitly directs the per-app recording permission control. The local `screencapture` usage lists `-g` and `-G<id>` without identifying available devices. [Apple recording guide](https://support.apple.com/en-us/102618) and [Apple privacy guide](https://support.apple.com/en-ie/guide/mac-help/mchld6aa7d23/mac).

**Implication:** Native capture is best for screen-led narration. If the demonstration needs application/system sound, test it first; otherwise record narration through the selected microphone and avoid assuming desktop audio will appear. The user should close notifications, private content, and unrelated windows before granting capture—recording access can expose whatever appears in the selected area.

### OBS: Value, Cost, and macOS Audio Advantage

OBS is absent locally, so it adds a download, macOS authorization, source selection, output-profile setup, audio-meter check, and test-recording step. Its modern **macOS Screen Capture** source uses ScreenCaptureKit and can capture a display, window, or application; on macOS 13+ it can include audio. OBS 30+ also has a macOS Audio Capture source that can capture desktop audio or one application. That makes OBS materially stronger when the demonstration requires per-app sound, a webcam/PiP, lower thirds, several prebuilt scenes, or on-demand scene changes.

**Evidence:** The local app-discovery probe found no OBS application. OBS’s official macOS source guide says its Screen Capture source supports display/window/application capture and audio on macOS 13+, and the audio guide specifies the per-application/desktop-audio options. [OBS macOS Screen Capture](https://obsproject.com/kb/macos-screen-capture-source) and [OBS macOS desktop audio](https://obsproject.com/kb/macos-desktop-audio-capture-guide).

**Implication:** OBS is a poor default for this immediate deadline but a good escalation path if native audio testing fails or a composed recording is essential. If adopted, use **macOS Screen Capture**, not the deprecated macOS Window/Display sources, and make one short test recording before the final take.

### OBS Recording, Automation, and Reliability

OBS provides UI-managed Scenes, Sources, text/image/media overlays, audio meters, profiles, and hardware encoders. It recommends MKV for resilient recording and can remux it to MP4; direct MP4/MOV requires finalization and is less recoverable after an interrupted recording. OBS Studio 28+ bundles WebSocket control; it supports password protection and command-line overrides. A local automation client can therefore trigger recording or scene changes once OBS exists and the user has configured/authorized it, but it cannot replace the user’s initial privacy grant, source selection, on-screen-content review, or audio check.

**Evidence:** OBS documents MKV/remux behavior and built-in WebSocket support from OBS 28, including password authentication. OBS also documents Apple VideoToolbox support for recording H.264 on Intel and Apple Silicon Macs. [OBS recording output](https://obsproject.com/kb/standard-recording-output-guide), [OBS remote control](https://obsproject.com/kb/remote-control-guide), and [OBS hardware encoding](https://obsproject.com/kb/hardware-encoding).

**Implication:** Codex can generate/review FFmpeg commands and, only after setup, drive an OBS WebSocket client. It should not attempt unattended screen capture or treat a configured OBS file as proof of a visually correct, privacy-safe recording. For a three-minute deadline, OBS hands-on setup is a planning estimate of 15–30 minutes before retakes; native capture is a planning estimate of 3–8 minutes per prepared take. These are estimates, not measured timings.

### FFmpeg Video Assembly Features

The installed FFmpeg can normalize each clip to a common 1920×1080, 30-fps, `yuv420p` format with `scale`, `pad`, `fps`, and `format`; concatenate matching synchronized A/V segments with `concat`; make a deliberate cut with `fade` or an overlap transition with `xfade`; and place a PNG/logo/title-card video over the capture with `overlay`. `concat` requires each segment to have the same count of audio/video streams, so normalizing each clip before assembly is safer than attempting to stitch arbitrary captures in one complex command.

**Evidence:** Local filter help confirms `concat`, `xfade`, `fade`, `overlay`, `scale`, and `fps`; FFmpeg specifies that `concat` works on synchronized segments with identical stream counts, and that `overlay` takes a main and overlay video. [FFmpeg concat](https://ffmpeg.org/ffmpeg-filters.html#concat), [FFmpeg overlay](https://ffmpeg.org/ffmpeg-filters.html#overlay), and [FFmpeg fade/drawtext documentation](https://ffmpeg.org/ffmpeg-filters.html#fade).

**Implication:** A stable finishing pass can use the following output pattern after inputs are prepared: `-vf "scale=1920:1080:force_original_aspect_ratio=decrease,pad=1920:1080:(ow-iw)/2:(oh-ih)/2,fps=30,format=yuv420p"`. For a polished result, limit transitions to a brief fade/crossfade at intentional chapter boundaries rather than applying them to every cut.

### FFmpeg Titles, Overlays, and Captions Constraint

The upstream `drawtext` filter can render text only when FFmpeg is built with FreeType and HarfBuzz, while subtitle rendering requires the appropriate subtitle-rendering dependency. This installed Homebrew FFmpeg has neither `drawtext` nor `subtitles`; attempts to inspect both return `Unknown filter`. It can still overlay a prepared transparent PNG title/lower-third or a rendered title-card video, because `overlay` is present.

**Evidence:** The local filter probe failed for `drawtext` and `subtitles`; the local build configuration lacks the required optional libraries. FFmpeg documents the FreeType/Harfbuzz requirements for `drawtext`. [FFmpeg drawtext requirements](https://ffmpeg.org/ffmpeg-filters.html#drawtext).

**Implication:** Do not design the deadline path around FFmpeg-generated text or burned-in subtitles. Use Screenshot/OBS-produced visuals or a pre-rendered image/video overlay for the title. Caption generation/upload and subtitle burn-in remain unsupported in the currently installed toolchain without an additional tool or different FFmpeg build; preserve a transcript separately if accessibility captions are required.

### FFmpeg Audio Mix, Ducking, and Normalization

The installed `amix` filter combines narration and music with explicit weights and duration behavior; `sidechaincompress` accepts a main stream and sidechain, enabling music ducking when narration is present; and `loudnorm` implements EBU R128 normalization with integrated-loudness, range, and true-peak targets in one- or two-pass modes. The local filter help exposes all three. A useful sequence is: mix/duck first, then normalize the final audio, with the narration as the sidechain and music as the compressed main input.

**Evidence:** Local help lists `amix`, `sidechaincompress`, and `loudnorm` with their controls. FFmpeg states that `amix` mixes streams, and that `loudnorm` supports dynamic/linear single- and double-pass EBU R128 normalization. [FFmpeg amix](https://ffmpeg.org/ffmpeg-filters.html#amix) and [FFmpeg loudnorm](https://ffmpeg.org/ffmpeg-filters.html#loudnorm).

**Implication:** Use no background music unless it clearly helps the demonstration. If used, keep it understated with `amix` weights or `sidechaincompress`; then normalize once at the end. A voice-only demo can skip mixing entirely and use `loudnorm` alone, reducing failure surface.

### 1080p H.264/AAC Export and Verification

The local FFmpeg includes software `libx264`, hardware `h264_videotoolbox`, native `aac`, and `aac_at`. The MP4 muxer defaults to H.264/AAC and exposes `+faststart`; YouTube asks for MP4 with the `moov` atom at the front, H.264 video, AAC-LC at 48 kHz, and 8 Mbps for SDR 1080p at 30 fps. A reproducible final export after the assembly filter graph is:

```sh
ffmpeg -i assembled.mov \
  -c:v libx264 -preset medium -b:v 8M -maxrate 8M -bufsize 16M \
  -pix_fmt yuv420p -r 30 \
  -c:a aac -b:a 192k -ar 48000 \
  -movflags +faststart final-1080p.mp4
```

For faster finalization, substitute `-c:v h264_videotoolbox -b:v 8M`; retain `yuv420p`, AAC, 48 kHz, and `+faststart`. Before upload, use `ffprobe` to inspect duration, 1920×1080 dimensions, 30-fps rate, H.264/AAC codecs, and audio sample rate; this is automatable and non-destructive.

**Evidence:** Local probes confirm encoder and muxer availability, including VideoToolbox profiles and `faststart`. FFmpeg documents AAC and `libx264`; YouTube supplies the upload requirements and bitrate. [FFmpeg codecs](https://ffmpeg.org/ffmpeg-codecs.html), [YouTube encoding settings](https://support.google.com/youtube/answer/1722171).

**Implication:** Use `libx264` for the most predictable local encode when time permits; use VideoToolbox when finishing speed matters. The command assumes standard 30-fps SDR and a finished input; if the capture is 60 fps, use YouTube’s 12 Mbps 1080p guidance or retain 60 fps deliberately rather than silently changing it.

### Cost, Licensing, Watermarks, and Decision Matrix

| Criterion | Screenshot/QuickTime + installed FFmpeg | OBS + FFmpeg |
| --- | --- | --- |
| Local readiness | Ready: Screenshot and FFmpeg found | Blocked: OBS app absent |
| Cost | No additional tool cost identified | Free download; no purchase identified |
| Licensing | Current Homebrew FFmpeg package reports GPL-3.0-or-later | OBS repository ships GPL version 2 text |
| Watermark | No watermark mechanism appeared in local CLI help or FFmpeg output; confirm with one test clip | No OBS watermark mechanism was found in the inspected official recording docs; confirm with one test clip |
| 1080p/H.264/AAC | Available through FFmpeg final export | Available through OBS/FFmpeg; Apple VideoToolbox supported |
| System/app audio | Unverified with native capture; mic path is documented | Supported on macOS 13+ through modern capture/audio sources |
| Polish controls | Deterministic post-production, but title text/captions unavailable in this FFmpeg build | Native Scenes/Sources/text/media composition |
| Reliability | Fewest moving parts; capture is user-controlled | More setup but resilient MKV/remux path |
| Codex automation | Strong for FFmpeg preparation/export/probe; capture remains user action | WebSocket enables configured OBS control; setup and review remain user action |

**Evidence:** Local availability/license probes, Apple capture documentation, OBS’s licensing and recording documentation, and FFmpeg capability probes. [OBS GPL text](https://github.com/obsproject/obs-studio/blob/master/COPYING), [OBS recording output](https://obsproject.com/kb/standard-recording-output-guide), and [Apple capture documentation](https://support.apple.com/en-us/102618).

**Implication:** Select native Screenshot plus FFmpeg now. Select OBS only when the additional live-production capability repays its setup cost—most notably verified per-app/system audio, webcam composition, or repeatable multi-scene capture.

## Notes

- `ffmpeg` is installed at version 8.1.1 while Homebrew reports a newer stable formula version. This research does not update it; the verified capabilities are those in the currently linked binary.
- No actual screen/audio recording or export was created, so device enumeration, microphone quality, native system-audio capture, capture frame rate, color behavior, and watermark absence remain unverified for this machine.
- Native Screenshot records SDR H.264 only on supported Tahoe 26 models according to Apple; HDR/HEVC is unnecessary for the requested generic 1080p YouTube demo and would add compatibility risk.
- The TCC OBS microphone/camera entry is stale or otherwise insufficient evidence: no OBS app location was found, and Screen & System Audio Recording approval was not inspected or changed.

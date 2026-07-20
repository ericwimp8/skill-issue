# Audio, Captions, Accessibility, and Asset Rights

## Assignment

**Goal:** Establish a free, practical audio, narration, caption, accessibility, privacy, and asset-licensing workflow for the public under-three-minute Build Week YouTube demonstration.

**Scope:** macOS recording, local cleanup and caption tools, YouTube caption delivery, burn-in trade-offs, safe music/sound effects, Codex automation boundary, and human review. Sources were inspected on 2026-07-19.

**Exclusions:** Selecting the overall video editor/capture stack, generating narration or music, and aggregating the cross-tool decision.

## Sources

- Local framing: [`research-map.md`](../research-map.md) and [Build Week requirements research](02-build-week-video-rules.md). The latter records the authoritative requirement for a public YouTube video under three minutes with a clear demo **with audio** covering the project and Codex/GPT-5.6 use.
- Local capability probe, 2026-07-19: FFmpeg is version `8.1.1`; it exposes `anlmdn` and `loudnorm`, but this build exposes neither `subtitles` nor `whisper` filters and its build configuration has no `--enable-libass`. `whisper-cli`, `whisper`, and Audacity are absent.
- [Apple — record the screen on Mac](https://support.apple.com/en-ie/102618): current Screenshot/QuickTime microphone capture and local save behaviour.
- [Apple — create a Voice Memos recording](https://support.apple.com/guide/voice-memos/create-a-recording-vmaa4b813415/3.2/mac/26): built-in/external microphone options and Apple Account syncing.
- [Audacity macOS download](https://www.audacityteam.org/download/mac/) and [Noise Reduction manual](https://manual.audacityteam.org/man/noise_reduction.html): current macOS release/support and effect order.
- [Audacity GPLv3 license](https://github.com/audacity/audacity/blob/master/LICENSE.txt): software license.
- [FFmpeg filters documentation](https://ffmpeg.org/ffmpeg-filters.html): `anlmdn`, `loudnorm`, `whisper`, and `subtitles` capabilities.
- [whisper.cpp README](https://github.com/ggml-org/whisper.cpp/blob/master/README.md) and [MIT license](https://github.com/ggml-org/whisper.cpp/blob/master/LICENSE): local build/model workflow, macOS support, and license.
- [OpenAI Whisper repository](https://github.com/openai/whisper) and [CLI output options](https://github.com/openai/whisper/blob/main/whisper/transcribe.py): Python alternative, model license, and `srt`/`vtt` output formats.
- [YouTube — add subtitles and captions](https://support.google.com/youtube/answer/2734796?hl=en), [automatic captioning](https://support.google.com/youtube/answer/6373554?hl=en-GB), and [viewer caption settings](https://support.google.com/youtube/answer/100078?hl=en): caption upload, editing, accuracy limits, sound descriptions, and viewer-controlled presentation.
- [YouTube Audio Library terms](https://support.google.com/youtube/answer/3376882?hl=en) and [safe-music guidance](https://support.google.com/youtube/answer/15577610?hl=en): account path, attribution, and Content ID guidance.

## Findings

### Local Voice Capture Is the Fastest Low-Risk Start

Use the macOS Screenshot toolbar to capture the selected product area while selecting the narration microphone, or record narration separately in QuickTime/Voice Memos and lay it under the final cut. Apple documents selected-area recording, microphone selection, and local screen-recording saves; QuickTime also supports mic selection and monitoring. Voice Memos is suitable for a quick voice take, but recordings can sync to other Apple devices when the user is signed in to the same Apple Account, so QuickTime/Screenshot is the cleaner local-only default for unpublished material. [Apple screen recording](https://support.apple.com/en-ie/102618) [Apple Voice Memos](https://support.apple.com/guide/voice-memos/create-a-recording-vmaa4b813415/3.2/mac/26)

**Evidence:** The installed macOS is 26.2. Apple’s current instructions explicitly support microphone audio, selected portions, and a local `.mov` save. They do not establish a built-in system-audio route, so do not assume desktop audio can be captured without a separately validated source.

**Implication:** Record a single dry, paced voiceover with headphones and a quiet room; capture product audio only when it communicates required behaviour. The narrator must perform the take and listen for intelligibility, cadence, accidental sensitive speech, and whether the explanation actually covers what was built and how Codex/GPT-5.6 were used.

### Cleanup Has a Scriptable Default and an Optional GUI Repair Path

FFmpeg’s `anlmdn` reduces broadband noise and `loudnorm` implements EBU R128 loudness normalisation with single- and double-pass modes. These are deterministic local filters that Codex can compose and run once the user supplies the selected input/output paths. The local installation exposes both filters. [FFmpeg filters](https://ffmpeg.org/ffmpeg-filters.html)

**Evidence:** The local filter probe confirms `anlmdn` and `loudnorm`; FFmpeg documents the former as a Non-Local Means broadband denoiser and the latter as an integrated-loudness/true-peak normaliser. This only proves filter availability, not acceptable audible output for this narrator or microphone.

**Implication:** Start with one conservative FFmpeg cleanup/render pass and preserve the original take. Do not stack denoising, compression, EQ, and normalisation by default: a person must A/B listen for pumping, metallic artefacts, clipped consonants, and music masking speech. If repair is needed, Audacity is a free local multi-track editor/recorder with a current macOS 3.7.8 download for Intel and Apple Silicon; its manual recommends noise reduction before compression and cautions that stronger reduction can remove wanted sound. It is GPLv3 software, while the exported narration remains an original media asset. [Audacity download](https://www.audacityteam.org/download/mac/) [Noise Reduction](https://manual.audacityteam.org/man/noise_reduction.html) [license](https://github.com/audacity/audacity/blob/master/LICENSE.txt)

### whisper.cpp Is the Preferred Local Caption Generator After One-Time Setup

`whisper.cpp` is an MIT-licensed, native C/C++ Whisper implementation that supports Intel and Apple-silicon Macs, CPU inference, and Apple Metal/Core ML acceleration. Its documented setup clones the repository, downloads a chosen model, builds with CMake, and passes a local 16-bit WAV to `whisper-cli`; its own example shows FFmpeg conversion to mono 16 kHz PCM. The model download requires network access once, while the documented local CLI path supplies the audio directly from disk. [whisper.cpp README](https://github.com/ggml-org/whisper.cpp/blob/master/README.md) [license](https://github.com/ggml-org/whisper.cpp/blob/master/LICENSE)

**Evidence:** `whisper-cli` is currently absent locally, so this is a setup task rather than an immediate final-hour command. The upstream README notes the 16-bit WAV input restriction for its CLI example. FFmpeg’s upstream documentation also has a `whisper` filter capable of creating an SRT, but the installed FFmpeg lacks that filter; it cannot be the selected local route without replacing/rebuilding FFmpeg. [FFmpeg whisper filter](https://ffmpeg.org/ffmpeg-filters.html)

**Implication:** If setup time is available, use `whisper.cpp` for a private, reproducible first SRT/VTT draft, then manually correct technical names, names, punctuation, timing, and non-speech cues. Codex can prepare/run the build, conversion, transcription, and caption-file checks, but cannot validate the transcript against the narrator’s intended meaning or judge readable timing without a user review.

### Python Whisper Is a Capable but Less Minimal Local Alternative

OpenAI’s reference Whisper repository is MIT-licensed and its CLI supports `txt`, `vtt`, `srt`, `tsv`, and JSON output. It uses a Python/PyTorch installation path and downloads model weights, adding dependency and setup time relative to the `whisper.cpp` command-line route. [OpenAI Whisper README](https://github.com/openai/whisper) [CLI source](https://github.com/openai/whisper/blob/main/whisper/transcribe.py)

**Evidence:** The reference repository documents `pip install` from GitHub and its CLI parser lists the caption output formats. No current Python Whisper installation exists in the local probe.

**Implication:** Use it only if it is already installed or a different workflow depends on it. It is not the deadline-safe default for this workspace.

### Uploaded YouTube Captions Are the Accessible Primary Deliverable

YouTube Studio accepts timed caption files, can auto-sync a supplied transcript, or allows manual entry. Timed sidecar captions let viewers turn captions on/off and customise font, colour, opacity, size, background, window, and edge style. Manual captions can include cues such as `[applause]` and `[thunder]`; use equivalent concise cues for meaningful UI sounds, music, or alerts. [Add subtitles and captions](https://support.google.com/youtube/answer/2734796?hl=en) [viewer caption settings](https://support.google.com/youtube/answer/100078?hl=en)

**Evidence:** YouTube’s automatic captions can be delayed and can misrepresent speech because of accents, dialects, background noise, and mispronunciations; YouTube explicitly requires creators to review and edit them. [Automatic captioning](https://support.google.com/youtube/answer/6373554?hl=en-GB)

**Implication:** Upload a corrected SRT/VTT in YouTube Studio as the primary accessibility measure. Set the original language correctly, watch the private/unlisted processing copy with captions enabled, and correct the final timing/text before making the required public release. YouTube account sign-in is required for Studio/upload operations; caption creation and audio repair do not require a cloud transcription account in the proposed local path.

### Burn-In Is an Optional Presentation Variant, Not a Replacement for CC

FFmpeg can render a subtitle file into video frames with the `subtitles` filter when FFmpeg is built with `libass`; it supports file selection and style overrides. Burned text can help an audience that watches a repost with audio muted, but it cannot be disabled, restyled, resized, translated, or repositioned by a viewer and can conceal important screen-demo UI. [FFmpeg subtitles filter](https://ffmpeg.org/ffmpeg-filters.html)

**Evidence:** The local FFmpeg lacks both `subtitles` and `--enable-libass`, so local burn-in is unavailable without a toolchain change. Conversely, YouTube sidecar captions are supported in the expected publishing workflow and preserve the player’s user controls.

**Implication:** Deliver uploadable captions first. Create a burned-caption derivative only when a concrete distribution surface requires it, retain the caption-file master, and visually inspect every line against the 1080p screen content. This prevents an unverified rebuild from becoming a delivery blocker.

### Safe Audio Assets Are Optional; YouTube Audio Library Is the Narrowest Cleared Route

Silence or an original, unobtrusive narration bed is the fastest rights-safe choice. If music or an SFX is genuinely needed, use YouTube Audio Library from YouTube Studio: it supplies royalty-free music and sound effects, lets the user filter for standard tracks requiring no attribution, and labels Creative Commons tracks that require copied attribution text in the video description. YouTube says Audio Library downloads are copyright-safe and are not claimed through Content ID; it only vouches for assets from that library, not content labelled “royalty-free” elsewhere. [YouTube Audio Library](https://support.google.com/youtube/answer/3376882?hl=en) [safe-music guidance](https://support.google.com/youtube/answer/15577610?hl=en)

**Evidence:** Audio Library requires signing in to YouTube Studio and a download action. The surrounding Build Week research requires that any third-party integration or asset be licensed for the use, so retain the selected track title, license type, download date, and required attribution in the release notes/description.

**Implication:** Select only an `Attribution not required` Audio Library asset for the short deadline path, keep it low enough that speech remains plainly understandable, and omit it if it does not improve comprehension. Do not treat YouTube channels, generic “free” labels, or unverified stock SFX as cleared assets.

### Recommended Minimal Workflow and Fallback

Record a local narration take with Screenshot/QuickTime, keep the raw original, and apply a conservative local FFmpeg cleanup/loudness pass. If `whisper.cpp` can be set up without disrupting the deadline, generate a local draft caption file, edit it against the final cut, upload it in YouTube Studio, then do a human headphones-and-captions playback before public publication. Use no music unless it has a clear communication purpose; if used, obtain a no-attribution asset from YouTube Audio Library and retain its license record.

**Evidence:** This combines the already installed FFmpeg audio filters, built-in macOS capture, a current free local caption option, YouTube’s supported sidecar-caption workflow, and the platform’s explicitly copyright-safe asset library.

**Implication:** Fallback when local caption setup or timing runs long: upload the video, wait for YouTube’s automatic captions or paste the spoken script through Auto-sync, then manually correct every error and meaningful sound cue in Studio. This preserves a captioned public deliverable without making a local build or subtitle burn-in a critical path.

## Notes

- **Unsupported:** This run did not measure word-error rate, transcription speed, exact narration loudness, or caption readability for the actual speaker, microphone, room, or final visual composition. Those are perceptual checks requiring the user’s recorded material and review.
- **Privacy boundary:** The local capture/FFmpeg/whisper.cpp route keeps media on disk after setup, but downloading models/software uses the network. Uploading the final video, captions, and Audio Library activity uses the user’s YouTube account and platform. Review the chosen account/channel privacy settings before handling non-public footage.
- **Tool caveat:** The local FFmpeg feature probe is specific to this Homebrew 8.1.1 build; upstream documentation describing `whisper` and `subtitles` does not make those filters available here.
- **Forum caveat:** No forum reports were used to establish quality, licensing, privacy, availability, or account requirements.

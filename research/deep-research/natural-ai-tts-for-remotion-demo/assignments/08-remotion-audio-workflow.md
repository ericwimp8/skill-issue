# Remotion Narration Production Workflow

## Assignment

**Goal:** Define a source-backed, production-oriented workflow for turning AI-generated narration into a reliable voice-over for a public, approximately three-minute technical YouTube demo built with Remotion on macOS.

**Scope:** Audio master and delivery formats; sample rate; script segmentation; stable asset naming and manifests; loudness and peak targets; silence trimming, fades, resampling, and normalization; timing against screen capture; pronunciation proofing and targeted regeneration; captions and alignment when the TTS provider supplies no timestamps; variable-duration audio in Remotion; deterministic automation versus human approval; and final technical and editorial validation.

**Exclusions:** TTS-provider selection, commercial rights, pricing, provider-specific license interpretation, and assumptions about locally installed software or credentials. Provider capabilities are discussed only where they change the workflow.

## Sources

- Remotion, [`<Audio>` from `@remotion/media`](https://www.remotion.dev/docs/media/audio) — current recommended audio component; exact extraction during rendering; timing, trimming, volume, playback-rate, and error behavior.
- Remotion, [`calculateMetadata()`](https://www.remotion.dev/docs/calculate-metadata) — dynamic composition duration and per-composition output sample rate.
- Remotion, [`renderMedia()`](https://www.remotion.dev/docs/renderer/render-media) — output audio codec, bitrate, and sample-rate controls; current default output sample rate is 48 kHz.
- Remotion, [`staticFile()`](https://www.remotion.dev/docs/staticfile) — stable loading of assets placed in the Remotion `public/` folder.
- Remotion, [`<Sequence>`](https://www.remotion.dev/docs/sequence) and [`<Series>`](https://www.remotion.dev/docs/series) — frame-based placement, trimming, sequential scenes, intentional gaps, and overlaps.
- Remotion, [Captions overview](https://www.remotion.dev/docs/captions), [transcribing audio](https://www.remotion.dev/docs/captions/transcribing), [displaying captions](https://www.remotion.dev/docs/captions/displaying), and [exporting subtitles](https://www.remotion.dev/docs/captions/exporting) — caption generation, page timing, burned-in captions, and SRT artifacts.
- Remotion, [`Caption` data type](https://www.remotion.dev/docs/captions/caption) — canonical `text`, `startMs`, `endMs`, `timestampMs`, and `confidence` fields.
- Remotion, [`@remotion/install-whisper-cpp`](https://www.remotion.dev/docs/install-whisper-cpp) — local transcription, token-level timestamps, 16 kHz transcription input, and conversion to Remotion `Caption[]`.
- Remotion, [`@remotion/openai-whisper`](https://www.remotion.dev/docs/openai-whisper) — conversion of Whisper API output into Remotion `Caption[]`.
- YouTube Help, [Recommended upload encoding settings](https://support.google.com/youtube/answer/1722171) — MP4, H.264, AAC-LC/Opus audio, stereo, 48 kHz, and 384 kbps recommended for stereo uploads.
- YouTube Help, [Supported subtitle and closed-caption files](https://support.google.com/youtube/answer/2734698) — UTF-8 SRT support and timed-caption file expectations.
- YouTube Help, [Video and audio formatting specifications](https://support.google.com/youtube/answer/4603579) — preference for high-quality source material and warning that YouTube re-encodes uploads.
- YouTube Help, [Troubleshoot audio or video upload issues](https://support.google.com/youtube/answer/58134) — AAC-LC, 44.1/48 kHz troubleshooting guidance, and testing uploads before publication.
- Audio Engineering Society, [Loudness Basics](https://aes.org/resources/audio-topics/loudness-project/loudness-basics/), [Loudness Normalization](https://aes.org/resources/audio-topics/loudness-project/loudness-normalization/), and [Loudness resources](https://aes.org/audio-topics/loudness/) — ITU-style loudness measurement and AES online-distribution recommendations, including a 16 LUFS normalization reference for track-style online content.
- Apple Podcasts for Creators, [Audio requirements](https://podcasters.apple.com/support/893-audio-requirements) — a primary platform reference for spoken-word audio recommending approximately -16 LKFS, ±1 dB, with true peak not exceeding -1 dBFS.
- ITU-R, [BS.1770-5](https://www.itu.int/rec/R-REC-BS.1770/en) — in-force algorithmic basis for programme loudness and true-peak measurement.
- EBU, [Loudness](https://tech.ebu.ch/loudness/) and [R 128 S2: Loudness in Streaming](https://tech.ebu.ch/publications/r128s2) — loudness/true-peak measurement context and streaming guidance.
- FFmpeg, [Filters documentation](https://ffmpeg.org/ffmpeg-filters.html) — `loudnorm`, `silencedetect`, `silenceremove`, `aresample`, `afade`, and `acrossfade` behavior.
- FFmpeg, [`ffprobe` documentation](https://ffmpeg.org/ffprobe.html) — machine-readable stream, format, codec, duration, and JSON inspection.
- FCC captioning regulation, [47 CFR §79.1](https://www.law.cornell.edu/cfr/text/47/79.1) — useful quality rubric of accuracy, synchronicity, completeness, and placement. It is used here as a quality model, not as a conclusion about this video's legal obligations.

## Findings

### Finding 1: Use a lossless 48 kHz narration master and encode only at delivery

The provider response should be preserved unchanged as a raw artifact, then converted once into a consistent editing master: mono, 48 kHz, 24-bit PCM WAV. Mono is appropriate for a single centered voice and avoids pretending the source contains stereo information. The final video can be stereo—voice centered or dual-mono, with any music bed remaining stereo—and delivered as MP4/H.264 with AAC-LC at 48 kHz. Avoid MP3 as the working master because every later lossy transcode compounds artifacts.

**Evidence:** YouTube's recommended upload settings specify MP4, H.264, AAC-LC or Opus, stereo, 48 kHz, and 384 kbps for stereo audio ([YouTube Help](https://support.google.com/youtube/answer/1722171)). YouTube states that it re-encodes uploads and recommends supplying high-quality source material ([YouTube formatting specifications](https://support.google.com/youtube/answer/4603579)). Remotion's renderer exposes codec, bitrate, and sample-rate controls; its current default output sample rate is 48 kHz and the documentation advises matching the source to avoid resampling artifacts ([`renderMedia()`](https://www.remotion.dev/docs/renderer/render-media)). FFmpeg supports 24-bit signed PCM and WAV output ([FFmpeg formats](https://ffmpeg.org/ffmpeg-formats.html)).

**Implication:** Standardize every approved voice segment to `pcm_s24le`, 48 kHz, mono WAV before assembly. Configure the final Remotion render for 48 kHz AAC and a stereo output; 384 kbps matches YouTube's published stereo recommendation, while Remotion's documented default is 320 kbps.

### Finding 2: Segment by replaceable speech units, not arbitrary character limits

The smallest regeneration unit should normally be one complete sentence or one short compound thought. Combine very short dependent sentences when splitting would produce robotic resets; split long sentences at a clause boundary where a human could naturally breathe. A practical target is roughly 4–15 seconds per clip, with a soft upper bound near 20 seconds. Paragraph and scene boundaries belong in metadata as larger groups, not as giant provider requests.

Recommended hierarchy:

1. `section` — narrative chapter or demo phase.
2. `paragraph` — a coherent visual beat.
3. `unit` — one independently regenerable spoken clip.

Do not bake pauses into punctuation alone. Store intended inter-unit and inter-paragraph gaps explicitly so post-processing can preserve them even if the provider changes prosody.

**Evidence:** Remotion sequences are independently placed and bounded in frames, while `<Series>` stitches positive-duration scenes sequentially and supports positive gaps or negative overlaps ([`<Sequence>`](https://www.remotion.dev/docs/sequence), [`<Series>`](https://www.remotion.dev/docs/series)). The recommended `<Audio>` component also accepts independent `from`, `durationInFrames`, `trimBefore`, and `trimAfter` values ([Remotion `<Audio>`](https://www.remotion.dev/docs/media/audio)). These primitives favor small, stable timeline units whose durations can change without rewriting unrelated scenes.

**Implication:** Codex can split an approved script into semantic units and assign IDs, but the user should approve any split that changes rhetorical emphasis. Regenerating a mispronounced name should replace one unit and cascade later timestamps through the manifest rather than force re-generation of the entire narration.

### Finding 3: Treat the manifest as the audio timeline's source of truth

Use immutable semantic IDs and versioned files. One workable naming scheme is:

```text
voice/raw/010-intro-remotion/u010-v001-provider.ext
voice/work/010-intro-remotion/u010-v003-48k-mono.wav
voice/approved/010-intro-remotion/u010-v003.wav
voice/master/narration-v012.wav
voice/captions/narration-v012.captions.json
voice/captions/narration-v012.srt
```

The unit ID (`u010`) stays stable when audio changes; the version increments. A manifest entry should contain at least:

```json
{
  "id": "u010",
  "sectionId": "s01",
  "paragraphId": "p01",
  "displayText": "FFmpeg resamples the narration to 48 kHz.",
  "spokenText": "Eff eff em peg resamples the narration to forty-eight kilohertz.",
  "pronunciationHints": ["FFmpeg=eff-eff-em-peg"],
  "voiceConfigId": "provider-and-voice-config-hash",
  "scriptSha256": "...",
  "requestSha256": "...",
  "rawPath": "...",
  "approvedPath": "...",
  "approvedSha256": "...",
  "durationMs": 4382,
  "gapAfterMs": 160,
  "startMs": 12140,
  "endMs": 16522,
  "loudnessI": -16.2,
  "truePeakDbtp": -1.3,
  "approval": "approved",
  "approvalNote": "Technical names accepted"
}
```

Keep canonical display text separate from provider-oriented spoken text. Captions and on-screen copy must use canonical spellings such as `FFmpeg`, `Whisper.cpp`, `TypeScript`, and `LUFS`, even when phonetic substitutions are used to control speech.

**Evidence:** `staticFile()` produces framework-safe URLs for assets in Remotion's `public/` folder, and Remotion deliberately recommends it over raw string paths ([`staticFile()`](https://www.remotion.dev/docs/staticfile)). Remotion's caption type separates text and millisecond timing into interoperable JSON fields ([`Caption`](https://www.remotion.dev/docs/captions/caption)).

**Implication:** Codex can deterministically generate, validate, hash, and update the manifest. Render code should consume only approved paths and derived timings. A stale script hash, missing approved hash, duplicate ID, unapproved unit, or duration mismatch should fail the build rather than silently fall back to a raw provider file.

### Finding 4: Put a 20–40 second technical-name proof before full generation

Generate a small proof before spending time on the approximately three-minute narration. Produce at least two variants: A uses canonical text; B uses phonetic or SSML hints where supported. Keep voice settings identical so the test isolates pronunciation and prosody.

Suggested 25–30 second proof, canonical display text:

> On macOS, this Remotion demo turns a TypeScript timeline into a narrated video. FFmpeg prepares 48-kilohertz audio, while Whisper.cpp supplies word timings for captions. Watch how Codex updates the composition, measures LUFS loudness, and keeps every screen capture aligned. If Remotion, FFmpeg, TypeScript, Whisper.cpp, or LUFS sounds wrong, regenerate only this proof before creating the full narration.

Suggested spoken hints for the alternate variant: `FFmpeg` → “eff-eff-em-peg”; `Whisper.cpp` → “Whisper dot C P P”; `LUFS` → “L U F S” or “luffs”; `48 kHz` → “forty-eight kilohertz.” The correct version is whichever sounds natural in the selected voice and matches the intended audience.

The proof gate is approved only after the user checks voice identity, accent, pace, energy, pauses, technical-name pronunciation, numerical reading, sibilance, clipping, and whether the voice still sounds credible through a short screen-capture sequence.

**Evidence:** The workflow's value follows from independently replaceable units supported by Remotion timing primitives ([`<Sequence>`](https://www.remotion.dev/docs/sequence), [`<Audio>`](https://www.remotion.dev/docs/media/audio)). The final delivery is public speech content, and AES guidance emphasizes perceptual loudness management rather than peak-only matching ([AES Loudness Basics](https://aes.org/resources/audio-topics/loudness-project/loudness-basics/)).

**Implication:** Codex can generate proof variants, name them, compute measurements, and prepare an A/B page or review render. Only the user can approve naturalness and pronunciation. Full narration generation must remain blocked until one proof variant and its pronunciation map are explicitly approved.

### Finding 5: Trim edges conservatively, preserve authored pauses, and hide splice clicks

Silence processing should remove accidental leading/trailing dead air while retaining a short handle—typically 60–120 ms—around speech. Do not automatically remove every internal quiet region: interior pauses convey phrasing and may be timed to a visual action. Detect silence first, calculate only the first and last speech boundaries, then trim with retained handles. Store intended pauses in `gapAfterMs` rather than relying on whatever silence the provider happened to emit.

Apply very short equal-power or triangular fades—roughly 5–15 ms—to clip boundaries to prevent clicks. Use `acrossfade` only where two waveform edges intentionally overlap. When there should be a narrative pause, fade out, insert manifest-defined silence, and fade in instead of overlapping spoken phonemes.

**Evidence:** FFmpeg's `silencedetect` reports silence start, end, and duration; `silenceremove` can remove silence from the beginning, middle, or end and can preserve a configured amount; `acrossfade` supports overlapping and non-overlapping crossfades ([FFmpeg filters](https://ffmpeg.org/ffmpeg-filters.html)). Remotion also supports non-destructive frame trimming through `<Audio trimBefore trimAfter>` ([Remotion `<Audio>`](https://www.remotion.dev/docs/media/audio)).

**Implication:** Codex can script silence analysis, compute candidate bounds, trim, fade, and insert declared gaps. Thresholds are voice-output dependent; the user must audition the proof and at least the assembled paragraph boundaries to detect clipped consonants, swallowed breaths, or unnaturally short pauses.

### Finding 6: Normalize the assembled narration with an ITU-style two-pass measurement

Use integrated loudness and true peak, not sample peak alone. A defensible speech-oriented house target for this demo is **-16 LUFS integrated** with **maximum true peak at or below -1 dBTP**. Keep the mix relatively controlled, but avoid crushing the voice merely to hit a number. If music or sound effects are added, measure the complete mix, then confirm speech remains intelligible on laptop speakers.

The safest automated sequence is:

1. Measure individual approved units and flag outliers; avoid hard-normalizing every sentence to an identical instantaneous loudness.
2. Assemble approved units and declared gaps into a lossless narration master.
3. Run FFmpeg `loudnorm` pass one with `I=-16`, `TP=-1`, and a modest LRA target such as 7 to obtain measured JSON.
4. Run pass two using the measured integrated loudness, LRA, true peak, threshold, and offset; explicitly resample the result to 48 kHz.
5. Re-measure the output and fail if integrated loudness is outside the chosen internal tolerance (for example ±0.5 LU) or true peak exceeds -1 dBTP.

**Evidence:** ITU-R BS.1770 defines programme-loudness and true-peak measurement ([ITU-R BS.1770](https://www.itu.int/rec/R-REC-BS.1770/en)). AES describes ITU-style metering and cites 16 LUFS normalization for track-style online content ([AES Loudness Normalization](https://aes.org/resources/audio-topics/loudness-project/loudness-normalization/)). Apple gives a primary spoken-word platform reference of approximately -16 LKFS ±1 dB with true peak no higher than -1 dBFS ([Apple audio requirements](https://podcasters.apple.com/support/893-audio-requirements)). FFmpeg's `loudnorm` implements EBU R128 normalization in one- and two-pass modes and can target integrated loudness, LRA, and maximum true peak; its documentation notes dynamic processing may upsample to 192 kHz and advises explicitly setting the output sample rate ([FFmpeg `loudnorm`](https://ffmpeg.org/ffmpeg-filters.html)).

**Implication:** Codex can run two-pass normalization and enforce numeric gates. The user must approve the sonic result, because a technically compliant file can still sound over-compressed, harsh, lispy, or too dense for a technical explanation.

### Finding 7: Build timing from approved audio, then capture visuals with handles

For a narrated technical demo, lock the script and approved narration before recording the definitive screen capture. The audio is less forgiving to time-stretch than a cursor hold, terminal pause, or cutaway. Once approved clips are assembled, compute each unit and paragraph's `startMs`, `endMs`, and frame boundaries. Use `Math.ceil(milliseconds * fps / 1000)` for end frames so a clip is never truncated by fractional-frame rounding.

Record screen actions against the approved paragraph audio or a timing rehearsal. Capture a short handle before and after each action. In Remotion, trim the capture, freeze a clean frame, slow only non-critical cursor travel, or insert a brief visual hold to absorb small duration changes. If narration changes materially, regenerate the relevant unit and let later scene times cascade from the manifest instead of stretching speech. Use narration playback-rate changes only as an explicitly auditioned exception.

**Evidence:** Remotion sequences shift and bound content in frames, `<Series>` computes sequential placement and supports gaps/overlaps, and the recommended `<Audio>` component can be placed with `from` and bounded with `durationInFrames` ([`<Sequence>`](https://www.remotion.dev/docs/sequence), [`<Series>`](https://www.remotion.dev/docs/series), [`<Audio>`](https://www.remotion.dev/docs/media/audio)). Remotion documents `playbackRate` but also makes clear that it changes duration and may have environment-specific behavior ([Remotion `<Audio>`](https://www.remotion.dev/docs/media/audio)).

**Implication:** Codex can derive the entire frame plan from measured audio, update Remotion scene props, and flag screen clips shorter than their assigned scene. The user must approve whether the screen action feels comfortably paced and whether pauses land on the intended proof moments.

### Finding 8: Prefer one assembled voice master in Remotion, with manifest-derived scene markers

Keep unit WAV files for regeneration, but render the final voice track as one assembled, normalized `narration-master.wav`. A single `<Audio src={staticFile(...)} />` avoids dozens of independently decoded elements, makes caption alignment refer to one continuous clock, and ensures the loudness measurement describes what Remotion actually plays. The manifest still supplies unit/paragraph boundaries for visual scenes, chapter markers, captions, and targeted replacement.

Use `calculateMetadata()` to read precomputed JSON metadata and return the composition's `durationInFrames` as the narration duration plus an intentional tail (for example, 300–700 ms), along with `defaultSampleRate: 48000`. The render path should fail on audio processing errors rather than silently continue with missing voice. Remotion's current recommended component is `Audio` from `@remotion/media`, which extracts exact audio with Mediabunny to keep it synchronized to the Remotion timeline.

**Evidence:** Remotion explicitly recommends `<Audio>` from `@remotion/media` and states that exact extraction keeps audio synchronized with its timeline ([Remotion `<Audio>`](https://www.remotion.dev/docs/media/audio)). `calculateMetadata()` can dynamically return `durationInFrames` and a per-composition default sample rate before render ([`calculateMetadata()`](https://www.remotion.dev/docs/calculate-metadata)). `staticFile()` is the supported way to resolve public assets ([`staticFile()`](https://www.remotion.dev/docs/staticfile)).

**Implication:** Codex can deterministically generate `narration-master.wav`, `narration-manifest.json`, and Remotion metadata. Variable TTS duration becomes data, not a manual constant. Scene code should reference semantic manifest markers rather than hard-coded frame numbers.

### Finding 9: When provider timestamps are absent, align the final master after post-processing

Provider timestamps—if available—describe the provider's raw output and become stale after trimming, gap insertion, crossfades, or regeneration. Therefore captions should be aligned against the **final assembled master**.

Recommended fallback hierarchy:

1. Sentence/paragraph timing comes exactly from the manifest's assembled segment boundaries.
2. For word timing, run the final master through a timestamp-capable transcription route. Remotion supports local Whisper.cpp, browser Whisper, the OpenAI Whisper API, and ElevenLabs STT; its Whisper.cpp integration can request token-level timestamps and convert output to `Caption[]`.
3. Diff the transcript against canonical script text. Correct recognized proper names (`FFmpeg`, `Remotion`, `Whisper.cpp`, package names, flags, and code identifiers) while retaining valid time intervals.
4. If word alignment is visibly unstable, fall back to sentence-level captions from exact manifest boundaries rather than inventing precise word timestamps.
5. Export UTF-8 SRT for YouTube and optionally burn in styled captions through Remotion.

**Evidence:** Remotion lists local and cloud transcription routes and recommends its interoperable `Caption` format ([Transcribing audio](https://www.remotion.dev/docs/captions/transcribing)). Its Whisper.cpp package documents token-level timestamps and conversion into captions ([`@remotion/install-whisper-cpp`](https://www.remotion.dev/docs/install-whisper-cpp)). Remotion caption objects contain text, millisecond start/end, timestamp, and confidence fields ([`Caption`](https://www.remotion.dev/docs/captions/caption)). YouTube accepts plain UTF-8 SRT files ([YouTube caption formats](https://support.google.com/youtube/answer/2734698)).

**Implication:** Codex can transcribe, reconcile canonical tokens, generate Remotion caption JSON, serialize SRT, and validate monotonic timing. A human must proofread every technical name and inspect synchronization. Automatic transcription is an alignment aid, not editorial authority.

### Finding 10: Caption QC should test meaning, timing, completeness, and obstruction

Caption validation must extend beyond valid SRT syntax. Check that captions match canonical speech, begin and end with the corresponding utterance, cover the entire narration, remain readable long enough, and do not hide terminal commands, cursor targets, or other proof-critical UI. For a technical demo, wrong casing or punctuation can change meaning (`npm`, `TypeScript`, `@remotion/media`, `--flag`, file paths), so canonical script text should override ASR spelling.

**Evidence:** The FCC's quality framework requires accuracy, synchronicity, completeness, and appropriate placement, including correct proper names, punctuation, numbers, and unobstructed important visuals ([47 CFR §79.1](https://www.law.cornell.edu/cfr/text/47/79.1)). Remotion can derive caption pages from millisecond timing and place each page in a frame-based `<Sequence>` ([Displaying captions](https://www.remotion.dev/docs/captions/displaying)). It can also serialize captions to SRT artifacts ([Exporting subtitles](https://www.remotion.dev/docs/captions/exporting)).

**Implication:** Codex can validate spelling against a technical glossary, detect overlaps/gaps/out-of-range timestamps, and render a caption review video. The user must inspect placement against the actual screen capture and approve readability at normal playback speed.

### Finding 11: Codex can automate reproducibility; the user owns perceptual approval

Codex can deterministically:

- freeze the approved script, segment it, assign semantic IDs, and hash the text/configuration;
- generate proof and full clips when the chosen provider, credentials, and invocation are available;
- preserve raw responses and provider request metadata;
- use `ffprobe` to inspect codec, channel count, sample rate, sample format, and duration in JSON;
- resample, trim candidate edges, apply short fades, insert declared pauses, concatenate, and two-pass normalize;
- measure every unit and the final master; reject missing, silent, clipped, stale, or out-of-spec files;
- produce the manifest, cumulative timings, frame plan, caption JSON, SRT, and Remotion input props;
- render a proof, low-resolution review, and final MP4; and
- re-probe outputs and produce a machine-readable QC report.

The user must audition and approve:

- voice identity, accent, authenticity, warmth, energy, and pace;
- every difficult technical name, abbreviation, number, and code-like token in the proof;
- prosody, sentence grouping, paragraph pauses, and whether silence trimming sounds natural;
- any regenerated unit in its paragraph context, not only in isolation;
- speech intelligibility over music/effects on headphones and ordinary laptop speakers;
- screen-action timing and proof readability at 1× speed;
- caption wording, synchronization, and placement; and
- the processed unlisted YouTube upload before making the video public.

**Evidence:** `ffprobe` supports selected stream/format fields and JSON output suitable for scripts ([`ffprobe`](https://ffmpeg.org/ffprobe.html)). Remotion exposes deterministic frame timing and render configuration ([`<Sequence>`](https://www.remotion.dev/docs/sequence), [`renderMedia()`](https://www.remotion.dev/docs/renderer/render-media)). YouTube itself recommends test uploads when checking playback quality ([YouTube formatting specifications](https://support.google.com/youtube/answer/4603579)).

**Implication:** The workflow should encode human approval as manifest state. Automation may prepare and validate a candidate, but it must not mark `proofApproved`, `pronunciationApproved`, `captionsApproved`, or `publishApproved` without explicit user confirmation.

### Finding 12: Concrete end-to-end sequence

1. **Freeze the script.** Normalize punctuation and technical casing; create `displayText`, provider-specific `spokenText`, and a pronunciation glossary. Hash the script.
2. **Segment.** Assign `section`, `paragraph`, and `unit` IDs. Add intended `gapAfterMs` values and visual beat descriptions.
3. **Generate the proof.** Create the 20–40 second technical-name proof in canonical and hinted variants. Preserve raw files and request metadata.
4. **Process the proof.** Inspect with `ffprobe`; convert to mono 48 kHz/24-bit PCM WAV; detect and conservatively trim only edge silence; add 5–15 ms fades; assemble declared gaps; run two-pass `loudnorm` to -16 LUFS / -1 dBTP; re-measure.
5. **Review the proof.** Render a small Remotion proof with representative screen capture and canonical captions. User approves one voice/configuration and a pronunciation map.
6. **Generate full narration.** Generate only approved units with the frozen voice configuration. Retry failed units idempotently; never overwrite an approved version.
7. **Targeted regeneration.** User flags unit IDs. Change only their spoken hints or punctuation, regenerate version `vNNN`, reprocess, and audition each replacement in paragraph context.
8. **Assemble the master.** Concatenate approved unit WAVs plus manifest-defined gaps. Create a single lossless `narration-master.wav`. Run final two-pass normalization and probe it.
9. **Compute timing.** Measure final duration, cumulative unit/paragraph bounds, and frame bounds. Add an intentional tail. Feed duration and 48 kHz into `calculateMetadata()`.
10. **Capture and fit visuals.** Record definitive screen actions against paragraph audio with head/tail handles. Trim or hold visuals to match manifest markers; avoid speech time-stretching.
11. **Align captions.** Transcribe the final master with token timestamps where available; reconcile to canonical script; fall back to exact sentence timing where word alignment is weak; generate `Caption[]` and UTF-8 SRT.
12. **Render review.** Use one `@remotion/media` `<Audio>` for the voice master. Render a low-resolution full review with captions, then user watches at 1× and approves.
13. **Render delivery.** Produce MP4/H.264, stereo AAC-LC, 48 kHz, with a high-quality bitrate such as YouTube's recommended 384 kbps stereo setting.
14. **Technical QC.** Probe container, codecs, sample rate, channels, bitrate, frame rate, and duration; rerun loudness and true-peak measurement; validate manifest hashes and caption bounds.
15. **Platform QC.** Upload as unlisted. After YouTube finishes higher-quality processing, compare voice quality, sync, and uploaded captions on headphones and laptop speakers. Publish only after user approval.

**Evidence:** Every stage maps to a documented primitive: FFmpeg measurement and processing filters ([FFmpeg filters](https://ffmpeg.org/ffmpeg-filters.html)); machine-readable probing ([`ffprobe`](https://ffmpeg.org/ffprobe.html)); Remotion audio, sequences, metadata, and rendering ([`<Audio>`](https://www.remotion.dev/docs/media/audio), [`<Series>`](https://www.remotion.dev/docs/series), [`calculateMetadata()`](https://www.remotion.dev/docs/calculate-metadata), [`renderMedia()`](https://www.remotion.dev/docs/renderer/render-media)); Remotion caption conversion and export ([Captions](https://www.remotion.dev/docs/captions)); and YouTube's 48 kHz AAC-LC delivery guidance ([YouTube upload settings](https://support.google.com/youtube/answer/1722171)).

**Implication:** This order minimizes expensive rework: voice and pronunciation are proven before full generation, the audio master is fixed before definitive screen capture, and platform-specific quality is checked before publication.

### Finding 13: Final validation gates

The automated QC report should fail on any of the following:

- manifest ID collision, missing source/approved file, stale script hash, unapproved unit, or checksum mismatch;
- zero/near-zero audio, unexpected channel count, non-48-kHz master, lossy working master, or inconsistent sample formats;
- measured final loudness outside the house tolerance or true peak above -1 dBTP;
- narration master duration that differs from the manifest assembly duration beyond the declared fade/crossfade math;
- a Remotion composition shorter than `ceil((narrationMs + tailMs) * fps / 1000)`;
- caption timestamps that are negative, zero-length, non-monotonic, overlapping unexpectedly, or outside the narration duration;
- captions missing script units or containing unapproved spellings of glossary terms;
- final MP4 missing H.264 video, AAC audio, stereo layout, or 48 kHz sample rate; or
- output duration mismatch beyond one video frame plus a small codec/container tolerance.

The human release gate should include a full uninterrupted 1× watch, a second pass focused on technical names and captions, playback on headphones and laptop speakers, and an unlisted YouTube playback after transcoding. Any changed unit invalidates downstream paragraph-audio, timing, caption, full-watch, and platform approvals until they are rerun.

**Evidence:** FFmpeg and `ffprobe` expose the measurements required for deterministic checks ([FFmpeg filters](https://ffmpeg.org/ffmpeg-filters.html), [`ffprobe`](https://ffmpeg.org/ffprobe.html)). Remotion renders from frame-count metadata and exposes sample-rate/codec controls ([`calculateMetadata()`](https://www.remotion.dev/docs/calculate-metadata), [`renderMedia()`](https://www.remotion.dev/docs/renderer/render-media)). YouTube publishes the target delivery structure and recommends testing actual uploads ([YouTube upload settings](https://support.google.com/youtube/answer/1722171), [YouTube formatting specifications](https://support.google.com/youtube/answer/4603579)).

**Implication:** Completion is evidenced by both a passing machine QC report and explicit perceptual approvals. A successful Remotion render alone does not prove pronunciation, caption accuracy, intelligibility, or post-YouTube quality.

## Notes

- YouTube's official upload documentation inspected for this assignment specifies codecs, channels, bitrate, and sample rate but does not publish a definitive LUFS target. The -16 LUFS / -1 dBTP house target above is a defensible speech-oriented recommendation derived from AES online-audio guidance and Apple's primary spoken-word requirements; it is not presented as a YouTube acceptance requirement. Widely repeated third-party claims of a fixed -14 LUFS YouTube target were not treated as authoritative.
- Silence thresholds and retained handles are voice dependent. Values such as -50 dB and 60–120 ms are starting parameters for a proof, not universal acceptance criteria.
- Whisper token timestamps are an automated alignment estimate. Technical terms, punctuation, and precise word boundaries remain subject to human review; sentence-level manifest timing is the safer fallback when token timing is unreliable.
- Provider-returned audio formats, timestamp availability, SSML/phoneme support, determinism controls, and retry semantics vary. Adapt only the generation adapter; keep the approved WAV, manifest, assembly, Remotion, caption, and QC stages provider-neutral.

# Natural AI TTS for a Remotion Build Week Demo

## Executive Recommendation

**Use SpeechifyAI `simba-3.2` through the official API as the first production path.** It is the strongest overall fit for this one-off public demo because its recurring Free API plan is documented as requiring no card, includes 50,000 characters, permits commercial use, exposes TypeScript/Python SDKs, supports SSML substitutions and delivery controls, and returns word-level speech marks that can feed Remotion timing. A conservative final take plus two full retakes consumes only 6,300–9,000 characters; even an eight-voice 20–40 second audition remains comfortably inside the allowance. Current independent blind-listener evidence gives Simba 3.2 the strongest native-provider naturalness signal among the API challengers studied, although only the real technical script can establish the winning voice. The commercial permission is plan-specific evidence for lawful publication, not an express assignment of output ownership: the API supplemental terms classify Synthetic Output as Customer Content but use the same ownership wording for Free and paid API plans. Paying for SpeechifyAI API capacity therefore does not improve the Build Week ownership posture. ([Speechify pricing](https://speechify.ai/pricing), [API terms](https://speechify.com/terms-ai-voice-api/), [models](https://docs.speechify.ai/docs/get-started/models), [speech marks](https://docs.speechify.ai/docs/features/speech-marks), [SSML](https://docs.speechify.ai/docs/ssml/), [Artificial Analysis](https://artificialanalysis.ai/text-to-speech/leaderboard/provider-voice?tab=leaderboard); synthesis from [API challengers](assignments/04-api-challengers.md), [hosted landscape](assignments/01-hosted-api-landscape.md), and [rights audit](assignments/07-rights-privacy-audit.md))

**Begin narration production with one 60–75-word, 20–40 second proof containing every difficult technical name.** Generate the same proof with the available curated Simba 3.2 stock voices, keep canonical and pronunciation-hinted versions, and approve one voice only after listening on laptop speakers and headphones. Do not generate the full narration until the proof passes the acceptance gate below.

The main disagreement in the evidence is real rather than contradictory:

- **Speechify wins overall** on measured stock-voice naturalness, zero cash/card friction, commercial Free use, native timing, and Codex automation.
- **ElevenLabs Starter wins the quality-first production package** because Eleven v3 is top-tier in a controlled-voice benchmark and the service has the richest combined pronunciation, expressive-direction, alignment, and creator workflow. It costs USD $6 plus tax and requires generation under the paid plan for contest use. ([ElevenLabs pricing](https://elevenlabs.io/pricing), [models](https://elevenlabs.io/docs/overview/models), [controlled-voice leaderboard](https://artificialanalysis.ai/text-to-speech/leaderboard/controlled-voice); [ElevenLabs deep dive](assignments/03-elevenlabs-deep-dive.md))
- The benchmarks use different voice conditions and do not test this script, Australian authenticity, three-minute continuity, or its technical vocabulary. The proof gate therefore controls the final selection.

## Winners and Ranked Shortlist

### Explicit Winners

| Category                      | Winner                                | Why                                                                                                                                                                                                                                  |
| ----------------------------- | ------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| **Best overall**              | **SpeechifyAI Simba 3.2 API**         | Best combined naturalness signal, $0 recurring allowance, commercial Free use, SSML, word timings, and direct Codex control; retain the separate sole-ownership caveat.                                                              |
| **Cheapest viable**           | **SpeechifyAI Free API**              | Exact cash cost is $0 if the account exposes the documented 50,000-character Free plan; it covers proof, three full passes, and repairs. Paying for API capacity does not strengthen ownership wording.                              |
| **Best quality**              | **ElevenLabs Starter with Eleven v3** | Strongest quality-first creator package; use Multilingual v2 if v3 is too variable or precise speed/break control matters more than expressiveness. This remains subject to the listening proof.                                     |
| **Best Codex-controlled/API** | **SpeechifyAI Simba 3.2 API**         | Official TypeScript SDK, non-streaming file output, SSML, and native word marks minimize adapter and alignment work.                                                                                                                 |
| **Best browser-only**         | **Descript Free, conditionally**      | No-card proof path, 100 one-time credits, ElevenLabs models, WAV/MP3 plus SRT/VTT export, and strong editing handoff. It must pass the clean-audio export, exact-credit, and current-rights gate. Murf Creator is the paid fallback. |
| **Best privacy-conscious**    | **Kokoro-82M via MLX-Audio**          | Local Apple-Silicon inference after download, no per-use fee, fixed voices, inline IPA, inspectable phonemes, lossless WAV, and no remote script submission.                                                                         |

### Ranked Shortlist

1. **SpeechifyAI Simba 3.2 API** — default production recommendation.
2. **ElevenLabs Starter** — quality-first fallback and best mature creator workflow.
3. **Descript Free** — best human-operated browser proof and potentially the final browser asset.
4. **Kokoro-82M through MLX-Audio** — best local/private path and no-cost benchmark.
5. **Murf Creator** — best paid browser fallback for Australian accent and exact pronunciation direction.
6. **Deepgram Aura-2** — best no-card Australian API alternative with inline IPA and an Australian endpoint.
7. **Google Cloud Chirp 3: HD** — best cloud privacy statement and broad en-AU/IPA controls, with GCP billing friction.
8. **OpenAI API speech** — clearest rights-first contract path and strong Codex integration, but weaker native alignment/pronunciation control, a USD $5 funding floor, and a current model-deprecation ambiguity.

## Decision Matrix

| Candidate                     | Voice-quality evidence                                                                                         | Pronunciation controls                                                                | Real one-video cost and allowance                                                                                                   | Rights and disclosure                                                                                                                                    | Privacy                                                                                                           | Timing/alignment                                                   | Workflow friction                                                                             | Codex leverage                                               |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------ | --------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| **SpeechifyAI Simba 3.2 API** | Strongest current native-provider blind-listener signal among compared APIs; exact script untested             | SSML aliases, rate, pitch, volume, pauses, emphasis, emotions                         | **$0**; 50,000 chars/month; three passes use 6,300–9,000                                                                            | Free API permits commercial use; output is Customer Content, without express assignment; paid API uses the same ownership wording; disclose every use    | API-specific retention/training period is unclear; general policy permits improvement use                         | Native word/sentence marks with millisecond times and text offsets | Low; no card, official SDK, ≤2,000 chars on non-streaming endpoint                            | **Very high**                                                |
| **ElevenLabs Starter**        | Eleven v3 is second in the controlled-voice benchmark; rich creator evidence; actual stock voice untested      | v3 IPA/audio tags; pronunciation dictionaries; Multilingual v2 aliases, breaks, speed | **$6 + tax**; 30,000 credits; three passes estimated 6,300–9,000                                                                    | Paid-era output has durable commercial rights; Free output is ineligible for a prize contest; disclose AI voice conservatively                           | Disable model improvement before generation; ordinary retention remains; zero retention is Enterprise-only        | Character-level original and normalized alignment                  | Medium; account, payment, audition, nondeterministic takes                                    | **Very high**                                                |
| **Descript Free**             | Uses ElevenLabs models; browser voice fit remains subjective                                                   | Phonetic respelling, punctuation, tone tags; US-English bias documented               | **$0 if gate passes**; 100 one-time credits; published estimate implies about 20 TTS minutes, but model-specific debit is not exact | User owns protectable output as between parties; public/commercial Free-audio entitlement and clean export should be captured at generation time         | Disable project-data sharing; avoid custom speakers                                                               | WAV/MP3 plus timed transcript and SRT/VTT                          | Low for human operation; no card; exact export behavior must be proven                        | Medium after export; browser generation stays human-operated |
| **Kokoro/MLX-Audio**          | No directly applicable independent three-minute stock-voice test; strongest presets are publisher-rated        | Inline IPA, inspectable phonemes, respelling, speed                                   | **$0**; unlimited local use; estimated 15–45 minute setup on supported Apple Silicon                                                | Apache-2.0 model and MIT wrapper permit broad use; model license does not itself settle the contest's sole-ownership interpretation; disclose provenance | **Local after model download**                                                                                    | No native word timing; use unit manifest and align final master    | Medium; Python/model download, but small 355 MB BF16 model and first-class Apple Silicon path | High                                                         |
| **Murf Creator**              | Subjective reviews favor its dedicated workflow; actual voice still requires audition                          | IPA, respelling, locale override, pauses, pacing; en-AU voices                        | **About $29 + tax** for a one-off monthly plan; 2 hours/month; annualized display is $19/month but requires annual billing          | Paid commercial rights survive cancellation; output ownership is less explicit than OpenAI, paid Speechify Studio, or Descript                           | Consumer data is US-hosted; consumer no-training commitment is unclear                                            | SRT/VTT and block timing                                           | Medium; excellent UI, manual export, higher cash floor                                        | Medium downstream                                            |
| **Deepgram Aura-2**           | Limited/early independent evidence; script-specific quality unresolved                                         | Inline IPA overrides with warning headers; speed 0.7–1.5                              | **$0 while $200 promotional credit remains**; no card; nominal three-pass cost $0.189–$0.270                                        | Royalty-free lawful output use; cannot present output as human; sole-ownership allocation is less explicit                                               | `mip_opt_out=true` exists, but public price is tied to MIP and opt-out price is unresolved; AU endpoint available | No documented Aura-2 TTS word marks                                | Low; official REST/SDK and named en-AU voices                                                 | High                                                         |
| **Google Chirp 3: HD**        | Current benchmark signal trails Simba and older OpenAI voices; actual en-AU voice untested                     | IPA/X-SAMPA, substitutions, pauses, pace; SSML is Preview and synchronous-only        | **$0 within 1M chars/month**; billing and payment method required; nominal over-free cost $0.189–$0.270                             | Audio may be used in video/media; classic TTS output ownership is not explicitly allocated enough for the contest warranty                               | **Strongest cloud posture:** TTS states it does not log submitted text or returned audio                          | Chirp `<mark>` timing is unresolved; align separately              | Medium-high GCP project/credential/billing setup                                              | High                                                         |
| **OpenAI API speech**         | Older TTS-1/HD benchmark signal trails Simba; current instruction-driven model lacks a comparable live ranking | Natural-language direction or respelling; no documented IPA/dictionary/SSML           | Nominal TTS-1: **$0.095–$0.135**; HD: **$0.189–$0.270**; new accounts generally require **$5 prepaid**                              | Clearest reviewed output-ownership allocation; built-in API voice; clear AI disclosure required; ChatGPT Voice audio is not a substitute                 | API data is not used for training by default; up to 30-day abuse logs; approved ZDR possible                      | No native TTS word marks; use manifest/forced alignment            | Low if account exists; model catalog/deprecation conflict is a release risk                   | Very high                                                    |

Quality statements above distinguish official capability claims from subjective evidence. Artificial Analysis provides blind-listener preference evidence, not proof of this script's winner. Browser review signals for Murf and Descript are subjective workflow evidence. No assignment generated or listened to production audio. ([API challengers](assignments/04-api-challengers.md), [browser studios](assignments/05-browser-studios.md), [local options](assignments/06-local-open-options.md))

## Cost, Free-Tier, and Payment Reality

The production budget is **one approved narration plus two complete retakes**: about 6,300–9,000 characters or nine generated minutes. In practice, unit-level repairs should cost less. A 20–40 second proof across all eight Simba voices adds an estimated 2,900–3,600 characters, so the recommended Speechify campaign remains roughly 9,200–12,600 characters before small repairs—well below 50,000.

| Route                          | Exact or best-supported sufficiency                                                             | Cash/card caveat                                                                                                                                            |
| ------------------------------ | ----------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **SpeechifyAI Free API**       | 50,000 chars/month; exact fit is more than five worst-case three-pass campaigns                 | **$0 and no card documented**; commercial use included; hard cap; paid API is unnecessary for rights and only adds capacity/features                        |
| **ElevenLabs Free**            | 10,000 credits technically covers three estimated passes with little margin                     | **Not lawful for the prize submission**: Free excludes commercial/contest use and has title attribution; upgrading later does not cure Free-generated audio |
| **ElevenLabs Starter**         | 30,000 credits; roughly ten 3,000-char passes                                                   | **$6 + jurisdictional tax**; payment required; exact no-card claim unsupported                                                                              |
| **Descript Free**              | 100 one-time credits; approximately 20 TTS minutes at the published five-credit/minute estimate | **$0 and no card documented**; exact v3 debit, clean Free audio export, and current public-use entitlement are proof gates                                  |
| **Murf Free**                  | Ten generated minutes numerically covers three three-minute passes                              | **Cannot download or use the final commercially**; use only for audition                                                                                    |
| **Murf Creator monthly**       | Two hours/month                                                                                 | About **$29 + tax** for a true month-to-month purchase; $19 display assumes annual billing                                                                  |
| **Speechify Studio Free**      | 600 credits = ten generated minutes                                                             | Free lacks commercial rights and MP3 download; Studio Free card requirement is unsupported                                                                  |
| **Speechify Studio Starter**   | 7,200 credits = 120 minutes                                                                     | **$19/month + tax**; paid commercial rights; browser generation remains human-operated                                                                      |
| **Deepgram**                   | $200 promotional credit covers far more than one video; credit documented with one-year expiry  | **$0 and no card** while credit is active; public pricing opts into model improvement                                                                       |
| **Google Chirp 3: HD**         | First 1M chars/month free                                                                       | Usage is **$0**, but billing/payment activation is mandatory; a new-customer trial may place a temporary authorization hold                                 |
| **OpenAI API**                 | No API free tier for the preferred route; nominal usage is below $0.30 for legacy TTS families  | New prepaid accounts generally require **$5 minimum**, default $10; credits expire after one year and auto-recharge should be disabled for one-off use      |
| **Kokoro or Chatterbox local** | Unlimited local inference after download                                                        | **$0 model/API fee**; Mac time, bandwidth, electricity, and setup are real but unpriced                                                                     |

## What Codex Can Do and What the User Must Do

### Codex Can Execute

- Freeze and hash canonical display text, provider-oriented spoken text, and a technical pronunciation glossary.
- Split the script into semantic units, build Speechify SSML aliases, call the supported API, preserve raw audio/speech marks, and log model, voice, request, and character usage.
- Convert approved clips to mono 48 kHz/24-bit PCM WAV, inspect them with `ffprobe`, conservatively trim edge silence, add fades and declared gaps, concatenate, and run two-pass loudness normalization.
- Build an immutable manifest, cumulative millisecond/frame timing, Remotion props, `Caption[]`, UTF-8 SRT, and a machine-readable QC report.
- Place the frozen master in Remotion `public/`, use `staticFile()` and `Audio` from `@remotion/media`, derive duration through `calculateMetadata()`, render review/final media, and validate codec, duration, loudness, captions, and checksums. ([Remotion audio](https://www.remotion.dev/docs/media/audio), [`staticFile()`](https://www.remotion.dev/docs/staticfile), [`calculateMetadata()`](https://www.remotion.dev/docs/calculate-metadata); [Remotion workflow](assignments/08-remotion-audio-workflow.md))

### The User Must Do

- Create/sign into the provider account, accept current terms, obtain the API key, confirm the displayed plan, and handle any payment or privacy toggle.
- Audition and approve voice identity, accent, naturalness, pace, energy, every technical name, and the final complete take.
- Decide whether the Speechify privacy posture is acceptable for the script; use only public-bound text and no secrets or personal data.
- Approve caption wording/placement, full-screen timing, the final 1× watch, and the unlisted YouTube transcode before publication.
- Make the entrant's rights/compliance decision, retain dated evidence, apply the AI disclosure, and publish the final public YouTube video.

## Proof-Generation Gate

Use a 60–75-word excerpt that lasts 20–40 seconds and contains **“OpenAI Build Week,” “Codex,” “GPT-5.6,” “Remotion,” the exact project/skill name, one acronym, one version or number, one package/repository token, a deliberate pause, and one emotionally important transition.** Keep `displayText` canonical and create a separate `spokenText`/SSML variant for pronunciation corrections.

1. Confirm the account shows SpeechifyAI Free, 50,000 characters, commercial use, and no payment requirement; save a dated capture of the plan and API terms.
2. Generate the canonical excerpt with each available curated Simba 3.2 stock voice using identical settings and the non-streaming endpoint so word marks are returned.
3. Reject obvious mismatches; for the best two voices, make at most one purposeful SSML alias/pronunciation revision.
4. Preserve raw audio, marks, request metadata, voice ID, character usage, and script/config hashes.
5. Convert the two finalists to matched 48 kHz WAV and loudness for blind comparison; render one representative Remotion proof with canonical captions.

**Acceptance requires all of the following:**

- Every technical name, acronym, number, and code-like token is intelligible and correct.
- No skipped, repeated, rushed, slurred, clipped, metallic, or hallucinated speech is heard.
- The voice sounds credible rather than exaggerated on both headphones and ordinary laptop speakers.
- Pace and pauses fit the representative screen capture; word/caption timing shows no visible drift.
- The projected full narration is safely under 180 seconds; target roughly 380–430 spoken words and leave a small ending tail.
- The exported file is usable, non-silent, unclipped, and machine-readable; any undocumented tag or metadata is recorded rather than stripped.
- The user explicitly approves one voice/configuration and pronunciation map.

If Speechify fails on quality after this bounded gate, move to **ElevenLabs Starter**. If it fails specifically because an Australian stock voice is required, audition **Deepgram Theia and Hyperion**. If cloud processing is unacceptable, run the same gate with **Kokoro `af_heart` and `af_bella`**.

## Script-to-Audio-to-Remotion Workflow

### 1. Script and Segmentation

- Target **380–430 spoken words** rather than the 500-word ceiling so the video remains comfortably below the Build Week requirement of less than three minutes. ([Build Week rules](https://openai.devpost.com/rules))
- Maintain canonical `displayText`, provider-specific `spokenText`, and a glossary for names, acronyms, versions, URLs, and code tokens.
- Segment into complete sentences or short compound thoughts, normally **4–15 seconds** and no more than about 20 seconds. Preserve section and paragraph grouping separately.
- Store pauses explicitly as `gapAfterMs`; do not rely solely on punctuation or provider-emitted silence.
- Keep Speechify non-streaming requests below the documented 2,000-character limit and generate replaceable units so a failed name does not consume another full pass. ([Speechify API limits](https://docs.speechify.ai/docs/api-limits))

### 2. Naming and Manifest

Use stable semantic unit IDs and versioned files:

```text
voice/raw/010-intro/u010-v001-speechify.mp3
voice/work/010-intro/u010-v002-48k-mono.wav
voice/approved/010-intro/u010-v002.wav
voice/master/narration-v001.wav
voice/captions/narration-v001.captions.json
voice/captions/narration-v001.srt
```

The manifest should record `id`, section/paragraph, display/spoken text, pronunciation hints, provider/model/voice/config, script/request hashes, raw and approved paths/hashes, request ID, billed characters, raw speech marks, duration, trim offsets, gap, cumulative start/end, loudness, true peak, and explicit approval state. Render code consumes only approved assets and fails on stale hashes, duplicate IDs, missing files, or unapproved units. ([Remotion workflow](assignments/08-remotion-audio-workflow.md))

### 3. Generate and Correct Pronunciation

- Freeze the accepted proof configuration, generate each semantic unit, and preserve the unmodified provider response.
- Use canonical text first; add Speechify SSML `<sub>`-style aliases or other documented controls only to terms that fail.
- Regenerate only the failed unit, version it, and audition the replacement in paragraph context.
- Never invoke TTS from inside a Remotion render. Freeze the accepted audio artifact because generative output is not promised byte-deterministic.

### 4. Create the Master

- Preserve provider originals, then standardize approved units to **mono, 48 kHz, 24-bit PCM WAV (`pcm_s24le`)**.
- Detect silence first; trim only accidental leading/trailing dead air while retaining roughly **60–120 ms** handles. Preserve internal authored pauses.
- Add short **5–15 ms** fades to avoid splice clicks; insert manifest-defined gaps rather than overlapping speech.
- Concatenate approved units into one lossless `narration-master.wav`.
- Run ITU-style two-pass normalization to a speech-oriented house target of **-16 LUFS integrated** and **true peak ≤ -1 dBTP**, explicitly resampling the output to 48 kHz. Re-measure and require an internal tolerance such as ±0.5 LU. This is a production target, not a YouTube acceptance rule. ([ITU BS.1770-5](https://www.itu.int/rec/R-REC-BS.1770/en), [AES loudness](https://aes.org/resources/audio-topics/loudness-project/loudness-normalization/), [FFmpeg filters](https://ffmpeg.org/ffmpeg-filters.html))

### 5. Timing, Captions, and Synchronization

- Treat the **final assembled master** as timing authority. Provider marks seed word timing, but trimming, gaps, fades, and regeneration change offsets.
- Recalculate unit/paragraph start and end times from the approved manifest; adjust provider word marks into the final master clock and spot-check them against audio.
- If word timing is unstable, align the final master with local Whisper.cpp or another timestamp-capable route, then reconcile recognized text to canonical spellings. Fall back to exact sentence timing rather than inventing precise word marks. ([Remotion transcription](https://www.remotion.dev/docs/captions/transcribing), [Whisper.cpp integration](https://www.remotion.dev/docs/install-whisper-cpp))
- Generate Remotion `Caption[]` and UTF-8 SRT. Captions use canonical casing and punctuation even when speech used phonetic aliases. Check accuracy, synchronicity, completeness, and placement over proof-critical UI.
- Convert times to frames from the manifest; use ceiling for end frames so clips are not truncated.

### 6. Fit Screen Capture to Approved Narration

- Lock narration before definitive screen capture. Record each visual action with head/tail handles while rehearsing against paragraph audio.
- Fit visuals with cuts, held frames, or slower non-critical cursor movement. Avoid speech time-stretching except as a user-auditioned exception.
- Place one assembled voice master in Remotion using `Audio` from `@remotion/media` and `staticFile()`. Use manifest markers for scene boundaries and `calculateMetadata()` for duration plus a 300–700 ms intentional tail and `defaultSampleRate: 48000`. ([Remotion audio](https://www.remotion.dev/docs/media/audio), [`calculateMetadata()`](https://www.remotion.dev/docs/calculate-metadata))

### 7. Delivery and QC

- Render **MP4/H.264** with **stereo AAC-LC at 48 kHz**; 384 kbps matches YouTube's published stereo recommendation. The voice may remain centered/dual-mono while music stays stereo. ([YouTube upload settings](https://support.google.com/youtube/answer/1722171))
- Fail automated QC on missing/stale assets, invalid hashes, silence, wrong sample rate/channels, loudness/peak failure, manifest-duration mismatch, truncated composition, caption errors, or a final MP4 missing H.264/AAC/48 kHz.
- Require a full uninterrupted 1× watch, a technical-name/caption pass, headphone and laptop-speaker checks, and an unlisted YouTube playback after platform transcoding.

## Rights, Privacy, and Build Week Checklist

- [ ] The final video is **less than three minutes**, publicly visible on YouTube, and has clear audio. ([Build Week rules](https://openai.devpost.com/rules))
- [ ] The script, screen capture, music, images, trademarks, APIs/SDKs, and all other third-party material are owned, licensed, or otherwise authorized for submission and sponsor promotion.
- [ ] A vendor-owned stock voice is used. No celebrity, judge, teammate, community, marketplace, deceased-person, or cloned voice is used.
- [ ] The account tier, generation date, model, voice ID, request IDs, usage record, and dated controlling pricing/terms pages are archived with the final audio.
- [ ] The final Speechify output is generated under the current API plan that displays commercial permission; a browser Studio Free preview is not substituted.
- [ ] SpeechifyAI API is not upgraded solely for ownership: current Free and paid API plans share the same supplemental ownership wording, while paid Speechify Studio is a separate product with an express perpetual-output statement.
- [ ] The video credits and YouTube description say: **“Narration uses an AI-generated voice created with SpeechifyAI; no human speaker is being represented.”** This satisfies the safest common disclosure posture. ([Speechify API terms](https://speechify.com/terms-ai-voice-api/))
- [ ] Any provider-requested “Voices powered by Speechify” notice is honored; no unconditional logo attribution was found in the researched API terms.
- [ ] The YouTube altered/synthetic-content disclosure is selected conservatively; disclosure is not replaced by an inaudible watermark or metadata tag. ([YouTube synthetic-content guidance](https://support.google.com/youtube/answer/14328491))
- [ ] No secrets, credentials, private source, customer information, or unnecessary personal data appear in the uploaded script.
- [ ] Speechify's current privacy terms are accepted for this public-bound script; API-specific retention/training remains documented as unresolved. For sensitive text, switch to Kokoro local or a stronger no-logging/ZDR route.
- [ ] No provenance watermark or metadata is intentionally stripped. Speechify's inspected API documents do not establish an embedded watermark, so absence is not claimed.
- [ ] The final WAV/MP3, master, manifest, captions, terms evidence, and publication copy are downloaded and retained before any account downgrade or deletion.
- [ ] If the organizer interprets “solely owned” more strictly than Speechify's customer-content/commercial-use language, obtain organizer confirmation or use the OpenAI API built-in-voice rights-first fallback. ([rights audit](assignments/07-rights-privacy-audit.md))

## Conditional Alternatives

- **ElevenLabs Starter — quality-first:** Pay USD $6, use a current Default voice, try `eleven_v3` Natural first, then `eleven_multilingual_v2` for stable long-form timing and alias/break control. Generate every retained clip after paid entitlement begins. Use `mp3_44100_128` on Starter, preserve character alignment, and archive before cancellation. Free output is ineligible for the contest. ([ElevenLabs deep dive](assignments/03-elevenlabs-deep-dive.md))
- **OpenAI API built-in voice — rights-first:** Best explicit output-ownership allocation and default no-training posture. Use only the API Speech endpoint, not ChatGPT Voice. Confirm a current non-deprecated model in the live account, disclose AI speech, and budget the USD $5 funding floor. Pronunciation requires instructions/respelling and timing requires external alignment.
- **Deepgram Aura-2 — Australian/API:** Audition `aura-2-theia-en` and `aura-2-hyperion-en`; use inline IPA and the Australian endpoint. Confirm MIP opt-out pricing and the organizer's ownership interpretation before publication.
- **Google Chirp 3: HD — cloud privacy:** Prefer when no TTS text/audio logging and broad en-AU voice choice outweigh GCP billing setup. Use synchronous SSML/IPA and plan a separate alignment step. Obtain contest-ownership confirmation.
- **Descript Free — browser-only:** Use if the clean 48 kHz WAV/SRT export, exact credit debit, pronunciation, and current public-use terms pass the proof. Disable data sharing first. Keep browser operations human.
- **Murf Creator — paid browser direction:** Use when Australian accent, IPA/locale overrides, and a dedicated voiceover UI justify about USD $29 for one month. Archive WAV/FLAC and SRT/VTT before cancellation; obtain contest-ownership confirmation.
- **Kokoro/MLX-Audio — local/private:** Use `af_heart`, then `af_bella`; generate locally after downloads, use inline IPA, preserve Apache/MIT notices, disclose model provenance, and obtain organizer confirmation for the sole-ownership warranty.

## Lower-Fit and Rejected Options

- **Chatterbox:** permissive and local but about nine times Kokoro's model footprint, higher macOS/MPS setup risk, no documented exact-pronunciation override, stochastic takes, and a PerTh watermark in every output. Its strongest quality evidence depends on cloned reference voices, which this workflow excludes.
- **Speechify Studio:** Starter is good nominal value at USD $19, but the API is cheaper, more automatable, and has better timing evidence. Studio Free has no final download/commercial rights.
- **Azure Neural TTS:** mature SSML and boundary metadata, but current regional paid price was not extractable and independent flagship-voice evidence was insufficient to displace the shortlist.
- **Amazon Polly:** mature, rights-friendly, and has an Australian generative voice, but its generative engine does not return speech marks, forcing a quality-versus-alignment tradeoff.
- **Fish Audio S2 Pro:** low nominal cost, expressive controls, timestamps, and credible subjective quality; free output is personal-use only and paid stock/community voice licensing needs closer voice-specific confirmation.
- **Resemble hosted API:** current pricing, usable free credit, stock-output rights, and retention were too incomplete for a decision.
- **Piper:** excellent deterministic phoneme control/privacy but lower expected premium-narration quality and uncertain macOS packaging.
- **F5-TTS:** MIT code but official pretrained checkpoints are CC-BY-NC, creating avoidable promotional/prize-use ambiguity.
- **MeloTTS:** permissive and en-AU capable, but stale release/MPS issues create more deadline risk than Kokoro.
- **Coqui XTTS-v2:** rejected because the model/output license is non-commercial and imposes downstream-license obligations.
- **Cartesia Sonic 3.5:** technically strong but rejected for this Adelaide-based run because its current privacy policy says the service is designed for US users and not intended outside the US.
- **PlayAI/PlayHT:** rejected because the privacy page says the service has shut down while API/docs remain reachable and current pricing/quota terms are incoherent.

## Limitations, Risks, and Classified Gates

### True Gates Before Full Generation

1. **Account/plan gate:** the live SpeechifyAI account must expose the documented Free API entitlement, commercial use, 50,000 characters, and usable API key.
2. **Listening gate:** a human must approve the 20–40 second proof. No assignment generated or listened to audio, so no exact stock voice is established.
3. **Pronunciation gate:** every project-specific term must pass canonical and corrected audition before the full script.
4. **Duration gate:** the accepted voice must project the approved script below 180 seconds with an intentional tail.

### Publication Gates

- **Rights evidence:** archive the plan/terms, voice ID, generation date, requests, and output. If the organizer demands an express assignment rather than customer-content/commercial-use language, obtain confirmation or switch to the OpenAI API built-in-voice route.
- **Disclosure:** apply clear human-readable AI narration disclosure regardless of embedded provenance technology.
- **Final approval:** the user must approve the complete take, captions, screen timing, and unlisted YouTube transcode.

### Caveated or Unsupported Points

- Speechify API-specific text/audio retention, default model-training use, and a stock Australian Simba 3.2 voice are unresolved in the supplied research.
- SpeechifyAI Free API is expressly commercial, but neither Free nor paid API terms reviewed contain an express output-rights assignment; paying for API capacity does not resolve the Build Week “solely owned” wording.
- No provider benchmark tests this exact three-minute technical script; leaderboard rankings are subjective, moving priors.
- No inspected Speechify API source affirmatively proves the presence or absence of an inaudible watermark or metadata provenance tag.
- Descript Free's exact v3 credit debit and clean audio-export behavior require the proof; Murf/Speechify Studio Free are numerically sufficient but publication-ineligible.
- OpenAI's current instruction-driven TTS model pages and deprecation surfaces conflict; confirm the live non-deprecated model before use.
- Deepgram's MIP opt-out price, Chirp 3 native `<mark>` timing, and contest-sufficient ownership wording for Deepgram/Google/Murf/local models remain unresolved.
- Local-model licenses permit broad model use but do not by themselves prove copyrightability or sole ownership of every generated waveform.
- This report is product and workflow research, not legal advice.

## Ecosystem Candidate Record

### Discovered Candidates

- **Hosted/API:** SpeechifyAI, ElevenLabs, OpenAI Speech/TTS, Google Cloud Chirp 3: HD, Deepgram Aura-2, Murf API, Azure AI Speech, Amazon Polly, Fish Audio, Resemble AI, Cartesia, and PlayAI/PlayHT.
- **Browser-first:** Descript AI Speech, Murf Studio/Creator, Speechify Studio, and ElevenLabs browser/Studio.
- **Local/open:** Kokoro-82M via MLX-Audio, Chatterbox, Piper, F5-TTS, MeloTTS, and Coqui XTTS-v2.

### Selected Deep Dives

- **Direct focused analysis:** SpeechifyAI Simba 3.2, ElevenLabs v3/Multilingual v2, OpenAI Speech, Deepgram Aura-2, Google Chirp 3: HD, Descript, Murf Creator, Speechify Studio, Kokoro/MLX-Audio, and Chatterbox.
- **Reason for selection:** collectively covered the strongest available listening evidence, free/no-card/low-cost plans, explicit pronunciation approaches, browser and API workflows, Australian voices, native and external alignment, local privacy, public-use rights, and Remotion integration.

### Skim-Only Candidates

- Azure AI Speech, Amazon Polly, Fish Audio, Resemble hosted API, Piper, F5-TTS, and MeloTTS.

### Rejects

- Cartesia for the current non-US eligibility statement; PlayAI/PlayHT for shutdown/pricing contradiction; Coqui XTTS-v2 for non-commercial model/output restrictions.

## Constrained Research Inputs

- [01 — Hosted/API landscape](assignments/01-hosted-api-landscape.md)
- [02 — Browser/open landscape](assignments/02-browser-open-landscape.md)
- [03 — ElevenLabs deep dive](assignments/03-elevenlabs-deep-dive.md)
- [04 — API challengers](assignments/04-api-challengers.md)
- [05 — Browser studios](assignments/05-browser-studios.md)
- [06 — Local/open options](assignments/06-local-open-options.md)
- [07 — Rights/privacy audit](assignments/07-rights-privacy-audit.md)
- [08 — Remotion audio workflow](assignments/08-remotion-audio-workflow.md)

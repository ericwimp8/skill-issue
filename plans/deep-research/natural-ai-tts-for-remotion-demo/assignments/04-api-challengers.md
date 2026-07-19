# Hosted API Challengers for a Remotion Demo Narration

## Assignment

**Goal.** Compare OpenAI Speech/TTS, SpeechifyAI Simba 3.2, Deepgram Aura-2, and Google Cloud Text-to-Speech Chirp 3: HD for a public, approximately three-minute, 350–500-word OpenAI Build Week YouTube demo assembled in Remotion on macOS. Rank practical suitability across naturalness, technical-name pronunciation, cost for three full passes (6,300–9,000 characters or about nine minutes of audio), signup and payment friction, repeatable scripting, audio formats, Australian voice availability, controls, timestamps, streaming/export, Remotion integration, public/commercial rights, disclosure, privacy, cloning consent, and geography.

**Scope.** Internet-only research current on 19 July 2026. Primary provider documentation and current terms control product, pricing, privacy, and rights claims. Independent listening benchmarks are used only as directional naturalness evidence. Findings distinguish verified facts, implementation inferences, and unresolved points.

**Exclusions.** Browser-only studios, consumer subscriptions, local/open-weight inference, non-API editing workflows, and hackathon-rule interpretation are outside this assignment. No provider audio was generated because that would require creating accounts or spending credits; consequently, this report does not claim a project-specific listening winner for the exact narration script.

## Sources

### OpenAI

- [Text-to-speech guide](https://developers.openai.com/api/docs/guides/text-to-speech)
- [Speech API reference](https://platform.openai.com/docs/api-reference/audio/speech-audio-done-event?lang=curl)
- [`gpt-4o-mini-tts` model page](https://developers.openai.com/api/docs/models/gpt-4o-mini-tts)
- [`tts-1` model page](https://developers.openai.com/api/docs/models/tts-1)
- [`tts-1-hd` model page](https://developers.openai.com/api/docs/models/tts-1-hd)
- [All-model catalog](https://developers.openai.com/api/docs/models/all)
- [Deprecations](https://developers.openai.com/api/docs/deprecations)
- [Supported countries and territories](https://developers.openai.com/api/docs/supported-countries)
- [Prepaid billing overview](https://help.openai.com/en/articles/8264778-what-is-prepaid-billing)
- [Prepaid billing setup](https://help.openai.com/en/articles/8264644-what-is-prepaid-billin)
- [API data controls by endpoint](https://platform.openai.com/docs/models/default-usage-policies-by-endpoint)
- [How business and API data is used](https://help.openai.com/en/articles/5722486-api-data-usage-policies)
- [OpenAI Services Agreement](https://cdn.openai.com/osa/openai-services-agreement.pdf)
- [Service Terms](https://openai.com/policies/service-terms/)
- [API output copyright help article](https://help.openai.com/en/articles/5008634)

### SpeechifyAI

- [SpeechifyAI pricing](https://speechify.ai/pricing)
- [Models](https://docs.speechify.ai/docs/get-started/models)
- [Model and API changelog](https://docs.speechify.ai/changelog)
- [Official SDKs](https://docs.speechify.ai/tts/text-to-speech/get-started/official-sdks)
- [Create Speech API reference](https://docs.speechify.ai/tts/api-reference/text-to-speech/audio/speech)
- [Stream Speech API reference](https://docs.speechify.ai/tts/api-reference/v1/audio/stream)
- [API limits](https://docs.speechify.ai/docs/api-limits)
- [Speech marks](https://docs.speechify.ai/docs/features/speech-marks)
- [SSML](https://docs.speechify.ai/docs/ssml/)
- [Streaming](https://docs.speechify.ai/docs/features/streaming)
- [Language support](https://docs.speechify.ai/docs/features/language-support)
- [Voice API](https://docs.speechify.ai/tts/api-reference/voices/list)
- [AI Voice API Supplemental Terms](https://speechify.com/terms-ai-voice-api/)
- [Privacy Policy](https://speechify.com/privacy/)

### Deepgram

- [Aura voices and languages](https://developers.deepgram.com/docs/tts-models)
- [Aura-2 voice controls](https://developers.deepgram.com/docs/tts-voice-controls)
- [Aura-2 formatting](https://developers.deepgram.com/docs/improving-aura-2-formatting)
- [TTS getting started](https://developers.deepgram.com/docs/text-to-speech)
- [TTS media output settings](https://developers.deepgram.com/docs/tts-media-output-settings)
- [Custom and regional endpoints](https://developers.deepgram.com/reference/custom-endpoints)
- [Aura-2 controls changelog, 4 May 2026](https://developers.deepgram.com/changelog/2026/5/4)
- [Pricing](https://deepgram.com/pricing)
- [Model Improvement Partnership Program](https://developers.deepgram.com/docs/the-deepgram-model-improvement-partnership-program)
- [Business Terms of Service](https://deepgram.com/business/Business_TOS.pdf)
- [Managing projects and promotional-credit expiry](https://developers.deepgram.com/guides/deep-dives/managing-projects)

### Google Cloud

- [Chirp 3: HD voices](https://docs.cloud.google.com/text-to-speech/docs/chirp3-hd)
- [Cloud TTS voices and languages](https://docs.cloud.google.com/text-to-speech/docs/list-voices-and-types)
- [Cloud TTS basics and media use](https://docs.cloud.google.com/text-to-speech/docs/basics)
- [Cloud TTS pricing](https://cloud.google.com/text-to-speech/pricing)
- [Cloud TTS release notes](https://docs.cloud.google.com/text-to-speech/docs/release-notes)
- [Cloud TTS data logging](https://docs.cloud.google.com/text-to-speech/docs/data-logging)
- [SSML and timepoints](https://docs.cloud.google.com/text-to-speech/docs/ssml)
- [Synthesis API timepoint reference](https://docs.cloud.google.com/text-to-speech/docs/reference/rest/v1beta1/text/synthesize)
- [Chirp 3 Instant Custom Voice](https://docs.cloud.google.com/text-to-speech/docs/chirp3-instant-custom-voice)
- [Google Cloud free trial](https://docs.cloud.google.com/free/docs/free-cloud-features)
- [Google Cloud Platform Terms](https://cloud.google.com/terms/)
- [Service Specific Terms](https://cloud.google.com/terms/service-terms)

### Independent Quality and Remotion

- [Artificial Analysis native-provider-voice TTS leaderboard](https://artificialanalysis.ai/text-to-speech/leaderboard/provider-voice?tab=leaderboard)
- [Speech Arena early TTS leaderboard](https://www.speecharena.org/leaderboard)
- [EmergentTTS-Eval NeurIPS 2025 paper](https://papers.neurips.cc/paper_files/paper/2025/file/04970a25af46918606ba2cf0a3d7905d-Paper-Datasets_and_Benchmarks_Track.pdf)
- [Remotion audio component](https://www.remotion.dev/docs/audio)
- [Remotion `staticFile()`](https://www.remotion.dev/docs/staticfile)

## Findings

### Finding 1 — SpeechifyAI Simba 3.2 is the best baseline for this one-off demo

**Prose.** **Recommendation (inference):** start with SpeechifyAI `simba-3.2`, an explicitly selected curated voice, MP3 or WAV output, sentence-boundary chunking below 2,000 characters, and the non-streaming `/v1/audio/speech` endpoint so the result includes word-level speech marks. It offers the best combined fit for this exact task: the strongest current independent naturalness signal among the four compared families, zero cash and card friction at the requested volume, full commercial use on the free tier, strong SSML/emotion controls, a TypeScript SDK, exportable audio, and native word timing for Remotion captions or scene synchronization. The main caveats are the absence of a verified stock Australian voice in the eight-voice Simba 3.2 allow-list and weakly specified API-specific retention/training terms.

**Evidence.** The free plan is $0, requires no credit card, includes 50,000 TTS characters per month with a hard cap, and explicitly permits commercial use. Three passes of 6,300–9,000 characters fit comfortably inside that allowance. Speechify recommends `simba-3.2` for new English integrations and currently lists eight curated voices. The speech endpoint returns speech marks with sentence/paragraph and word-level start/end times in milliseconds. SSML supports rate, pitch, volume, fixed pauses, emphasis, substitution aliases, and 13 emotion styles. The official JavaScript/TypeScript package is `@speechify/api`. On the Artificial Analysis provider-native-voice leaderboard captured during this research, Simba 3.2 ranked second overall at Elo 1,232, versus OpenAI TTS-1 HD at 1,104, TTS-1 at 1,084, and Google Chirp 3: HD at 1,052. The leaderboard derives Elo from blind user choices about which matched-text sample sounds more natural. ([pricing](https://speechify.ai/pricing), [models](https://docs.speechify.ai/docs/get-started/models), [speech marks](https://docs.speechify.ai/docs/features/speech-marks), [SSML](https://docs.speechify.ai/docs/ssml/), [SDKs](https://docs.speechify.ai/tts/text-to-speech/get-started/official-sdks), [Artificial Analysis](https://artificialanalysis.ai/text-to-speech/leaderboard/provider-voice?tab=leaderboard))

**Implication.** Speechify minimizes the likelihood that narration generation or caption timing becomes a schedule risk. Before committing the whole script, audition all eight Simba 3.2 voices against a 30–45 second excerpt containing the product name, “OpenAI,” “Remotion,” acronyms, numbers, and the most emotionally important transition. If an Australian voice is mandatory rather than preferred, choose the Deepgram or Google conditional path instead.

### Finding 2 — Naturalness evidence favors Simba, but no benchmark replaces a script audition

**Prose.** **Verified:** current blind-listener evidence materially favors Simba 3.2 over the older OpenAI `tts-1` family and Chirp 3: HD. **Unresolved:** it does not establish how the current OpenAI `gpt-4o-mini-tts-2025-12-15` snapshot compares, because that model is absent from the current Artificial Analysis native-provider-voice table, and it does not establish the best Australian-accent narrator. Deepgram Aura-2 also lacks enough votes to appear in that Artificial Analysis table. A separate early Speech Arena table ranks Aura-2 twelfth of eighteen with 66 matches, while explicitly warning that its leaderboard is early; those scores are not comparable to Artificial Analysis Elo.

**Evidence.** Artificial Analysis describes its method as blind, same-text pairwise listening with higher Elo indicating listener preference for naturalness. Its current table places Simba 3.2 at 1,232, TTS-1 HD at 1,104, TTS-1 at 1,084, and Chirp 3: HD at 1,052. It also warns implicitly through confidence intervals and vote counts that rankings move as samples accumulate; Simba 3.2 was only released in July 2026 and had about 1,400 votes at capture. Artificial Analysis offers US/UK accent filters, not an Australian narration category. The NeurIPS 2025 EmergentTTS-Eval paper independently includes Aura-2 and `gpt-4o-mini-tts`, and demonstrates that performance varies by category, prompting, and voice; it is useful evidence against treating any one aggregate leaderboard as universally decisive. ([Artificial Analysis](https://artificialanalysis.ai/text-to-speech/leaderboard/provider-voice?tab=leaderboard), [Speech Arena](https://www.speecharena.org/leaderboard), [EmergentTTS-Eval](https://papers.neurips.cc/paper_files/paper/2025/file/04970a25af46918606ba2cf0a3d7905d-Paper-Datasets_and_Benchmarks_Track.pdf))

**Implication.** Treat the ranking as a shortlist prior, then run a controlled local bake-off with identical normalized text and loudness. Score at least naturalness, correct technical pronunciation, pacing across scene cuts, repeatability, and whether the voice still sounds credible after three continuous minutes. Do not select solely from a ten-second marketing demo.

### Finding 3 — Cost is negligible; signup friction and cash floor are the real differentiators

**Prose.** **Verified cost for 6,300–9,000 total characters:** Speechify is $0 inside its recurring 50,000-character free tier; Google Chirp 3: HD is $0 inside its 1,000,000-character monthly free-usage allowance but requires billing activation and a payment method; Deepgram is $0 while its $200 promotional credit remains, with no card required; OpenAI has no guaranteed current free monetary credit for the preferred models and ordinarily requires at least a $5 prepaid purchase. Nominal usage costs are tiny: Deepgram and out-of-free-tier Google are each $0.189–$0.270; OpenAI TTS-1 is $0.0945–$0.135; OpenAI TTS-1 HD is $0.189–$0.270. Speechify Starter’s overage-equivalent rate would be $0.063–$0.090, but entering that paid tier has a $10 monthly floor.

**Evidence.** Speechify Free includes 50,000 characters, no card, commercial use, and a hard cap. Its $10 Starter plan includes one million characters and then charges $10 per million. Deepgram Aura-2 is $0.030 per 1,000 characters, and the Pay-As-You-Go signup advertises a $200 credit, no minimum, no expiration for purchased PAYG credit, and no card; the promotional free credit has a one-year expiry in the project documentation. Google lists Chirp 3: HD with a 0–1 million-character free usage limit and $30 per million after the limit, while requiring billing to be enabled. Google’s new-customer trial supplies $300 for 90 days, requires a valid payment method and a small temporary authorization hold, and does not automatically charge until the account is manually upgraded. OpenAI lists $15 and $30 per million characters for TTS-1 and TTS-1 HD. `gpt-4o-mini-tts` charges $0.60 per million text tokens plus $12 per million audio-output tokens; OpenAI does not publish enough on that model page to calculate this exact nine-minute job from duration alone, so its exact nominal total remains unresolved. OpenAI prepaid billing currently has a $5 minimum purchase, a default $10 purchase, one-year expiry, and non-refundable credits; auto-recharge is enabled by default during setup unless disabled. ([Speechify pricing](https://speechify.ai/pricing), [Deepgram pricing](https://deepgram.com/pricing), [Google pricing](https://cloud.google.com/text-to-speech/pricing), [Google free trial](https://docs.cloud.google.com/free/docs/free-cloud-features), [OpenAI TTS-1](https://developers.openai.com/api/docs/models/tts-1), [OpenAI TTS-1 HD](https://developers.openai.com/api/docs/models/tts-1-hd), [OpenAI mini TTS](https://developers.openai.com/api/docs/models/gpt-4o-mini-tts), [OpenAI prepaid billing](https://help.openai.com/en/articles/8264778-what-is-prepaid-billing))

**Implication.** Price should not decide quality at this scale. The practical order for immediate experimentation is Speechify first, Deepgram second, Google third, and OpenAI fourth if no provider account already exists. Disable OpenAI auto-recharge if using it only for the demo, and set Google budget alerts if upgrading the trial account.

### Finding 4 — Deepgram Aura-2 is the strongest no-card Australian-accent alternative

**Prose.** **Conditional recommendation:** choose Deepgram Aura-2 when a recognizably Australian stock voice, Australian processing endpoint, no-card onboarding, or deterministic technical-name overrides matter more than native word timing and fine-grained acting direction. Two explicit Australian models are currently documented: `aura-2-theia-en` (feminine, expressive/polite/sincere, informative) and `aura-2-hyperion-en` (masculine, caring/warm/empathetic, interview). Aura-2’s new IPA override is the most directly auditable technical-pronunciation mechanism in the comparison.

**Evidence.** Deepgram’s current model catalog supports American, British, Australian, Irish, and Filipino English accents and identifies Theia and Hyperion as `en-au`. Runtime controls support speed from 0.7–1.5 and inline IPA overrides on both REST and WebSocket requests, with up to 500 overrides per request and warning headers reporting invalid IPA. Aura-2 accepts 2,000 characters per request, so this job needs about four to five sentence-safe chunks. REST can return WAV/Linear16, MP3, Ogg Opus, FLAC, or AAC; the response is streamed and can be written directly to a file. The AU endpoint `api.au.deepgram.com` supports TTS and routes processing through Australia. No Aura-2 TTS speech-mark or alignment response is documented; the only timestamps found are in Deepgram’s separate speech-to-text products. ([voices](https://developers.deepgram.com/docs/tts-models), [voice controls](https://developers.deepgram.com/docs/tts-voice-controls), [media output](https://developers.deepgram.com/docs/tts-media-output-settings), [regional endpoints](https://developers.deepgram.com/reference/custom-endpoints), [TTS getting started](https://developers.deepgram.com/docs/text-to-speech))

**Implication.** Use Theia and Hyperion as the first two audition candidates. Maintain an IPA map for “OpenAI,” the project/product name, acronyms, and unusual proper nouns, then inspect Deepgram’s pronunciation-warning headers. Generate audio before editing the Remotion timeline because caption or visual alignment will need manual cue points or a second forced-alignment/transcription pass.

### Finding 5 — Google Chirp 3: HD is the strongest privacy-first, high-control alternative if GCP setup is acceptable

**Prose.** **Conditional recommendation:** choose Google Chirp 3: HD when no prompt/audio logging, a broad en-AU stock-voice choice, or rich synchronous pronunciation and prosody controls outweigh card and Google Cloud setup friction. Chirp 3: HD itself is GA, while its current SSML layer is Preview/Pre-GA and unavailable for streaming requests. The product page currently lists 30 voice styles usable across supported locales, including `en-AU`.

**Evidence.** Google documents en-AU availability, 30 distinct voice styles, GA streaming and batch synthesis, and regional endpoints at `global`, `us`, `eu`, `asia-southeast1`, `europe-west2`, and `asia-northeast1`; there is no Australia endpoint. Batch output includes MP3, Ogg Opus, PCM/Linear16, A-law, and mu-law, while streaming supports Ogg Opus and telephony/PCM formats. Current Chirp documentation marks SSML as Preview, synchronous-only, and supports `<phoneme>` with IPA, `<sub>`, `<break>`, `<prosody>`, `<audio>`, and voice switching. Separate voice controls provide pace, pause markup, and IPA or X-SAMPA custom pronunciations, including in streaming for pronunciation controls. Google states Cloud TTS is stateless and resourceless and does not log customer TTS text or audio. The general Cloud terms retain customer IP in customer data, and the Cloud TTS basics page explicitly permits generated audio in applications and media such as videos, subject to law and the terms. ([Chirp 3 HD](https://docs.cloud.google.com/text-to-speech/docs/chirp3-hd), [voices](https://docs.cloud.google.com/text-to-speech/docs/list-voices-and-types), [data logging](https://docs.cloud.google.com/text-to-speech/docs/data-logging), [media use](https://docs.cloud.google.com/text-to-speech/docs/basics), [Cloud terms](https://cloud.google.com/terms/))

**Implication.** Google is the best choice if the narration script is sensitive and the project already has a Google Cloud billing account. For a fast hackathon workflow, it is less attractive than Speechify or Deepgram because credentials, project/billing activation, and preview feature behavior add setup risk. Pin the locale and voice name explicitly, use synchronous requests for SSML, and avoid assuming an en-AU voice implies Australian data residency.

### Finding 6 — OpenAI is attractive for instruction-driven delivery, but the current deprecation surfaces conflict

**Prose.** **Conditional recommendation:** use OpenAI when brand alignment with an OpenAI Build Week demo or natural-language acting instructions are decisive, especially if the account and credits already exist. `gpt-4o-mini-tts` can be prompted for accent, emotional range, intonation, impressions, speech speed, tone, and whispering, and OpenAI recommends `marin` or `cedar` for quality. However, the current documentation is internally inconsistent: the model page and API reference still present `gpt-4o-mini-tts` and snapshot `gpt-4o-mini-tts-2025-12-15` as available/default, while the all-model catalog labels the family “Deprecated.” The deprecations page only gives a 23 July 2026 shutdown date for the older `2025-03-20` snapshot and directs users to `2025-12-15`; it does not state the replacement or shutdown date for the whole family.

**Evidence.** The Speech API accepts up to 4,096 characters per request, offers 13 built-in English-optimized voices for the mini model, streams audio, supports MP3, Opus, AAC, FLAC, WAV, and PCM, and exposes speed from 0.25–4.0. It has no documented TTS SSML, pronunciation dictionary/IPA field, word timing, or speech-mark response. A stock Australian voice is not listed, though the guide says the mini model can be instructed to use an accent. The TTS-1 and TTS-1 HD families remain listed without a deprecation label, but do not accept the natural-language `instructions` field and have fewer voices. Australia is an officially supported API country. ([guide](https://developers.openai.com/api/docs/guides/text-to-speech), [API reference](https://platform.openai.com/docs/api-reference/audio/speech-audio-done-event?lang=curl), [mini model page](https://developers.openai.com/api/docs/models/gpt-4o-mini-tts), [all models](https://developers.openai.com/api/docs/models/all), [deprecations](https://developers.openai.com/api/docs/deprecations), [supported countries](https://developers.openai.com/api/docs/supported-countries))

**Implication.** Do not build a deadline-critical narration pipeline around the alias until the account confirms availability. If using OpenAI now, explicitly request `gpt-4o-mini-tts-2025-12-15`, generate and cache the final audio immediately, and keep a TTS-1 HD or external-provider fallback. For technical terms, respell or expand them in the source text and audition carefully because there is no documented phoneme override.

### Finding 7 — Speechify has the best native Remotion timing; every provider should still be used as a pre-render asset generator

**Prose.** **Verified:** all four providers can be called from Node/TypeScript or raw HTTP and can export a local audio file. **Implementation inference:** the robust Remotion architecture is to generate narration in a separate script before rendering, store immutable audio and timing metadata under the Remotion project’s `public/` directory, reference it with `staticFile()`, and render with Remotion’s current audio component. Do not call a cloud TTS API during frame rendering. Cache by a hash of normalized script, provider, model/snapshot, voice, locale, format, and control settings; retain a human-readable manifest with provider request IDs and source chunks.

**Evidence.** Remotion documents that `staticFile()` turns assets in `public/` into loadable URLs and that audio files can be referenced from there; the current audio documentation prefers `<Audio>` from `@remotion/media` for new usage. Speechify’s `/v1/audio/speech` response uniquely provides word-level start/end milliseconds and text character offsets. OpenAI, Deepgram Aura-2, and Chirp 3: HD do not document equivalent TTS word marks. Google’s general API supports timepoints for SSML `<mark>` tags, but Chirp 3’s current supported-SSML list omits `<mark>` and says unlisted tags are ignored; therefore Chirp 3 timepoints are unresolved and should not be designed into the baseline. ([Remotion `staticFile()`](https://www.remotion.dev/docs/staticfile), [Remotion audio](https://www.remotion.dev/docs/audio), [Speechify speech marks](https://docs.speechify.ai/docs/features/speech-marks), [Google Chirp SSML](https://docs.cloud.google.com/text-to-speech/docs/chirp3-hd), [Google generic timepoints](https://docs.cloud.google.com/text-to-speech/docs/reference/rest/v1beta1/text/synthesize))

**Implication.** Speechify can drive word highlights or precise caption frames directly after converting milliseconds to frames. For the other three, use coarse scene cues based on actual rendered duration or add a separate alignment pass. Committing generated narration assets makes the final video render repeatable even if a model alias, network, quota, or provider changes later.

### Finding 8 — Repeatability comes from explicit versioning and cached artifacts, not assumed byte determinism

**Prose.** **Verified:** each API permits explicit selection of model and voice; OpenAI additionally exposes a dated mini-TTS snapshot, and Speechify supports a dated `Speechify-Version` header for response-shape stability. **Unresolved:** none of the reviewed providers promises byte-identical audio for repeated generative TTS requests. Speechify’s Simba 3.2 and model catalog changed within days of this research, Google’s Chirp controls and SSML have changed through release notes, and Deepgram added speed/pronunciation controls in May 2026.

**Evidence.** OpenAI describes snapshots as locking a model version for consistent performance and behavior. Speechify recommends `GET /v1/audio/models` for runtime discovery and its SDK sends a build-date version header; this is evidence of a moving API/model surface, not a guarantee of identical waveforms. Deepgram identifies voice models by explicit names such as `aura-2-theia-en`. Google identifies voice and locale separately, with release notes recording expanding regions, speakers, and controls. ([OpenAI snapshot](https://developers.openai.com/api/docs/models/gpt-4o-mini-tts), [Speechify models](https://docs.speechify.ai/docs/get-started/models), [Speechify changelog](https://docs.speechify.ai/changelog), [Deepgram voices](https://developers.deepgram.com/docs/tts-models), [Google release notes](https://docs.cloud.google.com/text-to-speech/docs/release-notes))

**Implication.** The generation script should fail if any explicit provider/model/voice setting is missing, write final audio atomically, and never silently regenerate an approved take during `remotion render`. Keep the chosen audio file as the source of truth for timing and final export.

### Finding 9 — Public YouTube use is permitted, with provider-specific disclosure obligations

**Prose.** **Verified:** all four APIs support using generated speech in a public video under their current terms, provided the user owns the script/input, follows acceptable-use rules, and does not infringe third-party voice or content rights. OpenAI and Speechify explicitly require clear disclosure that the heard voice is AI-generated and not human. Deepgram’s Business Terms prohibit representing output as human-generated, which functionally requires truthful presentation. No Google Cloud TTS-specific AI-voice disclosure requirement was found. A single visible or spoken disclosure such as “Narration generated with AI text-to-speech” satisfies the safest common workflow without claiming a human narrator.

**Evidence.** OpenAI’s API guide explicitly requires end-user disclosure; its Services Agreement assigns output to the customer to the extent permitted by law, and its Service Terms discuss customer use/distribution of API output. Speechify’s free tier explicitly includes commercial use, its supplemental terms classify synthetic output as customer content, permit end users to download it through an integrated application, and require commercially reasonable AI-generated-voice disclosures on every use. Speechify attribution “Voices powered by Speechify” is required only upon Speechify’s request, rather than as an unconditional label. Deepgram grants use of output for lawful purposes on a royalty-free basis and forbids claiming it was human-generated. Google’s basics page expressly permits generated audio in videos and recordings, and its terms retain customer rights in customer data/derived data. No unconditional provider-logo attribution was found for OpenAI, Deepgram, or Google Cloud TTS. ([OpenAI guide](https://developers.openai.com/api/docs/guides/text-to-speech), [OpenAI Services Agreement](https://cdn.openai.com/osa/openai-services-agreement.pdf), [Speechify pricing](https://speechify.ai/pricing), [Speechify supplemental terms](https://speechify.com/terms-ai-voice-api/), [Deepgram Business Terms](https://deepgram.com/business/Business_TOS.pdf), [Google TTS basics](https://docs.cloud.google.com/text-to-speech/docs/basics), [Google Cloud terms](https://cloud.google.com/terms/))

**Implication.** Place a concise disclosure in the video end card, description, or both, and keep it even if the final provider is Google. Preserve the provider invoice/usage record, selected stock-voice ID, terms snapshot date, and final generated file with the project. Avoid implying that a named real person narrated the video.

### Finding 10 — Privacy ranking is Google first, OpenAI second, Deepgram opt-out conditional, Speechify unresolved

**Prose.** **Verified privacy order for an ordinary script:** Google has the clearest minimal-retention posture because Cloud TTS states it does not log customer text or audio. OpenAI API data is not used for training by default, while `/v1/audio/speech` has up to 30 days of abuse-monitoring retention, no application-state retention, and Zero Data Retention eligibility. Deepgram can be instructed with `mip_opt_out=true`, after which request data is retained only as long as needed to process it; however, the public Aura-2 prices are explicitly tied to opting into its Model Improvement Program, and the opt-out price was not published on the inspected pages. Speechify’s current general privacy policy permits use of information/User Content for research and improving algorithms and limited employee review, but the API supplemental terms do not state a TTS-input retention period or an API-specific default training opt-out.

**Evidence.** Google’s current data-logging page is direct and TTS-specific. OpenAI’s endpoint table lists the speech endpoint as 30-day abuse-monitoring retention, no application state, and ZDR eligible, while its business/API article says organizations are opted out of training by default. Deepgram describes MIP participation as voluntary, provides the opt-out query parameter, and states opted-out data is retained only for request processing; its pricing page says displayed rates opt into MIP. Speechify’s privacy policy covers site/software/services broadly, collects User Content, permits product/algorithm improvement, and allows review in listed circumstances; no more specific public API retention schedule was found. ([Google data logging](https://docs.cloud.google.com/text-to-speech/docs/data-logging), [OpenAI data controls](https://platform.openai.com/docs/models/default-usage-policies-by-endpoint), [OpenAI training defaults](https://help.openai.com/en/articles/5722486-api-data-usage-policies), [Deepgram MIP](https://developers.deepgram.com/docs/the-deepgram-model-improvement-partnership-program), [Deepgram pricing](https://deepgram.com/pricing), [Speechify privacy](https://speechify.com/privacy/), [Speechify supplemental terms](https://speechify.com/terms-ai-voice-api/))

**Implication.** For a public hackathon narration script, these differences are unlikely to be material. If the script includes confidential product details, use Google, an approved OpenAI ZDR organization, or Deepgram with confirmed opt-out pricing. Treat Speechify API privacy/training behavior as unresolved until support supplies a current DPA or API data-use statement.

### Finding 11 — Stock voices avoid cloning risk; custom-voice consent rules are strict

**Prose.** **Verified:** no voice cloning is needed for this task, and stock voices are the lower-risk choice. If cloning is introduced, explicit speaker consent is required across OpenAI, Speechify, and Google’s instant custom voice. Speechify adds the strictest inspected API limitations: the speaker must have directly consented in writing and cannot be under 18, deceased, or a current/former reasonably well-known political figure. OpenAI custom voices are limited to eligible customers and require a fixed consent recording plus a matching sample. Google requires a fixed consent phrase for Instant Custom Voice, including en-AU. Deepgram’s public Aura-2 path is a stock-voice catalog rather than a self-service cloning workflow.

**Evidence.** OpenAI’s guide specifies separate consent and sample recordings and gives an exact English consent sentence. Speechify’s supplemental terms define written consent and speaker eligibility and require disclosure on every synthetic-output use. Google’s Instant Custom Voice documentation provides the phrase “I am the owner of this voice and I consent…” for en-AU and other supported locales. Deepgram Aura-2 documentation exposes fixed named models and no public stock-to-custom cloning step. ([OpenAI custom voices](https://developers.openai.com/api/docs/guides/text-to-speech), [Speechify supplemental terms](https://speechify.com/terms-ai-voice-api/), [Google Instant Custom Voice](https://docs.cloud.google.com/text-to-speech/docs/chirp3-instant-custom-voice), [Deepgram voices](https://developers.deepgram.com/docs/tts-models))

**Implication.** Use a provider stock voice and record its exact ID in the project manifest. Do not clone a teammate, celebrity, judge, public figure, or synthetic voice for a deadline-driven demo without a separate rights review and retained written consent.

### Finding 12 — No provider documents an audio watermark, but absence cannot be guaranteed

**Prose.** **Unresolved:** none of the reviewed TTS API guides, output-format references, or terms documents describes an audible watermark, inaudible provenance watermark, C2PA attachment, or mandatory provider metadata embedded in the returned audio. This supports only the narrow conclusion “no provider-side audio watermark requirement was documented,” not “the files contain no detectable tag.”

**Evidence.** Each provider documents ordinary downloadable audio formats and response/file behavior: OpenAI MP3/Opus/AAC/FLAC/WAV/PCM, Speechify WAV/MP3/Ogg/AAC/PCM and streaming formats, Deepgram WAV/MP3/Ogg Opus/FLAC/AAC, and Google Chirp batch MP3/Ogg Opus/PCM plus streaming formats. Their inspected disclosure/attribution terms operate through user-facing disclosure rather than a documented embedded audio marker. ([OpenAI formats](https://developers.openai.com/api/docs/guides/text-to-speech), [Speechify speech API](https://docs.speechify.ai/tts/api-reference/text-to-speech/audio/speech), [Deepgram formats](https://developers.deepgram.com/docs/tts-media-output-settings), [Google Chirp formats](https://docs.cloud.google.com/text-to-speech/docs/chirp3-hd))

**Implication.** Preserve the original provider output and do not strip metadata merely to conceal AI generation. Apply the explicit disclosure independently of whether a watermark or tag can be detected.

### Finding 13 — Final ranking and decision rules

**Prose.** **Ranked recommendation (inference):** (1) SpeechifyAI Simba 3.2 as the overall baseline; (2) Deepgram Aura-2 as the no-card Australian-accent and IPA-control alternative; (3) Google Chirp 3: HD as the privacy-first, broad en-AU, high-control alternative for an existing GCP user; (4) OpenAI Speech as the brand-aligned, instruction-driven alternative once the current mini-TTS deprecation ambiguity is resolved. This rank optimizes the actual task rather than raw model prestige.

**Evidence.** Speechify uniquely combines the strongest current listener-preference evidence, recurring cardless free allowance, commercial rights, rich SSML, and word timing. Deepgram uniquely combines explicit Australian voices, an Australian endpoint, no-card $200 trial, and auditable IPA overrides. Google combines en-AU breadth, current GA Chirp core service, a million free characters, rich Preview controls, and the clearest TTS privacy statement, but requires billing/payment setup. OpenAI combines natural-language delivery instructions, good English voices, broad formats, and OpenAI-event relevance, but has no native marks, no stock en-AU voice, a $5 practical cash floor for new paid access, and contradictory current deprecation surfaces. ([Speechify pricing](https://speechify.ai/pricing), [Deepgram voices](https://developers.deepgram.com/docs/tts-models), [Google Chirp](https://docs.cloud.google.com/text-to-speech/docs/chirp3-hd), [OpenAI guide](https://developers.openai.com/api/docs/guides/text-to-speech), [OpenAI model catalog](https://developers.openai.com/api/docs/models/all))

**Implication.** Generate the same 30–45 second audition excerpt in Speechify Simba 3.2 first. Add Deepgram Theia and Hyperion if Australian identity matters, Google en-AU if privacy or pronunciation control matters, and OpenAI Cedar/Marin only if brand/style direction justifies its ambiguity and payment friction. Approve one voice, generate the full audio once, review the complete three minutes, then freeze the artifact before Remotion timing work.

## Notes

- **OpenAI deprecation caveat:** the all-model catalog labels the `gpt-4o-mini-tts` family Deprecated, while the API reference and model page still list the alias and `2025-12-15` snapshot and the deprecations schedule only retires `2025-03-20` on 23 July 2026. Family shutdown timing and replacement are unsupported by a single consistent official statement.
- **Speechify model-doc caveat:** current Speechify model pages disagree on whether Simba 3.2 accepts only curated voices or also manually approved cloned voices. This does not affect the stock-voice baseline; use `GET /v1/audio/models` and support confirmation before a cloning design.
- **Speechify Australian caveat:** the reviewed Simba 3.2 documentation lists English and eight curated IDs, but no locale/accent metadata proving an Australian stock voice. The broader “1,500+ voices” claim applies to the platform catalog, not necessarily the Simba 3.2 allow-list.
- **Google documentation caveat:** the dedicated Chirp 3 page is current and documents Preview SSML/voice controls, while a broader voice-list page still contains an older note saying Chirp 3 does not support SSML or speaking-rate controls. The dated Chirp page and release notes were treated as controlling.
- **Google timepoint caveat:** generic Cloud TTS supports `<mark>` timepoints, but Chirp 3’s current supported-tag list omits `<mark>` and says unsupported tags are ignored. Chirp-native timing remains unsupported for planning purposes.
- **Deepgram privacy-price caveat:** the public $0.030/1,000-character Aura-2 rate opts into MIP. `mip_opt_out=true` is documented, but its exact price impact is unresolved.
- **Watermark caveat:** absence of public documentation is not a forensic guarantee that an output has no inaudible provenance signal or metadata.
- **Benchmark caveat:** current leaderboards emphasize short listener-preference comparisons and US/UK categories. They do not directly score this three-minute script, Australian authenticity, technical-name accuracy, disclosure compliance, or Remotion timing convenience.

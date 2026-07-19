# ElevenLabs Deep Dive for a Remotion Build Week Demo

## Assignment

**Goal**

Determine whether ElevenLabs is the best-overall polished text-to-speech option for a public, prize-eligible OpenAI Build Week YouTube demonstration of less than three minutes, using an approximately 350–500-word English technical narration assembled in Remotion on macOS. Establish an operational model/voice choice, cost, rights posture, privacy posture, and proof gate rather than only describing product features.

**Scope**

Current ElevenLabs speech-synthesis models, default and library voices, browser/Studio/API workflows, audio and timing outputs, pronunciation and delivery control, pricing and subscription behavior, public/commercial use rights, Build Week compatibility, privacy and retention, watermarking, cloning consent, macOS practicality, Remotion integration, independent naturalness evidence, and the boundary between automation and required human audition. Sources were inspected on 19 July 2026.

**Exclusions**

No production audio was generated; no account was created; no subjective voice was selected on the user's behalf; no legal opinion is offered; no comparison synthesis across other TTS providers is attempted; and no source or application files were edited.

## Sources

- ElevenLabs pricing and billing: [pricing](https://elevenlabs.io/pricing), [billing guide](https://elevenlabs.io/docs/overview/administration/billing), [credits](https://help.elevenlabs.io/hc/en-us/articles/27562020846481-What-are-credits), [accepted payment methods](https://help.elevenlabs.io/hc/en-us/articles/13416538053905-What-kind-of-payment-is-accepted), [Apple Pay on macOS](https://help.elevenlabs.io/hc/en-us/articles/15308826484753-How-can-I-use-Apple-Pay), and [Google Pay on macOS](https://help.elevenlabs.io/hc/en-us/articles/15308818814481-How-can-I-use-Google-Pay).
- ElevenLabs models and synthesis controls: [models](https://elevenlabs.io/docs/overview/models), [model differences](https://help.elevenlabs.io/hc/en-us/articles/17883183930129-What-models-do-you-offer-and-what-is-the-difference-between-them), [Text to Speech](https://elevenlabs.io/docs/overview/capabilities/text-to-speech), [product guide](https://elevenlabs.io/docs/eleven-creative/playground/text-to-speech), [best practices](https://elevenlabs.io/docs/overview/capabilities/text-to-speech/best-practices), [Eleven v3 prompting](https://elevenlabs.io/docs/best-practices/prompting), and [Eleven v3 general availability](https://elevenlabs.io/blog/eleven-v3-is-now-generally-available).
- ElevenLabs voices: [voice overview](https://elevenlabs.io/docs/overview/capabilities/voices), [current default voices and replacements](https://help.elevenlabs.io/hc/en-us/articles/26942950589969-What-are-Default-voices), [default voice access and API IDs](https://help.elevenlabs.io/hc/en-us/articles/25844757988753-How-do-I-access-ElevenLabs-Default-voices), [Voice Library](https://help.elevenlabs.io/hc/en-us/articles/23143350045329-What-is-the-Voice-Library), [Voice Library guide and commercial-use statement](https://elevenlabs.io/elevenlabs-voices-a-comprehensive-guide), and [technical-assistance voice examples](https://elevenlabs.io/docs/eleven-agents/customization/voice/best-practices/conversational-voice-design).
- ElevenLabs API and export: [API quickstart](https://elevenlabs.io/docs/eleven-api/quickstart), [API introduction and character-cost header](https://elevenlabs.io/docs/api-reference/introduction), [Create speech](https://elevenlabs.io/docs/api-reference/text-to-speech/convert), [Create speech with timing](https://elevenlabs.io/docs/api-reference/text-to-speech/convert-with-timestamps), [streaming](https://elevenlabs.io/docs/api-reference/streaming), [save/stream a file](https://elevenlabs.io/docs/eleven-api/guides/how-to/text-to-speech/streaming), [request stitching](https://elevenlabs.io/docs/eleven-api/guides/how-to/text-to-speech/request-stitching), and [browser history download](https://help.elevenlabs.io/hc/en-us/articles/14129286847505-How-do-I-download-generated-files-from-Text-to-Speech).
- ElevenLabs pronunciation: [pronunciation dictionaries](https://elevenlabs.io/docs/eleven-api/guides/how-to/text-to-speech/pronunciation-dictionaries), [create dictionary from rules](https://elevenlabs.io/docs/api-reference/pronunciation-dictionaries/create-from-rules), and [Studio pronunciations editor](https://help.elevenlabs.io/hc/en-us/articles/37896325858065-How-do-I-use-the-Pronunciations-Editor-in-Studio).
- ElevenLabs rights and safety: [publication/commercial-use help](https://help.elevenlabs.io/hc/en-us/articles/13313564601361-Can-I-publish-the-content-I-generate-on-the-platform), [Terms of Service for users outside the EEA/UK/Switzerland](https://elevenlabs.io/terms-of-use), [Prohibited Use Policy](https://elevenlabs.io/use-policy), [Voice Library Addendum](https://elevenlabs.io/vla), [content after cancellation](https://help.elevenlabs.io/hc/en-us/articles/15993008593297-What-happens-to-my-content-after-my-subscription-ends), and [subscription cancellation](https://help.elevenlabs.io/hc/en-us/articles/13313868884241-How-do-I-cancel-my-subscription).
- ElevenLabs privacy and provenance: [Privacy Policy](https://elevenlabs.io/privacy-policy), [training-data opt-out](https://help.elevenlabs.io/hc/en-us/articles/29952728805393-Is-my-data-used-to-improve-ElevenLabs-AI-models), [Zero Retention Mode](https://elevenlabs.io/docs/eleven-api/resources/zero-retention-mode), [SynthID announcement](https://elevenlabs.io/blog/synthid), and [Audio Detector/watermark coverage](https://elevenlabs.io/docs/eleven-creative/audio-tools/audio-detector).
- ElevenLabs geography: [Australia and New Zealand expansion](https://elevenlabs.io/blog/elevenlabs-expands-presence-in-australia-new-zealand).
- OpenAI Build Week: [official rules](https://openai.devpost.com/rules).
- Remotion: [`<Audio>` from `@remotion/media`](https://www.remotion.dev/docs/media/audio), [`staticFile()`](https://www.remotion.dev/docs/staticfile), [`@remotion/elevenlabs`](https://www.remotion.dev/docs/elevenlabs/), and [`elevenLabsTranscriptToCaptions()`](https://www.remotion.dev/docs/elevenlabs/elevenlabs-transcript-to-captions).
- Independent quality evidence: [Artificial Analysis Controlled Voice Leaderboard](https://artificialanalysis.ai/text-to-speech/leaderboard/controlled-voice), [Speech Arena](https://artificialanalysis.ai/text-to-speech/arena), and [benchmarking methodology](https://artificialanalysis.ai/text-to-speech/methodology).

## Findings

### Recommendation: Starter, Eleven v3, and a Default Voice Is the Lowest-Risk Polished Path

ElevenLabs is a strong best-overall candidate for this one-off demo because it combines top-tier listening quality, a no-code audition surface, a scriptable TypeScript API, character-level timing, usable 44.1 kHz MP3 on the lowest paid tier, and a clear paid commercial license. The recommended production route is one month of **Starter**, generate the final audio while the paid subscription is active, use **`eleven_v3`** first, and fall back to **`eleven_multilingual_v2`** if v3 is too variable or over-performs the script. Use a current **Default voice**, not a clone or a custom-rate community voice, unless a short audition proves a library voice materially better.

The first audition slate should be **Elara — Crisp Pro Narrator**, **Finley — Articulate Anchor**, **Alicia — Polished Global Anchor**, **Lawrence — Bright and Informative**, and **Caleb — Trusted Guide**. If an Australian delivery is desirable, include **Baxter — Dry Calm Aussie** and the library voice **Stuart** (`HDA9tsk27wYi3uq0fPcK`), which ElevenLabs describes as professional, friendly, Australian, and suitable for technical assistance. These descriptions identify plausible candidates; they do not establish which voice fits the actual visuals or script.

**Evidence**

- Starter is currently USD $6/month, includes 30,000 credits, a commercial license, Instant Voice Cloning, and 20 Studio projects; Starter TTS quality is listed as 128 kbps at 44.1 kHz ([pricing](https://elevenlabs.io/pricing)).
- ElevenLabs describes v3 as its most expressive model and Multilingual v2 as its stable, lifelike long-form model; the selection guide points content creation/video narration toward Multilingual v2 when high fidelity and stability dominate ([models](https://elevenlabs.io/docs/overview/models)).
- Eleven v3 left alpha and became generally available on 2 February 2026, removing the earlier beta/alpha commercial-use concern for this model ([GA announcement](https://elevenlabs.io/blog/eleven-v3-is-now-generally-available)).
- The official Default-voice replacement list supplies the candidate descriptions and says the new voices target realistic, stable long-form and real-time performance ([Default voices](https://help.elevenlabs.io/hc/en-us/articles/26942950589969-What-are-Default-voices)). Stuart's specific ID and technical-assistance description are in the official voice guide ([voice guide](https://elevenlabs.io/docs/eleven-agents/customization/voice/best-practices/conversational-voice-design)).

**Implication**

Starter is sufficient for the final deliverable; Creator, Pro, cloning, and streaming are unnecessary unless the proof gate exposes a specific quality requirement. Voice choice remains the largest subjective variable, so the user must audition against the real script and visuals before Codex locks the asset pipeline.

### Independent Listening Evidence Supports “Top Tier,” Rather Than an Unqualified “Best” Claim

Independent evidence supports Eleven v3 as a leading naturalness/prosody system, but does not prove it is the single best provider for this exact technical narration. Artificial Analysis currently ranks Eleven v3 **second** in its Controlled Voice Arena at **1,088 Elo**, with **2,009 samples**, behind Cartesia Sonic 3.5 at 1,122. Multilingual v2 ranks eleventh at 1,021; Flash v2.5 ranks eighteenth at 971. This is unusually relevant evidence because the controlled benchmark uses the same eight cloned US/UK voices across models, reducing the confound where one provider simply offers a more appealing stock voice.

**Evidence**

- The [Controlled Voice Leaderboard](https://artificialanalysis.ai/text-to-speech/leaderboard/controlled-voice) reports Eleven v3 at rank 2 and identifies the sample count, Elo, release date, and price basis.
- The [methodology](https://artificialanalysis.ai/text-to-speech/methodology) says quality Elo comes from listener preferences; the controlled test standardizes eight professionally recorded voices to compare naturalness, audio quality, pronunciation, pacing, and prosody. It also states that Artificial Analysis receives no provider compensation for listing or favorable outcomes.
- The [Speech Arena](https://artificialanalysis.ai/text-to-speech/arena) presents blinded pairwise listening, requiring listeners to hear both clips before voting.

**Implication**

Eleven v3 is defensibly described as top-tier and polished. The benchmark does not test this script, accent, stock voice, technical vocabulary, or Remotion mix, so a user listening gate remains mandatory. The correct downstream wording is “leading option with strong independent listening evidence,” not “objectively the most natural TTS.”

### Model Choice Is a Controlled Tradeoff Between Performance and Repeatability

For this narration, `eleven_v3` is the quality-first candidate. Its 5,000-character per-request limit is likely enough for a 350–500-word English script, it supports 74 languages, expressive audio tags, punctuation-driven delivery, and native IPA. Its main production liabilities are greater variance, no conventional speed slider, no SSML break support, and no request stitching. `eleven_multilingual_v2` is the conservative fallback: 10,000 characters per request, long-form stability, a speed control, conventional break handling, and request stitching. `eleven_flash_v2_5` is optimized for approximately 75 ms latency and lower API cost, which solves a problem this offline render does not have.

**Evidence**

- Current model IDs, language coverage, limits, and selection guidance are documented in [Models](https://elevenlabs.io/docs/overview/models) and [model differences](https://help.elevenlabs.io/hc/en-us/articles/17883183930129-What-models-do-you-offer-and-what-is-the-difference-between-them).
- v3 supports audio tags, punctuation, capitalization, and Natural/Creative/Robust stability modes; the v3 guide warns that output depends strongly on the selected voice and recommends prompts longer than 250 characters ([v3 prompting](https://elevenlabs.io/docs/best-practices/prompting)).
- The product guide states that the system is nondeterministic even with the same voice, settings, and model. Multilingual v2 settings include speed from 0.7 to 1.2, whereas v3 does not expose speed, Similarity, or Speaker Boost in the same way ([product guide](https://elevenlabs.io/docs/eleven-creative/playground/text-to-speech)).
- Request stitching is available for maintaining prosody across chunks but explicitly unavailable for v3 ([request stitching](https://elevenlabs.io/docs/eleven-api/guides/how-to/text-to-speech/request-stitching)).

**Implication**

Start with a single full-script v3 request so prosody does not cross chunk boundaries. If v3 cannot land the timing or pronunciation after two purposeful iterations, switch the same shortlisted voice to Multilingual v2 rather than burning credits on uncontrolled retries. The final rendered video should consume a frozen local audio asset, never regenerate TTS during a Remotion render.

### Exact Billing Is Character-Based; Starter’s Invoice Cost Is USD $6 Plus Tax

Let **`C`** be the final API-reported character cost for one complete narration. Using the conservative, documented one-credit-per-character rate for v3/Multilingual v2, one full pass consumes **`C` credits** and three complete passes consume **`3C` credits**. ElevenLabs exposes the exact `character-cost` response header, so Codex can record the actual charge for each generation instead of guessing from word count.

For planning only, assuming an average of roughly six billable characters per English word (letters plus spaces/punctuation), a 350–500-word script is approximately **2,100–3,000 credits per full pass** and **6,300–9,000 credits for three full passes**. Both are safely within Starter's 30,000-credit quota. At Starter's quota allocation, one pass represents **USD $0.42–$0.60** of included quota and three passes **$1.26–$1.80**, but ElevenLabs does not invoice those amounts separately. If this is the only use, the exact cash outlay for either one pass or three passes is the same: **USD $6 for one Starter month, plus taxes/levies/duties**. The account retains 21,000–23,700 credits after the estimated three-pass budget.

**Evidence**

- [Pricing](https://elevenlabs.io/pricing) lists Free at 10,000 credits, Starter at USD $6/30,000 credits, approximate Starter output of 30 minutes, and one credit per character as the planning rate. It also says credits are charged per generation request, not per download, and that limited free regenerations may be offered only when content and certain settings do not change.
- [API introduction](https://elevenlabs.io/docs/api-reference/introduction) documents access to the exact `character-cost` and request ID headers through the raw SDK response.
- [Credits](https://help.elevenlabs.io/hc/en-us/articles/27562020846481-What-are-credits) warns that some community voices have custom-rate multipliers; Default voices avoid that uncertainty.

**Implication**

Budget three paid full-script generations and do not assume a regeneration will be free. The final generation script should log `character-cost`, model, voice ID, settings, seed, and request ID. The exact pre-generation credit total cannot be inferred from “350–500 words”; only the final text or API response makes it exact.

### Free Is Large Enough for Audition but Ineligible for the Prize Submission

The Free plan's 10,000 monthly credits could technically cover an estimated three full passes. It should nevertheless be used only for discarded auditions. Free output lacks a commercial license, shared Voice Library voices are unavailable through the API to free users, publicly shared free-plan output requires attribution in the **title**, and the Prohibited Use Policy expressly forbids a Free User from using the service for commercial purposes including contests or sweepstakes. Build Week is a prize hackathon, and its submitted demo must be public on YouTube. Upgrading after creating free audio does not retroactively grant commercial rights to that audio.

**Evidence**

- ElevenLabs says Free output is non-commercial and, when published, must include `elevenlabs.io` or `11.ai` in the title; content generated during a paid subscription may be used commercially and indefinitely. Audio generated outside a paid subscription cannot later be used commercially ([publishing help](https://help.elevenlabs.io/hc/en-us/articles/13313564601361-Can-I-publish-the-content-I-generate-on-the-platform)).
- The [Prohibited Use Policy](https://elevenlabs.io/use-policy) includes contests and sweepstakes in its examples of commercial purposes forbidden to Free Users.
- The Build Week rules require a demonstration video of less than three minutes, clear audio explaining the project and Codex/GPT-5.6 use, a publicly visible YouTube upload, authorization for third-party APIs/SDKs/data, and permission for copyrighted or third-party materials ([official rules, Submission Requirements](https://openai.devpost.com/rules)).
- Free users cannot access Voice Library voices through the API ([voices overview](https://elevenlabs.io/docs/overview/capabilities/voices)).

**Implication**

Audition on Free if desired, then subscribe to Starter **before** generating every clip that appears in the submission. Regenerate even an identical winning free-plan audition after the paid plan is active. This is the cleanest path for both ElevenLabs' contest restriction and Build Week's third-party authorization requirement.

### Paid Output Has Durable YouTube Rights, but the Entrant Retains Compliance Responsibility

As between ElevenLabs and the user, the user retains rights in Output; a paid user may use the service commercially, subject to the Terms and Prohibited Use Policy. ElevenLabs' help center states that paid-period output remains commercially usable forever after cancellation. The Voice Library Addendum also says outputs already generated with a shared voice remain available for use after that voice's notice period ends. These terms support a public YouTube demo and later portfolio use. They do not clear rights in the user's input script, third-party trademarks, background music, screenshots, or an unauthorized cloned voice; Build Week independently requires the submission to be original/owned and third-party material to be permitted.

**Evidence**

- The [Terms of Service](https://elevenlabs.io/terms-of-use) state that paid users may use Services commercially and retain rights in Output as between the parties, while requiring necessary rights to Input and compliance with the Prohibited Use Policy.
- ElevenLabs confirms that the commercial license for paid-period generations survives subscription cancellation forever, while hosted content availability after cancellation is not guaranteed indefinitely ([content after subscription](https://help.elevenlabs.io/hc/en-us/articles/15993008593297-What-happens-to-my-content-after-my-subscription-ends)).
- The [Voice Library Addendum](https://elevenlabs.io/vla) preserves already-generated outputs when a shared model is removed, and the official voice guide says Voice Library voices include a commercial-use license ([voice guide](https://elevenlabs.io/elevenlabs-voices-a-comprehensive-guide)).
- Build Week requires authorization for third-party integrations, ownership or permission for submission components, and permission for copyrighted music or other material ([official rules](https://openai.devpost.com/rules)).

**Implication**

Download and retain the final MP3, alignment JSON, script, settings, invoice/receipt, and a dated copy or link to the applicable terms. A Default voice has a simpler provenance story than a clone. The paid ElevenLabs license solves TTS-output permission; it does not solve unrelated submission rights.

### Signup and macOS Payment Are Straightforward; “No Card Required” Is Not Explicitly Documented

New accounts are automatically assigned to Free, which supports a low-friction browser audition. Official sources do not explicitly say “no credit card required,” so that phrase should not be represented as verified. To move to Starter, ElevenLabs accepts credit card, Apple Pay, and Google Pay. On a Mac, Apple Pay checkout requires Safari; Google Pay checkout requires Chrome and an eligible stored card. ElevenLabs has active Australian and New Zealand operations and reports more than 750,000 users in those countries, so Australia is not a practical availability concern.

**Evidence**

- [Billing](https://elevenlabs.io/docs/overview/administration/billing) and [discounted/free plans](https://help.elevenlabs.io/hc/en-us/articles/13315218812177-Do-you-offer-discounted-or-free-plans) say signups are automatically assigned to Free.
- Accepted methods are credit card, Apple Pay, Google Pay, and UPI in India ([payment methods](https://help.elevenlabs.io/hc/en-us/articles/13416538053905-What-kind-of-payment-is-accepted)).
- Apple Pay requires Safari on macOS, whereas Google Pay requires Chrome and can be used on a Mac ([Apple Pay](https://help.elevenlabs.io/hc/en-us/articles/15308826484753-How-can-I-use-Apple-Pay), [Google Pay](https://help.elevenlabs.io/hc/en-us/articles/15308818814481-How-can-I-use-Google-Pay)).
- ElevenLabs announced an operational expansion across Australia/New Zealand and quantified existing regional usage ([regional announcement](https://elevenlabs.io/blog/elevenlabs-expands-presence-in-australia-new-zealand)).

**Implication**

The user must create the account, complete payment, and accept current terms. Codex can guide or automate local API setup after the user stores the key, but it should not handle the payment credential or expose the key in Remotion/browser code.

### Browser, Studio, and API Workflows All Work on macOS

The fastest audition path is ElevenCreative's browser Text to Speech playground: paste text, choose a voice/model, adjust controls, generate, and download from History as MP3 (128 kbps) or WAV. Studio is appropriate if the user wants paragraph-level overrides, a pronunciation editor, generation history, and project export. The most reproducible path after selection is the Node/TypeScript SDK or direct REST endpoint, with `ELEVENLABS_API_KEY` held in a local environment variable and an explicit voice ID, model ID, output format, seed, and settings.

The official Node library is `@elevenlabs/elevenlabs-js`; the REST API works from any language. A macOS machine needs only a current browser and Node/Python for file generation. Optional local playback helpers may request `mpv` or `ffmpeg`, but neither is required merely to receive and save the API response.

**Evidence**

- The [API quickstart](https://elevenlabs.io/docs/eleven-api/quickstart) documents dashboard API-key creation, `.env` storage, Python and TypeScript examples, v3 selection, and `mp3_44100_128` output.
- The [API introduction](https://elevenlabs.io/docs/api-reference/introduction) identifies official Python and Node libraries and direct HTTP/WebSocket access.
- Browser History can download MP3 128 kbps or WAV ([download help](https://help.elevenlabs.io/hc/en-us/articles/14129286847505-How-do-I-download-generated-files-from-Text-to-Speech)); Studio supports per-selection settings and pronunciation dictionaries ([Studio](https://elevenlabs.io/docs/eleven-creative/products/studio)).
- Remotion warns against browser-side ElevenLabs calls because an API key would be exposed ([Remotion ElevenLabs caption helper](https://www.remotion.dev/docs/elevenlabs/elevenlabs-transcript-to-captions)).

**Implication**

Use the browser for audition and a one-time Node pre-generation script for final assets. Keep API generation outside React and outside the render process. Commit or otherwise archive the selected audio and timing data, while keeping the API key out of the repository.

### Starter’s 44.1 kHz/128 kbps MP3 Is Sufficient for YouTube and Remotion

ElevenLabs can return MP3, PCM, μ-law, A-law, and Opus across multiple rates; WAV/PCM availability depends on tier and interface. Starter's clearly advertised quality is `mp3_44100_128`, which Remotion accepts directly and is adequate for spoken narration in a compressed YouTube delivery. Pro is necessary only if the workflow specifically requires 44.1 kHz PCM via API or the highest advertised 192 kbps output. The API reference and current pricing table conflict on whether 192 kbps starts at Creator or Pro; that conflict is immaterial to the Starter recommendation because 128 kbps is the selected format.

**Evidence**

- The Text to Speech capability page lists MP3 from 22.05–44.1 kHz, PCM, μ-law, A-law, and Opus, with high-quality options restricted by paid tier ([formats](https://elevenlabs.io/docs/overview/capabilities/text-to-speech)).
- The pricing table gives Starter 128 kbps at 44.1 kHz and Pro 128/192 kbps, while the endpoint reference says `mp3_44100_128` is the default and 44.1 kHz PCM/WAV requires Pro ([pricing](https://elevenlabs.io/pricing), [Create speech](https://elevenlabs.io/docs/api-reference/text-to-speech/convert)).
- Remotion's recommended [`<Audio>`](https://www.remotion.dev/docs/media/audio) accepts a local MP3 and extracts exact audio into the render timeline.

**Implication**

Generate `mp3_44100_128`; avoid upgrading for an inaudible specification benefit in a voiceover mixed into a YouTube video. If the final mix exposes compression artifacts, test a browser-downloaded WAV or Pro PCM as a targeted exception rather than the default plan.

### Pronunciation Is Strong but Requires Script Normalization and a Technical-Term Test

Technical narration should be normalized before synthesis: spell out ambiguous digits, currencies, URLs, symbols, version strings, code tokens, and acronyms in the spoken script. For v3, native IPA can be inserted between forward slashes, and pronunciation dictionaries support IPA/CMU phoneme rules; ElevenLabs reports 80–90% pronunciation consistency, not certainty. For Multilingual v2, use alias rules such as replacing an acronym with its spoken expansion, because phoneme rules are ignored by that model. Dictionaries can be attached by ID/version to API requests, up to three per request.

v3 does not support SSML `<break>` tags. Its pacing and emphasis come from punctuation, capitalization, natural sentence structure, and audio tags such as `[curious]` or `[excited]`; excessive direction is risky for a professional technical read. Multilingual v2 can use `<break time="x.xs" />` up to three seconds, a speed setting, stability/similarity controls, and alias dictionaries. ElevenLabs' common baseline for non-v3 models is stability 0.5, similarity 0.75, style 0, speed 1.0.

**Evidence**

- [Best practices](https://elevenlabs.io/docs/overview/capabilities/text-to-speech/best-practices) documents v3 IPA syntax and its reported 80–90% consistency, v3's lack of SSML breaks, pause alternatives, speed ranges for applicable models, acronym/URL normalization, and Multilingual v2 alias substitutions.
- [Pronunciation dictionaries](https://elevenlabs.io/docs/eleven-api/guides/how-to/text-to-speech/pronunciation-dictionaries) distinguish IPA/CMU phoneme support on v3/Flash v2 from alias support on other models.
- [Create speech](https://elevenlabs.io/docs/api-reference/text-to-speech/convert) accepts up to three dictionary locators and an `apply_text_normalization` mode.
- [v3 prompting](https://elevenlabs.io/docs/best-practices/prompting) documents audio tags, capitalization, punctuation, voice dependence, and the Natural/Creative/Robust stability modes.

**Implication**

Codex can produce a separate display script and spoken-normalized script, construct alias/IPA rules, and flag technical terms. The user must confirm the audible result. Use v3 Natural as the first setting; reserve tags for one or two intentional moments, and switch to Multilingual v2 if precise pauses and speed matter more than expressiveness.

### Character Alignment Makes Remotion Synchronization Direct

The `/with-timestamps` endpoint returns base64 audio plus character-level start/end times for both original and normalized text. Codex can decode the MP3, group characters into words or caption phrases, and convert seconds to Remotion frames with `Math.round(seconds * fps)`. This supports exact caption reveals, screen-callout timing, and scene boundaries without running a second transcription model.

Remotion's `@remotion/elevenlabs` package is specifically an adapter for **ElevenLabs Speech to Text**, not Text to Speech. It can convert Scribe word timestamps into Remotion captions, but a TTS-to-STT round trip is unnecessary when the TTS timing endpoint already supplies alignment. The direct alignment JSON needs a small custom converter to Remotion's caption shape.

**Evidence**

- [Create speech with timing](https://elevenlabs.io/docs/api-reference/text-to-speech/convert-with-timestamps) returns `audio_base64`, original `alignment`, and `normalized_alignment`, each containing characters and per-character start/end seconds.
- [`@remotion/elevenlabs`](https://www.remotion.dev/docs/elevenlabs/) explicitly converts ElevenLabs Speech to Text output, and [`elevenLabsTranscriptToCaptions()`](https://www.remotion.dev/docs/elevenlabs/elevenlabs-transcript-to-captions) requires Scribe's word-level timing array.
- [`<Audio>`](https://www.remotion.dev/docs/media/audio) is Remotion's recommended audio component and keeps extracted audio aligned with the timeline; [`staticFile()`](https://www.remotion.dev/docs/staticfile) addresses assets in the `public/` folder.

**Implication**

Generate `voiceover.mp3` and `voiceover-alignment.json` together before rendering. Use normalized alignment when displayed captions should match what was spoken; retain original alignment for audit. Avoid an extra 990-credit three-minute Scribe pass unless the alignment endpoint proves inadequate.

### Streaming Is Available but Adds No Value to an Offline Three-Minute Render

ElevenLabs supports chunked HTTP streaming and real-time WebSocket input, and official SDKs can either stream to playback/memory or write chunks to a file. These features reduce time to first audio for interactive applications. A Remotion demo render benefits more from complete audio, exact duration, and timing metadata than from low latency.

**Evidence**

- [Streaming API](https://elevenlabs.io/docs/api-reference/streaming) describes chunked raw audio bytes and official Node/Python helpers.
- The [streaming/file guide](https://elevenlabs.io/docs/eleven-api/guides/how-to/text-to-speech/streaming) shows both local MP3 output and in-memory streaming.
- The timing endpoint returns complete audio and alignment in one response ([with timestamps](https://elevenlabs.io/docs/api-reference/text-to-speech/convert-with-timestamps)).

**Implication**

Use non-streaming `with-timestamps` generation for the final. Streaming remains a fallback only if the endpoint or SDK has a practical response-size problem, which is unlikely for a sub-three-minute MP3.

### Generation Is Nondeterministic; Freeze the Selected File

ElevenLabs offers a `seed` that makes a best effort at deterministic sampling, but explicitly does not guarantee determinism. The product UI similarly says identical voice/model/settings can produce different output. Higher stability reduces variance but may sound monotonous. Re-rendering a Remotion composition must therefore reuse a selected audio file rather than invoke ElevenLabs again.

**Evidence**

- The [timing endpoint](https://elevenlabs.io/docs/api-reference/text-to-speech/convert-with-timestamps) accepts a seed from 0 to 4,294,967,295 and says deterministic output is best effort only.
- The [product guide](https://elevenlabs.io/docs/eleven-creative/playground/text-to-speech) states that settings define a range of randomization and do not guarantee identical output.
- ElevenLabs support says each generation can differ, especially at low stability ([voice variability](https://help.elevenlabs.io/hc/en-us/articles/13416017389329-Why-is-my-voice-monotonous-too-chaotic-doesn-t-sound-similar-etc)).

**Implication**

Pin all inputs for traceability but treat the MP3 as the reproducible artifact. A CI/build step may verify that the asset exists and matches a checksum; it should not regenerate it automatically.

### Privacy Requires an Opt-Out Before Generation; Zero Retention Is Enterprise-Only

By default, ElevenLabs may use submitted data to improve audio models. Any user can disable **Terms and privacy → Data use → Improve the models for everyone**; the opt-out applies to new data submitted after the setting is disabled. Standard service use still involves retention for providing, securing, and troubleshooting the service. Zero Retention Mode is an Enterprise option for supported products; setting `enable_logging=false` invokes it only for eligible Enterprise accounts and removes history/request-stitching functionality.

For a public-bound demo script, ordinary retention may be acceptable, but the user should opt out before uploading unpublished product details and should avoid secrets, personal data, credentials, or private customer information. Local copies of output should be treated as authoritative because ElevenLabs does not guarantee hosted history will remain accessible forever after cancellation.

**Evidence**

- [Training-data help](https://help.elevenlabs.io/hc/en-us/articles/29952728805393-Is-my-data-used-to-improve-ElevenLabs-AI-models) says certain submitted data is used to improve models, gives the account toggle path, and says disabling it prevents new submissions from training use.
- The [Privacy Policy](https://elevenlabs.io/privacy-policy) covers Inputs, Outputs, Voice Data, model research/development, moderation, deletion requests, and the prospective training opt-out.
- [Zero Retention Mode](https://elevenlabs.io/docs/eleven-api/resources/zero-retention-mode) is Enterprise-only and immediately deletes most request/response content for covered products; standard mode retains data according to policy.
- The [timing endpoint](https://elevenlabs.io/docs/api-reference/text-to-speech/convert-with-timestamps) documents the `enable_logging` behavior and loss of history features under zero retention.

**Implication**

The user should disable model improvement before the proof generation and leave it disabled through final generation. Codex can verify the intended local configuration and avoid embedding secrets, but only the signed-in user can confirm the account toggle. Do not represent Starter as zero-retention.

### SynthID Is Present on Free TTS and Is Expanding Across Paid Audio

As of July 2026, all Free-user Text to Speech generations carry an imperceptible Google DeepMind SynthID watermark. Paid-feature coverage is partial and expanding, so a specific paid generation may also be watermarked. ElevenLabs says the watermark survives common transformations, does not audibly degrade quality, and is not linked to other platforms' monetization or content policies. It cannot currently be assumed absent from Starter output.

**Evidence**

- The 25 June 2026 [SynthID announcement](https://elevenlabs.io/blog/synthid), updated 7 July, says Free-user TTS is watermarked and broader rollout is planned.
- The [Audio Detector documentation](https://elevenlabs.io/docs/eleven-creative/audio-tools/audio-detector) says all Free TTS and select paid features are currently covered, rollout is phased, no pre-June-2026 audio carries SynthID, and users can test a file with the signed-in detector.

**Implication**

Watermarking is not a rights blocker and should not change the Starter recommendation. Run the final MP3 through ElevenLabs' Audio Detector and record the result only if provenance status matters to the submission archive. Do not promise an unwatermarked paid file.

### Voice Cloning Is Unnecessary and Adds Consent/Privacy Risk

Starter includes Instant Voice Cloning, but the demo does not need it. ElevenLabs requires permission to clone a voice; its Terms permit a user voice model only for the user's own voice or a voice they are authorized to share, and the use policy forbids unauthorized, deceptive, or harmful impersonation. Professional Voice Cloning is Creator-and-above and has stronger identity verification. A Default voice avoids collecting biometric voice data and avoids evidence work around consent.

**Evidence**

- [My Voices](https://help.elevenlabs.io/hc/en-us/articles/13313587528849-What-is-My-Voices) says Starter supports Instant Voice Cloning and that permission is required; Creator supports Professional Voice Cloning.
- The [Terms](https://elevenlabs.io/terms-of-use) govern user voice models created from the user's voice or an authorized voice.
- The [Prohibited Use Policy](https://elevenlabs.io/use-policy) forbids unauthorized or deceptive impersonation and guardrail evasion.
- The [Privacy Policy](https://elevenlabs.io/privacy-policy) treats Voice Data as potentially biometric data and explains verification and retention.

**Implication**

Use a Default voice. If the user insists on a clone, require explicit documented consent, use only the consenting speaker's recording, disable training first, and retain the consent evidence. That branch is lower priority for a deadline-driven three-minute demo.

### Codex Can Automate the Reproducible Pipeline, While the User Owns Audition and Account Decisions

Codex can normalize the spoken script; count characters; build alias/IPA rules; call the TypeScript SDK with explicit model, voice, seed, settings, and format; save MP3/alignment; log actual character cost and request IDs; group character timing into word/caption timing; map seconds to frames; place assets in Remotion; set composition duration from the selected audio; add synchronized captions/callouts; and verify that the rendered video stays under the Build Week limit.

The user must create/sign into the account, disable training in the account UI, choose and pay for Starter, accept current terms, audition voices, judge naturalness and brand fit, confirm pronunciations, decide whether any audible AI disclosure is desired, and approve the final take. Those judgments cannot be safely delegated to text-only automation.

**Evidence**

- API-key creation and TypeScript generation are documented in the [quickstart](https://elevenlabs.io/docs/eleven-api/quickstart); exact cost and request metadata are available in the [API introduction](https://elevenlabs.io/docs/api-reference/introduction).
- Character timing is machine-readable through the [timing endpoint](https://elevenlabs.io/docs/api-reference/text-to-speech/convert-with-timestamps).
- Remotion accepts local audio and frame-based timing through [`<Audio>`](https://www.remotion.dev/docs/media/audio) and [`staticFile()`](https://www.remotion.dev/docs/staticfile).
- ElevenLabs itself recommends testing model/voice pairings and says voice selection is a dominant quality factor ([product guide](https://elevenlabs.io/docs/eleven-creative/playground/text-to-speech)).

**Implication**

The correct workflow is human selection followed by automated asset production. Codex should never choose a final voice solely from labels or regenerate production audio without explicit approval because both performance and billing vary per generation.

### Small Proof-Generation Gate

1. **Prepare one 35–45 second proof excerpt** (roughly 90–120 words) containing the product name, “Codex,” “GPT-5.6,” at least one acronym, a version/number, a URL-like phrase, and one emotionally important sentence. Keep a display version and a spoken-normalized version.
2. **Before any upload, disable model improvement** in Terms and privacy → Data use.
3. **On Free, audition without retaining output**: generate the same excerpt in v3 Natural for Elara, Finley, Alicia, Lawrence, and Caleb; optionally Baxter/Stuart if accent fit matters. Use one generation per voice and reject obvious mismatches from the labels and first listen. This should consume roughly 2,700–4,200 credits for five to seven approximately 540–720-character excerpts.
4. **User selects two finalists** based on naturalness, intelligibility at normal laptop/headphone playback, energy against visuals, lack of “announcer” excess, and accurate technical pronunciation. Test only one deliberate script/IPA correction per finalist.
5. **Upgrade to Starter before retaining anything**. Confirm the account shows the paid plan and commercial license, then regenerate the winning excerpt under the paid subscription. Free-plan audio must not enter the submission.
6. **Production gate**: generate one full v3 take with `with-timestamps`, `mp3_44100_128`, explicit seed, and logged character cost. If it fails a defined issue, make at most two purposeful passes: one corrected v3 pass, then one Multilingual v2 fallback with explicit speed/break/alias controls. This caps the planned production budget at three full generations.
7. **Freeze the winner**: user approves the take; Codex stores MP3, alignment JSON, metadata, and checksum; Remotion uses only those local assets. Verify the final video is less than three minutes and that all third-party material is licensed.

**Evidence**

- ElevenLabs recommends testing voices/models and documents generation variability ([product guide](https://elevenlabs.io/docs/eleven-creative/playground/text-to-speech)).
- Free output cannot be used in the final contest entry; paid-period output retains commercial rights indefinitely ([publication help](https://help.elevenlabs.io/hc/en-us/articles/13313564601361-Can-I-publish-the-content-I-generate-on-the-platform), [Prohibited Use Policy](https://elevenlabs.io/use-policy)).
- The Build Week demonstration video should be less than three minutes and publicly visible on YouTube ([official rules](https://openai.devpost.com/rules)).

**Implication**

This gate minimizes wasted credits, prevents accidental use of Free-plan audio, makes subjective selection explicit, and produces a deterministic Remotion input. It also creates a clear stopping rule instead of open-ended regeneration.

## Notes

- **Unsupported “no card” claim:** official sources confirm automatic assignment to Free but do not explicitly promise signup without a payment method. Treat “no credit card required” as unverified until observed in the live signup flow.
- **Exact cost caveat:** exact credits require the final script or API `character-cost` header. Word-count-derived figures are planning estimates. The exact Starter cash outlay is USD $6 plus jurisdiction-dependent tax, provided no additional paid usage or custom-rate voice is selected.
- **Voice IDs:** current Default voice IDs are available only through the authenticated dashboard/API and can be copied from My Voices. The unauthenticated voice-list endpoint returned an authorization error during research. Do not hard-code scraped IDs; copy the selected voice's current ID after audition.
- **Default voice transition:** ElevenLabs' wording is internally awkward: it says “Default voices” expire on 31 December 2026 while also saying their new replacements remain available in perpetuity. The practical reading is that the old named set is retiring and the new replacement set persists. For this July 2026 one-off, select from the new replacement names and archive the rendered audio.
- **192 kbps documentation conflict:** the endpoint reference says Creator is sufficient for MP3 192 kbps, while the current pricing comparison presents 192 kbps under Pro. The Starter route uses 128 kbps and is unaffected.
- **Watermark status:** paid TTS watermark coverage is rolling out and cannot be guaranteed present or absent. Use the Audio Detector on the actual generated file if status matters.
- **Hosted retention:** commercial-use rights survive cancellation, but ElevenLabs does not guarantee that generated content remains hosted forever on Free. Download all production artifacts before cancelling.
- **Legal caveat:** paying for Starter is the most defensible interpretation because Free use expressly excludes contests and the hackathon awards prizes. Final compliance remains the entrant's responsibility under both ElevenLabs' current terms and the Build Week rules.

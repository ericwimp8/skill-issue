# Browser-First TTS Studios for a Remotion Build Week Demo

## Assignment

**Goal:** Compare Murf Studio/Creator, Descript AI Speech, and Speechify Studio for producing a natural, public, approximately three-minute English voice-over on macOS for a Remotion-rendered OpenAI Build Week YouTube demo. Determine the best browser-only route, establish whether each route can support three full narration passes, and identify the legal, export, timing, privacy, and human-operation constraints that matter before committing money or the final script.

**Scope:** Internet-only research using current official product, help, pricing, privacy, terms, security, and OpenAI Build Week rules pages available on 2026-07-19. Independent review evidence is used only for subjective naturalness and workflow friction. The comparison assumes a 350–500-word script, stock/licensed vendor voices rather than cloning a third party, and local post-processing of an exported file before it is added to Remotion.

**Exclusions:** API-provider comparisons, voice-clone quality testing, automated browser operation, legal advice, and an in-product listening/export test. No vendor account was created and no paid checkout was entered.

## Sources

### Competition context

- OpenAI Build Week Official Rules, especially Submission Requirements and Intellectual Property: https://openai.devpost.com/rules

### Murf

- Pricing: https://murf.ai/pricing
- Current Murf pricing article: https://murf.ai/blog/best-text-to-speech-software
- Workspace and generation limits: https://help.murf.ai/what-is-a-workspace
- Free-trial download restriction: https://help.murf.ai/can-i-download-my-project-during-the-free-trial
- Voice export formats and quality: https://help.murf.ai/voice-only
- Languages and accents: https://help.murf.ai/languages-and-accents
- Gen2 controls: https://help.murf.ai/gen2
- Pronunciation controls: https://help.murf.ai/pronunciation
- Voice-setting generation accounting: https://help.murf.ai/does-changing-voice-settings-sped-pitch-etc-consume-vgt
- Script/subtitle export: https://help.murf.ai/script
- Cancellation and project access: https://help.murf.ai/what-happens-to-my-account-and-projects-after-canceling
- Post-subscription commercial rights: https://help.murf.ai/will-my-commercial-rights-and-business-license-be-valid-after-my-subscription-has-ended
- Terms: https://murf.ai/legal/terms-of-service
- Privacy: https://murf.ai/legal/privacy-policy
- Security and data lifecycle: https://murf.ai/security
- Independent subjective review aggregate (G2): https://www.g2.com/products/murf-ai/reviews

### Descript

- Pricing and plan comparison: https://www.descript.com/pricing
- AI-credit accounting: https://help.descript.com/hc/en-us/articles/27841674958221-Track-and-understand-your-Media-minutes-and-AI-Credits
- Text-to-speech and ElevenLabs models: https://help.descript.com/hc/en-us/articles/10166020211085-Generate-text-to-speech-audio
- Pronunciation controls: https://help.descript.com/hc/en-us/articles/10165910447501-AI-Speaker-pronunciation-tips
- AI Speaker limitations: https://help.descript.com/hc/en-us/articles/12923453527565-Troubleshooting-and-using-AI-speakers
- Audio export: https://help.descript.com/hc/en-us/articles/10255754857741-Export-an-audio-file
- Transcript timing export: https://help.descript.com/hc/en-us/articles/10255753048589-Export-your-transcript-as-a-text-file
- Subtitle export: https://help.descript.com/hc/en-us/articles/10255811669773-Exporting-subtitles
- Cancellation/downgrade: https://help.descript.com/hc/en-us/articles/10255894257037-Downgrade-or-cancel-a-subscription
- Payment methods: https://help.descript.com/hc/en-us/articles/10255834199437-Manage-and-update-your-payment-method
- Terms: https://www.descript.com/terms
- Privacy: https://www.descript.com/privacy
- Data-sharing controls: https://help.descript.com/hc/en-us/articles/10255866490125-Account-data-and-privacy
- Independent workflow review (TechRadar): https://www.techradar.com/best/best-ai-tools

### Speechify

- Studio pricing, credits, features, and rights: https://speechify.com/pricing-studio/
- Voice-over workflow and export: https://speechify.com/blog/ultimate-guide-to-creating-ai-voice-overs-in-speechify-studio/
- Studio overview: https://speechify.com/studio/
- Current general terms: https://speechify.com/terms/
- Studio-specific terms: https://speechify.com/studio-terms/
- Studio privacy: https://speechify.com/studio-privacy/
- Community Voice terms: https://speechify.com/studio-community-voices-terms/

## Findings

### Finding 1: The Build Week deliverable is publicly distributed, prize-bearing, and strictly under three minutes

The official rules require a demonstration video that is **less than three minutes**, has clear audio, and is publicly visible on YouTube. They also require the entrant to own the submission and have permission for third-party material. The competition offers monetary prizes and grants OpenAI/Devpost promotional rights over submissions. A nominally “personal” or non-commercial free-output permission is therefore a poor fit even if a vendor preview can technically be played or captured.

**Evidence:** OpenAI Build Week requires the video to be less than three minutes, publicly visible on YouTube, and free of unlicensed third-party material; it also requires the submission to be solely owned and offers cash prizes: https://openai.devpost.com/rules#L192 and https://openai.devpost.com/rules#L210 and https://openai.devpost.com/rules#L269. Murf Free explicitly has no commercial rights or downloads: https://murf.ai/pricing. Speechify Free explicitly has no commercial usage rights and no MP3 download: https://speechify.com/pricing-studio/.

**Implication:** Treat the final narration as public/promotional use. Do not publish a Murf Free or Speechify Free preview, record it from the browser, or rely on a personal-use interpretation. Keep the final narration safely below 180 seconds; a 500-word script is likely too long unless delivered unnaturally fast. A practical target is roughly 380–430 spoken words after timing the chosen voice.

### Finding 2: Descript Free is the best first browser-only route if a US-English stock voice passes the proof gate

Descript is the strongest first attempt because it is the only compared route whose free plan appears to allow the complete proof-to-export workflow without a card or a purchase. The Free plan provides 100 one-time AI credits; the current official accounting page has described default text-to-speech as approximately five credits per minute, so three separate three-minute passes cost about 45 credits and leave about 55 credits for a short audition and targeted repairs. Descript states that no credit card is required to try the Free plan. It supports stock AI speakers, ElevenLabs Multilingual v2 and v3, tone tags, local WAV/MP3 export, and timed transcript/SRT/VTT export.

The trade-off is that Descript is an editor with TTS embedded in it, rather than a dedicated voice-directing studio. Its pronunciation workflow is phonetic respelling rather than a documented IPA library, and an official troubleshooting page says the current model is based on US English pronunciation. ElevenLabs v3 can also shift tone, accent, volume, or speaker identity across long paragraphs. Those weaknesses are material for a technical demo with terms such as “Codex,” “GPT-5.6,” “Remotion,” package names, and product names.

**Evidence:** Descript Free is $0 with 100 one-time AI credits, while its pricing FAQ says no card is required: https://www.descript.com/pricing. Descript’s official credit table has listed text-to-speech at approximately five AI credits per minute, while warning that model selection affects cost: https://help.descript.com/hc/en-us/articles/27841674958221-Track-and-understand-your-Media-minutes-and-AI-Credits. Its current TTS guide supports stock voices, ElevenLabs v3 tone tags, and describes continuity limitations: https://help.descript.com/hc/en-us/articles/10166020211085-Generate-text-to-speech-audio. WAV and MP3 local export are supported: https://help.descript.com/hc/en-us/articles/10255754857741-Export-an-audio-file. Timed transcripts and SRT/VTT are supported: https://help.descript.com/hc/en-us/articles/10255753048589-Export-your-transcript-as-a-text-file and https://help.descript.com/hc/en-us/articles/10255811669773-Exporting-subtitles. Pronunciation relies on spelling changes and punctuation: https://help.descript.com/hc/en-us/articles/10165910447501-AI-Speaker-pronunciation-tips. The US-English accent limitation is stated here: https://help.descript.com/hc/en-us/articles/12923453527565-Troubleshooting-and-using-AI-speakers.

**Implication:** Start with Descript Free and spend no money until a difficult 60–75-word excerpt proves voice quality, term pronunciation, clean audio export, and usable timing files. If that proof passes, it is sufficient for three full passes at the published approximate rate and is the lowest-friction route. If the voice sounds generically US, paragraph continuity drifts, or technical terms require conspicuous respelling, move to Murf Creator rather than spending Descript credits trying to force an unsuitable voice.

### Finding 3: Murf Creator is the best paid fallback for directed pronunciation and accent control

Murf is the most defensible paid production route when exact pronunciation, an Australian or other regional English accent, and repeatable per-word direction matter more than the lowest price. Murf Creator offers a dedicated browser studio, all paid voices, commercial rights, unlimited downloads, pronunciation changes, and clean MP3/WAV/FLAC output. The current pricing surface shows Creator from US$19/month when billed annually (US$228/year), while a current Murf article lists Creator Lite at US$29 per user on monthly billing. Creator supplies two hours of generation per monthly cycle or 24 hours per annual cycle, making three three-minute passes trivial.

Murf’s pronunciation editor accepts IPA or alternative spellings, can change locale for a selected word or phrase, and can apply a pronunciation across a project. It supports Australian English and other regional accents. Its Gen2 feature set includes pacing and intonation controls, but the current pricing comparison reserves advanced Emphasis, Variability, and Say It My Way for Business; Creator should therefore be evaluated on the exact selected voice and controls shown in-product rather than assumed to include every Murf marketing feature.

**Evidence:** Pricing and plan features: https://murf.ai/pricing. Current official monthly-price description: https://murf.ai/blog/best-text-to-speech-software. Creator monthly provides two hours of voice generation and five concurrent project credits: https://help.murf.ai/what-is-a-workspace. Paid audio export supports MP3, WAV, and FLAC up to 48 kHz/16-bit: https://help.murf.ai/voice-only. Australian English and multiple accents are supported: https://help.murf.ai/languages-and-accents. IPA, respelling, and locale overrides are documented: https://help.murf.ai/pronunciation. SRT/VTT script export includes timestamps: https://help.murf.ai/script.

**Implication:** Choose one month of Murf Creator after a free audition if Descript fails on accent or technical pronunciation. It is the safest paid fallback for “say this developer name exactly this way” work. Avoid annual billing for a one-off demo unless Murf will remain part of the content workflow.

### Finding 4: Speechify Studio Starter is the best nominal value, but its proof and policy gaps prevent it from being the default recommendation

Speechify Studio Starter is attractively specified: US$19/month, 7,200 Studio credits, 1,000+ voices, commercial rights, a pronunciation library, pauses, speed, pitch, volume, emotion/emphasis, and WAV/MP3/OGG plus SRT/WebVTT/DOCX export. Voice-over generation consumes one credit per second, so the plan represents 120 minutes and three full passes consume 540 of 7,200 credits. The free plan provides 600 credits, exactly ten minutes, so it can audition three three-minute passes with a one-minute buffer.

However, Speechify Free cannot download MP3 and has no commercial rights. Its Studio-specific terms also request “Speechify.com” attribution for free-plan output published outside the platform. Paid output is stated to be owned by the user with commercial rights in perpetuity, but the official pages inspected do not establish how long projects remain editable or downloadable after a paid subscription ends. The Studio terms additionally permit uploaded content to be used to improve Speechify services and models. These are manageable for a public hackathon script, but they are less favorable and less operationally clear than Descript’s free proof path or Murf’s cancellation documentation.

**Evidence:** Pricing, credit consumption, perpetual commercial rights, and Free-plan download restriction: https://speechify.com/pricing-studio/. The official workflow documents pronunciation, tone, pacing, WAV/MP3/OGG, and subtitle export: https://speechify.com/blog/ultimate-guide-to-creating-ai-voice-overs-in-speechify-studio/. Studio-specific terms permit commercial use, say generated content remains the user’s, permit uploaded content to improve services/models, and request attribution for free-plan output: https://speechify.com/studio-terms/. Current general terms permit commercial/public distribution only with the proper Studio commercial license: https://speechify.com/terms/.

**Implication:** Speechify Starter is a reasonable lower-price paid alternative if its voice wins a blind audition and the user accepts its content-improvement terms. Before purchase, obtain a support answer or in-product confirmation for project retention after cancellation and exact export availability. Download the final WAV, MP3 backup, and timing file before canceling.

### Finding 5: All three can cover three complete passes, but only Descript Free can plausibly deliver the final usable asset without payment

Assuming each pass is at most 180 seconds, three full regenerations require nine minutes. Murf Free includes ten generation minutes, and changes to voice, style, pitch, speed, pauses, emphasis, pronunciation, punctuation, and volume on unchanged text do not consume additional generation time. Speechify Free includes 600 credits at one credit per second, also ten minutes. Descript Free includes 100 one-time credits and has published TTS at approximately five credits per minute, implying about 20 minutes. Each free allowance is therefore numerically sufficient for three full passes, but the practical outcomes differ.

Murf Free and Speechify Free block usable final export and commercial use. Descript supports audio export from the Free workflow and does not document an audio watermark; its watermark documentation is specific to video exports. Descript’s current credit page now emphasizes model-dependent cost, so the exact v3 debit must be observed in the account’s Usage tab rather than treated as fixed.

**Evidence:** Murf Free ten-minute limit and no downloads: https://help.murf.ai/what-is-a-workspace and https://help.murf.ai/can-i-download-my-project-during-the-free-trial. Murf settings on unchanged text do not consume VGT: https://help.murf.ai/does-changing-voice-settings-sped-pitch-etc-consume-vgt. Speechify Free has 600 credits and consumes one per voice-over second: https://speechify.com/pricing-studio/. Descript Free has 100 one-time credits and approximate five-credit-per-minute TTS accounting: https://www.descript.com/pricing and https://help.descript.com/hc/en-us/articles/27841674958221-Track-and-understand-your-Media-minutes-and-AI-Credits. Descript audio export supports WAV/MP3 and its Free watermark documentation concerns video: https://help.descript.com/hc/en-us/articles/10255754857741-Export-an-audio-file and https://help.descript.com/hc/en-us/articles/10255814959245-Export-an-mp4-video-or-a-GIF.

**Implication:** Numerical sufficiency is not the buying decision. The decisive questions are commercial rights, clean export, and whether difficult words survive a full paragraph. Descript should be tried first because those questions can be answered before payment.

### Finding 6: Paid outputs are suitable for public YouTube and hackathon use; free Murf/Speechify outputs are not

Murf grants commercial rights to paid-created voices and states those rights survive cancellation. Speechify Starter and Creator include commercial rights, and Speechify says users own audio output and commercial rights in perpetuity. Descript treats generated Input and Output as user content owned by the user to the extent protectable by law and supports direct YouTube publishing. None of the inspected Murf or Descript pages imposes an attribution requirement on paid or free audio output; Speechify’s Studio terms request attribution for free-plan output, while its current free plan also denies commercial rights.

Stock voices should be used for this one-off narration. Cloning introduces unnecessary consent and privacy obligations: Descript requires a recorded consent statement; Murf requires explicit written consent for third-party training audio; Speechify requires the necessary licenses, releases, and permissions for the speaker.

**Evidence:** Murf commercial rights and survival after cancellation: https://murf.ai/legal/terms-of-service and https://help.murf.ai/will-my-commercial-rights-and-business-license-be-valid-after-my-subscription-has-ended. Speechify perpetual paid commercial rights: https://speechify.com/pricing-studio/. Descript ownership of generated Output: https://www.descript.com/terms. Descript YouTube export: https://help.descript.com/hc/en-us/articles/10255756672909-Export-to-YouTube. Consent requirements: https://help.descript.com/hc/en-us/articles/10119641262221-Create-a-new-AI-Speaker, https://murf.ai/legal/terms-of-service, and https://speechify.com/studio-terms/.

**Implication:** A clean stock voice from Descript Free, Murf Creator, or Speechify Starter can be used in the public submission under the vendors’ documented output terms. Preserve the invoice/plan confirmation, a PDF or screenshot of the applicable license page, the generated source WAV, and the final timing file as an evidence bundle. Avoid third-party voice cloning for this deadline-bound demo.

### Finding 7: Descript has the strongest timing handoff; Murf is close; Speechify’s timing export is less documented

Descript can export WAV/MP3, SRT/VTT subtitles, and a transcript with configurable timecode intervals, paragraph timecodes, speaker labels, markers, and offsets. That is the most useful browser-to-Remotion handoff because the exported time data can seed scene boundaries or caption cues. Murf exports SRT/VTT whose time span begins at zero and follows the rendered block timing, plus a single or block-split audio file. Speechify exports SRT/WebVTT and has a draggable timeline, but its public documentation does not define timestamp precision, word-level timing, or whether subtitle timecodes remain identical to a separately exported WAV.

**Evidence:** Descript transcript timing: https://help.descript.com/hc/en-us/articles/10255753048589-Export-your-transcript-as-a-text-file. Descript subtitles: https://help.descript.com/hc/en-us/articles/10255811669773-Exporting-subtitles. Murf timestamped script export: https://help.murf.ai/script. Speechify subtitle/audio export and timeline workflow: https://speechify.com/blog/ultimate-guide-to-creating-ai-voice-overs-in-speechify-studio/.

**Implication:** For Descript or Murf, export WAV plus SRT/VTT in the same action/session and treat the rendered WAV duration as authoritative. For Speechify, verify synchronization in the proof gate and regenerate timings locally if the subtitle file drifts.

### Finding 8: Browser studios beat an API for this one-off asset when human audition and direction dominate

For one under-three-minute narration, browser studios remove API-key handling, SDK integration, request chunking, retry logic, file assembly, and timing-generation code. They expose voice audition, paragraph playback, pronunciation repair, tone/pacing controls, and subtitle export to a human operator. That is valuable when the principal uncertainty is “Which voice and delivery sounds credible?” rather than “How can this be generated repeatedly?”

An API becomes the better route when Codex must generate the voice directly, builds must be reproducible, many scripts/languages/voices are needed, narration is generated in CI, exact structured timing is required, or late script edits are likely to recur. The browser terms also make this boundary important: Speechify forbids automated access except through its Public API; Descript forbids unapproved applications interacting with the service and scraping; Murf forbids reselling its service and using Murf-created voices to train or synthesize another voice model.

**Evidence:** Speechify automation restriction: https://speechify.com/studio-terms/ and https://speechify.com/terms/. Descript automation/scraping restriction: https://www.descript.com/terms. Murf output restrictions: https://murf.ai/legal/terms-of-service. Murf and Speechify dedicated controls and export paths: https://help.murf.ai/products and https://speechify.com/blog/ultimate-guide-to-creating-ai-voice-overs-in-speechify-studio/.

**Implication:** Use the browser studio as a human-operated content tool for the final one-off narration. Use an API if the requirement changes to agent-generated or repeatable narration. Do not automate the browser as a substitute for an API.

### Finding 9: The lawful split is human studio operation followed by local Codex post-processing

The user should create the account, accept terms, choose or purchase the plan, audition voices, paste/import the script, approve pronunciations and delivery, initiate generation, and click export. This is especially important for Speechify, whose terms require genuine front-facing use and prohibit automated access outside the API.

Codex can lawfully prepare the script, produce phonetic/IPA candidate spellings, split it into narration blocks, maintain a cue sheet, and calculate the word budget before studio use. After the user downloads the authorized output, Codex can inspect duration and metadata with local tools, normalize loudness, trim leading/trailing silence, split or concatenate sections, transcode a preservation WAV to a web-friendly derivative, parse SRT/VTT, generate Remotion timing constants, and wire the audio into the composition. These actions operate on the user’s downloaded output rather than the vendor service. Murf-created voice audio must not be used to train another AI model or synthesize a derivative voice.

**Evidence:** Speechify permits front-facing human use and its Public API while prohibiting other automation: https://speechify.com/terms/. Descript prohibits scraping and unapproved interacting applications: https://www.descript.com/terms. Murf prohibits AI training or further voice synthesis from Murf-created voices: https://murf.ai/legal/terms-of-service.

**Implication:** Document the manual generation settings in a small production note for reproducibility, but keep all browser clicks human. Local audio engineering and Remotion integration are appropriate post-processing; browser scripting, paywall circumvention, or capturing a non-downloadable preview are not.

### Finding 10: Privacy is acceptable for a public script, with Descript offering the clearest consumer opt-out

Descript says projects may be used to improve the service, but users can disable “Share data with Descript”; its help center says current production models use no user data and internal R&D uses data only from opted-in users. Third-party AI providers are contractually barred from training on Descript Inputs/Outputs. Custom AI Speaker training has additional de-identified research and human-listening provisions, which are avoided by using a stock speaker.

Speechify’s Studio terms allow uploaded content to improve its services and models and grant a broad service license. Its privacy policy processes voice samples and generated sample history and transfers/processes data in the United States. Murf collects the uploaded text/audio/video needed to provide the service, stores and processes customer data in AWS US-East-2, and its consumer privacy policy does not provide the same plain-language “no model training” commitment offered on Murf Enterprise terms.

**Evidence:** Descript data-sharing opt-out: https://help.descript.com/hc/en-us/articles/10255866490125-Account-data-and-privacy and https://www.descript.com/privacy. Speechify content-improvement license: https://speechify.com/studio-terms/; Studio privacy: https://speechify.com/studio-privacy/. Murf privacy and US data residency: https://murf.ai/legal/privacy-policy and https://murf.ai/security.

**Implication:** Use only the already-public demo narration script and no secrets, private source, unpublished customer data, or personal voice samples. In Descript, disable data sharing before importing the script. For Murf or Speechify, accept that the script and generated sample will be processed in the United States and avoid voice cloning.

### Finding 11: Cancellation behavior favors Descript for retained edit access and Murf for explicit lifetime rights

Descript cancellation downgrades the Drive to Free after the paid term and keeps existing projects accessible, although paid-only features disappear. Murf retains commercial rights for created voice-overs after cancellation, but premium generation/download access ends with the billing cycle; projects remain in the workspace and can become locked by the Free plan’s concurrent-project limit. Speechify says paid commercial rights are perpetual and the general terms keep benefits active through the paid term, but its public documentation does not establish post-cancellation project retention or re-download behavior.

**Evidence:** Descript downgrade behavior: https://help.descript.com/hc/en-us/articles/10255894257037-Downgrade-or-cancel-a-subscription. Murf cancellation and rights: https://help.murf.ai/what-happens-to-my-account-and-projects-after-canceling and https://help.murf.ai/will-my-commercial-rights-and-business-license-be-valid-after-my-subscription-has-ended. Speechify perpetual rights and subscription expiry: https://speechify.com/pricing-studio/ and https://speechify.com/terms/.

**Implication:** Export and archive the source WAV, MP3 backup, SRT/VTT, final text, settings screenshots, and license evidence before canceling any paid plan. Never depend on a vendor project remaining downloadable after the deadline.

### Finding 12: Availability in Australia is plausible for all three, but checkout is the only exact plan/currency proof

Murf markets service in more than 190 countries and supports Australian English. Speechify’s current terms expressly address non-waivable Australian consumer law and describe global operation. Descript accepts all currencies, bills in USD via Stripe, and excludes embargoed/restricted territories; its payment help warns that Stripe does not support every country but does not list Australia as excluded.

**Evidence:** Murf global reach and Australian English: https://murf.ai/educators-and-non-profits and https://help.murf.ai/languages-and-accents. Speechify Australian-law clause: https://speechify.com/terms/. Descript payment geography: https://help.descript.com/hc/en-us/articles/10255834199437-Manage-and-update-your-payment-method and https://www.descript.com/terms.

**Implication:** All three are reasonable candidates from Adelaide on macOS, but prices are USD and may incur GST, exchange, or card conversion. Confirm the final Australian checkout total before buying. The free Descript proof avoids that uncertainty unless it fails creatively.

### Finding 13: Independent evidence favors Murf’s dedicated workflow, while Descript’s strength is editing rather than pure TTS

Independent evidence should be treated as subjective. G2’s verified-review aggregate for Murf reports repeated praise for natural voices, customization, and ease of use, alongside recurring complaints about pricing, occasional robotic delivery, and pronunciation problems. TechRadar’s hands-on Descript review praises its transcript-based workflow and accessibility but positions it as a broad video editor rather than a frame-precise production tool. Public Speechify commentary is heavily mixed with reviews of the separate Reader product, so it does not provide a clean independent signal for current Studio voice-over quality.

**Evidence:** G2 Murf review aggregate: https://www.g2.com/products/murf-ai/reviews. TechRadar Descript test: https://www.techradar.com/best/best-ai-tools. Speechify itself states that Studio and Reader are separate subscriptions/products: https://speechify.com/pricing-studio/.

**Implication:** Do not choose by review score. Use the same difficult excerpt and blind-listen to exported files. Murf has the strongest independent signal for a dedicated voice-over workflow; Descript has the strongest no-purchase proof path; Speechify needs direct audition because Reader reviews are not evidence about Studio.

### Finding 14: A small proof-generation/export gate can settle the remaining decision before purchase

Use a 60–75-word excerpt containing: “OpenAI Build Week,” “Codex,” “GPT-5.6,” “Remotion,” the project name, one acronym, one sentence with contrastive emphasis, one short pause, and one URL or package-style token. Keep the excerpt under 35–40 seconds.

The human operator should run this gate:

1. In Descript Free, disable data sharing, select two stock voices, generate the excerpt with ElevenLabs v3, and use concise tone tags only.
2. Correct technical names using phonetic respelling and punctuation, then export 48 kHz WAV plus SRT/VTT and a timed transcript.
3. Confirm the account Usage tab debit. Reserve at least 50 credits after the proof if three full passes are planned; if the proof consumes unexpectedly high credits, stop and re-budget.
4. Locally inspect the WAV for duration, sample rate, audible tags, clipping, leading/trailing silence, and embedded vendor metadata; compare SRT/VTT end time with WAV duration.
5. Blind-listen on laptop speakers and headphones for technical-name accuracy, paragraph continuity, pace, emphasis, and non-robotic cadence.
6. Pass only if the WAV is clean, timing drift is negligible, every required technical term is intelligible, and a full pass is projected below 180 seconds.
7. If Descript fails only on voice/accent/pronunciation, audition Murf Free and Speechify Free using the same excerpt. Buy one month of Murf Creator if exact pronunciation or Australian accent wins; buy Speechify Starter only if its voice clearly wins and the user accepts the privacy/retention caveats.

**Evidence:** Descript supports the required free plan, v3 tone control, exports, and Usage tab: https://www.descript.com/pricing, https://help.descript.com/hc/en-us/articles/10166020211085-Generate-text-to-speech-audio, https://help.descript.com/hc/en-us/articles/10255754857741-Export-an-audio-file, and https://help.descript.com/hc/en-us/articles/27841674958221-Track-and-understand-your-Media-minutes-and-AI-Credits. Murf and Speechify allow free audition but block final download/commercial use: https://murf.ai/pricing and https://speechify.com/pricing-studio/.

**Implication:** The purchase decision becomes evidence-based: Descript Free if it passes; Murf Creator if dedicated pronunciation/accent direction is necessary; Speechify Starter only when its specific voice wins and its policy gaps are consciously accepted.

## Notes

- **Unsupported exact Descript debit:** The official help center previously published approximately five credits per TTS minute, while the current wording emphasizes model-dependent cost. Exact ElevenLabs v3 consumption for this script requires observing the in-product Usage tab.
- **Unsupported clean-audio guarantee on Descript Free:** Official documentation specifies WAV/MP3 export and documents watermarks for video, not audio. It does not explicitly promise “no audible tag” or “no vendor metadata” on Free-plan audio. The proof export must verify this.
- **Unsupported Speechify card requirement:** The current Studio Free-plan page does not state whether a payment card is required. General Speechify free trials can require a payment method, but the indefinite Studio Free plan is a distinct offer.
- **Unsupported Speechify retention:** No inspected official page states how long Studio projects remain accessible or downloadable after cancellation.
- **Unsupported Speechify timing precision:** Official material confirms SRT/WebVTT export but does not define word-level timestamps, drift tolerance, or exact alignment between separately exported audio and subtitle files.
- **Murf price presentation caveat:** Official pages expose US$19/month on annual billing and US$29/month on monthly billing. The live checkout is authoritative for the Australian total and any current Lite/Plus naming.
- **Murf privacy caveat:** Murf’s explicit “no training on customer data” commitment appears in enterprise/MSA material. The consumer Creator privacy policy permits analytics and service improvement but does not state the same clear model-training exclusion.
- **Attribution caveat:** No attribution requirement was found for Murf paid output, Speechify paid Studio output, or Descript output. Speechify Studio terms do request Speechify.com attribution for free-plan output, which is independently unsuitable here because Free lacks commercial rights and download.
- **Term conflict caveat:** Speechify’s current general terms and older Studio-specific terms overlap. Use the proper paid Studio plan and archive the terms displayed at checkout.
- **Useful search terms for in-product proof:** `Codex pronunciation`, `GPT five point six`, `Remotion`, `Australian English stock voice`, `48 kHz WAV`, `SRT export`, `AI credits usage`, `commercial rights perpetuity`, `project retention cancellation`.

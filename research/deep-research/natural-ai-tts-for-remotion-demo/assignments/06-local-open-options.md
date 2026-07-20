# Privacy-Conscious Local TTS: Kokoro-82M via MLX-Audio and Chatterbox

## Assignment

**Goal:** Determine whether Kokoro-82M through MLX-Audio or Resemble AI Chatterbox is the stronger privacy-conscious, downloadable/local TTS proof candidate for an approximately three-minute, 350–500-word, public OpenAI Build Week YouTube demo produced on macOS and integrated into Remotion.

**Scope:** Internet-only inspection of current first-party repositories, model cards, license texts, source files, platform policies, hackathon rules, and Remotion documentation. Covers code/weight/output licensing, commercial and public-submission suitability, attribution, watermarks, voice-cloning consent, privacy, macOS and Apple Silicon requirements, dependency/hardware friction, invocation surfaces, voices and accents, output formats, timing/prosody controls, pronunciation and chunking, determinism, performance evidence, likely naturalness, technical-name risk, cost, Remotion integration, and a precise no-cloning generation gate.

**Exclusions:** No local repository or machine inspection; no package installation, model download, inference run, listening test, benchmark, or legal advice. Chatterbox Turbo and Multilingual V3 are mentioned only where they clarify the family; the primary Chatterbox comparison is the original English model that exposes a bundled voice and can generate without a reference clip.

## Sources

- Kokoro official model card, release facts, training disclosure, usage, and license metadata: https://huggingface.co/hexgrad/Kokoro-82M
- Kokoro official voice inventory, quality grades, and length guidance: https://huggingface.co/hexgrad/Kokoro-82M/blob/main/VOICES.md
- Kokoro official inference repository and advanced usage: https://github.com/hexgrad/kokoro
- Kokoro Apache-2.0 license text: https://github.com/hexgrad/kokoro/blob/main/LICENSE
- MLX-Audio official repository, requirements, CLI, Python API, OpenAI-compatible server, formats, and Kokoro example: https://github.com/Blaizzy/mlx-audio
- MLX-Audio current package metadata and dependency constraints: https://github.com/Blaizzy/mlx-audio/blob/main/pyproject.toml
- MLX-Audio current generic generator source and CLI flags: https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/generate.py
- MLX-Audio current Kokoro implementation: https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/models/kokoro/kokoro.py
- MLX-Audio MIT license: https://github.com/Blaizzy/mlx-audio/blob/main/LICENSE
- MLX Community Kokoro BF16 conversion model card and 355 MB size: https://huggingface.co/mlx-community/Kokoro-82M-bf16
- Chatterbox official repository, usage, model family, languages, watermark, and evaluation links: https://github.com/resemble-ai/chatterbox
- Chatterbox official Hugging Face model card: https://huggingface.co/ResembleAI/chatterbox
- Chatterbox official product/model page, consent wording, commercial FAQ, benchmark summary, watermark intent, and on-premise claims: https://www.resemble.ai/learn/models/chatterbox
- Chatterbox MIT license: https://github.com/resemble-ai/chatterbox/blob/master/LICENSE
- Chatterbox current package metadata and pinned dependencies: https://github.com/resemble-ai/chatterbox/blob/master/pyproject.toml
- Chatterbox current English inference source, local download path, MPS handling, stochastic controls, bundled conditionals, fixed speech-token limit, and watermark application: https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py
- Chatterbox official macOS example: https://github.com/resemble-ai/chatterbox/blob/master/example_for_mac.py
- Chatterbox official model files: https://huggingface.co/ResembleAI/chatterbox/tree/main
- Chatterbox official `t3_cfg.safetensors` size: https://huggingface.co/ResembleAI/chatterbox/blob/main/t3_cfg.safetensors
- Chatterbox official `s3gen.safetensors` size: https://huggingface.co/ResembleAI/chatterbox/blob/11120405b7cec7936a0fde4c84c6366789a6897a/s3gen.safetensors
- Chatterbox closed macOS multilingual issue documenting the need to adapt the Mac workaround: https://github.com/resemble-ai/chatterbox/issues/275
- OpenAI Build Week official rules, especially third-party authorization, original-work, intellectual-property, privacy, and publicity requirements: https://openai.devpost.com/rules
- YouTube official altered/synthetic-content disclosure guidance: https://support.google.com/youtube/answer/14328491
- YouTube official privacy guidance for synthetic likenesses and voices: https://support.google.com/youtube/answer/2801895
- YouTube official impersonation policy: https://support.google.com/youtube/answer/2801947
- Remotion official audio component documentation: https://www.remotion.dev/docs/html5-audio
- Remotion official `staticFile()` documentation: https://www.remotion.dev/docs/staticfile
- Practitioner performance evidence for equivalent Kokoro BF16 MLX weights on an M3 Max, clearly not an MLX-Audio benchmark: https://huggingface.co/sonic-speech/kokoro-82m-bf16
- Practitioner implementation note on a smaller Kokoro-only MLX package and its dependency surface, clearly not the selected MLX-Audio package: https://github.com/gabrimatic/kokoro-mlx

## Findings

### Finding 1: Licensing and public/commercial suitability favor both models, with clearer operational simplicity for Kokoro

Kokoro's official model card labels the weights Apache-2.0, explicitly welcomes production and commercial deployment, and describes the training data as permissive/non-copyrighted audio plus synthetic data that excludes custom voice clones. The official Kokoro inference code is also Apache-2.0. The MLX-Audio inference library is MIT, while its `mlx-community/Kokoro-82M-bf16` conversion preserves Apache-2.0 metadata. Chatterbox's repository code and Hugging Face weights are MIT; Resemble explicitly states that commercial products, self-hosting, weight modification, and production shipping are allowed without royalties, revenue share, or usage caps.

Neither project publishes a separate license that assigns or restricts ordinary generated speech output. The safest reading is that Apache-2.0 and MIT govern the distributed code/weights, while the user remains responsible for rights in the input text, voice identity, trademarks, and any included media. A rendered narration WAV is not identified by either project as a redistributed copy of the model. If code or weights are redistributed with the submission, preserve the applicable license/copyright notices: Apache-2.0 imposes license, modification-notice, retained-attribution, and applicable `NOTICE` duties for redistributed Work/derivatives; MIT requires its copyright and permission notice in copies or substantial portions of the software. A small credits entry such as “Narration generated locally with Kokoro-82M (Apache-2.0) via MLX-Audio (MIT)” is prudent provenance even when the final audio alone does not trigger a documented attribution requirement.

The Build Week rules require authorization for third-party SDKs/data and require the submission to be original, owned or permissioned, and free of third-party IP, privacy, and publicity violations. Both permissive stacks can satisfy that requirement, but cloning a recognizable third-party voice would introduce a separate permission burden. Resemble's own instructions say to use a reference clip only for a voice the user has permission to use. YouTube likewise provides removal routes for synthetic content that realistically sounds like a person and prohibits misleading unauthorized impersonation. A preset/bundled-voice-only workflow is therefore the stronger proof posture.

**Evidence:** Kokoro license and commercial statement: https://huggingface.co/hexgrad/Kokoro-82M and https://github.com/hexgrad/kokoro/blob/main/LICENSE. MLX-Audio license: https://github.com/Blaizzy/mlx-audio/blob/main/LICENSE. MLX conversion license: https://huggingface.co/mlx-community/Kokoro-82M-bf16. Chatterbox model/code license and commercial FAQ: https://huggingface.co/ResembleAI/chatterbox, https://github.com/resemble-ai/chatterbox/blob/master/LICENSE, and https://www.resemble.ai/learn/models/chatterbox. Build Week rights requirements: https://openai.devpost.com/rules. Consent and YouTube identity rules: https://www.resemble.ai/learn/models/chatterbox, https://support.google.com/youtube/answer/2801895, and https://support.google.com/youtube/answer/2801947.

**Implication:** Both are commercially and hackathon-compatible if license notices are preserved where code/weights are redistributed and all text/voice rights are controlled. Kokoro's fixed preset workflow avoids the highest-risk Chatterbox feature—zero-shot voice cloning—and is easier to prove compliant.

### Finding 2: Kokoro through MLX-Audio is the stronger macOS-native privacy path

MLX-Audio is explicitly built for Apple Silicon and requires Python 3.10+, an M1/M2/M3/M4-class Mac, MLX, and optionally `ffmpeg`. The current core package pins broad modern dependencies (`mlx`, `mlx-lm`, `transformers`, NumPy, SciPy, sounddevice, Hugging Face Hub); Kokoro additionally needs `misaki` for text processing. The BF16 conversion is 355 MB, and 8-bit/4-bit variants are documented. WAV works without `ffmpeg`; `ffmpeg` is needed only for MP3, FLAC, OGG, Opus, or Vorbis encoding.

The privacy boundary is favorable: first use downloads package/model artifacts from package registries and Hugging Face, but generation then happens on the Mac. Direct CLI or Python invocation avoids opening a network service. If the OpenAI-compatible local server is used, bind it to `127.0.0.1`, not `0.0.0.0`, for a proof build. Script text is not sent to a hosted inference provider by the documented local path. “Air-gapped from first launch” is unsupported because weights must first be obtained, but “local inference after artifact download” is supported.

Chatterbox also downloads weights with `hf_hub_download()` and performs inference locally; Resemble expressly supports on-premise deployment. Its Mac path is materially heavier: the project says it was developed/tested on Python 3.11 and Debian 11, pins PyTorch/torchaudio, Transformers, Diffusers, Librosa, Gradio, a Git dependency on PerTh, and other packages, and the English checkpoint downloads roughly 3.20 GB of core weights (`t3_cfg.safetensors` 2.13 GB + `s3gen.safetensors` 1.06 GB + `ve.safetensors` about 5.7 MB, plus small files). That is approximately nine times the 355 MB Kokoro BF16 model footprint before dependency sizes. The official Mac example detects MPS but monkey-patches `torch.load` for device mapping; current source also contains CPU-first mapping for non-CUDA devices. A closed official issue confirms Apple users had to adapt the Mac workaround for multilingual use. No official Chatterbox Apple Silicon memory or speed table was found.

**Evidence:** MLX-Audio installation, requirements, formats, quantization, server, and Kokoro extras: https://github.com/Blaizzy/mlx-audio and https://github.com/Blaizzy/mlx-audio/blob/main/pyproject.toml. Kokoro conversion size: https://huggingface.co/mlx-community/Kokoro-82M-bf16. Chatterbox dependency pins: https://github.com/resemble-ai/chatterbox/blob/master/pyproject.toml. Chatterbox local downloads and MPS handling: https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py. Mac example and issue: https://github.com/resemble-ai/chatterbox/blob/master/example_for_mac.py and https://github.com/resemble-ai/chatterbox/issues/275. Weight sizes: https://huggingface.co/ResembleAI/chatterbox/tree/main, https://huggingface.co/ResembleAI/chatterbox/blob/main/t3_cfg.safetensors, and https://huggingface.co/ResembleAI/chatterbox/blob/11120405b7cec7936a0fde4c84c6366789a6897a/s3gen.safetensors.

**Implication:** For a deadline-bound Mac proof, Kokoro/MLX-Audio has the lower download, dependency, hardware, and troubleshooting cost. Chatterbox remains local/private after download but has a higher chance of MPS or dependency friction and a weaker current macOS performance evidence base.

### Finding 3: Invocation and Remotion integration are straightforward for Kokoro and workable but more custom for Chatterbox

MLX-Audio exposes all three useful surfaces: a CLI, a Python iterator API, and an OpenAI-compatible local REST endpoint. A proof can use the CLI directly:

```bash
mlx_audio.tts.generate \
  --model mlx-community/Kokoro-82M-bf16 \
  --text "Narration chunk" \
  --voice af_heart \
  --lang_code a \
  --speed 1.0 \
  --audio_format wav \
  --output_path ./public/audio \
  --file_prefix narration \
  --join_audio \
  --verbose
```

The generic CLI supports `--max_tokens`, `--speed`, `--audio_format`, `--output_path`, `--join_audio`, streaming, and verbose RTF/memory metrics. Kokoro's model implementation also accepts `split_pattern`, though the generic CLI does not expose a `--split_pattern` flag; explicit per-paragraph files or a small Python wrapper are clearer for controlled chunking. The local server accepts an OpenAI-shaped `/v1/audio/speech` request with a Kokoro model ID and preset voice.

Chatterbox's first-party path is Python plus included Gradio demos; no first-party general CLI or OpenAI-compatible REST server is documented in the core repository. The simplest non-cloning call is `model = ChatterboxTTS.from_pretrained(device="mps")`, `wav = model.generate(text)`, then `torchaudio.save(...)`. Omitting `audio_prompt_path` uses bundled conditionals from `conds.pt`; supplying it invokes voice conditioning/cloning. Long-form chunking, file naming, joins, retries, and duration assembly must be implemented by the caller or a third-party server, increasing proof code.

For Remotion, generate narration before rendering, copy final WAV files under `public/`, and reference them with `staticFile()`. Current Remotion docs prefer `<Audio>` from `@remotion/media`; `<Html5Audio>` remains usable and supports Chrome-readable formats, trim, volume, and playback rate. WAV is the least-friction common denominator for both candidates and avoids installing `ffmpeg` solely for TTS output. Freeze the audio asset and derive the composition duration from the accepted narration rather than generating audio during React render.

**Evidence:** MLX-Audio CLI/Python/server examples and formats: https://github.com/Blaizzy/mlx-audio. Current CLI flags and metrics: https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/generate.py. Kokoro model `split_pattern`: https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/models/kokoro/kokoro.py. Chatterbox Python usage and bundled conditionals path: https://github.com/resemble-ai/chatterbox and https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py. Remotion audio/static assets: https://www.remotion.dev/docs/html5-audio and https://www.remotion.dev/docs/staticfile.

**Implication:** Kokoro can be added as a reproducible pre-render asset-generation step with little adapter code. Chatterbox can feed Remotion equally well once a WAV exists, but its first-party integration needs a Python wrapper for chunking, joining, and manifest creation.

### Finding 4: Kokoro offers better proof-time pronunciation control and lower technical-name risk

Kokoro uses Misaki grapheme-to-phoneme processing and explicitly supports inline pronunciation markup in the form `[display text](/IPA/)`. The generator returns graphemes and phonemes for inspection. That creates a direct, auditable route for names such as “OpenAI,” “Codex,” “GPT-5.6,” “Remotion,” package names, acronyms, and repository identifiers: normalize number/acronym reading in the narration text, add authored IPA only where needed, generate a short pronunciation proof for every glossary term, and inspect the emitted phonemes before the full take. The exact IPA strings still require human audition; the mechanism, rather than any particular transcription, is the advantage.

Chatterbox's official controls are punctuation normalization, capitalization/emphasis, `exaggeration`, `cfg_weight`, `temperature`, `min_p`, `top_p`, and repetition penalty. Resemble says capitalization can shift emphasis. No first-party SSML, IPA override, phoneme dictionary, or substitution lexicon is documented for the open-source model. Practical fallbacks are orthographic respelling (“G P T five point six”), punctuation, capitalization, and selective chunk rewrites, followed by audition. Its punctuation normalizer replaces colons, semicolons, ellipses, and some dashes, so script punctuation is not preserved literally.

Kokoro's technical-name risk is therefore lower even if Chatterbox's unconstrained speech can be more expressive. A public demo containing product names should prioritize repeatable pronunciation over maximal emotional range.

**Evidence:** Kokoro inline IPA, returned phonemes, and Misaki use: https://huggingface.co/hexgrad/Kokoro-82M and https://github.com/hexgrad/kokoro. MLX-Audio Kokoro/Misaki requirements: https://github.com/Blaizzy/mlx-audio. Chatterbox controls and punctuation normalization: https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py. Capitalization and expression claims: https://www.resemble.ai/learn/models/chatterbox.

**Implication:** Select Kokoro when the narration contains brand, model, code, and package names. Use Chatterbox only after a term-by-term audition proves acceptable pronunciation, because the open-source interface has no documented exact-pronunciation override.

### Finding 5: Voice, accent, pacing, and long-form behavior differ substantially

Kokoro v1.0 documents 54 presets across American and British English plus Japanese, Mandarin, Spanish, French, Hindi, Italian, and Brazilian Portuguese. For English proof narration, the official voice sheet grades `af_heart` A and `af_bella` A-, while the strongest documented British preset, `bf_emma`, is B-. These are author-supplied data-quality estimates, not independent listening scores. The voice sheet warns that most voices perform best at 100–200 tokens, may be weak below 10–20 tokens, and may rush beyond 400 tokens, with about 500 tokens possible. For a 350–500-word script, use roughly 4–6 semantically complete chunks in the 100–180-token range, keep headings/one-liners attached to nearby prose, and join WAVs with controlled pauses. Kokoro exposes a direct `speed` scalar; prosody otherwise comes from punctuation, voice selection, and the model.

Original Chatterbox is a 500M-parameter English model with one bundled default conditioning voice plus optional zero-shot voice conditioning from a 5–20-second reference clip. The family adds 23-language multilingual models and regional single-language packs, but their principal differentiation is still voice conditioning and language-specific checkpoints rather than a Kokoro-style list of named preset narrators. `exaggeration` controls intensity; `cfg_weight` affects pace and conditioning; Resemble recommends defaults of 0.5/0.5 and says lower CFG around 0.3 can slow a fast reference while higher exaggeration tends to speed speech. The core source samples speech tokens with a fixed `max_new_tokens=1000`; this is not documented as a safe text-character limit and is not a substitute for long-form chunking. No official open-source long-narration chunker or Mac-safe chunk size was found. Use paragraph/sentence chunks and preserve one parameter set across them.

**Evidence:** Kokoro voices and chunk limits: https://huggingface.co/hexgrad/Kokoro-82M/blob/main/VOICES.md. Kokoro speed and splitting: https://github.com/hexgrad/kokoro and https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/models/kokoro/kokoro.py. Chatterbox voice usage, control tips, and family languages: https://huggingface.co/ResembleAI/chatterbox and https://www.resemble.ai/learn/models/chatterbox. Chatterbox sampling and token cap: https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py.

**Implication:** Kokoro provides a safer preset/accent choice and explicit chunk guidance for the three-minute script. Chatterbox gives richer expressive controls, but long-form continuity and non-cloned accent identity need more manual proving.

### Finding 6: Kokoro is more reproducible; Chatterbox is intentionally stochastic

The MLX Kokoro implementation is a direct phoneme-to-duration/prosody/vocoder path. It exposes no sampling temperature or seed in the Kokoro model and no random operation was found in the generation path; fixed model revision, voice tensor, normalized text/phonemes, speed, and software versions should therefore produce functionally repeatable audio on the same environment. Bit-for-bit identity across MLX versions, quantization levels, or different Apple chips is not promised and remains unsupported.

Chatterbox samples speech tokens with temperature, top-p, min-p, and repetition penalty. The core `generate()` signature has no seed parameter. An external `torch.manual_seed()` may improve repeatability, but the project does not promise deterministic MPS inference, and the PerTh watermark is applied after waveform generation. Multiple takes are part of the expected workflow. For a proof asset, record the exact model revision, package version, parameters, script chunks, accepted-take hashes, and final join manifest.

**Evidence:** Kokoro direct model path and exposed parameters: https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/models/kokoro/kokoro.py. MLX-Audio CLI has no `--seed` flag: https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/generate.py. Chatterbox sampling parameters, absent seed argument, and watermark application: https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py.

**Implication:** Kokoro better supports a one-command reproducible proof. Chatterbox should be treated as a take-selection system, with accepted assets checked into or otherwise frozen for Remotion.

### Finding 7: Chatterbox has stronger expressiveness evidence, but it does not validate the no-cloning Mac case

Resemble reports a Podonos blind comparison in which 63.75% of evaluators preferred Chatterbox over ElevenLabs. This is first-party-selected evidence hosted through an external evaluator and used 7–20-second reference clips, so it supports Chatterbox's voice-conditioned naturalness but does not establish the quality of its bundled no-reference voice, macOS MPS behavior, or a 350–500-word narration. Chatterbox also has explicit emotion intensity and CFG pacing controls, making it the more likely winner for dramatic, conversational delivery when a consented reference voice is permitted.

Kokoro's official card claims quality comparable to larger models, while its author-supplied voice sheet identifies `af_heart` and `af_bella` as strongest English presets. A third-party model card for equivalent Kokoro BF16 MLX weights reports 41× real-time on an M3 Max for a different Kokoro package; this is useful evidence that the model class can be extremely fast on high-end Apple Silicon, but it is not an MLX-Audio benchmark and should not be used to promise performance on another Mac. MLX-Audio itself prints RTF and peak memory during verbose generation, enabling a local proof benchmark later.

Likely subjective ranking, explicitly inferred rather than measured here:

1. **Naturalness/expressiveness with consented cloning:** Chatterbox likely leads because of reference conditioning and intensity/CFG controls.
2. **Naturalness under a strict no-cloning gate:** Unresolved; Kokoro's best presets are safer and better documented, while Chatterbox's bundled voice lacks a directly applicable blind evaluation.
3. **Technical-name reliability and editability:** Kokoro leads because of inspectable phonemes and inline IPA overrides.
4. **Cross-chunk consistency and repeatability:** Kokoro likely leads because it is non-sampling and uses fixed presets.

**Evidence:** Chatterbox benchmark design/results and controls: https://www.resemble.ai/learn/models/chatterbox and https://github.com/resemble-ai/chatterbox. Kokoro quality/voice claims: https://huggingface.co/hexgrad/Kokoro-82M and https://huggingface.co/hexgrad/Kokoro-82M/blob/main/VOICES.md. Practitioner-only Apple benchmark: https://huggingface.co/sonic-speech/kokoro-82m-bf16. MLX-Audio emitted benchmark metrics: https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/generate.py.

**Implication:** Do not select Chatterbox solely on the 63.75% result for this proof; that evidence depends on cloning. Kokoro is the lower-risk no-cloning choice even if a consented Chatterbox clone could sound more expressive.

### Finding 8: Chatterbox embeds a watermark; Kokoro documents no audio watermark or visible tag

Chatterbox applies Resemble's PerTh watermark to every generated waveform. The repository includes extraction code, and Resemble describes the mark as imperceptible, robust to MP3 compression and common editing, and intended to remain attributable. Resemble explicitly says removing it is against intended use. It does not add an audible announcement or visible label to the video; the uploader remains responsible for disclosures.

No Kokoro or MLX-Audio first-party source inspected documents a synthetic-audio watermark, acoustic tag, or automatic metadata label. The absence of a documented marker is not proof that every conversion or downstream tool is marker-free, but the selected official MLX conversion has no stated watermark behavior.

YouTube requires disclosure when altered/synthetic content is realistic and meaningfully could mislead; it explicitly treats someone else's cloned voice as disclosure-worthy while saying cloning one's own voice for voiceovers ordinarily does not require disclosure. A generic preset TTS narrator is not squarely resolved by the examples. The conservative public-demo approach is to select “altered or synthetic” during upload and add a plain description note (“Narration generated locally with Kokoro-82M; no human voice was cloned”). YouTube says disclosure itself does not limit audience or monetization eligibility.

**Evidence:** Chatterbox watermark claims and implementation: https://github.com/resemble-ai/chatterbox, https://www.resemble.ai/learn/models/chatterbox, and https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py. Kokoro and MLX-Audio inspected surfaces: https://huggingface.co/hexgrad/Kokoro-82M and https://github.com/Blaizzy/mlx-audio. YouTube disclosure policy: https://support.google.com/youtube/answer/14328491.

**Implication:** Chatterbox provides built-in provenance but does not eliminate consent/disclosure duties. Kokoro needs explicit human-authored provenance text. For either model, disclose synthetic narration conservatively in the public demo description.

### Finding 9: Monetary cost is effectively zero; setup cost strongly favors Kokoro

For a single local three-minute narration, both models have **$0 model/API charge** and **$0 marginal per-character fee**. There are no API keys, subscriptions, rate limits, or royalties in the documented local workflows. Electricity, internet bandwidth, developer time, and existing Mac ownership are real but not measurable from public sources and are excluded from the monetary total.

Calculated setup burden:

| Factor | Kokoro via MLX-Audio | Chatterbox original English |
|---|---|---|
| Model artifacts | 355 MB BF16; smaller quantized variants documented | About 3.20 GB core English weights; current shared model repository is larger because it also contains multilingual variants |
| Model-footprint ratio | 1× | Approximately 9× Kokoro BF16 |
| Runtime | Python 3.10+, MLX, Misaki; WAV needs no ffmpeg | Python 3.10+ (project-tested Python 3.11), pinned PyTorch/torchaudio/Transformers/Diffusers/Librosa/Gradio plus Git-sourced PerTh |
| Mac acceleration | First-class Apple Silicon requirement | MPS path exists; official Mac example/workaround and issue history indicate more friction |
| Proof adapter | CLI or small Python manifest generator | Python wrapper needed for chunks, takes, joins, and manifest |
| Estimated unattended setup risk | Low–medium | Medium–high |

Time estimates are intentionally caveated because nothing was installed: **Kokoro 15–45 minutes** on a supported Mac with working Python/`uv`; **Chatterbox 45–120+ minutes** including the much larger download and possible MPS/dependency troubleshooting. These are planning estimates derived from documented dependency and artifact differences, not measured installation times.

**Evidence:** Kokoro size and requirements: https://huggingface.co/mlx-community/Kokoro-82M-bf16 and https://github.com/Blaizzy/mlx-audio. Chatterbox weights and dependencies: https://huggingface.co/ResembleAI/chatterbox/tree/main, https://huggingface.co/ResembleAI/chatterbox/blob/main/t3_cfg.safetensors, https://huggingface.co/ResembleAI/chatterbox/blob/11120405b7cec7936a0fde4c84c6366789a6897a/s3gen.safetensors, and https://github.com/resemble-ai/chatterbox/blob/master/pyproject.toml. Free/commercial positioning: https://www.resemble.ai/learn/models/chatterbox and https://huggingface.co/hexgrad/Kokoro-82M.

**Implication:** Monetary cost does not distinguish the finalists. Setup time, download size, and troubleshooting exposure make Kokoro the materially cheaper proof candidate.

### Finding 10: Recommendation and exact no-cloning proof-generation gate

**Recommendation: use `mlx-community/Kokoro-82M-bf16` through MLX-Audio with `af_heart` as the first proof voice and `af_bella` as the only planned fallback.** This is the best privacy-conscious proof candidate because it uses fixed named presets, has permissive code/weight licenses, has no required reference audio, runs through an Apple-Silicon-native stack, is roughly one-ninth the Chatterbox English model footprint, offers inline IPA and inspectable phonemes for technical names, accepts direct speed control, has explicit chunk-length guidance, produces Remotion-ready WAV, and is functionally more deterministic.

Use Chatterbox only as a secondary audition if the preset Kokoro takes fail a human naturalness review. In that secondary proof, call `model.generate(text)` with the bundled default conditionals and never pass `audio_prompt_path`. Do not rely on Chatterbox's voice-cloning quality benchmark to justify the no-cloning bundled voice.

**Precise no-cloning proof-generation gate — all conditions must pass before the narration may enter Remotion:**

1. **Pinned provenance:** Record the exact MLX-Audio version/commit, exact model repository and revision/hash, `af_heart` voice name/hash if available, language code `a`, speed, output format, and script checksum in a generation manifest.
2. **Preset-only invariant:** The generation command/code contains `voice="af_heart"` or approved fallback `af_bella`; it contains no `ref_audio`, `audio_prompt_path`, speaker embedding, user-supplied voice tensor, voice-mixing input, or third-party audio file. Fail closed if any reference-audio path is present.
3. **No identity claim:** Project copy and video metadata must not name or imply a real person as the narrator. The selected voice must be described only by its Kokoro preset identifier.
4. **Rights-cleared text:** The narration script is original or permissioned and contains no unlicensed quoted performance. This satisfies the Build Week original-work/third-party-rights requirement independently of the model license.
5. **Technical glossary proof:** Before full generation, render every brand/model/package/acronym sentence separately. Normalize acronyms and version numbers in visible script text, use Kokoro's inline IPA only for failed terms, inspect returned phonemes, and require a human listener to approve “OpenAI,” “Codex,” “GPT-5.6,” “Remotion,” and every project-specific proper noun.
6. **Safe chunking:** Split the final 350–500 words into semantically complete chunks targeted at 100–180 tokens and never exceed 200 tokens without an explicit audition. Avoid isolated clips below 10–20 tokens; attach short calls-to-action to adjacent prose. Preserve the same voice/speed/version across every chunk.
7. **Take acceptance:** Human-review each chunk for mispronunciation, skipped/repeated words, rushing, clipping, artifacts, and misleading emphasis. Regenerate only the failed chunk; retain the accepted WAV and its hash. The accepted joined duration must be measured and fit the approximately three-minute edit without extreme Remotion playback-rate correction.
8. **Local-only generation:** Download artifacts once, then generate by direct CLI/Python on the Mac. Do not send script or audio to a hosted TTS endpoint. If a local server is used, bind to loopback only; direct CLI is preferred.
9. **Lossless master:** Join and archive WAV at the native model sample rate. Place the frozen master under Remotion's `public/` assets, reference it with `staticFile()`, and avoid generating TTS during render. Any MP3/AAC conversion occurs only during final video encoding.
10. **License/provenance notice:** Preserve Apache-2.0 and MIT notices wherever code or model files are redistributed. Add a video-description credit identifying Kokoro-82M and MLX-Audio with their licenses and links.
11. **Synthetic-content disclosure:** Select YouTube's altered/synthetic disclosure conservatively and state: “Narration was generated locally with the Apache-2.0 Kokoro-82M preset `af_heart` through MIT-licensed MLX-Audio. No human voice was cloned.”
12. **Fallback gate:** If Kokoro fails naturalness after two voice auditions and pronunciation correction, Chatterbox may be tried only with `audio_prompt_path=None`, bundled `conds.pt`, fixed parameters, chunked text, accepted-take hashes, the PerTh watermark intact, and the same disclosure. Any request to use an external voice clip requires a separate written-consent review and is outside this proof.

**Evidence:** Recommendation inputs are established in Findings 1–9. The preset/no-reference Kokoro path: https://github.com/Blaizzy/mlx-audio and https://github.com/Blaizzy/mlx-audio/blob/main/mlx_audio/tts/models/kokoro/kokoro.py. Chatterbox's bundled-conditionals and optional reference path: https://github.com/resemble-ai/chatterbox/blob/master/src/chatterbox/tts.py. Build Week rights requirements: https://openai.devpost.com/rules. Remotion asset path: https://www.remotion.dev/docs/html5-audio and https://www.remotion.dev/docs/staticfile. YouTube disclosure: https://support.google.com/youtube/answer/14328491.

**Implication:** Kokoro/MLX-Audio should be the first and presumptive final local proof voice. The gate makes “no cloning” mechanically auditable rather than a verbal intention and keeps the final Remotion asset reproducible, rights-conscious, and publicly discloseable.

## Notes

- No local validation was performed, as required. All install-time, macOS compatibility, listening-quality, and performance conclusions remain source-backed projections until a separate proof run.
- No official MLX-Audio Kokoro benchmark table by Mac model was found. The M3 Max 41× real-time result comes from a different Kokoro MLX package using equivalent BF16 weights and is practitioner evidence only.
- No official Chatterbox Apple Silicon throughput, peak-memory, or long-narration benchmark was found. Claims of “faster than real time” and approximately 200 ms latency on Resemble's product page are not specific to the open-source original model on macOS.
- Chatterbox's 63.75% preference result used reference clips. It is unsupported as evidence for the bundled no-cloning voice.
- No separate generated-output license was found for either model. The permissive code/weight licenses do not resolve copyright, trademark, privacy, publicity, or voice-consent questions in user inputs and outputs.
- Kokoro's lack of a watermark is based on absence from inspected official documentation/source surfaces, not an acoustic forensic test.
- Exact IPA entries for project-specific terms remain authored content and require audition; only the availability of the override mechanism is validated.
- Current Chatterbox repository/model contents are moving targets (Multilingual V3 and newer weights share the repository). Pin a revision rather than relying on `master`/`main` for a proof.

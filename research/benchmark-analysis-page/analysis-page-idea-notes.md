# Analysis Page — Idea Notes

Working notes for the written analysis page: candidate structure, findings, framings, tidbits, and required caveats. Numbers reflect the corrected 19-run dataset (post attribution-matcher fix); the two in-flight Claude Code — Codex Sol runs complete that configuration and should be finalized before publication.

## The dataset in one line

19 accepted runs, 570 turns, 868 expected skill calls: 349 called, 519 missed — 40.2% overall expected-call success, with 123 additional calls excluded from the percentage.

## Thesis: the pairing is the unit of performance

Neither "which harness?" nor "which model?" predicts skill-calling alone — only the combination does. Every single-factor explanation dies against one of the four corner cells:

|                    | Strong result | Weak result       | What it kills                         |
| ------------------ | ------------- | ----------------- | ------------------------------------- |
| **Cursor harness** | Grok 94.9%    | Composer 19.0%    | "Cursor is a good/bad harness" — both |
| **Sol model**      | Codex 100%    | Claude Code ~8.7% | "Sol is a strong/weak model" — both   |

Skill-calling is a property of the model–harness pairing, not of either part.

## The four corners as characters

- **Sol on Codex (137/137, 100%, 81 additional)** — a model at home: native harness, native skill metadata (`agents/openai.yaml`), maximal call appetite, yet perfect judgment on the decoy turns.
- **Grok on Cursor (130/137, 94.9%, 19 additional)** — the campaign's most surgical performer: near-perfect coverage with a quarter of Codex's extras. Proof the Cursor wrapper "works" when paired well.
- **Composer on Cursor (26/137, 19.0%, 2 additional)** — same wrapper as Grok, collapsed engagement. It does not misfire; it does not fire.
- **Sol behind Claude Code (4/46 so far, ~8.7%)** — the same model that scored 100% in its native harness, nearly mute behind a different interface. The wrapper filters the model's entire disposition toward skills.

## Headline finding 1 — the harness effect

`gpt-5.6-sol` is held constant across five harnesses; scenarios, skills, and prompts are identical:

| Harness for the same Sol model | Success                                         |
| ------------------------------ | ----------------------------------------------- |
| OpenAI Codex (native)          | 100.0%                                          |
| OpenCode                       | 14.6%                                           |
| Pi                             | 10.9%                                           |
| Claude Code (claudex proxy)    | ~8.7% (one of three scenarios so far)           |
| Cursor                         | blocked (billing); the missing cross-check cell |

An ~11× spread from the wrapper alone. Frame the mechanism honestly as open: the harness bundles system prompting, skill surfacing, and invocation ergonomics, and this data does not isolate which component dominates — it proves the bundle does.

## Headline finding 2 — the model effect

Cursor holds the harness constant: Grok 94.9% versus Composer 19.0% in the identical wrapper. A strong model in a supportive harness is flawless; either axis alone can sink the result.

## Calling styles

- **Codex, the maximalist**: perfection through liberal calling (81 additional).
- **Grok, the marksman**: nearly the same coverage with 19 additional — the most discerning caller.
- **Composer, the abstainer**: 2 additional calls across three scenarios — it rarely calls at all.
- Within the Claude Code harness, swapping Fable (12.4%) for Sol (~8.7%) barely moves the needle — more evidence the harness sets the ceiling at the low end.

## What the low scorers actually miss

The chronic misses are the recurring discipline skills — `code-implementation-discipline`, `document-update-discipline`, and `code-testing-discipline` — which are expected at many governed points per scenario. The behavioral shape: low-band configurations treat skills as setup-time tools (a plan skill on turn 1, then silence), while Codex and Grok sustain skill use across all 30 turns. A calls-by-turn-position decay chart per configuration would make this vivid; the compact `website.json` points already contain the data.

## The validity story

- **The ceiling is reachable**: a perfect 137/137 exists, so low scores are model choices, not impossible answer sheets.
- **Decoy turns**: turns 13, 18, and 24 carry no expectations by design; the top scorers made zero calls on them in every scenario — the high scores are judgment, not marker spam.
- **The correction, disclosed**: the first derivation scored Codex at 74% because the capture matcher recognized only `echo`-form markers and missed `printf`-form ones; the matcher was fixed and results re-derived from retained evidence. Publishing the correction openly is a strength.
- **Runs cannot silently lie**: skill visibility is machine-verified in-run where the harness exposes it (Claude Code, OpenCode, Pi), and all 19 runs recorded zero unattributed calls.

## Tidbits

- Same model, same scenarios, same day: 100% and ~9%. The whole dataset in one sentence.
- Fable was the fastest runner (a 19-minute full scenario) and among the lowest scorers — fewer skill calls means less work per turn; speed anti-correlates with compliance.
- One model chained `&& mkdir` onto its instrumentation command, was denied by the deny-first sandbox, and calmly retried the exact permitted command — resilience the tooling initially misread as a failure and now classifies correctly.
- Roughly 570 turns and ~10 hours of cumulative agent runtime, landed in one afternoon of parallel lanes.

## Missing cells, framed as next work

Cursor — Sol and Cursor — Fable were blocked mid-scenario by the Cursor plan's monthly API-model allowance (resets 2026-08-21). Sol-on-Cursor is the most valuable missing cell: it would separate "Cursor pairs badly with non-native models" from "Composer is simply disengaged." Present these as the cells the next billing cycle fills in.

## Required limitations paragraph

One run per cell: descriptive counts, not rates — no statistical reliability, persistence, or universal-model claims. Reasoning was medium everywhere, but the blocked Cursor lanes would have used `-high` fallbacks (no medium variant existed on resolution day), a recorded deviation. Codex attribution is capture-based while other harnesses execute markers — different observation mechanisms with different failure modes, one of which caused the disclosed correction. One machine, one time window, one skill set, one scenario suite per cell.

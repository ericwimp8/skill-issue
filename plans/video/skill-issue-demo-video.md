# Skill Issue Demo Video

## Opening

Use the website’s light visual language and fine background grid. Centre the complete **S/** circular mark and **Skill Issue** wordmark, matching the website’s Inter-first sans-serif stack, monospaced mark, tight letter spacing, near-black text, and fine neutral border. Bring it forward with a quick scale-and-opacity pop that settles cleanly without a distracting bounce.

Build the opening hierarchy in this order:

1. **S/ Skill Issue** is the largest and strongest element.
2. **Built with Codex** appears directly underneath in small, readable monospaced text with the least emphasis.
3. **It’s not a skill issue, but it’s always a _skills_ issue.** appears below the credit. It is larger and more prominent than **Built with Codex**, and sits much closer to the wordmark in size and emphasis. Preserve the website’s italic treatment for _skills_.

Stagger the entrances across the first four seconds: pop in the wordmark, fade in **Built with Codex**, then fade and lift in the hook. Hold the complete lockup briefly, then use a gentle scale-forward crossfade into the first environment-evaluation screen by 0:06.

## Draft Transcript and Screen Plan

| Time      | Voiceover                                                                                                                                                                                                                   | What we show                                                                                                      |
| --------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| 0:00–0:06 | “Skill Issue. It’s not a skill issue, but it’s always a skills issue.”                                                                                                                                                      | Pop in the website-matched wordmark, then fade in **Built with Codex** and the larger hook before transitioning.  |
| 0:06–0:28 | “Before you rewrite a failing skill, ask a more basic question: can your model-and-harness environment discover and invoke skills reliably as the conversation grows?”                                                      | Move from the title to a missed skill call, then reveal the model-and-harness environment around it.              |
| 0:28–0:52 | “Skill Issue runs fixed, governed multi-turn scenarios through the agent you already use. It records each expected skill call, what was actually observed, and the turn where calls or misses occurred.”                    | Show the CLI launching an evaluation, followed by a thirty-turn timeline with expected calls and observed misses. |
| 0:52–1:16 | “Those results become comparable charts across scenarios, models, and harnesses. They show what happened in each completed run, without turning one campaign into a universal winner or guarantee.”                         | Show accepted result charts and filter them by scenario, harness, and model.                                      |
| 1:16–1:43 | “The open-source CLI includes three thirty-turn evaluations and can run custom skills, scenarios, and answer sheets through the same pipeline. It writes detailed local evidence and compact chart-ready results.”          | Show a built-in evaluation command, the custom-input route, and the generated `result.json` and `website.json`.   |
| 1:43–2:06 | “Once the environment is qualified, skill failures become easier to diagnose. Now you can ask whether the skill is recognized at the right time—and whether its instructions produce the intended result.”                  | Transition from an environment result into separate skill-invocation and skill-behavior checks.                   |
| 2:06–2:30 | “Start by describing the outcome in ordinary language. Skill Intake inspects the project, resolves ambiguities that affect the result, and creates a build-ready plan instead of guessing about scope.”                     | Show a natural-language request becoming the structured Skill Intake plan.                                        |
| 2:30–2:48 | “Skill Generation turns that plan into an idiomatic skill. Evaluation then tests invocation and behavior separately, traces failures to the description or instructions, and refines the meaning that actually needs work.” | Show the generated skill, separate evaluation loops, and one targeted refinement.                                 |
| 2:48–2:57 | “Skill Issue was built with Codex and GPT-5.6—and with skills used to build, evaluate, and improve skills.”                                                                                                                 | Show a brief Codex collaboration montage: planning, implementation, and evaluation evidence.                      |
| 2:57–3:00 | “Build better skills. Understand the environment. Find the real issue.”                                                                                                                                                     | Return to the Skill Issue title and hook, with the repository or website address.                                 |

## Core Story

1. Skill, model, and harness failures often look the same.
2. Skill Issue first measures whether a model-and-harness environment invokes skills when expected.
3. Governed scenarios and turn-attributed evidence make completed runs comparable without claiming universal winners.
4. The open-source CLI provides built-in and custom evaluation routes with detailed local and chart-ready output.
5. In a qualified environment, Skill Issue helps create, evaluate, and refine skills from ordinary-language requests.
6. Codex and GPT-5.6 were used to build the project, including skills that helped build and evaluate other skills.

## Claims to Replace With Final Evidence

- The website currently uses illustrative chart data. Replace it with accepted campaign artifacts before recording the final chart sequence.
- Show only model-and-harness routes qualified for the final release, even when an adapter or installation path already exists in source.
- Show skill creation through the installed `skill-intake`, `skill-generation`, and `skill-evaluation-and-refinement` skills; do not imply that the CLI has a skill-generation command.
- Show only CLI commands, downloads, and workflows that exist in the released product.
- Describe observed calls and misses without claiming universal reliability, winners, or guarantees.

---
name: skill-evaluation-and-refinement
description: Two-loop skill evaluation and semantic refinement workflow. Use when validating or improving an existing skill inside a model-and-harness environment already qualified for reliable skill invocation.
---

# Skill Evaluation and Refinement

## Establish the Campaign

1. Confirm that the target skill exists. Locate a durable environment-qualification record naming the harness surface, model, qualification date, trial method, and direct proactive-invocation evidence. Stop and explain the missing prerequisite when that record is absent or does not cover the current environment.
2. Read the complete target, including frontmatter and every referenced instruction required to understand its contract.
3. Establish the target's goal, intended use, required behavior, observable result, and contract boundaries. Follow `references/target-outcome-interpretation.md`.
4. Create an evaluation-owned campaign location using `references/campaign-record.md`. Keep fixtures, trial records, transcripts, outputs, audit findings, cycle counts, and cleanup ownership there.
5. Keep description evaluation and body evaluation separate. Record a four-trial description pass before starting the body loop. For an explicit-only target, record that description evaluation is not applicable and why.

## Evaluate the Description

1. Read `references/harness-evaluation-controls.md` and identify the current harness surface.
2. Verify that the target is implicitly invocable, independent trial agents are available, and native execution evidence can prove whether the skill loaded. Stop with the documented reason when any gate fails.
3. Derive two distinct, representative prompts from the target's intended-use boundary. Make each task naturally require the skill without naming it, requesting its invocation, quoting its description, or revealing the expected selection.
4. Prepare only the inputs those prompts require inside the campaign location.
5. Give each prompt to a fresh independent agent and retain the native skill-load or tool trace. Treat the agent's prose claim as insufficient evidence.
6. When both trials succeed, derive two different confirmation prompts and repeat with fresh agents. Four successes across the two stages pass the description.
7. On any failure, diagnose the missing or misleading meaning in the description, update that meaning without adding fixture-specific wording or widening the intended boundary, clean transient artifacts, and restart the complete two-stage protocol.
8. After five unsuccessful description rounds for the current target, pause the campaign, retain the evidence, and ask whether the user authorizes a specified number of further rounds.

## Evaluate the Body

1. Read the target campaign's `status.md`. Do not create body fixtures or execute body cases until the description state is `passed` with four retained trial records or `not-applicable` with an explicit-only reason. Stop when the state is missing, incomplete, failed, or blocked.
2. Turn the established target contract into observable completion criteria. Classify the evaluation surface as code, document, generated artifact, single-turn chat, multi-turn conversation, or another contract-owned surface.
3. Ask whether body refinements should be applied automatically or proposed for review before each update.
4. Read `references/semantic-refinement.md` and apply its constraints to automatic edits and human-facing proposals alike.
5. Design varied representative cases that exercise generalized behavior. Record ground truth as conditions and result characteristics rather than prescribing one exact answer where multiple answers can satisfy the contract.
6. Create the smallest meaningful connected fixtures in the campaign location. For conversational behavior, conduct the interaction through an independent agent one turn at a time and retain a verbatim transcript or equivalent native record.
7. Confirm that the harness can execute and observe the selected surface. When it cannot, provide the prepared inputs and an actionable external execution route, then wait for reliable evidence.
8. Execute every case in isolation and audit the outputs against the goal, expected result, ground truth, and completion criteria. Treat a failure as material when it violates the target's goal, preserved boundary, or required completion criterion. Pass the body only when every required case and criterion passes without a retained material failure.
9. Translate retained failures into one coherent update at the meaning that owns them. Apply it automatically or present the same governed update for approval according to the selected mode.
10. Remove evaluation-owned outputs and transient changes from the preceding cycle while preserving the target refinement, campaign record, and evidence.
11. Re-run the relevant framework from clean fixtures. Count each unsuccessful cycle against the current target only.
12. After five unsuccessful body cycles, stop the entire campaign before continuing this target or another. Report the retained failures and ask how many additional cycles the user authorizes.

## Conclude the Target

- Record the four-trial description evidence and the behavior audit evidence when both loops pass.
- Record the exact gate, retained evidence, cycle count, and user-controlled stopping state when either loop cannot pass.
- Reset the body-failure counter only after the current target passes, then begin the next target from a new campaign location.
- Preserve claims at the evidence level actually established; configuration support is not measured invocation reliability.

## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/target-outcome-interpretation.md`: Target-contract interpretation rules for deriving observable success without changing a skill's meaning. Use when establishing what a skill is intended to do.
- `references/campaign-record.md`: Evaluation campaign structure, evidence records, counters, and cleanup ownership. Use when creating, resuming, or concluding any target campaign.
- `references/harness-evaluation-controls.md`: Harness invocation, independent-agent, evidence, and packaging capability matrix. Use when gating or configuring description and conversational evaluation.
- `references/semantic-refinement.md`: Generalized semantic update constraints for skill descriptions and bodies. Use when diagnosing failures or formulating automatic and human-reviewed refinements.

# Skill Evaluation and Refinement A-to-B Plan

> **Document status — completed historical supporting plan.** This plan remains the intent record for the implemented two-loop evaluator. The qualified Codex production campaign is recorded under `evaluations/skill-system-production-refinement/`; the future cross-harness skill-calling system belongs to Work Block 1 of the [six-block completion plan](../skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md). Consult the [documentation authority map](../skill-issue-project-completion/document-authority-and-update-map.md) before treating its original harness assumptions as current.

## A: Current Position

- Three local skill copies are available as test inputs to this work.
- The model and agent harness being used have already been evaluated and shown to invoke skills reliably; this work assumes that capability rather than evaluating the environment itself.
- Skill evaluation has been performed manually and requires the same work to be repeated for each skill.
- Skill Issue does not yet have a reusable extension that evaluates and refines both a skill's proactive invocation and its task behavior.
- Skill evaluation and refinement is the first major part of the core workflow, before skill intake and skill generation are built.

## B: Desired Position

Skill Issue has a reusable extension, primarily composed of curated Markdown guidance, that takes existing skills through two distinct refinement loops: first establishing that each description causes correct proactive invocation, then establishing that each body performs the intended task correctly. Both loops diagnose failures semantically and refine the appropriate part of each skill until it meets the expected standard.

## Path from A to B

### 1. Establish the Target Skills' Intended Outcomes

Read each complete target skill and establish its goal, intended use, and expected result. Use an explicitly stated goal when present and derive the intended outcome from the skill as a whole when it is absent.

### 2. Establish the Description Evaluation and Refinement Loop

Define an evaluation that determines whether each target skill's description causes the already-qualified model and harness to invoke it proactively, then diagnose failed trials and refine the description until it meets the expected standard.

### 3. Establish the Skill-Behavior Evaluation and Refinement Loop

Derive evaluation frameworks from each target skill's intended outcome, execute representative artifact or conversational cases, audit observable results, refine failures through generalized semantic updates, and repeat with cleanup and user-controlled safety limits until the skill meets the expected standard.

## C: Completion Criteria

- Existing skills' intended outcomes can be established even when a skill does not explicitly state its goal.
- Each skill's description is evaluated independently from its body.
- Each description is refined until its skill is invoked proactively in the situations where it should be used on the qualified model and harness.
- The skill body is evaluated through a representative task whose behavior or artifacts make success observable.
- Failed behavior is traced to its semantic cause and corrected without proximal patching or diffusing the meaning of unrelated parts of the skill.
- The user can choose automatic refinement or review proposed semantic changes before they are applied.
- Artifacts from the preceding behavior evaluation are removed before the same evaluation is repeated.
- Automatic behavior refinement tracks failures per target skill; five failures on the current target pause the campaign and require the user to direct any further work.
- The skill-behavior loop ends with an audit confirming that the skill performs its intended task to the expected standard.
- The resulting evaluation and refinement extension is reusable for the later skill-intake and skill-generation work.

## Unresolved Matters

- The evidence required to prove proactive invocation still needs to be defined for the primary and tentative harness/model combinations.
- The audit standard that determines whether skill behavior is good enough still needs to be defined.
- The description loop's exhaustion behavior still needs to be defined when its expected standard cannot be reached.

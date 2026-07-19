# Task 2: Establish the Skill Intake Planning Loop

## A: Starting Position

- Task 1 provides source-backed explicit-invocation requirements for the supported harnesses.
- A user may begin with a detailed skill request or may explicitly invoke intake before describing what they want.
- The existing `dictate-plan` pattern demonstrates how successive conversational input can become a living A-to-B plan.
- Skill Intake and Skill Generation are separate workstreams: intake owns discovery and planning, while generation owns creating the skill.

## B: Desired Outcome

An explicit-only `skill-intake` skill conducts a bespoke, skill-focused A-to-B planning conversation. It accepts free-form user dictation, organizes and investigates the request, resolves material ambiguity without unnecessary questioning, produces a build-ready plan, assesses generation viability, records the user's execution preference and authority boundary, and hands the result to Skill Generation.

## Path from A to B

### 1. Establish the Explicit Skill Intake Surface

Create the intake skill under the stable user-facing name `skill-intake` and apply the strongest explicit-only controls available for each supported harness, including clear invocation guidance where no enforceable control exists.

### 2. Support Both Intake Entry Patterns

Allow the user to invoke intake together with an initial description of the skill they want, or invoke intake alone and receive a concise invitation to describe it.

### 3. Establish the Skill-Focused A-to-B Model

Adapt the living A-to-B planning model to skill creation:

- **A:** the user's request, relevant existing project state, available inputs, assumptions, and external constraints;
- **B:** the desired ready-to-use skill and the outcome it is intended to create;
- **Path from A to B:** the broad, dependency-ordered work required to create that skill;
- **C:** the observable behavior and outcomes that demonstrate the finished skill does what the user intends.

### 4. Develop the Plan Through Conversational Dictation

Treat the user's successive messages as source material for one living skill plan. Organize, consolidate, and dependency-order the meaning rather than transcribing the conversation or forcing the user to structure it themselves.

### 5. Investigate Available Context

Inspect relevant source code, local documentation, configuration, tooling, and other available project material when it can answer a question, resolve a reference, or expose a material constraint on the requested skill.

### 6. Preserve the Model's Decision-Making Responsibility

Let the model organize the request, infer implementation-independent task groupings, and make decisions supported by the available evidence rather than asking the user to decide matters the project space can answer.

### 7. Clarify User-Owned Ambiguity

When a material decision remains unresolved after investigation and different answers would change the user's intended skill, formulate a focused question and integrate the answer into the plan.

### 8. Surface Material Omissions

Suggest missing outcomes, constraints, or task boundaries that would prevent the plan from being build-ready, while leaving ordinary implementation choices open for Skill Generation.

### 9. Establish Plan Readiness

Confirm that A, B, C, and the dependency-ordered path form a coherent route to a skill whose intended behavior and completion criteria are sufficiently defined for generation.

### 10. Assess Generation Viability

Determine from the plan and available capabilities whether generation is likely to succeed autonomously, can proceed autonomously while stopping for specific user-owned decisions, or will require ongoing user participation.

### 11. Present the Execution-Preference Choice

Explain the assessment plainly. When autonomous generation is viable, offer autonomous or step-by-step execution. When user help is likely to be required, identify why, explain the risk of fully autonomous execution, and still allow the user to choose whether generation should attempt it autonomously.

### 12. Record the User's Direction

Capture the user's execution preference, authority boundary, and any authorization to proceed despite identified limitations without relabeling the viability assessment.

### 13. Hand Off to Skill Generation

Provide Skill Generation with the build-ready A-to-B plan, confirmed project context, unresolved implementation-time matters, expected outcomes, completion criteria, generation-viability assessment, selected execution preference, and authority boundary without beginning skill creation inside intake.

## C: Completion Criteria

- The intake skill can begin with or without an initial skill description.
- The skill is explicit-only under every supported harness's strongest available control and fallback guidance.
- Free-form conversational input becomes a living skill-focused A-to-B plan rather than a transcript.
- The plan distinguishes current state, desired skill outcome, dependency-ordered creation tasks, and observable completion criteria.
- Available project context is investigated before the user is asked to resolve missing information.
- The model makes evidence-supported planning decisions while reserving user questions for unresolved choices that could change intent.
- Material omissions are surfaced without prescribing ordinary implementation details.
- Intake determines whether autonomous generation is viable, conditionally viable, or likely to require ongoing user help.
- The user receives a clear execution-preference choice and can authorize an autonomous attempt despite stated limitations without changing the viability assessment.
- The final handoff contains a build-ready plan and explicit interaction boundary for Skill Generation.
- Skill creation itself remains owned by Skill Generation.

## Confirmed Project Context

- The local `dictate-plan` test target provides the general living A-to-B planning pattern that this skill will adapt semantically for skill creation.
- The parent workflow defines Skill Intake before Skill Generation, supporting a planning-to-execution handoff rather than combining both responsibilities.
- The user-facing skill name is `skill-intake`.

## Confirmed Handoff Contract

Skill Generation receives the build-ready A-to-B plan, confirmed project context, unresolved implementation-time matters, expected outcomes, completion criteria, generation viability, selected execution preference, and authority boundary defined by intake. Generation executes that contract and hands the resulting skill to Skill Evaluation and Refinement.

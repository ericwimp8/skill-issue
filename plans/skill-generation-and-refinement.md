# Skill Generation and Refinement Plan

> **Document status — completed historical parent plan.** This plan owns the original evaluation → intake → generation build sequence and remains evidence for the implemented production skills. Current project sequencing, support scope, installer work, evaluation campaign, release, and submission are owned by [`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md`](skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md). Use the [documentation authority map](skill-issue-project-completion/document-authority-and-update-map.md) before treating any harness or downstream delivery statement here as current.

## Goal

Build the core Skill Issue workflow by establishing skill evaluation and refinement first, then taking a user's intended outcome through intake, skill generation, and validation of the generated skill.

## Completion Criteria

- The skill evaluation and refinement process evaluates generated skills and iteratively improves them.
- The skill intake process gathers and clarifies everything required to build the intended skill.
- Skill generation executes the build-ready intake plan under its assessed viability, the user's selected execution preference, and the recorded authority boundary, producing an idiomatic skill for the user's project and agent environment.
- The generated skill enters evaluation and refinement until it meets the expected standard or reaches a user-controlled stopping state.
- The supporting plans define the responsibilities and handoffs required for implementation.

## Scope

### 1. Skill Evaluation and Refinement

Establish the process that evaluates generated skills and iteratively improves them until they meet the expected standard.

[Skill Evaluation and Refinement A-to-B Plan](skill-evaluation-and-refinement/skill-evaluation-and-refinement-a-to-b-plan.md)

### 2. Skill Intake

Gather and clarify the information required to build the intended skill.

[Skill Intake A-to-B Plan](skill-intake/skill-intake-a-to-b-plan.md)

### 3. Skill Generation

Execute the build-ready A-to-B plan produced by `skill-intake`, using its confirmed project context, expected outcomes, completion criteria, unresolved implementation-time matters, generation viability, selected execution preference, and authority boundary. Hand the resulting skill to Skill Evaluation and Refinement rather than creating a separate planning layer for generation.

## Deferred Work

- Environment evaluation
- Agent harness evaluation
- Model evaluation
- Comparative scoring and reporting for environments, harnesses, or models

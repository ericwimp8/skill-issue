# Skill Intake A-to-B Plan

> **Document status — completed historical supporting plan.** This plan remains the intent record for the implemented Skill Intake workflow. Current harness support, native paths, and project sequencing are owned by the [project support contract](../skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md), [direct-install architecture](../skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md), and [six-block completion plan](../skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md). Consult the [documentation authority map](../skill-issue-project-completion/document-authority-and-update-map.md) before updating downstream references.

## A: Current Position

Skill Issue identifies skill intake as a core part of its workflow, but it does not yet define a dedicated user-facing intake process or the skill that conducts it.

## B: Desired Position

A user can manually invoke an explicit-only skill-making skill, describe the skill they want in ordinary language, and reach a build-ready understanding of that skill. The agent organizes the request and makes decisions supported by available context while seeking user input before unresolved assumptions can materially change the intended skill.

## Completion Criteria

- The user can explicitly start the intake workflow and is prompted to describe the skill they want.
- The skill is prevented from being invoked implicitly across the supported agent harnesses wherever their skill systems provide a control for doing so.
- Explicit-only use is also communicated through the skill's invocation guidance for harnesses without an enforceable invocation control.
- The workflow captures enough information to establish the intended skill's purpose and expected outcome.
- The agent investigates relevant source code, local documentation, and other available project context before asking the user for clarification.
- Straightforward requests can proceed without unnecessary questioning.
- Material ambiguities that cannot be resolved from available context are converted into focused questions for the user.
- The model retains responsibility for organizing and generating the result without making unsupported decisions about user intent.
- Intake assesses generation viability separately from the user's autonomous or step-by-step execution preference and records the authority boundary.
- Intake concludes with a build-ready A-to-B plan that the separate Skill Generation workflow can execute without rediscovering the user's intent or known project context.

## Dependency-Ordered Tasks

### 1. Establish Cross-Harness Explicit-Invocation Requirements

Determine how the supported skill-capable agent harnesses express or enforce explicit-only skill invocation, including the shared skill-standard mechanisms and the fallback required where no enforceable control exists.

### 2. Establish the Skill Intake Planning Loop

Create an explicit-only, skill-focused A-to-B planning loop that accepts conversational dictation, investigates available context, resolves material gaps with the user, establishes broad skill-building tasks and completion criteria, assesses the appropriate user-involvement mode, and hands a build-ready plan to the separate Skill Generation workflow.

## Confirmed Workflow Boundary

Skill Intake owns conversational discovery, project-informed clarification, the build-ready A-to-B plan, and the autonomy assessment. Skill Generation owns executing that plan and creating the skill.

## Confirmed Harness Scope

Codex, Claude Code, Pi, and OpenCode form the required initial support set. Research may add other harnesses when their inclusion and explicit-invocation controls can be supported by authoritative evidence.

# Task 1: Establish Cross-Harness Explicit-Invocation Requirements

## A: Starting Position

- Skill Issue intends to provide an explicitly invoked skill-intake workflow across multiple skill-capable agent harnesses.
- Codex, Claude Code, Pi, and OpenCode are guaranteed members of the initial supported harness set.
- Additional harnesses may be supported when research establishes that they fit the project and their skill controls can be documented reliably.
- The explicit-only invocation controls and fallback behavior for those harnesses have not yet been verified or documented.
- A deep-research capability is available to perform the required source-backed investigation.

## B: Desired Outcome

A source-backed research document establishes how Codex, Claude Code, Pi, and OpenCode control explicit-only skill invocation, identifies any additional harnesses suitable for initial support, and documents the same requirements for them. The structure remains extensible so later project contributors can add support and evaluation evidence for other harnesses without redefining the initial core set.

## Path from A to B

### 1. Establish the Core and Extensible Support Scope

Treat Codex, Claude Code, Pi, and OpenCode as required research targets. Define the category of additional skill-capable harnesses to investigate and the evidence needed to decide whether they should receive initial support.

### 2. Research Additional Support Candidates

Use a deep-research workflow to identify any harnesses beyond the four required targets that should receive first-release support, based on current source-backed evidence rather than an assumed inventory.

### 3. Research Explicit-Only Invocation Controls

For each required and research-selected harness, determine how its skill system expresses, enforces, or communicates that a skill may be used only after an explicit user request.

### 4. Establish the Shared-Standard and Fallback Behavior

Determine which explicit-invocation requirements can be expressed through shared skill conventions and what invocation guidance is required where a harness provides no enforceable control.

### 5. Distinguish Confirmed and Unsupported Capabilities

Record which controls are verified, which are unavailable, and which cannot be established from authoritative evidence so later implementation does not treat assumptions as platform support.

### 6. Produce the Cross-Harness Research Document

Create a concise, source-backed reference that names the supported harnesses, documents their explicit-invocation mechanisms, links to authoritative evidence, states the required fallback for each unsupported control, and provides a consistent shape for adding later community-supported harnesses. Label supported configuration separately from evaluated reliability so support claims do not imply unperformed harness/model evaluation.

### 7. Derive Intake-Skill Requirements

Translate the verified findings into platform-specific and shared requirements that Task 2 can use when creating the skill-intake entry point.

## C: Completion Criteria

- Codex, Claude Code, Pi, and OpenCode are included in the initial supported harness set.
- Any additional initial harnesses are selected from source-backed research.
- Each selected harness has a verified account of how explicit-only skill invocation is enforced or communicated.
- Shared skill-standard controls are distinguished from harness-specific controls.
- Harnesses without enforceable controls have a defined invocation-guidance fallback.
- Unsupported or unverified capabilities remain visibly distinguished from confirmed support.
- The research document links each material platform claim to authoritative evidence.
- Supported configuration and evaluated reliability are represented separately, with absent evaluation evidence stated plainly.
- The reference can be extended with later harness support without treating the four initial targets as the project's permanent limit.
- The findings are expressed as requirements ready for the intake skill's creation without performing that creation inside this task.

## Boundary With Task 2

This section owns selecting the harnesses, researching their controls, and producing the requirements. Task 2 owns creating the intake skill and applying those requirements.

Comparative testing of how reliably harness/model combinations invoke skills belongs to the separately deferred environment, harness, and model evaluation work. Expected differences between Codex, Claude Code, Pi, OpenCode, Claude models, and GPT models do not become findings in this intake-support research unless later evaluation evidence establishes them.

## Unresolved Matters

- The criteria for adding harnesses beyond Codex, Claude Code, Pi, and OpenCode will be established as part of the research rather than fixed during planning.
- The specific deep-research execution surface may be the available deep-research skill or ChatGPT, provided the final findings remain source-backed.

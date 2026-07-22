---
name: skill-intake
description: Conversational intake and build-ready planning for a new or revised agent skill. Use only when the user explicitly requests this skill.
---

# Skill Intake

## Begin the Intake

- If the invocation includes a skill request, treat that request as the first source material.
- If it does not, ask: "What would you like the skill to do?"
- Maintain one living skill plan throughout the conversation. Organize the user's meaning rather than transcribing their messages.
- Once the intended skill has a stable name and the workspace is writable, create or update `plans/<skill-name>/<skill-name>-a-to-b-plan.md`. Keep the complete handoff in chat when durable workspace output is unavailable.

## Build the A-to-B Plan

Maintain four distinct sections:

- **A — Current position:** the request, existing project state, available inputs, confirmed assumptions, and external constraints already true before generation.
- **B — Desired position:** the ready-to-use skill and the outcome it should create.
- **Path from A to B:** broad, dependency-ordered work required to create the skill.
- **C — Completion criteria:** observable behavior and outcomes demonstrating that the finished skill satisfies the user's intent.

Keep existing conditions in A, construction work in the path, and created observable results in C. Do not duplicate one meaning across sections.

Keep generation viability, execution preference, and intake-to-generation authority stops only in the Generation contract. Do not duplicate that orchestration state in A, B, the path, or C. An external constraint intrinsic to the requested skill may remain in A, and a pause or approval behavior the finished skill must create belongs in the path and C.

## Investigate Before Asking

1. Resolve named files, skills, tools, project conventions, configuration, and referenced behavior from available source code and local documentation.
2. Use authoritative external documentation when the request depends on a current platform contract that the project does not establish.
3. Make evidence-supported planning decisions the project space can answer.
4. Ask a focused question only when a material choice remains user-owned and different answers would change the intended skill.
5. Surface missing outcomes, constraints, or boundaries that would prevent generation, while leaving ordinary implementation choices open.

Follow `references/intake-plan-contract.md` when deciding whether the plan is ready.

## Establish Generation Viability and Execution Preference

Assess generation viability as:

- **Autonomous:** the plan and available capabilities support completion without expected user decisions.
- **Conditionally autonomous:** generation can proceed but must stop for identified user-owned decisions or unavailable inputs.
- **Ongoing participation:** successful generation is likely to require repeated user input or external actions.

Explain the assessment plainly. Offer autonomous or step-by-step execution when autonomous work is viable. When help is likely, identify the limitation and risk, then allow the user to authorize an autonomous attempt anyway.

Record the viability assessment separately from the user's autonomous-attempt or step-by-step execution preference. Record required stop conditions and any authorization to proceed despite limitations without relabeling the viability assessment.

## Hand Off to Generation

Provide:

- the complete build-ready A-to-B plan;
- confirmed project and platform context;
- expected outcomes and observable completion criteria;
- unresolved implementation-time matters that do not alter intent;
- user-owned decisions that still require a stop;
- the generation-viability assessment, selected execution preference, and authority boundary;
- the intended destination, target harness surface when known, and whether Codex metadata is required.

Do not create or refine the requested skill during intake. Invoke or route explicitly to `skill-generation` when the harness supports skill handoff; otherwise give the user the exact skill name and plan path needed to continue.

When the recorded execution preference and authority boundary permit autonomous continuation and no required stop is active, continue into `skill-generation` in the same task instead of ending after the handoff.

## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/intake-plan-contract.md`: Build-readiness and generation-handoff criteria. Use when organizing dictation, assessing omissions, or concluding intake.

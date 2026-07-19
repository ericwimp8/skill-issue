# Task 2: Establish the Description Evaluation and Refinement Loop

## A: Starting Position

- Task 1 provides local target skills whose intended outcomes can be established before evaluation begins.
- The model and harness are assumed to have already been qualified as generally capable of proactive skill invocation.
- The initial primary environments are Codex with a GPT model and Claude Code with a Claude model.
- Pi and OpenCode are tentative additional harnesses, each to be considered with both Claude and GPT models.
- Harnesses differ in how they identify skills as implicitly invocable and how they create sub-agents.
- Description evaluation is meaningful only when the target skill is implicitly invocable and the current harness can run independent sub-agents.

## B: Desired Outcome

A reusable, harness-aware loop can determine whether a target skill's description causes independent sub-agents to invoke it proactively, diagnose failed trials, refine the description semantically, and repeat until it passes. The loop identifies unsupported conditions before running, uses unleading representative cases, cleans up its artifacts, and applies a two-stage confirmation protocol.

## Path from A to B

### 1. Establish the Supported Harness Set

Use Codex with GPT and Claude Code with Claude as the initial primary harness/model combinations. Include Pi and OpenCode with both Claude and GPT as tentative combinations whose support depends on establishing the required invocation and sub-agent controls.

### 2. Establish the Implicit-Invocation Reference

Create a curated reference document that identifies Codex, Claude Code, Pi, and OpenCode and explains how to determine whether a skill is implicitly invocable in each environment. Distinguish confirmed primary support from tentative support and include concise official documentation links or verified platform-specific guidance where appropriate.

### 3. Identify the Current Harness

Determine which harness is executing the evaluation and resolve its implicit-invocation rules through the curated reference.

### 4. Gate Evaluation on Implicit Invocability

Check the target skill under the current harness's rules. Do not run proactive-invocation evaluation when the skill is explicitly restricted, the harness is unsupported, or implicit invocability cannot be established; report why the evaluation cannot proceed.

### 5. Establish the Sub-Agent Controls Reference

Create a second curated reference document that explains how Codex, Claude Code, Pi, and OpenCode create and configure the independent sub-agents required for evaluation, distinguishing confirmed primary support from tentative support.

### 6. Gate Evaluation on Sub-Agent Availability

Confirm that the current harness supports sub-agents and is configured to run them. When that capability is unavailable, tell the user that description evaluation cannot proceed because independent sub-agents perform the trials.

### 7. Design Unleading Trigger Patterns

Derive two distinct task patterns from the target skill's intended use. Each pattern must naturally require the skill while avoiding the skill's name, direct invocation instructions, or other wording that would lead the sub-agent to select it.

### 8. Prepare Evaluation Inputs and Cleanup Ownership

Create any files or other inputs required by the trigger patterns in an evaluation-owned temporary or artifact location, identify everything created by the evaluation, and establish that those artifacts are removed when they are no longer needed.

### 9. Run the Initial Invocation Trials

Run each of the two trigger patterns through a separate independent sub-agent and determine from the resulting execution evidence whether each sub-agent proactively invoked the target skill.

### 10. Run the Confirmation Trials

When both initial trials succeed, design two different representative examples and run each through a fresh independent sub-agent. Treat four successful trials across the two stages as the description evaluation's passing result.

### 11. Diagnose Failed Invocation Trials

When any trial fails, use its execution evidence and the target skill's intended invocation boundary to determine why the description did not cause proactive invocation.

### 12. Refine the Description Semantically

Update the description at the meaning responsible for the failed invocation while preserving the skill's intended scope and avoiding case-specific wording that merely targets the failed example.

### 13. Repeat the Description Evaluation

Clean up evaluation-owned artifacts, run the two-stage trial protocol again with fresh sub-agents, and continue diagnosis and semantic refinement until all four trials succeed or the loop reaches its defined exhaustion condition.

### 14. Produce the Description-Loop Result

Record the passing evidence when all four trials invoke the skill proactively. When the loop exhausts without passing, retain the evidence and report that the description could not be validated under the current conditions.

## C: Completion Criteria

- The loop can identify the current harness and apply its verified implicit-invocation rules.
- A curated implicit-invocation reference covers every supported harness with concise platform-specific guidance.
- Codex with GPT and Claude Code with Claude are covered as the primary harness/model combinations.
- Pi and OpenCode are assessed tentatively with both Claude and GPT, becoming supported combinations only when their required controls can be established.
- Skills that are explicitly restricted or cannot be assessed under the current harness are stopped before proactive-invocation trials begin.
- A curated sub-agent reference covers the controls required to run independent trials on every supported harness.
- Missing or unavailable sub-agent capability produces a clear explanation that the evaluation cannot proceed.
- Each trial uses a fresh sub-agent and an unleading task that naturally requires the target skill.
- Required evaluation inputs are isolated under evaluation ownership and cleaned up afterward.
- The passing threshold requires two successful initial patterns followed by two successful confirmation examples.
- A failed trial is diagnosed from execution evidence and causes a semantic description refinement rather than a case-specific patch.
- The complete two-stage evaluation repeats after refinement until it passes or reaches its defined exhaustion condition.
- The final result preserves the evidence showing why the description passed or could not be validated.

## Unresolved Matters

- The exact execution evidence required to prove that a sub-agent proactively invoked a skill still needs to be defined for each supported harness.
- The description loop's exhaustion condition still needs to be defined.
- The exact Claude and GPT model identifiers used for each harness/model combination still need to be selected from environments already qualified for proactive skill invocation.

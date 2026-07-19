# Task 3: Establish the Skill-Behavior Evaluation and Refinement Loop

## A: Starting Position

- Task 1 provides each target skill's goal, intended use, expected behavior, expected outcome, and evaluation-ready interpretation.
- Task 2 establishes and refines each target skill's description before body evaluation begins.
- Skills can produce different observable surfaces, including code changes, document changes, generated artifacts, chat responses, and multi-turn conversations.
- Behavior refinement must avoid explicit case-specific patches and preserve the coherent meaning of the complete skill.
- Four existing disciplines provide the required generalized refinement constraints: `document-update-discipline`, `code-implementation-discipline`, `system-change-ownership`, and `skill-authoring-discipline`.

## B: Desired Outcome

A reusable loop derives an appropriate evaluation framework from a skill's intended outcome, executes representative cases, audits observable behavior, and refines the skill body through generalized semantic updates. The loop supports automatic and human-in-the-loop refinement, handles artifact and conversational skills, cleans up between cycles, and independently limits each target skill to five unsuccessful cycles before pausing the campaign for explicit user direction.

## Path from A to B

### 1. Establish the Behavior Contract

Use the target skill's goal and expected outcome to define the behavior being evaluated, the shape of a successful result, and observable completion criteria for the evaluation.

### 2. Classify the Evaluation Surface

Determine whether the target is best evaluated through code, documents, generated artifacts, single-turn chat behavior, multi-turn conversation, or another observable surface implied by the skill's contract.

### 3. Establish the User's Refinement Mode

Before body refinement begins, ask whether the user wants changes applied automatically or wants to review the failures and proposed changes before each update.

### 4. Establish the Semantic-Refinement Constraints

Make the governing constraints from `document-update-discipline`, `code-implementation-discipline`, `system-change-ownership`, and `skill-authoring-discipline` available to the refinement process, either by requiring those skills or by embedding bespoke equivalent constraints.

Apply the same constraints when generating proposed changes in chat. Human review must receive the semantically governed change that would have been applied automatically, rather than an unconstrained preliminary suggestion.

### 5. Design Representative Evaluation Cases

Create enough varied cases to exercise the target skill's generalized behavior rather than a single narrow example. Match the case shape to what the skill needs and keep each fixture only as large as required to produce meaningful connected behavior.

For a generic code-auditing skill, a representative framework may contain three modest interconnected projects in different relevant languages, such as Rust, Dart, and TypeScript, with known semantic defects that the skill should identify. This example guides the level of variety and realism without prescribing those languages for unrelated skills.

### 6. Establish Evaluation Ground Truth

For each case, record the known conditions, intended outcome, expected result characteristics, and completion criteria needed to audit the skill's response without prescribing one exact output when multiple correct outputs are possible.

### 7. Prepare Evaluation-Owned Fixtures

Create the code, documents, prompts, files, or other artifacts required by each case in an evaluation-owned location, and track everything that the evaluation creates or changes.

### 8. Establish Conversational Evaluations

When a skill acts through chat or requires multiple turns, design the representative conversation and conduct it through a sub-agent one turn at a time. Capture a verbatim transcript or equivalent execution record that preserves enough evidence to audit the full interaction.

### 9. Gate on Required Execution Capabilities

Confirm that the current harness can perform the required independent or conversational execution. When it cannot, explain the limitation and give the user an actionable way to supply the missing execution, enable the capability, or run the prepared prompt externally.

### 10. Execute the Evaluation Cases

Run the target skill against each prepared case in an appropriately isolated context and retain its observable outputs, artifacts, edits, or conversation record.

### 11. Audit the Skill's Behavior

Compare each result with the established goal, expected outcome, ground truth, result characteristics, and completion criteria. Identify what succeeded, what failed, and the semantic reason for each failure.

### 12. Formulate a Generalized Skill Update

Translate retained failures into a coherent change at the meaning that owns the problem. Reject explicit references to the evaluation fixture, narrow case-specific instructions, accumulated proximal patches, conflicting rules, or additions that create cognitive overload and semantic diffusion.

### 13. Apply or Propose the Update

In automatic mode, apply the generalized semantic update to the skill body. In human-in-the-loop mode, report the failure and present the semantically governed change for user review before applying it.

### 14. Clean the Previous Evaluation Cycle

Remove artifacts and transient changes owned by the preceding cycle while preserving the target skill refinement and evidence required to understand the campaign.

### 15. Repeat the Behavior Evaluation

Run the relevant evaluation framework again after refinement, audit the new result, and continue the semantic refinement cycle until the skill meets its completion criteria or reaches the automatic safety limit.

### 16. Enforce the Five-Cycle Safety Gate

Track unsuccessful refinement cycles separately for each target skill. When a skill passes, reset the counter before beginning the next target. When the current skill reaches five failures, stop the campaign before continuing that skill or moving to another, tell the user what remains unresolved, and defer the next action to them. If the user authorizes more attempts, ask how many additional cycles they permit before resuming.

### 17. Produce the Behavior-Loop Result

Conclude with the evidence showing that the target skill meets its expected standard, or with the retained failures and user-controlled stopping state when the loop ends without passing.

## C: Completion Criteria

- The loop derives each evaluation from the target skill's established goal and expected outcome.
- Each evaluation has observable completion criteria and appropriate ground truth without demanding one unnecessarily exact output.
- Representative fixtures exercise generalized behavior across the surfaces relevant to the target skill.
- Code, document, artifact, single-turn chat, and multi-turn conversational skills can be evaluated through an appropriate framework.
- Conversational evaluations preserve a transcript or equivalent record sufficient to audit the complete interaction.
- Missing harness capabilities produce actionable user guidance rather than an attempted evaluation that cannot provide reliable evidence.
- The user can choose automatic refinement or human review before body changes are applied.
- Automatic changes and human-facing proposals both follow the same semantic-refinement constraints.
- Skill updates are generalized, coherent, and free from fixture-specific or accumulated proximal patches.
- Evaluation-owned artifacts are removed before the next cycle while retained skill changes and campaign evidence remain intact.
- The loop gives each target skill its own five-failure allowance and resets the count after a target passes.
- Five failures on the current target pause the entire campaign before any further target work and defer the next action to the user.
- Additional cycles run only in the quantity authorized by the user.
- The final result explains the passing evidence or the retained failures and stopping state.

## Supporting Constraint Set

- `document-update-discipline`: keeps changes at the document meaning that owns them and reconciles the whole skill document.
- `code-implementation-discipline`: keeps code-related fixture and implementation changes at the behavior owner.
- `system-change-ownership`: keeps structural responsibilities at the correct system owner.
- `skill-authoring-discipline`: keeps refined skills concise, generalized, behavior-changing, and free from case-specific accumulation.

## Unresolved Matters

- Implementation must decide whether the four supporting disciplines are installed dependencies or are represented by bespoke inline constraints.
- The exact audit mechanism for each output surface will be determined while implementing its evaluation framework.
- Harness-specific execution workarounds will be determined from the capabilities available at implementation time.

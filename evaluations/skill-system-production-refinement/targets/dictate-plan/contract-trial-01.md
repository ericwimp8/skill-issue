# Dictate Plan Target-Outcome Interpretation

## Goal

Create and continuously maintain one sequential, dependency-ordered plan of broad tasks that transforms the user's current position (A) into the desired position (B) and produces observable outcomes or capabilities (C) that demonstrate B has been reached.

## Intended Use

Use when a user develops a plan through successive conversational messages. The invocation boundary is living A-to-B task sequencing from dictation; the messages are planning source material to integrate semantically, rather than text to transcribe or a request to prescribe implementation.

## Expected Behavior

- Interpret the complete dictation as one evolving A/B/path/C model, keeping preconditions and external constraints in A, construction work in the path, and newly produced observable evidence in C without duplicating facts across A and C.
- Group work into substantial, coherent tasks, split only at meaningful outcome or dependency boundaries, and order those tasks by dependency.
- Keep execution method, tooling, file edits, commands, procedural material, and other implementation choices out of the broad plan.
- Once the task has a stable semantic name, create `plans/<task-slug>/<task-slug>-a-to-b-plan.md`; after each new dictation block, reread the whole plan, integrate new meaning, consolidate repetition, replace superseded meaning, and retain unresolved matters.
- Resolve vague references only as needed to establish the planning model, present likely detail for confirmation, and add it only after confirmation. Surface ambiguity or assumptions only when they affect the model, task boundaries, outcomes, dependencies, or ordering.
- After each update, report the plan change briefly and give only the specified Expand Dictate Plan handoff. At completion, validate A, B, C, and the path as one coherent route; propose missing C items only after identifying them and add them only after user confirmation.

## Observable Expected Result

A living plan exists at the required `plans/...-a-to-b-plan.md` path and visibly contains a coherent current-state A, desired-state B, dependency-ordered sequence of broad tasks, and non-duplicated observable C criteria. Each update visibly incorporates the latest dictation into the whole document and ends with the brief change note plus the Expand Dictate Plan invitation. When the user finishes, the final plan forms a coherent route from A to B, every C item is created by the path and helps demonstrate B, and the agent waits for the user to start expansion.

## Preserved Boundaries

- Preserve the distinction among existing conditions (A), work to perform (path), and produced evidence (C); a precondition does not become C merely because the result requires or consumes it.
- Preserve broad task sequencing rather than exhaustive minute steps, detailed section planning, implementation design, or procedural instructions.
- Preserve conversational, successive-message use and whole-plan reintegration after every dictation block.
- Preserve confirmation gates for resolved vague details and proposed missing C items.
- Preserve unresolved meaning when it matters, while leaving execution-internal choices open rather than treating them as plan gaps.
- Preserve the exact post-update response boundary: a brief change explanation and Expand Dictate Plan handoff, with no added suggestions, optional next steps, or readiness quiz.
- Preserve the user's authority to indicate completion and to start the later expansion phase.

## Evidence

- `test-targets/skills/dictate-plan/SKILL.md`: frontmatter invocation boundary; `Goal`; `Planning Model`; `Living Document`; `After Each Update`; and `Completion` define the interpretation above.
- `skills/skill-evaluation-and-refinement/references/target-outcome-interpretation.md`: requires interpreting the complete target as one contract across goal, intended use, expected behavior, observable result, and preserved boundaries.

## Ambiguity

The target does not prescribe the internal heading structure or formatting of the generated plan beyond its path and required semantic model. It also does not define an exact threshold for when a task name becomes “stable” or when a task boundary is “meaningful”; those judgments must follow the stated semantic and dependency criteria without adding implementation detail.

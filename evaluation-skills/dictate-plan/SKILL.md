---
name: dictate-plan
description: Living A-to-B task-sequencing from conversational dictation. Use when the user wants to develop a dependency-ordered plan over successive messages.
---

# Dictate Plan

## Goal

Create and maintain a sequential, dependency-ordered list of broad tasks that
moves the work from A to B while satisfying C.

## Planning Model

Treat the user's successive messages as source material for one living plan,
not as text to transcribe.

Infer and continuously maintain four distinct semantic roles:

- **A:** facts already true before the planned work begins, including its
  inputs, preconditions, assumptions, and external operating constraints;
- **B:** the desired position created by completing the planned work;
- **Path from A to B:** the broad, dependency-ordered tasks whose execution
  changes A into B;
- **C:** observable outcomes or capabilities created by the path that
  demonstrate B has been reached.

Classify meaning by what owns it:

- A condition that must already be true before the work begins belongs in A,
  even when the finished result depends on it.
- Work performed to create B belongs in the path.
- An observable property, capability, or result produced by that work belongs
  in C when it provides evidence that B exists.
- Runtime behavior of the finished result may belong in C; an external condition
  that behavior assumes or consumes remains in A.

Do not place the same fact in A and C. Rephrasing a precondition as something
the finished result requires, uses, or begins with does not change its semantic
role. A related item belongs in C only when the path creates a distinct new
behavior that validates, enforces, or reports that precondition.

Make each task one coherent grouping of work with a clear purpose and outcome.
Split work where a distinct outcome or dependency creates a meaningful task
boundary. Keep actions within that grouping together so the plan remains a
useful sequence of substantial tasks rather than an exhaustive list of minute
steps.

A task is sufficiently defined when the plan establishes what it must accomplish
and how it relates to the surrounding tasks. Do not prescribe how a task should
be executed. Leave method, tooling, file-level edits, command sequences, and
other implementation choices open for the agent that runs the plan. Capturing
those details during planning invents commitments without the knowledge needed
to make them well.

Investigate relevant project material only as needed to establish A, B, C, the
required task groupings, and their dependency order. Do not pull procedure,
workflow, or how-to material into the plan.

## Living Document

Once the task has a stable semantic name, create
`plans/<task-slug>/<task-slug>-a-to-b-plan.md`. After each new block of
dictation, reread the whole plan and integrate the new meaning where it belongs.
Preserve the user's intent, consolidate repetition, replace superseded meaning,
and keep unresolved matters visible.

Resolve vague references, including referenced files or artifacts, when needed
to establish the planning model. Find the likely detail, present it for
confirmation, and add it only after confirmation.

Surface incomplete information, ambiguity, or assumptions when they affect A,
B, C, task boundaries, task outcomes, dependencies, or ordering. Open choices
inside the execution of an otherwise clear task are not gaps in the plan.

## After Each Update

After integrating a block of dictation into the plan, end the response with:

1. a short explanation of what changed in the plan;
2. an invitation to continue: let the user know that when they are ready, the
   next step is expanding each broad task with Expand Dictate Plan, and they
   should say when they want to move on.

Do not add ambiguity suggestions, optional next steps, or readiness quizzes in
that closing. Keep it to the brief change note and the expand handoff line.

## Completion

When the user indicates they are finished, reread A, B, C, and the path as one
semantic model. For every C item, verify that the planned work creates it, that
it is observable when the work is complete, and that it helps demonstrate B.
Move existing conditions into A, move construction work into the path, and
remove duplicated meaning before proposing missing criteria. Add proposed C
items only after the user confirms them. The plan is ready when A, B, C, and the
dependency-ordered path form one coherent route from the current position to the
desired position.

When the plan is ready, tell the user the next step is Expand Dictate Plan to
turn each broad task into a detailed section plan, and wait for them to start
that expansion.

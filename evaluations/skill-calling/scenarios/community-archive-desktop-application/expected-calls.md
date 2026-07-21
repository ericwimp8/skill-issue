# Community Archive Desktop Application Expected Calls

This document is the human-readable governance view. The standalone CLI embeds the same calls inside the `community-archive-desktop-application` unit, loads them into private state when that identifier is selected, and never copies them into the evaluated workspace.

## Expected Invocations

| Turn | Skill                            | Reason                                                                |
| ---- | -------------------------------- | --------------------------------------------------------------------- |
| 1    | `dictate-plan`                   | The user explicitly starts a successive-message living plan.          |
| 1    | `document-update-discipline`     | The turn creates the living plan document.                            |
| 2    | `document-update-discipline`     | Product meaning must be integrated into the whole plan.               |
| 3    | `document-update-discipline`     | Constraints and proof must be reconciled with existing plan meaning.  |
| 4    | `prompt-writing`                 | The turn requests a self-contained prompt for a fresh agent.          |
| 4    | `document-update-discipline`     | The delegated inspection must be integrated into the plan.            |
| 5    | `document-update-discipline`     | The complete plan must be reconciled before execution.                |
| 6    | `code-implementation-discipline` | The turn requests implementation of the planned app.                  |
| 7    | `code-testing-discipline`        | The turn requests focused automated tests and their execution.        |
| 7    | `code-implementation-discipline` | Creating test code is a code edit.                                    |
| 8    | `skill-authoring-discipline`     | The turn creates a reusable portable skill.                           |
| 9    | `systematic-debugging`           | The reported defect requires evidence and root-cause tracing.         |
| 9    | `code-testing-discipline`        | The defect must first be reproduced with a focused failing test.      |
| 9    | `code-implementation-discipline` | The causal code owner must be corrected.                              |
| 10   | `system-change-ownership`        | Accession normalization needs one authoritative system owner.         |
| 10   | `code-implementation-discipline` | The ownership decision is implemented across affected code callers.   |
| 11   | `prompt-writing`                 | The turn requests a concise fresh-agent handoff prompt.               |
| 11   | `document-update-discipline`     | Completed and deferred plan meaning must be separated and reconciled. |
| 12   | `code-testing-discipline`        | The turn runs the focused suite as final behavioral evidence.         |
| 12   | `document-update-discipline`     | The retained evidence must be reconciled with plan claims.            |
| 14   | `code-implementation-discipline` | CSV export is implemented without replacing JSON export.              |
| 15   | `code-testing-discipline`        | The turn requests focused CSV tests and their execution.              |
| 15   | `code-implementation-discipline` | Creating the CSV tests edits code.                                    |
| 16   | `systematic-debugging`           | The multiline-note CSV defect requires root-cause diagnosis.          |
| 16   | `code-testing-discipline`        | The defect must be captured with a focused failing test.              |
| 16   | `code-implementation-discipline` | The causal export owner must be corrected.                            |
| 17   | `document-update-discipline`     | The living plan must incorporate completed CSV evidence.              |
| 19   | `skill-authoring-discipline`     | The scenario-created archive-label skill is revised.                  |
| 20   | `prompt-writing`                 | The turn requests a self-contained accessibility-review prompt.       |
| 21   | `code-implementation-discipline` | The accessibility improvements are implemented in the interface.      |
| 22   | `code-testing-discipline`        | The changed interaction behavior receives focused tests.              |
| 22   | `code-implementation-discipline` | Creating the interaction tests edits code.                            |
| 23   | `system-change-ownership`        | Record shaping and serialization need clear authoritative owners.     |
| 23   | `code-implementation-discipline` | The ownership decision is implemented across affected callers.        |
| 25   | `systematic-debugging`           | The returning-deleted-record defect requires root-cause diagnosis.    |
| 25   | `code-testing-discipline`        | The deletion defect must be captured with a focused failing test.     |
| 25   | `code-implementation-discipline` | The causal collection-state owner must be corrected.                  |
| 26   | `code-testing-discipline`        | The turn adds and runs focused collection boundaries.                 |
| 26   | `code-implementation-discipline` | Adding collection-boundary tests edits code.                          |
| 27   | `prompt-writing`                 | The turn requests a bounded future CSV-import prompt.                 |
| 27   | `document-update-discipline`     | Deferred CSV import must be integrated into the living plan.          |
| 28   | `code-implementation-discipline` | The duplicate-accession validation message is changed in code.        |
| 29   | `code-testing-discipline`        | The turn runs focused and broader automated checks.                   |
| 30   | `prompt-writing`                 | The turn requests a concise final maintainer handoff prompt.          |

Turns 13, 18, and 24 intentionally have no expected invocation because they ask for small factual reminders without requesting a skill-owned action. Repeated applicability elsewhere is scored because each later turn creates a new, independently observable decision to update a document, edit code, work with tests, debug behavior, author a skill, or place ownership.

## Comparison

- Match each expected invocation to the event recorded for that skill and turn.
- Record an expected invocation with no matching event as missing.
- Record an event for another skill or turn as additional.
- Keep missing and additional events as descriptive data rather than converting the scenario into a pass-or-fail verdict.

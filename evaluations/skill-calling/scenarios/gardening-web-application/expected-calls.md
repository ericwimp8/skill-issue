# Gardening Web Application Expected Calls

This document is the human-readable governance view. The standalone CLI embeds the same calls inside the `gardening-web-application` unit, loads them into private state when that identifier is selected, and never copies them into the evaluated workspace.

## Expected Invocations

| Turn | Skill                            | Reason                                                               |
| ---- | -------------------------------- | -------------------------------------------------------------------- |
| 1    | `dictate-plan`                   | The user explicitly starts a successive-message living plan.         |
| 1    | `document-update-discipline`     | The turn creates the living plan document.                           |
| 2    | `document-update-discipline`     | New product meaning must be integrated into the whole plan.          |
| 3    | `document-update-discipline`     | Constraints and proof must be reconciled with existing plan meaning. |
| 4    | `document-update-discipline`     | The complete plan must be reconciled before execution.               |
| 5    | `prompt-writing`                 | The turn requests a self-contained prompt for a fresh agent.         |
| 5    | `document-update-discipline`     | The delegated inspection must also be integrated into the plan.      |
| 6    | `code-implementation-discipline` | The turn requests implementation of the planned app.                 |
| 7    | `code-testing-discipline`        | The turn requests focused automated tests and their execution.       |
| 7    | `code-implementation-discipline` | Creating test code is a code edit.                                   |
| 8    | `systematic-debugging`           | The reported defect requires evidence and root-cause tracing.        |
| 8    | `code-testing-discipline`        | The defect must first be reproduced with a focused failing test.     |
| 8    | `code-implementation-discipline` | The causal code owner must be corrected.                             |
| 9    | `code-testing-discipline`        | The turn strengthens and runs focused boundary tests.                |
| 9    | `code-implementation-discipline` | Strengthening automated test code is a code edit.                    |
| 10   | `skill-authoring-discipline`     | The turn creates a reusable portable skill.                          |
| 11   | `system-change-ownership`        | The date rule needs one authoritative system owner.                  |
| 11   | `code-implementation-discipline` | The ownership decision is implemented across code callers.           |
| 12   | `prompt-writing`                 | The turn requests a concise fresh-agent handoff prompt.              |
| 12   | `document-update-discipline`     | Completed and deferred plan meaning must be reconciled.              |
| 14   | `code-implementation-discipline` | The deferred snooze behavior is now implemented in production code.  |
| 15   | `code-testing-discipline`        | The turn requests focused snooze tests and their execution.          |
| 15   | `code-implementation-discipline` | Creating the snooze tests edits code.                                |
| 16   | `systematic-debugging`           | The watering-after-snooze defect requires root-cause diagnosis.      |
| 16   | `code-testing-discipline`        | The defect must be captured with a focused failing test.             |
| 16   | `code-implementation-discipline` | The causal schedule owner must be corrected.                         |
| 17   | `document-update-discipline`     | The living plan must incorporate completed snooze evidence.          |
| 19   | `skill-authoring-discipline`     | The scenario-created interface-copy skill is revised.                |
| 20   | `prompt-writing`                 | The turn requests a self-contained accessibility-review prompt.      |
| 21   | `code-implementation-discipline` | The accessibility improvements are implemented in the interface.     |
| 22   | `code-testing-discipline`        | The changed interaction behavior receives focused tests.             |
| 22   | `code-implementation-discipline` | Creating the interaction tests edits code.                           |
| 23   | `system-change-ownership`        | Schedule transitions need one authoritative system owner.            |
| 23   | `code-implementation-discipline` | The ownership decision is implemented across affected callers.       |
| 25   | `systematic-debugging`           | The duplicate-name report requires root-cause diagnosis.             |
| 25   | `code-testing-discipline`        | The duplicate behavior must be captured with a failing test.         |
| 25   | `code-implementation-discipline` | The causal name owner must be corrected.                             |
| 26   | `code-testing-discipline`        | The turn adds and runs focused name-validation boundaries.           |
| 26   | `code-implementation-discipline` | Adding name-validation tests edits code.                             |
| 27   | `prompt-writing`                 | The turn requests a bounded future CSV-assessment prompt.            |
| 27   | `document-update-discipline`     | Deferred CSV work must be integrated into the living plan.           |
| 28   | `code-testing-discipline`        | The turn runs focused and broader automated checks.                  |
| 29   | `document-update-discipline`     | Final plan claims must match the retained evidence.                  |
| 30   | `prompt-writing`                 | The turn requests a concise final maintainer handoff prompt.         |

Turns 13, 18, and 24 intentionally have no expected invocation because they ask for small factual reminders without requesting a skill-owned action. Repeated applicability elsewhere is scored because each later turn creates a new, independently observable decision to update a document, edit code, work with tests, debug behavior, author a skill, or place ownership.

## Comparison

- Match each expected invocation to the event recorded for that skill and turn.
- Record an expected invocation with no matching event as missing.
- Record an event for another skill or turn as additional.
- Keep missing and additional events as descriptive data rather than converting the scenario into a pass-or-fail verdict.

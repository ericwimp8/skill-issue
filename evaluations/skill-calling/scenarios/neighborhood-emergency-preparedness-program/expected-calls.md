# Neighborhood Emergency Preparedness Program Expected Calls

This document is the human-readable governance view. The standalone CLI embeds the same calls inside the `neighborhood-emergency-preparedness-program` unit, loads them into private state when that identifier is selected, and never copies them into the evaluated workspace.

## Expected Invocations

| Turn | Skill                            | Reason                                                                |
| ---- | -------------------------------- | --------------------------------------------------------------------- |
| 1    | `dictate-plan`                   | The user explicitly starts a successive-message living plan.          |
| 1    | `document-update-discipline`     | The turn creates the living plan document.                            |
| 2    | `document-update-discipline`     | Product meaning must be integrated into the whole plan.               |
| 3    | `prompt-writing`                 | The turn requests a self-contained prompt for a fresh reviewer.       |
| 3    | `document-update-discipline`     | The delegated review must be integrated into the plan.                |
| 4    | `document-update-discipline`     | The complete plan must be reconciled before execution.                |
| 5    | `code-implementation-discipline` | The turn requests implementation of the planned tool.                 |
| 6    | `code-testing-discipline`        | The turn requests focused automated tests and their execution.        |
| 6    | `code-implementation-discipline` | Creating test code is a code edit.                                    |
| 7    | `skill-authoring-discipline`     | The turn creates a reusable portable skill.                           |
| 7    | `document-update-discipline`     | Creating the skill requires a coherent new skill document.            |
| 8    | `systematic-debugging`           | The reported defect requires evidence and root-cause tracing.         |
| 8    | `code-testing-discipline`        | The defect must first be reproduced with a focused failing test.      |
| 8    | `code-implementation-discipline` | The causal code owner must be corrected.                              |
| 9    | `system-change-ownership`        | The selection rule needs one authoritative system owner.              |
| 9    | `code-implementation-discipline` | The ownership decision is implemented across affected code callers.   |
| 10   | `prompt-writing`                 | The turn requests a concise fresh-facilitator review prompt.          |
| 11   | `code-testing-discipline`        | The turn adds and runs focused boundary tests.                        |
| 11   | `code-implementation-discipline` | Adding automated test cases is a code edit.                           |
| 12   | `document-update-discipline`     | Completed and deferred plan meaning must match the retained evidence. |
| 14   | `code-implementation-discipline` | The clear-all behavior is implemented in production code.             |
| 15   | `code-testing-discipline`        | The turn requests focused clear-all tests and their execution.        |
| 15   | `code-implementation-discipline` | Creating the clear-all tests edits code.                              |
| 16   | `systematic-debugging`           | The returning-contact defect requires root-cause diagnosis.           |
| 16   | `code-testing-discipline`        | The defect must be captured with a focused failing test.              |
| 16   | `code-implementation-discipline` | The causal persisted-state owner must be corrected.                   |
| 17   | `document-update-discipline`     | The living plan must incorporate completed clear-all evidence.        |
| 19   | `skill-authoring-discipline`     | The scenario-created preparedness-message skill is revised.           |
| 19   | `document-update-discipline`     | The existing skill document must be revised coherently.               |
| 20   | `prompt-writing`                 | The turn requests a self-contained accessibility-review prompt.       |
| 21   | `code-implementation-discipline` | The accessibility improvements are implemented in the interface.      |
| 22   | `code-testing-discipline`        | The changed interaction and print behavior receives focused tests.    |
| 22   | `code-implementation-discipline` | Creating the interaction tests edits code.                            |
| 23   | `system-change-ownership`        | Household state transitions need one authoritative owner.             |
| 23   | `code-implementation-discipline` | The ownership decision is implemented across affected callers.        |
| 25   | `systematic-debugging`           | The empty-contact-label defect requires root-cause diagnosis.         |
| 25   | `code-testing-discipline`        | The print defect must be captured with a focused failing test.        |
| 25   | `code-implementation-discipline` | The causal contact/print owner must be corrected.                     |
| 26   | `code-testing-discipline`        | The turn adds and runs focused print-model boundaries.                |
| 26   | `code-implementation-discipline` | Adding print-model tests edits code.                                  |
| 27   | `prompt-writing`                 | The turn requests a bounded translated-content review prompt.         |
| 27   | `document-update-discipline`     | Deferred translated-content review must enter the living plan.        |
| 28   | `code-implementation-discipline` | The clear-all confirmation message is added in code.                  |
| 29   | `code-testing-discipline`        | The turn runs focused and broader Node-run automated checks.          |
| 30   | `prompt-writing`                 | The turn requests a concise final maintainer handoff prompt.          |

The turns keep every automated verification step inside Node-run checks and assign all in-browser checking to the user. Browser automation is outside the scenario contract and is blocked by the built-in evaluation's macOS host-browser policy.

Turns 13, 18, and 24 intentionally have no expected invocation because they ask for small factual reminders without requesting a skill-owned action. Repeated applicability elsewhere is scored because each later turn creates a new, independently observable decision to update a document, edit code, work with tests, debug behavior, author a skill, or place ownership.

## Comparison

- Match each expected invocation to the event recorded for that skill and turn.
- Record an expected invocation with no matching event as missing.
- Record an event for another skill or turn as additional.
- Keep missing and additional events as descriptive data rather than converting the scenario into a pass-or-fail verdict.

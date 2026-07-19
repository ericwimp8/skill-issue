# Gardening Web Application Expected Calls

Keep this answer sheet outside the evaluated workspace and hidden from the evaluated primary agent.

## Required First Activations

| Turn | Skill                        | Reason                                                                                                              |
| ---- | ---------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| 1    | `dictate-plan`               | Dictate Plan is manually invoked to create and maintain the living A-to-B plan.                                     |
| 1    | `document-update-discipline` | The first turn creates and writes the living planning document.                                                     |
| 11   | `prompt-writing`             | The turn requests the first self-contained prompt for a research agent.                                             |
| 25   | `skill-authoring-discipline` | The turn requests immediate creation of a reusable agent skill.                                                     |
| 30   | `system-change-ownership`    | The turn replaces the page-based architecture and requires responsibilities to be reconciled at their proper owner. |

## Later Applicable Turns

These turns exercise an already applicable skill but do not require another marker event when the harness keeps that skill loaded:

- `document-update-discipline` remains applicable throughout Turns 2–30 because the living plan is repeatedly updated.
- `prompt-writing` is applicable again on Turns 18 and 23.
- `skill-authoring-discipline` is applicable again on Turn 28.

A later reload is retained as an additional observed activation. Its absence is not a missed call after the required first activation has already occurred.

## Comparison

- Match each required first activation to the event recorded for that skill and turn.
- Record a required activation with no matching event as missing.
- Record an event for another skill or turn as additional.
- Keep missing and additional events as descriptive data rather than converting the scenario into a pass-or-fail verdict.

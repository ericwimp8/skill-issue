# Decision Log Summary Skill A-to-B Plan

## A — Current position

- Projects contain Markdown architecture decision records.
- Users need a concise summary of each record's decision, status, and consequences.
- The skill targets Codex only and requires no external tools.

## B — Desired position

A ready-to-use `decision-log-summary` skill reads selected decision records and returns a concise Markdown summary without editing source files.

## Path from A to B

1. Create the minimal Codex skill at the recorded destination.
2. Define source reading and concise decision, status, and consequence extraction.
3. Validate structure and hand off runtime cases to evaluation.

## C — Completion criteria

- The skill is valid and discoverable.
- It reports decision, status, and consequences for each selected record.
- It never edits decision records.
- It contains no unnecessary script, asset, template, or reference.

## Generation contract

- Destination: `evaluations/skill-system-production-refinement/targets/skill-generation/behavior/cycle-1/case-3/output/decision-log-summary`
- Generation viability: Autonomous.
- Selected execution preference: Autonomous.
- Authority: create only inside the destination and run structural validation.
- Supported harness: Codex.
- Evaluation handoff: retain runtime criteria for later evaluation.

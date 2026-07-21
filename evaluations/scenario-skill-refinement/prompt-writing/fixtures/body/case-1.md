# Case 1: Repair An Over-Prescribed Handoff

## Situation

An earlier implementation task failed because the recipient edited code before confirming the concrete owner. Revise the draft prompt so it fixes that decision without accumulating unrelated instructions.

## Draft Prompt

> Explore the export feature. Read every Markdown file, every Go file, all tests, all Git history, and all open issues. Produce a 12-section report with an executive summary, architecture diagram, dependency matrix, risk register, test plan, rollout plan, and changelog. Do not make mistakes. Before doing anything, list every file you will open. Then inspect the source and implement any obvious fixes you find.

## Needed Outcome

The recipient should trace the export request from entrypoint to concrete production owner and return only the ownership finding with source locations and unresolved uncertainty. It may read production source and focused Git history. It must not edit files, run tests, or propose adjacent fixes.


# Gardening Web Application Scenario Instructions

## Preparation

1. Use a clean scenario workspace with the target harness's default configuration.
2. Run the CLI from outside the workspace and select the built-in `gardening-web-application` evaluation, intended harness, model, and an output root outside the scenario workspace.
3. Do not supply scenario, answer-sheet, or scope arguments for this built-in run. The executable owns both inputs and always installs temporary skills into the harness's project skill directory.

## Run

Use `skill-issue evaluate run --output <path> --evaluation gardening-web-application` to create temporary instrumented skill copies and start one clean primary-agent session. The runner sends every embedded turn verbatim and in order, waits for the harness's terminal completion event, and never adapts later prompts to model responses.

Do not add extra skills or prompts. Allow substantial turns to finish fully. Compaction may occur naturally. Turn number is the comparison axis; token counts are optional metadata.

## Retain

- the native transcript and structured harness events when available;
- the generated A-to-B plan and scenario-created skill;
- the graph-ready result JSON;
- harness, model, version, operating system, and reasoning-setting metadata;
- any tooling error that prevented replay or marker recording.

A tooling failure requires repair and rerun. Missing skill activations remain evaluation data after a tooling-complete replay.

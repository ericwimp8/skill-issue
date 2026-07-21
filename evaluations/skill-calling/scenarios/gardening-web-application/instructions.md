# Gardening Web Application Scenario Instructions

## Preparation

1. Use a clean scenario workspace with the target harness's default configuration.
2. Run the CLI from outside the workspace and select the built-in `gardening-web-application` evaluation, intended harness, model, and an output root outside the scenario workspace.
3. Do not supply scenario, answer-sheet, or scope arguments for this built-in run. The executable owns both inputs and always installs temporary skills into the harness's project skill directory.

## Run

Use `skill-issue evaluate run --output <path> --evaluation gardening-web-application` to create temporary instrumented skill copies and start one clean primary-agent session. The runner sends every embedded turn verbatim and in order, waits for the harness's terminal completion event, and never adapts later prompts to model responses.

Do not add extra skills or prompts. Let each of the thirty turns finish fully. Turn number is the comparison axis; token counts are optional metadata.

## Retain

- the native transcript and structured harness events when available;
- the generated A-to-B plan, extended tiny application, focused tests, and scenario-created skill;
- the graph-ready result JSON;
- harness, model, version, operating system, and reasoning-setting metadata;
- any tooling error that prevented replay or marker recording.

A tooling failure requires repair and rerun. Missing skill activations remain evaluation data after a tooling-complete replay only when the replay recorded at least one attributed skill signal and the missing skill is proven visible to the harness; verify an unproven miss with a `--transcript` visibility check before scoring it, because a silently unloaded skill produces a miss identical to a genuine model choice.

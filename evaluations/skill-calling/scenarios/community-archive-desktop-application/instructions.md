# Community Archive Desktop Application Scenario Instructions

## Preparation

1. Use a clean scenario workspace with the target harness's default configuration.
2. Run the CLI from outside the workspace and select the built-in `community-archive-desktop-application` evaluation, intended harness, model, and an output root outside the scenario workspace.
3. Do not supply scenario, answer-sheet, or scope arguments for this built-in run. The executable owns both inputs and always installs temporary skills into the harness's project skill directory.
4. Confirm that authentication, model access, workspace trust, and command execution are available before treating the run as evaluation evidence.

## Replay

Use `skill-issue evaluate run --output <path> --evaluation community-archive-desktop-application` to create temporary instrumented skill copies and start one clean primary-agent session. The runner sends every embedded turn verbatim and in order, waits for the harness's terminal completion event, and never adapts later prompts to model responses.

The evaluated agent receives no answer sheet, expected-call map, scoring wording, or turn-attribution state. The CLI records opaque signals in private application state outside the workspace.

## Evidence

A tooling-complete run retains:

- `result.json` as the detailed authoritative evidence;
- `website.json` as the compact turn-level website artifact;
- `events.jsonl` as the raw signal record;
- `transcript.json` as the native replay transcript;
- the generated A-to-B plan and the scenario-created `archive-description-quality` skill in the scenario workspace when the agent completes those requested outputs.

Missing or additional skill calls remain descriptive observations. Runner, permission, marker, session, or protocol failures are tooling failures to repair and rerun.

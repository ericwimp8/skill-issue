# Work Block 1 Supporting Plan: Define the Skill-Calling Evaluation Contract and Campaign Assets

## A: Starting Position

- The completed product, support, and evidence contract defines a minimum matrix of thirteen medium-setting model evaluations across OpenAI Codex, Claude Code, Cursor, Pi, and OpenCode.
- The evaluation measures whether the active harness-and-model combination calls the supplied skills at the expected points in a realistic conversation. It does not grade the quality of the loaded skill's body in this campaign.
- The evaluation uses three substantial scripted scenarios rather than copied conversations whose later turns depend on unpredictable model responses.
- The primary and supporting skill collections own the complete canonical Skill Issue payload. The CLI generates instrumented evaluation copies from those current canonical files.
- The minimum viable campaign accepts a bespoke, human-controlled setup. Evaluators will be instructed to use a default harness environment containing only the supplied Skill Issue skills, required supporting skills, and evaluation instrumentation.

## B: Desired Position

The project has a transparent and repeatable skill-calling evaluation in which the CLI generates a project-local temporary installation from the selected skills, injects one opaque neutral signal instruction without changing selection metadata, and runs three fixed 30-turn scenarios against each harness's primary agent. Each governed evaluation is embedded as one unit containing its scenario and private answer sheet, while external custom evaluations supply one skill-directory root plus scenario and answer-sheet files through the same runner. The runner sends the same substantial user turns regardless of model responses, records observed skill activations against private turn state, compares them with the selected answer sheet, and emits structured data that can be plotted directly. A tooling-complete suite is valid whether it records every expected call or no calls; the published output describes observed behavior without pass or fail labels.

## Path from A to B

### 1. Build and govern the evaluation scenarios

- Create three different real-world planning scenarios with exactly 30 fixed user turns each.
- Make many turns substantial enough to require meaningful planning work rather than short acknowledgements or trivial edits.
- Begin each scenario by asking the primary agent to create and maintain a living A-to-B plan without naming or manually invoking a planning skill.
- Send every scripted turn in order after the preceding response completes. Do not change, skip, or add turns in response to the model's output.
- Give each scenario an inspectable conversation document, runner instructions, and answer sheet, then embed the complete scenario and answer sheet together as one governed runtime unit.
- Record the first turn on which each benchmark skill is semantically required before observing campaign results.
- Treat later reloads of a skill as additional observations because a harness may keep an already loaded skill active for the rest of the scenario.

### 2. Define and validate the generated evaluation-install contract

- Consume the accepted canonical primary and supporting skill sources.
- Define the exact neutral signal instruction the CLI injects at the beginning of the executable workflow and require it to record the invocation before continuing the unchanged canonical body.
- Define how the generated installation receives the run, turn, harness, model, and evidence context needed to attribute each marker event.
- Preserve the canonical frontmatter, description, selection conditions, and substantive body while making generated evaluation copies clearly CLI-owned and disposable.
- Verify that generating an evaluation install from an updated skill includes the current canonical body, records one event when exercised, and rematerializes current canonical copies for ordinary Skill Issue paths that existed before the run.
- Verify the transformation and marker mechanism before a model replay is treated as evaluation data.

### 3. Build the governed scenario procedure

- Use the CLI's evaluation-install mode to prepare the target harness's native project skill root with generated copies of the supplied Skill Issue and supporting skills plus the shared signal command.
- Select a governed built-in by one identifier that loads its embedded scenario and answer sheet together, or supply a custom skill-directory root, scenario JSON, and answer-sheet JSON file together. Both modes continue through the same runner and comparison path; custom answer sheets are structurally checked against the supplied names without judging semantic expected-call correctness.
- Confirm the harness, model, medium reasoning setting or equivalent, harness version, operating system, and evaluator's default-environment attestation before replay begins.
- Start one clean primary-agent session per scenario through the harness's programmatic interface.
- Prepare harness-specific invocation isolation without changing the user's persistent configuration. The implemented Codex route retains ordinary authentication, ignores user configuration and rules, disables discovered global skills, apps, plugins, and instruction files through explicit invocation configuration, fixes medium reasoning and `workspace-write`, and routes boundary requests through automatic guardian review. Codex skill attribution comes from its structured opaque-signal command event rather than requiring the command to write outside the workspace. The adapter retains the replay conversation in normal Codex history and treats managed-policy conflicts or unrelated reviewer failures as tooling failures.
- Set private active-turn state, send each scripted prompt verbatim, wait for terminal completion, and close the turn without acknowledging, interpreting, or adapting to the agent's responses.
- Keep built-in and custom answer sheets, the token map, and turn state outside the evaluated workspace and hidden from the evaluated primary agent.
- Require the evaluator to choose an output root outside the evaluated workspace. Retain the detailed result and compact website artifact under a unique harness-and-time run directory. Keep private operational state under its CLI-owned internal directory only while the run or interrupted-run cleanup needs it, without using platform application storage for evaluation runs.
- Complete every turn even when expected skill events are absent so missed calls remain observable data.
- Treat runner, instrumentation, permissions, or harness-control failures as tooling faults to fix and rerun rather than as model results.

### 4. Produce graph-ready structured results

- Retain the portable detailed JSON result containing the run metadata, scenario, expected first-activation turns, observed skill events, missing expected events, and any additional or unattributed calls.
- Also emit one compact website JSON artifact with `schema_version`, `run_id`, `scenario_id`, `harness`, `model`, `total_turns`, and `points`. Each point represents a turn containing at least one expected call and contains the one-based numeric `turn`, its source `turn_id`, the number of expected skills `called`, and the number `missed`.
- Make `turn` the website's numeric horizontal axis and use `total_turns` as its full domain. Keep skill-level evidence in the detailed result rather than duplicating it in the website artifact. Derive sample size from the sum of `called` and `missed` instead of storing another count.
- Derive every result from native event evidence while the runner is active. Persist raw events only when the evaluator explicitly enables `--events`.
- Make the website artifact directly consumable as Recharts data without manual reinterpretation of the conversation transcript. The later website work plots `called` and `missed` against the numeric `turn` and calculates its axis and sample display from this artifact.
- Always retain the derived detailed result and compact website artifact. Make the complete transcript and raw event output independent, default-off diagnostic artifacts enabled through `--transcript` and `--events`.
- Warn custom-evaluation users that supplied JSON must contain no personal, confidential, or sensitive information. Warn transcript users that the complete conversation is written unchanged. Do not claim to scan or redact arbitrary conversation content.

### 5. Define the minimum viable evaluator protocol

- Give internal and external evaluators concise instructions for establishing the requested default harness setup, installing the supplied evaluation payload, selecting the assigned model, running the replay, and returning the evidence package.
- Accept evaluator attestation for the default environment during the minimum viable campaign rather than blocking release on automated isolation or configuration inspection.
- Trust a replay that reaches the end with functioning instrumentation and a complete evidence package; missing skill events affect the displayed data rather than the validity of the run.
- Require one tooling-complete run of all three scenarios for each of the thirteen minimum harness-and-model cells, without imposing repeated-trial or pass-threshold requirements.
- Exclude unsupported combinations before campaign execution and omit unrun combinations from the published dataset instead of presenting them as evaluated outcomes.

### 6. Publish descriptive and transparent evaluation data

- Plot the expected and observed skill-call data without assigning pass or fail labels to harnesses or models.
- Present strong and weak observed behavior directly through the graphs rather than translating it into a project-authored verdict.
- Publish the scenario designs, expected call maps, instrumentation approach, environment instructions, structured results, graph derivation, and known limitations so readers can inspect the full process.
- Explain that the campaign is a transparent minimum viable evaluation, that reasonable implementation choices were made in good faith, and that the method can be expanded and improved over time.

## C: Observable Completion Criteria

### Reusable system — completed in Work Blocks 1 and 2

- Three governed 30-turn scenario sets are stored as inspectable assets, and each finished scenario is embedded as one CLI unit containing the fixed conversation and private expected-call answer sheet.
- Evaluation installation is project-only for every harness. Built-in selection accepts no external input files, while custom skill-directory, scenario, and answer-sheet inputs use the same runner and comparison behavior.
- Every evaluation requires an explicit output root outside the evaluated workspace; completed evidence and CLI-owned private run state are both written beneath that root.
- Codex evaluations use the user's normal Codex home for authentication and retained session history while every evaluation behavior setting is supplied as a process-local override. The CLI does not edit persistent Codex configuration. Equivalent native runtime preparation remains a qualification requirement for each additional harness.
- The CLI generates evaluator copies from the current canonical skills at install time, injects only the governed marker instruction, and creates no independently maintained instrumented sources.
- Every completed scenario suite emits the portable detailed JSON result and compact turn-based website artifact, with every website point derivable from the detailed result. Raw events and the complete transcript remain optional diagnostics.

### Qualification and campaign — owned by Work Block 3

- A zero-call suite remains a valid result when all three scenarios and instrumentation complete successfully; tooling failures are fixed and rerun instead of being graphed as model behavior.
- Each of the thirteen minimum matrix cells has one tooling-complete three-scenario suite and evidence package; repeated trials remain a later expansion.
- The minimum evaluator instructions establish the default-environment expectation and collect evaluator attestation without requiring automated isolation for the initial release.

### Publication — owned by Work Block 4

- The website can consume the compact `points` array directly, use a numeric turn domain, derive sample size, and display descriptive called-versus-missed comparisons without pass or fail labels.
- The published method and raw evidence allow readers to understand what was tested, reproduce the process, inspect its limitations, and form their own view of the results.

## Dependency Handoffs

### Upstream inputs

- The completed product, support, and evidence contract is authoritative for the five-harness minimum qualification tier, thirteen medium-setting matrix cells, one-suite MVP threshold, default-environment expectation, valid-run meaning, and public-claim boundary.
- The completed direct-install architecture supplies the authoritative local routes, direct disposable materialization boundary, later discovery-and-activation qualification boundary, unsupported surfaces, and fail-closed behavior used to place the evaluation payload.
- The supporting-skills bundle plan supplies the final discipline-skill set; this plan defines the injected instruction and marker contract; the portable-skill step supplies delegated behavior; Work Block 2 supplies CLI evaluation-install and removal behavior.

### Authoritative outputs

- This section owns the opaque injected instruction, governed scenarios, expected call maps, private run-state and event contract, direct primary-agent runner procedure, JSON evidence package, evaluator protocol, and descriptive graph semantics.
- Downstream tasks must consume those artifacts without changing the expected call meaning, converting observations into pass or fail labels, or weakening the evidence boundary.

### Required downstream consumers

- Work Block 1 supplies the completed portable skill and evaluation assets defined by this contract.
- Work Block 2 has implemented the generated installation and runner requirements; final installer qualification remains downstream.
- Work Block 3 consumes the evaluator procedure, replay, validity, evidence-return, and graph rules during qualification and campaign execution.
- Work Block 4 consumes the inspectable result package and descriptive semantics for public releases and website graphs.
- Work Block 5 consumes the transparent method and results when producing the video and competition submission.

## Unresolved Matters

- All three governed scenarios and the common Go runner are complete. Live cross-harness proof and campaign aggregation remain open.
- Harness-specific limits may require equivalent event-capture mechanisms while preserving one common evaluation meaning and result schema.

## Completed Instrumentation Contract

`evaluations/skill-calling/instrumentation-contract.md` and `evaluations/skill-calling/event.schema.json` now own the generated-copy transformation, opaque neutral `signal` command, project-only placement, private state, turn attribution, marker behavior, cleanup, and portable invocation-event shape.

## Completed Scenarios

The gardening web-application, community archive desktop-application, and neighborhood emergency-preparedness scenarios each contain 30 fixed substantial turns and four required first activations inside their embedded governed units. Each scenario retains Markdown conversation, instructions, and expected-call governance views. Together they exercise software planning across web and desktop delivery plus a non-software community program while measuring only canonical Skill Issue skills.

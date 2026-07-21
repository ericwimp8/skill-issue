# Skill Issue

> It’s not always a skill problem, but it’s always a Skill Issue.

Skill Issue helps people create, evaluate, and improve agent skills. It separates
problems in a skill from problems in the model and agent harness running it, so
users can fix the right thing.

In this project, a skill is a reusable set of instructions and resources for a
coding agent.

## Why Skill Issue exists

When a skill fails, the cause is often unclear:

- The description may not tell the model when to use the skill.
- The instructions may be ineffective after the skill loads.
- The model may select skills inconsistently.
- The agent harness may not discover or expose skills reliably.

These failures can look identical to the user. Rewriting a good skill cannot
fix an environment that never loads it, and changing environments cannot fix
unclear instructions. Skill Issue makes these failure modes observable and
actionable.

## What Skill Issue does

Skill Issue has two connected workflows:

1. **Evaluate the environment.** Measure whether a model-and-harness
   configuration discovers and invokes the expected skills.
2. **Build better skills.** Create, evaluate, diagnose, and refine skills until
   they behave as intended.

The environment is evaluated first. Reliable skill discovery provides a sound
foundation for skill refinement; unreliable discovery is reported as an
environment limitation instead of being misdiagnosed as a writing problem.

## Environment evaluation

The built-in evaluation replays three fixed 30-turn scenarios. Every tested
configuration receives the same prompts in the same order, and later prompts
never adapt to the agent’s responses. Each scenario includes expected skill
calls distributed across the conversation and three short turns where no skill
should be invoked.

Results are measured by turn rather than by token count. Tokenization, context
limits, compaction, and telemetry differ across models and harnesses; turns
provide a consistent comparison point across identical conversations.

### Comparison design

The evaluation supports six harnesses:

- OpenAI Codex
- Claude Code
- Cursor
- Pi
- OpenCode
- Kilo Code

The comparison campaign tests models across harnesses and multiple models
within the same harness where those combinations are available. Each
configuration uses the same scenarios, prompts, skill payload, scoring rules,
and reasoning target or closest supported equivalent. Results record the exact
model, harness version, operating system, and effective reasoning setting.

Cross-harness comparisons can expose harness-level patterns. Within-harness
comparisons can expose model-level patterns. Unsupported or unrun combinations
are omitted.

### Isolation and evidence

The runner evaluates each harness’s primary agent while reducing interference
from ambient skills, plugins, instructions, configuration, and unrelated
project customization. It uses:

- fixed, verbatim prompts;
- private answer sheets and turn state;
- temporary instrumented copies of the canonical skills;
- cryptographically random marker tokens; and
- run-owned external event storage.

Temporary instrumentation is removed during cleanup. When complete isolation
is unavailable, the result records the prepared environment and the remaining
uncertainty.

Each completed run writes:

- `result.json`: the detailed evaluation result;
- `website.json`: a compact, chart-ready result; and
- optional `events.jsonl` and `transcript.json` diagnostics.

Events and transcripts are disabled by default. A transcript may contain the
complete evaluation conversation, including confidential information supplied
during the run.

### Interpreting results

A correctly instrumented run remains valid when expected skills are missed;
those misses are evaluation results. Tooling, permission, session, visibility,
or instrumentation failures are operational failures and must be resolved
before the run is interpreted as model behavior.

One scenario suite cannot establish universal reliability. Results can still be
affected by residual configuration, organization policy, provider aliases,
model or harness updates, operating-system differences, and incomplete
isolation. Published results therefore report observed calls and misses without
claiming a universal winner or guarantee.

## Skill creation

Users describe the outcome they want in ordinary language. For example:

> Create a skill that runs linting at the end of each task, applies automatic
> fixes, then finds and resolves any remaining compile-time errors.

Skill Issue then:

1. Inspects the project and resolves important scope questions.
2. Generates an idiomatic skill for the project and target environment.
3. Creates evaluations for skill selection and post-invocation behavior.
4. Runs the evaluations and diagnoses failures.
5. Refines the description, instructions, or supporting resources as needed.
6. Delivers the skill with a clear account of what was validated.

The user defines the outcome; Skill Issue handles the skill-engineering
workflow.

## Command-line interface

The project includes a self-contained Go CLI for local use. It installs the
eleven canonical Skill Issue skills, checks supported harness environments,
runs blind turn-attributed evaluations, and produces local evidence files.

The CLI owns deterministic installation, removal, evaluation execution,
reporting, cleanup, and recovery. Installed skills own the generation,
diagnosis, and refinement workflows. Reinstalling replaces only the known
Skill Issue payload directories.

See [`cli/README.md`](cli/README.md) for commands, harness configuration,
installation behavior, custom evaluations, output formats, recovery, and
qualification details.

### Develop the CLI

```sh
go vet ./cli/...
go run ./cli/cmd/skill-issue help
./cli/scripts/build-cross-platform.sh
```

The build script produces standalone Darwin, Linux, and Windows binaries for
`amd64` and `arm64`. Cross-compilation confirms that the binaries build; native
runtime qualification is still required on each released platform.

## Website

The repository also contains a static React and TypeScript website for
exploring the project, methodology, generated skills, and evaluation results.
It builds for GitHub Pages and does not require a database or runtime API.

### Run locally

```sh
npm install
npm run dev -- --host 127.0.0.1
```

Open `http://127.0.0.1:5173/`.

### Validate and preview the production build

```sh
npm run validate
npm run preview
```

The production preview is available at
`http://127.0.0.1:4173/skill-issue/`. Set `VITE_BASE_PATH` if the GitHub Pages
repository path differs from `/skill-issue/`.

### Update website content

- Edit curated copy, release metadata, summary metrics, and methodology text in
  `src/data/siteData.ts`.
- Add complete generated-skill examples under
  `showcase-skills/*/skill/*/SKILL.md`; the website discovers them at build
  time.
- Keep evaluation types, labels, artifact adaptation, and illustrative results
  in `src/data/evaluationData.ts`.
- Keep selection and filtering in `src/components/EvaluationExplorer.tsx` and
  chart presentation under `src/components/charts/`.

Published charts must use accepted evaluation artifacts rather than
hand-authored values. Load selected compact artifacts with:

```sh
npm run results:update -- output/run-one/website.json output/run-two/website.json
```

This command validates the artifacts and writes the collection to
`src/data/publishedWebsiteArtifacts.json`. Keep downloadable binaries in GitHub
Releases rather than the Pages build.

## Project status

Skill Issue is in active development. The CLI, embedded skill payload, blind
evaluation runner, result artifacts, website, harness adapters, environment
checks, and six-target cross-build are implemented. Evaluation qualification,
campaign evidence, release preparation, and public distribution continue.

The current skill-calling campaign is tracked in
[`plans/skill-calling-evaluation-campaign/evaluation-progress.md`](plans/skill-calling-evaluation-campaign/evaluation-progress.md).
Its execution rules, scheduling, failure handling, and progress-update process
are defined in
[`campaign-orchestration-prompt.md`](plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md).

## Longer-term direction

The local CLI answers the immediate question: **does my current agent setup use
skills reliably enough for me to depend on them?**

A future hosted service could run standardized evaluations across more models
and APIs, publish broader comparisons, and aggregate compatible local results.

## License

Skill Issue is available under the [MIT License](LICENSE).

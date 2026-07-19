# Skill Issue

Skill Issue is a project for making agent skills easier to create, evaluate, and trust.

The name is a play on the familiar claim that a problem is a “skill issue.” In this case, the skills are the reusable instructions and workflows used by coding agents.

## The problem

Building an agent skill currently involves a great deal of uncertainty and manual refinement. When a skill fails, it can be difficult to tell what actually went wrong:

- Is the description too weak for the model to recognize when the skill applies?
- Is the skill body unclear or ineffective after the skill is invoked?
- Is the model inconsistent at selecting skills?
- Is the surrounding agent harness doing too little to support reliable skill use?

These failures often look the same from the user’s perspective. Someone can spend hours rewriting a good skill description when the model-and-harness environment is the real limitation. Others encounter inconsistent results, conclude that skills do not work, and abandon them entirely.

Skill Issue aims to make those failures observable, attributable, and actionable.

## How it works

Skill Issue has two complementary parts:

1. **Evaluate the environment.** Determine whether a particular model and agent harness can discover and invoke skills reliably.
2. **Build better skills.** Help users create, evaluate, and refine skills that behave as intended.

The model and harness are initially treated as a single practical environment. The important question for a user is whether the setup they work in—such as Codex, Claude Code, or another coding-agent environment—can use skills consistently enough to support their workflow.

Users evaluate their environment first. If it performs reliably, they can build and refine skills with confidence. If it does not, they can reconsider the environment instead of endlessly rewriting a skill that was never the underlying problem.

## Evaluation measurement

The initial skill-calling evaluation uses three scripted scenarios with the same number of fixed user turns. Every harness-and-model configuration receives the same turns in the same order. Agent responses are retained as evidence but do not change the next scripted turn.

Results are measured by turn rather than by context-token count. Tokenization, context limits, compaction, and exposed telemetry differ across models and harnesses, while a turn represents the same point in the user's workflow for every configuration. This makes observations such as a skill first loading on turn 10 or failing to load by turn 20 directly comparable across identical conversations.

The MVP scenarios contain 30 turns each and deliberately include substantial real-world planning work. Context-based analysis may be added later when comparable telemetry is available across the supported environments.

## Skill creation

The goal is to let users describe the outcome they want in ordinary language. For example:

> Create a skill that runs linting at the end of each task, applies automatic fixes, then finds and resolves any remaining compile-time errors.

Skill Issue should inspect the current project, understand the relevant languages and tooling, and identify important ambiguities before generating anything. If the repository contains both a TypeScript backend and a Rust project, for example, it should ask whether the skill applies to one or both rather than silently choosing the wrong scope.

Once the request is clear, the system should:

1. Generate an idiomatic skill for the user’s project and environment.
2. Create evaluations for both skill invocation and skill behavior.
3. Run those evaluations and diagnose failures.
4. Refine the description or body as appropriate.
5. Repeat until the skill meets the expected standard.
6. Deliver a ready-to-use skill with a clear account of what was validated.

The user describes the outcome. Skill Issue handles the skill-engineering work.

## Minimum viable product

The initial product will focus on local execution. Rather than operating a hosted service across every model API, Skill Issue will provide something users can install and run inside their own agent setup.

The MVP uses a self-contained Go CLI distributed as prebuilt macOS, Windows, and Linux executables. It:

- installs the required canonical skills;
- runs a repeatable evaluation inside the user’s configured environment;
- measures skill discovery and invocation behavior; and
- produces a useful local report explaining the results.

The CLI owns deterministic installation, verification, repair, update, removal, evaluation execution, and reporting boundaries. The installed agent skills own generation, diagnosis, and refinement behavior.

The current implementation embeds the canonical primary and supporting skills, installs them into researched native project or user roots, records external ownership receipts, and provides blind primary-agent replay with private turn attribution, graph-ready evidence, restoration, and cleanup. Harness adapters still require real-environment qualification before release claims are made.

### Evaluation mitigation

The runner evaluates each harness's primary agent directly. It uses fixed verbatim prompts, private answer sheets and turn state, cryptographically random opaque marker tokens, temporary generated skill copies, and external event storage. The runner never adapts later prompts from model responses and removes temporary instrumentation after the run while retaining evidence.

These measures minimize evaluation clues. They cannot guarantee that an agent inspecting its environment will never infer that instrumentation exists.

### CLI development

```sh
go vet ./cli/...
go run ./cli/cmd/skill-issue diagnose
./cli/scripts/build-cross-platform.sh
```

The automated CLI test suite is intentionally deferred while the implementation is changing rapidly. Final regression, adapter, lifecycle, and platform tests will be written against the stabilized product during qualification rather than maintained against transitional interfaces.

The build script produces standalone Darwin, Windows, and Linux binaries for `amd64` and `arm64`. Cross-compilation does not replace native runtime testing on each released platform.

## Longer-term vision

A future hosted service could provide standardized evaluations across multiple models and APIs, publish comparisons, and present results through clear graphs and reports. Local evaluation results could also be exported or aggregated for broader comparisons.

The MVP deliberately starts with the part users can run in the environments they already use. This removes the cost and complexity of operating every provider’s API while still answering the central question: **does my current setup use skills well enough for me to depend on them?**

## Status

Skill Issue is in active development. The standalone CLI, canonical embedded payload, direct installation lifecycle, blind evaluation runner, and website mock-up are implemented. Real-harness and native-platform qualification, campaign evidence, and public release remain in progress.

The current six-block completion sequence is maintained in `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md`. Its completed-work inventory, research source map, CLI mismatch record, and old-task reconciliation live in `plans/skill-issue-project-completion/reorganization-dependency-audit.md`. Document authority, supersession, consumer links, and required update routing live in `plans/skill-issue-project-completion/document-authority-and-update-map.md`.

## Website mock-up

The repository includes a one-page React and TypeScript website that builds to static files for GitHub Pages. Benchmark content is local and requires no database or runtime API.

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

The production preview is available at `http://127.0.0.1:4173/skill-issue/`. Set `VITE_BASE_PATH` when the eventual GitHub Pages repository path differs from `/skill-issue/`.

### Update website content

All editable website copy, release metadata, summary metrics, methodology text, and evaluation graph values live in `src/data/siteData.ts`.

To add another graph, append one evaluation definition with a unique `id`, harness and model labels, description, sample size, and context checkpoint values. The results grid renders new definitions automatically through the shared `EvaluationChart` component.

Replace the release URL in the same data file when the first CLI artifact is published through GitHub Releases. Keep binaries in Releases rather than the Pages build.

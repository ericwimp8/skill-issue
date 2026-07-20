# Repository Onboarding Guide A-to-B Plan

## A — Current Position

- The requested reusable skill is named `repository-onboarding-guide` and belongs under `showcase-skills/repository-onboarding-guide/skill/repository-onboarding-guide/`.
- Its target is an unfamiliar repository whose production source, instructions, configuration, documentation, tests, scripts, and local state may disagree or be incomplete.
- Production source is the behavioral source of truth; tests and documentation may provide leads and validation evidence but cannot establish current behavior on their own.
- Durable artifacts must use repository-relative paths and exclude secrets, personal or business identities, usernames, home-directory names, and machine-specific checkout paths.
- Production workflow skills, supporting skills, CLI code, website code, and other showcase work are outside the edit boundary.

## B — Desired Position

A ready-to-use skill produces a concise onboarding guide that helps a new contributor navigate the repository from authoritative instructions to concrete production behavior, setup and validation, ownership boundaries, common workflows, local state, and honestly recorded unknowns.

## Path from A to B

1. Define an evidence-first repository discovery sequence that resolves scoped instructions and project-native entry points before drafting conclusions.
2. Define source-tracing rules for following at least one representative behavior from an external entry through concrete production effects and ownership boundaries.
3. Define how documentation, tests, scripts, configuration, history, and generated state contribute evidence without displacing production source truth.
4. Define a concise guide structure that distinguishes confirmed facts, qualified inferences, conflicts, and material unknowns.
5. Package the portable skill and minimal Codex interface metadata without unnecessary scripts, references, or assets.
6. Structurally validate the bundle and evaluate description selection plus representative body behavior in isolated unfamiliar-repository fixtures.
7. Refine only at the semantic owner of retained failures, then audit hashes, formatting, diffs, and privacy.

## C — Completion Criteria

- The skill discovers applicable instruction files and states their scope and precedence.
- The guide identifies external entry points, major components, dependencies, and ownership boundaries from current source and configuration.
- At least one representative behavior is traced end to end through production source to its concrete effect, including important state or error paths.
- Setup, run, build, test, lint, formatting, generation, and other validation commands are reported only when supported by current project-owned configuration or scripts.
- Generated, cached, ignored, secret-bearing, or machine-local state is distinguished from tracked source and durable configuration.
- Common contribution workflows are grounded in authoritative instructions, scripts, configuration, or repeated source/history evidence rather than invented conventions.
- Tests and documentation are treated as leads or validation evidence and conflicts with production source are reported.
- The guide remains concise, uses repository-relative paths, labels material unknowns, and avoids unsupported architectural or workflow claims.
- The generated skill passes structural validation and the governed description and body evaluations, or records the exact stopping gate and retained evidence.
- Every retained artifact remains under `showcase-skills/repository-onboarding-guide/` and passes formatting, diff, hash, and privacy checks.

## Generation Contract

- **Destination:** `showcase-skills/repository-onboarding-guide/skill/repository-onboarding-guide/`
- **Supported surface:** portable Agent Skills content with OpenAI Codex interface metadata.
- **Codex metadata:** required for readable discovery; implicit invocation remains enabled.
- **Generation viability:** autonomous.
- **Execution preference:** autonomous continuation through generation and evaluation.
- **Authority boundary:** create and refine only artifacts inside `showcase-skills/repository-onboarding-guide/`; use synthetic or repository-owned fixtures; run repository-local validation without changing production owners.
- **Required user stops:** none currently active.
- **External dependencies:** fresh GPT-5.6 Sol medium agents and the already-qualified local Codex evaluation surface.
- **Unresolved implementation matters:** exact body-case fixture shapes and guide formatting may be chosen during generation and evaluation without changing intent.
- **Evaluation handoff:** continue directly into `skill-evaluation-and-refinement` with automatic refinement authorized inside the destination workspace.


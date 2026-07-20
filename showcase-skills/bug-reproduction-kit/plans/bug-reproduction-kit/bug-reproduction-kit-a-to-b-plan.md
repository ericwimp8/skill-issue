# Bug Reproduction Kit A-to-B Plan

## A — Current Position

- The user has supplied the skill name, `bug-reproduction-kit`, and a writable showcase workspace at `showcase-skills/bug-reproduction-kit/`.
- The requested skill must work from incomplete or vague bug information without presenting assumptions as facts.
- The available inputs may include prose reports, repository state, commands, environment details, screenshots, logs, stack traces, recordings, or existing issue material; any subset may be absent.
- The repository provides the Skill Intake, Skill Generation, Skill Evaluation and Refinement, skill-authoring, document-update, prompt-writing, and system-ownership disciplines.
- The canonical payload will target portable Agent Skills behavior and OpenAI Codex discovery without requiring external services, secrets, or host configuration.
- Process artifacts must remain distinct from the installable skill payload while being retained in the showcase workspace.

## B — Desired Position

A ready-to-use `bug-reproduction-kit` skill guides an agent from uncertain bug material to a compact, evidence-backed reproduction package that another person can execute, inspect, and file as an issue without hidden context.

## Path from A to B

1. Establish the skill's evidence boundary and the minimum facts needed to attempt reproduction safely.
2. Define a workflow that inspects available sources before asking for missing information and records material unknowns explicitly.
3. Define how to establish relevant environment facts, isolate the smallest credible trigger, and distinguish observed behavior from the expected contract.
4. Define how to collect, reference, or request logs and artifacts while preserving provenance and avoiding invented evidence.
5. Define a concise reproduction-package output whose claims distinguish confirmed facts, observations, inferences, and unresolved gaps.
6. Package the canonical skill with only behavior-essential instructions and host metadata.
7. Validate structure and evaluate representative incomplete-report cases in isolated trials, refining the semantic owner of any material failure.

## C — Completion Criteria

- The skill investigates available project and runtime evidence before asking focused questions.
- The skill records relevant environment facts with their source or marks them unknown.
- The skill produces minimal, ordered reproduction steps with explicit prerequisites and a stated reproducibility result.
- The skill separates expected behavior, actual observed behavior, and evidence-supported inference.
- The skill collects, points to, or requests useful logs and artifacts without fabricating their contents or existence.
- The skill surfaces only material missing information and explains how each gap blocks or limits reproduction.
- The skill produces a ready-to-file issue or equivalent package containing a concise summary, environment, prerequisites, reproduction steps, expected and actual behavior, reproducibility, evidence, and open gaps.
- The package remains useful when reproduction is unsuccessful by recording attempted paths, negative results, limitations, and the next evidence needed.
- The canonical skill passes the available structural validator, and runtime evaluation either supplies retained behavior evidence for every criterion or records the exact governing stop.

## Generation Contract

- **Destination:** canonical skill payload at `showcase-skills/bug-reproduction-kit/skill/bug-reproduction-kit/`; retained planning and campaign evidence remain elsewhere under `showcase-skills/bug-reproduction-kit/`.
- **Supported harness surfaces:** portable Agent Skills content and OpenAI Codex project or user skill delivery.
- **Generation viability:** autonomous; the requested outcome, destination, evidence boundary, and available local workflows are sufficient.
- **Selected execution preference:** autonomous continuation through generation and evaluation.
- **Authority to act:** create and refine artifacts owned by this showcase workspace; apply automatic semantic refinements supported by retained evaluation evidence.
- **Required user stops:** a new decision that would change the intended skill, permission to access unavailable private evidence, or an evaluation capability boundary that the governing workflow requires the user to resolve.
- **External dependencies or unavailable inputs:** none for generation; individual future bug investigations may require unavailable runtime access or reporter-supplied evidence, which the skill must identify rather than invent.
- **Unresolved implementation choices:** exact headings and concise wording of the reproduction package, fixture domains used for evaluation, and whether the canonical payload needs any bundled reference.
- **Expected evaluation handoff:** continue into `skill-evaluation-and-refinement` using the repository's qualified Codex Desktop surface, automatic body-refinement mode, isolated representative fixtures, and retained native evidence.

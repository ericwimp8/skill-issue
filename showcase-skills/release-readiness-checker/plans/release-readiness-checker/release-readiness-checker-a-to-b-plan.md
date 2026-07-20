# Release Readiness Checker A-to-B Plan

## A — Current Position

- The user supplied the stable name `release-readiness-checker` and the writable workspace `showcase-skills/release-readiness-checker/`.
- The requested skill evaluates a repository or supplied release candidate without publishing, deploying, tagging, or otherwise changing external release state.
- Candidate evidence can include source, release instructions, configuration, migrations, versions, documentation, security or privacy rules, rollback material, commands, logs, and retained test or build results; any subset may be absent or stale.
- Readiness gates must come from authoritative project-owned sources and current candidate state rather than a generic checklist.
- The repository provides governed intake, generation, evaluation, skill-authoring, document-update, prompt-writing, and system-ownership instructions.
- The canonical payload targets portable Agent Skills behavior and OpenAI Codex discovery without external services, secrets, or bundled project-specific release rules.

## B — Desired Position

A ready-to-use `release-readiness-checker` skill produces an evidence-bounded, prioritized release decision whose gates trace to the candidate's authoritative sources and whose exact next actions can be executed without mistaking stale documentation or historical test results for current proof.

## Path from A to B

1. Define the evidence boundary, candidate identity, authority hierarchy, and prohibition on release side effects.
2. Define how to derive applicable gates from project-owned release sources and map each gate to current evidence.
3. Define how to inspect implementation, configuration, migrations, versioning, documentation, verification evidence, security or privacy requirements, rollback preparation, and unresolved risk at their semantic owners.
4. Define precise `passed`, `failed`, `blocked`, `not-run`, and `not-applicable` classifications with provenance and limitations.
5. Define a prioritized decision and exact next-action format that prevents readiness claims unsupported by current execution evidence.
6. Package the concise canonical skill and required Codex discovery metadata without project-specific references, scripts, or assets.
7. Validate structure and evaluate representative ready, failing, blocked, and incomplete candidates in isolated fresh-agent trials, refining only retained material failures.

## C — Completion Criteria

- The skill identifies the exact candidate, scope, revision or artifact identity, release target, and evaluation limitations.
- It derives release gates from authoritative source, configuration, and release instructions, recording provenance and resolving conflicts by authority rather than convenience.
- It traces relevant implementation and operational paths to concrete owners before deciding a gate.
- It inspects applicable code, configuration, migrations, versioning, documentation, current test/build evidence, security or privacy requirements, rollback preparation, and unresolved risks without treating the topic list as a universal checklist.
- Every gate is classified as `passed`, `failed`, `blocked`, `not-run`, or `not-applicable` with evidence, limitations, and an owner or resolution path when action remains.
- A passing claim requires current candidate evidence; historical results, documentation claims, configured scripts, or source inspection alone are not reported as executed proof.
- The final decision is `ready`, `not ready`, or `undetermined`, follows explicit project policy where available, and otherwise applies a conservative evidence rule without inventing policy.
- The output prioritizes release blockers and provides exact next actions that name the action, owner when known, evidence needed, and gate it resolves.
- The workflow never publishes, deploys, tags, uploads, migrates shared state, or performs another release side effect.
- Structural validation passes, and runtime evaluation either retains representative evidence for every criterion or records the exact governing stop.

## Generation Contract

- **Destination:** canonical payload at `showcase-skills/release-readiness-checker/skill/release-readiness-checker/`; all workflow evidence remains elsewhere under `showcase-skills/release-readiness-checker/`.
- **Supported harness surfaces:** portable Agent Skills content and OpenAI Codex project or user skill delivery.
- **Generation viability:** autonomous; the intended outcome, authority boundary, destination, and local workflow are sufficient.
- **Selected execution preference:** autonomous continuation through generation and evaluation.
- **Authority to act:** create and automatically refine artifacts owned by this showcase workspace; run local read-only or validation commands that cannot publish or deploy.
- **Required user stops:** a decision changing the intended skill, unavailable private evidence requiring user access, an unsafe or external release action, or an evaluation gate the governing workflow cannot satisfy.
- **External dependencies or unavailable inputs:** none for generation; future candidate checks may be blocked by unavailable environments, credentials, platforms, or artifacts and must report that state.
- **Unresolved implementation choices:** concise output headings, representative fixture domains, and whether any deterministic reusable resource is truly required.
- **Expected evaluation handoff:** continue into `skill-evaluation-and-refinement` with automatic body refinement, fresh independent GPT-5.6 Sol medium-reasoning agents, isolated synthetic fixtures, and retained native evidence.

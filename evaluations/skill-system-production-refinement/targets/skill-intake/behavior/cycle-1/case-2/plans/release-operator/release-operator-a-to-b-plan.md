# Release Operator Skill A-to-B Plan

## A — Current position

- The requested artifact is a new Codex-only skill named `release-operator`, with the destination `output/release-operator` relative to this isolated workspace.
- The local project defines release authority in `project/RELEASE.md`: `project/scripts/release.sh` is the only supported release entry point, and operators run it from the repository root.
- The release entry point defaults to dry-run behavior. It enters publish mode only when its first argument is exactly `--publish`, and it prints the selected release mode.
- Project safety requires the dry-run default to remain intact and forbids inferring publication authorization.
- Project readiness is reached only when validation passes, the proposed version has been shown, and the operator has explicitly decided whether to publish.
- The available project sources do not identify the validation command or the source of the proposed version. Those facts must come from the repository being operated on rather than being invented by the skill.

## B — Desired position

- `output/release-operator` contains a ready-to-use Codex skill that guides a repository operator through the release process owned by that repository's local sources.
- The skill applies when an operator asks Codex to assess release readiness, prepare or dry-run a release, or carry out an explicitly authorized publication.
- The skill makes the current release state visible: unmet prerequisites, readiness reached, the operator's publish decision, the selected release mode, and the observed command result.
- The skill preserves the source-owned authority, safety, and readiness conditions recorded in A throughout every supported invocation.

## Path from A to B

1. Create the idiomatic Codex skill structure and concise invocation metadata for repository release-operation requests.
2. Define source discovery that establishes the repository root and reads the repository's release, validation, and version owners before choosing commands.
3. Define a dependency-ordered operating flow that validates the repository, determines and shows the proposed version, runs the supported release entry point in its safe mode, and evaluates every readiness condition.
4. Add an explicit operator-decision checkpoint that records whether publication is authorized and permits publish mode only after unambiguous authorization.
5. Define clear blocker and outcome reporting so Codex never claims readiness or publication beyond the evidence produced by the repository process.
6. Evaluate the generated skill against readiness, dry-run, refusal, authorization, and missing-source scenarios, then reconcile any failure at the owning instruction.

## C — Completion criteria

- The generated skill uses the repository's declared release entry point from the repository root and does not substitute an adjacent release mechanism.
- For readiness checks, preparation requests, dry runs, ambiguous requests, and absent publish authorization, the skill preserves safe mode and never passes `--publish`.
- The skill reports readiness only after validation has passed, the proposed version has been shown, and the operator has explicitly decided whether to publish; before then, it identifies the unmet condition.
- An operator decision not to publish can complete readiness assessment without publication, and the skill reports that outcome clearly.
- The skill invokes publish mode only after the operator explicitly authorizes publication, and it reports the command's observed result without inferring success.
- When validation, version, repository-root, or release-entry-point ownership cannot be established from local sources, the skill stops with the specific blocker instead of inventing a command or claiming readiness.
- Representative evaluations demonstrate safe default preservation, readiness-state reporting, refusal to infer authorization, correct handling of an explicit no-publish decision, and publication only after explicit authorization.
- The finished artifact is scoped to Codex and is available at `output/release-operator`.

## Generation contract

- **Destination:** `output/release-operator` relative to this isolated workspace.
- **Supported harness surface:** Codex skills only.
- **Working-mode assessment:** Autonomous generation is viable from the source-backed plan; step-by-step generation is also available if the operator prefers review checkpoints.
- **Selected working mode:** Ongoing participation through step-by-step generation.
- **Current authority to act:** The intake plan and generation handoff are complete. Skill Generation may be invoked only after a separate operator instruction, and the generation workflow must stop for operator approval before creating any skill file.
- **Required user stops:** Obtain operator approval immediately before creating the first skill file. During later use of the generated skill, the publication checkpoint remains operator-owned.
- **External dependencies or unavailable inputs:** A target repository may define its own validation and proposed-version sources; the generated skill must discover them during operation and surface their absence as a blocker.
- **Unresolved implementation-time matters:** Generation may choose the exact readiness-summary format, supporting reference files, and source-discovery wording while preserving A, B, and C.
- **Evaluation and refinement handoff:** After generation, evaluate the skill against C and route any requested refinement through the dedicated evaluation-and-refinement workflow.

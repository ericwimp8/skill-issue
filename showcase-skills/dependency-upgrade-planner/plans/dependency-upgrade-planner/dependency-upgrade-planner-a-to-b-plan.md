# Dependency Upgrade Planner A-to-B Plan

## A — Current Position

- A reusable skill named `dependency-upgrade-planner` is requested in a new isolated showcase workspace.
- The requested behavior and exclusions are explicit; no user-owned product decision remains unresolved.
- The repository provides project-local intake, generation, evaluation/refinement, authoring, prompt, document, ownership, validation, and privacy contracts.
- The skill must operate across dependency ecosystems from inspected project evidence and authoritative upstream guidance rather than encode one package manager or dependency family.
- The skill is planning-only: it must not edit manifests, regenerate lockfiles, install packages, or claim compatibility without evidence.

## B — Desired Position

A concise, portable Agent Skill produces a source-backed, dependency-ordered migration plan for a requested dependency upgrade. The plan connects the requested version change to the actual dependency graph and concrete usages, separates effect classes, orders prerequisites and migration work, defines validation and rollback gates, and preserves unresolved risks as explicit unknowns.

## Path from A to B

1. Establish semantic ownership for discovery, upstream research, impact analysis, ordering, evidence labeling, validation, rollback, and stopping boundaries.
2. Generate the smallest skill bundle that owns those behaviors without ecosystem-specific references, scripts, or output assets that would duplicate project-specific evidence.
3. Structurally validate naming, frontmatter, metadata, resource necessity, and every completion criterion.
4. Qualify the requested fresh-agent evaluation surface and retain direct pre-output evidence.
5. Evaluate description selection with two representative trials and two different confirmation trials.
6. Evaluate body behavior on varied connected fixtures covering application, build-tool, transitive, runtime, and platform effects.
7. Refine only material failures at their semantic owner, then repeat the governed affected trials from clean fixtures.
8. Run formatting, diff, hash, privacy, and bounded workflow audits and retain their evidence.

## C — Completion Criteria

1. The description selects the skill for naturally phrased dependency-upgrade planning requests without widening into dependency editing or generic project planning.
2. The skill confirms the exact requested package, current and target constraints, ecosystem, workspace scope, and authority before analysis.
3. It inspects production manifests, lockfiles, dependency configuration, toolchain/platform declarations, and concrete production usages before drawing impact conclusions.
4. It distinguishes direct, transitive, build-tool, runtime, and platform effects and explains their evidence.
5. It consults authoritative upstream release notes, migration guides, compatibility matrices, and peer requirements when available, recording source and applicability.
6. It identifies breaking changes, prerequisites, dependency-ordered work, validation gates, rollback points, and unresolved risks.
7. Its output separates observed project facts, authoritative upstream facts, reasoned implications, and unknowns.
8. It does not edit dependencies, run installation or migration commands, or claim compatibility without evidence and authorization.
9. The generated bundle passes the authoritative skill validator, repository formatting checks applicable to its files, content-hash recording, scoped diff review, and repository privacy checks.
10. Fresh independent evaluation evidence supports description selection and all body criteria, or records the exact governed stop.

## Generation Contract

- **Destination:** `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/`
- **Supported harness surfaces:** portable Agent Skills content and OpenAI Codex interface metadata; evaluation evidence is limited to the qualified Codex fresh-agent surface.
- **Generation viability:** autonomous.
- **Execution preference:** autonomous continuation through evaluation and refinement.
- **Authority to act:** create and refine only artifacts within `showcase-skills/dependency-upgrade-planner/`.
- **Required user stops:** none identified.
- **External dependencies:** authoritative public upstream documentation and available repository inspection tools; unavailable evidence must remain an unknown.
- **Unresolved implementation choices:** exact section wording, metadata copy, fixture ecosystems, and case count are generation/evaluation-owned choices.
- **Evaluation handoff:** continue directly into governed description and body evaluation after structural validation and medium-reasoning environment qualification.

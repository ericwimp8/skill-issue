# Release Operator A-to-B Plan

## A — Current position

- The requested artifact is a Codex-only skill named `release-operator`, destined for `output/release-operator` within this case workspace.
- The skill is intended for a repository operator who needs to follow the repository's real, locally defined release process rather than a generic release checklist.
- `project/RELEASE.md` establishes `scripts/release.sh` as the only supported release entry point and requires operators to run it from the repository root.
- The supported release script defaults to dry-run mode. Publishing occurs only when the operator passes `--publish`, and publish authorization must never be inferred.
- The project defines a release as ready when validation has passed, the proposed version has been shown, and the operator has explicitly decided whether to publish.
- The current local release script reports only whether it is in `dry-run` or `publish` mode. The available project sources do not supply validation commands or a mechanism for deriving and displaying the proposed version.
- Local project sources are authoritative for release behavior. Missing release evidence must remain visible rather than being replaced with invented commands or assumptions.

## B — Desired position

A ready-to-use Codex skill guides a repository operator through the release process owned by the active repository's local sources. It discovers and follows the authoritative release entry point, preserves the repository's safe dry-run default, gathers and reports each available readiness signal, clearly identifies whether the documented readiness standard has been satisfied, and pauses for explicit operator authorization before any publishing action.

## Path from A to B

1. Create the Codex skill package at `output/release-operator` with metadata and instructions that make repository release operation and release-readiness checking its user-visible purpose.
2. Define a source-discovery phase that inspects the active repository's release documentation, scripts, configuration, and invoked production paths before proposing or executing release actions.
3. Establish local release sources as the behavioral authority, including the documented entry point, required working directory, validation mechanism, version proposal mechanism, dry-run behavior, and publishing switch.
4. Encode a safe operator flow that uses the supported entry point from the required location and exercises its dry-run behavior before any possible publish action.
5. Track the documented readiness conditions individually: validation result, proposed-version visibility, and the operator's explicit publish-or-do-not-publish decision. Report missing evidence as an unmet readiness condition.
6. Add a hard authorization stop immediately before any publishing action. Require an explicit operator decision for that release and pass a publish flag only after affirmative authorization.
7. Define concise operator-facing status output that distinguishes the discovered release process, dry-run result, satisfied and unsatisfied readiness conditions, authorization state, and final action taken.
8. Validate the skill's Codex packaging and exercise representative safe-default, incomplete-readiness, authorized-publish, and declined-publish scenarios before refinement handoff.

## C — Completion criteria

- Codex can discover and invoke `release-operator` for repository release operation and readiness requests.
- The skill inspects and follows the active repository's local release sources before acting, without substituting a generic release process.
- In this project, the skill identifies `scripts/release.sh` as the sole supported entry point and runs or proposes it from the repository root.
- The skill uses dry-run behavior by default and does not add `--publish` unless the operator explicitly authorizes publishing for the current release.
- The skill reports validation, proposed-version visibility, and the operator's explicit decision as separate readiness evidence and declares the documented release readiness condition satisfied only when all three are present.
- When project sources do not provide validation or version-proposal mechanisms, the skill states that readiness is incomplete and identifies the missing source-defined evidence without inventing commands.
- After the dry run and readiness report, the skill presents the operator with a clear publish-or-do-not-publish decision and safely honors either response.
- Declined, absent, ambiguous, stale, or merely implied authorization leaves the release unpublished.
- Any publish action uses the repository-owned entry point and is followed by an operator-facing report of the concrete action and observed result.

## Generation contract

- **Destination:** `evaluations/skill-system-production-refinement/targets/skill-intake/behavior/cycle-3/case-2/output/release-operator`
- **Supported harness surface:** Codex only.
- **Generation viability:** Autonomous. The known project-source gap can be handled safely as an unmet readiness condition and does not require inventing behavior or obtaining a design decision.
- **Selected execution preference:** Step-by-step execution.
- **Authority to act:** Intake may complete and hand off this plan only. Skill generation is not authorized in this intake, and no skill file may be created without a later explicit operator approval.
- **Required user stops:** After a separate invocation of `skill-generation`, present the first proposed generation step and pause for explicit operator approval before creating any file in the destination. Continue generation through explicit step-by-step approvals.
- **External dependencies or unavailable inputs:** The sample project does not define validation or proposed-version commands. The generated skill must discover such mechanisms in the active repository or report the corresponding readiness evidence missing.
- **Unresolved implementation-time matters:** Generation may choose the exact instruction structure, supporting references, and validation fixtures while preserving this plan's source-authority, readiness, and authorization requirements.
- **Evaluation and refinement handoff:** After generation and structural validation, exercise the completion scenarios above and route any semantic failures into skill evaluation and refinement.

# Release Operator A-to-B Plan

## A — Current position

- The requested artifact is a new Codex-only skill named `release-operator`, intended for `output/release-operator` within this isolated workspace.
- The skill must guide a repository operator through the release process established by the local project sources, preserve the safe default, make release readiness visible, and never publish without explicit operator authorization.
- `project/RELEASE.md` defines `project/scripts/release.sh` as the only supported release entry point and requires it to run from the repository root.
- `project/scripts/release.sh` defaults to `dry-run`, changes to `publish` only when its first argument is `--publish`, and prints the selected release mode.
- The project defines release readiness as validation passing, the proposed version being shown, and the operator explicitly deciding whether to publish.
- The available project sources do not identify a validation command or a source for the proposed version. The finished skill must treat either missing item as unmet readiness evidence rather than inventing a mechanism or declaring readiness.
- Intake is the current phase, and the operator has selected step-by-step generation. Skill Generation has not been invoked and must remain paused pending a separate handoff; when invoked, it must pause for operator approval before creating any skill file.

## B — Desired position

A ready-to-use Codex skill at `output/release-operator` guides release work from current repository sources and presents a clear, evidence-based readiness state. It uses the supported release entry point from the repository root, keeps dry-run behavior as the default, identifies missing validation or version evidence, and requests an explicit publish decision only at the project-defined authority boundary. A publish invocation occurs only after affirmative operator authorization; otherwise the workflow remains non-publishing.

## Path from A to B

1. Create the Codex skill package at the recorded destination with metadata that routes repository release-operation requests to `release-operator`.
2. Encode a source-first discovery step that establishes the repository root and reads the project-owned release documentation and concrete release entry point before directing or executing release actions.
3. Build the safe release workflow around `scripts/release.sh`: invoke it from the repository root without `--publish` for the default dry-run and preserve its reported mode as execution evidence.
4. Add a readiness assessment that separately reports validation evidence, the proposed version, and the operator's publish decision; keep readiness unmet when a project-defined validation command, version source, or successful evidence is unavailable.
5. Add the publication authority stop so the workflow presents the exact publish action and waits for affirmative operator authorization before invoking `scripts/release.sh --publish`; retain the dry-run or stop when authorization is absent or declined.
6. Document the operator-facing sequence, readiness output, blocked-state behavior, and authorization boundary without introducing alternate release entry points or unsupported project mechanisms.
7. Validate the generated skill on Codex and evaluate representative dry-run, incomplete-readiness, authorized-publish, declined-publish, and ambiguous-authorization scenarios; refine any behavior that weakens the source-owned process or safety boundary.

## C — Completion criteria

- The generated artifact is a valid Codex skill named `release-operator` at `output/release-operator`.
- On a release-operation request, the skill resolves the repository root and uses local project sources to identify `scripts/release.sh` as the sole supported release entry point.
- Without explicit publish authorization, the skill runs or directs the release entry point in its default dry-run form, never adds `--publish`, and clearly reports the dry-run result.
- The skill displays a readiness assessment covering validation status, proposed version, and explicit operator decision, with the evidence or unmet reason visible for each item.
- The skill does not declare release readiness when validation has not passed, the proposed version is unavailable, or the required operator decision has not been made.
- When the project sources do not define validation or version retrieval, the skill identifies the missing project-owned evidence and stops short of readiness rather than inventing commands or values.
- Before any publish invocation, the skill presents the publish action and obtains affirmative operator authorization; absent, ambiguous, or declined authorization cannot produce a publish action.
- After affirmative authorization, the skill uses `scripts/release.sh --publish` from the repository root and reports the script's concrete result.
- The skill introduces no alternate release entry point and does not broaden the release process beyond locally established project authority.

## Generation contract

- **Destination:** `output/release-operator` within this isolated workspace.
- **Supported harness:** Codex only.
- **Generation viability:** Autonomous. The available sources and plan are sufficient to generate and validate the skill. The absent project-defined validation command and version source are intentional runtime readiness blockers that the skill can report safely.
- **Execution preference:** Step-by-step. Generation must present each proposed step and wait for operator approval before acting, beginning with a mandatory approval pause before any skill file is created.
- **Authority to act:** A future Skill Generation invocation may inspect this isolated workspace and prepare its first proposed generation step. It may create or update files only inside the destination skill after the operator approves file creation, and may run only approved, non-publishing validation and evaluation steps. Generation has no authority to invoke a real publish action.
- **Required user stops:** Do not invoke Skill Generation during intake. When generation is invoked separately, stop before creating the first skill file and at each subsequent step-by-step approval boundary. In the finished skill, stop before every `--publish` invocation until the operator gives explicit affirmative authorization for that publish action.
- **External dependencies or unavailable inputs:** No external dependency is required for generation. Current project sources provide no validation command or proposed-version source, so the finished skill must expose those gaps whenever no further repository source resolves them.
- **Evaluation and refinement handoff:** After generation, exercise the completion-criteria scenarios and refine the skill until the safe default, readiness evidence, source-owned entry point, and publish authority stop hold consistently.

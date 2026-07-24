# Plugin Dependency State

## Assignment

**Goal:** Establish the current pinned Skill Issue plugin dependency truth, the completed dependency migration, remaining sync or integration work, and dependency ordering relevant to project-state consolidation.

**Scope:** Inspect the submodule declaration and gitlink, initialized dependency checkout, standalone-plugin ownership contracts, CLI payload embedding and installation paths, dependency-sync and CLI-check workflows, repository instructions, current project-local integrations, and relevant local commit history.

**Exclusions:** Broad production lifecycle analysis, broad plan inventory, website analysis, remote-state verification, workflow execution, submodule initialization or updates, Git-state changes, and standalone-plugin modification.

## Sources

- `.gitmodules:1-3`; parent index gitlink and initialized checkout status for `dependencies/codex-skill-issue-plugin`.
- `AGENTS.md:43-71`; `.repository-privacy.md:1-11`.
- `dependencies/codex-skill-issue-plugin/AGENTS.md:1-27`; `dependencies/codex-skill-issue-plugin/README.md:25-39`.
- Pinned standalone commit `62c3921aca885df01ab9c8ca6d7ff20381b0c038` (`fix: use secret-free plugin synchronization`) and its parent `7b82668547614624ccacf8df20cb4f78c2092d1e` (`feat: publish standalone Skill Issue plugin`).
- `dependencies/codex-skill-issue-plugin/.agents/plugins/marketplace.json:1-20`; `dependencies/codex-skill-issue-plugin/plugins/skill-issue/.codex-plugin/plugin.json:1-30`.
- `dependencies/codex-skill-issue-plugin/plugins/skill-issue/skills/skill-intake/SKILL.md:1-15`; `dependencies/codex-skill-issue-plugin/plugins/skill-issue/skills/skill-intake/agents/openai.yaml:1-6`.
- `dependencies/codex-skill-issue-plugin/plugins/skill-issue/skills/skill-generation/SKILL.md:1-35`; `dependencies/codex-skill-issue-plugin/plugins/skill-issue/skills/skill-evaluation-and-refinement/SKILL.md:1-35`.
- `bundle.go:1-6`; `cli/internal/payload/assets/manifest.json:1-50`; `cli/internal/payload/payload.go:19-75`, `cli/internal/payload/payload.go:129-181`, and `cli/internal/payload/payload.go:227-263`.
- `cli/internal/lifecycle/lifecycle.go:57-93`; `cli/internal/installer/installer.go:54-68`, `cli/internal/installer/installer.go:319-379`, `cli/internal/installer/installer.go:469-515`, and `cli/internal/installer/installer.go:561-594`.
- `cli/internal/harness/harness.go:136-175`; `cli/README.md:59-67`; `README.md:168-181`.
- `cli/scripts/local-cli.sh:1-97`; `cli/scripts/build-cross-platform.sh:1-39`; `.github/workflows/cli-checks.yml:1-40`; `.github/workflows/sync-skill-issue-plugin.yml:1-52`.
- Parent migration commits `a7ee8ab0199076d38a297a370755fffc952f4382` (`refactor: package skills as marketplace plugin`) and `1ebbe4aaf7d65427e68aa9f1084893dbf4456075` (`refactor: source lifecycle skills from plugin submodule`); current parent commit `4979bd035680fd6e1142eabcd67b352d9221d713`; protected snapshot commit/tag `b64d9b5ff095191b971e4c2a953fbbc8bf3c352e` / `build-week-2026`.
- `.codex/skills/skill-intake:1`, `.codex/skills/skill-generation:1`, and `.codex/skills/skill-evaluation-and-refinement:1`; verified with `git ls-files -s`, `readlink`, and filesystem resolution checks.
- `showcase-skills/api-change-impact-mapper/workflow-prompt.md:1`; `showcase-skills/ci-failure-triage/workflow-prompt.md:4`; `showcase-skills/incident-timeline-builder/workflow-prompt.md:4-9`; `showcase-skills/release-readiness-checker/workflow-prompt.md:4-9`; `showcase-skills/repository-onboarding-guide/workflow-prompt.md:5-11`.
- Local read-only Git probes: `git submodule status`, `git ls-files -s`, `git ls-tree`, `git status --porcelain=v2`, branch-tree presence checks, parent and dependency logs, and dependency `main...origin/main` comparison.

## Findings

### Finding 1: The current local pin is exact, initialized, and clean

The parent index records gitlink `62c3921aca885df01ab9c8ca6d7ff20381b0c038`. The initialized dependency checkout is at the same commit, reports branch `main`, and has no dependency worktree changes. Within the local dependency clone, `main` and the locally stored `origin/main` both resolve to the same commit with a `0 0` divergence count. The plugin manifest at that commit identifies version `0.1.0+codex.20260723022002` and exposes the `plugins/skill-issue/skills/` directory (`dependencies/codex-skill-issue-plugin/plugins/skill-issue/.codex-plugin/plugin.json:1-14`).

**Evidence:** `.gitmodules:1-3` declares the canonical GitHub URL and checkout path. `git ls-files -s dependencies/codex-skill-issue-plugin` returned mode `160000` and commit `62c3921...`; `git submodule status` and `git -C dependencies/codex-skill-issue-plugin rev-parse HEAD` returned the same commit without a leading mismatch marker or dirty state. Commit `1ebbe4a...` introduced the gitlink already pointing at `62c3921...`, and no later parent commit changes that path.

**Implication:** The current local CLI source is pinned reproducibly to one initialized standalone-plugin revision; no local pointer update is presently indicated. Because this run was local-only and did not fetch, “current” means current local index and local remote-tracking state, not independently confirmed present-day GitHub `main`.

### Finding 2: Plugin authorship and CLI consumption have distinct semantic owners

The standalone repository owns plugin content, marketplace metadata, plugin version/cachebuster, and its three lifecycle skills. This parent repository owns the exact commit selection, embedding boundary, payload composition, harness-specific metadata adaptation, installer destinations, and CLI release/build behavior. The parent embeds only the standalone repository's `plugins/skill-issue/skills` subtree; it does not embed or install the standalone marketplace manifest or `.codex-plugin/plugin.json`.

**Evidence:** The dependency contract says all plugin-content changes belong under the standalone repository's `plugins/skill-issue/` and forbids editing through the submodule checkout (`dependencies/codex-skill-issue-plugin/AGENTS.md:1-8`). The parent repeats that boundary and defines synchronization rather than local editing (`AGENTS.md:43-61`). `bundle.go:5` embeds the dependency's skills subtree alongside parent-owned supporting skills and built-in evaluations. The parent payload manifest names three dependency-sourced lifecycle skills and eight parent-sourced components (`cli/internal/payload/assets/manifest.json:5-49`). Marketplace and plugin package metadata remain solely in the dependency (`dependencies/codex-skill-issue-plugin/.agents/plugins/marketplace.json:1-20`; `dependencies/codex-skill-issue-plugin/plugins/skill-issue/.codex-plugin/plugin.json:1-30`).

**Implication:** Skill-body, reference, `agents/openai.yaml`, plugin-version, or marketplace changes must originate in the standalone repository. Changes to what the CLI includes, how a harness receives files, how installation replaces directories, or which parent revision is released belong in this repository. Consolidation should preserve this split rather than recreate a second plugin-content owner here.

### Finding 3: The CLI consumes the pinned skill tree end to end

Go compilation embeds the dependency skill tree into `skillissue.CanonicalSkills`. The payload manifest supplies stable component IDs and repository-relative source paths. `payload.Skills` walks the embedded paths, copies every file except `.DS_Store`, requires and validates `SKILL.md` frontmatter, and validates local reference closure. The lifecycle service delegates installation to the installer, which resolves a harness-native skill root, applies harness metadata where supported, stages allowed files, atomically replaces each known skill directory, and verifies expected files exist.

**Evidence:** `bundle.go:3-6` owns the embedded filesystem. `cli/internal/payload/payload.go:19-20` embeds the manifest; `cli/internal/payload/payload.go:129-181` resolves every manifest component against `CanonicalSkills`; `cli/internal/payload/payload.go:227-263` checks referenced-file closure. `cli/internal/lifecycle/lifecycle.go:57-93` routes `install` to `installer.Service.Install`. `cli/internal/installer/installer.go:54-68` loads canonical skills and applies metadata; `cli/internal/harness/harness.go:136-175` selects the native root and filters harness-owned metadata; `cli/internal/installer/installer.go:561-594` performs staged replacement. The documented behavior matches the source (`cli/README.md:59-67`).

**Implication:** A changed gitlink affects future CLI binaries without a separate copying step, provided the submodule is initialized at the indexed commit before compilation. The effective CLI payload is the combination of the exact dependency tree, parent manifest, root embed declaration, and installer/harness code; checking only the plugin package manifest is insufficient.

### Finding 4: Build and CI paths were migrated to preserve the exact pin

Development builds include the dependency path in dirty-state detection. Known-good builds archive committed parent `HEAD`, resolve that commit's gitlink, verify the dependency commit exists in the initialized checkout, archive that exact dependency commit into the temporary source tree, and then compile. CLI CI initializes submodules before formatting, vetting, and tests. Cross-platform release builds compile the resulting embedded payload with `-trimpath`.

**Evidence:** `cli/scripts/local-cli.sh:48-62` includes the submodule path in development state calculation. `cli/scripts/local-cli.sh:66-97` derives `HEAD:dependencies/codex-skill-issue-plugin`, refuses an unavailable commit, and archives exactly that object. `.github/workflows/cli-checks.yml:23-40` checks out submodules before CLI verification. `cli/scripts/build-cross-platform.sh:24-29` uses `go build -trimpath`. Repository instructions require submodule initialization before use (`AGENTS.md:63-71`) and preserve committed `HEAD` as the known-good baseline (`AGENTS.md:31-41`).

**Implication:** Dependency ordering for a trustworthy binary is: accepted parent commit and gitlink first, initialized matching dependency object second, then build. A development build from an uncommitted pointer/content state is labeled dirty; a known-good build cannot silently substitute the dependency checkout's current branch tip for the commit recorded by parent `HEAD`.

### Finding 5: The ownership migration is committed, but project-local path migration is incomplete

Commit `a7ee8ab...` first moved the lifecycle skills into an in-repository marketplace plugin at `plugins/skill-issue/`. Commit `1ebbe4a...` then added the submodule and sync workflow, redirected root embedding, payload sources, build preparation, documentation, and CI, and deleted the parent-owned plugin copy. That second commit did not update three tracked `.codex/skills` symlinks created by `a7ee8ab...`, so all three now point at deleted `plugins/skill-issue/...` paths. Five retained showcase workflow prompts also still instruct agents to read those deleted paths.

**Evidence:** `1ebbe4a...` adds `.gitmodules` and the gitlink, changes `bundle.go` and the first three payload sources to `dependencies/codex-skill-issue-plugin/...`, adds submodule-aware CI/build logic, and deletes `.agents/plugins/marketplace.json` plus `plugins/skill-issue/`. Git history shows `.codex/skills/skill-intake`, `skill-generation`, and `skill-evaluation-and-refinement` were last changed only by `a7ee8ab...`; their line-1 targets remain `../../plugins/skill-issue/skills/...`. Read-only resolution checks classify all three as broken. The five showcase prompt citations in Sources name the same absent path.

**Implication:** The compiled CLI dependency migration is complete, while repository-local discovery and retained workflow-invocation material are inconsistent with it. Project-local use through those symlinks is blocked, and the five prompts cannot follow their named source paths. Consolidation must decide whether `.codex/skills` should retarget the pinned dependency or defer to plugin installation, then update or remove the stale links accordingly; the showcase prompts must reference the surviving canonical route consistent with that decision.

### Finding 6: The sync workflow is secret-free and pointer-scoped, with a branch-routing blocker

The standalone repository deliberately removed its cross-repository, secret-backed updater in pinned commit `62c3921...`. The parent now owns a manually dispatched workflow that reads standalone `main`, exits if the pin matches, otherwise checks out the exact latest commit, runs `go test ./...`, creates `automation/skill-issue-plugin-<12-char-commit>`, commits only the dependency path, and pushes that branch. This respects the no-direct-write rule. However, the workflow file exists on `codex/post-submission-development` and is absent from protected/default `main` and `build-week-2026`; the documented `gh workflow run ...` command supplies no `--ref`. The current instructions also require `main` to remain unchanged and say only that the automation branch is opened for review, without naming its target base.

**Evidence:** Pinned commit `62c3921...` deletes `.github/workflows/update-skill-issue.yml` from the standalone repository and replaces secret/deploy-key instructions with the parent-triggered process. `.github/workflows/sync-skill-issue-plugin.yml:23-52` implements latest-commit comparison, testing, commit-specific branching, and pointer-only staging. `AGENTS.md:2-12` protects `main`; `AGENTS.md:50-61` documents the unqualified dispatch and review step. Local tree probes returned the sync workflow absent at `main`/`build-week-2026` and present at `codex/post-submission-development`.

**Implication:** Before the next standalone-plugin publication, the team must prove the workflow can be dispatched from the intended non-`main` line and establish the automation PR's base branch. Local source alone cannot validate GitHub's workflow-dispatch availability or default-ref behavior. Until that routing contract is resolved, the documented normal sync path is operationally unsupported; manual gitlink changes remain explicitly recovery-only.

### Finding 7: Next dependency actions have a clear order

1. Reconcile repository-local integration: choose the surviving canonical route for project-local lifecycle skills, repair or remove the three broken `.codex/skills` links, and update the five showcase prompts so every named path resolves.
2. Resolve the sync branch contract: confirm how `sync-skill-issue-plugin.yml` is dispatched while absent from protected `main`, and name the intended non-`main` review/merge base.
3. For any future plugin change, edit and validate only in the standalone repository, update the plugin cachebuster when required, commit and push standalone `main`, then invoke the resolved parent sync route (`dependencies/codex-skill-issue-plugin/AGENTS.md:10-24`).
4. Review the generated automation branch for the expected gitlink only; confirm the commit exists in the standalone repository and run parent CLI verification with submodules initialized before acceptance (`AGENTS.md:58-71`; `.github/workflows/sync-skill-issue-plugin.yml:43-52`).
5. After the pointer and integration fixes are accepted on the active non-`main` branch, build/select a new known-good CLI only from that committed revision (`cli/scripts/local-cli.sh:66-97`).

**Evidence:** The owner contracts, current broken paths, workflow implementation, branch protection, CLI checks, and known-good archive behavior cited in Findings 2, 4, 5, and 6 establish these prerequisites and handoffs.

**Implication:** The current pin itself does not need movement based on local evidence. The immediate recovery work is parent-repository integration and sync-route correctness; a future content update should follow only after those prerequisites prevent another partially migrated or unreviewable dependency state.

## Notes

- Remote freshness is unsupported in this local-only run. Local `origin/main` for the dependency equals the pin, but no fetch or GitHub query was performed.
- GitHub workflow-dispatch semantics were not externally checked. The branch-routing issue is therefore a source-backed operational blocker, not a claim that dispatch definitely fails.
- No build, test, sync, submodule update, or workflow was run. Validation evidence here comes from production source, commit history, object/index inspection, and non-destructive path probes.
- The parent repository cannot resolve the submodule commit with its own object database; the initialized dependency checkout does contain and check out the exact indexed commit. This is normal gitlink separation, not evidence of a missing dependency.

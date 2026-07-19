# Completed Foundation: Direct Harness Installation Architecture

## A: Starting Position

- Skill Issue has selected nine harness targets: GitHub Copilot, Claude Code, OpenAI Codex, Cursor, Google Antigravity / Gemini CLI, Grok, OpenCode, Kilo Code, and Pi.
- The product direction is fixed: Skill Issue will use its own CLI to install and remove canonical skills directly rather than shipping and maintaining a separate native plugin for every harness.
- The pure-Go CLI foundation is complete under `cli/`. It already owns stable lifecycle command routing, operating-system and architecture detection, embedded-payload plumbing, and atomic ownership receipts, but its canonical payload and harness adapters remain intentionally unimplemented.
- The internet-only direct-installation research is complete in `plans/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md`, supported by the ten reports under `plans/deep-research/harness-direct-installation-architecture/assignments/`.
- That research establishes a portable Agent Skill payload, one shared receipt-backed filesystem lifecycle, and host-specific adapters for paths, activation, verification, trust, policy, configuration, and unsupported surfaces.
- Several claims remain deliberately caveated: Claude's exact enterprise managed-skills subdirectory, Kilo Cloud's conflicting project path, Grok Build's loose-skill schema and collision behavior, Antigravity lifecycle details, and specified remote or cloud discovery surfaces. These gaps must produce explicit caveated or unsupported outcomes rather than guessed behavior.

## B: Desired Position

The completed research is adopted as the authoritative implementation contract for all later payload and CLI work. One strict portable Skill Issue payload is materialized into each harness's documented native project or user root through a shared transactional lifecycle, while narrowly scoped adapters own surface detection, exact paths, capability gates, activation, discovery, verification, trust, policy, and reversible configuration behavior. Every operation distinguishes filesystem materialization from host discovery and successful activation, preserves unrelated user state through an external ownership ledger, and fails closed when official evidence or live host proof is insufficient.

## Path from A to B

### 1. Adopt the canonical portable payload contract

- Use one immutable Agent Skill directory as the source payload for every supported direct-install target.
- Require the directory basename to match a unique lowercase hyphenated `name`, an exact-case `SKILL.md`, required YAML `name` and non-empty `description`, the established 64-character name limit, and a complete relative-file closure for every referenced resource.
- Permit portable supporting directories such as `scripts/`, `references/`, `assets/`, `examples/`, and `templates/` when the skill requires them.
- Keep host-specific frontmatter, plugin manifests, hooks, MCP or LSP configuration, permission widening, and other nonportable behavior outside the canonical core unless a later explicit host overlay is separately justified and owned.
- Install independent copies into each harness's documented native root by default rather than coupling multiple harnesses to one shared physical directory.

### 2. Adopt the compatibility matrix and adapter classifications

- Treat GitHub Copilot, Claude Code, OpenAI Codex, Cursor, OpenCode, local Kilo Code, and Pi as common filesystem adapters with host-specific descriptors.
- Treat Antigravity Desktop, Antigravity CLI, and Gemini CLI as one bespoke Google adapter family whose three surfaces retain distinct project roots, user roots, commands, and lifecycle behavior.
- Treat Grok Build as a bespoke caveated adapter: its project and user roots are documented, but installation is accepted only when `grok inspect --json` proves the canonical candidate was discovered.
- Treat Grok selected as a model inside Cursor as a Cursor installation, not as a Grok Build installation.
- Keep Kilo Cloud, Codex cloud, Cursor Background Agents, Claude cloud or Cowork propagation, and other unvalidated remote surfaces outside supported direct installation until authoritative evidence or live proof establishes their contracts.
- Offer only the project or user scopes documented for the exact detected surface; do not manufacture a universal system-wide scope.

### 3. Adopt the shared lifecycle and adapter contract

- Implement the shared lifecycle through `detect`, `enumerateScopes`, `preview`, `install`, `activate`, `verify`, `diagnose`, `update`, `repair`, `rollback`, and `uninstall` capabilities.
- Make the shared engine validate, preview, lock, stage, revalidate, back up owned state, atomically materialize where supported, record ownership, activate, verify, diagnose, update, repair, roll back, and uninstall.
- Make each adapter own only exact surface detection, path resolution, version and capability gates, activation or restart behavior, host discovery, explicit-use verification, trust and policy blockers, collision interpretation, and any reversible structural configuration patch.
- Record `materialized`, `discovered`, and `activated` as distinct outcomes. A successful file write is never sufficient proof that the host can see or use the skill.
- Return `restart-required`, `verification-pending`, `caveated`, `unsupported`, policy-denied, or another precise capability result when the harness cannot complete a lifecycle stage.

### 4. Enforce ownership, conflicts, and reversible state

- Keep an ownership ledger outside installed skill content containing the adapter and host versions, scope, root, final path, canonical name, source and payload versions, digests, file and executable inventory, timestamps, backup identity, exact configuration fragment, and the latest materialization, discovery, and activation results.
- Store no credentials and never treat mutable installed content as the only ownership proof.
- Enumerate every documented root that can provide the canonical name before writing, reject a same-name foreign candidate, never merge skill directories, and replace only ledger-owned state or a foreign target the user explicitly authorizes.
- Avoid configuration edits for normal native-root installation. Any advanced extra-path or visibility change must be structurally parsed, atomic, narrowly reversible, recorded in the ledger, and refused when policy owns the value or faithful preservation is unavailable.
- Make update replace the complete owned directory so stale files cannot survive, repair stop on absent or foreign ownership, rollback require a verified owned backup, and uninstall remove only receipt-owned state and its exact recorded configuration fragment.

### 5. Enforce trust, policy, and fail-closed boundaries

- Preview every file and identify executable helpers before installation, but do not execute helpers during installation.
- Preserve each harness's sandbox, approvals, network controls, project trust, enterprise policy, skill permissions, and managed configuration. Installation must not grant new capabilities or weaken policy.
- Capability-probe current host behavior as well as version-checking it. Apply the documented floors for OpenCode, Pi, Kilo CLI, and optional preview `gh skill` behavior without treating a version number alone as discovery proof.
- Treat managed-policy denial, project-trust refusal, unwritable roots, foreign collisions, missing required host capabilities, failed live discovery, and absent verified rollback state as true blockers for the requested operation.
- Do not silently fall back to a plugin, guessed path, duplicate root, remote catalog, configuration weakening, or another harness surface when the requested route is blocked.

### 6. Reconcile the contract into downstream implementation

- Make Work Block 2 build the strict portable payload through the existing embedded-payload owner and implement the lifecycle through the existing command, lifecycle-manager, and receipt boundaries rather than creating another installer architecture.
- Make Work Block 3 require native discovery and activation evidence before an installation route is considered evaluator-ready or end-to-end qualified.
- Make Work Block 4 publish an explicit support matrix that distinguishes a materializable target, a live-verified local adapter, a caveated surface, an unsupported remote surface, and an environment that remains untested.
- Preserve the research synthesis and assignment reports as the source material for implementation; verify a caveated claim against official documentation or the live target before promoting it into a requirement or support claim.

## C: Observable Completion Criteria

- The research synthesis and its assignments are named as the authoritative source material for direct-install implementation.
- One strict portable payload contract and one receipt-backed transactional lifecycle are defined for the nine selected targets.
- Every selected local harness surface has a documented native project or user route, adapter classification, activation behavior, discovery and explicit-use verification path, and limitation boundary.
- The adapter contract distinguishes materialization, discovery, and activation and can return precise pending, caveated, blocked, or unsupported outcomes.
- Ownership, collision handling, configuration changes, update, repair, rollback, partial recovery, and uninstall preserve unrelated user state and stop on ambiguous ownership.
- Claude managed paths, Kilo Cloud, Grok loose-skill semantics, Antigravity lifecycle gaps, and unvalidated remote surfaces remain explicit caveats instead of inferred implementation requirements.
- Work Blocks 2, 3, and 4 identify and consume their local consequences from this architecture without repeating the research or creating competing owners.

## Dependency Handoffs

### Upstream inputs

- The completed product, support, and evidence contract is authoritative for the selected nine, the unified five-harness minimum release tier, and the environments requiring complete installation, workflow, and evaluation proof.
- `plans/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md` is the authoritative synthesis for the direct-install architecture.
- `plans/deep-research/harness-direct-installation-architecture/assignments/` contains the nine surface reports and the cross-harness adapter contract that preserve source-level evidence and caveats.
- The completed CLI foundation record and `cli/README.md` define the existing CLI implementation owners this architecture must extend.

### Authoritative outputs

- This section owns the decision to use CLI-managed direct installation and removal rather than separate native harness plugins.
- This section owns the strict portable payload boundary, local-surface compatibility matrix, adapter classifications, shared lifecycle contract, materialized/discovered/activated distinction, ownership and conflict rules, and fail-closed limitation policy.
- The research synthesis remains the detailed path, command, version-gate, verification, trust, and policy reference. This section integrates those findings into the project plan without replacing the cited research evidence.

### Required downstream consumers

- Work Block 2 consumes the portable payload contract, canonical identity rules, complete referenced-file closure, host-overlay boundary, shared lifecycle, adapter classifications, exact native paths, capability gates, ledger requirements, conflict rules, activation, verification, caveats, and safe removal behavior.
- Work Block 3 consumes the supported local routes and the requirement for materialization, discovery, and activation evidence during evaluator preparation and end-to-end qualification.
- Work Block 4 consumes the support-state distinctions, caveats, provenance, digests, and honest public-claim boundaries for release artifacts and documentation.

## Unresolved Matters

- Claude's exact enterprise managed-skills subdirectory is unvalidated and is not an implementation target. Managed policy remains authoritative and may block ordinary user or project installation.
- Kilo Cloud documents `.kilocode/skills/` while current local Kilo documents `.kilo/skills/`; Cloud remains unsupported until live proof or corrected official documentation resolves the conflict.
- Grok Build documents direct skill roots and `grok inspect`, but not the loose-skill entry filename, frontmatter schema, naming validation, collision precedence, or host-native repair, rollback, and uninstall. Its adapter must fail closed unless live inspection proves discovery.
- Antigravity hot reload, collision order, and loose-skill transaction behavior remain unvalidated; its adapter must report those limitations rather than infer them from Gemini CLI.
- Codex cloud discovery, Cursor Background Agent discovery, and Claude cloud or Cowork propagation of local skills remain unvalidated remote surfaces.

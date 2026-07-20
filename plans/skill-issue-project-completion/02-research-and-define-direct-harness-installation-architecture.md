# Completed Foundation: Direct Harness Installation Architecture

## A: Starting Position

- Skill Issue has selected nine harness targets: GitHub Copilot, Claude Code, OpenAI Codex, Cursor, Gemini CLI, Grok, OpenCode, Kilo Code, and Pi.
- The product direction is fixed: Skill Issue will use its own CLI to install and remove canonical skills directly rather than shipping and maintaining a separate native plugin for every harness.
- The pure-Go CLI foundation is complete under `cli/`. Work Block 2 has since populated its canonical payload, implemented all nine native project and user roots, added direct disposable materialization, and built the evaluation runner. Detection, confirmation, preview, the Grok fail-closed preflight, and final qualification remain unfinished.
- The internet-only direct-installation research is complete in `research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md`, supported by the ten reports under `research/deep-research/harness-direct-installation-architecture/assignments/`.
- That research establishes a portable Agent Skill payload and host-specific adapters for paths, activation, verification, trust, policy, configuration, and unsupported surfaces. The implemented ordinary lifecycle is intentionally narrower than the research proposal: repeated installation replaces the known payload paths, and uninstallation removes them.

## B: Desired Position

The completed research is adopted as the authoritative path and support-boundary source for later payload and CLI work. One strict portable Skill Issue payload is materialized into each harness's documented native project or user root through a direct disposable lifecycle. Harness adapters own detection, exact paths, and any concrete preflight required by their real installation route. The CLI reports filesystem materialization; Work Block 3 separately proves host discovery and activation in qualified environments.

## Path from A to B

### 1. Adopt the canonical portable payload contract

- Use one immutable Agent Skill directory as the source payload for every supported direct-install target.
- Require the directory basename to match a unique lowercase hyphenated `name`, an exact-case `SKILL.md`, required YAML `name` and non-empty `description`, the established 64-character name limit, and a complete relative-file closure for every referenced resource.
- Permit portable supporting directories such as `scripts/`, `references/`, `assets/`, `examples/`, and `templates/` when the skill requires them.
- Keep host-specific frontmatter, plugin manifests, hooks, MCP or LSP configuration, permission widening, and other nonportable behavior outside the canonical core unless a later explicit host overlay is separately justified and owned.
- Install independent copies into each harness's documented native root by default rather than coupling multiple harnesses to one shared physical directory.

### 2. Adopt the compatibility matrix and adapter classifications

- Treat GitHub Copilot, Claude Code, OpenAI Codex, Cursor, OpenCode, local Kilo Code, and Pi as common filesystem adapters with host-specific descriptors.
- Treat Gemini CLI as the supported Google adapter. Antigravity Desktop and `agy` remain outside the current runner.
- Treat Grok Build as a bespoke fail-closed adapter: its project and user roots are documented, but installation is accepted only when `grok inspect --json` proves the canonical candidate was discovered.
- Treat Grok selected as a model inside Cursor as a Cursor installation, not as a Grok Build installation.
- Keep Kilo Cloud, Codex cloud, Cursor Background Agents, Claude cloud or Cowork propagation, and other remote surfaces outside supported direct installation.
- Offer only the project or user scopes documented for the exact detected surface; do not manufacture a universal system-wide scope.

### 3. Adopt the direct lifecycle and adapter contract

- Implement detection, scope enumeration, preview, install, and uninstall through the existing command, harness, and installer owners.
- Make detection present likely harnesses and the researched project and user destinations for user confirmation rather than silently choosing a target.
- Build preview from the same payload, native-path, executable-helper, and blocker facts used by installation.
- Make each adapter own exact surface detection, path resolution, and only the concrete blockers or restart information required by that installation route.
- Treat successful materialization as an installation result without implying that the host discovered or activated the skill. Work Block 3 records discovery and activation evidence.

### 4. Preserve unrelated harness content

- Replace only the known Skill Issue payload directories beneath the selected native root and leave every unrelated file and skill directory untouched.
- Treat the embedded payload as disposable. Repeated `install` replaces its known paths and is the only reinstall, update, or repair behavior.
- Make `uninstall` remove the same known Skill Issue paths directly without inspecting or deleting unrelated content.
- Avoid configuration edits for normal native-root installation.
- Store no ordinary installation receipt, hash inventory, backup, rollback state, or mutable application data.

### 5. Enforce trust, policy, and fail-closed boundaries

- Preview every destination, canonical skill directory and file, executable helper, same-name collision, and concrete blocker before installation, but do not execute helpers during installation.
- Preserve each harness's sandbox, approvals, network controls, project trust, enterprise policy, skill permissions, and managed configuration. Installation must not grant new capabilities or weaken policy.
- Probe a command, permission, or version only when the concrete installation route depends on it. Do not create generic capability metadata that no operation consumes.
- Treat unwritable roots and a concrete adapter preflight failure as blockers for installation.
- Do not silently fall back to a plugin, guessed path, duplicate root, remote catalog, configuration weakening, or another harness surface when the requested route is blocked.

### 6. Reconcile the contract into downstream implementation

- Make Work Block 2 build the strict portable payload through the existing embedded-payload owner and implement the direct lifecycle through the existing command, lifecycle-manager, and installer boundaries rather than creating another installer architecture.
- Make Work Block 3 require native discovery and activation evidence before an installation route is considered evaluator-ready or end-to-end qualified.
- Make Work Block 4 publish an explicit support matrix that distinguishes a materializable target, a live-verified local adapter, a fail-closed route, an unsupported remote surface, and an environment that remains untested.
- Preserve the research synthesis and assignment reports as implementation evidence. Excluded surfaces remain outside the product unless the support contract is explicitly expanded.

## C: Observable Completion Criteria

- The research synthesis and its assignments are named as the authoritative source material for direct-install implementation.
- One strict portable payload contract and one direct disposable lifecycle are defined for the nine selected targets.
- Every selected local harness surface has a documented native project or user route, adapter classification, and limitation boundary.
- Installation reports direct materialization without claiming discovery or activation; Work Block 3 owns that live evidence.
- Reinstallation and uninstallation replace or remove only the known Skill Issue payload paths and preserve unrelated user state.
- Claude managed paths, Kilo Cloud, and remote surfaces remain outside supported direct installation. Grok remains fail-closed unless native inspection proves discovery.
- Work Blocks 2, 3, and 4 identify and consume their local consequences from this architecture without repeating the research or creating competing owners.

## Dependency Handoffs

### Upstream inputs

- The completed product, support, and evidence contract is authoritative for the selected nine, the unified five-harness minimum release tier, and the environments requiring complete installation, workflow, and evaluation proof.
- `research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md` is the authoritative synthesis for the direct-install architecture.
- `research/deep-research/harness-direct-installation-architecture/assignments/` contains the nine surface reports and the cross-harness adapter contract that preserve source-level evidence and support boundaries.
- The completed CLI foundation record and `cli/README.md` define the existing CLI implementation owners this architecture must extend.

### Authoritative outputs

- This section owns the decision to use CLI-managed direct installation and removal rather than separate native harness plugins.
- This section owns the strict portable payload boundary, local-surface compatibility matrix, adapter classifications, shared lifecycle contract, ownership and conflict rules, installation-versus-qualification boundary, and fail-closed limitation policy.
- The research synthesis remains the detailed path, command, version-gate, verification, trust, and policy reference. This section integrates those findings into the project plan without replacing the cited research evidence.

### Required downstream consumers

- Work Block 2 consumes the portable payload contract, canonical identity rules, complete referenced-file closure, direct lifecycle, adapter classifications, exact native paths, concrete preflight rules, support boundaries, and safe removal behavior.
- Work Block 3 consumes the supported local routes and the requirement for materialization, discovery, and activation evidence during evaluator preparation and end-to-end qualification.
- Work Block 4 consumes the support-state distinctions, provenance, and honest public-claim boundaries for release artifacts and documentation.

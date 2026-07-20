---
name: dependency-upgrade-planner
description: Source-backed, dependency-ordered migration planning for requested dependency upgrades. Use when assessing or planning a dependency version change in an existing project without changing dependency state.
---

# Dependency Upgrade Planner

## Establish the Upgrade Contract

- Confirm the requested dependency, current constraint and resolved version, target version or range, ecosystem, affected workspace, delivery constraints, and planning authority.
- Treat absent target versions, ambiguous package identities, or unclear workspace boundaries as unresolved inputs. Ask only when repository evidence and authoritative upstream sources cannot settle a choice that changes the plan materially.
- Keep all dependency changes, installs, lockfile regeneration, codemods, and migrations outside scope unless the user separately authorizes implementation.

## Trace the Current System

1. Read production manifests and the active lockfile before relying on documentation, tests, or package-manager summaries.
2. Trace workspace declarations, overrides, resolutions, patches, registries, build configuration, runtime and platform constraints, generated code, and deployment packaging that can change dependency selection or compatibility.
3. Locate concrete production imports, API calls, configuration keys, plugins, adapters, generated artifacts, and startup or build entry points using the dependency.
4. Follow relevant wrappers and abstractions to the concrete usage or effect. Record tests only as validation surfaces after production behavior is understood.
5. Identify the current dependency path and classify each affected relationship as direct, transitive, build-tool, runtime, or platform. Preserve overlaps rather than forcing one label.

## Establish Upstream Requirements

- Prefer the dependency owner's release notes, migration guide, version policy, compatibility matrix, peer requirements, and supported runtime or platform declarations. Use package registry metadata and source repositories to close gaps.
- Record the exact source, version span, publication or release identity when available, and why each upstream statement applies to this project.
- Distinguish documented breaking changes and prerequisites from reasoned project implications. Mark missing, conflicting, stale, or target-version-inapplicable guidance as an unresolved risk.
- Verify compatibility claims across both sides of every affected edge: the upgraded dependency's requirements and the dependent package, toolchain, runtime, or platform's supported range.

## Build the Dependency-Ordered Plan

1. Compare the current project evidence with applicable upstream changes and eliminate changes unrelated to concrete project usage.
2. Identify prerequisites first: runtime, language, package manager, compiler, SDK, platform, peer dependency, configuration, or intermediate-version requirements.
3. Group migration work by the behavior owner that must change, then order groups by dependency edges rather than file proximity.
4. Place lockfile refreshes, generated-code updates, codemods, and package-manager operations only after their prerequisites and affected source or configuration decisions are explicit.
5. Attach validation gates to each reversible stage and define the evidence required to proceed. Include focused static checks, builds, tests, runtime probes, platform checks, and release checks only where the repository and upstream contract make them relevant.
6. Define rollback points around state-changing stages: files or artifacts to restore, versions or constraints to reinstate, generated state to discard, and conditions that trigger rollback.
7. Keep blocked decisions and unverified compatibility visible. Do not turn an unknown into an assumed task merely to make the sequence complete.

## Produce the Migration Plan

Write one coherent plan containing:

- requested change, scope, constraints, and inspected evidence;
- current dependency graph and concrete usage summary;
- applicable upstream requirements with authoritative sources;
- impact matrix separating direct, transitive, build-tool, runtime, and platform effects;
- breaking changes, prerequisites, and unresolved risks;
- dependency-ordered migration stages with behavior owners and affected surfaces;
- validation gates, required evidence, rollback points, and stop conditions; and
- explicit exclusions and follow-up questions.

Label statements as observed project facts, authoritative upstream facts, reasoned implications, or unknowns. Cite repository locations and upstream sources close to the claims they support. End with the smallest safe first stage and the evidence needed before any implementation begins.

## Completion Gate

- Confirm every planned task traces to a concrete project usage, dependency edge, applicable upstream requirement, or explicit unresolved risk.
- Confirm prerequisites precede dependent work and each state-changing stage has validation and rollback coverage.
- Confirm the plan makes no unsupported compatibility, successful-build, successful-test, or runtime claim.
- Report inspection limitations and unavailable authoritative evidence instead of silently broadening certainty.

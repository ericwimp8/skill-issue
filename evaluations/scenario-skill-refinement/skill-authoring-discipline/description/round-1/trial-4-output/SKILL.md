---
name: dependency-upgrade-review
description: Dependency upgrade review constraints and workflow. Use when assessing proposed or completed dependency upgrades.
---

# Dependency Upgrade Review

## Review

Establish the repository's ecosystem, dependency sources, resolved versions, and validation conventions before assessing the upgrade.

Evaluate these concerns separately:

- **Compatibility risk:** Identify affected APIs, runtimes, toolchains, transitive dependencies, and supported version ranges.
- **Security relevance:** Confirm whether an advisory applies to the repository's resolved dependency and exposure, and whether the upgrade reaches a fixed version.
- **Migration work:** Identify required source, configuration, build, data, or operational changes from authoritative release and migration guidance.

Trace each conclusion to repository evidence or authoritative dependency documentation. Label unresolved questions and assumptions.

Use the repository's own package-management and validation commands. Recommend an upgrade only with the necessary migration and verification work stated explicitly.

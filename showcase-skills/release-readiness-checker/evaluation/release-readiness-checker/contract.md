# Release Readiness Checker Evaluation Contract

- **Target:** `showcase-skills/release-readiness-checker/skill/release-readiness-checker/SKILL.md`
- **Goal:** decide whether an exact release candidate is ready using authoritative project-derived gates and current evidence, without publishing or deploying.
- **Intended use:** repository pre-release reviews, candidate go/no-go assessments, and release evidence audits.
- **Expected behavior:** establish candidate identity and limits; derive and trace applicable gates; inspect current evidence; preserve commands and limitations; classify every gate; and produce a conservative prioritized decision with exact next actions.
- **Expected result:** a standalone release-readiness report with scoped candidate identity, `ready`, `not ready`, or `undetermined` decision, gate-level statuses, risks, next actions, and evidence provenance.
- **Preserved boundaries:** no release side effects, invented policy, inferred execution, stale-proof substitution, generic universal checklist, or hidden incomplete gate.
- **Evaluation surface:** generated Markdown report from isolated synthetic repository fixtures.

## Observable Criteria

1. Candidate identity, target, scope, authority sources, and evaluation limitations are explicit.
2. Gates derive from authoritative project sources and conflicts remain visible.
3. Relevant implementation and operational paths are traced to concrete owners.
4. Applicable code, configuration, migrations, versioning, documentation, verification, security or privacy, rollback, and risk evidence is inspected contextually.
5. Every gate uses exactly one defined status with candidate-specific evidence and limitations.
6. Current execution evidence is distinguished from source support, configured capability, documentation, and historical results.
7. The decision follows project policy or a disclosed conservative rule and never treats `blocked` or `not-run` critical gates as ready.
8. Prioritized next actions name the action, owner when known, required evidence, and gate resolved.
9. The report preserves exact command evidence and identifies deliberately unrun checks.
10. No action publishes, deploys, tags, uploads, promotes, signs, submits, migrates shared state, or mutates external release state.

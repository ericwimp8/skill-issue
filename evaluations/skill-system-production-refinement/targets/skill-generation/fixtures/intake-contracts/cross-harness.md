# Explicit Risk Review Skill A-to-B Plan

## A — Current position

- Users explicitly request a read-only review of a risk document.
- The initial supported surfaces are Codex, Claude Code, Pi, and OpenCode.
- The skill must never be selected implicitly and must never edit the reviewed document.

## B — Desired position

A portable `explicit-risk-review` skill reports evidence-backed risks and severity from a supplied document, with the strongest documented explicit-only control on each supported harness.

## Path from A to B

1. Create one canonical skill body and concise explicit-only description.
2. Add only the host metadata required for the four supported surfaces.
3. Validate canonical structure, metadata ownership, and the read-only boundary.
4. Hand runtime selection and behavior proof to evaluation.

## C — Completion criteria

- The canonical skill reports risks, severity, and evidence without editing source documents.
- Codex uses `policy.allow_implicit_invocation: false` in `agents/openai.yaml`.
- Claude Code and Pi receive supported explicit-only metadata overlays.
- OpenCode uses the explicit-only description fallback without unsupported frontmatter claims.
- Platform-specific metadata does not duplicate the canonical body.
- Structural checks pass; runtime reliability remains deferred.

## Generation contract

- Destination: `evaluations/skill-system-production-refinement/targets/skill-generation/behavior/cycle-1/case-4/output/explicit-risk-review`
- Generation viability: Autonomous.
- Selected execution preference: Autonomous.
- Authority: create only inside the destination and supporting platform-metadata subdirectory.
- Supported harnesses: Codex, Claude Code, Pi, OpenCode.
- Evaluation handoff: retain explicit-only invocation and read-only behavior for runtime proof.

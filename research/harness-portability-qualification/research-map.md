# Remaining Harness Portability Qualification Research Map

## Goal

Reconstruct the qualified evaluation pattern implemented for OpenAI Codex, Cursor, Claude Code, and Pi, then apply the same two qualification gates to GitHub Copilot, Gemini CLI, Grok Build, and OpenCode without expanding Skill Issue into harness installation or environment management.

## Scope And Output

- Source scope: local plus internet.
- Active researcher concurrency: four.
- Total researcher budget: sixteen.
- Assignment root: `research/harness-portability-qualification/assignments/`.
- Final aggregation target: `research/harness-portability-qualification/remaining-harness-portability-qualification-audit.md`.
- Final shape: best-supported direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers.

## Research Domains

1. Existing qualified implementation: concrete runtime, setup, replay, attribution, process ownership, and cleanup paths for Codex, Cursor, Claude Code, and Pi.
2. Existing product contract: support boundaries, installation ownership, completion criteria, local candidate research, and retained smoke evidence.
3. Candidate technical viability: clean/default isolation, temporary skills, noninteractive operation, resumable conversation, protocol evidence, permissions, cancellation, and bounded cleanup.
4. Candidate Codex-subscription viability: direct access, supported provider mechanisms, safe proxy or compatibility arrangements, and authentication ownership.
5. Cross-checking: reconcile official documentation, primary source code, local production behavior, and unresolved evidence gaps without speculative workarounds.

## Discovery And Fan-Out

### Wave 1: Pathway Discovery

- `01-existing-runtime-implementation.md`: trace the four concrete adapters end to end.
- `02-product-support-and-setup-contract.md`: extract controlling support and qualification contracts.
- `03-local-candidate-research-copilot-gemini.md`: map existing local evidence for Copilot and Gemini CLI.
- `05-existing-auth-and-subscription-patterns.md`: reconstruct qualified authentication and Codex-access arrangements.

### Wave 2: Candidate Deep Dives

Two narrow assignments per candidate: one for the first technical gate and one for the second Codex-subscription gate. Official documentation and primary repositories are the preferred sources. Discovery results determine exact source targets and caveats.

### Wave 3: Candidate Evidence Cross-Checks

One assignment per candidate audits the preceding local and upstream documents, resolves conflicts where possible, classifies unsupported claims, and recommends one allowed decision vocabulary outcome.

## Candidate Classification

- Deep-dive: the four named candidates, because each is an explicit campaign target.
- Skim-only: adjacent provider bridges or proxy arrangements that are officially documented but fail a core safety or ownership constraint.
- Reject: unofficial credential copying, permanent configuration mutation, invasive home-directory substitution, installation or repair ownership, and undocumented compatibility claims.

## Expected Evidence

- Repository-relative production paths and concrete symbols.
- Official documentation URLs, versioned primary repository paths, commits or releases when material.
- Safe read-only local executable evidence only when already available.
- Explicit separation of documented capability, source-code evidence, observed local evidence, inference, and unresolved uncertainty.
- Candidate decisions: proceed to implementation; blocked pending named evidence or access; or remove from the supported campaign.

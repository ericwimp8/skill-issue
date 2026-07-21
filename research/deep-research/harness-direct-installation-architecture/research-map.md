# Harness Direct Installation Architecture Research Map

## Run Contract

- Goal: define the direct CLI-managed install, verify, update, repair, rollback, diagnostics, and removal lifecycle for Skill Issue across GitHub Copilot, Claude Code, OpenAI Codex, Cursor, Google Antigravity / Gemini CLI, Grok, OpenCode, and Pi.
- Product boundary: one Skill Issue cross-platform lifecycle CLI installs a canonical payload directly; per-harness native plugins and package-manager delivery are outside the chosen architecture.
- Source scope: internet only.
- Active researcher concurrency: 5.
- Total researcher budget: 10.
- Final aggregation target: `research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md`.
- Final output shape: best-supported direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers.

## Research Domains

### Harness Installation Contracts

Establish the authoritative filesystem, metadata, configuration, activation, discovery, verification, collision, trust, version-gate, and reversible lifecycle contract for each selected harness. Each harness is a narrow assignment because its supported scopes and surfaces can differ materially.

### Cross-Harness Adapter Architecture

Compare the eight validated harness contracts after their assignments complete. Identify common declarative filesystem behavior, genuine bespoke lifecycle requirements, unsupported targets, and the minimum shared adapter contract for later CLI implementation.

## Discovery Waves

### Wave 1: Major Harness Contracts

Run five researchers in parallel:

1. `assignments/01-github-copilot-direct-installation.md` — GitHub Copilot skills or closest equivalent across IDE and CLI surfaces.
2. `assignments/02-claude-code-direct-installation.md` — Claude Code project and user skill lifecycle.
3. `assignments/03-openai-codex-direct-installation.md` — Codex CLI/app skill lifecycle and discovery behavior.
4. `assignments/04-cursor-direct-installation.md` — Cursor project and user rules/skills lifecycle.
5. `assignments/05-google-antigravity-gemini-cli-direct-installation.md` — Google Antigravity and Gemini CLI customization surfaces and differences.

Expected evidence: current official documentation, first-party repositories where authoritative, documented paths and manifests, activation/discovery requirements, and current version or feature gates.

### Wave 2: Remaining Harness Contracts

Backfill three researchers after Wave 1 finishes:

6. `assignments/06-grok-direct-installation.md` — determine whether a standalone installable Grok coding harness exists and reject Cursor-model conflation.
7. `assignments/07-opencode-direct-installation.md` — OpenCode project and user skill lifecycle.
9. `assignments/09-pi-direct-installation.md` — Pi package, skill, prompt, extension, or closest-equivalent direct lifecycle.

Expected evidence: the same lifecycle contract as Wave 1, with special emphasis on standalone-product identity for Grok and package-manager-versus-direct-filesystem boundaries for Pi.

### Wave 3: Ranked Adapter Synthesis Evidence

After all eight harness assignments finish, run:

10. `assignments/10-cross-harness-adapter-contract.md` — inspect the eight assignment documents, cross-check their primary citations where necessary, rank common-adapter, bespoke-adapter, and unsupported candidates, and propose the shared adapter contract.

Expected evidence: a traceable comparison table derived from the eight harness documents, validation of disputed or cross-cutting claims against primary sources, and explicit deep-dive/skim/reject classification for any competing installation interpretations.

## Fan-Out Decisions

- One researcher per harness prevents broad domain assignments from hiding surface-specific installation differences.
- Google Antigravity and Gemini CLI remain paired because the requested target names them as one family; the assignment must separate materially distinct surfaces rather than assume shared behavior.
- Grok remains its own evidence lane because unsupported classification is a possible result and must be proven from current first-party sources.
- The tenth assignment is reserved for cross-harness comparison after primary harness evidence exists; no budget remains for uncontrolled branches.
- Lower-priority sources, speculative integrations, native plugin packaging, or undocumented paths must be classified as skim-only, rejected, unsupported, or blocked within the owning assignment.

## Shared Assignment Schema

Every assignment document must contain:

- `# <Research Assignment Name>`
- `## Assignment` with `Goal`, `Scope`, and `Exclusions`
- `## Sources` with inspected URLs, source names, versions, commits, symbols, publications, or docs
- `## Findings` with repeatable finding headings; each finding includes prose, `Evidence`, and `Implication`
- `## Notes` only for relevant dead ends, caveats, unsupported observations, or useful search terms

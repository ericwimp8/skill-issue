# Manual Skill Invocation Across Eight Harnesses: Research Map

## Run Parameters

- Goal: determine how a user manually invokes the installed skill `dictate-plan` in eight named harnesses.
- Source scope: internet only.
- Active researcher concurrency: 8.
- Total researcher budget: 8.
- Final aggregation target: `research/deep-research/manual-skill-invocation-eight-harnesses/manual-skill-invocation-eight-harnesses-deep-research.md`.
- Requested synthesis shape: best-supported answer or direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers.

## Research Domains

- Dedicated skill invocation surfaces: slash commands, skill commands, mentions, palettes, or other documented selectors.
- Explicit natural-language invocation: documented or best-supported wording when no dedicated selector exists.
- Evidence quality: current official product documentation first, then first-party repositories, announcements, or maintainers' materials.

## Discovery Wave

Each harness receives one narrow assignment. The researcher first maps the product's current skill terminology and official documentation pathway, then checks only the manual invocation question. No cross-harness discovery assignment is needed because the caller fixed the complete candidate set and the budget equals the eight harnesses.

## Assignments

| # | Harness | Source targets | Expected evidence | Output |
|---|---|---|---|---|
| 01 | GitHub Copilot | GitHub Docs and first-party GitHub sources | Dedicated invocation syntax or explicit wording | `assignments/01-github-copilot.md` |
| 02 | Claude Code | Anthropic Docs and first-party Anthropic sources | Dedicated invocation syntax or explicit wording | `assignments/02-claude-code.md` |
| 03 | OpenAI Codex | OpenAI Docs and first-party OpenAI sources | Dedicated invocation syntax or explicit wording | `assignments/03-openai-codex.md` |
| 04 | Cursor | Cursor Docs and first-party Cursor sources | Dedicated invocation syntax or explicit wording | `assignments/04-cursor.md` |
| 05 | Google Antigravity or Gemini CLI | Google/Gemini official docs and first-party repositories | Product-specific invocation syntax or explicit wording | `assignments/05-google-antigravity-gemini-cli.md` |
| 06 | Grok Build | xAI official docs and first-party xAI sources | Dedicated invocation syntax or explicit wording | `assignments/06-grok-build.md` |
| 07 | OpenCode | Official OpenCode docs and first-party repository | Dedicated invocation syntax or explicit wording | `assignments/07-opencode.md` |
| 09 | Pi | Official Pi docs and first-party repository | Dedicated invocation syntax or explicit wording | `assignments/09-pi.md` |

## Fan-Out Decisions

- Deep dives: all eight caller-selected harnesses, one assignment each.
- Skim-only candidates: none; the caller requires a direct answer for every harness.
- Rejected candidates: none; no ecosystem discovery beyond the fixed harness list is in scope.
- Mid-run branches: disallowed by the exhausted researcher budget; unsupported product terminology or absent official evidence must be recorded as uncertainty in the owning assignment.

## Exclusions

Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated skill behavior are outside the research scope.

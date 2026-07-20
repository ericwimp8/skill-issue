# Harness Skill Delivery

## Document Authority

This is the current generation reference for selected harness paths, canonical payload boundaries, and host metadata placement. The completed direct-install architecture owns the exact installation contract. Historical native-plugin packaging research remains ecosystem context and does not define current delivery requirements.

## Selection Rule

Keep one strict canonical skill directory. Generate portable Agent Skill content and only the resources its behavior requires. Place host-specific metadata in an explicit overlay, and let the Skill Issue CLI materialize the finished payload into the exact target surface's documented native project or user root.

A generated skill does not install the Skill Issue CLI, authorize accounts, grant permissions, add secrets, trust a project, or configure unrelated policy.

| Harness surface     | Native delivery roots                                                                                    | Generation implication                                                                                                                              |
| ------------------- | -------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| GitHub Copilot CLI  | Project `.github/skills/<name>/`; user `~/.copilot/skills/<name>/`                                       | Emit standard skill content and only supported Copilot frontmatter; keep host lifecycle outside the skill body                                      |
| Claude Code         | Project `.claude/skills/<name>/`; user `~/.claude/skills/<name>/` or `$CLAUDE_CONFIG_DIR/skills/<name>/` | Emit standard content plus supported invocation metadata when required; do not target guessed managed-policy paths                                  |
| OpenAI Codex        | Project or user `.agents/skills/<name>/`                                                                 | Keep canonical frontmatter strict and place Codex policy or UI metadata in optional `agents/openai.yaml`                                            |
| Cursor              | Project `.cursor/skills/<name>/`; user `~/.cursor/skills/<name>/`                                        | Emit portable content and supported Cursor frontmatter only; do not infer Background Agent discovery from local support                             |
| Antigravity Desktop | Project `.agents/skills/<name>/`; user `~/.gemini/config/skills/<name>/`                                 | Keep the portable skill self-contained and keep desktop-specific permissions or plugin components outside the canonical body                        |
| Antigravity CLI     | Project `.agents/skills/<name>/`; user `~/.gemini/antigravity-cli/skills/<name>/`                        | Keep this surface distinct from Gemini CLI and desktop global roots                                                                                 |
| Gemini CLI          | Project `.gemini/skills/<name>/`; user `~/.gemini/skills/<name>/`                                        | Keep extension commands, policies, hooks, agents, settings, and MCP outside ordinary skill generation                                               |
| Grok Build          | Project `.grok/skills/<name>/`; user `~/.grok/skills/<name>/`                                            | Generate the strict portable candidate without inventing Grok-specific frontmatter; accept delivery only after live `grok inspect --json` discovery |
| OpenCode            | Project `.opencode/skills/<name>/`; user `~/.config/opencode/skills/<name>/`                             | Use standard portable frontmatter and keep permission or plugin configuration outside the skill body                                                |
| Kilo Code local     | Project `.kilo/skills/<name>/`; user `~/.kilo/skills/<name>/`                                            | Use standard frontmatter and bundled resources; do not duplicate into `.kilocode/skills` or claim Kilo Cloud support                                |
| Pi                  | Project `.pi/skills/<name>/`; user `~/.pi/agent/skills/<name>/`                                          | Use Pi-supported Agent Skills metadata; keep extensions or packages separate unless the intake contract genuinely requires them                     |

The selected-nine product boundary contains one Google lineage entry whose desktop, Antigravity CLI, and Gemini CLI surfaces require distinct adapters. Grok selected as a model inside Cursor uses Cursor delivery; Grok Build is the separate standalone harness target.

## Portable Payload Boundary

- Require an exact-case `SKILL.md`, lowercase hyphenated name matching the directory, required `name` and `description`, and complete referenced-file closure.
- Keep portable supporting files relative and contained within the skill directory.
- Keep plugin manifests, hooks, MCP or LSP configuration, permission widening, and unrelated host configuration outside the canonical skill.
- Prefer documented native roots. Advanced extra-path configuration belongs to the CLI adapter and requires structural, reversible handling.

## Explicit-Only Overlay

For `skill-intake`, preserve the canonical explicit-only description and apply `packaging/overlays/explicit-only-agent-skill.yaml` only to documented compatible targets. Codex uses `agents/openai.yaml`. Other surfaces use the description guidance unless their current authoritative contract establishes an enforceable field.

## Authoritative References

- Product support boundary: `../../../plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`.
- Direct-install architecture: `../../../plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md`.
- Detailed source evidence: `../../../research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md` and its `assignments/` directory.

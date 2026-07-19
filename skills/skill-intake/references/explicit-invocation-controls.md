# Explicit-Only Invocation Controls

## Document Authority

This is the current product reference for keeping Skill Intake user-controlled across the selected harnesses. The selected-nine boundary is owned by the completed product support contract; native paths and caveats are owned by the completed direct-install architecture. Update this reference when either owner changes.

Historical plugin-packaging and JetBrains-era reports remain research evidence and must not be used as the current installation or support matrix.

## Shared Requirement

Skill Intake is user-controlled. Apply the strongest native control available and keep this sentence in its description on every surface: **Use only when the user explicitly requests skill intake or directly invokes this skill.** The canonical `SKILL.md` keeps portable frontmatter. Render `packaging/overlays/explicit-only-agent-skill.yaml` only for targets that support those additional fields.

Configuration support does not establish measured invocation reliability.

## Harness Matrix

| Harness surface     | Native skill roots                                                                                       | Explicit-only implementation                                                                                                                                 | Support classification                                                          |
| ------------------- | -------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------- |
| OpenAI Codex        | Project or user `.agents/skills/<name>/`                                                                 | Set `policy.allow_implicit_invocation: false` in `agents/openai.yaml`; retain explicit-only description guidance                                             | Enforced through Codex metadata                                                 |
| Claude Code         | Project `.claude/skills/<name>/`; user `~/.claude/skills/<name>/` or `$CLAUDE_CONFIG_DIR/skills/<name>/` | Set `disable-model-invocation: true`; retain user invocation                                                                                                 | Enforced on documented local skill surfaces                                     |
| GitHub Copilot CLI  | Project `.github/skills/<name>/`; user `~/.copilot/skills/<name>/`                                       | Set `disable-model-invocation: true` and leave user invocation available                                                                                     | Enforced on documented CLI skill surfaces                                       |
| Cursor              | Project `.cursor/skills/<name>/`; user `~/.cursor/skills/<name>/`                                        | Set `disable-model-invocation: true` and invoke the named skill explicitly                                                                                   | Enforced on documented native skill surfaces; verify installed-version behavior |
| Antigravity Desktop | Project `.agents/skills/<name>/`; user `~/.gemini/config/skills/<name>/`                                 | No documented enforceable per-skill field; use the explicit-only description and named invocation guidance                                                   | Guidance fallback                                                               |
| Antigravity CLI     | Project `.agents/skills/<name>/`; user `~/.gemini/antigravity-cli/skills/<name>/`                        | No documented enforceable per-skill field; use the explicit-only description and named invocation guidance                                                   | Guidance fallback; keep separate from Gemini CLI                                |
| Gemini CLI          | Project `.gemini/skills/<name>/`; user `~/.gemini/skills/<name>/`                                        | No documented enforceable per-skill field; use the explicit-only description and direct invocation guidance                                                  | Guidance fallback                                                               |
| Grok Build          | Project `.grok/skills/<name>/`; user `~/.grok/skills/<name>/`                                            | The retained official corpus does not establish a loose-skill explicit-only field; use description guidance and require live `grok inspect --json` discovery | Caveated guidance fallback                                                      |
| OpenCode            | Project `.opencode/skills/<name>/`; user `~/.config/opencode/skills/<name>/`                             | Use the explicit-only description. `permission.skill` may ask or deny access but does not prevent model selection                                            | Guidance fallback                                                               |
| Kilo Code local     | Project `.kilo/skills/<name>/`; user `~/.kilo/skills/<name>/`                                            | Only standard Agent Skills fields are documented; use explicit-only description and named invocation                                                         | Guidance fallback; Kilo Cloud unsupported                                       |
| Pi                  | Project `.pi/skills/<name>/`; user `~/.pi/agent/skills/<name>/`                                          | Set `disable-model-invocation: true`; users invoke `/skill:skill-intake`                                                                                     | Enforced on documented local surfaces                                           |

## Packaging Boundary

The selected nine provide a documented local skill route or a caveated route explicitly named above. Preserve one canonical skill body and render only supported host metadata. Do not claim that skill installation also installs a standalone CLI, authorizes accounts, grants permissions, adds secrets, trusts a project, or configures unrelated policy.

Grok Build remains caveated until live inspection proves the installed candidate. Kilo Cloud and the remote surfaces excluded by the direct-install contract are not implied by local support.

## Authoritative References

- Product support boundary: `../../../plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`.
- Direct-install paths and caveats: `../../../plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md`.
- Source-backed harness details: `../../../plans/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md` and its `assignments/` directory.

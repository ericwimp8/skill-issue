# Harness Evaluation Controls

## Document Authority

This is the current product reference for description-evaluation gates and independent-agent controls. The selected-nine boundary and minimum qualification matrix are owned by `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`. Native paths, discovery, activation, trust, and caveats are owned by `plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md` and its research synthesis.

Update this reference whenever either owning contract changes. Historical plugin-packaging and JetBrains-era reports remain research evidence and do not define the current matrix.

Current sub-agent support, launch behavior, and skill-request behavior are recorded in `plans/research/harness-subagent-launch/harness-subagent-launch-reference.md`.

## Evidence Standard

Count proactive invocation only when the harness exposes native evidence that the target skill loaded before its instructions affected the response. Suitable evidence includes a skill-load event, tool trace, execution transcript, debug log, structured inspection result, or equivalent harness record that identifies the skill. Agent prose claiming that it followed the skill is insufficient.

Run description trials only when all four conditions hold:

1. the exact harness surface can discover the installed target;
2. the target is configured as implicitly invocable;
3. a fresh independent agent or isolated equivalent can run each trial;
4. native evidence can distinguish proactive invocation from ordinary reasoning.

Treat model identifiers and measured reliability as environment-evaluation inputs. This reference establishes configuration capability and evaluation gates only.

## Planned Qualification Classes

- **Minimum qualification tier:** OpenAI Codex, Claude Code, Cursor, Pi, and OpenCode, subject to the four gates above for the exact installed version and model.
- **Additional selected targets:** GitHub Copilot, the Google Antigravity or Gemini CLI family, Grok Build, and local Kilo Code when their installed surfaces provide every required gate.
- **Caveated or unsupported surfaces:** Grok description trials until its loose-skill schema, independent-agent route, and native invocation evidence are proven live; Kilo Cloud; Codex cloud; Cursor Background Agents; Claude cloud or Cowork propagation; and other remote surfaces not established by the direct-install contract.

Record exact model identifiers and reasoning settings in the environment result rather than hard-coding them into this workflow.

## Capability Matrix

| Harness surface     | Native skill roots                                                                                       | Implicit-invocation control                                                                                                      | Independent-agent path                                                 | Description-evaluation status                                                                                    |
| ------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| OpenAI Codex        | Project or user `.agents/skills/<name>/`                                                                 | `agents/openai.yaml` may set `policy.allow_implicit_invocation`; enable it for description trials                                | Natural-language delegation handled by Codex                           | Minimum-tier candidate when native skill-load evidence is retained                                               |
| Claude Code         | Project `.claude/skills/<name>/`; user `~/.claude/skills/<name>/` or `$CLAUDE_CONFIG_DIR/skills/<name>/` | Leave `disable-model-invocation` absent or false for trials                                                                      | `Agent` tool; skills may use `context: fork`                           | Minimum-tier candidate when the transcript identifies skill loading                                              |
| Cursor              | Project `.cursor/skills/<name>/`; user `~/.cursor/skills/<name>/`                                        | Leave `disable-model-invocation` absent or false for trials                                                                      | Automatic or explicit parent-Agent delegation                          | Minimum-tier candidate only when the installed surface exposes skill-load evidence                               |
| Pi                  | Project `.pi/skills/<name>/`; user `~/.pi/agent/skills/<name>/`                                          | Leave `disable-model-invocation` absent or false for trials                                                                      | Official extension's `subagent` tool                                   | Minimum-tier candidate, configuration-dependent and trust-gated                                                  |
| OpenCode            | Project `.opencode/skills/<name>/`; user `~/.config/opencode/skills/<name>/`                             | Description discovery is model-driven; `permission.skill` can allow, ask, or deny access but is not a proactive-selection switch | `Task` tool or explicit `@name` request                                | Minimum-tier candidate with version-gated native trace evidence                                                  |
| GitHub Copilot CLI  | Project `.github/skills/<name>/`; user `~/.copilot/skills/<name>/`                                       | Leave `disable-model-invocation` absent or false for trials                                                                      | Model delegation; `/fleet` in CLI; forked context in supported VS Code | Additional target until installed-version isolation and traceability are proven                                  |
| Antigravity Desktop | Project `.agents/skills/<name>/`; user `~/.gemini/config/skills/<name>/`                                 | Relevance discovery; no cross-version enforceable per-skill switch is established                                                | `invoke_subagent` through the parent agent                             | Additional target; gate on exact desktop surface and native evidence                                             |
| Antigravity CLI     | Project `.agents/skills/<name>/`; user `~/.gemini/antigravity-cli/skills/<name>/`                        | Relevance discovery; no cross-version enforceable per-skill switch is established                                                | `invoke_subagent` where available                                      | Additional target; do not borrow Gemini CLI controls                                                             |
| Gemini CLI          | Project `.gemini/skills/<name>/`; user `~/.gemini/skills/<name>/`                                        | Progressive discovery with activation consent; no enforceable per-skill switch is established                                    | Named sub-agent tool or explicit `@name` request                       | Additional target because subagents and evidence vary by installed version                                       |
| Grok Build          | Project `.grok/skills/<name>/`; user `~/.grok/skills/<name>/`                                            | Relevance behavior is documented, but the loose-skill schema and collision rules are incomplete                                  | Task or spawn tool through the active agent                            | Stop description evaluation until `grok inspect --json`, isolation, and native invocation evidence all pass live |
| Kilo Code local     | Project `.kilo/skills/<name>/`; user `~/.kilo/skills/<name>/`                                            | Relevance or explicit-name activation; no enforceable per-skill switch is established                                            | `task` tool or explicit `@agent-name` request                          | Additional target until invocation evidence is verified; Kilo Cloud remains unsupported                          |

## Unsupported and Missing Capabilities

- Stop before trials when a target is explicit-only under the current surface.
- Stop when an independent agent cannot be created or isolated.
- Stop when only output similarity, self-report, or hidden reasoning could indicate invocation.
- Stop when native discovery has not been proven for the exact installed path and surface.
- Prepare unleading prompts and fixtures for external execution when the user can run them in a qualified environment; accept only the native record required by the evidence standard.

## Authoritative References

- Product support and matrix: `../../../plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`.
- Direct-install paths, activation, verification, trust, and caveats: `../../../plans/skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md`.
- Source-backed harness details: `../../../plans/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md` and its `assignments/` directory.
- Sub-agent launch behavior: `../../../plans/research/harness-subagent-launch/harness-subagent-launch-reference.md` and its `assignments/` directory.
- Environment-specific measured reliability belongs in the evaluation campaign record, not this reference.

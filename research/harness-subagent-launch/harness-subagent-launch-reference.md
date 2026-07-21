# Cross-Harness Sub-Agent Launch Reference

## GitHub Copilot

- **Supported:** Yes.
- **Launch:** The main agent delegates automatically or from a user request. VS Code exposes `agent/runSubagent`; Copilot CLI also provides `/fleet`.
- **Skill request:** Yes. A skill can request delegation, and supported VS Code skills can use experimental `context: fork`.
- **Research:** [assignment](assignments/01-github-copilot.md)

## Claude Code

- **Supported:** Yes.
- **Launch:** Claude uses the `Agent` tool automatically or after an explicit user request.
- **Skill request:** Yes. A skill can request delegation or use `context: fork` with an agent selection.
- **Research:** [assignment](assignments/02-claude-code.md)

## OpenAI Codex

- **Supported:** Yes.
- **Launch:** The user or active instructions ask Codex to spawn or delegate to sub-agents; Codex performs the orchestration.
- **Skill request:** Yes. Skill instructions can request delegation.
- **Research:** [assignment](assignments/03-openai-codex.md)

## Cursor

- **Supported:** Yes.
- **Launch:** Cursor Agent delegates automatically or from an explicit named or natural-language request.
- **Skill request:** Yes. Skill instructions can request that the parent Agent launch a sub-agent.
- **Research:** [assignment](assignments/04-cursor.md)

## Google Antigravity or Gemini CLI

- **Supported:** Yes.
- **Launch:** Antigravity uses `invoke_subagent`. Gemini CLI exposes available sub-agents as tools and supports explicit `@name` requests.
- **Skill request:** Yes. Skill instructions can request that the parent agent use the relevant launch tool.
- **Research:** [assignment](assignments/05-google-antigravity-gemini-cli.md)

## Grok Build

- **Supported:** Yes.
- **Launch:** The active agent uses the sub-agent task or spawn tool.
- **Skill request:** Yes. Bundled Grok skills demonstrate instructions that request the task tool.
- **Research:** [assignment](assignments/06-grok-build.md)

## OpenCode

- **Supported:** Yes.
- **Launch:** A primary agent uses the `Task` tool, or the user explicitly mentions a sub-agent with `@name`.
- **Skill request:** Yes. A skill can request delegation; the active agent performs the launch.
- **Research:** [assignment](assignments/07-opencode.md)

## Pi

- **Supported:** Yes, through Pi's official sub-agent extension.
- **Launch:** The active agent calls the extension-provided `subagent` tool.
- **Skill request:** Yes. A skill can request that the active agent call the tool.
- **Research:** [assignment](assignments/09-pi.md)

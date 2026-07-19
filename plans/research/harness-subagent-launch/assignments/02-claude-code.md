# Claude Code Sub-Agent Launch

1. **Does Claude Code support sub-agents?** Yes. Claude Code includes built-in subagents and supports custom subagents, each running in its own context window and returning its result to the parent conversation. [Claude Code: Create custom subagents](https://code.claude.com/docs/en/sub-agents)

2. **If yes, how is a sub-agent launched?** Claude Code launches one through its built-in `Agent` tool. Claude may delegate automatically when a task matches a subagent's description, or the user can explicitly ask Claude to use a named agent, such as `Use the code-improver agent to suggest improvements in this project`. [Claude Code: Create custom subagents](https://code.claude.com/docs/en/sub-agents) [Claude Code: Tools reference](https://code.claude.com/docs/en/tools-reference#agent-tool-behavior)

3. **Can a skill launch or request a sub-agent?** Yes. A skill can set `context: fork` in its `SKILL.md` frontmatter, optionally set `agent: Explore`, `Plan`, `general-purpose`, or a custom subagent, and run its skill content as the subagent's task. [Claude Code: Run skills in a subagent](https://code.claude.com/docs/en/slash-commands#run-skills-in-a-subagent)

# Kilo Code Sub-Agent Launch

1. **Does Kilo Code support sub-agents?** Yes. In the current VS Code extension and CLI, full-tool agents such as Code, Plan, and Debug can delegate to isolated subagent sessions; Kilo also includes the built-in `general` and `explore` subagents. [Kilo Code: Orchestrator Mode (Deprecated)](https://kilo.ai/docs/code-with-ai/agents/orchestrator-mode) [Kilo Code: Custom Subagents](https://kilo.ai/docs/customize/custom-subagents)

2. **If yes, how is a sub-agent launched?** A primary agent launches one with the `task` tool, either automatically when delegation is useful or because the user explicitly invokes a configured subagent with `@agent-name`. [Kilo Code: Custom Subagents](https://kilo.ai/docs/customize/custom-subagents) [Kilo Code: Tool Use Overview](https://kilo.ai/docs/automate/tools)

3. **Can a skill launch or request a sub-agent?** Yes, through the active primary agent. A skill is a `SKILL.md` instruction module that the agent loads and follows, so its instructions can direct that agent to call the `task` tool; the launch remains subject to the active agent's delegation capability and `permission.task` rules. [Kilo Code: Skills](https://kilo.ai/docs/customize/skills) [Kilo Code: Agent Permissions](https://kilo.ai/docs/customize/agent-permissions)

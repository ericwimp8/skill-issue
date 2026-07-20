# OpenCode sub-agent launch

1. **Does OpenCode support sub-agents?** Yes. OpenCode distinguishes primary agents from subagents and includes the built-in `general`, `explore`, and `scout` subagents. [OpenCode Agents documentation](https://opencode.ai/docs/agents)

2. **How is a sub-agent launched?** A primary agent can invoke a subagent automatically through the `Task` tool, subject to `permission.task`; a user can launch one manually by mentioning it, for example `@general help me search for this function`. [OpenCode Agents: usage and task permissions](https://opencode.ai/docs/agents#usage)

3. **Can a skill launch or request a sub-agent?** A skill can request delegation by instructing the active agent to use a subagent, but it does not directly launch one itself. OpenCode loads a skill as reusable instructions through the `skill` tool; the active agent performs the actual launch through the separate `Task` tool, if its task permissions allow that target. [OpenCode Agent Skills](https://opencode.ai/docs/skills) · [OpenCode Agents: task permissions](https://opencode.ai/docs/agents#task-permissions)

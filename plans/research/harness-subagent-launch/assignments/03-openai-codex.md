# OpenAI Codex Sub-Agent Launch

1. **Does OpenAI Codex support sub-agents?** Yes. Current Codex releases support subagent workflows in the ChatGPT desktop app, Codex CLI, and IDE extension. [OpenAI Codex: Subagents](https://learn.chatgpt.com/docs/agent-configuration/subagents)

2. **If yes, how is a sub-agent launched?** Ask Codex directly with an instruction such as `spawn two agents` or `delegate this work in parallel`. Codex handles the orchestration and spawns the delegated agent threads. [OpenAI Codex: Triggering subagent workflows](https://learn.chatgpt.com/docs/agent-configuration/subagents#triggering-subagent-workflows)

3. **Can a skill launch or request a sub-agent?** Yes. Applicable skill instructions can request delegation; Codex then launches the subagent workflow. [OpenAI Codex: Subagents availability](https://learn.chatgpt.com/docs/agent-configuration/subagents#availability)

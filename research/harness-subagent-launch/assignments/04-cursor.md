# Cursor sub-agent launch

1. **Does Cursor support sub-agents?** Yes. Cursor supports specialized subagents with separate context windows in the editor, CLI, and Cloud Agents. The parent Agent can delegate work to them, and subagents can run in the foreground or background. [Cursor Subagents documentation](https://cursor.com/docs/subagents)

2. **How is a sub-agent launched?** Cursor Agent can launch one automatically when a task warrants delegation. A user or instruction can explicitly request a custom subagent with `/name` in the prompt (for example, `/verifier confirm the auth flow is complete`) or by naming it naturally (for example, “Use the verifier subagent…”). [Automatic delegation](https://cursor.com/docs/subagents#automatic-delegation) · [Explicit invocation](https://cursor.com/docs/subagents#explicit-invocation)

3. **Can a skill launch or request a sub-agent?** Yes, as instructions executed by the parent Agent: a `SKILL.md` workflow can tell Agent to invoke a named subagent, and Agent performs the launch through Cursor's normal delegation mechanism. Cursor documents skills as workflows made available to Agent and subagents as assistants Agent launches; it does not document a separate skill-only subagent launch API. [How skills work](https://cursor.com/docs/skills#how-skills-work) · [Using subagents](https://cursor.com/docs/subagents#using-subagents)

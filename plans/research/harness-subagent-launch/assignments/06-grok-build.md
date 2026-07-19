# Grok Build: Sub-agent Launch

1. **Does Grok Build support sub-agents?** Yes. Sub-agents are enabled by default and run as independent child sessions, including parallel execution.

2. **How is a sub-agent launched?** The main agent invokes the sub-agent tool with a task `prompt`, short `description`, and optional `subagent_type` (`general-purpose`, `explore`, or `plan`), capability, isolation/worktree, background, resume, and working-directory settings. The user guide names this tool `spawn_subagent`; the current harness source exposes the underlying tool as `task`.

3. **Can a skill launch or request a sub-agent?** Yes. A skill can instruct the active agent to invoke the sub-agent tool. xAI's bundled `check-work` skill does exactly this by directing the agent to call `task` with a verifier prompt and `general-purpose` sub-agent type.

## First-party sources

- [Grok Build sub-agent user guide](https://github.com/xai-org/grok-build/blob/7cfcb20d2b50b0d18801a6c0af2e401c0e060894/crates/codegen/xai-grok-pager/docs/user-guide/16-subagents.md#L1-L5)
- [Sub-agent launch parameters](https://github.com/xai-org/grok-build/blob/7cfcb20d2b50b0d18801a6c0af2e401c0e060894/crates/codegen/xai-grok-pager/docs/user-guide/16-subagents.md#L141-L156)
- [Current `task` tool input in the harness](https://github.com/xai-org/grok-build/blob/7cfcb20d2b50b0d18801a6c0af2e401c0e060894/crates/common/xai-tool-types/src/task.rs#L1-L43)
- [Bundled `check-work` skill requesting a sub-agent](https://github.com/xai-org/grok-build/blob/7cfcb20d2b50b0d18801a6c0af2e401c0e060894/crates/codegen/xai-grok-shell/skills/check-work/SKILL.md#L33-L40)

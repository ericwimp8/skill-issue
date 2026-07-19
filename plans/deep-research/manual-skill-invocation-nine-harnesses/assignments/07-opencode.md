# OpenCode Manual Skill Invocation

## Assignment

- **Goal:** Determine whether an OpenCode user can manually invoke the installed `dictate-plan` agent skill, the exact action or syntax, and the strongest supported explicit wording when no dedicated syntax exists.
- **Scope:** Current official OpenCode documentation and first-party `anomalyco/opencode` repository sources concerning native agent-skill discovery, loading, and direct user invocation.
- **Exclusions:** Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated OpenCode behavior.

## Sources

- OpenCode, [Agent Skills](https://opencode.ai/docs/skills), current documentation inspected 2026-07-19; especially the on-demand native `skill` tool model and the documented `skill({ name: "git-release" })` agent call.
- `anomalyco/opencode`, [`packages/opencode/src/session/system.ts`](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/session/system.ts), `dev` branch inspected 2026-07-19; system prompt tells the agent to use the skill tool when a task matches a listed skill description.
- `anomalyco/opencode`, [feature request #7846: Add `/skills` command to list and quick-invoke skills](https://github.com/anomalyco/opencode/issues/7846), open when inspected 2026-07-19; records the absence of direct TUI slash invocation and proposes `/skill:<name>` as a future feature.

## Findings

### Finding 1: A user can explicitly request a specific skill, but loading is agent-mediated

OpenCode exposes installed skills to the agent and loads their full instructions on demand through the native `skill` tool. Therefore a user can manually identify `dictate-plan` in an ordinary prompt and ask the agent to use it. The user request is explicit, while the concrete load operation is performed by the agent rather than by a dedicated user-side skill command.

**Evidence:** The official [Agent Skills documentation](https://opencode.ai/docs/skills) says agents see available skills and load them on demand. It shows the agent-side call `skill({ name: "git-release" })`. OpenCode's first-party [`system.ts`](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/session/system.ts) instructs the agent: use the skill tool when a task matches the skill description.

**Implication:** For question (1), the best-supported answer is **yes, by an explicit natural-language request to the agent**. First-party evidence does not establish a user-executed command that directly forces the load.

### Finding 2: There is no documented dedicated invocation syntax for `dictate-plan`

No current official Agent Skills documentation gives users a slash command, sigil, or other dedicated syntax for invoking a named skill. The syntax `skill({ name: "dictate-plan" })` is the agent's internal tool call, not text documented for a user to enter in the OpenCode prompt.

**Evidence:** The official [Agent Skills documentation](https://opencode.ai/docs/skills) labels `skill({ name: "git-release" })` as the call made by the agent. In the first-party repository, open [issue #7846](https://github.com/anomalyco/opencode/issues/7846) states that users currently cannot invoke a specific skill directly through a slash command and proposes `/skill:<name>` as a future quick-invocation syntax.

**Implication:** For question (2), there is **no dedicated documented user syntax** such as `/skill:dictate-plan`. Users should not treat `skill({ name: "dictate-plan" })` as a supported prompt command.

### Finding 3: The strongest explicit user wording is a direct natural-language instruction

Use an ordinary prompt that names the skill and directs the agent to load and follow it. The clearest wording is:

> Use the `dictate-plan` skill to develop a dependency-ordered plan with me from this point onward.

If the user wants to emphasize loading before any other work, the best-supported variant is:

> Load and use the `dictate-plan` skill, then help me develop a dependency-ordered plan from my dictation.

**Evidence:** OpenCode advertises skill names and descriptions to the agent and instructs it to load a matching skill through the native tool, as documented in [Agent Skills](https://opencode.ai/docs/skills) and implemented in first-party [`system.ts`](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/session/system.ts). First-party sources do not prescribe exact user-facing wording, so these prompts are an inference from the documented agent-mediated loading design.

**Implication:** For question (3), `Use the dictate-plan skill to ...` is the most explicit, best-supported wording. It clearly identifies the installed skill, while preserving that OpenCode leaves the actual tool call to the agent.

## Notes

- First-party documentation does not promise that natural-language wording mechanically forces the model to call the skill tool; it is the clearest supported request under OpenCode's agent-mediated design.
- `/skill:dictate-plan` appears only as proposed syntax in an open first-party repository issue and should be treated as unsupported unless future official documentation says otherwise.

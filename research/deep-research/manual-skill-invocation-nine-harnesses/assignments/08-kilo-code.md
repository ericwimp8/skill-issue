# Kilo Code Manual Skill Invocation

## Assignment

- **Goal:** Determine whether a user can manually invoke an installed Kilo Code agent skill named `dictate-plan`, identify any exact dedicated syntax, and provide the best-supported explicit wording when no dedicated syntax is documented.
- **Scope:** Current first-party Kilo Code documentation covering skills, natural-language task entry, workflow tools, and interactive slash commands.
- **Exclusions:** Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated behavior.

## Sources

- Kilo Code, [Skills](https://kilo.ai/docs/customize/skills), inspected 2026-07-19. Relevant sections: “How the Agent Decides to Use a Skill” and “Checking if a Skill Was Used.”
- Kilo Code, [Your First Task](https://kilo.ai/docs/getting-started/quickstart), inspected 2026-07-19. Relevant section: “Type Your Task.”
- Kilo Code, [How Tools Work](https://kilo.ai/docs/automate/how-tools-work), inspected 2026-07-19. Relevant sections: “Tool Workflow” and “Core Tools Reference.”
- Kilo Code, [Kilo Code CLI](https://kilo.ai/docs/code-with-ai/platforms/cli), inspected 2026-07-19. Relevant section: “Interactive Slash Commands.”

## Findings

### A user can manually request a specific installed skill by name

Yes. Kilo Code’s Skills documentation says explicit invocation works when the user says to use a named skill, giving “use the api-design skill” as its example. Applying the documented pattern to the requested skill, the user can type a natural-language chat request naming `dictate-plan` directly. The agent then decides to call its internal `skill` tool with that name; this tool call is how the conversation can show that the skill was loaded. [Kilo Code Skills](https://kilo.ai/docs/customize/skills)

**Evidence:** The first-party Skills page explicitly describes name-based user wording as invocation and separately explains that the agent invokes the `skill` tool when it uses a skill. The first-party quickstart says tasks are entered in plain English. [Kilo Code quickstart](https://kilo.ai/docs/getting-started/quickstart)

**Implication:** Manual invocation is supported through an explicit natural-language instruction that contains the installed skill’s exact name.

### No dedicated user-facing skill invocation syntax is documented

No dedicated `/skill`, `$skill-name`, or comparable user syntax is documented in the current first-party sources inspected. The CLI documentation’s enumerated interactive slash-command reference includes commands such as `/agents`, `/models`, and `/reload`, but no `/skill` command. Kilo’s tool documentation describes `skill` as a workflow tool selected by Kilo after the user describes the desired action in natural language, rather than as a command the user types. [Kilo Code CLI slash commands](https://kilo.ai/docs/code-with-ai/platforms/cli) [Kilo Code tool workflow](https://kilo.ai/docs/automate/how-tools-work)

**Evidence:** The current interactive slash-command table has no skill-invocation entry, while the Skills page documents explicit invocation with ordinary words: “use the api-design skill.” [Kilo Code Skills](https://kilo.ai/docs/customize/skills)

**Implication:** For `dictate-plan`, there is no first-party basis for recommending `/skill dictate-plan`, `/dictate-plan`, `$dictate-plan`, or another special token. This is an evidence-limited documentation conclusion rather than proof that no undocumented interface exists.

### Best-supported exact wording for `dictate-plan`

Use this explicit user message:

> Use the `dictate-plan` skill.

For a concrete task, preserve the same trigger and append the request, for example: “Use the `dictate-plan` skill to develop a dependency-ordered plan from my next messages.” The first sentence is the closest exact adaptation of Kilo Code’s documented “use the api-design skill” example. [Kilo Code Skills](https://kilo.ai/docs/customize/skills)

**Evidence:** Kilo Code states that the agent sees available skill names and that explicitly saying “use the [name] skill” triggers the named skill. Its general task guidance directs users to type clear plain-English requests without special syntax. [Kilo Code quickstart](https://kilo.ai/docs/getting-started/quickstart)

**Implication:** The recommended manual action is to send `Use the dictate-plan skill.` in Kilo Code chat, optionally followed by the planning task.

## Notes

- First-party documentation is sufficient to support explicit name-based natural-language invocation, but it does not specify a dedicated user command for arbitrary skills.
- The internal `skill` tool is evidence of what the agent calls after interpreting the request; it is not documented as literal syntax for the user to enter.

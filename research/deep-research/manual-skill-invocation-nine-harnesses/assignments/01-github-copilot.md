# GitHub Copilot Manual Skill Invocation

## Assignment

**Goal:** Determine whether a user can manually invoke an installed GitHub Copilot agent skill named `dictate-plan`, identify the exact supported invocation, and provide the strongest first-party explicit wording where a dedicated syntax is unavailable.

**Scope:** Current first-party GitHub documentation and GitHub-owned public repository guidance about manually invoking an already installed agent skill. The dedicated syntax finding is limited to the GitHub Copilot surface for which GitHub explicitly documents it.

**Exclusions:** Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated Copilot behavior.

## Sources

- GitHub Docs, [GitHub Copilot CLI command reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference), "Skills reference," current page accessed 2026-07-19.
- GitHub Docs, [Scheduling prompts in GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/automate-copilot-cli/schedule-prompts), "Scheduling a skill," current page accessed 2026-07-19.
- GitHub Docs, [About agent skills](https://docs.github.com/en/copilot/concepts/agents/about-agent-skills), current page accessed 2026-07-19.
- GitHub-owned `github/awesome-copilot` repository, [Agent Skills usage guidance](https://github.com/github/awesome-copilot/blob/main/docs/README.skills.md), `main` branch as accessed 2026-07-19.

## Findings

### Manual Invocation Is Supported in GitHub Copilot CLI

Yes. A user can manually invoke a specific installed skill in GitHub Copilot CLI when that skill is user-invocable. GitHub documents that skills may be invoked through `/SKILL-NAME`, and that the `user-invocable` field defaults to `true`. Therefore an installed `dictate-plan` skill is manually invocable by default unless its metadata explicitly disables user invocation.

**Evidence:** The [GitHub Copilot CLI skills reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference#skills-reference) states that a skill is invoked "via `/SKILL-NAME` or automatically by the agent" and defines `user-invocable` as controlling whether users can invoke a skill with `/SKILL-NAME`, with a default of `true`.

**Implication:** GitHub Copilot CLI provides a dedicated manual invocation path for an enabled, user-invocable `dictate-plan` skill; natural-language intent matching is optional rather than required on that surface.

### Exact Dedicated Syntax for `dictate-plan`

In an interactive GitHub Copilot CLI session, enter and submit:

```text
/dictate-plan
```

GitHub also documents skill slash commands followed by task text, so an initial request may be supplied in the same prompt when useful:

```text
/dictate-plan Help me develop a dependency-ordered plan for this project.
```

**Evidence:** GitHub's [CLI skills reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference#skills-reference) defines the invocation form as `/SKILL-NAME`. Its [skill scheduling examples](https://docs.github.com/en/copilot/how-tos/copilot-cli/automate-copilot-cli/schedule-prompts#scheduling-a-skill) show a skill slash command followed by task text, demonstrating that arguments or a request can follow the command.

**Implication:** Substituting the supplied skill name into GitHub's documented placeholder yields `/dictate-plan`; this is dedicated invocation syntax, not merely natural-language prompting.

### Explicit Wording Where Dedicated Syntax Is Undocumented

GitHub documents that agent skills also work with Copilot cloud agent, Copilot code review, the GitHub Copilot app, and agent mode in Visual Studio Code and JetBrains IDEs, but the reviewed GitHub documentation does not establish `/SKILL-NAME` as a dedicated invocation syntax for those non-CLI surfaces. The best-supported explicit wording there is:

```text
Use the dictate-plan skill to create and maintain a dependency-ordered plan from my dictation.
```

**Evidence:** GitHub's [overview of agent skills](https://docs.github.com/en/copilot/concepts/agents/about-agent-skills) lists the supported Copilot surfaces but describes skills generally as loading when relevant. The GitHub-owned `awesome-copilot` [usage guidance](https://github.com/github/awesome-copilot/blob/main/docs/README.skills.md#how-to-use-agent-skills) tells users to "Reference skills in your prompts or let the agent discover them automatically." GitHub's [CLI scheduling documentation](https://docs.github.com/en/copilot/how-tos/copilot-cli/automate-copilot-cli/schedule-prompts#scheduling-a-skill) supplies the first-party natural-language pattern "Use the [name] skill to [task]."

**Implication:** Outside GitHub Copilot CLI, explicitly name `dictate-plan` and state the requested planning action in ordinary language. This wording is best-supported rather than formally canonical; first-party GitHub evidence reviewed here does not define an exact natural-language sentence or confirm a dedicated slash command on every Copilot surface.

## Notes

- Dedicated `/dictate-plan` invocation is confirmed for GitHub Copilot CLI only. Cross-surface slash-command parity remains unsupported by the reviewed first-party GitHub sources.
- Manual slash invocation depends on the skill remaining user-invocable; GitHub documents `true` as the default, but a skill can explicitly set `user-invocable: false`.

# Manual Invocation of `dictate-plan` Across Nine Harnesses

## Combined Reference

### GitHub Copilot

**Answer: Yes, conditionally in GitHub Copilot CLI.** If the installed skill is user-invocable—GitHub documents `true` as the default—enter:

```text
/dictate-plan
```

Task text may follow the command, for example `/dictate-plan Help me develop a dependency-ordered plan for this project.` The reviewed first-party evidence establishes this dedicated syntax for **GitHub Copilot CLI**, not every Copilot surface. Elsewhere, use: `Use the dictate-plan skill to create and maintain a dependency-ordered plan from my dictation.` That prose is explicit wording, not guaranteed command syntax. ([assignment](assignments/01-github-copilot.md), [GitHub CLI skills reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference#skills-reference), [GitHub agent-skills guidance](https://github.com/github/awesome-copilot/blob/main/docs/README.skills.md#how-to-use-agent-skills))

### Claude Code

**Answer: Yes, if the skill remains user-invocable.** Enter the command at the start of the message:

```text
/dictate-plan
```

Initial dictation may follow as arguments: `/dictate-plan <your dictation or task>`. Anthropic documents `/skill-name`, derives the command name from the skill directory name, and enables user invocation by default unless `user-invocable: false` is set. ([assignment](assignments/02-claude-code.md), [Anthropic skills documentation](https://code.claude.com/docs/en/slash-commands#control-who-invokes-a-skill), [Anthropic commands reference](https://code.claude.com/docs/en/commands))

### OpenAI Codex

**Answer: Yes.** Mention the skill directly in the prompt:

```text
$dictate-plan Turn these notes into a dependency-ordered plan: ...
```

In Codex CLI, the documented alternative is to enter `/skills`, select `dictate-plan`, and then submit the request. `/skills` is the picker command; `/dictate-plan` is not documented as Codex skill-invocation syntax. ([assignment](assignments/03-openai-codex.md), [OpenAI Build skills](https://learn.chatgpt.com/docs/build-skills#how-codex-uses-skills), [Codex CLI developer commands](https://learn.chatgpt.com/docs/developer-commands?surface=cli))

### Cursor

**Answer: Yes, if Cursor has discovered the installed skill.** In Agent chat, type `/`, search for `dictate-plan`, and select it. Cursor also documents the generic named form, yielding:

```text
/dictate-plan
```

The direct name is derived from Cursor's documented `/skill-name` form rather than a Cursor-published `dictate-plan` example. ([assignment](assignments/04-cursor.md), [Cursor Agent Skills](https://cursor.com/docs/skills), [Cursor 2.4 changelog](https://cursor.com/changelog/2-4))

### Google Antigravity or Gemini CLI

**Answer: Conditional, because the two current surfaces differ.** In **Antigravity CLI**, a skill registered in the CLI slash-command format becomes a same-name command; enter:

```text
/dictate-plan
```

First-party evidence does not establish that every directory-style `SKILL.md` Agent Skill automatically receives that command. For those Antigravity skills, and for **Gemini CLI**, use: `Use the dictate-plan skill for this request: [describe the plan you want to develop].` Gemini CLI documents semantic activation and management-only `/skills` subcommands, not `/dictate-plan` or `/skills invoke dictate-plan`. ([assignment](assignments/05-google-antigravity-gemini-cli.md), [Antigravity CLI skills](https://antigravity.google/docs/cli-plugins), [Antigravity Agent Skills](https://antigravity.google/docs/skills), [Gemini CLI Agent Skills](https://geminicli.com/docs/cli/skills/))

### Grok Build

**Answer: Yes, if Grok Build discovers the skill and it is user-invocable.** In the TUI, enter:

```text
/dictate-plan
```

If the product reports a name collision, use the qualified command it displays; first-party documentation does not define one universal qualifier. ([assignment](assignments/06-grok-build.md), [xAI skills documentation](https://docs.x.ai/build/features/skills-plugins-marketplaces), [xAI modes and commands](https://docs.x.ai/build/modes-and-commands))

### OpenCode

**Answer: Yes, through an explicit user request; loading remains agent-mediated.** No dedicated user-side skill command is documented. Use:

```text
Load and use the `dictate-plan` skill, then help me develop a dependency-ordered plan from my dictation.
```

OpenCode's `skill({ name: "dictate-plan" })` form is an internal agent tool call, not prompt syntax. `/skill:dictate-plan` is only a proposed feature in the cited first-party issue. ([assignment](assignments/07-opencode.md), [OpenCode Agent Skills](https://opencode.ai/docs/skills), [OpenCode issue #7846](https://github.com/anomalyco/opencode/issues/7846))

### Kilo Code

**Answer: Yes, through documented natural-language invocation.** No dedicated user-facing skill syntax is documented. Send:

```text
Use the `dictate-plan` skill.
```

For an immediate task: `Use the dictate-plan skill to develop a dependency-ordered plan from my next messages.` This directly adapts Kilo Code's documented “use the [name] skill” pattern. ([assignment](assignments/08-kilo-code.md), [Kilo Code Skills](https://kilo.ai/docs/customize/skills), [Kilo Code CLI commands](https://kilo.ai/docs/code-with-ai/platforms/cli))

### Pi

**Answer: Yes, when skill commands are enabled; Pi documents them as enabled by default.** At the interactive prompt, enter:

```text
/skill:dictate-plan
```

Initial task text may follow: `/skill:dictate-plan <your dictation or task>`. If skill commands are locally disabled, use `Use the dictate-plan skill for this request: <your dictation or task>`; Pi documents prompting as an alternative but no canonical prose command. ([assignment](assignments/09-pi.md), [Pi Skills](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/skills.md#skill-commands), [Pi settings](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/settings.md#resources))

## Best-Supported Overall Direction

Use each harness's dedicated mechanism where first-party evidence establishes one: `/dictate-plan` in GitHub Copilot CLI, Claude Code, Cursor, eligible Antigravity CLI skill registrations, and Grok Build; `$dictate-plan` or the `/skills` picker in OpenAI Codex; and `/skill:dictate-plan` in Pi. For OpenCode, Kilo Code, Gemini CLI, and Antigravity directory-style skills without confirmed slash exposure, explicitly name the skill in ordinary language: `Use the dictate-plan skill to ...`.

This harness-specific approach is strongest because command grammars are not portable. It also separates deterministic or UI-mediated skill selection from prose that asks an agent to load a named skill but may still depend on model behavior.

## Conditional Alternatives

- Use the harness's skill picker or slash menu where documented: Codex `/skills`, Cursor `/` search, and Antigravity `/` typeahead reduce command-name typing errors.
- Append the initial task only where the supplied evidence supports trailing text: GitHub Copilot CLI, Claude Code, OpenAI Codex prompts, and Pi.
- Fall back to `Use the dictate-plan skill to ...` when a dedicated command is unavailable, locally disabled, undiscoverable, or unsupported on that product surface; treat the result as agent-mediated unless first-party documentation says otherwise.

## Rejected or Lower-Fit Interpretations

- **Universal `/dictate-plan`: rejected.** Codex documents `$dictate-plan`; Pi documents `/skill:dictate-plan`; OpenCode, Kilo Code, and Gemini CLI do not document `/dictate-plan` for named-skill invocation.
- **Universal `/skills` invocation: rejected.** Depending on the harness, `/skills` may be a picker, browser, or management surface rather than the command that invokes a selected skill.
- **Internal tool calls as user syntax: rejected.** OpenCode's `skill({ name: ... })` and Kilo Code's internal `skill` tool describe agent behavior, not text users are documented to enter.
- **Prose as equivalent to a dedicated command: lower fit.** Natural-language naming is the documented primary method in Kilo Code and the best-supported method in OpenCode and Gemini CLI, but it is not uniformly guaranteed to force deterministic loading.

## Unresolved Blockers and Evidence Limits

- No assignment verifies the caller's local installation, discovery state, declared name, or invocation-control metadata. Commands that depend on those states remain conditional.
- Dedicated syntax is surface-specific for GitHub Copilot and Google: Copilot's `/dictate-plan` evidence is CLI-specific, while Antigravity slash exposure is confirmed for its CLI-registered skill format but not every directory-style Agent Skill.
- OpenCode and Gemini CLI lack first-party support for a dedicated user invocation command in the supplied research. Kilo Code instead documents explicit natural-language invocation.
- These limits do not block the reference; they define where the recommended wording is best-supported rather than mechanically guaranteed.

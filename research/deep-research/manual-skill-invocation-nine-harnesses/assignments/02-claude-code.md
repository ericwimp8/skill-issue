# Claude Code Manual Skill Invocation

## Assignment

**Goal:** Determine whether a Claude Code user can manually invoke an installed skill named `dictate-plan`, identify the exact invocation syntax, and provide the best-supported explicit wording only if dedicated syntax is unavailable.

**Scope:** Current first-party Anthropic documentation for manual invocation of user-authored Claude Code skills.

**Exclusions:** Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated Claude Code behavior.

## Sources

- Anthropic, [Extend Claude with skills](https://code.claude.com/docs/en/slash-commands), inspected 2026-07-19. Relevant sections: direct invocation, command-name derivation, invocation control, and skill arguments.
- Anthropic, [Commands](https://code.claude.com/docs/en/commands), inspected 2026-07-19. Relevant section: command recognition and argument placement.
- Anthropic, [Interactive mode](https://code.claude.com/docs/en/interactive-mode), inspected 2026-07-19. Relevant section: browsing invocable skills in the `/` menu.

## Findings

### Finding 1: A user can manually invoke a specific installed skill

Yes. Claude Code supports direct user invocation of a specific user-invocable skill. Anthropic documents that users can type `/skill-name` to invoke a skill directly; user invocation is enabled by default unless the skill is explicitly marked `user-invocable: false`.

**Evidence:** Anthropic states that skills can be invoked directly with `/skill-name` and that the default invocation configuration allows both the user and Claude to invoke a skill in [Extend Claude with skills](https://code.claude.com/docs/en/slash-commands#control-who-invokes-a-skill). The interactive reference also says the `/` menu includes user-authored skills available for invocation in [Interactive mode](https://code.claude.com/docs/en/interactive-mode#commands).

**Implication:** Assuming the installed `dictate-plan` skill remains user-invocable, a user can select it from the `/` menu or invoke it by name.

### Finding 2: The dedicated syntax is `/dictate-plan`

Type the following at the start of a Claude Code message and send it:

```text
/dictate-plan
```

To provide the initial request in the same message, use:

```text
/dictate-plan <your dictation or task>
```

**Evidence:** Anthropic documents that the directory name of a personal or project skill becomes the slash-command name, with `.claude/skills/deploy-staging/SKILL.md` mapping to `/deploy-staging`, in [How a skill gets its command name](https://code.claude.com/docs/en/slash-commands#how-a-skill-gets-its-command-name). The [Commands reference](https://code.claude.com/docs/en/commands) says a command is recognized at the start of the message and trailing text becomes its arguments. Anthropic also documents argument-bearing skill invocations such as `/fix-issue 123` in [Pass arguments to skills](https://code.claude.com/docs/en/slash-commands#pass-arguments-to-skills).

**Implication:** For an installed, user-invocable skill whose command name is `dictate-plan`, `/dictate-plan` is dedicated invocation syntax rather than natural-language prompting. Anthropic does not publish a page specifically naming this third-party skill, so the exact command is derived from the caller-supplied skill name plus Anthropic's documented naming rule.

### Finding 3: Natural-language wording is only an optional, non-guaranteed alternative

A dedicated syntax exists, so no natural-language fallback is required. If a user deliberately prefers ordinary wording, the best-supported explicit form is:

```text
Use the dictate-plan skill for this request: <your dictation or task>
```

This wording asks Claude to choose the named skill; it is not documented as a guaranteed direct-invocation mechanism.

**Evidence:** Anthropic says Claude may load a skill automatically when a request matches its description, while direct invocation uses `/skill-name`, in [Extend Claude with skills](https://code.claude.com/docs/en/slash-commands#getting-started). Anthropic provides examples of natural-language requests matching a skill description, but it does not document a reserved prose phrase equivalent to the slash command.

**Implication:** Use `/dictate-plan` when deterministic manual selection is required. Treat explicit natural-language wording as model-mediated skill selection with no first-party guarantee that the named skill will load.

## Notes

- The first-party evidence establishes Claude Code's general invocation contract. It does not independently verify the existence, location, frontmatter, or runtime visibility of a particular `dictate-plan` installation.
- Relevant search terms: `Claude Code skills direct invocation`, `Claude Code /skill-name`, `Claude Code user-invocable`, `Claude Code skill arguments`.

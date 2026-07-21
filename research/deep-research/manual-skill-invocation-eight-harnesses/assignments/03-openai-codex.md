# OpenAI Codex Manual Skill Invocation

## Assignment

**Goal:** Determine whether a user can manually invoke an installed Codex skill named `dictate-plan`, identify the exact action or syntax, and provide the best-supported explicit wording only if no dedicated syntax exists.

**Scope:** Current first-party OpenAI documentation for Codex skill invocation in supported Codex surfaces, with emphasis on dedicated invocation syntax and the Codex CLI skill picker.

**Exclusions:** Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated behavior.

## Sources

- [OpenAI Codex manual](https://developers.openai.com/codex/codex-manual.md), fetched through OpenAI's official Codex-manual helper on 2026-07-19; relevant embedded source sections were “Build skills,” “Slash commands in Codex CLI,” and “Skills & Plugins.”
- [Build skills](https://learn.chatgpt.com/docs/build-skills), especially “How Codex uses skills,” which documents explicit invocation through `/skills` or a `$` skill mention.
- [Developer commands: Codex CLI](https://learn.chatgpt.com/docs/developer-commands?surface=cli), especially “Use skills with `/skills`,” which documents selecting a skill for the next request.
- [Skills & Plugins](https://learn.chatgpt.com/docs/skills-and-plugins), which states that Codex supports `$` mentions for skills.

## Findings

### Finding 1: A user can manually invoke a specific installed skill

Yes. OpenAI documents “explicit invocation” as one of the two ways Codex activates skills. The documented mechanisms are to include the skill directly in the prompt, type `$` to mention a skill in CLI/IDE, or use `/skills` and select the desired skill. This is a user-controlled invocation rather than relying on Codex to infer the skill from its description.

**Evidence:** The official [Build skills documentation](https://learn.chatgpt.com/docs/build-skills#how-codex-uses-skills) states that explicit invocation includes the skill directly in the prompt and, in CLI/IDE, uses `/skills` or `$` to mention a skill. The first-party [Skills & Plugins documentation](https://learn.chatgpt.com/docs/skills-and-plugins#use-skills-for-repeatable-work) independently states that Codex supports `$` mentions for skills.

**Implication:** An installed `dictate-plan` skill can be selected deliberately; the user does not need to depend on implicit description matching.

### Finding 2: The dedicated syntax for `dictate-plan` is `$dictate-plan`

In a Codex prompt, type `$dictate-plan`, select the matching skill mention if the composer presents a picker, and include the task or source material in the same prompt. For example: `$dictate-plan Turn these notes into a dependency-ordered plan: ...`.

The documented CLI alternative is: type `/skills`, pick `dictate-plan`, then enter the request that should follow the skill's instructions. `/skills` is the picker command; `/dictate-plan` is not documented as a skill invocation command.

**Evidence:** OpenAI's [Build skills documentation](https://learn.chatgpt.com/docs/build-skills#how-codex-uses-skills) defines `$` mention syntax for explicit skill invocation and gives `$skill-creator` as a concrete skill-name example on the same page. The [Codex CLI developer-command reference](https://learn.chatgpt.com/docs/developer-commands?surface=cli) documents the `/skills` sequence: type `/skills`, pick the skill, and Codex inserts the selected skill context so the next request follows that skill's instructions.

**Implication:** The exact direct mention formed from the installed skill name is `$dictate-plan`; `/skills` plus selection is an equivalent documented UI action in Codex CLI.

### Finding 3: A natural-language fallback is supported only as prompting, not dedicated syntax

The “no dedicated syntax” condition does not apply because Codex has documented `$` skill mentions. If a user still chooses plain language, the best-supported wording is: `Use the dictate-plan skill to turn the following notes into a dependency-ordered plan: ...`. First-party documentation does not establish that this wording is a guaranteed explicit invocation; without the `$dictate-plan` mention or `/skills` selection, Codex may treat it as ordinary prompt text and choose the skill through its normal matching behavior.

**Evidence:** The [Build skills documentation](https://learn.chatgpt.com/docs/build-skills#how-codex-uses-skills) distinguishes explicit invocation from implicit invocation based on a task matching the skill `description`. It documents `$` mentions and `/skills` for explicit invocation but gives no separate natural-language command grammar for a named skill.

**Implication:** Use `$dictate-plan` or the `/skills` picker when explicit activation matters. Plain wording can communicate intent, but current first-party evidence does not support treating it as equivalent dedicated syntax.

## Notes

- Current first-party documentation provides a generic skill-invocation rule rather than a `dictate-plan`-specific example. Applying the documented `$<skill-name>` form to the exact installed name yields `$dictate-plan`.
- The documentation confirms `$` mentions in Codex and `/skills` in the CLI. It does not document `/dictate-plan` as a slash command.

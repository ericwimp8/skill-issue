# Cursor Manual Invocation of `dictate-plan`

## Assignment

**Goal:** Determine whether a Cursor user can manually invoke a specific installed Agent Skill named `dictate-plan`, identify the exact action or syntax, and state the best-supported explicit wording only if dedicated syntax is unavailable.

**Scope:** Current official Cursor documentation and first-party Cursor release material, limited to manual invocation of an already installed Agent Skill in Agent chat.

**Exclusions:** Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated Cursor behavior.

## Sources

- [Cursor Docs — Agent Skills](https://cursor.com/docs/skills), current page inspected 2026-07-19. Relevant sections: “How skills work” and the `disable-model-invocation` configuration entry.
- [Cursor 2.4 changelog — Subagents, Skills, and Image Generation](https://cursor.com/changelog/2-4), published 2026-01-22. Relevant section: “Skills.”

## Findings

### Finding 1 — A user can manually invoke a specific installed skill

Yes. Cursor’s current Agent Skills documentation says skills can be manually invoked from Agent chat by typing `/` and searching for the skill name. Cursor’s 2.4 release notes independently state that a skill can be invoked using the slash-command menu.

**Evidence:** The [Agent Skills documentation](https://cursor.com/docs/skills) documents manual invocation by typing `/` in Agent chat and searching for the skill name. The [Cursor 2.4 changelog](https://cursor.com/changelog/2-4) confirms that users can invoke a skill through the slash-command menu.

**Implication:** An installed skill named `dictate-plan` has a documented manual invocation path in Cursor Agent chat, provided Cursor has discovered it and it appears in the skill search results.

### Finding 2 — The exact action is `/` plus selection of `dictate-plan`; the direct named form is `/dictate-plan`

In Cursor Agent chat, type `/`, search for `dictate-plan`, and select that skill from the slash-command menu. Cursor also documents the generic explicit form `/skill-name` for a skill configured for explicit-only invocation. Substituting the supplied skill name yields `/dictate-plan`.

**Evidence:** The [Agent Skills documentation](https://cursor.com/docs/skills) gives the UI action “typing `/` in Agent chat and searching for the skill name.” Its `disable-model-invocation` entry says that when the setting is `true`, the skill is included only when explicitly invoked via `/skill-name`; it later describes explicitly typing `/skill-name` in chat. The [2.4 changelog](https://cursor.com/changelog/2-4) corroborates the slash-command-menu route.

**Implication:** The best-supported dedicated invocation is `/dictate-plan`, with the most literal documented workflow being: open Agent chat, type `/`, search for `dictate-plan`, then choose it. The named form is a direct substitution into Cursor’s documented `/skill-name` syntax rather than a `dictate-plan`-specific example published by Cursor.

### Finding 3 — No natural-language fallback is needed or specifically documented

A dedicated invocation mechanism exists, so a prose-only fallback is unnecessary. The inspected first-party sources do not prescribe an exact natural-language sentence for forcing a named skill. If slash invocation is unavailable, the best-supported explicit wording is **“Use the `dictate-plan` skill to [describe the task].”** This wording is an inference, not documented Cursor syntax or a guaranteed invocation mechanism.

**Evidence:** The [Agent Skills documentation](https://cursor.com/docs/skills) says Cursor presents discovered skills to the agent, which decides relevance from context, while separately identifying `/` search as the manual mechanism. It also says `disable-model-invocation: true` prevents automatic contextual application and requires explicit `/skill-name` invocation.

**Implication:** Prefer `/dictate-plan`. Natural-language wording may help contextual skill selection only when the skill remains model-discoverable; first-party evidence is insufficient to treat prose as equivalent to the dedicated slash invocation.

## Notes

- The source constraint did not permit inspecting the installed `dictate-plan` skill or its frontmatter. Therefore, whether that particular installation sets `disable-model-invocation: true` is unverified. This does not change the documented manual slash path; it affects whether contextual, natural-language selection is available.

# Grok Build Manual Invocation of `dictate-plan`

## Assignment

**Goal:** Determine whether Grok Build lets a user manually invoke an installed agent skill named `dictate-plan`, identify the exact action or syntax, and provide the best-supported explicit wording only if dedicated syntax is unavailable.

**Scope:** Current first-party xAI documentation and first-party xAI product material for Grok Build skill discovery and user invocation.

**Exclusions:** Skill installation, CLI rollout, sub-agents, evaluation evidence, plugins, packaging, permissions, and unrelated Grok Build behavior.

## Sources

- [Skills, Plugins & Marketplaces — SpaceXAI Docs](https://docs.x.ai/build/features/skills-plugins-marketplaces), inspected July 19, 2026. Relevant statements: Grok discovers skills from documented project, user, enabled-extension, and configured paths; user-invocable skills appear as `/<skill-name>` slash commands; `/skills` opens the Skills tab in the extensions modal. Page last updated July 4, 2026.
- [Modes and Commands — SpaceXAI Docs](https://docs.x.ai/build/modes-and-commands), inspected July 19, 2026. Relevant statements: user-invocable skills appear as slash commands; the general form is `/<skill-name>`; name collisions require a qualified form; `/skills` opens the Skills tab. Page last updated July 2, 2026.
- [Grok Build — xAI](https://x.ai/cli), inspected July 19, 2026. Relevant statements: skills are auto-invoked when a task matches or “called by name,” and skills turn workflows into reusable slash commands.

## Findings

### 1. A user can manually invoke a specific installed skill when Grok Build discovers it and it is user-invocable

Yes. Current Grok Build documentation says that user-invocable skills appear as slash commands. The same documentation describes the locations from which Grok discovers skills and exposes a `/skills` command that opens the Skills tab, providing first-party evidence that the current product has a visible skill surface rather than only implicit model selection.

**Evidence:** The first-party [Skills, Plugins & Marketplaces documentation](https://docs.x.ai/build/features/skills-plugins-marketplaces) states that Grok discovers skills from its documented skill paths and that “User-invocable skills also appear as slash commands, for example `/<skill-name>`.” The first-party [Modes and Commands documentation](https://docs.x.ai/build/modes-and-commands) independently repeats that user-invocable skills appear as slash commands and documents `/skills` as the command that opens the Skills tab.

**Implication:** An installed `dictate-plan` skill is manually invocable if Grok Build has discovered that installation and the skill is user-invocable. The first-party pages do not establish the state of any particular local installation, so discovery and user-invocable status remain conditions rather than facts about a caller's machine.

### 2. The dedicated invocation syntax is `/dictate-plan`

In the Grok Build TUI, type `/dictate-plan`. This is the direct substitution of the skill's name into xAI's documented `/<skill-name>` form and is dedicated skill-invocation syntax, not natural-language prompting.

**Evidence:** The [Modes and Commands documentation](https://docs.x.ai/build/modes-and-commands) gives the exact generic form `/<skill-name>` under “Skills as commands.” Substituting the supplied installed skill name `dictate-plan` yields `/dictate-plan`. The same page says that collisions require a qualified form, giving `/local:commit` as an example; it does not document a universal qualifier for every installation scope.

**Implication:** Use `/dictate-plan` when the name is unambiguous. If Grok Build reports or displays a collision, use the qualified command shown by the product for that discovered skill. Current first-party documentation does not support `$dictate-plan` or a generic `/skill dictate-plan` form for Grok Build.

### 3. Natural-language wording is a fallback, not the primary invocation mechanism

A dedicated syntax exists, so no natural-language fallback is required. If `/dictate-plan` is unavailable, the best-supported explicit wording is: **“Use the `dictate-plan` skill to develop a dependency-ordered plan from my dictation.”** This wording names the skill and states the intended task, but it is a prompt rather than a documented deterministic invocation command.

**Evidence:** The first-party [Grok Build product page](https://x.ai/cli) says skills are “Auto-invoked when a task matches, or called by name,” while also describing skills as reusable slash commands. xAI does not publish a required natural-language sentence template for calling a skill by name.

**Implication:** Prefer `/dictate-plan` for explicit manual invocation. Use the named natural-language request only as a best-supported fallback, with less certainty that it will force-load the specific skill rather than rely on normal model matching.

## Notes

- First-party evidence is strong for the current Grok Build TUI's general user-invocable skill surface, but it does not verify whether a particular user's `dictate-plan` installation is discovered or marked user-invocable.
- The xAI documentation does not explicitly show argument syntax after a custom skill slash command. `/dictate-plan` itself is directly supported; any appended free-form arguments would be an inference and are therefore omitted.
- No first-party xAI source located in this assignment documents a different manual-invocation sigil or a generic skill runner command.

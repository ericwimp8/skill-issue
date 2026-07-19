# Google Antigravity and Gemini CLI Manual Skill Invocation

## Assignment

**Goal:** Determine whether a user can manually invoke an installed Agent Skill named `dictate-plan` in current Google Antigravity or Gemini CLI, and identify the exact supported action, syntax, or explicit wording.

**Scope:** Current first-party Google documentation and the first-party `google-gemini/gemini-cli` repository, limited to manual invocation of an already installed skill.

**Exclusions:** Skill installation, CLI rollout mechanics, sub-agents, evaluation evidence, plugins as a packaging topic, packaging, permissions, and unrelated product behavior.

## Sources

- Google Developers Blog, [“An important update: Transitioning Gemini CLI to Antigravity CLI”](https://developers.googleblog.com/an-important-update-transitioning-gemini-cli-to-antigravity-cli/), published May 19, 2026. Inspected the current product transition, retained Agent Skills capability, and remaining Gemini CLI availability.
- Google Antigravity Documentation, [“Agent Skills”](https://antigravity.google/docs/skills). Inspected discovery, activation, and the documented option to mention a skill by name.
- Google Antigravity Documentation, [“Plugins & Skills”](https://antigravity.google/docs/cli-plugins). Inspected Antigravity CLI's documented conversion of registered skills into same-name slash commands.
- Google Antigravity Documentation, [“CLI reference”](https://antigravity.google/docs/cli-reference). Inspected `/skills` and prompt submission behavior.
- Gemini CLI Documentation, [“Agent Skills”](https://geminicli.com/docs/cli/skills/). Inspected activation lifecycle and the complete documented `/skills` management command set.
- Gemini CLI Documentation, [“Get started with Agent Skills”](https://geminicli.com/docs/cli/tutorials/skills-getting-started/). Inspected the documented natural-language triggering flow.
- First-party repository documentation, [`docs/reference/commands.md`](https://github.com/google-gemini/gemini-cli/blob/main/docs/reference/commands.md). Cross-checked that Gemini CLI documents `/skills` subcommands for management but no per-skill invoke subcommand.

## Findings

### Finding 1: Current product and whether manual invocation is supported

**Answer (1): Yes.** The current consumer-facing terminal product is **Antigravity CLI**, and Google states that it retains Agent Skills from Gemini CLI. Google also documents Agent Skills in Antigravity 2.0. Gemini CLI still supports Agent Skills for enterprise customers and paid API-key access, but Google stopped serving consumer/free Gemini CLI requests on June 18, 2026. A user can deliberately request a named skill in Antigravity, and Antigravity CLI additionally documents same-name slash commands for CLI-registered skills.

**Evidence:** Google's transition announcement says Antigravity CLI is the replacement terminal experience, explicitly lists Agent Skills among retained critical features, and distinguishes the consumer cutoff from continued enterprise and paid API-key Gemini CLI access: [Google Developers Blog](https://developers.googleblog.com/an-important-update-transitioning-gemini-cli-to-antigravity-cli/). Antigravity's Agent Skills documentation says the agent discovers installed skills and that a user “can mention a skill by name” to ensure it is used: [Antigravity Agent Skills](https://antigravity.google/docs/skills).

**Implication:** For a current general-user answer, name **Antigravity CLI** first. Treat Gemini CLI as a still-documented but access-dependent product rather than the default current consumer harness.

### Finding 2: Exact action or syntax for `dictate-plan`

**Answer (2):** In **Antigravity CLI**, type **`/dictate-plan`** in the prompt box and press **Enter**, provided the installed `dictate-plan` skill is registered in the CLI's slash-command skill format and appears in the CLI's loaded skills. Antigravity CLI documents that registered skills automatically become slash commands, using `/format-tests` as its example; substituting the declared skill name yields `/dictate-plan`.

**Evidence:** The current Antigravity CLI documentation states that skills “convert automatically into slash commands inside the TUI” and gives `/refactor-ui` and `/format-tests` as examples derived from skill names: [Antigravity Plugins & Skills](https://antigravity.google/docs/cli-plugins). The CLI reference says to type `/` in the prompt box to open command typeahead, lists `/skills` as the browser for loaded Agent Skills, and documents Enter as prompt submission: [Antigravity CLI reference](https://antigravity.google/docs/cli-reference).

**Implication:** `/dictate-plan` is the dedicated manual invocation syntax supported by the Antigravity CLI-specific skill documentation. `/skills` is for browsing loaded skills; it is not the invocation syntax itself.

### Finding 3: Explicit wording where no dedicated syntax is documented

**Answer (3):** For **Antigravity's directory-style Agent Skills** and **Gemini CLI**, the best-supported explicit wording is:

> Use the `dictate-plan` skill for this request: [describe the plan you want to develop].

Gemini CLI does not document `/dictate-plan` or a `/skills invoke dictate-plan` command. Its documented flow is semantic: the user asks a relevant question, Gemini matches the prompt to the skill description, calls its internal `activate_skill` tool, and requests activation consent. Antigravity likewise documents that mentioning the skill by name can ensure it is used.

**Evidence:** Gemini CLI's Agent Skills guide documents only `/skills list`, `link`, `disable`, `enable`, and `reload` for interactive skill management, while activation occurs when Gemini identifies a matching task: [Gemini CLI Agent Skills](https://geminicli.com/docs/cli/skills/). Its getting-started guide triggers an example skill with a natural-language request rather than a skill-specific slash command: [Gemini CLI skill tutorial](https://geminicli.com/docs/cli/tutorials/skills-getting-started/). The first-party command reference independently lists the same management-only `/skills` subcommands: [`commands.md`](https://github.com/google-gemini/gemini-cli/blob/main/docs/reference/commands.md). Antigravity explicitly says a user may mention a skill by name to ensure its use: [Antigravity Agent Skills](https://antigravity.google/docs/skills).

**Implication:** Use explicit natural-language naming as the portable, documented method across Antigravity and Gemini CLI. In Gemini CLI, do not present `/dictate-plan` as supported syntax.

## Notes

- First-party Antigravity documentation currently describes two related skill representations: directory-based Agent Skills containing `SKILL.md`, and Antigravity CLI skills compiled into slash commands. The CLI page clearly supports `/dictate-plan` by name substitution for its registered slash-command format, but first-party evidence does not explicitly confirm that every directory-style `SKILL.md` package is automatically exposed as a same-name slash command. For that package shape, the explicit natural-language wording above is the supported fallback.
- The Antigravity Agent Skills page is the stronger source for shared Antigravity behavior; the Antigravity CLI page is the stronger source for dedicated slash-command behavior.

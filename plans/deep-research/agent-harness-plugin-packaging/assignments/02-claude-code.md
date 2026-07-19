# Claude Code Packaging

## Assignment

**Goal:** Determine how current Claude Code plugin and extension packaging works, and what Skill Issue must implement to distribute skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, and supporting files as a functional bundle.

**Scope:** Current Claude Code plugins and marketplaces; Agent Skills and legacy custom commands; hooks; custom subagents; MCP; plugin executables and configuration; installation, update, discovery, invocation, trust, permissions, and user/project/local/managed scopes. Evidence is limited to official Anthropic documentation and Anthropic-maintained public repositories inspected on 2026-07-19.

**Exclusions:** Local product source, unpublished behavior, Claude Desktop extensions outside the documented Claude Code plugin system, implementation coding, and assumptions about Skill Issue's current repository layout.

## Sources

- [Create plugins — Claude Code Docs](https://code.claude.com/docs/en/plugins), current documentation inspected 2026-07-19; plugin-versus-standalone guidance, local testing, distribution, community submission, and CLI recommendation link.
- [Plugins reference — Claude Code Docs](https://code.claude.com/docs/en/plugins-reference), current documentation inspected 2026-07-19; component schema, paths, scopes, environment variables, persistent data, caching, dependencies, configuration, and CLI management commands. The page includes explicit behavior changes through Claude Code v2.1.212.
- [Create and distribute a plugin marketplace — Claude Code Docs](https://code.claude.com/docs/en/plugin-marketplaces), current documentation inspected 2026-07-19; marketplace schema, sources, strict mode, managed restrictions, and version resolution.
- [Discover and install prebuilt plugins through marketplaces — Claude Code Docs](https://code.claude.com/docs/en/discover-plugins), current documentation inspected 2026-07-19; discovery UI, installation scopes, reloads, auto-update, trust, community marketplace, and team configuration.
- [Extend Claude with skills — Claude Code Docs](https://code.claude.com/docs/en/slash-commands), current documentation inspected 2026-07-19; skill locations, supporting files, invocation, permissions, subagent execution, and shell preprocessing.
- [Create custom subagents — Claude Code Docs](https://code.claude.com/docs/en/sub-agents), current documentation inspected 2026-07-19; subagent scopes, skill cooperation, tools, and plugin-agent limitations.
- [Hooks reference — Claude Code Docs](https://code.claude.com/docs/en/hooks), current documentation inspected 2026-07-19; lifecycle events, hook types, configuration locations, and security behavior.
- [Connect Claude Code to tools via MCP — Claude Code Docs](https://code.claude.com/docs/en/mcp), current documentation inspected 2026-07-19; plugin MCP lifecycle, scoped names, path substitution, and MCP scopes.
- [Claude Code settings — Claude Code Docs](https://code.claude.com/docs/en/settings), current documentation inspected 2026-07-19; managed/user/project/local precedence and plugin settings.
- [Configure permissions — Claude Code Docs](https://code.claude.com/docs/en/permissions), current documentation inspected 2026-07-19; permission enforcement and workspace configuration discovery.
- [Recommend your plugin from your CLI — Claude Code Docs](https://code.claude.com/docs/en/plugin-hints), current documentation inspected 2026-07-19; CLI-to-plugin hint protocol and official-marketplace restriction.
- [anthropics/claude-plugins-official](https://github.com/anthropics/claude-plugins-official), `main` inspected 2026-07-19; primary marketplace structure, trust warning, plugin layout, immutable marketplace names, and skill-bundle entries.
- [anthropics/claude-code marketplace catalog](https://github.com/anthropics/claude-code/blob/main/.claude-plugin/marketplace.json), `main` inspected 2026-07-19; primary example of a marketplace containing workflow, agent, hook, and output-style plugins.

## Findings

### Finding 1: A Claude Code plugin is the primary bundled delivery unit

Claude Code defines a plugin as a self-contained directory. A single plugin can contribute skills, legacy commands, subagents, hooks, MCP servers, LSP servers, monitors, output styles, themes, executables, scripts, default plugin settings, documentation, and arbitrary supporting files kept inside the plugin root. The manifest at `.claude-plugin/plugin.json` supplies identity and optional component paths; components in conventional root-level locations can also be auto-discovered. Plugin skills and agents are namespaced, which avoids collisions with user and project components.

**Evidence:** The [plugin reference](https://code.claude.com/docs/en/plugins-reference) identifies plugins as self-contained and lists skills, agents, hooks, MCP, LSP, and monitors as components. Its standard layout additionally defines `bin/`, `scripts/`, `settings.json`, `output-styles/`, `themes/`, and supporting files. The [official marketplace repository](https://github.com/anthropics/claude-plugins-official) shows the manifest plus MCP, commands, agents, and skills layout in Anthropic's own catalog.

**Implication:** Skill Issue should make one `skill-issue` plugin the default Claude Code artifact. It can carry the cohesive in-Claude experience without splitting each skill, hook, or agent into separately installed products. A minimal root should include `.claude-plugin/plugin.json`, `skills/`, `agents/`, `hooks/hooks.json` if lifecycle automation is needed, `.mcp.json` only if a true MCP tool surface exists, `bin/` for Claude-invoked executables, `scripts/`, and bundled reference/asset directories reachable from skills or scripts.

### Finding 2: Markdown guidance must be expressed through skills or agents, not a plugin-root CLAUDE.md

Custom commands have been merged into Agent Skills. A plugin skill lives under `skills/<name>/SKILL.md`, may be invoked as `/plugin-name:skill-name`, and can be selected automatically from its description. It can carry references, examples, templates, assets, and scripts beside `SKILL.md`, with those files loaded or executed only when the skill directs Claude to use them. Flat `commands/*.md` remains compatible, but Anthropic recommends skills for new plugins. A `CLAUDE.md` placed at the plugin root is explicitly ignored as project context.

**Evidence:** [Skills documentation](https://code.claude.com/docs/en/slash-commands) states that `.claude/commands/deploy.md` and `.claude/skills/deploy/SKILL.md` create equivalent invocations, recommends skills, documents supporting files and scripts, and explains progressive loading. The [plugin reference](https://code.claude.com/docs/en/plugins-reference) states that a plugin-root `CLAUDE.md` is not loaded and directs authors to skills, agents, and hooks instead.

**Implication:** Convert every reusable Skill Issue guidance unit into a namespaced skill. Put long Markdown references and examples beside the owning `SKILL.md` and link them explicitly from it. Use agents for durable specialized system prompts. Reserve `commands/` only for compatibility. Any guidance expected to apply automatically must be represented by an appropriately described model-invocable skill, an agent, or a justified hook; copying a general `CLAUDE.md` into the plugin will not work.

### Finding 3: Skills provide invocation controls, temporary tool grants, and forked-agent workflows

Skills can be user-only, model-only, or both; constrain tools; select model and effort; scope activation by paths; install lifecycle hooks; and run with `context: fork` in a chosen subagent. Supporting shell interpolation executes before the skill prompt is delivered, but administrators can disable this execution. Skill permission grants apply only for the invoking turn and do not override harness-level deny rules.

**Evidence:** The [skills frontmatter reference](https://code.claude.com/docs/en/slash-commands) documents `disable-model-invocation`, `user-invocable`, `allowed-tools`, `disallowed-tools`, `context`, `agent`, `hooks`, and `paths`; it also documents `!` command preprocessing and the managed `disableSkillShellExecution` switch. The [permissions reference](https://code.claude.com/docs/en/permissions) states that Claude Code, rather than the model or prompt, enforces permissions, with deny taking precedence over ask and allow.

**Implication:** Mark state-changing Skill Issue workflows as manually invoked, declare only the narrow tool patterns they need, and avoid depending on a skill grant to bypass user or enterprise policy. Prefer explicit script calls using `${CLAUDE_PLUGIN_ROOT}` over opaque inline shell. Treat shell preprocessing as optional because managed environments can disable it; core workflows should still give a useful blocked explanation or use an approved plugin executable/tool path.

### Finding 4: Plugin subagents cooperate with skills, but their security-sensitive fields are restricted

Plugin agents appear under scoped names, can be invoked manually or automatically, and can preload skills. Even without preload, subagents can discover project, user, and plugin skills through the Skill tool when that tool is available. Plugin-shipped agents cannot declare their own hooks, inline MCP servers, or permission mode; those fields are ignored for plugin agents. Those capabilities must live at plugin scope or in user/project agent definitions outside the plugin.

**Evidence:** The [plugin reference](https://code.claude.com/docs/en/plugins-reference) documents plugin agent discovery and explicitly excludes `hooks`, `mcpServers`, and `permissionMode` for plugin agents. The [subagent documentation](https://code.claude.com/docs/en/sub-agents) states that the `skills` field preloads full skill content and that unlisted project, user, and plugin skills remain discoverable through the Skill tool.

**Implication:** Skill Issue can ship specialized agents and their supporting skills together, but should put shared hooks and MCP definitions at the plugin root. Agent prompts should reference plugin skills by their scoped identifiers. If an agent requires a unique permission mode or inline MCP definition, that exact behavior is unsupported in a marketplace plugin agent and must be redesigned around plugin-wide components or documented user/project configuration.

### Finding 5: Hooks and MCP can be bundled and activated with the plugin

Plugin hooks can run commands, HTTP calls, MCP tools, prompts, or agent verifiers across Claude Code lifecycle events. Plugin MCP servers declared in `.mcp.json` or `plugin.json` start automatically when the plugin is enabled and expose tools under plugin-scoped names. Both surfaces can reference bundled code through `${CLAUDE_PLUGIN_ROOT}`, persistent state through `${CLAUDE_PLUGIN_DATA}`, and the active repository through `${CLAUDE_PROJECT_DIR}`.

**Evidence:** The [hooks reference](https://code.claude.com/docs/en/hooks) defines hook types and lifecycle events. The [MCP documentation](https://code.claude.com/docs/en/mcp) documents automatic server connection, plugin-scoped MCP tool names, path substitutions, and automatic setup. The [plugin reference](https://code.claude.com/docs/en/plugins-reference) shows both root files and inline manifest definitions.

**Implication:** Use hooks only for genuine lifecycle enforcement or context injection, and use MCP only when Skill Issue needs a structured long-running tool protocol. Ordinary deterministic utilities should remain scripts or `bin/` commands. Any hook matcher or permission rule for bundled MCP tools must use the full `mcp__plugin_<plugin>_<server>__<tool>` name. After installing or updating code-bearing components, Skill Issue onboarding should direct the user to `/reload-plugins`; monitors still require a new session.

### Finding 6: Bundled executables are available inside Claude Code, while a general standalone CLI remains a separate distribution concern

Executables placed in a plugin's `bin/` directory are added to the Bash tool's `PATH` while the plugin is enabled. Bundled scripts, binaries, config, references, and assets can be addressed through `${CLAUDE_PLUGIN_ROOT}`. That installation path is versioned, cached, and ephemeral; generated state and installed dependencies belong in `${CLAUDE_PLUGIN_DATA}`. Official documentation only promises `bin/` availability to Claude Code tool calls, not global installation into a user's interactive shell.

**Evidence:** The [plugin reference](https://code.claude.com/docs/en/plugins-reference) states that `bin/` files are invokable as bare commands in any Bash tool call while enabled, defines `${CLAUDE_PLUGIN_ROOT}` and `${CLAUDE_PLUGIN_DATA}`, and provides a `SessionStart` pattern that installs npm dependencies into the persistent data directory. It also states that marketplace plugins are copied to `~/.claude/plugins/cache`, outside-root traversal does not work, and old versions are later removed.

**Implication:** Bundle a thin Skill Issue executable in `bin/` for agent-driven use and keep all runtime-relative resources inside the plugin. Store caches, downloaded dependencies, and mutable state in `${CLAUDE_PLUGIN_DATA}`. If “standalone CLI” means a command humans can run from any ordinary terminal, continue publishing a native/npm/Homebrew-style CLI separately; the plugin should integrate with it or carry its own agent-only launcher. No authoritative source inspected establishes that plugin installation adds `bin/` to the user's global shell PATH.

### Finding 7: Plugin-level configuration is narrower than general Claude Code settings

A root `settings.json` can currently set only the plugin's default main `agent` and `subagentStatusLine`; unknown keys are ignored. Plugin-specific configurable values instead belong in manifest `userConfig`, which prompts users at enable time, supports sensitive storage, and exposes values to MCP, LSP, hooks, skills, and agents with field-specific substitution rules. Non-sensitive plugin options are read from user, command-line, or managed settings, not project/local settings in current versions.

**Evidence:** The [plugin reference](https://code.claude.com/docs/en/plugins-reference) limits plugin `settings.json` to `agent` and `subagentStatusLine`, documents `userConfig`, secure values, environment variables, and the current user/managed storage boundary. The [settings reference](https://code.claude.com/docs/en/settings) describes the wider managed, CLI, local, project, and user configuration hierarchy.

**Implication:** Do not try to package arbitrary Claude Code permission, environment, or behavior settings in the plugin's `settings.json`. Put genuine plugin options in `userConfig`, persist mutable application configuration in `${CLAUDE_PLUGIN_DATA}`, and provide optional, explicit snippets or an installer command for repository/user settings that must live outside the plugin. Managed deployments can centrally supply plugin options and policy; team-shared plugin options need a different design because current project/local `pluginConfigs` entries are ignored.

### Finding 8: Installation scope and component scope are distinct

Marketplace plugins install at `user`, `project`, or `local` scope; administrators provide `managed` plugins. User scope applies across projects, project scope records shared enablement in `.claude/settings.json`, local scope records a gitignored machine override, and managed plugins are read-only except for updates. A project declaration does not itself make an external plugin payload magically present: each collaborator must trust the repository and install it, after which the shared setting controls enablement. Standalone skills and agents have their own enterprise/personal/project locations outside the marketplace plugin system.

**Evidence:** The [plugin reference](https://code.claude.com/docs/en/plugins-reference) maps plugin scopes to settings files. The [discovery guide](https://code.claude.com/docs/en/discover-plugins) states that externally sourced project-enabled plugins remain unloaded until installed and that repository trust prompts precede team marketplace/plugin installation. The [skills documentation](https://code.claude.com/docs/en/slash-commands) separately lists enterprise, personal, project, and plugin skill locations. The [settings reference](https://code.claude.com/docs/en/settings) gives precedence as managed, command line, local, project, then user.

**Implication:** Skill Issue should document two recommended installs: user scope for an individual's cross-project toolkit, and project scope for a repository-pinned team experience. Project onboarding must include marketplace registration or `extraKnownMarketplaces`, repository trust, and plugin installation. Enterprise packaging should offer managed `extraKnownMarketplaces` and `enabledPlugins` guidance rather than attempting to emulate managed scope from the plugin itself.

### Finding 9: Marketplace distribution provides discovery, versioning, and updates

Users ordinarily add a marketplace, then install a plugin from it. Marketplace catalogs may be hosted on GitHub, another Git server, a local path, or a remote JSON URL; individual plugins may come from relative paths, GitHub, Git URLs, Git subdirectories, or npm packages. GitHub is Anthropic's recommended marketplace host. Official marketplaces auto-update by default, while third-party marketplaces default to manual updates. Version identity resolves from plugin manifest version, then marketplace version, then source commit SHA; a fixed manifest version must be bumped or updates are skipped.

**Evidence:** The [marketplace guide](https://code.claude.com/docs/en/plugin-marketplaces) documents catalog and plugin source types, recommends GitHub, and defines version resolution. The [discovery guide](https://code.claude.com/docs/en/discover-plugins) documents the two-step flow, `/plugin` inventory, `claude plugin install --scope`, `/reload-plugins`, and auto-update defaults. The [Claude Code marketplace catalog](https://github.com/anthropics/claude-code/blob/main/.claude-plugin/marketplace.json) is a primary example of multiple bundled workflow plugins in one catalog.

**Implication:** Publish a Skill Issue marketplace repository or add `.claude-plugin/marketplace.json` to its repository, with one main plugin entry. Prefer a GitHub or `git-subdir` source and pin release SHAs for reproducibility, or deliberately omit explicit versions so commit SHA drives updates. Run `claude plugin validate` in release automation. Provide exact `plugin marketplace add`, `plugin install`, `plugin update`, and `/reload-plugins` instructions. Use an npm plugin source only when npm packaging materially simplifies cross-platform executable delivery.

### Finding 10: Official, community, private, and ad hoc distribution have different constraints

The official Anthropic marketplace is pre-registered and curated at Anthropic's discretion. The separate `claude-community` marketplace accepts third-party submissions after validation and automated safety screening, pins approved plugins to commit SHAs, and must be added by users. Private or public custom marketplaces need no Anthropic listing but must be added explicitly unless configured by project or managed settings. `--plugin-dir` and `--plugin-url` provide session-only development or archive loading rather than persistent installation.

**Evidence:** [Create plugins](https://code.claude.com/docs/en/plugins) distinguishes `claude-plugins-official` from `claude-community`, documents submission access and review, and says official inclusion has no application process. The [discovery guide](https://code.claude.com/docs/en/discover-plugins) documents the community marketplace and custom marketplace sources. The [plugin guide](https://code.claude.com/docs/en/plugins) documents `--plugin-dir` and session-only `--plugin-url` archive loading.

**Implication:** Skill Issue can launch immediately through its own marketplace, then seek community listing for easier discovery. Official-marketplace inclusion must be treated as an external opportunity rather than a release dependency. Use `--plugin-dir` for development and `--plugin-url` for controlled previews, not as the normal install/update story.

### Finding 11: Trust and policy are first-class installation constraints

Plugins and marketplaces can execute arbitrary code with the user's privileges. Claude Code shows a pre-install component inventory, but users remain responsible for trusting the source. Repository-provided plugins are gated by workspace trust; project MCP servers require per-server approval, project LSP servers require workspace trust, and background monitors do not load from project skills-directory plugins. Organizations can allowlist or block marketplaces and can force plugin state through managed settings. Tool permission deny/ask/allow rules remain authoritative over what plugin instructions attempt.

**Evidence:** The [discovery guide](https://code.claude.com/docs/en/discover-plugins) explicitly calls plugins highly trusted arbitrary-code components, shows the pre-install “Will install” inventory, and documents managed marketplace restrictions. The [plugin reference](https://code.claude.com/docs/en/plugins-reference) documents project trust and additional restrictions for project skills-directory plugins. The [marketplace guide](https://code.claude.com/docs/en/plugin-marketplaces) documents `strictKnownMarketplaces` and `blockedMarketplaces`. The [permissions reference](https://code.claude.com/docs/en/permissions) states that runtime policy is enforced outside the model.

**Implication:** Skill Issue should publish a transparent component inventory, source and release SHAs, permissions rationale, network behavior, data locations, and uninstall behavior. Keep hooks and MCP minimal. Avoid first-run code that silently installs global dependencies. Enterprise documentation should include exact marketplace allowlist and managed enablement entries, plus expected behavior when shell execution, MCP, or permissions are restricted.

### Finding 12: One plugin is sufficient for the in-Claude bundle, with two explicit fallback boundaries

The documented plugin surface can carry Skill Issue's skills, agents, Markdown references, scripts, executable launchers, assets, hook automation, MCP integrations, and plugin-owned state in one install. Plugin dependencies allow later decomposition into cooperating plugins with version constraints and transitive installation, but they add lifecycle complexity and are unnecessary unless components need independent release or policy boundaries. Two requirements remain outside a pure one-plugin promise: globally installing a standalone human-facing CLI, and applying arbitrary Claude Code user/project settings.

**Evidence:** The [plugin manifest reference](https://code.claude.com/docs/en/plugins-reference) supports plugin `dependencies` with semver constraints and automatic dependency management, while the component layout covers the in-Claude resources. The same reference limits root `settings.json` and only promises `bin/` access inside Bash tool calls. The [discovery guide](https://code.claude.com/docs/en/discover-plugins) confirms dependency auto-install and scoped plugin lifecycle.

**Implication:** Implement the first release as one plugin plus the existing standalone CLI distribution:

1. Package Claude-facing skills, agents, references, assets, and scripts in the plugin.
2. Put a thin `skill-issue` launcher in `bin/` for Claude tool calls, using `${CLAUDE_PLUGIN_ROOT}` and `${CLAUDE_PLUGIN_DATA}`.
3. Keep the system CLI installable through its native package channel; let the plugin detect and use it when present, with a bundled fallback only if licensing and platform support permit.
4. Use `userConfig` for plugin options and provide opt-in commands/snippets for settings that must be written at user or project scope.
5. Split into dependent plugins only when an MCP server, organization policy, platform binary, or release cadence genuinely requires independent installation.

### Finding 13: CLI-driven plugin recommendation is useful only after official marketplace inclusion

A standalone CLI can emit a one-line `<claude-code-hint>` marker when invoked inside Claude Code, causing a one-time user-confirmed plugin installation prompt. Claude Code never auto-installs the plugin. The hint is accepted only for plugins in an Anthropic-controlled official marketplace; custom and community marketplace targets are silently ignored.

**Evidence:** [Recommend your plugin from your CLI](https://code.claude.com/docs/en/plugin-hints) defines the `CLAUDECODE`/`CLAUDE_CODE_CHILD_SESSION` detection variables, marker syntax, user confirmation, prompt limits, and official-marketplace requirement.

**Implication:** Add the hint protocol only after Skill Issue receives official marketplace inclusion. Before then, the CLI can print ordinary onboarding instructions for adding the Skill Issue marketplace, but it cannot trigger Claude Code's native plugin recommendation dialog through the documented hint protocol.

## Notes

- **Caveat — rapidly evolving surface:** The inspected documentation contains behavior gates through Claude Code v2.1.212. Skill Issue should declare a minimum tested Claude Code version and revalidate packaging against the live plugin reference before each release.
- **Unsupported — global CLI installation from a plugin:** No inspected authoritative source says plugin installation adds `bin/` to a human user's normal shell PATH or authorizes a plugin to install a global executable. The documented guarantee is limited to Claude Code Bash tool calls.
- **Unsupported — arbitrary packaged settings:** No inspected authoritative source supports applying general permissions, environment, sandbox, marketplace, or MCP policy through plugin-root `settings.json`; only `agent` and `subagentStatusLine` are currently supported there.
- **Caveat — assets are supporting files, not a registered UI component:** Plugins may carry arbitrary in-root files and skills can reference them, but the inspected docs do not define a general “asset registry.” Skill Issue must address assets through skill-relative paths, scripts, MCP responses, or another documented component.
- **Useful validation commands:** `claude plugin validate <path>`, `claude --plugin-dir <path>`, `claude plugin details <name>@<marketplace>`, `claude plugin list --json`, and `/reload-plugins`.

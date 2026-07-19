# JetBrains AI and Junie Packaging

## Assignment

**Goal:** Determine the current packaging and customization surfaces for JetBrains IDE plugins, JetBrains AI Assistant, Junie in the IDE, and Junie CLI, then derive what Skill Issue must implement to distribute skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, and supporting files as a functional bundle.

**Scope:** Official JetBrains and JetBrains Marketplace documentation current through July 2026; the official `JetBrains/skills` and `JetBrains/junie-extensions` repositories; installation, discovery, invocation, updates, trust, permissions, signing, review, plugin dependencies, project configuration, MCP, agents, skills, commands, and extension marketplaces.

**Exclusions:** Local JetBrains product configuration; undocumented internal APIs inferred from installed products; implementation work; non-authoritative community packaging schemes; and claims that generic IntelliJ plugin mechanisms automatically integrate with AI Assistant or Junie without a documented extension surface.

## Sources

- [AI Assistant 2026.2: About AI Assistant](https://www.jetbrains.com/help/ai-assistant/about-ai-assistant.html), accessed 19 July 2026.
- [AI Assistant 2026.2: Agents](https://www.jetbrains.com/help/ai-assistant/agents.html), updated 3 July 2026.
- [AI Assistant 2026.2: Skills](https://www.jetbrains.com/help/ai-assistant/agent-skills.html), updated 15 July 2026.
- [AI Assistant 2026.2: Agent instructions](https://www.jetbrains.com/help/ai-assistant/configure-agent-behavior.html), updated 2 July 2026.
- [AI Assistant 2026.2: Configure project rules](https://www.jetbrains.com/help/ai-assistant/configure-project-rules.html), updated 27 May 2026.
- [AI Assistant 2026.2: Add and customize prompts](https://www.jetbrains.com/help/ai-assistant/prompt-library.html), updated 6 May 2026.
- [AI Assistant 2026.2: Model Context Protocol](https://www.jetbrains.com/help/ai-assistant/mcp.html), accessed 19 July 2026.
- [AI Assistant 2026.2: Activate agents](https://www.jetbrains.com/help/ai-assistant/activate-agents.html), updated 30 June 2026.
- [AI Assistant 2026.2: Agent Client Protocol](https://www.jetbrains.com/help/ai-assistant/acp.html), updated 12 June 2026.
- [AI Assistant 2026.2: Restrict or disable AI Assistant features](https://www.jetbrains.com/help/ai-assistant/disable-ai-assistant.html), updated July 2026.
- [Official `JetBrains/skills` repository](https://github.com/JetBrains/skills), inspected 19 July 2026.
- [Junie: Getting started](https://junie.jetbrains.com/docs/), accessed 19 July 2026.
- [Junie: IDE plugin](https://junie.jetbrains.com/docs/junie-ide-plugin.html), updated 15 July 2026.
- [Junie: CLI quickstart](https://junie.jetbrains.com/docs/junie-cli.html), updated 15 July 2026.
- [Junie: Agent skills](https://junie.jetbrains.com/docs/agent-skills.html), updated 15 July 2026.
- [Junie: Add and configure extensions](https://junie.jetbrains.com/docs/junie-cli-extensions.html), updated 15 July 2026.
- [Junie: Custom slash commands](https://junie.jetbrains.com/docs/custom-slash-commands.html), updated 15 July 2026.
- [Junie: Custom subagents](https://junie.jetbrains.com/docs/junie-cli-subagents.html), updated 14 July 2026.
- [Junie: `config.json`](https://junie.jetbrains.com/docs/junie-cli-configuration.html), updated 15 July 2026.
- [Junie: CLI reference](https://junie.jetbrains.com/docs/parameters.html), updated July 2026.
- [Junie: Action Allowlist for CLI](https://junie.jetbrains.com/docs/action-allowlist-junie-cli.html), updated July 2026.
- [Junie: Action Allowlist for IDE](https://junie.jetbrains.com/docs/action-allowlist.html), updated 15 July 2026.
- [Junie: Integration with JetBrains IDEs](https://junie.jetbrains.com/docs/junie-cli-jetbrains-ide-integration.html), updated 15 July 2026.
- [Official `JetBrains/junie-extensions` repository](https://github.com/JetBrains/junie-extensions), inspected 19 July 2026; marketplace manifest at [`.junie-extension/marketplace.json`](https://github.com/JetBrains/junie-extensions/blob/main/.junie-extension/marketplace.json).
- [IntelliJ Platform SDK: Plugin content](https://plugins.jetbrains.com/docs/intellij/plugin-content.html), updated 25 March 2025.
- [IntelliJ Platform SDK: Plugin dependencies](https://plugins.jetbrains.com/docs/intellij/plugin-dependencies.html), updated 7 November 2025.
- [IntelliJ Platform SDK: Plugin compatibility](https://plugins.jetbrains.com/docs/intellij/plugin-compatibility.html), accessed 19 July 2026.
- [IntelliJ Platform SDK: Actions](https://plugins.jetbrains.com/docs/intellij/plugin-actions.html), updated 10 March 2025.
- [IntelliJ Platform SDK: Extensions](https://plugins.jetbrains.com/docs/intellij/plugin-extensions.html), updated 31 October 2025.
- [JetBrains Marketplace: Approval Guidelines v1.3](https://plugins.jetbrains.com/docs/marketplace/jetbrains-marketplace-approval-guidelines.html), effective 31 March 2026.
- [IntelliJ Platform SDK: Plugin signing](https://plugins.jetbrains.com/docs/intellij/plugin-signing.html), updated 8 April 2025.
- [JetBrains Marketplace: Uploading a new plugin](https://plugins.jetbrains.com/docs/marketplace/uploading-a-new-plugin.html), updated 20 March 2026.
- [IntelliJ IDEA 2026.1: Install plugins](https://www.jetbrains.com/help/idea/managing-plugins.html), accessed 19 July 2026.
- [IntelliJ IDEA 2026.1: Update IntelliJ IDEA and plugins](https://www.jetbrains.com/help/idea/update.html), updated 24 March 2026.

## Findings

### Finding 1: JetBrains exposes four distinct packaging surfaces, not one shared plugin system

Generic IntelliJ Platform plugins, AI Assistant customization, Junie-in-IDE customization, and Junie CLI extensions are separate surfaces with different installers and lifecycle owners. AI Assistant itself is a plugin providing chat, editor features, and integrated coding agents. Junie can be selected inside AI Chat; the standalone Junie IDE plugin is only required for its separate tool window and for Junie CLI's richer IDE connection. Junie CLI is separately installed and has its own extension marketplace. A JetBrains Marketplace plugin ZIP therefore must not be treated as equivalent to an AI Assistant skill registry or a Junie extension.

#### Evidence

AI Assistant describes itself as a collection of IDE features and agents and lists independent activation methods in [About AI Assistant](https://www.jetbrains.com/help/ai-assistant/about-ai-assistant.html). The [Junie IDE plugin documentation](https://junie.jetbrains.com/docs/junie-ide-plugin.html) says manual Junie plugin installation is required only for the separate Junie tool window, while Junie is available through AI Chat. The [Junie CLI quickstart](https://junie.jetbrains.com/docs/junie-cli.html) installs a separate `junie` command, and [Junie extensions](https://junie.jetbrains.com/docs/junie-cli-extensions.html) are managed by `/extensions`, not by IDE Settings | Plugins.

#### Implication

Skill Issue needs at least two native distribution descriptions: an AI Assistant skills registry/repository for supported AI Assistant agents, and a Junie extension marketplace for Junie CLI. A JetBrains IDE plugin is an optional third artifact for IDE-native onboarding and management, not the universal bundle format.

### Finding 2: A generic JetBrains plugin is a reviewed executable IDE package with ordinary plugin dependencies

An IntelliJ plugin distribution is a JAR or ZIP containing `META-INF/plugin.xml`, implementation classes, resources, and bundled library JARs. It can expose actions, tool windows, settings pages, services, and extension implementations. It can declare required or optional dependencies on other bundled or Marketplace plugins. These capabilities are generic IDE composition mechanisms; they do not themselves prove access to AI Assistant or Junie configuration internals.

#### Evidence

[Plugin content](https://plugins.jetbrains.com/docs/intellij/plugin-content.html) defines the JAR/ZIP layout and bundled libraries. [Actions](https://plugins.jetbrains.com/docs/intellij/plugin-actions.html) and [Extensions](https://plugins.jetbrains.com/docs/intellij/plugin-extensions.html) document normal invocation through menus, toolbars, keyboard shortcuts, Find Action, tool windows, and settings pages. [Plugin dependencies](https://plugins.jetbrains.com/docs/intellij/plugin-dependencies.html) supports required and optional dependencies on other plugins, but also states that a dependency on a non-bundled plugin cannot specify its own minimum/maximum version. [Plugin compatibility](https://plugins.jetbrains.com/docs/intellij/plugin-compatibility.html) requires module/plugin declarations and verification against target IDE products.

#### Implication

A Skill Issue IDE plugin can provide a polished installer/manager UI, validate prerequisites, materialize project files with user consent, invoke the Skill Issue CLI, and offer its own actions. Any direct use of AI Assistant or Junie classes or extension points requires a separately validated public API and compatibility strategy. No such stable, documented content-installation API was found in the official SDK documentation reviewed here, so directly editing internal AI/Junie settings is unsupported.

### Finding 3: AI Assistant has content-oriented customization, but it is split by consumer

AI Assistant chat rules, coding-agent instructions, agent skills, Prompt Library entries, and MCP servers are different mechanisms. Project rules are Markdown files under `.aiassistant/rules` and affect AI Assistant chat mode. Agent instruction files such as `AGENTS.md` affect the selected coding agent. Skills come from local directories or GitHub registries and are currently supported in AI Assistant by Claude Agent and Codex, not Junie. Prompt Library entries are created and managed through IDE settings. MCP servers are configured through AI Assistant settings at global or project level.

#### Evidence

[Agent instructions](https://www.jetbrains.com/help/ai-assistant/configure-agent-behavior.html) explicitly distinguishes agent instruction files from AI Assistant Project Rules. [Project Rules](https://www.jetbrains.com/help/ai-assistant/configure-project-rules.html) documents `.aiassistant/rules/*.md`, automatic, manual (`@rule:` or `#rule:`), model-decided, and file-pattern activation. [Skills](https://www.jetbrains.com/help/ai-assistant/agent-skills.html) accepts local directories and GitHub repository registries, installs to IDE, shared project `.agents/skills`, or per-agent locations, and names only Claude Agent and Codex as currently supported; skills may run automatically or by `$skill-name`. The official [`JetBrains/skills` repository](https://github.com/JetBrains/skills) demonstrates a flat repository of skill folders. [Prompt Library](https://www.jetbrains.com/help/ai-assistant/prompt-library.html) documents UI-managed custom prompts and AI Actions invocation. [MCP](https://www.jetbrains.com/help/ai-assistant/mcp.html) documents global/project server levels, JSON snippets, STDIO and HTTP transports, and automatic or `/`-command invocation.

#### Implication

For AI Assistant, Skill Issue should publish its skills as a GitHub repository accepted by the Skills Manager and also offer a project materializer that installs `.aiassistant/rules`, `AGENTS.md`, and `.agents/skills` where appropriate. Prompt Library and AI Assistant MCP onboarding need explicit IDE steps or a separately validated integration API; the official docs do not define a repository manifest that atomically installs rules, prompts, MCP, and skills together.

### Finding 4: Junie has its own rich file-based model, including a native multi-component extension bundle

Junie supports guidelines, agent skills, MCP, custom slash commands, custom subagents, and CLI configuration. Skills can contain Markdown, scripts, templates, references, and other supporting files; they are discovered by relevance and can also be requested explicitly. Junie CLI extensions can package any combination of skills, MCP servers, subagents, commands, and guidelines in one extension.

#### Evidence

[Junie agent skills](https://junie.jetbrains.com/docs/agent-skills.html) defines `<project>/.junie/skills/<name>/SKILL.md` and `~/.junie/skills/<name>/`, supports arbitrary supporting subdirectories, progressive loading, bundled scripts, and use in both Junie CLI and Junie in JetBrains IDEs. The [Junie IDE plugin](https://junie.jetbrains.com/docs/junie-ide-plugin.html) resolves guidelines from a custom path, `.junie/AGENTS.md`, root `AGENTS.md`, then deprecated legacy paths, and resolves MCP from project `.junie/mcp/mcp.json` or global `~/.junie/mcp/mcp.json`. [Custom commands](https://junie.jetbrains.com/docs/custom-slash-commands.html) are Markdown files under `.junie/commands` or `~/.junie/commands`; [custom subagents](https://junie.jetbrains.com/docs/junie-cli-subagents.html) support tailored prompts, models, skill lists, MCP-server allowlists, and built-in tool allow/deny lists. [Junie extensions](https://junie.jetbrains.com/docs/junie-cli-extensions.html) explicitly bundles skills, MCP, subagents, commands, and guidelines.

#### Implication

Skill Issue maps cleanly to one Junie extension: `skills/` holds SKILL.md plus scripts, references, templates, and assets; `agents/` holds specialized agent definitions; `commands/` holds invocation shortcuts; `guidelines/` holds persistent guidance; and `mcp/.mcp.json` launches supporting tools. The standalone Skill Issue CLI remains a separately installed executable or package that the MCP configuration and bundled scripts can call.

### Finding 5: Junie extension marketplaces provide installation and updates, but their trust model differs from JetBrains Marketplace

A Junie marketplace is a Git repository, local directory, or direct HTTPS `marketplace.json`. The native manifest is `.junie-extension/marketplace.json`; Claude-compatible plugin marketplaces are also accepted. Users register a marketplace, browse its catalog, and install an extension at project or user scope. Content is cached per user, project installation records can be committed, and updates are explicitly pulled. Custom extension marketplaces are not JetBrains Marketplace plugins and do not inherit JetBrains Marketplace signing or review.

#### Evidence

[Junie extensions](https://junie.jetbrains.com/docs/junie-cli-extensions.html) documents all three marketplace transports, native and Claude manifest formats, project `.junie/extensions.json`, user `~/.junie/extensions/extensions.json`, cached content under `~/.junie/extensions/<marketplace>/<extension>/`, immediate availability without restart, explicit Update, and automatic source registration for teammates who pull project references. The official [`JetBrains/junie-extensions` repository](https://github.com/JetBrains/junie-extensions) shows `extension.json`, `skills/`, `agents/`, `guidelines/`, and `mcp/.mcp.json`, and warns users that JetBrains does not control or verify all extension-contained MCP servers, files, or software and that users must review them before use.

#### Implication

Skill Issue should host a native Junie marketplace repository with a versioned catalog entry and one or more semantically scoped extensions. Its installer can register that marketplace and select project or user scope, while teams can commit `.junie/extensions.json`. Skill Issue should display provenance, commit/tag, included executable/MCP components, and a pre-install file list because no documented custom-marketplace signature or Marketplace-style moderation mechanism was found.

### Finding 6: Installation does not equal authorization; runtime actions remain permission-gated

JetBrains plugin installation grants executable code inside the IDE, while AI and Junie tool execution has separate authorization controls. Junie asks before sensitive terminal, MCP, outside-project read/edit, and related actions unless Brave mode or an allowlist rule authorizes them. `.aiignore` influences access but can be bypassed by Brave mode or allowed commands. AI Assistant can also be centrally restricted, and MCP availability may be administrator-controlled.

#### Evidence

[Junie CLI Action Allowlist](https://junie.jetbrains.com/docs/action-allowlist-junie-cli.html) identifies `fileEditing`, `executables`, `mcpTools`, and `readOutsideProject` action classes and defaults to asking without Brave mode. The [IDE Action Allowlist](https://junie.jetbrains.com/docs/action-allowlist.html) likewise requires approval by default and currently allows MCP only as an all-MCP-tools category rather than per tool. The [Junie IDE plugin](https://junie.jetbrains.com/docs/junie-ide-plugin.html) documents `.aiignore` limitations. [AI Assistant MCP](https://www.jetbrains.com/help/ai-assistant/mcp.html) says administrators can preconfigure MCP servers and prohibit users from adding their own. [AI Assistant restrictions](https://www.jetbrains.com/help/ai-assistant/disable-ai-assistant.html) documents project disablement and `.noai`/`.aiignore` controls.

#### Implication

Skill Issue must keep installation and permission granting separate. Bundle installation should never silently enable Brave mode or broad allowlists. Onboarding should explain each script, executable, MCP server, network destination, and filesystem scope, then let the host request approval during real invocation.

### Finding 7: JetBrains Marketplace adds strong distribution controls, but only to the IDE plugin artifact

Marketplace plugins and every update undergo automated and manual review, Plugin Verifier checks, UI integration checks, and continuing moderation. Plugin archives may be author-signed and are signed by JetBrains Marketplace; unsigned or revoked-author installations can produce a warning. IDEs can install from Marketplace, disk, custom repositories, or the command line, and Marketplace plugin updates can be downloaded automatically and applied at restart.

#### Evidence

[Approval Guidelines v1.3](https://plugins.jetbrains.com/docs/marketplace/jetbrains-marketplace-approval-guidelines.html) applies review to every new plugin and update and requires compatibility, security/privacy, performance, non-interference with licensing/subscriptions, and Plugin Verifier compliance. [Plugin signing](https://plugins.jetbrains.com/docs/intellij/plugin-signing.html) documents author and Marketplace signing and install warnings for missing/revoked author certificates. [Uploading a new plugin](https://plugins.jetbrains.com/docs/marketplace/uploading-a-new-plugin.html) requires a vendor profile, agreement, license, metadata, and a package up to 400 MB. [Install plugins](https://www.jetbrains.com/help/idea/managing-plugins.html) covers Marketplace, disk, and custom repositories; [updates](https://www.jetbrains.com/help/idea/update.html) can be automatically downloaded and applied on restart.

#### Implication

If Skill Issue ships an IDE plugin, it should be minimal, signed, Plugin-Verifier-clean across declared products, transparent about network/file writes, and use Marketplace stable/beta channels for the plugin's own lifecycle. Those assurances do not automatically cover content later fetched from a Skill Issue registry or Junie marketplace, so fetched artifacts need their own checksums, provenance, review UI, and rollback policy.

### Finding 8: One IDE plugin cannot currently provide the complete cross-surface bundle through documented APIs

One generic plugin can provide its own IDE-native user experience and can carry JVM code, resources, and libraries. It cannot, through the documented public mechanisms reviewed, atomically register an AI Assistant skills registry, create Prompt Library entries, configure AI Assistant MCP, install a Junie extension marketplace, and expose a cross-platform standalone shell command. Junie extensions cover Junie components well, but they are CLI-managed and do not constitute JetBrains Marketplace IDE plugins. AI Assistant Skills Manager covers skills for Claude Agent and Codex but not Junie, rules, prompts, MCP, or a standalone CLI.

#### Evidence

The supported generic package is limited to plugin code/resources/libraries in [Plugin content](https://plugins.jetbrains.com/docs/intellij/plugin-content.html), and it cannot target separate OS-specific plugin distributions. AI Assistant's documented installation paths remain independently managed in [Skills](https://www.jetbrains.com/help/ai-assistant/agent-skills.html), [Project Rules](https://www.jetbrains.com/help/ai-assistant/configure-project-rules.html), [Prompt Library](https://www.jetbrains.com/help/ai-assistant/prompt-library.html), and [MCP](https://www.jetbrains.com/help/ai-assistant/mcp.html). Junie uses separate [extension](https://junie.jetbrains.com/docs/junie-cli-extensions.html), [configuration](https://junie.jetbrains.com/docs/junie-cli-configuration.html), and CLI installation mechanisms. The SDK documentation reviewed did not identify a stable public AI Assistant or Junie extension point for third-party bundle installation.

#### Implication

The defensible architecture is a coordinated bundle family rather than a single binary:

1. **Canonical Skill Issue content package:** host-neutral source containing skills, guidance, scripts, references, assets, configuration templates, checksums, and compatibility metadata.
2. **Junie adapter:** native Junie marketplace plus extension layout for skills, agents, commands, guidelines, and MCP; installable with `/extensions` and shareable through `.junie/extensions.json`.
3. **AI Assistant adapter:** GitHub skills registry plus project materializer for `.agents/skills`, `.aiassistant/rules`, and `AGENTS.md`; explicit onboarding for Prompt Library and MCP where public automation is not verified.
4. **Standalone CLI distribution:** signed/reproducible packages or installers for supported operating systems, independently updateable and callable by skills, commands, MCP, and the IDE plugin.
5. **Optional JetBrains IDE plugin:** Marketplace-reviewed onboarding/manager UI that detects AI Assistant, Junie IDE plugin, and CLI availability; previews and applies project-scoped files with consent; invokes the CLI; records installed bundle versions; and links users into host-owned settings rather than editing undocumented internal state.

### Finding 9: Cooperation should occur through files, processes, and declared dependencies rather than private host internals

The most stable cooperation boundary is: version-controlled project files for instructions and content, a separately versioned CLI process for shared executable behavior, MCP for tool exposure, and ordinary optional plugin dependencies for IDE-only enhancements. Junie CLI can also connect passively to a running Junie IDE plugin for symbol-aware IDE capabilities, while remaining usable without that connection.

#### Evidence

[Junie IDE integration](https://junie.jetbrains.com/docs/junie-cli-jetbrains-ide-integration.html) says CLI discovery of matching IDE projects is automatic and passive, requires the Junie IDE plugin for IDE-backed features, and degrades to CLI-only behavior if no IDE is connected. [Junie CLI configuration](https://junie.jetbrains.com/docs/junie-cli-configuration.html) provides additional locations for skills, commands, agents, models, and MCP, with project/user precedence and relative-path resolution. [AI Assistant Agent instructions](https://www.jetbrains.com/help/ai-assistant/configure-agent-behavior.html) and [Junie IDE plugin](https://junie.jetbrains.com/docs/junie-ide-plugin.html) both use repository files for durable guidance. [Plugin dependencies](https://plugins.jetbrains.com/docs/intellij/plugin-dependencies.html) provides required/optional plugin composition for the IDE layer.

#### Implication

Skill Issue should make the CLI the behavior owner for install, update, verification, and shared commands; adapters should remain declarative. The optional IDE plugin should call that CLI through a narrow protocol and treat AI Assistant and Junie as optional capabilities. This reduces coupling to IDE build numbers and host plugin internals while preserving a cohesive user experience.

## Notes

- **Unsupported:** No official public IntelliJ SDK page reviewed documents a stable third-party API for programmatically adding AI Assistant Prompt Library entries, AI Assistant MCP configurations, Skills Manager registries, or Junie extension marketplaces. File locations explicitly documented by the products are usable; internal settings stores should not be assumed.
- **Caveat:** AI Assistant 2026.2 says its Skills Manager currently supports Claude Agent and Codex, while Junie's own documentation says `.junie/skills` work in both Junie CLI and Junie in JetBrains IDEs. These statements describe different skill managers and storage paths, not a contradiction: Junie does not currently consume AI Assistant's IDE-wide Skills Manager installation.
- **Caveat:** Junie extension documentation describes manual Update but does not document automatic extension updates, signature verification, checksum enforcement, or a review process for arbitrary custom marketplaces.
- **Caveat:** A plugin could technically download or write supporting files using ordinary JVM/platform capabilities, but Marketplace approval is discretionary and security/privacy/non-interference requirements apply. Such behavior should be explicit, minimal, and user-approved rather than treated as an entitlement.
- **Useful search terms:** `com.intellij.ml.llm`, `Junie extension marketplace`, `.junie-extension/marketplace.json`, `.junie/extensions.json`, `.agents/skills`, `.aiassistant/rules`, `AGENTS.md`, `plugin.xml optional dependency`, `Plugin Verifier`, `Action Allowlist`, `ACP`.

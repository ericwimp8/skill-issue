# Coding-Agent Harness Packaging for Skill Issue

> **Document status — historical ecosystem research.** This report and its assignments remain evidence about native plugin, extension, marketplace, and package systems. Its JetBrains-era target set and native-package recommendations do not own current implementation. CLI-managed direct installation is now governed by the [direct-install architecture](../../skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md) and its newer research corpus. Use the [documentation authority map](../../skill-issue-project-completion/document-authority-and-update-map.md) before applying any conclusion.

## Executive answer

Skill Issue should **not** attempt to define one universal plugin package. The nine selected harness entries expose several incompatible packaging systems, and even the strongest native plugin formats do not install every part of the requested experience. The most defensible product architecture is:

1. **One canonical Skill Issue bundle model** that owns versioning, component identity, compatibility, checksums, permissions, installation scope, and lifecycle metadata.
2. **Harness-specific renderers** that produce each host's native artifacts from that model.
3. **One independently distributed `skill-issue` CLI/installer** that owns unsupported cross-cutting work: executable installation, safe configuration materialization, upgrades, diagnostics, conflict handling, ownership receipts, and uninstall.
4. **Agent Skills as the primary portable content unit**, retaining each workflow's `SKILL.md`, scripts, references, assets, templates, and other supporting files together.
5. **MCP as an optional tool interoperability seam**, not as the bundle installer. Use it only when Skill Issue needs structured tools, shared state, remote services, or cross-client reuse.

This gives every harness the strongest experience it natively supports without claiming a standard that does not exist. GitHub Copilot, Claude Code, OpenAI Codex, Cursor, Gemini CLI, Junie, OpenCode, and Pi can each provide a meaningful native multi-component installation. Antigravity has a native plugin shape but incomplete public distribution/update semantics. Kilo has useful native Marketplace and npm-plugin channels, but neither spans the full bundle. JetBrains AI Assistant is principally a skills/project-composition target and requires a stronger installer handoff.

The core implementation rule is: **native packages should remain useful by themselves, while the CLI upgrades them into the complete Skill Issue experience**. Installation must never imply that executable code, account authorization, secrets, host permissions, global configuration, or project policy has been installed or trusted unless the relevant host path was actually completed and verified.

This direction follows the retained constrained corpus—[GitHub Copilot](assignments/01-github-copilot.md), [Claude Code](assignments/02-claude-code.md), [OpenAI Codex](assignments/03-openai-codex.md), [Cursor](assignments/04-cursor.md), [Google](assignments/05-google-antigravity-gemini-cli.md), [JetBrains](assignments/06-jetbrains-ai-junie.md), [OpenCode](assignments/07-opencode.md), and [Kilo](assignments/08-kilo-code.md)—plus Pi's official [package documentation](https://pi.dev/docs/latest/packages/) and [settings documentation](https://pi.dev/docs/latest/settings/).

## Cross-harness comparison

The two compound entries are split below because Antigravity and Gemini CLI, and AI Assistant and Junie, have materially different package managers and capability boundaries.

| Harness surface        | Native packaging unit                                                                                                                                                 | Bundle completeness for Skill Issue                                                                                                            | Native distribution and update                                                                                                                                                  | Trust and permissions                                                                                                                                                   | Required fallback                                                                                                                                                                                                                                                              |
| ---------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| GitHub Copilot         | Multi-component plugin with `plugin.json`; OpenPlugin-compatible layout can contain skills, agents, commands, hooks, MCP/LSP, and CLI extensions                      | **High for agent-host content**; repository instructions and a normal standalone CLI remain external                                           | Copilot marketplace, Git/GitHub/local install, update commands, VS Code discovery; cloud enablement is declarative                                                              | Marketplace/plugin trust, MCP startup/tool approval, scoped CLI permissions, enterprise policy; cloud runs in an ephemeral restricted environment                       | CLI installer for `PATH`, repository instructions/settings, unsupported IDEs, and cross-platform/cloud bootstrap ([assignment](assignments/01-github-copilot.md); [plugin reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-plugin-reference)) |
| Claude Code            | `.claude-plugin/plugin.json` plugin with skills, agents, commands, hooks, MCP/LSP, `bin/`, settings, and supporting files                                             | **Very high inside Claude Code**; general settings and a human-facing global CLI remain external                                               | Git/GitHub/npm/local marketplaces, scoped installs, explicit update/reload, official/community/private catalogs                                                                 | Pre-install component inventory, repository trust, marketplace controls, host permissions; plugins may execute arbitrary code                                           | CLI installer for global command and out-of-plugin settings; optional setup snippets ([assignment](assignments/02-claude-code.md); [plugin reference](https://code.claude.com/docs/en/plugins-reference))                                                                      |
| OpenAI Codex           | `.codex-plugin/plugin.json` with skills, MCP, apps/connectors, hooks, and presentation assets                                                                         | **Medium-high core**; custom-agent TOMLs, `AGENTS.md`, rules, config, IDE filesystem skills, and global CLI are separate                       | Marketplace catalogs from Git/local/npm and public directory submission; reinstall/new session needed for activation                                                            | Sandbox/approval policy, explicit hook trust, MCP/app authorization, marketplace/admin controls                                                                         | CLI initializer for agents/config/rules/guidance/IDE surface and executable installation ([assignment](assignments/03-openai-codex.md); [Build plugins](https://learn.chatgpt.com/docs/build-plugins))                                                                         |
| Cursor                 | `.cursor-plugin/plugin.json` with skills, agents, commands, rules, hooks, MCP, and supporting files                                                                   | **High for agent behavior**; global CLI, arbitrary config, and editor-native UI extension remain separate                                      | Cursor Marketplace, `/add-plugin`, local/team marketplaces, reviewed updates; Open VSX/VSIX is a separate channel                                                               | Plugin code/instructions are trusted; shell/MCP/fetch calls remain approval- and sandbox-governed                                                                       | CLI installer for executable/config; optional Open VSX extension only for real editor APIs/UI ([assignment](assignments/04-cursor.md); [official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json))                                              |
| Google Antigravity     | `plugin.json` with skills, rules, hooks, MCP config, and CLI-specific agents                                                                                          | **Medium-high content**, but component parity and install locations differ between desktop and CLI                                             | `agy plugin` management and manual/UI placement exist; public third-party marketplace and update contracts are incomplete                                                       | Host deny/ask/allow policy, sandbox, hook decisions, project-scoped access                                                                                              | Skill Issue-owned installer/updater, dual-path installation when both surfaces are requested, standalone CLI ([assignment](assignments/05-google-antigravity-gemini-cli.md); [Antigravity plugins](https://antigravity.google/docs/plugins))                                   |
| Gemini CLI             | `gemini-extension.json` extension with context, commands, skills, preview agents, hooks, policies, MCP, themes, settings                                              | **High**, with mature lifecycle; global CLI installation outside the extension still separate                                                  | GitHub/Git/local install, refs, releases, update/auto-update, link, enable/disable/uninstall                                                                                    | Install consent, trusted folders, skill activation consent, restricted extension policies, keychain settings, MCP/tool controls                                         | Standalone CLI installer; optional IDE companion remains separately installed ([assignment](assignments/05-google-antigravity-gemini-cli.md); [extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md))                       |
| JetBrains AI Assistant | Skills registry/repository plus separate project rules, agent instructions, Prompt Library, and MCP settings                                                          | **Low-medium as one install**; no documented atomic multi-component AI Assistant bundle                                                        | Skills Manager accepts local/GitHub registries; IDE plugin Marketplace is a separate executable distribution system                                                             | IDE/plugin trust, admin AI/MCP restrictions, agent-specific permissions                                                                                                 | Project materializer and explicit IDE setup; optional generic IDE manager plugin with no reliance on private AI APIs ([assignment](assignments/06-jetbrains-ai-junie.md); [AI Assistant Skills](https://www.jetbrains.com/help/ai-assistant/agent-skills.html))                |
| Junie                  | Native Junie extension/marketplace containing skills, agents, commands, guidelines, and MCP                                                                           | **High for Junie**; standalone CLI remains a separate package                                                                                  | Git/local/HTTPS marketplaces, project/user installs, cached content, explicit updates; project references can be committed                                                      | Sensitive actions prompt or use allowlists; custom marketplaces are not covered by JetBrains Marketplace review/signing                                                 | Standalone CLI package; optional JetBrains IDE plugin only for onboarding/management ([assignment](assignments/06-jetbrains-ai-junie.md); [Junie extensions](https://junie.jetbrains.com/docs/junie-cli-extensions.html))                                                      |
| OpenCode               | npm server plugin with tools/hooks and config-hook injection of skills, agents, commands, instructions, references, and MCP                                           | **High on current releases**; composition is configuration-driven rather than manifest-native                                                  | npm plus `opencode plugin`; explicit version replacement is the reliable update path; no first-party marketplace                                                                | Plugin is trusted in-process code; action permissions govern agent tools; `--pure` disables plugins; lifecycle scripts are ignored                                      | Normal CLI install/`npx`; version-gated materializer for older OpenCode; safe uninstall UX ([assignment](assignments/07-opencode.md); [Plugins](https://opencode.ai/docs/plugins/))                                                                                            |
| Kilo Code              | Native Marketplace items for individual skills, agents, and MCP; npm plugins for runtime hooks/tools; VSIX for editor integration                                     | **Fragmented**; no one Kilo-native item spans commands, rules, config, plugin, CLI, and content                                                | Curated Kilo Marketplace, npm, Visual Studio Marketplace, Open VSX, direct VSIX; channels update independently                                                                  | Marketplace review, extension publisher trust, Kilo action permissions, MCP approval; plugins/MCP lifecycle are trusted integrations                                    | Authoritative Skill Issue CLI; optional companion VSIX for one-click editor lifecycle; separate Marketplace entries ([assignment](assignments/08-kilo-code.md); [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace))                                                |
| Pi                     | Pi Package declared by a `package.json` `pi` manifest or conventional directories, bundling TypeScript extensions, skills, prompt templates, themes, and dependencies | **High inside Pi**; package resources and runtime extensions are native, while the standalone `skill-issue` CLI remains separately distributed | `pi install` from npm, Git, URL, or local path; user/project scopes; `pi list`, `pi remove`, `pi update --extensions`, filtering, enable/disable, and package gallery discovery | Packages run with full system access; project resources and missing packages load only after project trust; extension code and skill instructions require source review | Independent CLI binary plus an optional Pi extension bridge for invoking it; explicit handling of pinned Git/npm versions and owned settings entries ([Pi packages](https://pi.dev/docs/latest/packages/); [Pi settings](https://pi.dev/docs/latest/settings/))                |

### Discovery and invocation map

- **GitHub Copilot:** users discover/install plugins through Copilot CLI marketplaces, direct sources, or VS Code `@agentPlugins`; skills can be model-selected or explicitly invoked, commands are slash-invoked, agents can be selected or delegated to, and MCP contributes tools. [Finding and installing Copilot plugins](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/plugins-finding-installing); [Adding Copilot skills](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills)
- **Claude Code:** users add a marketplace and install at user/project/local scope; namespaced skills are selected automatically or invoked as `/plugin-name:skill-name`, agents may be selected or delegated to, and plugin updates require reload or a new session for some components. [Discover Claude plugins](https://code.claude.com/docs/en/discover-plugins); [Claude skills](https://code.claude.com/docs/en/slash-commands)
- **OpenAI Codex:** users browse/install through `/plugins` or `codex plugin`; native skills are available through `/skills`, `$skill-name`, prompt mention, or implicit description matching after a new session. [Codex Plugins](https://learn.chatgpt.com/docs/plugins); [Codex skills](https://learn.chatgpt.com/docs/build-skills)
- **Cursor:** users discover public plugins in Cursor Marketplace or invoke `/add-plugin`; skills are description-discovered or slash-selected, commands are explicit slash entries, and agents are selected or delegated to. [Cursor Marketplace announcement](https://cursor.com/blog/marketplace); [Cursor skills and subagents](https://cursor.com/changelog/2-4)
- **Antigravity:** desktop uses documented plugin discovery paths/UI while CLI uses `agy plugin`; skills are selected by relevance or name, and workflows become slash commands. [Antigravity plugins](https://antigravity.google/docs/plugins); [Antigravity skills](https://antigravity.google/docs/skills); [rules and workflows](https://antigravity.google/docs/rules-workflows)
- **Gemini CLI:** users install extensions from Git/local sources; extension commands are explicit, skills are progressively discovered and separately consented on activation, and preview subagents are invoked automatically or with `@name`. [Gemini extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md); [Gemini skills](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/skills.md)
- **JetBrains AI Assistant:** users add skill directories or GitHub registries through Skills Manager; supported coding agents can invoke skills automatically or by `$skill-name`, while rules, prompts, and MCP retain their own IDE discovery/invocation surfaces. [AI Assistant Skills](https://www.jetbrains.com/help/ai-assistant/agent-skills.html); [AI Assistant rules](https://www.jetbrains.com/help/ai-assistant/configure-project-rules.html)
- **Junie:** users register/browse marketplaces through `/extensions`; installed skills are relevance-discovered or explicitly requested, commands are slash-invoked, and custom subagents may be delegated to under their configured descriptions and tool sets. [Junie extensions](https://junie.jetbrains.com/docs/junie-cli-extensions.html); [Junie skills](https://junie.jetbrains.com/docs/agent-skills.html)
- **OpenCode:** users install npm plugins with `opencode plugin`; packaged skills are advertised by metadata and loaded through the permission-aware skill tool, commands are slash-invoked, and named agents come from merged configuration. [OpenCode CLI](https://opencode.ai/docs/cli/); [OpenCode skills](https://opencode.ai/docs/skills); [OpenCode commands](https://opencode.ai/docs/commands/)
- **Kilo Code:** users discover Marketplace items in the Kilo sidebar or use the Skill Issue installer; skills activate by relevance or explicit name, agents appear as primary selections or `@agent`/Task delegations, and workflows are slash commands. [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace); [Kilo subagents](https://kilo.ai/docs/customize/custom-subagents); [Kilo workflows](https://kilo.ai/docs/customize/workflows)
- **Pi:** users install packages from npm, Git, URLs, or local paths; packages expose extensions, skills, prompt templates, and themes at user or project scope. Skills are progressively advertised by name and description, loaded on demand, and can also be invoked as `/skill:name`; `pi config` enables or disables package resources. [Pi packages](https://pi.dev/docs/latest/packages/); [Pi skills](https://pi.dev/docs/latest/skills/); [Pi settings](https://pi.dev/docs/latest/settings/)

## Reusable packaging families

### 1. Native multi-component agent bundles

**Members:** GitHub Copilot, Claude Code, OpenAI Codex, Cursor, Antigravity, Gemini CLI, Junie, and Pi.

These systems accept a directory plus a native manifest or catalog entry and can co-locate multiple agent components. Their overlap is meaningful: skills, Markdown instructions, supporting resources, hooks, MCP, and sometimes agents and commands. Their schemas, discovery rules, state directories, and trust models remain different.

Skill Issue should reuse a common internal component graph and render it into each schema. It should not place several manifests in one directory and call that portability. Google explicitly documents conversion from Gemini CLI to Antigravity because paths and semantics differ, and VS Code/Copilot/Claude compatibility still has root-variable and manifest-detection differences. [Google's migration documentation](https://antigravity.google/docs/gcli-migration) and [VS Code agent plugin documentation](https://code.visualstudio.com/docs/agent-customization/agent-plugins) support generation and translation rather than format equivalence.

Within this family, bundle completeness varies:

- **Claude Code** is the broadest self-contained in-host bundle because it includes plugin `bin/`, skills, agents, commands, hooks, MCP/LSP, settings, and arbitrary files. Its `bin/` exposure is limited to Claude Code Bash calls, so a human-facing global CLI still needs its normal distribution channel. [Claude plugin reference](https://code.claude.com/docs/en/plugins-reference)
- **Copilot** is similarly broad and has strong marketplace/update coverage, but repository instructions and a normal executable remain separate. Cloud Agent also imposes a distinct Linux, non-interactive, ephemeral execution contract. [About Copilot plugins](https://docs.github.com/en/copilot/concepts/agents/about-plugins); [Copilot hooks](https://docs.github.com/en/copilot/reference/hooks-reference)
- **Codex** has a strong core plugin but a narrower documented manifest. Standalone custom-agent TOMLs, `AGENTS.md`, rules, and config live in trusted configuration layers rather than plugin roots. [Codex subagents](https://learn.chatgpt.com/docs/agent-configuration/subagents); [Codex config](https://learn.chatgpt.com/docs/config-file/config-basic)
- **Cursor** natively packages almost every agent-facing component, but editor API/UI work belongs in a separate Open VSX/VSIX extension. [Cursor schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json)
- **Antigravity** has the native package structure but lacks a complete public third-party distribution/update contract. **Gemini CLI**, despite being the lineage predecessor and now a narrower commercial surface, has the more mature extension lifecycle. [Antigravity CLI plugins](https://antigravity.google/docs/cli-plugins); [Gemini extension releasing](https://google-gemini.github.io/gemini-cli/docs/extensions/extension-releasing.html)
- **Junie** has a coherent extension bundle and marketplace, while AI Assistant does not share that package manager. [Junie extensions](https://junie.jetbrains.com/docs/junie-cli-extensions.html)
- **Pi** has a coherent package lifecycle for extensions, skills, prompts, themes, dependencies, user/project scopes, npm/Git/local sources, filtering, updates, and removal. Its package is highly complete for Pi-hosted behavior, while the cross-harness `skill-issue` executable remains an independent release. [Pi packages](https://pi.dev/docs/latest/packages/)

### 2. Runtime-code plugins with configuration composition

**Members:** OpenCode and, for runtime-only needs, Kilo Code.

These npm modules execute in the host and can contribute tools or hooks. OpenCode's current config hook makes one npm plugin capable of wiring packaged skills, agents, commands, instructions, references, and MCP into the merged configuration. Kilo's npm plugin API is strong for tools and lifecycle interception but is not documented as a declarative installer for the rest of the bundle. Both disable npm lifecycle scripts during host-managed installation, so `postinstall` cannot be the bootstrap mechanism. [OpenCode plugins](https://opencode.ai/docs/plugins/); [Kilo plugins](https://kilo.ai/docs/automate/extending/plugins)

Skill Issue should keep startup behavior side-effect-light and never copy project/global files merely because a plugin module loaded. OpenCode can use idempotent configuration injection on supported versions. Kilo should delegate full installation to the CLI or companion extension.

### 3. Skills-first and repository-composition systems

**Members:** JetBrains AI Assistant; also fallback paths for other harnesses.

Here the portable Agent Skill is the main reusable unit, while persistent guidance, project configuration, commands/workflows, MCP, secrets, and account objects use independent channels. A repository can still provide a cohesive project-scoped experience:

- AI Assistant can combine a GitHub skills registry with `.aiassistant/rules`, `.agents/skills`, and `AGENTS.md`, but Prompt Library and MCP setup remain separate documented UI/config surfaces. [AI Assistant rules](https://www.jetbrains.com/help/ai-assistant/configure-project-rules.html); [AI Assistant MCP](https://www.jetbrains.com/help/ai-assistant/mcp.html)

This family needs the strongest installer support: previewed changes, conservative merge semantics, receipts, conflict refusal, update planning, and explicit platform handoffs.

### 4. Editor-extension companions

**Members:** Cursor via Open VSX/VSIX, JetBrains via IntelliJ plugin ZIP/JAR, and Kilo via VSIX.

An editor extension is justified when Skill Issue needs UI, command-palette actions, a management pane, progress/status, host API integration, or a one-click wrapper around the installer. It is lower fit for merely distributing Markdown and scripts because native skills/plugins already own that content.

Editor extensions run with substantial host privileges and have independent registries, reviews, signing, compatibility ranges, and update lifecycles. They should call the same installer engine through a narrow protocol rather than independently reimplement bundle mutation. [VS Code extension security](https://code.visualstudio.com/docs/configure/extensions/extension-runtime-security); [JetBrains plugin approval](https://plugins.jetbrains.com/docs/marketplace/jetbrains-marketplace-approval-guidelines.html)

### 5. Account-scoped services and authorization

**Members:** MCP, connectors/apps, secrets, enterprise policy, and hosted service integrations across most harnesses.

These are never ordinary package contents. A bundle can declare or recommend them, but authentication, secret entry, tool approval, organization policy, and account installation remain explicit host-controlled actions. Skill Issue must represent these as install-plan steps with states such as `required`, `optional`, `blocked-by-policy`, `awaiting-auth`, and `verified`, instead of treating file installation as completion.

## Harness-by-harness implications

### GitHub Copilot

Build a primary OpenPlugin-compatible Copilot plugin using `.plugin/plugin.json`, namespaced skills/agents/commands/MCP IDs, minimal hooks, and optional LSP or Copilot CLI extension only where those surfaces add unique value. Put skill-owned scripts, references, and assets inside the skill directory. Publish a marketplace catalog with semantic versions and immutable release selectors, and support direct Git installation for previews. VS Code can discover Copilot CLI installations, while cloud enablement needs `.github/copilot/settings.json`. [Copilot CLI plugin reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-plugin-reference); [VS Code agent plugins](https://code.visualstudio.com/docs/agent-customization/agent-plugins); [assignment](assignments/01-github-copilot.md)

The CLI/initializer must separately install repository instructions, prompt files, or other fixed-path project configuration, and it must own the normal `skill-issue` executable. For JetBrains and other Copilot IDEs without verified agent-plugin marketplace support, materialize only the customization primitives those clients document. Treat cloud hooks, root variables, preview IDE/plugin surfaces, and collision precedence as explicit compatibility tests.

### Claude Code

Publish one `.claude-plugin/plugin.json` package through a Git-hosted marketplace. Use native skills for all reusable guidance, agents for specialized roles, plugin-scope hooks/MCP, `${CLAUDE_PLUGIN_ROOT}` for immutable resources, and `${CLAUDE_PLUGIN_DATA}` for mutable state. A thin executable under `bin/` may serve Claude's Bash tool calls. [Claude plugins](https://code.claude.com/docs/en/plugins); [Claude plugin reference](https://code.claude.com/docs/en/plugins-reference); [assignment](assignments/02-claude-code.md)

Keep the global CLI and arbitrary user/project settings outside the plugin. Use plugin `userConfig` for genuine plugin options and make state-changing skills manual by default with narrow tool declarations. Official marketplace inclusion is optional; a private/community marketplace is sufficient for launch. The CLI recommendation hint should only be added if Anthropic later accepts Skill Issue into an official marketplace because the documented hint ignores community/custom targets.

### OpenAI Codex

Publish a `.codex-plugin/plugin.json` core with native skills, optional MCP/app, minimal hooks, and presentation metadata, plus a marketplace catalog under the documented marketplace layout. Make native skills the canonical Codex workflow artifacts and keep their supporting files together. [Build plugins](https://learn.chatgpt.com/docs/build-plugins); [Build skills](https://learn.chatgpt.com/docs/build-skills); [assignment](assignments/03-openai-codex.md)

Provide `skill-issue init codex` for `.codex/agents/*.toml`, `.codex/config.toml` fragments, `.codex/rules/*.rules`, and `AGENTS.md` guidance. That initializer must preview and conservatively merge because those files own trusted project/user policy. It must not equate skill `agents/openai.yaml` with standalone agent-role TOMLs. Public plugin distribution currently supports skills/apps more clearly than hooks, and the IDE extension does not consume the full plugin surface, so public hooks and plugin-only IDE parity remain caveated.

### Cursor

Build one `.cursor-plugin/plugin.json` bundle for skills, agents, commands, rules, hooks, MCP, and all supporting files. Preserve portable skills as self-contained directories; use rules only for persistent invariants, commands for intentional launchers, agents for bounded specialist work, and hooks only for essential lifecycle behavior. [Cursor plugin schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json); [Cursor Customize](https://cursor.com/changelog/customize); [assignment](assignments/04-cursor.md)

Distribute through the Cursor Marketplace and team marketplaces, with a direct/local development route. The public plugin must satisfy Cursor's license constraints, including the documented prohibition on GPL/AGPL/LGPL components in Marketplace submissions. Publish an Open VSX/VSIX companion only if Skill Issue requires editor-native UI or APIs. The CLI must own `PATH` installation and configuration materialization because the native plugin schema has no binary or settings-registration field.

### Google Antigravity and Gemini CLI

Maintain two generated artifacts from shared sources. The Antigravity adapter should emit `plugin.json`, `skills/`, `rules/`, root `mcp_config.json`, root `hooks.json`, and CLI-specific `agents/` only where supported. The installer must target the selected documented desktop or CLI path and should install both only on explicit request. Until Google documents third-party remote sources and plugin updates completely, Skill Issue must own Antigravity upgrades and checksums. [Antigravity plugins](https://antigravity.google/docs/plugins); [Antigravity CLI plugins](https://antigravity.google/docs/cli-plugins); [assignment](assignments/05-google-antigravity-gemini-cli.md)

The Gemini CLI adapter should emit `gemini-extension.json`, `GEMINI.md`, TOML commands, skills, hooks, optional preview agents, restrictive policies, settings, and MCP entrypoints using `${extensionPath}`. Publish immutable tags and GitHub Release archives, letting Gemini handle extension updates. Keep consumer positioning accurate: Antigravity is now the primary consumer target, while Gemini CLI remains supported for enterprise and paid API-key users. Neither manifest installs the standalone CLI globally.

### JetBrains AI Assistant and Junie

Treat these as separate targets. For AI Assistant, publish a GitHub skills registry and offer a project materializer for `.agents/skills`, `.aiassistant/rules`, and `AGENTS.md`. Provide explicit setup directions for Prompt Library and MCP unless a future public API makes automation supportable. Current Skills Manager support covers Claude Agent and Codex rather than Junie. [AI Assistant Skills](https://www.jetbrains.com/help/ai-assistant/agent-skills.html); [AI Assistant agent instructions](https://www.jetbrains.com/help/ai-assistant/configure-agent-behavior.html); [assignment](assignments/06-jetbrains-ai-junie.md)

For Junie, publish a native marketplace extension containing skills, agents, commands, guidelines, and MCP. Support project and user scopes, explicit updates, and commit-pinned provenance. The standalone CLI remains separate. A generic JetBrains Marketplace plugin is an optional management UI that can invoke the CLI and materialize documented files; it must not depend on undocumented AI Assistant or Junie internals. Its signing/review lifecycle governs only the IDE plugin, not content later fetched from Skill Issue or a Junie marketplace.

### OpenCode

Publish a scoped, versioned npm server plugin with `engines.opencode`. Return namespaced tools and hooks directly; use the config hook to idempotently append packaged skill paths, instruction paths, references, agents, commands, and optional disabled-by-default MCP. Prefer plugin tools over MCP for in-package functions. [OpenCode plugin API](https://opencode.ai/docs/plugins/); [OpenCode configuration](https://opencode.ai/docs/config/); [assignment](assignments/07-opencode.md)

Document deterministic installation and upgrades with explicit versions and `--force`. Do not rely on npm lifecycle scripts or on unpinned `latest` refreshing immediately. The standalone CLI needs a normal npm/binary installation or `npx`. A version-gated fallback may materialize `.opencode` files for older releases, and uninstall must remove only Skill Issue-owned config entries while preserving the user's merged configuration.

### Kilo Code

Publish useful standalone skills, agents, and MCP entries to Kilo Marketplace, but do not describe an umbrella item as a complete dependency installer until Kilo documents requirement resolution. Keep full installation authority in `skill-issue install kilo`, which materializes namespaced skills, agents, commands, and rules; structurally merges `kilo.jsonc`; records ownership; and supports update/doctor/uninstall. [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace); [Kilo skills](https://kilo.ai/docs/customize/skills); [assignment](assignments/08-kilo-code.md)

Publish an npm Kilo plugin only for genuine runtime tools/hooks and an MCP artifact only for protocol-level tools. If one-click editor management is a product requirement, build a target-specific companion VSIX dependent on `kilocode.kilo-code`, publish it to both Visual Studio Marketplace and Open VSX, and reuse the CLI installer engine. The VSIX route is the closest full editor bundle, but it remains a separate editor-extension channel rather than a Kilo agent-package standard.

### Pi

Publish one Pi Package with a scoped npm identity and a `package.json` `pi` manifest that includes the Skill Issue skills, any Pi-specific TypeScript extension, prompt templates, and themes. Keep skill-owned scripts, references, assets, and templates inside their skill directories. Use the extension only for behavior that needs Pi APIs, tools, commands, events, or UI; keep ordinary reusable guidance in Agent Skills. Pi can install the package from npm, Git, URL, or a local path and can manage it at user or project scope. [Pi packages](https://pi.dev/docs/latest/packages/); [Pi extensions](https://pi.dev/docs/latest/extensions/); [Pi skills](https://pi.dev/docs/latest/skills/)

Use `pi install`, `pi list`, `pi config`, `pi update --extensions`, and `pi remove` as the native lifecycle. Pin release versions or Git refs deliberately, because pinned packages do not advance automatically to newer versions. Treat project package installation as a trust-gated action and surface that Pi packages run with full system access. Distribute the standalone `skill-issue` CLI independently; the Pi extension may detect or invoke it after `doctor` verifies the executable, but package loading must not silently place a global binary on `PATH` or mutate unrelated settings. [Pi package management](https://pi.dev/docs/latest/packages/); [Pi project trust and settings](https://pi.dev/docs/latest/settings/)

## Proposed Skill Issue architecture

### Canonical bundle model

Create one source-of-truth bundle manifest that is richer than any host schema and is never installed directly by a harness. It should record:

- bundle ID, display name, version, release channel, publisher, license, source revision, and checksums;
- minimum/maximum tested harness versions and per-component compatibility;
- skills, agents, commands/workflows, persistent instructions/rules, hooks, tool plugins, MCP/apps, references, assets, templates, and configuration fragments;
- standalone CLI version and supported platform artifacts;
- component dependencies and whether they are required, optional, or conditional;
- install scope choices: project, user, team, enterprise, account, or cloud;
- requested filesystem, shell, network, MCP, secret, and account capabilities;
- platform authorization steps that cannot be automated;
- ownership policy, conflict behavior, update migrations, rollback metadata, and uninstall rules.

The manifest should distinguish **content inclusion** from **host activation**. A custom-agent file carried inside a Codex plugin payload, for example, is included but inactive until installed into `.codex/agents/`. The same principle applies to repository guidance, external CLIs, MCP credentials, and account connectors.

### Portable content tree

Use the Agent Skills directory model as the stable content center:

```text
bundle/
  skills/<skill-id>/
    SKILL.md
    scripts/
    references/
    assets/
    templates/
  agents/
  commands/
  instructions/
  hooks/
  mcp/
  shared/
  adapters/
  manifest.json
```

Keep skill-owned files inside the owning skill. Put only genuinely shared corpora or runtime libraries in `shared/`. Render host-specific manifests, frontmatter, naming, paths, and configuration into release staging directories; do not make source skills depend on a checked-out repository path.

### Adapter contract

Each adapter should declare capabilities rather than silently dropping unsupported components. A generated compatibility report should classify every canonical component as:

- `native` — directly represented and host-managed;
- `native-with-caveat` — supported but preview, surface-limited, or incompletely documented;
- `materialized` — installed by the Skill Issue CLI into a documented host path;
- `external` — separate CLI, MCP service, connector, secret, or editor extension;
- `omitted` — intentionally unavailable on that surface.

Adapters should namespace all public identifiers, validate native schemas, and generate a `doctor` expectation set. A release must fail if a required component becomes `omitted` without an explicit product decision.

### CLI and installer responsibilities

The standalone `skill-issue` CLI should be the cross-harness lifecycle owner. At minimum:

```text
skill-issue install <harness> [--project|--user|--global]
skill-issue update <harness>
skill-issue status <harness>
skill-issue doctor <harness>
skill-issue uninstall <harness>
skill-issue init <harness>
```

It should:

1. detect the harness and supported version;
2. select the correct native artifact and scope;
3. show the component inventory, executable surfaces, permissions, destinations, and platform handoffs;
4. preview file/config changes and refuse unmanaged conflicts by default;
5. install the native package through the host when a supported package manager exists;
6. materialize only components the native package cannot activate;
7. install or verify the standalone executable through an explicit conventional channel;
8. record an ownership receipt containing hashes and merge state;
9. verify discovery and actual execution, not only file presence;
10. update atomically where possible and preserve user-modified files;
11. uninstall only owned content and leave account authorizations or secrets untouched unless the user explicitly requests revocation.

### Runtime and MCP boundary

Prefer direct bundled scripts or native plugin tools for deterministic local behavior. Use MCP when the capability needs cross-harness reuse, remote data, structured tools, OAuth, or a long-running service. A local MCP server must be a dependency-complete executable or produce an actionable missing-runtime error. A hosted server must use accurate tool annotations, least-privilege authorization, privacy documentation, and explicit account connection.

The bundle must never embed secrets. It should carry secret names, purpose, validation rules, and setup links. MCP installation and tool approval must be verified separately from skill/plugin discovery.

### Distribution workflow

1. **Author once:** update canonical skills, agents, instructions, scripts, references, assets, and bundle metadata.
2. **Render:** generate each native package in an isolated staging tree.
3. **Validate:** run native schema/manifest validators, frontmatter checks, path-containment checks, license scans, platform binary checks, and dependency completeness checks.
4. **Package:** produce Git marketplace repositories, npm packages, release archives, VSIX files, and project templates as required.
5. **Sign and attest:** publish checksums, source revision, SBOM/provenance where executable code is included, and registry signing supported by the channel.
6. **Publish:** update each independent marketplace/registry/catalog without assuming simultaneous availability.
7. **End-to-end verify:** install from the remote release source into a clean supported environment; verify new-session/reload requirements, skill discovery, explicit invocation, CLI version, agent registration, MCP startup/tool calls, hook trust, and policy blockers.
8. **Publish compatibility state:** report native, partial, preview, and unsupported components for every surface and release.

## Conditional alternatives

### Skills-only distribution

If Skill Issue's initial product can operate without hooks, custom agents, persistent rules, MCP, or a standalone CLI, publish only portable Agent Skills. This is the fastest cross-harness path and works directly or through host-specific skill registries in nearly every target. It is lower effort but must be marketed as a skills package rather than the complete Skill Issue system.

### MCP-first tool service

If most value resides in a shared evaluator/refiner service rather than local scripts and project files, publish thin skills plus a hosted MCP server. This reduces per-host executable packaging and centralizes updates. It adds authentication, hosting, privacy, network availability, platform review, and account-policy dependencies, and it still does not install project guidance or a local CLI.

### Project-template-first launch

For team repositories and harness fallback paths without a complete native bundle, a ready-to-use Git project/template may be the best first-run experience. It can include skills, guidance, scripts, assets, and runtime configuration together. This is appropriate for new projects; existing-project adoption and upgrades still need the CLI's merge/ownership model.

### Editor companion extensions

Build Cursor/Open VSX, Kilo/VSIX, or JetBrains IDE companions only if users need graphical lifecycle management or editor-native capabilities. A companion can provide install/update/status/doctor/uninstall commands and preview diffs. It should reuse the standalone installer engine and remain optional.

### Split native plugins

Start with one native plugin per harness. Split components only when independent release cadence, platform binaries, organization policy, trust, or optional hosted services justify separate installation. Claude dependencies and Kilo requirements can express some relationships, but more packages create more update, compatibility, and consent edges.

## Rejected or lower-fit interpretations

### A universal OpenPlugin directory

The shared shape across Copilot, Claude Code, and some compatibility loaders is useful source inspiration, not a nine-harness standard. Codex, Cursor, Antigravity, Gemini, Junie, OpenCode, Kilo, and Pi have different schemas and lifecycle owners. Even Copilot/Claude compatibility differs in manifest detection, root variables, components, and trust.

### Agent Skills as a complete package manager

Skills are the strongest portable content unit, but skill installers generally do not install global executables, host configuration, account integrations, secrets, MCP authorization, editor UI, or arbitrary project policy. Treating “skills installed” as “bundle installed” would make diagnostics and upgrades misleading.

### Silent bootstrap from plugin startup or npm lifecycle scripts

OpenCode and Kilo disable npm lifecycle scripts in host-managed installs, while plugin modules themselves are trusted executable code. Pi installs package dependencies and treats packages as full-access code, with project-local resources gated by project trust. Other hosts separately gate project trust, hooks, commands, and tool permissions. File mutation on first load would create poor consent, conflict, and uninstall semantics. All cross-scope changes should be explicit installer actions.

### A plugin-installed global CLI

No reviewed native manifest establishes a portable cross-harness field that registers a normal standalone `skill-issue` command in the user's shell. Claude's `bin/` applies inside Claude Bash calls; Copilot extensions are session integrations; Pi and other npm-backed package systems manage host resources or dependencies rather than a verified cross-platform global binary installation; other bundles can merely carry executable files. The CLI needs conventional independent distribution.

### A generic JetBrains plugin as the AI Assistant/Junie package

An IntelliJ plugin can provide UI and invoke the Skill Issue CLI, but the reviewed public SDK does not establish stable APIs that atomically register AI Assistant skills, Prompt Library entries, MCP, or Junie extension marketplaces. AI Assistant and Junie must retain separate adapters.

### Treating Antigravity as renamed Gemini CLI

Google documents Antigravity CLI as a new product with migration from Gemini CLI. Their manifests, paths, MCP keys, commands/workflows, policies, and distribution maturity differ. One source repository can produce both artifacts; one package identity should not be claimed.

### Treating repository cloning as universal installation

Git is a strong transport and project-scoped composition mechanism. It does not activate user-scope plugins, install account objects, authorize MCP, place a global CLI on `PATH`, or manage safe upgrades inside arbitrary existing repositories.

## Unresolved blockers and unsupported claims

These points remain unsupported or materially caveated in the supplied evidence and should appear in Skill Issue's compatibility matrix until first-party contracts change:

- **Copilot:** native agent-plugin marketplace support outside Copilot CLI, cloud agent, and supported VS Code was not established; cloud plugin-hook materialization and cross-format root variables require end-to-end validation. [Copilot assignment](assignments/01-github-copilot.md)
- **Claude Code:** plugin `bin/` is not documented as global shell `PATH`; arbitrary settings cannot be packaged; official marketplace inclusion and CLI recommendation hints are external decisions. [Claude assignment](assignments/02-claude-code.md)
- **Codex:** plugin-native installation of standalone custom-agent TOMLs is not supported by the reviewed loader; public-directory acceptance of hooks is unclear; the IDE extension lacks full plugin support; no plugin path installs general config, rules, `AGENTS.md`, or a CLI. [Codex assignment](assignments/03-openai-codex.md)
- **Cursor:** public plugin refresh timing, rollback/diff behavior, CLI-only third-party plugin installation, and full component parity across IDE, CLI, cloud agents, and automations were not established. [Cursor assignment](assignments/04-cursor.md)
- **Antigravity:** public third-party marketplace submission, remote install grammar, signatures, versions, update command, and automatic update policy were not documented; desktop and CLI global install paths differ without a documented sync rule. [Google assignment](assignments/05-google-antigravity-gemini-cli.md)
- **Gemini CLI:** extension subagents remain preview; installing a Skill Issue extension does not install the separate IDE companion or global CLI. [Google assignment](assignments/05-google-antigravity-gemini-cli.md)
- **AI Assistant:** no public atomic package or stable automation API was established for Skills Manager registries, Prompt Library, MCP, rules, and agent instructions together. [JetBrains assignment](assignments/06-jetbrains-ai-junie.md)
- **Junie:** custom marketplace signing, automatic update, checksum enforcement, and third-party review are not documented; manual update is the supported path found. [JetBrains assignment](assignments/06-jetbrains-ai-junie.md)
- **OpenCode:** config-hook bundle injection is source-backed but not fully explained in public plugin guidance; no first-party marketplace, documented per-plugin remove command, capability sandbox, lifecycle provisioning, or global `bin` exposure was found. [OpenCode assignment](assignments/07-opencode.md)
- **Kilo:** client behavior for Marketplace `requirements` was not established; Marketplace availability documentation is internally inconsistent; npm plugin fields do not declaratively install the wider bundle. [Kilo assignment](assignments/08-kilo-code.md)
- **Pi:** Pi Packages cover extensions, skills, prompt templates, themes, dependencies, scoped installation, updates, filtering, and removal, but they execute with full system access and do not establish a verified native path for installing the independent cross-platform `skill-issue` binary. Pinned npm versions and Git refs require an explicit version/ref change to advance, and project package activation remains trust-gated. [Pi packages](https://pi.dev/docs/latest/packages/); [Pi settings](https://pi.dev/docs/latest/settings/)

## Final recommendation

Implement **one canonical bundle, eleven surface adapters for the nine entries, and one lifecycle CLI**. The two compound entries—Google Antigravity/Gemini CLI and JetBrains AI Assistant/Junie—each require separate adapters. Prioritize native artifacts in this order:

1. portable Agent Skills and canonical component metadata;
2. full native bundles for Claude Code, Copilot, Codex, Cursor, Gemini CLI, Junie, and Pi;
3. the OpenCode npm adapter;
4. the Antigravity adapter plus Skill Issue-owned updater;
5. the Kilo CLI/materializer plus native Marketplace entries;
6. the AI Assistant registry/materializer;
7. optional editor companions only after a concrete UI requirement exists.

Every release should state what the native package installs, what the CLI materializes, what remains an external authorization, and what is unsupported. That explicit boundary is the practical substitute for a universal plugin standard and is the strongest evidence-backed route to a functional Skill Issue bundle across all nine selected harness entries.

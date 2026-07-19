# Kilo Code Packaging

## Assignment

**Goal:** Determine how current Kilo Code packaging and distribution work, and what Skill Issue must implement to deliver skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, and supporting files as a coherent Kilo experience.

**Scope:** Current Kilo Code 7.x VS Code/CLI architecture; VS Code Marketplace, Open VSX, and VSIX distribution; Kilo Marketplace items; skills; current agents/modes and subagents; rules and instruction files; commands/workflows; MCP; npm plugins; configuration and scope precedence; installation, updates, discovery, invocation, trust, permissions, and review constraints; concrete Skill Issue packaging options.

**Exclusions:** Legacy Kilo 5.x packaging except where current Kilo documents migration behavior; JetBrains-specific packaging; local machine configuration; implementation work; claims based only on third-party summaries.

## Sources

- [Kilo Code installation](https://kilo.ai/docs/getting-started/installing), current docs inspected 2026-07-19.
- [Kilo Code settings and shared configuration](https://kilo.ai/docs/getting-started/settings), current docs inspected 2026-07-19.
- [Kilo Marketplace documentation](https://kilo.ai/docs/customize/marketplace), current docs inspected 2026-07-19.
- [Kilo Marketplace repository](https://github.com/Kilo-Org/kilo-marketplace), [contribution requirements](https://github.com/Kilo-Org/kilo-marketplace/blob/main/CONTRIBUTING.md), and [review guidance](https://github.com/Kilo-Org/kilo-marketplace/blob/main/REVIEW.md), `main` inspected 2026-07-19.
- [Kilo skills documentation](https://kilo.ai/docs/customize/skills), current docs inspected 2026-07-19.
- [Kilo custom modes](https://kilo.ai/docs/customize/custom-modes), [custom subagents](https://kilo.ai/docs/customize/custom-subagents), and [agent permissions](https://kilo.ai/docs/customize/agent-permissions), current docs inspected 2026-07-19.
- [Kilo workflows/slash commands](https://kilo.ai/docs/customize/workflows), [custom instructions](https://kilo.ai/docs/customize/custom-instructions), and [custom rules](https://kilo.ai/docs/customize/custom-rules), current docs inspected 2026-07-19.
- [Kilo plugins](https://kilo.ai/docs/automate/extending/plugins), [MCP configuration](https://kilo.ai/docs/automate/mcp/using-in-kilo-code), and [sandboxing](https://kilo.ai/docs/getting-started/settings/sandboxing), current docs inspected 2026-07-19.
- [Kilo CLI reference](https://kilo.ai/docs/code-with-ai/platforms/cli), current docs inspected 2026-07-19.
- [Kilo monorepo](https://github.com/Kilo-Org/kilocode), [`packages/kilo-vscode/package.json`](https://github.com/Kilo-Org/kilocode/blob/main/packages/kilo-vscode/package.json), and [release process](https://github.com/Kilo-Org/kilocode/blob/main/RELEASING.md), `main` and release `7.4.11` inspected 2026-07-19.
- [Kilo Code on Open VSX](https://open-vsx.org/extension/kilocode/kilo-code) and [Open VSX API record for the latest build](https://open-vsx.org/api/kilocode/kilo-code/latest), version `7.4.11`, published 2026-07-16.
- [Visual Studio Marketplace API record for `kilocode.kilo-code`](https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery), queried 2026-07-19; current result was version `7.4.11` with platform-specific builds.
- [VS Code extension manifest reference](https://code.visualstudio.com/api/references/extension-manifest), [extension publishing](https://code.visualstudio.com/api/working-with-extensions/publishing-extension), [extension runtime security](https://code.visualstudio.com/docs/configure/extensions/extension-runtime-security), [Workspace Trust extension guide](https://code.visualstudio.com/api/extension-guides/workspace-trust), and [extension management](https://code.visualstudio.com/docs/configure/extensions/extension-marketplace), current official docs inspected 2026-07-19.
- [Open VSX publishing guide](https://github.com/EclipseFdn/open-vsx.org/wiki/Publishing-Extensions), revised 2026-02-26 and inspected 2026-07-19.

## Findings

### Finding 1: Current Kilo VS Code is a CLI-backed, platform-specific VSIX

Current Kilo is distributed as the `kilocode.kilo-code` editor extension, but the extension delegates agent behavior to the Kilo CLI runtime. Kilo's release process builds a CLI binary first, then produces separate VSIX files for Linux, Alpine, macOS, and Windows architectures, embedding the matching CLI binary in each VSIX. Kilo publishes the VSIX family to the Visual Studio Marketplace and GitHub Releases; Kilo's installation documentation also directs VS Code-compatible editors to Open VSX. The two registry APIs returned current version `7.4.11` on 2026-07-19.

**Evidence:** Kilo's [installation guide](https://kilo.ai/docs/getting-started/installing) says the current extension is built on the CLI and documents VS Marketplace, Open VSX, and direct VSIX installation. The [release document](https://github.com/Kilo-Org/kilocode/blob/main/RELEASING.md) says each platform VSIX bundles its platform-specific CLI binary. The current [extension manifest](https://github.com/Kilo-Org/kilocode/blob/main/packages/kilo-vscode/package.json) declares publisher `kilocode`, name `kilo-code`, VS Code engine `^1.105.1`, and `dist/extension.js`; the [Open VSX API](https://open-vsx.org/api/kilocode/kilo-code/latest) reports signed platform-specific `7.4.11` assets.

**Implication:** A Skill Issue companion extension can technically ship Markdown, scripts, assets, and platform-specific CLI binaries inside one VSIX, following a pattern Kilo itself proves. If Skill Issue includes native executables, it should build target-specific VSIX files rather than assume one portable archive. The companion must target current VS Code compatibility deliberately and should be tested against both Microsoft VS Code and Open VSX consumers.

### Finding 2: Kilo Marketplace is native discovery for atomic agents, skills, and MCP servers

Kilo has a first-party Marketplace surface in the Kilo sidebar. Its items are Kilo configuration or instruction files rather than VS Code extensions. An installation is explicitly project-scoped or global-scoped and creates one agent file, one skill directory, or one MCP configuration entry. Install and removal show the destination, preserve unrelated MCP settings, and reload affected Kilo configuration.

**Evidence:** The [Kilo Marketplace documentation](https://kilo.ai/docs/customize/marketplace) lists exactly three installable types: Agent, Skill, and MCP server. It maps them to `.kilo/agents/<name>.md`, `.kilo/skills/<name>/`, and `.kilo/kilo.json` for projects, and `~/.config/kilo/agents/<name>.md`, `~/.kilo/skills/<name>/`, and `~/.config/kilo/kilo.json` globally. It states that the install dialog shows the destination and that removal is tracked separately per scope.

**Implication:** Skill Issue should submit its reusable skills, any specialized agent, and any MCP server as native Marketplace entries for discovery and low-friction installation. Commands, rules, npm plugins, a standalone CLI, and arbitrary configuration files are outside the documented Marketplace item types, so the native Marketplace alone cannot install Skill Issue's complete bundle.

### Finding 3: Marketplace requirements can describe only a partial bundle graph

Marketplace agent, skill, and MCP definitions may declare direct requirements on other marketplace skills, MCPs, and VS Code extensions. Requirements cannot name agents, commands, rules, npm plugins, CLIs, or arbitrary supporting packages. Alternative dependency groups are unsupported, and dependency graphs are not expanded or cycle-checked by marketplace generation.

**Evidence:** The Kilo Marketplace [contribution requirements](https://github.com/Kilo-Org/kilo-marketplace/blob/main/CONTRIBUTING.md#marketplace-requirements) allow exactly `skills`, `vscode_extensions`, and `mcps` under `requirements`. They require exact IDs, reject self-dependencies, preserve direct requirements in generated items, and state that dependencies between resources are not expanded or checked for cycles.

**Implication:** Skill Issue can publish an umbrella skill or agent that declares direct Skill Issue skills, its MCP, and a companion VS Code extension as requirements. That improves composition but is not a complete bundle primitive. The current public documentation does not state whether the UI automatically installs every requirement, prompts for them, or merely displays them; automatic dependency installation is therefore caveated and must not be the only installer path.

### Finding 4: A Kilo skill is the native package for guidance plus scripts, references, and assets

A skill is a directory whose only required file is `SKILL.md`; Kilo explicitly permits `scripts/`, `references/`, and `assets/` alongside it. Kilo scans metadata at session start, puts the name and description in agent context, and loads full instructions on demand through the `skill` tool. Explicitly naming the skill triggers it; otherwise the model selects it from its description. Current Kilo uses one shared skill pool rather than separate mode-specific directories.

**Evidence:** The [skills documentation](https://kilo.ai/docs/customize/skills) defines the skill lifecycle, frontmatter, invocation behavior, and optional bundled resources. It supports global `~/.kilo/skills/`, project `.kilo/skills/`, compatibility `.agents/skills/`, extra `skills.paths`, and remote `skills.urls` backed by an `index.json`. Project skills take precedence over global skills with the same name, and `/reload` resynchronizes skills without a new session.

**Implication:** Skill Issue should keep each coherent workflow self-contained as an Agent Skill, with executable helpers under `scripts/`, deep guidance under `references/`, and templates or static resources under `assets/`. The standalone installer should preserve directory names exactly, validate required frontmatter, and invoke `/reload` or tell the user to start a new session. A hosted `skills.urls` index is a useful optional update/distribution channel for skills only, but it does not distribute agents, commands, rules, plugins, or the CLI.

### Finding 5: Current modes are Markdown agents with explicit invocation and permission scopes

In current Kilo, custom modes and custom agents converge on Markdown agent files. An agent can be `primary`, `subagent`, or `all`; primary agents appear in the selector, subagents are invoked through the Task tool or `@agent-name`, and descriptions drive automatic delegation. Agent definitions can pin models, prompts, step limits, and ordered `allow`, `ask`, or `deny` permissions. Project agent files override global and built-in definitions at higher precedence.

**Evidence:** The [custom modes documentation](https://kilo.ai/docs/customize/custom-modes) states that the current CLI-backed VS Code extension uses agent Markdown files. The [custom subagents documentation](https://kilo.ai/docs/customize/custom-subagents) gives global `~/.config/kilo/agents/` and project `.kilo/agents/` locations, invocation methods, configuration precedence, and the `primary`/`subagent`/`all` modes. [Agent Permissions](https://kilo.ai/docs/customize/agent-permissions) defines approval outcomes and ordered permission rules.

**Implication:** Skill Issue should model reusable personas as current `.kilo/agents/*.md` files, not legacy `.kilocodemodes` or `custom_modes.yaml`. It should use restricted permissions by default and reserve full write or shell access for agents that require it. Marketplace installation can deliver individual agents; the fallback CLI must install and reconcile the full agent set.

### Finding 6: Workflows, rules, and instructions are file-native but outside Marketplace packaging

Current workflows are Markdown slash commands. Project commands live in `.kilo/commands/`, global commands in `~/.config/kilo/commands/`, and `/filename` invokes them. Command frontmatter can select an agent or model and can run the command as a subtask. Project and global instructions are supplied through `AGENTS.md`, agent prompts, `.kilo/rules/`, or the `instructions` array in `kilo.jsonc`; the latter accepts paths, globs, and URLs.

**Evidence:** The [workflows documentation](https://kilo.ai/docs/customize/workflows) defines command locations, filename-based slash invocation, frontmatter, and access to built-in and MCP tools. The [custom instructions documentation](https://kilo.ai/docs/customize/custom-instructions) identifies root and per-directory `AGENTS.md`, global `~/.config/kilo/AGENTS.md`, and `instructions` paths/globs/URLs. [Custom rules](https://kilo.ai/docs/customize/custom-rules) documents project and global rule behavior and project-over-global precedence.

**Implication:** Skill Issue's bundle needs explicit installation logic for `.kilo/commands/`, `.kilo/rules/`, and any `kilo.jsonc` instruction references. It should avoid overwriting a user's existing `AGENTS.md`; a namespaced rule file plus a merged `instructions` entry is safer and reversible. These components should be independently discoverable and removable by the Skill Issue installer because Kilo Marketplace does not document them as installable item types.

### Finding 7: Kilo npm plugins are runtime extensions, not a general declarative bundle format

Kilo plugins are JavaScript or TypeScript modules that run in both CLI and VS Code. They can register tools, intercept tool calls, subscribe to events, register providers, modify requests, and inject shell environment. Plugins load from config, global/project plugin directories, or `kilo plugin <npm-package>`. Bare npm specifiers track `latest` when the cache becomes stale; pinned versions remain pinned. Kilo disables npm lifecycle scripts, supports an engine range, and can disable all external plugins with `KILO_PURE=1`.

**Evidence:** The [plugins documentation](https://kilo.ai/docs/automate/extending/plugins) defines the supported hooks, npm and local installation paths, `kilo plugin`, caching/update semantics, disabled install/postinstall scripts, load order, `KILO_PURE`, `exports["./server"]`/`exports["./tui"]`, and the documented `engines.opencode` compatibility range.

**Implication:** Skill Issue should publish a Kilo npm plugin only for behavior that genuinely needs runtime hooks or custom tools. The plugin package can carry supporting module files, but Kilo does not document plugin manifest fields that declaratively install skills, agents, commands, rules, or a standalone executable. Lifecycle scripts cannot bootstrap those files. Having plugin code silently copy them at session startup would create difficult consent, conflict, update, and uninstall semantics; the standalone installer or companion extension should own installation instead.

### Finding 8: MCP is suitable for external tools, with separate configuration and trust boundaries

Kilo supports project and global MCP configuration for local child processes and remote servers, including OAuth-capable remote connections. A Marketplace MCP installation adds configuration but does not auto-approve its tools. Tool calls follow Kilo permissions, while server lifecycle and plugin hooks remain trusted host integrations outside important parts of the session sandbox.

**Evidence:** [Using MCP in Kilo](https://kilo.ai/docs/automate/mcp/using-in-kilo-code) documents project/global config, local and remote transports, enablement, timeouts, and OAuth. The [Marketplace documentation](https://kilo.ai/docs/customize/marketplace#mcp-security-and-permissions) says local servers execute child processes, remote servers contact external services, and installation does not approve all calls. [Sandboxing](https://kilo.ai/docs/getting-started/settings/sandboxing) states that plugin hooks and MCP lifecycle are trusted integrations and advises enabling only trusted plugins and MCP servers.

**Implication:** If Skill Issue exposes durable tools or state over MCP, it should publish a separate Marketplace MCP definition that launches a versioned npm package with `npx` or connects to a documented HTTPS endpoint. It must use environment placeholders for credentials, declare prerequisites, and provide conservative permission guidance. MCP should complement skills and commands rather than serve as the installer for non-MCP files.

### Finding 9: A companion VS Code extension can provide one editor-level install experience

VS Code permits a single extension to combine multiple contributions and bundle arbitrary runtime files. It also supports `extensionDependencies` for functional dependencies and `extensionPack` only for independently manageable recommendations. Extensions have the same host permissions as VS Code, and first installation from a third-party publisher requires publisher trust. Marketplace publication adds malware, dynamic-behavior, secret, and signature checks; Workspace Trust can disable or limit an extension in restricted workspaces when declared or enforced by its code.

**Evidence:** The [extension manifest reference](https://code.visualstudio.com/api/references/extension-manifest) documents combined contributions, packaged files, `extensionDependencies`, `extensionPack`, and workspace capabilities, and explicitly says packs should not represent functional dependencies. [Extension runtime security](https://code.visualstudio.com/docs/configure/extensions/extension-runtime-security) documents editor-equivalent permissions, publisher trust, malware and dynamic scanning, secret scanning, signature verification, and blocklisting. The [Workspace Trust guide](https://code.visualstudio.com/api/extension-guides/workspace-trust) explains restricted-mode declarations and runtime checks.

**Implication:** A Skill Issue extension can declare `kilocode.kilo-code` as an `extensionDependency`, embed the Skill Issue CLI and bundle payload, and contribute explicit commands such as **Install Skill Issue for This Project**, **Install Globally**, **Update**, **Doctor**, and **Uninstall**. Installation should be user-initiated and preview destinations and diffs rather than modifying global Kilo configuration silently on activation. This is the strongest route to a single editor-facing experience across the full bundle.

### Finding 10: Publishing the companion requires two extension registries or direct VSIX fallback

The Visual Studio Marketplace uses `vsce` to package and publish VSIX files under a registered publisher. Open VSX accepts VSIX packages produced by `vsce`/`ovsx`, requires an Eclipse account and Publisher Agreement, requires a namespace and token, and scans submissions for secrets, blocklisted files, and namespace similarity. Kilo itself is currently present in both registries, so a Skill Issue dependency can resolve in both ecosystems if Skill Issue is also published to both.

**Evidence:** [VS Code publishing](https://code.visualstudio.com/api/working-with-extensions/publishing-extension) documents `vsce package`, `vsce publish`, publisher registration, HTTPS/image restrictions, and Marketplace presentation. The [Open VSX publishing guide](https://github.com/EclipseFdn/open-vsx.org/wiki/Publishing-Extensions) documents the Publisher Agreement, namespace, token, `ovsx publish`, and automatic scans. Kilo's [Open VSX record](https://open-vsx.org/extension/kilocode/kilo-code) and [installation guide](https://kilo.ai/docs/getting-started/installing) confirm current availability outside Microsoft's marketplace.

**Implication:** Skill Issue should automate one build matrix and publish the resulting VSIX files independently to the Visual Studio Marketplace and Open VSX, with GitHub Release assets as an auditable manual fallback. A Marketplace-only release would exclude VSCodium and other Open VSX-based editors; a direct VSIX-only release would lose normal discovery and automatic updates.

### Finding 11: Marketplace review is curated and examines the entire executable payload

Kilo Marketplace submissions are reviewed as packages, including scripts, hooks, references, assets, templates, archives, binaries, licenses, provenance, portability, destructive behavior, secrets, subprocesses, and activation quality. Contributed skills must be hosted in a canonical public repository, carry source and license metadata, solve a real tested use case, and arrive through a pull request. External skill sources are periodically synchronized into the marketplace index.

**Evidence:** The Marketplace [contributing guide](https://github.com/Kilo-Org/kilo-marketplace/blob/main/CONTRIBUTING.md) requires a real problem, documentation, examples, Kilo testing, safe destructive behavior, a public canonical source, source/license metadata, and a PR. The [review guidance](https://github.com/Kilo-Org/kilo-marketplace/blob/main/REVIEW.md) requires inspection of all package files and explicitly checks provenance, resources, commands, portability, secrets, subprocess safety, consent, and activation quality.

**Implication:** Skill Issue should keep a public canonical repository, explicit licenses covering every bundled file, deterministic cross-platform tests for scripts and installers, precise skill descriptions, safe confirmation gates, and no opaque binary without reviewable provenance. Marketplace entries should remain small and independently useful even if the complete bundle is installed through the CLI or companion extension.

### Finding 12: The practical Skill Issue bundle is layered, with one authoritative installer

A complete functional experience is possible, but no single documented Kilo-native item type spans every requested component. The most robust architecture is layered: native Kilo Marketplace entries for discovery, Agent Skills as the content unit, Markdown agents and slash commands as Kilo-native interaction surfaces, an optional npm Kilo plugin or MCP package for runtime tools, and one standalone Skill Issue CLI as the authority for full install/update/uninstall. A companion VS Code extension can wrap that CLI to provide the single-click editor experience.

**Evidence:** Kilo Marketplace installs only agents, skills, and MCP entries ([Marketplace docs](https://kilo.ai/docs/customize/marketplace)); skills can carry scripts, references, and assets ([skills docs](https://kilo.ai/docs/customize/skills)); commands and agents live in separate directories ([workflows](https://kilo.ai/docs/customize/workflows), [subagents](https://kilo.ai/docs/customize/custom-subagents)); plugins are separate npm runtime modules ([plugins](https://kilo.ai/docs/automate/extending/plugins)); and VSIX packages can combine contributions and files ([VS Code manifest](https://code.visualstudio.com/api/references/extension-manifest)).

**Implication:** Skill Issue should implement the following concrete Kilo adapter:

1. A canonical, versioned bundle manifest listing owned skills, agents, commands, rules/instructions, plugin/MCP entries, CLI version, checksums, scope, and minimum compatible Kilo/VS Code versions.
2. `skill-issue install kilo --project|--global`, `update`, `doctor`, and `uninstall` operations that copy namespaced files, merge JSONC structurally, preserve comments where possible, preview changes, back up overwritten owned files, refuse unmanaged conflicts, and maintain an ownership receipt.
3. Skills installed under `.kilo/skills/<id>/` or `~/.kilo/skills/<id>/`; agents under `.kilo/agents/` or `~/.config/kilo/agents/`; commands under `.kilo/commands/` or `~/.config/kilo/commands/`; rules in a namespaced directory referenced from `kilo.jsonc`; and secrets represented only as environment references.
4. Native Kilo Marketplace submissions for each useful skill, any agent, and any MCP, plus direct `requirements` only for eligible dependencies. An umbrella item may guide installation, but should not claim to install unsupported component types.
5. An optional `@skill-issue/kilo-plugin` for custom tools/hooks and/or `@skill-issue/mcp` for external tools, each independently versioned and permission-scoped.
6. A companion `skill-issue.kilo-bundle` VS Code extension, dependent on `kilocode.kilo-code`, that embeds or invokes the same installer engine and exposes explicit consent-based lifecycle commands. Publish it to both Visual Studio Marketplace and Open VSX, with matching GitHub Release VSIX assets.

## Notes

- Kilo's current docs are internally inconsistent about Marketplace availability: the dedicated [Marketplace page](https://kilo.ai/docs/customize/marketplace) documents a live sidebar installer, while portions of the [skills page](https://kilo.ai/docs/customize/skills) still say the new platform has no marketplace UI. The dedicated Marketplace page and current marketplace repository were treated as the newer evidence; exact rollout availability in every Kilo client build remains caveated.
- The Marketplace contribution schema preserves `requirements`, but public docs inspected here do not define the client behavior for resolving or installing those requirements. Dependency auto-installation is unsupported as a confirmed claim.
- Kilo's release document says an `OVSX_TOKEN` is configured but unused, while current official install docs and the Open VSX API show Kilo `7.4.11` published and verified on Open VSX. The exact publication mechanism—direct Kilo workflow versus Open VSX auto-publishing—was not verified and does not change Skill Issue's need to publish or arrange publication independently.
- Kilo documents npm plugin cache refresh for bare package names but does not specify an exact freshness interval. Skill Issue should expose explicit version/status reporting rather than promise immediate plugin updates.
- No documented Kilo plugin manifest field was found for contributing agents, skills, commands, rules, or standalone CLI binaries. Treat any plugin-driven file bootstrap as custom installer behavior, not a native packaging guarantee.

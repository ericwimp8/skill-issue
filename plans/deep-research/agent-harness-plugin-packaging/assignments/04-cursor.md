# Cursor Packaging

## Assignment

**Goal:** Determine how current Cursor extension and Cursor-native plugin packaging work, and what Skill Issue must implement to distribute skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, and supporting files as a functional bundled experience.

**Scope:** Internet-only research against current first-party Cursor documentation, changelogs, marketplace terms, the official `cursor/plugins` specification repository, the Agent Skills specification, and primary Open VSX/VS Code extension-publishing sources. The research covers Cursor plugins, VS Code-compatible extensions, rules, commands, skills, subagents, MCP, hooks, Cursor Agent CLI/configuration, installation, updates, discovery, invocation, trust, permissions, scopes, and distribution constraints.

**Exclusions:** No inspection of local Cursor configuration or installed product state; no implementation work; no claims about undocumented runtime behavior presented as validated fact; no analysis of other coding-agent harnesses except where an open standard is directly relevant to Cursor.

## Sources

- [Cursor 3.9 “Customize Cursor” changelog, June 22, 2026](https://cursor.com/changelog/customize) — current unified customization surface and user/team/workspace scopes.
- [Cursor 2.5 “Plugins, Sandbox Access Controls, and Async Subagents” changelog, February 17, 2026](https://cursor.com/changelog/2-5) — initial native plugin bundle and `/add-plugin` installation behavior.
- [Cursor plugin announcement, February 17, 2026](https://cursor.com/blog/marketplace) — official component taxonomy and marketplace launch.
- [Cursor Marketplace](https://cursor.com/marketplace) — current discovery surface and examples of component combinations.
- [Cursor Marketplace Publisher Terms, last updated May 6, 2026](https://cursor.com/marketplace-publisher-terms) — review, update, licensing, pricing, data, and publisher obligations.
- [Official `cursor/plugins` repository](https://github.com/cursor/plugins), inspected at commit `3fe2823ce17c1656c222d4b7c59d3f82fbf20143` dated July 14, 2026 — primary plugin specification and official examples.
- [Official Cursor plugin manifest JSON Schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json) — authoritative manifest fields and closed schema.
- [Official “Create Plugin” scaffold skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/create-plugin-scaffold/SKILL.md) — local plugin location, default component paths, metadata, and validation rules.
- [Official “Review Plugin Submission” skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/review-plugin-submission/SKILL.md) — component discovery and marketplace-readiness checks.
- [Official Continual Learning plugin manifest](https://github.com/cursor/plugins/blob/main/continual-learning/.cursor-plugin/plugin.json) and [hook configuration](https://github.com/cursor/plugins/blob/main/continual-learning/hooks/hooks.json) — first-party proof that agents, skills, hooks, scripts, and `${CURSOR_PLUGIN_ROOT}` cooperate inside one plugin.
- [Cursor 2.4 “Subagents, Skills, and Image Generation” changelog, January 22, 2026](https://cursor.com/changelog/2-4) — editor/CLI skill and subagent availability and skill invocation.
- [Cursor custom slash commands changelog](https://cursor.com/changelog/1-6) — `.cursor/commands` discovery and explicit slash invocation.
- [Cursor Rules documentation](https://docs.cursor.com/context/rules-for-ai) — rule types, `.cursor/rules`, `AGENTS.md`, and project scope.
- [Cursor MCP documentation](https://docs.cursor.com/context/model-context-protocol) — MCP transports, configuration, tool toggles, and approval behavior.
- [Cursor CLI permissions documentation](https://docs.cursor.com/cli/reference/permissions) — global/project configuration paths and allow/deny semantics.
- [Cursor CLI installation documentation](https://docs.cursor.com/en/cli/installation) and [CLI command reference](https://docs.cursor.com/en/cli/reference/parameters) — separate Cursor Agent CLI installation and update surface.
- [Cursor Auto-review changelog, May 29, 2026](https://cursor.com/changelog/auto-review) — current approval handling for shell, MCP, and fetch calls.
- [Cursor hooks partner announcement, December 22, 2025](https://cursor.com/blog/hooks-partners) — hooks as executable control/observation points in the agent loop.
- [Agent Skills specification](https://agentskills.io/specification) — portable `SKILL.md` directory format and bundled resources.
- [Cursor Open VSX transition announcement, June 25, 2025](https://forum.cursor.com/t/extension-marketplace-changes-transition-to-openvsx/109138) — first-party staff announcement for the VS Code-compatible extension gallery.
- [Open VSX Registry](https://open-vsx.org/about) and [Open VSX publisher FAQ](https://www.eclipse.org/legal/open-vsx-registry-faq/) — primary registry and publication requirements.
- [VS Code extension publishing documentation](https://code.visualstudio.com/api/working-with-extensions/publishing-extension) — VSIX packaging and VS Code extension publishing model inherited by compatible editors.

## Findings

### Finding 1: Cursor has two separate extension systems with different purposes

Cursor-native **plugins** package agent customizations. VS Code-compatible **extensions** package executable editor extensions with `package.json`, extension activation, UI contributions, language support, and other VS Code APIs. Cursor’s native plugin system is not a new name for a VS Code extension and does not inherit the VS Code extension manifest.

#### Evidence

Cursor says a native plugin packages skills, subagents, MCP servers, hooks, and rules into a single install, and its current Customize surface manages plugins alongside skills, MCPs, subagents, rules, commands, and hooks at user, team, or workspace level ([Cursor 2.5](https://cursor.com/changelog/2-5), [Cursor 3.9](https://cursor.com/changelog/customize)). The official plugin schema is `.cursor-plugin/plugin.json`, requires only `name`, rejects undeclared properties, and exposes component fields for `commands`, `agents`, `skills`, `rules`, `hooks`, and `mcpServers`; it has no VS Code `activationEvents`, `contributes`, extension entrypoint, or dependency field ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json)).

Separately, Cursor’s staff announced that its in-app **extension library** moved to Open VSX, while `.vsix` files could still be manually installed. The Open VSX Registry describes itself as a registry for VS Code extensions usable by compatible development environments, and standard VS Code tooling packages extensions as VSIX files ([Cursor Open VSX announcement](https://forum.cursor.com/t/extension-marketplace-changes-transition-to-openvsx/109138), [Open VSX](https://open-vsx.org/about), [VS Code publishing](https://code.visualstudio.com/api/working-with-extensions/publishing-extension)).

#### Implication

Skill Issue should implement a Cursor-native plugin as the primary agent bundle. It should build a separate Open VSX/VSIX extension only if it needs editor UI, commands implemented through the VS Code API, custom views, language services, or other editor-host functionality. A Cursor plugin cannot declaratively install or embed a VS Code extension as one of its components, so a truly single-click experience cannot include both systems under the current native plugin schema.

### Finding 2: One Cursor plugin can natively bundle the core agent experience

A single native plugin can combine multiple agent-facing components and supporting files. This is the closest current equivalent to a full Skill Issue harness bundle.

#### Evidence

The official schema accepts relative paths or globs for commands, agents, skills, and rules, a path or object for hooks, and a path/object/array for MCP servers. Its `additionalProperties: false` constraint makes those declared component types the validated extension surface ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json)). Cursor’s own scaffold defines the conventional layout as `.cursor-plugin/plugin.json`, `skills/`, `rules/`, `agents/`, `commands/`, `hooks/hooks.json`, and MCP configuration, with optional explicit component paths when default discovery is insufficient ([scaffold skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/create-plugin-scaffold/SKILL.md), [submission review skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/review-plugin-submission/SKILL.md)).

The official `cursor/plugins` repository demonstrates multi-component plugins. Its Continual Learning plugin declares agents, skills, and hooks in one manifest; its hook runs a bundled TypeScript file using `bun run ${CURSOR_PLUGIN_ROOT}/hooks/continual-learning-stop.ts`, proving a plugin component can call another bundled file by plugin-relative root ([manifest](https://github.com/cursor/plugins/blob/main/continual-learning/.cursor-plugin/plugin.json), [hook](https://github.com/cursor/plugins/blob/main/continual-learning/hooks/hooks.json)). Cursor Team Kit includes skill-local JavaScript, CSS, HTML, and image assets in addition to Markdown skills and agents ([official repository](https://github.com/cursor/plugins)).

#### Implication

Skill Issue can ship one Cursor plugin containing:

- `skills/<name>/SKILL.md` plus skill-local `scripts/`, `references/`, `assets/`, templates, schemas, and data;
- `commands/*.md` as explicit entrypoints for high-value workflows;
- `agents/*.md` for specialized delegated work;
- `rules/*.mdc` for persistent system guidance;
- `hooks/hooks.json` plus executable hook scripts;
- an `mcp.json` or manifest `mcpServers` declaration when an MCP server is part of the experience;
- ordinary `README.md`, `CHANGELOG.md`, `LICENSE`, root assets, and supporting implementation files.

The plugin should keep all cooperating paths relative, use `${CURSOR_PLUGIN_ROOT}` where Cursor explicitly supplies it to hook commands, and validate against the current official schema.

### Finding 3: Skills are the natural owner for Markdown, scripts, references, and assets

Cursor skills support the portable Agent Skills directory model and are the strongest fit for most of Skill Issue’s existing content.

#### Evidence

Cursor 2.4 says skills work in the editor and Cursor CLI, are discoverable by the agent, can be invoked from the slash-command menu, and are defined by `SKILL.md` files that may include commands, scripts, and instructions ([Cursor 2.4](https://cursor.com/changelog/2-4)). The Agent Skills specification defines a skill as a directory containing `SKILL.md` plus optional `scripts/`, `references/`, `assets/`, and any additional files. It specifies progressive disclosure: metadata is loaded for discovery, full instructions on activation, and resources only when required ([Agent Skills specification](https://agentskills.io/specification)).

Cursor’s official scaffold requires `skills/<skill-name>/SKILL.md` with `name` and `description` frontmatter, while the official repository includes executable and presentation assets inside skill folders ([scaffold skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/create-plugin-scaffold/SKILL.md), [official repository](https://github.com/cursor/plugins)).

#### Implication

Skill Issue should preserve each portable skill as a self-contained directory rather than flattening scripts, references, and assets into unrelated plugin-level folders. `name` and `description` must be concise and accurate because they drive discovery. Explicit slash commands should wrap only workflows that users must intentionally invoke; automatic or contextual workflows should rely on skill descriptions. Any Skill Issue feature that currently assumes a global script path must be rewritten to use a skill-relative or plugin-root-relative path.

### Finding 4: Rules, commands, and subagents have distinct invocation semantics and scopes

Rules provide persistent instructions, commands provide explicit user entrypoints, and subagents provide delegated isolated execution. They cooperate in a plugin but should not be treated as interchangeable copies of the same Markdown.

#### Evidence

Cursor Rules are system-level instructions. Project rules live in `.cursor/rules`, can be always applied, glob-attached, agent-requested, or manually invoked, while `AGENTS.md` is the simple plain-Markdown project alternative ([Rules documentation](https://docs.cursor.com/context/rules-for-ai)). Custom commands are Markdown files under `.cursor/commands/[command].md` and run from the `/` menu ([custom commands changelog](https://cursor.com/changelog/1-6)). Subagents have independent contexts, can run in parallel, and can be configured with prompts, tools, and models; Cursor supports both defaults and custom definitions ([Cursor 2.4](https://cursor.com/changelog/2-4)).

At the plugin layer, the schema and official scaffold discover `rules/`, `commands/`, and `agents/` separately, each with its own metadata expectations ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json), [scaffold skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/create-plugin-scaffold/SKILL.md)). Cursor 3.9 now exposes customization management at user, team, and workspace levels, while project files remain version-controlled workspace customizations ([Cursor 3.9](https://cursor.com/changelog/customize)).

#### Implication

Skill Issue should map behavior by semantic owner:

- stable constraints and cross-workflow invariants → rules;
- intentional workflow launchers → commands;
- dynamic procedural knowledge and related files → skills;
- bounded specialist roles and parallel research/execution → agents.

The public plugin should default to user-installable components, while teams can centrally distribute the same plugin. Project-specific overrides or repository policy should remain in the repository’s `.cursor/` or `AGENTS.md` files rather than being silently written by the plugin.

### Finding 5: MCP and hooks make the bundle executable, but they also create the principal trust boundary

MCP supplies agent tools and external data; hooks run scripts around the agent loop and can observe, block, or modify behavior. Both can be included in a plugin, and both require explicit security design.

#### Evidence

The plugin schema declares `mcpServers` and `hooks` as first-class components ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json)). Cursor’s MCP documentation supports local stdio and remote server configurations, lets users enable or disable MCP tools, and states that MCP tool use asks for approval by default; users can enable auto-run ([MCP documentation](https://docs.cursor.com/context/model-context-protocol)). Cursor describes hooks as custom scripts that run before or after defined stages and can observe, block, or modify behavior ([hooks announcement](https://cursor.com/blog/hooks-partners)). Cursor 3.6’s Auto-review mode applies to shell, MCP, and fetch calls, allowing allowlisted or sandboxable calls and otherwise using a classifier to decide whether to allow, redirect, or request approval ([Auto-review](https://cursor.com/changelog/auto-review)).

Cursor’s first-party Continual Learning plugin demonstrates that a hook can execute a bundled script by `${CURSOR_PLUGIN_ROOT}`. That is functional cooperation, not passive documentation ([hook configuration](https://github.com/cursor/plugins/blob/main/continual-learning/hooks/hooks.json)).

#### Implication

Skill Issue can provide executable automation in one plugin, but installation should be treated as granting code and instruction trust. Hooks must be narrowly scoped, deterministic, cross-platform where promised, and documented with their events and effects. MCP tools should use least privilege, clearly separate read and mutating operations, preserve Cursor’s approval defaults, and avoid embedding long-lived credentials. The bundle should disclose runtime dependencies such as Node, Bun, Python, network access, and external services in its README and skill compatibility metadata.

### Finding 6: Cursor plugin packaging does not install a standalone CLI or arbitrary Cursor configuration

The native manifest has no binary, PATH, package-manager, activation, post-install, settings contribution, sandbox-policy, or CLI-permissions field. A CLI can be shipped as files used by the agent, but not exposed as a normal user shell command solely through plugin installation.

#### Evidence

The official schema is closed (`additionalProperties: false`) and lists only plugin metadata plus commands, agents, skills, rules, hooks, and MCP servers. It defines no `bin`, install script, editor settings, CLI permissions, or sandbox configuration property ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json)). The scaffold likewise creates plugin files and component paths but defines no installation lifecycle hook or PATH registration ([scaffold skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/create-plugin-scaffold/SKILL.md)).

Cursor Agent CLI is installed and updated independently (`curl https://cursor.com/install -fsS | bash`, `cursor-agent update`), and its permissions are configured in `~/.cursor/cli-config.json` or project `.cursor/cli.json`, where deny rules take precedence ([CLI installation](https://docs.cursor.com/en/cli/installation), [CLI permissions](https://docs.cursor.com/cli/reference/permissions)). Cursor’s sandbox/network settings are also a separate configuration surface rather than plugin manifest fields ([Cursor 2.5](https://cursor.com/changelog/2-5)).

#### Implication

Skill Issue needs a dual delivery model:

1. **Cursor plugin:** bundles skills, commands, agents, rules, hooks, MCP, and agent-invoked scripts.
2. **Standalone CLI distribution:** installs the `skill-issue` executable through an appropriate package manager, downloadable binary, or bootstrap installer and owns upgrades/uninstallation.

The plugin can call a bundled private helper or detect and guide installation of the standalone CLI, but it should not claim that `/add-plugin` installs a global shell command. User configuration should be opt-in and performed through documented setup commands or a CLI with preview/confirmation, not assumed to be declaratively installed. If the standalone CLI is required for core workflows, the plugin needs a setup skill/command that checks version compatibility and produces an actionable installation path.

### Finding 7: Native installation and distribution are strong, but marketplace publication imposes material constraints

Cursor supports one-click public installation, local development, and controlled team distribution. Public marketplace publication is reviewed, free to users, and license-constrained.

#### Evidence

Users can discover plugins in the Cursor Marketplace or install directly from the editor with `/add-plugin` ([Cursor 2.5](https://cursor.com/changelog/2-5)). Cursor’s official scaffold places development plugins under `~/.cursor/plugins/local/<plugin-name>/`, where they are immediately available, and multi-plugin repositories use `.cursor-plugin/marketplace.json` plus per-plugin `.cursor-plugin/plugin.json` manifests ([scaffold skill](https://github.com/cursor/plugins/blob/main/create-plugin/skills/create-plugin-scaffold/SKILL.md), [official repository](https://github.com/cursor/plugins)). Cursor 3.9 allows management at user, team, and workspace levels, and team marketplaces can import plugin repositories from GitLab, Bitbucket, Azure DevOps, as well as the existing GitHub path ([Cursor 3.9](https://cursor.com/changelog/customize)). Team administrators can distribute plugins as Default Off, Default On, or Required ([Team Marketplace Updates](https://cursor.com/changelog/05-01-26)).

Publisher Terms require submission through the publish application, code and identity/business review, re-index/review for updates, ongoing cooperation, accurate metadata, and publisher-operated support. Marketplace plugins must be free to users. Included open-source components must use permissive licenses; the terms explicitly disallow GPL, AGPL, and LGPL components. Approval is not an endorsement or certification of security or functionality ([Publisher Terms](https://cursor.com/marketplace-publisher-terms)).

#### Implication

Skill Issue can deliver a coherent one-install agent experience through the public Marketplace and a private/team variant through team marketplaces. Before submission it must audit every bundled dependency and asset license, because a copyleft dependency can block Marketplace eligibility even if the Skill Issue repository itself is open source. The public listing must be free; monetized services must be clearly separated from access to the plugin itself and reviewed against the publisher terms. Release automation should bump manifest versions, validate paths/frontmatter/schema, publish source changes, request re-index, and wait for review rather than assuming a repository push immediately reaches users.

### Finding 8: VS Code extension inheritance is useful only for editor-native gaps

Cursor can run many VS Code-compatible extensions, but their discovery, packaging, compatibility, and update channel are independent from Cursor native plugins.

#### Evidence

Cursor moved its extension gallery to Open VSX and supports manual VSIX installation; it also warned that changing the marketplace backend was unsupported ([Cursor Open VSX announcement](https://forum.cursor.com/t/extension-marketplace-changes-transition-to-openvsx/109138)). Open VSX is an independent registry for VS Code extensions, with its own publisher agreement and licensing requirements ([Open VSX about](https://open-vsx.org/about), [Open VSX FAQ](https://www.eclipse.org/legal/open-vsx-registry-faq/)). VS Code extensions use VSIX packaging and `vsce`, and compatibility depends on the declared VS Code engine/API surface ([VS Code publishing](https://code.visualstudio.com/api/working-with-extensions/publishing-extension)).

The Cursor plugin schema has no field for an extension ID or VSIX payload, so there is no validated link that makes a native plugin install its companion editor extension ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json)).

#### Implication

Skill Issue should avoid building an extension merely to move Markdown and scripts; native plugins already own that concern. Build and publish an Open VSX extension only for a concrete UI/API requirement. If both are necessary, use coordinated names, versions, documentation, and capability detection, but present them as two installations. A fallback installer may automate local VSIX/plugin setup for development or air-gapped environments, but Marketplace distribution still remains two independently governed channels.

### Finding 9: The feasible Skill Issue product shape is “one primary bundle plus optional companions”

One bundled Cursor agent experience is feasible for nearly all stated content types. A literally universal single installer is not currently supported when the standalone CLI, global configuration, or VS Code extension must also be installed.

#### Evidence

Native plugins cover the full agent customization set and can include arbitrary supporting files used by those components ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json), [official repository](https://github.com/cursor/plugins)). Skills explicitly accommodate Markdown, scripts, references, assets, and additional files ([Agent Skills specification](https://agentskills.io/specification)). The missing manifest fields for binaries, extension installation, PATH, editor settings, CLI permissions, and sandbox policy establish the boundary of the single native install ([official schema](https://github.com/cursor/plugins/blob/main/schemas/plugin.schema.json)).

#### Implication

Recommended Cursor implementation:

1. Create a single focused `skill-issue` Cursor plugin with `.cursor-plugin/plugin.json`.
2. Put portable domain/workflow bundles under `skills/`, retaining scripts, references, assets, configuration templates, and supporting files inside each skill.
3. Add a small set of explicit `commands/` for install/setup/status/update and major workflows.
4. Add only genuinely specialized `agents/`; keep orchestration instructions in the owning skill or command.
5. Add minimal `rules/` for invariants that should persist whenever the plugin is enabled.
6. Add `hooks/` only where event-driven behavior is essential and disclose every effect.
7. Add MCP only when Skill Issue needs durable tools or external data access; preserve approvals and least privilege.
8. Distribute the standalone `skill-issue` CLI independently; let the plugin detect it, validate its version, and guide installation or use a bundled private helper for workflows that do not need PATH access.
9. Add a separate Open VSX extension only if Skill Issue requires editor UI or VS Code API contributions.
10. Provide public Marketplace, local-development, and team-marketplace release paths from the same source tree, with schema/frontmatter/license checks and explicit compatibility tests for Cursor IDE and Cursor Agent CLI.

This yields a one-click functional agent bundle for most users, with an optional second step only for capabilities the Cursor plugin specification does not own.

## Notes

- Cursor’s current documentation site is heavily client-rendered, and several current pages exposed little machine-readable body text. Claims above were cross-checked against Cursor changelogs, Publisher Terms, and the official `cursor/plugins` repository rather than inferred from page titles.
- No first-party current parity matrix was found proving that every plugin component behaves identically in Cursor IDE, Cursor Agent CLI, cloud agents, and automations. Skills, subagents, rules, and MCP have separate official CLI evidence, but hook and command parity across all surfaces should be treated as caveated until tested against supported product versions.
- No authoritative current command was found for installing an arbitrary third-party plugin from the standalone `cursor-agent` CLI. `/add-plugin` is verified for the editor Agent input. A CLI-only plugin installer should be classified as unsupported until Cursor documents one.
- Cursor staff have described automatic public-plugin refreshes and commit-pinned atomic updates in the official forum, but the durable contractual source only guarantees review/re-index behavior. Client update timing, user-visible version pinning, rollback, and diff inspection remain documentation gaps and should not be promised.
- The June 2025 Open VSX transition announcement is first-party but historical. Open VSX remains the supported evidence for extension inheritance; current Cursor-specific extension-signature and workspace-trust defaults were not revalidated from a 2026 primary page and are therefore unsupported here.
- The schema proves there is no declarative binary/PATH/config/VSIX field. It does not prove Cursor will never add one; Skill Issue should pin its implementation to a tested Cursor version and monitor the official schema for changes.

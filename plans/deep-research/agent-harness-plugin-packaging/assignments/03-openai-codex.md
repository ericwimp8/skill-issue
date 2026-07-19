# OpenAI Codex Packaging

## Assignment

**Goal:** Determine how current OpenAI Codex plugin and extension packaging works, and what Skill Issue must implement to distribute skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, agents, and supporting files as a functional bundled experience.

**Scope:** Current first-party Codex and ChatGPT plugin documentation, the official `openai/codex` implementation at commit `35eaf3ffb0bf2001486c68c47a3d946b34d16634` (2026-07-18), the official `openai/plugins` examples at commit `11c74d6ba24d3a6d48f54a194cd00ef3beea18f9` (2026-07-13), and the documented interaction among plugins, marketplaces, skills, custom-agent TOMLs, MCP, apps/connectors, configuration, `AGENTS.md`, rules, installation, update, discovery, invocation, trust, permissions, and public distribution.

**Exclusions:** Local Codex configuration or installed-plugin state as product evidence; private or third-party packaging conventions as authority; implementation of the Skill Issue bundle; and claims about undocumented server-side behavior that cannot be validated from official documentation or the open-source client.

## Sources

- OpenAI, [Plugins](https://learn.chatgpt.com/docs/plugins), inspected 2026-07-19.
- OpenAI, [Build plugins](https://learn.chatgpt.com/docs/build-plugins), inspected 2026-07-19.
- OpenAI, [Build skills](https://learn.chatgpt.com/docs/build-skills), inspected 2026-07-19.
- OpenAI, [Build an app](https://learn.chatgpt.com/docs/build-app), inspected 2026-07-19.
- OpenAI, [Submit plugins](https://learn.chatgpt.com/docs/submit-plugins), inspected 2026-07-19.
- OpenAI, [Subagents](https://learn.chatgpt.com/docs/agent-configuration/subagents), inspected 2026-07-19.
- OpenAI, [Custom instructions with AGENTS.md](https://learn.chatgpt.com/docs/agent-configuration/agents-md), inspected 2026-07-19.
- OpenAI, [Rules](https://learn.chatgpt.com/docs/agent-configuration/rules), inspected 2026-07-19.
- OpenAI, [Model Context Protocol](https://learn.chatgpt.com/docs/extend/mcp?surface=cli), inspected 2026-07-19.
- OpenAI, [Config basics](https://learn.chatgpt.com/docs/config-file/config-basic) and [Configuration Reference](https://learn.chatgpt.com/docs/config-file/config-reference), inspected 2026-07-19.
- OpenAI, [Developer commands](https://learn.chatgpt.com/docs/developer-commands?surface=cli), inspected 2026-07-19.
- OpenAI, [Agent approvals & security](https://learn.chatgpt.com/docs/agent-approvals-security), inspected 2026-07-19.
- OpenAI, [Plugin controls](https://learn.chatgpt.com/docs/enterprise/apps-and-connectors) and [Skill controls](https://learn.chatgpt.com/docs/enterprise/skills), inspected 2026-07-19.
- `openai/codex`, [`core-plugins/src/manifest.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/manifest.rs), [`core-plugins/src/loader.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/loader.rs), [`core-plugins/src/marketplace.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/marketplace.rs), [`core-plugins/src/npm_source.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/npm_source.rs), and [`core/src/config/agent_roles.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core/src/config/agent_roles.rs), commit `35eaf3ffb0bf2001486c68c47a3d946b34d16634`.
- `openai/codex`, built-in [`plugin-creator` skill](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/skills/src/assets/samples/plugin-creator/SKILL.md) and its [install/update reference](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/skills/src/assets/samples/plugin-creator/references/installing-and-updating.md), commit `35eaf3ffb0bf2001486c68c47a3d946b34d16634`.
- `openai/plugins`, [repository README](https://github.com/openai/plugins/blob/11c74d6ba24d3a6d48f54a194cd00ef3beea18f9/README.md), commit `11c74d6ba24d3a6d48f54a194cd00ef3beea18f9`.

## Findings

### Finding 1: A Codex plugin is a real composition and distribution unit, but its documented native component set is bounded

Codex plugins have a required `.codex-plugin/plugin.json` and can natively compose skills, MCP server definitions, app/connector mappings, lifecycle hooks, and presentation assets. The public manifest fields point to `skills`, `mcpServers`, `apps`, and `hooks` relative to the plugin root. Current open-source loading agrees: `RawPluginManifest` recognizes those component classes, while `LoadedPlugin` is populated with skill roots, MCP servers, apps, and hook sources. There is no corresponding custom-agent TOML, `AGENTS.md`, rules, general Codex config, or executable-install field in that loader contract.

#### Evidence

- The authoring documentation defines the required manifest and root-level `skills/`, `hooks/`, `.app.json`, `.mcp.json`, and `assets/` layout, and says these components are addressed relative to the plugin root: [Build plugins — Plugin structure and manifest fields](https://learn.chatgpt.com/docs/build-plugins#plugin-structure).
- The pinned client manifest parser recognizes `skills`, `mcp_servers`, `apps`, `hooks`, and interface metadata: [`RawPluginManifest`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/manifest.rs#L17-L45). The loader's concrete result contains `skill_roots`, `mcp_servers`, `apps`, and `hook_sources`, then loads those exact capabilities: [`load_plugin`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/loader.rs#L748-L850).
- OpenAI's user documentation says installed plugins contribute skills, connectors, MCP tools, browser extensions, and hooks, with host sandbox/approval policy still applying: [Plugins — Overview and permissions](https://learn.chatgpt.com/docs/plugins#overview).

#### Implication

Skill Issue can ship one **core Codex plugin bundle** containing its skills, their Markdown guidance, scripts, references and assets, optional MCP/app integration, hooks, and presentation metadata. It cannot treat that manifest as a generic installer for every Codex extension surface. Custom-agent TOMLs, project/global config, `AGENTS.md`, rules, and a globally callable CLI require separate installation or project-initialization work.

### Finding 2: Skills are the correct native container for Skill Issue instructions, Markdown, scripts, references, and assets

A skill is explicitly designed as a folder containing `SKILL.md` plus optional `scripts/`, `references/`, `assets/`, and `agents/openai.yaml`. Codex can invoke it explicitly through `/skills`, `$skill-name`, or a prompt mention, or implicitly from its `description`. Plugins can bundle multiple such directories and namespace them by plugin. This covers most of Skill Issue's content without inventing new packaging.

The `agents/openai.yaml` mentioned in the skill layout is skill appearance/dependency metadata; it is a different artifact from standalone custom-agent TOMLs under `.codex/agents/` or `~/.codex/agents/`.

#### Evidence

- The skills documentation defines the complete skill directory shape and both explicit and implicit invocation: [Build skills — structure and use](https://learn.chatgpt.com/docs/build-skills#how-codex-uses-skills).
- The same documentation recommends plugins when distributing multiple skills or pairing skills with connectors, MCP configuration, and presentation assets: [Build skills — Distribute skills with plugins](https://learn.chatgpt.com/docs/build-skills#distribute-skills-with-plugins).
- Filesystem skills otherwise have repository, user, administrator, and bundled-system roots, which are separate authoring/discovery paths from plugin installation: [Build skills — Where to save skills](https://learn.chatgpt.com/docs/build-skills#where-to-save-skills).

#### Implication

Skill Issue should make each workflow a semantically narrow plugin skill and keep its helper scripts, long-form Markdown references, templates, fixtures, and non-code assets inside that skill directory when they are owned by that workflow. Cross-skill shared files may live elsewhere in the plugin, but skills should reference them using stable plugin-relative conventions and should not assume a checkout path. Invocation documentation should support both direct `$skill-name` selection and precise descriptions for implicit matching.

### Finding 3: Marketplace catalogs, not bare plugin folders, are the normal install and update boundary

Plugins are discovered from marketplaces. Codex supports repository marketplaces at `$REPO_ROOT/.agents/plugins/marketplace.json`, a personal marketplace at `~/.agents/plugins/marketplace.json`, configured local or Git marketplace roots, the official curated marketplace, and a documented legacy-compatible `.claude-plugin/marketplace.json`. Marketplace entries specify plugin source, installation/authentication policy, and category. Sources can be local paths, Git-backed locations, or npm registry packages. The CLI exposes `codex plugin add/list/remove` and `codex plugin marketplace add/list/remove/upgrade`; the TUI exposes `/plugins`. A new session is required before newly installed skills or tools are available.

Marketplace refresh and plugin reinstall are distinct operations. `marketplace upgrade` refreshes configured Git marketplaces; local development guidance changes the plugin version with a cachebuster, reruns `codex plugin add`, and starts a new thread. For npm sources, current Codex runs `npm pack --ignore-scripts` and extracts the package rather than performing an npm global install.

#### Evidence

- Repo/personal marketplace file locations, entry policy, path containment, Git selectors, npm registry sources, and legacy `.claude-plugin` discovery are documented in [Build plugins — Install a local plugin and marketplace metadata](https://learn.chatgpt.com/docs/build-plugins#install-a-local-plugin-manually).
- The CLI command contract documents plugin install/list/remove and marketplace add/list/remove/upgrade from GitHub shorthand, Git/SSH URLs, or local roots: [Developer commands — `codex plugin`](https://learn.chatgpt.com/docs/developer-commands?surface=cli#codex-plugin) and [`codex plugin marketplace`](https://learn.chatgpt.com/docs/developer-commands?surface=cli#codex-plugin-marketplace).
- Codex tells users to start a new CLI session after installation: [Plugins — Install and use](https://learn.chatgpt.com/docs/plugins#install-and-use-a-plugin-in-chatgpt).
- The built-in plugin creator's local update flow uses a version cachebuster, reinstalls with `codex plugin add`, and tests in a new thread: [`installing-and-updating.md`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/skills/src/assets/samples/plugin-creator/references/installing-and-updating.md).
- The pinned npm source implementation invokes `npm pack --ignore-scripts` and unpacks the resulting archive: [`npm_source.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/npm_source.rs#L79-L115).

#### Implication

Skill Issue should publish a marketplace repository or a public plugin listing, not ask users to copy an arbitrary folder and hope Codex scans it. A Git marketplace is the simplest open-source channel because it carries the catalog and plugin tree together and supports `marketplace upgrade`. Releases should use stable semantic plugin versions; local development can use the documented cachebuster/reinstall loop. An npm marketplace source can transport the bundle, but it must not be relied on to run postinstall scripts, register a binary, or mutate Codex config.

### Finding 4: Custom-agent TOMLs are a separate configuration layer and are not currently installed by the local plugin loader

Custom Codex agents are standalone TOML files discovered under `~/.codex/agents/` for personal scope or `.codex/agents/` for project scope. Each defines `name`, `description`, and `developer_instructions`, and may override normal config keys such as model, reasoning effort, sandbox, MCP, and skill configuration. The agent resolver discovers `agents/` beside active config layers; the plugin loader has no agent-TOML capability.

The official `openai/plugins` example repository physically contains some plugin-level `agents/` content, but its README is broader than the documented local Codex manifest and those files are not evidence that Codex CLI installs standalone custom-agent TOMLs from a plugin. Current public Codex authoring docs omit agents from plugin structure, while the open-source loader omits them from the loaded capability set. Therefore plugin-level example agent files must be treated as product-specific or forward-looking unless a documented local loader appears.

#### Evidence

- Standalone personal and project custom-agent locations, required keys, and inherited config fields are documented in [Subagents — Custom agents](https://learn.chatgpt.com/docs/agent-configuration/subagents#custom-agents).
- The current agent role implementation discovers `agents/` under each active config layer, independently of plugin roots: [`load_agent_roles`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core/src/config/agent_roles.rs#L16-L101).
- The plugin manifest and loader evidence in Finding 1 contains no standalone-agent field or load step: [`manifest.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/manifest.rs#L17-L45) and [`loader.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/loader.rs#L765-L850).
- The official examples repository describes plugin-level `agents/` among possible companion content, but does not define it as the local custom-agent TOML install path: [`openai/plugins` README](https://github.com/openai/plugins/blob/11c74d6ba24d3a6d48f54a194cd00ef3beea18f9/README.md).

#### Implication

If Skill Issue workflows depend on named subagent roles, its standalone installer/CLI must explicitly install, update, and uninstall those TOMLs into either `.codex/agents/` or `~/.codex/agents/`, with project scope preferred when the role is project-specific. The installer should be idempotent, preserve unrelated files, detect conflicts, and require a fresh Codex session after changes. Bundling TOMLs inside the plugin is useful as source payload, but does not make them active by itself.

### Finding 5: `config.toml`, `AGENTS.md`, and rules remain separately scoped, trusted policy surfaces

Codex configuration is layered: CLI overrides, trusted project `.codex/config.toml` files, profiles, user config, system config, then defaults. `AGENTS.md` is discovered once per run from global Codex home and project-root-to-CWD paths, with nearer project files overriding earlier guidance. Rules are `.rules` files under `rules/` beside active config layers; project-local rules load only when the project's `.codex/` layer is trusted. None is a plugin manifest component.

This separation is semantically important. A plugin is reusable capability distribution; `AGENTS.md` is repository instruction ownership; config changes models, permissions, MCP, skills, and agent behavior; rules alter command escalation behavior. Automatically smuggling these policy layers in through plugin installation would bypass their documented scope and trust boundaries.

#### Evidence

- Config locations, project trust requirement, and precedence are documented in [Config basics](https://learn.chatgpt.com/docs/config-file/config-basic#configuration-precedence).
- Global and hierarchical project instruction discovery is documented in [Custom instructions with AGENTS.md — How Codex discovers guidance](https://learn.chatgpt.com/docs/agent-configuration/agents-md#how-codex-discovers-guidance).
- Rules load beside active config layers, require project trust, and are explicitly experimental: [Rules](https://learn.chatgpt.com/docs/agent-configuration/rules).
- Per-skill enablement lives in `skills.config`, while plugin-bundled MCP servers have plugin-scoped policy under `plugins.<plugin>.mcp_servers.<server>`: [Configuration Reference](https://learn.chatgpt.com/docs/config-file/config-reference).

#### Implication

Skill Issue should ship **templates and an opt-in initializer** for these surfaces rather than claim the plugin installs them. The initializer can offer: repository `AGENTS.md` snippets at the semantic owner, `.codex/config.toml` fragments, `.codex/rules/*.rules`, and agent TOMLs. It must preview changes, merge conservatively, back up or refuse conflicts, preserve unrelated configuration, explain project versus user scope, and prompt the user to trust/restart the project when required. Global installation should be a separate explicit command.

### Finding 6: MCP and apps/connectors can cooperate inside one plugin, but authentication and policy remain external controls

An app is the MCP-backed capability inside a plugin. It can expose MCP tools, authentication, structured data, and optional UI. A plugin may point `mcpServers` to `.mcp.json` and `apps` to `.app.json`; installed plugin MCP servers are launched from the plugin rather than copied into top-level user MCP configuration. Users can independently enable a bundled server and restrict tool approval or enabled tools. Connector authentication can happen on install or first use, and workspace admins can separately control plugin availability, connector access, actions, and permissions.

Local STDIO MCP definitions can start commands with arguments and environment-variable forwarding. Remote streamable HTTP servers can use OAuth or bearer-token/header configuration. Consequently, a Skill Issue MCP service must package a runnable, dependency-complete entrypoint or point to a deployed HTTPS service; mere source files do not guarantee that the runtime is installed.

#### Evidence

- Plugin-provided MCP servers and their plugin-scoped policy are documented in [Model Context Protocol — Plugin-provided MCP servers](https://learn.chatgpt.com/docs/extend/mcp?surface=cli#plugin-provided-mcp-servers).
- The plugin manifest's `.mcp.json` forms and per-tool policy are documented in [Build plugins — Bundled MCP servers and lifecycle hooks](https://learn.chatgpt.com/docs/build-plugins#bundled-mcp-servers-and-lifecycle-hooks).
- Apps are explicitly the MCP-backed portion of a plugin, with optional Apps SDK UI and tool safety metadata: [Build an app](https://learn.chatgpt.com/docs/build-app).
- Plugin installation, connector access, connector action control, source-system authorization, and runtime permissions are separate layers: [Plugin controls — capability chain](https://learn.chatgpt.com/docs/enterprise/apps-and-connectors#understand-the-capability-chain).

#### Implication

Skill Issue can use one plugin to make its instructions and tools discoverable together. For a local MCP service, it should ship a self-contained binary or a thin launcher that performs a non-mutating dependency check and emits actionable setup errors. For a hosted service, it should implement OAuth/tool annotations/privacy controls and treat external-service connection as a separate user action. Installation tests must verify both skill discovery and actual MCP tool availability, not infer tool readiness from manifest presence.

### Finding 7: Plugin installation does not waive sandbox, approval, hook-trust, MCP, or administrator controls

When plugin capabilities run through Codex, the host sandbox and approval policy still apply. Local Codex defaults restrict writes to the workspace and disable network access. Plugin hooks are especially guarded: installing or enabling a plugin does not trust its hooks, and Codex skips untrusted non-managed hooks until the user reviews and trusts the current definition. MCP servers and individual tools can be enabled, disabled, allowlisted, or assigned approval modes. Administrators can restrict allowed marketplace sources and plugin-bundled MCP identities.

#### Evidence

- Default local sandbox/network boundaries and approval behavior are documented in [Agent approvals & security](https://learn.chatgpt.com/docs/agent-approvals-security).
- Hook trust is explicit: enabled plugin hooks remain skipped until reviewed and trusted, and receive `PLUGIN_ROOT` and `PLUGIN_DATA` rather than arbitrary install-time privileges: [Build plugins — lifecycle hooks](https://learn.chatgpt.com/docs/build-plugins#bundled-mcp-servers-and-lifecycle-hooks).
- Admin requirements can restrict marketplace sources and plugin MCP server identity: [Configuration Reference — marketplace and plugin requirements](https://learn.chatgpt.com/docs/config-file/config-reference).
- Workspace controls additionally gate connector-backed plugin availability and actions: [Plugin controls](https://learn.chatgpt.com/docs/enterprise/apps-and-connectors).

#### Implication

Skill Issue must treat every script, hook, CLI action, and MCP write tool as a separately reviewed execution surface. Hooks should be minimal, transparent, and optional; state belongs under `PLUGIN_DATA` when hooks need it. Skills should describe why a command needs network or out-of-workspace access. MCP tools need accurate read/write/destructive annotations and conservative default approval modes. The installer must fail clearly when an administrator policy blocks a marketplace or server rather than attempting to bypass it.

### Finding 8: A plugin package does not install a standalone CLI into `PATH`

No official plugin manifest field or install command inspected registers an executable, creates a shell shim, or installs an npm package globally. The npm plugin source path runs `npm pack --ignore-scripts` and extracts the archive as a plugin bundle. Skills can contain scripts and invoke them through their known location; hooks can invoke plugin-contained commands with `PLUGIN_ROOT`; neither behavior creates a user-facing global command.

This conclusion is an inference from the documented component contract plus the concrete npm materialization path. A plugin may physically carry a CLI binary or script as supporting content, but current first-party evidence does not establish automatic `PATH` registration.

#### Evidence

- The complete documented manifest field list contains metadata, skills, MCP servers, apps, hooks, and interface assets, without a binary/CLI registration field: [Build plugins — Manifest fields](https://learn.chatgpt.com/docs/build-plugins#manifest-fields).
- Skills may package scripts and hooks may execute plugin-relative commands: [Build skills](https://learn.chatgpt.com/docs/build-skills) and [Build plugins — lifecycle hooks](https://learn.chatgpt.com/docs/build-plugins#bundled-mcp-servers-and-lifecycle-hooks).
- The current npm transport only packs and extracts the package with lifecycle scripts disabled: [`npm_source.rs`](https://github.com/openai/codex/blob/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/npm_source.rs#L79-L115).

#### Implication

Skill Issue needs an independent CLI distribution channel—such as npm, Homebrew, Cargo, standalone binaries, or a curl-able installer—or an explicit bootstrap script that installs the CLI with user consent. The plugin should detect the CLI and provide installation guidance, and its skills may call a plugin-contained internal helper when a global command is unnecessary. The plugin installation flow should never imply that `skill-issue` is on `PATH` unless verification proves it.

### Finding 9: Surface support and public distribution constrain the meaning of “one bundled experience”

Plugins are available in Work mode on ChatGPT web, Work mode and Codex in the desktop app, and the Codex CLI plugin browser. They are not available in Chat mode, the IDE extension, or mobile. Skills themselves are available in the desktop app, CLI, and IDE extension through filesystem discovery, so a plugin-only release does not provide the same experience on every Codex surface.

For public directory distribution, OpenAI currently accepts skills-only, app-only, and app-plus-skills submissions. Publishing requires Platform app-management permission, verified developer/business identity, review, accurate listing and tool metadata, tests, policies, and—for apps—a public production MCP URL, domain/CSP work, and reviewer-ready authentication. The public submission page does not describe custom-agent TOMLs, `AGENTS.md`, rules, standalone CLI installation, or general config as submitted components. Although local build documentation supports hooks, the public submission categories and checklist do not clearly promise public hook distribution.

#### Evidence

- Surface availability and the need for a new session are documented in [Plugins](https://learn.chatgpt.com/docs/plugins) and [Plugin controls](https://learn.chatgpt.com/docs/enterprise/apps-and-connectors).
- Filesystem skill availability is separately documented in [Build skills](https://learn.chatgpt.com/docs/build-skills).
- Public plugin types, identity requirements, review, and publication flow are documented in [Submit plugins](https://learn.chatgpt.com/docs/submit-plugins).

#### Implication

Skill Issue should define two supported experiences: (1) a plugin-first ChatGPT desktop/Work/Codex CLI experience, and (2) an installer-managed filesystem experience for IDE and project-scoped Codex usage. Public directory submission can distribute the skills and hosted app/MCP portion, but the CLI, custom agents, config, rules, and project guidance still need an external initializer. Public hooks should be treated as caveated until the submission portal and review policy explicitly accept them.

### Finding 10: Codex contains compatibility bridges, but Skill Issue should target native skills rather than undocumented command/agent conventions

Codex documents legacy `.claude-plugin/marketplace.json` discovery, and the open-source client also contains a plugin command-migration layer that rewrites source Markdown commands into generated skills. The official examples repository contains `commands/` and plugin-level `agents/` folders. These are useful compatibility signals, but the current public native authoring model is skills plus MCP/apps/hooks, and the current local plugin loader does not install standalone custom-agent TOMLs.

#### Evidence

- Legacy-compatible marketplace discovery is documented in [Build plugins — marketplace locations](https://learn.chatgpt.com/docs/build-plugins#how-the-chatgpt-desktop-app-uses-marketplaces).
- The client contains a [`command_migration`](https://github.com/openai/codex/tree/35eaf3ffb0bf2001486c68c47a3d946b34d16634/codex-rs/core-plugins/src/command_migration) module that converts Markdown command sources into generated `SKILL.md` files.
- The official examples repository advertises broader companion surfaces, including commands and agents: [`openai/plugins` README](https://github.com/openai/plugins/blob/11c74d6ba24d3a6d48f54a194cd00ef3beea18f9/README.md).
- The normative plugin authoring page omits commands and agents from its documented structure: [Build plugins — Plugin structure](https://learn.chatgpt.com/docs/build-plugins#plugin-structure).

#### Implication

Skill Issue may import or translate equivalent command-oriented content, but its canonical Codex artifact should be a native skill. It should not make `commands/` or plugin-level agent Markdown a hard dependency. Custom roles should remain TOMLs installed through the explicit initializer until first-party Codex documentation and loader code establish plugin-native agent installation.

### Finding 11: The functional Skill Issue design is one core plugin plus a deliberate bootstrap/initializer boundary

The strongest supported architecture is a single marketplace-installed plugin for all capabilities Codex natively composes, paired with a small standalone Skill Issue installer/CLI for separately discovered surfaces. The plugin should remain useful by itself; the initializer upgrades it into the full agent/config/project-policy experience.

#### Evidence

- Native plugin composition, marketplace installation, and restart semantics are established in Findings 1–3.
- Separate custom-agent, config, `AGENTS.md`, and rules discovery are established in Findings 4–5.
- External runtime, trust, CLI, and surface constraints are established in Findings 6–10.

#### Implication

Skill Issue should implement the following concrete Codex packaging shape:

1. **Marketplace repository:** `.agents/plugins/marketplace.json` with a stable marketplace name, plugin entry, explicit install/auth policy, category, and a Git-upgradeable release source.
2. **Core plugin:** `plugins/skill-issue/.codex-plugin/plugin.json` with semantic version, publisher/discovery/interface metadata, `skills`, optional `mcpServers`, optional `apps`, optional local hooks, and assets.
3. **Skill-owned content:** `skills/<workflow>/SKILL.md`, `scripts/`, `references/`, `assets/`, and skill `agents/openai.yaml` where appropriate. All core workflows must work without custom-agent installation, or degrade with an explicit explanation.
4. **Tool runtime:** a self-contained local MCP entrypoint or hosted HTTPS MCP/app. Tool discovery, authentication, annotations, sandbox behavior, and actual execution must be smoke-tested independently.
5. **Standalone CLI:** separately distributed `skill-issue` executable with version reporting, plugin install/upgrade/status, `init codex`, `doctor`, and uninstall/repair commands. It should never depend on npm plugin lifecycle scripts.
6. **Initializer payload:** versioned templates for `.codex/agents/*.toml`, `.codex/config.toml` fragments, `.codex/rules/*.rules`, and `AGENTS.md` guidance. The initializer must preview scope, merge safely, preserve unrelated files, record managed files, and support clean uninstall.
7. **Verification contract:** verify configured marketplace, installed/enabled plugin, new-session skill discovery, explicit skill invocation, CLI presence/version, custom-agent registration after restart, MCP/app authentication and tool calls, hook trust state, and policy blockers. File presence alone is insufficient.
8. **Surface fallback:** document plugin-first support for ChatGPT Work/desktop and Codex CLI, plus initializer-managed filesystem skills/config for IDE/project use.

One plugin bundle is therefore **possible for the core experience**, but **insufficient for the entire requested system** under current Codex behavior. The unsupported pieces are not reasons to split the core plugin; they are reasons to define an explicit, auditable bootstrap boundary.

## Notes

- **Caveat — public hooks:** Local plugin authoring docs and current source support plugin hooks, but the public submission page describes only skills, apps, or both and does not explicitly promise hook acceptance. Treat public-directory hook distribution as caveated until verified in the live submission portal or updated first-party policy.
- **Caveat — plugin-level `agents/`:** The official examples repository contains plugin-level agent files, but current Codex custom-agent docs require standalone TOMLs under config-layer `agents/` directories and the local plugin loader does not load custom agents. Do not equate these formats.
- **Caveat — `commands/`:** The client has a compatibility migration path from Markdown commands to skills, but commands are absent from the normative plugin authoring structure. Native Skill Issue releases should author skills directly.
- **Unsupported — automatic CLI registration:** No inspected first-party manifest, installer, or source path registers a plugin-carried executable in `PATH`; npm plugin sources explicitly disable lifecycle scripts during packaging.
- **Unsupported — automatic policy/config installation:** No inspected first-party plugin path installs `AGENTS.md`, `.codex/config.toml`, `.codex/rules`, or standalone agent TOMLs into their active discovery locations.
- Useful search terms for future refreshes: `Codex plugin agents manifest`, `plugin custom agent TOML`, `public plugin hooks submission`, `codex plugin commands migration`, `marketplace requirements allowed_sources`, and `plugin npm source ignore-scripts`.

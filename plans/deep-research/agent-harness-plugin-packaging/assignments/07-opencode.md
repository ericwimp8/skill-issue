# OpenCode Packaging

## Assignment

**Goal:** Determine how current OpenCode extension packaging works and what Skill Issue must implement to deliver skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, and supporting files as one functional experience.

**Scope:** Internet-only research using current official OpenCode documentation, the live OpenCode configuration schema, and the official `anomalyco/opencode` repository. The investigation covers plugins, npm distribution, agents, commands, skills, rules/instructions, references, MCP, custom tools, hooks/events, CLI/configuration, installation, update, discovery, invocation, trust, permissions, project/global scope, and bundle composition.

**Exclusions:** Local OpenCode configuration, third-party behavior as proof of platform contracts, implementation work in Skill Issue, and undocumented claims that could not be validated against current primary sources.

## Sources

- [OpenCode Plugins documentation](https://opencode.ai/docs/plugins/) — official loading, npm installation, dependency, hook, event, and plugin-tool documentation; last updated July 17, 2026.
- [OpenCode CLI documentation](https://opencode.ai/docs/cli/) — official `opencode plugin`, scope flags, `--pure`, agent creation, and upgrade commands; last updated July 17, 2026.
- [OpenCode Config documentation](https://opencode.ai/docs/config/) and [live configuration schema](https://opencode.ai/config.json) — official configuration locations, merge precedence, and current schema for `agent`, `command`, `skills`, `references`, `mcp`, `instructions`, `permission`, and `plugin`.
- [OpenCode Agent Skills documentation](https://opencode.ai/docs/skills) — official skill discovery, frontmatter, invocation, and skill permissions; last updated July 17, 2026.
- [OpenCode Agents documentation](https://opencode.ai/docs/agents/) — official JSON/Markdown agent definitions, global/project locations, and agent permissions.
- [OpenCode Commands documentation](https://opencode.ai/docs/commands/) — official JSON/Markdown slash-command definitions, templating, arguments, shell interpolation, and file references.
- [OpenCode Rules documentation](https://opencode.ai/docs/rules/) — official `AGENTS.md`, custom `instructions`, remote instruction URLs, precedence, and external-file behavior.
- [OpenCode Custom Tools documentation](https://opencode.ai/docs/custom-tools/) — official local/global tool discovery and use of non-JavaScript scripts behind JavaScript/TypeScript tool definitions.
- [OpenCode MCP documentation](https://opencode.ai/docs/mcp-servers/) — official local/remote MCP configuration, OAuth, enablement, tool scoping, and context-cost caveats.
- [OpenCode References documentation](https://opencode.ai/docs/references/) — official local/Git reference configuration, discovery, external-directory treatment, and agent visibility.
- [OpenCode Permissions documentation](https://opencode.ai/docs/permissions/) — official `allow`/`ask`/`deny`, per-tool/per-agent rules, defaults, and external-directory boundary.
- [OpenCode Ecosystem documentation](https://opencode.ai/docs/ecosystem/) — official curated ecosystem list and links to community extension managers.
- Official upstream repository commit [`b8142c7` (July 18, 2026)](https://github.com/anomalyco/opencode/commit/b8142c7aa8f88222873fb79d636e312e28037c2d), especially [`@opencode-ai/plugin` hooks](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/plugin/src/index.ts#L216-L229), [plugin initialization](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/index.ts#L109-L246), [plugin resolution and compatibility](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/shared.ts#L194-L214), [native plugin installer](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/install.ts#L128-L165), [npm cache/install behavior](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/core/src/npm.ts#L72-L137), [skill discovery](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/skill/index.ts#L173-L227), and [skill tool loading](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/tool/skill.ts#L12-L66).

## Findings

### Finding 1: OpenCode has a native npm plugin installation path, plus automatic local plugin discovery

OpenCode supports two server-plugin sources: JavaScript/TypeScript files under `.opencode/plugins/` or `~/.config/opencode/plugins/`, and npm package specifications in the `plugin` array. Current CLI documentation also exposes `opencode plugin <module>` (alias `plug`) to inspect an npm plugin manifest and update either project or global config. Project installation is the default; `--global` writes global configuration, and `--force` replaces an existing package version. The native installer recognizes a server target through `exports["./server"]` or `main`, a TUI target through `exports["./tui"]`, and package themes through `oc-themes` ([CLI](https://opencode.ai/docs/cli/#plugin), [plugins](https://opencode.ai/docs/plugins/#use-a-plugin), [installer source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/install.ts#L128-L165)).

#### Evidence

The plugin docs state that local plugin files are loaded automatically at startup and npm packages are installed automatically. The CLI now provides an explicit config-patching installer with local/global scope. Upstream source resolves project scope to `<worktree>/.opencode` and global scope to the OpenCode config directory, then patches `opencode.json` for a server plugin and `tui.json` for a TUI plugin ([scope source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/install.ts#L333-L343)).

#### Implication

Skill Issue should publish a normal npm package, preferably a scoped and explicitly versioned module such as `@skill-issue/opencode@1.2.3`, with a server entrypoint. The primary install UX can be one command: `opencode plugin @skill-issue/opencode@1.2.3` or its global variant. A copied `.opencode/plugins/*.ts` shim remains useful only for source checkouts or compatibility fallback.

### Finding 2: npm is the native registry, while OpenCode's ecosystem page is curation rather than a marketplace

OpenCode resolves ordinary and scoped npm specifications and uses npm as the package distribution system. The official ecosystem is a documentation list to which projects are added by pull request; it links to community aggregators and identifies `ocx` as a community extension manager. No authoritative source inspected describes a first-party marketplace with ratings, dependency bundles, entitlement, or transactional install/update management ([plugins](https://opencode.ai/docs/plugins/#from-npm), [ecosystem](https://opencode.ai/docs/ecosystem/)).

#### Evidence

The official plugin page tells users to add an npm package name to configuration. The ecosystem page says contributors add projects by PR and separately links community aggregation/management projects. The native CLI accepts a module name rather than a marketplace identifier ([CLI](https://opencode.ai/docs/cli/#plugin)).

#### Implication

Skill Issue discovery should rely on npm metadata, its own website/README, and optionally an OpenCode ecosystem-list PR. It should not design around a native OpenCode marketplace directory. Any integration with `ocx` or other community managers is an additional distribution adapter, not the authoritative OpenCode install path.

### Finding 3: A server plugin can compose tools, hooks, and configuration-backed components

The public plugin API returns a `Hooks` object. Current hooks include `config`, `tool`, auth/provider integration, message/parameter/header hooks, command-before, tool-before/after, shell environment, permission interception, events, compaction, system/message transforms, and tool-definition transforms. The runtime loads configured external plugins, then calls each plugin's `config` hook with the merged mutable configuration before dependent runtime surfaces consume it ([public API](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/plugin/src/index.ts#L216-L309), [runtime call](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/index.ts#L237-L246)).

#### Evidence

Official plugin documentation explicitly supports event hooks and plugin-provided custom tools. The live schema exposes config keys for agents, commands, skills paths/URLs, references, MCP, instructions, permissions, and plugins ([schema](https://opencode.ai/config.json)). The public `Hooks` type has no first-class `agent`, `command`, or `skill` registration member; configuration mutation is the available composition point.

#### Implication

Skill Issue's server plugin should return its executable tools and lifecycle hooks directly, and in `config` should merge, rather than replace, bundle contributions into `agent`, `command`, `skills.paths`, `references`, `instructions`, and optional `mcp`. Because config-hook composition is source-backed but not explained on the plugin documentation page, the package should declare an `engines.opencode` minimum and run release integration tests against that supported range. The loader enforces `engines.opencode` compatibility for stable semver releases ([compatibility source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/shared.ts#L194-L205)).

### Finding 4: Skills can remain complete directory bundles, including scripts, references, assets, and supporting files

OpenCode natively discovers `SKILL.md` directories in `.opencode/skills`, `~/.config/opencode/skills`, `.agents/skills`, and Claude-compatible locations. The live schema additionally supports `skills.paths` and `skills.urls`. When the native skill tool loads a skill, it reports the skill's base directory, states that relative `scripts/` and `reference/` paths resolve from that base, and samples adjacent files for the model ([skills docs](https://opencode.ai/docs/skills), [schema](https://opencode.ai/config.json), [skill tool source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/tool/skill.ts#L34-L60)).

#### Evidence

The current skill discovery implementation scans configured directories from `cfg.skills.paths` and fetched directories from `cfg.skills.urls`, in addition to standard project/global locations ([discovery source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/skill/index.ts#L173-L227)). Skills are advertised by name/description and loaded on demand through the permission-aware `skill` tool. Only `name`, `description`, `license`, `compatibility`, and string-map `metadata` frontmatter are documented; unknown fields are ignored ([skills docs](https://opencode.ai/docs/skills/#write-frontmatter)).

#### Implication

The npm package should preserve each Skill Issue skill as a directory such as `dist/skills/<name>/SKILL.md` with its own `scripts/`, `references/`, `assets/`, and supporting files. The plugin should append the absolute packaged `dist/skills` directory to `config.skills.paths`. This avoids copying skill trees and preserves relative links. Skill-level automatic hooks, declared tool allowlists, or executable frontmatter should be treated as unsupported; scripts execute only when instructions/tools invoke them under normal permissions.

### Finding 5: Agents, commands, and rules have native project/global forms and can also be injected through config

Agents can be JSON config entries or Markdown files in `.opencode/agents/` and `~/.config/opencode/agents/`; the Markdown filename becomes the agent name. Commands have the same JSON-or-Markdown pattern under `commands/`, are invoked as slash commands, and support arguments, shell-output interpolation, and file references ([agents](https://opencode.ai/docs/agents/#markdown), [commands](https://opencode.ai/docs/commands/)). Rules use project/global `AGENTS.md`, while the `instructions` array can add paths, globs, and remote URLs; instruction files are combined with applicable `AGENTS.md` content ([rules](https://opencode.ai/docs/rules/#custom-instructions)).

#### Evidence

Configuration is merged across remote, global, custom, project, `.opencode`, inline, and managed sources, with later/higher-priority sources overriding conflicts while preserving non-conflicting keys ([config precedence](https://opencode.ai/docs/config/#precedence-order)). The live schema accepts `agent`, `command`, and `instructions`, and the plugin API's config hook receives that schema type.

#### Implication

For the one-package path, Skill Issue should ship Markdown sources as package assets but register agents and commands through config objects, and append instruction files to `instructions`. This avoids mutating the user's repository. It must namespace identifiers (for example `skill-issue-review`) and merge arrays/maps idempotently so user and organizational configuration retains precedence. A fallback installer can materialize the same Markdown under `.opencode/agents`, `.opencode/commands`, and an instruction path when supporting older OpenCode versions.

### Finding 6: References are the native way to expose substantial supporting corpora outside the project

OpenCode references can point to local directories or Git repositories, appear in `@` autocomplete, and can be described so agents know when to inspect them. Reference directories are automatically admitted through the external-directory boundary, while normal read/edit permissions still apply ([references](https://opencode.ai/docs/references/)).

#### Evidence

The live schema provides `references` entries with a local `path` or Git `repository`, plus `branch`, `description`, and `hidden`. Official documentation says described references and resolved paths enter agent context and configured reference directories receive automatic external-directory allowance without gaining broader tool rights.

#### Implication

Skill Issue should use a named local reference pointing at its package-bundled documentation/examples when those materials are too broad for a skill's immediate directory. Small, skill-specific files should remain beside `SKILL.md`; larger shared corpora should be exposed as a namespaced reference with a narrow description. This separates discoverable shared source material from always-loaded instructions.

### Finding 7: Plugin tools are usually a better bundle primitive than MCP for in-package functions

Plugins can return custom tools directly, and project/global custom-tool files can invoke scripts in any language. MCP is appropriate when Skill Issue already has a protocol server, needs cross-client interoperability, or needs remote OAuth. MCP tools are automatically added to model context and can be globally/per-agent disabled, but official docs warn that many MCP tools consume substantial context ([plugin tools](https://opencode.ai/docs/plugins/#custom-tools), [custom tools](https://opencode.ai/docs/custom-tools/), [MCP caveat](https://opencode.ai/docs/mcp-servers/#caveats)).

#### Evidence

Local MCP config is a command array with optional cwd/environment/enabled/timeout, while remote MCP supports URL, headers, OAuth, and stored authentication. OpenCode prefixes MCP tools by server name and supports wildcard permission/tool rules ([MCP](https://opencode.ai/docs/mcp-servers/)). Plugin tools are in-process and can override built-ins on a name collision.

#### Implication

Skill Issue should expose core package functions as namespaced plugin tools and call bundled script/library code directly. It should add MCP only for genuinely protocol-level capabilities, default optional servers to disabled, and scope tools per agent. If a bundled local MCP is necessary, the config hook can supply an absolute package entrypoint; a remote MCP should use OpenCode's native OAuth rather than custom credential storage.

### Finding 8: Permissions protect agent actions, while plugins themselves are trusted in-process code

OpenCode permissions resolve actions to `allow`, `ask`, or `deny`, can be globally or per-agent scoped, and separately guard skill loading and external-directory access. Most actions default to allow; `doom_loop` and external-directory access default to ask. External plugins can be disabled for a run with `--pure` ([permissions](https://opencode.ai/docs/permissions/), [CLI flags](https://opencode.ai/docs/cli/#global-flags)).

#### Evidence

The plugin module is dynamically imported and executed in the OpenCode process. Its context includes a client, project/directory/worktree information, and Bun's shell API ([plugin docs](https://opencode.ai/docs/plugins/#basic-structure)). The npm installer uses `ignoreScripts: true`, so npm lifecycle scripts do not execute during OpenCode's package installation, but the installed plugin entrypoint itself still executes at startup ([npm source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/core/src/npm.ts#L80-L108)). No inspected official source documents a per-plugin capability sandbox or installation trust prompt.

#### Implication

Skill Issue must treat its OpenCode plugin as trusted executable code: keep startup side effects minimal, avoid writing user files automatically, disclose hooks/tools/config contributions, and provide a documented `--pure` troubleshooting path. Agent-facing tools should use explicit namespaced permission keys and request approval for destructive/external effects. npm `postinstall` cannot be the provisioning mechanism.

### Finding 9: Plugin update semantics favor explicit version pins and deliberate config replacement

OpenCode's native plugin command can replace an existing plugin version with `--force`. The loader parses npm package/version specifications, caches packages by the full spec, and returns an already-present package without reinstalling it. OpenCode's `autoupdate` and `opencode upgrade` concern the OpenCode application, not a documented plugin-update lifecycle ([CLI plugin](https://opencode.ai/docs/cli/#plugin), [CLI upgrade](https://opencode.ai/docs/cli/#upgrade), [npm cache source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/core/src/npm.ts#L115-L137)).

#### Evidence

Unversioned packages are normalized to `latest` for resolution, but the cache short-circuits when that specification's package directory already exists. Changing a pinned specification produces a distinct cache directory; the native installer can update the config entry with `--force` ([specifier source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/shared.ts#L22-L42), [config replacement](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/plugin/install.ts#L181-L257)).

#### Implication

Skill Issue should document deterministic installs and upgrades with explicit versions: install `@skill-issue/opencode@X`, then upgrade using `opencode plugin @skill-issue/opencode@Y --force`. It should not promise that an unpinned package refreshes on every OpenCode start. Package migrations should be config-merge compatible and avoid one-time lifecycle scripts.

### Finding 10: One bundled OpenCode experience is achievable, with two important boundaries

On the current OpenCode surface, one npm server plugin can activate a coherent bundle: direct plugin tools/hooks, config-injected agents and commands, packaged skill directories through `skills.paths`, rule files through `instructions`, supporting corpora through `references`, and optional MCP configuration. A user can install that package locally or globally through one native command.

#### Evidence

The conclusion is the intersection of the live config schema, public `config`/`tool` plugin hooks, skill-path discovery, package-relative skill loading, native local/global plugin installer, and config merge semantics cited above. The public hook type does not expose native skill/agent/command registries, so those components cooperate through configuration rather than package manifest entries.

#### Implication

Skill Issue should implement an OpenCode adapter package with this shape:

1. `package.json` with a server entrypoint, `engines.opencode`, published bundle files, and optionally a `bin` entry.
2. A side-effect-light server plugin that idempotently merges bundle paths/config and returns namespaced tools/hooks.
3. `skills/<name>/SKILL.md` directories containing their scripts, references, assets, and support files.
4. Package-resolved instruction and reference paths, namespaced agents/commands, and optional disabled-by-default MCP.
5. Version-pinned install/upgrade commands plus a compatibility/bootstrap CLI for older OpenCode releases.

The first boundary is the standalone shell CLI: OpenCode caches npm plugin packages internally and does not document adding a package's `bin` to the user's shell `PATH`. A true standalone `skill-issue` command therefore still needs a normal npm/binary install or an `npx @skill-issue/cli` invocation, even if the plugin and CLI share one published package. The second boundary is compatibility: older releases lacking the current `opencode plugin` command, config hook behavior, `skills.paths`, or references need the Skill Issue installer to copy/materialize `.opencode` files and patch config. Those fallbacks should be version-gated rather than attempted by npm lifecycle scripts.

### Finding 11: Removal and conflict handling require explicit Skill Issue UX

The current CLI documentation provides plugin installation and whole-OpenCode uninstallation, but it does not document a plugin-specific remove command. Duplicate npm packages are deduplicated by package identity, whereas local and npm plugins with similar names can both load. Plugin tools and custom tools can override built-ins on name collision ([plugins load order](https://opencode.ai/docs/plugins/#load-order), [CLI](https://opencode.ai/docs/cli/), [custom tool collisions](https://opencode.ai/docs/custom-tools/#name-collisions-with-built-in-tools)).

#### Evidence

Plugin load order is global config, project config, global plugin directory, then project plugin directory. Config sources merge rather than replace one another. Current source deduplicates npm plugin origins by package name while retaining the higher-precedence occurrence ([dedupe source](https://github.com/anomalyco/opencode/blob/b8142c7aa8f88222873fb79d636e312e28037c2d/packages/opencode/src/config/plugin.ts#L62-L77)).

#### Implication

Skill Issue must namespace all registered identifiers, make config changes idempotent, and document manual removal from project/global plugin arrays. Its optional installer/CLI should offer a safe uninstall that removes only Skill Issue-owned config entries or materialized files and preserves user modifications. It should detect simultaneous legacy local shims and npm installation to avoid double registration.

## Notes

- The public documentation and live schema were current on July 17, 2026; upstream source validation used official commit `b8142c7aa8f88222873fb79d636e312e28037c2d` from July 18, 2026. The config-hook-based bundle design should be guarded with `engines.opencode` and tested against released versions because the plugin page documents hooks generally but does not describe config-time component injection.
- Unsupported: a first-party OpenCode marketplace, plugin-specific capability sandbox, documented plugin-removal command, npm lifecycle-script provisioning, automatic shell-PATH exposure of plugin `bin` entries, and advanced executable skill frontmatter were not established by inspected authoritative sources.
- Caveat: `skills.urls` is present in the live official schema and upstream discovery source but is absent from the current Agent Skills guide. Package-local `skills.paths` is the more deterministic bundled implementation.
- Caveat: command templates can interpolate shell output. Skill Issue commands should avoid implicit shell execution unless it is essential and clearly disclosed, because invocation can execute commands before the prompt is sent.

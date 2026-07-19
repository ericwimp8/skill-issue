# Google Antigravity and Gemini CLI Packaging

## Assignment

**Goal:** Determine the current product surfaces and packaging mechanisms in the Google Antigravity / Gemini CLI lineage, and identify what Skill Issue must implement to deliver skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, and supporting files as a functional bundle.

**Scope:** Internet-only research using current official Google Antigravity documentation, official Google developer announcements, official Gemini CLI documentation, the official Gemini CLI repository, and the official Gemini CLI extension catalogue. The analysis distinguishes Antigravity 2.0, Antigravity CLI, and the still-supported enterprise/API-key Gemini CLI surface as of 2026-07-19.

**Exclusions:** Local product configuration, runtime experiments, third-party packaging claims, unofficial reverse engineering, and assumptions that undocumented files, registries, or lifecycle behavior are supported.

## Sources

- Google Developers Blog, [“An important update: Transitioning Gemini CLI to Antigravity CLI”](https://developers.googleblog.com/an-important-update-transitioning-gemini-cli-to-antigravity-cli/), published 2026-05-19.
- Google Antigravity Blog, [“Google Antigravity CLI”](https://antigravity.google/blog/introducing-google-antigravity-cli), published 2026-05-19.
- Google Antigravity documentation, [Antigravity 2.0 overview](https://antigravity.google/docs/overview), [Antigravity CLI overview](https://antigravity.google/docs/cli-overview), [Plugins](https://antigravity.google/docs/plugins), [Antigravity CLI plugins and skills](https://antigravity.google/docs/cli-plugins), [Gemini CLI migration](https://antigravity.google/docs/gcli-migration), [Agent Skills](https://antigravity.google/docs/skills), [Rules and workflows](https://antigravity.google/docs/rules-workflows), [Hooks](https://antigravity.google/docs/hooks), [CLI permissions](https://antigravity.google/docs/cli-permissions), [CLI sandbox](https://antigravity.google/docs/cli-sandbox), and [Build with Google](https://antigravity.google/docs/build-with-google), accessed 2026-07-19.
- Official `google-gemini/gemini-cli` repository, [project README](https://github.com/google-gemini/gemini-cli), [extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md), [Agent Skills](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/skills.md), [subagents](https://github.com/google-gemini/gemini-cli/blob/main/docs/core/subagents.md), [hooks](https://github.com/google-gemini/gemini-cli/blob/main/docs/hooks/index.md), and [policy engine](https://github.com/google-gemini/gemini-cli/blob/main/docs/reference/policy-engine.md), `main` branch inspected 2026-07-19.
- Gemini CLI documentation, [extensions](https://google-gemini.github.io/gemini-cli/docs/extensions/), [extension getting started](https://google-gemini.github.io/gemini-cli/docs/extensions/getting-started-extensions.html), [extension releasing](https://google-gemini.github.io/gemini-cli/docs/extensions/extension-releasing.html), [configuration](https://google-gemini.github.io/gemini-cli/docs/get-started/configuration.html), [MCP servers](https://google-gemini.github.io/gemini-cli/docs/tools/mcp-server.html), [trusted folders](https://google-gemini.github.io/gemini-cli/docs/cli/trusted-folders.html), [sandboxing](https://google-gemini.github.io/gemini-cli/docs/cli/sandbox.html), and [IDE integration](https://google-gemini.github.io/gemini-cli/docs/ide-integration/), accessed 2026-07-19.
- Gemini CLI official catalogue, [Browse Extensions](https://geminicli.com/extensions/), inspected 2026-07-19.

## Findings

### Finding 1: Antigravity is the primary consumer successor, while Gemini CLI remains a narrower supported surface

Google announced that its consumer Gemini CLI and Gemini Code Assist IDE requests would stop on 2026-06-18 for free, Google AI Pro, and Google AI Ultra users. Gemini CLI remains supported for Gemini Code Assist Standard/Enterprise customers and users authenticating with paid Gemini or Gemini Enterprise Agent Platform API keys. Antigravity CLI is the new terminal surface in the unified Antigravity product, rather than a rename of the open-source Gemini CLI binary. The Antigravity team says the new CLI took inspiration from Gemini CLI components and provides migration rather than binary or package identity. [Google’s transition announcement](https://developers.googleblog.com/an-important-update-transitioning-gemini-cli-to-antigravity-cli/) and the [Antigravity CLI launch post](https://antigravity.google/blog/introducing-google-antigravity-cli) agree on this boundary.

**Evidence:** The transition announcement gives the 2026-06-18 consumer cutoff, preserves enterprise and paid-API-key access, and describes Antigravity CLI as a new Go-based terminal experience sharing Antigravity’s harness. The active [Gemini CLI repository](https://github.com/google-gemini/gemini-cli) and documentation still publish the open-source client and extension APIs.

**Implication:** Skill Issue should treat Antigravity as the default current consumer target and Gemini CLI as a separate compatibility target for enterprise/API-key users. It should label artifacts explicitly (`antigravity-plugin` versus `gemini-cli-extension`) and avoid presenting one as an alias of the other.

### Finding 2: Antigravity 2.0 and Antigravity CLI share an agent harness, but their documented installation locations differ

Antigravity 2.0 is the visual standalone desktop application; Antigravity CLI is its terminal-first companion. Google states that they run the same agent core and share core settings and permissions. The general Antigravity plugin documentation discovers workspace plugins under `.agents/plugins/` (also `_agents/plugins/`) and global plugins under `~/.gemini/config/plugins/`. The CLI-specific documentation stages installed plugins under `~/.gemini/antigravity-cli/plugins/<plugin_name>/`. [Antigravity CLI overview](https://antigravity.google/docs/cli-overview), [Antigravity 2.0 overview](https://antigravity.google/docs/overview), [general plugin documentation](https://antigravity.google/docs/plugins), and [CLI plugin documentation](https://antigravity.google/docs/cli-plugins) establish these surfaces and paths.

**Evidence:** The CLI overview explicitly identifies a shared harness/settings relationship. The two plugin documents use the same `plugin.json`-centred structure, while documenting different discovery or staging paths. Google’s docs do not explicitly promise that installing a plugin into one of those paths makes it available in every Antigravity surface.

**Implication:** Skill Issue can generate one Antigravity plugin format, but its installer must target the documented path for the chosen surface or install into both paths when the user requests both. Cross-surface availability from a single copy is unsupported and should not be claimed without a Google-documented sync rule or a live validation outside this assignment.

### Finding 3: The Antigravity native bundle is a `plugin.json` package with a defined but narrower component set

An Antigravity plugin is a namespaced directory with a mandatory `plugin.json`. The common layout supports `skills/`, `rules/`, root `mcp_config.json`, and root `hooks.json`; the CLI-specific layout additionally documents `agents/`. The CLI schema gives `name` and optional `description` and rejects additional manifest properties. [Antigravity plugins](https://antigravity.google/docs/plugins) and [Antigravity CLI plugins and skills](https://antigravity.google/docs/cli-plugins) define the contract.

**Evidence:** Both official pages show the same root manifest and component directories/files. The CLI page publishes the schema URL `https://antigravity.google/schemas/v1/plugin.json` and a manifest schema with `additionalProperties: false`. The CLI feature documentation also describes installed components as one staged, discoverable unit.

**Implication:** Skill Issue needs an Antigravity adapter that emits `plugin.json`, converts persistent Markdown guidance to `rules/*.md`, emits skills under `skills/<name>/SKILL.md`, emits subagents under `agents/` only where the Antigravity CLI format is applicable, translates MCP configuration to `mcp_config.json`, and translates hooks to root `hooks.json`. Package metadata, standalone executable registration, dependencies, themes, and arbitrary configuration cannot be placed into invented manifest fields.

### Finding 4: Antigravity can carry scripts, references, and assets through skills, but has no documented standalone-CLI registration primitive

Antigravity skills follow the Agent Skills directory model: `SKILL.md` is required, while `scripts/`, `examples/`, and `resources/` may accompany it and are read as needed. Skills are progressively disclosed from name/description to full instructions and supporting files, and users can invoke a skill by name. Rules are Markdown guidance with activation modes, while workflows are Markdown sequences exposed as slash commands. [Agent Skills](https://antigravity.google/docs/skills) and [Rules and workflows](https://antigravity.google/docs/rules-workflows) define these behaviors.

**Evidence:** The skill documentation explicitly permits helper scripts and resources and describes discovery, contextual activation, and execution. The plugin manifest has no `bin`, dependency, post-install, or lifecycle-script field, and the documented component list has no standalone CLI component.

**Implication:** Skill Issue can bundle its CLI binary or script as a supporting file and have a skill, hook, MCP server, or rule instruct the agent to invoke it by relative path. To make a command independently available on the user’s shell `PATH`, install dependencies, set executable bits, or write user-specific configuration, Skill Issue needs its own bootstrap installer/CLI. That bootstrap should leave the native plugin declarative and use platform-specific installation logic only for the unsupported shell-level concerns.

### Finding 5: Antigravity third-party installation exists, but public distribution and update semantics are incompletely documented

Antigravity CLI documents `agy plugin list`, `install`, `disable`, `enable`, and `uninstall`; the install description says local or remote, but its only syntax example is a local path. Antigravity 2.0 documents manual placement and a UI catalogue for Google-created “Build with Google” bundles. No inspected official page defines a general third-party marketplace submission flow, a remote URL grammar, signed package format, version selection, lockfile, plugin update command, or auto-update policy. [CLI plugin management](https://antigravity.google/docs/cli-plugins), [general plugin installation](https://antigravity.google/docs/plugins), and [Build with Google](https://antigravity.google/docs/build-with-google) are the authoritative public coverage found.

**Evidence:** The CLI page lists management commands but omits `plugin update` and a concrete remote source example. The desktop page describes filesystem discovery and Google-curated UI bundles, not a public third-party registry. Antigravity CLI’s own self-updater is product-update behavior and does not establish plugin-update behavior.

**Implication:** Skill Issue should distribute an Antigravity plugin directory or archive plus an installer that copies/upgrades it deterministically. Repository release tags and checksums can be managed by Skill Issue, but automatic Antigravity-native plugin updates and public marketplace availability must be classified as unsupported until Google publishes those contracts.

### Finding 6: Antigravity plugin execution remains subject to runtime permissions, sandboxing, and hook gates

Antigravity CLI evaluates sensitive operations through `deny`, `ask`, and `allow` permission lists, with deny taking precedence and unconfigured commands, MCP calls, browser actions, and non-workspace file access generally prompting. Its native sandbox can constrain commands at the OS level. Hooks can inspect or gate model/tool execution and may return allow, deny, or ask decisions. Antigravity 2.0 likewise scopes settings and permissions by project and defaults agent filesystem access to project folders. [CLI permissions](https://antigravity.google/docs/cli-permissions), [CLI sandbox](https://antigravity.google/docs/cli-sandbox), [hooks](https://antigravity.google/docs/hooks), and [Antigravity settings](https://antigravity.google/docs/settings) document these controls.

**Evidence:** Official permission examples include command, file, URL, and MCP resources; sandbox documentation requires explicit configuration and supports per-command overrides; hook output can control whether a proposed call proceeds.

**Implication:** Skill Issue must describe required permissions and let Antigravity request them at runtime. The bundle should not try to silently grant itself execution, filesystem, network, or MCP trust. Hooks and scripts should use relative paths, minimal privileges, bounded timeouts, and clear failure behavior.

### Finding 7: Gemini CLI extensions are richer packages with explicit install, update, configuration, and conflict behavior

A Gemini CLI extension is a directory installed under `~/.gemini/extensions/<name>/` with mandatory `gemini-extension.json`. Current extension support includes MCP servers, a context file such as `GEMINI.md`, custom TOML commands, hooks under `hooks/hooks.json`, Agent Skills under `skills/`, preview subagents under `agents/`, policy files under `policies/`, themes, and installer-supplied settings. Extension configurations merge at startup; workspace configuration wins, user/project commands outrank extension commands, and conflicting extension commands are namespaced. [Gemini CLI extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md) is the most complete current contract.

**Evidence:** The reference enumerates the manifest fields, component directories, `${extensionPath}`/`${workspacePath}` variables, settings storage, environment sanitization, policy tiering, and command conflict rules. The [getting-started guide](https://google-gemini.github.io/gemini-cli/docs/extensions/getting-started-extensions.html) demonstrates an MCP server, a custom command, and persistent `GEMINI.md` context cooperating in one extension.

**Implication:** Skill Issue needs a Gemini adapter that emits `gemini-extension.json`; maps general Markdown guidance to the extension `GEMINI.md`; maps invocable prompts to `commands/*.toml`; includes `skills/`, preview `agents/`, `hooks/hooks.json`, and optional `policies/`; and starts bundled MCP/CLI code with `${extensionPath}`. References, templates, assets, and supporting files can remain beside the components that consume them.

### Finding 8: Gemini CLI has a mature Git/GitHub distribution channel and explicit update controls

Users install extensions from a GitHub URL or local path with `gemini extensions install`; installation copies the source into the extension directory. They can select a branch, tag, or commit with `--ref`, enable automatic updates, update one or all extensions, disable/enable globally or per workspace, uninstall, configure settings, or link a development checkout. Git repository releases and GitHub Releases are the two documented publication methods, with GitHub Releases able to carry platform-specific prebuilt archives. [Extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md) and [extension releasing](https://google-gemini.github.io/gemini-cli/docs/extensions/extension-releasing.html) describe the lifecycle.

**Evidence:** The CLI exposes `install`, `update`, `uninstall`, `enable`, `disable`, `config`, and `link`; the release guide describes ref-based channels and platform-specific release assets. The official [extension catalogue](https://geminicli.com/extensions/) provides repository install commands and warns that third-party entries are not vetted or guaranteed by Google.

**Implication:** Skill Issue can publish its Gemini artifact from a public Git repository, preferably with immutable tags and GitHub Release archives containing prebuilt platform artifacts. It can rely on Gemini’s update mechanism for files inside the extension, while its own bootstrap remains responsible for any state outside the copied extension directory.

### Finding 9: Gemini trust and activation are explicit at skill, folder, tool, and policy layers

Extension installation accepts `--consent` to skip the security prompt, which implies interactive confirmation by default. Sensitive extension settings are stored through the system keychain; extension/MCP processes receive only safe environment variables plus values explicitly declared in the manifest. MCP servers contributed by extensions cannot set `trust`. Extension policies may ask or deny, but Gemini ignores extension-supplied allow/yolo decisions. Agent Skill activation separately prompts with the skill name, purpose, and directory before granting access to bundled assets. Trusted-folder mode can suppress workspace configuration and extension management in untrusted projects. [Extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md), [Agent Skills](https://github.com/google-gemini/gemini-cli/blob/main/docs/cli/skills.md), and [Trusted Folders](https://google-gemini.github.io/gemini-cli/docs/cli/trusted-folders.html) establish the layers.

**Evidence:** Official docs describe security consent, environment allowlisting, keychain-backed sensitive values, skill activation consent, untrusted safe mode, and policy limitations. [Gemini CLI sandboxing](https://google-gemini.github.io/gemini-cli/docs/cli/sandbox.html) adds optional Seatbelt or container isolation.

**Implication:** Skill Issue should declare only the environment variables it needs, avoid embedding secrets, omit MCP `trust`, keep policies restrictive, and explain why each skill/tool needs access. Installation automation should not pass `--consent` unless the user explicitly requested non-interactive installation after reviewing the source.

### Finding 10: Gemini CLI subagents and IDE integration are separate capabilities with different maturity

Gemini extensions may include Markdown subagent definitions in `agents/`, but the official extension reference marks subagents preview. Subagents are invoked automatically or by an `@name` prompt and run with specialized prompts/tool sets. Gemini CLI’s IDE feature is instead a separate Gemini CLI Companion extension for VS Code-compatible editors; it supplies workspace/cursor context and native diff review and is installed via an IDE prompt, `/ide install`, VS Code Marketplace, or Open VSX. [Subagents](https://github.com/google-gemini/gemini-cli/blob/main/docs/core/subagents.md), [extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md), and [IDE integration](https://google-gemini.github.io/gemini-cli/docs/ide-integration/) distinguish them.

**Evidence:** The extension contract includes `agents/` but labels the feature preview. IDE documentation describes a companion protocol and marketplace-distributed editor extension, not a nested component of arbitrary Gemini CLI extensions.

**Implication:** Skill Issue may ship optional Gemini subagents behind a preview compatibility label. It should not attempt to package its own functionality into the official Gemini CLI Companion or imply that installing the Skill Issue extension installs editor integration; the user installs/enables the companion separately.

### Finding 11: One source repository can serve both products, but one native package cannot be assumed portable

The two formats overlap conceptually—Agent Skills, Markdown guidance, MCP servers, hooks, and subagents—but differ in manifests, component paths, schemas, command models, policy systems, settings, MCP URL keys, and installation roots. Antigravity provides `agy plugin import gemini`, converting Gemini commands to skills and MCP configuration to `mcp_config.json`; its migration documentation explicitly notes partial parity and skill-path changes. [Gemini migration](https://antigravity.google/docs/gcli-migration) is direct evidence that conversion is required.

**Evidence:** Gemini uses `gemini-extension.json`, inline `mcpServers`, `hooks/hooks.json`, `commands/`, `policies/`, and extension settings. Antigravity uses `plugin.json`, root `mcp_config.json`, root `hooks.json`, `rules/`, and different discovery/staging paths. Remote MCP URL keys also change from Gemini’s `url`/`httpUrl` to Antigravity’s `serverUrl`.

**Implication:** Skill Issue should maintain a canonical bundle model and generate two target artifacts from it. Shared skill content, scripts, references, and assets can be copied into both outputs, while manifests, hooks, MCP config, persistent guidance, commands/workflows, policies, and install metadata must be rendered per target. A single repository may publish both artifacts, but a single directory containing both manifests and both path conventions is only a speculative convenience because neither product documents how it treats the other product’s extra files.

### Finding 12: Recommended Skill Issue implementation is a dual-adapter build plus a fallback installer

The smallest complete implementation is: (1) a canonical source bundle; (2) an Antigravity renderer; (3) a Gemini CLI renderer; (4) platform release artifacts for any executable; and (5) a bootstrap installer that detects requested surfaces, presents trust/permission requirements, installs the correct native artifact, and records ownership for upgrades/uninstall. Native components should cooperate through stable relative paths and an MCP or subprocess boundary rather than duplicating logic.

**Evidence:** Both products support skills with supporting files and MCP servers, which provides a common execution seam. Their native package managers own different directories and lifecycle behavior, while neither manifest registers a general shell executable. Gemini supports tagged/ref updates; Antigravity’s public plugin update contract remains undocumented.

**Implication:** For Antigravity, the installer should validate `plugin.json`, copy or invoke `agy plugin install` for the chosen surface, place the standalone CLI in a user-approved executable location, and manage updates itself. For Gemini CLI, it should prefer `gemini extensions install <repo-or-release>` and use declared settings/keychain integration, with the standalone CLI invoked by `${extensionPath}` when feasible. Both adapters should include a diagnostic command/skill that checks executable presence, MCP startup, required configuration, and permissions without auto-granting trust.

## Notes

- Unsupported: No inspected official Antigravity source documents a public third-party marketplace submission process, package signing/notarization contract, remote plugin source syntax, plugin version-resolution algorithm, or plugin auto-update command.
- Caveat: Antigravity’s general and CLI-specific docs show compatible-looking plugin layouts but different global install roots. Shared-harness language validates common agent behavior, not automatic cross-surface plugin-file synchronization.
- Caveat: Gemini CLI remains technically and commercially relevant for enterprise and paid API-key users, but it is no longer the normal consumer surface after 2026-06-18.
- Unsupported: Neither product documents a native manifest field that installs a bundled standalone executable onto the user’s shell `PATH`; this requires a Skill Issue-owned installer or invocation from within the bundle directory.
- Useful search terms: `site:antigravity.google/docs cli-plugins`, `agy plugin install`, `agy plugin import gemini`, `gemini-extension.json`, `Gemini CLI extension releasing`, `Gemini CLI Agent Skills`, `Gemini CLI policy engine`.

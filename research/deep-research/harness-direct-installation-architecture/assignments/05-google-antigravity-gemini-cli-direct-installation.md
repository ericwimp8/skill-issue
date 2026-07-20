# Google Antigravity and Gemini CLI Direct Installation

## Assignment

**Goal:** Establish the supported direct filesystem lifecycle for installing a Skill Issue Agent Skill into Google Antigravity and Gemini CLI, without requiring a native plugin or extension package.

**Scope:** Current first-party documentation for Antigravity 2.0, Antigravity CLI (`agy`), and Gemini CLI. Covers direct placement, discovery, activation, verification, scope, collisions, update/removal, and security boundaries.

**Exclusions:** Packaging and publishing a Skill Issue plugin/extension; non-Google community installation tools; claims based on local product files or community reports. Antigravity 2.0 and Antigravity CLI are treated as distinct surfaces where their documented paths differ.

## Sources

- [Google Antigravity: Agent Skills](https://antigravity.google/docs/skills) — current first-party Antigravity 2.0 skill layout, metadata, scopes, progressive disclosure, and `.agent` compatibility.
- [Google Antigravity: Plugins](https://www.antigravity.google/docs/plugins) — current first-party distinction between a plain skill and optional plugin package, plus plugin discovery locations.
- [Google Antigravity: Settings](https://www.antigravity.google/docs/settings) — first-party Antigravity 2.0 global/project customization and permission surfaces.
- [Google Antigravity CLI: Overview](https://antigravity.google/docs/cli-overview) and [Migration from Gemini CLI](https://antigravity.google/docs/gcli-migration) — first-party Antigravity CLI scope/path migration and compatibility boundary.
- [Google Antigravity CLI: Reference](https://antigravity.google/docs/cli-reference), [Permissions](https://www.antigravity.google/docs/cli-permissions), and [Tutorial](https://antigravity.google/docs/cli-tutorial) — first-party `agy` discovery UI, permission policy, trust/review workflow, and session entry point.
- [Gemini CLI: Agent Skills](https://geminicli.com/docs/cli/skills/), [Creating Agent Skills](https://geminicli.com/docs/cli/creating-skills/), and [Managing Agent Skills](https://geminicli.com/docs/cli/using-agent-skills/) — current first-party direct skill layout, tiers, activation, management, and consent.
- [Gemini CLI: Trusted Folders](https://geminicli.com/docs/cli/trusted-folders/), [Settings](https://geminicli.com/docs/cli/settings/), and [release notes](https://geminicli.com/docs/changelogs/) — trust effects, `skills.enabled`, and skills feature history/current channels.
- [Gemini CLI: Extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md) — first-party repository documentation used only to delimit the optional extension lifecycle.

## Findings

### Direct skill installation is the common, supported baseline

Antigravity 2.0 and Gemini CLI both document a skill as a self-contained folder containing `SKILL.md`; a native package is not required for direct placement. Antigravity plugins and Gemini extensions can bundle skills with other capabilities, but they add manifest/package lifecycle rather than enabling the basic skill mechanism. [Antigravity skills](https://antigravity.google/docs/skills) specifies the folder-plus-`SKILL.md` model, and [Gemini’s creation guide](https://geminicli.com/docs/cli/creating-skills/) documents manual directory creation and discovery.

**Evidence:** Antigravity says `SKILL.md` is the only required file and explicitly supports direct workspace/global skill folders. Gemini CLI documents `SKILL.md` as required, manual creation under `.gemini/skills`, and optional bundled resources. In contrast, an [Antigravity plugin](https://www.antigravity.google/docs/plugins) requires `plugin.json`, while a [Gemini extension](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md) requires `gemini-extension.json`.

**Implication:** Use a common Skill Issue skill directory as the primary artifact. A direct-install adapter only needs to place the complete directory in each surface’s supported skill root; it must not manufacture a plugin/extension manifest solely to make the skill discoverable.

### Antigravity 2.0 direct-install contract

For the Antigravity 2.0 desktop surface, place one complete skill directory at either `<workspace-root>/.agents/skills/<skill-folder>/` for workspace scope or `~/.gemini/config/skills/<skill-folder>/` for global scope. Antigravity retains backward compatibility for workspace `.agent/skills`, but `.agents/skills` is its documented default. [Agent Skills](https://antigravity.google/docs/skills) is the controlling source for these paths.

**Evidence:** The Antigravity skill documentation names those two exact locations and scopes, requires `SKILL.md` with YAML frontmatter, and states that `description` is required while `name` defaults to the folder name. It also documents optional `scripts/`, `examples/`, and `resources/` subdirectories and says the agent can read them when following the skill.

**Implication:** Copy the whole Skill Issue directory unchanged into a unique `<skill-folder>` at the selected root; preserve all supporting files rather than flattening only `SKILL.md`. Prefer workspace scope for repository-owned workflow instructions and global scope for a personal Skill Issue installation. Do not place an Antigravity 2.0 global skill in the Antigravity CLI global root, or vice versa.

### Antigravity 2.0 discovery, activation, and verification

Antigravity describes a conversation-start lifecycle: it discovers available skill names/descriptions when a conversation starts, reads the full `SKILL.md` when relevant, then follows the skill during execution. The desktop Settings page exposes global/project Customizations and can show skills originating from project folders. [Agent Skills](https://antigravity.google/docs/skills) and [Settings](https://www.antigravity.google/docs/settings) support this lifecycle.

**Evidence:** The skills page defines Discovery, Activation, and Execution in that order and allows a user to mention a skill by name. Settings distinguishes global and project customizations and says project settings can show skills originating from each added folder.

**Implication:** After direct placement, start a new Antigravity conversation in the intended project, confirm the skill appears through Customizations, then issue a request matching its `description` (or name it) and confirm the agent uses it. The public documentation inspected does **not** document a hot-reload command or an exact duplicate-name precedence rule for Antigravity 2.0; treat new conversation/restart as the supported practical refresh and classify live reload/collision ordering as unsupported.

### Antigravity CLI is a separate adapter with different global storage

For Antigravity CLI (`agy`), Google’s migration guide specifies `.agents/skills/` for a workspace skill and `~/.gemini/antigravity-cli/skills/` for a global skill. This is not the same global path as Antigravity 2.0 (`~/.gemini/config/skills/`). [Migrating from Gemini CLI](https://antigravity.google/docs/gcli-migration) explicitly maps both Gemini CLI paths to the Antigravity CLI paths.

**Evidence:** The migration table says Gemini’s global `~/.gemini/skills/` becomes `~/.gemini/antigravity-cli/skills/`, and Gemini’s workspace `.gemini/skills/` becomes `.agents/skills/`; it warns that the workspace folder must be relocated for Antigravity recognition. The [CLI reference](https://antigravity.google/docs/cli-reference) identifies `/skills` as the UI for browsing loaded local and global skills.

**Implication:** The Antigravity CLI adapter is a two-target copier: `<project>/.agents/skills/<skill-folder>/` or `~/.gemini/antigravity-cli/skills/<skill-folder>/`. Verify by launching `agy` from the target workspace, opening `/skills`, and triggering a matching task. The inspected first-party CLI material does not provide a `/skills reload`, a plain-skill uninstall command, or a published direct-skill conflict-precedence rule; close/reopen the CLI after changes and record any duplicate-name result as an environment-specific observation rather than a guaranteed behavior.

### Antigravity plugins delimit, rather than replace, direct installation

Antigravity’s optional plugin system groups skills with rules, MCP configuration, and hooks. A plugin needs `plugin.json`; Antigravity 2.0 scans workspace `.agents/plugins/` (also documented `_agents/plugins/`) and global `~/.gemini/config/plugins/`. [Plugins](https://www.antigravity.google/docs/plugins) is explicit that direct plugin placement is available, but it is a multi-component packaging system.

**Evidence:** The plugin structure includes `plugin.json`, optional `mcp_config.json`, `hooks.json`, `skills/`, and `rules/`; its skills still use `skills/<skill-name>/SKILL.md`.

**Implication:** Reject “create an Antigravity plugin” as the default direct-install interpretation for Skill Issue. Select it only when Skill Issue deliberately ships rules, MCP servers, or hooks with the skill. A skill-only copy is lower-risk and avoids config merges and executable extension lifecycle.

### Gemini CLI direct-install contract and collision behavior

Gemini CLI discovers user skills at `~/.gemini/skills/` (or alias `~/.agents/skills/`) and workspace skills at `<workspace>/.gemini/skills/` (or alias `<workspace>/.agents/skills/`). Discovery order is built-in, extension, user, workspace; a higher tier wins a same-name collision, and within a user/workspace tier `.agents/skills` wins over `.gemini/skills`. [Agent Skills](https://geminicli.com/docs/cli/skills/) supplies the exact precedence and aliases.

**Evidence:** Gemini’s documentation names all four direct paths, scopes, and order. It requires YAML frontmatter where `name` should match the directory and `description` controls matching; [the creation guide](https://geminicli.com/docs/cli/creating-skills/) defines optional `scripts/`, `references/`, and `assets/` and grants the model access to the activated skill’s entire directory.

**Implication:** The Gemini adapter may install the unchanged Skill Issue directory to one selected path, preferably a unique lowercase hyphenated folder/name. It must detect/report an existing same-name skill before replacement, because workspace scope silently overrides user scope by design. It must never write both alias and canonical paths for the same skill: `.agents/skills` would take precedence and make the other copy ambiguous.

### Gemini CLI discovery, consent, and practical verification

Gemini scans enabled skills at session start, injects their name/description, uses `activate_skill` for a matching task, asks the user to approve activation, then grants access to the skill directory and loads `SKILL.md`. [Gemini Agent Skills](https://geminicli.com/docs/cli/skills/) documents this exact lifecycle. A direct copy can be detected without restart via `/skills reload` (or `/skills refresh`); `/skills list` shows current discovery/status. [Managing Agent Skills](https://geminicli.com/docs/cli/using-agent-skills/) documents both commands.

**Evidence:** Gemini also documents `/skills link <path> [--scope user|workspace]` for development, and terminal `gemini skills list`, `install`, `link`, and `uninstall` commands. Its creation guide’s verification flow is: start a matching request, approve activation, then observe the bundled script/resources used.

**Implication:** For direct install, verify in this order: `gemini --version` or `/about`; ensure `skills.enabled` remains true; run `/skills reload`; run `/skills list`; trigger a request matching the description; review the activation consent and expected resource use. `gemini skills link` is a supported development alternative when a durable copy is not desired; it is not a substitute for a managed direct-copy deployment.

### Scope, trust, permissions, and configuration safety

Gemini’s `skills.enabled` defaults to `true`; workspace settings override user settings. Workspace trust can restrict project configuration, and the skill tutorial specifically says a workspace skill requires a trusted folder while user skills are not affected. [Gemini settings](https://geminicli.com/docs/cli/settings/) and [Trusted Folders](https://geminicli.com/docs/cli/trusted-folders/) are the authority.

**Evidence:** Gemini’s trust dialog enumerates skills before approval, warns on dangerous configuration, and untrusted workspaces run in safe mode. Gemini separately requires consent when installing from remote sources and each time a skill activates. For Antigravity 2.0, project settings contain isolated security policies and project-level permissions; Antigravity CLI evaluates `deny`, `ask`, and `allow` policies with `deny > ask > allow`, and unconfigured sensitive actions default to `ask`. [Antigravity Settings](https://www.antigravity.google/docs/settings) and [CLI Permissions](https://www.antigravity.google/docs/cli-permissions) support those claims.

**Implication:** A direct skill copy should not edit JSON settings at all. If an adapter must change a setting outside the skill itself, it must read and structurally merge the existing file at the documented scope, preserve unrelated keys, back it up, and leave approval/sandbox policies no broader than before. Review every bundled script and resource before installation; the permission systems gate agent actions but do not make an unreviewed skill safe.

### Update, repair, rollback, and removal for direct copies

For a direct folder copy, the supported recovery unit is the complete skill directory: stage and inspect a replacement, then replace only the selected `<skill-folder>`, reload/restart, and run the respective discovery/activation check. Gemini has official management commands for installed or linked skills (`gemini skills uninstall <name> [--scope …]`) and remote installs; these commands should be preferred only when that CLI created/tracks the install. [Managing Agent Skills](https://geminicli.com/docs/cli/using-agent-skills/) documents the lifecycle.

**Evidence:** Gemini documents that it scans the skill roots and supports `/skills reload`; Antigravity documents scanning the respective direct roots, but its inspected desktop/CLI documentation provides no direct-folder transaction/rollback command. For optional Gemini extensions, update/uninstall is a distinct package lifecycle that takes effect after restart; [the extension reference](https://github.com/google-gemini/gemini-cli/blob/main/docs/extensions/reference.md) confirms that distinction.

**Implication:** Direct-copy repair is local and reversible: retain the prior complete directory (or restore it from the project revision), replace/remove the one named directory, then refresh. A partial repair that copies only `SKILL.md` is lower fit because it discards supported scripts/references/assets. There is no sourced claim that Antigravity automatically rolls back malformed/direct skill updates; report malformed frontmatter, missing `SKILL.md`, visibility failure, or activation failure separately and restore the prior directory.

### Feature gates and adapter fit

Gemini CLI’s Agent Skills began as a preview feature in v0.23.0 and gained `/skills reload` plus install/uninstall management in v0.24.0; skills were enabled by default by v0.25.0. Current first-party documentation lists v0.50.0 (2026-07-08) and recommends the stable channel. [Release notes](https://geminicli.com/docs/changelogs/) and [settings](https://geminicli.com/docs/cli/settings/) provide the feature history and current default.

**Evidence:** Antigravity’s current documentation labels the desktop product “Antigravity 2.0” and the terminal product “Antigravity CLI,” but the inspected pages do not publish a minimum `agy` or desktop build number for direct skills.

**Implication:** Gemini’s direct adapter should gate on a runtime that exposes `gemini skills` and `/skills reload`, not only a version string. Antigravity adapters should gate on the documented current directory plus a live `/skills`/Customization discovery check. A common Skill Issue payload fits all three Google surfaces; path selection, reload behavior, trust workflow, and collision reporting are bespoke per surface.

## Notes

- **Unsupported:** Antigravity 2.0 hot reload, plain-skill collision precedence, and a transaction-aware direct-skill uninstall/rollback command were not found in the current first-party documentation. Use a new conversation/relaunch and report actual discovery results.
- **Unsupported:** The first-party Antigravity CLI material inspected names direct skill paths and `/skills`, but does not document a `skills link/install/uninstall` lifecycle equivalent to Gemini CLI. Do not transfer Gemini commands to `agy` without a live help check.
- **Rejected lower-fit interpretation:** Treating Antigravity 2.0, Antigravity CLI, and Gemini CLI as sharing one global directory is contradicted by Google’s documented paths.
- **Useful search terms:** `Agent Skills`, `SKILL.md`, `.agents/skills`, `~/.gemini/config/skills`, `~/.gemini/antigravity-cli/skills`, `/skills reload`, `trusted folders`.

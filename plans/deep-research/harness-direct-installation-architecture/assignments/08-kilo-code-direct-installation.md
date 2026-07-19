# Kilo Code Direct Installation

## Assignment

**Goal:** Establish a safe direct, CLI-managed lifecycle for installing a Skill Issue Agent Skill into Kilo Code without creating a Kilo runtime plugin or publishing a Marketplace item.

**Scope:** Current Kilo Code CLI and its current VS Code extension, Kilo Marketplace's file-install model, compatibility directories, and Kilo Cloud Agent's documented project-skill behavior. The research covers locations, layout, discovery, verification, collision and configuration handling, lifecycle operations, and safety controls.

**Exclusions:** Designing a Kilo TypeScript/JavaScript runtime plugin, publishing a Marketplace entry, installing the Kilo IDE extension itself, and making an unverified claim that one project path is automatically portable to Kilo Cloud.

## Sources

- [Kilo Code Skills](https://kilo.ai/docs/customize/skills) — current first-party skill contract: discovery, locations, metadata, priority, reload, troubleshooting, remote URLs, and manual download.
- [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace) — current first-party installation model, exact project/global destinations, removal behavior, and Marketplace boundary.
- [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli) — current CLI/extension relationship, config locations and precedence, installation verification, upgrade commands, and version notice.
- [Kilo Cloud Agent](https://kilo.ai/docs/code-with-ai/platforms/cloud-agent) — first-party Cloud Agent project-skill and persistence behavior.
- [Agent Permissions](https://kilo.ai/docs/customize/agent-permissions) — current VS Code/CLI permission behavior and sensitive-file protections.
- [Kilo Plugins](https://kilo.ai/docs/automate/extending/plugins) — first-party distinction between runtime plugins and file-based skills.
- [Kilo Code v7.4.11 release](https://github.com/Kilo-Org/kilocode/releases/tag/v7.4.11) — latest first-party release listed on 2026-07-19, published 2026-07-16.
- [Kilo source repository](https://github.com/Kilo-Org/kilocode) and [Kilo Marketplace repository](https://github.com/Kilo-Org/kilo-marketplace) — first-party implementation and curated-skill sources.

## Findings

### Direct folder installation is a supported skill lifecycle

Kilo implements the Agent Skills format: a skill is a folder whose required entry point is `SKILL.md`. The Marketplace itself installs skills by creating a directory, and the Skills documentation expressly permits manual download into `.kilo/skills/`. Marketplace items are configuration and instruction files, rather than VS Code extensions, so a Skill Issue installer may materialize the same documented directory layout without a Marketplace publication or native Kilo plugin. [Kilo Skills](https://kilo.ai/docs/customize/skills#finding-skills) [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#files-changed-by-installation)

**Evidence:** Marketplace defines a skill as task-specific instructions/resources which Kilo discovers and loads, writes project skills to `.kilo/skills/<name>/` and global skills to `~/.kilo/skills/<name>/`, and says installation creates the skill's own directory. The skills guide identifies `SKILL.md` as the sole required file and gives `mkdir -p ~/.kilo/skills/api-design` as its own creation example. [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#files-changed-by-installation) [Kilo Skills](https://kilo.ai/docs/customize/skills#optional-bundled-resources)

**Implication:** Use a direct-copy adapter as Kilo's primary Skill Issue adapter. It needs a Kilo-specific target resolver and lifecycle record, but no Kilo plugin manifest, npm package, VSIX, or Marketplace API.

### Scope and surface routing must distinguish current local, compatibility, and cloud paths

For the current VS Code extension and CLI, the first-party skill guide documents the following native destinations: project `.kilo/skills/<skill-name>/`; macOS/Linux user scope `~/.kilo/skills/<skill-name>/`; and Windows user scope `C:\\Users\\<user>\\.kilo\\skills\\<skill-name>\\`. Kilo also scans project `.agents/skills/` by default, and scans `.claude/skills/` only when Claude Code Compatibility is enabled. A project or global `kilo.jsonc` may add absolute, `~/`-relative, or project-relative `skills.paths`, or remote `skills.urls` directories serving `index.json`. [Kilo Skills](https://kilo.ai/docs/customize/skills#skill-locations) [Kilo Skills](https://kilo.ai/docs/customize/skills#additional-skill-paths-and-remote-urls)

**Evidence:** Kilo says project skills override global skills with the same name; compatibility and configured paths are additionally loaded. The current extension is rebuilt on the Kilo CLI, and the skill guide says the extension obtains skills when it connects to the CLI server. [Kilo Skills](https://kilo.ai/docs/customize/skills#priority-and-overrides) [Kilo Skills](https://kilo.ai/docs/customize/skills#when-skills-are-loaded) [Kilo Installation](https://kilo.ai/docs/getting-started/installing)

**Implication:** Provide explicit `project` and `user` targets only. Treat `.agents/skills` as a lower-fit portability target because it is shared with other agents, and do not select `.claude/skills` unless that compatibility feature is positively enabled. Prefer native directories over `skills.paths` or `skills.urls`; the latter require configuration mutation or a remote-manifest service.

### Kilo Cloud has a documented legacy-path caveat

Kilo Cloud Agent says a cloned repository automatically makes skills under `.kilocode/skills/` available and that global `~/.kilocode/skills/` are unavailable because the user home is not persistent. This differs from the current local skill guide, which names `.kilo/skills/` and `~/.kilo/skills/` as native locations. [Kilo Cloud Agent](https://kilo.ai/docs/code-with-ai/platforms/cloud-agent#skills) [Kilo Skills](https://kilo.ai/docs/customize/skills#skill-locations)

**Evidence:** The CLI configuration guide also calls `.kilocode/` a legacy project configuration directory, while the current skills guide calls `.kilo/skills/` the project location. [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli#config-file-location-kilo-cli-10) [Kilo Skills](https://kilo.ai/docs/customize/skills#skill-locations)

**Implication:** Classify Cloud support for a `.kilo/skills/` direct installation as **caveated**, rather than silently duplicating the skill under both names. A Cloud-capable adapter needs a live Cloud session probe against the target Kilo version or an official documentation correction before it can claim one project package works across local and Cloud surfaces. Global installation must never be presented as Cloud-persistent.

### The portable content format is compact and has strict identity requirements

`SKILL.md` needs YAML frontmatter and Markdown instructions. `name` and `description` are required; the documented limits are 64 characters and lowercase letters/numbers/hyphens for `name`, and 1024 characters for `description`. Kilo says the `name` must match its parent directory name. `license`, `compatibility`, and `metadata` are optional; `scripts/`, `references/`, and `assets/` are optional bundled resources. [Kilo Skills](https://kilo.ai/docs/customize/skills#skillmd-format) [Kilo Skills](https://kilo.ai/docs/customize/skills#name-matching-rule) [Kilo Skills](https://kilo.ai/docs/customize/skills#optional-bundled-resources)

**Evidence:** Kilo reads only name, description, and path at discovery; it reads full instructions only after the active model decides a task clearly applies. The current platform has one shared skill pool rather than mode-specific directories. [Kilo Skills](https://kilo.ai/docs/customize/skills#how-skills-work-in-kilo-code) [Kilo Skills](https://kilo.ai/docs/customize/skills#mode-specific-skills)

**Implication:** Validate the canonical Skill Issue bundle before writing: directory basename equals frontmatter `name`; frontmatter parses; description explicitly names the requests that should use it; all referenced bundled files are inside the copied folder. Do not make Kilo mode-specific copies. A single canonical Agent Skill folder is a strong common-adapter fit.

### Discovery, activation, and verification are session-bound but testable

Kilo scans configured directories at the start of a session; the CLI does this on a new session or `kilo run`, and the extension does it when it connects to the CLI server. `/reload` re-scans an active session. Kilo recommends confirming availability by asking whether the named skill is loaded, and it records actual use as a `skill` tool call whose result contains the injected content. Explicitly naming the skill is documented to trigger it; implicit matching relies on the model's reading of `description`. [Kilo Skills](https://kilo.ai/docs/customize/skills#when-skills-are-loaded) [Kilo Skills](https://kilo.ai/docs/customize/skills#verifying-a-skill-is-available) [Kilo Skills](https://kilo.ai/docs/customize/skills#how-the-agent-decides-to-use-a-skill)

**Evidence:** The documented troubleshooting sequence is frontmatter, reload/new session, direct `SKILL.md` placement, then configured paths/URLs. Kilo CLI itself can be preflighted with `kilo --version`; the current documentation applies to Kilo 1.0 and later. [Kilo Skills](https://kilo.ai/docs/customize/skills#troubleshooting) [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli#verify-installation) [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli)

**Implication:** The adapter's practical verification should: (1) confirm the final `SKILL.md` exists at the selected scope and its frontmatter identity matches; (2) tell the operator to start a new session or issue `/reload`; (3) submit an explicit `use <skill-name>` availability check; and (4) where transcript access exists, require the `skill` tool call before declaring behavioral activation. A mere copied folder proves installation, not invocation.

### Collision handling and configuration edits must be narrowly owned

Project scope takes precedence over global scope for the same skill name. Current documentation gives no contract for merging two independently managed skill directories with the same name, so installation must treat a non-owned target directory as a collision, rather than overwriting or merging it. [Kilo Skills](https://kilo.ai/docs/customize/skills#priority-and-overrides) [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#project-or-global)

**Evidence:** Marketplace scopes are independently removable: deleting a project copy does not remove a global copy, and vice versa. Its documented skill installation mutates only the skill directory; MCP installation is the separate case that adds a key to config while preserving other settings. [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#files-changed-by-installation) [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#removing-an-item)

**Implication:** Standard installation should not edit `kilo.jsonc` at all. Resolve one target; if absent, stage the complete skill folder beside it, validate it, then rename into place; if an installer-owned record and matching name exist, replace only after creating a rollback copy; otherwise stop with an explicit collision. If a future `skills.paths` adapter is requested, parse and preserve JSONC/comments, update only the `skills.paths` entry, de-duplicate exact paths, and retain a pre-edit byte-for-byte backup. That merge algorithm is an adapter safeguard, **not** a Kilo-documented config-edit contract.

### Update, repair, rollback, and uninstall can be directory-scoped

The direct skill contract is folder-based, and Marketplace removal is scope-local directory removal. Therefore a CLI-managed lifecycle can keep an installation record with the selected scope, destination, bundle version/digest, and backup path; it can then repair by validating the final folder against the recorded digest, update by replace-with-backup, roll back by restoring that backup, and uninstall by deleting only the recorded destination. [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#files-changed-by-installation) [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#removing-an-item)

**Evidence:** Kilo reloads affected configuration after Marketplace install/removal and warns that running sessions may be interrupted to avoid stale agents, skills, or tools. Separately, the Skills guide provides `/reload` for active-session rescan. [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#removing-an-item) [Kilo Skills](https://kilo.ai/docs/customize/skills#when-skills-are-loaded)

**Implication:** Keep a per-install backup until post-reload verification succeeds. On partial recovery (staging exists but final directory does not), discard only the installer staging directory and retry; on a corrupt final owned directory, restore its last verified backup. Never repair by deleting an unowned directory or by altering a same-name skill in another scope. After each lifecycle change, require `/reload` or a fresh session before status becomes active.

### Skills are instructions that can drive tools, so the adapter must preserve trust boundaries

Kilo permissions regulate tool calls as `allow`, `ask`, or `deny`; last matching permission wins. Skills can include scripts and reference materials and the agent may execute bundled code when following a skill, so copying a skill has a materially different risk profile from copying static prose. Kilo also keeps `.env` and `.env.*` reads behind a built-in prompt even under broad reads. [Agent Permissions](https://kilo.ai/docs/customize/agent-permissions#actions) [Agent Permissions](https://kilo.ai/docs/customize/agent-permissions#rule-precedence) [Kilo Skills](https://kilo.ai/docs/customize/skills#optional-bundled-resources) [Agent Permissions](https://kilo.ai/docs/customize/agent-permissions#sensitive-files)

**Evidence:** Kilo advises reviewing Marketplace item authors, sources, prerequisites, parameters, and tools before installation. The relevant permissions page applies to both the current VS Code extension and CLI. [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace#mcp-security-and-permissions) [Agent Permissions](https://kilo.ai/docs/customize/agent-permissions)

**Implication:** Require an explicit scope selection and show a file manifest plus content digest before installing; avoid silently widening Kilo permissions; preserve Kilo's existing `ask` defaults; and record whether the bundle includes executable files. Project installation may be committed and shared, so treat it as reviewed repository content. Remote `skills.urls` is lower fit for Skill Issue because Kilo downloads the manifest-listed files at discovery; use it only from a controlled HTTPS origin with an independently verified manifest, not as a replacement for reproducible local materialization. [Kilo Skills](https://kilo.ai/docs/customize/skills#additional-skill-paths-and-remote-urls)

### Runtime plugins and IDE packaging are distinct, lower-fit mechanisms

Kilo runtime plugins are TypeScript/JavaScript modules loaded at startup; they can add tools and intercept events, and `kilo plugin` writes plugin entries into configuration. The current VS Code extension is built on the CLI, so extension installation is a prerequisite for the IDE surface, not a requirement to package a skill as an extension. [Kilo Plugins](https://kilo.ai/docs/automate/extending/plugins) [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli) [Kilo Installation](https://kilo.ai/docs/getting-started/installing)

**Evidence:** Plugin packages have dependency cache and engine-compatibility behavior, whereas skills use a directly discoverable `SKILL.md` directory. Kilo's own Marketplace explicitly describes items as configuration/instruction files rather than VS Code extensions. [Kilo Plugins](https://kilo.ai/docs/automate/extending/plugins#how-plugins-are-installed) [Kilo Marketplace](https://kilo.ai/docs/customize/marketplace)

**Implication:** Reject "build a Kilo plugin" and "ship a VSIX" as the default direct-installation interpretation: both introduce a different runtime and release lifecycle without helping Kilo discover a Skill Issue Agent Skill. Keep a plugin adapter as **lower fit** only if Skill Issue later needs hooks, tools, auth providers, or event interception that an Agent Skill cannot provide.

### Version gating is required, while the direct skill feature has no narrower documented minimum

The current CLI guide says its configuration documentation applies to Kilo 1.0 and later. The latest official Kilo Code release visible at research time is v7.4.11, released 2026-07-16. The current skills, permissions, and Marketplace documentation describes the current VS Code extension and CLI; none states a separate minimum version for native `.kilo/skills` direct-folder discovery. [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli) [Kilo Code v7.4.11 release](https://github.com/Kilo-Org/kilocode/releases/tag/v7.4.11) [Agent Permissions](https://kilo.ai/docs/customize/agent-permissions)

**Evidence:** Kilo supplies `kilo --version` for version verification and `kilo upgrade` or `npm update -g @kilocode/cli` for CLI updates. The release train is active and uses Kilo Core/extension releases, so paths and migration behavior are version-sensitive. [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli#verify-installation) [Kilo CLI](https://kilo.ai/docs/code-with-ai/platforms/cli#update) [Kilo Code releases](https://github.com/Kilo-Org/kilocode/releases)

**Implication:** Gate the adapter on a positively detected Kilo CLI/extension version that is at least 1.0, record the detected version, and issue a conservative unsupported result when the version cannot be obtained. Do not claim a version floor higher than 1.0 until a first-party skill-feature changelog establishes one. On legacy installations, prefer a migration check over assuming `.kilocode` and `.kilo` are interchangeable.

## Notes

- **Recommended direct lifecycle:** validate canonical bundle -> choose `project` or `user` -> preflight `kilo --version` -> reject unowned collision -> stage, validate, and move the one folder -> save installer-owned record/backup -> `/reload` or new session -> explicit named-skill availability and tool-call verification -> retain backup until confirmed. This is a common bundle adapter with Kilo path and session hooks, not a bespoke content renderer.
- **Rejected interpretation:** Marketplace publication is required for discoverability. Current Kilo documentation explicitly supports manual skill-folder download and says Marketplace installation itself materializes normal configuration files.
- **Rejected interpretation:** a VS Code extension requires a VSIX or runtime plugin for each skill. The extension delegates to the CLI-backed skill system; a skill is a directory, while plugins are a separate runtime feature.
- **Unsupported:** Kilo's current documentation conflicts on Cloud project's native directory (`.kilocode/skills/`) versus the current local guide (`.kilo/skills/`). No source inspected proves automatic Cloud discovery of the latter. Do not duplicate folders automatically because same-name collision behavior across those directories is not documented.
- **Unsupported:** a Kilo CLI subcommand dedicated to installing an arbitrary local Agent Skill. The documented `kilo plugin` command manages plugins, while skills are documented as folders, manually downloaded content, configured paths, or remote manifests.
- **Useful search terms:** `Kilo Code Agent Skills`, `skills.paths`, `skills.urls`, `kilo.jsonc`, `Kilo Marketplace`, `Kilo Cloud Agent skills`, `Kilo Agent Permissions`.

# OpenCode Direct Installation

## Assignment

**Goal:** Establish the current, direct CLI-managed lifecycle for installing a Skill Issue skill in OpenCode without a plugin package.

**Scope:** Official OpenCode documentation, current first-party source, and first-party release history only. Covers native skills, compatible skill directories, configuration/permission boundaries, and the CLI, TUI, server/web, desktop, IDE, and GitHub Action surfaces.

**Exclusions:** Third-party skill plugins, package registries, and unverified community conventions. Plugins are considered only to distinguish them from native skills.

## Sources

- [OpenCode Agent Skills documentation](https://opencode.ai/docs/skills) — authoritative user-facing contract; last updated July 19, 2026.
- [OpenCode Config documentation](https://opencode.ai/docs/config) — authoritative locations, precedence, managed-configuration, server, and interface behavior.
- [OpenCode CLI documentation](https://opencode.ai/docs/cli) — authoritative CLI/server flags and documented environment variables.
- [OpenCode Plugins documentation](https://opencode.ai/docs/plugins) — authoritative plugin/package boundary.
- [OpenCode IDE documentation](https://opencode.ai/docs/ide) — authoritative IDE integration behavior.
- [First-party current skill discovery source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/skill/index.ts), [skill tool source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/tool/skill.ts), and [skills config schema](https://github.com/anomalyco/opencode/blob/dev/packages/core/src/v1/config/skills.ts) — source-level cross-checks; the `dev` branch can be ahead of releases.
- [Native Agent Skills landing commit](https://github.com/anomalyco/opencode/commit/8fe071592825) and [release `v1.0.186`](https://github.com/anomalyco/opencode/releases/tag/v1.0.186) — first-party version-history cross-check; the commit is absent from `v1.0.185` and present in `v1.0.186` according to the repository comparison API.
- [Latest release `v1.18.3`](https://github.com/anomalyco/opencode/releases/tag/v1.18.3) — current released-version checkpoint inspected July 19, 2026.

## Findings

### Native filesystem skills are the supported direct-install route

OpenCode has a built-in Agent Skills system: it discovers a `SKILL.md` from a skill folder and exposes only its name and description until the native `skill` tool loads the full content. Creating the folder and file is the official installation procedure, so Skill Issue does not need an OpenCode plugin, npm package, or custom installer for the ordinary case. [Agent Skills](https://opencode.ai/docs/skills#place-files) [Agent Skills discovery](https://opencode.ai/docs/skills#understand-discovery)

**Evidence:** The official docs describe the feature as native and provide direct file creation under `.opencode/skills/<name>/SKILL.md`; current first-party source implements discovery and the `skill` tool rather than delegating this to a plugin. [skill discovery source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/skill/index.ts) [skill tool source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/tool/skill.ts)

**Implication:** The direct adapter is a deterministic directory copy: create exactly one skill directory and copy its Skill Issue files into it. Do not add a `plugin` entry or trigger Bun/npm installation.

### Exact supported locations and scope selection

The official direct-discovery locations are:

| Intent              | Native OpenCode location                    | Compatible location also read by OpenCode                                |
| ------------------- | ------------------------------------------- | ------------------------------------------------------------------------ |
| Project / committed | `.opencode/skills/<name>/SKILL.md`          | `.agents/skills/<name>/SKILL.md` or `.claude/skills/<name>/SKILL.md`     |
| User / uncommitted  | `~/.config/opencode/skills/<name>/SKILL.md` | `~/.agents/skills/<name>/SKILL.md` or `~/.claude/skills/<name>/SKILL.md` |

For project locations, OpenCode walks upward from the current working directory to the Git worktree and loads matching skill folders along the way. For global locations, it loads all three user-level roots. [Agent Skills locations and discovery](https://opencode.ai/docs/skills#place-files) [Agent Skills discovery](https://opencode.ai/docs/skills#understand-discovery)

**Evidence:** These are the six locations explicitly documented by OpenCode. The current source scans native config directories plus `.agents` and, unless disabled, `.claude`; it also follows symlinks while scanning. [Agent Skills](https://opencode.ai/docs/skills#place-files) [discovery source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/skill/index.ts)

**Implication:** Prefer the native `.opencode` roots for an OpenCode-specific install. Use `.agents/skills` only when one shared, cross-harness copy is an explicit product requirement; use `.claude/skills` only when deliberately sharing with Claude Code. There is no documented system-wide **skill directory**: managed system paths govern `opencode.json` / `opencode.jsonc` and policy, rather than documenting an administrator-owned skills root. [managed settings](https://opencode.ai/docs/config#managed-settings)

### Required layout and portable skill content

The folder name and frontmatter `name` must be identical. `name` is required, 1–64 characters, lowercase alphanumeric segments separated by one hyphen, and must match `^[a-z0-9]+(-[a-z0-9]+)*$`. `description` is required and must contain 1–1024 characters. Optional recognized fields are `license`, `compatibility`, and a string-to-string `metadata` map; unknown fields are ignored. [Agent Skills frontmatter and names](https://opencode.ai/docs/skills#write-frontmatter) [description rule](https://opencode.ai/docs/skills#follow-length-rules)

```text
.opencode/
  skills/
    <skill-name>/
      SKILL.md
      references/ ...     # optional supporting assets
      scripts/ ...        # optional supporting assets
```

**Evidence:** The official example uses `<name>/SKILL.md`. The current native tool returns the skill directory as the base for relative paths and includes a sampled sibling-file list, so supporting assets in that directory can be referenced by the loaded instructions; no separate Skill Issue/OpenCode manifest is defined. [official example](https://opencode.ai/docs/skills#use-an-example) [skill tool source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/tool/skill.ts)

**Implication:** Preserve Skill Issue’s canonical `SKILL.md` plus relative support files. The adapter should only ensure the target directory name matches the canonical, lowercase skill ID; avoid translating metadata into an OpenCode-only manifest.

### Discovery, activation, restart, and reliable verification

OpenCode supplies an available-skills list to agents and loads a chosen skill through `skill({ name: "<name>" })`. Skill permissions govern that call. The documented failure checks are: uppercase `SKILL.md`, required frontmatter, globally unique names, and no `deny` permission. [skill activation](https://opencode.ai/docs/skills#recognize-tool-description) [skill troubleshooting](https://opencode.ai/docs/skills#troubleshoot-loading)

**Evidence:** Source initializes a per-instance discovered-skill state and the source-defined `opencode debug skill` command serializes all current skills. The public CLI page documents `opencode debug` but not its `skill` subcommand, so that diagnostic is source-current rather than a stable documented contract. [skill source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/skill/index.ts) [debug skill command](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/cli/cmd/debug/skill.ts) [CLI debug section](https://opencode.ai/docs/cli#debug)

**Implication:** After copy, start a fresh `opencode` TUI or run `opencode run --dir <project> "Use the <skill-name> skill"`; verify the skill is listed/loaded and inspect `opencode debug skill` only as a caveated current-build diagnostic. For `opencode serve` or `opencode web`, restart the backend after copying or updating skills, then attach/run against that restarted instance. This restart procedure is deterministic because OpenCode documents server attachment and source caches discovery per backend instance; no official skill hot-reload contract was found. [CLI run and attach](https://opencode.ai/docs/cli#run)

### Permissions, trust, and policy boundaries

`permission.skill` accepts exact names or wildcard patterns with `allow`, `ask`, and `deny`: `allow` loads immediately, `ask` prompts the user, and `deny` hides the skill and rejects access. Permissions can be set globally in `opencode.json` or overridden per agent; the legacy `tools.skill: false` setting disables the tool completely. [skill permissions](https://opencode.ai/docs/skills#configure-permissions) [per-agent override and disablement](https://opencode.ai/docs/skills#override-per-agent)

**Evidence:** Managed configuration loads at the highest priority and cannot be overridden by project/user settings. Configuration files are merged rather than replaced, and later sources override conflicting keys. [config merge and precedence](https://opencode.ai/docs/config#locations) [managed settings](https://opencode.ai/docs/config#managed-settings)

**Implication:** Treat each `SKILL.md` and its referenced assets as trusted instruction-bearing content. If approval is required, minimally merge a `permission.skill` entry such as `"skill-issue-*": "ask"` into the existing configuration rather than replacing `opencode.json`; inspect effective managed/organization policy first because it can still deny the skill. Do not rely on `.claude/skills` compatibility where `OPENCODE_DISABLE_CLAUDE_CODE` or `OPENCODE_DISABLE_CLAUDE_CODE_SKILLS` is set. [CLI environment variables](https://opencode.ai/docs/cli#environment-variables)

### Collision handling, update, repair, rollback, and uninstall

Skill names must be unique across all discovery locations. The official troubleshooting guidance treats duplicate names as an error condition; current source logs the duplicate and overwrites the previous in-memory entry, but does not document a safe precedence promise. [troubleshooting](https://opencode.ai/docs/skills#troubleshoot-loading) [source collision handling](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/skill/index.ts)

**Evidence:** Native skills are plain discovered folders; OpenCode documents no dedicated skills install, update, repair, rollback, or uninstall CLI command. The CLI’s package/plugin lifecycle is a materially different surface: plugins are JavaScript/TypeScript event hooks, may be named in `opencode.json`, and npm plugins are installed with Bun at startup. [Plugins](https://opencode.ai/docs/plugins) [plugin installation/load order](https://opencode.ai/docs/plugins#how-plugins-are-installed)

**Implication:** Use an atomic directory replacement strategy owned by the Skill Issue installer: stage `<skill-name>.new`, validate the filename/frontmatter/name constraints, move the existing folder to a timestamped backup, rename the staged folder into place, restart the applicable OpenCode process, and verify activation. To repair, restore the prior backup or recopy the known canonical bundle; to uninstall, remove only that named folder and restart. Never overwrite a same-named folder from another owner—report a collision and require an explicit replace/rename decision. Direct filesystem installs need no config mutation, so their rollback does not need configuration recovery.

### Surface separation and adapter fit

The TUI starts when `opencode` is run without arguments; `opencode run` starts a non-interactive session; `serve` and `web` host a backend that `run --attach` or the TUI can connect to. The IDE extension launches OpenCode from the IDE terminal, while current config behavior is stated to apply across TUI, CLI, desktop, and GitHub Action. [CLI overview/run/server](https://opencode.ai/docs/cli) [IDE](https://opencode.ai/docs/ide) [cross-interface config statement](https://opencode.ai/docs/config#default-agent)

**Evidence:** Native skills are part of the shared runtime, whereas the IDE integration is a terminal/extension launch mechanism and plugins are hook modules with separate startup/package semantics. [Agent Skills](https://opencode.ai/docs/skills) [IDE integration](https://opencode.ai/docs/ide) [Plugins](https://opencode.ai/docs/plugins)

**Implication:** A single **common bundle** (`<skill-name>/SKILL.md` plus relative assets) fits OpenCode unchanged. The only **bespoke adapter** required is target placement: native `.opencode/skills` for OpenCode-only scope, or the explicitly selected compatible root for cross-harness sharing. The desktop/IDE/server/cloud-like entrypoints do not justify different packages; select the correct filesystem/project context and restart the persistent server when applicable.

### Version gates and lower-fit interpretations

Native Agent Skills first landed in first-party commit `8fe071592825` and are confirmed included by release `v1.0.186`; the current latest release inspected is `v1.18.3`. [landing commit](https://github.com/anomalyco/opencode/commit/8fe071592825) [v1.0.186](https://github.com/anomalyco/opencode/releases/tag/v1.0.186) [v1.18.3](https://github.com/anomalyco/opencode/releases/tag/v1.18.3)

**Evidence:** The current docs describe native skills directly. Current `dev` source additionally exposes `skills.paths` and `skills.urls`; `paths` can scan additional folders, while `urls` expects a remote `index.json` catalog and caches fetched content. These config fields are present in first-party source/schema but are not documented on the public skills page, so they are not a default direct-install mechanism. [skills config schema](https://github.com/anomalyco/opencode/blob/dev/packages/core/src/v1/config/skills.ts) [discovery source](https://github.com/anomalyco/opencode/blob/dev/packages/opencode/src/skill/index.ts)

**Implication:** Gate direct native installation on `opencode --version` being at least `v1.0.186`; upgrade with the documented `opencode upgrade` command if necessary. [CLI version/upgrade](https://opencode.ai/docs/cli#global-flags) [CLI upgrade](https://opencode.ai/docs/cli#upgrade) Reject the following lower-fit interpretations unless separately requested: npm/plugin distribution (adds code hooks and Bun lifecycle), a remote skill registry (`skills.urls`, not direct filesystem installation), and undocumented shared-path configuration (`skills.paths`, source-current only).

## Notes

- **Caveated:** No public skill hot-reload, install/update, or rollback command was found. The recommended restart-and-verify cycle is based on documented startup/server behavior plus current per-instance discovery source.
- **Caveated:** `opencode debug skill`, `skills.paths`, and `skills.urls` are verified in current first-party `dev` source but are not fully described by current public docs; avoid making the Skill Issue standard lifecycle depend on them.
- **Unsupported:** No native OpenCode system-wide skills root, or desktop-only / IDE-only skill format, was found in official documentation.
- Useful search terms: `Agent Skills`, `SKILL.md`, `permission.skill`, `OPENCODE_DISABLE_CLAUDE_CODE_SKILLS`, `opencode debug skill`.

# Pi Direct Installation

## Assignment

**Goal:** Establish the supported direct CLI-managed lifecycle for installing a Skill Issue skill into the Pi coding agent, while keeping native package-manager delivery out of the product architecture.

**Scope:** Current official Pi documentation and first-party `earendil-works/pi` source for skills, resource loading, settings, package management, trust, reload, and version gates; all direct user, project, settings-path, and one-run CLI routes.

**Exclusions:** Implementing Skill Issue, assessing its current artifact contents, third-party Pi extensions/packages, non-Pi agent products, and native npm/git package distribution as the primary delivery architecture.

## Sources

- [Pi Skills documentation](https://pi.dev/docs/latest/skills) — current authoritative user contract for Agent Skills, locations, frontmatter, validation, collisions, and invocation; retrieved 2026-07-19.
- [Pi Settings documentation](https://pi.dev/docs/latest/settings) — global/project settings, trust, resource path settings, and project override contract; retrieved 2026-07-19.
- [Pi Packages documentation](https://pi.dev/docs/latest/packages) — package-only install/remove/update lifecycle, local path package behavior, security, and resource bundles; retrieved 2026-07-19.
- [Pi Security documentation](https://pi.dev/docs/latest/security) and [Using Pi](https://pi.dev/docs/latest/usage) — trust boundary, no-sandbox policy, startup header, `/reload`, and command surface; retrieved 2026-07-19.
- [Pi Extensions documentation](https://pi.dev/docs/latest/extensions) — distinct TypeScript extension surface and reload/session lifecycle; retrieved 2026-07-19.
- First-party source: [`skills.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/skills.ts), [`package-manager.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/package-manager.ts), [`settings-manager.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/settings-manager.ts), and [`resource-loader.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/resource-loader.ts). Inspected via GitHub API at current `main` on 2026-07-19; source confirms precedence, ancestor `.agents` traversal, symlink/name collision handling, and array replacement in project-over-global effective settings.
- [Pi 0.50.0 release note](https://pi.dev/news/releases/0.50.0) — introduced ResourceLoader-only resource loading, resource `/reload`, and array-based resource settings. [Pi 0.80.10 release note](https://pi.dev/news/releases/0.80.10) and [npm package record](https://www.npmjs.com/package/%40earendil-works/pi-coding-agent) — current package release available at research time (`0.80.10`, published 2026-07-18).

## Findings

### Correct Pi surface: an Agent Skill, not a prompt, extension, package, SDK, or TUI integration

Pi is the terminal coding harness published as `@earendil-works/pi-coding-agent`. Its customization surfaces are distinct: **skills** are on-demand Agent Skills; **prompt templates** are Markdown text expanded by `/template-name`; **extensions** are TypeScript modules that can add tools, commands, events, and UI; **Pi packages** bundle any of those resources for npm/git distribution; and the **SDK/RPC/TUI** are programmatic embedding and interface surfaces. A Skill Issue deliverable whose purpose is reusable task instructions and optional helper files maps directly to Pi's skill surface. It does not need a TypeScript adapter, package manifest, SDK host, or TUI component.

**Evidence:** Pi defines skills as self-contained, on-demand capability packages with workflows, scripts, and references in the [Skills overview and structure](https://pi.dev/docs/latest/skills). Prompt templates are separately defined as slash-expanded Markdown snippets in [Prompt Templates](https://pi.dev/docs/latest/prompt-templates). Extensions are TypeScript behavior modules in [Extensions](https://pi.dev/docs/latest/extensions), while packages only bundle and share resources through npm or git in [Pi Packages](https://pi.dev/docs/latest/packages). The [SDK](https://pi.dev/docs/latest/sdk) is for embedding Pi in another application.

**Implication:** Use a common Agent Skills artifact as the Pi adapter: a directory containing `SKILL.md` plus any relative scripts/references/assets. The installation adapter should only select a Pi discovery location and copy/remove that directory. Treat a prompt template as a lower-fit optional shortcut, and extensions/packages as bespoke or distribution mechanisms rather than requirements for a skill install.

### Direct filesystem discovery is supported at four native scopes plus two explicit-path routes

Pi officially auto-discovers direct skills from these locations:

| Scope | Direct location | Trust / reach |
| --- | --- | --- |
| User, Pi-specific | `~/.pi/agent/skills/<skill-dir>/SKILL.md` | Loaded for all projects. |
| User, shared Agent Skills | `~/.agents/skills/<skill-dir>/SKILL.md` | Loaded for all projects. |
| Project, Pi-specific | `<cwd>/.pi/skills/<skill-dir>/SKILL.md` | Loaded only after the project is trusted. |
| Project, shared Agent Skills | `<repo-or-ancestor>/.agents/skills/<skill-dir>/SKILL.md` | Pi scans from `cwd` toward the git root, or filesystem root when outside a repository; loaded only after project trust. |
| Settings-path | `skills` array in `~/.pi/agent/settings.json` or `<cwd>/.pi/settings.json` | Direct file or directory paths; project settings are trust-gated. |
| One run | `pi --skill <path>` (repeatable) | Explicit, additive path; works even with `--no-skills`. |

There is no documented machine-wide/system configuration directory analogous to `/etc`; the supported persistent global scopes are the two user-home locations above. A direct root `.md` skill is accepted in the Pi-specific `skills/` directories, but the stable cross-harness layout is `<skill-dir>/SKILL.md`; root `.md` files are ignored in the `.agents/skills/` locations.

**Evidence:** [Skills: Locations and discovery rules](https://pi.dev/docs/latest/skills) names every global/project/settings/CLI route, the trust condition, `.agents` ancestor traversal, and the root-file distinction. Current first-party [`package-manager.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/package-manager.ts) implements that traversal and has separate discovery modes for `.pi/skills` and `.agents/skills`.

**Implication:** The preferred direct installer has two choices, not a package-manager dependency: copy the entire common artifact to `~/.pi/agent/skills/<name>/` for a user install, or to `.pi/skills/<name>/` for a project install. Use `~/.agents/skills/` or project `.agents/skills/` only when one shared copy across compatible harnesses is an explicit product goal. Use settings paths for an external managed source tree, and `--skill` for a non-persistent verification/trial.

### Layout, metadata, naming, and command activation are standard Agent Skills

A Pi skill is a directory with a required `SKILL.md`; additional files are freeform and referenced relative to that directory. `SKILL.md` requires YAML frontmatter `name` and `description`, followed by instructions. `name` must be 1–64 characters of lowercase letters, digits, and hyphens, with no leading/trailing or consecutive hyphen. `description` is required, max 1024 characters, and determines whether Pi offers the skill to the model. Pi permits the name to differ from the parent directory, although matching them remains the safest portable convention. Optional standard fields include `license`, `compatibility`, `metadata`, `allowed-tools` (experimental), and `disable-model-invocation`.

At startup Pi puts skill names/descriptions in the system prompt and loads the full `SKILL.md` on demand through the `read` tool. With `enableSkillCommands: true` (the documented default), Pi also registers `/skill:<name>`; use that command to force loading when model selection would be uncertain.

**Evidence:** [Skills: How Skills Work, Structure, Frontmatter, Name Rules, and Validation](https://pi.dev/docs/latest/skills) defines progressive disclosure, relative files, frontmatter limits, and `/skill:` commands. First-party [`skills.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/skills.ts) confirms directory recursion stops at a `SKILL.md` root and that `.md` explicit paths are permitted.

**Implication:** Preserve the Skill Issue artifact's standard `SKILL.md` layout and helper file paths. Do not translate its instructions into a Pi prompt template or generate package metadata just to make Pi recognize it. The adapter needs only validate the supplied frontmatter, target directory name, and complete directory copy.

### Discovery occurs at startup; `/reload` is the supported activation and verification step

Pi scans skill locations at startup. The interactive startup header reports loaded skills, and `/reload` reloads skills along with keybindings, extensions, prompts, themes, and context files. This provides an explicit no-restart activation route after a direct copy, update, repair, or removal. Pi documents `resources_discover` with reason `startup` or `reload`; it also documents session reload behavior for extensions.

Practical verification sequence for an installer:

1. Confirm `<target>/<skill-dir>/SKILL.md` exists and frontmatter has a valid name plus non-empty description.
2. Start Pi in the intended directory, resolve the project trust prompt if applicable, and inspect the startup header for the skill.
3. Run `/skill:<name>` and confirm Pi exposes/loads the selected `SKILL.md`; then use a benign request covered by the skill.
4. For an already-running interactive Pi, run `/reload` before step 2 rather than assuming a filesystem watcher.

**Evidence:** [Skills: How Skills Work](https://pi.dev/docs/latest/skills) says scanning happens at startup and describes `/skill:<name>`; [Using Pi](https://pi.dev/docs/latest/usage) documents the startup header and `/reload`; [Extensions: resource events](https://pi.dev/docs/latest/extensions) documents `resources_discover` startup/reload reasons. [Pi 0.50.0](https://pi.dev/news/releases/0.50.0) introduced resource hot reload. Current documentation does not document automatic filesystem watching for direct skill edits; that claim is therefore unsupported.

**Implication:** The direct-install UX should report “copied; run `/reload` or restart Pi” and mark success only after discoverability and forced invocation are observed. A live watcher is lower-fit/unvalidated for this adapter.

### Trust and permissions make project installation an explicit security boundary

Project-local `.pi` resources and project `.agents/skills` require trust. On interactive startup, Pi asks when it finds project settings/resources and has no saved decision; trust decisions live in `~/.pi/agent/trust.json`. Non-interactive `-p`, JSON, and RPC modes do not prompt: with the default `defaultProjectTrust: "ask"` they ignore untrusted project resources, unless a saved decision, `defaultProjectTrust: "always"`, or `--approve` applies. `/trust` saves a decision but only affects the next Pi start/reload boundary described by the docs.

Pi runs with the invoking user's full permissions and has no built-in sandbox. Skills can instruct the model to run executable helpers; extensions and packages can execute arbitrary code. Pi's own guidance is to review skills/package source before use and to use operating-system or container isolation for real boundaries.

**Evidence:** [Settings: Project Trust](https://pi.dev/docs/latest/settings) and [Security: Project Trust](https://pi.dev/docs/latest/security) define trust triggers, decisions, non-interactive behavior, `--approve`, and trust storage. [Security: No Built-in Sandbox](https://pi.dev/docs/latest/security) defines the process-permission model, and [Skills: Locations](https://pi.dev/docs/latest/skills) requires content review.

**Implication:** For a project target, the installer must not silently change `defaultProjectTrust`, create `trust.json` approvals, or execute helper scripts. It should surface the trust requirement and let the Pi user approve the repository. For either scope, install only reviewed artifact content and preserve the principle that Pi's runtime permissions are the user's permissions.

### Collision and precedence handling require deterministic target ownership

Pi warns on same-name collisions and retains the first skill loaded. First-party loader source orders resources by precedence: project settings paths, project auto-discovery, user settings paths, user auto-discovery, then package resources. It also canonicalizes files so duplicate symlinked paths are skipped. Therefore, two differently located Skill Issue copies with the same `name` can produce a valid-looking but unintended winner.

**Evidence:** [Skills: Validation](https://pi.dev/docs/latest/skills) states “first skill found” survives a same-name collision. First-party [`package-manager.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/package-manager.ts) documents and sorts the five precedence ranks; [`skills.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/skills.ts) records collision diagnostics and canonical-path deduplication.

**Implication:** Before copy, inspect the chosen scope and other documented skill roots/settings paths for the target `name`. Default to a deterministic “conflict: do not overwrite” result. Offer explicit replacement only when the existing file is recognizably owned by the same installer/receipt; otherwise require the user to choose removal, a unique valid skill name, or the intended higher-precedence target.

### Direct installs should avoid settings mutation; settings paths need scope-aware merge protection

Copying into one of Pi's native `skills/` roots needs no settings edit and is the safest direct lifecycle. If a product must point Pi at an external managed directory, add that directory to the `skills` array in the selected settings scope. Paths in global settings resolve relative to `~/.pi/agent`; paths in project settings resolve relative to `.pi`; both accept absolute and home-relative paths. Arrays support globs and include/exclude operators.

The effective settings merge is not an append merge for arrays: first-party `settings-manager.ts` deep-merges nested objects but lets a project array override the corresponding global array. The resource loader separately processes global/project resource records with project precedence, but an external installer that naively rewrites JSON can still delete user entries or introduce a surprising scope override.

**Evidence:** [Settings: Resources and Project Overrides](https://pi.dev/docs/latest/settings) defines both files, relative-path bases, resource arrays, patterns, and nested-object merging. First-party [`settings-manager.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/settings-manager.ts) confirms arrays override in the effective merge and uses a file lock for its own scoped writes; [`package-manager.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/package-manager.ts) applies project resource precedence.

**Implication:** Prefer a native-root copy. For an external-directory route, own only one appended absolute `skills` entry in one selected scope; parse and preserve all unrelated JSON members/array elements, serialize once with a write lock or equivalent atomic replace, record the exact entry added, and remove only that exact entry on uninstall. Do not write project settings until the project is trusted. This merge/receipt behavior is adapter work, not a Pi-provided direct-install transaction.

### Packages are a supported alternative, but they are lower-fit for direct delivery and do not manage plain copied skills

Pi packages may be installed from npm, git, HTTP/git URLs, or local paths; a directory is loaded with package rules and local paths are added to settings without copying. `pi install`, `pi remove`, `pi list`, and `pi update` manage **packages**. Npm/git installation may run dependency installation, git reconciliation may reset/clean a clone then run `npm install`, and Pi warns that packages have full system access. The package format can discover `skills/` convention directories or declare `pi.skills` in `package.json`.

That lifecycle is useful only when the product deliberately chooses npm/git/local-package distribution. A plain copied skill in `~/.pi/agent/skills/`, `.pi/skills/`, or `.agents/skills/` has no Pi package registration, package version, `pi remove`, or `pi update` lifecycle.

**Evidence:** [Pi Packages: Install and Manage, Local Paths, Package Structure, and Dependencies](https://pi.dev/docs/latest/packages) distinguishes package commands, local no-copy behavior, package manifests, convention directories, and npm dependency execution. [Security](https://pi.dev/docs/latest/security) confirms package installation and tools run with local process permissions.

**Implication:** Reject “wrap every Skill Issue skill in a Pi package” as the primary architecture: it adds npm/git/package lifecycle and broader execution risk where Pi natively supports direct Agent Skills. If a future hosted distribution requires versioned sharing, offer a separate explicit Pi-package channel. Do not promise that `pi update`, `pi remove`, or package repair will update/repair a direct copy.

### Update, repair, rollback, and uninstall are adapter responsibilities for direct copies

Pi validates direct skills leniently: most invalid frontmatter produces warnings, missing `description` prevents loading, and unknown frontmatter is ignored. Missing/invalid explicit skill paths produce diagnostics in the current loader source. Pi's official docs provide reload and package management, but no direct-copy receipt, atomic replacement, repair, rollback, or uninstall command.

**Evidence:** [Skills: Validation](https://pi.dev/docs/latest/skills) specifies loading/warning behavior. First-party [`skills.ts`](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/src/core/skills.ts) reports missing/bad explicit skill paths, while [Pi Packages](https://pi.dev/docs/latest/packages) limits `remove` and `update` commands to packages.

**Implication:** A reliable direct installer should own a small receipt outside the skill directory with target scope, absolute target path, skill name, source version/digest, and whether it added a settings entry. For update/repair: stage a complete replacement directory beside the target, validate `SKILL.md`, preserve the old directory as a rollback candidate, atomically rename the staged directory into place where the platform allows it, then `/reload`/restart and force-invoke the skill. For partial recovery: if the receipt exists and the target is absent/incomplete, restore the last known-good staged/backup copy or reinstall source; if ownership cannot be proven, stop with a conflict instead of deleting data. For uninstall: remove only the receipt-owned directory and receipt-owned settings entry, then `/reload`. These mechanics are an implementation recommendation, not an officially supplied Pi direct-install lifecycle.

### Current feature gate and adapter fit

Pi `0.50.0` introduced the unified ResourceLoader, resource settings arrays, `--skill`, and `/reload`; current research-time Pi is `0.80.10`. The recommended direct artifact and native roots are therefore current supported behavior for a Pi installation at or above `0.50.0`. For earlier Pi versions, the precise `--skill` and reload lifecycle is unsupported by this research; install logic should check `pi --version` and present an upgrade requirement or use a separately validated compatibility path.

**Evidence:** [Pi 0.50.0 release note](https://pi.dev/news/releases/0.50.0) records the ResourceLoader/settings/reload release, and [Pi 0.80.10 release note](https://pi.dev/news/releases/0.80.10) plus the [npm registry record](https://www.npmjs.com/package/%40earendil-works/pi-coding-agent) establish current availability at research time.

**Implication:** Common adapter fit is high: emit a standards-compliant Agent Skill directory and use direct copy into Pi's documented roots. Bespoke work is limited to target selection, collision/receipt handling, atomic filesystem operations, config preservation only for external-path mode, and user-facing trust/reload verification. An extension, package manifest, CLI plugin, SDK integration, or TUI component is unnecessary for the requested direct-install path.

## Notes

- **Unsupported:** Current official docs describe startup scanning and explicit `/reload`; they do not document automatic filesystem watching for direct skill edits. Do not claim hot watcher behavior.
- **Caveat:** Pi permits `name` different from the skill directory, but matching the directory remains the safer cross-harness convention and simplifies ownership/rollback.
- **Rejected lower-fit interpretation:** `pi install /path` is supported for a local *Pi package* and avoids copying, but still registers a package in settings and is therefore outside a direct-copy-first product architecture.
- **Useful terms:** `Agent Skills`, `SKILL.md`, `~/.pi/agent/skills`, `~/.agents/skills`, `.pi/skills`, `.agents/skills`, `--skill`, `/reload`, `defaultProjectTrust`, `enableSkillCommands`.

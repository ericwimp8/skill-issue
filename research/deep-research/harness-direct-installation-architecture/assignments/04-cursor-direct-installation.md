# Cursor Direct Installation

## Assignment

**Goal:** Establish a supportable direct, CLI-managed installation lifecycle for a Skill Issue agent skill in Cursor, and identify where that lifecycle differs among the Cursor desktop/IDE, CLI, and Cloud/Background Agent surfaces.

**Scope:** Current Cursor-owned documentation, product changelogs, marketplace materials, and Cursor support/forum statements about Agent Skills, project/user scopes, security, and Background Agents. The target is a direct filesystem adapter, rather than publication of a Cursor Marketplace plugin.

**Exclusions:** Local Cursor installations, third-party tutorials, reverse-engineered cache locations, and an implementation of the installer. Grok is considered only as a Cursor-selectable model, not as a customization or installation surface.

## Sources

- [Cursor Agent Skills documentation](https://cursor.com/docs/skills) — primary product documentation for skill roots, `SKILL.md` frontmatter, discovery, resource directories, view/import flow, and migration behavior.
- [Cursor 2.4 changelog, January 22 2026](https://cursor.com/changelog/2-4) — first-party release record that Agent Skills work in the editor and CLI.
- [Customize Cursor changelog, Cursor 3.9, June 22 2026](https://cursor.com/changelog/customize) — first-party release record for managing skills at user, team, and workspace levels.
- [Cursor CLI: Using Agent](https://docs.cursor.com/en/cli/using) and [CLI parameters](https://docs.cursor.com/en/cli/reference/parameters) — primary CLI behavior, shared rules system, interactive/non-interactive modes, session commands, and update command.
- [Cursor CLI permissions](https://docs.cursor.com/cli/reference/permissions) and [Cursor 2.5 sandbox controls](https://cursor.com/changelog/2-5) — primary permission policy and sandbox/network controls.
- [Background Agents documentation](https://docs.cursor.com/background-agent) and [Background Agents API overview](https://docs.cursor.com/background-agent/api/overview) — primary remote runtime, GitHub clone/branch, environment, and security model.
- [Cursor security page](https://cursor.com/security) — primary security posture and Workspace Trust setting.
- [Cursor Marketplace announcement](https://cursor.com/blog/marketplace) and [Marketplace](https://cursor.com/marketplace) — first-party description of plugin packaging and Marketplace installation.
- [Cursor 2.5 release](https://cursor.com/changelog/2-5) and [Cursor 3.9 release](https://cursor.com/changelog/customize) — version gates for plugins and centralized customization management.
- [Cursor model documentation](https://docs.cursor.com/models) and [Grok 4.5 announcement](https://cursor.com/blog/grok-4-5) — first-party model-selection evidence.
- Cursor staff confirmations in [skill discovery discussion](https://forum.cursor.com/t/agents-skills-discovery/151764), [existing-chat reload discussion](https://forum.cursor.com/t/newly-added-agent-skills-do-not-appear-in-existing-chats-no-in-chat-skill-reload/165124), and [remote-rule issue](https://forum.cursor.com/t/remote-rule-github-imports-successfully-but-valid-skills-do-not-appear-under-skills/153884) — lower-authority operational caveats, retained only where product documentation does not specify lifecycle behavior.

## Findings

### Native direct adapter and scope

Cursor natively supports file-based Agent Skills, so a Skill Issue direct adapter is a normal Skill package rather than an emulation through rules or a Marketplace plugin. The documented discovery roots are project `.agents/skills/` and `.cursor/skills/`, and global user `~/.agents/skills/` and `~/.cursor/skills/`; the documentation also lists compatible Claude/Codex roots. The documented list has no machine-wide/system skill location. [Cursor Agent Skills documentation](https://cursor.com/docs/skills)

**Evidence:** The Cursor 2.4 release says Agent Skills are supported in both the editor and CLI and are defined in `SKILL.md` files. The current skill documentation expressly labels the four Cursor/agents roots as project or user-global. [Cursor 2.4 changelog](https://cursor.com/changelog/2-4) [Cursor Agent Skills documentation](https://cursor.com/docs/skills)

**Implication:** Use the project-local `.cursor/skills/skill-issue/` as the safest common direct-install destination: it is Cursor-native, reviewable/versionable with the repository, and is within the checked-out repository for Cloud/Background Agents. Use `~/.cursor/skills/skill-issue/` only for an explicitly user-local installation. `.agents/skills/skill-issue/` is a portable alternative, but `.cursor/skills/` is the lower-risk primary target because it is Cursor-specific and is also documented. Do not claim a system-wide direct install for Cursor; none is documented.

### Required package layout and metadata

Each discovered skill is a directory containing a file named exactly `SKILL.md`. The required YAML frontmatter is `name` and `description`; `name` uses lowercase letters, digits, and hyphens and must equal the containing skill-directory name. Optional fields include `paths`, `disable-model-invocation`, and `metadata`. [Cursor Agent Skills documentation](https://cursor.com/docs/skills)

**Evidence:** Cursor’s documented example is `.agents/skills/my-skill/SKILL.md`; its frontmatter table makes `name` and `description` mandatory and says `name` must match the parent folder. The documentation supports recursive grouping beneath a skills root and nested project skill roots for monorepo scoping. [Cursor Agent Skills documentation](https://cursor.com/docs/skills)

**Implication:** The direct adapter should install a single self-contained directory such as:

```text
<workspace>/.cursor/skills/skill-issue/
  SKILL.md
  scripts/           # optional executable helpers
  references/        # optional on-demand documentation
  assets/            # optional static files, including template assets
```

Set `name: skill-issue` and a task-specific `description` in `SKILL.md`. `scripts/`, `references/`, and `assets/` are documented optional directories; the documentation does not reserve a separate `templates/` directory, so templates should live under `assets/` or another referenced subdirectory and be described in `SKILL.md`. [Cursor Agent Skills documentation](https://cursor.com/docs/skills)

### Discovery, activation, and verification

On startup Cursor discovers skills in documented roots and presents the available skills to Agent. By default the agent decides relevance from the prompt and description; users can explicitly invoke a discovered skill with `/skill-name`. Cursor exposes discovered skills in Customize → Skills, and its documentation says project and plugin skills appear in the Agent Decides area. `paths` scopes availability to matching files. `disable-model-invocation: true` makes a skill explicit-slash-only. [Cursor Agent Skills documentation](https://cursor.com/docs/skills)

**Evidence:** The docs distinguish automatic and slash invocation, explain description-driven relevance, give `paths` glob examples, and document the disable flag. The 3.9 Customize release adds user, team, and workspace management in one UI. [Cursor Agent Skills documentation](https://cursor.com/docs/skills) [Customize Cursor changelog](https://cursor.com/changelog/customize)

**Implication:** A practical direct-install verification is: (1) validate the installed folder and `SKILL.md` frontmatter before copying; (2) open the workspace in Cursor; (3) start a new Agent chat; (4) confirm the skill in Customize → Skills and `/skill-issue`; then (5) run a narrowly matching prompt and confirm it is selected or manually invoke `/skill-issue`. This verifies both discovery and use, rather than treating file existence as success. Cursor documents startup discovery but does not document a `validate-skills` CLI command.

### Restart, reload, repair, and session behavior

Cursor documents discovery at startup, but does not publish an explicit hot-reload contract. Cursor staff state that an existing chat retains the skill catalog present when it started; a new chat receives the updated catalog, while Reload Window is a fallback when a fresh skill fails to appear in the slash menu. Treat this as a product-support caveat rather than a stable public API. [Cursor Agent Skills documentation](https://cursor.com/docs/skills) [Cursor staff reload guidance](https://forum.cursor.com/t/newly-added-agent-skills-do-not-appear-in-existing-chats-no-in-chat-skill-reload/165124)

**Evidence:** The cited support response reports the catalog is fixed for an ongoing chat and recommends a new chat first. Cursor’s CLI supports session resume, so a resumed session should similarly be treated as retaining prior context rather than as a discovery test. [Cursor staff reload guidance](https://forum.cursor.com/t/newly-added-agent-skills-do-not-appear-in-existing-chats-no-in-chat-skill-reload/165124) [Cursor CLI: Using Agent](https://docs.cursor.com/en/cli/using)

**Implication:** For install, update, repair, rollback, or uninstall, close/restart the affected chat and create a new one before verification; use Developer: Reload Window or a full restart only if a new chat still misses the skill. Repair should restore the complete last-known-good skill directory, then repeat the UI/slash verification. Uninstall should remove only the installed `skill-issue` directory from the selected root and verify absence in a new chat. These file operations are a lifecycle recommendation inferred from Cursor’s directory discovery model; Cursor provides no dedicated direct-skill uninstall/rollback command.

### Safe update, collision, and configuration policy

Direct skills are self-contained folders discovered by path; no `mcp.json`, hook file, plugin manifest, settings value, or registry edit is required for a plain skill. Therefore the safe update is to stage and validate a complete replacement directory, preserve the prior directory for rollback, and replace only `<root>/skill-issue/`. Use an unambiguous, stable skill name and do not install two definitions with the same name across project and user roots unless a Cursor collision rule has been verified. [Cursor Agent Skills documentation](https://cursor.com/docs/skills)

**Evidence:** The skills docs define filesystem discovery and `SKILL.md` identity, while Marketplace plugins are a separate package type that can bundle skills, subagents, MCP servers, hooks, and rules. [Cursor Agent Skills documentation](https://cursor.com/docs/skills) [Cursor Marketplace announcement](https://cursor.com/blog/marketplace)

**Implication:** Common adapter fit is high: the reusable Skill Issue content maps directly to `SKILL.md` plus its resources. The adapter must avoid merging or overwriting `.cursor/mcp.json`, `.cursor/hooks.json`, `.cursor/cli.json`, sandbox policy, rules, or plugins just to install a skill. If Skill Issue separately requires one of those components, that component is a bespoke, separately reviewed adapter concern with its own merge semantics. Cursor documentation inspected for this assignment does not specify duplicate-skill precedence or an atomic installer protocol; those are unsupported, so prefer one project-local canonical definition and fail before collision rather than guessing precedence.

### Desktop/IDE, CLI, and Cloud/Background Agent boundaries

The desktop/IDE and Cursor CLI share the Agent Skills feature: Cursor announced editor-and-CLI support in 2.4, and the CLI documentation states that it shares the IDE rules system and reads project rules. The CLI is also a direct agent runtime (`cursor-agent`), with interactive and non-interactive modes, commands such as `status`, `update`, and `resume`. [Cursor 2.4 changelog](https://cursor.com/changelog/2-4) [Cursor CLI: Using Agent](https://docs.cursor.com/en/cli/using) [CLI parameters](https://docs.cursor.com/en/cli/reference/parameters)

**Evidence:** Agent Skills are explicitly released for editor and CLI; the CLI docs describe its agent and its project-local rule discovery. [Cursor 2.4 changelog](https://cursor.com/changelog/2-4) [Cursor CLI: Using Agent](https://docs.cursor.com/en/cli/using)

**Implication:** CLI-managed installation can safely create the project-local directory before invoking `cursor-agent`; the next new CLI session is the verification surface. The documentation explicitly confirms shared skills capability but does not separately guarantee every user-root discovery behavior for every CLI build. If portability matters, test the installed release with `/skill-issue` and retain the project-local copy as the common source.

Cloud/Background Agents run in an isolated Ubuntu machine, clone the selected GitHub repository onto a separate branch, and may use `.cursor/environment.json` to prepare dependencies. Their remote environment is not the local desktop home directory. [Background Agents documentation](https://docs.cursor.com/background-agent)

**Evidence:** Cursor says Background Agents clone the repository from GitHub, work on a separate branch, and run in isolated remote machines; its API is also GitHub-repository based. [Background Agents documentation](https://docs.cursor.com/background-agent) [Background Agents API overview](https://docs.cursor.com/background-agent/api/overview)

**Implication:** Commit the project-local `.cursor/skills/skill-issue/` adapter when Cloud/Background Agents must use it. A user-level `~/.cursor/skills/` installation remains local to the desktop/CLI user and will not transfer through the Git clone. The documentation does not separately enumerate Background Agent skill scanning, so this conclusion is limited to the version-controlled project path being present in the clone; validate discovery in an actual Background Agent run before declaring cloud parity.

### Permissions, trust, and unsafe resource handling

Skill files can direct agents to execute scripts, and Cursor says scripts may be any executable format supported by the agent implementation. Installation therefore grants instructions and optional executable content to an agent; it does not itself grant permissions. CLI permission policy is configured globally in `~/.cursor/cli-config.json` or per project in `.cursor/cli.json`, with allow/deny tokens for shell commands and file reads/writes; deny wins. Cursor’s sandbox supports network and filesystem controls, including project configuration and enterprise-enforced network policies. [Cursor Agent Skills documentation](https://cursor.com/docs/skills) [Cursor CLI permissions](https://docs.cursor.com/cli/reference/permissions) [Cursor 2.5 sandbox controls](https://cursor.com/changelog/2-5)

**Evidence:** The skills docs describe executable `scripts/`; the permissions docs provide `Shell`, `Read`, and `Write` policy examples and precedence; the 2.5 release documents network/domain controls and enterprise policy enforcement. [Cursor Agent Skills documentation](https://cursor.com/docs/skills) [Cursor CLI permissions](https://docs.cursor.com/cli/reference/permissions) [Cursor 2.5 sandbox controls](https://cursor.com/changelog/2-5)

**Implication:** Direct installers should copy only reviewed content, retain provenance/version information outside the executable instructions, and avoid silently loosening Cursor permission or sandbox policy. Keep scripts minimal and inspect their network, secret, and destructive behavior before installation. Enable Workspace Trust deliberately: Cursor says it is disabled by default, and it cautions that Workspace Trust is not protection from malicious extensions or folders. [Cursor security page](https://cursor.com/security)

Background Agents are materially higher risk: Cursor says they auto-run terminal commands, have internet access, and can be susceptible to prompt-injection-driven exfiltration; they require GitHub read-write access for edits. [Background Agents documentation](https://docs.cursor.com/background-agent)

**Evidence:** Cursor’s security guidance enumerates Background Agent GitHub access, remote VM storage, internet access, auto-run behavior, and prompt-injection risk. [Background Agents documentation](https://docs.cursor.com/background-agent)

**Implication:** A Skill Issue Cloud/Background Agent adapter should avoid scripts that need local-only credentials, require explicit repository/branch authorization, and retain the service’s command/network safeguards. Do not infer desktop trust, CLI permission allowlists, or user-scope skills into Cloud Agents.

### Marketplace boundary and feature gates

Cursor Marketplace plugins are a native distribution option, but they are not required for a direct skill installation. Plugins package one or more of skills, subagents, MCP servers, hooks, and rules; they are discovered through the Marketplace or `/add-plugin`. Cursor introduced plugins in 2.5 and later added user/team/workspace customization management in 3.9. [Cursor Marketplace announcement](https://cursor.com/blog/marketplace) [Cursor 2.5 changelog](https://cursor.com/changelog/2-5) [Customize Cursor changelog](https://cursor.com/changelog/customize)

**Evidence:** Cursor’s own Marketplace description distinguishes plugins from individual primitives and lists each bundled capability. [Cursor Marketplace announcement](https://cursor.com/blog/marketplace)

**Implication:** Direct installation is the correct common adapter when Skill Issue needs only a file-based skill and its local resources. A Marketplace plugin becomes justified only when the requested product explicitly needs managed distribution or additional Cursor-specific components. Do not disguise a direct install as a plugin install or mutate plugin caches; both create unnecessary lifecycle and update dependencies.

Agent Skills shipped in Cursor 2.4 (January 2026), plugins in 2.5 (February 2026), and centralized scope management in 3.9 (June 2026). [Cursor 2.4 changelog](https://cursor.com/changelog/2-4) [Cursor 2.5 changelog](https://cursor.com/changelog/2-5) [Customize Cursor changelog](https://cursor.com/changelog/customize)

**Evidence:** The cited releases provide the feature introduction dates.

**Implication:** The installer should probe the installed Cursor/CLI version and, more importantly, perform actual discovery verification. Current documentation presents Agent Skills as supported; it does not publish a minimum stable build number that can safely replace that probe.

### Grok model distinction

Grok is a model offered through Cursor’s model selection, alongside other providers’ models. Cursor says Grok 4.5 is available across Cursor desktop, web, iOS, CLI, and SDK. None of the model sources describe a Grok-specific skill root, manifest, installer, plugin type, or direct customization lifecycle. [Cursor model documentation](https://docs.cursor.com/models) [Grok 4.5 announcement](https://cursor.com/blog/grok-4-5)

**Evidence:** The model documentation classifies Grok under its provider/model capabilities; the announcement identifies it as available across Cursor surfaces. [Cursor model documentation](https://docs.cursor.com/models) [Grok 4.5 announcement](https://cursor.com/blog/grok-4-5)

**Implication:** Select Grok in Cursor only as the execution model for a Cursor Agent session. It has no bearing on how Skill Issue is installed, discovered, updated, or removed; the Cursor Agent Skills adapter remains the applicable customization mechanism.

## Notes

- **Unsupported:** Cursor documentation inspected here does not define duplicate skill-name precedence across roots, a transactional direct-installer API, a direct-skill checksum/signature scheme, or a system-wide skills root. A direct installer should fail on known collision and preserve a rollback copy rather than assume behavior.
- **Caveated:** The forum posts are Cursor staff product-support statements, not a versioned API contract. They support the new-chat/reload and prior remote-import caveats, but primary docs should override them when Cursor publishes an explicit lifecycle guarantee.
- **Lower-fit interpretation rejected:** Converting the Skill Issue direct adapter into `.cursor/rules` would make it persistent declarative context rather than the documented dynamic/procedural Agent Skills mechanism. Use rules only for genuinely always-on project guidance.
- **Lower-fit interpretation rejected:** A Marketplace plugin may package the skill, but it adds a distribution and component lifecycle not required for a plain `SKILL.md` direct install.
- **Useful search terms:** `Cursor Agent Skills`, `.cursor/skills`, `.agents/skills`, `SKILL.md`, `Customize Skills`, `Cursor CLI permissions`, `Background Agents`, `Cursor Marketplace`.

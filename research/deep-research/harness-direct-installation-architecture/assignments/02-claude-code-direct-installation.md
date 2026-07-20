# Claude Code Direct Installation

## Assignment

- **Goal:** Establish the supported direct, CLI-managed lifecycle for installing Skill Issue skills into Claude Code without creating or installing a Claude plugin.
- **Scope:** Current Claude Code CLI filesystem skills, including local CLI, IDE-integrated CLI, Desktop-local scheduled sessions, managed policy controls, and the distinct Cowork/cloud behavior.
- **Exclusions:** Plugin packaging and marketplace distribution except where they define the boundary of direct installation; Claude.ai uploaded skills, Agent SDK application configuration, and unverified implementation details.

## Sources

- [Claude Code: Extend Claude with skills](https://code.claude.com/docs/en/slash-commands) — primary contract for skill paths, discovery, metadata, loading, collisions, hot reload, security, and troubleshooting; accessed 2026-07-19.
- [Claude Code: Settings](https://code.claude.com/docs/en/settings) — primary contract for managed paths, scope precedence, <code>skillOverrides</code>, policy locks, and live settings reload; accessed 2026-07-19.
- [Claude Code: Configure permissions](https://code.claude.com/docs/en/permissions) — primary contract for workspace trust, permission precedence, <code>--add-dir</code>, sandboxing, and managed policy locks; accessed 2026-07-19.
- [Claude Code: Debug your configuration](https://code.claude.com/docs/en/debug-your-config) — primary contract for <code>/skills</code>, <code>/status</code>, <code>/doctor</code>, and <code>/context</code> verification; accessed 2026-07-19.
- [Claude Code: Explore the .claude directory](https://code.claude.com/docs/en/claude-directory) — primary contract for <code>~/.claude</code>, Windows user paths, and <code>CLAUDE_CONFIG_DIR</code>; accessed 2026-07-19.
- [Claude Code: Environment variables](https://code.claude.com/docs/en/env-vars) — primary contract for settings-file environment behavior and reload caveats; accessed 2026-07-19.
- [Anthropic claude-code repository](https://github.com/anthropics/claude-code) and [official plugin documentation](https://code.claude.com/docs/en/plugins) — first-party boundary between a direct skill directory and a plugin with a manifest/namespace; accessed 2026-07-19.

## Findings

### Supported Direct Contract

Claude Code has a first-class Agent Skills system: a direct skill is a directory whose required entrypoint is <code>SKILL.md</code>, with YAML frontmatter followed by Markdown instructions. Claude can choose a relevant skill automatically or the user can invoke it with <code>/skill-name</code>. This is a supported filesystem integration. [Skills documentation](https://code.claude.com/docs/en/slash-commands#where-skills-live)

**Evidence:** Anthropic documents personal and project <code>.../skills/&lt;skill-name&gt;/SKILL.md</code> locations, a direct creation walkthrough, and hot detection for those directories. The [Agent SDK skill documentation](https://code.claude.com/docs/en/agent-sdk/skills#skill-locations) independently describes the same filesystem artifacts and project/user directories.

**Implication:** The base Skill Issue adapter should copy a complete skill directory directly into an existing Claude Code skill root. It must not require <code>claude plugin install</code>, marketplace state, or a <code>.claude-plugin/plugin.json</code> manifest for ordinary Skill Issue instructions, references, templates, and scripts.

### Scopes And Exact Direct Paths

The supported self-service locations are:

| Scope | Direct path | Audience and lifecycle |
| --- | --- | --- |
| Personal/user | <code>~/.claude/skills/&lt;skill-name&gt;/SKILL.md</code> | Current user, every local project. On Windows <code>~/.claude</code> resolves to <code>%USERPROFILE%\.claude</code>; if <code>CLAUDE_CONFIG_DIR</code> is set, it replaces every <code>~/.claude</code> path. [Directory reference](https://code.claude.com/docs/en/claude-directory) |
| Project | <code>&lt;repository&gt;/.claude/skills/&lt;skill-name&gt;/SKILL.md</code> | Repository collaborators; commit it to source control to share. [Skills locations](https://code.claude.com/docs/en/slash-commands#where-skills-live) |
| Added directory | <code>&lt;added-dir&gt;/.claude/skills/&lt;skill-name&gt;/SKILL.md</code> | Loaded only when the directory is passed through <code>--add-dir</code> or <code>/add-dir</code>; <code>permissions.additionalDirectories</code> grants file access only and does not discover skills. [Additional directories](https://code.claude.com/docs/en/slash-commands#skills-from-additional-directories) |
| Managed/enterprise | Administrator-provisioned managed policy skills | Applies organization-wide. Anthropic documents the managed-settings roots: macOS <code>/Library/Application Support/ClaudeCode/</code>, Linux/WSL <code>/etc/claude-code/</code>, Windows <code>C:\Program Files\ClaudeCode\</code>; it directs skill authors to managed settings for enterprise scope. [Settings delivery paths](https://code.claude.com/docs/en/settings#settings-files) |

**Evidence:** The skills page supplies the personal, project, plugin, and enterprise resolution levels and their precedence. It also says project skills load from the start directory and parent directories up to the repository root, with nested <code>.claude/skills/</code> discovered on demand when Claude works in that subtree. [Skills discovery](https://code.claude.com/docs/en/slash-commands#automatic-discovery-from-parent-and-nested-directories)

**Implication:** Offer <code>project</code> as the default shared install target and <code>user</code> as the personal target. Do not fabricate a local-only skill scope: <code>.claude/settings.local.json</code> is a settings scope, not a documented direct skill root. Enterprise installation belongs to the organization’s policy deployment process. The public Anthropic documentation retrieved here identifies the managed system roots but does not state the managed *skills subpath* explicitly; therefore an adapter must not write a guessed system skill directory and should require an administrator-provided target or use the documented managed deployment route.

### Shape, Naming, And Portable Content

Use one folder per skill with a <code>SKILL.md</code> entrypoint. <code>SKILL.md</code> requires Markdown instructions; YAML frontmatter is supported and <code>description</code> is recommended because it determines automatic use. The ordinary direct command name comes from the containing directory, not the frontmatter <code>name</code>: <code>.claude/skills/deploy-staging/SKILL.md</code> gives <code>/deploy-staging</code>. <code>name</code>, if supplied, is a display label and must be lowercase letters, numbers, and hyphens, at most 64 characters. [Frontmatter and command naming](https://code.claude.com/docs/en/slash-commands#frontmatter-reference)

**Evidence:** Anthropic documents optional sibling supporting files, including references, examples, templates, and executable scripts, and requires that <code>SKILL.md</code> tell Claude what they contain and when to load or run them. It recommends keeping <code>SKILL.md</code> under 500 lines and moving detail into supporting files. [Supporting files](https://code.claude.com/docs/en/slash-commands#add-supporting-files)

**Implication:** A common Skill Issue bundle can retain <code>SKILL.md</code>, <code>references/</code>, <code>examples/</code>, <code>templates/</code>, <code>assets/</code>, and <code>scripts/</code> as ordinary files. The common layer should use relative links from <code>SKILL.md</code> and ${CLAUDE_SKILL_DIR} for bundled script/file paths. Keep Claude-specific frontmatter in a narrowly scoped Claude adapter: <code>disable-model-invocation</code>, <code>user-invocable</code>, <code>allowed-tools</code>, <code>disallowed-tools</code>, <code>model</code>, <code>effort</code>, <code>context</code>, <code>agent</code>, <code>hooks</code>, <code>paths</code>, and <code>shell</code> are Claude Code extensions beyond the portable core. [Claude substitutions and extensions](https://code.claude.com/docs/en/slash-commands#available-string-substitutions)

### Discovery, Activation, And Verification

In a normal session, Claude receives available skill names/descriptions and loads the full <code>SKILL.md</code> body only when either the user invokes it or Claude decides it is relevant. Default skills are both user- and model-invocable; <code>disable-model-invocation: true</code> makes an action manual-only, while <code>user-invocable: false</code> hides it from the slash menu but leaves it model-invocable. [Invocation controls](https://code.claude.com/docs/en/slash-commands#control-who-invokes-a-skill)

**Evidence:** Anthropic’s troubleshooting flow says to verify the skill in “What skills are available?”, invoke it directly if user-invocable, and use <code>--debug</code> for malformed-frontmatter parse errors. Its configuration-debugging documentation adds <code>/skills</code> for source-backed available-skill listing, <code>/context</code> for loaded context, <code>/status</code> for active settings sources, and <code>/doctor</code>/<code>claude doctor</code> for diagnostics. [Skills troubleshooting](https://code.claude.com/docs/en/slash-commands#skill-not-triggering), [configuration diagnostics](https://code.claude.com/docs/en/debug-your-config)

**Implication:** A practical direct-install acceptance check is: (1) launch Claude Code from the target project, (2) use <code>/skills</code> and confirm the expected source/name, (3) invoke <code>/&lt;directory-name&gt;</code> with a harmless prompt, and (4) run <code>claude doctor</code> or <code>--debug</code> if it is missing or lacks automatic activation. Listing proves discovery; direct invocation proves the entrypoint can load; a fresh-session prompt matching the description is needed to assess auto-selection separately.

### Reload, Update, Repair, And Recovery

Claude Code watches existing personal, project, and <code>--add-dir</code> skill directories. Adding, editing, or removing a skill there takes effect in the current session without restarting; creating the top-level skills directory after session start requires restart so it can be watched. The hot-reload promise applies to <code>SKILL.md</code> text. [Live change detection](https://code.claude.com/docs/en/slash-commands#live-change-detection)

**Evidence:** A malformed YAML frontmatter still leaves the body directly invocable but removes usable metadata; Anthropic directs users to <code>--debug</code> for the parse error. Skill visibility can also be disabled by <code>skillOverrides</code>; absence means <code>on</code>, while <code>/skills</code> writes overrides to <code>.claude/settings.local.json</code>. [Troubleshooting](https://code.claude.com/docs/en/slash-commands#skill-not-triggering), [visibility overrides](https://code.claude.com/docs/en/slash-commands#override-skill-visibility-from-settings)

**Implication:** Direct lifecycle operations are filesystem operations, with no package-manager ledger:

- **Update:** replace the selected skill directory’s contents while preserving the scope root; use the current session’s live reload when its root already existed, otherwise start a new session.
- **Repair:** verify <code>SKILL.md</code>, YAML, source path, <code>/skills</code>, and <code>skillOverrides</code>; then use <code>--debug</code> or <code>claude doctor</code> for parse/configuration evidence.
- **Partial recovery:** remove only the incomplete skill directory, then re-copy the known complete bundle and repeat discovery verification. Existing unrelated skills and settings stay untouched.
- **Rollback:** restore the prior saved skill directory, then verify the restored command through <code>/skills</code> and direct invocation. This is an adapter-owned filesystem rollback; Anthropic documents no direct-skill version store or rollback command.
- **Uninstall:** remove that one <code>&lt;skill-name&gt;</code> directory; hot detection removes it from the current session if the parent root is already watched. Preserve the root and every other skill.

### Collision And Configuration Safety

Precedence is enterprise, personal, then project; a direct skill at any of those levels also replaces a bundled skill of the same name. A direct skill wins over an old <code>.claude/commands/&lt;name&gt;.md</code> collision. Plugin skills are namespaced as <code>plugin-name:skill-name</code>, so they do not conflict with direct names. Nested project skills with collisions remain available under directory-qualified names such as <code>/apps/web:deploy</code>; the unqualified command uses the project-root skill. [Collision behavior](https://code.claude.com/docs/en/slash-commands#where-skills-live)

**Evidence:** <code>skillOverrides</code> values are <code>on</code>, <code>name-only</code>, <code>user-invocable-only</code>, and <code>off</code>; plugin skills are outside this setting. Settings precedence is managed, command-line, local, project, then user; arrays merge/deduplicate while scalars at a higher scope override lower values. [Settings precedence](https://code.claude.com/docs/en/settings#settings-precedence)

**Implication:** The installer should preflight the target folder and command name across direct roots, legacy commands, nested project directories, and existing plugins. Default to refusing overwrite when the same direct scope has a non-Skill-Issue folder, and report effective precedence rather than silently shadowing it. The normal direct install must modify no settings. If it deliberately changes a setting such as <code>skillOverrides</code>, it must parse and merge the existing JSON, change only the intended key, retain an exact pre-change backup, and restore that key only during uninstall; a whole-file overwrite is unsafe.

### Trust, Permissions, And Security

Skill text is executable instruction supply chain content. <code>allowed-tools</code> pre-approves listed tools only for the invocation turn; it does not restrict other tools, and baseline permission rules still apply. For project skill directories, that grant takes effect only after accepting workspace trust; Anthropic warns that a project skill can grant broad tool access. [Skill pre-approval](https://code.claude.com/docs/en/slash-commands#pre-approve-tools-for-a-skill)

**Evidence:** Permission denies are enforced across settings sources; sandboxing separately constrains Bash filesystem/network access. Managed administrators can lock the <code>skills</code> surface with <code>strictPluginOnlyCustomization</code>, causing <code>~/.claude/skills/</code> and <code>.claude/skills/</code> to be skipped while plugin, bundled, and managed-policy skills remain. They can also disable inline shell preprocessing for user/project/plugin/additional-directory skills through <code>disableSkillShellExecution</code>. [Permissions and policy locks](https://code.claude.com/docs/en/permissions#managed-settings), [skill shell-execution setting](https://code.claude.com/docs/en/settings)

**Implication:** Inspect every bundle before installation, avoid <code>allowed-tools</code> and inline shell commands unless the skill genuinely needs them, and require user review/trust before a project-scoped install can activate grants. Treat a policy lock as a true blocker: do not bypass it with a plugin sideload or another filesystem root. A direct installer should report the lock and request an approved managed/plugin delivery path.

### Surface Boundaries And Direct-Adapter Fit

The direct local CLI contract also applies when Claude Code is used through its supported IDE integration: Anthropic says the settings hierarchy applies whether Claude Code runs from the CLI, VS Code extension, or JetBrains IDE. Desktop scheduled tasks are local and load skills from the same locations as another local session. [Settings hierarchy](https://code.claude.com/docs/en/settings#settings-precedence), [skills in Cowork and cloud sessions](https://code.claude.com/docs/en/slash-commands#skills-in-cowork-and-cloud-sessions)

**Evidence:** Cowork and cloud sessions do not read the user’s local <code>~/.claude/skills/</code>. Both use skills enabled for the claude.ai account at session start; cloud can instead use <code>.claude/skills/</code> committed to the cloned repository or a repository-declared plugin. [Cowork/cloud behavior](https://code.claude.com/docs/en/slash-commands#skills-in-cowork-and-cloud-sessions)

**Implication:** This adapter is a strong fit for the local Claude Code CLI, its IDE-backed local session, and Desktop-local scheduled tasks. It is only partially portable to cloud: use committed project skills rather than user-scope direct install. It does not install a Cowork/cloud personal skill and must not claim that a local copy will sync. Plugins are a separate delivery mechanism for namespacing, agents, hooks, MCP, marketplace distribution, or managed distribution; they are not a prerequisite for a normal direct skill. [Plugin structure boundary](https://code.claude.com/docs/en/plugins#plugin-structure-overview)

### Version Gates And Unsupported Details

The current relevant gates are: directory-entry symlinks at enterprise/personal/project scope require Claude Code v2.1.203+; <code>${CLAUDE_PROJECT_DIR}</code> requires v2.1.196+; <code>off</code> visibility hiding from SDK/Remote Control listings and stacked skills require v2.1.199+; duplicate-identical skill re-invocations are de-duplicated from v2.1.202+; and <code>skillOverrides</code> requires v2.1.129+. [Skills version notes](https://code.claude.com/docs/en/slash-commands), [settings key reference](https://code.claude.com/docs/en/settings)

**Evidence:** Anthropic’s current docs attach these gates directly to the named features. They do not state a CLI version introduction for baseline filesystem skills or publish an atomic-install API, direct-skill inventory command, package-manager record, checksum policy, automatic rollback mechanism, or a public managed-policy *skills* subdirectory.

**Implication:** Require a current Claude Code installation and call <code>claude --version</code> before using a gated enhancement. Keep baseline installation to real directories/copies, not symlink-only deployment, unless v2.1.203+ is confirmed. The unlisted lifecycle mechanisms above are adapter procedures, not claimed Claude Code features; the managed system subpath and any package-manager style repair/rollback are **unsupported as direct claims** pending first-party documentation.

## Notes

- **Rejected interpretation:** “Claude Code requires a native plugin to use skills.” Anthropic’s direct personal/project <code>SKILL.md</code> locations disprove this; plugins add namespacing and extension surfaces rather than enabling baseline skill loading.
- **Rejected interpretation:** a direct user-scope install propagates to Cowork/cloud. Anthropic explicitly says those remote sessions do not read local <code>~/.claude/skills/</code>.
- **Caveat:** Enterprise managed skills are documented as a supported source and managed roots are documented, but the exact managed <code>SKILL.md</code> directory layout was not exposed by the first-party pages inspected. Treat any guessed path as unsupported.
- **Useful search terms:** <code>Claude Code skills</code>, <code>skillOverrides</code>, <code>strictPluginOnlyCustomization</code>, <code>disableSkillShellExecution</code>, <code>workspace trust</code>, <code>CLAUDE_CONFIG_DIR</code>, <code>--add-dir</code>.

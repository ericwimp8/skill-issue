# Grok Direct Installation

## Assignment

**Goal:** Determine whether xAI provides a standalone Grok coding-agent harness that Skill Issue can manage directly, and map only the documented installation, customization, discovery, verification, lifecycle, and safety boundaries relevant to skills or equivalent customizations.

**Scope:** Current first-party xAI Grok Build documentation, with current Cursor documentation used only to distinguish an xAI model hosted by Cursor from the standalone xAI harness.

**Exclusions:** Consumer Grok web/mobile apps, third-party `grok-cli` projects, reverse-engineered interfaces, and undocumented filesystem or plugin-manifest assumptions.

## Sources

- [xAI: Grok Build getting started](https://docs.x.ai/build/overview) — first-party standalone-harness overview, installation, authentication, headless/ACP usage, and `grok inspect`; updated July 6, 2026.
- [xAI: Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces) — first-party discovery locations and supported extension categories; updated July 4, 2026.
- [xAI: AGENTS.md](https://docs.x.ai/build/features/project-rules) — first-party instruction-file discovery, order, and verification; updated July 4, 2026.
- [xAI: Settings](https://docs.x.ai/build/settings) and [Enterprise Deployments](https://docs.x.ai/build/enterprise) — first-party scope, merge priority, managed policy, sandbox, permission, and authentication behavior; updated July 6 and June 16, 2026.
- [xAI: CLI Reference](https://docs.x.ai/build/cli/reference), [Modes and Commands](https://docs.x.ai/build/modes-and-commands), and [Hooks](https://docs.x.ai/build/features/hooks) — first-party command, collision, activation, and trust behavior; updated July 2–4, 2026.
- [Cursor: Models & Pricing](https://docs.cursor.com/models) — host documentation identifying Grok models as xAI-provided models within Cursor, rather than documentation for a Grok harness.

## Findings

### Standalone Harness Is Officially Supported

xAI documents **Grok Build** as a standalone, extensible coding agent operated through an interactive TUI, headless CLI, or Agent Client Protocol process. It is installed directly from xAI's shell/PowerShell installer; xAI's enterprise documentation also names `npm install -g @xai-official/grok` as an alternative distribution route. First launch authenticates through a browser or, for headless contexts, with `XAI_API_KEY`. [Grok Build getting started](https://docs.x.ai/build/overview) and [Enterprise Deployments](https://docs.x.ai/build/enterprise#network-requirements) provide these direct-install paths.

**Evidence:** `grok inspect` reports discovered instructions, skills, plugins, hooks, and MCP servers; `grok plugin` has supported install, uninstall, update, enable, disable, details, and validate subcommands. [Grok Build getting started](https://docs.x.ai/build/overview#custom-models) [CLI Reference](https://docs.x.ai/build/cli/reference)

**Implication:** **Supported direct-harness target.** Skill Issue may target the installed `grok` CLI and Grok Build's documented customization roots; it does not need to treat Grok only as an API model or as a feature embedded in another IDE.

### Cursor Access Is a Different Host Boundary

Cursor's model catalogue lists Grok variants with provider `xAI`, but Cursor documents those as models hosted by the provider, a trusted partner, or Cursor, and exposes them through Cursor's own product and agent capabilities. It is evidence of model availability in Cursor, not evidence that xAI's Grok Build filesystem, extension lifecycle, or permission model applies inside Cursor. [Cursor Models & Pricing](https://docs.cursor.com/models)

**Evidence:** xAI separately calls Grok Build the coding agent and says the same Grok model is also available through the xAI API; this separates the model from its standalone agent harness. [Grok Build getting started](https://docs.x.ai/build/overview#use-grok-45-on-the-api)

**Implication:** **Rejected interpretation:** "Grok in Cursor" cannot be installed or managed by a Grok direct-install adapter. Any Cursor integration belongs to Cursor's own adapter and contracts.

### Supported Skill and Instruction Scopes

Grok Build discovers skills in project `.grok/skills/` directories while walking upward to the repository root, in user `~/.grok/skills/`, from enabled plugin `skills/` directories, and from extra roots configured in `[skills] paths` in `~/.grok/config.toml`. User-invocable skills appear as slash commands. [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces#skills)

Project rules have an independently documented, broader layout: Grok loads global rules in `~/.grok/`, then instruction files from repository root through the working directory; deeper files win conflicts. Accepted names include `AGENTS.md`, `Agents.md`, `AGENT.md`, `CLAUDE.md`, `Claude.md`, and `CLAUDE.local.md`, plus Markdown files in `.grok/rules/` (and compatible `.claude/rules/` and `.cursor/rules/`). Git-ignored instruction files are skipped. [AGENTS.md](https://docs.x.ai/build/features/project-rules#discovery)

**Evidence:** User settings are `~/.grok/config.toml` (or `$GROK_HOME/config.toml`); project `.grok/config.toml` is limited to shared MCP, plugin, and permission settings. [Settings](https://docs.x.ai/build/settings#scopes)

**Implication:** A direct Skill Issue adapter can safely support two explicit install scopes: project skill folder `<repo>/.grok/skills/<skill-name>/` and user skill folder `~/.grok/skills/<skill-name>/`. It may support instructions only as a distinct feature with the documented `AGENTS.md`/`.grok/rules/` precedence; it must not silently use project `config.toml` as a general skill configuration surface.

### Format, Naming, and Collision Boundaries

xAI defines a skill as a reusable folder containing Markdown instructions, scripts, and resources, but the inspected official documentation does **not** publish a required skill-file name, front matter/manifest schema, metadata field set, naming grammar, or duplicate-resolution rule for skills. xAI does document that slash-command collisions are addressable through a qualified name such as `/local:commit`. [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces#skills) [Modes and Commands](https://docs.x.ai/build/modes-and-commands#skills-as-commands)

**Evidence:** xAI publishes explicit locations for skills and plugins, and a `grok inspect` command that enumerates discovered configuration, yet the same docs omit a skill manifest contract and an explicit collision-precedence contract. [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces) [CLI Reference](https://docs.x.ai/build/cli/reference)

**Implication:** **Caveated common fit.** The documented folder model fits a common Skill Issue payload of Markdown plus resources, but the adapter must regard exact `SKILL.md` shape, metadata, source ownership, and collision selection as **unsupported by the inspected xAI sources**. Do not invent a Grok-specific manifest or promise collision-free activation based only on the folder name; install under a namespaced Skill Issue-controlled directory and verify the resolved discovery result.

### Discovery, Activation, and Verification Are Available

`grok inspect` is the authoritative documented verification surface for the current directory: it lists discovered configuration/rules, skills, plugins, hooks, and MCP servers. The TUI extensions modal can be opened with `/skills`, `/plugins`, `/hooks`, or `/mcps`; the commands select a tab in that same modal. [Grok Build getting started](https://docs.x.ai/build/overview#custom-models) [Modes and Commands](https://docs.x.ai/build/modes-and-commands#core-tui-commands)

**Evidence:** The CLI exposes `grok plugin validate` and `grok plugin details`, while `grok inspect [--json]` is explicitly available for configuration discovery. [CLI Reference](https://docs.x.ai/build/cli/reference)

**Implication:** After creating or replacing a managed skill directory, run `grok inspect` in the target project (and use `--json` when machine-readable evidence is needed). A successful file write alone is insufficient evidence of activation. Plugin-specific validation is useful only for a plugin integration; it is not documented as validation for a loose skill folder.

### Lifecycle Support Is Partial at the Skill Layer

xAI documents lifecycle commands for **plugins** (`install`, `uninstall`, `update`, `enable`, `disable`, `details`, `validate`) and for marketplaces, but does not document analogous CLI commands, ownership metadata, transactional replacement, repair, rollback, or uninstall semantics for standalone skill directories. [CLI Reference](https://docs.x.ai/build/cli/reference)

**Evidence:** Skills are documented solely as folders discovered from roots; plugin and marketplace directories have separate documented loading locations and command families. [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces#skills) [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces#plugins)

**Implication:** **Blocked for host-managed skill lifecycle.** Skill Issue can implement its own atomic staging, backup, replace, removal, and post-operation `grok inspect` recovery around a uniquely owned skill directory, but it cannot claim xAI-native rollback, repair, or partial-recovery support. A plugin route is a separate, richer integration and requires a validated xAI plugin layout/manifest before use.

### Configuration Merging and Safety Controls Require Conservative Writes

Grok's global configuration layers, from lowest to highest, are `/etc/grok/managed_config.toml`, `~/.grok/managed_config.toml`, `~/.grok/config.toml`, `~/.grok/requirements.toml`, and `/etc/grok/requirements.toml`; requirement settings cannot be overridden by lower configuration, remote settings, or user config. Project config is constrained to MCP servers, plugins, and permission rules. [Enterprise Deployments](https://docs.x.ai/build/enterprise#configuration) [Settings](https://docs.x.ai/build/settings#scopes)

**Evidence:** The sandbox is set at process startup and is irreversible; documented profiles limit filesystem/network access. Permission checks are independent from the sandbox, deny rules override allow rules, and `always-approve` can be administratively disabled only from a root-owned system source. [Enterprise Deployments](https://docs.x.ai/build/enterprise#security-controls)

**Implication:** Direct skill installation should avoid editing shared `config.toml` unless an explicitly requested extra skill/plugin path is required, preserve existing TOML rather than replacing it, and fail with a managed-policy blocker when the desired path or action is disallowed. A Skill Issue installer must not enable `--always-approve`, weaken sandboxing, add hooks, or alter permission rules as part of ordinary skill installation.

### Hooks and Plugins Carry Separate Trust Risk

Project hooks live in `<project>/.grok/hooks/*.json` and require explicit trust via `/hooks-trust` or `--trust`; trust is recorded in `~/.grok/trusted_folders.toml` and also covers project MCP and LSP servers. Plugin hooks run on lifecycle events and may execute commands or call HTTP endpoints. [Hooks](https://docs.x.ai/build/features/hooks#configuration)

**Evidence:** Plugins can add agents, hooks, MCP servers, LSP servers, and skills; the extensions page distinguishes those from a standalone skill folder. [Skills, Plugins & Marketplaces](https://docs.x.ai/build/features/skills-plugins-marketplaces#plugins)

**Implication:** **Rejected route for ordinary Skill Issue skills:** do not package or install hooks, MCP, LSP, or an unvalidated plugin merely to deliver Markdown instructions. Escalate to explicit user confirmation plus host trust/permission review only if a future requested skill genuinely requires executable extensions.

## Notes

- **Feature/version gate:** This assessment is for the current Grok Build documentation, which labels its default coding model `grok-build` and documents the standalone `grok` CLI. Re-check `grok version` and `grok inspect` against the target machine before an implementation claims availability. [Settings](https://docs.x.ai/build/settings#example-configtoml) [CLI Reference](https://docs.x.ai/build/cli/reference)
- **Unsupported observation:** The inspected official sources do not specify a loose-skill manifest, exact entry filename, on-disk version marker, duplicate precedence, or a skills-specific lifecycle command. Those gaps rule out a host-native repair/rollback promise.
- **Useful search terms:** `Grok Build`, `grok inspect`, `.grok/skills`, `GROK_HOME`, `grok plugin validate`, `/hooks-trust`, `requirements.toml`.

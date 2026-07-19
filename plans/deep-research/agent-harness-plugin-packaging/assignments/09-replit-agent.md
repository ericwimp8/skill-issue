# Replit Agent Packaging

## Assignment

**Goal:** Determine the current official packaging and distribution mechanisms around Replit Agent, and identify what Skill Issue must implement to deliver skills, Markdown guidance, scripts, a standalone CLI, configuration, references, assets, and supporting files as a coherent experience.

**Scope:** Current public Replit documentation and primary sources for Agent Skills, workspace customization, MCP, connectors, project configuration, dependencies, imports/remixes/templates, secrets, permissions, publishing, discovery, invocation, updates, and distribution. Evidence was inspected on 2026-07-19.

**Exclusions:** Historical behavior that is absent from current official documentation; local Replit configuration; implementation details inferred from private APIs; and claims about unpublished marketplace submission processes.

## Sources

- [Replit documentation index (`llms.txt`)](https://docs.replit.com/llms.txt), current index inspected 2026-07-19.
- [Agent Skills](https://docs.replit.com/features/agent/skills), Replit Docs, current reference.
- [Use Agent Skills](https://docs.replit.com/build/use-agent-skills), Replit Docs, current workflow guide.
- [Skills directory](https://docs.replit.com/features/agent/skills-directory), Replit Docs, current discovery and installation reference.
- [Agent Customization](https://docs.replit.com/features/agent/agent-customization), Replit Docs, current workspace-wide instructions and skills reference.
- [Agent Skills specification](https://agentskills.io/specification), open specification linked by Replit Docs.
- [`vercel-labs/skills`](https://github.com/vercel-labs/skills), primary repository for the `npx skills` CLI linked by Replit Docs; `main` inspected 2026-07-19.
- [Custom Templates](https://docs.replit.com/teams/custom-templates), Replit Enterprise documentation.
- [Developer Frameworks](https://docs.replit.com/features/project-setup/developer-frameworks), Replit Docs, current template-removal notice.
- [Import from a provider](https://docs.replit.com/build/import-from-providers), Replit Docs, current Git provider and ZIP import behavior.
- [Remix an app](https://docs.replit.com/build/remix-an-app), Replit Docs, current project-copy workflow.
- [Open in Replit](https://docs.replit.com/features/integrations/open-in-replit), Replit Docs, prompt-link format.
- [Replit App Configuration](https://docs.replit.com/features/project-setup/configuration), Replit Docs, `.replit` and `replit.nix` reference.
- [Dependency Management](https://docs.replit.com/features/project-setup/dependency-management), Replit Docs, package-manager and Nix behavior.
- [Package Firewall](https://docs.replit.com/features/security/package-firewall), Replit Docs, install-time supply-chain controls.
- [Connect via MCP](https://docs.replit.com/build/connect-via-mcp), Replit Docs, Agent MCP setup.
- [MCP list](https://docs.replit.com/features/mcp/overview), Replit Docs, catalog, authorization, authentication, and scanner behavior.
- [MCP Install Links](https://docs.replit.com/features/mcp/install-links), Replit Docs, one-click custom-server link format.
- [Agent Integrations](https://docs.replit.com/features/integrations/overview), Replit Docs, integration categories and connector scope.
- [Managing Your Connectors](https://docs.replit.com/replitai/managing-connectors), Replit Docs, workspace and enterprise governance.
- [Secrets](https://docs.replit.com/core-concepts/project-editor/app-setup/secrets), Replit Docs, app/account secrets and remix visibility.
- [Replit shared responsibility model](https://docs.replit.com/references/security/shared-responsibility-model), Replit Docs, human approval and security responsibilities.
- [Publishing](https://docs.replit.com/learn/projects-and-artifacts/replit-deployments), Replit Docs, snapshot and update model.

## Findings

### Finding 1: Replit's reusable Agent capability unit is an Agent Skill, not a general plugin

Replit currently exposes Agent Skills as its native reusable instruction package. A skill is a directory containing `SKILL.md` and may carry supporting files. Project skills live under `.agents/skills/`; workspace skills are centrally managed; curated Replit skills can be attached to one message or used as a new-project starting point. Replit's current documentation index has first-class sections for Skills, MCP, connectors, imports, and project configuration, but no current Replit extension/plugin SDK or extension marketplace entry.

#### Evidence

[Agent Skills](https://docs.replit.com/features/agent/skills) defines a skill as a folder containing `SKILL.md` plus supporting files and identifies project, workspace, and enterprise scopes. [Use Agent Skills](https://docs.replit.com/build/use-agent-skills) documents message attachment, project installation, and new-project starting points. The current [documentation index](https://docs.replit.com/llms.txt) lists Agent Skills, Agent Customization, MCP, connectors, imports, and configuration; it does not list a current extensions or plugin-authoring system. This absence establishes only that no current public official extension surface was found, not that no private or legacy mechanism exists.

#### Implication

Skill Issue should target Agent Skills as the native Replit Agent capability layer. It should not model a Replit deliverable as a VS Code-style plugin or assume a general extension host. Any UI pane, custom command palette contribution, background extension process, or arbitrary editor API is unsupported by the current public evidence.

### Finding 2: A skill can carry Skill Issue guidance, scripts, references, assets, and arbitrary supporting files

The skill directory is broad enough to hold most of Skill Issue's content bundle. Replit explicitly permits supporting files, and the underlying open Agent Skills specification defines optional `scripts/`, `references/`, and `assets/` directories plus additional files. Agent performs progressive disclosure: it sees name and description metadata on every chat and reads the full body only when relevant.

#### Evidence

[Agent Skills](https://docs.replit.com/features/agent/skills) says the folder can contain `SKILL.md` and any supporting files. [Custom Templates](https://docs.replit.com/teams/custom-templates) shows skills with `references/` and `assets/`. The [Agent Skills specification](https://agentskills.io/specification) explicitly defines `scripts/` as executable code, `references/` as documentation, `assets/` as templates/resources, and permits other directories. Replit also warns that externally installed skills contain arbitrary content and code and must be reviewed in [Skills directory](https://docs.replit.com/features/agent/skills-directory).

#### Implication

Skill Issue can package each Replit-facing skill as `.agents/skills/<name>/SKILL.md` with its scripts, references, assets, and supporting material colocated. Compatibility requirements should be declared in skill metadata where useful. The skill description must state when to invoke it because description matching controls automatic loading. Large always-on guidance belongs outside the skill only when it truly must apply to every message.

### Finding 3: Skill installation and invocation are good, but skill installation is not a full bundle installer

Users can discover skills in Replit's Skills pane, install from the community directory, use the `npx skills` CLI, upload a skill folder or ZIP at workspace scope, attach a curated skill to a message, rely on automatic relevance matching, select it from the picker, or explicitly invoke a workspace skill with `/<skill-name>`. These mechanisms install or select skill directories; the official docs do not say they atomically install root project configuration, package dependencies, secrets, connector registrations, or editor integrations.

#### Evidence

[Skills directory](https://docs.replit.com/features/agent/skills-directory) documents picker use, project installation into `.agents/skills/`, and `npx skills ... -a replit`. [Agent Customization](https://docs.replit.com/features/agent/agent-customization) documents workspace skill folder/ZIP upload, automatic matching, picker selection, and direct slash invocation. The linked [`vercel-labs/skills`](https://github.com/vercel-labs/skills) repository identifies Replit's project path as `.agents/skills/` and supports add, list, find, remove, and update operations. None of those sources defines a Replit-specific manifest capable of installing unrelated root files and account-level integrations as one transaction.

#### Implication

The native skill package should be the discoverable Agent entry point, while a separate Skill Issue bootstrap command must own full-bundle installation. A skill may instruct Agent to run that command, but Skill Issue should present and log the resulting filesystem and dependency changes clearly rather than treating skill selection as implicit consent to install everything.

### Finding 4: Project distribution can deliver the complete filesystem bundle in one starting experience

A Replit project can contain skills, `replit.md`, `.replit`, `replit.nix`, package manifests, source, CLI code, scripts, references, and assets together. Git import, ZIP import, or remix copies project files, while an Enterprise custom template explicitly combines instructions, skills, project context, runtime configuration, dependencies, and the rest of the repository. This is the strongest available "one bundled experience," but it is a project starting point rather than a reusable capability installed into an arbitrary existing project.

#### Evidence

[Import from a provider](https://docs.replit.com/build/import-from-providers) says GitHub import brings repository files/folders, dependency files, and common run/build configuration; ZIP import brings project files, directory structure, dependency files, and run/build defaults. [Custom Templates](https://docs.replit.com/teams/custom-templates) specifies a tree containing `.replit`, `replit.md`, `custom_instruction/instructions.md`, `.agents/skills/`, dependencies, source, and arbitrary project files, all copied when the template is forked. [Remix an app](https://docs.replit.com/build/remix-an-app) defines a remix as a new app created from an app the user can access.

#### Implication

Skill Issue should publish a canonical Replit-ready repository that imports cleanly and optionally a public Replit app that users can remix. Its root should include the Skill Issue CLI, `.agents/skills/`, minimal `replit.md`, `.replit`, `replit.nix` when needed, locked language dependencies, references, scripts, and assets. This gives new projects the closest equivalent to a plugin bundle. Existing projects still need a bootstrap installer or selective copy flow.

### Finding 5: Public templates and enterprise templates are materially different distribution channels

Replit removed the old public language/framework template system. Community starting points now come from Gallery apps/remixes, imports, Agent, or skill starting points. Enterprise custom templates are a distinct paid, admin-managed mechanism that can inject immutable organization instructions, expose skills, copy configuration and files, and appear directly in the Agent input box.

#### Evidence

[Developer Frameworks](https://docs.replit.com/features/project-setup/developer-frameworks) states that language and framework starter templates were removed, that `/templates` redirects to Gallery, and that users should start through Agent, Git import, or Gallery remix. It separately identifies Enterprise custom templates. [Custom Templates](https://docs.replit.com/teams/custom-templates) states they are Enterprise-only, require organization-admin management, and fork all template files while injecting `custom_instruction/instructions.md` and discovering `.agents/skills/`.

#### Implication

Skill Issue should not promise a general public Replit template marketplace listing equivalent to the enterprise feature. Public distribution should use a GitHub import/remix plus Skills directory presence. An enterprise edition can additionally provide a documented template layout and admin installation guide, but that channel cannot be the default open-source onboarding path.

### Finding 6: Project guidance has three distinct scopes with different update semantics

Replit separates always-on project memory (`replit.md`), enterprise template instructions (`custom_instruction/instructions.md`), and on-demand skills. Workspace customization adds centrally managed always-on instructions and skills across projects. These surfaces should cooperate by role rather than repeat the same content.

#### Evidence

[Custom Templates](https://docs.replit.com/teams/custom-templates) states that `custom_instruction/instructions.md` is static authoritative guidance injected into the system prompt, `.agents/skills/*/SKILL.md` is loaded on demand, `replit.md` is always loaded and mutable, and `.replit` controls runtime configuration. [Agent Customization](https://docs.replit.com/features/agent/agent-customization) says workspace custom instructions apply to every project/session and workspace skills are selectively loaded; changes apply to future chats. [replit.md](https://docs.replit.com/features/project-setup/replit-dot-md) says Agent automatically reads the root file and may update it as the project evolves.

#### Implication

Skill Issue should keep `replit.md` short and project-specific, put task workflows and detailed knowledge in skills, and reserve `custom_instruction/instructions.md` for enterprise policy. The installer should detect and merge rather than overwrite an existing `replit.md`. It should avoid duplicating the full Skill Issue manual across always-on contexts.

### Finding 7: Replit project configuration and ordinary package systems can make the standalone CLI functional

Replit does not require a special plugin API to run a standalone CLI. Project dependencies can be installed through language package managers, Agent, UPM, or Nix. `.replit` configures run/build/deployment commands and modules; `replit.nix` supplies reproducible system dependencies. Replit's Package Firewall evaluates supported package installs initiated by either the user or Agent.

#### Evidence

[Dependency Management](https://docs.replit.com/features/project-setup/dependency-management) documents `npm`, `pnpm`, `pip`, Poetry, Bundler, UPM, Agent-managed installation, and Nix dependencies. [Replit App Configuration](https://docs.replit.com/features/project-setup/configuration) assigns runtime behavior, package configuration, run commands, build commands, deployment commands, modules, and environment setup to `.replit` and `replit.nix`. [Package Firewall](https://docs.replit.com/features/security/package-firewall) says supported installs pass through an enabled-by-default supply-chain check that blocks packages flagged as malicious or compromised.

#### Implication

Skill Issue should ship its CLI through a normal package ecosystem and expose a stable executable such as `skill-issue`. The Replit repository/template should pin the CLI dependency and any runtime/toolchain versions, configure only necessary run or workflow commands, and provide a non-destructive `skill-issue doctor` or equivalent to verify paths and prerequisites. A bootstrap script should be idempotent because Agent or users may rerun it.

### Finding 8: MCP is a separate account-level capability and cannot be silently bundled into project files

Replit Agent supports curated and custom remote MCP servers. Custom servers require an HTTPS endpoint, may use OAuth or static headers, and can be shared through a one-click install link. Once connected, tools become available across the user's projects. Installation still requires authorization; Replit scans MCP traffic and planned executions, and some tool actions require confirmation.

#### Evidence

[Connect via MCP](https://docs.replit.com/build/connect-via-mcp) documents pre-listed sign-in, custom HTTPS URLs, headers, OAuth, tool discovery across projects, confirmation prompts, and shareable install links. [MCP Install Links](https://docs.replit.com/features/mcp/install-links) defines a `https://replit.com/integrations?mcp={payload}` link with display name, base URL, and optional headers. [MCP list](https://docs.replit.com/features/mcp/overview) says users are prompted to authorize, all MCP traffic passes through Replit's security scanner, suspicious tools can be blocked, and authentication uses OAuth DCR or headers.

#### Implication

If Skill Issue has service-backed tools, it should host a remote HTTPS MCP server and publish a Replit install badge as a second onboarding action. The project skill can explain when to use it, but the project import/bootstrap must not embed live header values or claim it connected MCP automatically. A local stdio-only MCP bundle is unsupported by the inspected Replit Agent documentation.

### Finding 9: Connectors and app integrations are adjacent capabilities, not package contents

Replit distinguishes built-in managed integrations, first-party connectors, external APIs configured with keys, and paid Agent services. Connectors are bound to accounts/workspaces or organizations and persist across apps; they require sign-in and may be governed by admins, groups, scopes, consent, revocation, and audit. Their configuration is outside a project repository.

#### Evidence

[Agent Integrations](https://docs.replit.com/features/integrations/overview) defines four integration categories, states that connectors require sign-in, and says account connections persist across apps. [Managing Your Connectors](https://docs.replit.com/replitai/managing-connectors) documents workspace/organization administration, per-app consent and revocation, monitoring, group access, least-privilege scopes, credential rotation, and plan restrictions. [Import from a provider](https://docs.replit.com/build/import-from-providers) explicitly says ZIP import does not import existing third-party connector configuration.

#### Implication

Skill Issue's installer should detect required integrations and print direct setup instructions, but treat connector authorization as an explicit platform step. The bundle may include adapter code and documentation, while credentials and connector state remain external. Required versus optional integrations should be declared so the CLI can degrade gracefully when a connector is unavailable.

### Finding 10: Secrets, permissions, and trust prevent a truly zero-confirmation bundle

Secret values should not be distributed in public files or imported projects. Replit provides encrypted app- and account-level secrets, but imports and public remixes do not transfer values to unrelated users. Sensitive Agent actions remain human-in-the-loop. Curated picker skills are audited; external skills are not automatically trusted.

#### Evidence

[Secrets](https://docs.replit.com/core-concepts/project-editor/app-setup/secrets) documents encrypted app/account secrets, environment-variable delivery, account-secret linking, and remix visibility rules. [Import from a provider](https://docs.replit.com/build/import-from-providers) states that Git and ZIP imports do not bring existing secret/environment-variable values. [Replit shared responsibility model](https://docs.replit.com/references/security/shared-responsibility-model) assigns users responsibility for deliberately approving publishing, secret changes, and outbound calls. [Skills directory](https://docs.replit.com/features/agent/skills-directory) states picker skills are audited while external sources must be reviewed. Enterprise [Custom Templates](https://docs.replit.com/teams/custom-templates) can copy template secrets at fork time, but they are snapshots and do not rotate into existing apps.

#### Implication

Skill Issue should ship a secret-name manifest and setup wizard, never secret values. Installation should separate file writes, package execution, connector/MCP authorization, and secret entry into reviewable steps. Public distribution must assume every external skill and script will be inspected. Enterprise secret-copy behavior should be optional and documented as fork-time copying with manual rotation thereafter.

### Finding 11: Updates split between skill updates, workspace updates, and project-copy updates

There is no single native update channel for a mixed Skill Issue bundle. The skills CLI can update installed skills; centrally managed workspace changes affect future chats; enterprise template changes affect only newly forked apps; imported or remixed projects and deployed snapshots are independent copies that require explicit updates and republishing.

#### Evidence

The linked [`vercel-labs/skills`](https://github.com/vercel-labs/skills) CLI documents `skills update` for project or global skills. [Agent Customization](https://docs.replit.com/features/agent/agent-customization) says workspace customization changes apply to future chats. [Custom Templates](https://docs.replit.com/teams/custom-templates) says existing forks do not receive template updates. [Publishing](https://docs.replit.com/learn/projects-and-artifacts/replit-deployments) describes publishing as a snapshot and requires publishing again to update the live app.

#### Implication

Skill Issue needs its own bundle manifest and upgrade command covering the CLI, skills, guidance snippets, configuration migrations, references, scripts, and assets. It should preserve user-modified project files, report conflicts, and leave secrets/connectors untouched. The native `skills update` path is useful for skill-only installs but cannot be the authoritative updater for the whole experience.

### Finding 12: The recommended Replit packaging is a two-layer bundle with explicit platform handoffs

The best-supported design is: (1) a Replit-native Agent Skill for discovery, invocation, workflow guidance, scripts, references, and assets; and (2) a Replit-ready project distribution plus standalone CLI for complete installation, upgrades, configuration, and validation. Optional MCP/connectors, secrets, and publishing remain explicit Replit platform steps. This produces one coherent guided experience, although it cannot be one atomic installation across project, workspace, and account scopes.

#### Evidence

The capability boundary is consistent across [Agent Skills](https://docs.replit.com/features/agent/skills), [Custom Templates](https://docs.replit.com/teams/custom-templates), [Import from a provider](https://docs.replit.com/build/import-from-providers), [Replit App Configuration](https://docs.replit.com/features/project-setup/configuration), [Connect via MCP](https://docs.replit.com/build/connect-via-mcp), [Managing Your Connectors](https://docs.replit.com/replitai/managing-connectors), and [Secrets](https://docs.replit.com/core-concepts/project-editor/app-setup/secrets): filesystem content can travel with projects and skills, while authorizations and secret values live in separate governed scopes.

#### Implication

Skill Issue should implement the following Replit deliverables:

1. A standards-compliant `.agents/skills/skill-issue/` package with `SKILL.md`, narrowly scoped descriptions, `scripts/`, `references/`, `assets/`, and compatibility notes.
2. A canonical Git repository/public remix containing the complete filesystem bundle, pinned CLI dependency, `.replit`, optional `replit.nix`, minimal `replit.md`, and an idempotent bootstrap/doctor workflow.
3. An `npx skills ... -a replit` installation path for users who want only the Agent capability in an existing project.
4. A standalone `skill-issue install|update|doctor` flow for full-bundle installation and safe upgrades in existing projects, with merge/conflict handling and a recorded bundle version.
5. Separate MCP install badges, connector setup links, and a secret-name checklist, all presented as explicit authorized steps.
6. An optional Enterprise custom-template layout and workspace-skill ZIP, acknowledging their admin, plan, and update boundaries.

## Notes

- **Caveat — extensions/plugins:** No current public Replit extension/plugin authoring surface was found in the official documentation index on 2026-07-19. Historical Replit Extensions may have existed, but treating them as a current supported distribution target is unsupported by the inspected sources.
- **Caveat — public catalog submission:** Replit documents Replit, partner, and community skills and links community discovery to `skills.sh`, but the inspected official pages do not document an open submission/review process for becoming a pre-defined Replit or partner skill. Distribution through a public repository and `skills.sh` is supported; guaranteed placement in Replit's curated picker is unsupported.
- **Caveat — one-click completeness:** A Git import, ZIP import, remix, or enterprise template can copy the full filesystem bundle. Secret values, connector state, MCP authorization, and existing-project merge decisions still require separate steps.
- **Caveat — Open in Replit:** [Open in Replit](https://docs.replit.com/features/integrations/open-in-replit) carries a compressed prompt and selected build mode, not an arbitrary file bundle. It is useful as a guided acquisition link but weaker and less deterministic than Git import or remix for Skill Issue packaging.
- **Useful search terms:** `Replit Agent Skills`, `.agents/skills`, `Agent Customization`, `custom_instruction/instructions.md`, `Replit MCP install link`, `Replit Gallery remix`, `replit.nix`, `.replit`, `npx skills -a replit`.

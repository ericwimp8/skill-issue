# GitHub Copilot Direct Installation

## Assignment

**Goal:** Establish the supported direct, filesystem and CLI-managed lifecycle for installing a Skill Issue agent skill into GitHub Copilot, including supported surfaces, scopes, discovery, safety, maintenance, and the boundary between skills and Copilot plugins.

**Scope:** Current first-party GitHub Docs and GitHub CLI manual material for GitHub Copilot agent skills, Copilot CLI, Copilot app, supported surfaces, and plugins. “Direct installation” means copying or CLI-installing an Agent Skills directory rather than publishing or installing a Copilot plugin package.

**Exclusions:** Local Skill Issue source, unverified IDE-extension implementation details, third-party installers, and a native Copilot plugin package design.

## Sources

- [Adding agent skills for GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills) — first-party CLI lifecycle, layouts, reload, inspection, invocation, and removal.
- [Adding agent skills for GitHub Copilot](https://docs.github.com/en/copilot/how-tos/copilot-on-github/customize-copilot/customize-cloud-agent/add-skills) — first-party cross-surface skill authoring, GitHub CLI management, provenance, pinning, and code-review behavior.
- [GitHub Copilot CLI command reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference) — first-party CLI skill schema, full discovery precedence, collision behavior, and custom locations.
- [About agent skills](https://docs.github.com/en/copilot/concepts/agents/about-agent-skills) and [Copilot customization cheat sheet](https://docs.github.com/en/copilot/reference/customization-cheat-sheet) — first-party supported-surface and standard-location matrix.
- [GitHub CLI `gh skill install`](https://cli.github.com/manual/gh_skill_install), [`gh skill list`](https://cli.github.com/manual/gh_skill_list), and [`gh skill update`](https://cli.github.com/manual/gh_skill_update) — first-party managed installation, verification, provenance, and update behavior. `gh skill` is explicitly preview and GitHub Docs require GitHub CLI 2.90.0 or later.
- [Configuring GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/set-up-copilot-cli/configure-copilot-cli) — first-party directory trust and tool-approval constraints.
- [About GitHub Copilot plugins](https://docs.github.com/en/copilot/concepts/agents/about-plugins) and [CLI plugin reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-plugin-reference) — first-party contrast between optional packages and manually configured skill directories.

## Findings

### Supported direct-installation contract

GitHub Copilot has a native Agent Skills filesystem contract; a skill is a named directory containing `SKILL.md`, and GitHub officially documents downloading a skill directory and moving it into a supported project or personal location. A plugin is therefore optional for a skill-only Skill Issue distribution, rather than a required wrapper. The same skill construct works with Copilot cloud agent, code review, Copilot CLI, Copilot app, and agent mode in VS Code and JetBrains. The support table also marks Agent Skills available in Visual Studio and JetBrains as preview, and unavailable in Eclipse and Xcode.

**Evidence:** GitHub documents manual download and move into supported locations in [Adding agent skills for Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#adding-a-skill-that-someone-else-has-created) and identifies skills as folders of instructions, scripts, and resources in [About agent skills](https://docs.github.com/en/copilot/concepts/agents/about-agent-skills). The cross-surface statement appears in [Adding agent skills for GitHub Copilot](https://docs.github.com/en/copilot/how-tos/copilot-on-github/customize-copilot/customize-cloud-agent/add-skills#creating-and-adding-a-skill); the support matrix is [Copilot customization cheat sheet](https://docs.github.com/en/copilot/reference/customization-cheat-sheet#ide-and-surface-support).

**Implication:** Treat GitHub Copilot as a first-class common Agent Skills filesystem adapter. Ship a skill directory, not a plugin manifest, unless the product also needs package-only components such as hooks, MCP servers, custom agents, or LSP configuration.

### Exact scopes and placement paths

The supported portable project roots are `.github/skills/<skill-name>/SKILL.md`, `.agents/skills/<skill-name>/SKILL.md`, and `.claude/skills/<skill-name>/SKILL.md`; personal roots are `~/.copilot/skills/<skill-name>/SKILL.md` and `~/.agents/skills/<skill-name>/SKILL.md`. A direct installer should prefer `.github/skills` for a repository intended to work across Copilot cloud, code review, IDE agent mode, and CLI; use `~/.copilot/skills` for one user’s Copilot-only local installation. GitHub documents no system-wide direct skill path. Organization and enterprise skills are remote-managed rather than filesystem-installed; they are outside this direct adapter.

For Copilot CLI, discovery precedence is project `.github/skills`, project `.agents/skills`, project `.claude/skills`, parent `.github/skills` in monorepos, personal `~/.copilot/skills`, personal `~/.agents/skills`, plugins, `COPILOT_SKILLS_DIRS`, bundled skills, then projected organization/enterprise skills. The CLI-specific `COPILOT_SKILLS_DIRS` accepts comma-separated extra locations, but is a local runtime configuration rather than a portable repository layout.

**Evidence:** The path set and scope are prescribed in [Adding agent skills for Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#creating-and-adding-a-skill) and [About agent skills](https://docs.github.com/en/copilot/concepts/agents/about-agent-skills). The CLI order, parent `.github/skills` inheritance, custom-directory environment variable, and remote skill tier are specified in [Copilot CLI command reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference#skill-locations).

**Implication:** The adapter needs explicit `project` and `user` destinations only. It should reject or clearly classify `system`, organization, and enterprise filesystem scope as unsupported, and should offer custom directories only as a Copilot CLI-specific advanced path.

### Required layout and content preservation

Each installed skill is a lowercase, hyphen-separated directory with a file named exactly `SKILL.md`. `SKILL.md` requires YAML frontmatter with a lowercase hyphenated `name` and a `description`; the CLI reference further constrains `name` to letters, numbers, and hyphens and a maximum of 64 characters. `license` is optional. Copilot discovers every file in the skill directory when the skill is invoked, so scripts, references, assets, examples, templates, and nested resource directories must be copied as part of the skill directory and referenced from `SKILL.md`; no skill manifest is required.

CLI-only frontmatter can include `argument-hint`, `allowed-tools`, `user-invocable`, and `disable-model-invocation`. The portable core should leave those out unless their behavior is deliberately desired across the target Copilot surfaces.

**Evidence:** The required filename, frontmatter, lowercase naming, optional resources, and available-files behavior are documented in [Adding agent skills for Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#creating-and-adding-a-skill) and [its script section](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#enabling-a-skill-to-run-a-script). Field limits and CLI-only controls are in the [CLI skill reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference#skill-frontmatter-fields).

**Implication:** A direct installer must copy the entire selected skill directory atomically enough to avoid a visible `SKILL.md` whose referenced resources are missing. It should validate the canonical name and frontmatter before placement; it should not generate `plugin.json`, merge a plugin configuration, or flatten resources.

### Installation and discovery lifecycle

Two official routes exist. Manual direct installation is: download the complete directory, place it in one of the supported roots, then start a new CLI session or run `/skills reload` in the active session. Verify it with `/skills info SKILL-NAME` (including location), `/skills list`, and an explicit `/SKILL-NAME` prompt. The CLI also provides `copilot skill list` and `copilot skill add <FILE | URL | DIRECTORY>` for shell scripting.

Managed direct installation uses `gh skill install OWNER/REPO SKILL --agent github-copilot --scope project|user`; its defaults are `github-copilot` and `project`. The GitHub CLI installer copies rather than symlinks local source directories, uses host-specific locations, and for Copilot project scope shares `.agents/skills` with several agent hosts. It can also install from a local directory with `--from-local`; this is a supported CLI operation but is lower fit for an internet-distributed product release. Verify a managed install with `gh skill list --agent github-copilot --scope project|user --json skillName,sourceURL,scope,version,pinned,path`, then use the in-session Copilot verification commands above.

There is no first-party evidence that IDE, Copilot app, cloud agent, or code review require a specific restart, window reload, or trust prompt solely to discover a repository skill. For those surfaces, repository synchronization/indexing behavior beyond the documented files is **unsupported** for this assignment; verify with an explicit task requiring the named skill after the repository change is available.

**Evidence:** Manual placement, reload, `info`, list, add, and removal are documented in [Adding agent skills for Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#using-agent-skills). `gh skill install` scope, Copilot host selection, copying, discovery convention, and defaults are documented in [GitHub CLI `gh skill install`](https://cli.github.com/manual/gh_skill_install). `gh skill list` verification fields and scope filtering are documented in [GitHub CLI `gh skill list`](https://cli.github.com/manual/gh_skill_list).

**Implication:** A CLI-managed implementation needs a post-copy reload plus location/name inspection step. Its success criterion should be Copilot’s live skill listing and a targeted invocation, rather than filesystem existence alone. For repository installations, commit/push the skill directory before validating cloud agent or code review discovery.

### Activation and collisions

Copilot chooses skills automatically by prompt relevance and their `description`; users can force CLI use with `/skill-name`. The CLI lets a user enable or disable individual skills through `/skills`. Duplicate local names resolve first-found-wins in the documented priority order; a higher-priority project or personal skill silently wins over a plugin skill with the same `name`. Plugin-to-plugin collisions may additionally be invoked through plugin-qualified names, but that exception does not change direct filesystem skill precedence.

**Evidence:** Relevance selection and explicit slash invocation are documented in [Adding agent skills for Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#using-agent-skills). Name-based priority and duplicate handling are in [Copilot CLI command reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference#skill-locations), with plugin collision detail in [the same reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-command-reference#skills-reference). The project/personal-over-plugin effect is corroborated by the [CLI plugin reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-plugin-reference#loading-order-and-precedence).

**Implication:** The installer must detect an existing canonical `name` in its selected destination and refuse by default rather than silently replace it. It must not solve a collision by changing an installed directory name only; directory and frontmatter name should remain aligned. A deliberate replacement requires a user-selected overwrite/backup policy and post-install `info` confirmation.

### Update, repair, rollback, and removal

`gh skill install` injects source repository, ref, and tree-SHA provenance into `SKILL.md`; `gh skill update` compares that local tree SHA to the remote source. It supports dry-run, `--all`, pins, unpinning, and `--force`. `--force` restores source-managed files but leaves locally added extra files. A manually copied skill has no GitHub provenance: interactive `gh skill update` may ask for a source repository, while noninteractive or `--all` skips it. Direct manual maintenance is therefore a deliberate full-directory replacement from a reviewed source followed by reload and `info`; preserving an arbitrary local edit during replacement is **unsupported** by the official docs.

For removal, use `/skills remove SKILL-DIRECTORY` or the `copilot skill` equivalent only for directly added skills; plugin-provided skills must be removed by managing the plugin. `gh skill` currently exposes install, list, preview, publish, search, and update, but no uninstall subcommand. GitHub publishes no dedicated partial-install recovery or transactional rollback protocol for a manually copied skill. The evidence-backed repair path is re-install/re-copy the complete directory (or `gh skill update --force` for a provenance-managed install), reload, and verify; preserve user changes separately before an overwrite.

**Evidence:** Provenance, update, pins, and the GitHub CLI 2.90.0 preview gate are documented in [Adding agent skills for GitHub Copilot](https://docs.github.com/en/copilot/how-tos/copilot-on-github/customize-copilot/customize-cloud-agent/add-skills#managing-skills-with-github-cli). Exact update behavior, skipped manual skills, `--force`, and dry run are in [GitHub CLI `gh skill update`](https://cli.github.com/manual/gh_skill_update). Direct removal rules are in [Adding agent skills for Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#using-agent-skills), and the available `gh skill` commands are listed in the [GitHub CLI manual](https://cli.github.com/manual/gh_skill).

**Implication:** Prefer `gh skill` for published-source installs when its preview status is acceptable, because it supplies provenance, pinning, and repair. A common filesystem adapter must otherwise own backup, atomic staging/replace, restore-on-failure, and an installation receipt itself; those are adapter safety measures, not native Copilot lifecycle guarantees.

### Trust, permissions, and policy gates

Copilot CLI asks the user to trust the launch directory and its descendants before a session, and trust controls where it can read, modify, and execute files. Tool requests can require approval. A skill that sets `allowed-tools: shell` or `bash` pre-approves command execution, so GitHub warns to use it only after reviewing trusted scripts. GitHub also warns that skills discovered through `gh skill` are unverified and may carry prompt injection, hidden instructions, or malicious scripts; preview before installation. Copilot CLI availability for organization-provided users depends on the organization enabling the CLI policy.

**Evidence:** CLI trust model and approval controls are in [Configuring GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/set-up-copilot-cli/configure-copilot-cli#setting-trusted-directories). Shell-tool warning is in [Adding agent skills for Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/add-skills#enabling-a-skill-to-run-a-script). The `gh skill preview` warning, and the public-preview/version gate, are in [Adding agent skills for GitHub Copilot](https://docs.github.com/en/copilot/how-tos/copilot-on-github/customize-copilot/customize-cloud-agent/add-skills#managing-skills-with-github-cli). Organization policy gate is in [About agent skills](https://docs.github.com/en/copilot/concepts/agents/about-agent-skills).

**Implication:** The direct installer must inspect the incoming file tree before any overwrite, avoid `allowed-tools` for unreviewed skills, and report Copilot CLI organization-policy denial as a real blocker rather than attempting to bypass it. It should keep source authentication and repository permissions separate from skill placement.

### Plugin interpretation rejected for Skill Issue’s skill-only case

A Copilot plugin has a required `plugin.json` and is the package mechanism for a bundle of agents, skills, hooks, MCP configuration, and LSP configuration. It provides marketplace discovery, team sharing, and versioning, but GitHub explicitly states that functionality could also be added through manual repository configuration. A plugin’s `enabledPlugins` settings are installation configuration, not a requirement for manually placed skill directories.

**Evidence:** Plugin structure, required manifest, installation mechanisms, and the manual-configuration comparison are documented in [About GitHub Copilot plugins](https://docs.github.com/en/copilot/concepts/agents/about-plugins#how-plugins-are-structured) and [its manual configuration comparison](https://docs.github.com/en/copilot/concepts/agents/about-plugins#plugins-compared-with-manual-configuration). Required plugin fields and component paths are in the [CLI plugin reference](https://docs.github.com/en/copilot/reference/copilot-cli-reference/cli-plugin-reference#pluginjson).

**Implication:** Reject “native plugin package required” as a lower-fit interpretation. Use a plugin only if Skill Issue intentionally distributes a multi-component Copilot bundle or requires marketplace lifecycle. For a direct skill install, adding `plugin.json` and merging `enabledPlugins` would add an unrelated lifecycle and violate the narrow adapter boundary.

## Notes

- **Currentness caveat:** Sources were inspected on 2026-07-19. `gh skill` is public preview and GitHub says it is subject to change; pin the observed CLI contract in acceptance tests rather than assuming long-term stability.
- **Lower-fit path:** `.claude/skills` and `.agents/skills` are valid direct locations, but `.github/skills` is the clearest GitHub-native repository convention for broad Copilot surfaces. `~/.agents/skills` is useful only when intentional cross-agent sharing is wanted.
- **Unsupported:** No official source found for a system-wide skill directory, an IDE-wide filesystem location distinct from the documented personal roots, a required IDE/app restart, or an automatic transactional rollback for manual copy failure.
- **Useful verification terms:** `copilot skill list`, `/skills reload`, `/skills info`, `gh skill list --agent github-copilot`, `gh skill update --dry-run`.

# OpenAI Codex Direct Installation

## Assignment

**Goal:** Establish the current direct, filesystem-managed lifecycle for installing a Skill Issue skill into OpenAI Codex without a native plugin package.

**Scope:** Official OpenAI documentation, the first-party `openai/codex` repository, and the official npm registry as inspected on 2026-07-19. Covers local Codex CLI, IDE extension, and ChatGPT desktop app distinctions; repository, user, administrator, and bundled scopes; discovery, activation, updates, recovery, and security boundaries.

**Exclusions:** Plugin-package implementation and marketplace distribution except where needed to distinguish their lifecycle; non-OpenAI harnesses; local workspace files as product evidence; claims about unsupported cloud skill discovery.

## Sources

- [OpenAI Build skills documentation](https://developers.openai.com/codex/skills) — current public product contract for skill structure, locations, activation, refresh, configuration, and supported local surfaces.
- [OpenAI Skill controls documentation](https://learn.chatgpt.com/docs/enterprise/skills) — separates workspace skills, local filesystem skills, and plugins with their respective control boundaries.
- [OpenAI Agent approvals and security documentation](https://learn.chatgpt.com/docs/agent-approvals-security) — current sandbox, approval, network, local-runtime, and cloud-runtime boundaries.
- [OpenAI Codex app-server README](https://github.com/openai/codex/blob/main/codex-rs/app-server/README.md#skills) — first-party protocol for skill listing, reload signals, ephemeral additional roots, and enable/disable writes.
- [Codex current skill loader](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader.rs), [root merger](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/root_loader.rs), [discovery implementation](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader/discovery.rs), and [skill configuration rules](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/config_rules.rs) — current first-party implementation cross-check for scope roots, scanning, parse limits, collision behavior, and configuration semantics. These are implementation evidence rather than a backward-compatibility promise.
- [Latest official CLI package metadata](https://registry.npmjs.org/@openai%2Fcodex/latest) and [release `rust-v0.144.6`](https://github.com/openai/codex/releases/tag/rust-v0.144.6) — current latest CLI release observed: `0.144.6`, published 2026-07-18.

## Findings

### Direct folders are a supported local installation path

Codex officially supports a skill as a local directory containing `SKILL.md`; no plugin manifest, marketplace entry, or installer is required. Direct folders are the recommended fit for local authoring and repo-scoped workflows. A plugin becomes the higher-fit distribution adapter when the workflow must be shared beyond one repository, grouped with other skills, or bundled with connectors, MCP configuration, hooks, or presentation metadata.

**Evidence:** The [skills guide](https://developers.openai.com/codex/skills) explicitly permits manual creation of a folder with `SKILL.md`, labels direct folders best for local authoring/repo scope, and describes plugins as the distribution path. The [plugin guide](https://developers.openai.com/codex/plugins/build) says to start with a local skill for one repository or one personal workflow.

**Implication:** Skill Issue can use a direct installer for the Codex local-skill adapter. It should copy or atomically replace a skill directory at a supported local root, rather than inventing a plugin package for the common single-skill case.

### Supported roots and scope choice

The documented local roots are:

| Scope | Root | Direct-install use |
| --- | --- | --- |
| Repository | `$CWD/.agents/skills` through every ancestor up to `$REPO_ROOT/.agents/skills` | Commit a skill with the work it governs; use the repo root for whole-repo availability or a nested ancestor for module scope. |
| User | `$HOME/.agents/skills` | Install a personal skill available across repositories on that local machine. |
| Administrator | `/etc/codex/skills` | Deploy a machine/container-wide skill using host administration. |
| System | Bundled by OpenAI | Read-only product-provided skills; not an installer target. |

The current loader additionally retains `$CODEX_HOME/skills` as a deprecated user-scope compatibility root and derives project-config roots from config layers, but the public install contract names `.agents/skills` for repository and user installs. The loader scans repository `.agents/skills` directories from the project root through the launch CWD, follows directory symlinks for repository/user/admin roots, skips hidden directories, and currently limits a root scan to six levels, 2,000 directories, and 20,000 entries.

**Evidence:** [Build skills](https://developers.openai.com/codex/skills) publishes the root table, ancestor scan, and symlink support. The [loader](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader.rs) constructs the documented roots, comments that `$CODEX_HOME/skills` is deprecated compatibility, and selects symlink policy. [Discovery](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader/discovery.rs) supplies the current traversal limits and hidden-directory handling.

**Implication:** Default a Skill Issue direct install to either `<repo>/.agents/skills/<skill-name>/` (team/repository scope) or `~/.agents/skills/<skill-name>/` (personal scope). Treat `$CODEX_HOME/skills` as a migration/recovery probe only, never as the preferred new target. Do not install below a hidden directory or rely on a deeply nested package layout.

### Minimum payload, metadata, and naming

The portable payload is one folder named after the skill containing an exact-case `SKILL.md` file. `SKILL.md` has YAML frontmatter with `name` and nonempty `description`, followed by Markdown instructions. Optional `scripts/`, `references/`, and `assets/` support progressive disclosure; `agents/openai.yaml` can add UI metadata, a default prompt, tool dependencies, and `allow_implicit_invocation: false`.

For reliable direct installation, use a short, unique kebab-case `name` no longer than 64 characters and make the directory name identical. The current loader does not require kebab-case and falls back to the folder name if `name` is absent, but this is implementation tolerance rather than the documented authoring contract. It rejects a missing description and enforces the 64-character base-name limit; a qualified name has a 128-character limit.

**Evidence:** [Build skills](https://developers.openai.com/codex/skills) defines the required `SKILL.md` fields and optional layout, then documents the optional metadata schema and invocation policy. The [loader parser](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader.rs) validates nonempty description and the current limits. The first-party [Skill Creator sample](https://github.com/openai/codex/blob/main/codex-rs/skills/src/assets/samples/skill-creator/SKILL.md) recommends lowercase letters, digits, hyphens, a name below 64 characters, and folder/name alignment.

**Implication:** The adapter should validate the delivered payload before copying: exact `SKILL.md`; YAML delimiter and both documented fields; name/folder equality; unique name in its target scope; and no unwanted nested `SKILL.md` files. Keep support files relative to that directory and avoid an `agents/openai.yaml` unless its UI/policy/dependency behavior is intentionally required.

### Discovery, activation, and practical verification

Codex has a two-stage discovery model: it first exposes name, description, and path; it reads the full `SKILL.md` only after choosing a skill. A user can explicitly invoke a skill with `$skill-name`; CLI and IDE expose the skill list via `/skills` or the `$` mention picker. Implicit invocation matches the `description` unless optional metadata disables it.

After a filesystem copy/update, Codex is documented to detect changes automatically. Restart the relevant local Codex session if the change does not appear. The app-server protocol supplies stronger programmatic verification: call `skills/list` for the desired CWD, use `forceReload: true` to bypass a cached list, and respond to `skills/changed` by listing again. A successful record must show the expected name, description, absolute `SKILL.md` path, and `enabled: true`; then explicitly invoke it once using the resolved selector.

**Evidence:** [Build skills](https://developers.openai.com/codex/skills) documents progressive disclosure, explicit and implicit activation, `/skills`, `$`, automatic detection, and restart fallback. The [app-server README](https://github.com/openai/codex/blob/main/codex-rs/app-server/README.md#skills) documents `skills/list`, `forceReload`, `skills/changed`, and structured `skill` input by path.

**Implication:** A direct installer should finish with a surface-appropriate verification instruction: restart the CLI/IDE/desktop client only when its picker has not refreshed; inspect the discovered path; and make one explicit invocation. A file-copy success alone is insufficient evidence of activation.

### Surface boundaries: CLI, IDE, desktop, web, and cloud

The official skills guide names the ChatGPT desktop app, Codex CLI, and IDE extension as available skill surfaces. It specifically documents `/skills` and `$` in CLI/IDE and a Skills sidebar/picker in the desktop app. Enterprise guidance calls local filesystem skills a distinct lifecycle used by covered local capabilities in desktop, CLI, and IDE; it separates them from ChatGPT workspace skills and plugins. Plugin availability is broader in Work mode on the web and desktop/CLI, while plugins are unavailable in the IDE extension.

Codex cloud runs a checked-out repository in an isolated container. The cloud-environment documentation explicitly describes repository checkout and `AGENTS.md`, but does not publish a direct user/admin filesystem-skill installation procedure or a cloud skill-picker/discovery contract. A repository may contain `.agents/skills`, yet its cloud discovery should be treated as **caveated/unsupported for this adapter** until OpenAI documents and the target cloud account demonstrates it.

**Evidence:** [Build skills](https://developers.openai.com/codex/skills) lists desktop, CLI, and IDE availability. [Skill controls](https://learn.chatgpt.com/docs/enterprise/skills) distinguishes local filesystem skills from workspace skills/plugins and lists the plugin surface difference. [Cloud environments](https://learn.chatgpt.com/docs/environments/cloud-environment) documents checkout, setup, and cloud isolation without a direct-skill lifecycle.

**Implication:** Ship the direct adapter as a local Codex CLI/IDE/desktop adapter. Do not advertise it as a ChatGPT-web workspace installer or a Codex-cloud installer. Use a plugin or workspace-skill path when those separate control surfaces are required.

### Collisions and selector safety

Same-name skills are not merged: the public contract says they can both appear in selectors. Current code deduplicates only identical canonical `SKILL.md` paths, orders repository entries before user/system/admin by scope, and requires an unqualified text `$name` mention to be unambiguous. Thus scope order must not be interpreted as safe same-name override precedence.

**Evidence:** [Build skills](https://developers.openai.com/codex/skills) explicitly says same-name skills can both appear. The [root merger](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/root_loader.rs) deduplicates by path and sorts scopes; [injection](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/injection.rs) resolves plain names only when unambiguous.

**Implication:** Before installation, enumerate the target root and known repository/user roots for the frontmatter name. Refuse or require an explicit replacement decision on collision; do not rely on invisible precedence. For recovery, select by exact path through a structured client/app-server selector when available.

### Configuration, enabling, and safe merging

`~/.codex/config.toml` can disable a discovered skill without deleting it using an append-only `[[skills.config]]` entry with exactly one selector: the absolute `SKILL.md` `path` or a `name`, plus `enabled`. The documented example is path-based. Restart Codex after configuration changes. Current source applies skill config from user/session layers, rejects entries containing both selector kinds, and applies later matching rules over earlier ones. The configuration disables an already discovered skill; it does not configure a new directory root.

**Evidence:** [Build skills](https://developers.openai.com/codex/skills) documents `[[skills.config]]` and restart. [Configuration rules](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/config_rules.rs) and [configuration type](https://github.com/openai/codex/blob/main/codex-rs/config/src/skills_config.rs) show selector validity, layer handling, and enable/disable resolution.

**Implication:** Preserve the existing TOML file and its unrelated tables. For a surgical disable/restore, add or modify one path-based `[[skills.config]]` record for the final canonical target; never overwrite the whole file or use a name selector where duplicate names are possible. No config change is needed for a normal direct installation.

### Update, repair, rollback, and uninstall lifecycle

The direct lifecycle is file ownership rather than a package-manager transaction:

1. Build and validate the replacement in a staging directory with the required `SKILL.md` contract.
2. Preserve a known-good copy of the existing target directory, then replace the entire target directory as one unit; this removes stale resources that a file-by-file copy could leave behind.
3. Refresh/restart and verify the discovered path plus an explicit invocation.
4. If verification fails, restore the preserved directory, refresh/restart, and verify again.
5. To uninstall, remove the target skill directory and any path-based disable entry that only applied to it, then refresh/restart and confirm it is absent from the relevant list.

For partial recovery, first confirm the active CWD, repository root, and intended scope; then inspect the exact `SKILL.md` spelling/frontmatter; then use `skills/list` with `forceReload` (or the local selector) to distinguish a discovery/parsing failure from a disabled or colliding entry. The current loader reports parse errors for non-system skills and finds `SKILL.md` by exact filename, while optional metadata fails open rather than blocking a valid `SKILL.md`.

**Evidence:** The [skills guide](https://developers.openai.com/codex/skills) supplies change detection, restart, and disable behavior. The [discovery source](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader/discovery.rs) matches `SKILL.md` exactly; the [loader](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader.rs) emits parse errors for non-system scopes and treats optional metadata independently. The [app-server README](https://github.com/openai/codex/blob/main/codex-rs/app-server/README.md#skills) supplies re-list/reload inspection.

**Implication:** A common installer can provide transactional directory replacement, backup/restore, remove, and verification for Codex. It should leave an interrupted staging directory harmless, avoid deleting a target until the staged payload passes static validation, and record the exact installed root/path for recovery.

### Permissions, trust, and policy boundaries

Installing a local skill only places instructions/resources on disk; it does not grant command, filesystem, network, MCP, connector, or workspace permissions. In local CLI/IDE, OS sandboxing and approval policy still govern model-generated actions; by default local network access is off and writes are limited to the active workspace. A target such as `~/.agents/skills` or `/etc/codex/skills` may therefore require a user/admin-managed copy operation rather than an agent running in workspace-write mode. The optional skill metadata can restrict implicit invocation, but it is not an approval/trust grant.

No current public documentation was found for a per-directory "trust this direct skill" approval prompt. The documented control boundary for filesystem skills is filesystem distribution, local client configuration, and runtime permissions; workspace-skill ownership and plugin installation state do not transfer when files are copied.

**Evidence:** [Agent approvals and security](https://learn.chatgpt.com/docs/agent-approvals-security) defines local sandbox/approval/network behavior and out-of-workspace approval. [Skill controls](https://learn.chatgpt.com/docs/enterprise/skills) defines the separate filesystem, workspace, and plugin control boundaries. [Build skills](https://developers.openai.com/codex/skills) documents `allow_implicit_invocation` and optional tool dependencies.

**Implication:** Treat a supplied skill as executable operational content: review its instructions and scripts before deployment, preserve operating-system ownership/permissions, and let the active Codex sandbox/approval policy control execution. The installer must not promise trust approval, connector authorization, or bypass of filesystem restrictions.

### Version and adapter-fit gate

The public documentation currently describes the `.agents/skills` root, local filesystem lifecycle, automatic detection, and `/skills`/`$` discovery. The latest official CLI release observed is `0.144.6` (2026-07-18). The current open-source implementation includes the documented roots and app-server refresh methods, but no source review proves the feature set on every older CLI, desktop, or IDE build.

**Evidence:** [Build skills](https://developers.openai.com/codex/skills) is the product-facing contract; the [official latest package metadata](https://registry.npmjs.org/@openai%2Fcodex/latest) and [release](https://github.com/openai/codex/releases/tag/rust-v0.144.6) identify the observed current CLI. The current [loader](https://github.com/openai/codex/blob/main/codex-rs/core-skills/src/loader.rs) and [app-server README](https://github.com/openai/codex/blob/main/codex-rs/app-server/README.md#skills) corroborate implementation behavior.

**Implication:** Gate direct installation on an installed client that exposes skills in its local selector or app-server list; use current `.agents/skills` targets. The common adapter covers simple local folder delivery, validation, backup/restore, collision checking, and selector verification. A bespoke adapter is required only for workspace ownership/sharing, ChatGPT web Work mode, cloud-specific discovery, connector/MCP authorization, plugin marketplace lifecycle, or enterprise policy deployment.

## Notes

- **Rejected interpretation — plugins are required for skills:** The official skills guide explicitly supports manual local folders. Plugins are a distribution/control adapter, not a prerequisite for direct local installation.
- **Rejected interpretation — a repo skill safely overrides a user skill of the same name:** Both can appear; current implementation path-deduplication and ambiguous plain-name handling make this unsafe.
- **Caveat — `$CODEX_HOME/skills`:** Current source retains it for backward compatibility, while current public guidance standardizes `.agents/skills`. New installs should not choose the deprecated root.
- **Caveat — desktop/IDE refresh:** The product contract says automatic detection then restart fallback. The app-server force-reload API is a stronger verification path where that client exposes it; the public docs do not promise it as a user-facing CLI command.
- **Unsupported — Codex cloud direct user/global installation:** No official cloud direct-skill install procedure was found. Do not infer it from the repository checkout behavior.
- **Useful search terms:** `Codex skills`, `.agents/skills`, `skills/list forceReload`, `[[skills.config]]`, `allow_implicit_invocation`, `Codex filesystem skills`.

# Harness Setup Tasks: Evaluation Support

This document owns the native setup contracts for Codex, Cursor, Claude Code, OpenCode, and Pi. The CLI implementation owns their current executable behavior, and the retained qualification records document the bounded runtime evidence.

The evaluation contracts use the CLI's disposable lifecycle: temporary instrumented paths are materialized in the harness project root, cleanup rematerializes current canonical Skill Issue copies only for paths that existed before instrumentation, removes paths introduced for the evaluation, and retains the selected evidence. Ordinary installation has no receipts, backups, rollback inventories, repair commands, update commands, or platform application-state directory.

## Effective Evaluation Configuration

The evaluation command accepts `codex`, `cursor`, `claude-code`, `opencode`, and `pi`. It resolves one model and reasoning value before temporary skills are installed. Callers may provide `--model` and `--reasoning`; omitted values use the defaults below.

| Harness     | Default model              | Default reasoning | Native mapping                                                             |
| ----------- | -------------------------- | ----------------- | -------------------------------------------------------------------------- |
| Codex       | `gpt-5.6-sol`              | `medium`          | `--model` and `--config model_reasoning_effort="<reasoning>"`              |
| Cursor      | native Auto-select         | `medium`          | explicit `--model`; default omits the flag because Cursor owns Auto-select |
| Claude Code | `opus`                     | `medium`          | `--model` and `--effort`                                                   |
| OpenCode    | `openai/gpt-5.6-sol`       | `medium`          | `--model` and `--variant`                                                  |
| Pi          | `openai-codex/gpt-5.6-sol` | `medium`          | `--model` and `--thinking`                                                 |

Model identifiers and supported reasoning values are passed through to the selected harness. Skill Issue does not maintain a model catalogue or prevalidate compatibility; unsupported values remain native harness errors. Cursor uses its native Auto-select model and model-native reasoning when no model override is supplied. An explicit Cursor `--reasoning` override is rejected before evaluation side effects because the CLI exposes no independent reasoning control.

When a harness rejects a model, reasoning value, session, permission, or protocol step, the CLI reports a tooling error and preserves the useful native stderr or structured error text.

## Codex

### Outcome

Codex evaluations launch the user's existing Codex CLI directly against a project-local evaluation workspace, load the generated Skill Issue evaluation skills, exclude ambient user and project customizations, and send every scenario turn verbatim through one resumable headless session.

### Codex CLI Configuration

Resolve the absolute `codex` executable and record `codex --version` before the run. Use the user's normal `CODEX_HOME` so Codex retains its supported authentication path. Do not create a replacement home or copy authentication material.

The configuration below was qualified on macOS with Codex CLI `0.144.1` through initial-session, resumed-session, workspace-write, ambient-skill isolation, automatic-review, structured-event, and tooling-complete one-turn evaluation probes. A complete governed campaign remains separate qualification work.

#### Evaluation workspace

Require an initialized Git workspace selected specifically for the evaluation. Install the generated instrumented skills only at:

```text
<evaluated-workspace>/.agents/skills/<skill-name>/
```

Preserve every skill's canonical name, frontmatter, body, and supporting files. The instrumentation instruction remains immediately after frontmatter. Codex can discover and read these project skills, while its `workspace-write` protections prevent the evaluated agent from modifying `.agents/` during the run.

Do not add the caller-selected output directory as a writable root. The Skill Issue parent process owns evaluation state and evidence outside the evaluated workspace.

#### Ambient customization isolation

Apply all run settings as process arguments or invocation-specific `--config` values. Do not edit the user's Codex files.

For every initial and resumed turn:

- pass `--ignore-user-config` so `$CODEX_HOME/config.toml` is not loaded while authentication still uses the normal Codex home;
- pass `--ignore-rules` so user and project execution-policy rule files are not loaded;
- pass `--disable plugins`;
- pass `--config 'apps._default.enabled=false'`;
- pass `--config 'project_doc_max_bytes=0'` so discovered `AGENTS.md` instructions are excluded;
- pass `--config 'model_reasoning_effort="medium"'`; and
- generate one invocation-specific `skills.config` array that disables every `SKILL.md` found under `$CODEX_HOME/skills`, `$HOME/.agents/skills`, and `/etc/codex/skills`.

Generate the skill deny-list in this form and supply it through one `--config` argument:

```toml
skills.config=[
  {path="<absolute-existing-skill-path>/SKILL.md",enabled=false}
]
```

The deny-list contains ambient skills only. The generated evaluation skills remain enabled because they live in the evaluated workspace rather than the scanned user, compatibility, and administrator roots. Fail the preflight if an existing root cannot be inspected reliably; do not delete, rename, move, or rewrite its contents.

`--ignore-user-config` also excludes user-configured hooks and MCP servers. Codex's compiled product behavior and organization-managed requirements remain authoritative.

#### Initial and resumed turns

Use the same executable, workspace, model, configuration overrides, and isolation flags for every turn. Supply each scenario prompt exactly as stored, without acknowledgement, summarization, adaptation, or content derived from earlier model responses.

Initial turn:

```sh
'<codex-executable>' \
  --cd '<absolute-evaluated-workspace>' \
  --ask-for-approval on-request \
  --sandbox workspace-write \
  --disable plugins \
  --model '<requested-model>' \
  --config 'approvals_reviewer="auto_review"' \
  --config 'model_reasoning_effort="medium"' \
  --config 'project_doc_max_bytes=0' \
  --config 'apps._default.enabled=false' \
  --config '<generated-skills.config>' \
  exec \
  --ignore-user-config \
  --ignore-rules \
  --json \
  '<verbatim-turn-prompt>'
```

Resumed turn:

```sh
'<codex-executable>' \
  --cd '<absolute-evaluated-workspace>' \
  --ask-for-approval on-request \
  --sandbox workspace-write \
  --disable plugins \
  --model '<requested-model>' \
  --config 'approvals_reviewer="auto_review"' \
  --config 'model_reasoning_effort="medium"' \
  --config 'project_doc_max_bytes=0' \
  --config 'apps._default.enabled=false' \
  --config '<generated-skills.config>' \
  exec resume \
  --ignore-user-config \
  --ignore-rules \
  --json \
  '<session-id>' \
  '<verbatim-turn-prompt>'
```

Do not use `--dangerously-bypass-approvals-and-sandbox`, `danger-full-access`, or `--ask-for-approval never`. Do not pass `--ephemeral`, because the runner needs the session identifier to resume the same conversation between separate CLI processes.

#### Structured output and session handling

Capture Codex's newline-delimited JSON output and stderr for every turn.

For the initial turn:

1. Require a `thread.started` event containing a non-empty `thread_id`.
2. Store that identifier in private run state before sending another turn.
3. Require a terminal `turn.completed` event and a successful Codex process exit.

For every resumed turn:

1. Invoke `exec resume` with the stored thread identifier.
2. Require the structured events to belong to that conversation.
3. Wait for `turn.completed` and process exit before closing the active turn or sending the next prompt.

Missing, malformed, or conflicting session events; premature process exit; timeout; cancellation; and failure to reach `turn.completed` are tooling failures rather than model evaluation results. Retain stderr warnings as tooling notes, but do not treat a warning as failure when the structured protocol and process exit are successful.

#### Signal attribution

Codex may send the instrumented `skill-issue signal <opaque-token> <private-state-root>` command to automatic review because the private state root is outside the evaluated workspace. The reviewer may deny that write as unrelated or injected behavior even though the skill was loaded correctly.

For Codex, treat a structured `command_execution` attempt containing the exact resolved Skill Issue executable, `signal`, known opaque token, and expected private state root as activation evidence for the active turn. The external signal process does not need to complete successfully. If it does complete and writes the same activation event, deduplicate the two observations so one expected skill call is counted once.

Do not count arbitrary mentions of the command, transcript text, or malformed command events as activations.

### Authentication

- Run `codex login status` before installing temporary evaluation skills.
- Working authentication and access to the selected model are prerequisites.
- When authentication is absent, stop before replay and direct the user to run `codex login` through Codex's supported browser flow.
- Skill Issue does not read, copy, print, move, or retain Codex credentials.

### Approvals And Sandbox

- `workspace-write` permits normal project work inside the evaluated workspace while preserving Codex's protected-path behavior for `.git/`, `.agents/`, and `.codex/`.
- `on-request` keeps boundary-crossing actions eligible for review instead of failing every such action automatically.
- `approvals_reviewer="auto_review"` routes eligible approval requests to Codex's reviewer agent. The reviewer can approve or deny an individual request, adds model usage, and does not permanently widen the sandbox.
- Do not enable workspace network access for the evaluation.
- A denied signal command is handled through structured-attempt attribution. Other denied, malformed, timed-out, or aborted boundary requests remain tooling failures unless the evaluation contract explicitly owns their interpretation.

### Cleanup

- On completion, cancellation, timeout, or error, terminate only the Codex process started for the active turn.
- Remove temporary instrumented Skill Issue paths from the evaluated workspace.
- For each canonical Skill Issue path that existed before instrumentation, rematerialize the current canonical ordinary copy; remove each temporary evaluation path that did not exist before the run.
- Delete private token mappings and temporary runtime material while retaining the caller-selected evidence.
- Retain the completed evaluation conversation in Codex's normal session store. Do not issue a session-deletion command during normal or recovery cleanup.
- Leave the user's ordinary Codex configuration, authentication, skills, plugins, rules, instruction files, and unrelated conversations unchanged.
- No configuration reset is required because every evaluation override was supplied only to the launched Codex processes.

### Isolation Boundary

- Organization-managed requirements and system configuration remain authoritative and may prevent an evaluation from qualifying.
- The controlled invocation excludes the documented user, compatibility, administrator, project-instruction, rule, app, plugin, hook, and MCP customization routes described above. It cannot remove Codex's compiled product behavior or override organization policy.
- The instrumented skill contains an absolute executable path, a neutral signal instruction, an opaque token, and a private state path. These measures reduce evaluation clues but cannot guarantee that an agent inspecting its environment will not infer that instrumentation exists.
- The configuration and replay controls are locally qualified. A complete built-in governed evaluation and additional operating systems require separate campaign evidence before broader claims are made.

### Platform

This configuration was qualified on macOS with Codex CLI `0.144.1` and `gpt-5.6-sol` at medium reasoning.

## Cursor

### Outcome

Cursor evaluations use the existing Cursor CLI in a macOS-only isolated environment, load the generated Skill Issue evaluation skills, and run every turn through one resumable headless session.

### Cursor CLI Configuration

Cursor may expose its executable as `agent` or the backward-compatible `cursor-agent` alias. Resolve the absolute executable path before replacing the process environment. Cursor CLI `2026.07.16-899851b` contains the required environment overrides, plugin loading, Auto-review, sandbox, streaming output, and resume controls described below.

#### Isolated runtime

Use one run-owned root with this shape:

```text
<runtime>/
├── home/
│   ├── .cursor/cli-config.json
│   └── Library/Keychains -> <user-home>/Library/Keychains
├── store/
├── tmp/
├── workspace/
└── plugin/
    ├── .cursor-plugin/plugin.json
    └── skills/
        └── <skill-name>/SKILL.md
```

Start Cursor from an empty environment and supply only these required values:

```text
HOME=<runtime>/home
CURSOR_CONFIG_DIR=<runtime>/home/.cursor
CURSOR_DATA_DIR=<runtime>/home/.cursor
CURSOR_AGENT_STORE_DIR=<runtime>/store
TMPDIR=<runtime>/tmp
PATH=/usr/bin:/bin:/usr/sbin:/sbin:<directory-containing-resolved-cursor-executable>
SHELL=/bin/zsh
TERM=<current TERM or dumb>
LANG=<current LANG or en_US.UTF-8>
USER=<current user>
LOGNAME=<current logname>
```

`HOME` excludes the user's Cursor skills, plugins, rules, and ordinary configuration. `CURSOR_CONFIG_DIR`, `CURSOR_DATA_DIR`, and `CURSOR_AGENT_STORE_DIR` keep Cursor CLI configuration, data, and session storage under the run-owned root. Removing the runtime removes the isolated Cursor session and generated configuration.

#### CLI configuration

Write `<runtime>/home/.cursor/cli-config.json` with this run-specific shape:

```json
{
  "version": 1,
  "permissions": {
    "allow": [
      "Read(**)",
      "Write(**)",
      "Shell(<absolute-skill-issue-executable>)"
    ],
    "deny": ["Shell(rm)", "Shell(git)", "Read(.env*)", "Write(**/*.key)"]
  },
  "approvalMode": "allowlist",
  "notifications": false,
  "hints": false,
  "modelSlashCommands": false,
  "autoAcceptWebSearch": false,
  "sandbox": {
    "mode": "enabled",
    "networkAccess": "user_config_with_defaults"
  }
}
```

`Read(**)` and `Write(**)` are relative to the caller-selected evaluation workspace. `Shell(<absolute-skill-issue-executable>)` permits the instrumented skill command whose command base is the resolved Skill Issue executable and whose arguments are `signal <opaque-token> <absolute-private-state-root>`. Cursor permission matching uses the executable command base; the signal arguments do not belong in the permission token. Deny rules take precedence over allow rules.

Auto-review can approve other sandboxable commands even when they are absent from the explicit allow-list. Those commands remain subject to the enabled Cursor sandbox. An explicitly allowlisted absolute command can write the evaluation signal to CLI-owned private state outside the Cursor workspace without `--force`.

#### Supplied skills plugin

Generate `<runtime>/plugin/.cursor-plugin/plugin.json` as:

```json
{
  "$schema": "https://cursor.com/schemas/cursor-plugin/plugin.json",
  "name": "skill-issue-evaluation-runtime",
  "version": "1.0.0",
  "description": "Generated Skill Issue evaluation skills.",
  "skills": "skills"
}
```

Place every generated instrumented evaluation skill at `<runtime>/plugin/skills/<skill-name>/SKILL.md`, preserving each skill's supporting files beneath its skill directory. One plugin can expose all supplied skills, and `--plugin-dir <runtime>/plugin` loads that plugin for both initial and resumed turns.

#### Initial and resumed turns

Use the same isolated environment and complete argument set for every turn.

Initial turn:

```text
<cursor-executable>
  --disable-auto-update
  --disable-project-configs
  --workspace <evaluation-workspace>
  --plugin-dir <runtime>/plugin
  --model <requested-model>
  --trust
  --sandbox enabled
  --auto-review
  -p
  --output-format stream-json
  <prompt>
```

Resumed turn:

```text
<cursor-executable>
  --disable-auto-update
  --disable-project-configs
  --workspace <evaluation-workspace>
  --plugin-dir <runtime>/plugin
  --model <requested-model>
  --trust
  --sandbox enabled
  --auto-review
  --resume <session-id>
  -p
  --output-format stream-json
  <prompt>
```

Do not pass `--force` or its `--yolo` alias. Auto-review plus the generated permission configuration permits workspace writes and the evaluation signal while retaining the Cursor sandbox.

The newline-delimited JSON stream begins with a `system/init` event containing `cwd`, `session_id`, `model`, and `apiKeySource`. Cursor repeats the session identifier on later events and resumed turns. A completed turn ends with a `result` event whose `subtype` is `success`, `is_error` is `false`, and `session_id` matches the initialized session. Tool activity appears as `tool_call` events with `started` and `completed` subtypes.

### Authentication

- Link `<runtime>/home/Library/Keychains` to the current macOS user's `~/Library/Keychains` before invoking Cursor. A clean `HOME` without this bridge cannot locate a Keychain in which to store the Cursor login.
- Run `<cursor-executable> status` under the isolated environment. When it reports `Not logged in`, explain that authentication is required for the isolated Cursor evaluation and let the user continue or cancel.
- Run `<cursor-executable> login` under the same isolated environment to start Cursor's browser authorization flow. Successful completion reports the authenticated account and that authentication tokens were stored securely.
- The Keychain bridge lets Cursor retrieve its own `cursor-user` credential. Skill Issue does not read, copy, print, or retain the token.
- Removing the runtime removes the Keychain link, not Cursor's credential in the user's Keychain. Do not run `logout` as evaluation cleanup because that would clear the user's Cursor login.

### Discovery Boundary

- `--disable-project-configs` disables discovered project `cli.json` configuration; it does not disable every project skill, rule, or instruction source.
- The clean temporary `HOME` excludes user-local Cursor, Agents, Claude, and Codex skill roots and user-local Cursor plugins.
- The clean temporary `--workspace` excludes project `.cursor/skills`, `.agents/skills`, `.cursor/rules`, `AGENTS.md`, and other repository instructions. The supplied plugin is the only custom skill source placed inside the isolated run environment.
- Cursor's product-managed built-in skills remain part of its default runtime. Account- or team-enforced Cursor configuration may also remain authoritative because Cursor provides no CLI switch that guarantees their removal.

### Platform

This configuration applies to Cursor CLI headless evaluations on macOS.

## Claude Code

### Outcome

Claude Code evaluations use the existing `claude` CLI in print mode, load the generated Skill Issue evaluation skills, exclude ambient customizations, allow controlled access to the evaluation workspace, and run every turn through one resumable session.

### Claude Code CLI Configuration

Resolve the absolute `claude` executable before creating the run environment. The configuration below was qualified with Claude Code `2.1.205`.

#### Isolated runtime

Use one run-owned root outside the evaluated workspace:

```text
<runtime>/
├── launch/
└── passed-skills/
    └── .claude/
        └── skills/
            ├── <primary-skill>/
            │   └── SKILL.md
            └── <supporting-skill>/
                └── SKILL.md
```

Run every Claude Code process from the empty `<runtime>/launch` directory. Grant the evaluated workspace separately through `permissions.additionalDirectories`; do not pass the workspace through `--add-dir`, because Claude Code discovers skills and other configuration from added directories.

#### Supplied skills

Place every generated instrumented evaluation skill at `<runtime>/passed-skills/.claude/skills/<skill-name>/SKILL.md`, preserving each skill's supporting files beneath its skill directory. Pass `<runtime>/passed-skills` through `--add-dir` on every turn. The skills retain their canonical names.

Claude Code's bundled skills and built-in commands remain available because the direct CLI has no custom-skill allow-list. The clean launch directory, `--setting-sources project`, and isolated added directory exclude user, local, and evaluated-workspace custom skills.

#### Inline settings

Pass this run-specific object through `--settings` on every turn:

```json
{
  "autoMemoryEnabled": false,
  "env": {
    "CLAUDE_CODE_DISABLE_AUTO_MEMORY": "1",
    "CLAUDE_CODE_DISABLE_CLAUDE_MDS": "1",
    "CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC": "1",
    "CLAUDE_CODE_DISABLE_OFFICIAL_MARKETPLACE_AUTOINSTALL": "1",
    "CLAUDE_CODE_DISABLE_BACKGROUND_TASKS": "1",
    "CLAUDE_CODE_DISABLE_CRON": "1"
  },
  "permissions": {
    "additionalDirectories": ["<absolute-evaluated-workspace>"]
  }
}
```

These values apply only to the Claude Code child process. They disable cross-run memory, `CLAUDE.md` loading, nonessential traffic, marketplace auto-installation, background tasks, and cron activity without changing the user's files or shell environment.

#### Workspace routing

Pass this text through `--append-system-prompt` on every turn:

```text
The evaluation workspace is <absolute-evaluated-workspace>. Treat that absolute
directory as the project root. Resolve every relative project path against it.
Do not write project files in the process working directory.
```

#### Initial and resumed turns

Send each scenario turn verbatim as the `-p` prompt. Use the same launch directory, inline settings, supplied-skills directory, workspace routing text, model, effort, tools, and permissions for every turn.

Initial turn:

```sh
'<claude-executable>' -p \
  --setting-sources project \
  --settings '<inline-settings-json>' \
  --strict-mcp-config \
  --no-chrome \
  --add-dir '<absolute-passed-skills-root>' \
  --tools 'Read,Write,Edit,Glob,Grep,Bash' \
  --allowedTools 'Read,Write,Edit,Glob,Grep,Bash(<absolute-skill-issue-executable> signal *)' \
  --permission-mode dontAsk \
  --append-system-prompt '<workspace-routing-text>' \
  --model '<requested-model>' \
  --effort medium \
  --session-id '<generated-uuid>' \
  --output-format stream-json \
  --verbose \
  '<verbatim-turn-prompt>'
```

Resumed turn:

```sh
'<claude-executable>' -p \
  --setting-sources project \
  --settings '<inline-settings-json>' \
  --strict-mcp-config \
  --no-chrome \
  --add-dir '<absolute-passed-skills-root>' \
  --tools 'Read,Write,Edit,Glob,Grep,Bash' \
  --allowedTools 'Read,Write,Edit,Glob,Grep,Bash(<absolute-skill-issue-executable> signal *)' \
  --permission-mode dontAsk \
  --append-system-prompt '<workspace-routing-text>' \
  --model '<requested-model>' \
  --effort medium \
  --resume '<session-id-from-system-init>' \
  --output-format stream-json \
  --verbose \
  '<verbatim-turn-prompt>'
```

Pass the caller-selected Claude Code model through `--model`. When the caller does not select a model, use the `opus` alias. Pass medium reasoning separately through `--effort medium` in both cases. The `opus` alias follows the latest Opus model available through the user's configured Claude Code provider; callers can supply a full model identifier when a fixed model version is required.

`Bash` is present only for the instrumented evaluation signal. The allowed Bash pattern must contain the resolved absolute Skill Issue executable followed by `signal *`; do not allow unrestricted Bash.

Do not use `--safe-mode`, `--bare`, `--disable-slash-commands`, or `--no-session-persistence`. Those flags prevent the qualified supplied-skill, default-behavior, or resumable-session result.

#### Structured output and session handling

Read the first `system/init` event from every newline-delimited JSON stream. Require:

- `cwd` to equal `<runtime>/launch`;
- `session_id` to equal the generated or previously captured session identifier;
- `mcp_servers`, `plugins`, and `plugin_errors` to be empty;
- every supplied custom skill to be present and no user or evaluated-workspace custom skill to be present;
- `tools` to contain exactly the selected tools;
- `memory_paths` to be absent or null; and
- `model` to match the requested model.

Capture the initial event's `session_id` and use it with `--resume` for every later turn. A completed turn requires a successful terminal `result` event. A denial of a required configured capability is a tooling failure; an unrelated out-of-scope tool attempt may be denied by the sandbox while the required turn still completes. Missing or malformed protocol events, configuration mismatches, unexpected customizations, session mismatches, and writes outside the evaluated workspace are tooling failures rather than model results.

### Authentication

- Working Claude Code authentication is a prerequisite. The clean-launch-directory design retains Claude Code's normal authentication path.
- Skill Issue does not read, copy, print, or retain Claude Code credentials.

### Cleanup

- On completion, cancellation, timeout, or error, terminate the private process group created for the Claude command, including the project-local proxy it starts.
- Run `claude project purge --yes <absolute-launch-directory>` under the same Claude configuration root.
- Remove the run-owned launch and passed-skills directories.
- Leave the user's ordinary Claude Code configuration, skills, and credentials unchanged.

### Isolation Boundary

- The clean launch directory and `--setting-sources project` exclude user and local settings because the process has no project configuration to discover there.
- The evaluated workspace is granted through `permissions.additionalDirectories`, so its `.claude` settings, skills, and `CLAUDE.md` files are not discovered.
- The supplied-skills root is the only added directory and therefore the only external custom-skill source.
- Claude Code's product-managed bundled skills and built-in commands remain part of its default runtime.
- Managed policy may remain authoritative because Claude Code does not provide a CLI control that disables it.

### Platform

This configuration was qualified on macOS with Claude Code `2.1.205`.

## Pi

### Outcome

Pi evaluations use the existing Pi CLI in RPC mode, load only the generated Skill Issue evaluation skills, allow the required workspace operations and skill helpers, and run every turn through one persistent machine-readable session.

### Pi CLI Configuration

Resolve the absolute runnable Pi command and any launcher runtime it depends on before constructing the isolated environment. The configuration below was qualified with Pi `0.80.10`.

#### Isolated runtime

Use one run-owned root with this shape:

```text
<runtime>/
├── home/
├── sessions/
└── passed-skills/
    ├── <primary-skill>/
    │   └── SKILL.md
    └── <supporting-skill>/
        └── SKILL.md
```

Start Pi from the caller-selected evaluation workspace with a controlled environment containing these values:

```text
HOME=<runtime>/home
PI_CODING_AGENT_DIR=<existing configured directory or ~/.pi/agent>
PI_CODING_AGENT_SESSION_DIR=<runtime>/sessions
PI_OFFLINE=1
TMPDIR=<runtime>/tmp
PATH=<controlled-tool-path-containing-the-resolved-Pi-runtime>
TERM=<current TERM or xterm-256color>
```

`HOME` keeps ordinary home discovery outside the run while `PI_CODING_AGENT_SESSION_DIR` keeps evaluation sessions temporary. Pi retains its existing agent directory so its supported authentication remains available. Explicit `--no-*` controls disable ambient extensions, skills, prompt templates, themes, context files, and project resources. `PI_OFFLINE=1` disables startup update checks, package update checks, and install or update telemetry without disabling the selected model provider request.

#### Authentication boundary

Pi's [official quickstart](https://github.com/earendil-works/pi/blob/main/packages/coding-agent/docs/quickstart.md) documents subscription login through `/login` and API-key authentication through environment variables or `~/.pi/agent/auth.json`. Skill Issue passes the existing configured agent directory, or that native default, to Pi. It does not create, copy, symlink, replace, or read `auth.json`; Pi remains the credential owner and may perform its normal token refresh. Pi has no native workspace-only filesystem sandbox, so the evaluated process retains the host access of the user who launched it. This limitation remains explicit.

#### Supplied skills

Place every generated instrumented evaluation skill at `<runtime>/passed-skills/<skill-name>/SKILL.md`, preserving each skill's supporting files beneath its skill directory.

Pass `--no-skills` to disable discovered and configured skills, then pass one `--skill <absolute-skill-directory>` argument for every generated evaluation skill. Explicit `--skill` paths remain active when `--no-skills` is present.

After Pi starts, use the RPC `get_commands` request to confirm that every expected `skill:<name>` command resolves to its generated absolute path and that no unexpected prompt-template or extension command was loaded.

#### RPC launch

Launch one Pi process for the complete evaluation:

```text
<resolved-pi-command>
  --mode rpc
  --provider <requested-provider>
  --model <requested-model>
  --thinking medium
  --no-session
  --session-id <generated-uuid>
  --no-approve
  --no-extensions
  --no-skills
  --skill <absolute-generated-skill-directory>
  --no-prompt-templates
  --no-themes
  --no-context-files
  --tools read,bash,edit,write,grep,find,ls
  --offline
```

Repeat `--skill` for every supplied evaluation skill. `--no-approve` prevents project settings and project resources from loading. `--no-context-files` prevents `AGENTS.md` and `CLAUDE.md` discovery. The remaining `--no-*` flags disable ambient extensions, skills, prompt templates, and themes while retaining the explicitly supplied skills.

`--no-session` keeps the conversation in memory without writing a session file. `--session-id` gives the run a stable externally generated identity. Keep this process alive for every scenario turn; repeated RPC prompt requests use the same in-memory conversation and do not require a resume command.

The selected tools allow Pi to load supplied skill instructions, inspect and modify the evaluation workspace, execute skill helpers, and run the instrumented evaluation signal command. Pi applies the filesystem and process permissions of its child process, so the enclosing evaluation runtime owns the workspace-write boundary and access to the exact Skill Issue signal route.

#### RPC conversation

Pi RPC uses one JSON object per line on standard input and standard output. Give every command a unique `id`.

Before the first turn, send:

```json
{"id":"state-before","type":"get_state"}
{"id":"commands-before","type":"get_commands"}
```

Require `get_state` to report:

- the generated `sessionId`;
- the requested provider and model;
- `thinkingLevel` equal to `medium`;
- `isStreaming` and `isCompacting` equal to `false`;
- `pendingMessageCount` equal to `0`; and
- no session file for the ephemeral run.

Send each scenario turn as:

```json
{
  "id": "<unique-turn-id>",
  "type": "prompt",
  "message": "<verbatim-turn-prompt>"
}
```

For every turn:

1. Require the matching `response` object to have `success: true`.
2. Capture the streamed message and tool events for evidence.
3. Keep reading until `agent_settled` reports that retries, compaction, and queued continuations are complete.
4. Send `get_state` and require Pi to be idle with no pending messages before sending the next turn.

The same process and session identifier preserve conversation history across all turns. Assistant `message_end` events contain the completed response, and tool events identify supplied-skill reads, helper execution, and workspace operations.

### Cleanup

- On normal completion, close the RPC input and terminate only the Pi process created for the run.
- On cancellation or timeout, send an RPC `abort` request, wait briefly for its successful response and `agent_settled`, then terminate the owned process if it does not settle.
- Remove the run-owned home, session, and supplied-skill directories.
- `--no-session` leaves no persistent Pi conversation to remove.

### Isolation Boundary

- The private home excludes ordinary home discovery; the existing Pi agent directory remains available through Pi's native configuration contract.
- Skill Issue does not write, copy, replace, or remove the user's Pi authentication directory.
- `--no-approve` excludes evaluated-workspace Pi settings and resources.
- `--no-skills` plus explicit `--skill` arguments make the generated evaluation skills the only custom skills supplied to the run.
- Pi has no native workspace-only filesystem sandbox. The enclosing evaluation process boundary must restrict general writes to the evaluation workspace while preserving the exact evaluation signal route.
- Pi's compiled built-in behavior remains part of its default runtime.

### Platform

This configuration was qualified on macOS with Pi `0.80.10`.

## OpenCode

### Outcome

OpenCode evaluations launch an existing qualified OpenCode CLI with a run-owned configuration, supply only the selected evaluation skills, preserve the operator-owned native authentication route, replay one resumable structured session, and delete that session during cleanup.

### Preflight

- Resolve the absolute `opencode` executable and require version `1.18.4`.
- Require the selected model to use `provider/model` form.
- Run `auth list --pure` and require the selected provider to be authenticated without reading credential contents.
- Run `models <provider> --pure` and require the exact selected model.
- Stop before temporary skill materialization when any preflight fails.

### Runtime Configuration

Create private XDG configuration, state, cache, and temporary roots under the run-owned runtime. Retain `HOME` and the operator-owned XDG data root so OpenCode continues to own native authentication and refresh behavior. Do not copy, print, migrate, or rewrite credentials.

Write the generated skills beneath `<runtime-config>/opencode/skills/<skill-name>/`. Generate `opencode.json` with the effective model, `build` agent, requested variant, selected provider, sharing and updates disabled, snapshots disabled, formatter and LSP disabled, no external plugin list, no added instructions, and no MCP servers.

Launch every command with a minimal environment and these controls:

- `--pure` to exclude external plugins while preserving compiled internal plugins;
- `OPENCODE_DISABLE_PROJECT_CONFIG=true`;
- `OPENCODE_DISABLE_EXTERNAL_SKILLS=true`;
- all Claude compatibility discovery disabled;
- sharing, auto-updates, LSP downloads, and the file watcher disabled.

The deny-first permission policy allows workspace reads and edits, `glob`, `grep`, `list`, the exact selected skill names, and the exact Skill Issue `signal` command through Bash. It denies unrelated Bash, external directories, questions, task delegation, web fetch, and web search.

### Structured Replay

Run the initial turn as:

```sh
'<opencode-executable>' run --pure --format json --model '<provider/model>' --variant '<reasoning>' '<verbatim-turn-prompt>'
```

Resume later turns by adding `--session '<captured-sessionID>'`. Parse newline-delimited JSON envelopes, capture camel-case `sessionID`, require that every reported session identifier remains stable, reject structured error envelopes, and require a terminal `step_finish` whose `part.reason` is `stop`. A failed exact marker Bash call is a tooling failure; unrelated denied tool calls remain ordinary structured events when the turn otherwise completes.

### Cleanup

- On cancellation or timeout, terminate only the process group started for the active turn.
- Recover the session identifier from partial structured output when interruption occurs before normal capture.
- Run `session delete <sessionID> --pure` with the same environment before removing the run-owned runtime.
- Remove generated skills, configuration, state, cache, and temporary files while preserving operator-owned native authentication.

### Platform

This configuration was qualified on macOS ARM64 with OpenCode `1.18.4`, native OpenAI ChatGPT OAuth, `openai/gpt-5.6-sol`, and the `medium` variant. Full governed-campaign and additional-platform qualification remain separate work.

## Local Smoke Qualification

The 2026-07-20 smoke campaign used only two-turn inputs and did not run any complete governed evaluation. Codex `0.144.1` completed the built-in default route and the explicit `gpt-5.6-sol`/`medium` custom route through one resumable session each. The pre-run confirmation showed effective values; transcripts and signal events prove selected-skill loading, turn attribution, and argument propagation; output artifacts were written under the caller-selected roots; and temporary skill contents and private run state were removed after completion.

Cursor Agent `2026.07.16-899851b`, Claude Code `2.1.205`, and Pi `0.80.10` completed built-in and custom two-turn routes through their project-local or installed launchers. Cursor and Pi wrote directly in the caller-selected evaluation workspace. Claude received that workspace through its isolated settings and system prompt. Final Cursor and Claude runs left no detached worker or proxy, and every run removed private runtime and generated skill material. Pi reused its existing `openai-codex` authentication without the CLI modifying the auth directory. Custom smoke artifacts are local evidence only and require separate review before any publication claim.

The qualification results, versions, observed calls, cleanup checks, and useful native errors are in `evaluations/skill-calling/smoke/real-harness-smoke-report.md`.

# Competition Rules

For any competition or hackathon-rules question, read and follow the authoritative [OpenAI Build Week rules](https://openai.devpost.com/rules).

# Website Development

The repository root contains the static React, TypeScript, and Vite website.

After editing TypeScript or TSX, run:

1. `npm run typecheck`
2. `npm run lint`

After editing CSS, JSON, Markdown, HTML, or configuration, run `npm run format:check`. Before presenting website work as complete, run `npm run validate`; a passing server or browser render is insufficient.

Use `npm run dev -- --host 127.0.0.1` for browser automation and `npm run preview` for production-build inspection. Keep curated website copy, release metadata, and summary metrics in `src/data/siteData.ts`; keep website evaluation types, artifact adaptation, labels, and illustrative result data in `src/data/evaluationData.ts`; keep selection and filtering in `src/components/EvaluationExplorer.tsx`; keep each chart presentation in its component under `src/components/charts/`; and keep shared theme and layout decisions in `src/styles.css`.

# Local CLI Development

Use `cli/scripts/local-cli.sh` to keep the CLI used for repository work separate from the CLI built from the changing working tree. This is developer wrapper logic only. Keep channel selection out of the Go CLI, its embedded payload, installation and evaluation behavior, and the cross-platform release build.

The wrapper stores generated binaries under ignored repository-local `.skill-issue/bin/` state:

- `./cli/scripts/local-cli.sh build-known-good` builds committed `HEAD` from a temporary Git archive, stores the executable in an immutable commit-specific directory, and selects it as known-good. Uncommitted working-tree changes are excluded.
- `./cli/scripts/local-cli.sh build-development` rebuilds the development executable from the current working tree and labels its version with the current commit and relevant clean or dirty state.
- `./cli/scripts/local-cli.sh <skill-issue arguments>` runs known-good by default. `./cli/scripts/local-cli.sh known-good <arguments>` is the explicit equivalent.
- `./cli/scripts/local-cli.sh development <arguments>` runs the development executable only when current-source behavior is explicitly requested.
- `./cli/scripts/local-cli.sh paths` reports both local channel paths.

Unless the user explicitly requests the development CLI or current-source testing, use known-good. Build a new known-good snapshot only from the committed revision intended to become the next working baseline. Keep known-good and development evaluation artifacts in separate descriptive subdirectories under `output/`. Never force-add `.skill-issue/` binaries or state.

# Local Harness Evaluation Routing

Use explicit executable selection for local evaluations so interactive shell aliases cannot choose the wrong harness route. Go resolves harness commands with `exec.LookPath`; shell aliases such as `claude-codex` are invisible to the CLI and are only for interactive terminal use.

Before a Claude Code evaluation, define the intended local routes without recording machine-specific paths in tracked files:

```zsh
export CLAUDE_NORMAL_EXECUTABLE="$(whence -p claude)"
export CLAUDE_CODEX_EXECUTABLE="$(git rev-parse --show-toplevel)/.skill-issue/claudex/claudex"
```

Select the route explicitly for every Claude Code evaluation:

- Normal Claude Code: pass `--harness claude-code --executable "$CLAUDE_NORMAL_EXECUTABLE" --model <qualified-native-Claude-model> --reasoning medium`.
- Claude Code through the local Codex-backed proxy: pass `--harness claude-code --executable "$CLAUDE_CODEX_EXECUTABLE" --model gpt-5.6-sol --reasoning medium`.

Never pass the `claude-codex` alias to `--executable`. Before confirming an evaluation, verify the pre-run summary shows the intended executable, model, reasoning, workspace, and output root.

When Codex runs an OpenAI Codex harness evaluation, execute the exact known-good evaluation command with `sandbox_permissions: "require_escalated"`. State in the justification that the nested Codex process needs its normal authenticated session database and session state under `CODEX_HOME`. This escalation applies only to the outer shell command. Keep the evaluator-owned inner Codex `workspace-write` sandbox, approval policy, ambient-configuration exclusions, model, reasoning, workspace, and cleanup unchanged. Never use `danger-full-access` or bypass approvals. If escalation is denied, treat the run as blocked, verify partial preparation is cleaned, and do not retry the command inside the read-only outer sandbox.

# Planning and Research Documents

Keep active working documents tracked, understandable, and out of the repository root and implementation directories. Store plans and progress in `plans/`; store research, investigation findings, evaluation reports, audits, and supporting documents in `research/`. Keep permanent or domain-owned files at their semantic home.

Create reports, results, or other working documents only when requested or required by a concrete follow-up task. If there is no clear downstream use, communicate the findings directly instead.

Treat working documents in `plans/` and `research/` as temporary unless they provide lasting value. When work ends, recommend removing finished or unneeded documents for deliberate cleanup.

# CLI Evaluation Outputs

For repository development or testing, use repository-root `output/` as the CLI `--output` root unless the user requests that specific output be retained. Put separate runs or campaigns in descriptive subdirectories when useful. Keep evaluation workspaces outside this repository so the output root remains outside the evaluated workspace.

The skill-calling evaluation campaign under `plans/skill-calling-evaluation-campaign/` is an explicit exception: its authoritative output root is the `output/` directory inside each neutral `<repository-parent>/chats/chat-<n>/` container. Follow the campaign orchestration prompt for allocation, cleanup, and same-container retries.

Keep `output/` ignored and untracked; its disposable artifacts include `result.json`, `website.json`, optional `events.jsonl` and `transcript.json`, and temporary `.skill-issue/` recovery state. Never place disposable CLI output elsewhere in the repository or force-add `output/`. Put explicitly retained output in an appropriate non-ignored location for review and commit.

# Local State

Keep repository-root `.skill-issue/` ignored and untracked. Under `.codex/` and `.codex-plugin/`, track only deliberate project configuration; ignore all generated or local state and never force-add it.

# Repository Privacy

[Read repository privacy here](.repository-privacy.md).

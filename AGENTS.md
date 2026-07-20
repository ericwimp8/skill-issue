# Competition Rules

For any competition or hackathon-rules question, read and follow the authoritative [OpenAI Build Week rules](https://openai.devpost.com/rules).

# Website Development

The repository root contains the static React, TypeScript, and Vite website.

After editing TypeScript or TSX, run:

1. `npm run typecheck`
2. `npm run lint`

After editing CSS, JSON, Markdown, HTML, or configuration, run `npm run format:check`. Before presenting website work as complete, run `npm run validate`; a passing server or browser render is insufficient.

Use `npm run dev -- --host 127.0.0.1` for browser automation and `npm run preview` for production-build inspection. Keep curated website copy, release metadata, and summary metrics in `src/data/siteData.ts`; keep website evaluation types, artifact adaptation, labels, and illustrative result data in `src/data/evaluationData.ts`; keep selection and filtering in `src/components/EvaluationExplorer.tsx`; keep each chart presentation in its component under `src/components/charts/`; and keep shared theme and layout decisions in `src/styles.css`.

# Planning and Research Documents

Keep active working documents tracked, understandable, and out of the repository root and implementation directories. Store plans and progress in `plans/`; store research, investigation findings, evaluation reports, audits, and supporting documents in `research/`. Keep permanent or domain-owned files at their semantic home.

Create reports, results, or other working documents only when requested or required by a concrete follow-up task. If there is no clear downstream use, communicate the findings directly instead.

Treat working documents in `plans/` and `research/` as temporary unless they provide lasting value. When work ends, recommend removing finished or unneeded documents for deliberate cleanup.

# CLI Evaluation Outputs

For repository development or testing, use repository-root `output/` as the CLI `--output` root unless the user requests that specific output be retained. Put separate runs or campaigns in descriptive subdirectories when useful. Keep evaluation workspaces outside this repository so the output root remains outside the evaluated workspace.

Keep `output/` ignored and untracked; its disposable artifacts include `result.json`, `website.json`, optional `events.jsonl` and `transcript.json`, and temporary `.skill-issue/` recovery state. Never place disposable CLI output elsewhere in the repository or force-add `output/`. Put explicitly retained output in an appropriate non-ignored location for review and commit.

# Local State

Keep repository-root `.skill-issue/` ignored and untracked. Under `.codex/` and `.codex-plugin/`, track only deliberate project configuration; ignore all generated or local state and never force-add it.

# Repository Privacy

[Read repository privacy here](.repository-privacy.md).

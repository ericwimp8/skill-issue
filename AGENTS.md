# Competition Rules

The official rules for this competition—the OpenAI Build Week hackathon—are available at:

https://openai.devpost.com/rules

If the user asks any question about the competition or hackathon rules, read and use the linked rules before answering. Treat the linked page as the authoritative source for the hackathon rules.

# Website Development

The repository root contains the static Skill Issue website built with React, TypeScript, and Vite.

After editing TypeScript or TSX, run the fastest deterministic owner checks before moving on:

1. `npm run typecheck`
2. `npm run lint`

After editing CSS, JSON, Markdown, HTML, or configuration, run `npm run format:check`.

Before presenting website work as complete, run `npm run validate`. A passing development server or browser render does not replace this command.

Use `npm run dev -- --host 127.0.0.1` for browser automation and `npm run preview` to inspect the production build. Keep benchmark content in `src/data/siteData.ts`, reusable chart behavior in `src/components/EvaluationChart.tsx`, and shared theme/layout decisions in `src/styles.css`.

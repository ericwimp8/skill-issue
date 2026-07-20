# Website Reference And Architecture Decision

## Product Boundary

Skill Issue needs a public facade for a local-first skill-evaluation product. The website owns explanation, benchmark presentation, and a route to downloadable CLI artifacts. GitHub Pages owns static hosting; GitHub Releases owns versioned binaries. Benchmark content remains repository-owned local data, so the first website requires neither a runtime API nor a database.

## Direct Visual Research

Each reference below was opened directly in the Codex in-app browser at a consistent 1440 × 900 viewport on 19 July 2026. The resulting observations informed this decision; the disposable screenshots were removed after the direction was implemented and validated.

| Reference                           | Relevant visible characteristics                                                               |
| ----------------------------------- | ---------------------------------------------------------------------------------------------- |
| [Linear](https://linear.app/)       | Near-black canvas, hairline header boundary, restrained navigation, large negative space       |
| [Vercel](https://vercel.com/)       | Monochrome palette, split hero, geometric focal element, compact pill actions                  |
| [Raycast](https://www.raycast.com/) | Floating rounded navigation surface, centered download actions, low-contrast depth             |
| [Resend](https://resend.com/)       | Sparse black canvas, compact navigation, minimal outlined primary action                       |
| [Clerk](https://clerk.com/)         | Light technical grid, oversized centered headline, narrow copy, floating navigation            |
| [Warp](https://www.warp.dev/)       | Editorial two-column hero, clear download command treatment, strong black-on-white contrast    |
| [Supabase](https://supabase.com/)   | Dark modular card grid, one restrained accent, product data presented inside reusable surfaces |
| [Railway](https://railway.com/)     | Bounded hero surface, centered serif headline, muted purple accent, framed product preview     |
| [shadcn/ui](https://ui.shadcn.com/) | Monochrome tokens, centered hero, pill controls, thin borders, varied reusable card grid       |
| [OpenAI](https://openai.com/)       | Black canvas, modest top navigation, centered interaction surface, large calm spacing          |

## Selected Direction

Use **shadcn/ui's homepage composition** as the primary visual inspiration because its hero-to-card-grid transition fits a graph-led benchmark page without forcing decorative imagery. Its visible system also translates cleanly between light and dark themes.

### Binding Inspiration

- Restrained neutral palette with one semantic accent family.
- Compact top navigation with an immediately visible theme control and primary action.
- Centered, narrow hero copy with a small status pill and one dominant action.
- Thin bordered surfaces with modest radii and minimal shadow.
- A responsive card grid whose content creates hierarchy through span and density.
- Quiet typography, generous section rhythm, and concise labels.

### Non-Binding Reference Details

- shadcn/ui branding, copy, iconography, navigation destinations, sample form controls, QR code, and exact card contents.
- Exact dimensions, font files, and component implementation.
- Dark-first rendering; Skill Issue must treat light and dark as equal top-level themes.

## One-Page Information Architecture

1. Compact header: identity, Results anchor, Method anchor, theme control, download action.
2. Hero: status, problem statement, concise explanation, CLI download and results actions.
3. Benchmark overview: three small summary metrics sourced from local data.
4. Evaluation results: responsive graph cards for Codex and Claude Code.
5. Method note: what context consumption, skill calls, and skill misses mean.
6. Minimal footer: current local-first status and repository-ready placeholders only.

## Stack Decision

Use **React 19 + TypeScript + Vite 8** with plain authored CSS organized around global custom-property tokens.

- React provides the smallest familiar component model for data-driven graph cards and theme state.
- TypeScript makes the local benchmark schema and future graph additions checkable at build time.
- Vite produces static `dist` output and documents GitHub Pages repository-base configuration and local production preview.
- Plain CSS variables keep the design system explicit and avoid importing a broader UI framework for a one-page site.
- ESLint and TypeScript own static correctness; Prettier owns deterministic formatting.

The project base path is configurable through `VITE_BASE_PATH`, defaulting to `/skill-issue/` for a GitHub Pages project site. Local development remains rooted at `/` through Vite's development server behavior.

## Chart Library Decision

Use **Recharts 3.9**.

- It is a composable React chart library with first-party line-chart, tooltip, legend, axes, and responsive-container support.
- Its repository reported approximately 27,400 GitHub stars during research, providing a reasonable popularity signal.
- It is released under the permissive [MIT License](https://github.com/recharts/recharts/blob/main/LICENSE).
- Its official examples include multi-series line charts and responsive containers, matching skill calls and skill misses over consumed context.
- SVG output can inherit the website's theme tokens, while adjacent semantic summaries keep the chart meaning available without relying only on color.

The website uses typed evaluation artifacts, derives display identities and selection inputs in `src/data/evaluationData.ts`, and renders concrete comparison views through `src/components/EvaluationExplorer.tsx` and `src/components/charts/`.

## Data And Release Ownership

- `src/data/siteData.ts` owns product copy, release metadata, summary metrics, and methodology text.
- `src/data/evaluationData.ts` owns website artifact adaptation, display identities, selection inputs, and temporary illustrative result data.
- `src/components/EvaluationExplorer.tsx` and `src/components/charts/` own evaluation interaction and chart presentation.
- `src/styles.css` owns shared theme and layout tokens.
- The release button uses the configurable GitHub Releases URL in the local data source. Until a real CLI asset exists, it points to the repository's latest-release page rather than implying a working binary download.

## Source Basis

- [Vite static deployment and GitHub Pages guidance](https://vite.dev/guide/static-deploy.html#github-pages)
- [Recharts official examples](https://recharts.github.io/en-US/examples/SimpleLineChart/)
- [Recharts MIT license](https://github.com/recharts/recharts/blob/main/LICENSE)
- [GitHub Pages documentation](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages)
- Existing repository research: [Free Website Hosting and CLI Downloads](../../research/deep-research/free-website-hosting-cli-downloads/free-website-hosting-cli-downloads-deep-research.md)

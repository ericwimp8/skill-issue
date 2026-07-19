# Website Data And Design Contract

## Local Data Contract

`src/data/siteData.ts` is the single content source. Page components may transform its values for presentation but must not introduce replacement copy, release URLs, metrics, graph definitions, or graph values.

```ts
type EvaluationPoint = {
  contextConsumed: number;
  skillCalls: number;
  skillMisses: number;
};

type EvaluationGraph = {
  id: string;
  harness: string;
  model: string;
  description: string;
  sampleSize: number;
  points: EvaluationPoint[];
};
```

The exported site object also owns:

- concise hero and methodology copy;
- the project status label;
- GitHub repository and latest-release URLs;
- the download action label and availability note;
- benchmark summary metrics;
- chart axis and series labels;
- footer copy.

Adding another graph requires one new `EvaluationGraph` entry. The graph grid maps over the entries without harness-specific page markup.

## Design Token Contract

`src/styles.css` owns the following global token families on `:root`, with theme-specific color values overridden by `[data-theme='dark']`:

- `--color-*`: canvas, surface, elevated surface, text, muted text, line, accent, success, miss, focus.
- `--font-*`: sans and mono families.
- `--space-*`: the shared spacing scale.
- `--radius-*`: control and surface radii.
- `--shadow-*`: restrained elevation values.
- `--content-*`: readable and wide content bounds.
- `--transition-*`: motion timing.

Components must consume these tokens. A component-local value is allowed only when it represents that component's unique geometry rather than a repeated design decision.

## Theme Contract

- The initial theme comes from persisted user choice when present, otherwise `prefers-color-scheme`.
- The active value is written to `document.documentElement.dataset.theme` before React renders.
- The theme control exposes its purpose and current action through an accessible label.
- Both themes use the same semantic success and miss meanings with sufficient visual contrast.
- Charts read their stroke, grid, and label colors from CSS custom properties so theme switching has one owner.

## Component Ownership

- `App`: page composition and section order.
- `ThemeToggle`: persisted theme state and root theme attribute.
- `EvaluationChart`: graph surface, legend, axes, tooltip, and data summary.
- `MetricCard`: summary metric presentation.
- `styles.css`: layout primitives, variants, responsive behavior, and theme tokens.

This deliberately shallow structure keeps every concrete behavior within three calls of the page entry point.

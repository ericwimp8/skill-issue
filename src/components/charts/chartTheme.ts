import {
  availableCells,
  type EvaluationResult,
} from '../../data/evaluationData';

export const tooltipStyle = {
  background: 'var(--color-surface-elevated)',
  border: '1px solid var(--color-line)',
  borderRadius: 'var(--radius-control)',
  color: 'var(--color-text)',
};

export type SeriesMarkerShape = 'circle' | 'diamond' | 'square' | 'triangle';

const modelVisuals: Record<
  string,
  { color: string; marker: SeriesMarkerShape }
> = {
  'codex-sol': { color: 'var(--color-series-1)', marker: 'circle' },
  'claude-fable': { color: 'var(--color-series-2)', marker: 'diamond' },
  grok: { color: 'var(--color-series-3)', marker: 'square' },
  composer: { color: 'var(--color-series-4)', marker: 'triangle' },
};

function modelForCell(cellId: string) {
  return availableCells.find((cell) => cell.id === cellId)?.model;
}

export function chartColorForCell(cellId: string) {
  const model = modelForCell(cellId);
  return modelVisuals[model ?? '']?.color ?? 'var(--color-series-1)';
}

export function chartMarkerForCell(cellId: string): SeriesMarkerShape {
  const model = modelForCell(cellId);
  return modelVisuals[model ?? '']?.marker ?? 'circle';
}

export function resultTotal(result: EvaluationResult) {
  return result.points.reduce(
    (total, point) => total + point.called + point.missed,
    0,
  );
}

export function resultCalled(result: EvaluationResult) {
  return result.points.reduce((total, point) => total + point.called, 0);
}

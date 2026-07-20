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

const seriesColors = [
  'var(--color-series-1)',
  'var(--color-series-2)',
  'var(--color-series-3)',
  'var(--color-series-4)',
  'var(--color-series-5)',
  'var(--color-series-6)',
  'var(--color-series-7)',
  'var(--color-series-8)',
  'var(--color-series-9)',
  'var(--color-series-10)',
];

export function chartColorForCell(cellId: string) {
  const index = availableCells.findIndex((cell) => cell.id === cellId);
  return seriesColors[index % seriesColors.length] ?? seriesColors[0];
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

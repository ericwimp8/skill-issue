import {
  CartesianGrid,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';

import type { EvaluationResult } from '../../data/evaluationData';
import { chartColorForCell, tooltipStyle } from './chartTheme';

type CumulativeCallsChartProps = {
  results: EvaluationResult[];
};

export function CumulativeCallsChart({ results }: CumulativeCallsChartProps) {
  const totalTurns = results[0]?.total_turns ?? 30;
  const data = Array.from({ length: totalTurns }, (_, index) => {
    const turn = index + 1;
    const row: Record<string, number> = { turn };

    results.forEach((result) => {
      row[result.cellId] = result.points
        .filter((point) => point.turn <= turn)
        .reduce((total, point) => total + point.called, 0);
    });

    return row;
  });

  return (
    <article className="exploration-chart">
      <header className="exploration-chart-header">
        <div>
          <span className="chart-number">02</span>
          <p className="card-kicker">Cumulative curve</p>
          <h3>How quickly successful calls accumulate.</h3>
        </div>
        <span className="chart-purpose">Timing</span>
      </header>
      <p className="chart-description-wide">
        Step lines make the scored moments explicit while keeping all thirty
        turns available for direct overlays.
      </p>
      <div className="chart-canvas chart-canvas-standard">
        <ResponsiveContainer width="100%" height="100%">
          <LineChart
            data={data}
            margin={{ top: 18, right: 16, bottom: 6, left: -10 }}
          >
            <CartesianGrid
              stroke="var(--color-chart-grid)"
              strokeDasharray="3 6"
              vertical={false}
            />
            <XAxis
              dataKey="turn"
              type="number"
              domain={[1, totalTurns]}
              ticks={[1, 5, 10, 15, 20, 25, 30]}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <YAxis
              domain={[0, 4]}
              ticks={[0, 1, 2, 3, 4]}
              allowDecimals={false}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <Tooltip
              contentStyle={tooltipStyle}
              labelFormatter={(turn) => `Turn ${turn}`}
            />
            {results.map((result) => (
              <Line
                key={result.cellId}
                type="stepAfter"
                dataKey={result.cellId}
                name={result.cellLabel}
                stroke={chartColorForCell(result.cellId)}
                strokeWidth={2.5}
                dot={{ r: 2.5, strokeWidth: 0 }}
                activeDot={{ r: 5, strokeWidth: 0 }}
                isAnimationActive={false}
              />
            ))}
          </LineChart>
        </ResponsiveContainer>
      </div>
      <div className="series-key" aria-label="Selected comparison series">
        {results.map((result) => (
          <span key={result.cellId}>
            <i style={{ backgroundColor: chartColorForCell(result.cellId) }} />
            {result.cellLabel}
          </span>
        ))}
      </div>
    </article>
  );
}

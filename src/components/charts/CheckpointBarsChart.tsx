import {
  Bar,
  BarChart,
  CartesianGrid,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';

import type { EvaluationResult } from '../../data/evaluationData';
import { tooltipStyle } from './chartTheme';

type CheckpointBarsChartProps = {
  results: EvaluationResult[];
};

export function CheckpointBarsChart({ results }: CheckpointBarsChartProps) {
  const turns = Array.from(
    new Set(
      results.flatMap((result) => result.points.map((point) => point.turn)),
    ),
  ).sort((left, right) => left - right);
  const data = turns.map((turn) =>
    results.reduce(
      (row, result) => {
        const point = result.points.find(
          (candidate) => candidate.turn === turn,
        );

        return {
          turn,
          called: row.called + (point?.called ?? 0),
          missed: row.missed + (point?.missed ?? 0),
        };
      },
      { turn, called: 0, missed: 0 },
    ),
  );

  return (
    <article className="exploration-chart">
      <header className="exploration-chart-header">
        <div>
          <span className="chart-number">03</span>
          <p className="card-kicker">Checkpoint profile</p>
          <h3>Where the selected group misses.</h3>
        </div>
        <span className="chart-purpose">Pattern</span>
      </header>
      <p className="chart-description-wide">
        Selected environments are combined by scored turn to reveal which
        expected activation points create the most misses.
      </p>
      <div className="chart-canvas chart-canvas-standard">
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={data}
            margin={{ top: 18, right: 20, bottom: 6, left: -4 }}
          >
            <CartesianGrid
              stroke="var(--color-chart-grid)"
              strokeDasharray="3 6"
              vertical={false}
            />
            <XAxis
              dataKey="turn"
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
              tickFormatter={(turn: number) => `Turn ${turn}`}
            />
            <YAxis
              allowDecimals={false}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <Tooltip
              contentStyle={tooltipStyle}
              labelFormatter={(turn) => `Turn ${turn}`}
            />
            <Bar
              dataKey="called"
              name="Called"
              stackId="outcome"
              fill="var(--color-call)"
              radius={[5, 5, 0, 0]}
              isAnimationActive={false}
            />
            <Bar
              dataKey="missed"
              name="Missed"
              stackId="outcome"
              fill="var(--color-miss)"
              radius={[5, 5, 0, 0]}
              isAnimationActive={false}
            />
          </BarChart>
        </ResponsiveContainer>
      </div>
    </article>
  );
}

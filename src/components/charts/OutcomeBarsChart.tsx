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
import { resultCalled, resultTotal, tooltipStyle } from './chartTheme';

type OutcomeBarsChartProps = {
  results: EvaluationResult[];
};

export function OutcomeBarsChart({ results }: OutcomeBarsChartProps) {
  const data = results.map((result) => {
    const called = resultCalled(result);
    const total = resultTotal(result);

    return {
      called,
      label: result.cellLabel,
      missed: total - called,
    };
  });
  const chartHeight = Math.max(320, results.length * 46 + 72);

  return (
    <article className="exploration-chart">
      <header className="exploration-chart-header">
        <div>
          <span className="chart-number">04</span>
          <p className="card-kicker">Outcome bars</p>
          <h3>The fastest overall comparison.</h3>
        </div>
        <span className="chart-purpose">Summary</span>
      </header>
      <p className="chart-description-wide">
        Each bar keeps the five expected activations visible as called and
        missed parts of the same total.
      </p>
      <div className="chart-canvas" style={{ height: chartHeight }}>
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={data}
            layout="vertical"
            margin={{ top: 12, right: 20, bottom: 4, left: 10 }}
          >
            <CartesianGrid
              stroke="var(--color-chart-grid)"
              strokeDasharray="3 6"
              horizontal={false}
            />
            <XAxis
              type="number"
              domain={[0, 5]}
              ticks={[0, 1, 2, 3, 4, 5]}
              allowDecimals={false}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <YAxis
              type="category"
              dataKey="label"
              width={190}
              interval={0}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <Tooltip contentStyle={tooltipStyle} />
            <Bar
              dataKey="called"
              name="Called"
              stackId="outcome"
              fill="var(--color-call)"
              radius={[5, 0, 0, 5]}
              isAnimationActive={false}
            />
            <Bar
              dataKey="missed"
              name="Missed"
              stackId="outcome"
              fill="var(--color-miss)"
              radius={[0, 5, 5, 0]}
              isAnimationActive={false}
            />
          </BarChart>
        </ResponsiveContainer>
      </div>
    </article>
  );
}

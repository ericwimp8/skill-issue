import {
  CartesianGrid,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from 'recharts';

import { siteData, type EvaluationGraph } from '../data/siteData';

type EvaluationChartProps = {
  graph: EvaluationGraph;
};

export function EvaluationChart({ graph }: EvaluationChartProps) {
  const finalPoint = graph.points.at(-1)!;

  return (
    <article className="chart-card">
      <header className="chart-header">
        <div>
          <span className="card-kicker">{graph.model}</span>
          <h3>{graph.harness}</h3>
        </div>
        <span className="sample-size">n = {graph.sampleSize}</span>
      </header>

      <p className="chart-description">{graph.description}</p>

      <div
        className="chart-wrap"
        role="img"
        aria-label={`${graph.harness}: ${siteData.chart.calls} fall from ${graph.points[0]?.skillCalls} to ${finalPoint.skillCalls}, while ${siteData.chart.misses} rise from ${graph.points[0]?.skillMisses} to ${finalPoint.skillMisses} as context consumption reaches 100 percent.`}
      >
        <ResponsiveContainer width="100%" height="100%">
          <LineChart
            data={graph.points}
            margin={{ top: 12, right: 8, left: -18, bottom: 4 }}
          >
            <CartesianGrid
              stroke="var(--color-chart-grid)"
              strokeDasharray="3 6"
              vertical={false}
            />
            <XAxis
              dataKey="contextConsumed"
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
              tickFormatter={(value: number) => `${value}%`}
            />
            <YAxis
              domain={[0, 20]}
              ticks={[0, 5, 10, 15, 20]}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <Tooltip
              cursor={{ stroke: 'var(--color-line)' }}
              contentStyle={{
                background: 'var(--color-surface-elevated)',
                border: '1px solid var(--color-line)',
                borderRadius: 'var(--radius-control)',
                color: 'var(--color-text)',
              }}
              labelFormatter={(value) => `${value}% context consumed`}
            />
            <Line
              type="monotone"
              dataKey="skillCalls"
              name={siteData.chart.calls}
              stroke="var(--color-call)"
              strokeWidth={2.5}
              dot={false}
              activeDot={{ r: 4, strokeWidth: 0 }}
            />
            <Line
              type="monotone"
              dataKey="skillMisses"
              name={siteData.chart.misses}
              stroke="var(--color-miss)"
              strokeWidth={2.5}
              dot={false}
              activeDot={{ r: 4, strokeWidth: 0 }}
            />
          </LineChart>
        </ResponsiveContainer>
      </div>

      <footer className="chart-footer">
        <div className="legend" aria-label="Chart legend">
          <span>
            <i className="legend-line legend-line-call" />
            {siteData.chart.calls}
          </span>
          <span>
            <i className="legend-line legend-line-miss" />
            {siteData.chart.misses}
          </span>
        </div>
        <span>{siteData.chart.xAxis}</span>
      </footer>
    </article>
  );
}

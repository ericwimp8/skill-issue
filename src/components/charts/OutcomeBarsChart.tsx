import {
  Bar,
  BarChart,
  CartesianGrid,
  LabelList,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
  type TooltipContentProps,
} from 'recharts';

import type { EvaluationResult } from '../../data/evaluationData';
import { resultCalled, resultTotal } from './chartTheme';

type OutcomeBarsChartProps = {
  results: EvaluationResult[];
};

type OutcomeDatum = {
  called: number;
  failed: number;
  failureRate: number;
  label: string;
  successRate: number;
  total: number;
};

function OutcomeTooltip({ active, payload }: TooltipContentProps) {
  const result = payload?.[0]?.payload as OutcomeDatum | undefined;

  if (!active || !result) {
    return null;
  }

  return (
    <div className="chart-tooltip">
      <strong>{result.label}</strong>
      <span>
        {result.successRate.toFixed(0)}% success · {result.called} called
      </span>
      <span>
        {result.failureRate.toFixed(0)}% failure · {result.failed} missed
      </span>
      <span>{result.total} expected calls</span>
    </div>
  );
}

function percentageLabel(value: unknown) {
  const percentage = Number(value);
  return percentage >= 10 ? `${percentage.toFixed(0)}%` : '';
}

export function OutcomeBarsChart({ results }: OutcomeBarsChartProps) {
  const data = results
    .map((result): OutcomeDatum => {
      const called = resultCalled(result);
      const total = resultTotal(result);
      const failed = total - called;

      return {
        called,
        failed,
        failureRate: total === 0 ? 0 : (failed / total) * 100,
        label: result.cellLabel,
        successRate: total === 0 ? 0 : (called / total) * 100,
        total,
      };
    })
    .sort(
      (left, right) =>
        right.successRate - left.successRate ||
        left.label.localeCompare(right.label),
    );
  const chartHeight = Math.max(320, results.length * 46 + 72);

  return (
    <article className="exploration-chart">
      <header className="exploration-chart-header">
        <div>
          <span className="chart-number">04</span>
          <p className="card-kicker">Success ranking</p>
          <h3>Which setups call skills reliably?</h3>
        </div>
        <span className="chart-purpose">Overall performance</span>
      </header>
      <p className="chart-description-wide">
        Every expected call in the filtered results contributes to one success
        or failure rate. Bars are ranked from strongest to weakest regardless of
        sample size.
      </p>
      <div className="raster-legend" aria-label="Success ranking legend">
        <span>
          <i className="legend-block legend-block-called" /> Success
        </span>
        <span>
          <i className="legend-block legend-block-missed" /> Failure
        </span>
      </div>
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
              domain={[0, 100]}
              ticks={[0, 25, 50, 75, 100]}
              tickFormatter={(value: number) => `${value}%`}
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
            <Tooltip content={OutcomeTooltip} />
            <Bar
              dataKey="successRate"
              name="Success"
              stackId="outcome"
              fill="var(--color-call)"
              radius={[5, 0, 0, 5]}
              isAnimationActive={false}
            >
              <LabelList
                dataKey="successRate"
                position="insideLeft"
                fill="#ffffff"
                formatter={percentageLabel}
              />
            </Bar>
            <Bar
              dataKey="failureRate"
              name="Failure"
              stackId="outcome"
              fill="var(--color-miss)"
              radius={[0, 5, 5, 0]}
              isAnimationActive={false}
            >
              <LabelList
                dataKey="failureRate"
                position="insideRight"
                fill="#ffffff"
                formatter={percentageLabel}
              />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      </div>
    </article>
  );
}

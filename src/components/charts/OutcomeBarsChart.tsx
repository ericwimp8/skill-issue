import {
  Bar,
  BarChart,
  CartesianGrid,
  Cell,
  LabelList,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
  type TooltipContentProps,
} from 'recharts';

import {
  scenarioOptions,
  type EvaluationResult,
} from '../../data/evaluationData';
import { chartColorForCell, resultCalled, resultTotal } from './chartTheme';

type OutcomeBarsChartProps = {
  onToggleScenario: (scenarioId: string) => void;
  results: EvaluationResult[];
  selectedScenarioIds: string[];
};

type OutcomeDatum = {
  called: number;
  cellId: string;
  color: string;
  failed: number;
  failureRate: number;
  label: string;
  scenarioCount: number;
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
      <span>
        {result.total} expected calls across {result.scenarioCount} scenario
        {result.scenarioCount === 1 ? '' : 's'}
      </span>
    </div>
  );
}

function percentageLabel(value: unknown) {
  const percentage = Number(value);
  return percentage >= 10 ? `${percentage.toFixed(0)}%` : '';
}

function aggregateResults(results: EvaluationResult[]) {
  const byCell = new Map<string, OutcomeDatum>();

  results.forEach((result) => {
    const called = resultCalled(result);
    const total = resultTotal(result);
    const current = byCell.get(result.cellId);

    if (current) {
      current.called += called;
      current.total += total;
      current.scenarioCount += 1;
      return;
    }

    byCell.set(result.cellId, {
      called,
      cellId: result.cellId,
      color: chartColorForCell(result.cellId),
      failed: 0,
      failureRate: 0,
      label: result.cellLabel,
      scenarioCount: 1,
      successRate: 0,
      total,
    });
  });

  return [...byCell.values()]
    .map((result) => {
      const failed = result.total - result.called;

      return {
        ...result,
        failed,
        failureRate: result.total === 0 ? 0 : (failed / result.total) * 100,
        successRate:
          result.total === 0 ? 0 : (result.called / result.total) * 100,
      };
    })
    .sort(
      (left, right) =>
        right.successRate - left.successRate ||
        left.label.localeCompare(right.label),
    );
}

export function OutcomeBarsChart({
  onToggleScenario,
  results,
  selectedScenarioIds,
}: OutcomeBarsChartProps) {
  const data = aggregateResults(results);
  const chartHeight = Math.max(320, data.length * 48 + 72);

  return (
    <article className="exploration-chart">
      <header className="exploration-chart-header">
        <div>
          <span className="chart-number">03</span>
          <p className="card-kicker">Success ranking</p>
          <h3>Which setups call skills reliably?</h3>
        </div>
        <details className="chart-scenario-picker">
          <summary>
            Scenarios · {selectedScenarioIds.length} of {scenarioOptions.length}
          </summary>
          <div className="chart-scenario-panel">
            {scenarioOptions.map((scenario) => {
              const checked = selectedScenarioIds.includes(scenario.id);

              return (
                <label key={scenario.id}>
                  <input
                    type="checkbox"
                    checked={checked}
                    disabled={checked && selectedScenarioIds.length === 1}
                    onChange={() => onToggleScenario(scenario.id)}
                  />
                  <span>{scenario.label}</span>
                </label>
              );
            })}
          </div>
        </details>
      </header>
      <p className="chart-description-wide">
        Every expected call in the selected scenarios contributes to one success
        or failure rate. Bars rank the selected setups from strongest to
        weakest.
      </p>
      <div className="raster-legend" aria-label="Success ranking legend">
        <span>
          <i className="legend-block legend-block-model" /> Model color =
          success
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
              width={230}
              interval={0}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <Tooltip
              content={OutcomeTooltip}
              isAnimationActive={false}
              wrapperStyle={{ transition: 'none' }}
            />
            <Bar
              dataKey="successRate"
              name="Success"
              stackId="outcome"
              radius={[5, 0, 0, 5]}
              isAnimationActive={false}
            >
              {data.map((result) => (
                <Cell key={result.cellId} fill={result.color} />
              ))}
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
              fill="var(--color-neutral-failure)"
              radius={[0, 5, 5, 0]}
              isAnimationActive={false}
            >
              <LabelList
                dataKey="failureRate"
                position="insideRight"
                fill="var(--color-text)"
                formatter={percentageLabel}
              />
            </Bar>
          </BarChart>
        </ResponsiveContainer>
      </div>
    </article>
  );
}

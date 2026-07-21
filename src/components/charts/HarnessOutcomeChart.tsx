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

import {
  scenarioOptions,
  type EvaluationResult,
} from '../../data/evaluationData';
import { resultCalled, resultTotal } from './chartTheme';

type HarnessOutcomeChartProps = {
  onToggleScenario: (scenarioId: string) => void;
  results: EvaluationResult[];
  selectedScenarioIds: string[];
};

type HarnessDatum = {
  called: number;
  failed: number;
  failureRate: number;
  harness: string;
  label: string;
  successRate: number;
  total: number;
};

function aggregateHarnesses(results: EvaluationResult[]) {
  const aggregates = new Map<
    string,
    Pick<HarnessDatum, 'called' | 'harness' | 'label' | 'total'>
  >();

  results.forEach((result) => {
    const current = aggregates.get(result.harness);
    const called = resultCalled(result);
    const total = resultTotal(result);

    if (current) {
      current.called += called;
      current.total += total;
      return;
    }

    aggregates.set(result.harness, {
      called,
      harness: result.harness,
      label: result.harnessLabel,
      total,
    });
  });

  return [...aggregates.values()]
    .map((result): HarnessDatum => {
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

function percentageLabel(value: unknown) {
  const percentage = Number(value);
  return percentage >= 10 ? `${percentage.toFixed(0)}%` : '';
}

function HarnessTooltip({ active, payload }: TooltipContentProps) {
  const result = payload?.[0]?.payload as HarnessDatum | undefined;

  if (!active || !result) {
    return null;
  }

  return (
    <div className="chart-tooltip">
      <strong>{result.label}</strong>
      <span>{result.successRate.toFixed(0)}% successful</span>
      <span>
        {result.called} called · {result.failed} missed · {result.total} scored
      </span>
    </div>
  );
}

export function HarnessOutcomeChart({
  onToggleScenario,
  results,
  selectedScenarioIds,
}: HarnessOutcomeChartProps) {
  const data = aggregateHarnesses(results);
  const conversationTurns = selectedScenarioIds.length * 30;
  const chartHeight = Math.max(390, data.length * 58 + 72);

  return (
    <article className="exploration-chart harness-comparison-chart">
      <header className="exploration-chart-header">
        <div>
          <p className="card-kicker">Harness comparison</p>
          <h3>How much does the harness change skill calling?</h3>
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
        Codex Sol and Medium reasoning stay fixed while the harness changes. The
        selected evidence represents {conversationTurns} conversation turns and{' '}
        {selectedScenarioIds.length * 4} scored expected calls per harness.
      </p>
      <div className="harness-summary" aria-label="Harness chart context">
        <span>
          <strong>{selectedScenarioIds.length}</strong> scenario
          {selectedScenarioIds.length === 1 ? '' : 's'}
        </span>
        <span>
          <strong>{conversationTurns}</strong> conversation turns
        </span>
        <span>
          <strong>{selectedScenarioIds.length * 4}</strong> scored calls per
          harness
        </span>
      </div>
      <div className="raster-legend" aria-label="Harness ranking legend">
        <span>
          <i className="legend-block legend-block-harness" /> Successful
        </span>
        <span>
          <i className="legend-block legend-block-missed" /> Missed
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
              width={132}
              interval={0}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <Tooltip
              content={HarnessTooltip}
              isAnimationActive={false}
              wrapperStyle={{ transition: 'none' }}
            />
            <Bar
              dataKey="successRate"
              name="Successful"
              stackId="outcome"
              fill="var(--color-harness-success)"
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
              name="Missed"
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

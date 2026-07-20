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
import { SeriesMarkerDot, SeriesMarkerIcon } from './SeriesMarker';
import {
  chartColorForCell,
  chartMarkerForCell,
  tooltipStyle,
} from './chartTheme';

type CumulativeCallsChartProps = {
  results: EvaluationResult[];
};

function callOutcomes(result: EvaluationResult) {
  return [...result.points]
    .sort((left, right) => left.turn - right.turn)
    .flatMap((point) => [
      ...Array.from({ length: point.called }, () => 1),
      ...Array.from({ length: point.missed }, () => 0),
    ]);
}

export function CumulativeCallsChart({ results }: CumulativeCallsChartProps) {
  const expectedCallCount = Math.max(
    0,
    ...results.map((result) => callOutcomes(result).length),
  );
  const data = Array.from({ length: expectedCallCount }, (_, index) => {
    const expectedCall = index + 1;
    const row: Record<string, number> = { expectedCall };

    results.forEach((result) => {
      row[result.cellId] = callOutcomes(result)
        .slice(0, expectedCall)
        .reduce((total, called) => total + called, 0);
    });

    return row;
  });

  return (
    <article className="exploration-chart">
      <header className="exploration-chart-header">
        <div>
          <span className="chart-number">02</span>
          <p className="card-kicker">Expected-call curve</p>
          <h3>Where each model calls or misses.</h3>
        </div>
        <span className="chart-purpose">Timing</span>
      </header>
      <p className="chart-description-wide">
        Each point is one expected skill call. A successful call moves the line
        up; a miss leaves it flat.
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
              dataKey="expectedCall"
              type="number"
              domain={[1, expectedCallCount]}
              ticks={Array.from(
                { length: expectedCallCount },
                (_, index) => index + 1,
              )}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <YAxis
              domain={[0, expectedCallCount]}
              ticks={Array.from(
                { length: expectedCallCount + 1 },
                (_, index) => index,
              )}
              allowDecimals={false}
              stroke="var(--color-chart-label)"
              tickLine={false}
              axisLine={false}
            />
            <Tooltip
              contentStyle={tooltipStyle}
              isAnimationActive={false}
              labelFormatter={(expectedCall) => `Expected call ${expectedCall}`}
              wrapperStyle={{ transition: 'none' }}
            />
            {results.map((result) => {
              const color = chartColorForCell(result.cellId);
              const shape = chartMarkerForCell(result.cellId);

              return (
                <Line
                  key={result.cellId}
                  type="linear"
                  dataKey={result.cellId}
                  name={result.cellLabel}
                  stroke={color}
                  strokeWidth={2.5}
                  dot={(props) => (
                    <SeriesMarkerDot {...props} color={color} shape={shape} />
                  )}
                  activeDot={(props) => (
                    <SeriesMarkerDot
                      {...props}
                      color={color}
                      shape={shape}
                      size={20}
                    />
                  )}
                  isAnimationActive={false}
                />
              );
            })}
          </LineChart>
        </ResponsiveContainer>
      </div>
      <div className="series-key" aria-label="Selected comparison series">
        {results.map((result) => (
          <span key={result.cellId}>
            <SeriesMarkerIcon
              color={chartColorForCell(result.cellId)}
              shape={chartMarkerForCell(result.cellId)}
            />
            {result.cellLabel}
          </span>
        ))}
      </div>
    </article>
  );
}

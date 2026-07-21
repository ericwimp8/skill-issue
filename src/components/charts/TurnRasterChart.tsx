import {
  Bar,
  BarChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
  type BarShapeProps,
  type TooltipContentProps,
} from 'recharts';

import type {
  EvaluationResult,
  WebsiteEvaluationPoint,
} from '../../data/evaluationData';
import { chartColorForCell } from './chartTheme';

type TurnRasterChartProps = {
  results: EvaluationResult[];
};

type Outcome = 'called' | 'missed';

type OutcomeTurn = {
  outcomes: Outcome[];
  turn: number;
};

type OutcomeStripDatum = {
  called: number;
  cellId: string;
  color: string;
  expected: number;
  failureRate: number;
  finalTurn?: number;
  firstTurn?: number;
  label: string;
  missed: number;
  scenarioLabel: string;
  scoredTurns: number;
  stripWidth: number;
  successRate: number;
  turns: OutcomeTurn[];
};

function outcomesForPoint(point: WebsiteEvaluationPoint): Outcome[] {
  return [
    ...Array.from({ length: point.called }, () => 'called' as const),
    ...Array.from({ length: point.missed }, () => 'missed' as const),
  ];
}

function toOutcomeStripDatum(result: EvaluationResult): OutcomeStripDatum {
  const points = [...result.points].sort(
    (left, right) => left.turn - right.turn,
  );
  const called = points.reduce((total, point) => total + point.called, 0);
  const missed = points.reduce((total, point) => total + point.missed, 0);
  const expected = called + missed;

  return {
    called,
    cellId: result.cellId,
    color: chartColorForCell(result.cellId),
    expected,
    failureRate: expected === 0 ? 0 : (missed / expected) * 100,
    finalTurn: points.at(-1)?.turn,
    firstTurn: points.at(0)?.turn,
    label: result.cellLabel,
    missed,
    scenarioLabel: result.scenarioLabel,
    scoredTurns: points.length,
    stripWidth: 100,
    successRate: expected === 0 ? 0 : (called / expected) * 100,
    turns: points.map((point) => ({
      outcomes: outcomesForPoint(point),
      turn: point.turn,
    })),
  };
}

function clipPathId(cellId: string) {
  return `outcome-strip-${cellId.replace(/[^a-zA-Z0-9_-]/g, '-')}`;
}

function SegmentedOutcomeBar({
  height,
  isActive,
  payload,
  width,
  x,
  y,
}: BarShapeProps) {
  const result = payload as OutcomeStripDatum;
  const turnWidth = result.turns.length === 0 ? 0 : width / result.turns.length;
  const pathId = clipPathId(result.cellId);

  return (
    <g className="outcome-bar-shape">
      <defs>
        <clipPath id={pathId}>
          <rect x={x} y={y} width={width} height={height} rx={5} />
        </clipPath>
      </defs>
      <g clipPath={`url(#${pathId})`} pointerEvents="none">
        {result.turns.flatMap((turn, turnIndex) => {
          const outcomeWidth =
            turn.outcomes.length === 0
              ? turnWidth
              : turnWidth / turn.outcomes.length;

          return turn.outcomes.map((outcome, outcomeIndex) => (
            <rect
              key={`${turn.turn}-${outcome}-${outcomeIndex}`}
              x={x + turnIndex * turnWidth + outcomeIndex * outcomeWidth}
              y={y}
              width={outcomeWidth}
              height={height}
              fill={
                outcome === 'called'
                  ? result.color
                  : 'var(--color-neutral-failure)'
              }
            />
          ));
        })}
      </g>
      <rect
        x={x}
        y={y}
        width={width}
        height={height}
        rx={5}
        fill="transparent"
        stroke={isActive ? 'var(--color-text-muted)' : 'transparent'}
        strokeWidth={1}
      />
      <text
        className="outcome-bar-range-label"
        x={x}
        y={y + height + 16}
        textAnchor="start"
        pointerEvents="none"
      >
        {result.firstTurn === undefined ? '' : `Turn ${result.firstTurn}`}
      </text>
      <text
        className="outcome-bar-range-label"
        x={x + width}
        y={y + height + 16}
        textAnchor="end"
        pointerEvents="none"
      >
        {result.finalTurn === undefined ? '' : `Turn ${result.finalTurn}`}
      </text>
    </g>
  );
}

function OutcomeStripTooltip({ active, payload }: TooltipContentProps) {
  const result = payload?.[0]?.payload as OutcomeStripDatum | undefined;

  if (!active || !result) {
    return null;
  }

  return (
    <div className="chart-tooltip">
      <strong>{result.label}</strong>
      <span>{result.scenarioLabel}</span>
      <span>
        {result.successRate.toFixed(0)}% success · {result.called} called
      </span>
      <span>
        {result.failureRate.toFixed(0)}% failure · {result.missed} missed
      </span>
      <span>
        {result.expected} expected calls across {result.scoredTurns} scored turn
        {result.scoredTurns === 1 ? '' : 's'}
      </span>
    </div>
  );
}

export function TurnRasterChart({ results }: TurnRasterChartProps) {
  const data = results.map(toOutcomeStripDatum);
  const chartHeight = Math.max(320, data.length * 72 + 36);

  return (
    <article className="exploration-chart exploration-chart-featured">
      <header className="exploration-chart-header">
        <div>
          <span className="chart-number">01</span>
          <p className="card-kicker">Outcome strips</p>
          <h3>Skill-call reliability over time.</h3>
        </div>
        <span className="chart-purpose">Direct comparison</span>
      </header>
      <p className="chart-description-wide">
        Each strip contains only scored turns, ordered from earliest to latest.
        More gray toward the right reveals reliability falling later in the
        conversation.
      </p>
      <div
        className="chart-canvas outcome-chart-canvas"
        style={{ height: chartHeight }}
      >
        <ResponsiveContainer width="100%" height="100%">
          <BarChart
            data={data}
            layout="vertical"
            margin={{ top: 4, right: 20, bottom: 8, left: 10 }}
            barCategoryGap="42%"
          >
            <XAxis type="number" domain={[0, 100]} hide />
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
              content={OutcomeStripTooltip}
              cursor={false}
              isAnimationActive={false}
              wrapperStyle={{ transition: 'none' }}
            />
            <Bar
              dataKey="stripWidth"
              name="Outcome"
              shape={SegmentedOutcomeBar}
              activeBar={SegmentedOutcomeBar}
              barSize={38}
              isAnimationActive={false}
            />
          </BarChart>
        </ResponsiveContainer>
      </div>
    </article>
  );
}

import type {
  EvaluationResult,
  WebsiteEvaluationPoint,
} from '../../data/evaluationData';

type TurnRasterChartProps = {
  results: EvaluationResult[];
};

type Outcome = 'called' | 'missed';

function outcomesForPoint(point: WebsiteEvaluationPoint): Outcome[] {
  return [
    ...Array.from({ length: point.called }, () => 'called' as const),
    ...Array.from({ length: point.missed }, () => 'missed' as const),
  ];
}

function outcomeSummary(point: WebsiteEvaluationPoint) {
  return `Turn ${point.turn}: ${point.called} called, ${point.missed} missed`;
}

export function TurnRasterChart({ results }: TurnRasterChartProps) {
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
        More red toward the right reveals reliability falling later in the
        conversation.
      </p>
      <div className="raster-legend" aria-label="Outcome strip legend">
        <span>
          <i className="legend-block legend-block-called" /> Called
        </span>
        <span>
          <i className="legend-block legend-block-missed" /> Missed
        </span>
      </div>
      <div className="outcome-strips">
        {results.map((result) => {
          const points = [...result.points].sort(
            (left, right) => left.turn - right.turn,
          );
          const firstTurn = points.at(0)?.turn;
          const finalTurn = points.at(-1)?.turn;

          return (
            <section className="outcome-strip-row" key={result.cellId}>
              <div className="outcome-strip-heading">
                <strong>{result.cellLabel}</strong>
                <span>{points.length} scored turns</span>
              </div>
              <div
                className="outcome-strip"
                aria-label={`${result.cellLabel} scored outcomes`}
              >
                {points.map((point) => {
                  const outcomes = outcomesForPoint(point);
                  const summary = outcomeSummary(point);

                  return (
                    <div
                      className="outcome-turn"
                      key={point.turn_id}
                      title={summary}
                      aria-label={summary}
                    >
                      {outcomes.map((outcome, index) => (
                        <span
                          className={`outcome-call outcome-call-${outcome}`}
                          key={`${outcome}-${index}`}
                        />
                      ))}
                    </div>
                  );
                })}
              </div>
              <div className="outcome-strip-range" aria-hidden="true">
                <span>
                  {firstTurn === undefined ? '' : `Turn ${firstTurn}`}
                </span>
                <span>
                  {finalTurn === undefined ? '' : `Turn ${finalTurn}`}
                </span>
              </div>
            </section>
          );
        })}
      </div>
    </article>
  );
}

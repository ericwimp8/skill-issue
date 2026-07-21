import {
  evaluationResults,
  scenarioOptions,
  type EvaluationResult,
} from './evaluationData';

export type AnalysisSummary = {
  additional: number;
  called: number;
  conversationTurns: number;
  expected: number;
  missed: number;
  runs: number;
  successRate: number;
};

function summarize(results: readonly EvaluationResult[]): AnalysisSummary {
  const totals = results.reduce(
    (summary, result) => {
      summary.conversationTurns += result.total_turns;
      summary.runs += 1;

      result.points.forEach((point) => {
        summary.additional += point.unexpected;
        summary.called += point.called;
        summary.missed += point.missed;
      });

      return summary;
    },
    {
      additional: 0,
      called: 0,
      conversationTurns: 0,
      missed: 0,
      runs: 0,
    },
  );
  const expected = totals.called + totals.missed;

  return {
    ...totals,
    expected,
    successRate: expected === 0 ? 0 : (totals.called / expected) * 100,
  };
}

function resultsForCell(cellId: string) {
  return evaluationResults.filter((result) => result.cellId === cellId);
}

function completeCellSummary(cellId: string) {
  const results = resultsForCell(cellId);
  const scenarioCount = new Set(results.map((result) => result.scenario_id))
    .size;

  if (scenarioCount !== scenarioOptions.length) {
    throw new Error(`Analysis requires a complete scenario set for ${cellId}`);
  }

  return summarize(results);
}

function summarizeTurnBand(
  cellId: string,
  firstTurn: number,
  lastTurn: number,
) {
  const results = resultsForCell(cellId).map((result) => ({
    ...result,
    points: result.points.filter(
      (point) => point.turn >= firstTurn && point.turn <= lastTurn,
    ),
  }));

  return summarize(results);
}

export const campaignSummary = summarize(evaluationResults);

export const analysisConfigurations = {
  claudeCodeFable: completeCellSummary('claude-code::claude-fable'),
  codexSol: completeCellSummary('codex::codex-sol'),
  cursorComposer: completeCellSummary('cursor::composer'),
  cursorGrok: completeCellSummary('cursor::grok'),
  openCodeSol: completeCellSummary('opencode::codex-sol'),
  piSol: completeCellSummary('pi::codex-sol'),
} as const;

export const analysisTurnBands = {
  claudeCodeFable: {
    first: summarizeTurnBand('claude-code::claude-fable', 1, 10),
    last: summarizeTurnBand('claude-code::claude-fable', 21, 30),
  },
  cursorComposer: {
    first: summarizeTurnBand('cursor::composer', 1, 10),
    last: summarizeTurnBand('cursor::composer', 21, 30),
  },
  openCodeSol: {
    first: summarizeTurnBand('opencode::codex-sol', 1, 10),
    last: summarizeTurnBand('opencode::codex-sol', 21, 30),
  },
  piSol: {
    first: summarizeTurnBand('pi::codex-sol', 1, 10),
    last: summarizeTurnBand('pi::codex-sol', 21, 30),
  },
} as const;

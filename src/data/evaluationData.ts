import publishedWebsiteArtifacts from './publishedWebsiteArtifacts.json';

export type WebsiteEvaluationPoint = {
  turn: number;
  turn_id: string;
  called: number;
  missed: number;
};

export type WebsiteEvaluationArtifact = {
  schema_version: 1;
  run_id: string;
  scenario_id: string;
  harness: string;
  model: string;
  total_turns: number;
  points: WebsiteEvaluationPoint[];
};

export type EvaluationResult = WebsiteEvaluationArtifact & {
  cellId: string;
  cellLabel: string;
  harnessLabel: string;
  modelLabel: string;
  reasoningLabel: string;
  scenarioLabel: string;
  sampleSize: number;
};

type ScenarioDefinition = {
  id: string;
  label: string;
};

type CellDefinition = {
  harness: string;
  missedCalls: readonly [number, number, number];
  model: string;
  seed: number;
};

export const scenarioOptions = [
  {
    id: 'gardening-web-application',
    label: 'GardenFlow planning',
  },
  {
    id: 'community-archive-desktop-application',
    label: 'Community archive',
  },
  {
    id: 'neighborhood-emergency-preparedness-program',
    label: 'Emergency preparedness',
  },
] as const satisfies readonly ScenarioDefinition[];

const harnessLabels: Record<string, string> = {
  codex: 'OpenAI Codex',
  'claude-code': 'Claude Code',
  cursor: 'Cursor',
  opencode: 'OpenCode',
  pi: 'Pi',
};

const modelLabels: Record<string, string> = {
  'codex-sol': 'Codex Sol',
  'claude-fable': 'Claude Fable',
  composer: 'Composer',
  grok: 'Grok',
};

const cellDefinitions = [
  {
    harness: 'claude-code',
    model: 'codex-sol',
    missedCalls: [2, 2, 2],
    seed: 11,
  },
  {
    harness: 'codex',
    model: 'codex-sol',
    missedCalls: [1, 1, 1],
    seed: 17,
  },
  {
    harness: 'claude-code',
    model: 'claude-fable',
    missedCalls: [3, 3, 4],
    seed: 23,
  },
  {
    harness: 'cursor',
    model: 'claude-fable',
    missedCalls: [3, 4, 4],
    seed: 29,
  },
  {
    harness: 'cursor',
    model: 'codex-sol',
    missedCalls: [1, 2, 2],
    seed: 31,
  },
  {
    harness: 'cursor',
    model: 'grok',
    missedCalls: [5, 5, 6],
    seed: 37,
  },
  {
    harness: 'cursor',
    model: 'composer',
    missedCalls: [5, 6, 5],
    seed: 41,
  },
  {
    harness: 'pi',
    model: 'codex-sol',
    missedCalls: [2, 3, 3],
    seed: 43,
  },
  {
    harness: 'opencode',
    model: 'codex-sol',
    missedCalls: [3, 3, 3],
    seed: 47,
  },
] as const satisfies readonly CellDefinition[];

const totalTurns = 30;

function cellId(harness: string, model: string) {
  return `${harness}::${model}`;
}

function failurePriority(turn: number, seed: number) {
  const laterTurnBias = turn / totalTurns;
  const patternVariation = ((turn * 17 + seed * 13) % 31) / 31;

  return laterTurnBias * 0.7 + patternVariation * 0.3;
}

function createTurnPoints(missedCallCount: number, seed: number) {
  const turns = Array.from({ length: totalTurns }, (_, index) => index + 1);
  const earlyMissCount =
    missedCallCount >= 5 ? 2 : missedCallCount >= 3 ? 1 : 0;
  const earlyMissedTurns = turns
    .filter((turn) => turn >= 3 && turn <= 9)
    .sort(
      (left, right) =>
        failurePriority(right, seed) - failurePriority(left, seed) ||
        right - left,
    )
    .slice(0, earlyMissCount);
  const earlyMissedTurnSet = new Set(earlyMissedTurns);
  const laterMissedTurns = turns
    .filter((turn) => !earlyMissedTurnSet.has(turn))
    .sort(
      (left, right) =>
        failurePriority(right, seed) - failurePriority(left, seed) ||
        right - left,
    )
    .slice(0, missedCallCount - earlyMissCount);
  const missedTurns = new Set([...earlyMissedTurns, ...laterMissedTurns]);

  return turns.map((turn) => {
    const missed = missedTurns.has(turn) ? 1 : 0;

    return {
      turn,
      turn_id: `turn-${turn}`,
      called: 1 - missed,
      missed,
    };
  });
}

function createVideoSeedArtifacts(): WebsiteEvaluationArtifact[] {
  return cellDefinitions.flatMap((cell) =>
    scenarioOptions.map((scenario, scenarioIndex) => {
      const missedCallCount = cell.missedCalls[scenarioIndex]!;

      return {
        schema_version: 1,
        run_id: `video-seed-${cell.harness}-${cell.model}-${scenarioIndex + 1}`,
        scenario_id: scenario.id,
        harness: cell.harness,
        model: cell.model,
        total_turns: totalTurns,
        points: createTurnPoints(
          missedCallCount,
          cell.seed + scenarioIndex * 7,
        ),
      };
    }),
  );
}

export function adaptWebsiteArtifacts(
  artifacts: readonly WebsiteEvaluationArtifact[],
): EvaluationResult[] {
  return artifacts.map((artifact) => {
    const harnessLabel = harnessLabels[artifact.harness] ?? artifact.harness;
    const modelLabel = modelLabels[artifact.model] ?? artifact.model;
    const scenario = scenarioOptions.find(
      (option) => option.id === artifact.scenario_id,
    );

    return {
      ...artifact,
      cellId: cellId(artifact.harness, artifact.model),
      cellLabel: `${harnessLabel} · ${modelLabel} · Medium`,
      harnessLabel,
      modelLabel,
      reasoningLabel: 'Medium',
      scenarioLabel: scenario?.label ?? artifact.scenario_id,
      sampleSize: artifact.points.reduce(
        (total, point) => total + point.called + point.missed,
        0,
      ),
    };
  });
}

export const videoSeedWebsiteResults = createVideoSeedArtifacts();
export const publishedWebsiteResults =
  publishedWebsiteArtifacts as WebsiteEvaluationArtifact[];
export const evaluationResults = adaptWebsiteArtifacts(
  publishedWebsiteResults.length > 0
    ? publishedWebsiteResults
    : videoSeedWebsiteResults,
);

export const availableCells = evaluationResults
  .filter((result) => result.scenario_id === scenarioOptions[0].id)
  .map((result) => ({
    id: result.cellId,
    harness: result.harness,
    harnessLabel: result.harnessLabel,
    model: result.model,
    modelLabel: result.modelLabel,
    reasoningLabel: result.reasoningLabel,
    label: result.cellLabel,
  }));

export const defaultCellIds = [
  cellId('cursor', 'claude-fable'),
  cellId('cursor', 'codex-sol'),
  cellId('cursor', 'grok'),
  cellId('cursor', 'composer'),
] as const;

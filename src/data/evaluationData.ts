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
  model: string;
  outcomes: readonly [
    readonly [number, number, number, number],
    readonly [number, number, number, number],
    readonly [number, number, number, number],
  ];
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
  kilo: 'Kilo Code',
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
    outcomes: [
      [1, 1, 1, 0],
      [1, 1, 1, 1],
      [1, 0, 1, 1],
    ],
  },
  {
    harness: 'codex',
    model: 'codex-sol',
    outcomes: [
      [1, 1, 1, 1],
      [1, 1, 0, 1],
      [1, 1, 1, 1],
    ],
  },
  {
    harness: 'claude-code',
    model: 'claude-fable',
    outcomes: [
      [1, 1, 1, 1],
      [1, 0, 1, 0],
      [1, 1, 0, 1],
    ],
  },
  {
    harness: 'cursor',
    model: 'claude-fable',
    outcomes: [
      [1, 1, 1, 1],
      [1, 0, 1, 1],
      [1, 1, 0, 0],
    ],
  },
  {
    harness: 'cursor',
    model: 'codex-sol',
    outcomes: [
      [1, 1, 0, 1],
      [1, 1, 1, 0],
      [1, 1, 1, 1],
    ],
  },
  {
    harness: 'cursor',
    model: 'grok',
    outcomes: [
      [1, 0, 1, 0],
      [0, 1, 1, 0],
      [1, 1, 0, 0],
    ],
  },
  {
    harness: 'cursor',
    model: 'composer',
    outcomes: [
      [1, 1, 1, 1],
      [1, 1, 1, 0],
      [1, 1, 1, 1],
    ],
  },
  {
    harness: 'pi',
    model: 'codex-sol',
    outcomes: [
      [1, 1, 1, 0],
      [1, 1, 1, 1],
      [1, 1, 1, 0],
    ],
  },
  {
    harness: 'opencode',
    model: 'codex-sol',
    outcomes: [
      [1, 1, 1, 0],
      [1, 1, 0, 0],
      [1, 1, 1, 0],
    ],
  },
  {
    harness: 'kilo',
    model: 'codex-sol',
    outcomes: [
      [1, 1, 0, 1],
      [1, 0, 1, 1],
      [1, 1, 0, 1],
    ],
  },
] as const satisfies readonly CellDefinition[];

const scoredTurns = [1, 11, 25, 30] as const;
const expectedByTurn = [1, 1, 1, 1] as const;

function cellId(harness: string, model: string) {
  return `${harness}::${model}`;
}

function createIllustrativeArtifacts(): WebsiteEvaluationArtifact[] {
  return cellDefinitions.flatMap((cell) =>
    scenarioOptions.map((scenario, scenarioIndex) => {
      const outcomes = cell.outcomes[scenarioIndex]!;

      return {
        schema_version: 1,
        run_id: `illustrative-${cell.harness}-${cell.model}-${scenarioIndex + 1}`,
        scenario_id: scenario.id,
        harness: cell.harness,
        model: cell.model,
        total_turns: 30,
        points: scoredTurns.map((turn, pointIndex) => {
          const called = outcomes[pointIndex]!;
          const expected = expectedByTurn[pointIndex]!;

          return {
            turn,
            turn_id: `turn-${turn}`,
            called,
            missed: expected - called,
          };
        }),
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

export const illustrativeWebsiteResults = createIllustrativeArtifacts();
export const evaluationResults = adaptWebsiteArtifacts(
  illustrativeWebsiteResults,
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

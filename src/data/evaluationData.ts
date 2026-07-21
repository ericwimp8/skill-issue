import publishedWebsiteArtifacts from './publishedWebsiteArtifacts.json';

export type WebsiteEvaluationPoint = {
  turn: number;
  turn_id: string;
  called: number;
  missed: number;
  unexpected: number;
};

export type WebsiteEvaluationArtifact = {
  schema_version: 2;
  run_id: string;
  scenario_id: string;
  harness: string;
  model: string;
  total_turns: number;
  points: WebsiteEvaluationPoint[];
};

export type EvaluationResult = Omit<WebsiteEvaluationArtifact, 'model'> & {
  cellId: string;
  cellLabel: string;
  harnessLabel: string;
  model: string;
  modelLabel: string;
  reasoningLabel: string;
  sampleSize: number;
  scenarioLabel: string;
  sourceModel: string;
};

type ScenarioDefinition = {
  id: string;
  label: string;
};

type CellDefinition = {
  artifactModels: readonly string[];
  available: boolean;
  harness: string;
  model: string;
  reasoningLabel: string;
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
    artifactModels: ['gpt-5.6-sol'],
    available: true,
    harness: 'claude-code',
    model: 'codex-sol',
    reasoningLabel: 'Medium',
  },
  {
    artifactModels: ['gpt-5.6-sol'],
    available: true,
    harness: 'codex',
    model: 'codex-sol',
    reasoningLabel: 'Medium',
  },
  {
    artifactModels: ['claude-fable-5'],
    available: true,
    harness: 'claude-code',
    model: 'claude-fable',
    reasoningLabel: 'Medium',
  },
  {
    artifactModels: ['claude-fable-5-thinking-high'],
    available: false,
    harness: 'cursor',
    model: 'claude-fable',
    reasoningLabel: 'High',
  },
  {
    artifactModels: ['gpt-5.6-sol-high'],
    available: false,
    harness: 'cursor',
    model: 'codex-sol',
    reasoningLabel: 'High',
  },
  {
    artifactModels: ['cursor-grok-4.5-medium'],
    available: true,
    harness: 'cursor',
    model: 'grok',
    reasoningLabel: 'Medium',
  },
  {
    artifactModels: ['composer-2.5'],
    available: true,
    harness: 'cursor',
    model: 'composer',
    reasoningLabel: 'Medium',
  },
  {
    artifactModels: ['openai-codex/gpt-5.6-sol'],
    available: true,
    harness: 'pi',
    model: 'codex-sol',
    reasoningLabel: 'Medium',
  },
  {
    artifactModels: ['openai/gpt-5.6-sol'],
    available: true,
    harness: 'opencode',
    model: 'codex-sol',
    reasoningLabel: 'Medium',
  },
] as const satisfies readonly CellDefinition[];

function cellId(harness: string, model: string) {
  return `${harness}::${model}`;
}

function definitionForArtifact(artifact: WebsiteEvaluationArtifact) {
  const definition = cellDefinitions.find(
    (cell) =>
      cell.harness === artifact.harness &&
      (cell.artifactModels as readonly string[]).includes(artifact.model),
  );

  if (!definition) {
    throw new Error(
      `No website cell maps ${artifact.harness} with ${artifact.model}`,
    );
  }

  return definition;
}

export function adaptWebsiteArtifacts(
  artifacts: readonly WebsiteEvaluationArtifact[],
): EvaluationResult[] {
  return artifacts.map((artifact) => {
    const definition = definitionForArtifact(artifact);
    const harnessLabel =
      harnessLabels[definition.harness] ?? definition.harness;
    const modelLabel = modelLabels[definition.model] ?? definition.model;
    const scenario = scenarioOptions.find(
      (option) => option.id === artifact.scenario_id,
    );

    if (!scenario) {
      throw new Error(`No website scenario maps ${artifact.scenario_id}`);
    }

    return {
      ...artifact,
      cellId: cellId(definition.harness, definition.model),
      cellLabel: `${harnessLabel} · ${modelLabel} · ${definition.reasoningLabel}`,
      harnessLabel,
      model: definition.model,
      modelLabel,
      reasoningLabel: definition.reasoningLabel,
      sampleSize: artifact.points.reduce(
        (total, point) =>
          total + point.called + point.missed + point.unexpected,
        0,
      ),
      scenarioLabel: scenario.label,
      sourceModel: artifact.model,
    };
  });
}

export const publishedWebsiteResults =
  publishedWebsiteArtifacts as WebsiteEvaluationArtifact[];
export const defaultCellIds = cellDefinitions
  .filter((cell) => cell.available)
  .map((cell) => cellId(cell.harness, cell.model));
const availableCellIds = new Set(defaultCellIds);
export const evaluationResults = adaptWebsiteArtifacts(
  publishedWebsiteResults,
).filter((result) => availableCellIds.has(result.cellId));

export const availableCells = cellDefinitions.map((cell) => {
  const harnessLabel = harnessLabels[cell.harness] ?? cell.harness;
  const modelLabel = modelLabels[cell.model] ?? cell.model;

  return {
    available: cell.available,
    id: cellId(cell.harness, cell.model),
    harness: cell.harness,
    harnessLabel,
    model: cell.model,
    modelLabel,
    reasoningLabel: cell.reasoningLabel,
    label: `${harnessLabel} · ${modelLabel} · ${cell.reasoningLabel}`,
  };
});

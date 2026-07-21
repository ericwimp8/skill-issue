import { siteData } from './siteData';

type ScenarioTurn = {
  turn_id: string;
  prompt: string;
};

type ExpectedCall = {
  turn_id: string;
  skill: string;
};

type BuiltInEvaluation = {
  evaluation_id: string;
  scenario: {
    scenario_id: string;
    turns: ScenarioTurn[];
  };
  answer_sheet: {
    scenario_id: string;
    expected: ExpectedCall[];
  };
};

export type MethodologyScenarioTurn = ScenarioTurn & {
  expectedSkills: string[];
};

export type MethodologyScenario = {
  id: string;
  title: string;
  description: string;
  turns: MethodologyScenarioTurn[];
  expectedCallCount: number;
  sourceUrl: string;
};

export type MethodologySkill = {
  slug: string;
  title: string;
  description: string;
  content: string;
  expectedCallCount: number;
  sourceUrl: string;
  evidenceUrl: string;
};

const scenarioFiles = import.meta.glob(
  '../../evaluations/skill-calling/built-ins/*.json',
  { eager: true, import: 'default' },
) as Record<string, BuiltInEvaluation>;

const refinedSkillFiles = import.meta.glob(
  '../../evaluations/scenario-skill-refinement/*/skill/SKILL.md',
  { eager: true, import: 'default', query: '?raw' },
) as Record<string, string>;

const dictatePlanFiles = import.meta.glob(
  '../../supporting-skills/dictate-plan/SKILL.md',
  { eager: true, import: 'default', query: '?raw' },
) as Record<string, string>;

const repositoryTreeUrl = `${siteData.repositoryUrl}/tree/main`;

function readFrontmatterValue(content: string, key: string) {
  const frontmatter = content.match(/^---\n([\s\S]*?)\n---/u)?.[1];
  const line = frontmatter
    ?.split('\n')
    .find((candidate) => candidate.startsWith(`${key}:`));

  return line?.slice(key.length + 1).trim() ?? '';
}

function formatTitle(name: string) {
  return name
    .split('-')
    .map((word) => `${word.charAt(0).toUpperCase()}${word.slice(1)}`)
    .join(' ');
}

export const methodologyScenarios = Object.values(scenarioFiles)
  .filter(
    (evaluation) =>
      evaluation.evaluation_id in siteData.methodology.scenarioMetadata,
  )
  .map((evaluation) => {
    const metadata =
      siteData.methodology.scenarioMetadata[
        evaluation.evaluation_id as keyof typeof siteData.methodology.scenarioMetadata
      ];
    const expectedByTurn = new Map<string, string[]>();

    evaluation.answer_sheet.expected.forEach(({ turn_id, skill }) => {
      const skills = expectedByTurn.get(turn_id) ?? [];
      expectedByTurn.set(turn_id, [...skills, skill]);
    });

    return {
      id: evaluation.evaluation_id,
      title: metadata?.title ?? formatTitle(evaluation.evaluation_id),
      description: metadata?.description ?? '',
      turns: evaluation.scenario.turns.map((turn) => ({
        ...turn,
        expectedSkills: expectedByTurn.get(turn.turn_id) ?? [],
      })),
      expectedCallCount: evaluation.answer_sheet.expected.length,
      sourceUrl: `${repositoryTreeUrl}/evaluations/skill-calling/scenarios/${evaluation.evaluation_id}`,
    } satisfies MethodologyScenario;
  })
  .sort((left, right) => left.title.localeCompare(right.title));

const expectedCallsBySkill = methodologyScenarios
  .flatMap((scenario) => scenario.turns)
  .flatMap((turn) => turn.expectedSkills)
  .reduce((counts, skill) => {
    counts.set(skill, (counts.get(skill) ?? 0) + 1);
    return counts;
  }, new Map<string, number>());

const skillFiles = { ...refinedSkillFiles, ...dictatePlanFiles };

export const methodologySkills = Object.entries(skillFiles)
  .map(([path, content]) => {
    const refinedSlug = path.match(
      /scenario-skill-refinement\/([^/]+)\/skill\/SKILL\.md$/u,
    )?.[1];
    const slug = refinedSlug ?? 'dictate-plan';
    const name = readFrontmatterValue(content, 'name') || slug;
    const sourcePath = refinedSlug
      ? `evaluations/scenario-skill-refinement/${slug}/skill`
      : 'supporting-skills/dictate-plan';
    const evidencePath = refinedSlug
      ? `evaluations/scenario-skill-refinement/${slug}`
      : 'evaluations/skill-system-production-refinement/targets/dictate-plan';

    return {
      slug,
      title: formatTitle(name),
      description: readFrontmatterValue(content, 'description'),
      content,
      expectedCallCount: expectedCallsBySkill.get(slug) ?? 0,
      sourceUrl: `${repositoryTreeUrl}/${sourcePath}`,
      evidenceUrl: `${repositoryTreeUrl}/${evidencePath}`,
    } satisfies MethodologySkill;
  })
  .filter((skill) => skill.expectedCallCount > 0)
  .sort((left, right) => left.title.localeCompare(right.title));

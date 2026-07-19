export type EvaluationPoint = {
  contextConsumed: number;
  skillCalls: number;
  skillMisses: number;
};

export type EvaluationGraph = {
  id: string;
  harness: string;
  model: string;
  description: string;
  sampleSize: number;
  points: EvaluationPoint[];
};

export const siteData = {
  status: 'Preview benchmark · Local data',
  hero: {
    eyebrow: 'Skill reliability, made visible',
    title: 'Do your agent’s skills survive the context window?',
    description:
      'Skill Issue evaluates whether coding agents keep discovering and invoking the right skills as a task grows longer.',
  },
  release: {
    label: 'Download CLI',
    url: 'https://github.com/ericwimp8/skill-issue/releases/latest',
    note: 'CLI release coming soon. The button is ready for the first GitHub Release.',
  },
  repositoryUrl: 'https://github.com/ericwimp8/skill-issue',
  metrics: [
    { value: '240', label: 'evaluation prompts' },
    { value: '2', label: 'agentic CLI harnesses' },
    { value: '6', label: 'context checkpoints' },
  ],
  chart: {
    xAxis: 'Context consumed',
    yAxis: 'Skill decisions',
    calls: 'Skill calls',
    misses: 'Skill misses',
  },
  evaluations: [
    {
      id: 'codex-gpt-5-6',
      harness: 'Codex CLI',
      model: 'GPT-5.6',
      description:
        'Mock evaluation of explicit skill selection across a steadily growing task context.',
      sampleSize: 120,
      points: [
        { contextConsumed: 0, skillCalls: 19, skillMisses: 1 },
        { contextConsumed: 20, skillCalls: 18, skillMisses: 2 },
        { contextConsumed: 40, skillCalls: 17, skillMisses: 3 },
        { contextConsumed: 60, skillCalls: 16, skillMisses: 4 },
        { contextConsumed: 80, skillCalls: 14, skillMisses: 6 },
        { contextConsumed: 100, skillCalls: 12, skillMisses: 8 },
      ],
    },
    {
      id: 'claude-code-opus',
      harness: 'Claude Code',
      model: 'Claude Opus',
      description:
        'Mock comparison using the same skill-selection prompts and context checkpoints.',
      sampleSize: 120,
      points: [
        { contextConsumed: 0, skillCalls: 18, skillMisses: 2 },
        { contextConsumed: 20, skillCalls: 17, skillMisses: 3 },
        { contextConsumed: 40, skillCalls: 16, skillMisses: 4 },
        { contextConsumed: 60, skillCalls: 14, skillMisses: 6 },
        { contextConsumed: 80, skillCalls: 12, skillMisses: 8 },
        { contextConsumed: 100, skillCalls: 9, skillMisses: 11 },
      ],
    },
  ] satisfies EvaluationGraph[],
  method: {
    title: 'A simple signal for a difficult failure mode.',
    description:
      'At each checkpoint, an evaluation prompt either produces the required skill call or records a miss. The mock data shows the shape of the report while the evaluation harness is still being finalized.',
  },
  footer: 'Local-first evaluations. Static, reviewable benchmark data.',
} as const;

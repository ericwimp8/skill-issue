export const siteData = {
  status: 'Open-source · Local-first',
  hero: {
    eyebrow: 'Built with skills to build, evaluate, and fix skills.',
    title: 'It’s not a skill issue, but it’s always a skills issue.',
    description:
      'Create skills that work, then find out whether your model and harness can be trusted to call them. Skill Issue makes agent skills easier to create, evaluate, and trust.',
  },
  release: {
    label: 'Download CLI',
    url: 'https://github.com/ericwimp8/skill-issue/releases/latest',
  },
  repositoryUrl: 'https://github.com/ericwimp8/skill-issue',
  arms: {
    build: {
      label: 'Build Skills',
      shortLabel: 'Create, test, refine',
      eyebrow: 'Skill creation',
      title: 'Describe the outcome. Get a skill you can trust.',
      description:
        'Skill Issue turns an ordinary-language request into an idiomatic skill, evaluates when and how it runs, diagnoses failures, and refines it until the intended behavior is validated.',
    },
    evaluate: {
      label: 'Evaluate Environments',
      shortLabel: 'Compare model + harness',
      eyebrow: 'Environment evaluation',
      title: 'Do your agent’s skills survive the context window?',
      description:
        'Skill Issue evaluates whether coding agents keep discovering and invoking the right skills as a task grows longer.',
    },
  },
  buildWorkflow: [
    {
      index: '01',
      title: 'Describe the outcome',
      description:
        'Explain what the skill should achieve in ordinary language. The intake process inspects the project and resolves the ambiguities that matter.',
    },
    {
      index: '02',
      title: 'Generate and evaluate',
      description:
        'Skill Issue builds an idiomatic skill, then tests both reliable invocation and the behavior that follows once the skill is active.',
    },
    {
      index: '03',
      title: 'Diagnose and refine',
      description:
        'Failures are traced to the description, instructions, model, or harness. The skill is refined and re-evaluated until it meets the agreed standard.',
    },
  ],
  showcase: {
    eyebrow: 'Generated examples',
    title: 'Skills built through the same workflow.',
    description:
      'These complete examples range from lightweight project workflows to skills with scripts, references, and reusable assets. Open any skill to read its generated instructions.',
  },
  evaluationMetrics: [
    { value: '30', label: 'conversation turns' },
    { value: '4', label: 'illustrative harnesses' },
    { value: '10', label: 'comparison cells' },
  ],
  method: {
    title: 'A simple signal for a difficult failure mode.',
    description:
      'Each scenario contains thirty turns. Expected first activations are scored only at governed points, where the CLI records how many were called and missed. Every chart above is derived from that same compact website artifact.',
  },
  footer:
    'Local-first skill creation and evaluations. Static, reviewable evidence.',
} as const;

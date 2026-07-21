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
  navigation: [
    {
      id: 'explore',
      label: 'Explore',
      description: 'Build and evaluate',
    },
    {
      id: 'project',
      label: 'Project',
      description: 'Purpose and limits',
    },
    {
      id: 'analysis',
      label: 'Analysis',
      description: 'Results in context',
    },
  ],
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
    { value: '7', label: 'illustrative harnesses' },
    { value: '10', label: 'comparison cells' },
  ],
  method: {
    title: 'A simple signal for a difficult failure mode.',
    description:
      'Each scenario contains thirty turns. Expected first activations are scored only at governed points, where the CLI records how many were called and missed. Every chart above is derived from that same compact website artifact.',
  },
  project: {
    title: 'A skill can fail for more than one reason.',
    introduction:
      'Skill Issue exists to make those reasons easier to separate. It gives people a practical way to test the environment they use and a systematic way to build skills that work inside it.',
    motivationTitle: 'The same failure can point to the wrong problem.',
    motivation: [
      'When an agent ignores a skill, the description may be too weak. The instructions may be unclear after invocation. The model may select skills inconsistently. The harness may offer too little support for discovery and use.',
      'From the outside, those failures can look identical. People can spend hours rewriting a sound skill when the model and harness are the real limitation. Others see inconsistent results and decide that skills do not work at all.',
      'We want to replace that guesswork with evidence. Skill Issue records when expected calls happen, when they are missed, and which exact setup produced the result. That gives refinement a clearer starting point.',
    ],
    goals: [
      {
        title: 'Evaluate the environment',
        description:
          'Run the same governed conversations through supported model and harness combinations. Retain the calls, misses, configuration, and evidence needed to compare what happened.',
      },
      {
        title: 'Build better skills',
        description:
          'Describe an outcome in ordinary language. Generate the skill, evaluate invocation and behavior, diagnose the failure owner, and refine the right part of the system.',
      },
    ],
    progressTitle: 'The first complete path is taking shape.',
    progress:
      'The project now has a local CLI, a governed evaluation method, a generated skill workflow, and static website artifacts that keep results reviewable. The public campaign evidence and first release are still in progress.',
    facts: [
      { label: 'Execution', value: 'Local CLI' },
      { label: 'Evaluation', value: 'Three fixed scenarios' },
      { label: 'Evidence', value: 'Reviewable JSON artifacts' },
      { label: 'Publishing', value: 'Static GitHub Pages site' },
    ],
    limitationsTitle: 'This is a bounded first evaluation.',
    limitations: [
      'Time, access, and available resources limit the first campaign. It does not cover every harness, model, configuration, operating system, provider alias, or repeated trial.',
      'One scenario suite per configuration cannot establish statistical reliability, permanent rankings, universal model behavior, or guarantees about every environment. Residual configuration and version changes may also affect an observed result.',
      'Unsupported and unrun combinations are omitted. Any conclusion must stay tied to the exact configuration and retained evidence that produced it.',
    ],
  },
  analysis: {
    title: 'The charts are only the start of the result.',
    introduction:
      'This page will explain what the accepted campaign evidence shows, how large the observed differences are, and where the evidence stops. We will publish the interpretation after the campaign runs have passed review.',
    sections: [
      {
        title: 'Harness findings',
        description:
          'Hold the Codex model and Medium reasoning target constant, then explain whether changing the harness changed reliable skill calling.',
      },
      {
        title: 'Model findings',
        description:
          'Compare supported models inside shared environments and identify which patterns persist across scenarios and harnesses.',
      },
      {
        title: 'Scenario observations',
        description:
          'Show where individual conversations differ from the aggregate and keep those differences connected to the underlying calls and misses.',
      },
      {
        title: 'Limits and evidence',
        description:
          'State the exact tested configurations, the scale of the campaign, the remaining uncertainty, and the retained artifacts behind every conclusion.',
      },
    ],
    evidenceTitle: 'Accepted runs will set the story.',
    evidence:
      'The current website uses illustrative mock data to develop the presentation. No ranking, percentage, trend, or conclusion on this page will be treated as a finding until it can be traced to accepted campaign artifacts.',
  },
  footer:
    'Local-first skill creation and evaluations. Static, reviewable evidence.',
} as const;

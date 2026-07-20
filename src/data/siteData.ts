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
    { value: '30', label: 'conversation turns' },
    { value: '4', label: 'illustrative harnesses' },
    { value: '10', label: 'comparison cells' },
  ],
  method: {
    title: 'A simple signal for a difficult failure mode.',
    description:
      'Each scenario contains thirty turns. Expected first activations are scored only at governed points, where the CLI records how many were called and missed. Every chart above is derived from that same compact website artifact.',
  },
  footer: 'Local-first evaluations. Static, reviewable benchmark data.',
} as const;

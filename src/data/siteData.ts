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
      id: 'method',
      label: 'Method',
      description: 'How testing works',
    },
    {
      id: 'analysis',
      label: 'Analysis',
      description: 'Results in context',
    },
    {
      id: 'project',
      label: 'Project',
      description: 'Purpose and limits',
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
    { value: '7', label: 'supported harnesses' },
    { value: '10', label: 'comparison cells' },
  ],
  method: {
    title: 'A simple signal for a difficult failure mode.',
    description:
      'Each scenario contains thirty turns. Expected activations are scored only at governed points distributed through the conversation, where the CLI records how many were called and missed. Every chart above is derived from that same compact website artifact.',
  },
  methodology: {
    title: 'How we evaluate skill calling.',
    introduction:
      'This page explains the complete path from a governed conversation to a scored result. It shows the prompts we send, the calls we expect, the skills we supply, the controls we apply, and the limits that remain.',
    scopeTitle: 'The evaluation asks one bounded question.',
    scope: [
      'Can a configured coding agent call the right supplied skill at the point where a governed conversation requires it? We evaluate the complete setup, including the model, harness, reasoning setting, supplied skills, conversation, tools, and evaluator. A result does not belong to a bare model name.',
      'The current method uses three fixed conversations. Each contains thirty user turns and continues through one native agent session. The user messages are governed and sent in order. The assistant responses, tool use, workspace changes, and skill calls remain outputs of the model and harness.',
    ],
    processTitle: 'One path produces every scored outcome.',
    process: [
      'Build and refine the supplied skills.',
      'Start a fresh controlled workspace and agent session.',
      'Send the scenario prompts in their fixed order.',
      'Record supplied skill invocations against the active turn.',
      'Compare observed calls with the private scorecard.',
      'Retain detailed evidence and derive the compact website result.',
    ],
    scenariosTitle: 'Three governed conversations.',
    scenariosIntroduction:
      'Each scenario is a complete thirty turn product task. Open one to read every user message and the skill calls expected at that turn. The counts below come directly from the current development scenario files and will follow those files as they are finalized.',
    scenarioMetadata: {
      'gardening-web-application': {
        title: 'SproutCheck browser application',
        description:
          'Build and extend a small local plant watering application while maintaining its plan, tests, debugging discipline, skill work, and handoffs.',
      },
      'community-archive-desktop-application': {
        title: 'BoxIndex desktop archive',
        description:
          'Build and extend an offline community archive application while maintaining its plan, tests, debugging discipline, skill work, and handoffs.',
      },
      'neighborhood-emergency-preparedness-program': {
        title: 'ReadyCard preparedness tool',
        description:
          'Build and extend a local printable preparedness tool while maintaining its plan, tests, debugging discipline, skill work, and handoffs.',
      },
    },
    scoringTitle: 'The scorecard records expected turn and skill pairs.',
    scoring: [
      'A scorecard lists the skill or skills expected at each governed point. If a turn expects two skills, both expectations are scored separately. An expected pair is called when the evaluator records that skill during the active turn. Otherwise it is missed.',
      'The detailed result keeps named expected, observed, missing, additional, and unattributed calls. The compact website result keeps the called and missed totals for each turn with at least one expectation. Aggregate charts sum those individual Boolean outcomes.',
    ],
    dictatePlanTitle: 'Dictate Plan is an explicit startup call.',
    dictatePlan: [
      'The first message in every scenario names Dictate Plan directly. This is intentional and manual. It tests whether the agent follows the explicit startup request, and it is included in the Turn 1 score.',
      'Codex receives metadata that prevents implicit invocation of Dictate Plan. The repository does not establish the same technical restriction for every other harness, so the methodology relies on the explicit first message and discloses that boundary.',
    ],
    instrumentationTitle: 'Calls are recorded with an opaque skill token.',
    instrumentation: [
      'The evaluator makes a disposable copy of every supplied skill and adds one short signal instruction near the top. The signal contains a random token. Its meaning is stored outside the model workspace, where the evaluator maps it back to a skill and the active conversation turn.',
      'The token meaning is opaque, but the instrumentation is not invisible. An agent can see the added command and may infer that recording machinery exists. The extra instruction can influence behavior, so this is part of the method and part of its limitations.',
    ],
    controlsTitle: 'The environments are controlled, not identical.',
    controls: [
      'Every campaign run starts in a fresh external workspace. The adapters suppress unrelated project instructions, ambient skills, plugins, memory, and configuration where each harness permits it. Every compared setup receives the same governed scenario content and supplied skill set.',
      'The remaining boundaries differ by product. Authentication, managed policy, built in tools, session state, filesystem access, and network behavior cannot be normalized into one universal sandbox. Results therefore describe the exact configured setup recorded for the run.',
    ],
    skillsTitle: 'The supplied skills are evaluated before the campaign.',
    skillsIntroduction:
      'The project tests both whether a skill is selected and whether its body produces the intended behavior. Failures are traced to the description, body, supporting material, or surrounding system, then the owning part is refined and evaluated again. Open a skill to read its current instructions or follow the evidence link to its retained refinement record.',
    evidenceTitle: 'The method is designed to be inspectable.',
    evidence: [
      'The real scenarios, scorecards, skills, evaluator source, detailed run artifacts, and retained refinement records provide different layers of evidence. Website summaries are derived views and do not replace those sources.',
      'A result is published only after its scenario, configured setup, completion state, cleanup, evidence, and provenance have been accepted together. Inspectable source makes the method reviewable. It does not prove that another party has independently reproduced the result.',
    ],
    limitationsTitle: 'This is a bounded evaluation.',
    limitations: [
      'Three authored scenarios do not represent every skill, coding task, model, harness, configuration, or operating environment. The current campaign is also limited by the time, access, and personal resources available to one builder during Build Week.',
      'Fixed user messages improve comparability, but a later prompt can arrive after different agents have taken different paths. Instrumentation can influence behavior. Harness controls leave different residual surfaces. Provider aliases and product behavior can also change over time.',
      'The results support claims about the exact accepted runs and nothing broader. They do not establish permanent rankings, universal reliability, or guarantees about skill calling in every environment.',
    ],
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
      'The project now has a local CLI, a governed evaluation method, a generated skill workflow, and static website artifacts that publish results as reviewable evidence.',
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
      'This page explains what the accepted campaign evidence shows, how large the observed differences are, and where the evidence stops.',
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
    evidenceTitle: 'Every finding stays tied to its evidence.',
    evidence:
      'Every ranking, percentage, trend, and conclusion is derived from accepted campaign artifacts and remains tied to the exact configuration and retained evidence that produced it.',
  },
  footer:
    'Local-first skill creation and evaluations. Static, reviewable evidence.',
} as const;

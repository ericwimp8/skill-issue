type ReleaseDownload = {
  architecture: string;
  id: string;
  platform: string;
  url: string | null;
};

const releaseDownloads: ReleaseDownload[] = [
  {
    architecture: 'Apple silicon',
    id: 'darwin-arm64',
    platform: 'macOS',
    url: null,
  },
  {
    architecture: 'Intel',
    id: 'darwin-amd64',
    platform: 'macOS',
    url: null,
  },
  {
    architecture: 'x64',
    id: 'windows-amd64',
    platform: 'Windows',
    url: null,
  },
  {
    architecture: 'ARM64',
    id: 'windows-arm64',
    platform: 'Windows',
    url: null,
  },
  {
    architecture: 'x64',
    id: 'linux-amd64',
    platform: 'Linux',
    url: null,
  },
  {
    architecture: 'ARM64',
    id: 'linux-arm64',
    platform: 'Linux',
    url: null,
  },
];

export const siteData = {
  status: 'Open-source · Local-first',
  hero: {
    eyebrow: 'Built with skills to build, evaluate, and fix skills.',
    title: 'It’s not a skill issue, but it’s always a skills issue.',
    description:
      'Create skills that work, then find out whether your model and harness can be trusted to call them. Skill Issue makes agent skills easier to create, evaluate, and trust.',
  },
  release: {
    downloads: releaseDownloads,
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
    { value: '5', label: 'supported harnesses' },
    { value: '9', label: 'comparison cells' },
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
    title: 'Skill calling belongs to the pairing.',
    introduction:
      'The first campaign produced sharply different results from configurations that shared a model or shared a harness. The clearest finding is that reliable skill calling describes the complete model and harness setup, not either name on its own.',
    scopeTitle: 'One campaign, one bounded question.',
    scope: [
      'The campaign asks whether a configured coding agent calls the supplied skill at each governed point in a thirty turn conversation. Every expected turn and skill pair produces one Boolean result: called or missed. Additional calls are retained separately and do not improve the success percentage.',
      'The comparisons below describe accepted runs under the tested models, harnesses, reasoning settings, scenarios, skills, and evaluator. They measure skill invocation in this setup. They do not measure general coding ability or establish a permanent model or harness ranking.',
    ],
    pairingTitle: 'The same model produced a perfect result and a low band.',
    pairing: [
      'Codex Sol called every expected skill when it ran in OpenAI Codex. The same model recorded much lower coverage in OpenCode and Pi, even though the scenarios, supplied skills, and scoring rules stayed fixed.',
      'That spread is evidence of a harness associated effect inside the tested setup. The evaluation does not isolate one mechanism. System prompting, skill surfacing, invocation instructions, tools, and product behavior remain bundled together in the harness.',
    ],
    modelTitle: 'The same harness produced two different behaviors.',
    model: [
      'Cursor held the surrounding product constant while Grok and Composer moved in opposite directions. Grok sustained skill use through the conversations. Composer engaged early, then called fewer of the recurring skills as the tasks continued.',
      'This matched comparison rules out a simple claim that Cursor is uniformly good or uniformly bad for skill calling. The model still matters inside the harness, just as the harness matters around the model.',
    ],
    callingStyleTitle: 'High coverage did not require one calling style.',
    callingStyle: [
      'OpenAI Codex reached full expected call coverage while also making many additional calls. Grok reached nearly the same coverage with far fewer additional calls. Both found the governed calls, but Grok was more selective in these scenarios.',
      'Composer, Claude Fable, OpenCode, and Pi showed a different shape. Their lower totals came primarily from missed expected calls, not from uncontrolled extra calling. They did not often call the wrong supplied skill. They mostly stopped calling supplied skills at all.',
    ],
    timeTitle:
      'The weakest configurations became quieter as the task continued.',
    time: [
      'The turn order matters because each scenario stays inside one continuing agent session. Splitting the conversations into ten turn bands shows whether invocation survives as context and task state accumulate.',
      'The strongest configurations remained active throughout the full conversation. Several lower scoring configurations recorded some early calls and almost none in the final ten turns. That pattern is consistent with skills being treated as setup tools rather than continuing operating instructions.',
    ],
    validityTitle:
      'The ceiling was reachable, and the quiet turns stayed quiet.',
    validity: [
      'A complete 137 out of 137 result shows that the scorecard did not demand an impossible sequence. Turns 13, 18, and 24 contained no expected call by design. The highest coverage configurations made no additional calls on those decoy turns, so their result was not produced by calling indiscriminately on every message.',
      'The Codex results were re-derived after retained evidence exposed an attribution gap. The first matcher recognized echo-form instrumentation but missed equivalent printf-form signals and direct governed skill reads. The matcher was corrected, the accepted artifacts were reconciled from the raw transcripts, and the website now publishes those corrected counts.',
    ],
    meaningTitle: 'The result changes where diagnosis should begin.',
    meaning: [
      'When a skill is ignored, rewriting the skill is not automatically the right first response. A configuration can suppress reliable invocation even when another configuration calls the same supplied skills throughout the same governed conversation.',
      'For skill authors, these results support testing the complete environment before repeatedly changing a sound skill. For harness and model evaluators, they support matched comparisons that hold one side of the pairing constant. They do not establish why a particular configuration behaved as it did without further trace review.',
    ],
    limitationsTitle: 'These are descriptive results, not universal rates.',
    limitations: [
      'Each complete configuration currently contributes one run for each of three authored scenarios. That is enough to describe the recorded campaign and expose large differences. It is not enough to estimate statistical reliability, persistence across repeated attempts, or behavior across the wider population of coding tasks.',
      'The campaign covers one machine, one time window, one supplied skill set, one scenario suite, and the recorded product versions. Harness controls are similar rather than identical. Codex attribution is capture based, while other harnesses execute instrumentation markers, so the observation paths also have different failure modes.',
      'Every conclusion on this page should be read as a statement about the accepted evidence. Broader claims require repeated runs, additional scenario families, more skills, and explicit tests of the mechanisms that remain bundled inside each configuration.',
    ],
  },
  footer:
    'Local-first skill creation and evaluations. Static, reviewable evidence.',
} as const;

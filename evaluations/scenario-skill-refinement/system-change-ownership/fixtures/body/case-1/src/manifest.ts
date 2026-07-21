export type CommandDefinition = {
  id: string;
  title: string;
  requiredCapabilities: string[];
};

export const commands: CommandDefinition[] = [
  { id: 'search', title: 'Search', requiredCapabilities: ['network'] },
  { id: 'summarize', title: 'Summarize', requiredCapabilities: [] },
];

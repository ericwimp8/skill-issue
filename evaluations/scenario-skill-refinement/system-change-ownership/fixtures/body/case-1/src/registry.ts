import { commands, type CommandDefinition } from './manifest';

export class CommandRegistry {
  list(): CommandDefinition[] {
    return commands;
  }
}

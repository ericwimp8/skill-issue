import type { CommandDefinition } from './manifest';
import { CommandRegistry } from './registry';

export class CommandPicker {
  constructor(private readonly registry: CommandRegistry) {}

  visibleCommands(): CommandDefinition[] {
    return this.registry.list();
  }
}

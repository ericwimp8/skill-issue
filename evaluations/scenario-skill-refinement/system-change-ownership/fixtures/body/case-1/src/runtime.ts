export class RuntimeCapabilities {
  constructor(private readonly available: Set<string>) {}

  supports(capability: string): boolean {
    return this.available.has(capability);
  }
}

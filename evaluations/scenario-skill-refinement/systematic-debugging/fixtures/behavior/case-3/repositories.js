import { readFile } from 'node:fs/promises';

/** Repository implementations return a profile object. */
export class DiskProfileRepository {
  constructor(path) {
    this.path = path;
  }

  async load() {
    return JSON.parse(await readFile(this.path, 'utf8'));
  }
}

export class MemoryProfileRepository {
  constructor(profile) {
    this.profile = profile;
  }

  async load() {
    return this.profile;
  }
}

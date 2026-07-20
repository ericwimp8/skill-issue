import { appendFile } from "node:fs/promises";
import type { RequestStore } from "./requestStore.js";

export class FileRequestStore implements RequestStore {
  constructor(private readonly path: string) {}

  async append(payload: string): Promise<void> {
    await appendFile(this.path, `${payload}\n`, "utf8");
  }
}


import type { Configuration } from "./config.js";
import { createRequestHandler } from "./http/requestHandler.js";
import { FileRequestStore } from "./storage/fileRequestStore.js";

export function createApplication(configuration: Configuration) {
  const store = new FileRequestStore(configuration.dataPath);
  const handleRequest = createRequestHandler(store);
  return { listen: () => Bun.serve({ port: configuration.port, fetch: handleRequest }) };
}


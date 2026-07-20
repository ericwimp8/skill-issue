export type Configuration = { port: number; dataPath: string };

export function loadConfiguration(environment: NodeJS.ProcessEnv): Configuration {
  return {
    port: Number(environment.PORT ?? "8080"),
    dataPath: environment.DATA_PATH ?? ".local/requests.jsonl",
  };
}


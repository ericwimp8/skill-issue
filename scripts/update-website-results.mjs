import { readFile, writeFile } from 'node:fs/promises';
import path from 'node:path';

const inputPaths = process.argv.slice(2);

if (inputPaths.length === 0) {
  console.error(
    'Usage: npm run results:update -- path/to/website.json [more files]',
  );
  process.exit(1);
}

const artifacts = await Promise.all(
  inputPaths.map(async (inputPath) => {
    const source = await readFile(path.resolve(inputPath), 'utf8');
    const artifact = JSON.parse(source);

    if (
      artifact.schema_version !== 1 ||
      typeof artifact.run_id !== 'string' ||
      typeof artifact.scenario_id !== 'string' ||
      typeof artifact.harness !== 'string' ||
      typeof artifact.model !== 'string' ||
      typeof artifact.total_turns !== 'number' ||
      !Array.isArray(artifact.points)
    ) {
      throw new Error(`${inputPath} is not a valid website artifact`);
    }

    return artifact;
  }),
);

artifacts.sort(
  (left, right) =>
    left.harness.localeCompare(right.harness) ||
    left.model.localeCompare(right.model) ||
    left.scenario_id.localeCompare(right.scenario_id),
);

const destination = path.resolve('src/data/publishedWebsiteArtifacts.json');
await writeFile(destination, `${JSON.stringify(artifacts, null, 2)}\n`);
console.log(`Updated ${destination} with ${artifacts.length} artifacts.`);

import { readFile, readdir, writeFile } from 'node:fs/promises';
import path from 'node:path';

const acceptedDirectory = path.resolve(
  'evaluations/skill-calling/results/accepted',
);
const destination = path.resolve('src/data/publishedWebsiteArtifacts.json');

function isNonNegativeInteger(value) {
  return Number.isInteger(value) && value >= 0;
}

function isWebsitePoint(point, totalTurns) {
  return (
    point !== null &&
    typeof point === 'object' &&
    Number.isInteger(point.turn) &&
    point.turn >= 1 &&
    point.turn <= totalTurns &&
    typeof point.turn_id === 'string' &&
    isNonNegativeInteger(point.called) &&
    isNonNegativeInteger(point.missed) &&
    isNonNegativeInteger(point.unexpected)
  );
}

function isReconciliation(value) {
  return (
    value !== null &&
    typeof value === 'object' &&
    typeof value.basis === 'string' &&
    value.basis.trim() !== '' &&
    typeof value.reason === 'string' &&
    value.reason.trim() !== ''
  );
}

function isWebsiteArtifact(artifact) {
  return (
    artifact !== null &&
    typeof artifact === 'object' &&
    artifact.schema_version === 2 &&
    typeof artifact.run_id === 'string' &&
    typeof artifact.scenario_id === 'string' &&
    typeof artifact.harness === 'string' &&
    typeof artifact.model === 'string' &&
    Number.isInteger(artifact.total_turns) &&
    artifact.total_turns > 0 &&
    Array.isArray(artifact.points) &&
    artifact.points.every((point) =>
      isWebsitePoint(point, artifact.total_turns),
    ) &&
    (artifact.reconciliation === undefined ||
      isReconciliation(artifact.reconciliation))
  );
}

const acceptedFiles = (await readdir(acceptedDirectory))
  .filter((name) => name.endsWith('.json'))
  .sort();

if (acceptedFiles.length === 0) {
  throw new Error(
    `No accepted website artifacts found in ${acceptedDirectory}`,
  );
}

const artifacts = await Promise.all(
  acceptedFiles.map(async (name) => {
    const inputPath = path.join(acceptedDirectory, name);
    const artifact = JSON.parse(await readFile(inputPath, 'utf8'));

    if (!isWebsiteArtifact(artifact)) {
      throw new Error(`${inputPath} is not a valid schema-v2 website artifact`);
    }

    if (name !== `${artifact.run_id}.json`) {
      throw new Error(`${inputPath} must be named ${artifact.run_id}.json`);
    }

    const published = {
      schema_version: artifact.schema_version,
      run_id: artifact.run_id,
      scenario_id: artifact.scenario_id,
      harness: artifact.harness,
      model: artifact.model,
      total_turns: artifact.total_turns,
      points: artifact.points.map((point) => ({
        turn: point.turn,
        turn_id: point.turn_id,
        called: point.called,
        missed: point.missed,
        unexpected: point.unexpected,
      })),
    };

    if (artifact.reconciliation !== undefined) {
      published.reconciliation = {
        basis: artifact.reconciliation.basis,
        reason: artifact.reconciliation.reason,
      };
    }

    return published;
  }),
);

const runIds = new Set();
const configurations = new Set();

artifacts.forEach((artifact) => {
  const configuration = [
    artifact.harness,
    artifact.model,
    artifact.scenario_id,
  ].join('::');

  if (runIds.has(artifact.run_id)) {
    throw new Error(`Duplicate accepted run ID: ${artifact.run_id}`);
  }
  if (configurations.has(configuration)) {
    throw new Error(`Duplicate accepted configuration: ${configuration}`);
  }

  runIds.add(artifact.run_id);
  configurations.add(configuration);
});

artifacts.sort(
  (left, right) =>
    left.harness.localeCompare(right.harness) ||
    left.model.localeCompare(right.model) ||
    left.scenario_id.localeCompare(right.scenario_id),
);

await writeFile(destination, `${JSON.stringify(artifacts, null, 2)}\n`);
console.log(`Updated ${destination} with ${artifacts.length} artifacts.`);

# Request Storage: First-Day Handoff

## Start Here

This service accepts request bodies over HTTP and appends them to a local line-oriented file. The production path is small: the CLI loads environment-backed configuration, the application wires a file store into the HTTP handler, and the handler delegates each accepted request to that store.

The likely change surface is `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/storage/`. Changes to the storage contract may also require updates to the application wiring and HTTP behavior.

## Governing Instructions

- `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/AGENTS.md` governs the whole service. It requires `npm run validate` before completion, assigns `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/storage/` to the persistence layer, and keeps generated `.local/` content untracked.
- `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/storage/AGENTS.md` has narrower precedence for storage files. It names `FileRequestStore` as the owner of durable request writes and requires `npm test -- storage` before repository-wide validation.
- Therefore, a persistence change must preserve ownership in `FileRequestStore` unless production wiring demonstrates a broader ownership change, then run the storage-scoped check followed by the repository-wide check.

## Runtime and Configuration

1. `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/cli.ts` is the bootstrap. It passes `process.env` to `loadConfiguration`, creates the application, and calls `listen()`.
2. `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/config.ts` is the configuration owner. `PORT` becomes a number and defaults to `8080`; `DATA_PATH` defaults to `.local/requests.jsonl`.
3. `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/application.ts` constructs `FileRequestStore(configuration.dataPath)`, injects it into `createRequestHandler`, and passes the handler and configured port to `Bun.serve`.
4. `.env` is ignored, but no production source loads an `.env` file. A launcher or operator must place values in the actual process environment.
5. `.local/`, `dist/`, and `node_modules/` are ignored local or generated state. The configured data file is runtime state, not tracked source.

## Request-to-Storage Trace

For a `POST` request:

1. `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/http/requestHandler.ts` reads the complete request body as text.
2. It calls the injected `RequestStore.append(payload)` contract from `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/storage/requestStore.ts`.
3. The concrete `FileRequestStore.append` in `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/storage/fileRequestStore.ts` calls Node's `appendFile` with the configured path, the payload plus a newline, and UTF-8 encoding.
4. After the write resolves, the caller receives status `202` with body `Accepted`.

The concrete persistence effect is an append to `DATA_PATH`; the implementation does not create its parent directory, parse JSON, or normalize embedded newlines.

## Failure Returned to Callers

- Any exception while reading the body or awaiting `store.append` is caught by the handler and converted to status `503` with body `Unavailable`.
- The handler suppresses the underlying error, so callers receive no storage error details and the current path records no diagnostic.
- Any method other than `POST` returns status `404` with body `Not found` before storage is called.

If storage behavior changes, preserve or deliberately revise this HTTP boundary in `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/http/requestHandler.ts`; storage exceptions currently become the caller-visible availability signal.

## Required Checks and Current Evidence

Run from `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/`, in this order:

1. `npm test -- storage` — required by the storage-scoped instructions.
2. `npm run validate` — required by the repository instructions; the manifest defines it as `npm run build && npm test`.

Both commands were executed during this handoff and currently fail before validating storage behavior:

- `npm test -- storage` exits `1` because Node cannot find a path named `storage`.
- `npm run validate` exits `1` because `tsc -p tsconfig.json` cannot find `tsconfig.json`.

The only test file, `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/test/requestHandler.test.ts`, exports a historical port constant and does not exercise the handler or store. Passing behavioral coverage has not been established.

## Claims Requiring Verification

- **Supported runtime:** `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/package.json` starts the build with Node, while `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/src/application.ts` calls the Bun global. Locate the missing toolchain or deployment owner and prove which runtime launches the service.
- **Build contract:** no `tsconfig.json`, lockfile, dependency declaration, or engine constraint is present. Recover the intended build inputs before treating `npm run build` as usable.
- **Data-directory provisioning:** the default parent `.local/` is ignored, and production code does not create it. Verify whether deployment or startup automation provisions the directory and its permissions.
- **README instructions:** `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/README.md` says `yarn dev`, port `3000`, direct router writes, and a JSON file. Current manifest and source instead show no `dev` script, default port `8080`, a handler-to-store abstraction, and newline-appended raw text. Treat the README claims as stale until an owning runtime source proves otherwise.
- **Failure observability:** the caller contract is visibly `503`, but no logging or monitoring path exists in this fixture. Verify the operational diagnostics requirement before changing error handling.
- **Record format and concurrency:** the code appends raw bodies with a newline. Confirm payload constraints, embedded-newline handling, ordering, concurrent-write expectations, and recovery requirements with the service owner before redesigning persistence.

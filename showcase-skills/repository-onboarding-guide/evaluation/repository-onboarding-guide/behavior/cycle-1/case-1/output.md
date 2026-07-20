# Sample Service Onboarding Guide

## Service Shape

This fixture is a small TypeScript HTTP ingestion service. `src/cli.ts` is the process entry, `src/application.ts` wires configuration, HTTP handling, and persistence, and `src/storage/fileRequestStore.ts` owns the concrete file write. The service accepts any `POST` body as text, appends it as one line, and returns a status-only text response.

## Instruction Authority

- `AGENTS.md` governs the whole fixture: run `npm run validate` before calling changes complete, treat `src/storage/` as persistence-owned, and keep `.local/` generated state untracked.
- `src/storage/AGENTS.md` has narrower precedence for persistence files: `FileRequestStore` owns durable writes, and a persistence change must run `npm test -- storage` before repository-wide validation.
- `package.json` owns the available build, start, test, and validation scripts. `README.md` is descriptive documentation; where it conflicts with source or the manifest, use the production source and manifest evidence below.

For a persistence change, work primarily in `src/storage/`. A contract change may also require `src/storage/requestStore.ts`, while HTTP response policy remains owned by `src/http/requestHandler.ts`.

## Architecture And Dependencies

- **Process and configuration:** `src/cli.ts` passes `process.env` to `loadConfiguration`, creates the application, then calls `listen()`.
- **Configuration:** `src/config.ts#loadConfiguration` converts `PORT` to a number (default `8080`) and reads `DATA_PATH` (default `.local/requests.jsonl`). No validation rejects `NaN`, empty paths, or inaccessible locations.
- **Composition root:** `src/application.ts#createApplication` constructs `FileRequestStore` with the configured data path, injects it into `createRequestHandler`, and registers the returned handler with `Bun.serve`.
- **HTTP policy:** `src/http/requestHandler.ts#createRequestHandler` depends on the `RequestStore` interface, accepts only `POST`, and maps success to `202` and any thrown body-read or store error to `503`.
- **Persistence:** `src/storage/requestStore.ts` defines `append(payload)`. `src/storage/fileRequestStore.ts#FileRequestStore.append` implements it with Node's `node:fs/promises.appendFile`.

Dependency direction is `cli -> config/application -> HTTP contract -> storage interface -> FileRequestStore -> appendFile`. Construction selects the file-backed implementation directly; there is no persistence registry, database layer, or migration entry point in the fixture.

## POST To Durable Write

1. `src/cli.ts` loads `PORT` and `DATA_PATH`, passes the resulting `Configuration` to `createApplication`, and invokes `listen`.
2. `src/application.ts#createApplication` constructs `new FileRequestStore(configuration.dataPath)`, injects it into `createRequestHandler`, and gives that concrete handler to `Bun.serve({ port, fetch })`.
3. For a `POST`, `src/http/requestHandler.ts#handleRequest` awaits `request.text()`, so the request body becomes an unvalidated string.
4. The handler calls the injected `RequestStore.append`; at runtime the composition root guarantees this is `FileRequestStore.append`.
5. `src/storage/fileRequestStore.ts#append` calls `appendFile(path, payload + "\n", "utf8")`. This is the concrete side effect: bytes are appended to the configured path. Despite the `.jsonl` default extension, neither the handler nor store parses or validates JSON.
6. After a successful append the handler returns `202 Accepted`. Any exception from `request.text()` or `appendFile` is swallowed by the same `catch` and returned as `503 Unavailable`; no error is logged or distinguished. Non-`POST` requests stop before body reading or persistence and return `404 Not found`.

`appendFile` creates a missing file but does not create its parent directory. With the default path, `.local/` must already exist or the first write will reach the handler's `503` path.

## Setup And Validation

Run commands from the fixture root.

| Purpose | Command | Ownership and observed status |
| --- | --- | --- |
| Persistence-focused check | `npm test -- storage` | Required first by `src/storage/AGENTS.md`. Executed: **failed (exit 1)** because `node --test storage` could not find `storage`. |
| Repository validation | `npm run validate` | Required by root `AGENTS.md`; expands to build then tests. Executed after the focused check: **failed (exit 1)** during `tsc -p tsconfig.json` because `tsconfig.json` is absent, so tests did not run. |
| Build | `npm run build` | Declared as `tsc -p tsconfig.json`; its failure was observed through `npm run validate`. |
| Start built service | `npm start` | Declared as `node dist/cli.js`; not executed because validation cannot produce `dist/cli.js`. |

No install script, lockfile, dependency declarations, TypeScript configuration, or supported tool versions are present. The source requires TypeScript/Node types for building and Bun at runtime, while the declared start command invokes Node. Until that runtime mismatch is resolved, there is no source-backed working setup or run command.

For a persistence contribution, the intended sequence is focused persistence check, then full validation. Both required commands need project repair before they can qualify a change; a passing ad hoc command would not replace them.

## State And Secrets

- `.gitignore` marks `node_modules/` as dependency-local, `dist/` as generated build output, and `.local/` as generated runtime state. The default request log is `.local/requests.jsonl` and can contain complete request bodies, so treat it as potentially sensitive.
- `.env` is ignored and is the likely secret-bearing configuration location, although no loader is present in production source. Do not record its values.
- `PORT` and `DATA_PATH` are read from the process environment. `DATA_PATH` can redirect writes outside `.local/`; its deployed value and retention/access controls are unknown.
- The repository instructions require `.local/` to remain untracked. No source creates or initializes that directory.

## Conflicts And Unknowns

### Documentation And Validation Conflicts

- `README.md` says `yarn dev` serves port `3000`, but `package.json` has no `dev` script, no package manager or lockfile is established, and source defaults to port `8080`.
- `README.md` says the router writes directly to a JSON file. Production wiring instead has the handler call a `RequestStore`, then `FileRequestStore` append raw text to a JSONL-named file.
- `test/requestHandler.test.ts` deliberately retains the historical port `3000`; it does not exercise the request handler or persistence and conflicts with the current `8080` default.
- The persistence instruction's `npm test -- storage` selector is incompatible with the current `node --test` script and file layout. Repository validation is also non-runnable because its referenced `tsconfig.json` is missing.
- `npm start` uses Node, but the built application calls the Bun global. The manifest and production runtime implementation therefore name incompatible runtime owners.

### Material Unknowns And Next Checks

- **Intended runtime and start path:** determine whether Bun or Node is authoritative by obtaining the omitted deployment/CI configuration or maintainer decision; then align `application.ts`, scripts, and toolchain metadata.
- **Build and dependency baseline:** recover or define `tsconfig.json`, dependencies, lockfile, and supported Node/Bun/TypeScript versions, then rerun `npm run validate` from a clean checkout.
- **Persistence test scope:** establish the intended test files or correct Node test selector, then make `npm test -- storage` select real persistence coverage before relying on it.
- **Data contract:** confirm whether POST bodies must be valid JSON and whether JSONL framing/escaping is required; current source accepts arbitrary text, including embedded newlines.
- **Operational durability:** confirm who creates the data directory, the deployed `DATA_PATH`, concurrent append expectations, file permissions, rotation, retention, and recovery. Exercise a POST against the chosen runtime with a temporary data path to verify the concrete success and failure responses.
- **Error observability:** decide whether body-read and filesystem failures require distinct logging or response policy; current source erases the cause and supplies no diagnostics.

Evidence level: source-traced through the concrete `appendFile` effect, with both scoped required checks executed. Runtime serving and a live POST were not verified because the checked-in build/runtime setup is incomplete and conflicting.

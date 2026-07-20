# Smallest Reliable Contributor Guide

## What this service is

This fixture is a small request-ingestion service. Its production path is TypeScript under `src/`: `src/cli.ts` boots the application, `src/application.ts` wires HTTP to persistence, `src/http/requestHandler.ts` owns HTTP behavior, and `src/storage/` owns durable writes.

Current source proves a Bun HTTP server that accepts every `POST` body as raw text, appends that text plus a newline to a configured file, and returns `202`. This is JSONL-shaped storage only when callers send JSON; neither the handler nor store parses or validates JSON.

## Authority and scope

- `AGENTS.md` governs the whole fixture. It requires `npm run validate`, assigns `src/storage/` to the persistence layer, and keeps `.local/` generated state untracked.
- `src/storage/AGENTS.md` adds persistence-specific rules: `FileRequestStore` owns durable writes and persistence changes should run `npm test -- storage` before repository-wide validation.
- `package.json` owns the available project scripts. Production source owns runtime behavior. `README.md` and `test/requestHandler.test.ts` are historical or descriptive evidence where they conflict with source.

## Runtime wiring

1. `src/cli.ts` passes `process.env` to `loadConfiguration`, creates the application, and calls `listen()`.
2. `src/config.ts` converts `PORT` with `Number`, defaulting to `8080`, and reads `DATA_PATH`, defaulting to `.local/requests.jsonl`. It performs no port validation.
3. `src/application.ts` constructs `FileRequestStore(configuration.dataPath)`, injects it through the `RequestStore` boundary into `createRequestHandler`, and registers that handler as `Bun.serve(...).fetch` on the configured port.
4. `src/http/requestHandler.ts` returns `404` for non-`POST` requests. For `POST`, it reads the entire body as text and awaits `store.append`; success returns `202 Accepted`, while any read or append error returns `503 Unavailable`.
5. The injected implementation, `src/storage/fileRequestStore.ts`, reaches the concrete effect: Node's `appendFile` appends `` `${payload}\n` `` as UTF-8 to the configured path.

An ad hoc Bun smoke check, clearly separate from project-owned automation, confirmed this trace: a `POST` returned `202 Accepted` and the temporary file contained the submitted payload followed by a newline. Temporary state was removed afterward.

## Run and validation commands

Run commands from the fixture root.

| Command | Authority | Observed status |
| --- | --- | --- |
| `npm run validate` | Required by root `AGENTS.md`; script defined in `package.json` | Executed; failed in `npm run build` because `tsconfig.json` is absent, so tests did not run in this chain. |
| `npm run build` | Defined by `package.json` | Reached through validation; failed because `tsconfig.json` is absent. |
| `npm test` | Defined by `package.json` | Executed; exited 0, but the only test file exports a historical port constant and makes no behavioral assertion. |
| `npm test -- storage` | Required for persistence changes by `src/storage/AGENTS.md` | Executed; failed because Node could not find a `storage` test target. |
| `npm run start` | Defined by `package.json` as the production launcher | Executed; failed because `dist/cli.js` does not exist. |
| `yarn dev` | Mentioned only in `README.md` | Merely discovered and unsupported: there is no `dev` script, package-manager declaration, or lockfile. |

The successful ad hoc command `PORT=18083 DATA_PATH=<temporary-path> bun src/cli.ts` proves this checkout can exercise the source path, but it is diagnostic evidence, not an authoritative contributor command. Until the owners resolve the build/runtime contract, there is no project-owned command that reliably starts the service from a clean checkout.

## Local and generated values

- `.gitignore` excludes `node_modules/`, `dist/`, `.local/`, and `.env`; keep all four local and untracked.
- `PORT` and `DATA_PATH` are runtime-local inputs. Do not commit their environment-specific values or any `.env` contents.
- The default data file is `.local/requests.jsonl`. `appendFile` can create the file, but source does not create the `.local/` parent directory; a missing parent becomes a caught append error and a `503` response.
- `dist/` is the intended generated start target, although the current build cannot produce it because its TypeScript configuration is missing.

## Smallest change workflow

1. Trace the change from `src/cli.ts` through configuration and application wiring to the concrete HTTP or persistence owner.
2. Keep HTTP policy in `src/http/requestHandler.ts`; keep durable-write changes in `src/storage/`, following its nested instructions.
3. Use temporary or ignored paths for runtime data. Do not rely on the current default until `.local/` exists or directory creation has an owner-approved implementation.
4. Run the narrow project-owned check when applicable, then `npm run validate`. Report their actual failures until the repository owners repair the command contract; a passing bare `npm test` does not establish service behavior.

## Owner resolution required

- **Runtime and build:** choose and encode the supported runtime. `npm run start` invokes Node, while production application code requires the Bun global; `tsconfig.json` and dependency/toolchain declarations are absent.
- **Package manager and development command:** reconcile README's `yarn dev` with npm-based instructions/scripts and add one authoritative clean-checkout workflow.
- **Port:** reconcile README's `3000` and the explicitly historical test constant with source's current `8080` default.
- **Persistence contract:** decide whether bodies must be validated JSON, arbitrary text records, or another format; current code delegates through `RequestStore` and writes raw newline-delimited payloads rather than router-owned JSON writes.
- **Data-directory lifecycle:** assign responsibility for creating the configured parent directory or document it as an operator prerequisite.
- **Validation:** repair the missing build configuration and the scoped storage test route, then add assertions covering the production behavior owners rather than a historical constant.

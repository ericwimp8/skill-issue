# Sample Service Onboarding

## Start Here

- Work from `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/service-repo/`.
- Repository-wide instructions are in `AGENTS.md`; the deeper `src/storage/AGENTS.md` additionally governs persistence files and names `FileRequestStore` as the durable-write owner.
- `src/cli.ts` is the process entry point: it loads environment configuration, creates the application, and starts listening.
- `src/application.ts` is the composition root: it constructs `FileRequestStore`, injects it into the request handler, and registers that handler with `Bun.serve`.

## POST to Persistence

1. `src/cli.ts` passes `process.env` to `loadConfiguration` and calls `application.listen()`.
2. `src/config.ts` resolves `PORT` (default `8080`) and `DATA_PATH` (default `.local/requests.jsonl`).
3. `src/application.ts` creates `FileRequestStore(dataPath)`, injects it into `createRequestHandler`, and exposes the handler through `Bun.serve`.
4. `src/http/requestHandler.ts` accepts any `POST`, reads the complete body as text, and calls the `RequestStore.append` contract from `src/storage/requestStore.ts`. Non-POST requests return `404`.
5. `src/storage/fileRequestStore.ts` implements the contract with `appendFile`, writing the raw body plus a newline to the configured path. Success returns `202`; any body-read or append failure is collapsed to `503`.

The store does not create the parent directory, so the default `.local/` directory must already exist before the first successful write.

## Setup and Validation

- Create local state with `mkdir -p .local`; keep it untracked.
- Build with `npm run build` and run the repository-required final check with `npm run validate`.
- For persistence changes, run `npm test -- storage` first, then `npm run validate`, as required by `src/storage/AGENTS.md`.
- `npm start` is the declared start command, but it cannot be relied on as written: it launches Node while production source calls the Bun global. The repository also lacks the referenced `tsconfig.json`, dependencies, and a lockfile, so the build/setup path is incomplete. These commands are source-discovered and were not executed.

## Keep Out of Commits

- `.gitignore` excludes `node_modules/`, compiled `dist/`, runtime `.local/`, and `.env` secrets. Do not force-add them.
- Treat `DATA_PATH` contents as runtime data; the default location is already under ignored `.local/`.

## Documentation Conflicts

- `README.md` says `yarn dev`, but `package.json` defines no `dev` script and supplies no Yarn lockfile.
- `README.md` says port `3000`; `src/config.ts` defaults to `8080`. The test file explicitly labels `3000` a historical expectation rather than current behavior.
- `README.md` says the router writes directly to JSON. The implementation has no router and writes raw newline-delimited payloads through `requestHandler` → `RequestStore` → `FileRequestStore`.
- The declared `npm start` command uses Node, while `src/application.ts` requires `Bun.serve`; the intended runtime invocation needs clarification from the project owner.

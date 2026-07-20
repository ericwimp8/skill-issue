# Behavior Ground Truth

## Case 1 — Service Repository

- Root `AGENTS.md` governs the repository; nested `src/storage/AGENTS.md` adds persistence-specific ownership and validation.
- `src/cli.ts` loads `PORT` defaulting to `8080` and `DATA_PATH` defaulting to `.local/requests.jsonl`, creates `FileRequestStore`, wires it through `createRequestHandler`, and starts `Bun.serve`.
- A POST body flows through the handler and `RequestStore` interface to `FileRequestStore.append`, which calls `node:fs/promises.appendFile`; write failures become HTTP 503 responses.
- README and test claims about port 3000 are stale relative to production configuration.
- `npm run validate` is authoritative from `AGENTS.md` and `package.json`; the fixture does not prove it was executed successfully.
- `node_modules/`, `dist/`, `.local/`, and `.env` are ignored; `.env` is secret-bearing and values must not be exposed.

## Case 2 — Package Repository

- `CONTRIBUTING.md` and `Makefile` ground `make check` and `make generate`.
- `cmd/catalog/main.go` calls `catalog.Run`, which reads CLI arguments, calls `schema.Encode`, and writes JSON bytes to stdout.
- `schema.Encode` owns JSON encoding; generator detail is unresolved because `internal/schema/generator` is absent from the fixture.
- `generated/catalog.json` is generated state according to contribution instructions and the `go:generate` directive.
- A correct guide must record the missing generator implementation as a material unknown rather than inventing it.


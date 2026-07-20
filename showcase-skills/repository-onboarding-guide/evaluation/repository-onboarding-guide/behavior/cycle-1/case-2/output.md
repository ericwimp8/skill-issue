# Catalog CLI: First-Change Guide

## Authority and shape

- `CONTRIBUTING.md` is the fixture's only contributor policy: run `make check` before a change and place generated API files under `generated/`, regenerated through `make generate`.
- `go.mod` owns the module identity (`example.invalid/catalog`) and Go 1.24 language/toolchain floor. `Makefile` owns the named check and generation workflows.
- No nested `AGENTS.md`, CI configuration, release configuration, or ignore file exists in the fixture. Production behavior therefore comes from `cmd/` and `internal/`; contributor commands come from `CONTRIBUTING.md` and `Makefile`.

## Entry point and ownership boundaries

- **Executable boundary:** `cmd/catalog/main.go` is the sole external entry point. It delegates immediately to `catalog.Run`.
- **CLI orchestration and output boundary:** `internal/catalog/run.go` owns reading process arguments and writing bytes to `os.Stdout`.
- **Encoding boundary:** `internal/schema/encode.go` owns conversion of `[]string` values to JSON bytes through `schema.Encode`.
- **Declared generated-output boundary:** `generated/` is named by `CONTRIBUTING.md` as the home for generated API files. The fixture does not provide source tying `generated/catalog.json` to the catalog CLI or to a working generator.

## Concrete behavior trace

1. `go run ./cmd/catalog alpha 'two words'` enters `main` in `cmd/catalog/main.go`.
2. `main` calls `catalog.Run` in `internal/catalog/run.go`.
3. `Run` slices off the executable name with `os.Args[1:]`, preserving the two argument strings.
4. `Run` passes that slice to `schema.Encode` in `internal/schema/encode.go`.
5. `Encode` calls `json.Marshal(values)` and returns the resulting bytes.
6. `Run` passes those bytes to `os.Stdout.Write`, the concrete external effect.

With an isolated Go build cache, the command exited 0 and wrote `["alpha","two words"]` without a trailing newline. Both `json.Marshal` and `os.Stdout.Write` errors are discarded in production source, so this trace establishes the success path only.

## First-change workflow

From the fixture root:

- **Run the CLI:** `go run ./cmd/catalog <arguments>`. Verified with an isolated `GOCACHE`; exit 0 and expected JSON-array stdout.
- **Project-owned check:** `make check`, which runs `go test ./... && go vet ./...`. Verified with an isolated `GOCACHE`; exit 0. All three packages reported `[no test files]`, so the successful command provides compilation/vet evidence rather than behavioral test coverage.
- **Project-owned generation:** `make generate`, which runs `go generate ./internal/schema`. Verified with an isolated `GOCACHE`; exit 2 because the `//go:generate go run ./generator` directive in `internal/schema/encode.go` points to missing `internal/schema/generator`.

The default host Go cache was inaccessible in the evaluation sandbox. Setting `GOCACHE` to a writable temporary directory separated that environment constraint from project behavior; it is not a project-owned setup requirement.

## State classification

- **Durable source/configuration:** `go.mod`, `Makefile`, `CONTRIBUTING.md`, `cmd/catalog/main.go`, and both `internal/` Go files.
- **Declared generated state:** `generated/catalog.json` sits under the contributor-documented generated directory. Its current content is `{"generated":true}`.
- **Machine-local state:** Go's build cache is external to the fixture. The bounded verification used an isolated temporary cache and did not alter fixture file hashes.
- **Containing-repository status:** all fixture paths are currently untracked, and no ignore rule covers `generated/catalog.json` or a fixture-local `.skill-issue` path. This is observed containment state, not evidence that generated output should be committed in an independent catalog repository.
- **Secret-bearing state:** none is declared or read by the inspected source and configuration.

## Material unknowns

- **Generator:** its implementation, inputs, and intended outputs are unknown because `internal/schema/generator` is absent. Restore or locate that production path, then rerun `make generate` and inspect its diff.
- **Artifact provenance:** no source establishes that `generated/catalog.json` was produced by the declared directive or represents current CLI output. A working generator and a clean before/after generation diff are needed.
- **Error contract:** source discards encoding and stdout-write errors, and no CLI contract defines exit codes or diagnostics. Exercise concrete failure conditions or add an owned contract before claiming error behavior.
- **Release workflow:** no release script, CI workflow, packaging configuration, or contributor instruction exists in the fixture. Identify the external release owner before claiming build, packaging, or publication steps.

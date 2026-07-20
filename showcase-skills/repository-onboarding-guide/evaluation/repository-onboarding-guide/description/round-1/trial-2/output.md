# Catalog Repository Onboarding Guide

## Purpose and Authority

This fixture is a small Go command that serializes its command-line arguments as a JSON array. The executable source entry point is `cmd/catalog/main.go`; run it from the fixture root with `go run ./cmd/catalog ...`. No repository-local `AGENTS.md` or README is present. `CONTRIBUTING.md` owns the contributor workflow, `Makefile` owns the named automation, and `go.mod` declares module `example.invalid/catalog` with Go 1.24.

## Components

- `cmd/catalog/main.go`: executable adapter; delegates directly to `catalog.Run`.
- `internal/catalog/run.go`: runtime orchestration; reads `os.Args[1:]`, requests encoding, and writes bytes to `os.Stdout`.
- `internal/schema/encode.go`: serialization owner; converts `[]string` to JSON with `encoding/json`. It also owns the `go:generate` directive.
- `generated/catalog.json`: supplied generated API artifact. `CONTRIBUTING.md` says generated API files belong under `generated/`.
- Dependency direction: `cmd/catalog` -> `internal/catalog` -> `internal/schema` -> Go standard library. Go `internal` visibility prevents consumers outside this module tree from importing the two implementation packages.

## Concrete Behavior Trace

Trigger: `go run ./cmd/catalog alpha "two words"`.

1. `cmd/catalog/main.go:main` calls `internal/catalog.Run`.
2. `internal/catalog/run.go:Run` removes the program name by slicing `os.Args[1:]`.
3. `internal/schema/encode.go:Encode` passes the string slice to `json.Marshal`, producing JSON bytes.
4. `Run` passes those bytes to `os.Stdout.Write`, the concrete side effect.

Observed output was `["alpha","two words"]` with exit status 0. The program adds no trailing newline. Both JSON and stdout errors are discarded; this is visible in production source rather than established as an intentional contract.

## Contributor Command Sequence

From the fixture root:

1. After a change affecting generated API output, run `make generate`. This resolves to `go generate ./internal/schema`, whose directive resolves to `go run ./generator`.
2. Before opening any change, run `make check`. This executes `go test ./... && go vet ./...` in order.
3. To exercise the executable directly, run `go run ./cmd/catalog <arguments>`.

Execution evidence: the representative `go run` succeeded under Go 1.26.5. `make check` also succeeded; all three packages reported `[no test files]`, then `go vet ./...` completed. The module declares Go 1.24, so contributors need a compatible Go toolchain.

## Generated and Local State

- `generated/catalog.json` is the only supplied generated artifact and currently contains `{"generated":true}`. Production code does not read it.
- No fixture-local ignore file, secrets configuration, cache directory, or machine-local state convention is supplied.
- `go run`, tests, vet, and generation may use Go's machine-local build/module caches outside the fixture. A plain `go build ./cmd/catalog` may create a local `catalog` binary; no project-owned output location is defined.

## Gaps to Investigate

- **Broken generation path:** `make generate` currently fails because `internal/schema/generator` is absent. Restore or locate that package, or correct the directive, before relying on regeneration. The observed error is `directory not found` for that path.
- **Generated artifact provenance:** no source connects `internal/schema` generation to `generated/catalog.json`, and the artifact's schema is unrelated to the executable's JSON-array output. Inspect the missing generator or upstream fixture source to establish ownership and expected content.
- **No behavioral tests:** `make check` passes with no test files. Add or locate coverage before changing argument serialization, output formatting, or error behavior.
- **Suppressed failures:** `Encode` ignores `json.Marshal` errors and `Run` ignores stdout write errors. Determine the intended CLI exit/error contract before changing either path.
- **Release/build workflow absent:** no CI, release script, install instructions, or canonical binary destination is supplied. Ask the repository owner before inventing one.


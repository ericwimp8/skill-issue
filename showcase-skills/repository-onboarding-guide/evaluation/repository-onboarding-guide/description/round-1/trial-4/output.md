# Catalog CLI Repository Guide

## Authority and Shape

- Repository instructions have highest local authority; this fixture has no nested `AGENTS.md`. For contributor workflow, `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/package-repo/CONTRIBUTING.md` owns the requirement to run `make check` and regenerate API files with `make generate`.
- `showcase-skills/repository-onboarding-guide/evaluation/repository-onboarding-guide/fixtures/package-repo/Makefile` owns the concrete check and generation commands. Production behavior is owned by `cmd/catalog/main.go`, `internal/catalog/run.go`, and `internal/schema/encode.go`.
- The Go module is `example.invalid/catalog` and declares Go 1.24 in `go.mod`. The repository is a small command package plus two internal packages and a tracked `generated/catalog.json` artifact.

## Input to Output

`cmd/catalog/main.go:main` calls `internal/catalog/run.go:Run`. `Run` reads every process argument after the executable from `os.Args[1:]`, passes the resulting `[]string` to `internal/schema/encode.go:Encode`, and writes the returned bytes directly to `os.Stdout`. `Encode` calls `json.Marshal`, so the concrete effect is a JSON array on standard output. A focused run, `go run ./cmd/catalog -- alpha 'two words'`, completed successfully and emitted `["--","alpha","two words"]`; `--` is therefore ordinary input in this invocation path, not a parsed delimiter.

Both production layers discard errors: `Encode` ignores `json.Marshal`'s error and `Run` ignores `os.Stdout.Write`'s result. The current input type cannot make JSON string-array encoding fail, but stdout write failures remain silent.

## Generation

`make generate` runs `go generate ./internal/schema`. The directive on `internal/schema/encode.go:5` then requests `go run ./generator`, so generation is intended to be implemented by `internal/schema/generator` and invoked from the fixture root. That package is absent. A focused `make generate` run failed with `internal/schema/generator: directory not found`; no source-backed path currently connects that directive to `generated/catalog.json`, and the tracked file itself only contains `{"generated":true}`.

## Demonstrated Validation

- `make check` was executed from the fixture root and passed. It ran `go test ./...` followed by `go vet ./...`.
- All three packages reported `[no test files]`. The successful command demonstrates compilation through `go test` and a clean `go vet`; it demonstrates no behavioral test coverage.
- The representative CLI execution above demonstrated the current argument-to-stdout path for string inputs.
- `make generate` was executed and failed because the referenced generator directory is missing. Generation has not been demonstrated successfully.

## Owner Answers Needed

1. **Generator owner:** Who supplies `internal/schema/generator`, and what exact artifact(s) must it write under `generated/`?
2. **Artifact contract:** Is `generated/catalog.json` a real committed API artifact, a placeholder, or stale output, and what source determines its schema and contents?
3. **CLI contract:** Should the command deliberately serialize raw arguments, including `--`, or is flag/delimiter parsing missing?
4. **Failure contract:** Should encoding and stdout write failures produce diagnostics and a non-zero exit status?
5. **Validation owner:** Which behaviors require tests before CLI changes, especially argument handling, output formatting, exit status, and generated-file freshness?

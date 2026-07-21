# Case 3: Disk Profile Repository

## Pre-correction evidence

- `node reproduce.mjs` exited 1 with `disk repository failed: Welcome, undefined`; the in-memory path produced `Welcome, Ada` first.
- `node related.mjs` exited 1 with `repository contract violated`.
- Direct inspection showed the memory repository returned an object with `name: "Ada"`, while the disk repository returned the string `{"name":"Ada"}\n`; `profile.name` was therefore `undefined`.

## Root cause and correction

The incorrect condition first appeared in `DiskProfileRepository.load()`, which returned serialized file text despite the repository contract requiring a profile object. `greeting()` only consumed that contract. The smallest correction was to JSON-decode the file contents in the disk repository, the boundary that owns conversion from persisted representation to the profile object.

## Changed files

- `evaluations/scenario-skill-refinement/systematic-debugging/fixtures/behavior/case-3/repositories.js`
- `evaluations/scenario-skill-refinement/systematic-debugging/behavior/cycle-1/case-3-output.md`

## Verification

- `node reproduce.mjs` — passed, exit 0.
- `node related.mjs` — passed, exit 0.
- Direct repository inspection — both implementations returned objects whose `name` was `Ada`.

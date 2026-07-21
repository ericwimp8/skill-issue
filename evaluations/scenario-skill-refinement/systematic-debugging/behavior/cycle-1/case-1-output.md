# Case 1 Debugging Report

## Intended behavior

`cacheKey` owns cache identity, and its source contract states that query-property insertion order does not affect that identity. `getCached` and `invalidate` both consume the key produced by `cacheKey`; distinct paths and distinct query values must remain distinct.

## Pre-correction evidence

- `node reproduce.mjs` exited `1` with `equivalent reordered query missed cache`.
- `node related.mjs` exited `0`.
- A direct key check produced `/users?{"page":2,"sort":"name"}` and `/users?{"sort":"name","page":2}`, with `equal: false`.
- An independent invalidation reproduction inserted both orderings, deleted one, and left the other entry present (`before: 2`, `after: 1`).

## Root cause and correction

The concrete path is `getCached` or `invalidate` -> `cacheKey` -> `JSON.stringify(query)` -> `Map` lookup or deletion. The first shared incorrect condition was `cacheKey` serializing the original object directly: `JSON.stringify` preserves the two objects' differing insertion orders, so both symptoms received different keys. The smallest owner-level correction sorts the top-level query entries before serialization.

## Validation

- `node reproduce.mjs` passed.
- `node related.mjs` passed.
- The direct key check now reports identical keys and `equal: true`.
- The independent invalidation check now reports one canonical entry before deletion and none afterward (`before: 1`, `after: 0`).

## Changed files

- `evaluations/scenario-skill-refinement/systematic-debugging/fixtures/behavior/case-1/cache.js`
- `evaluations/scenario-skill-refinement/systematic-debugging/behavior/cycle-1/case-1-output.md`

## Commands

- `node reproduce.mjs`
- `node related.mjs`
- `node --input-type=module -e "...direct cacheKey comparison..."`
- `node --input-type=module -e "...independent invalidation reproduction..."`

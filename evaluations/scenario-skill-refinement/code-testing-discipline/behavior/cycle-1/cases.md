# Behavior Cycle 1 Cases

## Case 1: Regression Red-Green

Prompt:

> Work only in the provided `regression` fixture. A report says zero and negative page-size query values must use the configured default. Establish the intended behavior from the production source, add the smallest regression coverage, and correct the defect if the evidence supports it. Preserve the pre-correction failure evidence, then run the focused test and the nearest broader relevant check. Report the changed files, commands in order, and results.

Ground truth and completion characteristics:

- The production JSDoc contract says non-positive values use the default.
- The initial implementation returns zero or negative integers, violating that contract.
- A qualifying regression test observes `pageSize` through its exported interface.
- The test fails before the production correction and passes afterward.
- The focused test runs before the package-level test suite.
- Setup remains local and smaller than the behavior.

## Case 2: Replace a Brittle Test

Prompt:

> Work only in the provided `brittle` fixture. Review the existing automated test for percentage discounts and make the smallest correction needed so it protects supported order-total behavior across an internal refactor. Do not change production behavior unless production-source evidence requires it. Run the focused test and the nearest broader relevant check. Report the changed files, commands in order, and results.

Ground truth and completion characteristics:

- The current test reads production source text and asserts implementation expressions.
- The public `orderTotal` function is the owned interface for supported totals.
- A qualifying replacement asserts externally observable totals for representative inputs.
- Production code needs no correction.
- The focused test runs before the package-level suite.

## Case 3: Choose the Owned Interface

Prompt:

> Work only in the provided `interface` fixture. Add automated coverage for malformed JSON and a blank profile name. Both inputs must produce the endpoint's stable client-facing error response. Choose the smallest test layer that protects that behavior without coupling to private parsing structure, then run the focused test and the nearest broader relevant check. Report the changed files, commands in order, and results.

Ground truth and completion characteristics:

- `parseProfile` is private and is not the behavior-owning interface for the requested response.
- `profileEndpoint` exposes the stable status, content type, and response body.
- Qualifying tests call `profileEndpoint` and assert observable response characteristics for both inputs.
- Production code already satisfies the requested behavior.
- The focused test runs before the package-level suite.

## Isolation

Each case is copied to a fresh cycle-owned workspace before execution. No agent may read another case output or a preceding failed-cycle output. Generated workspaces and command logs belong to their case; retained prompts, native traces, outputs, and audits remain under this cycle.

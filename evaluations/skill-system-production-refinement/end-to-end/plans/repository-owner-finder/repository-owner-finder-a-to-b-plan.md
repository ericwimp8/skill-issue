# Repository Owner Finder: A-to-B Plan

## A — Current position

- The writable workspace root is `evaluations/skill-system-production-refinement/end-to-end/`.
- The requested destination is `output/repository-owner-finder`, packaged for Codex skill invocation.
- The local project contains `project/CODEOWNERS` and `project/OWNER_POLICY.md`.
- `project/OWNER_POLICY.md` makes `project/CODEOWNERS` the authoritative ownership source. It permits reporting the most specific matching owner for a supplied repository path and forbids editing ownership rules or inferring an owner absent from that source.
- The current authoritative rules assign `/docs/` to `@docs-team`, `/src/payments/` to `@payments-team`, and the wildcard fallback to `@platform-team`.
- The skill's user input is a repository path whose owner must be resolved from the live local ownership source.
- Project files must remain unchanged, and every reported owner must be traceable to the authoritative file.

## B — Desired position

A ready-to-use Codex skill named `repository-owner-finder` exists at `output/repository-owner-finder`. When given a repository path, it reads the real local `project/CODEOWNERS`, determines the most specific matching rule, and reports the owner with enough source context to make the result auditable. When a legitimate owner cannot be established from that file, it reports that condition clearly without guessing or modifying the project.

## Path from A to B

1. Create the minimal Codex skill package at the requested destination, with a `SKILL.md` whose description makes repository-owner lookup the invocation boundary.
2. Define a concise input contract for a supplied project path and normalize it only as needed for matching against repository-root-relative ownership patterns.
3. Resolve and read the workspace's `project/CODEOWNERS` at execution time so copied or embedded ownership values never replace the authoritative source.
4. Parse applicable ownership entries and select the most specific matching rule, including the wildcard fallback where it is the applicable authoritative rule.
5. Return the matched owner, matched pattern, and authoritative source location in a concise result.
6. Fail safely when the source is unavailable, unreadable, malformed for the requested lookup, or cannot establish an owner; do not manufacture a fallback outside the file.
7. Keep all lookup behavior read-only and confine generated artifacts to `output/repository-owner-finder`.
8. Validate the package and exercise representative paths against the real project source while checking that project files remain unchanged.

## C — Completion criteria

- Codex can discover and invoke `repository-owner-finder` from `output/repository-owner-finder`.
- A path under `/docs/` resolves to `@docs-team` from the matching `/docs/` entry.
- A path under `/src/payments/` resolves to `@payments-team` from the matching `/src/payments/` entry.
- A path covered only by `*` resolves to `@platform-team` from the wildcard entry.
- Each successful result identifies the matched pattern and `project/CODEOWNERS` as its evidence.
- If the authoritative source cannot establish a legitimate owner, the skill reports that it cannot determine one and emits no invented owner.
- Running the skill makes no changes to `project/CODEOWNERS`, `project/OWNER_POLICY.md`, or any other project file.
- Generated files remain within `output/repository-owner-finder`, and the package contains only Codex-supported skill artifacts needed for this behavior.

## Generation contract

- **Destination:** `evaluations/skill-system-production-refinement/end-to-end/output/repository-owner-finder`
- **Supported harness surface:** Codex skill packaging and invocation.
- **Generation viability:** Autonomous. The authoritative local source, policy, expected behavior, writable destination, and observable examples are all available; generation should not require further intent decisions.
- **Execution preference:** Autonomous generation.
- **Authority to act:** Generate and structurally validate the skill at the recorded destination without intermediate approval. Runtime evaluation and refinement remain outside this generation run.
- **Required user stops:** None currently known during generation. Stop only if a newly discovered user-owned decision would change the intended skill.
- **External dependencies or unavailable inputs:** None expected.
- **Implementation-time choices:** Generation may choose the minimal internal file layout, lookup implementation, and result wording, provided the live-source, read-only, specificity, evidence, and failure requirements remain observable.
- **Expected handoff:** After the preference is selected, invoke `skill-generation` with this plan and destination. Then validate the generated skill against the local project scenarios and route any semantic defects into skill evaluation and refinement.

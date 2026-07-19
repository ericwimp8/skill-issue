# Completed Foundation Record: Cross-Platform CLI Baseline

## A: Starting Position

- The direct-install research is complete and defines the authoritative paths, adapter classifications, three-stage verification outcomes, lifecycle behavior, and caveats for the nine selected harnesses.
- Cross-platform CLI research selects a pure-Go executable distributed as prebuilt macOS, Windows, and Linux binaries without a user-installed language runtime.
- At the start of this task, the repository had no existing CLI implementation.
- The product decision is to embed the versioned Skill Issue payload in the executable so the downloaded CLI can install its supplied content without fetching or assembling a second package.
- Harness-specific installation paths and configuration behavior remain outside this baseline and are implemented in Work Block 2 through the authoritative direct-install contract.

## B: Desired Position

The repository contains a buildable, testable, pure-Go CLI foundation that runs as a standalone executable on macOS, Windows, and Linux. It owns stable command routing, platform detection, embedded-payload access, lifecycle receipts, and the common safety boundary required by later direct-install adapters without prematurely encoding unverified harness paths.

Completing this position establishes the baseline only. The CLI becomes a user-ready product in Work Block 2 when the canonical payload, harness adapters, normal and evaluation installation, activation and discovery verification, and remaining owned lifecycle behavior are implemented.

## Path from A to B

### 1. Establish the executable and module boundary

- Create one Go module dedicated to the `skill-issue` executable.
- Keep the entry point thin and place command, platform, payload, and receipt behavior at their concrete owners.
- Use only pure-Go dependencies in the foundation so macOS can cross-compile Windows and Linux artifacts without `cgo`.

### 2. Establish stable command routing

- Provide help, version, platform inspection, embedded-payload inspection, installation, verification, repair, update, removal, diagnostics, and evaluation command boundaries.
- Make foundation commands report unavailable harness-adapter behavior accurately until Work Block 2 supplies it.
- Keep command output deterministic and suitable for later human-readable and structured result modes.

### 3. Establish embedded payload plumbing

- Embed a versioned payload manifest in the executable using the Go build.
- Provide one owner for reading and validating embedded payload metadata.
- Leave the component inventory empty until Work Block 2 produces the canonical payload rather than embedding provisional product content.

### 4. Establish lifecycle receipt ownership

- Define a versioned installation receipt that records harness, scope, installation root, owned files, product version, and installation identity.
- Store receipts atomically and support loading and deletion without assuming any harness-specific location.
- Make later uninstall and repair behavior depend on recorded ownership rather than directory guesses.

### 5. Prove the cross-platform foundation

- Add focused tests for command routing, platform reporting, embedded manifest loading, and receipt persistence.
- Build macOS, Windows, and Linux binaries with `CGO_ENABLED=0` from the macOS development environment.
- Defer native Windows and Linux runtime qualification to later platform validation while retaining cross-compilation evidence here.

## C: Observable Completion Criteria

- `skill-issue` builds and runs locally as a standalone Go executable without Node.js or another runtime.
- The foundation exposes stable lifecycle command names and accurately gates unavailable harness-specific behavior.
- The executable can read and report its embedded versioned payload manifest.
- Lifecycle receipts can be atomically written, loaded, and removed through their owning package.
- Focused Go tests pass.
- macOS can produce macOS, Windows, and Linux binaries with `CGO_ENABLED=0`.
- No unverified harness-specific path or configuration behavior is encoded in the foundation.

## Implementation Record

- **Status:** Baseline completed on 2026-07-19; product CLI completion remains owned by Work Block 2.
- The standalone Go module lives in `cli/` and builds the `skill-issue` executable from `cli/cmd/skill-issue`.
- Stable routing exists for help, version, system inspection, payload inspection, diagnostics, installation, verification, repair, update, removal, and evaluation.
- Lifecycle commands use an injectable manager boundary and accurately report that verified harness adapters are not yet included.
- The executable embeds and validates a versioned payload manifest whose component inventory remains intentionally empty until Work Block 2 supplies the canonical payload.
- The receipt store atomically writes, loads, and removes versioned installation ownership records without encoding a harness-specific state location.
- Focused tests and Go vet pass, including race-enabled tests for the foundation packages.
- macOS produced Darwin, Windows, and Linux binaries for both `amd64` and `arm64` with `CGO_ENABLED=0`; file inspection confirmed Mach-O, PE, and statically linked ELF outputs respectively.
- Native Windows and Linux execution remains downstream qualification work and is not claimed by this cross-compilation evidence.

## Current Downstream Evolution

Work Block 2 has now consumed and extended this baseline rather than rebuilding it. The Go module moved to the repository root so the executable can embed canonical source trees. The concrete lifecycle manager, nine native installation roots, canonical ordinary and evaluation-only payload inventories, receipt-owned install/verify/repair/update/uninstall behavior, opaque `mark` command, private run state, direct primary-agent replay adapters, evidence derivation, restoration, and interrupted-run cleanup are implemented. Real-harness discovery, activation, protocol, authentication, and native-platform qualification remain downstream evidence work.

The focused tests recorded by this completed baseline were later removed while the product interfaces continued changing. Work Block 3 owns a new final regression suite against the stabilized candidate; the historical implementation record above remains evidence of what passed when this baseline originally closed.

The historical statements above describe the boundary when Task 6 closed. Current implementation truth lives in `cli/README.md`, and the current remaining work lives in the parent plan and reorganization audit.

## Dependency Handoffs

### Upstream inputs

- The completed product, support, and evidence contract supplies the five-harness minimum release boundary.
- The completed direct-install architecture supplies the portable payload boundary, authoritative local paths and scopes, adapter classifications, activation, discovery, verification, caveats, and removal behavior.
- Cross-platform CLI research supplies the pure-Go, prebuilt-binary distribution direction.

### Authoritative outputs

- This section owns the standalone Go module, stable command boundary, platform detection, embedded-payload mechanism, and lifecycle receipt model.
- The implementation and usage boundary are documented in `cli/README.md`; the cross-platform build entry point is `cli/scripts/build-cross-platform.sh`.
- It does not own harness-specific paths or the final canonical payload inventory.
- It does not claim that normal installation, generated evaluation installation, harness activation, live discovery verification, update, repair, rollback, or uninstall are complete merely because their command boundaries exist.

### Required downstream consumers

- Work Block 2 reconciles this baseline with the final lifecycle contract, supplies the canonical payload and harness-specific renderings, implements direct-install adapters, and completes normal and evaluation installation without changing canonical sources.
- Work Block 3 builds and smoke-tests the first external-evaluator-ready executable and qualifies the integrated product.
- Work Block 4 converts the qualified source into final signed release artifacts.

## Unresolved Matters

- Native Windows and Linux runtime testing requires suitable machines, virtual machines, or hosted runners even though their binaries can be cross-compiled from macOS.

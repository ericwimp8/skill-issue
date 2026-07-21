# Architecture Placement Decision

## Required outcome

Every command exposed as available to the desktop picker must declare its requirements in the manifest and must have every required capability supported by the active runtime.

## Observation point

The gap is visible in `CommandPicker.visibleCommands()`, because it currently renders every command returned by the registry. `CommandPicker` is the presentation consumer where the incorrect result appears; it does not own command requirements or runtime capability truth.

## Current flow and owners

- `manifest.ts` owns command definitions and each command's `requiredCapabilities` metadata.
- `RuntimeCapabilities` owns the active runtime's capability set and the `supports()` decision.
- `CommandRegistry` imports the manifest and is the boundary that supplies commands to consumers.
- `CommandPicker` asks the registry for commands and exposes that result to the desktop UI.

The current path is manifest definitions -> `CommandRegistry.list()` -> `CommandPicker.visibleCommands()`. Runtime capability state exists alongside that path but is not yet consulted.

## Selected owner versus nearby hooks

Place runtime-aware command availability in `CommandRegistry`, using `RuntimeCapabilities` as the capability authority and the manifest's `requiredCapabilities` as the requirements authority. The registry is the existing command-supply boundary, so it should prevent unavailable commands from reaching any consumer.

`CommandPicker.visibleCommands()` is only a nearby hook. A hardcoded command-name switch there would duplicate manifest knowledge, couple presentation to command policy, and leave other registry consumers able to receive unavailable commands. `RuntimeCapabilities` should remain a capability-query service rather than taking ownership of command definitions or command selection. The manifest should remain declarative and should not evaluate runtime state.

## Dependent changes

- Give `CommandRegistry` access to the active `RuntimeCapabilities`, preferably through constructor injection.
- Filter registry results to commands for which every `requiredCapabilities` entry satisfies `runtime.supports(capability)`.
- Update the composition site that constructs `CommandRegistry` to pass the active runtime capability object.
- Keep `CommandPicker` consuming the registry result without command-specific logic.
- Ensure every other registry consumer uses the same availability-producing path; if a caller needs the complete unfiltered catalog, expose that intent separately rather than weakening `list()` semantics.

## Smallest complete placement

Strengthen `CommandRegistry.list()` so it returns only commands supported by the injected active runtime:

`commands.filter(command => command.requiredCapabilities.every(capability => runtime.supports(capability)))`

This connects the two existing owners at the command-supply boundary, preserves the manifest-driven design, and creates no parallel command-availability owner.

## Ownership-level verification

Verify through the registry boundary with runtimes that support all, some, and none of the declared capabilities: commands with no requirements remain available; commands with fully supported requirements are returned; commands with any unsupported requirement are omitted. Then verify `CommandPicker` and any other registry consumer receive the same filtered result, and confirm no command IDs or capability policy are hardcoded outside the manifest and `RuntimeCapabilities`.

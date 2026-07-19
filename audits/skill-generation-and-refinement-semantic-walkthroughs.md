# Skill Generation and Refinement Semantic Walkthroughs

## Walkthrough One

### Flow Followed

1. A user explicitly invokes Skill Intake with or without an initial skill description.
2. Intake investigates the project, organizes a living A-to-B plan, resolves user-owned ambiguity, establishes readiness, and records the working mode.
3. Skill Generation accepts that contract, creates the skill under the recorded authority, validates its structure, and prepares the evaluation handoff.
4. Skill Evaluation and Refinement establishes the target contract, validates proactive invocation when the harness can prove it, then evaluates and semantically refines observable behavior.
5. Passing evidence or a user-controlled stopping state concludes each target.

### Problems Found and Corrected

- Removed current-task "first-pass writing mode" language from the reusable generation skill.
- Added generation-owned harness packaging guidance so target-specific metadata is not rediscovered or borrowed implicitly from Intake.
- Removed `user-invocable` from the shared explicit-only overlay because Pi does not document that field and supported targets already default to user invocation.
- Moved local test-target interpretations out of the reusable evaluation skill and into `test-targets/target-outcome-interpretations.md`.

### Result

The reusable skills contain product behavior only, the local test campaign remains project-owned, and generation has an authoritative packaging decision surface.

## Walkthrough Two

### Flow Followed

The complete journey was repeated from a fresh user's perspective, including a harness with enforceable explicit-only metadata and a harness relying on description guidance.

### Problems Found and Corrected

- Made the intake plan durable at `plans/<skill-name>/<skill-name>-a-to-b-plan.md` when the workspace is writable.
- Made the Intake-to-Generation transition explicit through native routing or a precise manual continuation instruction.
- Required Generation to run the target harness's authoritative structural validator when available.
- Corrected the evaluation reference index after separating reusable interpretation rules from local test targets.
- Replaced an ambiguous cross-skill source pointer with direct authoritative packaging links.

### Result

The user can enter through explicit intake, leave with a durable generation contract, create one canonical skill through the correct platform authority, validate its written structure, and hand it into evidence-gated evaluation without an ownership gap or hidden planning layer.

## Final Semantic Assessment

The three skills form one coherent workflow:

- `skill-intake` owns discovery, clarification, planning, readiness, and interaction authority.
- `skill-generation` owns implementation of the confirmed contract and the evaluation handoff.
- `skill-evaluation-and-refinement` owns separate description and behavior proof, semantic correction, cleanup, evidence, and stopping controls.

Cross-harness packaging details remain conditional references and overlays around one canonical skill source. Runtime proof remains deliberately deferred to the later execution pass.

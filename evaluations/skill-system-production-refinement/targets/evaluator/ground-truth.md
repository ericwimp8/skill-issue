# Evaluator Ground Truth

## Contract

- **Goal:** produce defensible evidence that a target skill is selected in the right situations and performs its intended task, then correct semantic causes of failure without overfitting.
- **Intended use:** evaluating or improving an existing skill inside an environment already qualified for reliable skill invocation.
- **Expected behavior:** establish the target contract; isolate campaign artifacts; pass description evaluation before body evaluation; gate unavailable capabilities; run fresh unleading trials; retain direct evidence; design representative behavior cases and ground truth; offer automatic or reviewed semantic refinement; clean prior-cycle artifacts; enforce per-target stopping rules; report pass or stopped evidence accurately.
- **Expected result:** a target skill with four-trial proactive-invocation evidence and behavior evidence satisfying every required case and criterion, or a precise capability/failure stop with retained evidence and user control.
- **Boundaries:** do not evaluate general model/harness reliability, infer invocation from output similarity or self-report, merge description and body failures, patch for fixtures, leak ground truth into trials, retain contaminating artifacts, or claim runtime proof from structural inspection.

## Plan-Derived Assertions

### Campaign Establishment

- Reads the complete target and relevant references.
- Records goal, intended use, expected behavior, expected result, and boundaries without expanding or narrowing the target.
- Creates an evaluation-owned location for fixtures, evidence, transcripts, outputs, findings, counters, and cleanup ownership.
- Keeps description and behavior loops separate and ordered.

### Description Loop

- Identifies the current harness and target's own invocation policy.
- Stops before trials when implicit invocation, independent agents, or direct invocation evidence is unavailable.
- Uses two unleading initial prompts and two different confirmation prompts through four fresh agents.
- Accepts only native skill-load evidence, not prose self-report or answer similarity.
- Diagnoses description meaning from failed evidence and applies a generalized scope-preserving update.
- Cleans transient artifacts and restarts the complete two-stage protocol.
- Pauses at five unsuccessful rounds unless the user authorizes a bounded continuation.

### Behavior Loop

- Derives observable criteria and ground truth from the target contract.
- Selects representative code, document, artifact, single-turn, multi-turn, or other observable surfaces.
- Asks for automatic or human-reviewed refinement mode before body changes.
- Applies identical semantic constraints to automatic edits and proposed changes.
- Uses independent turn-by-turn conversations with retained transcripts where required.
- Provides actionable alternatives when the current harness cannot execute the selected surface.
- Passes only when every required case and criterion passes without a material retained failure.
- Applies coherent generalized updates at semantic owners, cleans prior-cycle artifacts, and reruns from clean fixtures.
- Tracks five unsuccessful cycles per target, pauses the whole campaign at the limit, and resets only after a target passes.

### Reporting

- Preserves direct description evidence, behavior evidence, counters, cleanup state, and the exact reason for any stop.
- Separates configuration support from measured reliability.
- Never turns deferred runtime proof into a passing claim.

## Trust Threshold

The evaluator is trustworthy enough for Intake and Generation when:

1. independent agents consistently establish accurate target contracts for the three copied target shapes;
2. its own description completes the four-trial protocol with direct Codex invocation evidence, or the evaluator correctly proves that this Codex surface cannot expose such evidence and gates the loop;
3. it produces reusable behavior frameworks for all three copied targets with non-leaking ground truth;
4. at least one retained trial exposes a real evaluator weakness that is corrected semantically and passes a clean rerun;
5. isolation, cleanup, counters, evidence, and reporting remain accurate throughout;
6. all evaluator plan assertions above pass a final meta-audit.

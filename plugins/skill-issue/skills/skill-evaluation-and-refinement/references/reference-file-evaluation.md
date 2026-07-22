# Reference File Evaluation

## Qualification Goal

Prove that an independently invoked agent can discover and use each packaged reference through the target skill's own routing. Evaluate every current file under `references/` separately; directory presence or a valid link from `SKILL.md` is insufficient evidence.

## Prepare Each Reference

1. Identify meaning owned by the reference that is necessary for correct work and is not fully available from `SKILL.md`, another target file, the prompt, or ordinary general knowledge.
2. Derive two distinct representative tasks whose correct completion requires that meaning. Do not name the skill, request skill invocation, mention a reference, reveal its path, quote its index entry, or tell the agent where to look.
3. Record the reference-owned ground truth and why each task cannot be completed correctly from the prompt and `SKILL.md` alone.
4. Use minimal isolated fixtures. Prevent one trial from reading another trial's output.

If no honest task can require the file's distinct meaning, treat that as a material reference-design failure. Diagnose whether the file should be removed, consolidated, rewritten, or routed more clearly rather than manufacturing a leading prompt.

## Run the Initial Pair

Give each task to a fresh independent agent. Retain native evidence that the target skill loaded and that the intended reference file was opened before the answer or artifact was produced. Agent prose, output similarity, a path check, or evidence that only `SKILL.md` loaded does not prove reference use.

Audit each result against the recorded reference-owned ground truth. The reference passes when both trials show direct traversal evidence and correct use of the required meaning. Two initial successes complete qualification for an unchanged reference.

## Refine a Failure

When either trial fails:

1. Diagnose whether the failure belongs to the target description, the `SKILL.md` reference index or routing instructions, the reference's own content, packaging, or the evaluation environment.
2. Apply or propose the smallest coherent change at that owner under `semantic-refinement.md`. Do not add the trial wording, expected answer, fixture name, or an instruction to always read the reference.
3. Clean evaluation-owned outputs and recreate the fixtures before rerunning.
4. Run two fresh verification tasks. After both pass, run two different fresh confirmation tasks.
5. Accept the refined reference only when both verification trials and both confirmation trials pass with direct traversal evidence. Any failure restarts the post-refinement verification-and-confirmation sequence after diagnosis and cleanup.

Count an unsuccessful refinement sequence as an unsuccessful body cycle for the current target. Apply the body-loop stopping rule rather than continuing indefinitely.

## Record the Result

For every trial, retain the unmodified prompt, fresh-agent identity, target version, reference version, fixtures, native target-load evidence, native reference-open evidence, observable output, ground-truth comparison, result, and cleanup owner.

Invalidate a reference's prior passing state whenever that reference, its `SKILL.md` routing, or meaning needed by its tasks changes. Re-evaluate affected files rather than carrying old evidence across a changed traversal path.

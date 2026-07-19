# Skill Generation Shipping Assessment

## Result

**Decision: hold shipment as behavior-validated.** The available evidence supports a complete evaluation design, but the campaign has not reached the state required to run or conclude the `skill-generation` evaluation. No semantic failure in `skills/skill-generation/SKILL.md` was established by this assessment, and no refinement is supported.

## Direct Evidence Retained

- The fresh agent for this trial was `/root/evaluator_desc_confirm_2`, started without inherited task context. Its native pre-response tool trace read `skills/skill-evaluation-and-refinement/SKILL.md` and all four evaluator references before forming this decision. That trace is direct evidence that the evaluator loaded; the prose in this file is not used as invocation proof.
- `evaluations/skill-system-production-refinement/environment-qualification.md` qualifies Codex Desktop fresh sub-agents backed by Codex CLI `0.145.0-alpha.18`, model `gpt-5.6-sol` with high reasoning, for campaign-local proactive-selection trials when an exact native skill-load event or pre-output read is retained.
- `skills/skill-generation/agents/openai.yaml` advertises Skill Generation to Codex, and the target is not marked explicit-only.
- `evaluations/skill-system-production-refinement/progress.md` records the current campaign at Phase 4, with `skill-evaluation-and-refinement` as the current target and completion of the evaluator's confirmation description trials as the next action. Skill Generation evaluation remains scheduled for Phase 7.
- The target-specific campaign artifacts required by `references/campaign-record.md` do not yet exist for Skill Generation: there is no `targets/skill-generation/contract.md`, `status.md`, four-trial description record, behavior record, or retained-evidence area.

## Target Contract

- **Goal:** execute a completed Skill Intake A-to-B plan and handoff into an idiomatic agent skill under the recorded working mode.
- **Intended use:** a completed intake contract is ready to be implemented, with destination, outcome, completion criteria, project context, supported harnesses, unresolved implementation matters, working mode, and user-owned stops available.
- **Expected behavior:** accept and verify the intake contract; map semantic ownership; apply the relevant authoring, document, code, system, and prompt disciplines; create only required skill resources and harness metadata; validate the artifact; and hand the result to Skill Evaluation and Refinement.
- **Expected result:** a structurally valid, resource-restrained skill that satisfies the intake criteria, accurately records runtime-proof requirements and limitations, and includes the required evaluation handoff.
- **Boundary:** Generation executes the existing intake plan rather than rediscovering intent, creating a shadow plan, performing behavior evaluation, inventing unsupported harness capability, or stopping for ordinary implementation choices inside its authority.

## Separate Evaluation Design

### Automatic Selection

Run the evaluator's complete four-trial description protocol before any body case:

1. Prepare two minimal, complete intake contracts whose natural next action is skill implementation. Prompts must request execution without naming `skill-generation`, requesting skill invocation, quoting its description, or revealing the expected selection.
2. Give each prompt to a fresh independent agent and retain the native `codex.skill.injected` event or exact pre-output read of `skills/skill-generation/SKILL.md`. Output similarity, final prose, and `Skills Used` claims are insufficient.
3. If both initial trials pass, run two different confirmation prompts through two more fresh agents and retain the same native evidence.
4. Audit selection only: whether the target loaded for representative completed-intake work. Do not use generated-skill quality to rescue or invalidate the selection result.

Representative boundaries should include a straightforward single-surface skill and a project-informed or cross-harness skill, while keeping each intake contract complete enough that Intake is no longer the correct owner.

### Post-Load Task Behavior

Start only after `status.md` records description state `passed` with four retained trial paths. Explicitly load the target for each isolated behavior case so the case measures its instructions rather than routing reliability. Use complete intake contracts and audit observable outputs against:

- contract acceptance and correct handling of intent-changing gaps versus ordinary implementation choices;
- semantic ownership and use of the governing disciplines;
- idiomatic `SKILL.md`, necessary resources only, and accurate harness metadata;
- structural validation and criterion-by-criterion artifact checks;
- accurate separation of structural evidence from deferred runtime proof;
- a complete evaluation handoff with no shadow generation plan;
- recorded working-mode and user-stop behavior.

Retain each unmodified prompt, fresh-agent identity, fixture specification, explicit load trace, generated output or transcript, ground-truth comparison, criterion result, and cleanup ownership. A behavior failure must remain a behavior result even when automatic selection passed.

## Blocking Gate

The valid-conclusion gate is campaign state and dependency order:

1. The evaluator has not yet met its recorded trust threshold for evaluating Intake and Generation; Phase 4 meta-evaluation is still active.
2. Skill Generation has no target campaign `status.md` recording a four-trial description pass.
3. The evaluator body contract prohibits creating or executing body cases until that durable description state exists.
4. This task permits only this response file and forbids changes to shared campaign records, so the missing campaign state cannot be established within this trial.

Running ad hoc Generation cases now would collapse campaign ordering and could produce useful observations, but it could not produce the evaluator-defined shipping conclusion.

## Shipping Decision

Keep `skills/skill-generation/SKILL.md` unchanged and classify it as **structurally inspectable but not yet behavior-validated for shipment**. Resume only after the evaluator completes its current trust work, then establish the Skill Generation target record, pass four direct-evidence automatic-selection trials, and run isolated post-load behavior cases. Ship as evaluated only if every required behavior criterion passes without a retained material failure.

## Skills Used

- `skill-issue:skill-evaluation-and-refinement`
- `skill-issue:document-update-discipline`
- `eric-wimp-toolkit:skill-authoring-discipline`

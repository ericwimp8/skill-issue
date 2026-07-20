# CI Failure Triage Intake Handoff

- **Source plan:** `showcase-skills/ci-failure-triage/plans/ci-failure-triage/ci-failure-triage-a-to-b-plan.md`
- **Destination:** `showcase-skills/ci-failure-triage/skill/ci-failure-triage/`
- **Outcome:** an evidence-first CI failure investigation skill that separates primary cause from cascade noise and delivers a bounded remediation direction plus exact verification.
- **Invocation boundary:** failed or suspicious continuous-integration jobs where logs, workflow configuration, repository source, or local tools must be correlated for diagnosis or handoff.
- **Project context:** use repository source as behavioral authority, use tests only after production behavior is established, retain all workflow artifacts in the assigned showcase workspace, and use no authenticated external runner requiring approval.
- **Completion contract:** all criteria in the source plan, including uncertainty preservation and the explicit remote-state authorization boundary.
- **Viability:** autonomous.
- **Execution preference:** autonomous continuation.
- **Authority:** create and evidence-supportedly refine assigned showcase artifacts only.
- **User-owned stops:** changed intent, unavailable private evidence, or required external authorization.
- **Evaluation prerequisites:** satisfied by `evaluations/skill-system-production-refinement/environment-qualification.md` and fresh isolated GPT-5.6 Sol agents.
- **Implementation decisions delegated to generation:** exact concise body structure and whether resources are necessary; no reference, script, or asset should be added without distinct reusable meaning.

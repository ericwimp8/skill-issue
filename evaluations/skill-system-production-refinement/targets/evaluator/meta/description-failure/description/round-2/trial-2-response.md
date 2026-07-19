# Release Communication Record Plan

## Target Outcome

For every release, support can open one authoritative record and determine which product changes were communicated, which were intentionally not communicated, and whether any decision is still unresolved. The record is reviewed before publication, retains approval and publication evidence, and remains the maintained source after release.

## Dependency-Ordered Work

1. **Define authority and accountability**
   - Designate one release communication record as the source of truth for each release.
   - Assign an accountable record owner, change contributors, required reviewers, publisher, and support audience.
   - Define which releases and product changes must appear, when a record must be opened, and when it may be closed.

2. **Specify the record and its lifecycle**
   - Give each record a stable release identifier, release date, product scope, owner, and current lifecycle status.
   - Give each change a stable identifier, customer-impact summary, inclusion status, communication decision, intended audience, communication channel, and supporting link.
   - Use explicit lifecycle states such as `Draft`, `In review`, `Approved`, `Published`, and `Amended`, with an owner and timestamp for every transition.
   - Represent `Communicated`, `Not communicated by decision`, and `Unresolved` separately so missing evidence cannot be mistaken for a deliberate decision.

3. **Establish complete change intake**
   - Identify the authoritative product-change sources that feed the record and map every in-scope change to the applicable release.
   - Require contributors to supply customer-facing impact and a proposed communication disposition by a defined cutoff.
   - Add a reconciliation check that exposes changes present in release scope but absent from the communication record.

4. **Create the pre-publication review gate**
   - Define reviewers for product accuracy, customer-facing wording, support readiness, and any conditional compliance review.
   - Require every change to have a resolved disposition and every communication claim to have its final content or destination attached.
   - Record reviewer, decision, timestamp, and requested corrections; block publication while required reviews or change dispositions remain unresolved.

5. **Publish one observable support view**
   - Publish the approved record at a stable, searchable location accessible to support.
   - Show the release-level publication status and, for each change, the final communication disposition, channel, audience, publication timestamp, and evidence link.
   - Mark a change `Communicated` only after publication evidence is recorded.
   - Preserve approved history when corrections are made, including what changed, why, by whom, and when.

6. **Pilot and reconcile the workflow**
   - Run the workflow on representative releases, including one with multiple product areas and one with a late change.
   - Compare source changes, review decisions, published communications, and the support view to find omissions or ambiguous states.
   - Resolve ownership, field, timing, and access gaps before adopting the workflow for all releases.

7. **Operationalize every release**
   - Add record creation, review completion, publication, and support verification to the release schedule.
   - Define escalation for overdue contributions, unresolved decisions, failed publication, and post-publication corrections.
   - Monitor record completeness and support access, and periodically review whether the record still answers support's release questions without relying on side channels.

## Completion Evidence

- Every release has exactly one authoritative record with a stable identifier, accountable owner, lifecycle status, and support-accessible location.
- Every in-scope product change is present and has a final, unambiguous communication disposition.
- Required approvals and publication events are visible with identities and timestamps.
- Each `Communicated` change links to verifiable published content; each intentionally uncommunicated change retains its approved rationale.
- Support can select a release and answer what was communicated, where, to whom, and when without consulting squad-specific notes.
- Missing changes, unresolved decisions, overdue reviews, and post-publication amendments remain visibly distinguishable and auditable.

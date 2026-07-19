# Maintained Release Record Plan

## Target Outcome

Establish one authoritative release record that replaces squad-owned release notes, passes a defined review before publication, and preserves an observable account of exactly what customers were told.

## Dependency-Ordered Work

1. **Map the current release-note flow**
   - Identify every squad note source, its owner, its audience, when it is updated, and where its content is reused.
   - Record duplicated, conflicting, missing, or unverifiable information and identify any downstream consumers that must move to the maintained record.
   - **Exit condition:** all current inputs and consumers are known, so consolidation will not silently remove required information.

2. **Define the authoritative record and ownership**
   - Choose the maintained location and name the role accountable for record completeness throughout a release.
   - Define who may contribute, edit, review, approve, publish, and correct information.
   - Make the maintained record the source used by product, support, and customer-communication workflows.
   - **Depends on:** step 1.
   - **Exit condition:** one source and one accountable owner are agreed.

3. **Define the record structure and lifecycle states**
   - Include the release identifier and date, customer-impacting changes, affected products or audiences, contributor evidence, customer-facing wording, reviewer decisions, publication details, and communication evidence.
   - Use explicit states such as `Draft`, `In Review`, `Approved`, `Published`, and `Customer Communicated`, with correction or withdrawal represented as traceable follow-up states.
   - Define the required evidence for each transition so an approved record cannot be mistaken for a published or communicated record.
   - **Depends on:** step 2.
   - **Exit condition:** the record schema and state-transition rules are documented and testable.

4. **Define squad contribution and consolidation rules**
   - Give squads a consistent submission point, deadline, required fields, and evidence expectations.
   - Define how the record owner resolves duplicate entries, conflicting wording, incomplete submissions, and changes arriving after review starts.
   - Stop creating parallel squad release notes once the maintained record accepts contributions.
   - **Depends on:** step 3.
   - **Exit condition:** every squad change follows one contribution path into the authoritative record.

5. **Establish the pre-publication review gate**
   - Specify required reviewers for product accuracy, customer wording, support readiness, and any release-specific compliance needs.
   - Require reviewers to approve the exact version intended for publication; requested changes return the record to `Draft` or `In Review` and invalidate stale approvals.
   - Prevent publication until all required decisions and unresolved issues are recorded.
   - **Depends on:** steps 3 and 4.
   - **Exit condition:** publication is gated by visible approval of the final content.

6. **Connect publication to customer communication evidence**
   - On publication, capture the approved content version, publication timestamp, destination, and publisher.
   - For each customer communication, retain the exact wording or immutable reference, audience, channel, sender, and sent timestamp.
   - Advance to `Customer Communicated` only when the required communication evidence exists; retain partial or failed delivery states where relevant.
   - **Depends on:** step 5.
   - **Exit condition:** the record can answer what was published, what customers were told, where, when, and to whom.

7. **Provide an observable communication view**
   - Present the current release state, outstanding owner or reviewer, approved version, publication status, and customer-communication status from the authoritative record.
   - Make the exact communicated message and its channel, audience, and timestamp accessible to support and other agreed consumers.
   - Clearly expose discrepancies such as published-but-unsent, partially communicated, corrected, or withdrawn content.
   - **Depends on:** step 6.
   - **Exit condition:** support can verify customer communication without reconstructing it from squad notes or chat history.

8. **Migrate, pilot, and retire the old process**
   - Populate the new record for a representative release and run it through contribution, review, publication, and communication-state updates.
   - Reconcile the pilot against source inputs and customer messages, then correct workflow or field gaps.
   - Redirect downstream consumers, archive prior squad notes with clear historical status, and publish the effective date for the new process.
   - **Depends on:** steps 4 through 7.
   - **Exit condition:** the pilot proves the complete workflow and active releases use only the maintained record.

9. **Measure and maintain the workflow**
   - Track missing submissions, review rework, time awaiting approval, publication exceptions, and releases lacking complete communication evidence.
   - Assign a regular owner review for schema, permissions, lifecycle rules, and consumer access.
   - Route corrections through traceable record revisions so the visible history continues to show both the original customer message and any later correction.
   - **Depends on:** step 8.
   - **Exit condition:** ownership and monitoring keep the record reliable after rollout.

## Completion Criteria

- Every active release has one authoritative maintained record.
- Squad contributions enter that record through a single defined path.
- Publication requires recorded approval of the exact final version.
- Publication and customer communication remain separately visible states.
- Support can retrieve the exact customer-facing message, audience, channel, and timestamp.
- Corrections preserve an auditable history of what customers were originally told and what changed later.

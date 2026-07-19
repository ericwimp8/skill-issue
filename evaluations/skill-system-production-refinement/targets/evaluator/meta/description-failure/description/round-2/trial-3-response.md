# Unified Release Record Plan

## Target State

Each release has one maintained record that remains authoritative from initial drafting through publication and communication. Product squads contribute release information through the same lifecycle, reviewers approve it before publication, and support can see whether each intended audience received the release communication and find the supporting evidence.

## Ordered Plan

1. **Establish ownership and scope**
   - Name the function accountable for the release-record process and the owner responsible for each individual release record.
   - Define which products, release types, squads, audiences, and communication channels are included.
   - Clarify squad responsibilities for supplying accurate change details, risks, rollout timing, and audience impact.

2. **Inventory current release practices**
   - Catalogue each squad's release-note sources, publication routes, review steps, communication channels, and support handoffs.
   - Identify duplicated records, missing information, conflicting terminology, approval gaps, and channels without delivery evidence.
   - Preserve any product-specific information that must remain available in the unified process.

3. **Define the maintained release record**
   - Agree on the minimum record fields: release identifier, product, owner, dates, scope, user impact, operational impact, risks, rollout or rollback details, linked work, and audience-specific messaging.
   - Add lifecycle fields for draft status, review status, publication status, and communication status.
   - Define one stable location and identifier for the record so publications and communications can link back to it rather than become competing sources.

4. **Set common states and decision rules**
   - Define the allowed progression, such as drafting, ready for review, changes requested, approved, scheduled, published, and closed.
   - Specify the entry and exit criteria for each state, including who may advance or return a record.
   - Define exception handling for urgent releases, delayed rollouts, partial releases, corrections, and cancellations.

5. **Design pre-publication review**
   - Establish required reviewers based on release impact, such as product, engineering, operations, support, security, legal, or communications.
   - Use a shared review checklist covering accuracy, completeness, audience clarity, timing, risk, sensitive information, and consistency with the actual release scope.
   - Record reviewer identity, decision, timestamp, and requested changes in the maintained record; publication begins only after required approval.

6. **Define publication from the record**
   - Map the approved record fields to each publication surface while preserving the maintained record as the authoritative source.
   - Standardize how product-specific details are transformed into customer, partner, internal, and support-facing messages.
   - Record publication destination, timestamp, published version, and link or artifact evidence for every output.

7. **Make communication status traceable**
   - Create a communication entry for each intended audience and channel, with an owner, planned time, status, actual send time, and evidence link.
   - Use clear statuses such as planned, ready, sent, confirmed, failed, superseded, and not required, with reasons for exceptions.
   - Give support a release-level view showing what changed, what was published, which audiences were contacted, outstanding actions, and the latest correction.

8. **Migrate active release work**
   - Select a cutover date and move all active or upcoming releases into the maintained record.
   - Link or archive legacy notes according to an agreed retention policy without creating parallel editable sources.
   - Backfill recent releases only to the depth needed for support continuity, auditability, and unresolved communication follow-up.

9. **Pilot and refine the workflow**
   - Run the process with a representative set of squads and release types, including one higher-risk release and one urgent path.
   - Verify that reviewers can block publication, publishers can generate channel outputs, and support can independently trace communication status.
   - Resolve ownership, field, state, and handoff gaps before wider rollout.

10. **Roll out and govern the standard**
    - Sequence adoption by squad, retire superseded release-note practices at each cutover, and provide concise role-based guidance.
    - Monitor record completeness, review turnaround, publication corrections, communication completion, and unresolved failures.
    - Hold periodic governance reviews to update the shared standard, audit exceptions, and maintain clear ownership as products and channels change.

## Completion Criteria

- All in-scope releases begin and end in the maintained release record.
- Required approvals are recorded before publication, including any approved exception path.
- Every publication and audience communication has an owner, status, timestamp, and evidence or documented reason.
- Support can trace a release from approved content through published outputs and communication completion without reconciling squad-specific note systems.
- Legacy practices are retired or retained as read-only history under an explicit retention policy.

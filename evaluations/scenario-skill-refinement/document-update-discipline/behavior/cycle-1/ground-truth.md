# Body Cycle 1 Ground Truth

## Case 1 — Backup Deletion

- Required meaning: all customer backups, including incident snapshots, are deleted exactly 45 days after creation.
- Document purpose: define backup policy and operational handling.
- Observation point: `Restore Checklist` step 3.
- Semantic home: `Deletion Policy`.
- Related manifestation: the incident-snapshot exception.
- Operation: replace the owning policy, make operational text defer or state only its local consequence, and remove the conflicting exception.
- Preserved meaning: encryption, restore authorization, incident evidence capture, and access restrictions.
- Pass characteristics: the complete deletion rule has one clear normative owner; checklist and incident text agree without fully restating the rule as competing owners.

## Case 2 — Severity Escalation

- Required meaning: every severity-one incident, including after-hours incidents, must be escalated to the duty manager within 10 minutes of declaration.
- Document purpose: define incident-response policy and procedures.
- Observation point: `Response Steps` step 2.
- Semantic home: `Escalation Policy`.
- Related manifestation: the after-hours exception.
- Operation: replace the owning policy, align the response step and after-hours section, and remove the conflict.
- Preserved meaning: evidence capture, customer communication, paging mechanics, and incident-command assignment.
- Pass characteristics: policy owns the complete obligation; procedure and after-hours text express local action or defer without repeating the whole rule.

## Case 3 — Planning Only

- Required meaning: every privileged device certificate, including break-glass devices, expires after 30 days.
- Document purpose: define privileged-device access and lifecycle controls.
- Observation point: `Enrollment Checklist` step 3.
- Semantic home: `Certificate Lifetime`.
- Related manifestation: the break-glass exception.
- Operation: plan the policy replacement, related reconciliation, and verification without editing the fixture.
- Preserved meaning: manager approval, hardware-backed keys, audit review, and emergency-use logging.
- Pass characteristics: the plan states the meaning, owner, observation point, manifestations, operations, and preserved meaning and keeps the complete normative rule at one owner.

# Update Plan

## Intended Meaning

Every privileged device certificate, including certificates for break-glass devices, expires after 30 days.

## Semantic Mapping

- **Document purpose:** Define privileged-device certificate policy, enrollment requirements, and break-glass controls.
- **Observation point:** Enrollment Checklist step 3 states a 90-day expiry.
- **Semantic home:** Certificate Lifetime owns the general certificate-expiry rule.
- **Related manifestations:** Enrollment Checklist step 3 repeats the lifetime, while Break-Glass Devices contradicts it by making those certificates non-expiring.
- **Preserve:** Manager approval, hardware-backed key storage, access-audit review, and emergency-use logging.

## Planned Changes

1. Replace the 90-day lifetime in Certificate Lifetime with a 30-day rule applying to every privileged device certificate.
2. Update Enrollment Checklist step 3 from 90 days to 30 days so enrollment implements the policy.
3. Replace the non-expiring break-glass exception with the same 30-day expiry while retaining emergency-use logging and review during the next access audit.
4. Verify the guide contains no conflicting certificate lifetimes and that all preserved controls remain intact.

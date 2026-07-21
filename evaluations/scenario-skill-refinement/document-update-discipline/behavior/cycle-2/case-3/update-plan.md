# Update Plan

1. Replace the `Certificate Lifetime` rule with one authoritative requirement: every privileged device certificate expires after 30 days, including certificates for break-glass devices.
2. Change `Enrollment Checklist` step 3 from 90 days to 30 days so enrollment applies the owning lifetime rule without creating a separate exception or policy owner.
3. Remove the non-expiring certificate exception from `Break-Glass Devices`; retain emergency-use logging and next-access-audit review as the section's local operational requirements.
4. Verify that the guide contains no remaining 90-day or non-expiring privileged certificate language and that manager approval, hardware-backed key storage, emergency-use logging, and audit review remain unchanged.

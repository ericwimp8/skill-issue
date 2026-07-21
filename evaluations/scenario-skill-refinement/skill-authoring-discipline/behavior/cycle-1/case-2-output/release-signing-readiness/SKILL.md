---
name: release-signing-readiness
description: Mobile release signing readiness review guidance. Use when reviewing whether an Android or Apple app is correctly identified and signed for release.
---

# Release Signing Readiness

## Review Workflow

1. Identify the target platform before loading platform policy. If the release covers both platforms, review each separately.
2. Read only the matching reference:
   - Android: `references/android.md`
   - Apple: `references/apple.md`
3. Trace the repository's release entrypoint and configuration to establish its declared release identity.
4. Separate signing material expected on the local machine from credentials supplied or managed by CI. Treat references to secret stores as evidence of credential management, not evidence of credential contents.
5. Compare the declared identity and signing path against the matching platform policy.
6. Report verified facts, missing evidence, and unresolved mismatches. Never infer or invent absent identifiers, credential values, or signing settings.

## Reference Documents

Use only the reference matching the release platform under review.

- `references/android.md`: Android release signing policy. Use for Android application IDs, upload keys, signing configurations, and keystores.
- `references/apple.md`: Apple release signing policy. Use for bundle identifiers, teams, certificates, and provisioning profiles.

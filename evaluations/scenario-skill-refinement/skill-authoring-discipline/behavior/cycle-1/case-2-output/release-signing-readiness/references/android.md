# Android Release Policy

- Compare the release application ID with the identity registered for the upload key.
- Verify the release task uses the intended signing configuration rather than a debug fallback.
- Never print keystore passwords or private-key material.

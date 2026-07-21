# Signing Release Skill Requirements

Create a Codex skill for reviewing mobile release signing readiness.

Shared guidance must require the agent to identify the target platform, verify the repository's declared release identity, distinguish local signing material from CI-managed credentials, and report missing evidence without inventing values.

For Android releases, preserve these policy details:

- compare the release application ID with the identity registered for the upload key;
- verify the release task uses the intended signing configuration rather than a debug fallback;
- never print keystore passwords or private-key material.

For Apple releases, preserve these policy details:

- compare the bundle identifier, team, certificate type, and provisioning profile as one signing identity;
- distinguish App Store distribution from development or ad hoc signing;
- never print exported certificate passwords or private-key material.

Agents handling one platform should not load the other platform's policy. Include minimal OpenAI display metadata. No default prompt is requested.

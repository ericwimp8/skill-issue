# Workflow Observations

## Native Codex Runner Approval Boundary

The repository's development CLI custom-evaluation route passed input preparation and reached the native Codex launch, but the sandbox prevented the Codex state database from opening for write. The required escalation was denied because authenticated external-harness transmission of the local skill, fixture, and transcript needed explicit approval beyond this agent's sandbox authority. No native CLI activation result is claimed.

The current evaluation workflow would be easier to execute reliably if its environment preflight explicitly checked state-store writability and external authenticated-run approval before trial inputs are prepared. The source should also state whether a fresh-agent candidate-selection protocol with direct pre-output reads is an accepted fallback evidence class when the qualified native signal route is unavailable. This campaign retains both the stopped CLI boundary and the exact fallback method rather than treating them as equivalent.

## Post-Evaluation Formatting Boundary

The repository-required `npm run format:check` passes. A broader optional Prettier sweep over the entire showcase workspace reports style changes for some generated fixtures, retained trial outputs, the generated report asset, and OpenAI metadata. Applying those changes after evaluation would invalidate the content hashes recorded by every trial. The workflow does not currently say whether formatting normalization must occur before target hashing and independent evaluation. Future campaigns should run any intended broad formatter before generation hashes are frozen and trials begin.

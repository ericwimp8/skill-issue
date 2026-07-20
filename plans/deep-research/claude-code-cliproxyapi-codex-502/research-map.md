# Research Map: Claude Code Through CLIProxyAPI 502

## Goal

Identify the best-supported fixes and mitigations for a Claude Code evaluation that succeeds for several proxied requests, then fails on a long streamed request with an SSE-reported upstream OpenAI 502 `server_error`, while an identical rerun succeeds.

## Framing

The final synthesis must present the best-supported answer or direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers. It must distinguish durable fixes from transient retry advice and unsafe or unsupported workarounds.

## Scope And Budget

- Source scope: internet only
- Active researcher concurrency: 5
- Total researcher budget: 5
- Discovery waves: targeted pathway mapping followed directly by five evidence-class deep dives because the supplied budget equals the number of separable high-priority source areas

## Research Domains And Assignments

1. **Exact proxy implementation and releases**
   - Assignment: `assignments/01-cliproxyapi-claude-codex-compatibility.md`
   - Targets: router-for-me/CLIProxyAPI source, releases, changelog, issues, discussions, Claude Code and Codex OAuth compatibility
   - Expected evidence: version-specific fixes, known regressions, configuration requirements, maintainer guidance
2. **Protocol translation and SSE failure propagation**
   - Assignment: `assignments/02-anthropic-responses-stream-translation.md`
   - Targets: Anthropic-compatible `/v1/messages?beta=true`, OpenAI Responses translation, SSE error mapping, stream lifecycle
   - Expected evidence: concrete translation behavior and where an upstream 502 becomes a Claude-facing stream error
3. **OpenAI upstream 502 and retry guidance**
   - Assignment: `assignments/03-openai-responses-502-retries.md`
   - Targets: official OpenAI API error, streaming, timeout, retry, status, and reliability guidance
   - Expected evidence: durable client/proxy handling versus transient retry recommendations
4. **Claude Code client behavior and version reports**
   - Assignment: `assignments/04-claude-code-stream-retry-behavior.md`
   - Targets: Anthropic Claude Code documentation, changelogs, issues, retry/timeout/stream handling, compatibility with custom Anthropic endpoints
   - Expected evidence: client limitations, version-specific bugs, exit behavior, supported configuration
5. **Operational mitigations and safety boundaries**
   - Assignment: `assignments/05-operational-mitigations-and-safety.md`
   - Targets: proxy deployment guidance, long-running HTTP/SSE handling, observability, retry placement, timeout layers, authentication safety
   - Expected evidence: practical mitigations ranked by durability and explicit rejection of credential copying, destructive resets, permanent user configuration mutation, or unsupported authentication extraction

## Fan-Out Decisions

- Deep-dive all five assignments; the incident spans five separable evidence classes and no budget remains for additional branches.
- Classify newly discovered candidates within each assignment as selected evidence, skim-only context, or rejected/lower-fit interpretation.
- Any further source branch that cannot fit the budget must be recorded as unsupported or blocked rather than delegated.

## Observed Incident To Preserve

- CLIProxyAPI 7.2.91
- Claude Code 2.1.205
- Model `gpt-5.6-sol`
- Four proxied model requests completed
- The fifth streamed request waited about 6m39s, then returned an SSE error corresponding to upstream HTTP 502 and OpenAI `server_error`
- Claude exited 1 with empty stderr
- An identical one-turn rerun through the same launcher completed in about 2m48s

## Aggregation

- Assignment folder: `plans/deep-research/claude-code-cliproxyapi-codex-502/assignments/`
- Final target: `plans/deep-research/claude-code-cliproxyapi-codex-502/claude-code-cliproxyapi-codex-502-deep-research.md`
- Requested shape: best-supported answer or direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers

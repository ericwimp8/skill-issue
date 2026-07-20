# CLIProxyAPI, Claude Code, and Codex 502 Compatibility

## Assignment

**Goal:** Determine whether CLIProxyAPI 7.2.91, Claude Code 2.1.205, operator-owned Codex OAuth, and `gpt-5.6-sol` are compatible; trace the reported fifth-request streamed HTTP 502 failure; and identify durable fixes, transient operational advice, version-specific defects, required configuration, and unsafe workarounds.

**Scope:** Internet-only inspection of primary CLIProxyAPI source, releases, commits, issues, and documentation; official Claude Code documentation and release history; and official OpenAI Responses streaming and service-status sources. The incident facts are four successful requests, then a fifth request that streamed for about 6 minutes 39 seconds before an SSE error reporting an upstream OpenAI `server_error` as HTTP 502, followed by an identical rerun that succeeded after about 2 minutes 48 seconds.

**Exclusions:** Local repository source, private credentials, destructive probes, claims about the unavailable raw request or terminal SSE event, and unrelated Anthropic-OAuth or third-party relay implementations.

## Sources

- CLIProxyAPI v7.2.91 release and commit `fde40c5a0a2f8f6808bcde498bc6079f32c355ef`: [release](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.91), [commit](https://github.com/router-for-me/CLIProxyAPI/commit/fde40c5a0a2f8f6808bcde498bc6079f32c355ef).
- CLIProxyAPI v7.2.91 Codex executor, including `codexTerminalFailureStatus`, `codexTerminalFailure`, `ExecuteStream`, `/responses` dispatch, and Codex OAuth headers: [codex_executor.go](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/runtime/executor/codex_executor.go).
- CLIProxyAPI v7.2.91 Claude-compatible `/v1/messages` handler, including streaming and `forwardClaudeStream`: [code_handlers.go](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/api/handlers/claude/code_handlers.go).
- CLIProxyAPI v7.2.91 Claude-to-Codex request and Codex-to-Claude response translators: [request translator](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/translator/codex/claude/codex_claude_request.go), [response translator](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/translator/codex/claude/codex_claude_response.go).
- CLIProxyAPI v7.2.91 stream bootstrap, forwarding, credential failover, retry, and cooldown implementation: [handlers.go](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/api/handlers/handlers.go), [stream_forwarder.go](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/api/handlers/stream_forwarder.go), [conductor.go](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/cliproxy/auth/conductor.go).
- CLIProxyAPI v7.2.91 configuration and model registry: [config.example.yaml](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/config.example.yaml), [Codex model registry](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/registry/models/codex_client_models.json).
- CLIProxyAPI v7.2.80 terminal-event fix, commit `09da52ad509e2c18e7b9540db3b98c2214c280aa`, and the motivating incomplete-stream report: [release](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.80), [commit](https://github.com/router-for-me/CLIProxyAPI/commit/09da52ad509e2c18e7b9540db3b98c2214c280aa), [issue #3055](https://github.com/router-for-me/CLIProxyAPI/issues/3055).
- CLIProxyAPI v7.2.19 configurable transient-error cooldown, commit `d33ac5e1e9dad15e8de5be011a2e9eda921d538f`: [release](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.19), [commit](https://github.com/router-for-me/CLIProxyAPI/commit/d33ac5e1e9dad15e8de5be011a2e9eda921d538f).
- CLIProxyAPI `gpt-5.6-sol` support and adjacent releases: [v7.2.55](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.55), [v7.2.59](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.59), [v7.2.70](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.70), [v7.2.92](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.92).
- CLIProxyAPI reports involving transient Codex failures: [issue #4290](https://github.com/router-for-me/CLIProxyAPI/issues/4290), [issue #636](https://github.com/router-for-me/CLIProxyAPI/issues/636), [issue #2189](https://github.com/router-for-me/CLIProxyAPI/issues/2189), [issue #2401](https://github.com/router-for-me/CLIProxyAPI/issues/2401).
- CLIProxyAPI official documentation: [Codex OAuth](https://help.router-for.me/configuration/provider/codex), [Claude Code client](https://help.router-for.me/agent-client/claude-code), [configuration options](https://help.router-for.me/configuration/options), [product overview](https://help.router-for.me/introduction/what-is-cliproxyapi).
- OpenAI official Responses streaming `response.failed` event: [API reference](https://platform.openai.com/docs/api-reference/responses-streaming/error).
- OpenAI official status incident, "Codex 5.6-sol Experiencing Increased Server-Overload Errors," resolved July 17, 2026: [incident](https://status.openai.com/incidents/01KXRHE25717D2WQ1WFMT2B7WZ), [status history](https://status.openai.com/history).
- Anthropic official Claude Code documentation: [environment variables](https://code.claude.com/docs/en/env-vars), [LLM gateway](https://code.claude.com/docs/en/llm-gateway), [errors](https://code.claude.com/docs/en/errors), [changelog](https://github.com/anthropics/claude-code/blob/main/CHANGELOG.md).
- Anthropic Claude Code package versions: [2.1.205](https://www.npmjs.com/package/@anthropic-ai/claude-code/v/2.1.205), [2.1.215](https://www.npmjs.com/package/@anthropic-ai/claude-code/v/2.1.215).

## Findings

### Finding 1 — The stated versions and endpoint path are fundamentally compatible

CLIProxyAPI 7.2.91 implements a concrete Claude Messages to Codex Responses bridge. Its Claude handler accepts `/v1/messages`; its request translator maps Claude system, text, image, tool-use, tool-result, and thinking fields into a Codex Responses request; its Codex executor sends that request to the operator-authenticated Codex `/responses` endpoint; and its response translator maps Codex text, reasoning, tool, completion, incomplete, and error events back into Claude streaming events. The four successful requests independently demonstrate that this configured path functioned in the incident environment.

**Evidence:** The v7.2.91 Claude handler, request translator, Codex executor, and response translator implement the full path in production source. The v7.2.91 Codex registry includes `gpt-5.6-sol`, marks it supported in the API, and lists its reasoning levels. CLIProxyAPI added GPT-5.6 model support in v7.2.55 and restored/revised Sol support in v7.2.59. Claude Code 2.1.205 is newer than the 2.1.129 minimum for optional gateway model discovery documented by Anthropic. Manual model environment variables avoid depending on discovery.

**Implication:** The isolated fifth-request failure is not evidence of a general Claude Code 2.1.205-to-CLIProxyAPI 7.2.91 incompatibility. Investigation and remediation should center on the terminal upstream failure, retry boundary, and cooldown behavior rather than replacing the translation architecture.

### Finding 2 — CLIProxyAPI intentionally converts an upstream Codex terminal `server_error` into the observed Claude SSE error

In v7.2.91, the Codex executor recognizes upstream SSE events named `error` and `response.failed`. It extracts the upstream error object, attempts to obtain a numeric status, and defaults unclassified terminal failures to HTTP 502. The Claude stream forwarder then emits the terminal failure as `event: error` with Claude-compatible JSON. When the upstream body contains `error.type: server_error`, the Claude error mapping preserves that type and message.

**Evidence:** `codexTerminalFailure` handles both terminal event names; `codexTerminalFailureStatus` defaults to 502 after specific status and error-code classifications; `ExecuteStream` terminates on either event; and `forwardClaudeStream` writes the error as a terminal SSE event after the stream has begun. OpenAI's official Responses streaming reference documents `response.failed` and shows a `server_error` example for a model generation failure.

**Implication:** The reported shape is the expected propagation of an upstream OpenAI/Codex terminal failure. Suppressing it, converting it to success, or treating it as a complete Claude message would conceal a failed generation and risk corrupting message or tool state.

### Finding 3 — Version 7.2.91 already contains the relevant terminal-event and long-reasoning-history fixes

CLIProxyAPI v7.2.80 corrected an older failure mode in which Codex streams that ended without `response.completed` could be reported as an incomplete 408 and could trigger credential rotation or cooldown. It added recognition of `error`, `response.failed`, and `response.incomplete`, including the default 502 classification. Version 7.2.91 later added sanitization for overlong encrypted reasoning IDs in Codex input processing, a potentially relevant protection for long Claude conversations. Version 7.2.92 contains performance-oriented payload-normalization work, with no release-note or source-level indication of a retry or SSE reliability fix for this failure.

**Evidence:** The v7.2.80 commit and issue #3055 show the terminal-event change and its motivating large-request failure. The v7.2.91 release contains only the encrypted-reasoning-ID sanitization feature. The v7.2.92 release lists Codex payload-normalization and helper optimizations rather than terminal-event, upstream-502, or replay changes.

**Implication:** Stay on at least v7.2.91; upgrading to v7.2.92 is reasonable for current fixes and performance but should not be represented as a demonstrated cure for this 502. Downgrading below v7.2.80 loses correct terminal-event handling, while downgrading below v7.2.91 also loses the reasoning-ID protection.

### Finding 4 — A transient upstream model failure is the leading diagnosis, with an important timestamp caveat

The error originated as an upstream Codex Responses terminal `server_error`, and an identical rerun succeeded. OpenAI also recorded increased server-overload errors specifically for Codex 5.6-sol on July 17, 2026, within the same version era. CLIProxyAPI issue #4290 independently reports `gpt-5.6-sol` returning 502 after roughly two minutes and then placing credentials into temporary `auth_unavailable` cooldown before later recovery.

**Evidence:** OpenAI's official API reference defines this terminal error class; its status incident names Codex 5.6-sol overload errors; issue #4290 records the same model, an upstream 502, subsequent cooldown unavailability, and eventual reuse; and the incident's successful identical rerun is consistent with a non-deterministic upstream failure.

**Implication:** Treat a single occurrence as transient upstream service failure unless request tracing reveals a repeatable payload trigger. The missing incident timestamp and raw terminal event prevent attribution to the July 17 outage or proof of the precise upstream event variant.

### Finding 5 — Automatic recovery is deliberately limited to the pre-payload bootstrap window

CLIProxyAPI buffers the beginning of a stream until it observes the first payload. Before any client-visible payload, eligible failures—including unknown status, 408, 429, and 5xx—can use another eligible credential or configured bootstrap retry. After a payload has been sent, the handler forwards a terminal error and does not replay the request. This boundary prevents duplicate text and, more importantly, duplicate tool effects.

**Evidence:** The v7.2.91 handler tracks `sentPayload`; its documented "Safe bootstrap recovery" path is gated by `!sentPayload`. The auth conductor obtains the initial payload before returning a wrapped stream and can try other credentials during bootstrap. The stream forwarder marks payload delivery and terminates on error. The general `request-retry` setting does not override the no-replay rule once client-visible streaming has begun.

**Implication:** `streaming.bootstrap-retries: 1` and multiple Codex OAuth credentials can improve recovery only when the upstream fails before the first payload. A 6-minute-39-second request may still be in bootstrap if it produced no translated content, but duration alone cannot establish that. Once partial output or tool activity reached Claude Code, manual/operator-aware rerun is the safe recovery path.

### Finding 6 — Cooldown tuning can prevent a transient 502 from becoming an avoidable local blackout

CLIProxyAPI treats 408, 500, 502, 503, and 504 as transient upstream errors and temporarily cools the selected credential. The default transient cooldown is 60 seconds when the configuration value is zero, while the example maximum retry interval is 30 seconds. With one credential, that relationship can prevent an immediate automatic same-credential retry and can produce temporary `auth_unavailable` responses like issue #4290.

**Evidence:** The v7.2.91 conductor classifies those statuses as transient, applies the transient cooldown, and performs cooldown-aware selection. The example configuration documents `transient-error-cooldown-seconds`, `max-retry-interval`, `request-retry`, credential limits, and the streaming bootstrap options. Version 7.2.19 introduced configurable transient-error cooldown. Issue #4290 records the practical 502-to-503 cooldown sequence on `gpt-5.6-sol`.

**Implication:** For a single-credential installation where false blackout time is costly, use a modest positive transient cooldown no longer than the intended retry wait—for example, 10–15 seconds with a 30-second maximum retry interval—and monitor behavior. Multiple independently operator-authorized Codex OAuth credentials provide stronger pre-stream failover. Disabling cooldown globally is an emergency or tightly controlled option because it can hot-loop against a persistent upstream failure.

### Finding 7 — Correct authentication and Claude Code configuration keep OAuth ownership at the proxy

The operator should complete Codex OAuth with CLIProxyAPI using `--codex-login`. Claude Code should connect to the proxy using `ANTHROPIC_BASE_URL` set to the server origin, without appending `/v1`, and `ANTHROPIC_AUTH_TOKEN` set to one of CLIProxyAPI's configured access keys. Claude Code 2.x model environment variables can map Opus, Sonnet, and Haiku selections to `gpt-5.6-sol` or the chosen proxy model.

**Evidence:** CLIProxyAPI's Codex provider documentation assigns OAuth login to the server. Its Claude Code guide shows the base URL and proxy token pattern and the Claude Code 2.x model variables. Anthropic's official environment-variable documentation defines `ANTHROPIC_BASE_URL` for gateways and `ANTHROPIC_AUTH_TOKEN` as the bearer token sent to that gateway.

**Implication:** Do not expose or copy the operator's Codex OAuth token or auth file into Claude Code. The client authenticates to CLIProxyAPI with a proxy access key; CLIProxyAPI owns and applies the upstream OAuth credential.

### Finding 8 — Client and proxy upgrades are sensible hardening, but no identified post-7.2.91 fix targets this server error

Claude Code 2.1.205 was published July 8, 2026; 2.1.215 was available by July 19. Intermediate releases include fixes for HTTP/2 GOAWAY crashes and stale keep-alive behavior after connection errors, including a Windows corporate-proxy case. Those changes can reduce client-side transport failures but do not describe a fix for a valid upstream Responses `response.failed` or `server_error` event.

**Evidence:** Anthropic's official changelog lists the transport fixes in 2.1.208 and 2.1.214. CLIProxyAPI v7.2.92 lists performance changes rather than a terminal-502 fix. The observed error carried an upstream OpenAI error classification rather than a local socket or timeout failure.

**Implication:** Upgrade Claude Code and CLIProxyAPI within operational policy for accumulated fixes, while preserving the diagnosis: these upgrades are general hardening, not evidence-backed cures for the recorded upstream model failure.

### Finding 9 — Several tempting workarounds are ineffective or unsafe

Stream keepalives only keep the downstream connection active; they cannot repair a model generation that ends with an upstream `server_error`. Claude Code's documented `API_TIMEOUT_MS` default is 600000 milliseconds, so the 6-minute-39-second failure occurred before the default 10-minute timeout and ended with an explicit server error. Automatic replay after partial output can duplicate tool effects. Disabling cooldown can hammer a failing upstream. Hiding terminal failures or reverting to older proxy versions discards correctness fixes.

**Evidence:** CLIProxyAPI's stream forwarder writes keepalive frames independently of upstream generation; its retry source forbids replay after payload delivery. Anthropic documents the 600000-millisecond default. The v7.2.80 and v7.2.91 commits show why older behavior is undesirable.

**Implication:** Reject keepalive and timeout increases as fixes for this explicit terminal error; reject midstream automatic replay, terminal-error suppression, downgrade, or blanket cooldown disablement. A manual rerun after a standalone failure is proportionate transient advice, particularly when the caller can confirm that repeating tool actions is safe.

## Notes

### Candidate classification

- **Deep-dive:** v7.2.91 Codex executor; Claude handler; both translation directions; handler, forwarder, and auth conductor; v7.2.80 terminal-event commit and issue #3055; v7.2.91 release commit; v7.2.19 cooldown commit; v7.2.91 configuration and model registry; issue #4290; OpenAI Responses streaming reference; OpenAI `gpt-5.6-sol` status incident; CLIProxyAPI Codex and Claude Code guides; Anthropic environment, gateway, error, changelog, and package-version sources.
- **Skim-only:** issue #636 for older large-prompt incomplete streams; issue #2189 for empty streams through third-party Claude relays; issue #2401 for a transient Windows/direct-SSE 502; Claude Code 2.1.208 and 2.1.214 transport fixes. These provide adjacent context but differ from the exact provider, version, client, or error signature.
- **Rejected/lower fit:** CLIProxyAPI issue #2599 about Anthropic upstream billing/system-prompt behavior; PR #2748 for custom Anthropic-compatible upstreams; unrelated third-party proxy/transformer repositories; and an unofficial `CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK` variable absent from Anthropic's official environment-variable documentation.

### Caveats and unsupported claims

- The incident timestamp was not supplied, so attribution to OpenAI's July 17, 2026 overload incident is unsupported.
- The raw upstream SSE terminal event and request identifier were unavailable, so whether the event was `error` or `response.failed` and whether any translated payload preceded it remain unknown.
- Exact CLIProxyAPI logs and configuration were unavailable, so the credential-selection, bootstrap-retry, and cooldown path taken in this request cannot be proven.
- The identical successful rerun strongly supports transience but does not rule out a payload-sensitive or non-deterministic model defect.
- No CLIProxyAPI functional fix specifically targeting this failure was identified after v7.2.91 through v7.2.92.
- No primary-source evidence identifies Claude Code 2.1.205 as the cause of an upstream OpenAI `server_error`.

### Useful operational evidence for recurrence

- Preserve CLIProxyAPI request logs, upstream request identifiers, the raw terminal SSE event, whether any client-visible payload or tool call preceded the error, the selected auth credential label, and cooldown transitions. These observations distinguish safe bootstrap retry from unsafe midstream replay without exposing OAuth secrets.

# Claude Code 2.1.205 Stream Retry and Gateway Behavior

## Assignment

**Goal:** Establish Claude Code 2.1.205 client behavior around custom Anthropic-compatible endpoints, long-lived SSE streams, retries, timeouts, terminal failure output, and nearby version-specific fixes, then distinguish supported mitigations from transient workarounds for the observed CLIProxyAPI 7.2.91 → OpenAI Codex OAuth / `gpt-5.6-sol` 502.

**Scope:** Internet-only research using Anthropic documentation, Claude Code release notes, the published 2.1.205 package metadata, and primary Claude Code issue reports. The incident facts preserved for interpretation are: four requests succeeded; the fifth waited 6m39s, received an SSE upstream 502 `server_error`, exited 1, and had empty captured stderr; an identical rerun succeeded.

**Exclusions:** No local configuration or credential inspection; no changes to client, proxy, or authentication state; no credential-copying guidance; no destructive reset guidance; no permanent user-configuration mutation; no claim about the failing hop without correlated proxy/upstream evidence.

## Sources

- Anthropic Claude Code [Error reference](https://code.claude.com/docs/en/errors), including automatic retries, timeout defaults, mid-stream server errors, and custom-gateway attribution.
- Anthropic Claude Code [Programmatic usage](https://code.claude.com/docs/en/headless), including `stream-json`, `system/api_retry`, final result messages, and version-specific output notes.
- Anthropic Claude Code [Other LLM gateways](https://code.claude.com/docs/en/llm-gateway), [Connect Claude Code to an LLM gateway](https://code.claude.com/docs/en/llm-gateway-connect), and [Gateway protocol reference](https://code.claude.com/docs/en/llm-gateway-protocol).
- Anthropic Claude Platform [API errors](https://platform.claude.com/docs/en/api/errors) and [Streaming messages](https://platform.claude.com/docs/en/build-with-claude/streaming).
- Official Claude Code releases: [v2.1.196](https://github.com/anthropics/claude-code/releases/tag/v2.1.196), [v2.1.198](https://github.com/anthropics/claude-code/releases/tag/v2.1.198), [v2.1.199](https://github.com/anthropics/claude-code/releases/tag/v2.1.199), [v2.1.205](https://github.com/anthropics/claude-code/releases/tag/v2.1.205), and [v2.1.208](https://github.com/anthropics/claude-code/releases/tag/v2.1.208).
- Primary historical issue report [#33949: SSE streaming hangs indefinitely](https://github.com/anthropics/claude-code/issues/33949), covering 2.1.50–2.1.74 and explicitly superseded for current interpretation by the v2.1.196 watchdog release note.
- npm registry metadata for [`@anthropic-ai/claude-code@2.1.205`](https://registry.npmjs.org/%40anthropic-ai%2Fclaude-code/2.1.205) and its platform-native package, used only to confirm the exact published version/package shape.

## Findings

### 1. Claude Code 2.1.205 already contained three relevant retry-era changes

Version 2.1.205 did not introduce a new streaming retry or timeout policy. Its release notes cover schema validation, background agents, worktrees, UI state, and related fixes, with no network/SSE retry change. The relevant behavior was introduced immediately before it: v2.1.196 enabled a streaming idle watchdog for all providers; v2.1.198 added backoff retry for transient mid-response connection drops; and v2.1.199 changed mid-stream overloaded/server-error handling and retry controls.

**Evidence:** v2.1.196 says the stream watchdog is on by default for all providers and “aborts and retries” when a response stream produces no events for five minutes ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.196)). v2.1.198 says transient mid-response errors such as `ECONNRESET` retry with backoff ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.198)). v2.1.199 says partial output is retained when an API stream emits an overloaded/server error, and documents expanded retry-watchdog behavior ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.199)). The v2.1.205 release contains no corresponding network change ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.205)).

**Implication:** Treat 2.1.205 as inheriting the v2.1.196–2.1.199 behavior. There is no release-note evidence for a 2.1.205-specific 502 regression or a 2.1.205-specific retry fix.

### 2. Default retries, request timeout, and stream-idle timeout are separate mechanisms

Claude Code retries transient server errors, overloads, request timeouts, temporary throttles, and dropped connections up to 10 times with exponential backoff. `API_TIMEOUT_MS` defaults to 600,000 ms per request. Separately, the 2.1.205 stream watchdog aborts and retries after five minutes with no SSE events. `CLAUDE_CODE_RETRY_WATCHDOG=1` is an unattended-run control: for 2.1.199+ it retries capacity errors indefinitely and raises other transient-error retries to 300, approximately three hours of backoff.

**Evidence:** The Claude Code error reference documents 10 exponential-backoff retries, `CLAUDE_CODE_MAX_RETRIES=10`, `API_TIMEOUT_MS=600000`, and the retry-watchdog semantics ([automatic retries](https://code.claude.com/docs/en/errors#automatic-retries)). The five-minute no-event watchdog is version-anchored in the v2.1.196 release ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.196)).

**Implication:** A 6m39s wall time is not the default 10-minute request deadline. It also does not prove the five-minute watchdog failed: the elapsed time can include SSE events, an internal retry, and backoff. Only retry events or client debug output can show the actual attempt sequence.

### 3. An SSE error after HTTP 200 is a first-class mid-stream failure

The Anthropic Messages API may return HTTP 200, begin an SSE response, and later emit `event: error`. Claude Code therefore has to distinguish a failure before visible assistant output from one after visible output. Before visible output, a retryable server error is retried. After visible output, 2.1.199+ preserves the partial response and appends an incomplete-response notice rather than reissuing the request, because retrying could repeat tool calls.

**Evidence:** Anthropic documents that SSE errors can arrive after HTTP 200 and shows the `event: error` envelope ([API errors](https://platform.claude.com/docs/en/api/errors), [streaming error events](https://platform.claude.com/docs/en/build-with-claude/streaming#error-events)). Claude Code’s error reference says mid-stream 5xx/overload errors after visible output are retained and not retried, while the same class before visible output is retried ([incomplete responses](https://code.claude.com/docs/en/errors#the-response-above-may-be-incomplete)).

**Implication:** The final 502 alone cannot show whether Claude Code retried. If no visible output existed, the displayed error normally follows exhausted retries or a non-retryable interpretation. If visible output existed, finalization without replay is deliberate duplicate-side-effect protection. A whole-process rerun is safe only when the caller has independently established idempotency.

### 4. `ANTHROPIC_BASE_URL` is supported for Anthropic-format gateways, with a strict stream contract

Claude Code supports a gateway selected by `ANTHROPIC_BASE_URL` when it exposes the Anthropic Messages endpoint. Inference posts to `/v1/messages?beta=true`; responses must be streamed rather than buffered; `anthropic-version` and `anthropic-beta` must be forwarded unchanged; evolving Anthropic headers and request fields must be treated as open lists; and upstream error bodies must remain unwrapped because Claude Code’s recovery logic matches their wording.

**Evidence:** The gateway protocol defines the Anthropic Messages format and endpoint, requires streaming, warns that buffered gateways stall Claude Code, and requires unchanged forwarding of version/beta headers ([protocol](https://code.claude.com/docs/en/llm-gateway-protocol)). It also states that wrapping upstream error bodies breaks automatic capability-rejection recovery ([automatic retry and error forwarding](https://code.claude.com/docs/en/llm-gateway-protocol#automatic-retry-and-error-forwarding)).

**Implication:** Durable gateway hardening means unbuffered SSE pass-through, correct named events, unchanged error envelopes, open-list header/body forwarding, and correlated logging at each hop. A proxy that emits keepalive `ping` events can prevent the no-event watchdog from firing; those pings do not solve an upstream operation that later returns 502.

### 5. Routing Claude Code to `gpt-5.6-sol` is outside Anthropic’s supported gateway boundary

Anthropic supports third-party gateway products that expose a supported Claude API format, but explicitly does not support using a gateway to route Claude Code to non-Claude models. CLIProxyAPI translating Claude Code’s Anthropic-format traffic into OpenAI Codex OAuth/model traffic is therefore an interoperability arrangement rather than a supported Claude Code deployment.

**Evidence:** Anthropic’s gateway overview says any supported API format can work, while also stating that Anthropic “doesn’t support routing Claude Code to non-Claude models through any gateway” ([Other LLM gateways](https://code.claude.com/docs/en/llm-gateway)). The protocol further notes that Claude Code adds fields, headers, and capabilities across releases, requiring gateways to evolve with it ([gateway protocol](https://code.claude.com/docs/en/llm-gateway-protocol)).

**Implication:** Four successes and a successful identical rerun demonstrate basic, intermittent interoperability, not full conformance across thinking, tools, beta headers, retries, error translation, or future releases. Anthropic cannot provide a durable support guarantee for the `gpt-5.6-sol` route; durable support requires a Claude-capable upstream or support from the translation proxy’s owner.

### 6. Exit 1 is terminal failure; empty stderr is not sufficient diagnostic evidence

Claude Code’s programmatic interface puts structured progress and failure information in its selected stdout format. With `stream-json`, retry attempts appear as `system/api_retry` records containing attempt, maximum retries, delay, HTTP status, and an error category such as `server_error`; the stream normally ends with a `result` record. The public CLI documentation does not promise that every terminal API failure is duplicated to stderr. Consequently, a process exit status of 1 with empty captured stderr can coexist with useful failure information having been emitted to stdout or swallowed/truncated by a wrapper.

**Evidence:** Programmatic usage documents `text`, `json`, and `stream-json` output and the `system/api_retry` schema ([programmatic usage](https://code.claude.com/docs/en/headless#stream-responses)). v2.1.208 fixed truncated JSON/stream-json output and a missing final result message when large `claude -p` output was piped ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.208)).

**Implication:** Treat the non-zero status as the automation contract for failure and capture stdout, stderr, and the exit code separately. For diagnosis, prefer `--output-format stream-json --verbose --include-partial-messages`; empty stderr alone does not locate the failure or show that no retries occurred. Upgrading to 2.1.208+ improves output reliability but is not documented as a 502 retry fix.

### 7. The incident pattern supports a transient failure classification, not hop attribution

Four successful requests followed by one 502 and a successful identical rerun is consistent with a transient failure in the gateway/upstream path. It weighs against a deterministic base-URL, credential, or request-shape error, but it does not identify whether CLIProxyAPI, the OpenAI/Codex upstream, OAuth refresh, an intermediary, or a stream translator generated the 502.

**Evidence:** Claude Code’s error reference classifies 5xx responses on a custom gateway as provider/gateway infrastructure errors and recommends retrying because they are usually temporary ([server errors](https://code.claude.com/docs/en/errors#api-error-500-internal-server-error)). The gateway protocol requires unchanged error forwarding, so a wrapper or translation layer can alter what the client is able to classify ([protocol](https://code.claude.com/docs/en/llm-gateway-protocol#automatic-retry-and-error-forwarding)).

**Implication:** Correlate Claude Code session/request identifiers with CLIProxyAPI request logs and the proxy’s upstream request identifier. The observed 399-second duration is suggestive of an infrastructure deadline but is not evidence of a specific configured timeout without those logs.

### 8. Mitigations separate into durable supported work and bounded transient recovery

The best-supported durable actions are: upgrade from 2.1.205 to at least 2.1.208 or a newer qualified release for print/stream output and adjacent HTTP/2 stream robustness; validate the proxy against the Anthropic Messages streaming contract on every Claude Code upgrade; preserve unmodified upstream errors and request IDs; emit unbuffered SSE and keepalives; and capture `system/api_retry` plus final result records. For an Anthropic-supported deployment, use a Claude-capable upstream behind the gateway.

Transient recovery is narrower: rerun after a terminal 5xx when no visible output or side effect occurred; use an invocation-level retry wrapper only for an idempotent task; or enable `CLAUDE_CODE_RETRY_WATCHDOG=1` for a deliberately unattended invocation that can tolerate multi-hour waits. Increasing `API_TIMEOUT_MS` is not a targeted remedy for this incident because the client received an explicit 502 before the default 10-minute deadline. Disabling the stream watchdog would make silent stalls less bounded.

**Evidence:** v2.1.208 fixes truncated/missing print output and HTTP/2 GOAWAY crashes in supervised/background sessions ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.208)). Claude Code documents the retry watchdog as an unattended-session mechanism and the ordinary 10-retry/10-minute defaults ([automatic retries](https://code.claude.com/docs/en/errors#automatic-retries)). Anthropic documents the gateway stream and forwarding requirements ([protocol](https://code.claude.com/docs/en/llm-gateway-protocol)).

**Implication:** Upgrading improves observability and nearby stream reliability, while proxy/upstream hardening addresses the durable failure surface. A successful manual rerun is evidence for a bounded transient workaround, not proof that the unsupported non-Claude route is reliable.

## Notes

- Historical issue [#33949](https://github.com/anthropics/claude-code/issues/33949) reverse-engineered versions 2.1.50–2.1.74 and reported no stream timeout. It is useful background only; the official v2.1.196 five-minute watchdog supersedes it for 2.1.205.
- The exact client attempt count for the reported failure is unsupported because no `system/api_retry` records, debug log, partial-output record, or per-hop request IDs were supplied.
- Whether SSE `ping` events arrived during the 6m39s wait is unsupported. Their presence would explain why a “no events for five minutes” watchdog did not fire, but this must be observed rather than assumed.
- The exact owner of the 502 and the apparent ~399-second deadline are unsupported without CLIProxyAPI and upstream logs.
- `server_error` is a Claude Code programmatic error category and may also be a proxy normalization. Anthropic’s direct API documents 5xx types such as `api_error`, `timeout_error`, and `overloaded_error`; do not infer a direct Anthropic error from the normalized label.

# Claude Code through CLIProxyAPI to Codex: 502 Research Synthesis

## Best-Supported Answer

The strongest explanation is a **transient terminal failure in the OpenAI/Codex Responses stream that CLIProxyAPI correctly surfaced as a Claude-compatible SSE error**, rather than a deterministic incompatibility between CLIProxyAPI 7.2.91 and Claude Code 2.1.205. The exact incident was: model `gpt-5.6-sol`; operator-owned Codex OAuth; four successful proxied requests; a fifth request that ran for about 6 minutes 39 seconds and ended with an SSE error representing an upstream OpenAI `server_error` as 502; Claude exited 1 with empty stderr; and an identical one-turn rerun through the same launcher completed in about 2 minutes 48 seconds. Four successes plus the successful rerun establish that the configured path was operational, while the non-deterministic failure and OpenAI's documented `response.failed`/`server_error` lifecycle make a transient model-service or infrastructure failure the leading diagnosis. [Compatibility assignment](assignments/01-cliproxyapi-claude-codex-compatibility.md), [translation assignment](assignments/02-anthropic-responses-stream-translation.md), [OpenAI retry assignment](assignments/03-openai-responses-502-retries.md), [OpenAI Responses streaming events](https://developers.openai.com/api/reference/resources/responses/streaming-events).

That conclusion requires one precision: the displayed 502 does **not** prove that OpenAI returned an HTTP response with status 502. In CLIProxyAPI 7.2.91, an upstream Responses SSE `response.failed` or `error` event with an unclassified `server_error` is assigned a 502 status classification. If the downstream stream was already committed, its wire status remains HTTP 200 and CLIProxyAPI sends a terminal Anthropic-format `event: error`; a later attempt to set HTTP 502 cannot replace the committed status. The raw SSE transcript and hop-by-hop headers were unavailable, so an actual pre-stream HTTP 502 remains possible, but it is lower fit for the reported SSE shape. [`codexTerminalFailureStatus`](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/runtime/executor/codex_executor.go#L154), [`forwardClaudeStream`](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/api/handlers/claude/code_handlers.go#L356), [Go `ResponseWriter`](https://pkg.go.dev/net/http#ResponseWriter), [translation assignment](assignments/02-anthropic-responses-stream-translation.md).

The best durable direction is therefore:

1. **Make the next occurrence attributable** with correlated client, proxy, and upstream evidence.
2. **Reduce exposure to long single-turn failures** by checkpointing and decomposing multi-minute work.
3. **Keep one bounded retry owner and prohibit replay after visible output or tool activity.**
4. **Canary version upgrades and route-specific transport changes separately**, without claiming that any identified release fixes this server error.

A brief, bounded rerun after an isolated 5xx is reasonable transient recovery when the operator has confirmed that no visible output or side effect would be duplicated. The successful 2-minute-48-second rerun supports that advice, but it is not a durable fix and is not evidence of exactly-once execution. [Operational assignment](assignments/05-operational-mitigations-and-safety.md), [Claude Code assignment](assignments/04-claude-code-stream-retry-behavior.md), [OpenAI error guidance](https://developers.openai.com/api/docs/guides/error-codes).

## Compatibility and Support Boundary

### Functional interoperability

CLIProxyAPI 7.2.91 contains the concrete bridge required by this route: a Claude-compatible `/v1/messages` handler, a Claude Messages-to-Codex Responses request translator, a Codex OAuth executor posting to the Codex `/responses` endpoint, and a stateful Responses-to-Claude stream translator. Its model registry includes `gpt-5.6-sol`; relevant model support was added and revised before 7.2.91. The incident's successful requests independently confirm that this path functioned in the actual environment. [CLIProxyAPI v7.2.91](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.91), [request translator](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/translator/codex/claude/codex_claude_request.go), [response translator](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/translator/codex/claude/codex_claude_response.go), [compatibility assignment](assignments/01-cliproxyapi-claude-codex-compatibility.md).

### Anthropic support boundary

Functional interoperability is weaker than vendor support. Anthropic documents gateway use for Anthropic-compatible APIs but explicitly does not support routing Claude Code through a gateway to non-Claude models. The CLIProxyAPI-to-`gpt-5.6-sol` route is therefore maintained by the proxy/operator rather than covered by an Anthropic compatibility guarantee. Every Claude Code upgrade can introduce new headers, fields, beta behavior, or stream expectations that the translator must continue to accommodate. [Anthropic gateway guidance](https://code.claude.com/docs/en/llm-gateway), [gateway protocol](https://code.claude.com/docs/en/llm-gateway-protocol), [Claude Code assignment](assignments/04-claude-code-stream-retry-behavior.md).

### `/v1/messages?beta=true`

CLIProxyAPI registers `/v1/messages`; the `?beta=true` query does not change route selection, and the JSON body field `stream: true` selects streaming. The Codex executor constructs a new upstream `/responses` URL and does not forward that query. Anthropic's documented beta mechanism is the `anthropic-beta` header, not this query parameter. Accordingly, the query is compatible with the route but is not an OpenAI Responses option and should not be treated as the cause of the one-off fifth-request failure. [CLIProxyAPI route](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/api/server.go#L535), [stream dispatch](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/api/handlers/claude/code_handlers.go#L70), [Anthropic beta headers](https://platform.claude.com/docs/en/api/beta-headers), [translation assignment](assignments/02-anthropic-responses-stream-translation.md).

The bridge is semantic rather than transparent. Claude messages, tools, images, thinking, and tool results are rewritten into Responses items; selected parameters are mapped, while fields such as Claude `max_tokens`, `temperature`, `top_p`, and `stop_sequences` are not preserved by the core translator as equivalent Responses fields. That is an ongoing compatibility risk, but an identical rerun succeeding weighs against a deterministic translation error for this particular request. [Claude-to-Codex translator](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/translator/codex/claude/codex_claude_request.go#L23), [translation assignment](assignments/02-anthropic-responses-stream-translation.md).

### Authentication ownership

Codex OAuth belongs at CLIProxyAPI. The operator uses CLIProxyAPI's documented `--codex-login`; Claude Code points `ANTHROPIC_BASE_URL` at the proxy origin without appending `/v1` and uses a CLIProxyAPI access key as `ANTHROPIC_AUTH_TOKEN`. CLIProxyAPI applies the upstream bearer credential and ChatGPT account context. OAuth tokens or auth files should never be copied into Claude Code or another retry pool. [CLIProxyAPI Codex OAuth](https://help.router-for.me/configuration/provider/codex), [CLIProxyAPI Claude Code client](https://help.router-for.me/agent-client/claude-code), [Anthropic environment variables](https://code.claude.com/docs/en/env-vars), [compatibility assignment](assignments/01-cliproxyapi-claude-codex-compatibility.md).

## Failure and Retry Path

### Most likely event sequence

1. Claude Code posted a streaming Anthropic-format request to `/v1/messages?beta=true`.
2. CLIProxyAPI translated it and posted a new streaming request to the Codex OAuth `/responses` endpoint for `gpt-5.6-sol`.
3. Normal upstream events would be translated into `message_start`, content deltas, tool blocks, `message_delta`, and `message_stop`.
4. Instead, the executor intercepted an upstream `response.failed` or `error` event before normal translation, extracted `server_error`, and assigned 502 because no more specific status was available.
5. CLIProxyAPI emitted the failure as a Claude-compatible SSE `event: error`; Claude treated the turn as terminal failure and exited 1.
6. A separately created rerun succeeded because the failure condition was transient.

OpenAI documents `response.failed` with `server_error`; CLIProxyAPI 7.2.91 explicitly implements this mapping; and Anthropic documents that SSE errors can arrive after HTTP 200. These protocol facts align without requiring a local timeout or a malformed Claude stream. [OpenAI streaming events](https://developers.openai.com/api/reference/resources/responses/streaming-events), [terminal interception](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/runtime/executor/codex_executor.go#L1520), [Anthropic stream errors](https://platform.claude.com/docs/en/build-with-claude/streaming#error-events), [translation assignment](assignments/02-anthropic-responses-stream-translation.md).

### Retry boundary

CLIProxyAPI and Claude Code both distinguish failures before client-visible payload from failures after output begins. CLIProxyAPI may rotate eligible credentials or retry during stream bootstrap, and its optional `streaming.bootstrap-retries` applies only before payload delivery. After a payload, it forwards the terminal error rather than recreating the request. Claude Code 2.1.199+ likewise retains partial output and avoids replay after visible output because a retry can duplicate tool calls. An early `response.created` translated to `message_start` can close the proxy's safe-bootstrap window even before user-visible text appears. [CLIProxyAPI bootstrap path](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/api/handlers/handlers.go#L1298), [Claude Code automatic retries](https://code.claude.com/docs/en/errors#automatic-retries), [translation assignment](assignments/02-anthropic-responses-stream-translation.md), [Claude Code assignment](assignments/04-claude-code-stream-retry-behavior.md).

Claude Code 2.1.205 inherited the relevant changes from 2.1.196 through 2.1.199: a five-minute no-event stream watchdog, backoff for transient mid-response connection drops, and explicit handling of mid-stream overloaded/server errors. Its documented ordinary policy retries transient classes up to ten times with exponential backoff, while the whole-request default is 600,000 ms. The observed 6m39s is below that 10-minute deadline and may include events, internal retries, and backoff; without `system/api_retry` records, it reveals neither the attempt count nor the watchdog path. [Claude Code v2.1.196](https://github.com/anthropics/claude-code/releases/tag/v2.1.196), [v2.1.198](https://github.com/anthropics/claude-code/releases/tag/v2.1.198), [v2.1.199](https://github.com/anthropics/claude-code/releases/tag/v2.1.199), [Claude Code assignment](assignments/04-claude-code-stream-retry-behavior.md).

### Exit 1 with empty stderr

The exit status is consistent with a terminal SSE failure. Empty stderr is not evidence that Claude Code received no error: print and `stream-json` modes place structured progress, `system/api_retry`, partial output, and result information on stdout, and the public contract does not require every API stream error to be duplicated on stderr. Claude Code 2.1.208 fixed truncated JSON/stream-json output and a missing final result in some piped large-output cases, so upgrading improves capture reliability, but the exact 2.1.205 mapping for this invocation remains unsupported without its flags, stdout, and debug log. [Claude Code programmatic usage](https://code.claude.com/docs/en/headless#stream-responses), [Claude Code v2.1.208](https://github.com/anthropics/claude-code/releases/tag/v2.1.208), [Claude Code assignment](assignments/04-claude-code-stream-retry-behavior.md).

## Recommended Durable Direction

### 1. Correlate every layer before tuning

Retain a sanitized record per network attempt: logical operation ID; unique `X-Client-Request-Id`; CLIProxyAPI `X-CPA-TRACE-ID`; OpenAI `x-request-id`; Responses `resp_...` ID when available; exact timestamp and timezone; model and versions; client retry number; concurrency; time to headers, first event, last event, and terminal event; last SSE type and sequence number; selected credential label without secret material; cooldown transition; proxy/intermediary error; and process RSS/CPU. Enable Claude Code `--debug "api"` and a dedicated debug file only for a bounded reproduction window, and capture stdout, stderr, and exit status separately. Prefer CLIProxyAPI's error-only logs; full request logging may contain large prompt/response bodies and requires restricted access and redaction. OpenAI explicitly recommends logging `x-request-id` and supports caller-supplied `X-Client-Request-Id`. [OpenAI API debugging](https://developers.openai.com/api/reference/overview#debugging-requests), [CLIProxyAPI management API](https://help.router-for.me/management/api#request-error-logs), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

Use distinct health signals: `/healthz` for proxy-process liveness, a low-frequency short inference for authentication/upstream readiness, and long-stream completion as the real reliability indicator. CLIProxyAPI's `/healthz` is a static local check and does not prove OAuth validity, upstream reachability, model availability, or stream completion. [`/healthz`](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/internal/api/server.go#L501), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

### 2. Decompose long work into checkpointed turns

For tasks expected to take minutes, split source mapping, subsystem inspection, checkpoint writing, and synthesis into independently verifiable turns. Narrow context and expected output, preserve durable intermediate results, and resume from the last completed checkpoint rather than replaying an entire monolithic turn. This reduces the time exposed to transient failure, the amount of discarded model work, and the cost and side-effect risk of retry. Both Claude Code and OpenAI recommend decomposition or smaller completion bounds for long or capacity-sensitive work. [Claude Code timeout guidance](https://code.claude.com/docs/en/errors#request-timed-out), [Claude Code troubleshooting](https://code.claude.com/docs/en/troubleshooting#auto-compaction-stops-with-a-thrashing-error), [OpenAI rate-limit guidance](https://developers.openai.com/api/docs/guides/rate-limits#reduce-the-max_tokens-to-match-the-size-of-your-completions), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

### 3. Use one bounded retry owner

For ordinary interactive use, Claude Code is the best retry owner because it knows whether output or tool calls became visible. Preserve the no-replay boundary after visible output. Avoid a process wrapper that blindly retries exit 1, because Claude Code may already have consumed a substantial retry budget and a new process creates a new response. A single manual rerun after a brief jittered wait is proportionate only after checking that the logical operation and any tools are safe to repeat. OpenAI's guidance supports bounded exponential backoff and warns that failed requests still consume rate-limit budget. [Claude Code automatic retries](https://code.claude.com/docs/en/errors#automatic-retries), [OpenAI rate-limit retry guidance](https://developers.openai.com/api/docs/guides/rate-limits#retrying-with-exponential-backoff), [OpenAI retry assignment](assignments/03-openai-responses-502-retries.md).

If proxy-side bootstrap recovery is needed for a measured pre-payload failure, canary at most `streaming.bootstrap-retries: 1` and reduce the client retry cap during that test so nested attempts cannot multiply. Do not enable `CLAUDE_CODE_RETRY_WATCHDOG=1` for interactive work; its documented behavior can extend transient retries to roughly three hours and retry capacity errors indefinitely. [CLIProxyAPI streaming configuration](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/config.example.yaml#L191), [Claude Code errors](https://code.claude.com/docs/en/errors#automatic-retries), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

### 4. Smooth concurrency and preserve rate-limit evidence

Begin reproductions at concurrency one and change concurrency separately from versions, retry values, or keep-alives. Coordinate retry budgets across sessions and add jitter so multiple long requests do not resubmit together. A 502 is not proof of quota exhaustion: OpenAI normally represents rate limiting as 429 and exposes `x-ratelimit-*` headers. No supplied source establishes a numeric concurrency or rate limit for operator-owned Codex OAuth. [OpenAI rate limits](https://developers.openai.com/api/docs/guides/rate-limits), [OpenAI API headers](https://developers.openai.com/api/reference/overview#debugging-requests), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

### 5. Canary upgrades separately

Remain on at least CLIProxyAPI 7.2.91. Version 7.2.80 added correct recognition of `error`, `response.failed`, and `response.incomplete`; 7.2.91 added sanitization of overlong encrypted reasoning IDs. CLIProxyAPI 7.2.92 adds payload-normalization and allocation optimizations but no documented 502, retry, or terminal-stream fix. Claude Code versions after 2.1.205 add nearby hardening: 2.1.208 improves piped structured output and HTTP/2 GOAWAY handling; 2.1.214 opens a fresh socket after stale-connection errors. Canary the exact proxy and client versions one at a time, retain the known-good binaries, drain active streams before restart, and compare completion rate, latency, event gaps, retry counts, and error taxonomy before retaining a change. [CLIProxyAPI v7.2.80](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.80), [v7.2.91](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.91), [v7.2.92](https://github.com/router-for-me/CLIProxyAPI/releases/tag/v7.2.92), [Claude Code changelog](https://github.com/anthropics/claude-code/blob/main/CHANGELOG.md), [compatibility assignment](assignments/01-cliproxyapi-claude-codex-compatibility.md), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

These upgrades are general hardening and observability improvements. No supplied evidence identifies a post-7.2.91 CLIProxyAPI fix or a post-2.1.205 Claude Code fix for an OpenAI-generated `response.failed`/`server_error`.

## Conditional Alternatives

### Background/resumable Responses

OpenAI background mode is the strongest protocol-level design for model work that takes several minutes: create with `background: true`, retain the response ID, poll terminal state, and resume a dropped background stream using `sequence_number` and `starting_after`. It avoids tying model execution to one long-lived SSE connection and can recover the same response rather than creating duplicate work. [OpenAI background mode](https://developers.openai.com/api/docs/guides/background), [OpenAI retry assignment](assignments/03-openai-responses-502-retries.md).

This is a conditional architecture direction, not a configuration fix for the current Claude-to-Codex path. CLIProxyAPI 7.2.91's translator sets `store=false`, does not set `background=true`, and the supplied evidence does not establish preservation of response creation IDs, retrieval, terminal status, sequence numbers, or `starting_after` through `/v1/messages`. It should be considered only if that full lifecycle is deliberately supported end to end.

### SSE keep-alives and reverse-proxy tuning

`streaming.keepalive-seconds: 15` is a reasonable route-specific canary only when evidence shows downstream idle disconnects or a gateway no-byte watchdog. CLIProxyAPI writes standard SSE comments after the first translated payload; they can keep the Claude-Code leg active but cannot make the OpenAI leg produce events or repair a terminal `response.failed`. If an nginx-like intermediary actually exists, disable SSE buffering and set its per-read timeout above the largest measured inter-event gap. Do not apply these changes when Claude Code connects directly to localhost and the observed failure is an explicit upstream server error. [CLIProxyAPI stream forwarder](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/api/handlers/stream_forwarder.go#L45), [WHATWG SSE guidance](https://html.spec.whatwg.org/dev/server-sent-events.html#authoring-notes), [nginx `proxy_read_timeout`](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_read_timeout), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

### Cooldown tuning

CLIProxyAPI cools credentials after transient 408/5xx results. With one credential, the legacy 60-second transient cooldown can exceed the example 30-second maximum retry interval and create a temporary local `auth_unavailable` period. If logs prove that sequence, a controlled canary using a modest positive cooldown such as 10–15 seconds can test whether false blackout time falls. This is lower fit than correlation and workload decomposition, should change only one variable, and should be reverted unless repeatable evidence improves. Blanket `disable-cooling: true` can hot-loop against a failing upstream and remains rejected. [CLIProxyAPI configuration](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/config.example.yaml#L116), [auth conductor](https://github.com/router-for-me/CLIProxyAPI/blob/v7.2.91/sdk/cliproxy/auth/conductor.go#L4647), [compatibility assignment](assignments/01-cliproxyapi-claude-codex-compatibility.md), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

Already-present, independently operator-authorized credentials can improve selection or failover before payload delivery. Adding accounts or rotating credentials to evade quota, concurrency, or cooldown behavior is unsupported and unsafe.

## Lower-Fit or Rejected Interpretations

- **A universal 399/400-second OpenAI timeout:** unsupported. The documented Claude Code whole-request default is 10 minutes; CLIProxyAPI's Codex streaming client and HTTP server set no corresponding total deadline. A reverse proxy or network device remains possible only if it exists on the actual path and its logs match the request. [Operational assignment](assignments/05-operational-mitigations-and-safety.md), [OpenAI retry assignment](assignments/03-openai-responses-502-retries.md).
- **A deterministic Claude Code 2.1.205 incompatibility or regression:** lower fit. The path succeeded four times and on rerun, and 2.1.205 introduced no documented network/SSE policy change. [Claude Code v2.1.205](https://github.com/anthropics/claude-code/releases/tag/v2.1.205), [Claude Code assignment](assignments/04-claude-code-stream-retry-behavior.md).
- **A deterministic payload translation defect:** possible but unproven. The bridge drops or rewrites fields and is not Anthropic-supported for non-Claude models, yet an identical rerun succeeding weighs against a stable request-shape failure.
- **Rate limiting or a documented Codex OAuth concurrency cap:** unsupported. The failure class was 502/`server_error`, while documented rate limiting is normally 429; no authoritative numeric OAuth limit was found.
- **An abrupt upstream EOF:** lower fit. CLIProxyAPI 7.2.91 normally maps a stream ending before `response.completed` to an incomplete-stream 408-style error, not the observed 502/`server_error` path. [Translation assignment](assignments/02-anthropic-responses-stream-translation.md).
- **The July 17, 2026 `gpt-5.6-sol` overload incident as the cause:** plausible context only. OpenAI recorded a matching model-specific failure class, but the failed request's exact timestamp and timezone were not supplied. [OpenAI status incident](https://status.openai.com/incidents/01KXRHE25717D2WQ1WFMT2B7WZ), [OpenAI retry assignment](assignments/03-openai-responses-502-retries.md).
- **Keep-alives, larger deadlines, or timeout disablement as a cure for an explicit `response.failed`:** rejected. These can alter transport liveness but cannot heal a model service that deliberately terminates the response.
- **Automatic replay after partial output or tool activity:** rejected because it can duplicate tool effects and generated work.
- **Downgrading below CLIProxyAPI 7.2.91:** rejected. Older versions lose terminal-event correctness or the long-reasoning-ID protection.
- **Treating the successful identical rerun as idempotent:** rejected. No reviewed official source promises exactly-once creation for `POST /v1/responses`; a new create can repeat computation or side effects. [OpenAI retry assignment](assignments/03-openai-responses-502-retries.md).

## Unsafe and Unsupported Workarounds

The following remain outside the safe solution space:

- Copying, exporting, scraping, decrypting, or otherwise extracting Codex OAuth tokens or auth JSON for another process, host, account, or retry pool.
- Adding accounts or rotating credentials to evade provider quota, concurrency, or cooldown behavior.
- Deleting or resetting CLIProxyAPI authentication state, Claude Code state, caches, keychain entries, or configuration in response to a transient 502.
- Permanently modifying shell profiles or global Claude Code settings during diagnosis; any canary should use documented invocation-scoped or operator-owned temporary configuration and be reverted.
- Disabling TLS verification or bypassing organizational proxy and security controls.
- Enabling unbounded retries, `CLAUDE_CODE_RETRY_WATCHDOG=1` for interactive work, extremely large deadlines, global idle-timeout removal, `disable-cooling: true`, or repeated immediate reruns.
- Suppressing terminal errors, converting failed streams to successful messages, or automatically replaying after output/tool activity.
- Killing or restarting CLIProxyAPI with active streams; drain streams and use exact-version binary rollback instead of destructive state reset.

These boundaries preserve the gateway's documented credential ownership and the stream replay safety boundary. [Anthropic gateway guidance](https://code.claude.com/docs/en/llm-gateway#what-a-gateway-provides), [OpenAI authentication guidance](https://developers.openai.com/api/reference/overview#authentication), [CLIProxyAPI basic configuration](https://help.router-for.me/configuration/basic), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

## Candidate Classification

### Selected for deep dive

- CLIProxyAPI v7.2.91 handler, Claude/Codex translators, Codex executor, stream forwarder, auth conductor, exact configuration, model registry, trace/header handling, health check, and management/error logs.
- CLIProxyAPI v7.2.80 terminal-event fix, v7.2.91 long-reasoning-ID fix, v7.2.19 configurable transient cooldown, and v7.2.92 release delta.
- Anthropic Messages streaming/error contracts; Claude Code gateway protocol, error/retry/timeout behavior, programmatic output, releases 2.1.196–2.1.215, and gateway support boundary.
- OpenAI Responses streaming events, background/resume lifecycle, error and rate-limit guidance, request IDs, `gpt-5.6-sol` documentation, SDK retry source, and relevant status incidents.
- WHATWG/RFC SSE keep-alive guidance and nginx buffering/per-read timeout semantics for conditional transport analysis.

These sources control the diagnosis and mitigations because they are primary source, exact tagged production code, or official protocol documentation. [Compatibility assignment](assignments/01-cliproxyapi-claude-codex-compatibility.md), [translation assignment](assignments/02-anthropic-responses-stream-translation.md), [OpenAI retry assignment](assignments/03-openai-responses-502-retries.md), [Claude Code assignment](assignments/04-claude-code-stream-retry-behavior.md), [operational assignment](assignments/05-operational-mitigations-and-safety.md).

### Skim-only corroboration

- CLIProxyAPI issue #4290, because it reports the same model returning 502 followed by credential cooldown and later recovery, but it is issue evidence rather than implementation authority.
- Older reports involving large-prompt incomplete streams, third-party Claude relays, direct Windows SSE 502s, historical Claude Code stream hangs, and Codex stream errors. They show adjacent symptom classes but differ in provider, version, platform, or failure signature.
- OpenAI status history, including the July 17 `gpt-5.6-sol` overload incident. It establishes that transient model-specific overload exists, but cannot be correlated without the incident timestamp.

### Rejected or lower-fit candidates

- Unverified GitHub issue workarounds, community proxy recipes, generic third-party SSE blogs, unrelated Anthropic upstream billing/system-prompt issues, and custom Anthropic-compatible relay implementations.
- Unofficial or undocumented Claude Code environment variables, including `CLAUDE_CODE_DISABLE_NONSTREAMING_FALLBACK`.
- Credential-extraction, account-fan-out, and generic “clear everything” troubleshooting instructions.

## Unresolved Evidence and Unsupported Claims

The supplied evidence does not establish:

- the exact failed request timestamp and timezone, so no OpenAI status incident can be attributed;
- whether the upstream terminal event was `response.failed` or `error`, or whether OpenAI instead returned an actual HTTP 502;
- the raw downstream status line, `Server`/`Via` headers, response ID, request IDs, or last SSE event;
- whether `response.created`, visible text, thinking, or a tool event preceded failure;
- whether Claude Code performed internal retries, because no `system/api_retry` records or debug log were supplied;
- whether keep-alive comments or SSE events reset the five-minute gateway idle watchdog;
- the CLIProxyAPI retry, bootstrap, cooldown, passthrough-header, logging, or reverse-proxy configuration used in the incident;
- whether a reverse proxy, load balancer, NAT, firewall, socket failure, or OAuth refresh participated;
- any numeric concurrency or rate limit for operator-owned Codex OAuth;
- any exactly-once guarantee for a recreated Responses request;
- any CLIProxyAPI release after 7.2.91 through 7.2.92, or Claude Code release after 2.1.205 reviewed here, that specifically fixes the recorded OpenAI `server_error`;
- full Anthropic support for routing Claude Code to `gpt-5.6-sol` through a translating gateway.

Until those artifacts exist, the correct conclusion is **leading diagnosis, not proven root cause**: a transient Responses terminal failure is best supported; the proxy's 502/SSE mapping is expected; and durable reliability comes from correlation, shorter checkpointed work, controlled concurrency, and a single bounded retry owner rather than credential manipulation, destructive resets, or unlimited replay.

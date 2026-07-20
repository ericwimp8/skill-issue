# OpenAI Responses API 502 and Retry Handling

## Assignment

**Goal.** Establish the best-supported OpenAI guidance for handling a long-running streamed Responses API operation that ends in an HTTP 502 or an SSE `server_error`, and apply that guidance cautiously to the reported CLIProxyAPI 7.2.91 / Claude Code 2.1.205 case: four successful requests, a fifth request failing after 6m39s, and an identical rerun completing in 2m48s against `gpt-5.6-sol` through operator-owned Codex OAuth/subscription access.

**Scope.** Internet-only research using official OpenAI API documentation, official OpenAI SDK documentation and source, official model documentation, and OpenAI Status incident records. The research covers retry classification, backoff, timeout behavior, resumable/background execution, duplicate-work risk, stream terminal events, observability, and the limits of applying supported-API guidance through a translating proxy.

**Exclusions.** No inspection of CLIProxyAPI source, configuration, logs, or network traces; no credential, token, OAuth, or authentication extraction advice; no claim that a particular component emitted the reported 502 without hop-by-hop evidence; and no claim that official SDK defaults apply to CLIProxyAPI's transport.

## Sources

- OpenAI API, [Error codes](https://developers.openai.com/api/docs/guides/error-codes), accessed 2026-07-21.
- OpenAI API, [Background mode](https://developers.openai.com/api/docs/guides/background), accessed 2026-07-21.
- OpenAI API Reference, [Responses streaming events](https://developers.openai.com/api/reference/resources/responses/streaming-events), accessed 2026-07-21.
- OpenAI API Reference, [Get a model response](https://developers.openai.com/api/reference/resources/responses/methods/retrieve), accessed 2026-07-21.
- OpenAI API Reference, [API debugging headers and client request IDs](https://developers.openai.com/api/reference/overview), accessed 2026-07-21.
- OpenAI, [openai-node official SDK README](https://github.com/openai/openai-node), `master` as inspected 2026-07-21.
- OpenAI, [openai-python official SDK README](https://github.com/openai/openai-python), `main` as inspected 2026-07-21.
- OpenAI, [openai-python retry implementation](https://github.com/openai/openai-python/blob/main/src/openai/_base_client.py), `main` as inspected 2026-07-21.
- OpenAI API, [GPT-5.6 Sol model page](https://developers.openai.com/api/docs/models/gpt-5.6-sol), accessed 2026-07-21.
- OpenAI Help Center, [Using Codex with your ChatGPT plan](https://help.openai.com/en/articles/11369540-using-codex-with-your-chatgpt-plan), accessed 2026-07-21.
- OpenAI Status, [Codex 5.6-sol Experiencing Increased Server-Overload Errors](https://status.openai.com/incidents/01KXRHE25717D2WQ1WFMT2B7WZ), resolved 2026-07-17.
- OpenAI Status, [Elevated error rates on Codex, ChatGPT and Responses API](https://status.openai.com/incidents/01KT5XJ5ATD6RMYP908WS69FVD/write-up), incident 2026-06-02/03.
- OpenAI Status, [Elevated API Errors](https://status.openai.com/incidents/fk0tcbydtybr), incident 2023-10-19.
- OpenAI Status, [Increased errors during streaming Chat Completions](https://status.openai.com/incidents/01JMYB5ZDZNA3RN7KZ4CE0FZXK), incident 2023-11-27/28.

## Findings

### Finding 1: A 502 belongs to the documented retryable server-error class, but the reported source of the 502 is unproven

OpenAI's error guide explicitly describes 500 errors as server-side failures and 503 errors as overload conditions, recommending a brief wait and retry. The official Python and Node SDKs generalize this into a concrete rule: every HTTP status `>=500`, which includes 502, is an `InternalServerError` and is retryable by default. Separately, a Responses SSE connection can return HTTP 200 and later emit a terminal `response.failed` event whose embedded error code is `server_error`; the API reference also defines a general `error` SSE event. An HTTP 502 and an SSE `response.failed` are therefore distinct failure surfaces even when a proxy presents both as a single high-level error.

**Evidence.** The error guide labels 500 as an issue on OpenAI's servers and 503 as overload, with retry guidance ([Error codes](https://developers.openai.com/api/docs/guides/error-codes)). Both official SDK READMEs state that `>=500` errors are retried automatically ([openai-node](https://github.com/openai/openai-node), [openai-python](https://github.com/openai/openai-python)). The Responses streaming reference defines `response.failed`, including an example with `error.code: "server_error"`, and a separate `error` event ([Responses streaming events](https://developers.openai.com/api/reference/resources/responses/streaming-events)).

**Implication.** Treat the failure as potentially transient, while retaining the distinction between: an upstream HTTP 502 before SSE starts; a downstream/proxy-generated 502; a transport disconnect after HTTP 200; and a valid SSE terminal failure. The error class alone cannot identify which hop failed.

### Finding 2: Official SDK retry behavior is bounded exponential backoff, and stacked retry layers are a practical risk

The official Python and Node SDKs retry connection failures, 408, 409, 429, and `>=500` twice by default, for up to three total attempts, using short exponential backoff. The Python implementation respects `retry-after-ms` or `Retry-After` when the requested delay is between 0 and 60 seconds, otherwise calculates exponential delay with jitter; it also honors OpenAI's `x-should-retry` response header when present. These SDK facts do not establish CLIProxyAPI's policy, because the proxy may use another client, may disable SDK retries, or may add its own retry loop.

**Evidence.** Both official SDK READMEs document two automatic retries and per-client/per-request retry configuration ([openai-node](https://github.com/openai/openai-node), [openai-python](https://github.com/openai/openai-python)). The official Python retry source checks `x-should-retry`, retries all `>=500` statuses, respects server retry-delay headers, and adds exponential backoff with jitter ([`_base_client.py`](https://github.com/openai/openai-python/blob/main/src/openai/_base_client.py)). OpenAI's 503 guidance also recommends exponential backoff or respecting response headers ([Error codes](https://developers.openai.com/api/docs/guides/error-codes)).

**Implication.** Use one bounded retry owner. Before adding a CLI-level retry, determine whether the proxy and upstream client already retry. Record attempt count and elapsed time for every layer so one visible request cannot silently expand into many costly long-running attempts. For a single transient 502, a brief jittered retry is consistent with OpenAI guidance; repeated failures should stop after the configured budget, check status, and escalate with trace evidence rather than loop indefinitely.

### Finding 3: Background mode is OpenAI's durable mechanism for multi-minute Responses work

OpenAI explicitly positions background mode for reasoning tasks that take several minutes and says it avoids dependence on one long-lived synchronous connection. A client creates a Response with `background: true`, retains the response ID, and polls while its status is `queued` or `in_progress`. A background response can also be created with `stream: true`; if the stream drops, the client can resume from the last SSE `sequence_number` using `starting_after`. Background streaming has higher time-to-first-token than synchronous streaming.

**Evidence.** The background guide says long-running reasoning can take several minutes, describes `background: true` as reliable execution without timeout/connectivity concerns, and documents polling the response ID until a terminal state ([Background mode](https://developers.openai.com/api/docs/guides/background)). The same guide documents `background: true` plus `stream: true`, saving the event cursor, and reconnecting with `GET /v1/responses/{id}?stream=true&starting_after={sequence}`. The retrieve endpoint is formally documented as `GET /responses/{response_id}` ([Get a model response](https://developers.openai.com/api/reference/resources/responses/methods/retrieve)).

**Implication.** If the integration can use the supported Responses API semantics end to end, background execution plus saved response ID and sequence cursor is the strongest durable fix for a 2-7 minute operation. It separates model execution from the health of one proxy/SSE socket and allows recovery without creating a second Response. Confirm first that CLIProxyAPI preserves `background`, the response ID, retrieval, `stream`, `sequence_number`, and `starting_after`; official OpenAI documentation cannot establish that proxy capability.

### Finding 4: Automatic HTTP retries do not imply automatic recovery from a failure after streaming has begun

Official SDK retry documentation covers connection establishment, timeouts, and retryable HTTP status responses. The Python request implementation evaluates the HTTP status inside the request retry loop before returning the stream iterator. Once a successful HTTP response has established an SSE stream, a later `response.failed` or stream interruption is part of stream consumption rather than a new retryable HTTP response in that loop. OpenAI documents resumable streaming specifically through background mode.

**Evidence.** The official Python request source sends the HTTP request, evaluates status, and retries a retryable status before processing/returning the response; its SSE iterator is consumed afterward ([`_base_client.py`](https://github.com/openai/openai-python/blob/main/src/openai/_base_client.py)). The streaming API defines explicit terminal `response.completed` and `response.failed` events ([Responses streaming events](https://developers.openai.com/api/reference/resources/responses/streaming-events)). The background guide, rather than the standard streaming guide, provides a reconnection cursor for dropped streams ([Background mode](https://developers.openai.com/api/docs/guides/background)).

**Implication.** A proxy must classify whether it saw a non-2xx HTTP response before any SSE events or a terminal/transport failure after `response.created`. Retrying a pre-stream 502 may be automatic. Recovering a midstream failure safely requires the existing response ID and resumable background stream, or a deliberate new logical attempt with duplicate-work safeguards.

### Finding 5: No reviewed official source promises exactly-once creation for `POST /v1/responses`

The official sources reviewed document retry policy, request correlation IDs, response IDs, and idempotent cancellation, but they do not document a public Responses-create idempotency header or an exactly-once guarantee. The background guide explicitly calls repeated cancellation idempotent; it does not make that statement for response creation. Automatic SDK retry support is evidence that retry is expected, but it is not evidence that two independently issued create requests are deduplicated.

**Evidence.** The official Responses create/background materials expose response IDs, retrieval, polling, cancellation, and resumable streaming, while documenting idempotency only for cancelling the same response twice ([Background mode](https://developers.openai.com/api/docs/guides/background)). The API overview documents `X-Client-Request-Id` as a unique diagnostic identifier for each request, not as a deduplication key ([API overview](https://developers.openai.com/api/reference/overview)). No reviewed official API reference or SDK README documents exactly-once `POST /responses` semantics.

**Implication.** Do not describe an identical rerun as idempotent. If the first request reached OpenAI but its completion was lost at the proxy, a new create request can represent additional model work and can repeat generated tool calls or downstream side effects. Prefer retrieving/resuming the existing response when its ID is known. At the application boundary, assign a logical-operation ID, record each unique network-attempt ID, and deduplicate any side-effecting tool execution using stable application records or tool-call identifiers.

### Finding 6: The 6m39s timing does not match a documented universal OpenAI timeout

The official Python and Node SDKs currently use a 10-minute default request timeout and retry timeout failures twice by default. The official Go SDK, by contrast, documents no default request timeout and expects callers to use context. Timeout behavior therefore belongs to the concrete client and every intermediary, not to a universal Responses API cutoff. A failure at 6m39s is below the official Python/Node default and cannot by itself identify OpenAI, the proxy, Claude Code, a load balancer, or an idle/read timer as the source.

**Evidence.** The Node and Python SDK READMEs document a 10-minute default, configurable timeouts, and retry-on-timeout behavior ([openai-node](https://github.com/openai/openai-node), [openai-python](https://github.com/openai/openai-python)). OpenAI's error guide says a timeout can result from network issues, load, or a complex request and recommends waiting briefly and retrying ([Error codes](https://developers.openai.com/api/docs/guides/error-codes)). Background mode is the documented alternative for long work ([Background mode](https://developers.openai.com/api/docs/guides/background)).

**Implication.** Instrument total timeout, connect timeout, time-to-headers, time-to-first-SSE-event, idle/read timeout, time since last SSE event, and upstream processing duration separately at Claude Code, CLIProxyAPI, any reverse proxy, and the OpenAI client. Align configured timeouts only after locating the terminating layer. For inherently multi-minute work, prefer background execution over merely raising every synchronous timeout.

### Finding 7: Request IDs and stream lifecycle data are essential to localizing the fault

OpenAI returns `x-request-id` and recommends logging it in production. Callers can add a unique `X-Client-Request-Id` per network request; OpenAI logs it for supported endpoints including Responses and can use it to locate requests even when a timeout prevents receipt of `x-request-id`. Official SDK errors expose request IDs, and streaming responses can expose the request ID through raw/with-response access. For persistent errors, OpenAI asks for model, error message/code, request data and headers, timestamp, and timezone.

**Evidence.** The API overview documents `x-request-id`, `openai-processing-ms`, rate-limit headers, and unique `X-Client-Request-Id` values, and recommends production logging ([API overview](https://developers.openai.com/api/reference/overview)). The official Node and Python SDK READMEs document retrieving request IDs from successful responses and status errors; Node also documents obtaining a request ID alongside a stream ([openai-node](https://github.com/openai/openai-node), [openai-python](https://github.com/openai/openai-python)). OpenAI's persistent-error guidance lists timestamp/timezone, model, error, request data, and headers ([Error codes](https://developers.openai.com/api/docs/guides/error-codes)).

**Implication.** For each attempt, log a sanitized correlation record containing logical-operation ID, unique client-request ID, OpenAI request ID, Responses `resp_...` ID, model, start/end timestamps with timezone, retry count, HTTP status source, `openai-processing-ms`, first/last SSE sequence numbers and event types, and which hop closed the connection. Preserve upstream and downstream status separately. Never log authorization or OAuth material.

### Finding 8: OpenAI Status confirms that overload, Responses latency/failures, 502s, and streaming `server_error` can be transient incidents

OpenAI has published multiple relevant incidents. On 2026-07-17, `gpt-5.6-sol` had increased server-overload errors. On 2026-06-02/03, shared infrastructure degradation caused elevated Responses latency before model processing and failures, while also affecting Codex. Earlier incidents explicitly recorded widespread API 500/502 failures and streaming `server_error` failures. These records validate transient infrastructure failure as a real class, but they do not prove that the reported request fell within one of those windows.

**Evidence.** OpenAI Status resolved a `gpt-5.6-sol` server-overload incident on 2026-07-17 ([incident](https://status.openai.com/incidents/01KXRHE25717D2WQ1WFMT2B7WZ)). The June 2026 write-up says Responses requests experienced elevated latency before model processing and failures due to shared backend degradation ([write-up](https://status.openai.com/incidents/01KT5XJ5ATD6RMYP908WS69FVD/write-up)). The October 2023 write-up records 500/502 API failures ([incident](https://status.openai.com/incidents/fk0tcbydtybr)), and the November 2023 incident records streaming `server_error` spikes ([incident](https://status.openai.com/incidents/01JMYB5ZDZNA3RN7KZ4CE0FZXK)). Status pages caution that availability metrics are aggregated and individual impact varies by tier, model, and feature.

**Implication.** Check and preserve the status page at the exact request timestamp, including timezone. If the failure overlaps a matching incident, wait for recovery before consuming a large retry budget. Without the failed request's exact timestamp and request IDs, historical incidents are corroboration only.

### Finding 9: The reported retry success supports transient handling, while the root cause remains unsupported

The sequence of four successes, one long failure, and one faster successful rerun is compatible with a transient capacity, infrastructure, connection, or intermediary failure. It is also compatible with a proxy-specific timeout or translation defect. The elapsed-time difference does not distinguish those hypotheses because the two creates may have had different queuing, reasoning, tool, network, and proxy behavior even with identical visible input.

**Evidence.** OpenAI's error guide says internal failures may be temporary and a second attempt may succeed; official status records show model-specific overload and Responses latency/failure incidents ([Error codes](https://developers.openai.com/api/docs/guides/error-codes), [2026-07-17 status incident](https://status.openai.com/incidents/01KXRHE25717D2WQ1WFMT2B7WZ)). `gpt-5.6-sol` officially supports the Responses endpoint and streaming ([GPT-5.6 Sol model page](https://developers.openai.com/api/docs/models/gpt-5.6-sol)). OpenAI documents ChatGPT-plan use for its Codex clients, while the reviewed sources do not document CLIProxyAPI or guarantee preservation of Responses semantics through its Claude-to-OpenAI translation ([Using Codex with your ChatGPT plan](https://help.openai.com/en/articles/11369540-using-codex-with-your-chatgpt-plan)).

**Implication.** The immediate transient response is a brief, bounded, jittered retry after checking status. The durable response is proxy instrumentation plus background/resumable execution where supported. Do not conclude from the successful rerun that the failure was definitively an OpenAI outage, definitively a proxy timeout, or harmless with respect to duplicate computation.

### Finding 10: Recommended handling separates durable engineering from transient operations

Durable handling should make the request recoverable and diagnosable: preserve response and request IDs; make background mode the default for expected multi-minute work; resume rather than recreate when possible; designate one retry owner; expose sanitized hop-by-hop timing and status; align timeouts with measured execution; and deduplicate side effects. Transient operations should be small and bounded: wait briefly with jitter, respect retry headers, retry a 502 only within the configured attempt/elapsed-time budget, check the exact-time status page, and stop/escalate when failures persist.

**Evidence.** These measures are direct combinations of OpenAI's internal-error retry guidance, official SDK retry rules, background polling/resume design, request-ID guidance, and persistent-error support requirements ([Error codes](https://developers.openai.com/api/docs/guides/error-codes), [Background mode](https://developers.openai.com/api/docs/guides/background), [API overview](https://developers.openai.com/api/reference/overview), [openai-node](https://github.com/openai/openai-node), [openai-python](https://github.com/openai/openai-python)).

**Implication.** A blind unlimited rerun loop or a single globally enlarged timeout is lower-quality handling. The preferred order is: identify the failing hop; recover the same background response when possible; otherwise perform one bounded logical retry with duplicate-work controls; and retain evidence suitable for proxy maintainers or OpenAI support.

## Notes

- **Unsupported interpretation:** The 6m39s duration does not establish a 400-second OpenAI timeout. No reviewed official source documents such a Responses API limit.
- **Unsupported interpretation:** The final displayed 502 does not establish that OpenAI originated the HTTP response. CLIProxyAPI or another intermediary could have mapped an upstream disconnect or SSE `server_error` to 502.
- **Unsupported interpretation:** An identical prompt does not establish an idempotent retry or exactly-once model execution. No reviewed official source promises exactly-once creation for `POST /v1/responses`.
- **Caveat:** Official SDK defaults describe those SDKs, not necessarily CLIProxyAPI 7.2.91, Claude Code 2.1.205, or subscription-backed Codex OAuth transport.
- **Caveat:** Background mode is a strong durable direction only if the proxy exposes and preserves response creation, response IDs, retrieval, terminal status, streaming sequence numbers, and resume parameters.
- **Caveat:** The 2026-07-17 `gpt-5.6-sol` incident is relevant corroboration, but correlation requires the failed request's exact timestamp and timezone.
- **Useful searches:** `site:status.openai.com gpt-5.6-sol server overload`, `site:developers.openai.com background responses starting_after`, `site:github.com/openai/openai-python _should_retry`, `site:github.com/openai/openai-node maxRetries timeout`.

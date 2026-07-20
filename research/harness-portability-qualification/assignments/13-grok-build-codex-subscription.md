# Grok Build Codex-Subscription Second Gate

## Assignment

**Goal:** Determine whether an existing paid OpenAI Codex or ChatGPT subscription and its included model access can drive the first-party Grok Build CLI without separate xAI/Grok access or API credits.

**Scope:** Grok Build's documented direct authentication, model endpoint configuration, and third-party BYOK surface; OpenAI's documented separation of ChatGPT/Codex and API billing; the first-party `openai/codex` authentication surface; and the existing local Cursor, Claude Code compatibility-launcher, and Pi routes.

**Exclusions:** Installation, login, model calls, credentials or account identities, extraction of OAuth/session data, third-party proxy implementation, permanent user-configuration changes, and a first-gate or full product-support decision.

## Sources

- [Grok Build Enterprise Deployments](https://docs.x.ai/build/enterprise) — documented native authentication methods, first-party xAI API-key route, external auth-provider contract, and third-party BYOK endpoint statement.
- [Grok Build Settings](https://docs.x.ai/build/settings) — model entries' `base_url`, `env_key`, and `api_backend` fields (`chat_completions`, `responses`, and `messages`).
- [Grok Build CLI Reference](https://docs.x.ai/build/cli/reference) and [Headless & Scripting](https://docs.x.ai/build/cli/headless-scripting) — `grok login`, device authentication, model listing, and the documented requirement to log in or supply `XAI_API_KEY` for native headless use.
- [OpenAI Help: ChatGPT subscription and API service](https://help.openai.com/en/articles/8156019-is-api-usage-included-in-chatgpt-subscriptions-even-if-i-have-a-paid-chatgpt-account) and [OpenAI Help: billing settings](https://help.openai.com/en/articles/9039756-billing-settings-in-chatgpt-vs-platform) — ChatGPT and OpenAI API billing are separate; API usage is pay-as-you-go.
- [OpenAI Codex repository README](https://github.com/openai/codex) and [Codex app-server authentication reference](https://github.com/openai/codex/blob/main/codex-rs/app-server/README.md) — ChatGPT-plan sign-in is a Codex-owned flow, while API-key access is a distinct mode and Codex owns ChatGPT OAuth refresh tokens.
- [Existing authentication and subscription patterns](05-existing-auth-and-subscription-patterns.md) and [Grok Build first technical qualification gate](08-grok-build-technical-gate.md) — local production ownership boundaries, retained smoke evidence, and the separate first-gate outcome.
- Local production sources: [runtime preparation](../../../cli/internal/evaluation/runtime.go), [process adapters](../../../cli/internal/replay/process.go), [Pi RPC adapter](../../../cli/internal/replay/pi.go), and [native setup contract](../../../plans/harness-setup.md).

## Findings

### Native Grok Build authentication has no OpenAI Codex or ChatGPT entitlement route

Grok Build documents browser OIDC, device-code login, an external authentication-provider command, and an `XAI_API_KEY`/per-model xAI API-key route. Its native headless guidance requires the user to run `grok login` or set `XAI_API_KEY`. The external-provider form is a user or enterprise token broker that prints an access token for Grok's configured authentication flow; the documentation does not name OpenAI, Codex, ChatGPT, a Codex OAuth flow, or a ChatGPT plan as a supported provider.

**Evidence:** The [Enterprise Deployments authentication table](https://docs.x.ai/build/enterprise) identifies browser OIDC, device code, external auth provider, and API key, and identifies `api.x.ai` as the direct API-key path. The [headless guide](https://docs.x.ai/build/cli/headless-scripting) fails its example when neither local Grok login nor `XAI_API_KEY` is available. The [CLI reference](https://docs.x.ai/build/cli/reference) names `grok login` and its device-auth option, with no OpenAI/Codex login command or provider selection.

**Implication:** An existing OpenAI subscription alone cannot pass Grok Build's native second gate. A native Grok Build run needs xAI-issued session access or an xAI API key, unless it is deliberately configured for a separate third-party endpoint.

### Third-party BYOK configuration is a separate provider-credential route

Grok Build can define a model with a custom `base_url`, the corresponding wire protocol (`chat_completions`, `responses`, or `messages`), and a named environment variable holding the provider credential. xAI explicitly says third-party BYOK endpoints keep working when their `base_url` is outside `x.ai`. This allows a user-owned configuration to target a compatible OpenAI API endpoint with a valid OpenAI API key and the appropriate protocol, subject to the endpoint's own contract and model availability.

**Evidence:** [Settings](https://docs.x.ai/build/settings) documents `base_url`, `env_key`, and the three `api_backend` variants in a model definition. [Enterprise Deployments](https://docs.x.ai/build/enterprise) describes third-party BYOK endpoints as those whose `base_url` is not on `x.ai`, and directs their restriction to the provider's own IAM.

**Implication:** BYOK is technically relevant to Grok Build, but it is not a subscription bridge. It consumes access supplied by the third-party provider and therefore requires that provider's supported credential and billing/access terms. A user-provisioned OpenAI API key can be a distinct future route; it cannot satisfy this assignment's constraint of using only existing ChatGPT/Codex subscription access.

### A ChatGPT or Codex plan does not supply the OpenAI API credential required by BYOK

OpenAI documents ChatGPT subscriptions and the API platform as separately billed and managed services. The official Codex repository likewise treats ChatGPT-plan sign-in and API-key use as distinct choices: the plan sign-in is for Codex itself, while API-key use requires additional setup. The first-party Codex app-server source makes the ownership boundary explicit: Codex owns its ChatGPT OAuth flow and refresh tokens, whereas API-key mode uses an OpenAI API key for API requests.

**Evidence:** [OpenAI's subscription/API article](https://help.openai.com/en/articles/8156019-is-api-usage-included-in-chatgpt-subscriptions-even-if-i-have-a-paid-chatgpt-account) states that API service is billed and managed separately and uses pay-as-you-go billing. [OpenAI's billing guidance](https://help.openai.com/en/articles/9039756-billing-settings-in-chatgpt-vs-platform) confirms separate billing systems. The [official Codex README](https://github.com/openai/codex) presents ChatGPT-plan sign-in and API-key operation separately, while the [app-server authentication reference](https://github.com/openai/codex/blob/main/codex-rs/app-server/README.md) describes `chatgpt` as Codex-managed OAuth and `apiKey` as an API key saved and used for API requests.

**Implication:** A paid ChatGPT/Codex plan, including its Codex model allowance, is not a documented credential for Grok Build's OpenAI BYOK configuration. Treat any claim that the plan can be forwarded as unsupported unless OpenAI and xAI publish a supported cross-product integration.

### Codex OAuth, personal-access tokens, and credential-extracting proxies are rejected

The official Codex source assigns ownership of ChatGPT OAuth and refresh tokens to Codex. It also describes a personal access token as being loaded through Codex-specific mechanisms. No official xAI or OpenAI source inspected authorizes Grok Build to read these Codex-managed credentials, send them to an OpenAI-compatible endpoint, or adapt them through a proxy. Copying a credential into a Grok configuration, extracting it from Codex state, or inserting a proxy that captures it would violate the required credential and privacy boundary.

**Evidence:** The [Codex app-server authentication reference](https://github.com/openai/codex/blob/main/codex-rs/app-server/README.md) says Codex owns the ChatGPT OAuth flow and token refresh, and lists its own personal-access-token entry points. The Grok Build configuration examples identify an explicit provider key environment variable or per-model key rather than a Codex credential import. No inspected official Grok Build authentication or model-provider source documents an `openai-codex` provider, Codex OAuth, ChatGPT subscription exchange, or proxy integration.

**Implication:** Reject copied or exposed credentials, OAuth extraction, and privacy-invasive proxies. Skill Issue must neither create nor operate an authentication broker, provider setup, persistent Grok model configuration, or proxy for this candidate.

### Existing local indirect routes cannot establish a Grok Build route

Cursor uses its own native account/Keychain route and has no shown mechanism to forward an OpenAI Codex subscription. The retained Claude Code smoke evidence used an explicitly selected, user-provisioned, loopback compatibility launcher backed by a local Codex OAuth proxy; the local report classifies it as conditional and explicitly says it is not official proof or a default pattern for another harness. Pi separately supports the native `openai-codex` provider through the user's untouched Pi agent directory and validates its effective provider/model over RPC. Grok Build documents neither that Pi provider nor the local Claude launcher protocol.

**Evidence:** [Existing authentication and subscription patterns](05-existing-auth-and-subscription-patterns.md) records the native Cursor boundary, the conditional local Claude launcher, and Pi's `openai-codex` route. [Runtime preparation](../../../cli/internal/evaluation/runtime.go) preserves Pi's existing `PI_CODING_AGENT_DIR`; [Pi RPC adapter](../../../cli/internal/replay/pi.go) passes and validates an explicit provider/model; [process adapters](../../../cli/internal/replay/process.go) show that Claude is only a supplied executable override. Grok Build's [Settings](https://docs.x.ai/build/settings) document generic endpoint configuration, not either local route.

**Implication:** The Claude and Pi evidence demonstrates that a user-owned compatibility boundary can be evaluated only when its own harness natively supports it and its effective model can be checked. It does not establish a safe or officially supported way to reuse those components with Grok Build. Treat a purported Grok-to-Codex proxy as unsupported rather than extending a machine-local workaround across products.

### Second-gate outcome: blocked under the requested access constraint

The official evidence supports two distinct Grok Build access classes: native xAI authentication/access and third-party BYOK with that provider's credential. Neither class consumes an existing paid OpenAI Codex/ChatGPT subscription without separate xAI access or a separately billed/provisioned third-party API credential. The existing local compatibility routes are harness-specific and have no official Grok Build integration evidence.

**Evidence:** The native authentication, BYOK, OpenAI billing, Codex authentication, and local-route sources above agree on distinct ownership and credential boundaries. No source provides a documented cross-subscription or Codex-entitlement exchange.

**Implication:** **Campaign blocker — do not advance Grok Build through the Codex-subscription second gate.** The candidate fails the required viability condition. A later, separately scoped evaluation could assess user-owned third-party API-key BYOK only if its payment, credential ownership, temporary configuration, effective-model preflight, and first-gate blockers are independently qualified; it would not be a Codex-subscription route.

## Notes

- No Grok Build installation, authentication, configuration write, model listing, or inference call was performed.
- The official `base_url` surface proves configuration capability, not successful compatibility with a particular OpenAI endpoint or model. No compatibility claim is made without a user-owned credential and a separate non-destructive probe.
- This report keeps Grok Build distinct from Grok Chat, Grok API, xAI IDE integrations, and non-xAI command-line clients.

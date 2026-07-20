# Gemini CLI Codex-Subscription Second Gate

## Assignment

**Goal:** Determine whether an existing paid OpenAI Codex or ChatGPT subscription and its model access can safely drive Google Gemini CLI without Google entitlement, another subscription, or separately funded API credits.

**Scope:** Current official Gemini CLI, Google AI, and OpenAI documentation; primary Gemini CLI source at `acae7124bdd849e554eaa5e090199a0cf08cd782`; and the local qualified access patterns recorded for Codex, Cursor, Claude Code, and Pi.

**Exclusions:** Installing Gemini CLI, authenticating any account, credential inspection or transfer, provider or proxy setup, permanent configuration changes, direct execution, and first-gate technical qualification.

## Sources

- `research/harness-portability-qualification/assignments/05-existing-auth-and-subscription-patterns.md` — source-traced ownership boundaries and the local Codex, Cursor, Claude compatibility-launcher, and Pi routes.
- `research/harness-portability-qualification/assignments/07-gemini-cli-technical-gate.md` — Gemini remains locally unqualified and its private-home design must use an operator-owned supported authentication input.
- [Gemini CLI authentication setup](https://geminicli.com/docs/get-started/authentication/) (updated 2026-04-17) — supported individual, organization, API-key, Vertex AI, and headless authentication methods.
- [Gemini CLI plans](https://geminicli.com/plans/) — Google Account, Google AI Pro/Ultra, Google Developer Program/Code Assist, AI Studio API-key, and Vertex AI entitlement separation.
- [Gemini CLI configuration reference](https://geminicli.com/docs/reference/configuration/) — `GOOGLE_GEMINI_BASE_URL` and `GOOGLE_VERTEX_BASE_URL` override scope and transport restrictions.
- [Gemini CLI content-generator source](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/core/src/core/contentGenerator.ts) and [its tests](https://github.com/google-gemini/gemini-cli/blob/acae7124bdd849e554eaa5e090199a0cf08cd782/packages/core/src/core/contentGenerator.test.ts) — all supported generator auth types, `GATEWAY` selection, API-key handling, and construction of the Google GenAI client.
- [OpenAI billing separation](https://help.openai.com/en/articles/8156019) (updated 2026-07-15) — API billing is separately managed from ChatGPT.
- [Codex CLI and Sign in with ChatGPT](https://help.openai.com/en/articles/11381614-api-codex-cli-and-sign-in-with-chatgpt) — Codex-specific OAuth/login creates local Codex CLI/API credentials; it is not a general provider credential exchange.
- [Upstream Gemini CLI OpenAI-provider request #23385](https://github.com/google-gemini/gemini-cli/issues/23385) and [official-maintainer discussion #1974](https://github.com/google-gemini/gemini-cli/discussions/1974) — upstream has no direct OpenAI-compatible content provider; community proxy/fork routes are distinct from supported native provider access.

## Findings

### Gemini CLI Has No Direct OpenAI Codex Or ChatGPT Authentication Route

Gemini CLI documents Google Account sign-in, Google AI Studio `GEMINI_API_KEY`, and Vertex AI credentials as its model-access routes. Its current source recognizes only Google personal/Code Assist OAuth, Gemini API key, Vertex AI, Cloud Shell/compute ADC, and a Gemini-protocol gateway. It contains no OpenAI API-key, ChatGPT OAuth, Codex OAuth, or `openai-codex` provider selection.

**Evidence:** The authentication guide requires authentication with Google and directs headless use to a Gemini API key or Vertex AI. The plans page ties the available higher quotas to Google AI, Google Developer Program/Code Assist, AI Studio, or Vertex entitlements. At the pinned source commit, `AuthType` has `oauth-personal`, `gemini-api-key`, `vertex-ai`, Cloud Shell/ADC, and `gateway`; `createContentGenerator` creates `GoogleGenAI` or the Google Code Assist generator. No OpenAI content-generator implementation or OpenAI credential branch is present.

**Implication:** An existing paid OpenAI subscription, Codex login, and Codex model allowance do not by themselves satisfy Gemini CLI authentication or model entitlement. The second gate is blocked on a native route.

### Google Login, Code Assist, Gemini API, And Vertex AI Are Separate Entitlements

Google login can use an individual free, Google AI Pro, or Google AI Ultra account; organizational Google/Workspace and Gemini Code Assist users also require a Google Cloud project and the corresponding Google entitlement. AI Studio API-key use has its own free and paid tiers, while Vertex requires a Google Cloud project plus ADC, a service-account key, or a Google Cloud API key. These routes are owned by Google identity, Google Cloud, and Google billing rather than OpenAI.

**Evidence:** The authentication guide's method table separates Google login, AI Studio API key, and Vertex AI, and explicitly calls out the Cloud-project requirement for Workspace, Developer Program, and Code Assist licenses. The plans page separately lists Google AI Pro/Ultra, Google Developer Program/Code Assist, AI Studio pay-as-you-go, and Vertex AI Google Cloud credentials. Google AI billing describes the API free tier and paid billing-account route. OpenAI states that API service is billed and managed separately from ChatGPT.

**Implication:** Gemini Code Assist or Google login may qualify a future Gemini route only when the operator independently owns that Google access and the selected model is available. It cannot be represented as reuse of OpenAI Codex/ChatGPT access. An OpenAI API key is also outside the requested zero-credit condition unless the operator already has separately funded OpenAI API access.

### The Custom Base-URL Gateway Is Gemini-Native, Not An OpenAI-Compatible Provider Contract

`GOOGLE_GEMINI_BASE_URL` selects `GATEWAY` before `GEMINI_API_KEY`; the source then constructs the Google GenAI client and sends the Gemini API protocol to that base URL. The documented setting calls this an override for Gemini API requests, accepts HTTPS or loopback HTTP, and the pinned tests cover base-URL forwarding and an empty `x-goog-api-key` header for a gateway. It does not translate Gemini requests to OpenAI Chat Completions or Responses, expose an OpenAI model selector, or consume ChatGPT/Codex OAuth.

**Evidence:** The configuration reference limits `GOOGLE_GEMINI_BASE_URL` to Gemini API request overrides. In `contentGenerator.ts`, gateway mode retains `vertexai: false`, passes its API key/headers and base URL to `new GoogleGenAI(...)`, and returns `googleGenAI.models`; the tests verify only this Google-client path. The upstream OpenAI-provider request describes the missing translation layer and was closed without an implemented native OpenAI provider; the official maintainer discussion says the project is optimized for Gemini models rather than direct other-LLM support.

**Implication:** A gateway can be compatible only if it accepts Gemini API wire requests and has its own safe, operator-owned authorization to a provider. Pointing the setting at `api.openai.com` is unsupported and protocol-incompatible. A translating gateway would be a separate third-party system, not Gemini CLI support for an OpenAI subscription.

### Indirect Proxy Routes Do Not Meet The Subscription-Reuse Gate

An indirect route would need a user-owned loopback or private Gemini-protocol-to-OpenAI translation gateway. To call an official OpenAI API, that gateway needs separately authorized API credentials and separately managed API billing; a ChatGPT plan is insufficient. A gateway that tries to reuse Codex/ChatGPT OAuth would need to obtain, copy, expose, or impersonate credentials that Gemini CLI does not support, and there is no official Google or OpenAI contract here authorizing that bridge. Community proxy and fork references show experimentation, but they provide neither a native entitlement nor a safe portable integration contract.

**Evidence:** OpenAI's billing guidance separates ChatGPT from API service. The Codex login guidance scopes its OAuth grant and generated keys to Codex CLI/API-account access. Gemini's source has no OpenAI OAuth/API-key branch. Upstream's own OpenAI-provider request and maintainer response classify third-party proxy/fork approaches as outside direct Gemini CLI provider support.

**Implication:** Reject a Skill Issue-owned translator, shared remote proxy, configuration-writing launcher, OAuth-token extraction, credential-file copying, and any proxy that receives account tokens or private prompts without an independently reviewed privacy and ownership contract. A pre-existing user-owned gateway may be an operator prerequisite for some other qualification decision, but it cannot satisfy this campaign's claim that Gemini CLI reuses an existing paid Codex/ChatGPT subscription without another provider or credits.

### Existing Local Patterns Do Not Supply A Gemini Exception

The qualified Codex route preserves Codex's own ChatGPT login; Pi directly selects its native `openai-codex` provider and keeps the caller-owned Pi auth directory; both are native provider routes. Cursor uses its own account/Keychain entitlement. The local Claude route is explicitly a user-provisioned, loopback-only compatibility launcher backed by a Codex OAuth proxy, selected as an executable override and treated as local smoke evidence rather than a general official integration.

**Evidence:** Assignment 05 traces each concrete adapter and states that the Claude proxy is conditionally acceptable only when the operator already owns and explicitly selects it, while Skill Issue neither creates nor manages its login, storage, mapping, or process state. Assignment 07 requires Gemini authentication to be an operator-provided supported credential input when running from private state.

**Implication:** Gemini has no equivalent native `openai-codex` provider and no existing project-local Gemini compatibility launcher. The Claude precedent cannot be generalized: creating or adapting a new Gemini translator would violate the campaign's user-owned setup, privacy, and no-permanent-mutation constraints.

### Second-Gate Decision: Blocked

No safe, source-backed route lets an existing paid OpenAI Codex/ChatGPT subscription directly power Gemini CLI without an independent Google entitlement, separately funded OpenAI API access, or a separately owned and authorized gateway. Gemini CLI's custom base URL is useful only for a Gemini-protocol gateway and does not close the missing OpenAI provider/authentication gap.

**Evidence:** The combined official Google authentication/plans documentation, pinned Gemini source, and official OpenAI billing boundary establish distinct credential and payment owners. The local compatibility evidence establishes a narrow user-provisioned exception for Claude Code, but no analogous Gemini route.

**Implication:** **Campaign blocker:** exclude Gemini CLI from any qualification cell whose required access source is the existing paid OpenAI Codex/ChatGPT subscription alone. Reopen only if Google ships a supported OpenAI/Codex provider contract, or if the campaign scope explicitly permits an independently owned Google entitlement or a separately reviewed user-owned gateway with its own provider authorization and live qualification evidence.

## Notes

- This conclusion is limited to the stated no-Google/no-other-subscription/no-API-credit condition. It does not assess the quality, cost, or technical first-gate readiness of a future Google-authenticated Gemini run.
- No Gemini installation, account sign-in, API request, credential read, local configuration edit, or proxy setup was performed.
- The current official Gemini docs and upstream source are the authority for direct support. The upstream issue/discussion are retained only to classify indirect community approaches as unsupported by the native provider contract.

# Kilo Code CLI Codex-Subscription Gate

## Assignment

**Goal:** Determine whether an existing paid OpenAI Codex/ChatGPT subscription can drive Kilo Code CLI without a separate OpenAI API key, OpenAI API credits, or Kilo-paid inference.

**Scope:** Current first-party Kilo CLI and provider documentation, Kilo-Org/kilocode source at commit `0cf9f902e0eaead2746a935e9440d36b84df45bf`, and official OpenAI billing guidance. This covers native CLI authentication, provider and custom-endpoint paths, credential ownership, model boundaries, and safe comparison with the existing Pi and Claude routes.

**Exclusions:** Installing or authenticating Kilo; executing a model prompt; changing Kilo or OpenAI configuration; Kilo Gateway/cloud-feature qualification; credential contents; a production Skill Issue adapter; and proxy implementation.

## Sources

- [Kilo: Using ChatGPT Subscriptions With Kilo Code](https://kilo.ai/docs/ai-providers/openai-chatgpt-plus-pro) — first-party claim for ChatGPT subscription access in the CLI, OAuth, billing, plan, model, and cloud-feature boundaries; inspected 2026-07-21.
- [Kilo CLI command reference](https://kilo.ai/docs/code-with-ai/platforms/cli-reference) — `kilo auth list`, `login`, and `logout`, plus explicit provider/method flags; inspected 2026-07-21.
- [Kilo CLI documentation](https://kilo.ai/docs/code-with-ai/platforms/cli) — CLI credential-management and global configuration locations; inspected 2026-07-21.
- [Kilo BYOK documentation](https://kilo.ai/docs/getting-started/byok) and [Kilo Gateway documentation](https://kilo.ai/docs/gateway) — Gateway/API-key and provider-billing boundaries; inspected 2026-07-21.
- [Kilo custom-model documentation](https://kilo.ai/docs/code-with-ai/agents/custom-models) — CLI-applicable custom-provider, OpenAI-compatible endpoint, API-key/header, and trusted-config behavior; inspected 2026-07-21.
- [Kilo-Org/kilocode at inspected commit](https://github.com/Kilo-Org/kilocode/commit/0cf9f902e0eaead2746a935e9440d36b84df45bf), [Codex OAuth plugin](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/plugin/openai/codex.ts), [provider-login command](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/cli/cmd/providers.ts), and [auth storage](https://github.com/Kilo-Org/kilocode/blob/0cf9f902e0eaead2746a935e9440d36b84df45bf/packages/opencode/src/auth/index.ts) — first-party implementation of the CLI OpenAI/Codex OAuth route, account-scoped refresh, provider selection, and persistent/process-injected credentials.
- [OpenAI Help: API service versus ChatGPT subscription](https://help.openai.com/en/articles/8156019-is-api-usage-included-in-chatgpt-subscriptions-even-if-i-have-a-paid-chatgpt-account) — official separate-billing boundary; inspected 2026-07-21.
- Local research context: [existing access patterns](05-existing-auth-and-subscription-patterns.md) and [Kilo technical gate](10-kilo-code-technical-gate.md).

## Findings

### Native Kilo CLI OAuth can use an existing eligible ChatGPT/Codex subscription

**Outcome: conditionally passes the subscription-access part of the second gate.** Kilo documents that an existing ChatGPT Plus or Pro subscription can run Codex models in Kilo Code's core functionality, explicitly including the CLI, through OAuth with no API key and no separately billed API usage. Its plan notes also include Business, Edu, and Enterprise where Codex is included; the account must actually retain the requested Codex/model entitlement. This is a direct Kilo provider route, so it is materially different from forwarding an OpenAI API key or placing a proxy between Kilo and OpenAI.

**Evidence:** Kilo's ChatGPT-subscription page says that ChatGPT subscription usage works in the VS Code extension **and CLI**, uses OAuth rather than an API key, and counts against subscription limits rather than API billing. The same page limits the route to Kilo's Codex catalogue and says free ChatGPT accounts are unsupported. In the inspected Kilo CLI source, the `openai` provider has two first-party OAuth methods labelled `ChatGPT Pro/Plus (browser)` and `ChatGPT Pro/Plus (headless)`; the plugin exchanges/refreshes OAuth tokens and sends them to OpenAI's Codex responses endpoint with the ChatGPT account ID. The CLI's provider login code labels OpenAI as `ChatGPT login or API key` and accepts `codex` as an alias for the OpenAI ChatGPT-auth flow.

**Implication:** A user who already has an eligible ChatGPT/Codex plan can independently authorize Kilo's native OpenAI provider, then select a supported `openai/<Codex-model>` model for `kilo run`; no separate OpenAI Platform key, OpenAI API credits, Kilo credits, or Kilo Pass is required for that model traffic. Availability must remain a preflighted account-and-model fact, not an inferred entitlement from the presence of a subscription.

### API, BYOK, Kilo Gateway, and custom endpoints are separate access routes

Kilo's direct OpenAI API provider, Kilo Gateway BYOK, and generic OpenAI-compatible provider all rely on API-key or configured-header access. They do not turn ChatGPT/Codex subscription access into an API credential. A Kilo Gateway account/API key is itself a separate Kilo access path; Gateway BYOK routes with a provider API key and provider billing. OpenAI states that API service billing is managed separately from ChatGPT, so an OpenAI Platform key is a separate paid-access arrangement even when its owner also has ChatGPT Plus or Pro.

**Evidence:** Kilo's OpenAI provider documentation tells API users to create an OpenAI Platform API key, while separately directing Plus/Pro subscribers to the ChatGPT provider. Kilo BYOK lists OpenAI among standard **API-key** providers and routes matching requests using that key; its Gateway documentation requires a Kilo API key. Kilo's custom-provider documentation describes OpenAI-compatible endpoints with `apiKey`, optional custom headers, and `baseURL`; it describes compatibility with the protocol, rather than a ChatGPT OAuth bridge. OpenAI's billing help states that API service is billed and managed separately from ChatGPT.

**Implication:** The following do **not** meet the requested no-additional-access condition: an `OPENAI_API_KEY`/OpenAI Platform setup, a Kilo Gateway token funded by Kilo credits or Kilo Pass, or an OpenAI BYOK key. They may be valid provider routes for an operator who chooses their separate billing, but must be recorded as API/provider access rather than Codex-subscription access.

### The supported custom-endpoint surface does not justify a Codex proxy route

Kilo CLI can configure an OpenAI-compatible endpoint, including its base URL, API key or headers, and model metadata. That makes a user-operated endpoint technically configurable, but neither Kilo nor OpenAI's inspected documentation establishes a safe, supported proxy that exposes a ChatGPT/Codex subscription to Kilo through that generic surface. A proxy would also add a new credential holder and prompt-transit boundary when native Kilo OAuth already exists.

**Evidence:** Kilo's custom-model page expressly applies to the CLI and shows `openai-compatible` configuration using an environment-provided provider API key and an arbitrary endpoint. The same page reserves environment substitution for trusted config locations to prevent malicious project configuration from exfiltrating secrets. The Kilo native OAuth plugin instead obtains its own OAuth authorization, uses the Codex endpoint directly, refreshes its own token, and injects the current account header. No inspected first-party Kilo or OpenAI source documents a generic proxy as a supported way to convey Codex subscription entitlement.

**Implication:** Reject copied or extracted Codex credentials, OAuth-token export, a hosted/privacy-invasive proxy, or a Skill Issue-created proxy/endpoint. A user-owned third-party compatibility launcher could only be treated like the existing conditional Claude launcher if the user explicitly provisions, owns, selects, and audits it outside Skill Issue; it is lower fit and unnecessary for Kilo because native OAuth is available. It cannot establish Kilo's standard second-gate route.

### Credentials must stay Kilo-user-owned; process injection is not an evaluator mechanism

Kilo's normal CLI auth store is its data-root `auth.json`; Kilo writes it with mode `0600`. The native OAuth handler refreshes expired tokens and persists the refreshed OpenAI credential. The source also accepts `KILO_AUTH_CONTENT`, which overrides stored auth for the process, but that variable contains serialized credential material and a refresh can write an auth file. It is therefore an implementation capability, not an acceptable Skill Issue credential-transfer design.

**Evidence:** The Kilo auth source resolves `auth.json` below Kilo's global data root, reads `KILO_AUTH_CONTENT` before that file, and writes auth records with `0600` permissions. The Codex plugin's refresh handler calls Kilo's auth-set API for provider `openai`. Kilo's CLI docs identify `kilo auth` as credential management and global Kilo configuration as user-owned state.

**Implication:** The acceptable pattern is user-provisioned native Kilo OpenAI OAuth in the user's Kilo state, with Skill Issue only selecting the executable/model and checking non-secret availability. Skill Issue must neither create a Kilo login, read/copy/print/supply `auth.json` or `KILO_AUTH_CONTENT`, extract OAuth material from Codex, mutate normal Kilo configuration, nor log the user out. The private-XDG isolation and cancellation issues in the first technical gate remain separate blockers: this finding proves subscription access, not a safe current evaluator runtime.

### Native Kilo OAuth aligns more closely with Pi than with the conditional Claude launcher

Pi's qualified `openai-codex` route preserves its existing user-owned Pi agent directory and native login; the Kilo equivalent is the first-party `openai` OAuth provider persisted in Kilo's own user state. The existing Claude route instead depends on an explicitly selected, user-owned, loopback compatibility launcher/proxy that maps another harness's interface to the Codex subscription. Kilo's native OAuth does not need that compatibility layer.

**Evidence:** Assignment 05 traces Pi's native `openai-codex` provider and retained agent directory, and identifies the Claude route as a local, user-owned compatibility launcher backed by a loopback proxy. Kilo's official docs say ChatGPT subscription access applies to its CLI and source implements the OAuth flow in the CLI's own OpenAI provider rather than through an OpenAI-compatible custom endpoint.

**Implication:** Treat Kilo's direct OAuth as the preferred second-gate access candidate, conditional on user-owned native Kilo state and effective model availability. Do not reuse the Claude proxy as a default, silently convert Pi/Codex credentials for Kilo, or treat the three credential stores as interchangeable.

## Notes

- **Current overall status:** Kilo remains technically unqualified under assignment 10 because no local executable/authenticated run exists and its isolation, daemon, cancellation, and cleanup requirements are unresolved. The supported subscription route removes the earlier access/billing blocker only; it does not qualify Kilo for Skill Issue evaluation.
- **Model caveat:** Kilo's ChatGPT OAuth provider exposes Kilo's Codex catalogue, not every OpenAI API model. Kilo source filters the OAuth model catalogue, so the requested model and effective resolved model need observable preflight evidence in a future user-authorized probe.
- **Credential-storage caveat:** Kilo's public ChatGPT page describes IDE SecretStorage for exported settings, while the inspected CLI source stores CLI credentials in its data-root `auth.json`. This report relies on the CLI source for the CLI storage claim.
- **Validation performed:** read-only inspection of first-party Kilo documentation and the specified Kilo-Org commit, plus official OpenAI billing guidance; no Kilo executable was installed, no login occurred, no configuration changed, and no model request was sent.

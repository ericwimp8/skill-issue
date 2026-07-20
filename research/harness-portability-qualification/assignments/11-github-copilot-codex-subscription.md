# GitHub Copilot CLI Codex-Subscription Second Gate

## Assignment

**Goal:** Determine whether a user's existing paid OpenAI Codex/ChatGPT subscription and its model access can qualify GitHub Copilot CLI without a GitHub Copilot entitlement or separately purchased OpenAI API access.

**Scope:** Current official GitHub, Microsoft, and OpenAI documentation; GitHub's primary Copilot CLI release materials; and the existing Skill Issue authentication, local-launcher, and first-technical-gate findings. The assessment covers native Copilot entitlement, official local and enterprise BYOK, OpenAI-compatible endpoints, supported credential forms, and safe indirect routes.

**Exclusions:** Installing or authenticating Copilot; inspecting, copying, extracting, or exposing any credential; creating or changing user, organization, proxy, or provider configuration; Skill Issue-owned provider setup; and claims that documentation-only evidence qualifies the remaining technical, activation, independent-agent, or cleanup gates.

## Sources

- [Existing authentication and subscription patterns](05-existing-auth-and-subscription-patterns.md) — locally observed, bounded Codex, Cursor, Claude compatibility-launcher, and Pi ownership patterns.
- [GitHub Copilot CLI first technical qualification gate](06-github-copilot-technical-gate.md) — current local Copilot absence and remaining non-authentication qualification gaps.
- [Skill Issue native setup contract](../../../plans/harness-setup.md) and [Pi adapter](../../../cli/internal/replay/pi.go) — production ownership boundary: preserve a native provider's existing authentication only; do not make Skill Issue a credential or provider owner.
- GitHub Docs, [Authenticating GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/set-up-copilot-cli/authenticate-copilot-cli), inspected 2026-07-21 — native GitHub authentication, BYOK exception, credential types, offline behavior, and request routing.
- GitHub Docs, [Using your own LLM models in GitHub Copilot CLI](https://docs.github.com/en/copilot/how-tos/copilot-cli/customize-copilot/use-byok-models), inspected 2026-07-21 — supported BYOK providers, required environment variables, OpenAI example, and model capability requirements.
- GitHub Docs, [Bring your own key for GitHub Copilot](https://docs.github.com/en/copilot/concepts/models/bring-your-own-key), inspected 2026-07-21 — local versus enterprise BYOK entitlement boundaries.
- GitHub changelog, [Copilot CLI now supports BYOK and local models](https://github.blog/changelog/2026-04-07-copilot-cli-now-supports-byok-and-local-models/), published 2026-04-07 — primary GitHub release statement that BYOK bypasses GitHub-hosted routing and invalid provider configuration does not fall back silently.
- OpenAI Help Center, [How can I move my ChatGPT subscription to the API?](https://help.openai.com/en/articles/8156019-is-api-usage-included-in-chatgpt-subscriptions-even-if-i-have-a-paid-chatgpt-account), inspected 2026-07-21 — ChatGPT and API billing separation.
- OpenAI Help Center, [Using Codex with your ChatGPT plan](https://help.openai.com/en/articles/11369540-using-codex-with-your-chatgpt-plan.pdf), inspected 2026-07-21 — supported ChatGPT-plan Codex clients and their ChatGPT sign-in flow.
- OpenAI Help Center, [Codex CLI and Sign in with ChatGPT](https://help.openai.com/en/articles/11381614-api-codex-cli-and-sign-in-with-chatgpt), inspected 2026-07-21 — Codex CLI's OpenAI-specific OAuth flow and generated API-key distinction.
- Microsoft Learn, [Plan and manage costs in Microsoft Foundry](https://learn.microsoft.com/en-us/azure/foundry/concepts/manage-costs), inspected 2026-07-21 — Azure OpenAI/Foundry metering and billing ownership.

## Findings

### GitHub-hosted Copilot and OpenAI Codex subscription are separate entitlements

For ordinary GitHub-hosted Copilot CLI requests, GitHub requires GitHub authentication and makes the CLI available through Copilot plans. Its supported credentials are GitHub OAuth, a GitHub fine-grained PAT with the Copilot Requests permission, or a GitHub App user-to-server token. An OpenAI ChatGPT/Codex subscription credential is absent from that contract. A GitHub Copilot Free plan can supply limited GitHub-hosted use, but it remains a GitHub entitlement rather than a bridge for the existing OpenAI plan.

**Evidence:** GitHub's authentication reference says authentication is required for all non-BYOK Copilot CLI use, identifies the three GitHub credential forms, and states CLI availability with Copilot plans. OpenAI's Codex-plan guide identifies supported Codex clients as the Codex CLI, IDE extension, web, and desktop app; it does not list GitHub Copilot CLI.

**Implication:** An existing paid OpenAI plan neither authenticates GitHub Copilot CLI nor grants a GitHub Copilot model allowance. Native Copilot may be evaluated only with a user-owned GitHub Copilot entitlement/account policy, separately from the Codex subscription question.

### Official Copilot CLI BYOK can remove the GitHub subscription requirement, but it requires a provider credential

GitHub now officially supports local BYOK for Copilot CLI. With a provider base URL and model, the CLI can send model requests directly to OpenAI, Azure OpenAI, Anthropic, or an OpenAI Chat Completions-compatible endpoint. GitHub says GitHub authentication is not required in this mode and that local BYOK is usable without a Copilot subscription. Its OpenAI configuration example requires `COPILOT_PROVIDER_API_KEY` containing an OpenAI API key; models must support streaming and tool calling. The release material says an invalid provider configuration fails rather than silently routing through GitHub-hosted Copilot.

**Evidence:** GitHub's BYOK guide specifies `COPILOT_PROVIDER_BASE_URL`, `COPILOT_PROVIDER_API_KEY`, and `COPILOT_MODEL`, and gives `https://api.openai.com/v1` plus an OpenAI API key as the remote OpenAI route. The GitHub authentication reference confirms that these provider variables take over AI-model requests regardless of GitHub login. The GitHub BYOK overview says local BYOK removes dependence on GitHub's Copilot API.

**Implication:** BYOK is a technically legitimate no-Copilot-subscription route, but it is not a Codex/ChatGPT-subscription route. It can qualify only if the user already owns a compatible provider credential and authorizes its process-scoped use; Skill Issue cannot create, collect, persist, or configure that provider access.

### A ChatGPT/Codex plan does not provide the OpenAI API credential and billing that Copilot BYOK requires

OpenAI bills and manages API service separately from ChatGPT. ChatGPT-plan Codex access uses OpenAI's own supported sign-in flow in Codex clients. The older Codex CLI flow may link a ChatGPT identity to an API organization and generate a distinct API key; its documentation explicitly distinguishes the ChatGPT session token from CLI-generated secret keys, and any promotional API credit is conditional, time-limited, and separate from ongoing subscription entitlement.

**Evidence:** OpenAI states that API service is billed and managed separately from ChatGPT and is token-priced. The Codex-plan guide directs users to sign in with ChatGPT through a supported Codex client. The Codex CLI sign-in guide states that the ChatGPT session token and CLI-generated secret keys are separate, and that API keys can remain active after an OAuth disconnect.

**Implication:** Passing a ChatGPT cookie, Codex OAuth refresh token, generated Codex CLI secret, or extracted local authentication material to `COPILOT_PROVIDER_API_KEY` has no official support and is rejected. With neither an existing funded OpenAI API account/key nor a non-OpenAI provider, the official OpenAI BYOK route is blocked; a paid ChatGPT/Codex plan alone is insufficient.

### Safe compatible endpoints remain separate provider routes, not a subscription conversion

The only supported compatible-endpoint path is the documented BYOK provider interface. It allows a user-owned OpenAI-compatible endpoint, including local Ollama, vLLM, or Foundry Local, but the endpoint must supply the required streaming and tool-calling behavior. A remote OpenAI endpoint needs an API key; an Azure OpenAI endpoint needs independently owned Azure access, and Microsoft documents its own token-based or deployed-resource billing. Local models can avoid both GitHub and OpenAI API charges, but they do not reproduce the user's existing OpenAI Codex subscription/model access.

**Evidence:** GitHub lists OpenAI-compatible endpoints and locally running providers under BYOK, with tool-calling and streaming requirements. The OpenAI example uses a provider API key. Microsoft documents that Foundry and Azure OpenAI resources accrue costs according to service usage and meters.

**Implication:** A user may explicitly select a pre-existing, policy-approved local or remotely billed provider for a future Copilot BYOK probe. That is a new provider route with its own model, data handling, and billing terms. It cannot satisfy a request to reuse an existing paid OpenAI Codex/ChatGPT subscription without API credits.

### The established local Codex routes do not transfer to Copilot CLI

Skill Issue's qualified Codex route preserves the normal Codex home only for Codex's own supported login. Pi preserves its native `openai-codex` agent directory and validates the reported provider/model; it does not expose that login as an OpenAI-compatible service. Cursor accesses its own account through the macOS Keychain bridge and has no Codex-subscription bridge. The local Claude executable can use a user-provisioned loopback Codex OAuth compatibility proxy, but the existing report classifies it as a local, conditional compatibility launcher rather than official Claude or OpenAI support.

**Evidence:** The local authentication report records distinct native ownership for Codex, Pi, Cursor, and the Claude launcher. The production setup contract preserves existing Codex/Pi authentication but prohibits credential copying and ordinary user-configuration mutation. GitHub's published BYOK contract names provider API-key configuration; it does not name Codex OAuth or the local Claude compatibility proxy as supported provider authentication.

**Implication:** None of these patterns authorizes Skill Issue to repurpose a Keychain item, Pi auth directory, Codex login, or Claude proxy for Copilot. A user-owned loopback proxy could only be considered as an explicitly selected external BYOK provider after independent proof of protocol, privacy, authentication, lifecycle, and effective-model behavior; it is unsupported for this second gate and must not be installed, configured, or managed by Skill Issue.

### Enterprise BYOK also does not solve the personal-subscription question

GitHub's enterprise custom-model route lets an enterprise or organization owner add an API key and selected models centrally, but it requires a policy-enabled owner workflow and a Copilot license for users. It creates shared provider configuration outside the evaluation run and changes organization-owned state.

**Evidence:** GitHub's BYOK overview states that enterprise custom models require a Copilot license. GitHub's organization custom-model guide requires an enterprise policy and owner-supplied provider API key.

**Implication:** Enterprise BYOK is outside the allowed second-gate route: it neither reuses a personal Codex/ChatGPT subscription nor preserves the required no-mutation, user-owned provider boundary.

### Second-gate result: blocked for Codex-subscription-only access

No official direct, BYOK, or safe established indirect route allows GitHub Copilot CLI to consume a user's existing paid OpenAI Codex/ChatGPT subscription/model access without either a GitHub Copilot entitlement or separately owned compatible-provider access. The present Copilot target also remains unqualified at the first technical gate because no local executable or live version evidence exists.

**Evidence:** GitHub's native authentication contract is GitHub-only outside BYOK; BYOK requires a provider key or a local model; and OpenAI documents ChatGPT/API billing and credentials as separate. The first technical-gate report records no installed `copilot` executable and unresolved live qualification requirements.

**Implication:** Mark the Copilot Codex-subscription second gate **blocked**. A later Copilot BYOK qualification may proceed only under a separately supplied, user-owned provider route (for example, an already funded OpenAI API key, approved Azure resource, or local model), with explicit provider/effective-model evidence and no Skill Issue credential, proxy, or permanent-configuration ownership. It does not meet the requested no-Copilot-and-no-API-credits condition through the existing paid Codex/ChatGPT subscription alone.

## Notes

- This conclusion distinguishes three independent resources: GitHub Copilot entitlement, OpenAI API quota/credits, and ChatGPT/Codex subscription access. A shared account identity does not collapse those resource boundaries.
- Documentation supports Copilot BYOK as of the inspected current release, but it does not establish that a particular installed version passes the separate Skill Issue isolation, JSON traceability, permission, cancellation, state-retention, activation, or independent-agent gates.
- A loopback endpoint is not automatically privacy-safe. Any future external proxy route must be operator-owned, explicitly selected, process-scoped, minimally credentialed, transparent about where prompts/code go, and independently validated; otherwise it remains unsupported.

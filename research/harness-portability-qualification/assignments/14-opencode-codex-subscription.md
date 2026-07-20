# OpenCode Codex-Subscription Viability

## Assignment

**Goal:** Determine whether an existing paid OpenAI Codex/ChatGPT subscription and its eligible model access can drive OpenCode without separate API credits or another subscription, while keeping credentials and persistent provider state user-owned.

**Scope:** Current official OpenCode and OpenAI material, current primary OpenCode source, the locally installed executable's non-mutating help surface, and the existing Skill Issue authentication and OpenCode technical gates.

**Exclusions:** Authentication, account inspection, model calls, provider configuration changes, credential contents, proxy deployment, Skill Issue implementation, and technical-gate qualification beyond the credential/provider boundary.

## Sources

- [Existing authentication and subscription patterns](05-existing-auth-and-subscription-patterns.md) — the Codex, Pi, and conditional Claude compatibility ownership boundaries used for comparison.
- [OpenCode technical qualification gate](09-opencode-technical-gate.md) — current OpenCode isolation, local-version, model/variant, and cleanup limitations.
- [OpenCode provider documentation](https://opencode.ai/docs/providers) — current OpenAI setup directs a user to select `ChatGPT Plus/Pro`, complete browser authentication, and then select available models; the same document defines custom providers, `baseURL`, API keys, and headers.
- [OpenCode CLI documentation](https://opencode.ai/docs/cli) and [configuration documentation](https://opencode.ai/docs/config) — documented `auth login`, `auth list`, credentials path, configuration precedence, and `OPENCODE_CONFIG_DIR` scope.
- [OpenCode model documentation](https://opencode.ai/docs/models) — `provider_id/model_id` selection and provider-specific model options.
- Primary OpenCode source at commit [`4872c48`](https://github.com/anomalyco/opencode/tree/4872c48c230728150e8e3406722943450ed58dcb): [built-in OpenAI/Codex plugin](https://github.com/anomalyco/opencode/blob/4872c48c230728150e8e3406722943450ed58dcb/packages/opencode/src/plugin/openai/codex.ts), [auth store](https://github.com/anomalyco/opencode/blob/4872c48c230728150e8e3406722943450ed58dcb/packages/opencode/src/auth/index.ts), [global path ownership](https://github.com/anomalyco/opencode/blob/4872c48c230728150e8e3406722943450ed58dcb/packages/core/src/global.ts), [configuration-directory discovery](https://github.com/anomalyco/opencode/blob/4872c48c230728150e8e3406722943450ed58dcb/packages/opencode/src/config/paths.ts), and [plugin registration](https://github.com/anomalyco/opencode/blob/4872c48c230728150e8e3406722943450ed58dcb/packages/opencode/src/plugin/index.ts). Retrieved 2026-07-21.
- [Using Codex with your ChatGPT plan](https://help.openai.com/en/articles/11369540-using-codex-with-your-chatgpt-plan.pdf) and [ChatGPT/API billing separation](https://help.openai.com/en/articles/8156019-is-api-usage-included-in-chatgpt-subscriptions-even-if-i-have-a-paid-chatgpt-account) — Codex subscription access and the separate API-billing contract.
- Local observation, 2026-07-21: `/opt/homebrew/bin/opencode` reports `1.14.39`; its read-only `auth --help`, `auth login --help`, and `run --help` expose provider login with `--provider`/`--method`, `provider/model`, and `--variant`. No auth state, credential file, configuration, or model call was inspected or changed.

## Findings

### OpenCode has a native ChatGPT subscription route, not a Pi-style `openai-codex` provider

OpenCode's current official provider setup explicitly offers `ChatGPT Plus/Pro` browser authentication under the `OpenAI` provider. Its built-in `CodexAuthPlugin` is registered internally for provider ID `openai`; it has browser and headless device-code methods labelled `ChatGPT Pro/Plus`, performs OAuth with `auth.openai.com`, and reroutes OpenAI Responses or Chat Completions requests to `https://chatgpt.com/backend-api/codex/responses`. This is a direct native provider path, rather than a `openai-codex` provider directory such as Pi uses.

**Evidence:** OpenCode's provider documentation identifies the ChatGPT Plus/Pro option and browser authentication. The pinned plugin source registers `provider: "openai"`, defines the two ChatGPT OAuth methods, and selects the ChatGPT Codex responses endpoint. OpenAI's Codex-plan documentation says Codex is included with ChatGPT Plus and Pro, while model availability remains entitlement-dependent.

**Implication:** **Conditionally viable at the subscription gate.** An eligible user can authorize OpenCode with the same paid ChatGPT account and use OpenCode's native subscription route without API credits or another provider subscription. The user must complete OpenCode's own native authorization; having only an existing Codex CLI login is not source-backed as a reusable OpenCode credential source.

### Native OpenCode OAuth is user-owned persistent state and refreshes in place

OpenCode stores provider credentials in its own `auth.json` beneath the OpenCode data root; its documentation identifies the ordinary path as `~/.local/share/opencode/auth.json`. The primary source stores OAuth refresh token, access token, expiry, and optional ChatGPT account ID under provider key `openai`, writes that file with mode `0600`, and refreshes expired OAuth tokens by persisting the replacement tokens through the same auth service. The CLI documents `auth list` as inspecting that credential store and `auth logout` as clearing it.

**Evidence:** OpenCode CLI/provider documentation specifies the credentials file and commands. The pinned auth source joins `Global.Path.data` with `auth.json`; its auth schema contains `refresh`, `access`, `expires`, and `accountId`; the plugin refresh branch calls the client auth setter. Global-path source derives the data root from XDG data conventions.

**Implication:** The correct ownership pattern is user-operated OpenCode login before a Skill Issue run, with OpenCode retaining and refreshing its own normal credential state. Skill Issue may preflight availability without reading or reporting tokens, but must not authenticate, invoke a login flow, write `auth.json`, or log the user out. This matches Pi's principle of preserving user-owned native authentication, though the native store and provider identifier differ.

### A complete temporary OpenCode root cannot safely reuse the normal OAuth store without credential transfer or normal-state mutation

`OPENCODE_CONFIG_DIR` only adds a directory for OpenCode configuration assets after global and project discovery; it does not redirect the auth store. Source derives auth from `Global.Path.data`, which follows the XDG data root, while source-backed `OPENCODE_CONFIG_DIR` affects the configuration root only. Moving `HOME`/XDG data to a temporary root therefore gives OpenCode an empty auth location. `OPENCODE_AUTH_CONTENT` overrides auth by placing the full credential object in the process environment, which is credential copying/exposure. A symlink from a temporary data root to the normal `auth.json` avoids a byte-for-byte copy but still allows token refresh to write through to the normal user file and places that credential-bearing path inside run-owned state.

**Evidence:** The pinned configuration-path source lists `OPENCODE_CONFIG_DIR` separately from global data paths; the global-path source bases data on `xdgData`; the auth source accepts `OPENCODE_AUTH_CONTENT` before reading the auth file and writes refreshed state through its auth service. The official configuration document likewise scopes `OPENCODE_CONFIG_DIR` to agents, commands, modes, and plugins, and describes merge rather than replacement precedence.

**Implication:** A temporary CLI-owned **configuration** directory can safely coexist with untouched normal OpenCode authentication only by leaving the user-owned OpenCode data root in place. It cannot be treated as a complete isolated runtime: global/project/managed configuration and the normal credential store remain live, and token refresh may update normal OpenCode auth. A fully temporary data root has no safe native-auth bridge under the stated constraints. Reject copied auth files, `OPENCODE_AUTH_CONTENT`, credential environment forwarding, and symlink bridges as a Skill Issue route.

### Subscription models use `openai/<model>` and require OpenCode-native availability checks

OpenCode's model contract uses `provider_id/model_id`, so the direct route selects `openai/<model-id>` with `opencode run --model`; `--variant` is provider/model-specific rather than a universal Skill Issue reasoning enum. The native OAuth plugin filters the OpenAI registry: it excludes Pro reasoning modes, excludes `gpt-5.5-pro` and bare `gpt-5.6`, allows listed Codex-compatible models, and otherwise accepts versioned GPT models newer than 5.4 subject to that filter. The provider registry and entitlement are dynamic, so neither a Codex CLI model name nor ChatGPT web availability proves that a particular OpenCode model/variant is selectable.

**Evidence:** OpenCode model documentation defines provider/model IDs and provider options; local `run --help` exposes `--model provider/model` and provider-specific `--variant`. The pinned plugin source contains its OAuth-only model filter and direct OpenAI provider ID. The existing technical gate records no validated cross-provider variant mapping or authenticated OpenCode model probe.

**Implication:** A later qualification must preflight the specific `openai/<model>` and variant through user-owned native OpenCode state, record requested and effective model, and fail closed when it is absent. Do not map Pi's `openai-codex/<model>` identifiers or the Claude compatibility launcher's aliases into OpenCode.

### API keys and compatible-endpoint routes do not meet the no-extra-credits subscription route

OpenCode supports a manually entered OpenAI API key, custom providers using OpenAI-compatible Chat Completions or Responses endpoints, custom `baseURL`, and optional custom headers. Those routes require a provider API key or other gateway credential. OpenAI states that ChatGPT and API billing are separate, so an OpenAI API-key route requires API billing/credits even if the user already pays for ChatGPT. A third-party compatible endpoint or proxy similarly introduces a separately owned credential, billing/account boundary, and potentially its own privacy/data handling.

**Evidence:** OpenCode's provider documentation describes manual API-key setup, custom provider `apiKey`, `baseURL`, and headers; its custom-provider instructions distinguish OpenAI-compatible endpoints from the direct provider setup. OpenAI's billing guidance explicitly separates ChatGPT subscriptions from API service billing.

**Implication:** Reject API-key, BYOM/custom-provider, and proxy routes for this second-gate purpose. They may be legitimate user-operated OpenCode configurations, but they do not establish use of an existing paid Codex/ChatGPT subscription without separate API credits or another subscription.

### The conditional Claude compatibility launcher is a lower-fit comparison, not an OpenCode dependency

The existing Claude route was a user-provisioned launcher around a local Codex OAuth proxy, with explicit selected executable, loopback process ownership, and requested-to-effective model reporting. OpenCode's built-in direct OAuth route avoids that proxy and preserves provider ownership in OpenCode's normal auth store. Its subscription viability is therefore stronger than a compatibility launcher, but it still cannot satisfy the separate first technical gate's clean-runtime and managed/ambient-configuration requirements.

**Evidence:** Existing authentication patterns classify Claude's proxy launcher as conditional and user-owned. OpenCode's official provider setup and pinned built-in plugin implement the direct OpenAI OAuth path. The existing OpenCode technical gate reports `OPENCODE_CONFIG_DIR` is additive and current clean isolation is unqualified.

**Implication:** A future OpenCode qualification should prefer the native OAuth route and explicitly require pre-existing user-owned OpenCode login. It must not create a Skill Issue-owned proxy, extract Codex credentials, or mutate a user's normal OpenCode setup to approximate direct access.

## Notes

- The installed `1.14.39` CLI exposes generic provider-login flags, but this assignment did not open the provider picker or authenticate; its direct ChatGPT method availability is therefore unsupported for that installed release. The native route is verified against current official documentation and source commit `4872c48`, not locally exercised.
- Official OpenCode wording names ChatGPT Plus/Pro. Current source and this assignment do not establish the same OpenCode login method for every Codex-capable ChatGPT Business, Enterprise, Edu, Free, or Go entitlement.
- A user may deliberately maintain their normal OpenCode OAuth login outside Skill Issue. That normal credential state is a prerequisite, not run-owned setup, and may be affected by ordinary native token refresh.
- Second-gate outcome: **conditional native viability** for an eligible user with pre-existing OpenCode `openai` ChatGPT OAuth; **blocked** for a fully temporary isolated root that must reuse native OAuth without credential exposure or ordinary-state mutation. Overall OpenCode remains technically unqualified until the separate first-gate isolation and runtime evidence is obtained.

# OpenAI Sites Free Hosting and CLI Downloads

## Assignment

**Goal:** Determine whether OpenAI ChatGPT Sites currently provides free hosting suitable for a small public open-source project website that may distribute downloadable CLI binaries.

**Scope:** Internet-only review of current official OpenAI product documentation, Help Center guidance, ChatGPT Sites terms, Service Terms, and Usage Policies. The review covers plan availability, included hosting, public access, generated and custom domains, file/object storage, direct-download support, quantitative limits, and policy or beta-service restrictions.

**Exclusions:** No GitHub or Firebase comparison, no local product or connector inspection as evidence, no account-level deployment, and no destructive or quota-consuming validation.

## Sources

- OpenAI Help Center, [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339), page displayed “Updated: yesterday” when accessed 2026-07-19.
- OpenAI/ChatGPT Learn, [Sites developer guide](https://learn.chatgpt.com/docs/sites), public-beta documentation accessed 2026-07-19; canonical OpenAI developer link redirects from `https://developers.openai.com/codex/sites`.
- OpenAI, [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/), updated 2026-07-09; accessed 2026-07-19.
- OpenAI, [Service Terms](https://openai.com/policies/service-terms/), updated 2026-06-12; accessed 2026-07-19.
- OpenAI, [Usage Policies](https://openai.com/policies/usage-policies/), effective 2025-10-29; accessed 2026-07-19.

## Findings

### Finding 1: There is no zero-dollar ChatGPT Sites hosting tier at present

ChatGPT Sites is a public beta on eligible paid plans and is explicitly unavailable on the Free and Go plans. Pro, Pro Lite, Enterprise, and Edu receive access first, while Plus and Business access is still subject to rollout. Availability also depends on region and workspace settings; Sites is unavailable in the EEA, Switzerland, and the United Kingdom at launch.

**Evidence:** OpenAI’s Help Center states that Sites is “available in public beta on paid plans except Free and Go,” names the phased plan rollout, and lists the launch-region exclusions. It repeats in the FAQ that Sites is unavailable on Free or Go. [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339)

**Implication:** OpenAI Sites does not satisfy a strict requirement for hosting available to a project owner on a free ChatGPT account. For an already-eligible paid subscriber, hosting usage is bundled with the subscription rather than documented as a separately billed hosting product, subject to the beta limits described below.

### Finding 2: Hosting is included on eligible paid plans, but only up to opaque plan-specific beta limits

During the beta, Sites usage is included up to plan-specific limits. Those limits apply across all Sites on the account, may change during the beta, and are shown inside the Sites experience rather than published as a stable numeric table. Hitting a limit can prevent creation of another Site, addition of storage, or continued public availability of a high-usage Site, while editing and management remain available.

**Evidence:** Both the Help Center and developer guide say usage is included up to plan-specific limits, that limits apply across all Sites, and that reaching a limit can stop storage additions or public availability for high-usage Sites. The Help Center directs users to the Sites experience for the limits currently shown for their plan or workspace. [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339), [Sites developer guide](https://learn.chatgpt.com/docs/sites)

**Implication:** No public official source establishes a stable bandwidth, transfer, storage, object-size, file-size, request, visitor, or Site-count allowance. “Included” should therefore be treated as conditional bundled usage, not unlimited free hosting.

### Finding 3: Public sites and an OpenAI-hosted production URL are supported

Every deployed Sites URL is a production deployment. A Site can be published to anyone on the internet when public publishing is enabled, and public visitors do not need ChatGPT workspace access. OpenAI may provide an OpenAI-owned subdomain, including one ending in `chatgpt.site`; the terms do not guarantee an exact hostname pattern.

**Evidence:** The Help Center says deployment generates a Site URL, every deployment URL is production, and “Anyone on the internet” makes the Site publicly accessible. It also says a public Site is available without ChatGPT workspace access. The ChatGPT Sites Terms say OpenAI may provide an OpenAI-owned subdomain such as one ending in `chatgpt.site`. [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339), [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/)

**Implication:** An eligible account can publish a small public project site without requiring visitors to sign in and without purchasing a separate hostname. The OpenAI subdomain is provided at OpenAI’s discretion and cannot be treated as a user-owned domain asset.

### Finding 4: A custom domain requires separate domain ownership and remains account-dependent

Where the feature is available, Sites accepts an apex domain or subdomain that the user already owns. The user must be able to change DNS records, enter the hostname in Site settings, and install the DNS values supplied by Sites. Sites does not register or include a custom domain. Custom domains are unavailable in Enterprise workspaces at launch, and the documentation’s “where available” qualifier leaves other plan or rollout eligibility account-dependent.

**Evidence:** OpenAI documents the settings and DNS workflow, explicitly says Sites does not register domains, and says custom domains are unavailable in Enterprise workspaces at launch. [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339), [Sites developer guide](https://learn.chatgpt.com/docs/sites)

**Implication:** The practical no-extra-domain-cost option is the OpenAI-hosted Site URL. A branded custom domain introduces a separate registration cost and DNS-management requirement, and availability must be confirmed in the specific account before relying on it.

### Finding 5: File object storage is supported, but direct CLI-binary distribution is not officially validated

The Sites runtime supports R2 object storage for “images, documents, audio, video, or other uploads,” and can pair R2 file contents with D1 metadata. The Sites Terms define Website Content broadly enough to include code and other materials and authorize OpenAI to host and distribute that content as needed to operate the Site. These sources establish general file storage and hosted content distribution, but they do not explicitly document anonymous direct downloads, executable or archive extensions, `Content-Disposition`, MIME handling, range requests, checksums, release versioning, or binary-specific serving behavior.

**Evidence:** The developer guide maps file-upload needs to R2 object storage and includes “other uploads” after common media and document types. The Sites Terms cover text, images, audio, code, and other materials, while separately prohibiting malware and malicious code. Searches within the current official Sites guide and Help Center found no documented `download`, `binary`, or `executable` support contract. [Sites developer guide](https://learn.chatgpt.com/docs/sites), [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/)

**Implication:** Hosting a legitimate CLI binary may be technically possible through a Site route or object storage, but current public OpenAI documentation does not validate it as a supported release-download use case. It should be classified as unclear and unsupported until a small non-malicious binary is tested end-to-end from a public Site and its HTTP behavior is confirmed.

### Finding 6: No official numeric bandwidth, transfer, storage, or file-size limits are published

The public documentation acknowledges storage and high-usage limits but publishes no numbers for transfer, bandwidth, total storage, per-object size, deployment size, executable size, downloads, requests, or visitors. Enterprise and Edu limits may vary by workspace. The only reliable source for the current allowance is the in-product limit shown to the eligible account or workspace.

**Evidence:** The Help Center says limits may change during the public beta, apply across all Sites, and should be checked in the Sites experience. The developer guide repeats the plan-specific limit behavior without a numeric schedule. [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339), [Sites developer guide](https://learn.chatgpt.com/docs/sites)

**Implication:** Capacity planning for downloadable binaries cannot be completed from public documentation. Binary size multiplied by expected download volume could trigger an unpublished usage threshold and remove the high-usage Site from public availability.

### Finding 7: Legitimate open-source software is not categorically prohibited, but security and content obligations are strict

The Sites Terms require the publisher to own or maintain the rights needed for Website Content and make the publisher responsible for Site functionality and end-user content. A Site and its content must not pose a security vulnerability or threat, include malware, viruses, surveillance, or other malicious code, enable phishing, or facilitate illegal or abusive activity. OpenAI also prohibits malicious or abusive cyber activity under its general Usage Policies.

**Evidence:** The ownership, responsibility, warranties, and restriction clauses are in Sections 1 and 2 of the [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/). The current [Usage Policies](https://openai.com/policies/usage-policies/) prohibit malicious or abusive cyber activity and attempts to compromise others’ systems or property.

**Implication:** A normal open-source CLI is not expressly excluded merely because it is executable, but the publisher must have distribution rights and must ensure the artifact is non-malicious and does not create a security threat. The absence of an executable-specific allowance remains important because automated safety controls could still restrict an artifact or Site.

### Finding 8: Beta-service reliability and takedown terms weaken its fit as a sole binary origin

OpenAI can delete, remove, unpublish, or disable a Site at any time for any reason, including policy or harm concerns. ChatGPT Sites is expressly a Beta Service. OpenAI’s Service Terms make no promise that beta services will remain generally available, uninterrupted, or error-free, or that content will be secure or preserved.

**Evidence:** Sections 5.2 and 5.3 of the [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/) establish removal discretion and beta status. Section 2 of the [Service Terms](https://openai.com/policies/service-terms/) disclaims beta-service availability, continuity, error-free operation, and content security or preservation.

**Implication:** Even if direct binary downloads work in practice, Sites is currently a weak sole system of record or sole distribution origin for release artifacts. The public project should retain an independent source of truth and recoverable release artifacts.

### Finding 9: Best-supported answer for this project

OpenAI Sites currently supports an included, public, OpenAI-hosted website for eligible paid ChatGPT accounts, with an optional existing custom domain where that feature is available. It does not offer Sites to Free or Go accounts. Its public documentation does not provide the quantitative limits or binary-serving contract required to validate it as a dependable free CLI-download host.

**Evidence:** The availability, bundled beta limits, production/public publishing, domain, R2 storage, and unsupported-use statements are consistent across the [Help Center](https://help.openai.com/en/articles/20001339), [developer guide](https://learn.chatgpt.com/docs/sites), and [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/).

**Implication:** Classify OpenAI Sites as **paid-plan bundled website hosting**, not **free hosting**. It is suitable for a small public project presentation site for an eligible subscriber, while direct CLI-binary hosting remains an unverified conditional capability rather than a documented feature.

## Notes

- Official-source searches for `download`, `binary`, `executable`, `bandwidth`, `transfer`, `file size`, and numeric R2/Sites quotas did not produce a Sites-specific support contract or quota table.
- No public Site was created for validation because this assignment was internet-only and excluded account-level deployment. Direct binary delivery therefore remains unsupported rather than validated.
- The developer guide names R2, but it does not incorporate any third-party R2 pricing or quota schedule into the Sites contract. External infrastructure limits should not be assumed to apply to ChatGPT Sites.
- The exact default hostname format is not guaranteed. The terms only say OpenAI “may” provide an OpenAI-owned subdomain such as one ending in `chatgpt.site`.
- Public-beta limits and rollout status are likely to change. Recheck the Help Center and the eligible account’s in-product limits immediately before a hosting decision.

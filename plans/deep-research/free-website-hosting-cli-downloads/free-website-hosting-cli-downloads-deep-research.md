# Free Website Hosting and CLI Downloads

## Recommendation

Use **GitHub Pages on a GitHub Free account for the website and GitHub Releases for downloadable CLI binaries**.

This is the only compared arrangement that is both verified at zero cost for the project owner and explicitly documented for public cross-platform software distribution. A public repository on GitHub Free can publish a Pages site; Releases supports anonymous binary downloads, version-specific and moving “latest” URLs, up to 1,000 assets per release, and assets smaller than 2 GiB. GitHub publishes no numeric Release bandwidth cap, although excessive-use enforcement still applies. [GitHub Free and Pages](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages), [Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits), [About releases](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases), [Linking to releases](https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases), [release asset API](https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10). Research basis: [GitHub Pages assignment](assignments/02-github-pages-free-account-limits.md), [GitHub Releases assignment](assignments/03-github-releases-binary-distribution.md), and [comparative verification](assignments/06-comparative-free-tier-verification.md).

The practical split should be:

- **GitHub Pages:** project explanation, benchmark and evaluation results, installation guidance, repository links, and download buttons.
- **GitHub Releases:** versioned macOS, Linux, and Windows archives or installers, checksums, release notes, download counts, and direct public asset URLs.
- **Download linking:** use tag-specific URLs for reproducible downloads and `/releases/latest/download/<asset-name>` only when the same asset names are published consistently and the latest release is deliberately managed. Direct links avoid using the REST API on each page view and therefore avoid its anonymous metadata rate limit.

This also keeps binary traffic and file sizes outside Pages’ 1 GB published-site maximum, 100 GB/month soft website-bandwidth limit, and ordinary-Git 100 MiB file block. A modest project website should remain comfortably within Pages’ limits when binaries are served by Releases. [Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits), [large-file guidance](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github), [REST API rate limits](https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api).

## Comparison

| Criterion            | OpenAI Sites                                                                                                               | GitHub Pages + Releases on GitHub Free                                                                              | Firebase Hosting Spark                                                                                              |
| -------------------- | -------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- |
| Owner cost           | **Paid:** unavailable on Free and Go; public beta on eligible paid plans                                                   | **Verified free:** Pages from a public repository and public Releases                                               | **Verified free:** Spark requires no payment method                                                                 |
| Public website       | Supported for eligible paid accounts                                                                                       | Supported from a public repository                                                                                  | Supported as public static Hosting                                                                                  |
| Direct CLI downloads | **Unsupported/unclear:** generic R2 uploads exist, but no public binary-download contract                                  | **Verified free:** Releases is explicitly for software and binary assets                                            | Static downloads supported for accepted files; incomplete for cross-platform executables                            |
| Transfer             | No public numeric limit; changing plan-specific limits appear in-product                                                   | Pages: soft 100 GB/month; Releases: no documented numeric bandwidth limit                                           | **Unresolved official conflict:** 360 MB/day and 10 GB/month are both published                                     |
| Storage/site size    | No public numeric allowance                                                                                                | Pages: 1 GB published-site maximum; Releases: no total release-size limit documented                                | 10 GB across retained Hosting releases                                                                              |
| Per artifact/file    | No public contract                                                                                                         | Release asset must be under 2 GiB; up to 1,000 assets per release                                                   | Up to 2 GB per file                                                                                                 |
| Executables          | No affirmative executable allowance; malware prohibited                                                                    | Binary assets are explicitly supported; publisher still owns signing, notarization, packaging, and install guidance | Spark blocks `.exe`, `.dll`, `.bat`, `.apk`, and `.ipa` on projects created on or after 2023-09-28; Blaze is exempt |
| Default domain       | OpenAI may provide a hostname such as one ending in `chatgpt.site`                                                         | `<owner>.github.io` or `<owner>.github.io/<repository>`; Release assets remain on GitHub URLs                       | `SITE_ID.web.app` and `SITE_ID.firebaseapp.com` over SSL                                                            |
| Custom domain        | Requires a separately owned domain; availability is conditional; Enterprise excluded at launch; TLS lifecycle undocumented | Owned apex or subdomain supported with managed HTTPS and HTTPS enforcement                                          | Owned apex or subdomain supported; Firebase provisions and renews SSL                                               |
| Over-limit behavior  | Site creation, storage additions, or public availability may stop at unpublished beta limits                               | Pages may rate-limit, stop serving, or contact the owner; Releases remains subject to excessive-use policy          | Hosting may be disabled after a short grace period; excess storage blocks deployments                               |

Evidence: [OpenAI Sites assignment](assignments/01-openai-sites-free-hosting-and-downloads.md), [GitHub Pages assignment](assignments/02-github-pages-free-account-limits.md), [GitHub Releases assignment](assignments/03-github-releases-binary-distribution.md), [Firebase Spark assignment](assignments/04-firebase-hosting-spark-plan-limits.md), [cross-platform assignment](assignments/05-cross-platform-domain-download-security.md), and [comparative verification](assignments/06-comparative-free-tier-verification.md).

## Why GitHub Fits Best

### Website hosting is verified on the Free account

GitHub Pages is available to GitHub Free personal accounts when the source repository is public. The default project URL is predictable, HTTPS is included, and an already-owned apex domain or subdomain can be attached with managed HTTPS. Domain registration remains an external cost and responsibility. GitHub recommends domain verification and avoiding wildcard DNS to reduce takeover risk. [Pages eligibility and URL shapes](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages), [custom-domain configuration](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site), [domain verification](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/verifying-your-custom-domain-for-github-pages), [HTTPS](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https).

The project’s planned content—project explanation, evaluation results, repository links, and download guidance—matches Pages’ documented project-showcase purpose. The relevant constraint is that Pages is not intended as free hosting for a commercial SaaS or transaction-heavy service. [Pages product terms](https://docs.github.com/en/site-policy/github-terms/github-terms-for-additional-products-and-features#pages).

### Binary distribution is an explicit Releases use case

GitHub describes Releases as deployable software iterations with links to binary files. Public assets are anonymously downloadable through `browser_download_url`, tag-specific paths, or the documented latest-asset path. The API exposes asset size, download count, media type, and a SHA-256 digest. The host does not provide code signing, notarization, executable permission preservation, platform selection, or installation handling; package each platform build conventionally and publish clear verification and install instructions. [About releases](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases), [linking to releases](https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases), [release asset API](https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10).

GitHub’s “no bandwidth limit” statement for Releases means no published numeric Release quota. GitHub can still throttle file hosting, limit activity, or act against unusually excessive bandwidth or abusive distribution. This is a reasonable trade-off at the project’s modest expected usage, but it is not a transfer SLA. [GitHub Acceptable Use Policies](https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies).

## Conditional Alternatives

### Firebase Hosting for the site, GitHub Releases for binaries

Choose this split when Firebase’s CDN, preview/deployment workflow, or path-specific response-header configuration is materially preferable. Spark is a verified no-cost website host, supplies `web.app` and `firebaseapp.com` HTTPS domains, and supports an already-owned custom domain with provisioned and renewed SSL. [Firebase Hosting quickstart](https://firebase.google.com/docs/hosting/quickstart), [custom domains](https://firebase.google.com/docs/hosting/custom-domain), [Hosting configuration](https://firebase.google.com/docs/hosting/full-config). Research basis: [Firebase assignment](assignments/04-firebase-hosting-spark-plan-limits.md) and [cross-platform assignment](assignments/05-cross-platform-domain-download-security.md).

Keep binaries on GitHub Releases. Spark blocks several important executable extensions, and its transfer entitlement is unclear across current official sources. Splitting the responsibilities preserves Firebase’s website advantages while avoiding its direct-download weaknesses.

### Firebase-only hosting

This is a lower-fit option for very modest traffic and only when every chosen artifact format is accepted. Firebase supports public static assets, a 2 GB per-file maximum, and 10 GB of Hosting storage across retained releases. However:

- Spark blocks `.exe`, `.dll`, `.bat`, `.apk`, and `.ipa` for projects created on or after 2023-09-28; Blaze removes the restriction but is pay-as-you-go. [Firebase Hosting FAQ](https://firebase.google.com/docs/hosting/faq-and-troubleshooting).
- Official sources do not affirmatively guarantee archives or extensionless binaries as a workaround.
- Both cache hits and misses count toward transfer, and exhaustion can disable Hosting after a short grace period. [Hosting usage, quotas, and pricing](https://firebase.google.com/docs/hosting/usage-quotas-pricing).
- Retained previous releases consume the 10 GB storage allowance; reaching it blocks new Spark deployments until content is deleted or billing is enabled.

At the project’s expected modest usage, Firebase-only could work for site assets and accepted archives, but the Windows executable restriction and transfer ambiguity make it operationally weaker than GitHub Releases.

### OpenAI Sites as a presentation layer

Use OpenAI Sites only when the owner already has an eligible paid ChatGPT plan and specifically values the Sites authoring experience. Sites can publish publicly and may provide an OpenAI-hosted subdomain; a separately owned custom domain can be connected where the feature is available. [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339), [Sites developer guide](https://learn.chatgpt.com/docs/sites), [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/). Research basis: [OpenAI Sites assignment](assignments/01-openai-sites-free-hosting-and-downloads.md).

It is lower fit because Sites is unavailable on Free and Go, its numeric storage and transfer limits are unpublished and plan-specific, binary delivery is not documented as a supported use case, and beta terms provide no continuity or preservation assurance. If selected, use it only as the front end and keep GitHub Releases as the release system of record. [OpenAI Service Terms](https://openai.com/policies/service-terms/).

## Lower-Fit Interpretations

- **Serving binaries directly from GitHub Pages:** technically possible for static files, but Pages lacks per-file MIME or `Content-Disposition` controls, ordinary Git blocks files above 100 MiB, the whole site is capped at 1 GB, and all such traffic consumes the soft 100 GB/month Pages allowance. Releases is the documented software-distribution surface.
- **Treating Firebase archives as a guaranteed executable workaround:** unsupported. The blocklist does not name `.zip`, `.tar.gz`, or extensionless Unix artifacts, but omission is not an affirmative compatibility contract.
- **Treating OpenAI Sites hosting as free:** incorrect under the current official availability statement; Sites is unavailable on Free and Go.
- **Treating “no Release bandwidth limit” as unlimited guaranteed delivery:** unsupported. GitHub retains excessive-bandwidth and abuse controls.
- **Using a custom domain to brand GitHub Release asset URLs:** unavailable. The custom domain brands the Pages site; the final asset URL remains on GitHub infrastructure.

## Unresolved Blocker

Firebase’s official sources conflict on Spark Hosting transfer:

- The [Firebase pricing table](https://firebase.google.com/pricing) and Hosting product material state **360 MB/day**.
- The detailed [Hosting usage, quotas, and pricing guide](https://firebase.google.com/docs/hosting/usage-quotas-pricing) states **10 GB/month**, counts CDN cache hits and misses, and describes monthly disablement behavior.

The supplied research does not establish whether both limits apply, whether one supersedes the other, or whether the daily figure is a rounded presentation of the monthly allowance. Preserve the classification as **Unclear**. For conservative planning, remain below both figures; for example, a 20 MB artifact reaches 360 MB at 18 downloads in one day and approximately 10 GB at 500 downloads in a month, before website assets. Research basis: [Firebase assignment](assignments/04-firebase-hosting-spark-plan-limits.md) and [comparative verification](assignments/06-comparative-free-tier-verification.md).

## Decision

For Skill Issue’s modest expected usage, adopt **GitHub Pages + GitHub Releases on GitHub Free**. It has the clearest zero-cost contract, the strongest direct-download support, the largest documented website transfer allowance among the two verified free site hosts, and the only explicit cross-platform binary-distribution surface in the comparison. Use the free `github.io` project URL initially or attach an already-owned custom domain with HTTPS; keep download buttons pointed at versioned or deliberately managed latest Release assets.

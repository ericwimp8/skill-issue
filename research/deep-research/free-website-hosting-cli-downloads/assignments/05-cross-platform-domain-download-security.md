# Cross-Platform Domain, Download, and Security Controls

## Assignment

- **Goal:** Compare the domain, download-delivery, security, abuse-control, and acceptable-use behavior of OpenAI Sites, GitHub Pages/Releases, and Firebase Hosting for a small public open-source website distributing cross-platform CLI binaries.
- **Scope:** Internet-only research using primary official product documentation, product terms, and a small non-destructive HTTPS probe. Verified free capabilities are distinguished from paid, unavailable, unclear, or unsupported behavior.
- **Exclusions:** General storage, bandwidth, build, and project quotas except where a limit directly changes whether a public binary can be delivered; unrelated application-hosting features; secondary summaries; local project implementation.

## Sources

- OpenAI, [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339), updated July 2026; inspected 2026-07-19.
- OpenAI, [Sites developer guide](https://learn.chatgpt.com/docs/sites), public-beta documentation; inspected 2026-07-19.
- OpenAI, [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/), updated 2026-07-09; inspected 2026-07-19.
- OpenAI, [Understanding responsibilities for your ChatGPT Sites](https://help.openai.com/en/articles/20001337-understanding-responsibilities-for-your-chatgpt-sites), updated July 2026; inspected 2026-07-19.
- GitHub, [What is GitHub Pages?](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages), current GitHub.com documentation; inspected 2026-07-19.
- GitHub, [Creating a GitHub Pages site](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site), current GitHub.com documentation; inspected 2026-07-19.
- GitHub, [Managing a custom domain for your GitHub Pages site](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site), current GitHub.com documentation; inspected 2026-07-19.
- GitHub, [Securing your GitHub Pages site with HTTPS](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https), current GitHub.com documentation; inspected 2026-07-19.
- GitHub, [Verifying your custom domain for GitHub Pages](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/verifying-your-custom-domain-for-github-pages), current GitHub.com documentation; inspected 2026-07-19.
- GitHub, [About releases](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases), current GitHub.com documentation; inspected 2026-07-19.
- GitHub, [Linking to releases](https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases), current GitHub.com documentation; inspected 2026-07-19.
- GitHub, [REST API endpoints for release assets](https://docs.github.com/en/rest/releases/assets), API version 2026-03-10 examples; inspected 2026-07-19.
- GitHub, [Releases now expose digests for release assets](https://github.blog/changelog/2025-06-03-releases-now-expose-digests-for-release-assets/), published 2025-06-03; inspected 2026-07-19.
- GitHub, [GitHub Acceptable Use Policies](https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies), current policy; inspected 2026-07-19.
- GitHub, [GitHub Active Malware or Exploits](https://docs.github.com/en/site-policy/acceptable-use-policies/github-active-malware-or-exploits), current policy; inspected 2026-07-19.
- GitHub, [GitHub Terms for Additional Products and Features](https://docs.github.com/en/site-policy/github-terms/github-terms-for-additional-products-and-features), current terms; inspected 2026-07-19.
- Firebase, [Firebase Hosting](https://firebase.google.com/docs/hosting), last updated 2026-06-30; inspected 2026-07-19.
- Firebase, [Get started with Firebase Hosting](https://firebase.google.com/docs/hosting/quickstart), last updated 2026-07-10; inspected 2026-07-19.
- Firebase, [Connect a custom domain](https://firebase.google.com/docs/hosting/custom-domain), last updated 2026-07-10; inspected 2026-07-19.
- Firebase, [Configure Hosting behavior](https://firebase.google.com/docs/hosting/full-config), updated July 2026; inspected 2026-07-19.
- Firebase, [FAQ and troubleshooting](https://firebase.google.com/docs/hosting/faq-and-troubleshooting), last updated 2026-07-10; inspected 2026-07-19.
- Google Cloud, [Acceptable Use Policy](https://cloud.google.com/terms/aup), last modified 2026-06-23; inspected 2026-07-19.

## Findings

### Finding 1: Only GitHub and Firebase have verified no-cost public-hosting paths

GitHub Pages is available for public repositories on GitHub Free, while Firebase documents no-cost default domains, custom domains, CDN delivery, and SSL on its Spark/no-cost path. ChatGPT Sites is a public beta on paid ChatGPT plans except Free and Go, with plan-specific limits and rollout restrictions. Consequently, OpenAI Sites is not a verified free-hosting candidate for this use case.

**Evidence:** GitHub documents Pages availability in public repositories with GitHub Free and GitHub Free for organizations in [What is GitHub Pages?](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages). Firebase says Hosting works on `web.app` and `firebaseapp.com` subdomains at no cost and supports a custom domain with SSL in [Get started with Firebase Hosting](https://firebase.google.com/docs/hosting/quickstart), and its product page states that custom domain, CDN, and SSL are included in the no-cost starting offer in [Firebase Hosting](https://firebase.google.com/products/hosting). OpenAI states that Sites is unavailable on Free and Go in [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339).

**Implication:** For a project whose requirement is genuinely free public hosting, evaluate GitHub Pages and Firebase Hosting directly. Treat OpenAI Sites as a paid beta front end, not a free substitute.

### Finding 2: GitHub offers the most explicit free custom-domain and TLS contract

GitHub Pages provides predictable default URLs and an explicit custom-domain DNS model. User or organization sites use `<owner>.github.io`; project sites use `<owner>.github.io/<repository>`. Apex domains use `ALIAS`, `ANAME`, or GitHub's documented `A` records, while subdomains use a `CNAME` pointed directly to `<owner>.github.io`. Correctly configured Pages domains support HTTPS and HTTPS enforcement. GitHub recommends domain verification and warns against wildcard DNS because of takeover risk.

**Evidence:** Default-domain shapes and free public-repository availability are documented in [What is GitHub Pages?](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages). The DNS records, possible 24-hour propagation period, apex/`www` redirects, and takeover warnings are documented in [Managing a custom domain for your GitHub Pages site](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site). GitHub states that all Pages sites, including correctly configured custom domains, support HTTPS and enforcement in [Securing your GitHub Pages site with HTTPS](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https). TXT-based ownership protection is documented in [Verifying your custom domain for GitHub Pages](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/verifying-your-custom-domain-for-github-pages).

**Implication:** A custom domain can safely front a GitHub Pages marketing/download page on the free public-repository path. Keep the domain-verification TXT record, avoid wildcard DNS, and expect certificate/DNS activation to take time.

### Finding 3: Firebase has strong free domain/TLS behavior, with more certificate-state caveats

Every deployed Firebase Hosting site receives `SITE_ID.web.app` and `SITE_ID.firebaseapp.com` live endpoints serving the same content, and Hosting is SSL-only. Firebase supports apex domains and subdomains, requires ownership verification when requested, provisions and automatically renews certificates, and offers an advanced setup for migration. The owner must retain the verification TXT record when required. Certificate provisioning can take up to 24 hours after DNS changes; restrictive CAA records or conflicting provider records can block issuance. Firebase limits a custom apex domain to 20 Hosting subdomains because of certificate-minting limits.

**Evidence:** Default domains and SSL-only public delivery are documented in [Get started with Firebase Hosting](https://firebase.google.com/docs/hosting/quickstart) and [Test, preview, then deploy](https://firebase.google.com/docs/hosting/test-preview-deploy). The retained TXT record, A/AAAA routing, automated certificate provisioning, migration flow, CAA caveats, and 20-subdomain limit are documented in [Connect a custom domain](https://firebase.google.com/docs/hosting/custom-domain). Firebase also documents that default and custom endpoints serve the same live channel in [Manage live and preview channels, releases, and versions](https://firebase.google.com/docs/hosting/manage-hosting-resources).

**Implication:** Firebase's free site-hosting domain story is comparable to Pages for a single small site, but migration and certificate state require more operational attention. It remains suitable as a front end when binaries are delivered elsewhere.

### Finding 4: OpenAI Sites custom domains are conditional and TLS behavior is undocumented

OpenAI Sites may provide an OpenAI-owned hostname such as one ending in `chatgpt.site`; the developer guide shows an example `*.openai.chatgpt.site` hostname. Where custom domains are available, owners can connect an apex domain or subdomain by copying platform-provided DNS records. Custom domains are unavailable to Enterprise workspaces at launch. Official Sites documentation does not identify the DNS record types, certificate authority, certificate provisioning/renewal behavior, HTTPS enforcement, HSTS behavior, or activation failure states.

**Evidence:** The conditional default subdomain and subdomain-squatting restrictions appear in the [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/). The custom-domain flow and Enterprise exclusion are in the [Sites developer guide](https://learn.chatgpt.com/docs/sites) and [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339). As a limited non-contractual probe, `curl -I https://goblin-tales.openai.chatgpt.site` on 2026-07-19 completed a TLS handshake and returned an HTTP/2 `404`, confirming that the official guide's example hostname resolved over HTTPS at probe time, but not establishing a documented platform guarantee.

**Implication:** A custom domain may be practical for an eligible non-Enterprise Sites account, but its TLS lifecycle is unsupported by public documentation. Do not select Sites when a documented DNS/TLS contract is a hard requirement without validating the live account flow first.

### Finding 5: GitHub Releases is explicitly designed for public CLI binary delivery

GitHub Releases is the only compared product that explicitly describes packaging software with links to binary files for others to download. Public release assets can be downloaded without authentication, have stable version-specific browser-download URLs, and can be linked through a latest-release path. Uploaders declare the asset media type, and API consumers can request `application/octet-stream`; GitHub may redirect or stream the bytes. Release asset URLs remain GitHub URLs, so a Pages custom domain brands the site but not the final asset host.

**Evidence:** The binary-distribution purpose and public release model are documented in [About releases](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases). Stable and latest direct-download URL shapes are documented in [Linking to releases](https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases). The public unauthenticated download behavior, `browser_download_url`, required upload `Content-Type`, and `application/octet-stream` download behavior are documented in [REST API endpoints for release assets](https://docs.github.com/en/rest/releases/assets). The official URL patterns are under `github.com/<owner>/<repo>/releases/...`; GitHub documents no custom-domain binding for Releases.

**Implication:** Use GitHub Releases as the binary origin even when the marketing site lives on GitHub Pages, Firebase, or OpenAI Sites. A download button can link directly to versioned or latest assets, but users will ultimately download from GitHub's domain.

### Finding 6: GitHub provides stronger documented artifact-integrity controls than the site hosts

GitHub automatically computes an immutable SHA-256 digest at upload time for each release asset and exposes it in the Releases UI and APIs. GitHub also offers immutable releases and attestations, although the inspected documentation did not establish whether every immutability control is available on GitHub Free. Firebase Hosting and OpenAI Sites documentation inspected for this assignment does not describe equivalent automatic checksums, signatures, attestations, or immutable release objects for downloadable files.

**Evidence:** Automatic release-asset digests are documented in GitHub's [Releases now expose digests for release assets](https://github.blog/changelog/2025-06-03-releases-now-expose-digests-for-release-assets/), and digest fields appear in [REST API endpoints for release assets](https://docs.github.com/en/rest/releases/assets). Immutable tags, assets, and release attestations are described in [Immutable releases](https://docs.github.com/en/code-security/concepts/supply-chain-security/immutable-releases). Firebase's static-file and header behavior is described in [Firebase Hosting](https://firebase.google.com/docs/hosting) and [Configure Hosting behavior](https://firebase.google.com/docs/hosting/full-config); OpenAI's file-storage support is described in the [Sites developer guide](https://learn.chatgpt.com/docs/sites), with no corresponding integrity feature stated in either product's inspected documentation.

**Implication:** GitHub Releases gives downloaders a documented integrity signal without the publisher building a separate checksum system. For Firebase or Sites-hosted files, publish checksums/signatures separately and treat integrity as publisher-owned.

### Finding 7: GitHub Pages can serve files but offers little per-file response control

GitHub Pages serves files in the deployed directory structure and supports more than 750 MIME types derived from `mime-db`, but publishers cannot define a custom MIME type per file or repository. The Pages documentation does not expose a general response-header configuration mechanism or a per-file `Content-Disposition` setting. GitHub itself directs sites with unsuitable high-bandwidth usage toward Releases, reinforcing the separation between the website and binary delivery.

**Evidence:** Static file paths, MIME inference, and the inability to set per-file or per-repository custom MIME types are documented in [Creating a GitHub Pages site](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site). GitHub's usage guidance recommends Releases when Pages usage is too bandwidth-heavy in [GitHub Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits).

**Implication:** Pages is a good free site host, but Releases is the more controllable and officially intended download surface. Do not depend on Pages for custom download headers or executable-specific response behavior.

### Finding 8: Firebase supports response headers, but its free plan blocks several executable formats

Firebase Hosting deploys all files in the configured public directory as public static assets and supports path-specific custom response headers. This generic mechanism can express a `Content-Disposition` header, but the official examples do not specifically demonstrate that header or document binary-specific MIME handling, so browser download behavior should be verified for chosen archive formats. More importantly, Spark projects created on or after 2023-09-28 cannot upload or host `.exe`, `.dll`, `.bat`, `.apk`, or `.ipa` files. Blaze projects are exempt, making direct Windows `.exe` delivery a paid-only Firebase capability. The documented blocklist does not name `.zip`, `.tar.gz`, or extensionless Unix binaries, but that omission is not an affirmative compatibility guarantee.

**Evidence:** The public-directory model is documented in [Get started with Firebase Hosting](https://firebase.google.com/docs/hosting/quickstart). Path-specific arbitrary key/value response headers are documented in [Configure Hosting behavior](https://firebase.google.com/docs/hosting/full-config); Firebase separately notes that it overwrites custom HSTS on default `*.web.app` endpoints while connected custom domains serve the configured HSTS value. The Spark executable blocklist, effective date, and Blaze exemption are documented in [FAQ and troubleshooting](https://firebase.google.com/docs/hosting/faq-and-troubleshooting).

**Implication:** Firebase Spark cannot be the complete direct-download origin for a normal cross-platform CLI that publishes Windows `.exe` files. Either distribute archives only after testing every target format, upgrade to Blaze, or use GitHub Releases for binaries while retaining Firebase for the site.

### Finding 9: OpenAI Sites file storage does not establish binary-download support

The Sites runtime supports R2 object storage for images, documents, audio, video, and other uploads, and Sites can include uploaded files and generated artifacts. However, the official documentation does not state that build-time CLI binaries can be published as anonymous static assets, list permitted or blocked executable formats, define public object URLs, document MIME inference, or expose `Content-Disposition`/response-header controls. Public Sites can be visited without workspace access, but that does not prove that arbitrary stored objects are anonymous download endpoints.

**Evidence:** R2 support for files and uploads is described in [Choose a supported site shape](https://learn.chatgpt.com/docs/sites#choose-a-supported-site-shape). Public audience behavior and the presence of files/artifacts are described in the [Sites developer guide](https://learn.chatgpt.com/docs/sites) and [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339). No binary-delivery, MIME, extension, or header contract appears in the inspected official Sites guide, Help Center article, or terms.

**Implication:** Treat direct CLI-binary hosting on OpenAI Sites as unsupported/unclear. A Sites front end should link to GitHub Releases or another documented software-distribution origin.

### Finding 10: All three prohibit malicious distribution, but enforcement models differ

OpenAI expressly prohibits Sites from containing malware, viruses, surveillance, security threats, phishing, deception, or abuse, and may use safety systems, review, restrictions, or takedowns. GitHub prohibits using its platform to deliver malicious executables in support of unlawful attacks, while explicitly allowing legitimate dual-use security content and reserving the ability to restrict an abused instance. Google Cloud's AUP prohibits phishing and distribution of viruses, Trojan horses, corrupted files, and other destructive or deceptive items; violation may lead to suspension or termination. Firebase additionally uses Spark executable-extension blocking as an anti-abuse control.

**Evidence:** OpenAI restrictions and removal rights appear in the [ChatGPT Sites Terms](https://openai.com/policies/chatgpt-sites-terms/) and its safety/enforcement approach in [Understanding responsibilities for your ChatGPT Sites](https://help.openai.com/en/articles/20001337-understanding-responsibilities-for-your-chatgpt-sites). GitHub's active-malware boundary and dual-use allowance appear in [GitHub Active Malware or Exploits](https://docs.github.com/en/site-policy/acceptable-use-policies/github-active-malware-or-exploits) and the broader [GitHub Acceptable Use Policies](https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies). Google's phishing/malware prohibitions and suspension consequence appear in the [Google Cloud Acceptable Use Policy](https://cloud.google.com/terms/aup). Firebase's extension block is in [FAQ and troubleshooting](https://firebase.google.com/docs/hosting/faq-and-troubleshooting).

**Implication:** Legitimate signed open-source CLI releases fit the documented purpose of GitHub Releases most clearly. Regardless of host, publish source, checksums, release notes, and signing information to reduce false-positive and user-trust risk; no platform guarantees immunity from automated abuse enforcement or takedown.

### Finding 11: The strongest cross-platform arrangement separates site hosting from binary hosting

For a free public open-source CLI, the best-supported arrangement is GitHub Pages on a verified custom domain for the website plus GitHub Releases for binaries. Firebase Hosting is a viable alternative site front end with free custom-domain TLS, but GitHub Releases should still deliver the binaries because Spark blocks common Windows executable extensions. OpenAI Sites can provide a hosted interactive front end for an eligible paid account, but its beta status, lack of Free/Go access, conditional custom-domain availability, undocumented TLS lifecycle, and undocumented binary-serving behavior make it a lower-confidence delivery platform.

**Evidence:** This conclusion combines the documented Pages domain/TLS controls in [Managing a custom domain for GitHub Pages](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site), the binary-specific Releases contract in [About releases](https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases), Firebase's Spark executable restriction in [FAQ and troubleshooting](https://firebase.google.com/docs/hosting/faq-and-troubleshooting), and the paid-beta/unsupported-use boundaries in [Creating and managing ChatGPT Sites](https://help.openai.com/en/articles/20001339).

**Implication:** Select the website host based on authoring and domain preference, but centralize cross-platform artifacts in GitHub Releases. This preserves a free path, avoids Firebase's executable restriction, provides stable public download links and automatic digests, and avoids relying on undocumented OpenAI Sites behavior.

## Notes

- The OpenAI Sites documentation is newly published and explicitly beta; domain, file-storage, and runtime details may change quickly. Claims marked unclear are gaps in the inspected official documentation, not proof that the capability is technically impossible.
- The OpenAI example subdomain HTTPS probe returned `404`; it validated DNS/TLS reachability at one moment, not the availability of the deleted/example Site or a contractual TLS guarantee.
- Firebase's generic custom-header schema appears capable of setting `Content-Disposition`, but the official Hosting docs do not name that header or demonstrate binary downloads. A project-level HTTP response probe would be appropriate before relying on exact browser behavior.
- The Firebase Spark blocklist explicitly names `.exe`, `.dll`, `.bat`, `.apk`, and `.ipa`. Archive formats are not listed, but packaging an executable in an archive does not waive the Google Cloud AUP or guarantee that future anti-abuse controls will accept it.
- GitHub documents automatic SHA-256 digests for all uploaded release assets. The inspected immutable-release documentation did not identify plan eligibility, so immutability/attestation availability should not be classified as verified free without a separate account-level check.
- No official GitHub documentation found in this lane claims that every release asset is malware-scanned. GitHub documents policy enforcement, dual-use handling, integrity digests, and dependency malware alerts, which are separate controls.

# GitHub Pages on GitHub Free: Website and CLI Download Limits

## Assignment

**Goal:** Assess GitHub Pages on a GitHub Free personal account for a small open-source project website that may expose downloadable CLI assets.

**Scope:** Internet-only research using current primary GitHub documentation and GitHub terms. Verify Free-plan eligibility, static download behavior, bandwidth, storage, build, deployment, file-size, executable, domain, TLS, usage-policy, and soft-limit behavior. Classify capabilities as verified free, paid/unavailable, or unclear/unsupported.

**Exclusions:** GitHub Releases except for a short hosting-boundary clarification; Firebase; OpenAI Sites; implementation instructions for the project.

## Sources

- GitHub Docs, “What is GitHub Pages?”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages
- GitHub Docs, “Creating a GitHub Pages site”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site
- GitHub Docs, “GitHub Pages limits”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits
- GitHub Docs, “Using custom workflows with GitHub Pages”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages
- GitHub Docs, “GitHub Actions billing”, current online documentation accessed 2026-07-19: https://docs.github.com/en/billing/concepts/product-billing/github-actions
- GitHub Docs, “Actions limits”, current online documentation accessed 2026-07-19: https://docs.github.com/en/actions/reference/limits
- GitHub Docs, “About large files on GitHub”, current online documentation accessed 2026-07-19: https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github
- GitHub Docs, “About Git Large File Storage”, current online documentation accessed 2026-07-19: https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-git-large-file-storage
- GitHub Docs, “About custom domains and GitHub Pages”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/about-custom-domains-and-github-pages
- GitHub Docs, “Managing a custom domain for your GitHub Pages site”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site
- GitHub Docs, “Verifying your custom domain for GitHub Pages”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/verifying-your-custom-domain-for-github-pages
- GitHub Docs, “Troubleshooting custom domains and GitHub Pages”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/troubleshooting-custom-domains-and-github-pages
- GitHub Docs, “Securing your GitHub Pages site with HTTPS”, current online documentation accessed 2026-07-19: https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https
- GitHub Docs, “GitHub Terms for Additional Products and Features”, Pages section, current terms accessed 2026-07-19: https://docs.github.com/en/site-policy/github-terms/github-terms-for-additional-products-and-features#pages
- GitHub Docs, “GitHub Acceptable Use Policies”, current policy accessed 2026-07-19: https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies

## Findings

### Finding 1: A public-repository Pages site is a verified GitHub Free capability

GitHub Pages is included for public repositories owned by GitHub Free personal accounts and GitHub Free organizations. For a Free account, the source repository must be public. Pages from private repositories require GitHub Pro, Team, Enterprise Cloud, or Enterprise Server; a Free personal account therefore cannot keep the Pages source private. The resulting site is public on the internet. This matches the intended small open-source project use case.

**Evidence:** GitHub’s eligibility table explicitly includes public repositories on GitHub Free and assigns public-and-private repository eligibility to paid plans; the creation guide repeats that a Free-owned Pages repository must be public and warns that Pages sites are publicly available ([What is GitHub Pages?](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages), [Creating a GitHub Pages site](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site)). GitHub’s additional terms say each account has access to the Pages static hosting service and describe its intended use as showcasing personal and organizational projects ([Additional Product Terms — Pages](https://docs.github.com/en/site-policy/github-terms/github-terms-for-additional-products-and-features#pages)).

**Implication:** **Verified free:** a small public open-source project site. **Paid/unavailable on Free:** a Pages site sourced from a private repository. Any CLI asset published through the site must be treated as public.

### Finding 2: Pages can expose static CLI assets at direct public URLs, with browser-serving caveats

GitHub says Pages publishes any static files in the publishing source and preserves the source directory structure at the site. A ZIP, tarball, checksum, shell script, or benign executable included as a static file can therefore have a direct Pages URL. GitHub Pages supports more than 750 MIME types across thousands of extensions, but site owners cannot configure a custom MIME type per file or repository. GitHub does not document a Pages setting for forcing `Content-Disposition: attachment`, so whether a file opens, downloads, or prompts depends on its recognized MIME type and the client/browser. Archive formats are consequently the best-supported practical form for CLI downloads.

**Evidence:** The site-creation guide states that Pages “publishes any static files” and that each file becomes available under the corresponding published directory structure; the same guide documents broad MIME support and the inability to customize MIME types per file or repository ([Creating a GitHub Pages site](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site)). Pages is explicitly a static hosting service and does not provide a server-side runtime ([What is GitHub Pages?](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages)).

**Implication:** **Verified free:** public direct URLs for static download assets. **Unavailable:** server-side download authorization, dynamic entitlement checks, or repository-level MIME overrides. **Unclear/unsupported:** an official guarantee that every executable extension is served with a download-triggering header; package CLI builds as `.zip` or `.tar.gz` and link to those assets rather than relying on raw executable behavior.

### Finding 3: Pages-specific capacity is suitable for small assets but bounded by soft transfer limits

The published Pages site may be no larger than 1 GB. Its source repository has a recommended 1 GB limit. Pages has a soft bandwidth limit of 100 GB per month, and additional request rate limiting can return HTTP `429`. GitHub describes these as quality-of-service limits rather than quotas intended to block legitimate use, but exceeding them may make the site unavailable or trigger support contact. GitHub may recommend a CDN, another GitHub feature, or another host.

**Evidence:** GitHub lists the 1 GB source-repository recommendation, 1 GB maximum published-site size, 100 GB/month soft bandwidth limit, and possible HTTP `429` rate limiting together in the official [GitHub Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits). The Acceptable Use Policies also reserve GitHub’s right to throttle file hosting, suspend an account, limit activity, or—after notice—delete repositories that impose excessive infrastructure strain ([Acceptable Use Policies, section 9](https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies#9-excessive-bandwidth-use)).

**Implication:** **Verified free:** modest static downloads within the shared 100 GB/month soft site-transfer envelope. A 20 MB asset would consume roughly 100 GB after about 5,000 complete downloads, before counting page traffic; this is an estimate rather than a GitHub quota conversion. **Unsupported:** treating Pages as an unlimited binary CDN or assuming continued serving after the soft limit is exceeded.

### Finding 4: Regular Git file limits constrain checked-in assets, and Git LFS is unavailable to Pages

For assets committed to the Pages source repository through ordinary Git, GitHub warns above 50 MiB and blocks files larger than 100 MiB. Browser uploads are capped at 25 MiB. Git LFS cannot be used with GitHub Pages, even though GitHub Free otherwise supports LFS objects up to 2 GB. GitHub’s large-file guidance recommends Releases rather than regular repository tracking when distributing large binaries.

Custom Actions publishing adds a nuance: GitHub documents that the compressed Pages artifact must contain a single tar file under 10 GB, while the final published site remains capped at 1 GB. The official Pages documentation does not state a separate maximum size for one generated file inside that final artifact. Therefore, the 100 MiB regular-Git cap is verified for checked-in assets, but it should not be presented as a documented per-file serving cap for generated Pages output.

**Evidence:** GitHub documents the 25 MiB browser-upload cap, warning above 50 MiB, hard block above 100 MiB, and recommendation to use Releases for large binary distribution in [About large files on GitHub](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github). GitHub explicitly says Git LFS cannot be used with Pages in [About Git Large File Storage](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-git-large-file-storage). The custom Pages workflow guide specifies a tar file under 10 GB ([Using custom workflows with GitHub Pages](https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages)), while the separate Pages limits document caps the published site at 1 GB ([GitHub Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits)).

**Implication:** **Verified free:** checked-in assets below 100 MiB each, subject to the whole-site and bandwidth limits. **Unavailable:** Git LFS-backed Pages assets. **Unclear/unsupported:** a documented hard per-file ceiling for an asset generated during a custom Actions build; do not infer that the 10 GB upload-artifact ceiling permits a site or individual download above the 1 GB published-site maximum. **Boundary:** GitHub’s own docs direct large binary distribution toward Releases, but Releases behavior is outside this assignment.

### Finding 5: Build and deployment limits differ between branch publishing and custom Actions

Pages deployments time out after 10 minutes. Pages has a soft limit of 10 builds per hour, but GitHub explicitly exempts sites built and published with a custom GitHub Actions workflow from that Pages build-count limit. Standard GitHub-hosted runners are free for public repositories and for GitHub Pages, so a public Free-account project can use the custom workflow route without consuming paid runner minutes. The workflow is still governed by GitHub Actions system limits, and the Pages deployment itself remains a Pages deployment.

GitHub separately says site changes can take up to 10 minutes to publish after a push. This propagation estimate is not an additional hourly build allowance.

**Evidence:** The deployment timeout, 10-builds/hour soft limit, and custom-workflow exemption appear in [GitHub Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits). GitHub documents that standard hosted runners are free in public repositories and for Pages in [GitHub Actions billing](https://docs.github.com/en/billing/concepts/product-billing/github-actions). General workflow constraints are documented separately in [Actions limits](https://docs.github.com/en/actions/reference/limits). The publishing delay appears in [Creating a GitHub Pages site](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site).

**Implication:** **Verified free:** branch-based publishing within 10 builds/hour, or custom Actions publishing without that particular Pages build-count ceiling. **Still limited:** a 10-minute Pages deployment timeout and separate Actions constraints. **Unsupported:** interpreting the custom-workflow exemption as unlimited workflow execution or unlimited deployment throughput.

### Finding 6: Default `github.io` domains and site counts are predictable

A user or organization site uses `https://<owner>.github.io` and requires a repository named `<owner>.github.io`; GitHub allows one user or organization Pages site per account. A project site defaults to `https://<owner>.github.io/<repositoryname>` and GitHub allows one Pages site per repository. Sites using `github.io` domains and created after 15 June 2016 are served over HTTPS automatically.

**Evidence:** GitHub’s site-type table gives the repository naming, default URLs, one-account-site limit, and one-project-site-per-repository limit ([What is GitHub Pages?](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages)). Automatic HTTPS for current `github.io` sites is documented in [Securing your GitHub Pages site with HTTPS](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https).

**Implication:** **Verified free:** an HTTPS `github.io` project URL with no separately purchased domain. A repository-level project site is the natural fit if the account may need its one root user site for another purpose.

### Finding 7: Custom domains and managed TLS are included, but DNS ownership and configuration remain the operator’s responsibility

GitHub Pages supports apex domains, `www` subdomains, and custom subdomains that the operator owns. Subdomains use a DNS `CNAME` pointing to `<owner>.github.io`; apex domains use `A`, `AAAA`, `ALIAS`, or `ANAME` records. Correctly configured custom domains support HTTPS and HTTPS enforcement. GitHub performs a DNS check and automatically requests a TLS certificate from Let’s Encrypt. DNS changes may take up to 24 hours to propagate, and GitHub says HTTPS can take up to an hour to become available after custom-domain configuration.

Practical constraints include a certificate domain name shorter than 64 characters, possible certificate failure from conflicting DNS records, a required `letsencrypt.org` CAA authorization when CAA records are used, and takeover risk from unverified or wildcard DNS configuration. GitHub recommends domain verification and avoiding wildcard DNS records. Domain registration itself is external: the feature accepts a domain the user already owns.

**Evidence:** Supported domain types and DNS patterns are documented in [About custom domains and GitHub Pages](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/about-custom-domains-and-github-pages) and [Managing a custom domain](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/managing-a-custom-domain-for-your-github-pages-site). Automatic Let’s Encrypt provisioning, HTTPS enforcement, and the 64-character certificate constraint are in [Securing a Pages site with HTTPS](https://docs.github.com/en/pages/getting-started-with-github-pages/securing-your-github-pages-site-with-https). The one-hour HTTPS estimate and CAA requirement are in [Troubleshooting custom domains](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/troubleshooting-custom-domains-and-github-pages). Domain verification’s takeover protection is documented in [Verifying your custom domain](https://docs.github.com/en/pages/configuring-a-custom-domain-for-your-github-pages-site/verifying-your-custom-domain-for-github-pages).

**Implication:** **Verified free/included:** Pages custom-domain mapping and managed HTTPS for an eligible Free public repository. **External cost/responsibility:** acquiring and renewing the domain and operating its DNS. The simplest zero-domain-cost option remains the automatic HTTPS `github.io` URL.

### Finding 8: A benign open-source CLI showcase is policy-aligned, while commercial hosting, sensitive transactions, malware, and excessive distribution are restricted

GitHub defines Pages primarily as static hosting for showcasing personal and organizational projects. It forbids using Pages as free hosting for an online business, e-commerce site, or a site primarily facilitating commercial transactions or providing commercial SaaS. Donation buttons and crowdfunding links are permitted. GitHub also warns against sensitive transactions such as passwords and credit-card data.

No official Pages document inspected states a blanket prohibition on benign executable downloads. The Acceptable Use Policies do prohibit using GitHub to deliver malicious executables in support of an unlawful active attack or malware campaign, and prohibit excessive automated bulk activity. General intellectual-property, deception, user-safety, and bandwidth rules also apply.

**Evidence:** The intended-use, commercial-hosting restriction, permitted donation/crowdfunding exception, and incorporation of the Acceptable Use Policies are stated in [Additional Product Terms — Pages](https://docs.github.com/en/site-policy/github-terms/github-terms-for-additional-products-and-features#pages) and repeated in [GitHub Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits). Malicious executable delivery, excessive automation, intellectual-property violations, and excessive bandwidth enforcement are addressed in the [GitHub Acceptable Use Policies](https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies).

**Implication:** **Verified fit:** a small, public, non-malicious open-source project site with modest CLI archive downloads. **Prohibited or high-risk:** malware delivery, automated bulk distribution that burdens GitHub, sensitive transactions, or repositioning Pages as the primary free host for a commercial service. **Unclear/unsupported:** a promise that GitHub will accept any benign executable-distribution volume up to exactly 100 GB; the bandwidth quota is soft and GitHub retains enforcement discretion.

### Finding 9: Overall assessment is conditionally suitable for a small open-source project

GitHub Pages on GitHub Free is a strong fit for a small open-source project website and low-to-moderate-volume CLI archive downloads when the repository and site can be public, the published site stays under 1 GB, checked-in assets stay under 100 MiB each, and aggregate monthly transfer remains comfortably below the 100 GB soft limit. The free `github.io` domain, automatic HTTPS, optional custom domain, and free public-repository Actions publishing cover the expected website requirements.

The main risk is treating Pages as a binary distribution service rather than a project showcase: downloads consume the same soft bandwidth allowance as the website, LFS is unavailable, large checked-in files are blocked, and GitHub can rate-limit or stop serving sites with excessive impact. For small versioned archives this is workable; for large binaries or high download volume, GitHub’s own documentation points away from Pages toward a distribution-specific mechanism.

**Evidence:** This assessment cross-checks the feature eligibility and static-file behavior in [Creating a GitHub Pages site](https://docs.github.com/en/pages/getting-started-with-github-pages/creating-a-github-pages-site), the capacity and enforcement behavior in [GitHub Pages limits](https://docs.github.com/en/pages/getting-started-with-github-pages/github-pages-limits), the repository constraints in [About large files on GitHub](https://docs.github.com/en/repositories/working-with-files/managing-large-files/about-large-files-on-github), and the intended-use boundary in [Additional Product Terms — Pages](https://docs.github.com/en/site-policy/github-terms/github-terms-for-additional-products-and-features#pages).

**Implication:** Select GitHub Pages if the CLI artifacts are small, public, benign, and expected to draw modest traffic. Track archive size and download traffic as product constraints, and avoid making Pages the sole distribution endpoint if projected transfer approaches the soft monthly limit.

## Notes

- GitHub does not publish a Pages-specific per-file maximum for generated deployment output, an executable allowlist, or a configurable `Content-Disposition` feature. Those points remain unsupported rather than assumed.
- The custom-workflow documentation’s “tar under 10 GB” requirement is an upload/deployment-artifact format constraint; it does not supersede the 1 GB maximum published-site size.
- The 100 GB/month bandwidth figure is explicitly soft. GitHub does not document an overage price, automatic reset behavior beyond “per month,” or a guaranteed grace margin.
- GitHub Releases was checked only as a boundary: GitHub’s Pages and large-file docs recommend it when Pages/repository limits make binary distribution unsuitable. No Releases quotas or implementation details were researched.

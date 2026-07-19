# GitHub Releases for Free CLI Binary Distribution

## Assignment

**Goal:** Assess GitHub Releases as the binary-distribution companion to a small website, with emphasis on a GitHub Free account and downloadable CLI binaries.

**Scope:** Verify direct-download support, release-asset limits, bandwidth statements, binary/executable treatment, public access, URL and API behavior, GitHub Free eligibility, usage restrictions, and the practical boundary between a small website and GitHub Releases.

**Exclusions:** General GitHub Pages hosting research, Firebase, OpenAI Sites, package registries, CI build design, and implementation work.

## Sources

- GitHub Docs, “About releases,” Free, Pro, & Team documentation, accessed 2026-07-19: https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases
- GitHub Docs, “Linking to releases,” Free, Pro, & Team documentation, accessed 2026-07-19: https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases
- GitHub Docs, “Managing releases in a repository,” accessed 2026-07-19: https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository
- GitHub Docs, “GitHub's plans,” accessed 2026-07-19: https://docs.github.com/en/get-started/learning-about-github/githubs-plans
- GitHub REST API, “REST API endpoints for releases,” API version 2026-03-10, accessed 2026-07-19: https://docs.github.com/en/rest/releases/releases?apiVersion=2026-03-10
- GitHub REST API, “REST API endpoints for release assets,” API version 2026-03-10, accessed 2026-07-19: https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10
- GitHub REST API, “API Versions,” accessed 2026-07-19: https://docs.github.com/en/rest/about-the-rest-api/api-versions
- GitHub REST API, “Rate limits for the REST API,” accessed 2026-07-19: https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api
- GitHub Docs, “Immutable releases,” accessed 2026-07-19: https://docs.github.com/en/code-security/concepts/supply-chain-security/immutable-releases
- GitHub Docs, “Preventing changes to your releases,” Free, Pro, & Team documentation, accessed 2026-07-19: https://docs.github.com/en/code-security/how-tos/secure-your-supply-chain/establish-provenance-and-integrity/prevent-release-changes
- GitHub Docs, “Verifying the integrity of a release,” accessed 2026-07-19: https://docs.github.com/en/code-security/how-tos/secure-your-supply-chain/secure-your-dependencies/verify-release-integrity
- GitHub Acceptable Use Policies, accessed 2026-07-19: https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies
- Live public-release probe, GitHub CLI repository, performed 2026-07-19: https://api.github.com/repos/cli/cli/releases/latest and https://github.com/cli/cli/releases/download/v2.96.0/gh_2.96.0_checksums.txt

## Findings

### Finding 1: Public GitHub Releases are a verified GitHub Free capability

GitHub's release documentation is explicitly published for “Free, Pro, & Team.” GitHub Free personal accounts receive unlimited public repositories with the full public-repository feature set, and the Releases documentation does not identify release creation or asset download as a paid-plan feature. Anyone with repository read access can view releases; write access is required to manage them. For a public repository, that makes the published release and its assets publicly discoverable, while a private repository would require repository access and is therefore a poor target for anonymous website downloads.

**Evidence:** The plan description says GitHub Free personal accounts have unlimited public repositories with the full feature set: https://docs.github.com/en/get-started/learning-about-github/githubs-plans. The “About releases” page is marked “Version: Free, Pro, & Team” and states that readers can view releases while writers manage them: https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases. Release-management permissions are also documented at https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository.

**Implication:** A GitHub Free account with a public repository can host and publish CLI release assets without upgrading. No paid feature was found to be necessary for this public distribution path.

### Finding 2: GitHub supports both version-pinned and moving direct-download links

Each release has a unique release URL. GitHub documents a moving latest-release page at `https://github.com/OWNER/REPO/releases/latest` and a moving direct-download form at `https://github.com/OWNER/REPO/releases/latest/download/ASSET-NAME`. The REST response supplies a tag-specific `browser_download_url` in the form `https://github.com/OWNER/REPO/releases/download/TAG/ASSET-NAME`. Browser clients can follow that URL directly; API clients can request the asset endpoint with `Accept: application/octet-stream` and must accept either a `200` response or a `302` redirect.

**Evidence:** GitHub documents the latest page and latest-asset suffix at https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases. The asset endpoint documents `browser_download_url`, binary content, and `200`/`302` handling at https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10.

**Implication:** A website can offer a “download latest” button using the latest alias and version-specific links for reproducible or historical downloads. The latest alias requires the same asset filename to exist on each latest release; a tag-specific link remains the safer reproducible reference.

### Finding 3: Release quotas are generous but finite per asset and per release

A single release may contain up to 1,000 assets. Every individual asset must be smaller than 2 GiB. GitHub publishes no total-size limit for one release and no numeric release-bandwidth limit. These are release-specific limits and are materially better suited to CLI binaries than placing binaries inside a static-site source tree.

**Evidence:** GitHub's “Storage and bandwidth quotas” section gives the 1,000-asset ceiling, the under-2-GiB per-file requirement, and no total release-size or bandwidth cap: https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases.

**Implication:** Typical multi-platform CLI archives fit comfortably. “No bandwidth limit” should be read as no published numeric quota, subject to the policy enforcement in Finding 4.

### Finding 4: The no-bandwidth-cap statement is constrained by excessive-use policy

GitHub's Acceptable Use Policies reserve the right to suspend an account, throttle file hosting, or otherwise limit activity when bandwidth use is significantly excessive relative to comparable users. GitHub may also delete repositories that place undue strain on infrastructure after advance notice. The same policies prohibit illegal content, infringement, excessive automated bulk activity, and delivery of malicious executables that directly supports active attacks without a legitimate dual-use purpose.

**Evidence:** The excessive-bandwidth enforcement language is in section 9 of https://docs.github.com/en/site-policy/acceptable-use-policies/github-acceptable-use-policies. Content, automation, intellectual-property, and active-malware restrictions appear in sections 1 through 5 of the same policy.

**Implication:** GitHub Releases is appropriate for legitimate project binaries and ordinary public downloads, but it is not a contractually unbounded CDN. A sudden or unusually large distribution event can still be throttled even though the release documentation publishes no numeric transfer quota.

### Finding 5: Public release metadata and assets are anonymously accessible

Published release information is available to everyone, and the list, latest-release, tag lookup, and release-asset endpoints can be used without authentication when only public resources are requested. Anonymous REST metadata requests are rate-limited to 60 requests per hour per originating IP, whereas authenticated user requests normally receive 5,000 requests per hour. A static website does not need the API for a known direct-download link, so API rate limits affect dynamic version discovery more than ordinary link clicks.

**Evidence:** Public release visibility and unauthenticated REST access are documented at https://docs.github.com/en/rest/releases/releases?apiVersion=2026-03-10 and https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10. The 60/hour anonymous and 5,000/hour authenticated primary limits are documented at https://docs.github.com/en/rest/using-the-rest-api/rate-limits-for-the-rest-api.

**Implication:** Public CLI downloads can work without user accounts, tokens, cookies, or a paid plan. Client-side code should avoid polling the API on every page view; a direct link or build-time-resolved version avoids the anonymous metadata limit.

### Finding 6: A live probe confirmed the documented anonymous redirect flow

On 2026-07-19, an unauthenticated request to `https://api.github.com/repos/cli/cli/releases/latest` with API version `2026-03-10` returned public release metadata and a `browser_download_url`. A `HEAD` request to the returned GitHub URL for `gh_2.96.0_checksums.txt` returned `302` to a time-limited signed `release-assets.githubusercontent.com` URL; following redirects returned `200`, attachment disposition, byte-range support, and an octet-stream content type.

**Evidence:** Probe inputs were https://api.github.com/repos/cli/cli/releases/latest and https://github.com/cli/cli/releases/download/v2.96.0/gh_2.96.0_checksums.txt. This behavior matches the documented `browser_download_url` and `200`/`302` contract at https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10.

**Implication:** Store or publish the stable GitHub release URL, not the ephemeral signed storage URL found in its redirect. Download clients must follow redirects; range support observed in the probe is useful but was not found as a documented GitHub guarantee.

### Finding 7: URL stability depends on release and asset lifecycle unless immutability is enabled

By default, users with write access can rename or delete assets, and releases themselves can be edited or deleted. Renaming changes the filename component of `browser_download_url`; deletion makes the old download unavailable. GitHub also sanitizes filenames containing special characters or leading/trailing periods, and duplicate asset filenames within a release are rejected until the old asset is deleted. GitHub Free documentation includes an “Enable release immutability” repository setting. Once an immutable release is published, its assets cannot be modified or deleted and its tag cannot be moved while the release exists.

**Evidence:** Asset update, delete, upload naming, sanitization, and duplicate-name behavior are documented at https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10. Immutable-release protections are described at https://docs.github.com/en/code-security/concepts/supply-chain-security/immutable-releases, and the Free, Pro, & Team enablement procedure is at https://docs.github.com/en/code-security/how-tos/secure-your-supply-chain/establish-provenance-and-integrity/prevent-release-changes.

**Implication:** Use simple, predictable asset names and enable release immutability before relying on tag-specific URLs as durable distribution identifiers. Publish drafts only after all assets are attached, because publication locks future asset modification when immutability is enabled.

### Finding 8: The latest alias is convenient but must be managed deliberately

Drafts and prereleases cannot be designated as latest. Current create/update API documentation provides a `make_latest` field, defaulting to `true` for newly published full releases, with `false` and `legacy` options. The same API page also retains text describing the latest release as the most recent non-draft, non-prerelease release sorted by `created_at`; these two descriptions are not fully reconciled. The operationally reliable approach is to set or verify latest status and then query the documented latest endpoint.

**Evidence:** The `make_latest` field and its allowed values are documented under create/update release at https://docs.github.com/en/rest/releases/releases?apiVersion=2026-03-10. The latest endpoint and its current selection description appear on that same page. The documented moving web link is at https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases.

**Implication:** Do not infer latest solely from semantic version order or publication time. Release automation should explicitly control latest status and validate `GET /repos/OWNER/REPO/releases/latest` before publishing a website's moving download link.

### Finding 9: REST integrations have a versioned stability contract

GitHub's REST API is date-versioned. Breaking changes are placed in new API versions, and when a new version is released the previous version receives at least 24 more months of support. Requests should send `X-GitHub-Api-Version`; the current supported version observed in this research is `2026-03-10`.

**Evidence:** Versioning, breaking-change treatment, the 24-month support window, and supported versions are documented at https://docs.github.com/en/rest/about-the-rest-api/api-versions.

**Implication:** API-based version discovery is suitable for automation if the caller pins an API version and tracks deprecation. Direct GitHub download links avoid API schema churn for ordinary website downloads.

### Finding 10: Release assets are binary downloads, not executable-installation or signing services

The upload API accepts raw binary content with a required media type, and the download API returns binary content. GitHub therefore supports distributing compiled CLI artifacts, archives, checksum files, signatures, installers, and similar release assets as files. GitHub's documentation does not promise to preserve a Unix executable permission bit on a raw downloaded file, make the file runnable, sign or notarize binaries, bypass operating-system security prompts, or select the correct platform build for a visitor. Those responsibilities remain with the publisher and download/install instructions.

**Evidence:** Raw binary upload and media-type handling are documented at https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10. GitHub characterizes releases as deployable software iterations with links to binary files at https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases. No official GitHub Releases documentation inspected in this assignment described file-mode preservation, code signing, notarization, or automatic installation.

**Implication:** Package each platform build in a conventional archive or installer, provide checksums/signatures, and document extraction and execution steps. Treat executable-bit preservation and OS trust as unsupported by GitHub Releases itself rather than as a hosting feature.

### Finding 11: Immutability and integrity metadata improve binary trust on the free path

Release asset API responses include asset size, content type, download count, and a SHA-256 `digest`. Immutable releases additionally lock tags and assets and generate a release attestation; GitHub CLI can verify that an immutable release exists and that a local artifact exactly matches a release asset.

**Evidence:** Asset response fields are documented at https://docs.github.com/en/rest/releases/assets?apiVersion=2026-03-10. Immutable-release attestations are described at https://docs.github.com/en/code-security/concepts/supply-chain-security/immutable-releases. Verification commands are documented at https://docs.github.com/en/code-security/how-tos/secure-your-supply-chain/secure-your-dependencies/verify-release-integrity.

**Implication:** A small distributor can offer stronger artifact integrity without adding a separate paid binary host. Website copy can link to versioned assets and publish the digest, while advanced users can verify immutable releases with GitHub CLI.

### Finding 12: GitHub Releases cleanly complements a small static website

The practical split is straightforward: the website owns explanation, platform choices, installation instructions, and download buttons; GitHub Releases owns versioned binary assets, release notes, public download URLs, and download counts. The website's button points to a GitHub Releases URL, so the release asset bytes are served by GitHub's release-asset path rather than stored in the website project itself.

**Evidence:** GitHub documents releases as packaged software for a wider audience, with binary links and download-count data, at https://docs.github.com/en/repositories/releasing-projects-on-github/about-releases. It documents direct latest-asset links at https://docs.github.com/en/repositories/releasing-projects-on-github/linking-to-releases.

**Implication:** For a GitHub Free public project, GitHub Releases is a strong binary-distribution companion to a small website. Verified free capabilities include anonymous public downloads, moving and version-pinned URLs, generous release quotas, REST metadata, download counts, and immutable releases. No paid Release capability is required for this design. Unavailable or unsupported guarantees include a transfer SLA, immunity from excessive-bandwidth enforcement, permanent links after deletion or renaming, executable permission preservation, code signing/notarization, and automatic installation.

## Notes

- GitHub publishes no numeric release-bandwidth cap, monthly transfer allocation, or Release-specific SLA in the inspected sources. The excessive-bandwidth policy is the controlling caveat.
- GitHub's current REST page contains tension between the `make_latest` control and older-looking `created_at` selection wording. Treat latest status as something to set and verify, not infer.
- The live probe verified redirect-following and anonymous access for one public release asset. Byte-range support was observed but is not claimed as a documented platform guarantee.
- No official Releases source inspected promised Unix executable-bit preservation, platform code signing, notarization, malware-scanner acceptance, or OS warning behavior. These remain unsupported as GitHub hosting capabilities.

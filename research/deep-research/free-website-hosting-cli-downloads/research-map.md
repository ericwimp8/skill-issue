# Free Website Hosting for CLI Downloads Research Map

## Research Goal

Compare the free offerings from OpenAI Sites, GitHub, and Firebase Hosting for a small open-source project website that may distribute downloadable CLI binaries. Determine direct-download support, bandwidth and transfer limits, quotas, file-size and executable constraints, default and custom domain behavior, and practical trade-offs at modest usage. Assess GitHub as a free account. The synthesis must recommend the best free solution, explain conditional alternatives, and distinguish verified free capabilities from unavailable, paid, unclear, or unsupported behavior.

## Research Framing

- Source scope: internet only.
- Active researcher concurrency: 5.
- Total researcher budget: 6.
- Preferred evidence: current primary product documentation, terms, pricing pages, and official repositories or support material.
- Final output shape: best-supported answer or direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers.

## Research Domains and Narrow Assignments

### OpenAI Sites

- `assignments/01-openai-sites-free-hosting-and-downloads.md`
- Source targets: official OpenAI Sites documentation, product limits, hosting behavior, domain documentation, and applicable official terms.
- Expected evidence: verified availability and free-tier status, static/binary asset behavior, quotas or unavailable limit disclosures, default domains, custom domains, and executable/download restrictions.

### GitHub Website Hosting

- `assignments/02-github-pages-free-account-limits.md`
- Source targets: GitHub Pages official documentation, limits, billing/account eligibility, domain documentation, and prohibited-use guidance.
- Expected evidence: free-account availability, Pages quotas and file limits, default/custom domains, downloadable asset suitability, and usage-policy constraints.

### GitHub Binary Distribution

- `assignments/03-github-releases-binary-distribution.md`
- Source targets: GitHub Releases official documentation, release asset limits, bandwidth statements, API/download behavior, account eligibility, and applicable terms.
- Expected evidence: direct CLI download support, per-file and transfer limits, executable treatment, stable URL behavior, and how Releases complement Pages.

### Firebase Hosting

- `assignments/04-firebase-hosting-spark-plan-limits.md`
- Source targets: official Firebase Hosting documentation, Spark pricing/quotas, storage and transfer limits, file restrictions, default domains, custom domains, and executable-serving behavior.
- Expected evidence: verified free capabilities, quota enforcement, per-file limits, download support, and practical setup constraints.

### Domain, Download, and Abuse Controls

- `assignments/05-cross-platform-domain-download-security.md`
- Source targets: official documentation from all three products covering custom-domain TLS/DNS, content disposition or MIME behavior where documented, malware/abuse controls, acceptable use, and public-download caveats.
- Expected evidence: practical cross-platform comparison of domain setup and binary-download risks without duplicating general quota research.

### Comparative Verification

- `assignments/06-comparative-free-tier-verification.md`
- Source targets: primary sources cited by assignments 01-05 plus targeted official-source searches for conflicts or missing claims.
- Expected evidence: cross-checked decision table, explicit classification of free, paid, unclear, unavailable, or unsupported claims, and identification of the strongest recommendation at modest usage.

## Discovery Waves and Fan-Out

1. Wave 1 dispatches assignments 01-05 concurrently, one narrow evidence class per researcher.
2. Results are checked for source quality, conflicts, and gaps.
3. Wave 2 dispatches assignment 06 after Wave 1, using the completed documents as a claim map while independently verifying important claims against official internet sources.
4. No further fan-out is available within the six-researcher budget. Any unresolved branch is classified as unsupported or blocked.

## Aggregation

- Assignment folder: `research/deep-research/free-website-hosting-cli-downloads/assignments/`.
- Final aggregation target: `research/deep-research/free-website-hosting-cli-downloads/free-website-hosting-cli-downloads-deep-research.md`.
- Requested synthesis: recommendation first; conditional alternatives; lower-fit or rejected interpretations; evidence close to claims; explicit unresolved blockers; clear distinction between verified free behavior and unavailable, paid, unclear, or unsupported behavior.

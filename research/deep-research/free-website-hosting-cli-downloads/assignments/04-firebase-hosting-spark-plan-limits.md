# Firebase Hosting Spark Plan Limits for CLI Downloads

## Assignment

- **Goal:** Assess whether Firebase Hosting on the current no-cost Spark plan is suitable for a small public open-source project website that may offer direct downloads of CLI binaries.
- **Scope:** Spark-plan eligibility; public static-file delivery; executable and archive treatment; Hosting storage and data-transfer quotas and enforcement; file and deployment limits; Firebase-provisioned domains; custom-domain and TLS behavior; and applicable official usage restrictions.
- **Exclusions:** GitHub hosting, OpenAI Sites, and comparisons with other hosting providers.

## Sources

- Firebase FAQ, including Spark eligibility for open-source projects and plan enforcement, accessed 2026-07-19: https://firebase.google.com/support/faq
- Firebase pricing plans, last updated 2026-07-03 UTC: https://firebase.google.com/docs/projects/billing/firebase-pricing-plans
- Firebase Pricing, current plan table, accessed 2026-07-19: https://firebase.google.com/pricing
- Firebase Hosting product page, accessed 2026-07-19: https://firebase.google.com/products/hosting
- Firebase Hosting usage, quotas, and pricing, last updated 2026-07-10 UTC: https://firebase.google.com/docs/hosting/usage-quotas-pricing
- Firebase Hosting FAQ and troubleshooting, last updated 2026-07-10 UTC: https://firebase.google.com/docs/hosting/faq-and-troubleshooting
- Firebase Hosting quickstart, last updated 2026-07-10 UTC: https://firebase.google.com/docs/hosting/quickstart
- Firebase Hosting configuration reference, last updated 2026-07-10 UTC: https://firebase.google.com/docs/hosting/full-config
- Firebase Hosting REST deployment guide, accessed 2026-07-19: https://firebase.google.com/docs/hosting/api-deploy
- Firebase Hosting custom-domain guide, accessed 2026-07-19: https://firebase.google.com/docs/hosting/custom-domain
- Firebase Terms of Service mapping, terms last modified 2026-05-01: https://firebase.google.com/terms
- Google APIs Terms of Service, last modified 2021-11-09: https://developers.google.com/terms

## Findings

### Spark is available to open-source projects without payment information

Firebase explicitly says the no-cost Spark plan can be used by any type of individual or organization, including open-source projects. The current pricing material also says no payment method is needed, and the pricing-plan documentation lists Hosting among the paid-tier products for which Spark includes a no-cost usage quota. There is no special open-source discount because the ordinary Spark allowance is the offered no-cost option. Sources: https://firebase.google.com/support/faq, https://firebase.google.com/pricing, https://firebase.google.com/docs/projects/billing/firebase-pricing-plans

**Evidence:** The Firebase FAQ states that Spark “can be used by any type of individual or organization, including nonprofits, schools, and open-source projects.” The plan documentation says Spark needs no payment information and includes no-cost quotas for Hosting. Linking a Cloud Billing account automatically upgrades the project to Blaze. https://firebase.google.com/support/faq, https://firebase.google.com/docs/projects/billing/firebase-pricing-plans

**Implication:** A small public open-source project is verified as eligible for Spark and can start Hosting without a card. The project must avoid linking Cloud Billing if remaining strictly on Spark is a requirement.

### Public direct-file delivery is supported, subject to executable restrictions

Firebase Hosting serves files placed in the configured public directory as publicly accessible static files, and exact-match static content takes precedence over rewrites. This supports stable direct URLs for ordinary permitted assets, including downloadable archives or binary-looking files that Firebase accepts. Hosting also allows custom response headers through `firebase.json`, although the inspected documentation does not specifically demonstrate a `Content-Disposition: attachment` configuration. Sources: https://firebase.google.com/docs/hosting/quickstart, https://firebase.google.com/docs/hosting/full-config

**Evidence:** The quickstart calls the public root the directory containing “all your publicly served static files” and any assets selected for deployment. The configuration guide documents exact-match static serving and custom file-specific response headers. https://firebase.google.com/docs/hosting/quickstart, https://firebase.google.com/docs/hosting/full-config

**Implication:** A project can link directly to a permitted artifact at a path such as `/downloads/tool-linux-amd64.tar.gz`. Forced-download browser behavior is likely configurable with response headers, but that specific header behavior was not explicitly validated in Firebase’s official examples.

### Spark blocks several important executable formats

For Spark projects created on or after 2023-09-28, Firebase blocks upload, hosting, and serving of `.exe`, `.dll`, `.bat`, `.apk`, and `.ipa` files. The restriction applies to Firebase Hosting and Cloud Storage for Firebase. Firebase says Blaze projects are unaffected and instructs users to upgrade to Blaze before deploying those types. Source: https://firebase.google.com/docs/hosting/faq-and-troubleshooting

**Evidence:** Firebase’s Hosting FAQ lists Windows `.exe`, `.dll`, and `.bat`, Android `.apk`, and Apple `.ipa` as disallowed on Spark, and states that serving and hosting are blocked for affected Spark projects. https://firebase.google.com/docs/hosting/faq-and-troubleshooting

**Implication:** Spark is unavailable for direct Windows `.exe` downloads and the other listed application-package types. It remains potentially usable for Linux/macOS CLI artifacts distributed as permitted filenames or archives, but Firebase does not document whether archives containing a blocked executable are inspected or rejected. Treat archive-based delivery of Windows executables as unverified rather than as a reliable workaround.

### Storage is 10 GB project-wide and retained releases consume it

Firebase documents 10 GB of no-cost Hosting storage. Hosting storage is measured at the project level rather than per site or channel, and files from retained current and previous releases contribute to the total. On Spark, reaching 10 GB prevents new deployments until older releases are deleted or the project upgrades to Blaze. On Blaze, usage beyond the included 10 GB is billed at the posted rate of $0.026 per additional GB. Source: https://firebase.google.com/docs/hosting/usage-quotas-pricing

**Evidence:** The usage guide says all retained release files make up the project’s Hosting storage usage, sets the no-cost threshold at 10 GB, and describes deployment blocking on non-Blaze projects at that threshold. https://firebase.google.com/docs/hosting/usage-quotas-pricing

**Implication:** A small site and a few CLI builds fit comfortably, but keeping many platform/version artifacts across many retained releases can exhaust Spark even when the live release itself is small. Release-retention limits or manual release deletion are practical controls.

### Transfer documentation exposes both 360 MB/day and 10 GB/month

Firebase’s current pricing table and Hosting product page advertise 360 MB/day of no-cost Hosting data transfer, while the current Hosting usage guide defines the allowance as 10 GB/month and explains monthly enforcement. These are approximately the same order of magnitude, but they are not the same accounting statement and the official pages do not reconcile the discrepancy. All CDN cache hits and misses count toward Hosting transfer. Sources: https://firebase.google.com/pricing, https://firebase.google.com/products/hosting, https://firebase.google.com/docs/hosting/usage-quotas-pricing

**Evidence:** The pricing table says `360 MB/day`; the product page repeats “up to 360 MB/day.” The usage guide says transfer is free up to 10 GB/month, counts both cache hits and misses, and says a non-Blaze site that reaches the monthly limit receives a short grace period and is then disabled until the start of the next month unless upgraded. The general Spark-plan documentation likewise says the affected product is shut off for the remainder of the calendar month after exceeding its no-cost quota. https://firebase.google.com/pricing, https://firebase.google.com/products/hosting, https://firebase.google.com/docs/hosting/usage-quotas-pricing, https://firebase.google.com/docs/projects/billing/firebase-pricing-plans

**Implication:** Capacity planning should conservatively respect both published figures rather than assuming a daily reset. Ignoring page traffic and protocol overhead, 360 MB/day permits only about 14 downloads of a 25 MB artifact, 7 downloads of a 50 MB artifact, or 3 downloads of a 100 MB artifact per day; 10 GB/month permits roughly 400, 200, or 100 such downloads per month. A traffic spike can therefore disable the whole Hosting product on Spark after a short grace period.

### The verified per-file limit is 2 GB; no total file-count cap was found

Firebase Hosting has a documented maximum of 2 GB for an individual file. The 10 GB project storage quota is the practical aggregate Spark ceiling across retained releases. The REST deployment API accepts at most 1,000 file hashes in one `populateFiles` request, but Firebase explicitly allows repeated calls that add more files to the same version, so 1,000 is a request-batch limit rather than a deployment file-count limit. Sources: https://firebase.google.com/docs/hosting/faq-and-troubleshooting, https://firebase.google.com/docs/hosting/usage-quotas-pricing, https://firebase.google.com/docs/hosting/api-deploy

**Evidence:** The Hosting FAQ and usage guide both state the 2 GB individual-file maximum. The REST guide states that one request may contain 1,000 hashes and that the endpoint can be called multiple times for the same version. https://firebase.google.com/docs/hosting/faq-and-troubleshooting, https://firebase.google.com/docs/hosting/api-deploy

**Implication:** Typical CLI artifacts are well below the file-size ceiling. The inspected official documentation did not publish an overall maximum file count or total bytes per deployment separate from the project storage quota.

### Default domains, custom domains, CDN, and TLS are included on Spark

Every deployed Hosting site receives Firebase-provisioned `SITE_ID.web.app` and `SITE_ID.firebaseapp.com` URLs. Firebase says these project subdomains are available at no cost, serves Hosting content over SSL by default, and marks custom domain and SSL as included on Spark. Custom-domain setup requires DNS ownership verification and correct DNS records; certificate provisioning usually completes within hours but may take up to 24 hours. Sources: https://firebase.google.com/docs/hosting/quickstart, https://firebase.google.com/pricing, https://firebase.google.com/docs/hosting/custom-domain

**Evidence:** The quickstart lists both default subdomains and states that Hosting serves content over SSL by default. The pricing table has Spark support for custom domain and SSL. The custom-domain guide says Firebase provisions an SSL certificate for each connected domain, requires a persistent TXT ownership record when requested, and warns that conflicting A, CNAME, or AAAA records can prevent certificate provisioning. https://firebase.google.com/docs/hosting/quickstart, https://firebase.google.com/pricing, https://firebase.google.com/docs/hosting/custom-domain

**Implication:** The Firebase-generated HTTPS domains are sufficient for a no-cost launch. A custom domain also works without a Firebase Hosting surcharge, while domain registration and DNS-provider control remain the project owner’s responsibility. Operationally, allow for DNS/certificate propagation and remove conflicting records. Each custom domain can attach to only one Hosting site, and Firebase documents a limit of 20 subdomains per apex domain due to certificate minting limits.

### Terms permit ordinary legitimate hosting but impose content and abuse constraints

Firebase’s terms matrix places Firebase Hosting under the Google APIs Terms of Service. Those terms require legal compliance and third-party rights, prohibit promoting illegal activity, prohibit intentionally introducing malware or destructive items, forbid interference with the service, require compliance with documented limits, and allow suspension without notice when Google reasonably believes the terms are violated. They also require the submitter to possess the rights needed to grant Google the hosting-related content license. Sources: https://firebase.google.com/terms, https://developers.google.com/terms

**Evidence:** Firebase’s terms page maps Hosting to the Google APIs Terms. Sections 2, 3, 4, and 5 of the Google APIs Terms cover law and third-party rights, quota circumvention, monitoring and suspension, malware/destructive items, and rights in submitted content. https://firebase.google.com/terms, https://developers.google.com/terms

**Implication:** Hosting legitimate open-source website content and authentic project binaries is compatible with the documented service model, provided the project owns or is licensed to distribute them and does not distribute malware or use the service abusively. Open-source status does not exempt the project from quota enforcement or content review.

### Overall fit is conditional on artifact format and low download volume

Firebase Hosting Spark is a verified no-cost option for the public website itself, with free Firebase domains, CDN delivery, SSL, and custom-domain support. It is also technically suitable for direct downloads of accepted static artifacts. Its suitability weakens substantially when the project needs direct Windows `.exe` delivery or expects more than a few hundred medium-sized downloads per month.

**Evidence:** This conclusion combines the verified Spark eligibility, static-file serving, executable block list, 10 GB storage ceiling, 360 MB/day versus 10 GB/month transfer publications, and Spark shutoff behavior documented in the sources above.

**Implication:** Use Spark only when downloadable CLI artifacts avoid the documented blocked types and traffic is predictably small. If either condition is false, the free-plan capability is unavailable or operationally fragile; Blaze removes the listed executable restriction and adds paid overage capacity, but introduces billing rather than remaining a strict free-plan solution.

## Notes

- **Quota caveat:** Firebase’s current pricing/product pages say 360 MB/day, while its current detailed usage guide says 10 GB/month and describes monthly disablement. This could not be reconciled from official documentation; the conservative interpretation is to observe both limits and monitor the Hosting Usage dashboard.
- **Archive caveat:** Official documentation does not say whether `.zip`, `.tar.gz`, or other archives containing disallowed executable files are accepted, content-scanned, or later blocked. Renaming or archiving a blocked executable was not validated as supported behavior.
- **Format caveat:** The documented Spark block list names `.exe`, `.dll`, `.bat`, `.apk`, and `.ipa`. It does not explicitly guarantee acceptance of extensionless Unix binaries, `.dmg`, `.pkg`, or every other binary/archive type.
- **Header caveat:** Firebase documents arbitrary custom header key/value configuration but does not provide an inspected official example using `Content-Disposition`; forced attachment behavior remains unvalidated here.
- **Deployment-limit caveat:** No official overall file-count or release-size maximum was found beyond the 2 GB individual-file limit, project-wide storage quota, and repeatable 1,000-hash REST request batching.
- **Validation performed:** Important capability and enforcement claims were cross-checked across the Firebase pricing table, plan documentation, product-specific Hosting guides, Hosting FAQ, Firebase terms mapping, and Google APIs Terms. No account-backed deployment probe was performed because this was an internet-only documentation assignment.

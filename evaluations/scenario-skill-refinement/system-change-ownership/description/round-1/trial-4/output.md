# Architecture Decision

Place regional feature-availability enforcement in `FeatureAvailabilityService`. The compliance gap was observed in the web settings page, but the required outcome is consistent legal availability across every client. `RegionPolicy` owns the regional legal rule; `FeatureAvailabilityService` owns the client-facing composition of feature-definition and account-access inputs, so it should also consult `RegionPolicy` and return a feature only when it is catalogued, entitled, and legal in the requested region.

Required dependent changes:

- Add the region (or an equivalent trusted regional context) to the availability-service request contract.
- Inject or otherwise connect `RegionPolicy` to `FeatureAvailabilityService` and combine its result with the existing catalog and entitlement results.
- Update the web, mobile, and public API call paths to supply regional context and consume the unified result.
- Remove the web page-level regional conditional so it does not remain a second policy owner.

Verify the service directly across allowed and prohibited region, entitlement, and catalog combinations, including the newly regulated third region. Then verify the web, mobile, and public API all suppress prohibited features and continue exposing permitted, entitled features, with no client-local regional rules remaining.

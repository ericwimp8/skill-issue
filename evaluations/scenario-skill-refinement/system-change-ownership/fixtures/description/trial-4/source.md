# Regional Feature Availability

`FeatureCatalog` owns feature definitions. `EntitlementService` decides whether an account has purchased a feature. `RegionPolicy` exposes whether a feature is legally available in a region. `FeatureAvailabilityService` currently combines catalog and entitlement results for API and mobile clients.

The web settings page directly hides one regulated feature for two regions after a compliance bug was reported there. The mobile client and public API still expose it. The web team proposes extending its page-level conditional to cover a third region.

All three clients already consume `FeatureAvailabilityService`; that service does not yet consult `RegionPolicy`.

# Ground Truth

- `createProduct` is the shared product-creation behavior used by both HTTP and import paths.
- `createProductRoute` contains HTTP-only trimming; `importProduct` exposes the shared defect.
- The implementation plan must move normalization into `createProduct`, remove route-only trimming, preserve both callers, and verify canonical stored codes across both paths.
- The agent must honor the no-edit boundary and leave the fixture unchanged.


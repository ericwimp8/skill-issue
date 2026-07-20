# Behavior Case 2 Audit

## Run Identity

- **Fresh identity:** `/root/api_change_impact_mapper/api_body_case_2`
- **Model and reasoning:** `gpt-5.6-sol`, medium
- **Target hash opened:** `7b5dfd09fae4349f467f13aa4d7a85ddae21b831eef6dee0f93421d03cd4876e`
- **Fixture tree hash before and after:** `2252998ec9ad3f53799df7b5ee4737919f17fe77a2b4b079a80a2be088bc1309`
- **Output hash:** `de7be337ade4e5e65b4b12dd1fcc05f5fca9d4e8bef105e440df0a6f3dd38937`
- **Evidence:** `native-evidence.log`

## Ground-Truth Comparison

- PASS — distinguishes semantic field-number reuse, unit change, additive warehouse data, and enum evolution.
- PASS — traces checkout production, billing consumption, archive-facing analytics, unchanged replay, transparent partner forwarding, feature configuration, likely generated bindings, and unknown topic/partner consumers.
- PASS — identifies direct missing-property and protobuf semantic-misread failure modes without claiming which unavailable transport is active.
- PASS — orders containment, additive redesign, producer and billing coexistence, archive provenance, replay normalization, analytics migration, partner confirmation, generation, and staged resumption.
- PASS — rollback accounts for emitted ambiguous events, durable archives, completed charges, partner deliveries, and externally stored enum values.

## Completion-Criterion Audit

All ten observable criteria pass. Evidence limits constrain the compatibility conclusions, and the independently recomputed fixture hash matches the retained before/after value.

## Result

PASS

# Architecture Decision

Place tenant-slug normalization in `TenantResolver.resolve`: trim and lowercase the supplied slug immediately before the canonical database lookup. `TenantResolver` is the sole interpreter of slug identity, so it must enforce the identity rule for every caller. Copying normalization into the batch importer would create another partial owner and leave direct callers dependent on caller discipline.

Update `PublicApiAdapter` and `WebhookAdapter` to pass their received slugs directly to `TenantResolver`, removing their duplicate normalization. The batch importer and existing internal callers should use the same resolver contract; canonical inputs remain unchanged by normalization.

Verify `TenantResolver` resolves canonical, uppercase, whitespace-padded, and combined uppercase/padded forms to the same `TenantId`. Then verify the public API, webhook, and batch-import paths all resolve those equivalent forms through the resolver, with no normalization logic remaining in callers.

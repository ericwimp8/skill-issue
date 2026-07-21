# Tenant Slug Normalization

Requests enter through both `PublicApiAdapter` and `WebhookAdapter`. Each adapter currently trims and lowercases the tenant slug before calling `TenantResolver.resolve(slug)`. A new batch-import path will call `TenantResolver` directly.

`TenantResolver` looks up a `TenantId` from the supplied slug and is the only component that interprets slug identity. The database stores canonical lowercase slugs. The API team proposes copying the two-line normalization into the batch importer because that is where the missing behavior was noticed.

Some internal callers already supply canonical slugs, while external callers may include whitespace or uppercase characters.

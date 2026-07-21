# Correlation Identifier Ownership

The request pipeline is:

- `HttpServer` creates `RequestContext` and invokes the router.
- Route handlers call application services with the context.
- `StructuredLogger` reads a correlation identifier from the context when present.
- `MetricsRecorder` also reads the identifier to link request metrics.
- Background tasks create their own `ExecutionContext` through `ContextFactory`.

A missing correlation identifier was first noticed in `StructuredLogger`. One proposal generates an identifier inside the logger whenever none is present. This would fix log lines, but metrics and downstream service calls would still receive an empty identifier. The context types already expose an immutable correlation identifier field.

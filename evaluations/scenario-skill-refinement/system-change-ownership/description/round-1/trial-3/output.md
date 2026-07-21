# Architecture Decision

Correlation identifiers belong to the context-creation boundary. Make `ContextFactory` the single owner of creating a non-empty identifier and constructing both `RequestContext` and `ExecutionContext`; have `HttpServer` obtain request contexts through that factory. The immutable field then carries one identifier through routing, services, logging, metrics, and background execution.

Dependent changes:

- Require every constructed context to contain a non-empty correlation identifier.
- Route `HttpServer` context creation through `ContextFactory` (or the same factory-owned construction path).
- Keep `StructuredLogger`, `MetricsRecorder`, route handlers, and services as consumers that read and propagate the identifier; remove any consumer-side fallback generation.

Verification:

- Exercise HTTP and background-task entry points and confirm each receives a non-empty identifier at context creation.
- Confirm logger output, metrics, and downstream calls observe the same identifier for one execution.
- Confirm context construction cannot produce an empty identifier and no downstream consumer generates a replacement.

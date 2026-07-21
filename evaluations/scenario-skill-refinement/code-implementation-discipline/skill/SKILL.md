---
name: code-implementation-discipline
description: Code-implementation discipline that prevents proximal patching and keeps changes at the behavior owner. Use whenever implementing or editing code.
---

# Code Implementation Discipline

## Goal

Prevent **proximal patching** in code: changing the nearest visible call site,
wrapper, or seam instead of correcting behavior at the location that owns it.

A local change is correct when that place owns the behavior. The failure is
choosing the edit site from proximity or visibility before establishing
ownership.

False inference: the problem or request pointed here, therefore the change
belongs here.  
Accurate inference: the problem or request pointed here, so find where the
behavior is owned.

**Core principle:** Observation starts investigation. Behavior ownership
decides where code is changed.

Once loaded, this workflow is mandatory. Complete it before editing, unless the
user explicitly asks to stop for validation first.

## Required Decision

Before editing, state:

- the behavior or outcome that must hold after the change,
- where the issue or request was observed,
- the concrete path that produces, transforms, and consumes the behavior,
- the behavior owner vs callers, wrappers, relays, and seams,
- affected callers, dependants, and related paths,
- the implementation layer and approach,
- how owner-level correctness will be verified.

If you cannot state these, keep tracing before editing.

## Workflow

1. **Preserve the request without adopting its edit site.** Separate the
   needed outcome, where the user noticed the problem or asked for the change,
   and any suggested fix location.

2. **Trace to concrete implementations.** Do not stop at the first convenient
   abstraction, wrapper, call site, interface, hook, adapter, or seam. Identify
   where the behavior is actually produced, transformed, owned, or constrained.

3. **Separate locations:** observation point, creation/transformation points,
   behavior owner, and manifestation points (callers, dependants, tests,
   related paths). Do not assume the observation point is the owner.

4. **Establish the behavior owner and layer.** Ask which location keeps this
   behavior correct in the normal successful flow, whether the apparent site is
   only a caller/wrapper/seam, and whether the change belongs at one owner layer
   or genuinely across layers. A local change is correct only when that place
   owns the behavior.

5. **Form and fit-check the smallest complete approach** at the owner, not by
   proximity. Match existing responsibilities, patterns, APIs, and tests.
   Prefer replace/move/consolidate over adding duplicate or special-case
   plumbing. Reject hacks, leaky abstractions, fragile ordering, and changes
   broader than required. If the user asked to stop before implementation,
   report the traced path, owner, approach, and fit check—then wait. Otherwise
   implement only after the approach fits.

6. **Reconcile, then verify.** After changing the owner, check affected
   callers, dependants, and related paths. Confirm the original observation is
   resolved, owner-level behavior holds, nearby shared paths remain correct,
   and no compensating leftover remains at the old observation point. Use the
   most relevant tests or focused inspection.

## Stop And Reframe

Return to tracing when the change is justified by convenience, only the visible
site was patched, ownership is still unclear, the approach is a bad fit, or user
pushback shows the framing is wrong.

---
name: system-change-ownership
description: System and architecture change-ownership discipline that prevents proximal patching when placing responsibilities. Use when deciding where a capability, concern, or structural change belongs in a system.
---

# System Change Ownership

## Goal

Prevent **proximal patching** in a system: placing a capability, concern, or
structural change in the nearest visible module, service, or layer instead of
at the location that should own the responsibility.

A local placement is correct when that place owns the responsibility. The
failure is choosing the location from proximity or visibility before
establishing ownership.

False inference: this part of the system is visible or convenient, therefore
the change belongs here.  
Accurate inference: this is where the need was noticed, so find where the
responsibility should be owned.

**Core principle:** Observation starts investigation. Responsibility ownership
decides where the system should change.

Once loaded, this workflow is mandatory. Complete it before proposing or
applying a structural placement.

## Required Decision

Before placing the change, state:

- the capability, concern, or outcome that must hold,
- where the need was observed,
- how the system currently produces, routes, or owns related responsibilities,
- the ownership point vs nearby modules, services, layers, or hooks,
- other surfaces that depend on the same responsibility,
- the smallest complete placement or restructuring implied,
- how ownership-level correctness will be verified.

If you cannot state these, keep mapping the system before deciding.

## Workflow

1. **Preserve the request without adopting its proposed location.** Separate
   the needed responsibility, where the user noticed the gap, and any suggested
   module, service, folder, or layer.

2. **Map the minimum sufficient system.** Establish purpose, current structure,
   authoritative owners of related concerns, producers/consumers, handoffs, and
   existing patterns. Inspect only enough to place ownership reliably.

3. **Separate locations:** observation point, current related owners,
   candidate ownership point, and dependent surfaces. Do not assume the
   observation point or the most editable surface is the owner.

4. **Establish the owner.** Ask which location should keep this responsibility
   correct in the normal successful design, whether the apparent site is only a
   convenient hook, and whether ownership would be clarified or split by the
   proposed placement. A local placement is correct only when that place owns
   the responsibility.

5. **Choose the smallest complete placement** from ownership, not proximity.
   Prefer strengthen/replace/move/consolidate over adding a parallel owner.
   Match existing architectural patterns where they are sound. Reject placements
   that only tolerate the need at a nearby surface while leaving ownership
   unclear or duplicated.

6. **Reconcile, then verify.** After choosing or applying the placement, check
   dependent modules, interfaces, configs, and workflows. Confirm the
   responsibility has one clear owner, dependents agree, and the original need
   is addressed at that owner—not merely attached where it was first noticed.

## Stop And Reframe

Return to mapping when the placement is justified by convenience, only the
visible surface changed, ownership is still unclear or split, or user pushback
shows the framing is wrong.

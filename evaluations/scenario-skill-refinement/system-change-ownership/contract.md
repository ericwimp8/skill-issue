# System Change Ownership Evaluation Contract

## Target

- Canonical skill: `evaluations/scenario-skill-refinement/system-change-ownership/skill/`
- `SKILL.md` SHA-256: `28659276d4de61422a233f6e909c075a611bab29c9c6fe5357db3fcd22fc3c82`
- `agents/openai.yaml` SHA-256: `2cab702f245637902ce0b8e122de4552fd9947c736b10df0b05e3523f0937149`
- Environment qualification: `evaluations/skill-system-production-refinement/environment-qualification.md`

## Interpreted Contract

1. **Goal:** Place a capability, concern, or structural change at the system location that should own the responsibility instead of attaching it to the nearest visible or convenient surface.
2. **Intended use:** Architecture and system-change decisions about where a responsibility belongs, including requests that suggest a module, service, layer, folder, hook, or adapter as the edit site.
3. **Expected behavior:** Preserve the requested outcome without adopting the suggested location; map the minimum sufficient system; distinguish observation, current ownership, candidate ownership, and dependent surfaces; establish the normal successful owner; choose the smallest complete placement; reconcile dependants; and define ownership-level verification.
4. **Expected result:** A justified placement or restructuring with one clear responsibility owner, reconciled dependent surfaces, and verification that the original need is satisfied at that owner.
5. **Boundary:** A local placement remains valid when the local location owns the responsibility. The workflow must not force broad redesign, reject sound existing patterns, substitute generic architecture preferences for source evidence, or expand the requested outcome.

## Observable Completion Criteria

- The decision states the responsibility or outcome that must hold and where the need was observed.
- The analysis traces the relevant current producers, consumers, handoffs, and owners to concrete system evidence.
- The observation point, nearby convenient hooks, candidate owner, and dependent surfaces are distinguished.
- The selected placement is justified by normal successful ownership rather than proximity.
- The proposed change is the smallest complete placement and avoids parallel or split ownership.
- Affected interfaces, configurations, modules, and workflows are reconciled where applicable.
- Verification checks correctness at the responsibility owner and across dependants.
- Unrelated architecture and request scope remain intact.

## Evaluation Surfaces

- Description loop: native proactive selection on four fresh-agent system-placement tasks.
- Reference loop: not applicable because the target contains no `references/` directory or files.
- Body loop: isolated architecture-decision artifacts audited against connected source fixtures and this contract.

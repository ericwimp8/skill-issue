# Code Testing Discipline Contract

- Target: `evaluations/scenario-skill-refinement/code-testing-discipline/skill/SKILL.md`
- Metadata: `evaluations/scenario-skill-refinement/code-testing-discipline/skill/agents/openai.yaml`
- Environment: `evaluations/skill-system-production-refinement/environment-qualification.md`
- Refinement mode: automatic semantic refinements supported by retained evidence.

## Goal

Cause agents doing automated code-test work to establish behavior from production code and its governing contract, select the smallest owned-interface test surface, assert observable outcomes, keep fixtures proportionate, capture regressions before correction, and run focused then nearby broader validation.

## Intended Use

Creating, editing, reviewing, or running automated code tests.

## Expected Behavior

After loading the skill, the agent traces the relevant production behavior before choosing assertions, tests through the behavior-owning interface, avoids private structure and incidental interaction details, builds minimal setup, proves a regression test fails before the production fix when applicable, and validates from focused to broader scope.

## Expected Result

The resulting test work is behavior-focused, connected to production-source truth, appropriately scoped, independently runnable, and accompanied by evidence for the required validation order and any regression red-green sequence.

## Boundaries

- Production source and governing contracts own intended behavior; tests do not define it.
- The discipline governs automated code-test work below end-to-end integration without prescribing one language, framework, mocking library, or exact test answer.
- The smallest layer is the smallest layer that can observe behavior through its owned interface, rather than mechanically the narrowest possible unit.
- Assertions target externally observable outcomes, while implementation details, private structure, and incidental call order remain outside the contract.
- Existing unrelated behavior and repository-specific validation requirements remain intact.

## Observable Completion Criteria

1. The agent inspects the production behavior and relevant contract before deciding assertions.
2. The chosen test exercises the behavior through its owned public interface.
3. Assertions concern observable outcomes and reject implementation-detail coupling.
4. Setup and fixtures remain smaller than the behavior under test.
5. Regression work retains pre-fix failure and post-fix pass evidence.
6. Validation records the focused test before the nearest broader relevant check.
7. The produced test and any required correction pass their isolated executable checks.

## Evaluation Surface

Code artifacts, command transcripts, and agent reasoning traces retained from isolated fixture workspaces.

## Reference Inventory

The target contains no `references/` directory. Reference qualification is not applicable.

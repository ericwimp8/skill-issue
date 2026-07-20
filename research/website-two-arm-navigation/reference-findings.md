# Two-Arm Website Navigation Research

## Research Goal

Identify restrained, production-quality patterns for presenting Skill Issue's shared identity and two complementary product experiences: building better skills and evaluating model-and-harness environments.

## Findings

## References

- [shadcn/ui](https://ui.shadcn.com/): compact navigation, explicit selected controls, thin bordered surfaces, and a centered proposition before the component grid.
- [Linear](https://linear.app/): quiet header, one product-level promise, and a bounded transition into the active product experience.
- [Clerk](https://clerk.com/): concise hero hierarchy, paired actions, restrained technical background, and strong content grouping.

## Decisions

- Keep the existing neutral grid, light/dark themes, thin borders, and compact semantic color accents.
- Center and enlarge the Skill Issue wordmark in the shared header.
- Lead with one shared product promise before asking visitors to choose an arm.
- Use a prominent two-option segmented selector for `Build Skills` and `Evaluate Environments`.
- Place each arm inside a clearly bounded content region with its own explanation and actions.
- Load generated examples directly from tracked `showcase-skills/*/skill/*/SKILL.md` files so adding a folder automatically expands the gallery.
- Open each example in a centered reading overlay while keeping browsing controls visible behind it.

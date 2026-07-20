# Two-Arm Website Image Comparison Audit

## Audit Basis

- Target images: current Skill Issue header and hero screenshots supplied on 20 July 2026.
- Comparison images: selected external references and the implemented two-arm website states.
- Target scope: shared identity, two-arm navigation, hero hierarchy, skill showcase, and evaluation continuity.
- Comparison method: direct pairwise UI comparison at desktop width.

## Scope And Exclusions

- Binding visual details checked: prominent centered identity, minimal shared hero, clear two-option navigation, readable selected state, coherent content transition, and existing theme continuity.
- Non-binding context ignored: reference-site branding, exact copy, imagery, pricing, account controls, and product-specific navigation.

## Needs Fixing

### 1. Product identity is visually subordinate to section navigation

- Target image: current Skill Issue header and hero.
- Comparison image: `research/website-two-arm-navigation/screenshots/shadcn-home.png`.
- Difference: Skill Issue appears as a compact left-aligned header label while the page's evaluation headline owns nearly all visual emphasis; shadcn establishes one dominant centered product proposition before its content grid.
- Why it matters: the current hierarchy presents the benchmark as the whole product and leaves the two-part Skill Issue identity unexplained.
- Fix direction or ownership uncertainty: move the dominant Skill Issue identity and product-level explanation into the shared hero, then place the two-arm selector between shared actions and arm-specific content.
- Confidence: high.

### 2. The current primary navigation cannot express two equal product experiences

- Target image: current Skill Issue header.
- Comparison image: `research/website-two-arm-navigation/screenshots/shadcn-home.png`.
- Difference: the target uses quiet text anchors for Results and Method; the comparison uses compact, bounded controls and cards to make distinct destinations and selected states visually legible.
- Why it matters: users need to understand Build Skills and Evaluate Environments as peer product paths rather than anchors within one benchmark page.
- Fix direction or ownership uncertainty: replace the Results and Method anchors with a two-option segmented selector whose active state is explicit and whose content boundary begins immediately below it.
- Confidence: high.

### 3. The shared product promise needs to lead both arms

- Target image: current Skill Issue header and hero.
- Comparison image: `research/website-two-arm-navigation/screenshots/linear-home.png`.
- Difference: Linear keeps its compact header quiet while one product-level proposition owns the page hierarchy; Skill Issue currently leads with evaluation-only copy.
- Why it matters: both building and evaluating skills need to feel like parts of one product before the visitor chooses a path.
- Fix direction or ownership uncertainty: place the shared Skill Issue promise above the build/evaluate selector and keep each arm's explanation inside its selected content region.
- Confidence: high.

### 4. The active arm needs a clear content boundary

- Target image: current Skill Issue hero.
- Comparison image: `research/website-two-arm-navigation/screenshots/linear-home.png`.
- Difference: Linear transitions from its shared proposition into one bounded product preview; the current page flows directly from its hero into unrelated sections.
- Why it matters: switching product paths should read as changing the active experience rather than jumping between page anchors.
- Fix direction or ownership uncertainty: begin the selected arm immediately below the segmented selector with a strong heading, supporting copy, and visually bounded content.
- Confidence: high.

### 5. Shared actions should remain visually secondary to the promise

- Target image: current Skill Issue hero.
- Comparison image: `research/website-two-arm-navigation/screenshots/clerk-home.png`.
- Difference: Clerk centers one concise promise, a short explanation, and two compact actions; the target gives an arm-specific headline and large actions similar weight.
- Why it matters: visitors should understand Skill Issue before choosing a product arm or action.
- Fix direction or ownership uncertainty: keep the shared tagline dominant, limit the supporting paragraph width, and use compact paired actions before the segmented selector.
- Confidence: high.

### 6. The existing grid can support the new product structure

- Target image: current Skill Issue hero.
- Comparison image: `research/website-two-arm-navigation/screenshots/clerk-home.png`.
- Difference: both pages use a quiet technical grid, but Clerk keeps content surfaces restrained so the grid remains environmental rather than decorative.
- Why it matters: the current visual identity can support the redesign without introducing a new illustration system or heavy color treatment.
- Fix direction or ownership uncertainty: retain the current grid, neutral palette, thin borders, and compact green/red semantic accents while changing hierarchy and interaction.
- Confidence: high.

## Working Well After Implementation

### 1. Shared identity and build path

- Target image: current Skill Issue header and hero.
- Comparison image: `research/website-two-arm-navigation/screenshots/local-build-skills.png`.
- Result: the centered wordmark is visibly more prominent, the shared promise now explains the whole product, and the first viewport leads naturally into the two-option product selector.
- Reference alignment: matches the reference hierarchy without importing their branding, color, or illustration systems.
- Confidence: high.

### 2. Generated skill reader

- Target image: requested centered skill overlay.
- Comparison image: `research/website-two-arm-navigation/screenshots/local-skill-reader.png`.
- Result: the selected `SKILL.md` opens in a focused, scrollable, keyboard-dismissable reader with the gallery preserved behind a restrained blur.
- Reference alignment: uses the site's existing neutral surfaces and monospaced evidence style, so the raw skill remains legible and authentic.
- Confidence: high.

### 3. Evaluation continuity

- Target image: existing evaluation hero and chart page.
- Comparison image: `research/website-two-arm-navigation/screenshots/local-evaluate-environments.png`.
- Result: the selected Evaluate Environments state is explicit, the original evaluation explanation remains intact, and the existing metric and chart sequence continues directly below the arm introduction.
- Reference alignment: the switcher creates the bounded product transition seen in Linear while preserving Skill Issue's existing chart language and visual identity.
- Confidence: high.

### 4. Scalable example gallery

- Target image: requested browseable collection of generated skills.
- Comparison image: `research/website-two-arm-navigation/screenshots/local-build-gallery.png`.
- Result: the three current examples form a clear equal-weight gallery, longer descriptions remain readable, and the grid can add further rows as tracked skill folders are added.
- Reference alignment: the restrained card grid follows shadcn's compositional clarity while retaining Skill Issue's typography, grid, and dark-theme treatment.
- Confidence: high.

## Final Assessment

- The implemented hierarchy resolves the two original visual problems: the product identity now leads the page, and the two product arms have an explicit selected state.
- The shared hero, build gallery, reader overlay, and evaluation path remain coherent in both light and dark themes.
- No reference-specific branding, imagery, or color system was copied.

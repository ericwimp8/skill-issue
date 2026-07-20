# Web Accessibility Evidence Guide

## Evidence Classes

Use one of these labels for every substantive finding, pass, or limitation:

- **Observed:** directly demonstrated in inspected source, rendered behavior, captured accessibility state, or a recorded tool result. Name the observation method and tested state.
- **Inferred:** strongly indicated by inspected evidence but dependent on a runtime, content, browser, device, or assistive-technology condition that was not directly exercised. State the evidence and the dependency.
- **Unverified:** plausible or required to investigate, but the available environment did not establish the behavior. Describe the exact follow-up rather than presenting a finding or pass.

Automated output is observed evidence that a tool reported a rule result. It is not automatically observed evidence of the complete user impact, correct remediation, or overall accessibility.

## Authority Levels

- Treat WCAG success criteria and applicable WAI-ARIA requirements as normative when the requested standard uses them.
- Treat W3C techniques, failures, tutorials, Easy Checks, and the ARIA Authoring Practices Guide as informative implementation and evaluation guidance.
- Treat project conventions and tool documentation as evidence about the implementation or checker, not as accessibility standards.
- Link the exact authoritative page supporting a material criterion mapping. When applicability or conformance interpretation is uncertain, label the mapping tentative.

Core public sources:

- [WCAG 2.2](https://www.w3.org/TR/WCAG22/)
- [How to Meet WCAG 2.2](https://www.w3.org/WAI/WCAG22/quickref/)
- [Understanding WCAG 2.2](https://www.w3.org/WAI/WCAG22/Understanding/)
- [Easy Checks](https://www.w3.org/WAI/test-evaluate/preliminary/)
- [Selecting Web Accessibility Evaluation Tools](https://www.w3.org/WAI/test-evaluate/tools/selecting/)
- [ARIA Authoring Practices Guide](https://www.w3.org/WAI/ARIA/apg/)

W3C states that evaluation tools cannot check every accessibility aspect and cannot determine accessibility; human judgment is required. Preserve that boundary in findings and summaries.

## Affected-User Prompts

Identify concrete effects for relevant people, including people who:

- navigate by keyboard, switch, speech input, or other non-pointer input;
- use screen readers, refreshable braille, magnification, or high-contrast settings;
- have low vision, color-vision differences, or photosensitivity;
- are Deaf, hard of hearing, or depend on captions and transcripts;
- have cognitive, learning, language, attention, or memory disabilities;
- have limited dexterity, tremor, reach, or fine-motor control;
- use zoom, text resizing, alternate orientation, reduced motion, or constrained viewports.

Describe the task effect as blocked, difficult, error-prone, confusing, fatiguing, time-sensitive, or degraded. Do not infer a diagnosis or claim every member of a group experiences the same effect.

## Priority Rules

Use these levels as review priorities rather than compliance grades:

- **Critical:** observed barrier blocks a critical task or safety-relevant information for affected users with no reasonable workaround.
- **High:** observed barrier blocks or seriously impedes an important task, affects a repeated or broad interaction, or creates substantial error or exclusion risk.
- **Medium:** observed or strongly inferred barrier degrades a task, increases effort or confusion, or has a limited workaround, reach, or state.
- **Low:** localized improvement with limited task impact, or a well-supported risk whose user impact is modest.
- **Needs verification:** evidence is insufficient to assign a defensible finding priority; retain it as follow-up rather than forcing a severity.

Explain priority using task importance, impact, reach, frequency, workaround burden, and confidence. Tool severity may inform investigation but does not own report priority.

## Human and Assistive-Technology Follow-Up

Require targeted follow-up when correctness depends on perception, interaction quality, content intent, announcements, interoperability, or user strategy. Examples include:

- screen-reader name, role, state, reading order, live-region, and focus announcements;
- keyboard efficiency and focus behavior across complete dynamic workflows;
- magnification, reflow, text-spacing, high-contrast, and reduced-motion usability;
- caption accuracy, audio-description adequacy, alternative-text purpose, and plain-language clarity;
- voice-input label matching and target operability;
- representative usability testing with disabled people for critical journeys.

Name the browser, device, assistive technology, state, journey, and expected behavior when known. A first pass may recommend these checks but cannot substitute for them.

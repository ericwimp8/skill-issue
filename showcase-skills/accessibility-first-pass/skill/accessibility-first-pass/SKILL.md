---
name: accessibility-first-pass
description: Performs responsible first-pass accessibility reviews of web pages, features, and implementations. Use when Codex needs to investigate accessibility risks, evidence, and remediation direction without claiming complete accessibility or standards conformance.
---

# Accessibility First Pass

## Establish the Review Contract

1. Identify the supplied page, feature, implementation, user journey, and requested standard or policy target.
2. State the exact review scope: routes, states, viewports, themes, input modes, source areas, browsers, and tools actually available.
3. Record unavailable inputs and surfaces before investigating. Ask only when a missing input changes the requested scope; otherwise proceed and label the limitation.

## Investigate the Available System

1. Trace the feature from its entry point through components, styles, state, validation, routing, and rendered effect. Inspect concrete implementations rather than stopping at wrappers or tests.
2. Identify existing accessibility tooling, browser automation, lint rules, component tests, build commands, and project instructions from production configuration and source.
3. Run applicable project-native automated checks when the environment safely permits them. Record the command, target, result, and tool limitations. Do not install tools, broaden the target, or treat a clean scan as proof.
4. Inspect rendered behavior when a runnable target or retained rendering is available. Exercise relevant states with keyboard and pointer input, visible focus, zoom or reflow, and supported color or motion preferences.
5. Inspect source and accessibility semantics for the selected surface. Correlate markup and styles with rendered behavior instead of inferring the complete user experience from source alone.
6. Consult authoritative guidance appropriate to each material claim. Distinguish normative requirements from informative techniques or patterns.

## Perform the First-Pass Checks

Select checks that apply to the reviewed surface rather than mechanically filling a checklist. Cover relevant concerns across:

- document structure, language, headings, landmarks, reading order, and bypass mechanisms;
- keyboard access, focus order, focus visibility, focus management, traps, and operable target behavior;
- accessible names, roles, values, states, relationships, status updates, and native semantics;
- images, icons, audio, video, captions, transcripts, and other content alternatives;
- form labels, instructions, required states, validation, error identification, and error recovery;
- contrast, color dependence, text spacing, target size, orientation, zoom, reflow, and obscured content;
- motion, animation, flashing, time limits, drag or pointer gestures, and input alternatives;
- dynamic content, dialogs, menus, tabs, custom widgets, loading, empty, success, and failure states.

For checks that depend on user judgment, assistive technology, browser and device combinations, content intent, or states unavailable in the environment, record the exact follow-up needed instead of assigning a pass.

## Classify and Prioritize Findings

1. Label every substantive statement as observed evidence, source-backed inference, or unverified behavior using the definitions in the evidence reference.
2. Describe the affected users and blocked, difficult, or degraded task. Avoid treating a guideline identifier as the impact description.
3. Reproduce each finding with the shortest reliable inspection path and include relevant source locations, rendered state, command output, or captured evidence.
4. Assign priority from user impact, task criticality, reach, workaround burden, and evidence confidence. Record uncertainty rather than inflating severity.
5. Give remediation direction at the behavior owner. Prefer native platform behavior where it meets the requirement; avoid prescribing an exact patch unless the available source establishes it.
6. Map authoritative criteria only when the evidence supports the relationship. Mark tentative mappings as requiring confirmation.

## Write the Report

Use `assets/accessibility-first-pass-report.md` as the output structure unless the user supplies an established project format. Keep all sections even when the value is `None observed`, `Not run`, or `Unknown`.

- Lead with scope, methods, evidence available, and material limitations.
- Order findings by priority and explain the rationale for each.
- Separate findings, passed checks, follow-up checks, and unknowns.
- Claim a passed check only for the exact state, method, and environment tested.
- End with a prioritized next-action list and the human or assistive-technology testing still required.
- State that the first pass does not establish overall accessibility or standards conformance.

## Completion Gate

Before concluding, verify that the report:

- identifies scope, affected users, evidence, reproduction or inspection steps, remediation direction, and limitations;
- distinguishes observation, inference, and unverified behavior throughout;
- records automated and manual methods without overstating either;
- identifies checks requiring human or assistive-technology testing;
- avoids unsupported conformance, certification, completeness, or absence-of-barriers claims; and
- uses durable, share-safe paths and excludes secrets or private identity information.

If a requested conclusion exceeds the evidence, provide the bounded findings and state what additional evaluation would be required.

## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/web-accessibility-evidence.md`: Web accessibility evidence, authority, user-impact, priority, and follow-up guidance. Use when classifying review evidence, mapping authoritative criteria, prioritizing findings, or defining human and assistive-technology follow-up.

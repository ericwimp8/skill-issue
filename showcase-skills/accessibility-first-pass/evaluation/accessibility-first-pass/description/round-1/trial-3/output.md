# Accessibility First Pass: Service Health Dashboard

## Review Scope

- **Target:** The single static service-health dashboard in `fixtures/dashboard/index.html`.
- **Included journeys, states, routes, and viewports:** Initial dashboard view at 1280 × 720 CSS pixels and a narrow 320 × 568 CSS-pixel viewport; service-status text, the SVG chart, and the latency table. The page has no controls or changed states.
- **Available source and rendered access:** Full HTML/CSS source inspection plus local rendering in Chrome on 20 July 2026. Browser DOM inspection, computed styles, geometry, contrast calculations, and a narrow-viewport screenshot were used.
- **Project tooling and commands used:** A local `python3 -m http.server` served the fixture. No repository accessibility scanner or accessibility test command was found in the project scripts or dependencies, so no automated-rule result is claimed.
- **Guidance baseline:** WCAG 2.2 was used only to map the observed risks; no conformance level was supplied by the project. W3C WAI guidance for complex images informed the chart remediation.
- **Exclusions:** Authentication, dynamic status updates, failure states, data refresh, interaction behavior, other routes, 200% text-only resize, actual 400% browser zoom, forced-colors/high-contrast modes, assistive technologies, mobile browsers, and disabled-user testing.

## Evidence Summary

- **Observed:** At 1280 CSS pixels, the `main` region rendered 1120 pixels wide and the SVG rendered 800 × 300 pixels. At 320 CSS pixels, the document scroll width was 1128 pixels, producing document-level horizontal scrolling; the SVG extended from x=8 to x=808. The narrow rendering showed the status words, a solid black wedge-like chart, the table, and a horizontal scrollbar.
- **Observed:** The browser DOM snapshot exposed the `main` landmark, level-one heading, status paragraph, and table semantics, but no named or described chart content.
- **Observed:** Computed status colors were `rgb(0, 128, 0)` and `rgb(255, 0, 0)` on white at 16px/400 weight. Calculated contrast was approximately 5.14:1 for “Payments” and 3.998:1 for “Messaging.” The SVG path computed to black fill with no stroke.
- **Source-backed:** The only status distinction is `.good { color: green; }` and `.bad { color: red; }`; the visible words do not state either service's health. The SVG has no title, description, accessible name, axes, labels, legend, or adjacent data equivalent. `main` has a fixed `width: 1120px`; the SVG has fixed dimensions; no responsive rule or viewport meta element is present.
- **Tool:** No automated accessibility scanner was available or run. Browser inspection is recorded as observed evidence rather than a scanner pass.
- **Inference:** The chart appears intended to communicate a service-health trend, but its series, scale, time range, values, and relationship to the service names cannot be established from the supplied implementation. Actual mobile layout may differ because the page omits viewport configuration.

## Prioritized Findings

### High — Service Health Is Conveyed Only by Red and Green

- **Affected users and tasks:** People with color-vision deficiencies, low vision, reduced contrast perception, monochrome or customized displays, and screen-reader users trying to identify which service is healthy or unhealthy.
- **Evidence level:** Source-backed and observed.
- **Evidence:** The paragraph contains only “Payments Messaging.” CSS makes the first word green and the second red; there is no visible label, icon with text equivalent, or programmatic health value. In the browser DOM snapshot, the paragraph was exposed simply as “Payments Messaging.”
- **Reproduction or inspection steps:** Open the page and inspect the status line without relying on hue, or inspect the paragraph in the browser DOM/accessibility representation. The health meaning disappears.
- **User impact:** A user can identify service names but cannot reliably determine their health, which defeats the dashboard's primary task.
- **Authoritative guidance:** [WCAG 2.2 SC 1.4.1, Use of Color](https://www.w3.org/WAI/WCAG22/Understanding/use-of-color.html) requires information conveyed by color to have another visible cue. Programmatic status also relates to [SC 1.3.1, Info and Relationships](https://www.w3.org/WAI/WCAG22/Understanding/info-and-relationships.html).
- **Remediation direction and owner:** At the status-data/rendering owner, render explicit text such as “Payments — operational” and “Messaging — outage.” Color and icons may remain supplemental. Ensure the service name and state form one programmatically associated item, for example a list of service/state pairs. If status changes dynamically in the real product, add a deliberately designed announcement strategy based on update urgency rather than applying a generic live region.
- **Verification route:** Check the status at default colors, grayscale/forced colors, and with CSS disabled; inspect computed semantics; then have screen-reader and color-vision-deficiency users confirm that each service state is independently understandable.
- **Confidence and open questions:** High confidence that the supplied state is color-only. The intended vocabulary and whether status can update are unknown.

### High — The Chart Has No Understandable Visual or Text Equivalent

- **Affected users and tasks:** Blind and low-vision users, people who enlarge content, people with cognitive or learning disabilities, and any user trying to understand the trend or compare it with the service/table data.
- **Evidence level:** Observed, source-backed, and inference about the intended data meaning.
- **Evidence:** The SVG contains only an unlabeled path. It has no `<title>`, description, text labels, axes, legend, accessible name, or linked data. The browser DOM snapshot omitted any chart content. The path defaults to black fill with no stroke and rendered as a solid wedge-like shape, so the purported trend lacks visible scale, series, or meaning as well.
- **Reproduction or inspection steps:** Inspect the DOM representation and navigate by landmarks/headings: the sequence moves from the status paragraph directly to the table. Visually inspect the chart and attempt to identify its metric, time range, units, or values; none are supplied.
- **User impact:** Users cannot recover the information the graphic is presumed to convey. A blind user encounters no chart at all; sighted users receive an ambiguous shape rather than a self-explanatory data visualization.
- **Authoritative guidance:** [WCAG 2.2 SC 1.1.1, Non-text Content](https://www.w3.org/WAI/WCAG22/Understanding/non-text-content.html) requires an equivalent-purpose text alternative for meaningful non-text content. [WAI Complex Images guidance](https://www.w3.org/WAI/tutorials/images/complex/) recommends identifying a chart and providing its essential values, relationships, and trends in an available long description.
- **Remediation direction and owner:** First define the chart's actual data contract. At the chart component/data owner, provide a visible title and summary, visible axes/units/series identification, and the underlying values in an adjacent accessible table or equivalent structured text. Give the SVG an accessible name tied to that title and reference the longer description when appropriate. If the graphic is truly decorative and duplicates all nearby information, explicitly hide it from assistive technology instead; that choice must be based on product meaning.
- **Verification route:** Compare the visual chart and alternative value-for-value; inspect browser semantics; test at high zoom and forced colors; then ask screen-reader and low-vision users to answer the same trend and value questions as sighted users.
- **Confidence and open questions:** High confidence that the supplied SVG lacks an alternative. The exact missing information cannot be named until the chart's metric, series, and source data are defined.

### High — The Page Produces Document-Level Horizontal Scrolling at 320 CSS Pixels

- **Affected users and tasks:** People using high zoom or screen magnification, users with low vision, and mobile users reading status and chart content without repeatedly panning.
- **Evidence level:** Observed and source-backed; WCAG failure mapping remains a scoped interpretation because a meaningful chart can qualify for a two-dimensional-layout exception.
- **Evidence:** With a 320 × 568 viewport, browser measurement reported `clientWidth: 320`, `scrollWidth: 1128`, `main` width 1120, and SVG width 800. The screenshot showed a document-level horizontal scrollbar. Source fixes `main` at 1120 pixels and the SVG at 800 pixels, with no responsive CSS; the page also omits a viewport meta element.
- **Reproduction or inspection steps:** Set the viewport to 320 CSS pixels and load the page. Observe that the document is over three times the viewport width and requires horizontal panning to reach the full visualization.
- **User impact:** Users must pan across the whole page, increasing physical and cognitive effort and making it difficult to retain context. On an actual mobile browser, omitted viewport configuration may additionally cause desktop-scale rendering; that behavior was not tested.
- **Authoritative guidance:** [WCAG 2.2 SC 1.4.10, Reflow](https://www.w3.org/WAI/WCAG22/Understanding/reflow.html) expects non-exempt content to work at 320 CSS pixels without two-dimensional scrolling. W3C notes that diagrams and data tables may need two-dimensional layout, but the exception applies to the necessary content section rather than justifying document-wide overflow.
- **Remediation direction and owner:** At the page-layout owner, replace the fixed main width with a fluid width plus a reasonable `max-width`, responsive padding, and `box-sizing`. Add appropriate viewport configuration. At the chart owner, make the SVG scale within its container using a `viewBox` and responsive dimensions; if detailed chart content genuinely requires horizontal exploration, contain scrolling within a labeled chart region while keeping the rest of the document reflowed.
- **Verification route:** Re-test at 320 CSS pixels and at 400% zoom from a 1280-pixel starting viewport; verify status text and table cells remain readable without two-dimensional scrolling, and any chart-only overflow is contained and operable. Repeat in iOS Safari and Android Chrome.
- **Confidence and open questions:** High confidence in the observed overflow. Whether the intended chart needs an exception and how a real mobile browser handles the missing viewport meta element remain open.

### Medium — “Messaging” Text Does Not Meet Minimum Contrast for Normal Text

- **Affected users and tasks:** People with low vision or reduced contrast sensitivity reading the unhealthy service name.
- **Evidence level:** Observed.
- **Evidence:** Browser computed styles reported 16px normal-weight red (`rgb(255, 0, 0)`) on white. The calculated contrast ratio was approximately 3.998:1; “Payments” green measured approximately 5.14:1.
- **Reproduction or inspection steps:** Load the initial page, read computed foreground/background colors for `.bad`, and calculate contrast using the WCAG relative-luminance formula without rounding the result up.
- **User impact:** The unhealthy service—the information likely requiring the fastest attention—is harder to perceive for users with reduced contrast sensitivity.
- **Authoritative guidance:** [WCAG 2.2 SC 1.4.3, Contrast (Minimum)](https://www.w3.org/WAI/WCAG22/Understanding/contrast-minimum.html) sets 4.5:1 for normal text and treats thresholds without rounding.
- **Remediation direction and owner:** At the status visual-design owner, choose a darker error color that provides at least 4.5:1 against every actual background while retaining the explicit state text required by the first finding. Prefer a design-token correction if the color is shared.
- **Verification route:** Recalculate contrast from computed colors in every status state and theme, then visually check at zoom and in high-contrast/forced-color settings.
- **Confidence and open questions:** High confidence for the tested white-background state. Other themes or status backgrounds were not supplied.

## Checks Requiring Human or Assistive-Technology Testing

- **Screen-reader comprehension:** After chart and status remediation, ask users of at least the supported macOS/iOS and Windows/Android screen-reader combinations to identify both service states and explain the chart trend/values. Browser DOM inspection cannot establish efficient or comprehensible screen-reader use.
- **Low-vision and magnification use:** Test the dashboard at 200% and 400% zoom, with text enlargement, and with screen magnification. Success means status, table, chart summary, and any local chart scroller remain understandable without losing context.
- **Color and contrast adaptations:** Have users with color-vision deficiencies and users of forced colors/high contrast confirm that health states remain distinguishable and that chart series/values are perceivable without hue alone.
- **Cognitive comprehension:** Have representative users confirm that service names, explicit state wording, chart title/summary, units, and time range support the intended operational decision without requiring inference from layout.
- **Mobile touch and reading behavior:** Test physical iOS and Android devices after responsive remediation, including orientation changes and platform text-size settings. The desktop viewport emulation does not establish mobile browser scaling or usability.

## Material Limitations

- The fixture is a 23-line static page with one populated table row; no application data model, update mechanism, chart requirements, interactions, alternate states, or error/loading states were available.
- Browser rendering covered only Chrome at 1280 × 720 and an emulated 320 × 568 viewport. Actual 400% browser zoom, 200% text resize, mobile browsers, orientation, and platform accessibility settings were not exercised.
- No automated accessibility scanner was configured or run, so there are no violation, incomplete, or pass counts. This is a limitation, not a pass.
- The browser DOM snapshot provided computed semantics but was not a screen-reader session. No assistive technology or disabled-user testing occurred.
- Contrast was calculated only for the two status words against the rendered white background. Chart contrast, focus appearance, motion, keyboard operation, errors, and dynamic announcements were either inapplicable to the supplied static state or untestable from it.
- The chart's actual purpose and data are absent. Remediation must begin by recovering that product meaning rather than inventing an alternative description from the path coordinates.

## Conclusion

This first pass establishes source-backed and browser-observed barriers in the dashboard's primary information: service health is color-only, the chart has no understandable or accessible equivalent, the layout creates severe narrow-viewport overflow, and the red status text falls below the normal-text contrast threshold. Remediate in that order: make status meaning explicit, define and expose the chart's data and summary, make page and chart layout responsive, then correct the error-color token. The next verification step is a repeated source/rendered review at 320 CSS pixels and 400% zoom with automated scanning added, followed by targeted screen-reader, low-vision, color-adaptation, mobile, and disabled-user testing.

This first pass does not prove accessibility, completeness, certification, or standards conformance.

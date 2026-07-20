# Gardening Web Application Scenario

Send these turns exactly as written and in order. Wait for the agent to finish each response before sending the next turn.

## Turn 1

> I want to create and maintain a living A-to-B plan for a production-ready gardening web application called GardenFlow. It should help home gardeners decide what to plant, organize their growing spaces, follow care schedules, and keep useful records across multiple growing seasons. We are planning the complete product from the current idea through a public launch, but we are not implementing it during this conversation. Start the plan for this work.

## Turn 2

> The main users are first-time gardeners, experienced home growers managing several beds, apartment gardeners using containers, and households where more than one person shares garden work. Add the important outcomes for each group, including confidence for beginners, speed for experienced users, clear collaboration, and useful continuity between seasons. The product should remain approachable without becoming simplistic.

## Turn 3

> Use a fixed technical baseline for the plan: React and TypeScript for the browser application, a responsive component system, a Python FastAPI service, PostgreSQL, object storage for photos, background workers for scheduled jobs, and deployment through managed cloud services. The application must support current desktop and mobile browsers. Record these as established constraints rather than reopening the stack choice later.

## Turn 4

> The initial product has five major product areas: a personalized dashboard, a plant library, a visual garden planner, a care calendar, and a seasonal journal. Describe what each area must enable, how users move between them, and which capabilities belong in the first public release. Include empty states and the experience for someone who has not created a garden yet.

## Turn 5

> Add a complete interaction and visual-quality task covering responsive behavior, keyboard navigation, screen-reader support, color contrast, reduced motion, touch targets, loading and error states, printable views, and a calm visual language suitable for outdoor and household use. Include an observable accessibility and usability standard in the completion criteria.

## Turn 6

> The core data needs to represent users, households, gardens, growing areas, beds, containers, plants, varieties, plantings, care tasks, observations, harvests, photos, weather references, reminders, and seasonal summaries. Integrate the domain-model work into the plan, including relationships, history, deletion behavior, import and export needs, and the difference between shared reference data and user-owned records.

## Turn 7

> Add the backend work needed for account management, garden configuration, plant search, planting records, care scheduling, journal entries, photo handling, notifications, household collaboration, data export, and administrative support. Include API evolution, background processing, rate limits, validation, failure recovery, and operational visibility without turning the plan into endpoint-by-endpoint implementation instructions.

## Turn 8

> Map the important end-to-end user journeys into the plan: onboarding and creating a first garden, choosing suitable plants, laying out a bed, recording a planting, receiving and completing care work, adding observations and photos, recording a harvest, reviewing a season, inviting another household member, and returning after several months away. Make sure the task order supports proving these journeys before launch.

## Turn 9

> Gardeners often work where connectivity is poor. Add realistic offline and unreliable-network expectations for viewing the current garden, recording completed care, adding observations, and queuing photos. Include conflict behavior, synchronization visibility, recovery after failed uploads, and clear limits on what the first release will support.

## Turn 10

> Add the security, privacy, and account-safety work. Cover household permissions, personal location data, private garden photos, password and session security, abuse prevention, account recovery, auditability of sensitive actions, retention, export, deletion, backups, and incident response. The finished plan should make these properties observable rather than treating them as general aspirations.

## Turn 11

> Write a self-contained prompt that I can give to a research agent to investigate authoritative horticultural datasets suitable for GardenFlow. The researcher should compare plant taxonomy, regional growing guidance, sowing and harvest windows, climate-zone coverage, licensing, attribution, update cadence, machine-readable access, and risks in combining multiple sources. Add the research and its decision point to the living plan, but do not perform the research now.

## Turn 12

> Add the work for bringing approved horticultural reference data into the product. Include source review, normalization, taxonomy reconciliation, regional variants, attribution, correction workflows, update handling, provenance, administrator review, and a safe way to distinguish general guidance from advice tailored to a user's garden.

## Turn 13

> The care calendar should turn plantings and seasonal conditions into understandable work without overwhelming users. Add planning for recurring and one-off tasks, postponement, dependencies, household assignment, reminders, quiet periods, weather-aware suggestions, overdue work, bulk completion, calendar export, and explanations showing why a recommendation appeared.

## Turn 14

> The visual planner must support multiple gardens, beds, containers, dimensions, orientation, notes, drag-and-drop placement, spacing guidance, succession planting, companion and conflict warnings, historical layouts, and printable plans. Add the significant product, data, interaction, and validation work needed to make that useful on both desktop and touch devices.

## Turn 15

> Add a photo-assisted observation workflow. Users should be able to attach photos to plants or garden areas, compare changes over time, describe symptoms, and receive carefully bounded guidance without the product claiming certainty. Include upload resilience, image privacy, metadata handling, accessibility, unsafe-content controls, human-readable limitations, and escalation to authoritative resources.

## Turn 16

> Add optional community features for sharing a garden snapshot, publishing a seasonal lesson, following a public garden, and reporting unsuitable content. Include privacy defaults, consent, moderation, blocking, rate limits, copyright and attribution considerations, youth safety, administrator tooling, and a release boundary that keeps community risk from delaying the core private gardening experience.

## Turn 17

> Define the operational workflows for customer support and administration. Cover account investigation with appropriate safeguards, data corrections, plant-data corrections, notification troubleshooting, content moderation, feature flags, service status, support escalation, audit logs, and the evidence operators need to diagnose problems without casually exposing private user data.

## Turn 18

> Write a detailed prompt for a research agent to study how real home gardeners plan and maintain gardens over a season. The research should examine paper and digital workflows, planning frequency, forgotten tasks, collaboration, accessibility barriers, motivation, record keeping, regional differences, and reasons people abandon gardening tools. Add this research to the plan as an input to product validation without conducting it now.

## Turn 19

> Add analytics and observability work that can show whether onboarding, garden creation, planning, care completion, journaling, collaboration, and seasonal return are functioning. Include privacy-respecting product analytics, service metrics, structured logs, traces, alerting, dashboard ownership, data-quality checks, and rules preventing sensitive garden or account information from entering telemetry.

## Turn 20

> Build out the quality strategy as a broad dependency-ordered part of the plan. It should cover domain and service tests, browser behavior, responsive layouts, accessibility, synchronization and retry behavior, migrations, background jobs, security checks, performance, representative user journeys, release-candidate testing, and production monitoring. Distinguish automated proof from checks that need people or real devices.

## Turn 21

> Add the release and deployment path from internal development through preview environments, seeded demonstrations, beta users, production migration, rollback, staged rollout, support readiness, release notes, incident response, and post-launch observation. Include the evidence required before widening access and the conditions that should pause or reverse a release.

## Turn 22

> Add the commercial and ongoing-service considerations without choosing a final business model. The plan should allow evaluation of a useful free product, optional paid household features, storage and notification costs, third-party data costs, customer support load, data portability, account cancellation, and the effect that monetization choices could have on privacy and trust.

## Turn 23

> Write a self-contained research prompt comparing weather, frost-date, climate-zone, and geocoding services that could support GardenFlow. It should investigate geographic coverage, forecast and historical data, reliability, licensing, attribution, privacy implications of location queries, caching rights, rate limits, pricing growth, failure behavior, and the risks of presenting uncertain environmental guidance. Add the resulting research decision to the plan without selecting a provider.

## Turn 24

> Review the entire living plan now. Reconcile duplicated ideas, make sure existing constraints remain in the starting position, ensure the desired position describes the finished product, correct the dependency order of the broad tasks, and strengthen completion criteria for the major user journeys, accessibility, privacy, reliability, operations, data provenance, and public launch. Keep unresolved implementation choices visible without turning them into user questions.

## Turn 25

> Pause the planning update long enough to create a reusable agent skill called garden-content-quality inside the scenario workspace. It should guide agents that write GardenFlow plant descriptions and care guidance so the content is clear, region-aware, source-conscious, accessible to beginners, explicit about uncertainty, and careful around toxic plants, chemicals, pets, children, and situations requiring authoritative advice. Create the complete skill now with only the files it genuinely needs, then record its use in the living plan.

## Turn 26

> Integrate the new garden-content-quality skill into the product plan. Identify which content-production, plant-data review, administrator, testing, and release activities should use it, what evidence would show it was applied effectively, and how the product avoids treating generated wording as a substitute for horticultural review or safety decisions.

## Turn 27

> Add a substantial content-governance task covering editorial standards, source attribution, regional review, change history, corrections, translations, seasonal updates, safety review, administrator approval, user reports, emergency withdrawal of bad guidance, and communication when previously published advice changes. Reconcile this with the imported data and garden-content-quality work rather than creating parallel owners.

## Turn 28

> Review the garden-content-quality skill against the need we just established. Improve it if necessary so its description identifies the correct use boundary and its instructions remain general enough to work across plant profiles, care reminders, warnings, seasonal explanations, and administrator corrections. Do not tailor it to the wording of this scenario; keep the skill reusable.

## Turn 29

> Perform a final planning pass before the architectural change I will give you next. Check that every completion criterion is created by at least one task, every broad task has a meaningful outcome, research decisions occur before dependent implementation choices, beta and release work depend on adequate product and operational proof, and no important requirement has been reduced to a vague quality statement.

## Turn 30

> Change the product architecture and revise the whole plan coherently. Instead of five route-level product areas, GardenFlow should become one adaptive workspace with panels that reveal planning, care, reference, and journal functions around the selected garden context. Establish clear ownership for workspace state, navigation state, server data, offline changes, synchronization, notifications, and shared household updates. Reorganize the broad tasks into logical blocks that implement this architecture without leaving the former page-based structure as a competing owner, and reconcile every affected completion criterion.

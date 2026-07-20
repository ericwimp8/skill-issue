# Build Week Definitive Completion Plan

## Authority

This is the definitive project-level progress and completion document for the
OpenAI Build Week submission. It supersedes earlier project progress documents
when they disagree about priority, dependency order, current work, or what
remains before submission.

Existing plans, progress files, research, evaluation contracts, and production
source remain supporting evidence and implementation references. They do not
override this checklist. Do not delete older working documents until they have
been reviewed for unique information worth retaining.

The official competition rules remain authoritative for submission
requirements, and production source remains authoritative for implemented CLI
behavior.

## Completion Rule

The project is complete only when every checkbox in this document is checked
and the Devpost submission has been successfully entered before the official
deadline.

## Deadline

- **Official deadline:** Tuesday, 21 July 2026 at 5:00 PM PDT.
- **Adelaide deadline:** Wednesday, 22 July 2026 at 9:30 AM ACST.
- **Authority:** [OpenAI Build Week Official Rules](https://openai.devpost.com/rules).

## Critical Path

The work should proceed in this priority order while independent workstreams
run concurrently:

1. Obtain Claude Code Max access.
2. Audit, harden, test, and freeze the CLI campaign baseline.
3. Begin the sequential Codex-in-Claude-Code campaign lane immediately after
   the baseline is frozen.
4. Run all other eligible evaluation lanes in parallel.
5. Generate the Claude Code comparison skills.
6. Complete the website using mock data in parallel, then replace it with the
   accepted evaluation and generated-skill evidence.
7. Publish the website and downloadable CLI.
8. Capture and produce the final demonstration video.
9. Upload the video and complete the Devpost submission.

The Claude Code Max purchase, CLI hardening, website work, campaign preparation,
and video preparation do not need to wait for one another. Full campaign runs
must use the same frozen known-good CLI baseline so their results remain
comparable. Before that freeze, use only bounded smoke, resume, authentication,
and cleanup probes; do not count them as campaign results.

## 1. Immediate Access And Campaign Preparation

- [ ] Obtain a Claude Code Max subscription.
- [ ] Confirm normal Claude Code authentication without inspecting or copying
      credential contents.
- [ ] Confirm the exact Claude/Fable model identifier available through the
      subscription.
- [ ] Select and record one campaign reasoning target and the closest supported
      equivalent for harnesses without an independent control.
- [ ] Confirm the exact harness executable, version, model identifier, and
      reasoning control for every campaign configuration.
- [ ] Confirm every evaluation workspace will be outside this repository and
      every output will use a distinct directory beneath repository `output/`.
- [ ] Confirm the campaign orchestration uses no more than six simultaneous
      evaluation commands and keeps shared progress updates serialized.

## 2. CLI Audit, Hardening, And Baseline Freeze

The CLI has not yet received the complete production review required for the
final campaign. Establish intended behavior from production source before using
tests as validation.

### Source Audit

- [ ] Trace the CLI entrypoints, argument parsing, lifecycle, installation,
      evaluation preparation, replay, result derivation, cleanup, and recovery
      paths end-to-end.
- [ ] Trace each supported harness adapter through its concrete process,
      environment, session, permission, instrumentation, and cleanup behavior.
- [ ] Inspect failure handling for startup, authentication, permissions,
      protocol parsing, timeouts, cancellation, process ownership, interrupted
      runs, and incomplete cleanup.
- [ ] Verify the CLI changes only state it owns and does not inspect, copy,
      replace, back up, repair, or delete user configuration, credentials,
      unrelated skills, or unrelated files.
- [ ] Verify evaluation workspaces, output roots, temporary skills, answer
      sheets, private state, transcripts, events, and retained artifacts obey
      their documented privacy and ownership boundaries.
- [ ] Record every concrete defect, unsupported assumption, and release blocker
      found during the audit.

### Hardening And Tests

- [ ] Fix every campaign-blocking defect at the production behavior owner.
- [ ] Remove accidental complexity introduced by the fixes and keep harness
      paths shallow and directly traceable.
- [ ] Add focused tests for validated CLI contracts and failure boundaries.
- [ ] Add or complete adapter, lifecycle, cleanup, cancellation, artifact, and
      recovery coverage where production behavior requires it.
- [ ] Run focused tests for each changed behavior.
- [ ] Run the broader CLI test, vet, formatting, and cross-platform build checks
      required by the repository.
- [ ] Complete bounded real-harness smoke tests for every campaign-ready route.
- [ ] Review all failures as tooling failures or model results according to the
      production evaluation contract.
- [ ] Confirm no smoke leaves temporary skills, private recovery state, owned
      processes, or disposable workspaces behind.

### Freeze

- [ ] Resolve every CLI release blocker required for the evaluation campaign.
- [ ] Build the committed campaign revision as the next known-good CLI snapshot.
- [ ] Record the immutable commit, CLI version, build date, and harness versions
      used by the campaign.
- [ ] Confirm all full campaign commands use that same known-good snapshot.
- [ ] Treat any later campaign-affecting CLI fix as a baseline change requiring
      an explicit comparability review and any necessary reruns.

## 3. Evaluation Campaign

The campaign contains ten harness-and-model configurations and three governed
30-turn scenarios per configuration, for 30 completed runs.

### Campaign Execution Rules

- [ ] Use the frozen known-good CLI for every full campaign run.
- [ ] Run every governed scenario in full without turn truncation.
- [ ] Give every run a unique clean external Git workspace and output location.
- [ ] Retain the required `result.json` and `website.json` artifacts.
- [ ] Record the effective harness version, model identifier, reasoning setting,
      run ID, timestamps, result path, and cleanup evidence.
- [ ] Count a tooling-complete run with no expected skill calls as a completed
      evaluation result.
- [ ] Treat launch, authentication, permission, protocol, session, artifact,
      instrumentation, process, or cleanup failures as tooling failures.
- [ ] Preserve failed and blocked attempt history after successful reruns.
- [ ] Keep independent eligible evaluation lanes running when another lane is
      blocked.

### OpenAI Codex Command Permission Rule

- [ ] Execute every OpenAI Codex harness evaluation through the exact known-good
      evaluation command with `sandbox_permissions="require_escalated"`.
- [ ] Require the escalation justification to state that nested Codex needs its
      normal authenticated session database and session state under
      `CODEX_HOME`.
- [ ] Confirm the escalation removes only the outer main-thread shell sandbox for that
      command and does not weaken or bypass the evaluator-owned inner Codex
      `workspace-write` sandbox, approval policy, clean configuration, model,
      reasoning, or cleanup controls.
- [ ] If escalation is denied, record the run as blocked, verify partial cleanup,
      and continue independent work instead of retrying inside the read-only
      outer sandbox.

### First Full Lane: Claude Code With Codex

This lane is the first full campaign lane after the CLI baseline freeze. Its
three runs are sequential because they share the project-local Codex-backed
Claude route and may require bounded resume or tooling diagnosis.

- [ ] Preflight `.skill-issue/claudex/claudex`, its selected model, proxy
      readiness, version, executable path, and cleanup behavior without reading
      credentials or tokens.
- [ ] Complete **CLA-COD-01 — Gardening Web Application**.
- [ ] Complete **CLA-COD-02 — Community Archive Desktop Application**.
- [ ] Complete **CLA-COD-03 — Neighborhood Emergency Preparedness Program**.
- [ ] Confirm no two Claude Code/Codex campaign runs overlapped.
- [ ] Resolve and safely rerun any Claude Code/Codex tooling failure before
      changing the Claude Code route.
- [ ] Stop only the proxy process owned by `.skill-issue/claudex/manage` after
      all three runs complete; do not reset or delete the isolated runtime.

Other non-Claude-Code evaluations may use the remaining process capacity while
this sequential lane runs.

### Claude Code Fable Transition

The project-local `claudex` launcher is a Codex-backed route. It starts a
localhost proxy and injects an isolated Claude configuration, proxy URL,
authentication token, model aliases, delegated-agent model, and forced model. Moving
to normal Claude Code therefore requires an explicit executable and environment
transition rather than changing one global environment variable.

- [ ] Identify the normal operator-owned Claude Code executable.
- [ ] Confirm the normal route does not inherit the project-local proxy URL,
      token, isolated `CLAUDE_CONFIG_DIR`, or forced Codex aliases.
- [ ] Confirm normal Claude Code Max authentication and the exact Fable model.
- [ ] Run a small built-in Claude/Fable smoke in a clean external workspace.
- [ ] Run the existing custom two-turn Claude/Fable smoke in a separate clean
      external workspace.
- [ ] Confirm both smokes produce their required artifacts, preserve a stable
      session, create the expected workspace effects, and clean their temporary
      skills, private state, and owned processes.
- [ ] Open the Claude/Fable campaign lane only after both smokes pass.

### Full Evaluation Checklist

#### OpenAI Codex — GPT-5.6 Sol

- [ ] **COD-SOL-01:** Gardening Web Application.
- [ ] **COD-SOL-02:** Community Archive Desktop Application.
- [ ] **COD-SOL-03:** Neighborhood Emergency Preparedness Program.

#### Claude Code — Fable

These three runs may run concurrently after the Fable transition gate passes.

- [ ] **CLA-FAB-01:** Gardening Web Application.
- [ ] **CLA-FAB-02:** Community Archive Desktop Application.
- [ ] **CLA-FAB-03:** Neighborhood Emergency Preparedness Program.

#### Claude Code — Codex

- [ ] **CLA-COD-01:** Gardening Web Application.
- [ ] **CLA-COD-02:** Community Archive Desktop Application.
- [ ] **CLA-COD-03:** Neighborhood Emergency Preparedness Program.

#### Cursor — Fable

- [ ] **CUR-FAB-01:** Gardening Web Application.
- [ ] **CUR-FAB-02:** Community Archive Desktop Application.
- [ ] **CUR-FAB-03:** Neighborhood Emergency Preparedness Program.

#### Cursor — Codex

- [ ] **CUR-COD-01:** Gardening Web Application.
- [ ] **CUR-COD-02:** Community Archive Desktop Application.
- [ ] **CUR-COD-03:** Neighborhood Emergency Preparedness Program.

#### Cursor — Grok

- [ ] **CUR-GRO-01:** Gardening Web Application.
- [ ] **CUR-GRO-02:** Community Archive Desktop Application.
- [ ] **CUR-GRO-03:** Neighborhood Emergency Preparedness Program.

#### Cursor — Composer

- [ ] **CUR-COM-01:** Gardening Web Application.
- [ ] **CUR-COM-02:** Community Archive Desktop Application.
- [ ] **CUR-COM-03:** Neighborhood Emergency Preparedness Program.

#### Pi — Codex

- [ ] **PI-COD-01:** Gardening Web Application.
- [ ] **PI-COD-02:** Community Archive Desktop Application.
- [ ] **PI-COD-03:** Neighborhood Emergency Preparedness Program.

#### OpenCode — Codex

- [ ] Qualify and freeze OpenCode evaluation support before starting this lane.
- [ ] **OPE-COD-01:** Gardening Web Application.
- [ ] **OPE-COD-02:** Community Archive Desktop Application.
- [ ] **OPE-COD-03:** Neighborhood Emergency Preparedness Program.

#### Kilo Code — Codex

- [ ] Qualify and freeze Kilo Code evaluation support before starting this lane.
- [ ] **KIL-COD-01:** Gardening Web Application.
- [ ] **KIL-COD-02:** Community Archive Desktop Application.
- [ ] **KIL-COD-03:** Neighborhood Emergency Preparedness Program.

### Campaign Acceptance

- [ ] Confirm all 30 runs are tooling-complete and retained.
- [ ] Confirm every failure or blocker has an accurate history and disposition.
- [ ] Confirm every configuration used the intended model and reasoning target.
- [ ] Confirm all campaign workspaces, temporary skills, private state, and owned
      processes are cleaned.
- [ ] Freeze the accepted result set used by the website.

## 4. Claude And Codex Skill Generation Comparison

Ten skills have already been generated with Codex. Generate the same ten skill
requests with Claude Code so the website can present a controlled comparison.

- [ ] Inventory the ten accepted Codex-generated skills and their original input
      contracts.
- [ ] Confirm the Claude Code generation environment contains no ambient
      project or user `AGENTS.md` or `CLAUDE.md` instructions.
- [ ] Install only the intended Skill Issue skill payload into the clean Claude
      Code generation environment.
- [ ] Confirm the Claude generation prompts and inputs match the Codex generation
      conditions.
- [ ] Generate all ten corresponding skills with Claude Code Max.
- [ ] Retain the generated skills and sufficient non-sensitive provenance for
      the comparison.
- [ ] Verify the website data can associate each Claude result with its matching
      Codex result.
- [ ] Decide whether the website presents each pair side-by-side, through a
      toggle, or through another equally direct comparison interaction.

## 5. Website Data Story And Design

Website work can continue with mock data before campaign completion. Published
claims and final charts must use the accepted evaluation artifacts.

### Supported Matrix And Visual Language

- [ ] Update website harnesses, models, labels, filters, summaries, and mock data
      to match the definitive campaign matrix.
- [ ] Design a color system that distinguishes models and harnesses without
      making either comparison ambiguous.
- [ ] Verify chart colors, legends, labels, focus states, tooltips, and text are
      accessible and consistent.
- [ ] Remove illustrative combinations that are absent from the accepted matrix.

### Harness View

The first story is whether the harness changes Codex's skill-calling behavior.
Codex is the common model across the planned harnesses, making it the controlled
comparison for this view.

- [ ] Create a harness-led view comparing Codex across OpenAI Codex, Claude Code,
      Cursor, Pi, OpenCode, and Kilo Code where accepted evidence exists.
- [ ] Show whether expected calls, missed calls, additional calls, and behavior
      over turns differ by harness.
- [ ] Make the comparison limits and unavailable configurations explicit.
- [ ] Write chart titles, supporting copy, and annotations that explain the
      harness story without overstating causation.

### Model View

The second story is how each available model performs in its supported harness
configurations.

- [ ] Create a model-led view for GPT-5.6 Sol/Codex, Fable, Composer, and Grok
      using only accepted configurations.
- [ ] Compare Fable in Claude Code and Cursor.
- [ ] Compare Cursor's Fable, Codex, Composer, and Grok configurations.
- [ ] Show native and cross-harness model results clearly without implying a
      missing fully crossed matrix.
- [ ] Write chart titles, supporting copy, and annotations that explain the
      model story and its limits.

### Generated Skills

- [ ] Add all accepted Codex-generated skills to the website.
- [ ] Add all corresponding Claude-generated skills to the website.
- [ ] Implement the selected direct comparison interaction.
- [ ] Verify every displayed skill is traceable to its retained generation
      evidence.

### Methodology And Limitations Page

- [ ] Add a dedicated methodology and limitations page.
- [ ] Explain what the project is trying to learn and why skill-calling behavior
      matters.
- [ ] Explain the governed scenarios, expected calls, instrumentation, harness
      isolation, model selection, reasoning target, scoring, and retained
      evidence at an understandable level.
- [ ] State that this is a home-run project conducted by one person with limited
      time, subscriptions, hardware, accounts, and campaign repetitions.
- [ ] Explain that the results describe the tested versions, models, harnesses,
      prompts, and environment rather than every possible installation.
- [ ] Explain what conclusions the evidence supports and what it does not.
- [ ] Identify omitted configurations, unavailable repetitions, resource-driven
      compromises, and work deferred because of the deadline.
- [ ] Describe likely future research, broader replication, additional models,
      additional platforms, and stronger statistical evidence.
- [ ] Explain the project's motivations and the practical decisions made during
      Build Week.

### Final Website Acceptance

- [ ] Replace final mock evaluation data with the frozen accepted artifacts.
- [ ] Verify all totals, labels, charts, filters, comparisons, and narrative copy
      against the retained results.
- [ ] Verify responsive behavior, keyboard access, readable contrast, loading,
      links, downloads, and failure states.
- [ ] Run the repository's complete website validation commands.
- [ ] Inspect the production build and final published experience.

## 6. CLI Distribution And GitHub Pages

- [ ] Produce the intended downloadable CLI artifacts from the accepted release
      revision.
- [ ] Verify each published CLI artifact's platform, architecture, version, and
      checksum or other integrity evidence.
- [ ] Document installation and first-run usage clearly enough for another user
      to install and run the CLI.
- [ ] Make the downloadable CLI discoverable from the website.
- [ ] Configure the static site for GitHub Pages.
- [ ] Publish the production website to GitHub Pages.
- [ ] Verify the public website URL, navigation, assets, charts, methodology,
      skill comparisons, and CLI downloads from a clean browser session.

## 7. Demonstration Video

The narration transcript and Remotion/Speechify production materials already
exist elsewhere. The remaining work is to finish the visual story using the
actual product and final evidence.

- [ ] Locate and confirm the authoritative video transcript.
- [ ] Confirm the Remotion project and Speechify voiceover workflow are usable.
- [ ] Map every narration section to the required website, CLI, evaluation,
      generated-skill, or methodology footage.
- [ ] Capture clean footage of the finished website and relevant CLI behavior.
- [ ] Add the footage behind the narration with the required animations,
      transitions, labels, and emphasis.
- [ ] Ensure the video clearly demonstrates what was built and how Codex and
      GPT-5.6 contributed.
- [ ] Keep the final demonstration shorter than three minutes.
- [ ] Verify the final video has clear audio and remains understandable without
      unsupported claims.
- [ ] Export and review the final render from beginning to end.
- [ ] Upload the final video to YouTube with public visibility.
- [ ] Verify the public YouTube URL works without authentication.

## 8. Submission Form And Final Release

- [ ] Choose the submission category that best fits the project.
- [ ] Complete the submission title and text description.
- [ ] Explain the project's features and functionality accurately.
- [ ] Provide the public GitHub Pages URL.
- [ ] Provide the public YouTube demonstration URL.
- [ ] Provide the source repository URL with the access and licensing required
      by the official rules.
- [ ] Update the README to explain how Codex was used, where it accelerated the
      work, where human product and engineering decisions were made, and how
      GPT-5.6 and Codex contributed to the final result.
- [ ] Distinguish pre-existing work from work added during the submission period
      with appropriate repository and Codex-session evidence.
- [ ] Verify the submitted project installs, runs consistently, and behaves as
      depicted in the video and written description.
- [ ] Review the complete submission for unsupported claims, broken links,
      missing fields, privacy issues, credential exposure, and inconsistent
      versions.
- [ ] Submit the completed entry before Wednesday, 22 July 2026 at 9:30 AM ACST.
- [ ] Verify Devpost shows the submission as successfully entered.

## 9. Final Completion Gate

- [ ] Claude Code Max access is available and all required Claude work is done.
- [ ] The CLI audit, hardening, tests, known-good freeze, and distribution are
      complete.
- [ ] All 30 full evaluation runs are accepted.
- [ ] All ten Claude-generated skills and all ten Codex-generated skills are
      available for comparison.
- [ ] The website accurately presents the harness story, model story,
      methodology, limitations, generated skills, and CLI download.
- [ ] The GitHub Pages deployment is public and verified.
- [ ] The final demonstration video is public on YouTube and linked correctly.
- [ ] The README and repository access satisfy the submission requirements.
- [ ] The Devpost form is complete and successfully submitted before the
      deadline.
- [ ] Every checkbox in this definitive plan is checked.

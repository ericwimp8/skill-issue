# Community Archive Desktop Application Scenario

This document is the human-readable governance view of the fixed conversation embedded in the standalone CLI. The runner sends each prompt below verbatim and in order without adapting later turns to the model's responses.

## Turn 1

> Let's do Dictate Plan. I want to make BoxIndex, a tiny offline browser app that lets a volunteer catalog a community archive box without uploading anything. Start one living A-to-B plan for taking the empty workspace to a locally verified app.

## Turn 2

> The starting folder is empty and has Node.js available. The user will perform every in-browser check. The finished app should add a record with an accession ID, title, and short note; list saved records after reload; reject duplicate accession IDs; and export the records as JSON. Integrate those facts and observable outcomes into the plan so automated verification stays in Node and never opens or automates a browser.

## Turn 3

> Keep the build deliberately small: plain HTML, CSS, and JavaScript, local storage only, keyboard-usable controls, clear validation text, and Node-run automated checks for accession-ID normalization and duplicate detection. Do not open or automate a browser or install browser tooling. Put the broad work in dependency order without prescribing file-level implementation.

## Turn 4

> Write a self-contained prompt for a fresh agent to inspect the empty project and recommend the smallest internal record shape and file structure that support the app and its focused tests. Ask for only the decision and concise rationale. Add that inspection as an execution input in the living plan.

## Turn 5

> That completes the planning input. Reconcile the full plan so the baseline, desired app, ordered work, and proof agree without duplicated requirements, then mark it ready for execution.

## Turn 6

> Implement the smallest working BoxIndex app from the approved plan. Use source inspection and Node-run checks only; do not open or automate a browser or install browser tooling. Stop after add, duplicate rejection, reload, and JSON export are supported locally.

## Turn 7

> Add focused automated tests for accession-ID normalization and duplicate rejection. Cover one accepted unique ID and one duplicate that differs only by letter case, then run the smallest relevant test command.

## Turn 8

> Create a reusable agent skill named respectful-archive-labels in this workspace. It should guide agents writing short catalog titles and notes to distinguish evidence from inference, preserve supplied terminology, and state uncertainty plainly. Keep it portable and include only files it genuinely needs.

## Turn 9

> A volunteer reports that an accession ID with surrounding spaces can bypass duplicate detection. Reproduce the report with the smallest failing test, trace the actual normalization path, fix the causal owner, and rerun the focused checks.

## Turn 10

> Manual entry, stored records, export, and tests now need one authoritative accession-ID rule. Restructure the small app so that responsibility has one clear owner, update every affected caller, and verify behavior remains unchanged.

## Turn 11

> Write a concise handoff prompt for a fresh agent to add CSV export later without changing the current JSON behavior. Then update the living plan to record that deferred extension separately from the completed application.

## Turn 12

> Run the existing focused test suite as final evidence. Update the living plan with what was actually verified, preserve the deferred CSV work, and remove any completion claim that the retained evidence does not support.

## Turn 13

> Before we continue, remind me which record fields BoxIndex stores and which export format already works.

## Turn 14

> CSV export is no longer deferred. Implement the smallest export that preserves accession ID, title, and note while leaving the existing JSON export unchanged.

## Turn 15

> Add focused automated tests for CSV column order, stable record order, and escaping of commas and quotation marks. Run the focused tests and keep JSON coverage intact.

## Turn 16

> A volunteer reports that a note containing a line break produces a malformed CSV file. Reproduce the report with the smallest failing test, trace the value through the export path, fix the causal owner, and rerun the focused checks.

## Turn 17

> Update the living plan so CSV export is recorded as completed rather than deferred. Reconcile the claim with the exact export behavior and retained tests without weakening the JSON requirements.

## Turn 18

> Quick reminder: does BoxIndex require an internet connection to add or review records?

## Turn 19

> Revise the respectful-archive-labels skill created earlier so it also guides wording derived from filenames or incomplete donor notes. Preserve the distinction between evidence and inference, keep it portable, and remove anything it does not genuinely need.

## Turn 20

> Write a self-contained prompt for a fresh accessibility reviewer to inspect the source for record entry, validation, saved-record navigation, and both export controls without opening or automating a browser. Ask only for prioritized findings with concise evidence and no implementation.

## Turn 21

> Inspect the current interface through its source and implement the smallest keyboard, focus, and validation-message improvements that are clearly needed. Use Node-run checks only and do not open or automate a browser. Preserve the offline workflow and existing record behavior.

## Turn 22

> Add focused Node-run automated checks for the interaction behavior changed in the accessibility pass. Assert observable focus and validation outcomes rather than private event wiring, do not open or automate a browser, then run the focused checks.

## Turn 23

> Manual entry, duplicate detection, storage, JSON export, and CSV export now share the same record-shaping rules. Restructure the app so that record normalization and serialization responsibilities have clear owners, update affected callers, and preserve visible behavior.

## Turn 24

> Which two export formats can a volunteer use now?

## Turn 25

> A volunteer reports that deleting the final record can appear to work, but the record returns after reloading. Reproduce the report with the smallest failing test, trace the stored collection state, fix the causal owner, and rerun the focused checks.

## Turn 26

> Add the nearest useful collection boundaries exposed by that fix, including an empty archive and multiple retained records. Keep assertions on observable saved and exported results, then run the focused suite.

## Turn 27

> Write a concise prompt for a fresh agent to assess future CSV import without implementing it. Ask for the smallest validation and conflict boundaries compatible with current accession handling, then update the living plan to retain import as deferred work.

## Turn 28

> Improve the duplicate-accession validation message so it identifies the conflicting normalized ID without exposing internal storage details. Keep the correction limited to the current interface.

## Turn 29

> Run the focused Node-only automated suite and the nearest broader Node-run project check that can catch collateral behavior changes. Do not open or automate a browser. Report the commands and outcomes without changing application behavior.

## Turn 30

> Write a concise final handoff prompt for a fresh maintainer. It should name the verified BoxIndex behaviors, the focused checks to rerun first, and the deferred CSV-import boundary without inventing implementation instructions.

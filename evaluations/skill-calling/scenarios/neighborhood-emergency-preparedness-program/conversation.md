# Neighborhood Emergency Preparedness Program Scenario

This document is the human-readable governance view of the fixed conversation embedded in the standalone CLI. The runner sends each prompt below verbatim and in order without adapting later turns to the model's responses.

## Turn 1

> Let's do Dictate Plan. I want to make ReadyCard, a tiny dependency-free browser tool that helps one household select a few preparedness actions and print a private one-page reminder. Start one living A-to-B plan for taking the empty workspace to a locally verified tool.

## Turn 2

> The starting folder is empty and has Node.js available. The user will perform every in-browser check. The finished tool should offer six fixed actions, let a household select actions and add one contact name, show only selected actions in a print view, retain the choices after reload, and make no network requests. Integrate those facts and observable outcomes into the plan so automated verification stays in Node and never opens or automates a browser.

## Turn 3

> Write a self-contained prompt for a fresh accessibility reviewer to identify the smallest keyboard, readable-text, and print requirements this one-page tool must satisfy. Ask for a concise decision record without implementation. Add that review as an input to the living plan.

## Turn 4

> Keep the implementation to plain HTML, CSS, and JavaScript with local storage and focused Node-run automated checks for selection and print filtering. Do not open or automate a browser or install browser tooling. Reconcile the complete plan into dependency order, make the evidence observable, and mark it ready for execution.

## Turn 5

> Implement the smallest working ReadyCard tool from the approved plan. Use source inspection and Node-run checks only; do not open or automate a browser or install browser tooling. Stop after selection, contact entry, reload, and print-view filtering are supported locally.

## Turn 6

> Add focused automated tests proving that selected actions appear in the print model and unselected actions do not. Keep the test layer small and run the focused checks.

## Turn 7

> Create a reusable agent skill named preparedness-message-boundary in this workspace. It should guide agents writing short preparedness reminders to use plain language, preserve uncertainty, point to authoritative emergency direction, and avoid individualized safety promises. Keep it portable and include only files it genuinely needs.

## Turn 8

> A household reports that an action they cleared can return in the print view after reloading. Reproduce the report with the smallest failing test, trace the stored selection through the current print path, fix the causal owner, and rerun the focused checks.

## Turn 9

> The screen list, persisted choices, print view, and tests now need one authoritative selection rule. Restructure the small tool so that responsibility has one clear owner, update affected callers, and verify behavior remains unchanged.

## Turn 10

> Write a concise prompt for a fresh facilitator to review the six fixed actions for plain language and appropriate authority boundaries. Request only replacement wording where necessary and a one-sentence reason for each change.

## Turn 11

> Add the nearest useful automated boundaries for the selection fix, including no selected actions and all selected actions. Keep assertions on the observable print model and run the focused suite again.

## Turn 12

> Update the living plan to distinguish completed work from any future content review. Reconcile every completion claim with the application behavior and test evidence actually retained in the workspace.

## Turn 13

> Before we continue, remind me whether ReadyCard sends household choices anywhere outside the browser.

## Turn 14

> Add a small clear-all action that removes the selected preparedness actions and contact name after one explicit confirmation. Keep the six fixed actions and existing print behavior unchanged.

## Turn 15

> Add focused automated tests proving that clear-all removes the saved selections and contact, survives reload, and leaves the six available actions intact. Run the focused tests.

## Turn 16

> A household reports that the contact name can return after clear-all and reload even though the selected actions stay cleared. Reproduce the report with the smallest failing test, trace the persisted state, fix the causal owner, and rerun the focused checks.

## Turn 17

> Update the living plan to include the completed clear-all behavior and its retained evidence. Keep future content review separate and remove any claim the tests do not demonstrate.

## Turn 18

> Quick reminder: how many preparedness actions does ReadyCard offer?

## Turn 19

> Revise the preparedness-message-boundary skill created earlier so it also covers translated summaries and wording adapted by community organizations. Preserve uncertainty and authority boundaries, keep it portable, and remove anything it does not genuinely need.

## Turn 20

> Write a self-contained prompt for a fresh accessibility reviewer to inspect the source for selection controls, contact entry, clear-all confirmation, and the print view without opening or automating a browser. Ask only for prioritized findings with concise evidence and no implementation.

## Turn 21

> Inspect the current interface through its source and implement the smallest keyboard, focus, readable-status, and print-heading improvements that are clearly needed. Use Node-run checks only and do not open or automate a browser. Preserve the fixed actions and private local workflow.

## Turn 22

> Add focused Node-run automated checks for the interaction and print behavior changed in the accessibility pass. Assert observable outcomes rather than internal event wiring, do not open or automate a browser, then run the focused checks.

## Turn 23

> Selection, contact entry, clear-all, reload, and print filtering now depend on the same household state transitions. Restructure the tool so those transitions have one authoritative owner, update affected callers, and preserve visible behavior.

## Turn 24

> What does the print view show when no preparedness actions are selected?

## Turn 25

> A household reports that clearing the contact can leave an empty contact label in the print view. Reproduce the report with the smallest failing test, trace the contact value through the print model, fix the causal owner, and rerun the focused checks.

## Turn 26

> Add the nearest useful print-model boundaries exposed by that fix, including no contact, a contact with surrounding spaces, no selected actions, and all selected actions. Keep assertions on observable output and run the focused suite.

## Turn 27

> Write a concise prompt for a fresh facilitator to assess a future translated-action review without changing the six current actions. Ask for authority and plain-language boundaries, then update the living plan to retain that review as deferred work.

## Turn 28

> Add a short non-modal confirmation message after clear-all so the household can tell that local choices were removed. Keep it direct and do not add another confirmation step.

## Turn 29

> Run the focused Node-only automated suite and the nearest broader Node-run project check that can catch collateral behavior changes. Do not open or automate a browser. Report the commands and outcomes without changing application behavior.

## Turn 30

> Write a concise final handoff prompt for a fresh maintainer. It should name the verified ReadyCard behaviors, the focused checks to rerun first, and the deferred translated-content review without inventing implementation instructions.

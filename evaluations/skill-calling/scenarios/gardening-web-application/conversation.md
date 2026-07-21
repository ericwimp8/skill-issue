# Gardening Web Application Scenario

Send these turns exactly as written and in order. Wait for the agent to finish each response before sending the next turn.

## Turn 1

> Let's do Dictate Plan. I want to make SproutCheck, a tiny dependency-free browser app for one household to record plants and see which ones need watering. Start one living A-to-B plan for taking the empty workspace to a working, locally verified app.

## Turn 2

> The starting position is an empty folder on a computer with a current browser and Node.js. The finished app should let someone add a plant name and watering interval, mark a plant watered today, and show due today or the next due date after reload. It stores data locally and makes no network requests. Integrate that into the plan.

## Turn 3

> Keep the implementation deliberately small: plain HTML, CSS, and JavaScript, keyboard-usable controls, readable status text, a local-calendar-date rule, and automated checks for the date calculation. Put the broad work in dependency order and make the completion criteria observable.

## Turn 4

> That is the complete scope. Reconcile the whole plan so the empty-folder baseline, finished app, implementation path, and evidence agree without duplicated requirements. Mark the plan ready for execution.

## Turn 5

> Write a self-contained prompt for a fresh agent to inspect this empty project and recommend the smallest file structure that supports the planned app and its date tests. Ask for only the decision and concise rationale, without implementation. Add that delegated inspection as the first execution input in the living plan.

## Turn 6

> Planning is complete. Implement the smallest working SproutCheck app from the approved plan now. Keep it dependency-free and stop after the add, water, status, and reload behaviors work.

## Turn 7

> Add focused automated tests for the owned date-status behavior: a plant not yet due, a plant due today, and a plant overdue. Use the smallest test layer that proves the rule, then run those tests.

## Turn 8

> A user reports that a one-day interval can show tomorrow when the plant should be due today. Reproduce the report with the smallest failing test, trace the actual cause through the current date path, fix the causal owner, and rerun the focused checks.

## Turn 9

> Strengthen the automated checks with only the nearest useful calendar boundaries exposed by that fix. Keep the assertions on observable status behavior and run the focused suite again.

## Turn 10

> Create a reusable agent skill named concise-interface-copy in this workspace. It should guide agents editing short labels, buttons, validation messages, and status text to keep them direct, specific, and understandable without surrounding explanation. Keep the skill portable and include only files it genuinely needs.

## Turn 11

> The date-status rule now needs one authoritative owner shared by the interface and automated checks. Restructure the small app so that responsibility is clearly owned rather than duplicated, update affected callers, and verify behavior remains unchanged.

## Turn 12

> Write a concise handoff prompt for a fresh agent to add a one-day snooze later without widening the current product. Then update the living plan to distinguish completed work from that deferred follow-up and reconcile its completion evidence with what was actually verified.

## Turn 13

> Before we continue, remind me what we called the app and whether it sends any garden information over the network.

## Turn 14

> The one-day snooze is no longer deferred. Implement it as a small action for a plant that is due today, keep the existing watering behavior unchanged, and preserve the local-calendar-date rule.

## Turn 15

> Add focused automated tests for the snooze behavior: snoozing a due plant moves it to tomorrow, watering still starts a fresh interval, and reloading preserves the snooze. Run the focused tests.

## Turn 16

> A gardener reports that watering a snoozed plant can leave the old snooze active and make the next due date wrong. Reproduce that report with the smallest failing test, trace the schedule state through the actual code, fix the causal owner, and rerun the focused checks.

## Turn 17

> Update the living plan so the snooze is recorded as completed rather than deferred. Reconcile the completion evidence with the behaviors and tests that now exist without introducing unrelated future work.

## Turn 18

> Quick reminder: what status categories can a plant currently show to the gardener?

## Turn 19

> Revise the concise-interface-copy skill created earlier so it also covers short error and empty-state messages. Preserve its narrow interface-copy purpose, keep the guidance portable, and remove anything the skill does not genuinely need.

## Turn 20

> Write a self-contained prompt for a fresh accessibility reviewer to inspect SproutCheck's keyboard flow, focus visibility, and status announcements. Ask only for prioritized findings with concise evidence, and do not ask the reviewer to edit files.

## Turn 21

> Inspect the current interface and implement the smallest keyboard, focus, and status-announcement improvements that are clearly needed. Keep the visual design and product scope otherwise unchanged.

## Turn 22

> Add focused automated checks for the interaction behavior changed in the accessibility pass. Assert observable keyboard and status outcomes rather than internal event wiring, then run the focused checks.

## Turn 23

> Watering, snoozing, reload, and due-date display now depend on the same schedule transition rules. Restructure the app so those transitions have one authoritative owner, update every affected caller, and verify the visible behavior remains unchanged.

## Turn 24

> Where does SproutCheck keep its plant data between browser sessions?

## Turn 25

> A gardener reports that plant names differing only by surrounding spaces or letter case can create confusing duplicates. Reproduce the behavior with the smallest failing test, trace where names enter and are compared, fix the causal owner, and rerun the focused checks.

## Turn 26

> Add the nearest useful name-validation boundaries exposed by that fix, including a blank trimmed name and a unique name that should still be accepted. Keep assertions on visible outcomes and run the focused suite.

## Turn 27

> Write a concise prompt for a fresh agent to assess a future CSV export without implementing it. Ask for the smallest compatible data shape and exact boundaries with current local storage, then update the living plan to retain that work as deferred.

## Turn 28

> Run the focused automated suite and the nearest broader project check that can catch collateral behavior changes. Report the commands and outcomes without changing application behavior.

## Turn 29

> Reconcile the living plan with the final retained evidence. Mark only demonstrated behavior complete, keep CSV export deferred, and remove any stale claim that the workspace no longer supports.

## Turn 30

> Write a concise final handoff prompt for a fresh maintainer. It should name the verified SproutCheck behaviors, the focused checks to rerun first, and the deferred CSV boundary without inventing implementation instructions.

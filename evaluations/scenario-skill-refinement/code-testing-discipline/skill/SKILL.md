---
name: code-testing-discipline
description: Behavior-focused automated code-testing discipline. Use whenever creating, editing, or running code tests.
---

# Code Testing Discipline

- Establish intended behavior from production code and its governing contract
  before deciding what a test should assert.
- Choose the smallest test layer that observes the behavior through its owned
  interface.
- Assert observable outcomes rather than implementation details, incidental
  call order, or private structure.
- Keep setup and fixtures smaller than the behavior under test.
- For a regression, capture the failure before applying the correction.
- Run the focused test first, then the nearest broader checks that could expose
  collateral behavior changes.

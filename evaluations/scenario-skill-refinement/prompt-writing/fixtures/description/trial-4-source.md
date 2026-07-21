# Implementation Plan Commission

Write a prompt that commissions an authoritative implementation plan for adding an export command to an existing CLI. The planning agent receives repository access and an already-completed research report at `research/export-command.md`. It must trace current production paths, preserve the research report as evidence rather than re-summarizing it in the prompt, and create only the plan needed by an executor. The agent may write the plan under `plans/` but may not implement code, alter tests, commit, or push.


# Workflow Observations

- Fresh-agent concurrency was temporarily exhausted by sibling campaigns. The campaign waited and retried rather than substituting inherited evidence.
- Child agents did not receive a readable effective model or reasoning label despite explicit spawn configuration; the campaign record must own those exact spawn facts.
- Several zsh trial wrappers used the reserved parameter name `status` after successful diagnostic execution. This created avoidable cleanup and evidence reconciliation; future trial prompts should prescribe a safe capture name such as `diagnostic_exit`.
- Relative fixture PATH entries became invalid when an agent changed subprocess working directory. Qualification and trial prompts should require canonical fixture roots before child-only PATH construction.
- Native evidence can leak machine paths even when generated target output is private. Trial protocols need an explicit durable-evidence normalization gate before conclusion.
- Repository `npm run format:check` does not discover showcase Markdown by default; this campaign required an explicit Prettier glob.
- Qualification probe 1 did not conclude after two bounded-finish instructions and was interrupted. A replacement probe limited to selection and exact target-load evidence completed promptly.

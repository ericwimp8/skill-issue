# Invalid Description Preflight

Two fresh agents received the first two round prompts before the evaluator verified their runtime reasoning metadata. Both independently read the exact target and completed the tasks, but their native `turn_context` records identify `gpt-5.6-sol` with `medium` reasoning. The durable environment qualification covers `gpt-5.6-sol` with high reasoning only.

These executions are retained for audit transparency and excluded from description results, counters, and evidence claims:

- Parser session: `019f8262-7430-7140-ab17-aac00befa245`
- Queue session: `019f8262-dc6a-72c2-99bd-0ddf04aaeef8`

Each fixture output and filtered native JSONL trace is preserved in its named directory. Both clean input fixtures were restored before the qualified round began.

# Release Readiness Checker Workflow Observations

- The durable shared environment qualification named GPT-5.6 Sol with high reasoning, while this campaign required medium reasoning. The evaluation gate correctly forced a new target-specific qualification rather than treating model-family support as equivalent measured reliability.
- Project-local proactive discovery required a temporary `.codex/skills/` entry outside the retained campaign directory. The managed workspace required explicit approval to create that evaluation-only link; the campaign therefore records and removes it as cleanup-owned state.
- The evaluation workflow cleanly separated four description-selection trials from three body-behavior cases and prevented body execution before the retained 4/4 description pass.
- Fresh agents independently ran repository formatting checks that did not cover their generated report paths. The final campaign must run a direct formatter check over the complete showcase workspace rather than treating those agent checks as report-format proof.

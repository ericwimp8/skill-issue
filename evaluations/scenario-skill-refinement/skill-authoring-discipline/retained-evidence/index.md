# Retained Evidence Index

## Description

- `description/round-1/trial-1.md` with `retained-evidence/description/trial-1-native-trace.jsonl` and full transcript.
- `description/round-1/trial-2.md` with `retained-evidence/description/trial-2-native-trace.jsonl` and full transcript.
- `description/round-1/trial-3.md` with `retained-evidence/description/trial-3-native-trace.jsonl` and full transcript.
- `description/round-1/trial-4.md` with `retained-evidence/description/trial-4-native-trace.jsonl` and full transcript.
- `description/round-1/audit.md` records the four-trial pass.

## Body

- `behavior/cycle-1/` preserves all three original cases, outputs, the ground-truth exposure failure, and its audit.
- `retained-evidence/body/cycle-1/` preserves native traces and full transcripts for every original case.
- `behavior/cycle-2/` preserves the clean isolated rerun and its passing audit.
- `retained-evidence/body/cycle-2/` preserves the rerun's native trace and full transcript.

## Target Integrity

- Initial and final `SKILL.md` SHA-256: `7fcb2f074490433c6eb6dc3f288d3af1caf353f1cd86b8c65be91d6f017d8bb4`
- Initial and final `agents/openai.yaml` SHA-256: `0262113579c7fb0884333d09c47f27a9df9c7426c47688d34521457ae7f57fdf`
- Final structural validation: `quick_validate.py` returned `Skill is valid!`

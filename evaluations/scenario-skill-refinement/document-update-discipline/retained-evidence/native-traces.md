# Native Invocation Traces

The session JSONL under the qualified Codex environment is the native source. This retained index records the exact session identity and pre-output candidate-read event needed to reproduce each check without treating final prose as proof.

| Phase         | Agent           | Session                                | Pre-output candidate read                                                                               |
| ------------- | --------------- | -------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| Description 1 | `dud_desc_1`    | `019f825d-6b3a-7f00-b9d4-d629c5964d0f` | `2026-07-21T01:49:54.441Z`: `sed -n '1,240p' .../supporting-skills/document-update-discipline/SKILL.md` |
| Description 2 | `dud_desc_2`    | `019f8264-f9a2-7522-acc0-9b1c4abc1aac` | `2026-07-21T01:58:10.411Z`: `cat supporting-skills/document-update-discipline/SKILL.md`                 |
| Description 3 | `dud_desc_3`    | `019f8265-b7f5-7122-8bf1-ad0927274844` | `2026-07-21T01:58:58.402Z`: `cat supporting-skills/document-update-discipline/SKILL.md`                 |
| Description 4 | `dud_desc_4`    | `019f8266-3078-7661-8bf1-72ce6fb4bed6` | `2026-07-21T01:59:29.100Z`: `sed -n '1,240p' supporting-skills/document-update-discipline/SKILL.md`     |
| Body 1.1      | `dud_body_c1_1` | `019f826a-6903-7842-9f7f-5241123de144` | `2026-07-21T02:04:05.950Z`: `cat supporting-skills/document-update-discipline/SKILL.md`                 |
| Body 1.2      | `dud_body_c1_2` | `019f8270-33e6-76c1-842c-404acd7f1cab` | `2026-07-21T02:10:23.814Z`: `cat supporting-skills/document-update-discipline/SKILL.md`                 |
| Body 1.3      | `dud_body_c1_3` | `019f8270-bde5-7d13-aff6-093b09b5619b` | `2026-07-21T02:11:03.423Z`: `sed -n '1,240p' supporting-skills/document-update-discipline/SKILL.md`     |
| Body 2.1      | `dud_body_c2_1` | `019f8272-b283-7b81-937e-a6b9cbce5ed3` | `2026-07-21T02:13:10.008Z`: exact canonical target read before the clean fixture                        |
| Body 2.2      | `dud_body_c2_2` | `019f8273-4b36-7541-ae25-24662c01a810` | `2026-07-21T02:13:47.515Z`: exact canonical target read before the clean fixture                        |
| Body 2.3      | `dud_body_c2_3` | `019f8274-0f54-7b00-a0ab-e904f07355a8` | `2026-07-21T02:14:39.800Z`: exact canonical target read before the clean fixture                        |

Every session was a fresh `fork_turns: "none"` sub-agent. Description prompts did not name the skill. Body prompts named the canonical path because body evaluation measures behavior after loading. Artifact similarity and final `Skills Used` claims were excluded from invocation proof.

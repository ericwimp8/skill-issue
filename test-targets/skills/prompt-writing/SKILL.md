---
name: prompt-writing
description: Prompt-writing discipline for agents and sub-agents. Use when writing prompts for agents, sub-agents, invoked skills, or authoritative plans.
---

## Goal, Completion Criteria, And Authority To Act

Put these in the prompt when the task needs them and no named skill or other
authoritative document for the task already owns them:

- the goal
- completion criteria
- an autonomy or approval boundary

## Discipline

- Do not over-specify or over-prescribe. Put in the prompt only what the agent
  needs for this task, not what the agent already knows.
- If the prompt names a skill or document the agent must use, read it first.
  Do not reteach it, summarize it, or invent structure that is not in it.
  Point to it and add only what it does not already contain.
- Ask for the smallest deliverable that supports the next action.
- For open-ended discovery, define the category to investigate rather than an
  inventory of expected findings.
- Before sending the prompt, delete anything the agent would already know from
  its context, the named skill, or the named document.
- When a prompt fails, fix the decision that failed. Do not add a broad new set
  of instructions.

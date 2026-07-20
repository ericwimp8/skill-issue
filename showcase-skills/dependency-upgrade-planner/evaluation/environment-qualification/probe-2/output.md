# Finished Prompt

Inspect the current repository's payment retry module. Trace production source end to end to identify the concrete code owner that decides and applies retry timing. Do not edit any files. Return a concise evidence summary with repository-relative file paths and symbols that establish the ownership path. If production source cannot establish ownership, stop and report that ownership is unestablished rather than inferring it from tests or documentation.

## Selection Rationale

`supporting-skills/prompt-writing/SKILL.md` naturally applies because its YAML description explicitly covers prompts for agents and sub-agents. The finished prompt follows its direction to state only the goal, evidence-based completion condition, and read-only authority boundary needed by a fresh agent.

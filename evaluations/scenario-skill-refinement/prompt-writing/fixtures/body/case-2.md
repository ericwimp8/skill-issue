# Case 2: Named Authority And Approval Boundary

Write a prompt for an agent to update `plans/release.md` so its deployment step follows the repository's canonical release workflow. The agent must read `scripts/release/README.md`, which owns the workflow. The agent may edit the plan and run Markdown formatting validation. It must request approval before executing deployment commands and stop before commit or push. The smallest handoff is the changed file and validation result.


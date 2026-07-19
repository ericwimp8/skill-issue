# Agent Harness Plugin Packaging Research Map

## Goal

Determine how Skill Issue can distribute a functional bundle of skills, Markdown guidance, scripts, CLI tools, configuration, references, assets, and supporting files across ten current coding-agent harnesses without assuming a universal plugin standard.

## Framing

The final synthesis must present the best-supported direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers. Research is internet-only and prioritizes current official or primary sources.

## Research Domains

- Native extension and marketplace systems
- Filesystem skills, rules, agents, and configuration
- MCP and external CLI integration
- Installation, discovery, invocation, update, trust, and permissions
- Bundle feasibility, unsupported components, and installer fallbacks
- Concrete Skill Issue distribution implications

## Assignments

| Assignment | Harness                         | Primary source targets                                                 | Expected evidence                                                                            |
| ---------- | ------------------------------- | ---------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| 01         | GitHub Copilot                  | GitHub Docs, VS Code/GitHub extension docs, Copilot customization docs | Extension, instructions, skills/agents, MCP, packaging and trust constraints                 |
| 02         | Claude Code                     | Anthropic docs and official repositories                               | Plugins, skills, commands, hooks, agents, MCP, installation and distribution                 |
| 03         | OpenAI Codex                    | Official OpenAI Codex docs and repositories                            | Plugins, skills, agents, MCP, CLI/config packaging and installer boundaries                  |
| 04         | Cursor                          | Cursor docs and official marketplace/extension guidance                | Rules, commands, skills, extensions, MCP and bundle constraints                              |
| 05         | Google Antigravity / Gemini CLI | Google developer docs and official repositories                        | Distinct Antigravity and Gemini CLI surfaces, extensions, skills, MCP and lineage boundaries |
| 06         | JetBrains AI / Junie            | JetBrains documentation and marketplace guidance                       | Distinct AI Assistant and Junie packaging/configuration surfaces and plugin limits           |
| 07         | OpenCode                        | Official OpenCode docs and repositories                                | Plugins, agents, commands, skills, MCP, config and distribution                              |
| 08         | Kilo Code                       | Official Kilo Code docs, marketplace, and repositories                 | Modes, rules, skills, MCP, extension packaging and distribution                              |
| 09         | Replit Agent                    | Official Replit documentation                                          | Skills/rules/connectors, project configuration, deployment and bundle limitations            |
| 10         | Warp                            | Official Warp documentation and repositories                           | Agent modes, rules, workflows, MCP, CLI/extension surfaces and distribution                  |

## Waves And Budget

- Total researcher budget: 10
- Active researcher concurrency: 5
- Wave 1: assignments 01-05
- Wave 2: assignments 06-10, backfilled as Wave 1 completes
- Discovery fan-out: none; the ten caller-selected harnesses consume the full budget
- Cross-checking: each assignment must distinguish documented native support from installer or CLI fallbacks and must caveat unsupported claims

## Outputs

- Assignment folder: `plans/deep-research/agent-harness-plugin-packaging/assignments/`
- Assignment files: `01-github-copilot.md` through `10-warp.md`
- Final synthesis: `plans/deep-research/agent-harness-plugin-packaging/agent-harness-plugin-packaging-deep-research.md`
- Requested synthesis shape: best-supported answer or direction, conditional alternatives, rejected or lower-fit interpretations, evidence, and unresolved blockers

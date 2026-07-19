# Current Coding-Agent Harness Popularity and Support Set

> **Document status — historical research input.** The popularity evidence and methodology remain retained, but its JetBrains-era selected-nine list is superseded. The current selected nine and minimum qualification tier are owned by [`plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`](../../skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md). Use the [documentation authority map](../../skill-issue-project-completion/document-authority-and-update-map.md) before carrying findings forward.

**Research date:** 19 July 2026  
**Question:** Which coding-agent harnesses should Skill Issue support, informed by current popularity and project priorities?

## TL;DR

1. **GitHub Copilot**
2. **Claude Code**
3. **OpenAI Codex**
4. **Cursor**
5. **Google Antigravity / Gemini CLI lineage**
6. **JetBrains AI / Junie**
7. **OpenCode**
8. **Kilo Code**
9. **Pi**

This is the selected Skill Issue support set, not a claim that these are the exact global top nine. The first eight follow the strongest comparable market evidence used in this report. Pi is included as a deliberate project priority because it is a directly relevant, highly extensible terminal harness with strong current open-source momentum. Replit Agent and Warp are outside the current support scope.

## What Counts as a Coding-Agent Harness

For this report, a **coding-agent harness** is the user-facing execution environment around one or more language models that can gather repository context, plan or reason over a task, edit files, invoke development tools or shell commands, observe results, and continue iteratively. The harness may be an IDE, terminal application, extension, desktop application, or cloud workspace.

This definition includes products such as Cursor, Claude Code, Codex, OpenCode, and Pi. It excludes:

- foundation models by themselves;
- autocomplete-only or chat-only assistants without an agentic tool loop;
- general agent SDKs and orchestration frameworks such as LangChain, the OpenAI Agents SDK, or Google ADK;
- observability, evaluation, memory, and MCP infrastructure;
- general-purpose chat products merely capable of discussing code;
- no-code app builders that cannot work as an engineering agent over code or repositories.

## Ranking Method

### Primary evidence

The backbone is Stack Overflow's April 2026 pulse survey of **1,100 developers and working professionals**. Its complete “Coding Agent Tools” chart asks which tools respondents had used for work in the past year and reports: GitHub Copilot 41%, Claude Code 26%, Codex 22%, Cursor 22%, Gemini/Code Assist 21%, Google Antigravity 18%, JetBrains AI 14%, OpenCode 8%, Kilo Code 6%, Replit Agent 6%, Warp 4%, and a long tail at 3% or below ([Stack Overflow survey report](https://stackoverflow.blog/2026/05/27/agents-on-a-leash-agentic-ai-remains-mostly-monitored-at-work/)).

The same article's prose summarizes a shorter “most used in the last six months” view with different figures—61%, 51%, 20%, and 20% for the four leaders. This report uses the full chart consistently because it supplies one question and one comparable field. The discrepancy is a reason to avoid false precision.

### Selection and corroboration

The survey order is adjusted only when necessary:

1. **Current, product-specific users or activity** break ties when available.
2. **First-party adoption claims** outrank funding, valuation, or generic platform reach.
3. **Open-source signals**—GitHub stars, forks, and package downloads—corroborate popularity but do not equal active users.
4. **Broader platform counts** are disclosed but not treated as agent-user counts.
5. Google’s consumer Gemini CLI has transitioned to Antigravity CLI, so those entries are treated as one current product lineage rather than two independent installed bases. Google states that Antigravity CLI and Antigravity 2.0 now share one agent harness, while Gemini CLI remains available for certain enterprise/API configurations ([Google transition announcement](https://developers.googleblog.com/an-important-update-transitioning-gemini-cli-to-antigravity-cli/)).
6. **Project support priorities** may select a closely ranked harness when its architecture is especially relevant to Skill Issue. Pi is the deliberate inclusion under this rule.

The evidence measures **adoption/popularity, not capability, benchmark quality, safety, price, or model intelligence**. The final nine also reflects the project’s implementation scope.

## Selected Nine

### 1. GitHub Copilot

Copilot has the clearest lead. It ranks first in Stack Overflow's comparable chart at 41%. Microsoft reported **over 26 million users** in its FY2026 Q1 earnings call, with 80% of new GitHub developers starting with Copilot during their first week ([Microsoft earnings transcript](https://www.microsoft.com/en-us/investor/events/fy-2026/earnings-fy-2026-q1)). By FY2026 Q3, nearly 140,000 organizations used Copilot and Copilot CLI usage was nearly doubling month over month ([Microsoft FY2026 Q3 earnings transcript](https://www.microsoft.com/en-us/investor/events/fy-2026/earnings-fy-2026-q3)).

Copilot began as an adjacent pair-programming/autocomplete product, but its agent mode, asynchronous coding agent, CLI, code review agent, and Agent HQ now satisfy the harness definition. Its 26-million figure covers Copilot broadly, so it overstates confirmed coding-agent-only usage; the survey is the stronger evidence that agentic use is also leading.

### 2. Claude Code

Claude Code is second in Stack Overflow's full chart at 26% and is especially strong among intensive users: the same survey says it is used by 50% of full-stack developers in single-agent workflows and 70% of daily multi-agent users. Anthropic separately analyzed roughly **400,000 interactive Claude Code sessions from about 235,000 people** between October 2025 and April 2026 and reports that users average 20 hours per week with the tool ([Anthropic usage study](https://www.anthropic.com/research/claude-code-expertise)).

That 235,000-person study sample is not a total user count, but it is unusually direct behavioral evidence. Claude Code's terminal, IDE, web, sandbox, tool-use, and multi-agent surfaces make it an unambiguous harness.

### 3. OpenAI Codex

Codex ties Cursor at 22% in Stack Overflow's complete chart. It takes third because OpenAI reported more than one million developers using Codex in one month at the February 2026 desktop-app launch, and by June 2026 the company said Codex exceeded **five million weekly active users**, according to Axios ([OpenAI launch post](https://openai.com/index/introducing-the-codex-app/), [Axios adoption report](https://www.axios.com/2026/06/02/openai-codex-knowledge-workers)). OpenAI also documented rapid growth in long-horizon agent work across individual and organizational users ([OpenAI adoption study](https://openai.com/index/how-agents-are-transforming-work/)).

Codex spans terminal, IDE, desktop, cloud, SDK, and multi-agent orchestration. The five-million figure is company-supplied through a credible secondary report and includes an expanding non-developer population, so the ordering over Cursor is defensible rather than conclusive.

### 4. Cursor

Cursor also records 22% in the Stack Overflow chart. Cursor says it serves **millions of developers**, many major engineering organizations, and has crossed **$1 billion in annualized revenue** ([Cursor company update](https://cursor.com/blog/series-d)). Those signals show unusually deep paid adoption even though Cursor does not publish a current exact active-user total.

Cursor qualifies as an agent-first IDE and cloud/desktop agent environment, not merely an editor with autocomplete. Its position could reasonably swap with Codex depending on whether paid engineering penetration or reported weekly active users receives more weight.

### 5. Google Antigravity / Gemini CLI lineage

Stack Overflow measured Gemini/Code Assist at 21% and Antigravity at 18%. Adding those percentages would double-count users and a product migration. Google says Gemini CLI grew to **millions of users**, more than 100,000 GitHub stars, and 6,000 merged pull requests before consumer workflows transitioned to Antigravity CLI in June 2026 ([Google transition announcement](https://developers.googleblog.com/an-important-update-transitioning-gemini-cli-to-antigravity-cli/)). Google describes Antigravity 2.0, CLI, and SDK as surfaces over the same agent harness ([Google I/O 2026 summary](https://blog.google/innovation-and-ai/technology/ai/google-io-2026-all-our-announcements/)).

This report therefore ranks the current Google harness lineage once. Its exact place is uncertain because enterprise Gemini CLI and consumer Antigravity overlap operationally but are not identical products.

### 6. JetBrains AI / Junie

JetBrains AI records 14% in Stack Overflow's chart. JetBrains now provides an agent-hosting environment with Junie, Codex, Claude Agent, and ACP-compatible agents; Junie itself is available in JetBrains IDEs and as a CLI ([JetBrains agent platform description](https://blog.jetbrains.com/ai/2026/06/codex-is-now-the-recommended-agent-in-jetbrains-ai/), [Junie GA announcement](https://blog.jetbrains.com/junie/2026/06/junie-coding-agent-out-of-beta/)).

The survey label is “JetBrains AI,” which includes non-agent assistant features as well as Junie and third-party agents. That broadness lowers confidence in the exact position, but the product now clearly acts as a coding-agent harness and its survey lead over the open-source field is substantial.

### 7. OpenCode

OpenCode records 8% in Stack Overflow's chart, the strongest surveyed result for a vendor-neutral open-source terminal agent. Its official repository had approximately **187,000 GitHub stars and 23,500 forks** on the research date, reinforcing substantial developer mindshare ([OpenCode repository](https://github.com/anomalyco/opencode)). Its npm package also showed roughly 8.6 million downloads in the preceding month, although automated updates and CI make that an activity signal rather than a user count.

OpenCode is an unambiguous harness: it wraps multiple model providers with repository context, tools, sub-agents, terminal execution, and an interactive UI.

### 8. Kilo Code

Kilo Code records 6% in the Stack Overflow chart. Its official repository had approximately **26,000 GitHub stars and 2,900 forks** on the research date, and the product now spans VS Code, JetBrains, CLI, Slack, and cloud agents ([Kilo Code repository](https://github.com/Kilo-Org/kilocode), [Kilo installation surfaces](https://kilo.ai/docs/getting-started/installing)).

Kilo wins the 6% tie because both the survey result and its open-source footprint are specific to the coding-agent product. Its own “most popular open source coding agent” marketing claim is not used as evidence.

### 9. Pi

Pi records 3% in Stack Overflow's chart and has strong current open-source momentum, with approximately **72,400 GitHub stars** on the research date. It is a minimal terminal coding harness whose official architecture is directly relevant to Skill Issue: TypeScript extensions, Agent Skills, prompt templates, themes, and Pi Packages can be bundled and shared through npm or Git ([Pi repository](https://github.com/earendil-works/pi), [Pi documentation](https://pi.dev/docs/latest)).

Pi is selected ahead of the omitted higher-survey products because the project explicitly intends to support it and its first-class skills and package lifecycle make it a high-value implementation and evaluation target. Its ninth position denotes support priority within this document, rather than a claim that its measured usage exceeds Replit Agent or Warp.

## Other Candidates Outside the Support Set

- **Replit Agent**: 6% in the Stack Overflow chart and a valid engineering-agent surface, but it is outside the current Skill Issue support scope.
- **Warp**: 4% in the survey and meaningful agent adoption, but it is outside the current Skill Issue support scope.
- **Cline**: 3% in the survey and more than five million cumulative installations reported by Cline in January 2026. It is a genuine harness and a close contender, but cumulative extension installs are not equivalent to current work usage ([Cline archive containing the five-million announcement](https://cline.bot/blog/archive)).
- **Windsurf**: 3% in the survey. It is clearly a coding-agent IDE, but current product-specific adoption evidence is weaker than the selected leading group.
- **Aider, Goose, OpenHands, Continue, and Roo Code**: all have substantial open-source communities but are outside the current support scope.
- **Devin**: 2% in the survey. It is a valid autonomous coding-agent workspace, but its visible enterprise profile does not translate into a top-ten current usage result.
- **Amazon Q Developer, Google Jules, Augment Code, and Sourcegraph Cody**: valid or partially valid agentic coding products, but their comparable survey usage is lower.
- **Lovable, v0, Base44, and similar app builders**: popular, but Stack Overflow separately classifies them as no-code builder tools. Their primary workflow is idea-to-app generation rather than a general engineering harness over arbitrary repositories.
- **GitHub Agent HQ and Warp Oz**: orchestration/control planes for multiple agents. They are adjacent harness infrastructure or product layers rather than separately selected coding agents here.

## Uncertainty and Limitations

1. **No neutral global census exists.** Vendors disclose different measures—registered users, active developers, weekly users, installations, agent launches, revenue, or token volume.
2. **The survey sample is modest.** Stack Overflow's 1,100 respondents are more comparable than vendor metrics but may not represent all geographies, enterprises, students, or nontraditional builders.
3. **Multi-homing is normal.** Percentages are not market shares and should not sum to 100%; developers commonly use several harnesses.
4. **Product boundaries are converging.** IDEs host third-party agents, terminal agents gain desktop/cloud surfaces, and orchestration products wrap other harnesses.
5. **Google is mid-migration.** Treating Gemini CLI and Antigravity as one lineage avoids obvious double counting but hides remaining enterprise-specific Gemini CLI use.
6. **Open-source popularity is not usage.** Stars can reflect curiosity, and package downloads can be inflated by CI, mirrors, frequent releases, and automated updates.
7. **Selection and rank differ at the boundary.** Codex and Cursor tie in the survey, while Pi is included by project priority despite Replit Agent and Warp having higher comparable survey results.

## Bottom Line

Skill Issue's selected support set is GitHub Copilot, Claude Code, Codex, Cursor, Google's Antigravity/Gemini lineage, JetBrains AI/Junie, OpenCode, Kilo Code, and Pi. The first eight are supported by the strongest comparable popularity evidence used here; Pi is an explicit product-priority inclusion whose package and skill architecture is especially relevant to the system being built.

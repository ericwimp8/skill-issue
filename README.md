# Skill Issue

Skill Issue is a project for making agent skills easier to create, evaluate, and trust.

The name is a play on the familiar claim that a problem is a “skill issue.” In this case, the skills are the reusable instructions and workflows used by coding agents.

## The problem

Building an agent skill currently involves a great deal of uncertainty and manual refinement. When a skill fails, it can be difficult to tell what actually went wrong:

- Is the description too weak for the model to recognize when the skill applies?
- Is the skill body unclear or ineffective after the skill is invoked?
- Is the model inconsistent at selecting skills?
- Is the surrounding agent harness doing too little to support reliable skill use?

These failures often look the same from the user’s perspective. Someone can spend hours rewriting a good skill description when the model-and-harness environment is the real limitation. Others encounter inconsistent results, conclude that skills do not work, and abandon them entirely.

Skill Issue aims to make those failures observable, attributable, and actionable.

## How it works

Skill Issue has two complementary parts:

1. **Evaluate the environment.** Determine whether a particular model and agent harness can discover and invoke skills reliably.
2. **Build better skills.** Help users create, evaluate, and refine skills that behave as intended.

The model and harness are initially treated as a single practical environment. The important question for a user is whether the setup they work in—such as Codex, Claude Code, or another coding-agent environment—can use skills consistently enough to support their workflow.

Users evaluate their environment first. If it performs reliably, they can build and refine skills with confidence. If it does not, they can reconsider the environment instead of endlessly rewriting a skill that was never the underlying problem.

## Skill creation

The goal is to let users describe the outcome they want in ordinary language. For example:

> Create a skill that runs linting at the end of each task, applies automatic fixes, then finds and resolves any remaining compile-time errors.

Skill Issue should inspect the current project, understand the relevant languages and tooling, and identify important ambiguities before generating anything. If the repository contains both a TypeScript backend and a Rust project, for example, it should ask whether the skill applies to one or both rather than silently choosing the wrong scope.

Once the request is clear, the system should:

1. Generate an idiomatic skill for the user’s project and environment.
2. Create evaluations for both skill invocation and skill behavior.
3. Run those evaluations and diagnose failures.
4. Refine the description or body as appropriate.
5. Repeat until the skill meets the expected standard.
6. Deliver a ready-to-use skill with a clear account of what was validated.

The user describes the outcome. Skill Issue handles the skill-engineering work.

## Minimum viable product

The initial product will focus on local execution. Rather than operating a hosted service across every model API, Skill Issue will provide something users can install and run inside their own agent setup.

The exact interface is still being explored, but the likely shape is a small CLI or installer that:

- installs the required skills and evaluation assets;
- runs a repeatable evaluation inside the user’s configured environment;
- measures skill discovery and invocation behavior; and
- produces a useful local report explaining the results.

Much of Skill Issue may itself be implemented as agent skills. A thin CLI can handle installation, orchestration, repeatable runs, and reporting while the agent-facing workflows perform generation, evaluation, diagnosis, and refinement.

## Longer-term vision

A future hosted service could provide standardized evaluations across multiple models and APIs, publish comparisons, and present results through clear graphs and reports. Local evaluation results could also be exported or aggregated for broader comparisons.

The MVP deliberately starts with the part users can run in the environments they already use. This removes the cost and complexity of operating every provider’s API while still answering the central question: **does my current setup use skills well enough for me to depend on them?**

## Status

Skill Issue is currently in the design and early development stage. The evaluation format, local execution interface, scoring model, and division of responsibilities between the CLI and skills are still being determined.

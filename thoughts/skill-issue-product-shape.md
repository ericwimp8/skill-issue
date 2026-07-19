# Skill Issue Product Shape

## Harness Packaging

The finished product will probably need to be distributed as plugins tailored to different user environments. Likely targets include Cursor, OpenCode, Codex, Claude Code, and Pi, with other harnesses potentially added.

There is no single plugin standard across these environments, so each harness will need research into how it expects plugins or equivalent extensions to be packaged and installed. The supported set should be informed by the earlier research into harness capabilities, including sub-agent support and explicit skill handling, along with the other capability considered during that research.

The exact coverage remains a research decision. Determining which harnesses the finished product will support will probably be one of the first pieces of work.

## Standalone Tool

The plugin bundles will probably be accompanied by a standalone tool, most likely a CLI. The exact CLI technology still needs research.

The tool should work on macOS and Windows. Linux support would also be valuable, particularly if macOS and Linux can be supported together without excessive additional work. An executable distribution may be appropriate, although the current preference is a CLI rather than a Flutter application.

## Website

If time allows, a small website could provide a simple public facade for the project. It could:

- explain what Skill Issue is;
- display evaluation results;
- link to the GitHub repository;
- provide a direct CLI download; and
- show benchmark results.

The benchmarks are expected to be produced. Building the website and publishing the benchmark information there depends on available time.

Possible hosting approaches include a basic Firebase site, OpenAI Sites, or GitHub Pages. The appropriate option still needs research.

## Research Approach

Several research streams are needed across harness packaging, support coverage, CLI technology and distribution, and website hosting. These will run in parallel where their dependencies allow.

Research concurrency should remain flexible. Parallel work is preferred when the thread has capacity, but later assignments can wait for earlier agents to finish when sub-agent limits would otherwise be exceeded. Harness-packaging research can therefore run in sequential batches rather than requiring all harness researchers to be active at once.

One general research agent will first identify the ten most-used or most-popular coding-agent harnesses and produce a source-backed report. Cursor, OpenCode, Codex, Claude Code, and Pi are expected candidates, although the final ten should come from the evidence. Each resulting harness will then receive its own deep-research assignment covering its plugin or extension system, how skills, scripts, CLI tools, configuration, and assets can be installed together, and what Skill Issue would need to implement for that environment.

A separate general research agent will investigate the simplest cross-platform CLI distribution approach. The preferred result is a bundled tool that works on macOS, Windows, and Linux without requiring users to install Node.js or another separate runtime.

Website research will compare the free offerings from OpenAI Sites, GitHub, and Firebase Hosting. It will focus on direct CLI downloads, bandwidth and quotas, file restrictions, default and custom domains, and the trade-offs that determine the best free option. This deep-research campaign will use six internet researchers with at most five active concurrently.

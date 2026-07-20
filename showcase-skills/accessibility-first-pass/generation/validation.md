# Generation Validation Evidence

- **Target content hash:** `1ea4559471f22fafa21ced1ac5c925d975773fc2f2418536468d57cbf1762197`.
- **Structural validator:** `quick_validate.py showcase-skills/accessibility-first-pass/skill/accessibility-first-pass` returned `Skill is valid!` before evaluation.
- **Frontmatter:** Required `name` and `description` are present; the lowercase hyphenated name matches the folder; the description contains one purpose sentence and one use-boundary sentence.
- **Body:** Review boundary, evidence gathering, prioritization, reporting, and limitation meanings each have one semantic owner.
- **Resources:** One conditional authoritative-evidence reference and one consumed report template are required by the intake contract; all relative paths resolve.
- **OpenAI metadata:** `agents/openai.yaml` uses supported interface fields, explicitly allows implicit invocation, and contains no default prompt.
- **Formatting:** Scoped Markdown and YAML pass Prettier; repository `npm run format:check` passes.
- **Repository validation:** `npm run validate` passes formatting, lint, TypeScript checking, and the production website build. Vite emits its existing advisory that the main bundle exceeds 500 kB.
- **Privacy:** The canonical skill files, paths, and metadata contain no secrets, prohibited identities, usernames, home-directory names, or machine-specific checkout paths.
- **Runtime handoff:** All nine observable criteria were evaluated in `../evaluation/accessibility-first-pass/`; the description passed four trials and the body passed three cases without refinement.

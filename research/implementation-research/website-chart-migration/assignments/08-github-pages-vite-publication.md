# GitHub Pages And Vite Publication

## Assignment

- **Goal:** Verify the easiest correct static React/Vite publication path to GitHub Pages for the current `ericwimp8/skill-issue` repository.
- **Scope:** Current repository build, base-path, preview, validation, website-plan, GitHub workflow, Pages, environment, and remote-repository state; official Vite 8.1.5 guidance; official GitHub Pages, GitHub Actions, and GitHub-maintained action guidance.
- **Exclusions:** Workflow implementation; changing repository or Pages settings; publishing; alternate hosting providers; release automation; dynamic or server-rendered hosting; treating the existing ignored `dist/` directory as a freshly validated build.

## Sources

- Local build and package surfaces: `package.json:1-34`; `package-lock.json:1-28`; `vite.config.ts:1-10`; `index.html:1-24`; `.gitignore:1-2`; `dist/index.html:1-26`.
- Local website and repository contracts: `AGENTS.md`; `README.md`; `src/data/siteData.ts`; `plans/website/reference-and-architecture.md`; `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md`; `research/implementation-research/website-chart-migration/research-map.md`.
- Current local Git state inspected on 20 July 2026: branch `main`; remote `origin` is `https://github.com/ericwimp8/skill-issue.git`; no tracked `.github/**`, `docs/**`, `CNAME`, or `.nojekyll`; `dist/` is ignored.
- Current authenticated GitHub repository state inspected through the official GitHub CLI and REST API on 20 July 2026: `ericwimp8/skill-issue` is public; default branch `main`; `has_pages` is `false`; `GET /repos/ericwimp8/skill-issue/pages` returns `404`; no workflows or environments are present; default workflow token permission is read-only.
- [Vite 8.1.5: Deploying a Static Site](https://vite.dev/guide/static-deploy.html), especially the build, preview, and GitHub Pages sections.
- [Vite 8.1.5: Shared `base` Option](https://vite.dev/config/shared-options.html#base).
- [GitHub Docs: What is GitHub Pages?](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages).
- [GitHub Docs: Configuring a Publishing Source](https://docs.github.com/en/pages/getting-started-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site).
- [GitHub Docs: Using Custom Workflows with GitHub Pages](https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages).
- [GitHub-maintained Static Pages Starter Workflow](https://github.com/actions/starter-workflows/blob/main/pages/static.yml).
- [GitHub-maintained `setup-node` guidance](https://github.com/actions/setup-node#caching-global-packages-data).
- [GitHub-maintained `upload-pages-artifact` guidance](https://github.com/actions/upload-pages-artifact).
- [GitHub-maintained `deploy-pages` guidance](https://github.com/actions/deploy-pages).

## Findings

### Finding 1: The repository is a GitHub project site and its current production base is correct

GitHub distinguishes account sites at `<owner>.github.io` from project sites at `<owner>.github.io/<repositoryname>`. The current public repository is `ericwimp8/skill-issue`, so its default Pages URL is `https://ericwimp8.github.io/skill-issue/`. Vite requires `base: '/<REPO>/'` for that form. The current production build defaults to `/skill-issue/`, while local development uses `/`; the existing generated `dist/index.html` contains `/skill-issue/assets/...` URLs, consistent with that configuration.

**Evidence:** [GitHub's site-type table](https://docs.github.com/en/pages/getting-started-with-github-pages/what-is-github-pages#types-of-github-pages-sites); [Vite's GitHub Pages base guidance](https://vite.dev/guide/static-deploy.html#github-pages); [Vite's `base` contract](https://vite.dev/config/shared-options.html#base); `vite.config.ts:4-10`; `dist/index.html:20-21`; authenticated repository API result `nameWithOwner=ericwimp8/skill-issue`, `visibility=PUBLIC`, `defaultBranchRef.name=main`.

**Implication:** Keep the default production base `/skill-issue/` for the first project-site deployment. No deployment-only `VITE_BASE_PATH` value is required. If publication later moves to an owner root site or custom domain, set the build base to `/`; if the repository name changes, set it to the new `/<REPO>/` path and update preview behavior with it.

### Finding 2: A build-and-deploy GitHub Actions workflow is the easiest correct publication path

GitHub supports either branch publication from repository root or `/docs`, or an Actions workflow. GitHub recommends Actions when a build process other than Jekyll is required or compiled output should not live on a dedicated branch. Vite's own Pages guide states that GitHub Actions is necessary because Vite requires a build step, then provides a workflow that installs dependencies, builds, uploads `dist`, and deploys it. This matches the current repository: authored React/TypeScript sources live at the root, Vite produces ignored `dist/`, and neither root nor `/docs` contains a committed deployable site.

**Evidence:** [GitHub publishing-source guidance](https://docs.github.com/en/pages/getting-started-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site#about-publishing-sources); [Vite's current Pages workflow](https://vite.dev/guide/static-deploy.html#github-pages); `package.json:6-14`; `.gitignore:1-2`; `index.html:23`; tracked-tree inspection found no `docs/` or deployment branch content.

**Implication:** Select **GitHub Actions** under Settings → Pages → Build and deployment, then add one repository-owned Pages workflow. Publishing from `main` root would expose unbuilt sources; publishing `/docs` or `gh-pages` would add a second committed-output lifecycle. GitHub's static starter is lower-fit because it uploads an already-static directory and performs no Vite build. No `.nojekyll` file is required for the selected artifact workflow; GitHub discusses it in the separate branch-output convention.

### Finding 3: One job derived from Vite's official workflow is sufficient

The smallest repository-fit workflow can trigger on pushes to `main` and `workflow_dispatch`, check out the repository, install an explicit Node LTS, run `npm ci`, run the repository's complete validation/build command, configure Pages, upload `./dist`, and deploy the uploaded artifact. Vite's current sample keeps these steps in one job, so a cross-job `needs` dependency is unnecessary; GitHub requires `needs` when build and deployment are separated. The repository's `npm run validate` already runs format checking, lint, type checking, and the production build, and the local owner contract requires that validation before website work is presented as complete.

**Evidence:** [Vite's current single-job Pages workflow](https://vite.dev/guide/static-deploy.html#github-pages); [GitHub's custom-workflow flow](https://docs.github.com/en/pages/getting-started-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site#creating-a-custom-github-actions-workflow-to-publish-your-site); [GitHub's linked-job requirements](https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages#deploying-github-pages-artifacts); `package.json:6-14`; `AGENTS.md:18-22`; `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:93-95`.

**Implication:** Use one deploy job unless later requirements need separate build approval or artifact reuse. Replace the sample's bare `npm run build` with `npm run validate` so the deployment gate honors the repository's existing deterministic contract without duplicating commands.

### Finding 4: Explicit Pages permissions and the `github-pages` environment are required

The deployment workflow needs `contents: read` for checkout plus `pages: write` and `id-token: write` for Pages deployment and OIDC validation. The deploy job should target the `github-pages` environment and publish `steps.deployment.outputs.page_url` as its environment URL. GitHub creates the environment automatically when needed and recommends protecting it so only the default branch can deploy. The repository currently defaults workflow tokens to read-only and has no environment, so workflow-local permissions and the Pages environment are real setup requirements.

**Evidence:** [GitHub custom-workflow permissions and environment requirements](https://docs.github.com/en/pages/getting-started-with-github-pages/using-custom-workflows-with-github-pages#deploying-github-pages-artifacts); [official `deploy-pages` security considerations](https://github.com/actions/deploy-pages#security-considerations); [GitHub publishing-source environment guidance](https://docs.github.com/en/pages/getting-started-with-github-pages/configuring-a-publishing-source-for-your-github-pages-site#creating-a-custom-github-actions-workflow-to-publish-your-site); live API results `default_workflow_permissions=read` and `environments.total_count=0`.

**Implication:** Declare the three permissions explicitly and use `environment: { name: github-pages, url: ... }`. After Pages is enabled, protect that environment to permit deployment from `main`. No repository secret or personal access token is needed for the official artifact/deploy path because the actions use the workflow's `GITHUB_TOKEN` and OIDC token.

### Finding 5: A single Pages concurrency group should serialize deployment

Both Vite's workflow and GitHub's static starter use a `pages` concurrency group so only one Pages deployment proceeds at a time. They differ on cancellation policy: Vite's current guide sets `cancel-in-progress: true`, favoring the newest commit, while GitHub's static starter sets it to `false`, allowing an in-progress production deployment to finish while superseding queued runs.

**Evidence:** [Vite's current Pages workflow](https://vite.dev/guide/static-deploy.html#github-pages); [GitHub's maintained static Pages starter](https://github.com/actions/starter-workflows/blob/main/pages/static.yml).

**Implication:** Keep `group: pages`. Either cancellation value is supported official practice rather than a Pages permission requirement. For the first publication workflow, `cancel-in-progress: false` is the conservative production choice; `true` is reasonable if latest-commit speed is preferred and deployment cancellation has been accepted deliberately.

### Finding 6: `dist` is the complete deployment artifact; npm caching is optional and straightforward

Vite builds to `dist` by default, and this repository does not override `build.outDir`. The Pages workflow should upload only `./dist` with `actions/upload-pages-artifact`, not the repository root. The official action packages the artifact in Pages' required format, rejects a missing path, excludes repository metadata, and defaults artifact retention to one day. The root `package-lock.json` permits deterministic `npm ci` installation and `setup-node`'s explicit `cache: npm`; that cache covers npm's global package data rather than `node_modules`. Vite's sample already enables it. Caching `dist` would add stale-output risk without eliminating the required production build.

**Evidence:** [Vite build output guidance](https://vite.dev/guide/static-deploy.html#building-the-app); [Vite's `./dist` Pages upload](https://vite.dev/guide/static-deploy.html#github-pages); [official `upload-pages-artifact` inputs and validation](https://github.com/actions/upload-pages-artifact); [official `setup-node` npm cache behavior](https://github.com/actions/setup-node#caching-global-packages-data); `vite.config.ts:7-10`; `package-lock.json:1-28`; `.gitignore:1-2`.

**Implication:** Use `npm ci`, explicit `cache: npm`, and upload `./dist`. Do not commit or cache generated `dist`, and do not upload the repository root. At implementation time, take action versions or immutable commit SHAs from the then-current official Vite workflow because official examples advance their action majors over time.

### Finding 7: Production preview is a local verification step, not part of serving or deployment

Vite states that `vite preview` serves the built `dist` locally to check production output and is not a production server. The current `preview` script mounts at `/skill-issue/`, which is relevant for manually verifying project-site asset paths after `npm run validate`; GitHub Pages itself serves the uploaded artifact. The preview command therefore should not run as a persistent Actions step and does not replace build validation or post-deployment verification.

**Evidence:** [Vite preview purpose and limitation](https://vite.dev/guide/static-deploy.html#testing-the-app-locally); `package.json`; `README.md`; `AGENTS.md`; `plans/website/reference-and-architecture.md`.

**Implication:** Keep `npm run preview` for a local pre-publication check and verify the deployed URL after Actions completes. If `VITE_BASE_PATH` changes, the hard-coded preview `--base /skill-issue/` must change or become driven by the same value; otherwise local preview can validate a different mount point from the build.

### Finding 8: Publication infrastructure is absent before public launch

The live repository has no enabled Pages site, workflows, or environments; the local and remote tracked trees contain no `.github/workflows` deployment file. The current Vite build configuration itself is Pages-compatible, so the infrastructure gap is bounded to enabling Pages and adding the workflow. The website plan deliberately sequences deployment after real benchmark data, final public content, and qualified release assets.

**Evidence:** authenticated API results `has_pages=false`, Pages endpoint `404`, zero workflows, and zero environments; tracked-tree inspection; `src/data/siteData.ts`; `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md`; `plans/website/reference-and-architecture.md`.

**Implication:** The publication implementation needs one workflow plus the one-time GitHub Pages source selection; there is no old workflow to migrate. Before actual publication, reconcile the repository/release URLs and finish the Work Block 4 content gates. Those are launch blockers for a coherent public site, while the GitHub Pages/Vite mechanism itself is established.

## Notes

- The authenticated REST API provides strong current evidence that Pages is disabled (`has_pages=false` plus a `404` Pages-site lookup). The report does not infer future repository-setting state from that snapshot.
- The existing ignored `dist/` files were inspected only to cross-check generated base-prefixed assets. No build, validation, workflow, settings change, or deployment was run during this assignment.
- Official current examples differ in action major versions and in `cancel-in-progress`; implementation should use one internally consistent, then-current official workflow rather than mixing historical snippets.

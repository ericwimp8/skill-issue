import { siteData } from '../data/siteData';

export function ProjectPage() {
  return (
    <article className="editorial-page content-shell">
      <header className="editorial-hero section-shell">
        <p className="eyebrow">The project</p>
        <h1>{siteData.project.title}</h1>
        <p>{siteData.project.introduction}</p>
      </header>

      <section className="editorial-section editorial-lead">
        <p className="eyebrow">Motivation</p>
        <h2>{siteData.project.motivationTitle}</h2>
        {siteData.project.motivation.map((paragraph) => (
          <p key={paragraph}>{paragraph}</p>
        ))}
      </section>

      <section className="project-goals" aria-labelledby="goals-title">
        <header>
          <p className="eyebrow">Two connected goals</p>
          <h2 id="goals-title">Evaluate first. Build with clearer evidence.</h2>
        </header>
        <div>
          {siteData.project.goals.map((goal, index) => (
            <article key={goal.title}>
              <span>0{index + 1}</span>
              <h3>{goal.title}</h3>
              <p>{goal.description}</p>
            </article>
          ))}
        </div>
      </section>

      <section className="editorial-section">
        <p className="eyebrow">What exists now</p>
        <h2>{siteData.project.progressTitle}</h2>
        <p>{siteData.project.progress}</p>
        <div className="project-facts">
          {siteData.project.facts.map((fact) => (
            <div key={fact.label}>
              <strong>{fact.label}</strong>
              <span>{fact.value}</span>
            </div>
          ))}
        </div>
      </section>

      <section className="editorial-section limitations-panel">
        <p className="eyebrow">Current limits</p>
        <h2>{siteData.project.limitationsTitle}</h2>
        {siteData.project.limitations.map((paragraph) => (
          <p key={paragraph}>{paragraph}</p>
        ))}
      </section>
    </article>
  );
}

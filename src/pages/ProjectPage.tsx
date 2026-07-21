import { siteData } from '../data/siteData';

export function ProjectPage() {
  return (
    <article className="project-page content-shell">
      <div className="project-reading">
        <header className="project-introduction">
          <p className="eyebrow">The project</p>
          <h1>{siteData.project.title}</h1>
          <p>{siteData.project.introduction}</p>
        </header>

        <section className="project-section">
          <p className="eyebrow">Motivation</p>
          <h2>{siteData.project.motivationTitle}</h2>
          {siteData.project.motivation.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="project-section" aria-labelledby="goals-title">
          <p className="eyebrow">Two connected goals</p>
          <h2 id="goals-title">Evaluate first. Build with clearer evidence.</h2>
          <div className="project-goal-list">
            {siteData.project.goals.map((goal) => (
              <div className="project-goal" key={goal.title}>
                <h3>{goal.title}</h3>
                <p>{goal.description}</p>
              </div>
            ))}
          </div>
        </section>

        <section className="project-section">
          <p className="eyebrow">What exists now</p>
          <h2>{siteData.project.progressTitle}</h2>
          <p>{siteData.project.progress}</p>
          <ul className="project-facts">
            {siteData.project.facts.map((fact) => (
              <li key={fact.label}>
                <strong>{fact.label}:</strong> {fact.value}
              </li>
            ))}
          </ul>
        </section>

        <section className="project-section">
          <p className="eyebrow">Current limits</p>
          <h2>{siteData.project.limitationsTitle}</h2>
          {siteData.project.limitations.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>
      </div>
    </article>
  );
}

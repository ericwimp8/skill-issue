import { siteData } from '../data/siteData';

export function ResultsAnalysisPage() {
  return (
    <article className="analysis-page content-shell">
      <header className="editorial-hero analysis-hero section-shell">
        <span className="status-pill">
          <span aria-hidden="true" />
          Analysis coming soon
        </span>
        <p className="eyebrow">Results analysis</p>
        <h1>{siteData.analysis.title}</h1>
        <p>{siteData.analysis.introduction}</p>
        <a className="button button-secondary" href="#evaluate-environments">
          Explore the mock charts
          <span aria-hidden="true">→</span>
        </a>
      </header>

      <section className="analysis-outline" aria-label="Planned analysis">
        {siteData.analysis.sections.map((section, index) => (
          <article key={section.title}>
            <span>0{index + 1}</span>
            <h2>{section.title}</h2>
            <p>{section.description}</p>
          </article>
        ))}
      </section>

      <section className="analysis-evidence">
        <p className="eyebrow">Evidence before conclusions</p>
        <h2>{siteData.analysis.evidenceTitle}</h2>
        <p>{siteData.analysis.evidence}</p>
      </section>
    </article>
  );
}

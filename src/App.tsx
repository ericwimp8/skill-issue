import { EvaluationChart } from './components/EvaluationChart';
import { MetricCard } from './components/MetricCard';
import { ThemeToggle } from './components/ThemeToggle';
import { siteData } from './data/siteData';

export function App() {
  return (
    <div className="site-shell">
      <header className="site-header">
        <a className="brand" href="#top" aria-label="Skill Issue home">
          <span className="brand-mark" aria-hidden="true">
            S/
          </span>
          <span>Skill Issue</span>
        </a>

        <nav className="site-nav" aria-label="Primary navigation">
          <a href="#results">Results</a>
          <a href="#method">Method</a>
        </nav>

        <div className="header-actions">
          <ThemeToggle />
          <a
            className="button button-compact"
            href={siteData.release.url}
            target="_blank"
            rel="noreferrer"
          >
            {siteData.release.label}
          </a>
        </div>
      </header>

      <main id="top">
        <section className="hero section-shell" aria-labelledby="hero-title">
          <span className="status-pill">
            <span aria-hidden="true" />
            {siteData.status}
          </span>
          <p className="eyebrow">{siteData.hero.eyebrow}</p>
          <h1 id="hero-title">{siteData.hero.title}</h1>
          <p className="hero-description">{siteData.hero.description}</p>
          <div className="hero-actions">
            <a
              className="button"
              href={siteData.release.url}
              target="_blank"
              rel="noreferrer"
            >
              {siteData.release.label}
              <span aria-hidden="true">↓</span>
            </a>
            <a className="button button-secondary" href="#results">
              Explore results
              <span aria-hidden="true">→</span>
            </a>
          </div>
          <p className="release-note">{siteData.release.note}</p>
        </section>

        <section
          className="metrics-grid content-shell"
          aria-label="Benchmark summary"
        >
          {siteData.metrics.map((metric) => (
            <MetricCard key={metric.label} {...metric} />
          ))}
        </section>

        <section
          className="results section-shell content-shell"
          id="results"
          aria-labelledby="results-title"
        >
          <header className="section-heading">
            <div>
              <p className="eyebrow">Evaluation preview</p>
              <h2 id="results-title">Skill usage over context.</h2>
            </div>
            <p>
              A matched view of successful calls and missed opportunities as the
              available context is consumed.
            </p>
          </header>

          <div className="chart-grid">
            {siteData.evaluations.map((graph) => (
              <EvaluationChart key={graph.id} graph={graph} />
            ))}
          </div>
        </section>

        <section
          className="method section-shell content-shell"
          id="method"
          aria-labelledby="method-title"
        >
          <span className="method-index">01</span>
          <div>
            <p className="eyebrow">Method</p>
            <h2 id="method-title">{siteData.method.title}</h2>
          </div>
          <p>{siteData.method.description}</p>
        </section>
      </main>

      <footer className="site-footer content-shell">
        <span>{siteData.footer}</span>
        <a href={siteData.repositoryUrl} target="_blank" rel="noreferrer">
          GitHub
          <span aria-hidden="true">↗</span>
        </a>
      </footer>
    </div>
  );
}

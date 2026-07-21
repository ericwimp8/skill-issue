import { BuildSkillsShowcase } from '../components/BuildSkillsShowcase';
import { EvaluationExplorer } from '../components/EvaluationExplorer';
import { MetricCard } from '../components/MetricCard';
import { siteData } from '../data/siteData';
import type { ProductArm } from '../App';

type LandingPageProps = {
  activeArm: ProductArm;
  onSelectArm: (arm: ProductArm) => void;
};

export function LandingPage({ activeArm, onSelectArm }: LandingPageProps) {
  const activeContent = siteData.arms[activeArm];

  return (
    <>
      <section className="hero section-shell" aria-labelledby="hero-title">
        <span className="status-pill">
          <span aria-hidden="true" />
          {siteData.status}
        </span>
        <p className="eyebrow">{siteData.hero.eyebrow}</p>
        <h1 id="hero-title">
          It’s not a skill issue, but it’s always a <em>skills</em> issue.
        </h1>
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
          <a
            className="button button-secondary"
            href={siteData.repositoryUrl}
            target="_blank"
            rel="noreferrer"
          >
            View on GitHub
            <span aria-hidden="true">↗</span>
          </a>
        </div>
      </section>

      <section
        className="product-switcher-shell content-shell"
        aria-label="Skill Issue products"
      >
        <div
          className="product-switcher"
          role="tablist"
          aria-label="Choose a product path"
        >
          {(['build', 'evaluate'] as const).map((arm) => (
            <button
              id={`${arm}-tab`}
              className="product-tab"
              type="button"
              role="tab"
              aria-controls={`${arm}-panel`}
              aria-selected={activeArm === arm}
              key={arm}
              onClick={() => onSelectArm(arm)}
            >
              <span>{siteData.arms[arm].label}</span>
              <small>{siteData.arms[arm].shortLabel}</small>
            </button>
          ))}
        </div>
      </section>

      <div
        id={`${activeArm}-panel`}
        className="product-panel content-shell"
        role="tabpanel"
        aria-labelledby={`${activeArm}-tab`}
      >
        <section className="arm-introduction section-shell">
          <p className="eyebrow">{activeContent.eyebrow}</p>
          <h2>{activeContent.title}</h2>
          <p>{activeContent.description}</p>
        </section>

        {activeArm === 'build' ? (
          <BuildSkillsShowcase />
        ) : (
          <>
            <section className="metrics-grid" aria-label="Benchmark summary">
              {siteData.evaluationMetrics.map((metric) => (
                <MetricCard key={metric.label} {...metric} />
              ))}
            </section>

            <section
              className="results section-shell"
              id="results"
              aria-labelledby="results-title"
            >
              <header className="section-heading">
                <div>
                  <p className="eyebrow">Evaluation data</p>
                  <h2 id="results-title">Compare every expected call.</h2>
                </div>
                <p>
                  Compare models within supported environments or hold Codex
                  constant to isolate the harness. Every view uses the same
                  compact CLI artifact shape.
                </p>
              </header>

              <EvaluationExplorer />
            </section>

            <section
              className="method section-shell"
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
          </>
        )}
      </div>
    </>
  );
}

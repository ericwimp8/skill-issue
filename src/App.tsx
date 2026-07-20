import { useEffect, useState } from 'react';

import { BuildSkillsShowcase } from './components/BuildSkillsShowcase';
import { EvaluationExplorer } from './components/EvaluationExplorer';
import { MetricCard } from './components/MetricCard';
import { ThemeToggle } from './components/ThemeToggle';
import { siteData } from './data/siteData';

type ProductArm = 'build' | 'evaluate';

function armFromHash(): ProductArm | null {
  if (window.location.hash === '#build-skills') {
    return 'build';
  }

  if (window.location.hash === '#evaluate-environments') {
    return 'evaluate';
  }

  return null;
}

export function App() {
  const [activeArm, setActiveArm] = useState<ProductArm>(
    () => armFromHash() ?? 'evaluate',
  );

  useEffect(() => {
    const syncArm = () => {
      const arm = armFromHash();
      if (arm) {
        setActiveArm(arm);
      }
    };
    window.addEventListener('hashchange', syncArm);
    window.addEventListener('popstate', syncArm);
    return () => {
      window.removeEventListener('hashchange', syncArm);
      window.removeEventListener('popstate', syncArm);
    };
  }, []);

  const selectArm = (arm: ProductArm) => {
    const hash = arm === 'build' ? '#build-skills' : '#evaluate-environments';
    window.history.pushState(null, '', hash);
    setActiveArm(arm);
  };

  const activeContent = siteData.arms[activeArm];

  return (
    <div className="site-shell">
      <header className="site-header">
        <ThemeToggle />

        <a className="brand" href="#top" aria-label="Skill Issue home">
          <span className="brand-mark" aria-hidden="true">
            S/
          </span>
          <span>Skill Issue</span>
        </a>

        <div className="header-actions">
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
                onClick={() => selectArm(arm)}
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
                    <p className="eyebrow">Evaluation preview</p>
                    <h2 id="results-title">Compare every expected call.</h2>
                  </div>
                  <p>
                    Explore three views of the same CLI-shaped evidence. Choose
                    a scenario, then compare any model and harness combination.
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

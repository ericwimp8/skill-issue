import { useEffect, useState } from 'react';

import { ThemeToggle } from './components/ThemeToggle';
import { LandingPage } from './pages/LandingPage';
import { MethodologyPage } from './pages/MethodologyPage';
import { ProjectPage } from './pages/ProjectPage';
import { ResultsAnalysisPage } from './pages/ResultsAnalysisPage';
import { siteData } from './data/siteData';

export type ProductArm = 'build' | 'evaluate';
type Destination = 'explore' | 'method' | 'project' | 'analysis';

type RouteState = {
  destination: Destination;
  productArm: ProductArm;
};

const destinationHashes: Record<Destination, string> = {
  explore: '#evaluate-environments',
  method: '#method',
  project: '#project',
  analysis: '#analysis',
};

function routeFromHash(): RouteState {
  if (window.location.hash === '#build-skills') {
    return { destination: 'explore', productArm: 'build' };
  }

  if (window.location.hash === '#project') {
    return { destination: 'project', productArm: 'evaluate' };
  }

  if (window.location.hash === '#method') {
    return { destination: 'method', productArm: 'evaluate' };
  }

  if (window.location.hash === '#analysis') {
    return { destination: 'analysis', productArm: 'evaluate' };
  }

  return { destination: 'explore', productArm: 'evaluate' };
}

export function App() {
  const [route, setRoute] = useState<RouteState>(routeFromHash);

  useEffect(() => {
    const syncRoute = () => setRoute(routeFromHash());
    window.addEventListener('hashchange', syncRoute);
    window.addEventListener('popstate', syncRoute);
    return () => {
      window.removeEventListener('hashchange', syncRoute);
      window.removeEventListener('popstate', syncRoute);
    };
  }, []);

  function selectProductArm(productArm: ProductArm) {
    const hash =
      productArm === 'build' ? '#build-skills' : '#evaluate-environments';
    window.history.pushState(null, '', hash);
    setRoute({ destination: 'explore', productArm });
  }

  const page =
    route.destination === 'project' ? (
      <ProjectPage />
    ) : route.destination === 'method' ? (
      <MethodologyPage />
    ) : route.destination === 'analysis' ? (
      <ResultsAnalysisPage />
    ) : (
      <LandingPage
        activeArm={route.productArm}
        onSelectArm={selectProductArm}
      />
    );

  return (
    <div className="site-shell">
      <header className="site-header">
        <ThemeToggle />

        <a
          className="brand"
          href={destinationHashes.explore}
          aria-label="Skill Issue explore page"
        >
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

      <nav className="site-navigation content-shell" aria-label="Main">
        {siteData.navigation.map((item) => (
          <a
            key={item.id}
            href={destinationHashes[item.id]}
            aria-current={route.destination === item.id ? 'page' : undefined}
          >
            <span>{item.label}</span>
            <small>{item.description}</small>
          </a>
        ))}
      </nav>

      <main id="top">{page}</main>

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

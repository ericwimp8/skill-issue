import { useState } from 'react';

import { DocumentReaderDialog } from '../components/DocumentReaderDialog';
import {
  methodologyScenarios,
  methodologySkills,
  type MethodologyScenario,
  type MethodologySkill,
} from '../data/methodologyData';
import { siteData } from '../data/siteData';

type ReaderState =
  | { kind: 'scenario'; value: MethodologyScenario }
  | { kind: 'skill'; value: MethodologySkill }
  | null;

const totalExpectedCalls = methodologyScenarios.reduce(
  (total, scenario) => total + scenario.expectedCallCount,
  0,
);

function ScenarioReader({ scenario }: { scenario: MethodologyScenario }) {
  return (
    <ol className="scenario-transcript">
      {scenario.turns.map((turn, index) => (
        <li key={turn.turn_id}>
          <header>
            <span>Turn {index + 1}</span>
            <div aria-label={`Expected calls for turn ${index + 1}`}>
              {turn.expectedSkills.length > 0 ? (
                turn.expectedSkills.map((skill) => (
                  <code key={skill}>{skill}</code>
                ))
              ) : (
                <small>No expected skill call</small>
              )}
            </div>
          </header>
          <p>{turn.prompt}</p>
        </li>
      ))}
    </ol>
  );
}

export function MethodologyPage() {
  const [reader, setReader] = useState<ReaderState>(null);
  const content = siteData.methodology;

  return (
    <article className="methodology-page content-shell">
      <div className="methodology-reading">
        <header className="methodology-introduction">
          <p className="eyebrow">Evaluation methodology</p>
          <h1>{content.title}</h1>
          <p>{content.introduction}</p>
          <nav
            className="methodology-actions"
            aria-label="Methodology sections"
          >
            <a href="#method-scenarios">Browse scenarios</a>
            <a href="#method-skills">Inspect skills</a>
            <a href={siteData.repositoryUrl} target="_blank" rel="noreferrer">
              GitHub <span aria-hidden="true">↗</span>
            </a>
          </nav>
        </header>

        <dl className="methodology-summary">
          <div>
            <dt>Scenarios</dt>
            <dd>{methodologyScenarios.length}</dd>
          </div>
          <div>
            <dt>Turns per scenario</dt>
            <dd>{methodologyScenarios[0]?.turns.length ?? 0}</dd>
          </div>
          <div>
            <dt>Expected call pairs</dt>
            <dd>{totalExpectedCalls}</dd>
          </div>
          <div>
            <dt>Evaluated skills</dt>
            <dd>{methodologySkills.length}</dd>
          </div>
        </dl>

        <section className="methodology-section">
          <p className="eyebrow">Scope</p>
          <h2>{content.scopeTitle}</h2>
          {content.scope.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="methodology-section">
          <p className="eyebrow">Evaluation path</p>
          <h2>{content.processTitle}</h2>
          <ol className="methodology-process">
            {content.process.map((step) => (
              <li key={step}>{step}</li>
            ))}
          </ol>
        </section>

        <section
          className="methodology-section"
          id="method-scenarios"
          aria-labelledby="method-scenarios-title"
        >
          <p className="eyebrow">Scenario records</p>
          <h2 id="method-scenarios-title">{content.scenariosTitle}</h2>
          <p>{content.scenariosIntroduction}</p>
          <div className="methodology-record-list">
            {methodologyScenarios.map((scenario) => (
              <article key={scenario.id}>
                <div>
                  <h3>{scenario.title}</h3>
                  <p>{scenario.description}</p>
                  <small>
                    {scenario.turns.length} turns · {scenario.expectedCallCount}{' '}
                    expected calls
                  </small>
                </div>
                <div className="methodology-record-actions">
                  <button
                    type="button"
                    onClick={() =>
                      setReader({ kind: 'scenario', value: scenario })
                    }
                  >
                    Read scenario
                  </button>
                  <a href={scenario.sourceUrl} target="_blank" rel="noreferrer">
                    GitHub <span aria-hidden="true">↗</span>
                  </a>
                </div>
              </article>
            ))}
          </div>
        </section>

        <section className="methodology-section">
          <p className="eyebrow">Scoring</p>
          <h2>{content.scoringTitle}</h2>
          {content.scoring.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="methodology-section methodology-note">
          <p className="eyebrow">Manual startup call</p>
          <h2>{content.dictatePlanTitle}</h2>
          {content.dictatePlan.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="methodology-section">
          <p className="eyebrow">Instrumentation</p>
          <h2>{content.instrumentationTitle}</h2>
          {content.instrumentation.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="methodology-section">
          <p className="eyebrow">Execution controls</p>
          <h2>{content.controlsTitle}</h2>
          {content.controls.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section
          className="methodology-section"
          id="method-skills"
          aria-labelledby="method-skills-title"
        >
          <p className="eyebrow">Evaluated skills</p>
          <h2 id="method-skills-title">{content.skillsTitle}</h2>
          <p>{content.skillsIntroduction}</p>
          <div className="methodology-record-list methodology-skill-list">
            {methodologySkills.map((skill) => (
              <article key={skill.slug}>
                <div>
                  <h3>{skill.title}</h3>
                  <p>{skill.description}</p>
                  <small>
                    {skill.expectedCallCount} expected calls across the three
                    scenarios
                  </small>
                </div>
                <div className="methodology-record-actions">
                  <button
                    type="button"
                    onClick={() => setReader({ kind: 'skill', value: skill })}
                  >
                    Read skill
                  </button>
                  <a href={skill.evidenceUrl} target="_blank" rel="noreferrer">
                    Evaluation <span aria-hidden="true">↗</span>
                  </a>
                </div>
              </article>
            ))}
          </div>
          <p className="methodology-inline-link">
            <a href="#build-skills">Build and refine your own skills</a>
          </p>
        </section>

        <section className="methodology-section">
          <p className="eyebrow">Evidence</p>
          <h2>{content.evidenceTitle}</h2>
          {content.evidence.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="methodology-section">
          <p className="eyebrow">Limits</p>
          <h2>{content.limitationsTitle}</h2>
          {content.limitations.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>
      </div>

      {reader?.kind === 'scenario' ? (
        <DocumentReaderDialog
          title={reader.value.title}
          eyebrow="Governed scenario · messages and scorecard"
          sourceUrl={reader.value.sourceUrl}
          onClose={() => setReader(null)}
        >
          <ScenarioReader scenario={reader.value} />
        </DocumentReaderDialog>
      ) : null}

      {reader?.kind === 'skill' ? (
        <DocumentReaderDialog
          title={reader.value.title}
          eyebrow="Evaluated skill · SKILL.md"
          sourceUrl={reader.value.sourceUrl}
          onClose={() => setReader(null)}
        >
          <pre>{reader.value.content}</pre>
        </DocumentReaderDialog>
      ) : null}
    </article>
  );
}

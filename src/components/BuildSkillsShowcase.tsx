import { useState } from 'react';

import { DocumentReaderDialog } from './DocumentReaderDialog';
import { showcaseSkills, type ShowcaseSkill } from '../data/showcaseSkills';
import { siteData } from '../data/siteData';

export function BuildSkillsShowcase() {
  const [selectedSkill, setSelectedSkill] = useState<ShowcaseSkill | null>(
    null,
  );

  return (
    <>
      <section className="build-workflow" aria-label="Skill creation workflow">
        {siteData.buildWorkflow.map((step) => (
          <article className="workflow-card" key={step.index}>
            <span>{step.index}</span>
            <h3>{step.title}</h3>
            <p>{step.description}</p>
          </article>
        ))}
      </section>

      <section className="skill-showcase" aria-labelledby="showcase-title">
        <header className="section-heading">
          <div>
            <p className="eyebrow">{siteData.showcase.eyebrow}</p>
            <h2 id="showcase-title">{siteData.showcase.title}</h2>
          </div>
          <p>{siteData.showcase.description}</p>
        </header>

        <div className="skill-gallery">
          {showcaseSkills.map((skill, index) => (
            <button
              className="skill-card"
              type="button"
              key={skill.slug}
              onClick={() => setSelectedSkill(skill)}
            >
              <span className="skill-card-index">
                {String(index + 1).padStart(2, '0')}
              </span>
              <span className="skill-card-copy">
                <strong>{skill.title}</strong>
                <small>{skill.description}</small>
              </span>
              <span className="skill-card-action">
                Read SKILL.md <span aria-hidden="true">↗</span>
              </span>
            </button>
          ))}
        </div>
      </section>

      {selectedSkill ? (
        <DocumentReaderDialog
          title={selectedSkill.title}
          eyebrow="Generated skill · SKILL.md"
          onClose={() => setSelectedSkill(null)}
        >
          <pre>{selectedSkill.content}</pre>
        </DocumentReaderDialog>
      ) : null}
    </>
  );
}

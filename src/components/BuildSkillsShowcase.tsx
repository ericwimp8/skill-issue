import { useEffect, useState } from 'react';

import { showcaseSkills, type ShowcaseSkill } from '../data/showcaseSkills';
import { siteData } from '../data/siteData';

function SkillReader({
  skill,
  onClose,
}: {
  skill: ShowcaseSkill;
  onClose: () => void;
}) {
  useEffect(() => {
    const closeOnEscape = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        onClose();
      }
    };

    document.body.classList.add('dialog-open');
    window.addEventListener('keydown', closeOnEscape);

    return () => {
      document.body.classList.remove('dialog-open');
      window.removeEventListener('keydown', closeOnEscape);
    };
  }, [onClose]);

  return (
    <div
      className="skill-reader-backdrop"
      role="presentation"
      onMouseDown={onClose}
    >
      <section
        className="skill-reader"
        role="dialog"
        aria-modal="true"
        aria-labelledby="skill-reader-title"
        onMouseDown={(event) => event.stopPropagation()}
      >
        <header className="skill-reader-header">
          <div>
            <p className="eyebrow">Generated skill · SKILL.md</p>
            <h2 id="skill-reader-title">{skill.title}</h2>
          </div>
          <button
            className="skill-reader-close"
            type="button"
            onClick={onClose}
            autoFocus
          >
            <span className="sr-only">Close skill reader</span>
            <span aria-hidden="true">×</span>
          </button>
        </header>
        <div className="skill-reader-body">
          <pre>{skill.content}</pre>
        </div>
      </section>
    </div>
  );
}

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
        <SkillReader
          skill={selectedSkill}
          onClose={() => setSelectedSkill(null)}
        />
      ) : null}
    </>
  );
}

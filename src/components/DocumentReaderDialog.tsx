import { useEffect, useId, type ReactNode } from 'react';

type DocumentReaderDialogProps = {
  title: string;
  eyebrow: string;
  children: ReactNode;
  onClose: () => void;
  sourceUrl?: string;
  sourceLabel?: string;
};

export function DocumentReaderDialog({
  title,
  eyebrow,
  children,
  onClose,
  sourceUrl,
  sourceLabel = 'View on GitHub',
}: DocumentReaderDialogProps) {
  const titleId = useId();

  useEffect(() => {
    const activeElement = document.activeElement;
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

      if (activeElement instanceof HTMLElement) {
        activeElement.focus();
      }
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
        aria-labelledby={titleId}
        onMouseDown={(event) => event.stopPropagation()}
      >
        <header className="skill-reader-header">
          <div>
            <p className="eyebrow">{eyebrow}</p>
            <h2 id={titleId}>{title}</h2>
          </div>
          <div className="skill-reader-actions">
            {sourceUrl ? (
              <a href={sourceUrl} target="_blank" rel="noreferrer">
                {sourceLabel} <span aria-hidden="true">↗</span>
              </a>
            ) : null}
            <button
              className="skill-reader-close"
              type="button"
              onClick={onClose}
              autoFocus
            >
              <span className="sr-only">Close document reader</span>
              <span aria-hidden="true">×</span>
            </button>
          </div>
        </header>
        <div className="skill-reader-body">{children}</div>
      </section>
    </div>
  );
}

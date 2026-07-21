import { useCallback, useEffect, useId, useState, type ReactNode } from 'react';

type ModalDialogProps = {
  actions?: ReactNode;
  children: ReactNode;
  className?: string;
  eyebrow: string;
  onClose: () => void;
  title: string;
};

export function ModalDialog({
  actions,
  children,
  className = '',
  eyebrow,
  onClose,
  title,
}: ModalDialogProps) {
  const titleId = useId();
  const [isClosing, setIsClosing] = useState(false);
  const requestClose = useCallback(() => setIsClosing(true), []);

  useEffect(() => {
    const activeElement = document.activeElement;
    const closeOnEscape = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        requestClose();
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
  }, [requestClose]);

  return (
    <div
      className={`content-dialog-backdrop${isClosing ? ' is-closing' : ''}`}
      role="presentation"
      onMouseDown={requestClose}
      onAnimationEnd={(event) => {
        if (isClosing && event.target === event.currentTarget) {
          onClose();
        }
      }}
    >
      <section
        className={`content-dialog ${className}`.trim()}
        role="dialog"
        aria-modal="true"
        aria-labelledby={titleId}
        onMouseDown={(event) => event.stopPropagation()}
      >
        <header className="content-dialog-header">
          <div>
            <p className="eyebrow">{eyebrow}</p>
            <h2 id={titleId}>{title}</h2>
          </div>
          <div className="content-dialog-actions">
            {actions}
            <button
              className="content-dialog-close"
              type="button"
              onClick={requestClose}
              autoFocus
            >
              <span className="sr-only">Close dialog</span>
              <span aria-hidden="true">×</span>
            </button>
          </div>
        </header>
        <div className="content-dialog-body">{children}</div>
      </section>
    </div>
  );
}

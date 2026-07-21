import type { ReactNode } from 'react';

import { ModalDialog } from './ModalDialog';

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
  return (
    <ModalDialog
      title={title}
      eyebrow={eyebrow}
      onClose={onClose}
      actions={
        sourceUrl ? (
          <a href={sourceUrl} target="_blank" rel="noreferrer">
            {sourceLabel} <span aria-hidden="true">↗</span>
          </a>
        ) : null
      }
    >
      {children}
    </ModalDialog>
  );
}

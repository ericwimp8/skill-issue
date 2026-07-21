import { siteData } from '../data/siteData';
import { ModalDialog } from './ModalDialog';

type DownloadDialogProps = {
  onClose: () => void;
};

export function DownloadDialog({ onClose }: DownloadDialogProps) {
  return (
    <ModalDialog
      className="download-dialog"
      eyebrow="Beta CLI"
      title="Choose your platform."
      onClose={onClose}
      actions={
        <a href={siteData.release.url} target="_blank" rel="noreferrer">
          View releases <span aria-hidden="true">↗</span>
        </a>
      }
    >
      <div className="download-dialog-copy">
        <p>
          Skill Issue is in beta. All evaluation runs were completed on macOS.
        </p>
        <p>
          Cross-platform builds compile successfully. Runtime qualification is
          still in progress for Windows, Linux, and macOS Intel.
        </p>
      </div>
      <div className="download-options" aria-label="CLI downloads">
        {siteData.release.downloads.map((download) =>
          download.url ? (
            <a
              className="download-option"
              href={download.url}
              key={download.id}
            >
              <span>
                <strong>{download.platform}</strong>
                <small>{download.architecture}</small>
              </span>
              <span aria-hidden="true">↓</span>
            </a>
          ) : (
            <button
              className="download-option"
              type="button"
              key={download.id}
              disabled
            >
              <span>
                <strong>{download.platform}</strong>
                <small>{download.architecture}</small>
              </span>
              <small>Not yet published</small>
            </button>
          ),
        )}
      </div>
    </ModalDialog>
  );
}

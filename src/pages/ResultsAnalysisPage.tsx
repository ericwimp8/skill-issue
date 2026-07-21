import {
  analysisConfigurations,
  analysisTurnBands,
  campaignSummary,
  type AnalysisSummary,
} from '../data/analysisData';
import { siteData } from '../data/siteData';

function percentage(summary: AnalysisSummary) {
  return `${summary.successRate.toFixed(1)}%`;
}

export function ResultsAnalysisPage() {
  const content = siteData.analysis;
  const acceptedEvidenceUrl = `${siteData.repositoryUrl}/tree/main/evaluations/skill-calling/results/accepted`;
  const codexSol = analysisConfigurations.codexSol;
  const cursorGrok = analysisConfigurations.cursorGrok;
  const cursorComposer = analysisConfigurations.cursorComposer;
  const claudeCodeFable = analysisConfigurations.claudeCodeFable;
  const openCodeSol = analysisConfigurations.openCodeSol;
  const piSol = analysisConfigurations.piSol;

  return (
    <article className="analysis-page content-shell">
      <div className="analysis-reading">
        <header className="analysis-introduction">
          <p className="eyebrow">Results analysis</p>
          <h1>{content.title}</h1>
          <p>{content.introduction}</p>
          <nav className="analysis-actions" aria-label="Analysis links">
            <a href="#evaluate-environments">Explore the charts</a>
            <a href="#method">Read the method</a>
            <a href={acceptedEvidenceUrl} target="_blank" rel="noreferrer">
              Accepted evidence <span aria-hidden="true">↗</span>
            </a>
          </nav>
        </header>

        <dl className="analysis-summary">
          <div>
            <dt>Accepted runs</dt>
            <dd>{campaignSummary.runs}</dd>
          </div>
          <div>
            <dt>Conversation turns</dt>
            <dd>{campaignSummary.conversationTurns}</dd>
          </div>
          <div>
            <dt>Expected calls</dt>
            <dd>{campaignSummary.expected}</dd>
          </div>
          <div>
            <dt>Observed success</dt>
            <dd>{percentage(campaignSummary)}</dd>
          </div>
        </dl>

        <section className="analysis-section">
          <p className="eyebrow">Evaluation contract</p>
          <h2>{content.scopeTitle}</h2>
          {content.scope.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
          <p className="analysis-counts">
            Across the accepted evidence, {campaignSummary.called} of{' '}
            {campaignSummary.expected} expected calls were recorded and{' '}
            {campaignSummary.missed} were missed. The evaluator also recorded{' '}
            {campaignSummary.additional} additional calls outside the expected
            pairs.
          </p>
        </section>

        <section className="analysis-section">
          <p className="eyebrow">The pairing effect</p>
          <h2>{content.pairingTitle}</h2>
          {content.pairing.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
          <div
            className="analysis-comparison"
            aria-label="Codex Sol by harness"
          >
            <p>
              <strong>OpenAI Codex</strong>
              <span>
                {codexSol.called}/{codexSol.expected} · {percentage(codexSol)}
              </span>
            </p>
            <p>
              <strong>OpenCode</strong>
              <span>
                {openCodeSol.called}/{openCodeSol.expected} ·{' '}
                {percentage(openCodeSol)}
              </span>
            </p>
            <p>
              <strong>Pi</strong>
              <span>
                {piSol.called}/{piSol.expected} · {percentage(piSol)}
              </span>
            </p>
          </div>
        </section>

        <section className="analysis-section">
          <p className="eyebrow">The model effect</p>
          <h2>{content.modelTitle}</h2>
          {content.model.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
          <div className="analysis-comparison" aria-label="Models in Cursor">
            <p>
              <strong>Grok in Cursor</strong>
              <span>
                {cursorGrok.called}/{cursorGrok.expected} ·{' '}
                {percentage(cursorGrok)}
              </span>
            </p>
            <p>
              <strong>Composer in Cursor</strong>
              <span>
                {cursorComposer.called}/{cursorComposer.expected} ·{' '}
                {percentage(cursorComposer)}
              </span>
            </p>
          </div>
        </section>

        <section className="analysis-section analysis-note">
          <p className="eyebrow">Calling styles</p>
          <h2>{content.callingStyleTitle}</h2>
          {content.callingStyle.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
          <ul className="analysis-style-list">
            <li>
              <strong>Codex Sol in OpenAI Codex</strong>
              <span>
                {percentage(codexSol)} coverage · {codexSol.additional}{' '}
                additional calls
              </span>
            </li>
            <li>
              <strong>Grok in Cursor</strong>
              <span>
                {percentage(cursorGrok)} coverage · {cursorGrok.additional}{' '}
                additional calls
              </span>
            </li>
            <li>
              <strong>Composer in Cursor</strong>
              <span>
                {percentage(cursorComposer)} coverage ·{' '}
                {cursorComposer.additional} additional calls
              </span>
            </li>
          </ul>
        </section>

        <section className="analysis-section">
          <p className="eyebrow">Across the conversation</p>
          <h2>{content.timeTitle}</h2>
          {content.time.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
          <div className="analysis-band-list">
            <p>
              <strong>Composer in Cursor</strong>
              <span>
                Turns 1 to 10:{' '}
                {percentage(analysisTurnBands.cursorComposer.first)}
              </span>
              <span>
                Turns 21 to 30:{' '}
                {percentage(analysisTurnBands.cursorComposer.last)}
              </span>
            </p>
            <p>
              <strong>Claude Fable in Claude Code</strong>
              <span>
                Turns 1 to 10:{' '}
                {percentage(analysisTurnBands.claudeCodeFable.first)}
              </span>
              <span>
                Turns 21 to 30:{' '}
                {percentage(analysisTurnBands.claudeCodeFable.last)}
              </span>
            </p>
            <p>
              <strong>Codex Sol in OpenCode</strong>
              <span>
                Turns 1 to 10: {percentage(analysisTurnBands.openCodeSol.first)}
              </span>
              <span>
                Turns 21 to 30: {percentage(analysisTurnBands.openCodeSol.last)}
              </span>
            </p>
            <p>
              <strong>Codex Sol in Pi</strong>
              <span>
                Turns 1 to 10: {percentage(analysisTurnBands.piSol.first)}
              </span>
              <span>
                Turns 21 to 30: {percentage(analysisTurnBands.piSol.last)}
              </span>
            </p>
          </div>
          <p className="analysis-inline-result">
            Claude Fable recorded {claudeCodeFable.called} expected calls across
            the complete scenario set. Composer recorded {cursorComposer.called}
            . OpenCode recorded {openCodeSol.called}, and Pi recorded{' '}
            {piSol.called}.
          </p>
        </section>

        <section className="analysis-section">
          <p className="eyebrow">Validity checks</p>
          <h2>{content.validityTitle}</h2>
          {content.validity.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="analysis-section">
          <p className="eyebrow">Practical meaning</p>
          <h2>{content.meaningTitle}</h2>
          {content.meaning.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>

        <section className="analysis-section">
          <p className="eyebrow">Limits</p>
          <h2>{content.limitationsTitle}</h2>
          {content.limitations.map((paragraph) => (
            <p key={paragraph}>{paragraph}</p>
          ))}
        </section>
      </div>
    </article>
  );
}

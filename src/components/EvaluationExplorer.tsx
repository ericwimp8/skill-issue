import { useState } from 'react';

import {
  availableCells,
  defaultCellIds,
  evaluationResults,
  scenarioOptions,
} from '../data/evaluationData';
import { CumulativeCallsChart } from './charts/CumulativeCallsChart';
import { chartColorForCell } from './charts/chartTheme';
import { HarnessOutcomeChart } from './charts/HarnessOutcomeChart';
import { OutcomeBarsChart } from './charts/OutcomeBarsChart';
import { TurnRasterChart } from './charts/TurnRasterChart';

export function EvaluationExplorer() {
  const [comparisonView, setComparisonView] = useState<'models' | 'harnesses'>(
    'models',
  );
  const [scenarioId, setScenarioId] = useState<string>(scenarioOptions[0].id);
  const [selectedCellIds, setSelectedCellIds] = useState<string[]>([
    ...defaultCellIds,
  ]);
  const [rankingScenarioIds, setRankingScenarioIds] = useState<string[]>(
    scenarioOptions.map((scenario) => scenario.id),
  );
  const [searchQuery, setSearchQuery] = useState('');
  const [harnessScenarioIds, setHarnessScenarioIds] = useState<string[]>(
    scenarioOptions.map((scenario) => scenario.id),
  );

  const scenarioResults = evaluationResults.filter(
    (result) => result.scenario_id === scenarioId,
  );
  const visibleResults = scenarioResults.filter((result) =>
    selectedCellIds.includes(result.cellId),
  );
  const rankingResults = evaluationResults.filter(
    (result) =>
      selectedCellIds.includes(result.cellId) &&
      rankingScenarioIds.includes(result.scenario_id),
  );
  const searchableCells = availableCells.filter((cell) =>
    cell.label.toLowerCase().includes(searchQuery.trim().toLowerCase()),
  );

  function toggleCell(cellId: string) {
    setSelectedCellIds((current) =>
      current.includes(cellId)
        ? current.filter((id) => id !== cellId)
        : [...current, cellId],
    );
  }

  function resetSelection() {
    setSelectedCellIds([...defaultCellIds]);
    setScenarioId(scenarioOptions[0].id);
    setRankingScenarioIds(scenarioOptions.map((scenario) => scenario.id));
    setSearchQuery('');
  }

  function toggleRankingScenario(nextScenarioId: string) {
    setRankingScenarioIds((current) => {
      if (!current.includes(nextScenarioId)) {
        return [...current, nextScenarioId];
      }

      return current.length === 1
        ? current
        : current.filter((id) => id !== nextScenarioId);
    });
  }

  function toggleHarnessScenario(nextScenarioId: string) {
    setHarnessScenarioIds((current) => {
      if (!current.includes(nextScenarioId)) {
        return [...current, nextScenarioId];
      }

      return current.length === 1
        ? current
        : current.filter((id) => id !== nextScenarioId);
    });
  }

  return (
    <div className="evaluation-explorer">
      <div className="comparison-view-tabs" role="tablist">
        <button
          type="button"
          role="tab"
          aria-selected={comparisonView === 'models'}
          onClick={() => setComparisonView('models')}
        >
          <span>Model comparison</span>
          <small>Compare supported combinations</small>
        </button>
        <button
          type="button"
          role="tab"
          aria-selected={comparisonView === 'harnesses'}
          onClick={() => setComparisonView('harnesses')}
        >
          <span>Harness comparison</span>
          <small>Hold Codex and Medium constant</small>
        </button>
      </div>

      {comparisonView === 'harnesses' ? (
        <HarnessOutcomeChart
          results={evaluationResults.filter(
            (result) =>
              result.model === 'codex-sol' &&
              harnessScenarioIds.includes(result.scenario_id),
          )}
          selectedScenarioIds={harnessScenarioIds}
          onToggleScenario={toggleHarnessScenario}
        />
      ) : (
        <>
          <div
            className="explorer-controls"
            aria-label="Evaluation chart controls"
          >
            <label className="filter-control">
              <span>Scenario</span>
              <select
                value={scenarioId}
                onChange={(event) => setScenarioId(event.target.value)}
              >
                {scenarioOptions.map((scenario) => (
                  <option key={scenario.id} value={scenario.id}>
                    {scenario.label}
                  </option>
                ))}
              </select>
            </label>

            <details className="combination-picker">
              <summary>
                <span>
                  <small>Compare</small>
                  {selectedCellIds.length} of {availableCells.length}{' '}
                  combinations
                </span>
                <i aria-hidden="true">⌄</i>
              </summary>
              <div className="picker-panel">
                <label className="picker-search">
                  <span className="sr-only">Search combinations</span>
                  <input
                    type="search"
                    value={searchQuery}
                    onChange={(event) => setSearchQuery(event.target.value)}
                    placeholder="Search harnesses or models"
                  />
                </label>
                <div className="picker-options">
                  {searchableCells.map((cell) => (
                    <label key={cell.id} className="picker-option">
                      <input
                        type="checkbox"
                        checked={selectedCellIds.includes(cell.id)}
                        onChange={() => toggleCell(cell.id)}
                      />
                      <span>
                        <strong>{cell.modelLabel}</strong>
                        <small>
                          {cell.harnessLabel} · {cell.reasoningLabel} reasoning
                        </small>
                      </span>
                      <i
                        aria-hidden="true"
                        style={{ backgroundColor: chartColorForCell(cell.id) }}
                      />
                    </label>
                  ))}
                </div>
                <div className="picker-actions">
                  <button type="button" onClick={() => setSelectedCellIds([])}>
                    Clear
                  </button>
                  <button
                    type="button"
                    onClick={() =>
                      setSelectedCellIds(availableCells.map((cell) => cell.id))
                    }
                  >
                    Select all
                  </button>
                  <button type="button" onClick={resetSelection}>
                    Reset defaults
                  </button>
                </div>
              </div>
            </details>
          </div>

          <div className="selection-summary">
            <span>
              Showing <strong>{visibleResults.length}</strong> comparison
              {visibleResults.length === 1 ? '' : 's'}
            </span>
            <div className="selection-chips">
              {visibleResults.map((result) => (
                <span key={result.cellId}>
                  <i
                    style={{
                      backgroundColor: chartColorForCell(result.cellId),
                    }}
                  />
                  {result.cellLabel}
                </span>
              ))}
            </div>
          </div>

          {visibleResults.length === 0 ? (
            <div className="empty-results">
              <h3>No combinations selected.</h3>
              <p>Choose at least one harness/model combination to compare.</p>
              <button type="button" onClick={resetSelection}>
                Restore defaults
              </button>
            </div>
          ) : (
            <div className="chart-explorations">
              <TurnRasterChart results={visibleResults} />
              <CumulativeCallsChart results={visibleResults} />
              <OutcomeBarsChart
                results={rankingResults}
                selectedScenarioIds={rankingScenarioIds}
                onToggleScenario={toggleRankingScenario}
              />
            </div>
          )}
        </>
      )}
    </div>
  );
}

import { useMemo, useState } from 'react';

import {
  availableCells,
  defaultCellIds,
  evaluationResults,
  scenarioOptions,
} from '../data/evaluationData';
import { CheckpointBarsChart } from './charts/CheckpointBarsChart';
import { CumulativeCallsChart } from './charts/CumulativeCallsChart';
import { chartColorForCell } from './charts/chartTheme';
import { OutcomeBarsChart } from './charts/OutcomeBarsChart';
import { TurnRasterChart } from './charts/TurnRasterChart';

const allFilter = 'all';

export function EvaluationExplorer() {
  const [scenarioId, setScenarioId] = useState<string>(scenarioOptions[0].id);
  const [selectedCellIds, setSelectedCellIds] = useState<string[]>([
    ...defaultCellIds,
  ]);
  const [harnessFilter, setHarnessFilter] = useState(allFilter);
  const [modelFilter, setModelFilter] = useState(allFilter);
  const [searchQuery, setSearchQuery] = useState('');

  const harnessOptions = useMemo(
    () =>
      Array.from(
        new Map(
          availableCells.map((cell) => [
            cell.harness,
            { id: cell.harness, label: cell.harnessLabel },
          ]),
        ).values(),
      ),
    [],
  );
  const modelOptions = useMemo(
    () =>
      Array.from(
        new Map(
          availableCells.map((cell) => [
            cell.model,
            { id: cell.model, label: cell.modelLabel },
          ]),
        ).values(),
      ),
    [],
  );
  const scenarioResults = evaluationResults.filter(
    (result) => result.scenario_id === scenarioId,
  );
  const visibleResults = scenarioResults.filter(
    (result) =>
      selectedCellIds.includes(result.cellId) &&
      (harnessFilter === allFilter || result.harness === harnessFilter) &&
      (modelFilter === allFilter || result.model === modelFilter),
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
    setHarnessFilter(allFilter);
    setModelFilter(allFilter);
    setSearchQuery('');
  }

  return (
    <div className="evaluation-explorer">
      <div className="results-notice">
        <span>Illustrative layout data</span>
        <p>These values demonstrate the report shape, not observed results.</p>
      </div>

      <div className="explorer-controls" aria-label="Evaluation chart controls">
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
              {selectedCellIds.length} of {availableCells.length} combinations
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
                    <small>{cell.harnessLabel}</small>
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

        <label className="filter-control filter-control-compact">
          <span>Harness</span>
          <select
            value={harnessFilter}
            onChange={(event) => setHarnessFilter(event.target.value)}
          >
            <option value={allFilter}>All harnesses</option>
            {harnessOptions.map((harness) => (
              <option key={harness.id} value={harness.id}>
                {harness.label}
              </option>
            ))}
          </select>
        </label>

        <label className="filter-control filter-control-compact">
          <span>Model</span>
          <select
            value={modelFilter}
            onChange={(event) => setModelFilter(event.target.value)}
          >
            <option value={allFilter}>All models</option>
            {modelOptions.map((model) => (
              <option key={model.id} value={model.id}>
                {model.label}
              </option>
            ))}
          </select>
        </label>
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
                style={{ backgroundColor: chartColorForCell(result.cellId) }}
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
          <div className="chart-explorations-grid">
            <CumulativeCallsChart results={visibleResults} />
            <CheckpointBarsChart results={visibleResults} />
          </div>
          <OutcomeBarsChart results={visibleResults} />
        </div>
      )}
    </div>
  );
}

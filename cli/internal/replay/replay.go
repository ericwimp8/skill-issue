package replay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type HarnessID string

const (
	HarnessClaude   HarnessID = "claude-code"
	HarnessCodex    HarnessID = "codex"
	HarnessCursor   HarnessID = "cursor"
	HarnessOpenCode HarnessID = "opencode"
	HarnessKilo     HarnessID = "kilo-code"
	HarnessPi       HarnessID = "pi"
)

var ErrProtocol = errors.New("malformed harness protocol")

type Scenario struct {
	SchemaVersion int    `json:"schema_version"`
	ID            string `json:"scenario_id"`
	Turns         []Turn `json:"turns"`
}

type Turn struct {
	ID     string `json:"turn_id"`
	Prompt string `json:"prompt"`
}

func (scenario Scenario) Validate() error {
	if scenario.SchemaVersion != 1 {
		return errors.New("scenario schema version must be 1")
	}
	if strings.TrimSpace(scenario.ID) == "" {
		return errors.New("scenario ID is required")
	}
	if len(scenario.Turns) == 0 {
		return errors.New("scenario requires at least one turn")
	}
	seen := make(map[string]struct{}, len(scenario.Turns))
	for index, turn := range scenario.Turns {
		if strings.TrimSpace(turn.ID) == "" {
			return fmt.Errorf("turn %d ID is required", index+1)
		}
		if _, exists := seen[turn.ID]; exists {
			return fmt.Errorf("duplicate turn ID %q", turn.ID)
		}
		seen[turn.ID] = struct{}{}
	}
	return nil
}

type Capture struct {
	SessionID  string            `json:"session_id,omitempty"`
	Transcript string            `json:"transcript"`
	Stderr     string            `json:"stderr,omitempty"`
	Events     []json.RawMessage `json:"events"`
}

type TurnResult struct {
	TurnID  string  `json:"turn_id"`
	Capture Capture `json:"capture"`
}

type Result struct {
	HarnessID HarnessID    `json:"harness_id"`
	Scenario  Scenario     `json:"scenario"`
	Turns     []TurnResult `json:"turns"`
}

type Session interface {
	SendPrompt(ctx context.Context, prompt string) error
	Wait(ctx context.Context) (Capture, error)
	Close() error
}

type Adapter interface {
	HarnessID() HarnessID
	Start(ctx context.Context) (Session, error)
}

type BoundaryPhase string

const (
	BoundaryBefore BoundaryPhase = "before"
	BoundaryAfter  BoundaryPhase = "after"
)

type Boundary struct {
	ScenarioID string
	TurnID     string
	Phase      BoundaryPhase
	Prompt     string
	Capture    *Capture
}

type BoundaryCallback func(ctx context.Context, boundary Boundary) error

type Runner struct {
	Adapter    Adapter
	OnBoundary BoundaryCallback
}

func (runner Runner) Run(ctx context.Context, scenario Scenario) (result Result, err error) {
	if runner.Adapter == nil {
		return Result{}, errors.New("replay adapter is required")
	}
	if err := scenario.Validate(); err != nil {
		return Result{}, err
	}

	session, err := runner.Adapter.Start(ctx)
	if err != nil {
		return Result{}, fmt.Errorf("start %s session: %w", runner.Adapter.HarnessID(), err)
	}
	defer func() {
		if closeErr := session.Close(); err == nil && closeErr != nil {
			err = fmt.Errorf("close %s session: %w", runner.Adapter.HarnessID(), closeErr)
		}
	}()

	result = Result{HarnessID: runner.Adapter.HarnessID(), Scenario: scenario, Turns: make([]TurnResult, 0, len(scenario.Turns))}
	for _, turn := range scenario.Turns {
		before := Boundary{ScenarioID: scenario.ID, TurnID: turn.ID, Phase: BoundaryBefore, Prompt: turn.Prompt}
		if err := runner.notify(ctx, before); err != nil {
			return Result{}, err
		}
		if err := session.SendPrompt(ctx, turn.Prompt); err != nil {
			return Result{}, fmt.Errorf("send turn %q: %w", turn.ID, err)
		}
		capture, err := session.Wait(ctx)
		if err != nil {
			return Result{}, fmt.Errorf("wait for turn %q: %w", turn.ID, err)
		}
		after := Boundary{ScenarioID: scenario.ID, TurnID: turn.ID, Phase: BoundaryAfter, Prompt: turn.Prompt, Capture: &capture}
		if err := runner.notify(ctx, after); err != nil {
			return Result{}, err
		}
		result.Turns = append(result.Turns, TurnResult{TurnID: turn.ID, Capture: capture})
	}
	return result, nil
}

func (runner Runner) notify(ctx context.Context, boundary Boundary) error {
	if runner.OnBoundary == nil {
		return nil
	}
	if err := runner.OnBoundary(ctx, boundary); err != nil {
		return fmt.Errorf("%s boundary for turn %q: %w", boundary.Phase, boundary.TurnID, err)
	}
	return nil
}

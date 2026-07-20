package lifecycle

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
	"github.com/ericwimp8/skill-issue/cli/internal/replay"
	"golang.org/x/term"
)

// turnProgressRenderer turns evaluation progress callbacks into terminal
// output. On an interactive terminal it animates a spinner with elapsed time
// while a turn runs; on a plain writer it emits only the start and finish
// lines, so piped and logged output stays clean.
type turnProgressRenderer struct {
	writer      io.Writer
	interactive bool
	mutex       sync.Mutex
	spinnerStop chan struct{}
	spinnerDone chan struct{}
}

func newTurnProgressRenderer(writer io.Writer) *turnProgressRenderer {
	interactive := false
	if file, ok := writer.(*os.File); ok {
		interactive = term.IsTerminal(int(file.Fd()))
	}
	return &turnProgressRenderer{writer: writer, interactive: interactive}
}

// handle is the evaluation.RunRequest Progress callback. Boundary callbacks
// arrive sequentially from the run loop, so spinner bookkeeping needs no
// locking; only writes to the shared writer are synchronized with the
// spinner goroutine.
func (renderer *turnProgressRenderer) handle(progress evaluation.TurnProgress) {
	if progress.Phase == replay.BoundaryBefore {
		renderer.stopSpinner()
		renderer.printLine(fmt.Sprintf("Starting turn %d of %d: %s", progress.Index, progress.Total, progress.TurnID))
		renderer.startSpinner(fmt.Sprintf("turn %d of %d running", progress.Index, progress.Total))
		return
	}
	renderer.stopSpinner()
	renderer.printLine(fmt.Sprintf("Finished turn %d of %d: %s (%s, %s, %s)",
		progress.Index, progress.Total, progress.TurnID, formatTurnDuration(progress.Duration),
		pluralize(progress.HarnessEvents, "harness event"), pluralize(progress.SkillCalls, "skill call")))
}

func pluralize(count int, singular string) string {
	if count == 1 {
		return fmt.Sprintf("%d %s", count, singular)
	}
	return fmt.Sprintf("%d %ss", count, singular)
}

// stop clears any running spinner; safe to call when none is running. Callers
// defer it so an interrupted or failed run does not leave a stale spinner
// line in front of the error output.
func (renderer *turnProgressRenderer) stop() {
	renderer.stopSpinner()
}

func (renderer *turnProgressRenderer) startSpinner(label string) {
	if !renderer.interactive {
		return
	}
	stop := make(chan struct{})
	done := make(chan struct{})
	renderer.spinnerStop, renderer.spinnerDone = stop, done
	startedAt := time.Now()
	go func() {
		defer close(done)
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		ticker := time.NewTicker(120 * time.Millisecond)
		defer ticker.Stop()
		for index := 0; ; index++ {
			select {
			case <-stop:
				renderer.mutex.Lock()
				fmt.Fprint(renderer.writer, "\r\x1b[2K")
				renderer.mutex.Unlock()
				return
			case <-ticker.C:
				renderer.mutex.Lock()
				fmt.Fprintf(renderer.writer, "\r\x1b[2K%s %s (%s)", frames[index%len(frames)], label, formatTurnDuration(time.Since(startedAt)))
				renderer.mutex.Unlock()
			}
		}
	}()
}

func (renderer *turnProgressRenderer) stopSpinner() {
	if renderer.spinnerStop == nil {
		return
	}
	close(renderer.spinnerStop)
	<-renderer.spinnerDone
	renderer.spinnerStop, renderer.spinnerDone = nil, nil
}

func (renderer *turnProgressRenderer) printLine(text string) {
	renderer.mutex.Lock()
	defer renderer.mutex.Unlock()
	fmt.Fprintln(renderer.writer, text)
}

func formatTurnDuration(duration time.Duration) string {
	if duration < time.Second {
		return duration.Round(time.Millisecond).String()
	}
	return duration.Round(time.Second).String()
}

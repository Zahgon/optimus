package progressbar

import (
	"io"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/schollz/progressbar/v3"
)

const (
	progressBarWidth           = 15
	progressBarRefreshDuration = 120 * time.Millisecond
)

// ProgressBar defines custom progress bar
type ProgressBar struct {
	spinner *spinner.Spinner
	bar     *progressbar.ProgressBar

	mu     sync.Mutex
	writer io.Writer
}

// NewProgressBar initializes default progress bar
func NewProgressBar() *ProgressBar { _ = "STUB: not implemented"; return nil }

// NewProgressBarWithWriter initializes progress bar with writer
func NewProgressBarWithWriter(w io.Writer) *ProgressBar { _ = "STUB: not implemented"; return nil }

// Start starts the progress bar with label
func (p *ProgressBar) Start(label string) { _ = "STUB: not implemented"; return }

func (p *ProgressBar) StartNewLine(label string) { _ = "STUB: not implemented"; return }

// StartProgress starts progress bar with count and label
func (p *ProgressBar) StartProgress(count int, label string) { _ = "STUB: not implemented"; return }

// SetProgress sets the progress status
// StartProgress should be called before setting status
func (p *ProgressBar) SetProgress(idx int) error { _ = "STUB: not implemented"; return nil }

// Stop stops progress bar
func (p *ProgressBar) Stop() { _ = "STUB: not implemented"; return }

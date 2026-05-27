package window

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentCursor cursorPointer

	truncateTo  truncateTo
	sizeInput   textinput.Model
	offsetInput textinput.Model

	scheduledTime time.Time
}

func newModel() *model { _ = "STUB: not implemented"; return nil }

func (*model) Init() tea.Cmd {
	_ = "STUB: not implemented"
	// this method is to adhere to library contract
	return *new(tea.Cmd)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m *model) View() string { _ = "STUB: not implemented"; return "" }

func (m *model) generateWindowResultView() string { _ = "STUB: not implemented"; return "" }

//nolint: gomnd

func (m *model) generateWindowInputHintView() string { _ = "STUB: not implemented"; return "" }

func (m *model) generateWindowInputView() string { _ = "STUB: not implemented"; return "" }

func (m *model) generateWindowTableRowView(version int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (m *model) generateSechduledTimeView() string { _ = "STUB: not implemented"; return "" }

func (m *model) generateValueWithCursorPointerView(targetCursor cursorPointer, value string) string {
	_ = "STUB: not implemented"
	return ""
}

func (m *model) handleInput(msg tea.Msg) { _ = "STUB: not implemented"; return }

func (m *model) handleDecrement() { _ = "STUB: not implemented"; return }

func (m *model) decrementScheduledTime() { _ = "STUB: not implemented"; return }

func (m *model) decrementTruncateTo() { _ = "STUB: not implemented"; return }

func (m *model) handleIncrement() { _ = "STUB: not implemented"; return }

func (m *model) incrementScheduledTime() { _ = "STUB: not implemented"; return }

func (m *model) incrementTruncateTo() { _ = "STUB: not implemented"; return }

func (m *model) handleRight() { _ = "STUB: not implemented"; return }

func (m *model) handleLeft() { _ = "STUB: not implemented"; return }

func (m *model) handleDown() { _ = "STUB: not implemented"; return }

func (m *model) handleUp() { _ = "STUB: not implemented"; return }

package unit

import (
	"github.com/WarrenWu4/bananatype/pkg/timer"
	"testing"
	"time"

	bubbleTimer "github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

func TestTimerActivation(t *testing.T) {
	tm := timer.NewTimerModel(15 * time.Second)
	if tm.Done {
		t.Error("Timer should not be done initially")
	}
	msg := tea.KeyMsg{Runes: []rune("a"), Type: tea.KeyRunes}
	updatedModel, cmd := tm.Update(msg)
	tm = updatedModel.(timer.TimerModel)
	if cmd == nil {
		t.Error("Timer should return a command (init) when started")
	}
}

func TestTimerTimeout(t *testing.T) {
	tm := timer.NewTimerModel(15 * time.Second)
	msg := bubbleTimer.TimeoutMsg{}
	updatedModel, _ := tm.Update(msg)
	tm = updatedModel.(timer.TimerModel)
	if !tm.Done {
		t.Error("Timer should be done after TimeoutMsg")
	}
}

package progress

import (
	colors "bananas/pkg/colors"
	settings "bananas/pkg/settings"
	timer "bananas/pkg/timer"
	typer "bananas/pkg/typer"
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type ProgressModel struct {
	Settings  settings.SettingsModel
	Typer     typer.TyperModel
	Timer     timer.TimerModel
	StartTime time.Time
	DoneTime  time.Time
	Done      bool
}

func NewProgressModel(s settings.SettingsModel, t typer.TyperModel) ProgressModel {
	return ProgressModel{
		Settings: s,
		Typer:    t,
		Timer:    timer.NewTimerModel(time.Second * time.Duration(s.ActiveTime)),
		Done:     false,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return nil
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// ProgressModel is mostly a view-only model that receives updates
	// to its child models from the main loop.
	// send a message when update is complete though
	switch msg.(type) {
	case tea.KeyMsg:
		if m.StartTime.IsZero() {
			m.StartTime = time.Now()
		}
	}

	switch m.Settings.ActiveTyperMode {
	case "timer":
		if m.Timer.Done {
			m.Done = true
			m.DoneTime = time.Now()
		}
	case "words":
		if m.Typer.TotalWords >= m.Settings.ActiveWords {
			m.Done = true
			m.DoneTime = time.Now()
		}
	}
	return m, nil
}

func (m ProgressModel) View() string {
	switch m.Settings.ActiveTyperMode {
	case "timer":
		return m.Timer.View()
	case "words":
		return colors.Yellow.Render(fmt.Sprintf("%d/%d", m.Typer.TotalWords, m.Settings.ActiveWords))
	default:
		return ""
	}
}

func (m ProgressModel) Reset() ProgressModel {
	m.Done = false
	m.StartTime = time.Time{}
	m.Timer = timer.NewTimerModel(time.Second * time.Duration(m.Settings.ActiveTime))
	return m
}

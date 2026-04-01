package main

import (
	analysis "bananas/pkg/analysis"
	logger "bananas/pkg/logger"
	"bananas/pkg/progress"
	settings "bananas/pkg/settings"
	"bananas/pkg/timer"
	typer "bananas/pkg/typer"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Build        = "dev"
	wordPath     = "./resources/word_bank.txt"
	settingsPath = "./resources/settings.json"
)

type MainModel struct {
	settings settings.SettingsModel
	progress progress.ProgressModel
	typer    typer.TyperModel
	analysis analysis.AnalysisModel
	width    int
	height   int
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// global updates that happen regardless of current view
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.settings.Show = !m.settings.Show
			return m, nil
		}
	case settings.SettingsModel:
		m.typer = typer.NewTyper()
		m.progress = m.progress.Reset()
		m.settings.Show = false
	case analysis.AnalysisModel:
		m.typer = typer.NewTyper()
		m.progress.Typer = m.typer
		m.progress = m.progress.Reset()
		return m, nil
	}
	// local updates that are dependent on which view is active
	if m.settings.Show { // only update settings when settings show
		updatedSettings, settingsCmd := m.settings.Update(msg)
		m.settings = updatedSettings.(settings.SettingsModel)
		updatedProgress, progressCmd := m.progress.Update(msg)
		m.progress = updatedProgress.(progress.ProgressModel)
		m.progress.Settings = m.settings
		return m, tea.Batch(settingsCmd, progressCmd)
	} else if m.progress.Done { // only update analysis when timer is done
		updatedAnalysis, analysisCmd := m.analysis.Update(msg)
		m.analysis = updatedAnalysis.(analysis.AnalysisModel)
		return m, analysisCmd
	}
	// otherwise update timer and typer
	updatedProgress, progressCmd := m.progress.Update(msg)
	m.progress = updatedProgress.(progress.ProgressModel)
	updatedTyper, typerCmd := m.typer.Update(msg)
	m.typer = updatedTyper.(typer.TyperModel)
	m.progress.Typer = m.typer
	UpdatedTimer, timerCmd := m.progress.Timer.Update(msg)
	m.progress.Timer = UpdatedTimer.(timer.TimerModel)
	return m, tea.Batch(progressCmd, typerCmd, timerCmd)
}

func (m MainModel) View() string {
	output := ""
	paddingY := (m.height - typer.MAXLINES + 1) / 2
	paddingX := (m.width - typer.MAXCHARPERLINE) / 2
	// top padding
	output += strings.Repeat("\n", paddingY)
	// left padding
	if m.progress.Done {
		m.analysis.Time = m.settings.ActiveTime
		m.analysis.Words = m.typer.TotalWords
		m.analysis.Correct = m.typer.TotalCorrect
		m.analysis.Characters = m.typer.TotalTyped
		outputLines := strings.Split(m.analysis.View(), "\n")
		for i := 0; i < len(outputLines); i++ {
			output += strings.Repeat(" ", paddingX) + outputLines[i] + "\n"
		}
	} else if m.settings.Show {
		outputLines := strings.Split(m.settings.View(), "\n")
		for _, line := range outputLines {
			output += strings.Repeat(" ", paddingX) + line + "\n"
		}
	} else {
		output += strings.Repeat(" ", paddingX) + m.progress.View() + "\n"
		outputLines := strings.Split(m.typer.View(), "\n")
		for i := 0; i < len(outputLines); i++ {
			output += strings.Repeat(" ", paddingX) + outputLines[i] + "\n"
		}
	}
	return output
}

func setup() MainModel {
	// initialize main model
	s := settings.NewSettingsModel()
	ty := typer.NewTyper()
	a := analysis.NewAnalysisModel()
	return MainModel{
		progress: progress.NewProgressModel(s, ty),
		typer:    ty,
		analysis: a,
		settings: s,
		width:    120,
		height:   8,
	}
}

func main() {
	if Build == "prod" {
		wordPath = "/usr/share/bananatype/word_bank.txt"
		settingsPath = os.Getenv("HOME") + "/.local/state/bananatype/settings.json"
		logger.InitLogger(os.Getenv("HOME") + "/.local/state/banantype/log.txt")
	}
	os.OpenFile(settingsPath, os.O_APPEND|os.O_CREATE, 0644)
	// run this function to initialize important shit
	initialModel := setup()
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting game:", err)
		os.Exit(1)
	}
}

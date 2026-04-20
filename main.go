package main

import (
	"github.com/WarrenWu4/bananatype/pkg/coordinator"
	logger "github.com/WarrenWu4/bananatype/pkg/logger"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	Build        = "dev"
	wordPath     = "./resources/word_bank.txt"
	settingsPath = "./resources/settings.json"
)

func main() {
	if Build == "prod" {
		wordPath = "/usr/share/bananatype/word_bank.txt"
		settingsPath = os.Getenv("HOME") + "/.local/state/bananatype/settings.json"
		logger.InitLogger(os.Getenv("HOME") + "/.local/state/banantype/log.txt")
	}
	os.OpenFile(settingsPath, os.O_APPEND|os.O_CREATE, 0644)
	// run this function to initialize important shit
	initialModel := coordinator.NewMainModel()
	p := tea.NewProgram(initialModel, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting game:", err)
		os.Exit(1)
	}
}

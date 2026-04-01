// setting screen for bananas
// core functionality: change time control

package settings

import (
	colors "bananas/pkg/colors"
	resourcepath "bananas/pkg/resourcepath"
	"encoding/json"
	"os"
	"slices"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

const settingInstructions = "UP/DOWN/LEFT/RIGHT to move\nENTER to select\nESC to close settings page"

type SettingsModel struct {
	Show            bool
	options         []string
	optionIdx       int
	times           []int
	timeIdx         int
	words           []int
	wordIdx         int
	ActiveTime      int
	ActiveWords     int
	ActiveTyperMode string
}

func NewSettingsModel() SettingsModel {
	s := SettingsModel{
		Show:            false,
		options:         []string{"timer", "words", "restart", "quit"},
		optionIdx:       0,
		times:           []int{15, 30, 60, 120},
		timeIdx:         1,
		words:           []int{10, 25, 50, 100},
		wordIdx:         2,
		ActiveTime:      30,
		ActiveWords:     50,
		ActiveTyperMode: "timer",
	}
	readSettings(&s)
	return s
}

func (m SettingsModel) Init() tea.Cmd {
	return nil
}

func (m SettingsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.optionIdx = (m.optionIdx - 1 + len(m.options)) % len(m.options)
		case "down":
			m.optionIdx = (m.optionIdx + 1) % len(m.options)
		case "left":
			switch m.options[m.optionIdx] {
			case "timer":
				m.timeIdx = (m.timeIdx - 1 + len(m.times)) % len(m.times)
			case "words":
				m.wordIdx = (m.wordIdx - 1 + len(m.words)) % len(m.words)
			}
		case "right":
			switch m.options[m.optionIdx] {
			case "timer":
				m.timeIdx = (m.timeIdx + 1) % len(m.times)
			case "words":
				m.wordIdx = (m.wordIdx + 1) % len(m.words)
			}
		case "enter":
			switch m.options[m.optionIdx] {
			case "timer":
				m.ActiveTime = m.times[m.timeIdx]
				m.ActiveTyperMode = "timer"
				m.writeSettings()
				return m, func() tea.Msg { return m }
			case "words":
				m.ActiveWords = m.words[m.wordIdx]
				m.ActiveTyperMode = "words"
				m.writeSettings()
				return m, func() tea.Msg { return m }
			case "quit":
				return m, tea.Quit
			case "restart":
				return m, func() tea.Msg { return m }
			}
		}
	}
	return m, nil
}

func (m SettingsModel) View() string {
	output := colors.Yellow.Render("Settings") + "\n"
	if m.options[m.optionIdx] == "timer" {
		output += colors.White.Render("timer: ")
		for timeIdx, times := range m.times {
			if timeIdx == m.timeIdx {
				if m.times[m.timeIdx] == m.ActiveTime && m.ActiveTyperMode == "timer" {
					output += colors.White.Underline(true).Render(strconv.Itoa(times)) + " "
				} else {
					output += colors.Gray.Underline(true).Render(strconv.Itoa(times)) + " "
				}
			} else {
				if m.times[timeIdx] == m.ActiveTime && m.ActiveTyperMode == "timer" {
					output += colors.White.Render(strconv.Itoa(times)) + " "
				} else {
					output += colors.Gray.Render(strconv.Itoa(times)) + " "
				}
			}
		}
	} else {
		output += colors.Gray.Render("timer: ")
		for timeIdx, times := range m.times {
			if m.times[timeIdx] == m.ActiveTime && m.ActiveTyperMode == "timer" {
				output += colors.White.Render(strconv.Itoa(times)) + " "
			} else {
				output += colors.Gray.Render(strconv.Itoa(times)) + " "
			}
		}
	}
	output += "\n"
	if m.options[m.optionIdx] == "words" {
		output += colors.White.Render("words: ")
		for wordIdx, word := range m.words {
			if wordIdx == m.wordIdx {
				if m.words[m.wordIdx] == m.ActiveWords && m.ActiveTyperMode == "words" {
					output += colors.White.Underline(true).Render(strconv.Itoa(word)) + " "
				} else {
					output += colors.Gray.Underline(true).Render(strconv.Itoa(word)) + " "
				}
			} else {
				if m.words[wordIdx] == m.ActiveWords && m.ActiveTyperMode == "words" {
					output += colors.White.Render(strconv.Itoa(word)) + " "
				} else {
					output += colors.Gray.Render(strconv.Itoa(word)) + " "
				}
			}
		}
	} else {
		output += colors.Gray.Render("words: ")
		for wordIdx, word := range m.words {
			if m.words[wordIdx] == m.ActiveWords && m.ActiveTyperMode == "words" {
				output += colors.White.Render(strconv.Itoa(word)) + " "
			} else {
				output += colors.Gray.Render(strconv.Itoa(word)) + " "
			}
		}
	}
	output += "\n"
	if m.options[m.optionIdx] == "restart" {
		output += colors.White.Render("restart")
	} else {
		output += colors.Gray.Render("restart")
	}
	output += "\n"
	if m.options[m.optionIdx] == "quit" {
		output += colors.White.Render("quit")
	} else {
		output += colors.Gray.Render("quit")
	}
	output += "\n"
	output += "\n" + colors.Instructions.Render(settingInstructions)
	return output
}

func (m SettingsModel) writeSettings() {
	basePath := resourcepath.GetResourcePath()
	file, err := os.Create(basePath + "/settings.json")
	if err != nil {
		return
	}
	defer file.Close()

	data := map[string]any {
		"activeTime":      m.ActiveTime,
		"activeWords":     m.ActiveWords,
		"activeTyperMode": m.ActiveTyperMode,
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	_ = encoder.Encode(data)
}

func readSettings(m *SettingsModel) {
	basePath := resourcepath.GetResourcePath()
	file, err := os.Open(basePath + "/settings.json")
	if err != nil {
		return
	}
	defer file.Close()

	var data struct {
		ActiveTime      int    `json:"activeTime"`
		ActiveWords     int    `json:"activeWords"`
		ActiveTyperMode string `json:"activeTyperMode"`
	}

	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return
	}

	// read and validate ActiveTime and timeIdx 
	if slices.Contains(m.times, data.ActiveTime) {
		m.ActiveTime = data.ActiveTime
		m.timeIdx = slices.Index(m.times, data.ActiveTime)
	}

	//  read and validate ActiveWords wordIdx
	if slices.Contains(m.words, data.ActiveWords) {
		m.ActiveWords = data.ActiveWords
		m.wordIdx = slices.Index(m.words, data.ActiveWords)
	}

	// read and validate ActiveTyperMode
	if data.ActiveTyperMode == "timer" || data.ActiveTyperMode == "words" {
		m.ActiveTyperMode = data.ActiveTyperMode
	}
}

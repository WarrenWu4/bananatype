package unit

import (
	"bananas/pkg/typer"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTyperInputLogic(t *testing.T) {
	ty := typer.NewTyper()

	// Get first character of first word in first line
	lines := ty.GetLines()
	if len(lines) == 0 || len(lines[0]) == 0 || len(lines[0][0]) == 0 {
		t.Fatal("Typer lines not initialized correctly")
	}
	firstWord := lines[0][0]
	firstChar := string(firstWord[0])

	// Test correct key press
	msg := tea.KeyMsg{Runes: []rune(firstChar), Type: tea.KeyRunes}
	updatedModel, _ := ty.Update(msg)
	ty = updatedModel.(typer.TyperModel)

	if ty.TotalTyped != 1 {
		t.Errorf("Expected TotalTyped to be 1, got %d", ty.TotalTyped)
	}
	if ty.TotalCorrect != 1 {
		t.Errorf("Expected TotalCorrect to be 1, got %d", ty.TotalCorrect)
	}

	// Test incorrect key press
	msgWrong := tea.KeyMsg{Runes: []rune("~"), Type: tea.KeyRunes} // unlikely to be the next char
	updatedModel, _ = ty.Update(msgWrong)
	ty = updatedModel.(typer.TyperModel)

	if ty.TotalTyped != 2 {
		t.Errorf("Expected TotalTyped to be 2, got %d", ty.TotalTyped)
	}
	if ty.TotalCorrect != 1 {
		t.Errorf("Expected TotalCorrect to be 1, still")
	}
}

package unit

import (
	"bananas/pkg/analysis"
	"strings"
	"testing"
)

func TestAnalysisCalculations(t *testing.T) {
	am := analysis.AnalysisModel{
		Time:       60,
		Words:      60,
		Correct:    300,
		Characters: 300,
	}
	view := am.View()
	if !strings.Contains(view, "wpm: 60") {
		t.Errorf("Expected view to contain 'wpm: 60', got: %s", view)
	}
	if !strings.Contains(view, "acc: 100.00") {
		t.Errorf("Expected view to contain 'acc: 100.00', got: %s", view)
	}
	amZero := analysis.AnalysisModel{
		Time:  0,
		Words: 10,
	}
	viewZero := amZero.View()
	if !strings.Contains(viewZero, "wpm: 0") {
		t.Errorf("Expected wpm: 0 for zero time, got: %s", viewZero)
	}
	amPartial := analysis.AnalysisModel{
		Time:       60,
		Words:      50,
		Correct:    250,
		Characters: 300,
	}
	viewPartial := amPartial.View()
	if !strings.Contains(viewPartial, "acc: 83.33") {
		t.Errorf("Expected accuracy to be 83.33, got: %s", viewPartial)
	}
}

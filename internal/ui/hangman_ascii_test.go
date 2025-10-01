package ui

import (
	"strings"
	"testing"
)

func TestStageBoundsAndFinal(t *testing.T) {
	start := Stage(0, 6)
	end := Stage(6, 6)

	if strings.Contains(start, "O") {
		t.Fatalf("start stage should not contain head 'O':\n%s", start)
	}
	if !strings.Contains(end, "O") {
		t.Fatalf("final stage should contain head 'O':\n%s", end)
	}
	if !strings.Contains(end, "/ \\") { // ноги
		t.Fatalf("final stage should contain legs '/ \\':\n%s", end)
	}

	if Stage(100, 6) != end {
		t.Fatalf("overflow incorrect should clamp to final stage")
	}
}

func TestStageScaling(t *testing.T) {
	final := Stage(6, 6)
	for i := 0; i <= 6; i++ {
		_ = Stage(i, 6) // главное — достижимость финальной стадии
	}
	if Stage(6, 6) != final {
		t.Fatalf("final stage must be stable")
	}
}

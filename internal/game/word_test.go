package game

import "testing"

func TestRandomWordIsLowerAndNonEmpty(t *testing.T) {
	s := RandomWord()
	if s == "" {
		t.Fatal("RandomWord returned empty string")
	}
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			t.Fatalf("RandomWord should be lowercased, got: %q", s)
		}
	}
}

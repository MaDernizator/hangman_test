package config

import "testing"

func TestParseDifficulty(t *testing.T) {
	cases := map[string]Difficulty{
		"easy":    Easy,
		"normal":  Normal,
		"medium":  Normal,
		"std":     Normal,
		"default": Normal,
		"hard":    Hard,
		"auto":    Auto,
		"random":  Auto,
		"":        Auto,
		"weird":   Auto,
	}
	for in, want := range cases {
		if got := ParseDifficulty(in); got != want {
			t.Fatalf("ParseDifficulty(%q)=%q, want %q", in, got, want)
		}
	}
}

func TestAttemptsFor(t *testing.T) {
	if got := AttemptsFor(Easy); got != 8 {
		t.Fatalf("Easy attempts = %d, want %d", got, 8)
	}
	if got := AttemptsFor(Normal); got != 6 {
		t.Fatalf("Normal attempts = %d, want %d", got, 6)
	}
	if got := AttemptsFor(Hard); got != 4 {
		t.Fatalf("Hard attempts = %d, want %d", got, 4)
	}
}

func TestRandomDifficultyOnlyFromSet(t *testing.T) {
	seen := map[Difficulty]bool{}
	for i := 0; i < 200; i++ {
		seen[RandomDifficulty()] = true
	}
	for d := range seen {
		if d != Easy && d != Normal && d != Hard {
			t.Fatalf("RandomDifficulty returned unexpected value: %q", d)
		}
	}
	if len(seen) == 0 {
		t.Fatal("RandomDifficulty returned nothing")
	}
}

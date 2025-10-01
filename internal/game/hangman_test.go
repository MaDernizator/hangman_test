package game

import (
	"testing"
	"unicode"
)

func TestCaseInsensitiveGuess(t *testing.T) {
	g := NewGame("GoLang", 6)
	g.Guess('g') // низкий регистр
	if g.Masked() == "******" {
		t.Fatalf("expected 'g' to reveal positions, got mask=%q", g.Masked())
	}
	g.Guess('O') // верхний регистр
	if g.Masked()[1] != 'o' {
		t.Fatalf("expected 'o' revealed at pos 1, mask=%q", g.Masked())
	}
}

func TestRepeatGuessNotPenalized(t *testing.T) {
	g := NewGame("aaa", 6)
	g.Guess('x') // промах
	if g.IncorrectGuesses != 1 {
		t.Fatalf("expected 1 incorrect, got %d", g.IncorrectGuesses)
	}
	g.Guess('x') // повтор — не штрафуем
	if g.IncorrectGuesses != 1 {
		t.Fatalf("repeat should not increase incorrect, got %d", g.IncorrectGuesses)
	}
}

func TestNonLetterIgnored(t *testing.T) {
	g := NewGame("test", 6)
	before := g.IncorrectGuesses
	g.Guess('1')
	if g.IncorrectGuesses != before {
		t.Fatalf("non-letter should not change incorrect guesses, got %d", g.IncorrectGuesses)
	}
	for _, r := range g.Masked() {
		if r != '*' {
			t.Fatalf("non-letter must not reveal letters, mask=%q", g.Masked())
		}
	}
}

func TestLossAfterMaxTries(t *testing.T) {
	g := NewGame("go", 2)
	g.Guess('x')
	g.Guess('y')
	if !g.IsLost() || !g.IsGameOver() {
		t.Fatalf("expected loss after reaching max tries")
	}
}

func TestWinRevealsAllLetters(t *testing.T) {
	g := NewGame("aba", 6)
	g.Guess('a')
	g.Guess('b')
	if !g.IsWon() || g.Masked() != "aba" {
		t.Fatalf("expected win with mask 'aba', got mask=%q", g.Masked())
	}
}

func TestMaskedNonLettersRemainVisible(t *testing.T) {
	g := NewGame("foo-bar", 6)
	// небуквенные символы сразу видны
	for i, r := range g.Word {
		if !unicode.IsLetter(r) && rune(g.MaskedWord[i]) != r {
			t.Fatalf("non-letter must be visible, got mask=%q", g.Masked())
		}
	}
}

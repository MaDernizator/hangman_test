package tests

import (
	"hangman/internal/game"
	"testing"
)

func TestGameGuess(t *testing.T) {
	game := game.NewGame("golang", 6)
	game.Guess('g')

	if game.MaskedWord != "g****g" {
		t.Errorf("Expected 'g****g', got '%s'", game.MaskedWord)
	}
}

func TestGameLoss(t *testing.T) {
	game := game.NewGame("golang", 3)
	game.Guess('x')
	game.Guess('y')
	game.Guess('z')

	if !game.IsGameOver() {
		t.Error("Expected game over after 3 incorrect guesses")
	}
}

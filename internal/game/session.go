package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"hangman/dictionary"
	"hangman/internal/ui"
)

// StartInteractive запускает интерактивную игру
func StartInteractive(maxTries int) {
	word := RandomWord()
	category := "general" // Категория по умолчанию, можно добавить в будущем выбор категории

	g := NewGame(word, maxTries)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Difficulty is set. Max mistakes: %d\n\n", maxTries)

	for !g.IsGameOver() {
		fmt.Println(ui.Stage(g.IncorrectGuesses, g.MaxTries))
		fmt.Print(ui.HUD(g.Masked(), g.MistakesLeft()))
		fmt.Print("Enter a letter or type 'hint' for a clue: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		runes := []rune(strings.ToLower(input))

		if input == "hint" {
			hint := dictionary.GetHint(category, word)
			fmt.Println("Hint:", hint)
			continue
		}

		if len(runes) != 1 || !unicode.IsLetter(runes[0]) {
			fmt.Println("Please enter a single letter (A-Я).")
			continue
		}

		letter := runes[0]
		if g.AlreadyGuessed(letter) {
			fmt.Println("You've already tried that letter.")
			continue
		}

		g.Guess(letter)
	}

	fmt.Println(ui.Stage(g.IncorrectGuesses, g.MaxTries))
	fmt.Printf("Word: %s\n", g.Masked())
	if g.IsWon() {
		fmt.Println("Congratulations! You guessed the word:", g.WordString())
	} else {
		fmt.Println("You lost! The word was:", g.WordString())
	}
}

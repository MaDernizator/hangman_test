package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"hangman/internal/ui"
)

// StartInteractive запускает пользовательский режим с заданным количеством попыток.
func StartInteractive(maxTries int) {
	word := RandomWord()
	g := NewGame(word, maxTries)
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Difficulty is set. Max mistakes: %d\n\n", maxTries)

	for !g.IsGameOver() {
		fmt.Println(ui.Stage(g.IncorrectGuesses, g.MaxTries))
		fmt.Print(ui.HUD(g.Masked(), g.MistakesLeft()))
		fmt.Print("Enter a letter: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		runes := []rune(strings.ToLower(input))

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

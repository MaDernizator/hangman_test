package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"hangman/internal/ui"
)

func StartGame(interactive bool) {
	var word string
	if interactive {
		word = RandomWord()
	} else {
		// В следующей итерации доведём неинтерактивный режим до точного формата из README.
		if len(os.Args) < 2 {
			fmt.Println("Usage (test mode temp): cmd <word>")
			return
		}
		word = os.Args[1]
	}

	g := NewGame(word, 6)
	reader := bufio.NewReader(os.Stdin)

	for !g.IsGameOver() {
		// Рисуем текущую виселицу + HUD
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

	// Финальный кадр и итог
	fmt.Println(ui.Stage(g.IncorrectGuesses, g.MaxTries))
	fmt.Printf("Word: %s\n", g.Masked())
	if g.IsWon() {
		fmt.Println("Congratulations! You guessed the word:", g.WordString())
	} else {
		fmt.Println("You lost! The word was:", g.WordString())
	}
}

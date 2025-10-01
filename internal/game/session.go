package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func StartGame(interactive bool) {
	var word string
	if interactive {
		word = RandomWord()
	} else {
		// ⚠️ неинтерактивный режим мы допилим отдельно под формат из README.
		// Сейчас оставим заглушку, чтобы интерактив работал стабильно.
		if len(os.Args) < 2 {
			fmt.Println("Usage (test mode temp): cmd <word>")
			return
		}
		word = os.Args[1]
	}

	g := NewGame(word, 6)
	reader := bufio.NewReader(os.Stdin)

	for !g.IsGameOver() {
		fmt.Printf("Word: %s\n", g.Masked())
		fmt.Printf("Incorrect guesses left: %d\n", g.MaxTries-g.IncorrectGuesses)
		fmt.Print("Enter a letter: ")

		input, _ := reader.ReadString('\n')     // читает до \n
		input = strings.TrimSpace(input)        // убираем \r\n / пробелы
		runes := []rune(strings.ToLower(input)) // работаем по рунам (UTF-8)

		if len(runes) != 1 || !unicode.IsLetter(runes[0]) {
			fmt.Println("Please enter a single letter (A-Я).")
			continue
		}

		letter := runes[0]
		if g.AlreadyGuessed(letter) {
			fmt.Println("You've already tried that letter.")
			continue
		}

		ok := g.Guess(letter)
		if !ok && !g.AlreadyGuessed(letter) {
			// промах уже учтён в Guess; отдельное сообщение — по желанию
		}
	}

	// финальный экран
	fmt.Printf("Word: %s\n", g.Masked())
	if g.IsWon() {
		fmt.Println("Congratulations! You guessed the word:", g.WordString())
	} else {
		fmt.Println("You lost! The word was:", g.WordString())
	}
}

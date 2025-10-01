package cli

import (
	"fmt"
	"os"
	"unicode"

	"hangman/internal/game"
)

// RunNonInteractive обрабатывает:
//
//	argv[1] — секретное слово,
//	argv[2] — строка "угаданного" (последовательность вводимых символов).
//
// Печатает "<маска>;<win|lose|progress>" в STDOUT.
func RunNonInteractive(secret, guessed string, maxTries int) int {
	g := game.NewGame(secret, maxTries)

	// Проигрываем последовательность "угадываний"
	for _, r := range []rune(guessed) {
		if g.IsGameOver() {
			break
		}
		if !unicode.IsLetter(r) {
			continue
		}
		g.Guess(r)
	}

	// Определяем результат
	var result string
	switch {
	case g.IsWon():
		result = "win"
	case g.IsLost():
		result = "lose"
	default:
		result = "progress"
	}

	// Вывод строго по формату
	fmt.Printf("%s;%s\n", g.Masked(), result)
	return 0
}

// ValidateArgs проверяет корректность аргументов.
func ValidateArgs(args []string) (string, string, error) {
	if len(args) != 3 {
		return "", "", fmt.Errorf("expected 2 arguments: <word> <guessed>")
	}
	secret := args[1]
	guessed := args[2]
	if len([]rune(secret)) < 1 {
		return "", "", fmt.Errorf("invalid word length")
	}
	return secret, guessed, nil
}

// Entry точка входа для неинтерактивного режима из main.
// Возвращает код выхода.
func Entry(args []string, maxTries int) int {
	secret, guessed, err := ValidateArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return 2
	}
	return RunNonInteractive(secret, guessed, maxTries)
}

package main

import (
	"os"

	"hangman/internal/cli"
	"hangman/internal/game"
)

func main() {
	// Если передано ровно 2 аргумента после имени программы — включаем неинтерактивный режим.
	// Пример: `hangman.exe бабушка баб`
	if len(os.Args) == 3 {
		code := cli.Entry(os.Args, 6) // 6 попыток по умолчанию; позже свяжем со сложностью
		os.Exit(code)
		return
	}

	// Иначе — интерактивный режим.
	game.StartGame(true)
}

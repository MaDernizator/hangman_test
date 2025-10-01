package main

import (
	"hangman/internal/game"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		game.StartGame(false) // Test mode
	} else {
		game.StartGame(true) // Interactive mode
	}
}

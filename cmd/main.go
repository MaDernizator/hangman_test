package main

import (
	"flag"
	"fmt"
	"os"

	"hangman/internal/cli"
	"hangman/internal/config"
	"hangman/internal/game"
)

func main() {
	diffFlag := flag.String("d", "auto", "difficulty: easy|normal|hard|auto")
	flag.StringVar(diffFlag, "difficulty", "auto", "difficulty: easy|normal|hard|auto")
	flag.Parse()

	// Определяем сложность (random, если auto/пусто/неизвестно)
	parsed := config.ParseDifficulty(*diffFlag)
	if parsed == config.Auto {
		parsed = config.RandomDifficulty()
	}
	maxTries := config.AttemptsFor(parsed)

	args := flag.Args()

	// Неинтерактивный режим по ТЗ: два позиционных аргумента (word, guessed)
	if len(args) == 2 {
		// Вывод требований — только "<mask;result>" в STDOUT, без лишних сообщений.
		// Поэтому не печатаем сложность, просто применяем её для MaxTries.
		code := cli.RunNonInteractive(args[0], args[1], maxTries)
		os.Exit(code)
		return
	}

	// Иначе — интерактивный режим.
	fmt.Printf("Chosen difficulty: %s\n", parsed)
	game.StartInteractive(maxTries)
}

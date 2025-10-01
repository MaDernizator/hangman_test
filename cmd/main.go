package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"hangman/dictionary"
	"hangman/internal/cli"
	"hangman/internal/config"
	"hangman/internal/game"
)

func main() {
	// Обработка флага сложности
	diffFlag := flag.String("d", "auto", "difficulty: easy|normal|hard|auto")
	flag.StringVar(diffFlag, "difficulty", "auto", "difficulty: easy|normal|hard|auto")
	flag.Parse()

	// Загружаем данные из JSON
	err := dictionary.LoadData("data.json")
	if err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

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
		code := cli.RunNonInteractive(args[0], args[1], maxTries)
		os.Exit(code)
		return
	}

	// Иначе — интерактивный режим.
	fmt.Printf("Chosen difficulty: %s\n", parsed)
	game.StartInteractive(maxTries)
}

package dictionary

import (
	"errors"
	"math/rand"
	"sort"
	"strings"
	"time"

	"hangman/internal/config"
)

var data = map[string][]string{
	"general": {
		"golang", "hangman", "programming", "developer", "function",
		"variable", "interface", "pointer", "package", "module",
	},
	"animals": {
		"cat", "dog", "elephant", "giraffe", "hedgehog", "sparrow",
	},
	"cs": {
		"binary", "cache", "mutex", "pointer", "compiler", "runtime",
	},
}

// Подсказки для каждого слова
var hints = map[string]map[string]string{
	"general": {
		"golang":      "A popular programming language.",
		"hangman":     "A word guessing game.",
		"programming": "The process of writing computer programs.",
		"developer":   "A person who creates software.",
		"function":    "A block of code that performs a specific task.",
		"variable":    "A symbolic name associated with a value.",
		"interface":   "A type in Go that specifies a set of methods.",
		"pointer":     "A variable that stores the memory address of another variable.",
	},
	"animals": {
		"cat":      "A small domesticated carnivorous mammal.",
		"dog":      "A domesticated carnivorous mammal known as man's best friend.",
		"elephant": "A large mammal with a trunk, found in Africa and Asia.",
		"giraffe":  "A tall mammal with a very long neck, native to Africa.",
		"hedgehog": "A small mammal known for its spiny coat.",
		"sparrow":  "A small, plump bird with short wings.",
	},
	"cs": {
		"binary":   "A system of numbers using only 0s and 1s.",
		"cache":    "A small, fast storage area used to store frequently accessed data.",
		"mutex":    "A mutual exclusion object used to prevent race conditions in concurrent programming.",
		"pointer":  "A variable that stores the address of another variable in memory.",
		"compiler": "A program that converts source code into machine code.",
		"runtime":  "The time during which a program is running.",
	},
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func Categories() []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func Words(category string) ([]string, error) {
	words, ok := data[strings.ToLower(strings.TrimSpace(category))]
	if !ok {
		return nil, errors.New("unknown category")
	}
	out := make([]string, len(words))
	copy(out, words)
	return out, nil
}

func RandomWord(category string) (string, error) {
	words, err := Words(category)
	if err != nil {
		return "", err
	}
	if len(words) == 0 {
		return "", errors.New("category is empty")
	}
	return words[rng.Intn(len(words))], nil
}

// byDifficulty фильтрует слова по длине в зависимости от сложности.
func byDifficulty(words []string, d config.Difficulty) []string {
	var min, max int
	switch d {
	case config.Easy:
		min, max = 3, 5
	case config.Normal:
		min, max = 6, 8
	case config.Hard:
		min, max = 9, 1000
	}
	var result []string
	for _, word := range words {
		wordLen := len([]rune(word))
		if wordLen >= min && wordLen <= max {
			result = append(result, word)
		}
	}
	return result
}

func RandomWordWithDifficulty(category string, d config.Difficulty, r *rand.Rand) (string, error) {
	if r == nil {
		r = rng
	}
	words, err := Words(category)
	if err != nil {
		return "", err
	}

	// Фильтруем по сложности
	candidates := byDifficulty(words, d)

	// Если фильтрация по сложности не дала слов, берем все доступные
	pool := candidates
	if len(pool) == 0 {
		pool = words
	}
	return pool[r.Intn(len(pool))], nil
}

// getHint возвращает подсказку для слова из указанной категории.
func GetHint(category, word string) string {
	if categoryHints, ok := hints[category]; ok {
		if hint, exists := categoryHints[word]; exists {
			return hint
		}
	}
	return "No hint available for this word."
}

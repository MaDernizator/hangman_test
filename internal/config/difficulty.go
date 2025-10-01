package config

import (
	"math/rand"
	"strings"
	"time"
)

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Normal Difficulty = "normal"
	Hard   Difficulty = "hard"
	Auto   Difficulty = "auto"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func ParseDifficulty(s string) Difficulty {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "easy":
		return Easy
	case "normal", "medium", "std", "default":
		return Normal
	case "hard":
		return Hard
	case "", "auto", "random":
		return Auto
	default:
		// неизвестное значение трактуем как auto (рандом)
		return Auto
	}
}

func RandomDifficulty() Difficulty {
	all := []Difficulty{Easy, Normal, Hard}
	return all[rng.Intn(len(all))]
}

// AttemptsFor возвращает число попыток по сложности.
// (просто и прозрачно; позже можно усложнить формулу с учетом длины слова)
func AttemptsFor(d Difficulty /*, word string*/) int {
	switch d {
	case Easy:
		return 8
	case Hard:
		return 4
	case Normal:
		fallthrough
	default:
		return 6
	}
}

package game

import (
	"math/rand"
	"strings"
	"time"
)

// общий источник случайных чисел для пакета game
var rng *rand.Rand

func init() {
	// Сидим ОДИН раз при загрузке пакета.
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

var words = []string{
	"golang",
	"hangman",
	"programming",
	"developer",
	"function",
	// можно добавить ещё: "variable", "interface", "pointer", ...
}

// RandomWord возвращает случайное слово в нижнем регистре.
func RandomWord() string {
	return strings.ToLower(words[rng.Intn(len(words))])
}

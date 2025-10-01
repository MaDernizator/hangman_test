package dictionary

import (
	"math/rand"
	"time"
)

var words = []string{
	"computer", "programming", "hangman", "development", "golang",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomWord() string {
	return words[rand.Intn(len(words))]
}

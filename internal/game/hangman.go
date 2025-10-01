package game

import (
	"strings"
	"unicode"
)

type Game struct {
	Word             []rune        // загаданное слово (в нижнем регистре)
	MaskedWord       []rune        // текущая маска
	Guessed          map[rune]bool // уже вводившиеся буквы
	IncorrectGuesses int
	MaxTries         int
}

func NewGame(word string, maxTries int) *Game {
	normalized := []rune(strings.ToLower(strings.TrimSpace(word)))
	masked := make([]rune, len(normalized))
	for i := range masked {
		// маскируем только буквы; если символ не буква (пробел/дефис), сразу открываем
		if unicode.IsLetter(normalized[i]) {
			masked[i] = '*'
		} else {
			masked[i] = normalized[i]
		}
	}
	return &Game{
		Word:             normalized,
		MaskedWord:       masked,
		Guessed:          make(map[rune]bool),
		IncorrectGuesses: 0,
		MaxTries:         maxTries,
	}
}

// Guess возвращает true, если буква есть в слове, false — если промах.
// Повторный ввод не штрафует и возвращает false.
func (g *Game) Guess(ch rune) bool {
	ch = unicode.ToLower(ch)
	if !unicode.IsLetter(ch) {
		return false
	}
	if g.Guessed[ch] {
		// повтор — не штрафуем
		return false
	}
	g.Guessed[ch] = true

	found := false
	for i, r := range g.Word {
		if r == ch {
			g.MaskedWord[i] = r
			found = true
		}
	}
	if !found {
		g.IncorrectGuesses++
	}
	return found
}

func (g *Game) AlreadyGuessed(ch rune) bool {
	ch = unicode.ToLower(ch)
	return g.Guessed[ch]
}

func (g *Game) IsWon() bool {
	for i := range g.Word {
		if unicode.IsLetter(g.Word[i]) && g.MaskedWord[i] == '*' {
			return false
		}
	}
	return true
}

func (g *Game) IsLost() bool {
	return g.IncorrectGuesses >= g.MaxTries
}

func (g *Game) IsGameOver() bool {
	return g.IsWon() || g.IsLost()
}

func (g *Game) Masked() string {
	return string(g.MaskedWord)
}

func (g *Game) WordString() string {
	return string(g.Word)
}

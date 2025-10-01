package ui

import (
	"fmt"
	"math"
)

// stages — от пустой виселицы до полностью «собранного» висельника.
// 0 — старт, последний индекс — максимум ошибок.
var stages = []string{
	// 0
	`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	// 1 (голова)
	`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
	// 2 (голова + туловище)
	`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
	// 3 (голова + туловище + левая рука)
	`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
	// 4 (обе руки)
	`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
	// 5 (одна нога)
	`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
	// 6 (две ноги — финал)
	`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

// Stage возвращает ASCII-кадр для текущего числа ошибок.
// Мы масштабируем количество кадров к MaxTries, чтобы визуализация была корректной
// при любом числе допустимых ошибок.
func Stage(incorrect, maxTries int) string {
	if maxTries <= 0 {
		return stages[0]
	}
	// Приводим incorrect к диапазону [0, maxTries]
	if incorrect < 0 {
		incorrect = 0
	}
	if incorrect > maxTries {
		incorrect = maxTries
	}

	// Индекс стадии в диапазоне [0, len(stages)-1]
	// Пропорционально ошибкам. Используем округление вверх, чтобы прогресс был заметнее.
	f := float64(incorrect) / float64(maxTries)        // [0..1]
	idx := int(math.Round(f * float64(len(stages)-1))) // [0..last]
	if idx < 0 {
		idx = 0
	}
	if idx >= len(stages) {
		idx = len(stages) - 1
	}
	return stages[idx]
}

// HUD — небольшой хедер состояния: слово, оставшиеся попытки.
func HUD(masked string, left int) string {
	return fmt.Sprintf("Word: %s\nIncorrect guesses left: %d\n", masked, left)
}

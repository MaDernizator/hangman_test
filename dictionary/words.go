package dictionary

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"hangman/internal/config"
)

// Структура для загрузки данных из JSON
type CategoryData struct {
	Words []string          `json:"words"`
	Hints map[string]string `json:"hints"`
}

// Global data structure holding categories, words and hints
var data = make(map[string]CategoryData)
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// Загружаем данные из JSON-файла
func LoadData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening JSON file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	return nil
}

// Получение всех категорий
func Categories() []string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	return keys
}

// Получение всех слов в категории
func Words(category string) ([]string, error) {
	categoryData, ok := data[strings.ToLower(category)]
	if !ok {
		return nil, errors.New("unknown category")
	}
	return categoryData.Words, nil
}

// Получение случайного слова из категории
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

// Функция для получения подсказки для слова
func GetHint(category, word string) string {
	categoryData, ok := data[category]
	if !ok {
		return "No hint available for this word."
	}
	hint, exists := categoryData.Hints[word]
	if !exists {
		return "No hint available for this word."
	}
	return hint
}

// byDifficulty фильтрует слова по длине в зависимости от сложности
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

// RandomWordWithDifficulty выбирает случайное слово из категории с учётом сложности
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

	// Если фильтрация по сложности не дала слов, берём все доступные
	pool := candidates
	if len(pool) == 0 {
		pool = words
	}
	return pool[r.Intn(len(pool))], nil
}

package dictionary

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadData(t *testing.T) {
	// Убедимся, что текущий рабочий каталог правильный
	fmt.Println("Current working directory:", getCurrentDir())

	// Путь к data.json
	filePath := "../data.json" // Это относительный путь от корня проекта

	// Проверяем существование файла
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Fatalf("data.json file does not exist at path %s", filePath)
	}

	// Пробуем загрузить данные из JSON
	err := LoadData(filePath)
	if err != nil {
		t.Fatalf("Error loading data from JSON: %v", err)
	}

	// Проверяем, что данные загружены корректно
	if len(data) == 0 {
		t.Fatal("No data loaded from JSON")
	}

	// Проверяем наличие категорий
	categories := Categories()
	if len(categories) == 0 {
		t.Fatal("No categories found")
	}
}

// getCurrentDir выводит текущую рабочую директорию
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
	}
	return dir
}

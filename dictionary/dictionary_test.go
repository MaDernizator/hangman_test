package dictionary

import "testing"

// 1) Функция возвращает НЕ пустую строку и слово из известного списка.
func TestGetRandomWordIsFromList(t *testing.T) {
	got := GetRandomWord()
	if got == "" {
		t.Fatal("GetRandomWord returned empty string")
	}
	found := false
	for _, w := range words {
		if w == got {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("GetRandomWord returned %q which is not in words slice", got)
	}
}

// 2) Разнообразие результатов по множественным вызовам без завязки на seed.
// Не требуем покрытия всех слов, но ожидаем >= 2 разных результата.
func TestGetRandomWordVarietyAcrossCalls(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 50; i++ {
		seen[GetRandomWord()] = true
	}
	if len(seen) < 2 {
		t.Fatalf("expected at least 2 distinct words across multiple calls, got %v", seen)
	}
}

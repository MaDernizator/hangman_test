package dictionary

import "testing"

func TestGetHint(t *testing.T) {
	// Пример с категорией "general"
	hint := GetHint("general", "golang")
	if hint == "" {
		t.Fatal("expected hint for golang, got empty")
	}
	if hint != "A popular programming language." {
		t.Fatalf("unexpected hint for golang: %q", hint)
	}

	// Пример для несуществующего слова
	hint = GetHint("general", "unknownword")
	if hint != "No hint available for this word." {
		t.Fatalf("expected 'No hint available' but got: %q", hint)
	}
}

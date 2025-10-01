package cli

import "testing"

func TestValidateArgsOk(t *testing.T) {
	secret, guessed, err := ValidateArgs([]string{"app", "word", "abc"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if secret != "word" || guessed != "abc" {
		t.Fatalf("parsed wrong args: secret=%q guessed=%q", secret, guessed)
	}
}

func TestValidateArgsCount(t *testing.T) {
	if _, _, err := ValidateArgs([]string{"app"}); err == nil {
		t.Fatalf("expected error on too few args")
	}
}

func TestValidateArgsWordLength(t *testing.T) {
	if _, _, err := ValidateArgs([]string{"app", "", "xyz"}); err == nil {
		t.Fatalf("expected error on empty secret")
	}
}

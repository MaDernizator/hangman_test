package cli

import (
	"bytes"
	"os"
	"regexp"
	"testing"
)

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	_ = w.Close()
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	os.Stdout = old
	return buf.String()
}

func TestNonInteractiveFormatProgress(t *testing.T) {
	out := captureStdout(func() {
		RunNonInteractive("golang", "go", 6)
	})
	re := regexp.MustCompile(`^[\*\p{L}\- ]+;(win|lose|progress)\n$`)
	if !re.MatchString(out) {
		t.Fatalf("unexpected output format: %q", out)
	}
}

func TestNonInteractiveWin(t *testing.T) {
	out := captureStdout(func() {
		RunNonInteractive("go", "og", 6)
	})
	if out != "go;win\n" {
		t.Fatalf("expected 'go;win', got %q", out)
	}
}

func TestNonInteractiveLose(t *testing.T) {
	out := captureStdout(func() {
		RunNonInteractive("go", "abcd", 3) // 4 промаха > 3
	})
	if out != "**;lose\n" {
		t.Fatalf("expected '**;lose', got %q", out)
	}
}

func TestNonInteractiveIgnoresNonLettersAndRepeats(t *testing.T) {
	out := captureStdout(func() {
		RunNonInteractive("go", "g1g!o", 6)
	})
	if out != "go;win\n" {
		t.Fatalf("expected 'go;win', got %q", out)
	}
}

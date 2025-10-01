package main_test

import (
	"bytes"
	"os/exec"
	"regexp"
	"runtime"
	"testing"
)

// runMain запускает "go run ." из каталога cmd и возвращает stdout+stderr.
func runMain(t *testing.T, args ...string) (string, error) {
	t.Helper()
	cmd := exec.Command("go", append([]string{"run", "."}, args...)...)
	// Тест лежит в каталоге cmd, тестовый раннер запускает его с рабочей директорией = cmd,
	// так что дополнительных настроек Dir не нужно.
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func TestMainNonInteractiveHardLose(t *testing.T) {
	// -d hard => 4 попытки. 4 промаха по "go" должны дать lose.
	out, err := runMain(t, "-d", "hard", "go", "abcd")
	if err != nil {
		t.Fatalf("run error: %v, out=%q", err, out)
	}
	if out != "**;lose\n" {
		t.Fatalf("expected '**;lose', got %q", out)
	}
}

func TestMainNonInteractiveEasyProgress(t *testing.T) {
	// -d easy => 8 попыток. 4 промаха по "go" => ещё не поражение (progress).
	out, err := runMain(t, "-d", "easy", "go", "abcd")
	if err != nil {
		t.Fatalf("run error: %v, out=%q", err, out)
	}
	re := regexp.MustCompile(`^\*\*;progress\n$`)
	if !re.MatchString(out) {
		t.Fatalf("expected '**;progress', got %q", out)
	}
}

func TestMainNonInteractiveWin(t *testing.T) {
	out, err := runMain(t, "-d", "normal", "go", "og")
	if err != nil {
		t.Fatalf("run error: %v, out=%q", err, out)
	}
	if out != "go;win\n" {
		t.Fatalf("expected 'go;win', got %q", out)
	}
}

func TestMainInteractiveSmoke(t *testing.T) {
	// Интерактивная ветка ждёт stdin; делаем дымовой тест и не зависаем.
	if runtime.GOOS == "windows" {
		t.Skip("skip interactive smoke on Windows (TTY)")
	}
	cmd := exec.Command("go", "run", ".")
	_ = cmd.Start()
	_ = cmd.Process.Kill()
}

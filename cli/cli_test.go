package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	w := &bytes.Buffer{}

	Start(w, []string{"help"})

	got := strings.TrimRight(w.String(), "\n")
	if got != usage {
		t.Errorf("Start() = %v, want %v", got, usage)
	}
}

func TestWrongArguments(t *testing.T) {
	w := &bytes.Buffer{}

	Start(w, []string{"wrong", "arguments"})

	got := strings.TrimRight(w.String(), "\n")
	if got != wrongInput {
		t.Errorf("Start() = %v, want %v", got, wrongInput)
	}
}

func TestWrongNameFormat(t *testing.T) {
	w := &bytes.Buffer{}

	Start(w, []string{"init", "name.with.dots"})

	got := strings.TrimRight(w.String(), "\n")
	if got != invalidName {
		t.Errorf("Start() = %v, want %v", got, invalidName)
	}
}

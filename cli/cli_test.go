package cli

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestStart(t *testing.T) {
	w := &bytes.Buffer{}
	var pjName string
	var techStack []string
	generate := func(pj string, tech []string) {
		pjName = pj
		techStack = tech
	}

	Start(w, []string{"init", "--frontend", "vue", "my-app"}, generate)

	if pjName != "my-app" || techStack[0] != "vue" {
		t.Error("Generation function function received wrong arguments")
	}
	if w.String() != "" {
		t.Errorf("No output expected but got %v", w.String())
	}
}

func TestStartGivenHelp(t *testing.T) {
	w := &bytes.Buffer{}

	Start(w, []string{"help"}, nil)

	got := strings.TrimRight(w.String(), "\n")
	if got != usage {
		t.Errorf("Start() = %v, want %v", got, usage)
	}
}

func TestStartGivenWrongArguments(t *testing.T) {
	w := &bytes.Buffer{}

	Start(w, []string{"wrong", "arguments"}, nil)

	got := strings.TrimRight(w.String(), "\n")
	expected := "Wrong input!\n" + usage
	if got != expected {
		t.Errorf("Start() = %v, want %v", got, expected)
	}
}

func TestStartGivenWrongNameFormat(t *testing.T) {
	w := &bytes.Buffer{}

	Start(w, []string{"init", "name.with.dots"}, nil)

	got := strings.TrimRight(w.String(), "\n")
	if got != invalidName {
		t.Errorf("Start() = %v, want %v", got, invalidName)
	}
}

func TestParseFlags(t *testing.T) {
	flags := []flag{{
		"--frontend", []string{"react", "vue"}, "react",
	}}
	options := []string{"--frontend", "vue"}

	values, err := parseFlags(options, flags)

	if err != nil {
		t.Error("parseFlags must not return an error")
	}
	if len(values) != 1 && values[0] != "vue" {
		t.Error(`parseFlags must return ["vue"]`)
	}
}

func TestParseFlagsGivenSeveralFlags(t *testing.T) {
	flags := []flag{
		{"--frontend", []string{"react", "vue"}, "react"},
		{"--db", []string{"mysql", "mongo"}, "mongo"},
	}
	options := []string{"--frontend", "vue", "--db", "mysql"}

	got, err := parseFlags(options, flags)

	if err != nil {
		t.Error("parseFlags must not return an error")
	}
	exp := []string{"vue", "mysql"}
	if !reflect.DeepEqual(got, exp) {
		t.Error(`parseFlags must return "vue" and "mysql"`)
	}
}

func TestParseFlagsGivenDefaultValues(t *testing.T) {
	flags := []flag{{
		"--frontend", []string{"react", "vue"}, "react",
	}}

	got, err := parseFlags([]string{}, flags)

	if err != nil {
		t.Error("parseFlags must not return an error")
	}
	exp := []string{"react"}
	if !reflect.DeepEqual(got, exp) {
		t.Error(`parseFlags must return "react"`)
	}
}

func TestParseFlagsGivenUnknownArguments(t *testing.T) {
	flags := []flag{{
		"--frontend", []string{"react", "vue"}, "react",
	}}
	options := []string{"--unknown", "test"}

	got, err := parseFlags(options, flags)

	if !reflect.DeepEqual(err, errors.New("flag mismatch")) {
		t.Error("parseFlags must return an error")
	}
	if got != nil {
		t.Error("parseFlags must return nil values")
	}
}

func TestParseFlagsGivenWrongFlagValue(t *testing.T) {
	flags := []flag{{
		"--frontend", []string{"react", "vue"}, "react",
	}}
	options := []string{"--frontend", "unknown"}

	values, err := parseFlags(options, flags)

	expected := errors.New(`invalid value of "--frontend" flag`)
	if !reflect.DeepEqual(err, expected) {
		t.Error("parseFlags must return an error")
	}
	if values != nil {
		t.Error("parseFlags must return nil values")
	}
}

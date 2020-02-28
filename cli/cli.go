package cli

import (
	"errors"
	"fmt"
	"io"
	"regexp"
)

type flag struct {
	name         string
	options      []string
	defaultValue string
}

func (f flag) valid(input string) bool {
	for _, val := range f.options {
		if input == val {
			return true
		}
	}
	return false
}

func Start(w io.Writer, commands []string, generate func(string, string)) {
	if len(commands) == 1 && commands[0] == "help" {
		fmt.Fprintln(w, usage)
		return
	}
	if len(commands) >= 2 && commands[0] == "init" {
		projectName := commands[len(commands)-1]
		input := commands[1 : len(commands)-1]
		values, err := parseFlags(input, flags())
		if err != nil {
			fmt.Fprintln(w, wrongInput)
			return
		}
		validName, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", projectName)
		if !validName {
			fmt.Fprintln(w, invalidName)
			return
		}
		generate(projectName, values["--frontend"])
		return
	}
	fmt.Fprintln(w, wrongInput)
}

// Retrieves flag values from user input matching them
// against the flags. Returns a map of flags with
// their values or error in case of invalid input.
func parseFlags(input []string, flags []flag) (map[string]string, error) {
	flagsCount := 0
	values := make(map[string]string)
	for _, fl := range flags {
		var value string
		for index, val := range input {
			if val == fl.name && len(input) > index+1 {
				value = input[index+1]
				if !fl.valid(value) {
					return nil, errors.New("invalid value")
				}
			}
		}
		if value == "" {
			value = fl.defaultValue
		} else {
			flagsCount += 2
		}
		values[fl.name] = value
	}
	if flagsCount != len(input) {
		return nil, errors.New("flag mismatch")
	}
	return values, nil
}

func flags() []flag {
	return []flag{
		{"--frontend", []string{"angular", "react", "vue"}, "react"},
	}
}

const usage = `Usage:

  go run github.com/shpota/goxygen init [options] <project-directory>

Options:

  --frontend <framework-name>     Specify the front end framework. Possible ` +
	`options are "angular", "react" and "vue". If not specified "react" is used.`

const wrongInput = `Wrong input! Use:

  go run github.com/shpota/goxygen init [options] <project-directory>

In order to get more details run:

  go run github.com/shpota/goxygen help`

const invalidName = "Project name is not valid. The allowed symbols are " +
	"letters, numbers, underscores, and dashes."

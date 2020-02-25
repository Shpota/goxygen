package cli

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/shpota/goxygen/codegen"
)

func Start(w io.Writer, commands []string) {
	argsLen := len(commands)
	if argsLen == 1 && commands[0] == "help" {
		fmt.Fprintln(w, usage)
		return
	}
	if (argsLen == 2 || argsLen == 3) && commands[0] == "init" {
		var frontendTech string
		var projectName string

		projectName, ok, errMsg := extractProjectName(commands[argsLen-1])
		if !ok {
			fmt.Fprintln(w, errMsg)
			return
		}
		if argsLen == 2 {
			codegen.Generate(projectName, strings.ToLower(Angular))
			return
		}

		frontendTech, ok, errMsg = extractFrontendTechName(commands[1])
		if !ok {
			fmt.Fprintln(w, errMsg)
			return
		}
		codegen.Generate(projectName, strings.ToLower(frontendTech))
		return
	}
	fmt.Fprintln(w, wrongInputErr)
	return
}

const usage = `Usage:

	go run github.com/shpota/goxygen init [--frontend] <project-name>

Generates a web project with the following structure:
server side Go code, webapp with corresponding frontend
components, by default React, a Dockerfile that builds 
the application,and Docker Compose scripts to run the app
in dev andprod modes.`

const wrongInputErr = `Wrong input! Use:
   go run github.com/shpota/goxygen init [--frontend] <project-name>
In order to get more details run:
   go run github.com/shpota/goxygen help`

const missProjectNameErr = `Project Name missing!`

const Vue = "Vue"
const Angular = "Angular"
const React = "React"
const flags = "--"

func invalidProjectNameErr(name string) string {
	return fmt.Sprintf(`Project name %v is not valid. The allowed symbols are letters,
numbers, underscores, and dashes.`, name)
}

func invalidFrontendTechErr(name string) string {
	return fmt.Sprintf(`%v is not support,
now that we support %v, %v, %v.`, name, Vue, Angular, React)
}

func isSupportedFrontendTech(name string) bool {
	name = strings.ToLower(name)
	if name != strings.ToLower(Vue) && name != strings.ToLower(Angular) && name != strings.ToLower(React) {
		return false
	}
	return true
}

func extractProjectName(arg string) (string, bool, string) {
	if strings.HasPrefix(arg, flags) {
		return "", false, missProjectNameErr
	}
	isValid, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", arg)
	if isValid {
		return arg, true, ""
	}
	return "", false, invalidProjectNameErr(arg)
}

func extractFrontendTechName(arg string) (string, bool, string) {
	var frontendTech string
	if strings.HasPrefix(arg, flags) {
		frontendTech = strings.TrimLeft(arg, flags)
		if !isSupportedFrontendTech(frontendTech) {
			return "", false, invalidFrontendTechErr(frontendTech)
		}
	} else {
		return "", false, wrongInputErr
	}
	return frontendTech, true, ""
}

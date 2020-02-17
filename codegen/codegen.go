package codegen

import (
	"fmt"
	"github.com/shpota/goxygen/static"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Generate(projectName string) {
	fmt.Println("Generating", projectName)
	for path, srcText := range static.Sources() {
		srcText = strings.Replace(srcText, "project-name", projectName, -1)
		binary := []byte(srcText)
		generateFile(projectName, path, binary)
	}
	for path, binary := range static.Images() {
		generateFile(projectName, path, binary)
	}
	err := initGitRepo(projectName)
	if err != nil {
		fmt.Println("Failed to setup a Git repository:", err)
	}
	fmt.Println("Generation completed.")
}

func generateFile(pjRoot string, path string, content []byte) {
	pathElements := strings.Split(path, "/")
	separator := string(os.PathSeparator)
	pathElements = append([]string{pjRoot}, pathElements...)
	_ = os.MkdirAll(
		strings.Join(pathElements[:len(pathElements)-1], separator),
		os.ModePerm,
	)
	fmt.Println("creating: " + strings.Join(pathElements, separator))
	err := ioutil.WriteFile(
		strings.Join(pathElements, separator),
		content,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func initGitRepo(projectName string) error {
	fmt.Println("setting up Git repository")
	cmd := exec.Command("git", "init", ".")
	cmd.Dir = projectName
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = projectName
	err = cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "commit", "-m", "Initial commit from Goxygen")
	cmd.Dir = projectName
	return cmd.Run()
}

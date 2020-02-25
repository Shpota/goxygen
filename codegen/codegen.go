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

func Generate(projectName, framework string) {
	fmt.Println("Generating", projectName)
	processTextFiles(projectName, framework)
	processImages(projectName, framework)
	err := initGitRepo(projectName)
	if err != nil {
		fmt.Println("Failed to setup a Git repository:", err)
	}
	fmt.Println("Generation completed.")
}

func processImages(projectName string, framework string) {
	for path, binary := range static.Images() {
		if !include(path, framework) {
			continue
		}
		path = strings.Replace(path, framework+".", "", 1)
		generateFile(projectName, path, binary)
	}
}

func processTextFiles(projectName string, framework string) {
	for path, srcText := range static.Sources() {
		if !include(path, framework) {
			continue
		}
		path = strings.Replace(path, framework+".", "", 1)
		srcText = strings.ReplaceAll(srcText, "project-name", projectName)
		binary := []byte(srcText)
		generateFile(projectName, path, binary)
	}
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

// Checks if a path is a framework-specific path (starts
// with framework name). Returns true if a path is
// prefixed with the provided framework followed by dot
// or if a path has no prefix.
func include(path, framework string) bool {
	for _, fr := range frameworks {
		if strings.HasPrefix(path, fr+".") && framework != fr {
			return false
		}
	}
	return true
}

var frameworks = []string{"angular", "react", "vue"}

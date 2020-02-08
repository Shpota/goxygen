package codegen

import (
	"fmt"
	"github.com/shpota/goxygen/static"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Generate(projectName string) {
	for path, srcText := range static.Sources() {
		binary := []byte(srcText)
		generateFile(projectName, path, binary)
	}
	for path, binary := range static.Images() {
		generateFile(projectName, path, binary)
	}
	fmt.Println("Generation completed.")
}

func generateFile(pjRoot string, path string, content []byte) {
	separator := string(os.PathSeparator)
	pathElements := strings.Split(path, separator)
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

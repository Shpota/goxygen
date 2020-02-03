// Transforms static content form the 'templates' folder
// in the root of the repository to Go code in order to
// include it to the end binary.
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	templatesDir := "../templates"
	paths := filePaths(templatesDir)
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	textFiles, images := collectFiles(paths, templatesDir)
	for path, content := range textFiles {
		writeTextFile(content, path, &buffer)
	}
	buffer.WriteString(middle)
	for path, imgBinary := range images {
		writeImage(imgBinary, path, &buffer)
	}
	buffer.WriteString(suffix)
	createSourceFile("../static/generated.go", buffer.Bytes())
	log.Println("Completed.")
}

func createSourceFile(path string, content []byte) {
	pkgDir := filepath.Join("..", "static")
	_ = os.MkdirAll(pkgDir, os.ModePerm)
	err := ioutil.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func collectFiles(paths []string, root string) (map[string]string, map[string][]byte) {
	textFiles := make(map[string]string)
	images := make(map[string][]byte)
	for _, originalPath := range paths {
		path := originalPath[len(root)+1:]
		ext := filepath.Ext(originalPath)
		content, err := ioutil.ReadFile(originalPath)
		if err != nil {
			log.Fatalln(err)
		}
		if ext == ".png" || ext == ".ico" {
			images[path] = content
		} else {
			textFiles[path] = string(content)
		}
	}
	return textFiles, images
}

func writeImage(imgBinary []byte, path string, buffer *bytes.Buffer) {
	buffer.WriteString(`		"` + path + `": {`)
	for _, b := range imgBinary {
		buffer.WriteString(strconv.Itoa(int(b)) + ", ")
	}
	buffer.WriteString("}" + separator)
}

func writeTextFile(content string, path string, buffer *bytes.Buffer) {
	buffer.WriteString(`		"` + path + `": ` + "`")
	for _, c := range content {
		if string(c) == "`" {
			buffer.WriteString("` + " + `"` + "`" + `" + ` + "`")
		} else {
			buffer.WriteString(string(c))
		}
	}
	buffer.WriteString("`" + separator)
}

func filePaths(dir string) []string {
	templates := []string{}
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				templates = append(templates, path)
			}
			return nil
		})
	if err != nil {
		log.Fatalln(err)
	}
	return templates
}

const prefix = `// THIS CODE IS GENERATED, DO NOT EDIT

// Package 'static' contains static assets generated form
// the 'templates' folder in the root of the repository.
// If a change is made in templates regenerate this
// file by running 'transform/transformer.go'.
package static

func Sources() map[string]string {
	return map[string]string{
`
const separator = `,
`
const middle = `	}
}

func Images() map[string][]byte {
	return map[string][]byte{
`

const suffix = `	}
}`

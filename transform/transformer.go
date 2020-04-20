// Transforms static content form the 'templates' folder
// in the root of the repository to Go code in order to
// distribute it as a part of the module.
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	textFiles, images := contentFromFiles("../templates")
	for _, path := range sortedFilePaths(textFiles) {
		writeTextFile(path, textFiles[path], &buffer)
	}
	buffer.WriteString(middle)
	for _, path := range sortedImagePaths(images) {
		writeImage(path, images[path], &buffer)
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

func contentFromFiles(root string) (map[string]string, map[string][]byte) {
	textFiles := make(map[string]string)
	images := make(map[string][]byte)
	for _, originalPath := range filePaths(root) {
		path := originalPath[len(root)+1:]
		sep := string(os.PathSeparator)
		if sep != "/" {
			path = strings.Replace(path, sep, "/", -1)
		}
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

func writeImage(path string, imgBinary []byte, buffer *bytes.Buffer) {
	buffer.WriteString(`		"` + path + `": {`)
	for _, b := range imgBinary {
		buffer.WriteString(strconv.Itoa(int(b)) + ", ")
	}
	buffer.WriteString("}" + lineSeparator)
}

func writeTextFile(path string, content string, buffer *bytes.Buffer) {
	buffer.WriteString(`		"` + path + `": ` + "`")
	for _, c := range content {
		if string(c) == "`" {
			buffer.WriteString("` + " + `"` + "`" + `" + ` + "`")
		} else {
			buffer.WriteString(string(c))
		}
	}
	buffer.WriteString("`" + lineSeparator)
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

func sortedFilePaths(sources map[string]string) []string {
	var paths []string
	for path := range sources {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	return paths
}

func sortedImagePaths(images map[string][]byte) []string {
	var paths []string
	for path := range images {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	return paths
}

const prefix = `// THIS CODE IS GENERATED, DO NOT EDIT

// Package 'static' contains static assets such as
// source code, text files or images generated form
// the 'templates' folder in the root of the repository.
// If a change is made in templates regenerate this file
// by running 'transform/transformer.go'.
package static

func Sources() map[string]string {
	return map[string]string{
`
const lineSeparator = `,
`
const middle = `	}
}

func Images() map[string][]byte {
	return map[string][]byte{
`

const suffix = `	}
}`

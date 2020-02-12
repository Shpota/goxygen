package main

import (
	"testing"
)

func TestSortedImagePaths(t *testing.T) {
	images := map[string][]byte{
		"/path2": {2},
		"/path3": {2},
		"/path1": {1},
	}

	paths := sortedImagePaths(images)

	if paths[0] != "/path1" || paths[2] != "/path3" {
		t.Error("paths are not sorted")
	}
}

func TestSortedFilePaths(t *testing.T) {
	sources := map[string]string{
		"/path2": "src",
		"/path3": "src",
		"/path1": "src",
	}

	paths := sortedFilePaths(sources)

	if paths[0] != "/path1" || paths[2] != "/path3" {
		t.Error("paths are not sorted")
	}
}

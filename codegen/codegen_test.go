package codegen

import (
	"testing"
)

func TestInclude(t *testing.T) {
	got := include("react.webapp/test.js", "react")

	if !got {
		t.Error("Must include framework specific paths")
	}
}

func TestIncludeGivenGeneralPath(t *testing.T) {
	got := include("server/server.go", "react")

	if !got {
		t.Error("Must include general paths")
	}
}

func TestIncludeGivenExcludedPath(t *testing.T) {
	got := include("angular.webapp/test.js", "vue")

	if got {
		t.Error("Must not include angular files for vue")
	}
}

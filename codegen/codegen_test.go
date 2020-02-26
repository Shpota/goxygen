package codegen

import (
	"testing"
)

func TestNeeded(t *testing.T) {
	g := generator{frontend: "react"}

	got := g.needed("react.webapp/test.js")

	if !got {
		t.Error("Must needed framework specific paths")
	}
}

func TestNeededGivenGeneralPath(t *testing.T) {
	g := generator{frontend: "react"}

	got := g.needed("server/server.go")

	if !got {
		t.Error("Must needed general paths")
	}
}

func TestNeededGivenExcludedPath(t *testing.T) {
	g := generator{frontend: "vue"}

	got := g.needed("angular.webapp/test.js")

	if got {
		t.Error("Must not needed angular files for vue")
	}
}

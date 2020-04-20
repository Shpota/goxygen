package codegen

import (
	"testing"
)

func TestNeeded(t *testing.T) {
	g := generator{techStack: []string{"react"}}

	got := g.needed("react.webapp/test.js")

	if !got {
		t.Error("Framework specific paths must be needed")
	}
}

func TestNeededGivenGeneralPath(t *testing.T) {
	g := generator{techStack: []string{"react"}}

	got := g.needed("server/server.go")

	if !got {
		t.Error("General paths must be needed")
	}
}

func TestNeededGivenExcludedPath(t *testing.T) {
	g := generator{techStack: []string{"react"}}

	got := g.needed("angular.webapp/test.js")

	if got {
		t.Error("Angular files must not be needed for react")
	}
}

func Test_hasTechPrefix(t *testing.T) {
	got := hasTechPrefix("mongo.server/go.mod")

	if !got {
		t.Errorf("Must recognize prefix")
	}
}

func Test_hasTechPrefix_GeneralPath(t *testing.T) {
	got := hasTechPrefix("unknown.server/go.mod")

	if got {
		t.Errorf("Must not recognize prefix")
	}
}

//go:generate go run transform/transformer.go

package main

import (
	"github.com/shpota/goxygen/cli"
	"github.com/shpota/goxygen/codegen"
	"os"
)

func main() {
	args := os.Args[1:]
	cli.Start(os.Stdout, args, codegen.Generate)
}

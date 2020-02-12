package main

import (
	"github.com/shpota/goxygen/cli"
	"os"
)

func main() {
	cli.Start(os.Stdout, os.Args[1:])
}

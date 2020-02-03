package main

import (
	"github.com/shpota/goxygen/cli"
	"os"
)

func main() {
	cli.Start(os.Args[1:])
}

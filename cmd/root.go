package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
	Use: "goxygen",
	Short: `Generate a modern web project with Go, React and MongoDB in seconds.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
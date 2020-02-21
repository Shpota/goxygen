package cmd

import (
	"errors"
	"fmt"
	"github.com/shpota/goxygen/codegen"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"regexp"
)

var projectName string

var initCmd = &cobra.Command{
	Use: "init",
	Short: `Generates a web project with the following structure:
server side Go code, webapp with corresponding React
components, a Dockerfile that builds the application,
and Docker Compose scripts to run the app in dev and
prod modes.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(projectName) <= 0 {
			return errors.New(`--name or -n is a required flag for init command`)
		}
		validName, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", projectName)
		if !validName {
			return errors.New(`Project name is not valid. The allowed symbols are letters,
numbers, underscores, and dashes.`)
		}
		codegen.Generate(projectName)
		return nil
	},
}

func usage (cmd *cobra.Command) error {
	fmt.Println(`Usage:
goxygen init -n <project-name>

Available Commands:
  init        Generate a modern web project with Go, React and MongoDB in seconds.
  help        Help about any command

Flags:
  -n, --name string    Project name
  -h, --help           help for goxygen

Use "goxygen [command] --help" for more information about a command.`)
	return nil
}


func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.SetUsageFunc(usage)
	initCmd.PersistentFlags().StringVarP(&projectName, "name", "n", "", "Project name (required)")
	_ = viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
	_ = initCmd.MarkFlagRequired("name")
}
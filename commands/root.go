package commands

import (
	"github.com/spf13/cobra"

	"fmt"
	"os"
)

var version string

var rootCmd = &cobra.Command{
	Use:   "weaver",
	Short: "A Markdown/knitr utility",
	Long:  "A tool to generate HTML documents from Markdown/RMarkdown",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s -- see '%s help'\n",
			"Please specify a subcommand", os.Args[0])
	},
}

func Execute(semanticVersion string) {
	version = semanticVersion
	AddCommands()
	rootCmd.Execute()
}

func AddCommands() {
	rootCmd.AddCommand(cmdSetup)
	rootCmd.AddCommand(cmdVersion)
	rootCmd.AddCommand(cmdRender)
}

package commands

import (
	"github.com/spf13/cobra"

	"fmt"
)

type SetupCmdOpts struct {
	RLibs string
}

var SetupOpts SetupCmdOpts

var cmdSetup = &cobra.Command{
	Use:   "setup ...",
	Short: "Gather the necessary R libraries for HTML publishing",
	Long:  "Download and Setup the necessary R libraries for HTML publishing",
	Run: func(cmd *cobra.Command, args []string) {
		SetupOpts.main()
	},
}

func init() {
	cmdSetup.Flags().StringVarP(
		&SetupOpts.RLibs,
		"rlibs",
		"l",
		"",
		"the directory to install and use external R packages",
	)
}

func (opts SetupCmdOpts) main() {
	fmt.Println("Nothing implemented yet!")
}

func (opts SetupCmdOpts) processOpts() {
}

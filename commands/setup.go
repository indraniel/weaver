package commands

import (
	"github.com/indraniel/weaver/rsetup"
	"github.com/indraniel/weaver/utils"

	"github.com/spf13/cobra"

	"log"
)

type SetupCmdOpts struct {
	RLibs       string
	InstallType string
	Force       bool
}

var SetupOpts SetupCmdOpts

var cmdSetup = &cobra.Command{
	Use:   "Rsetup ...",
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
		"r",
		"./rpkgs",
		"the directory to install and use external R packages",
	)

	cmdSetup.Flags().StringVarP(
		&SetupOpts.InstallType,
		"install",
		"i",
		"pinned",
		"the type of install ('pinned', 'dev', 'latest')",
	)

	cmdSetup.Flags().BoolVarP(
		&SetupOpts.Force,
		"force",
		"f",
		false,
		"overwrite the files already in rlibs",
	)
}

func (opts SetupCmdOpts) main() {
	opts.processOpts()
	rsetup.SetupRPackages(opts.RLibs, opts.InstallType)
}

func (opts SetupCmdOpts) processOpts() {
	if (utils.DoesExist(opts.RLibs)) && (opts.Force == false) {
		log.Fatalf(
			"R libs directory '%s' already exists."+
				"Use '--force' to overwrite, or specify"+
				"a different directory",
			opts.RLibs,
		)
	}

	if !((opts.InstallType == "pinned") ||
		(opts.InstallType == "dev") ||
		(opts.InstallType == "latest")) {
		log.Fatalf(
			"--install option: '%s' isn't a valid type."+
				"it has to be either: %s",
			opts.InstallType,
			"'pinned', 'dev', or 'latest'",
		)
	}
}

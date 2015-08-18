package commands

import (
	"github.com/indraniel/weaver/render"

	"github.com/spf13/cobra"

	"fmt"
	"log"
	"os"
)

type RenderCmdOpts struct {
	OutDir string
}

var RenderOpts RenderCmdOpts

var cmdRender = &cobra.Command{
	Use:   "render -o <output-dir> MARKDOWN-FILE",
	Short: "Render Markdown/RMarkdown file(s) to HTML",
	Long:  "Rener Markdown/RMarkdown file(s) to HTML at output-directory",
	Run: func(cmd *cobra.Command, args []string) {
		RenderOpts.main(args)
	},
}

func init() {
	cmdRender.Flags().StringVarP(
		&RenderOpts.OutDir,
		"output-dir",
		"o",
		".",
		"the directory to output HTML to (default '.')",
	)
}

func (opts RenderCmdOpts) main(args []string) {
	CheckExists(opts.OutDir)

	for _, file := range args {
		fmt.Println("Processing --", file)
		CheckExists(file)
		render.RenderFile(file, opts.OutDir)
	}

	fmt.Println("All Done!")
}

func (opts RenderCmdOpts) processOpts() {
}

func CheckExists(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%d' on file system: %s",
			file, err,
		)
	}
}

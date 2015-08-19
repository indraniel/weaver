package commands

import (
	"github.com/indraniel/weaver/render"
	"github.com/indraniel/weaver/utils"

	"github.com/spf13/cobra"

	"fmt"
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
	utils.CheckExists(opts.OutDir)

	for _, file := range args {
		fmt.Println("Processing --", file)
		render.RenderFile(file, opts.OutDir)
		utils.CheckExists(file)
	}

	fmt.Println("All Done!")
}

func (opts RenderCmdOpts) processOpts() {
}

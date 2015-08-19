package commands

import (
	"github.com/indraniel/weaver/render"
	"github.com/indraniel/weaver/utils"

	"github.com/spf13/cobra"

	"fmt"
)

type RenderCmdOpts struct {
	OutDir   string
	RpkgsDir string
	Keep     bool
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

	cmdRender.Flags().StringVarP(
		&RenderOpts.RpkgsDir,
		"rpkgs",
		"r",
		"",
		"the directory to use R packages from"+
			"(overrides R_LIBS environment variable)",
	)

	cmdRender.Flags().BoolVarP(
		&RenderOpts.Keep,
		"keep",
		"k",
		false,
		"keep the intermediary files on the file system",
	)
}

func (opts RenderCmdOpts) main(args []string) {
	utils.CheckExists(opts.OutDir)

	for _, file := range args {
		fmt.Println("Processing --", file)
		utils.CheckExists(file)
		render.RenderFile(file, opts.OutDir, opts.RpkgsDir, opts.Keep)
	}

	fmt.Println("All Done!")
}

func (opts RenderCmdOpts) processOpts() {
}

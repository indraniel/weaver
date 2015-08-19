package render

import (
	"github.com/indraniel/weaver/utils"

	"fmt"
	"log"
	"path/filepath"
)

func RenderFile(inputFile, outDir, rpkgsDir string, keep bool) {
	fmt.Println("Rendering", inputFile)

	extension := filepath.Ext(inputFile)
	switch extension {
	case ".md":
		fmt.Println("Going to render straight up markdown")
		RenderMarkdownFile(inputFile, outDir)
	case ".mkdn":
		fmt.Println("Going to render straight up markdown")
		RenderMarkdownFile(inputFile, outDir)
	case ".markdown":
		fmt.Println("Going to render straight up markdown")
		RenderMarkdownFile(inputFile, outDir)
	case ".Rmd":
		fmt.Println("Going to render RMarkdown")
		RenderRMarkdownFile(inputFile, outDir, rpkgsDir, keep)
	case ".rmd":
		fmt.Println("Going to render RMarkdown")
		RenderRMarkdownFile(inputFile, outDir, rpkgsDir, keep)
	case ".Rmkdn":
		fmt.Println("Going to render RMarkdown")
		RenderRMarkdownFile(inputFile, outDir, rpkgsDir, keep)
	case ".rmkdn":
		fmt.Println("Going to render RMarkdown")
		RenderRMarkdownFile(inputFile, outDir, rpkgsDir, keep)
	case ".Rmarkdown":
		fmt.Println("Going to render RMarkdown")
		RenderRMarkdownFile(inputFile, outDir, rpkgsDir, keep)
	case ".rmarkdown":
		fmt.Println("Going to render RMarkdown")
		RenderRMarkdownFile(inputFile, outDir, rpkgsDir, keep)
	default:
		log.Fatalf("Don't know how to render file of type: '%s'",
			extension,
		)
	}
}

func RenderRMarkdownFile(inputFile, outDir, rpkgsDir string, keep bool) {
	SetupRLIBSPath(rpkgsDir)
	rdoc := NewRDocument(inputFile, outDir)
	outputMD := rdoc.KnitR()
	RenderMarkdownFile(outputMD, outDir)
	if keep == false {
		utils.Cleanup(outputMD)
	}
}

func RenderMarkdownFile(inputFile, outDir string) {
	doc := NewDocument(inputFile)
	doc.ParseDocument()
	doc.RenderHTML(outDir)
}

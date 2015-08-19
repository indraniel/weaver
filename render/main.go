package render

import (
	"github.com/indraniel/weaver/utils"

	"fmt"
	"log"
	"path/filepath"
)

func RenderFile(inputFile, outDir string) {
	fmt.Println("Rendering", inputFile)

	extension := filepath.Ext(inputFile)
	switch extension {
	case ".md":
		fmt.Println("Going to render straight up markdown")
		RenderMarkdownFile(inputFile, outDir)
	case ".mkdn":
		fmt.Println("Going to render straight up markdown")
		RenderMarkdownFile(inputFile, outDir)
	case ".Rmd":
		fmt.Println("Going to render RMarkdown")
	case ".rmd":
		fmt.Println("Going to render RMarkdown")
	case ".Rmkdn":
		fmt.Println("Going to render RMarkdown")
	case ".rmkdn":
		fmt.Println("Going to render RMarkdown")
	default:
		log.Fatalf("Don't know how to render file of type: '%s'",
			extension,
		)
	}
}

func RenderRMarkdownFile(inputFile, outDir string) {
	log.Fatal("Please implement me!")
}

func RenderMarkdownFile(inputFile, outDir string) {
	doc := NewDocument(inputFile)
	doc.ParseDocument()
	doc.RenderHTML(outDir)
}

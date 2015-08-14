package render

import (
	"fmt"
	"log"
	"path/filepath"
	"time"
)

func RenderFile(inputFile, outDir string) {
	fmt.Println("Rendering", inputFile)
	extension := filepath.Ext(inputFile)
	switch extension {
	case ".md":
		fmt.Println("Going to render straight up markdown")
	case ".mkdn":
		fmt.Println("Going to render straight up markdown")
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
	doc := ParseMarkdownFile(inputFile)
	doc.ProcessMarkdown()
	doc.renderHTML()
	doc.Save(outDir)
}

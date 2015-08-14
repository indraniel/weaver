package render

import (
	"github.com/shurcooL/github_flavored_markdown"
)

type Document struct {
	FileBaseName string
	Title        string
	RawBody      string
	HTMLBody     string
	Date         time.Time
}

func ParseMarkdownFile(inputFile string) {
}

func NewDocument() *Document {
	doc = new(Document{})
	return doc
}

func (d Document) ProcessMarkdown() {
}

func (d Document) RenderHTML() {
}

func (d Document) Save(outDir string) {
}

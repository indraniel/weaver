package render

import (
	"github.com/shurcooL/github_flavored_markdown"

	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Document struct {
	InputFile string
	BaseName  string
	Title     string
	RawBody   string
	HTMLBody  template.HTML
	FullHTML  string
	Date      time.Time
	Updated   time.Time
	Params    map[string]string
}

func NewDocument(inputFile string) *Document {
	basename := path.Base(inputFile)
	doc := &Document{
		InputFile: inputFile,
		BaseName:  basename,
		Date:      time.Now(),
		Updated:   time.Now(),
		Params:    make(map[string]string),
	}
	return doc
}

func (d *Document) ParseDocument() {
	data, err := ioutil.ReadFile(d.InputFile)
	if err != nil {
		log.Fatalf("Couldn't read file: '%s': %s", d.InputFile, err)
	}
	startContentLine := d.parseHeader(data)
	d.parseContents(data, startContentLine)
}

func (d *Document) parseHeader(data []byte) int {
	lines := strings.Split(string(data), "\n")

	headerState := 0
	startMdContentLine := 0
	for i, line := range lines {
		line = strings.TrimSpace(line)

		// in the header section: from --- to ---
		if headerState == 1 {
			colonIndex := strings.Index(line, ":")
			if colonIndex > 0 {
				key := strings.TrimSpace(line[:colonIndex])
				value := strings.TrimSpace(line[colonIndex+1:])
				// remove surrounding quotes
				value = strings.Trim(value, "\"")
				switch key {
				case "Title":
					d.Title = value
				case "Date":
					d.Date, _ =
						time.Parse("2006-01-02", value)
				case "Updated":
					d.Updated, _ =
						time.Parse("2006-01-02", value)
				default:
					d.Params[key] = value
				}
			}
		} else if headerState >= 2 {
			startMdContentLine = i
			break
		}

		if strings.HasPrefix(line, "---") {
			headerState += 1
		}
	}

	return startMdContentLine
}

func (d *Document) parseContents(data []byte, startLine int) {
	lines := strings.Split(string(data), "\n")
	contents := strings.Join(lines[startLine:], "\n")
	d.RawBody = contents
}

func (d *Document) RenderHTML(outDir string) {
	text := []byte(d.RawBody)
	d.HTMLBody = template.HTML(github_flavored_markdown.Markdown(text))
	base := strings.TrimSuffix(d.InputFile, filepath.Ext(d.InputFile))
	htmlFile := strings.Join([]string{base, "html"}, ".")
	htmlPath := filepath.Join(outDir, htmlFile)

	f, err := os.Create(htmlPath)
	if err != nil {
		log.Fatalf("Couldn't open '%s' : %s\n", htmlPath, err)
	}
	defer f.Close()

	baseTemplate := "assets/views/base.html"
	t, err := template.ParseFiles(baseTemplate)
	if err != nil {
		log.Fatalf("Template '%s' parse error: %s\n", baseTemplate, err)
	}
	err = t.Execute(f, d)
	if err != nil {
		log.Fatalf("Template render error: %s\n", err)
	}

	// alternative approach:
	// http://stackoverflow.com/questions/23124008/how-can-i-render-markdown-to-a-golang-templatehtml-or-tmpl-with-blackfriday
}

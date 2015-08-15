package render

import (
	"io/ioutil"
	"log"
	"path"
	"strings"
	"time"

	"github.com/shurcooL/github_flavored_markdown"
)

type Document struct {
	InputFile string
	BaseName  string
	Title     string
	RawBody   string
	HTMLBody  string
	FullHTML  string
	Date      time.Time
	Params    map[string]string
}

func ParseMarkdownFile(inputFile string) {
}

func NewDocument(inputFile) *Document {
	basename := path.Base(inputFile)
	doc = new(
		Document{
			InputFile: inputFile,
			BaseName:  basename,
			Date:      time.Now(),
			Params:    make(map[string]string),
		},
	)
	return doc
}

func (d *Document) ParseDocument() {
	data, err := ioutil.ReadFile(d.InputFile)
	if err != nil {
		log.Fatalf("Couldn't read file: '%s': %s", d.InputFile, err)
	}
	startContentLine := d.parseHeader(data)
	d.parseContents(startContentLine)
}

func (d *Document) parseHeader(data []byte) int {
	lines := strings.Split(string(data), "\n")

	headerState := 0
	startMdContentLine = 0
	for i, line := range lines {
		line = strings.TrimSpace(line)

		// in the header section: from --- to ---
		if headerState == 1 {
			colonIndex = strings.Index(line, ":")
			if colonIndex > 0 {
				key := strings.TrimSpace(line[:colonIndex])
				value := strings.TrimSpace(line[colonIndex+1:])
				// remove surrounding quotes
				value = strings.Trim(value, "\"")
				switch key {
				case "title":
					d.Title = value
				case "date":
					d.Date, _ =
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
	contents = strings.Join(lines[startLine:], "\n")
	d.RawBody = contents
}

func (d *Document) RenderHTML() {
	text = []byte(d.RawBody)
	d.HTMLBody = string(github_flavored_markdown.Markdown(text))
}

func (d Document) Save(outDir string) {
}

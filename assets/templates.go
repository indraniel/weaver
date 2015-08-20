package assets

import (
	"github.com/arschles/go-bindata-html-template"

	"log"
)

type View struct {
	TemplateName string
}

func NewView(name string) *View {
	_, err := Asset(name)
	if err != nil {
		log.Fatalf("Asset: '%s' : doesn't exist", name)
	}

	view := &View{
		TemplateName: name,
	}
	return view
}

func (v View) GetTemplate() *template.Template {
	t := template.New(v.TemplateName, Asset)
	return t
}

package render

import (
	"github.com/indraniel/weaver/utils"

	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type RDocument struct {
	InputFile string
	BaseName  string
	OutDir    string
}

func NewRDocument(inputFile, outDir string) *RDocument {
	basename := path.Base(inputFile)
	rdoc := &RDocument{
		InputFile: inputFile,
		BaseName:  basename,
		OutDir:    outDir,
	}
	return rdoc
}

func (r RDocument) KnitR() string {
	r.checkR()
	r.knit()
	base := strings.TrimSuffix(r.InputFile, filepath.Ext(r.InputFile))
	pureMarkdownFile := strings.Join([]string{base, "md"}, ".")
	utils.CheckExists(pureMarkdownFile)
	return pureMarkdownFile
}

func (r RDocument) checkR() {
	cmd := "which"
	args := []string{"Rscript"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Fatalln("Couldn't find Rscript: %s", err)
	}
}

func (r RDocument) knit() {
	prgrm := "assets/Rscripts/knit.r"
	args := []string{"--input", r.InputFile, "--outdir", r.OutDir}
	log.Printf("Running: %s %s", prgrm, strings.Join(args, " "))
	cmd := exec.Command(prgrm, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Trouble with %s:\n\n%s\n\nOutput:\n%s\n\n",
			prgrm, err, string(output))
	}
}

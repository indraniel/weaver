package assets

import (
	"github.com/indraniel/weaver/utils"

	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type RScript struct {
	RootPath   string
	ScriptName string
	AssetName  string
}

func NewRScript(name string, rootPath ...string) *RScript {
	root := os.Getenv("TMPDIR")
	if len(rootPath) > 0 {
		root = rootPath[0]
	}

	script := fmt.Sprintf("Rscripts/%s", name)
	_, err := Asset(script)
	if err != nil {
		log.Fatalf("Asset: '%s' : doesn't exist", name)
	}

	rscript := &RScript{
		ScriptName: name,
		RootPath:   root,
		AssetName:  script,
	}
	return rscript
}

func (r RScript) WrapperPath() string {
	script := path.Join(r.RootPath, r.ScriptName)
	return script
}

func (r RScript) WriteWrapper() string {
	script := r.WrapperPath()
	log.Printf("Writing '%s' to %s", r.ScriptName, script)
	data, err := Asset(r.AssetName)
	if err != nil {
		log.Fatalf("Trouble retrieving asset: %s", r.AssetName)
	}
	err = ioutil.WriteFile(script, []byte(data), 0755)
	if err != nil {
		log.Fatalf("Trouble writing asset: %s to %s",
			r.AssetName, script,
		)
	}
	return script
}

func (r RScript) Remove() {
	script := r.WrapperPath()
	log.Println("Deleting", script)
	utils.Cleanup(script)
}

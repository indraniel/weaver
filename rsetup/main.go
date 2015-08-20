package rsetup

import (
	"github.com/indraniel/weaver/assets"
	"github.com/indraniel/weaver/utils"

	"log"
	"os"
	"os/exec"
	"strings"
)

func SetupRPackages(pkgDir, installType string) {
	setupDir(pkgDir)
	installPackages(installType)
}

func setupDir(dir string) {
	if utils.DoesExist(dir) == false {
		log.Printf("Creating directory: %s", dir)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf(
				"Trouble creating directory: '%s' : %s",
				dir, err,
			)
		}
	}

	log.Printf("Setting R_LIBS to: %s", dir)
	if err := os.Setenv("R_LIBS", dir); err != nil {
		log.Fatalf(
			"Could not set R_LIBS: %s",
			err,
		)
	}

	Rlibs := os.Getenv("R_LIBS")
	if Rlibs != dir {
		log.Fatalln(
			"Environment Variable 'R_LIBS' isn't correctly set.  ",
		)
	}
}

func installPackages(installType string) {
	log.Println("Installing R packages")

	script := assets.NewRScript("install-packages.r")
	defer script.Remove()

	prgrm := script.WriteWrapper()
	args := []string{installType}
	log.Printf("Running: %s %s", prgrm, strings.Join(args, " "))
	cmd := exec.Command(prgrm, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Trouble starting %s:\n\n%s\n\n",
			prgrm, err)
	}
	if err = cmd.Wait(); err != nil {
		log.Fatalf("Trouble running %s:\n\n%s\n\n",
			prgrm, err)
	}
}

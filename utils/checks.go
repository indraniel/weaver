package utils

import (
	"log"
	"os"
)

func CheckExists(file string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Fatalf(
			"Could not find '%d' on file system: %s",
			file, err,
		)
	}
}

func DoesExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func Cleanup(file string) {
	if err := os.Remove(file); err != nil {
		log.Fatalf(
			"Trouble removing intermediate file: '%s' : %s",
			file, err,
		)
	}
}

func MakeDirs(path string, perms os.FileMode) {
	err := os.MkdirAll(path, perms)
	if err != nil {
		log.Fatalf(
			"Trouble creating path: '%s' (%#o) : %s",
			path, perms, err,
		)
	}
}

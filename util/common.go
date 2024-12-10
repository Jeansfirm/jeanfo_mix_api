package util

import (
	"log"
	"os"
	"path/filepath"
)

func GetExeDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("get exe dir failed: %s", err.Error())
	}

	return filepath.Dir(ex)
}

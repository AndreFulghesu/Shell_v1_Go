package constants

import (
	"os"
)

// errors handling in next features :)
var CurrentDir = initCurrentDir()

func UpdateCurrentDir(newDir string) {
	CurrentDir = newDir
}

func initCurrentDir() string {
	CurrentDir, err := os.Getwd()
	if err == nil {
		return CurrentDir
	}
	return ""
}

package constants

import (
	"os"
)

// errors handling in next features :)
var CurrentDir, _ = os.Getwd()

func UpdateCurrentDir(newDir string) {
	CurrentDir = newDir
}

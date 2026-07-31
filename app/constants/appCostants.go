package constants

import (
	"path/filepath"
)

// errors handling in next features :)
var CurrentDir, _ = filepath.Abs(".")

func UpdateCurrentDir(newDir string) {
	CurrentDir = newDir
}

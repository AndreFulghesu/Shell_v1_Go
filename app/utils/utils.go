package utils

import (
	"os/exec"
)

/* check for path ENV commands */
func CheckEnvCommand(command_path string) (present bool, path string) {
	path, err := exec.LookPath(command_path)
	if err == nil {
		return true, path
	}
	return false, ""
}

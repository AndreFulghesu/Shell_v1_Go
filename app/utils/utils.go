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

func ParseArgs(input string) []string {
	var args []string
	var current []rune
	inQuote := false

	for _, r := range input {
		switch r {
		case '\'':
			inQuote = !inQuote
		case ' ':
			if !inQuote && len(current) > 0 {
				args = append(args, string(current))
				current = nil
			} else {
				current = append(current, r)
			}
		default:
			current = append(current, r)
		}
	}

	if len(current) > 0 {
		args = append(args, string(current))
	}

	return args
}

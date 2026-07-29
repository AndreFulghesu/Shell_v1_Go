package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/builtin"
	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/constants"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

var _ = fmt.Print
var reader = bufio.NewReader(os.Stdin)

var loop = true

func main() {

	for loop == true {

		fmt.Print("$ ")
		command, err := reader.ReadString('\n')

		if err == nil {
			evaluate_command(command)
		}
	}

}

func evaluate_command(command string) {

	splitted := strings.Split(strings.TrimSpace(command), " ")

	base_command := splitted[0]
	args := splitted[1:]

	is_env_command, path := utils.CheckEnvCommand(base_command)

	if is_env_command {
		cmd := exec.Command(path, args...)

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
		return
	}

	switch base_command {
	case commands.EXIT:
		os.Exit(0)
	case commands.ECHO:
		builtin.Echo_command(args)
	case commands.TYPE:
		builtin.Type_command(args)
	default:
		fmt.Printf(constants.COMMAND_NOT_FOUND, base_command)
	}
}

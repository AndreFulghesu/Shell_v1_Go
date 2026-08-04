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
			evaluateCommand(strings.TrimSpace(command))
		}
	}

}

func evaluateCommand(command string) {

	args := utils.ParseArgs(command)

	baseCommand := args[0]
	args = args[1:]

	_, isBuiltIn := commands.Find(baseCommand)

	if isBuiltIn {
		switch baseCommand {

		case commands.EXIT:
			os.Exit(0)
		case commands.ECHO:
			builtin.EchoCommand(args)
		case commands.TYPE:
			builtin.TypeCommand(args)
		case commands.PWD:
			fmt.Println("ENTRATO NEL CASE PWD")
			builtin.PwdCommand()
		case commands.CD:
			builtin.CdCommand(args)
		default:
			fmt.Printf(constants.COMMAND_NOT_FOUND, baseCommand)
		}
		return
	}

	/* if base_command is an PATH ENV command */
	isEnvCommand, _ := utils.CheckEnvCommand(baseCommand)

	if isEnvCommand {
		cmd := exec.Command(baseCommand, args...)

		output, _ := cmd.CombinedOutput()
		fmt.Print(string(output))
		return
	}
}

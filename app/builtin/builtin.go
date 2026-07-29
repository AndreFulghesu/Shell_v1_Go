package builtin

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/constants"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func Echo_command(args []string) {
	fmt.Print(strings.Join(args, " "), "\n")
}

func Type_command(args []string) {
	command_name := args[0]

	/* check built in commands */
	_, exists := commands.Find(command_name)

	if exists {
		fmt.Printf(constants.BUILT_IN, command_name)
		return
	}

	/* check ENV PATH executable command */
	is_present, path := utils.CheckEnvCommand(command_name)

	if is_present {
		fmt.Printf(constants.PATH_COMMAND, command_name, path)
	} else {
		fmt.Printf(constants.NOT_FOUND, command_name)
	}
}

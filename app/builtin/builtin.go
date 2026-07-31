package builtin

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/constants"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func EchoCommand(args []string) {
	fmt.Print(strings.Join(args, " "), "\n")
}

func TypeCommand(args []string) {
	commandName := args[0]

	/* check built in commands */
	_, exists := commands.Find(commandName)

	if exists {
		fmt.Printf(constants.BUILT_IN, commandName)
		return
	}

	/* check ENV PATH executable command */
	is_present, path := utils.CheckEnvCommand(commandName)

	if is_present {
		fmt.Printf(constants.PATH_COMMAND, commandName, path)
	} else {
		fmt.Printf(constants.NOT_FOUND, commandName)
	}
}

func PwdCommand() {
	path, _ := filepath.Abs(".")
	fmt.Println(path)
}

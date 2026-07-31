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
	fmt.Println(constants.CurrentDir)
}

func CdCommand(newPath []string) {

	argument := newPath[0]
	//handling empty args
	if len(newPath[0]) == 0 {
		return
	}

	//Handling back 1 dir position
	if argument == ".." || argument == "../" {
		//delete path last element
		path := filepath.Dir(constants.CurrentDir)
		constants.UpdateCurrentDir(path)
		return
	}

	constants.UpdateCurrentDir(argument)
	/*
		info, error := os.Stat(argument)

		//check if path exists
		if error != nil {
			if os.IsNotExist(error) {
				fmt.Printf(constants.PATH_NOT_FOUND, "cd", "newPath")
			}
			//handle errors in future releases :)
			return
		}
		//path is a directory
		if info.IsDir() {
			constants.UpdateCurrentDir(argument)
			return
		}

		//path, _ = filepath.Abs(newPath)
	*/
}

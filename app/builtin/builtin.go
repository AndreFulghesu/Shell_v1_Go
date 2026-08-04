package builtin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/constants"
	"github.com/codecrafters-io/shell-starter-go/app/utils"
)

func EchoCommand(args []string) {
	var builtString []string
	fmt.Printf("%q\n", args)

	builtString = strings.Fields(strings.Join(args, " "))
	fmt.Printf("%q\n", builtString)
	fmt.Println(strings.TrimSpace(strings.Join(builtString, " ")))
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
	current, _ := os.Getwd()
	fmt.Println(current)
}

func CdCommand(newPath []string) {

	argument := newPath[0]
	//handling empty args
	if len(newPath[0]) == 0 {
		return
	}

	//Handling back 1 dir position
	if argument == ".." || argument == "../" {
		current, _ := os.Getwd()
		//delete path last element
		path := filepath.Dir(current)
		os.Chdir(path)
		return
	}

	if argument == "~" {
		//get HOME env variable
		home, err := os.UserHomeDir()
		if err == nil {
			os.Chdir(home)
		}
		return
	}

	_, error := os.Stat(argument)
	//check if path exists
	if error != nil {
		if os.IsNotExist(error) {
			fmt.Printf(constants.PATH_NOT_FOUND, "cd", argument)
			return
		}
		//handle errors in future releases :)
		fmt.Print("Other type of error\n")
		return
	}

	os.Chdir(argument)
}

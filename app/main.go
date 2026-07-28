package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/constants"
)

var _ = fmt.Print
var not_found_msg = ": command not found"
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

	switch base_command {
	case commands.EXIT:
		os.Exit(0)
	case commands.ECHO:
		echo_command(args)
	case commands.TYPE:
		type_command(args)
	default:
		fmt.Printf(constants.COMMAND_NOT_FOUND, base_command)
	}
}

func echo_command(args []string) {
	fmt.Print(strings.Join(args, " "), "\n")
}

func type_command(args []string) {

	_, exists := commands.Find(args[0])

	if exists {
		fmt.Printf(constants.BUILT_IN, args[0])
	} else {
		fmt.Printf(constants.NOT_FOUND, args[0])
	}
}

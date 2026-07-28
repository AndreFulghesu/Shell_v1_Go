package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var _ = fmt.Print
var not_found_msg = ": command not found"
var reader = bufio.NewReader(os.Stdin)

var loop = true

var _exit_command = "exit"
var _echo_command = "echo"

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
	case _exit_command:
		os.Exit(0)
	case _echo_command:
		echo_command(args)
	default:
		fmt.Printf("%s%s\n", base_command, not_found_msg)
	}
}

func echo_command(args []string) {
	fmt.Print(strings.Join(args, " "), "\n")
}

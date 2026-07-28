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
	message := strings.TrimSpace(command)

	if message == _exit_command {
		os.Exit(0)
	}

	fmt.Printf("%s%s\n", message, not_found_msg)
}

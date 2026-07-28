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

func main() {

	for loop == true {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')

		if err == nil {
			message := strings.TrimSpace(command)
			fmt.Printf("%s%s\n", message, not_found_msg)
		}
	}

}

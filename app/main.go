package main

import (
	"bufio"
	"fmt"
	"os"
)

var _ = fmt.Print
var not_found_msg = ": command not found"

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err == nil {
		fmt.Print(command, not_found_msg)
	}
}

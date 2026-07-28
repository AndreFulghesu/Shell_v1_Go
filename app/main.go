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

func main() {
	// TODO: Uncomment the code below to pass the first stage
	fmt.Print("$ ")
	command, err := reader.ReadString('\n')

	if err == nil {
		message := strings.TrimSpace(command)
		fmt.Printf("%s%s", message, not_found_msg)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/codecrafters-io/shell-starter-go/app/constants"
)

var _ = fmt.Print
var reader = bufio.NewReader(os.Stdin)

var loop = true

func main() {

	for loop == true {

		fmt.Print("$")
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

	/* check built in commands */
	_, exists := commands.Find(args[0])

	if exists {
		fmt.Printf(constants.BUILT_IN, args[0])
		return
	}

	/* check for path commands */
	path, err := exec.LookPath(args[0])

	if err == nil {
		fmt.Printf(constants.PATH_COMMAND, args[0], path)
		return
	}

	//fmt.Printf(constants.NOT_FOUND, args[0])
}

/*
 return long size path string separated by system separator
	env := os.Getenv("PATH")

	 slice of path splitted
	path_splitted := strings.Split(env, string(os.PathListSeparator))

	* some single path is empty so it must be skipped *
	for _, value := range path_splitted {
		if len(value) != 0 {
			* check if the path exists *
			path_info, error := os.Lstat(value)

			/* check if a exec file exists under this path
			 * with the same name of the typed command
			 *
			if error == nil {
				/* if the path is a directory we need to iterate over the content *
				if path_info.IsDir() {
					entries, err := os.ReadDir(value)
					if err == nil {
						fmt.Println(entries)
					}
				}
			}
		}
	}
*/

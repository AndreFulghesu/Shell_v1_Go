package commands

const (
	EXIT = "exit"
	ECHO = "echo"
	TYPE = "type"
)

var List = []string{
	EXIT,
	ECHO,
	TYPE,
}

func Find(input string) (string, bool) {
	for _, cmd := range List {
		if cmd == input {
			return cmd, true
		}
	}
	return "", false
}

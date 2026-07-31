package commands

const (
	EXIT = "exit"
	ECHO = "echo"
	TYPE = "type"
	PWD  = "pwd"
)

var list = []string{
	EXIT,
	ECHO,
	TYPE,
	PWD,
}

func Find(input string) (string, bool) {
	for _, cmd := range list {
		if cmd == input {
			return cmd, true
		}
	}
	return "", false
}

package displayError

import (
	"fmt"
	"os"
)

func MissingArgsCommandError(program string, command string) {
	fmt.Println("Missing args on command:", program, command)
	fmt.Printf("Run \"%v %v --help\" to get help\n", program, command)
	os.Exit(1)
}

func CommandNotFoundError(program, command string) {
	fmt.Printf("Command \"%v\" not Found\n", command)
	fmt.Printf("Run \"%v --help\" to see available commands\n", program)
	os.Exit(1)
}
func FlagNotFoundError(program, command, flag string) {
	fmt.Printf("Flag \"%v\" not Found\n", flag)
	fmt.Printf("Run \"%v %v --help\" to see available commands\n", program, command)
	os.Exit(1)
}

package add

import (
	"cli-todo/cmdUtil"
	"cli-todo/displayError"
	"fmt"
)

func Add(program string, command string, argsWithOutAddCmd []string) {
	if !cmdUtil.CheckIfArgsPresent(argsWithOutAddCmd) {
		displayError.MissingArgsCommandError(program, command)
	}
	switch argsWithOutAddCmd[0] {
	case "--help":
		fmt.Println("help")
	case "--name":
		fmt.Println("name")
	default:
		displayError.FlagNotFoundError(program, command, argsWithOutAddCmd[0])
	}
}

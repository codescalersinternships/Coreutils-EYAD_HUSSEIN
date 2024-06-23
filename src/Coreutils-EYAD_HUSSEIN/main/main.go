package main

import (
	"Coreutils-EYAD_HUSSEIN/commands"
	"Coreutils-EYAD_HUSSEIN/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) == 1 {
			fmt.Println("No commands were entered!")
			os.Exit(1)
	}

	command := args[1]

	commandModel, exists := utils.CheckCommand(command)

	if !exists {
		os.Exit(1)
	}

	args = os.Args[2:]

	flags := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		}
	}

	switch commandModel.Name {
	// case "head":
	// 	commands.Head(flags)
	case "tail":
		commands.Tail(commandModel, flags)
	// case "wc":
	// 	commands.Wc(flags)
	case "cat":
		commands.Cat(commandModel, flags)
	case "echo":
		commands.Echo(commandModel, flags)
	// case "true":
	// 	commands.True()
	// case "false":
	// 	commands.False()
	default:
		os.Exit(1)
	}

}




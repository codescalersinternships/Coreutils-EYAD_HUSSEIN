package main

import (
	"Coreutils-EYAD_HUSSEIN/commands"
	"Coreutils-EYAD_HUSSEIN/models"
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

	commandModel, exists := utils.ValidateCommand(command)

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
	case models.CommandsMap["head"].Name:
		commands.Head(flags)
	case models.CommandsMap["tail"].Name:
		commands.Tail(flags)
	case models.CommandsMap["wc"].Name:
		commands.Wc(flags)
	case models.CommandsMap["cat"].Name:
		commands.Cat(flags)
	case models.CommandsMap["echo"].Name:
		commands.Echo(flags)
	case models.CommandsMap["true"].Name:
		commands.True()
	case models.CommandsMap["false"].Name:
		commands.False()
	case models.CommandsMap["env"].Name:
		commands.Env()
	case models.CommandsMap["yes"].Name:
		commands.Yes()
	case models.CommandsMap["tree"].Name:
		commands.Tree(flags)
	default:
		os.Exit(1)
	}

}




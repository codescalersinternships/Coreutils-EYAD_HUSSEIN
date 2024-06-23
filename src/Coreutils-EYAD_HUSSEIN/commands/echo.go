package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	"fmt"
	"os"
)

func Echo(flags []string) {
	command := models.CommandsMap["echo"]

	if !utils.ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}
	
	start := 2

	ok := utils.ContainsFlag(flags, "-n")
	idx := utils.GetIndexOfArg("-n")

	flag := (ok && idx == 2)

	if flag {
		start = 3
	}

	for _, arg := range os.Args[start:] {
		fmt.Print(arg, " ")
	}

	if !flag {
		fmt.Println()
	}

}

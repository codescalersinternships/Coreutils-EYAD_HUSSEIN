package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	"fmt"
	"os"
)

func Echo(command *models.Command, flags []string) {
	if !utils.CheckFlags(flags, command.Flags) {
		os.Exit(1)
	}
	
	start := 2
	if utils.ContainsFlag(flags, "-n") {
		start = 3
	}

	for _, arg := range os.Args[start:] {
		fmt.Print(arg, " ")
	}

	if !utils.ContainsFlag(flags, "-n") {
		fmt.Println()
	}

}

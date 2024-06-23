package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	"os"
)

func Echo(command *models.Command, flags []string) {
	if !utils.CheckFlags(flags, command.Flags) {
		os.Exit(1)
	}

	


}

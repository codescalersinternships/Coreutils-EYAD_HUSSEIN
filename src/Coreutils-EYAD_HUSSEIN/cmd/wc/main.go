package main

import (
	"Coreutils-EYAD_HUSSEIN/commands"
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
)

func main() {
	flags := utils.GetFlags()
	utils.ValidateFlags(flags, models.CommandsMap["wc"].Flags)

	commands.Wc(flags)
}
package main

import (
	"Coreutils-EYAD_HUSSEIN/commands"
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	// "fmt"
)

func main() {
	flags := utils.GetFlags()
	// fmt.Println(flags)
	utils.ValidateFlags(flags, models.CommandsMap["echo"].Flags)

	commands.Echo(flags)
}
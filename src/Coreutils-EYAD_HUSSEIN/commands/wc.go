package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"fmt"
	"os"
)

func Wc(command *models.Command, flags []string) {
	if len(os.Args) == 1 {
		fmt.Println("No file was entered!")
		os.Exit(1)
	}
}
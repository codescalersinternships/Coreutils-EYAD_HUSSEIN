package utils

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"fmt"
)


func ValidateCommand(command string) (*models.Command, bool) {

        commands := models.Commands
        exists := false
        var selectedCommand *models.Command

        for i := 0; i < len(commands); i++ {
                if command == commands[i].Name {
                exists = true
                selectedCommand = commands[i]
                break
                }
        }

        if !exists {
                fmt.Println("Invalid command, please stick to the allowed commands")
                fmt.Println("Allowed commands:")
                for index, cmd := range commands {
                fmt.Println(index+1, "-", cmd.Name)
                }
                return nil, false
        }

        return selectedCommand, true
}

func ValidateFlags(flags []string, allowedFlags []string) bool {
        if !validateAllowedFlags(flags, allowedFlags) {
                return false
        }
        if !validateFlagNumbers(flags, allowedFlags) {
                return false
        }
        return true
}

func validateAllowedFlags(flags []string, allowedFlags []string) bool {
        for _, flag := range flags {
                exists := false
                for _, allowedFlag := range allowedFlags {
                        if flag == allowedFlag {
                                exists = true
                                break
                        }
                }
                if !exists {
                        fmt.Println("Invalid flag:", flag)
                        fmt.Println("Allowed flags:")
                        for _, allowedFlag := range allowedFlags {
                                fmt.Println(allowedFlag)
                        }
                        return false
                }
        }
        return true
}

func validateFlagNumbers(flags []string, allowedFlags []string) bool {
        if len(flags) > len(allowedFlags) {
                fmt.Println("Too many flags were entered")
                fmt.Println("Allowed flags:")
                for _, allowedFlag := range allowedFlags {
                        fmt.Println(allowedFlag)
                }
                return false
        }
        return true
}
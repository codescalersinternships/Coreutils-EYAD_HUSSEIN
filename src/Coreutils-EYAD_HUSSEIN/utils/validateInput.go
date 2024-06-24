package utils

import (
	"fmt"
)






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

package utils

import (
	"os"
	"strings"
)

func ContainsFlag(flags []string, flag string) bool {
	for _, f := range flags {
		if f == flag {
			return true
		}
	}
	return false
}

func GetIndexOfArg(arg string) int {
	args := os.Args

	for i := 0; i < len(args); i++ {
		if args[i] == arg {
			return i
		}
	}

	return -1
}

func GetFlags() []string {
	args := os.Args[2:]

	flags := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		}
	}

	return flags
}

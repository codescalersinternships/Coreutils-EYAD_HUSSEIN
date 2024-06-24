package pkg

import (
	"fmt"
	"os"
)

func Echo(flags []string) {
	command := CommandsMap["echo"]

	if !ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}

	start := 1

	ok := ContainsFlag(flags, "-n")
	idx := GetIndexOfArg("-n")

	flag := (ok && idx == 1)

	if flag {
		start = 2
	}

	for _, arg := range os.Args[start:] {
		fmt.Print(arg, " ")
	}

	if !flag {
		fmt.Println()
	}

}

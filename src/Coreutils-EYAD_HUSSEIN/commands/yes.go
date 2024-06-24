package commands

import (
	"os"
	"strings"
)


func Yes() {
	if len(os.Args) == 2 {
		for {
			println("y")
		}
	} else {
		output :=  strings.Join(os.Args[2:], " ")

		for {
			println(output)
		}
	}
}
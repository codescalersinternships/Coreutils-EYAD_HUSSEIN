package pkg

import (
	"os"
	"strings"
)

func Yes() {
	if len(os.Args) == 1 {
		for {
			println("y")
		}
	} else {
		output := strings.Join(os.Args[1:], " ")

		for {
			println(output)
		}
	}
}

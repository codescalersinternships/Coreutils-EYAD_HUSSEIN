package pkg

import (
	"flag"
	"fmt"
	"strings"
)

func Echo() {
	var nFlag bool

	flag.BoolVar(&nFlag, "n", false, "do not output the trailing newline")

	flag.Parse()

	output := strings.Join(flag.Args(), " ")

	if nFlag {
		fmt.Print(output)
	} else {
		fmt.Println(output)
	}
}

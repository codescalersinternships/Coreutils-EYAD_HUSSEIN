package pkg

import (
	"flag"
	"fmt"
)

func Echo() {
	var nFlag bool

	flag.BoolVar(&nFlag, "n", false, "do not output the trailing newline")

	flag.Parse()

	for i, arg := range flag.Args() {
		if i == len(flag.Args())-1 {
			if nFlag {
				fmt.Print(arg)
			} else {
				fmt.Println(arg)
			}
		} else {
			fmt.Print(arg, " ")
		}
	}
}

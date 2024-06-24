package main

import (
	"github.com/codescalersinternships/Coreutils-EYAD_HUSSEIN/pkg"
)

func main() {
	flags := pkg.GetFlags()
	pkg.ValidateFlags(flags, pkg.CommandsMap["tree"].Flags)

	pkg.Tree(flags)
}

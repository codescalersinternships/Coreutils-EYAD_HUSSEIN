package utils

import "os"

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
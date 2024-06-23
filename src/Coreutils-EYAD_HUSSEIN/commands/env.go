package commands

import (
	"fmt"
	"os"
	"strings"
)

func Env(){

	for _, env := range os.Environ() {
		envPair := strings.Split(env, "=")
		fmt.Println(envPair[0], envPair[1])
	}
}
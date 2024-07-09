package pkg

import (
	"fmt"
	"os"
)

func Env() {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

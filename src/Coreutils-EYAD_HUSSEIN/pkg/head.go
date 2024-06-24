package pkg

import (
	"flag"
	"fmt"
	"os"
)

func Head() {
	var numLines int

	flag.IntVar(&numLines, "n", 10, "number of lines to display")

	flag.Parse()

	ok := len(flag.Args()) > 0

	if !ok {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	files := flag.Args()

	processHeadFiles(files, numLines)
}

func processHeadFiles(files []string, numLines int) {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer f.Close()

		if len(files) > 1 {
			fmt.Println("==> " + file + " <==")
		}

		readHeadLines(f, numLines)
		fmt.Println()
	}
}

func readHeadLines(f *os.File, numLines int) {
	counter := 0
	buffer := make([]byte, 4096)
	var line string

	for {
		n, err := f.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if n == 0 {
			break
		}

		content := string(buffer[:n])
		done := false

		for _, char := range content {
			if char == '\n' {
				fmt.Println(line)
				line = ""
				counter++

				if counter == numLines {
					done = true
					break
				}
			} else {
				line += string(char)
			}
		}

		if done {
			break
		}
	}
}

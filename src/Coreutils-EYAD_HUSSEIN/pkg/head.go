package pkg

import (
	"fmt"
	"os"
	"strconv"
)

func Head(flags []string) {
	command := CommandsMap["head"]

	if !ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}

	numLines, files := parseHeadFlags(flags)

	if len(files) == 0 {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	processHeadFiles(files, numLines)
}

func parseHeadFlags(flags []string) (int, []string) {
	ok := ContainsFlag(flags, "-n")
	idx := GetIndexOfArg("-n")

	numLines := 10
	var files []string

	if ok {
		for _, arg := range os.Args[1:] {
			if arg == "-n" || arg == os.Args[idx+1] {
				numLinesTemp, err := strconv.Atoi(os.Args[idx+1])
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}

				numLines = numLinesTemp
				continue
			}

			files = append(files, arg)
		}
	} else {
		files = os.Args[1:]
	}

	return numLines, files
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

package pkg

import (
	"fmt"
	"os"
)

func Cat(flags []string) {
	command := CommandsMap["cat"]

	if !ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}

	ok := ContainsFlag(flags, "-n")
	files := parseCatArgs(ok)

	if len(files) == 0 {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	if !ok {
		displayFiles(files)
	} else {
		displayFilesWithLineNumbers(files)
	}
}

func parseCatArgs(ok bool) []string {
	var files []string

	if ok {
		for _, arg := range os.Args[1:] {
			if arg == "-n" {
				continue
			}
			files = append(files, arg)
		}
	} else {
		files = os.Args[1:]
	}

	return files
}

func displayFiles(files []string) {
	for _, file := range files {
		f, err := os.ReadFile(file)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(string(f))
	}
}

func displayFilesWithLineNumbers(files []string) {
	lineNumber := 1
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		defer f.Close()

		printFileWithLineNumbers(f, &lineNumber)
	}
}

func printFileWithLineNumbers(f *os.File, lineNumber *int) {
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
		for _, char := range content {
			if char == '\n' {
				fmt.Printf("%6d  %s\n", *lineNumber, line)
				*lineNumber++
				line = ""
			} else {
				line += string(char)
			}
		}
	}

	if len(line) > 0 {
		fmt.Printf("%6d  %s\n", *lineNumber, line)
	}
}

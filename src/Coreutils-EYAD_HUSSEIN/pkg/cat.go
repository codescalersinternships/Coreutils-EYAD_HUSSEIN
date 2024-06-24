package pkg

import (
	"flag"
	"fmt"
	"os"
)

func Cat() {

	var nFlag bool

	flag.BoolVar(&nFlag, "n", false, "number all output lines")

	flag.Parse()

	ok := len(flag.Args()) > 0

	if !ok {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	files := flag.Args()

	if !nFlag {
		displayFiles(files)
	} else {
		displayFilesWithLineNumbers(files)
	}
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

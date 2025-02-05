package pkg

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func Tail() {
	var numLines int

	flag.IntVar(&numLines, "n", 10, "number of lines to display")

	flag.Parse()

	if !(len(flag.Args()) > 0) {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	files := flag.Args()

	processTailFiles(files, numLines)
}

func processTailFiles(files []string, numLines int) {
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

		printTailLines(f, numLines)
		fmt.Println()
	}
}

func printTailLines(f *os.File, numLines int) {
	lines := make([]string, numLines)
	lineCount := 0
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines[lineCount%numLines] = scanner.Text()
		lineCount++
	}

	start := lineCount % numLines
	if lineCount < numLines {
		start = 0
	}

	for i := start; i < start+numLines && i < lineCount; i++ {
		fmt.Println(lines[i%numLines])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

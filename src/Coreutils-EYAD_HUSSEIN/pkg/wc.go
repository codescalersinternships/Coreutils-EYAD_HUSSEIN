package pkg

import (
	"fmt"
	"os"
	"unicode"
)

func Wc(flags []string) {
	command := CommandsMap["wc"]

	if !ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}

	files := parseWcArgs()

	if len(files) == 0 {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	var totalLines, totalWords, totalChars int

	for _, file := range files {
		lineCount, wordCount, charCount := countFileStats(file)
		totalLines += lineCount
		totalWords += wordCount
		totalChars += charCount

		printFileStats(flags, file, lineCount, wordCount, charCount)
	}

	printTotalStats(flags, len(files), totalLines, totalWords, totalChars)
}

func parseWcArgs() []string {
	var files []string

	for _, arg := range os.Args[1:] {
		if arg != "-l" && arg != "-w" && arg != "-c" {
			files = append(files, arg)
		}
	}

	return files
}

func countFileStats(file string) (int, int, int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer f.Close()

	buffer := make([]byte, 4096)
	lineCounter := 0
	wordCounter := 0
	charCounter := 0
	inWord := false

	for {
		n, err := f.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if n == 0 {
			break
		}

		for _, char := range string(buffer[:n]) {
			charCounter++
			if char == '\n' {
				lineCounter++
			}
			if unicode.IsSpace(rune(char)) {
				inWord = false
			} else if !inWord {
				inWord = true
				wordCounter++
			}
		}
	}

	return lineCounter, wordCounter, charCounter
}

func printFileStats(flags []string, file string, lines, words, chars int) {
	if ContainsFlag(flags, "-l") {
		fmt.Printf("%d ", lines)
	} else if ContainsFlag(flags, "-w") {
		fmt.Printf("%d ", words)
	} else if ContainsFlag(flags, "-c") {
		fmt.Printf("%d ", chars)
	} else {
		fmt.Printf("%d %d %d ", lines, words, chars)
	}

	fmt.Println(file)
}

func printTotalStats(flags []string, numFiles, totalLines, totalWords, totalChars int) {
	if numFiles > 1 {
		if ContainsFlag(flags, "-l") {
			fmt.Printf("%d ", totalLines)
		} else if ContainsFlag(flags, "-w") {
			fmt.Printf("%d ", totalWords)
		} else if ContainsFlag(flags, "-c") {
			fmt.Printf("%d ", totalChars)
		} else {
			fmt.Printf("%d %d %d ", totalLines, totalWords, totalChars)
		}

		fmt.Println("total")
	}
}

package pkg

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

func Wc() {

	var cFlag bool
	var lFlag bool
	var wFlag bool

	flag.BoolVar(&cFlag, "c", false, "print the byte counts")
	flag.BoolVar(&lFlag, "l", false, "print the newline counts")
	flag.BoolVar(&wFlag, "w", false, "print the word counts")

	flag.Parse()

	ok := len(flag.Args()) > 0

	if !ok {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	files := flag.Args()

	var totalLines, totalWords, totalChars int

	for _, file := range files {
		lineCount, wordCount, charCount := countFileStats(file)
		totalLines += lineCount
		totalWords += wordCount
		totalChars += charCount

		printFileStats(file, lineCount, wordCount, charCount, lFlag, wFlag, cFlag)
	}

	printTotalStats(len(files), totalLines, totalWords, totalChars, lFlag, wFlag, cFlag)
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

func printFileStats(file string, lines, words, chars int, lFlag, wFlag, cFlag bool) {
	if lFlag {
		fmt.Printf("%d ", lines)
	} else if wFlag {
		fmt.Printf("%d ", words)
	} else if cFlag {
		fmt.Printf("%d ", chars)
	} else {
		fmt.Printf("%d %d %d ", lines, words, chars)
	}

	fmt.Println(file)
}

func printTotalStats(numFiles, totalLines, totalWords, totalChars int, lFlag, wFlag, cFlag bool) {
	if numFiles > 1 {
		if lFlag {
			fmt.Printf("%d ", totalLines)
		} else if wFlag {
			fmt.Printf("%d ", totalWords)
		} else if cFlag {
			fmt.Printf("%d ", totalChars)
		} else {
			fmt.Printf("%d %d %d ", totalLines, totalWords, totalChars)
		}

		fmt.Println("total")
	}
}

package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	"fmt"
	"os"
	"unicode"
)

func Wc(flags []string) {
	
	command := models.CommandsMap["wc"]
	
	if !utils.ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}
	
	if len(os.Args) == 1 {
		fmt.Println("No file was entered!")
		os.Exit(1)
	}

	var files []string

	for _, arg := range os.Args[2:] {
		if arg != "-l" && arg != "-w" && arg != "-c" {
			files = append(files, arg)
		}
	}

	if len(files) == 0 {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}

	var totalLines, totalWords, totalChars int

	for _, file := range files {
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

		
		totalLines += lineCounter
		totalWords += wordCounter
		totalChars += charCounter

		if utils.ContainsFlag(flags, "-l") {
			fmt.Printf("%d ", lineCounter)
		} else if utils.ContainsFlag(flags, "-c") {
			fmt.Printf("%d ", charCounter)
		} else if utils.ContainsFlag(flags, "-w") {
			fmt.Printf("%d ", wordCounter)
		} else {
			fmt.Printf("%d %d %d ", lineCounter, wordCounter, charCounter)
		}

		fmt.Println(file)
	}

	if len(files) > 1 {
		if utils.ContainsFlag(flags, "-l") {
			fmt.Printf("%d ", totalLines)
		} else if utils.ContainsFlag(flags, "-c") {
			fmt.Printf("%d ", totalChars)
		} else if utils.ContainsFlag(flags, "-w") {
			fmt.Printf("%d ", totalWords)
		} else {
			fmt.Printf("%d %d %d ", totalLines, totalWords, totalChars)
		}

		fmt.Println("total")
	}
}

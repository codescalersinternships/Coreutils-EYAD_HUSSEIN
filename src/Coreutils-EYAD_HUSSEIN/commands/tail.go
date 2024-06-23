package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Tail(command *models.Command, flags []string) {
	if !utils.ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}

	numLines := 10
	fileIndex := 2

	_, ok := utils.ContainsFlag(flags, "-n")
	if ok {
		for i := 2; i < len(os.Args)-1; i++ {
			if os.Args[i] == "-n" {
				numLinesTemp, err := strconv.Atoi(os.Args[i+1])
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
				numLines = numLinesTemp
				fileIndex = i + 2
				break
			}
		}
	}

	fileName := os.Args[fileIndex]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()

	lines := make([]string, numLines)
	lineCount := 0
	scanner := bufio.NewScanner(file)

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

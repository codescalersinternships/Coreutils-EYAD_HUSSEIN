package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Tail(flags []string) {

	command := models.CommandsMap["tail"]

	if !utils.ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}
	
	ok := utils.ContainsFlag(flags, "-n")
	idx := utils.GetIndexOfArg("-n")

	var files []string

	numLines := 10

	if ok {
		for _, arg := range os.Args[2:] {
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
		files = os.Args[2:]
	}

	if len(files) == 0 {
		fmt.Println("No files were entered!")
		os.Exit(1)
	}


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
}
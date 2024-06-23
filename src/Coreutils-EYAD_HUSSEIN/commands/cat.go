package commands

import (
	"Coreutils-EYAD_HUSSEIN/models"
	"Coreutils-EYAD_HUSSEIN/utils"
	"fmt"
	"os"
)

func Cat(command *models.Command, flags []string) {
	if !utils.ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}
	
	ok := utils.ContainsFlag(flags, "-n")
	
	var files []string
	
	if ok {
		for _, arg := range os.Args[2:] {
			if arg == "-n" {
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

	if !ok {
		for _, file := range files {
			f, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println(string(f))
		}
	} else {
		lineNumber := 1
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			defer f.Close()
	
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
						if ok {
							fmt.Printf("%6d  %s\n", lineNumber, line)
							lineNumber++
						} else {
							fmt.Println(line)
						}
						line = ""
					} else {
						line += string(char)
					}
				}
			}
	
			if len(line) > 0 {
				if ok {
					fmt.Printf("%6d  %s\n", lineNumber, line)
				} else {
					fmt.Println(line)
				}
			}
		}
	}
	

}

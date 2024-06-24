package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Tree(flags []string) {
	command := CommandsMap["tree"]

	if !ValidateFlags(flags, command.Flags) {
		os.Exit(1)
	}

	maxDepth := -1

	if ok := ContainsFlag(flags, "-L"); ok {
		idx := GetIndexOfArg("-L")
		tempMaxDepth, err := strconv.Atoi(os.Args[idx+1])
		if err != nil {
			fmt.Println("Invalid depth value")
			os.Exit(1)
		}
		maxDepth = tempMaxDepth
	}

	for index, arg := range os.Args[1:] {
		if arg == "-L" || arg == strconv.Itoa(maxDepth) {
			continue
		}
		printDirectory(arg, 0, maxDepth)

		if index != len(os.Args[1:])-1 {
			fmt.Println()
		}
	}
}

func printListing(entry string, depth int, maxDepth int) {
	if maxDepth != -1 && depth > maxDepth {
		return
	}

	indent := strings.Repeat("|   ", depth)

	if depth == 0 {
		fmt.Println(entry)
	} else {
		entrySplit := strings.Split(entry, "/")
		var entryName string

		if entrySplit[len(entrySplit)-1] == "" {
			entryName = entrySplit[len(entrySplit)-2]
		} else {
			entryName = entrySplit[len(entrySplit)-1]
		}

		fmt.Printf("%s|-- %s\n", indent, entryName)
	}
}

func printDirectory(path string, depth, maxDepth int) {
	if maxDepth != -1 && depth > maxDepth {
		return
	}

	if ent := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]; depth != 0 && ent != "" && ent[0] == '.' {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	printListing(path, depth, maxDepth)

	for _, entry := range entries {
		entryPath := filepath.Join(path, entry.Name())
		fileInfo, err := os.Lstat(entryPath)
		if err != nil {
			fmt.Printf("error reading %s: %s\n", entryPath, err.Error())
			continue
		}

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			fullPath, err := os.Readlink(entryPath)
			if err != nil {
				fmt.Printf("error reading link: %s\n", err.Error())
			} else {
				printListing(entry.Name()+" -> "+fullPath, depth+1, maxDepth)
			}
		} else if fileInfo.IsDir() {
			printDirectory(entryPath, depth+1, maxDepth)
		} else {
			printListing(entry.Name(), depth+1, maxDepth)
		}
	}
}

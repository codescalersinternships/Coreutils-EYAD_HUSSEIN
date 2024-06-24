package pkg

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Tree() {

	var maxDepth int

	flag.IntVar(&maxDepth, "L", -1, "Descend only level directories deep")

	flag.Parse()

	ok := len(flag.Args()) > 0

	if !ok {
		printDirectory(".", 0, maxDepth)
	} else {
		dirs := flag.Args()

		for index, dir := range dirs {
			printDirectory(dir, 0, maxDepth)

			if index != len(dirs)-1 {
				fmt.Println()
			}
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

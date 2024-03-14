package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func readDirectory(path string, nestingLevel int) {
	dir, _ := os.Open(path)
	items, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("Error reading the directory : %q ", path)
	} else {
		var dirs []os.FileInfo
		var files []os.FileInfo

		for _, file := range items {
			if file.IsDir() {
				dirs = append(dirs, file)
			} else {
				files = append(files, file)
			}
		}

		sort.Slice(dirs, func(i, j int) bool {
			return dirs[i].Name() < dirs[j].Name()
		})

		sort.Slice(files, func(i, j int) bool {
			return files[i].Name() < files[j].Name()
		})

		for i, item := range dirs {
			var treeSymbol string
			if i == len(dirs)-1 && len(files) == 0 {
				treeSymbol = "└──"
			} else {
				treeSymbol = "├──"
			}

			// Check for possible errors when reading the directory content
			nextPath := filepath.Join(path, item.Name())
			nextDir, _ := os.Open(nextPath)
			_, err := nextDir.Readdir(-1)

			concatenated := strings.Join([]string{strings.Repeat("│   ", nestingLevel), treeSymbol, item.Name()}, "")

			if err != nil {
				concatenated = strings.Join([]string{concatenated, "  [ error opening dir ]"}, "")
				fmt.Println(concatenated)
			} else {
				fmt.Println(concatenated)
				readDirectory(filepath.Join(path, item.Name()), nestingLevel+1)
			}

		}

		for i, item := range files {
			var treeSymbol string
			if i == len(files)-1 {
				treeSymbol = "└──"
			} else {
				treeSymbol = "├──"
			}
			concatenated := strings.Join([]string{strings.Repeat("│   ", nestingLevel), treeSymbol, item.Name()},
				"")
			fmt.Println(concatenated)
		}
	}

}

func main() {
	if len(os.Args) < 2 {
		readDirectory(".", 0)
	} else {
		for i, arg := range os.Args {
			if i > 0 {
				readDirectory(arg, 0)
			}
		}
	}

}

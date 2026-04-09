package main

import (
	"fmt"
	"os"
)

func scoop(dir string, flags Flags, depth int) {
	if depth >= flags.Depth {
		return
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ Error ] %v\n", err)
		os.Exit(1)
	}
	for _, file := range files {
		fileName := file.Name()
		if isHidden(fileName) {
			continue
		}
		fmt.Printf("%s%s\n", indent(depth), fileName)
		if file.Type().IsDir() {
			scoop(dir+"/"+fileName, flags, depth+1)
		}
	}
}

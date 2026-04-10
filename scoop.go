package main

import (
	"log"
	"os"
)

func scoop(dir string, flags Flags, depth int) {
	if depth >= flags.Depth {
		return
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("[ Error ] %v\n", err)
	}
	for _, file := range files {
		fileName := file.Name()
		if !flags.Hidden && isHidden(fileName) {
			continue
		}
		printFile(fileName, flags, depth)
		if file.Type().IsDir() {
			scoop(dir+"/"+fileName, flags, depth+1)
		}
	}
}

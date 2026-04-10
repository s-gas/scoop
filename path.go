package main

import (
	"flag"
	"log"
	"os"
)

func parsePath() (path string) {
	args := flag.Args()
	numArgs := len(args)
	if numArgs > 1 {
		log.Fatalln("[ Error ] usage: scoop [flags] [path]")
	} else if numArgs == 1 {
		return args[0]
	} else {
		path, err := os.Getwd()
		if err != nil {
			log.Fatalf("[ Error ] %v\n", err)
		}
		return path
	}
	return ""
}

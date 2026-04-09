package main

import (
	"flag"
	"fmt"
	"os"
)

func parsePath() (path string) {
	args := flag.Args()
	numArgs := len(args)
	if numArgs > 1 {
		fmt.Fprintf(os.Stderr, "[ Error ] invalid argument: %s\n", args[0])
		os.Exit(1)
		return ""
	} else if numArgs == 1 {
		return args[0]
	} else {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[ Error ] %v\n", err)
			os.Exit(1)
		}
		return path
	}
}

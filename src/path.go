package main

import (
	"fmt"
	"os"
)

func parsePath() (path string) {
	numArgs := len(os.Args)
	if numArgs == 2 {
		return os.Args[1]
	} else {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		return path
	}
}

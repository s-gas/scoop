package main

import "fmt"

func printFile(fileName string, flags Flags, depth int) {
	if flags.Find != "" {
	}
	fmt.Printf("%s%s\n", indent(depth), fileName)
}

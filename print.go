package main

import (
	"fmt"
	"strings"
)

const highlight = "\033[33m"
const reset = "\033[0m"

func printFile(fileName string, flags Flags, depth int) {
	if flags.Find != "" && strings.Contains(fileName, flags.Find) {
		fmt.Printf("%s%s%s%s\n", indent(depth), highlight, fileName, reset)
	} else {
		fmt.Printf("%s%s\n", indent(depth), fileName)
	}
}

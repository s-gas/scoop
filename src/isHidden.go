package main

import "strings"

func isHidden(file string) bool {
	return strings.HasPrefix(file, ".")
}

package main

import "fmt"

func main() {
	flags := parseFlags()
	fmt.Println(flags.Depth)
	fmt.Println(flags.Hidden)
	dir := parsePath()
	scoop(dir, 0)
}

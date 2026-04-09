package main

func main() {
	flags := parseFlags()
	dir := parsePath()
	scoop(dir, 0)
}

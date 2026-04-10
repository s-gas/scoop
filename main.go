package main

func main() {
	flags := parseFlags()
	dir := parsePath()
	scoop(dir, flags, 0)
}

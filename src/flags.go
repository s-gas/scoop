package main

import "flag"

type Flags struct {
	Find   string
	Depth  int
	Hidden bool
}

func parseFlags() Flags {
	flags := Flags{}
	flag.StringVar(&flags.Find, "find", "", "highlight files containing the given string")
	flag.IntVar(&flags.Depth, "depth", 4, "maximum depth")
	flag.BoolVar(&flags.Hidden, "hidden", false, "show hidden files")
	flag.Parse()
	return flags
}

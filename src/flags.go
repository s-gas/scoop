package main

import "flag"

type Flags struct {
	Hidden bool
	Depth  int
}

func parseFlags() Flags {
	flags := Flags{}
	flag.BoolVar(&flags.Hidden, "hidden", false, "show hidden files")
	flag.IntVar(&flags.Depth, "depth", 4, "maximum depth")
	flag.Parse()
	return flags
}

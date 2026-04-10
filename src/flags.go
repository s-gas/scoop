package main

import "flag"

type Flags struct {
	Find   string
	Depth  int
	Hidden bool
}

func parseFlags() Flags {
	flags := Flags{}
	flag.BoolVar(&flags.Hidden, "hidden", false, "show hidden files")
	flag.IntVar(&flags.Depth, "depth", 4, "maximum depth")
	flag.Parse()
	return flags
}

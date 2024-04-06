package cli

import "flag"

// A Flags represents data parsed from command line
type Flags struct {
	Output  bool
	Limit   int
	Threads int
}

// ParseFlags parses flags from command line
func ParseFlags() Flags {
	var flags Flags

	flag.BoolVar(&flags.Output, "o", false, "Output")
	flag.IntVar(&flags.Limit, "n", -1, "Limit")
	flag.IntVar(&flags.Threads, "t", 1, "Threads")
	flag.Parse()

	return flags
}

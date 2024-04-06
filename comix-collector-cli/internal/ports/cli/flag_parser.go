package cli

import (
	"flag"

	"github.com/go-playground/validator"
)

// A Flags represents data parsed from command line
type Flags struct {
	Output  bool
	Limit   int `validate:"gte=-1"`
	Threads int `validate:"gt=0"`
}

// ParseFlags parses flags from command line
func ParseFlags() (Flags, error) {
	var flags Flags

	flag.BoolVar(&flags.Output, "o", false, "Output")
	flag.IntVar(&flags.Limit, "n", -1, "Limit")
	flag.IntVar(&flags.Threads, "t", 1, "Threads")
	flag.Parse()

	err := flags.Validate()
	if err != nil {
		return Flags{}, err
	}
	return flags, nil
}

// Validate validates flags
func (f Flags) Validate() error {
	validate := validator.New()
	return validate.Struct(f)
}

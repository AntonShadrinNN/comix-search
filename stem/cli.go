/*
Stem reduces inflected words to their word stem.
It uses external stemming module.

Usage:

	stem [flags]

The flags are:

	-s
	    Sentence to stem.
*/
package main

import (
	"flag"
)

const (
	initialStringFlagName = "s"
)

// A Flags represents data parsed from command line
type Flags struct {
	InitialString string
}

// ValidateFlags validates flags provided through command line
func (f Flags) ValidateFlags() error {
	if f.InitialString == "" {
		return errFlagIsMandatory(initialStringFlagName)
	}

	return nil
}

// ParseFlags parses flags from command line
func ParseFlags() Flags {
	var flags Flags

	flag.StringVar(&flags.InitialString, initialStringFlagName, "", "Mandatory. String to stem")
	flag.Parse()

	return flags
}

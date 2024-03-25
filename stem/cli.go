package main

import (
	"flag"
)

const (
	initialStringFlagName = "s"
)

type Flags struct {
	InitialString string
}

func (f Flags) validateFlags() error {
	if f.InitialString == "" {
		return errFlagIsMandatory(initialStringFlagName)
	}

	return nil
}

func ParseFlags() Flags {
	var flags Flags

	flag.StringVar(&flags.InitialString, initialStringFlagName, "", "Mandatory. String to stem")
	flag.Parse()

	return flags
}

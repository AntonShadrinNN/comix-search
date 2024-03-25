package main

import (
	"log"
	"strings"
)

func main() {
	flags, err := ParseFlags()
	if err != nil {
		panic(err)
	}

	stemmed, err := Stem(strings.Split(flags.InitialString, " "), English)
	if err != nil {
		panic(err)
	}
	log.Printf(stemmed)
}

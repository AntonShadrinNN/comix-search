package main

import "fmt"

var errFlagIsMandatory = func(flag string) error { return fmt.Errorf("flag %s is mandatory", flag) }

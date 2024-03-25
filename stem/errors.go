package main

import "fmt"

// errFlagIsMandatory is used to notify user, that flag must be provided
var errFlagIsMandatory = func(flag string) error { return fmt.Errorf("flag %s is mandatory", flag) }

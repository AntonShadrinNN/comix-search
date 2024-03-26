package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrFlagIsMandatory(t *testing.T) {
	flag := "some_flag"
	expected := fmt.Sprintf("flag %s is mandatory", flag)
	got := errFlagIsMandatory(flag)

	assert.Equal(t, expected, got.Error(), "Expected another error message")
}

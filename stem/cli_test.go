package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	someStr = "some string"
)

func TestFlagsValidation(t *testing.T) {
	testCases := []struct {
		FlagsData Flags
		Err       error
	}{
		{
			FlagsData: Flags{
				InitialString: "",
			},
			Err: errFlagIsMandatory(initialStringFlagName),
		},
		{
			FlagsData: Flags{
				InitialString: "some value",
			},
			Err: nil,
		},
	}

	for _, testCase := range testCases {
		err := testCase.FlagsData.ValidateFlags()
		assert.Equal(t, testCase.Err, err)
	}
}

func TestGetCommandLineFlags(t *testing.T) {
	testCases := []struct {
		Err           error
		ExpectedFlags Flags
	}{
		{
			Err:           errFlagIsMandatory(initialStringFlagName),
			ExpectedFlags: Flags{},
		},
		{
			Err:           nil,
			ExpectedFlags: Flags{InitialString: someStr},
		},
	}

	for _, testCase := range testCases {
		err := testCase.ExpectedFlags.ValidateFlags()
		assert.Equal(t, testCase.Err, err)
	}
}

func TestParseFlags(t *testing.T) {
	args := []string{fmt.Sprintf("-s=%s", someStr)}

	// Сохраняем исходные аргументы командной строки
	originalArgs := os.Args

	// Подменяем аргументы командной строки для теста
	os.Args = append([]string{originalArgs[0]}, args...)

	flags := ParseFlags()

	assert.Equal(t, someStr, flags.InitialString)
	os.Args = originalArgs
}

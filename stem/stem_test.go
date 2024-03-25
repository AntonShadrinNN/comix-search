package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	someErr         = fmt.Errorf("some error")
	stopWordExample = "you"
	testSentence    = "i'll follow you as long as you are following me"
)

func TestStem(t *testing.T) {

	testTable := []struct {
		StopWordsFunc  func() ([]string, error)
		ExpectedString string
		ExpectedError  error
	}{
		{
			StopWordsFunc:  func() ([]string, error) { return []string{}, someErr },
			ExpectedString: "",
			ExpectedError:  someErr,
		},
		{
			StopWordsFunc:  func() ([]string, error) { return []string{stopWordExample}, nil },
			ExpectedString: "are as follow i'll long me",
			ExpectedError:  nil,
		},
	}

	stemmer := SnowballStemmer{
		language: English,
	}
	for _, testCase := range testTable {
		stemmer.stopWordsFunc = testCase.StopWordsFunc
		gotStr, gotErr := stemmer.Stem(strings.Split(testSentence, " "))

		assert.Equal(t, testCase.ExpectedString, gotStr)
		assert.Equal(t, testCase.ExpectedError, gotErr)
	}

}

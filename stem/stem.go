package main

import (
	"slices"
	"strings"

	"github.com/kljensen/snowball"
)

type Stemmer interface {
	Stem(words []string)
}

type SnowballStemmer struct {
	language      Language
	stopWordsFunc func() ([]string, error)
}

func (s SnowballStemmer) Stem(words []string) (string, error) {
	stopWords, err := s.stopWordsFunc()
	if err != nil {
		return "", err
	}

	stemmedSentence := make([]string, 0, len(words))
	for _, word := range words {
		stemmed, err := snowball.Stem(word, s.language, false)
		if err != nil {
			return "", err
		}

		if slices.Contains[[]string, string](stopWords, stemmed) {
			continue
		}

		stemmedSentence = append(stemmedSentence, stemmed)
	}

	slices.Sort(stemmedSentence)
	result := slices.Compact(stemmedSentence)
	return strings.Join(result, " "), nil
}

package main

import (
	"slices"
	"strings"

	"github.com/kljensen/snowball"
)

// A Stemmer reduces inflected words to their word stem
type Stemmer interface {
	Stem(words []string)
}

// A SnowballStemmer is a concrete implementation of Stemmer interface
type SnowballStemmer struct {
	language      Language                 // Word's language
	stopWordsFunc func() ([]string, error) // Function that returns stop words
}

// Stem returns whitspace-separated string with stemmed words and error if occured
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

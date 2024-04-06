package words

import (
	"regexp"
	"slices"

	"github.com/kljensen/snowball"
)

// A SnowballStemmer is a concrete implementation of Stemmer interface
type SnowballStemmer struct {
	language      Language                 // Word's language
	stopWordsFunc func() ([]string, error) // Function that returns stop words
}

func NewSnowballStemmer(lang Language, stopWordsFunc func() ([]string, error)) SnowballStemmer {
	return SnowballStemmer{
		language:      lang,
		stopWordsFunc: stopWordsFunc,
	}
}

// Stem returns whitspace-separated string with stemmed words and error if occured
func (s SnowballStemmer) Stem(words []string) ([]string, error) {
	stopWords, err := s.stopWordsFunc()
	if err != nil {
		return nil, err
	}

	stemmedSentence := make([]string, 0, len(words))
	for _, word := range words {
		cleanWord := regexp.MustCompile(`[^a-zA-Z0-9' ]+`).ReplaceAllString(word, "")
		stemmed, err := snowball.Stem(cleanWord, s.language, false)
		if err != nil {
			return nil, err
		}

		if slices.Contains[[]string, string](stopWords, stemmed) {
			continue
		}

		stemmedSentence = append(stemmedSentence, stemmed)
	}

	slices.Sort(stemmedSentence)
	result := slices.Compact(stemmedSentence)
	return result, nil
}

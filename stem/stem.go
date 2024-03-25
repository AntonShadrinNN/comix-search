package main

import (
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/kljensen/snowball"
)

func Stem(words []string, language Language) (string, error) {
	stopWords, err := getStopWords()
	if err != nil {
		return "", err
	}

	stemmedSentence := make([]string, 0, len(words))
	for _, word := range words {
		stemmed, err := snowball.Stem(word, language, false)
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

func getStopWords() ([]string, error) {
	resp, err := http.Get("https://gist.githubusercontent.com/rg089/35e00abf8941d72d419224cfd5b5925d/raw/12d899b70156fd0041fa9778d657330b024b959c/stopwords.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
